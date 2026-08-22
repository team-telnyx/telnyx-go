#!/usr/bin/env python3
"""Deliver an immutable Go release handoff and prove the CLI run was created."""

from __future__ import annotations

import argparse
import json
import os
import re
import subprocess
import sys
import time
import uuid
from typing import Callable, NamedTuple

CLI_REPOSITORY = "team-telnyx/telnyx-cli-staging"
CLI_WORKFLOW = "promote-to-prod.yml"
SHA_RE = re.compile(r"^[0-9a-f]{40}$")
VERSION_RE = re.compile(r"^v4\.[0-9]+\.[0-9]+$")
HANDOFF_RE = re.compile(r"^[0-9a-z-]{8,80}$")


class HandoffError(RuntimeError):
    """The handoff could not be delivered or independently observed."""


class HandoffResult(NamedTuple):
    mode: str
    run_id: int
    run_url: str
    handoff_id: str


def validate_payload(payload):
    required = {"handoff_id", "go_version", "go_release_sha", "sdk_config_sha"}
    if set(payload) != required:
        raise HandoffError("handoff payload has missing or unknown fields")
    if not HANDOFF_RE.fullmatch(payload["handoff_id"]):
        raise HandoffError("handoff_id is malformed")
    if not VERSION_RE.fullmatch(payload["go_version"]):
        raise HandoffError("go_version must be an exact v4.X.Y tag")
    for key in ("go_release_sha", "sdk_config_sha"):
        if not SHA_RE.fullmatch(payload[key]):
            raise HandoffError(f"{key} must be a full lowercase commit SHA")


def _attempt(operation: Callable[[], None], attempts: int, sleep: Callable[[float], None]):
    errors = []
    for attempt in range(1, attempts + 1):
        try:
            operation()
            return errors
        except HandoffError as error:
            errors.append(str(error))
            if attempt < attempts:
                sleep(min(2 ** (attempt - 1), 8))
    raise HandoffError("; ".join(errors))


def _observe(client, handoff_id, attempts, sleep):
    marker = f"[handoff:{handoff_id}]"
    for attempt in range(attempts):
        candidates = [
            run
            for run in client.list_workflow_runs()
            if run.get("event") in {"repository_dispatch", "workflow_dispatch"}
            and marker in (run.get("display_title") or "")
        ]
        if len(candidates) > 1:
            raise HandoffError(f"multiple CLI workflow runs claim handoff {handoff_id}")
        if candidates:
            run = candidates[0]
            run_id = run.get("id")
            run_url = run.get("html_url")
            if not isinstance(run_id, int) or not isinstance(run_url, str):
                raise HandoffError("observed CLI workflow run has malformed identity")
            return run_id, run_url
        if attempt + 1 < attempts:
            sleep(10)
    return None


def deliver_and_observe(
    client,
    payload,
    *,
    dispatch_attempts=3,
    observation_attempts=30,
    sleep=time.sleep,
):
    validate_payload(payload)
    if dispatch_attempts < 1 or observation_attempts < 1:
        raise HandoffError("attempt counts must be positive")

    envelope = {"event_type": "telnyx-go-released", "client_payload": dict(payload)}
    mode = "repository_dispatch"
    try:
        _attempt(lambda: client.dispatch_repository(envelope), dispatch_attempts, sleep)
    except HandoffError as primary_error:
        mode = "workflow_dispatch"
        try:
            _attempt(lambda: client.dispatch_workflow(dict(payload)), dispatch_attempts, sleep)
        except HandoffError as fallback_error:
            raise HandoffError(
                "CLI handoff could not be delivered by repository or workflow dispatch: "
                f"repository={primary_error}; workflow={fallback_error}"
            ) from fallback_error

    primary_observation_attempts = min(observation_attempts, 6)
    observed = _observe(client, payload["handoff_id"], primary_observation_attempts, sleep)
    if observed is None and mode == "repository_dispatch":
        # A 204 without an observable run is not proof of delivery. Immediately
        # reconcile through the independently addressable workflow endpoint.
        mode = "workflow_dispatch"
        _attempt(lambda: client.dispatch_workflow(dict(payload)), dispatch_attempts, sleep)
        observed = _observe(client, payload["handoff_id"], observation_attempts, sleep)
    if observed is None:
        raise HandoffError(
            f"CLI handoff {payload['handoff_id']} was dispatched but no exact workflow run was observable"
        )
    return HandoffResult(mode, observed[0], observed[1], payload["handoff_id"])


class GhClient:
    def __init__(self, repository=CLI_REPOSITORY, workflow=CLI_WORKFLOW):
        self.repository = repository
        self.workflow = workflow

    def _api(self, method, path, payload=None):
        command = ["gh", "api", "--method", method, path]
        if payload is not None:
            command.extend(("--input", "-"))
        process = subprocess.run(
            command,
            input=(json.dumps(payload) if payload is not None else None),
            text=True,
            stdout=subprocess.PIPE,
            stderr=subprocess.PIPE,
            check=False,
        )
        if process.returncode != 0:
            detail = process.stderr.strip() or process.stdout.strip()
            raise HandoffError(f"GitHub API {method} {path} failed: {detail}")
        if not process.stdout.strip():
            return None
        try:
            return json.loads(process.stdout)
        except json.JSONDecodeError as error:
            raise HandoffError(f"GitHub API returned malformed JSON for {path}") from error

    def dispatch_repository(self, payload):
        self._api("POST", f"repos/{self.repository}/dispatches", payload)

    def dispatch_workflow(self, inputs):
        self._api(
            "POST",
            f"repos/{self.repository}/actions/workflows/{self.workflow}/dispatches",
            {"ref": "main", "inputs": inputs},
        )

    def list_workflow_runs(self):
        response = self._api(
            "GET",
            f"repos/{self.repository}/actions/workflows/{self.workflow}/runs?per_page=30",
        )
        runs = response.get("workflow_runs") if isinstance(response, dict) else None
        if not isinstance(runs, list):
            raise HandoffError("GitHub API did not return a workflow run list")
        return runs


def _write_output(path, result):
    if not path:
        return
    with open(path, "a", encoding="utf-8") as handle:
        handle.write(f"handoff_id={result.handoff_id}\n")
        handle.write(f"cli_run_id={result.run_id}\n")
        handle.write(f"cli_run_url={result.run_url}\n")
        handle.write(f"delivery_mode={result.mode}\n")


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--go-version", required=True)
    parser.add_argument("--go-release-sha", required=True)
    parser.add_argument("--sdk-config-sha", required=True)
    parser.add_argument("--handoff-id", default=f"go-{uuid.uuid4().hex}")
    parser.add_argument("--github-output", default=os.environ.get("GITHUB_OUTPUT"))
    args = parser.parse_args()
    payload = {
        "handoff_id": args.handoff_id,
        "go_version": args.go_version,
        "go_release_sha": args.go_release_sha,
        "sdk_config_sha": args.sdk_config_sha,
    }
    try:
        result = deliver_and_observe(GhClient(), payload)
    except HandoffError as error:
        parser.error(str(error))
    _write_output(args.github_output, result)
    print(json.dumps(result._asdict(), sort_keys=True))
    return 0


if __name__ == "__main__":
    sys.exit(main())
