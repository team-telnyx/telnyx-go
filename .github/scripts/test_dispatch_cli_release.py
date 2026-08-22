#!/usr/bin/env python3
"""Tests for durable and observable Go-to-CLI release handoff."""

import importlib.util
from pathlib import Path
import unittest

SCRIPT = Path(__file__).with_name("dispatch_cli_release.py")
SPEC = importlib.util.spec_from_file_location("dispatch_cli_release", SCRIPT)
assert SPEC and SPEC.loader
module = importlib.util.module_from_spec(SPEC)
try:
    SPEC.loader.exec_module(module)
except FileNotFoundError:
    module = None


class FakeClient:
    def __init__(self, repository_failures=0, workflow_failures=0, runs=None):
        self.repository_failures = repository_failures
        self.workflow_failures = workflow_failures
        self.runs = list(runs or [])
        self.repository_payloads = []
        self.workflow_inputs = []
        self.list_calls = 0

    def dispatch_repository(self, payload):
        self.repository_payloads.append(payload)
        if self.repository_failures:
            self.repository_failures -= 1
            raise module.HandoffError("transient repository dispatch failure")

    def dispatch_workflow(self, inputs):
        self.workflow_inputs.append(inputs)
        if self.workflow_failures:
            self.workflow_failures -= 1
            raise module.HandoffError("transient workflow dispatch failure")

    def list_workflow_runs(self):
        self.list_calls += 1
        return self.runs


def run(handoff_id="handoff-123", **client_kwargs):
    client = FakeClient(**client_kwargs)
    payload = {
        "handoff_id": handoff_id,
        "go_version": "v4.97.0",
        "go_release_sha": "1" * 40,
        "sdk_config_sha": "2" * 40,
    }
    result = module.deliver_and_observe(
        client,
        payload,
        dispatch_attempts=3,
        observation_attempts=2,
        sleep=lambda _: None,
    )
    return client, result


@unittest.skipIf(module is None, "implementation does not exist yet")
class DispatchCliReleaseTests(unittest.TestCase):
    def test_repository_dispatch_is_immediate_and_exact(self):
        run_record = {
            "id": 77,
            "event": "repository_dispatch",
            "display_title": "Promote CLI [handoff:handoff-123]",
            "html_url": "https://github.example/runs/77",
        }
        client, result = run(runs=[run_record])
        self.assertEqual(result.mode, "repository_dispatch")
        self.assertEqual(result.run_id, 77)
        self.assertEqual(client.repository_payloads[0]["client_payload"]["go_version"], "v4.97.0")
        self.assertEqual(client.repository_payloads[0]["client_payload"]["sdk_config_sha"], "2" * 40)
        self.assertEqual(client.workflow_inputs, [])

    def test_retries_recoverable_repository_delivery(self):
        record = {
            "id": 78,
            "event": "repository_dispatch",
            "display_title": "Promote CLI [handoff:handoff-123]",
            "html_url": "https://github.example/runs/78",
        }
        client, result = run(repository_failures=2, runs=[record])
        self.assertEqual(len(client.repository_payloads), 3)
        self.assertEqual(result.run_id, 78)

    def test_immediate_workflow_reconciliation_after_delivery_failure(self):
        record = {
            "id": 79,
            "event": "workflow_dispatch",
            "display_title": "Promote CLI [handoff:handoff-123]",
            "html_url": "https://github.example/runs/79",
        }
        client, result = run(repository_failures=3, runs=[record])
        self.assertEqual(result.mode, "workflow_dispatch")
        self.assertEqual(client.workflow_inputs[0]["handoff_id"], "handoff-123")
        self.assertEqual(client.workflow_inputs[0]["go_release_sha"], "1" * 40)

    def test_unobserved_primary_delivery_triggers_immediate_reconciliation(self):
        client = FakeClient(runs=[])
        payload = {
            "handoff_id": "handoff-456",
            "go_version": "v4.97.0",
            "go_release_sha": "1" * 40,
            "sdk_config_sha": "2" * 40,
        }

        original_dispatch_workflow = client.dispatch_workflow

        def dispatch_workflow(inputs):
            original_dispatch_workflow(inputs)
            client.runs = [{
                "id": 80,
                "event": "workflow_dispatch",
                "display_title": "Promote CLI [handoff:handoff-456]",
                "html_url": "https://github.example/runs/80",
            }]

        client.dispatch_workflow = dispatch_workflow
        result = module.deliver_and_observe(
            client, payload, dispatch_attempts=1, observation_attempts=1, sleep=lambda _: None
        )
        self.assertEqual(result.mode, "workflow_dispatch")
        self.assertEqual(result.run_id, 80)

    def test_fails_loud_when_neither_delivery_path_is_observable(self):
        client = FakeClient(repository_failures=2, workflow_failures=2)
        payload = {
            "handoff_id": "handoff-789",
            "go_version": "v4.97.0",
            "go_release_sha": "1" * 40,
            "sdk_config_sha": "2" * 40,
        }
        with self.assertRaisesRegex(module.HandoffError, "could not be delivered"):
            module.deliver_and_observe(
                client, payload, dispatch_attempts=2, observation_attempts=1, sleep=lambda _: None
            )

    def test_observation_requires_exact_handoff_identity(self):
        wrong = {
            "id": 81,
            "event": "repository_dispatch",
            "display_title": "Promote CLI [handoff:other]",
            "html_url": "https://github.example/runs/81",
        }
        client = FakeClient(workflow_failures=1, runs=[wrong])
        payload = {
            "handoff_id": "wanted",
            "go_version": "v4.97.0",
            "go_release_sha": "1" * 40,
            "sdk_config_sha": "2" * 40,
        }
        with self.assertRaises(module.HandoffError):
            module.deliver_and_observe(
                client, payload, dispatch_attempts=1, observation_attempts=1, sleep=lambda _: None
            )


if __name__ == "__main__":
    if module is None:
        raise SystemExit("dispatch_cli_release.py is missing (expected red test state)")
    unittest.main()
