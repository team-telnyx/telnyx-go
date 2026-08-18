#!/usr/bin/env python3
import unittest

from verify_go_release import AvailabilityError, validate_go_list, validate_module_file, validate_proxy_info, validate_tag

VERSION = "4.95.0"
TAG = "v" + VERSION
SHA = "a" * 40
MODULE = "github.com/team-telnyx/telnyx-go/v4"


class GoReleaseAvailabilityTests(unittest.TestCase):
    def test_accepts_exact_tag_proxy_and_module_resolution(self):
        validate_tag({"ref": "refs/tags/" + TAG, "object": {"type": "commit", "sha": SHA}}, TAG, SHA)
        validate_proxy_info({"Version": TAG, "Time": "2026-08-18T00:00:00Z"}, TAG)
        validate_module_file("module " + MODULE + "\n\ngo 1.22\n", MODULE)
        validate_go_list({"Path": MODULE, "Version": TAG}, MODULE, TAG)

    def test_rejects_tag_pointing_to_other_commit(self):
        with self.assertRaisesRegex(AvailabilityError, "commit"):
            validate_tag({"ref": "refs/tags/" + TAG, "object": {"type": "commit", "sha": "b" * 40}}, TAG, SHA)

    def test_rejects_wrong_proxy_version(self):
        with self.assertRaisesRegex(AvailabilityError, "version"):
            validate_proxy_info({"Version": "v4.94.0", "Time": "2026-08-18T00:00:00Z"}, TAG)

    def test_rejects_wrong_module_path(self):
        with self.assertRaisesRegex(AvailabilityError, "module path"):
            validate_module_file("module github.com/example/wrong/v4\n", MODULE)
        with self.assertRaisesRegex(AvailabilityError, "module path"):
            validate_go_list({"Path": "github.com/example/wrong/v4", "Version": TAG}, MODULE, TAG)

    def test_rejects_malformed_or_annotated_unresolved_tag(self):
        with self.assertRaisesRegex(AvailabilityError, "tag"):
            validate_tag({}, TAG, SHA)
        with self.assertRaisesRegex(AvailabilityError, "dereference"):
            validate_tag({"ref": "refs/tags/" + TAG, "object": {"type": "tag", "sha": SHA}}, TAG, SHA)


if __name__ == "__main__":
    unittest.main()
