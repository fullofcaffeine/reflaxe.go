#!/usr/bin/env python3
from __future__ import annotations

import json
import subprocess
import tempfile
import unittest
from pathlib import Path


REPO_ROOT = Path(__file__).resolve().parent.parent
SCRIPT = REPO_ROOT / "scripts" / "ci" / "perf-warning-summary.py"


class PerfWarningSummaryTests(unittest.TestCase):
    def run_summary(self, harness: str, warnings: list[str]) -> dict:
        with tempfile.TemporaryDirectory() as raw_dir:
            tmp = Path(raw_dir)
            comparison = tmp / "comparison.json"
            out_json = tmp / "warning_history.json"
            out_md = tmp / "warning_history.md"
            comparison.write_text(json.dumps({"warnings": warnings}) + "\n", encoding="utf-8")

            subprocess.run(
                [
                    "python3",
                    str(SCRIPT),
                    "--harness",
                    harness,
                    "--comparison",
                    str(comparison),
                    "--out-json",
                    str(out_json),
                    "--out-md",
                    str(out_md),
                ],
                cwd=REPO_ROOT,
                check=True,
            )

            self.assertTrue(out_md.read_text(encoding="utf-8").startswith("# Perf Warning History"))
            return json.loads(out_json.read_text(encoding="utf-8"))

    def test_go_profile_warning_groups_are_stable(self) -> None:
        report = self.run_summary(
            "go-profile",
            [
                "generic_overhead.metal.startup ratio +10.60% (current=1.400000, baseline=1.265823, budget=+10.00%)",
                "delta.channel.startup ratio +22.00% (current=1.220000, baseline=1.000000, budget=+15.00%)",
            ],
        )

        self.assertEqual(report["schemaVersion"], 1)
        self.assertEqual(report["warningCount"], 2)
        groups = {group["groupKey"]: group for group in report["groups"]}
        profile_group = next(group for group in groups.values() if group["kind"] == "profile")
        self.assertEqual(profile_group["case"], "generic_overhead")
        self.assertEqual(profile_group["profile"], "metal")
        self.assertEqual(profile_group["metric"], "startup")
        self.assertEqual(profile_group["thresholdPct"], 10.0)
        delta_group = next(group for group in groups.values() if group["kind"] == "delta")
        self.assertEqual(delta_group["case"], "channel")
        self.assertEqual(delta_group["profile"], "portable_vs_metal")
        self.assertEqual(delta_group["thresholdPct"], 15.0)

    def test_app_warning_groups_capture_app_variant_profile_and_metric(self) -> None:
        report = self.run_summary(
            "app-profile",
            [
                "fluxproxy::core::portable.startup_ratio_vs_pure rose 34.35% (current=1.472603, baseline=1.096080, budget=15.00%)",
                "delta.pulseforge::go_native.throughput_delta dropped 29.68% (current=0.770206, baseline=1.095259, budget=15.00%)",
            ],
        )

        self.assertEqual(report["warningCount"], 2)
        profile_group = next(group for group in report["groups"] if group["kind"] == "profile")
        self.assertEqual(profile_group["app"], "fluxproxy")
        self.assertEqual(profile_group["variant"], "core")
        self.assertEqual(profile_group["profile"], "portable")
        self.assertEqual(profile_group["metric"], "startup_ratio_vs_pure")
        self.assertEqual(profile_group["thresholdPct"], 15.0)

        delta_group = next(group for group in report["groups"] if group["kind"] == "delta")
        self.assertEqual(delta_group["app"], "pulseforge")
        self.assertEqual(delta_group["variant"], "go_native")
        self.assertEqual(delta_group["profile"], "portable_vs_metal")
        self.assertEqual(delta_group["metric"], "throughput_delta")
        self.assertEqual(delta_group["warnings"][0]["driftPct"], -29.68)


if __name__ == "__main__":
    unittest.main()
