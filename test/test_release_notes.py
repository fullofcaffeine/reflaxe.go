#!/usr/bin/env python3

from __future__ import annotations

import subprocess
import tempfile
import unittest
from pathlib import Path


ROOT = Path(__file__).resolve().parent.parent
BUILDER = ROOT / "scripts" / "release" / "build-release-notes.py"
SOURCE_SHA = "db5175ec27d1f41efa0ae15c9436b8c2e2893392"


class ReleaseNotesTest(unittest.TestCase):
    def test_notes_are_bounded_manifest_derived_and_release_aware(self) -> None:
        with tempfile.TemporaryDirectory() as temp:
            output = Path(temp) / "release-notes.md"
            proc = subprocess.run(
                [
                    "python3",
                    str(BUILDER),
                    "--version",
                    "0.54.0",
                    "--tag",
                    "v0.54.0",
                    "--source-sha",
                    SOURCE_SHA,
                    "--repository",
                    "fullofcaffeine/reflaxe.go",
                    "--manifest",
                    str(ROOT / "docs" / "compatibility-support-manifest.json"),
                    "--output",
                    str(output),
                ],
                cwd=ROOT,
                capture_output=True,
                text=True,
            )
            self.assertEqual(proc.returncode, 0, proc.stdout + proc.stderr)
            notes = output.read_text(encoding="utf-8")

        self.assertIn(
            "Haxe.Go is a pre-1.0 beta for pinned, application-qualified portable workloads",
            notes,
        )
        self.assertIn("`v0.54.0`", notes)
        self.assertIn(f"`{SOURCE_SHA}`", notes)
        self.assertIn("portable", notes)
        self.assertIn("Linux/amd64", notes)
        self.assertIn("Haxe 4.3.7", notes)
        self.assertIn("Go 1.25.12 and 1.26.5", notes)
        self.assertIn("metal", notes)
        self.assertIn("HTTP", notes)
        self.assertIn("named/reverse DNS", notes)
        self.assertIn("IPv6", notes)
        self.assertIn("untrusted", notes)
        self.assertIn("stable 1.x", notes)
        self.assertIn("Haxelib ZIP", notes)
        self.assertIn("SHA-256", notes)
        self.assertIn("provenance", notes)
        self.assertNotIn("published validated beta-baseline artifact", notes)
        self.assertNotIn("Published checksummed same-SHA Haxelib release assets`, `excluded`", notes)


if __name__ == "__main__":
    unittest.main()
