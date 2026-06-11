#!/usr/bin/env python3

from __future__ import annotations

import json
import unittest
from pathlib import Path


REPO_ROOT = Path(__file__).resolve().parent.parent
SCRIPT_PATH = REPO_ROOT / "scripts" / "ci" / "perf-go-profiles.sh"
APP_SCRIPT_PATH = REPO_ROOT / "scripts" / "ci" / "perf-apps.sh"
GO_PROFILE_BASELINE = REPO_ROOT / "scripts" / "ci" / "perf" / "go-profile-baseline.json"
APP_PROFILE_BASELINE = REPO_ROOT / "scripts" / "ci" / "perf" / "app-profile-baseline.json"


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

    def test_microbench_compiles_use_selective_hxrt_runtime(self) -> None:
        text = SCRIPT_PATH.read_text(encoding="utf-8")
        self.assertIn('GO_PERF_HXRT_FEATURES', text)
        self.assertIn('-D "reflaxe_go_hxrt_features=$hxrt_features"', text)

    def test_checked_in_perf_baselines_keep_toolchain_and_methodology(self) -> None:
        go_profile = SCRIPT_PATH.read_text(encoding="utf-8")
        app_profile = APP_SCRIPT_PATH.read_text(encoding="utf-8")

        self.assertIn("toolchain: current.toolchain", go_profile)
        self.assertIn("options: current.options", go_profile)
        self.assertIn('"toolchain": current_payload["toolchain"]', app_profile)
        self.assertIn('"params": current_payload["params"]', app_profile)

    def test_checked_in_perf_baseline_files_record_provenance(self) -> None:
        go_profile = json.loads(GO_PROFILE_BASELINE.read_text(encoding="utf-8"))
        app_profile = json.loads(APP_PROFILE_BASELINE.read_text(encoding="utf-8"))

        for baseline in (go_profile, app_profile):
            self.assertIn("toolchain", baseline)
            self.assertIn("haxe", baseline["toolchain"])
            self.assertIn("go", baseline["toolchain"])

        self.assertIn("options", go_profile)
        self.assertIn("params", app_profile)


if __name__ == "__main__":
	unittest.main()
