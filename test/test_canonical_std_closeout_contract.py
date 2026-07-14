#!/usr/bin/env python3

from __future__ import annotations

from collections import Counter
import json
from pathlib import Path
import shutil
import subprocess
import sys
import tempfile
import unittest


ROOT = Path(__file__).resolve().parent.parent
SMOKE_RUNNER = ROOT / "scripts" / "ci" / "run-isolated-haxelib-smoke.py"
CLOSEOUT_DOC = ROOT / "docs" / "canonical-std-migration-closeout.md"
DOCS_INDEX = ROOT / "docs" / "index.md"
START_HERE = ROOT / "docs" / "start-here.md"
HARDENING_DOC = ROOT / "docs" / "cross-overrides-and-hardening.md"
STATUS_PATH = ROOT / "test" / "canonical_std_layout_status.json"
LEDGER_PATH = ROOT / "docs" / "stdlib-provenance-ledger.json"
PACKAGE_PATH = ROOT / "package.json"
RELEASE_RUNNER_PATH = ROOT / "test" / "run-release-contracts.py"


EXPECTED_INVENTORY = {
    "canonicalOverrideSources": 62,
    "ledgerEntries": 82,
    "ownershipClasses": {
        "hxrt_binding": 10,
        "public_go_facade": 5,
        "staged_support": 5,
        "upstream_std_override": 62,
    },
    "packagedArchiveMembers": 225,
    "packagedManifestEntries": 224,
    "trackedCrossHxFiles": 0,
}


def tracked_paths(*patterns: str) -> set[str]:
    process = subprocess.run(
        ["git", "ls-files", "-z", "--", *patterns],
        cwd=ROOT,
        check=True,
        capture_output=True,
    )
    return {
        item.decode("utf-8")
        for item in process.stdout.split(b"\0")
        if item
    }


class CanonicalStdCloseoutContractTest(unittest.TestCase):
    maxDiff = None

    def test_closeout_evidence_is_documented_and_release_blocking(self) -> None:
        self.assertTrue(SMOKE_RUNNER.is_file())
        self.assertTrue(CLOSEOUT_DOC.is_file())

        package = json.loads(PACKAGE_PATH.read_text(encoding="utf-8"))
        scripts = package.get("scripts", {})
        contract = "python3 test/test_canonical_std_closeout_contract.py"
        self.assertEqual(contract, scripts.get("test:canonical-std-closeout"))
        self.assertIn("npm run test:canonical-std-closeout", scripts.get("test", ""))
        self.assertIn(
            "npm run test:canonical-std-closeout",
            scripts.get("test:changed", ""),
        )
        self.assertIn(
            '"test/test_canonical_std_closeout_contract.py"',
            RELEASE_RUNNER_PATH.read_text(encoding="utf-8"),
        )

        closeout_doc = " ".join(
            CLOSEOUT_DOC.read_text(encoding="utf-8").split()
        )
        for phrase in (
            "ordinary `.hx` under `std/go/_std`",
            "generated `.cross.hx` under `src`",
            "zero generated-Go or stdout snapshot changes",
            "test:stdlib-sweep:go-test",
            "run-isolated-haxelib-smoke.py",
        ):
            self.assertIn(phrase, closeout_doc)
        self.assertIn(
            "canonical-std-migration-closeout.md",
            DOCS_INDEX.read_text(encoding="utf-8"),
        )
        self.assertIn(
            "canonical-std-migration-closeout.md",
            START_HERE.read_text(encoding="utf-8"),
        )
        hardening_doc = " ".join(
            HARDENING_DOC.read_text(encoding="utf-8").split()
        )
        self.assertIn(
            "Package generation and isolated installed-package selection are now enforced",
            hardening_doc,
        )

    def test_declared_closeout_inventory_matches_every_tracked_std_source(self) -> None:
        status = json.loads(STATUS_PATH.read_text(encoding="utf-8"))
        self.assertEqual(2, status.get("schemaVersion"))
        self.assertEqual("required-green", status.get("sourceLayout"))
        self.assertEqual([], status.get("expectedViolationCodes"))
        self.assertEqual(
            "packaging-only-no-generated-output-change",
            status.get("snapshotChangePolicy"),
        )
        self.assertEqual(EXPECTED_INVENTORY, status.get("closeoutEvidence"))

        ledger = json.loads(LEDGER_PATH.read_text(encoding="utf-8"))
        entries = ledger.get("entries", [])
        ownership_counts = Counter(entry["ownershipClass"] for entry in entries)
        override_sources = {
            entry["path"]
            for entry in entries
            if entry["ownershipClass"] == "upstream_std_override"
        }
        canonical_sources = {
            path
            for path in tracked_paths("std/go/_std")
            if path.endswith(".hx")
        }
        cross_sources = tracked_paths("*.cross.hx")

        self.assertEqual(EXPECTED_INVENTORY["ledgerEntries"], len(entries))
        self.assertEqual(
            EXPECTED_INVENTORY["ownershipClasses"],
            dict(sorted(ownership_counts.items())),
        )
        self.assertEqual(canonical_sources, override_sources)
        self.assertEqual(
            EXPECTED_INVENTORY["canonicalOverrideSources"],
            len(canonical_sources),
        )
        self.assertEqual(set(), cross_sources)

    @unittest.skipUnless(
        shutil.which("haxe") and shutil.which("haxelib") and shutil.which("go"),
        "requires Haxe, Haxelib, and Go",
    )
    def test_real_zip_installs_and_runs_without_checkout_classpaths(self) -> None:
        process = subprocess.run(
            [sys.executable, str(SMOKE_RUNNER)],
            cwd=ROOT,
            capture_output=True,
            text=True,
            timeout=240,
        )
        self.assertEqual(0, process.returncode, process.stdout + process.stderr)
        evidence = json.loads(process.stdout)
        self.assertEqual(1, evidence.get("schemaVersion"))
        self.assertEqual(
            "haxe.go-canonical-std-isolated-package-smoke",
            evidence.get("kind"),
        )
        self.assertEqual(
            {
                "archiveMembers": EXPECTED_INVENTORY["packagedArchiveMembers"],
                "manifestEntries": EXPECTED_INVENTORY["packagedManifestEntries"],
                "stdlibOverrides": EXPECTED_INVENTORY["canonicalOverrideSources"],
                "version": "0.0.0",
            },
            evidence.get("package"),
        )
        self.assertEqual(
            {
                "checkoutClasspathsAbsent": True,
                "generatedPathLeaksAbsent": True,
                "goRun": "pass",
                "goTest": "pass",
                "haxeCompile": "pass",
                "haxelibInstall": "pass",
                "stdout": "pass",
            },
            evidence.get("checks"),
        )
        self.assertEqual(
            {
                "name": "stdlib/stringtools_cross_std_basic",
                "profile": "portable",
            },
            evidence.get("fixture"),
        )
        self.assertNotIn(str(ROOT), process.stdout)

    @unittest.skipUnless(shutil.which("haxelib"), "requires Haxelib")
    def test_invalid_archive_is_classified_as_a_package_failure(self) -> None:
        with tempfile.TemporaryDirectory(prefix="haxe-go-invalid-package-smoke-") as raw:
            invalid = Path(raw) / "invalid.zip"
            invalid.write_bytes(b"not a ZIP\n")
            process = subprocess.run(
                [sys.executable, str(SMOKE_RUNNER), "--archive", str(invalid)],
                cwd=ROOT,
                capture_output=True,
                text=True,
                timeout=30,
            )
        self.assertNotEqual(0, process.returncode)
        self.assertIn("ERROR [package]", process.stderr)


if __name__ == "__main__":
    unittest.main()
