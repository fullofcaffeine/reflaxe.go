#!/usr/bin/env python3

from __future__ import annotations

import json
import subprocess
import unittest
from pathlib import Path


ROOT = Path(__file__).resolve().parent.parent
SUMMARY_JSON = ROOT / "test" / ".test-cache" / "portable_parity_closure_summary.json"
POLICY_LOCKED_MODULES = {
    "haxe.CallStack": "target_sensitive_snapshot",
    "haxe.EntryPoint": "target_sensitive_snapshot",
    "haxe.MainLoop": "target_sensitive_snapshot",
    "haxe.NativeStackTrace": "target_sensitive_snapshot",
    "haxe.Timer": "target_sensitive_snapshot",
    "haxe.Ucs2": "target_sensitive_snapshot",
    "haxe.http.HttpJs": "explicit_exclusion",
    "haxe.http.HttpNodeJs": "explicit_exclusion",
    "sys.net.UdpSocket": "target_sensitive_snapshot",
    "sys.ssl.Certificate": "target_sensitive_snapshot",
    "sys.ssl.Digest": "target_sensitive_snapshot",
    "sys.ssl.Key": "target_sensitive_snapshot",
    "sys.ssl.Socket": "target_sensitive_snapshot",
}


class PortableParityClosureContractTest(unittest.TestCase):
    def test_unsupported_modules_are_reported_as_explicit_exclusions(self) -> None:
        subprocess.run(
            ["python3", "test/run-portable-parity-closure.py"],
            cwd=ROOT,
            check=True,
            stdout=subprocess.PIPE,
            stderr=subprocess.STDOUT,
            text=True,
        )
        summary = json.loads(SUMMARY_JSON.read_text(encoding="utf-8"))
        unsupported = [
            blocker
            for blocker in summary["remaining_blockers"]
            if blocker["status"] == "unsupported"
        ]
        self.assertTrue(unsupported, "expected at least one explicit unsupported module")
        for blocker in unsupported:
            self.assertEqual(
                blocker["next_step"],
                "keep explicit exclusion policy, or promote only if the module becomes portable-eligible on Go",
            )

    def test_no_compile_only_modules_remain(self) -> None:
        subprocess.run(
            ["python3", "test/run-portable-parity-closure.py"],
            cwd=ROOT,
            check=True,
            stdout=subprocess.PIPE,
            stderr=subprocess.STDOUT,
            text=True,
        )
        summary = json.loads(SUMMARY_JSON.read_text(encoding="utf-8"))
        self.assertEqual(summary["status_counts"]["compile-only"], 0)

    def test_target_sensitive_policy_is_documented_per_surface(self) -> None:
        known_gaps = (ROOT / "docs" / "known-gaps.md").read_text(encoding="utf-8")
        feature_matrix = (ROOT / "docs" / "feature-support-matrix.md").read_text(encoding="utf-8")

        self.assertIn("## Target-sensitive parity policy", known_gaps)
        self.assertIn("| Surface | Current evidence | Decision | Why this is not semantic-diff yet | Reopen trigger |", known_gaps)
        for module in sorted(POLICY_LOCKED_MODULES):
            if module.startswith("haxe.http.") or module == "haxe.Ucs2":
                continue
            self.assertIn(f"`{module}`", known_gaps)
        self.assertIn("`sys.ssl.Certificate` / `sys.ssl.Digest` / `sys.ssl.Key`", known_gaps)
        self.assertIn("`haxe.EntryPoint` / `haxe.MainLoop` / `haxe.Timer`", known_gaps)
        self.assertIn("target-sensitive means", known_gaps)
        self.assertIn("snapshot/runtime", known_gaps)
        self.assertIn("must not be promoted to `semantic-diff`", known_gaps)
        self.assertIn("event-loop-semantic-diff-spike.md", known_gaps)
        self.assertIn("Haxe `--interp` does not execute `EntryPoint.runInMainThread` callbacks", known_gaps)
        self.assertIn("Known non-semantic-diff stdlib surfaces", feature_matrix)

    def test_remaining_non_semantic_diff_surfaces_are_policy_locked(self) -> None:
        subprocess.run(
            ["python3", "test/run-portable-parity-closure.py"],
            cwd=ROOT,
            check=True,
            stdout=subprocess.PIPE,
            stderr=subprocess.STDOUT,
            text=True,
        )
        summary = json.loads(SUMMARY_JSON.read_text(encoding="utf-8"))
        self.assertEqual(summary["actionable_blocker_count"], 0)
        blockers = {blocker["module"]: blocker for blocker in summary["remaining_blockers"]}
        self.assertEqual(set(blockers), set(POLICY_LOCKED_MODULES))
        for module, expected_policy in POLICY_LOCKED_MODULES.items():
            blocker = blockers[module]
            self.assertFalse(blocker["actionable"], module)
            self.assertEqual(blocker["closure_policy"], expected_policy, module)


if __name__ == "__main__":
    unittest.main()
