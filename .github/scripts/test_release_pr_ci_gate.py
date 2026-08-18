#!/usr/bin/env python3
"""Static safety contracts for the DOT-2061 Go phase-2 rollout."""
import unittest
from pathlib import Path

ROOT = Path(__file__).resolve().parents[2]


class ReleasePRGateWorkflowTests(unittest.TestCase):
    def test_classifier_owns_every_full_ci_job(self):
        ci = (ROOT / ".github/workflows/ci.yml").read_text()
        self.assertIn("name: classify production CI", ci)
        self.assertIn("classify_production_ci.py --event-path", ci)
        self.assertEqual(ci.count("needs: classify-production-ci"), 3)
        self.assertEqual(ci.count("needs.classify-production-ci.outputs.run_full == 'true'"), 3)
        for name in ("name: lint", "name: build", "name: test"):
            self.assertIn(name, ci)
        for test in (
            "test_release_pr_auto_merge.py", "test_release_pr_ci_gate.py",
            "test_classify_production_ci.py", "test_validate_next_provenance.py",
            "test_verify_go_release.py",
        ):
            self.assertIn(test, ci)

    def test_release_workflow_verifies_tag_and_module_after_release(self):
        workflow = (ROOT / ".github/workflows/release-please.yml").read_text()
        self.assertIn("name: Verify Go tag and module availability", workflow)
        self.assertIn("verify_go_release.py", workflow)
        self.assertIn('--version "$VERSION"', workflow)
        self.assertIn('--release-sha "$RELEASE_SHA"', workflow)
        self.assertNotIn("release-pr-auto-merge.yml", workflow)
        self.assertNotIn("Dispatch exact-head release PR gate", workflow)

    def test_next_readiness_is_lightweight_and_fail_closed(self):
        workflow = (ROOT / ".github/workflows/next-readiness.yml").read_text()
        self.assertIn("branches: [next]", workflow)
        self.assertIn("name: next-readiness", workflow)
        self.assertIn("validate_next_provenance.py", workflow)
        self.assertIn("--expected-next", workflow)
        self.assertIn("MERGE_TOKEN: ${{ secrets.SDK_WRITE_TOKEN }}", workflow)
        self.assertIn("persist-credentials: false", workflow)
        self.assertNotIn("go build", workflow)
        self.assertNotIn("scripts/test", workflow)
        self.assertNotIn("scripts/lint", workflow)

    def test_readiness_uses_trusted_policy_and_never_merges(self):
        workflow = (ROOT / ".github/workflows/release-pr-readiness.yml").read_text()
        self.assertIn("pull_request_target:", workflow)
        self.assertIn("ref: ${{ github.event.repository.default_branch || 'main' }}", workflow)
        self.assertIn("persist-credentials: false", workflow)
        self.assertIn("--expected-head", workflow)
        self.assertIn("--dry-run", workflow)
        self.assertNotIn("--merge", workflow)

    def test_readiness_publishes_exact_head_status_fail_closed(self):
        workflow = (ROOT / ".github/workflows/release-pr-readiness.yml").read_text()
        self.assertIn("context=release-provenance", workflow)
        self.assertIn("state=pending", workflow)
        self.assertIn("STATE=failure", workflow)
        self.assertIn('[ "$STATE" = success ]', workflow)
        self.assertIn("statuses: write", workflow)


if __name__ == "__main__":
    unittest.main()
