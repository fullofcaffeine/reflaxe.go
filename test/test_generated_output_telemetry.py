#!/usr/bin/env python3
from __future__ import annotations

import importlib.util
import sys
import tempfile
import unittest
from pathlib import Path


REPO_ROOT = Path(__file__).resolve().parent.parent
RUN_EXAMPLES = REPO_ROOT / "test" / "run-examples.py"


def load_run_examples_module():
    spec = importlib.util.spec_from_file_location("run_examples", RUN_EXAMPLES)
    if spec is None or spec.loader is None:
        raise RuntimeError("failed to load run-examples.py")
    module = importlib.util.module_from_spec(spec)
    sys.modules["run_examples"] = module
    spec.loader.exec_module(module)
    return module


class GeneratedOutputTelemetryTests(unittest.TestCase):
    def test_collect_output_telemetry_records_size_and_largest_go_file(self) -> None:
        module = load_run_examples_module()
        with tempfile.TemporaryDirectory() as raw_dir:
            root = Path(raw_dir)
            (root / "main.go").write_text("package main\n", encoding="utf-8")
            (root / "nested").mkdir()
            (root / "nested" / "large.go").write_text("package main\n" + ("// x\n" * 20), encoding="utf-8")
            (root / "note.txt").write_text("ignored\n", encoding="utf-8")

            telemetry = module.collect_output_telemetry(root)

        self.assertEqual(telemetry["goFileCount"], 2)
        self.assertGreater(telemetry["totalGoBytes"], telemetry["largestGoFileBytes"])
        self.assertTrue(str(telemetry["largestGoFile"]).endswith("nested/large.go"))
        self.assertEqual(telemetry["goTestCommand"], "go test ./...")

    def test_telemetry_report_is_sorted_and_markdown_renderable(self) -> None:
        module = load_run_examples_module()
        result_a = module.CaseResult(
            "zeta/metal",
            True,
            "done",
            "ok",
            1.0,
            [{"caseId": "zeta/metal", "lane": "default", "goFileCount": 1, "totalGoBytes": 10, "largestGoFile": "z/main.go", "largestGoFileBytes": 10, "goTestElapsedMs": 2.5}],
        )
        result_b = module.CaseResult(
            "alpha/portable",
            True,
            "done",
            "ok",
            1.0,
            [{"caseId": "alpha/portable", "lane": "ci", "goFileCount": 2, "totalGoBytes": 20, "largestGoFile": "a/main.go", "largestGoFileBytes": 12, "goTestElapsedMs": 3.0}],
        )

        report = module.build_telemetry_report([result_a, result_b])
        self.assertEqual([entry["caseId"] for entry in report["entries"]], ["alpha/portable", "zeta/metal"])

        markdown = module.render_telemetry_markdown(report)
        self.assertIn("# Generated Output Telemetry", markdown)
        self.assertIn("alpha/portable", markdown)
        self.assertIn("zeta/metal", markdown)


if __name__ == "__main__":
    unittest.main()
