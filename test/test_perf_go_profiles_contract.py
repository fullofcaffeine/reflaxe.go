#!/usr/bin/env python3

from __future__ import annotations

import unittest
from pathlib import Path


REPO_ROOT = Path(__file__).resolve().parent.parent
SCRIPT_PATH = REPO_ROOT / "scripts" / "ci" / "perf-go-profiles.sh"


class PerfGoProfilesContractTest(unittest.TestCase):
    def test_tui_perf_lane_skips_missing_profile_compile_files(self) -> None:
        text = SCRIPT_PATH.read_text(encoding="utf-8")
        self.assertIn("compile_file=", text)
        self.assertIn("if [[ ! -f \"$compile_file\" ]]; then", text)
        self.assertIn("tui case ($profile) skipped", text)

    def test_tui_summary_profiles_are_derived_from_collected_metrics(self) -> None:
        text = SCRIPT_PATH.read_text(encoding="utf-8")
        self.assertIn("const tuiProfiles =", text)
        self.assertIn('metric.case === "tui"', text)
        self.assertIn('ratioTable("TUI Profile Spread', text)


if __name__ == "__main__":
    unittest.main()
