#!/usr/bin/env python3

from __future__ import annotations

import json
import shutil
import subprocess
import tempfile
import unittest
from pathlib import Path


ROOT = Path(__file__).resolve().parent.parent
RUNNER = ROOT / "scripts" / "vendor" / "verify-reflaxe-provenance.py"
MANIFEST = ROOT / "provenance" / "reflaxe" / "vendor-manifest.json"
PATCH = ROOT / "provenance" / "reflaxe" / "upstream-to-supplier.patch"
DOC = ROOT / "docs" / "vendor-reflaxe-provenance.md"
VENDOR = ROOT / "vendor" / "reflaxe"


class VendorReflaxeProvenanceTest(unittest.TestCase):
    def test_manifest_records_exact_origin_chain_and_tree_identity(self) -> None:
        manifest = json.loads(MANIFEST.read_text(encoding="utf-8"))

        self.assertEqual(manifest["schema_version"], 1)
        self.assertEqual(
            manifest["official_upstream"]["repository"],
            "https://github.com/SomeRanDev/reflaxe.git",
        )
        self.assertEqual(
            manifest["official_upstream"]["commit"],
            "430b4187a6bf4813cf618fc3a73ccf494a2ab9f5",
        )
        self.assertEqual(
            manifest["supplier_snapshot"]["repository"],
            "https://github.com/fullofcaffeine/reflaxe.rust.git",
        )
        self.assertEqual(
            manifest["supplier_snapshot"]["commit"],
            "f53bec2adae8ef000467e488b974f6514c1af98f",
        )
        self.assertEqual(
            manifest["supplier_snapshot"]["git_tree_sha1"],
            "a26137d0af5f297eb12e3750a62d0544a4755b76",
        )
        self.assertEqual(manifest["shipped_tree"]["file_count"], 63)
        self.assertRegex(manifest["shipped_tree"]["sha256"], r"^[0-9a-f]{64}$")
        self.assertEqual(len(manifest["shipped_tree"]["files"]), 63)
        self.assertRegex(manifest["patch_set"]["sha256"], r"^[0-9a-f]{64}$")
        self.assertEqual(len(manifest["patch_set"]["modified_source_files"]), 14)
        self.assertEqual(
            manifest["shipped_overlays"],
            [
                {
                    "kind": "license-restoration",
                    "sourceAuthority": "official_upstream",
                    "sourcePath": "LICENSE",
                    "path": "LICENSE",
                    "sha256": "68f17b4096da082538e5a6e2c050cd510f47fe33574522f3dc397750da600e79",
                }
            ],
        )

    def test_offline_verifier_checks_tree_and_patch_round_trip(self) -> None:
        proc = subprocess.run(
            ["python3", str(RUNNER)],
            cwd=ROOT,
            capture_output=True,
            text=True,
            timeout=30,
        )

        self.assertEqual(proc.returncode, 0, proc.stdout + proc.stderr)
        self.assertIn("63 files", proc.stdout)
        self.assertIn("1 overlay", proc.stdout)
        self.assertIn("patch round-trip: OK", proc.stdout)

    def test_verifier_rejects_tampered_vendor_source(self) -> None:
        temp_dir = tempfile.TemporaryDirectory()
        self.addCleanup(temp_dir.cleanup)
        copied_vendor = Path(temp_dir.name) / "reflaxe"
        shutil.copytree(VENDOR, copied_vendor)
        tampered = copied_vendor / "haxelib.json"
        tampered.write_text(tampered.read_text(encoding="utf-8") + "\n", encoding="utf-8")

        proc = subprocess.run(
            ["python3", str(RUNNER), "--vendor-dir", str(copied_vendor)],
            cwd=ROOT,
            capture_output=True,
            text=True,
            timeout=30,
        )

        self.assertEqual(proc.returncode, 1, proc.stdout + proc.stderr)
        self.assertIn("haxelib.json", proc.stderr)
        self.assertIn("digest mismatch", proc.stderr)

    def test_patch_and_policy_document_reconstruction(self) -> None:
        patch = PATCH.read_text(encoding="utf-8")
        self.assertIn("diff --git a/src/reflaxe/ReflectCompiler.hx", patch)
        self.assertIn("diff --git a/src/reflaxe/output/OutputManager.hx", patch)
        self.assertNotIn("/Users/", patch)
        self.assertNotIn("/tmp/", patch)

        doc = DOC.read_text(encoding="utf-8")
        self.assertIn("# Vendored Reflaxe Provenance", doc)
        self.assertIn("430b4187a6bf4813cf618fc3a73ccf494a2ab9f5", doc)
        self.assertIn("f53bec2adae8ef000467e488b974f6514c1af98f", doc)
        self.assertIn("a26137d0af5f297eb12e3750a62d0544a4755b76", doc)
        self.assertIn("npm run verify:vendor-reflaxe", doc)
        self.assertIn("npm run verify:vendor-reflaxe:reconstruct", doc)
        self.assertIn("14 modified framework source files", doc)
        self.assertIn("does not decide redistribution obligations", doc)


if __name__ == "__main__":
    unittest.main()
