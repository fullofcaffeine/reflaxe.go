#!/usr/bin/env python3

from __future__ import annotations

import json
import subprocess
import unittest
from pathlib import Path


ROOT = Path(__file__).resolve().parent.parent
SUMMARY_JSON = ROOT / "test" / ".test-cache" / "portable_parity_closure_summary.json"


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


if __name__ == "__main__":
    unittest.main()
