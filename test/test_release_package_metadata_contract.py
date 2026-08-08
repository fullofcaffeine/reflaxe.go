#!/usr/bin/env python3

from __future__ import annotations

import json
import unittest
from pathlib import Path


REPO_ROOT = Path(__file__).resolve().parent.parent
PACKAGE_JSON = REPO_ROOT / "package.json"
HAXELIB_JSON = REPO_ROOT / "haxelib.json"
README = REPO_ROOT / "README.md"
RELEASE_STATE_SCRIPT = REPO_ROOT / "scripts" / "release" / "check-release-state.sh"


class ReleasePackageMetadataContractTest(unittest.TestCase):
    def test_package_and_haxelib_versions_and_license_match(self) -> None:
        package = json.loads(PACKAGE_JSON.read_text(encoding="utf-8"))
        haxelib = json.loads(HAXELIB_JSON.read_text(encoding="utf-8"))

        self.assertEqual(package["version"], haxelib["version"])
        self.assertEqual(package["license"], haxelib["license"])

    def test_haxelib_url_points_at_release_repository(self) -> None:
        haxelib = json.loads(HAXELIB_JSON.read_text(encoding="utf-8"))
        self.assertEqual(haxelib["url"], "https://github.com/fullofcaffeine/reflaxe.go")

    def test_release_status_checks_haxelib_url(self) -> None:
        script = RELEASE_STATE_SCRIPT.read_text(encoding="utf-8")
        self.assertIn("HAXELIB_URL", script)
        self.assertIn("https://github.com/${REPO_SLUG}", script)

    def test_readme_documents_immutable_git_installation(self) -> None:
        readme = README.read_text(encoding="utf-8")
        for phrase in (
            "Install from the GitHub release tag",
            "haxelib git reflaxe.go "
            "https://github.com/fullofcaffeine/reflaxe.go.git v0.54.0",
            "lix install gh:fullofcaffeine/reflaxe.go#v0.54.0",
            "92d458e760a30bcb57f2cefb6202f0996fe1ac71",
            "development sentinel `0.0.0`",
            "GitHub release ZIP",
        ):
            self.assertIn(phrase, readme)


if __name__ == "__main__":
    unittest.main()
