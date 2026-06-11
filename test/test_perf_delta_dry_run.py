#!/usr/bin/env python3
from __future__ import annotations

import json
import subprocess
import tempfile
import unittest
from pathlib import Path


REPO_ROOT = Path(__file__).resolve().parent.parent
SCRIPT = REPO_ROOT / "scripts" / "ci" / "perf-delta-dry-run.py"


class PerfDeltaDryRunTests(unittest.TestCase):
    def run_report(self, harness: str, hard_failures: list[str]) -> dict:
        with tempfile.TemporaryDirectory() as raw_dir:
            tmp = Path(raw_dir)
            comparison = tmp / "comparison.json"
            out_json = tmp / "delta_hard_gate_dry_run.json"
            out_md = tmp / "delta_hard_gate_dry_run.md"
            comparison.write_text(
                json.dumps({"enforceDeltaBudget": False, "hardFailures": hard_failures}) + "\n",
                encoding="utf-8",
            )

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

            self.assertTrue(out_md.read_text(encoding="utf-8").startswith("# Delta Hard-Gate Dry Run"))
            return json.loads(out_json.read_text(encoding="utf-8"))

    def test_go_profile_delta_candidates_are_extracted_without_metal_candidates(self) -> None:
        report = self.run_report(
            "go-profile",
            [
                "string_overhead.metal.startup ratio +44.00% (current=1.440000, baseline=1.000000, budget=+25.00%)",
                "delta.channel.startup ratio +30.00% (current=1.300000, baseline=1.000000, budget=+25.00%)",
            ],
        )

        self.assertFalse(report["enforceDeltaBudget"])
        self.assertTrue(report["wouldFailIfEnforced"])
        self.assertEqual(report["candidateCount"], 1)
        candidate = report["candidates"][0]
        self.assertEqual(candidate["case"], "channel")
        self.assertEqual(candidate["profile"], "portable_vs_metal")
        self.assertEqual(candidate["metric"], "startup")
        self.assertEqual(candidate["thresholdPct"], 25.0)

    def test_app_delta_candidates_capture_app_and_variant(self) -> None:
        report = self.run_report(
            "app-profile",
            [
                "fluxproxy::core::metal.startup_ratio_vs_pure rose 40.00% (current=1.400000, baseline=1.000000, metal budget=30.00%)",
                "delta.pulseforge::go_native.throughput_delta dropped 29.68% (current=0.770206, baseline=1.095259, delta budget=25.00%)",
            ],
        )

        self.assertEqual(report["candidateCount"], 1)
        candidate = report["candidates"][0]
        self.assertEqual(candidate["app"], "pulseforge")
        self.assertEqual(candidate["variant"], "go_native")
        self.assertEqual(candidate["profile"], "portable_vs_metal")
        self.assertEqual(candidate["metric"], "throughput_delta")
        self.assertEqual(candidate["driftPct"], -29.68)


if __name__ == "__main__":
    unittest.main()
