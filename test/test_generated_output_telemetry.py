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
    def make_case(self, module):
        metadata = module.ExampleMetadata(
            example_id="demo",
            tier="capability-showcase",
            claim_bearing=True,
            profiles=("portable",),
            test_command="npm run test:examples",
            lanes={
                "default": module.ExampleLaneMetadata(
                    product_surfaces=("portable-compiler",),
                    evidence_modes=("go-build-run",),
                )
            },
        )
        return module.ExampleProfileCase(
            example="demo",
            profile="portable",
            example_dir=Path("demo"),
            compile_hxml=Path("compile.hxml"),
            compile_ci_hxml=Path("compile.ci.hxml"),
            out_dir=Path("out"),
            out_ci_dir=Path("out_ci"),
            expected_stdout=Path("expected.stdout"),
            expected_ci_stdout=Path("expected.ci.stdout"),
            generated_dir=Path("generated"),
            metadata=metadata,
        )

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

    def test_unfinished_or_compile_only_lane_cannot_publish_claim_evidence(self) -> None:
        module = load_run_examples_module()
        case = self.make_case(module)
        entry = module.annotate_telemetry(case, "default", {"goTestOk": True}, compile_only=True)
        self.assertTrue(entry["declaredClaimBearing"])
        self.assertFalse(entry["claimBearing"])
        self.assertEqual([], entry["productSurfaces"])
        self.assertEqual([], entry["evidenceModes"])
        self.assertEqual("skipped", entry["runtimeStatus"])
        self.assertEqual("diagnostic-only", entry["claimStatus"])

    def test_runtime_and_stdout_completion_activate_the_declared_lane_claim(self) -> None:
        module = load_run_examples_module()
        case = self.make_case(module)
        entry = module.annotate_telemetry(case, "default", {"goTestOk": True}, compile_only=False)
        module.complete_lane_telemetry(entry, runtime_ok=True, stdout_ok=False)
        self.assertFalse(entry["claimBearing"])
        self.assertEqual("failed", entry["stdoutStatus"])

        module.complete_lane_telemetry(entry, runtime_ok=True, stdout_ok=True)
        self.assertTrue(entry["claimBearing"])
        self.assertEqual(["portable-compiler"], entry["productSurfaces"])
        self.assertEqual(["go-build-run"], entry["evidenceModes"])
        self.assertEqual("supported", entry["claimStatus"])

    def test_failed_case_report_cannot_retain_an_earlier_lane_claim(self) -> None:
        module = load_run_examples_module()
        case = self.make_case(module)
        entry = module.annotate_telemetry(case, "default", {"goTestOk": True}, compile_only=False)
        module.complete_lane_telemetry(entry, runtime_ok=True, stdout_ok=True)
        report = module.build_telemetry_report(
            [module.CaseResult("demo/portable", False, "generated", "drift", 1.0, [entry])]
        )
        published = report["entries"][0]
        self.assertFalse(published["claimBearing"])
        self.assertEqual([], published["productSurfaces"])
        self.assertEqual("case-failed", published["claimStatus"])


if __name__ == "__main__":
    unittest.main()
