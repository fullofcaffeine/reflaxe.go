#!/usr/bin/env python3

from __future__ import annotations

import importlib.util
import json
import tempfile
import unittest
from pathlib import Path
from types import ModuleType


ROOT = Path(__file__).resolve().parent.parent
RUNNER = ROOT / "test" / "run-compiler-debt-ratchet.py"
POLICY = ROOT / "test" / "compiler_debt_policy.json"
POLICY_DOC = ROOT / "docs" / "compiler-debt-ratchet.md"


def load_runner() -> ModuleType:
    spec = importlib.util.spec_from_file_location("compiler_debt_ratchet", RUNNER)
    if spec is None or spec.loader is None:
        raise RuntimeError(f"cannot load {RUNNER}")
    module = importlib.util.module_from_spec(spec)
    spec.loader.exec_module(module)
    return module


class CompilerDebtRatchetTest(unittest.TestCase):
    @classmethod
    def setUpClass(cls) -> None:
        cls.runner = load_runner()

    def test_scanner_separates_source_runtime_and_generated_profiles(self) -> None:
        with tempfile.TemporaryDirectory() as tmp:
            root = Path(tmp)
            compiler = root / "src" / "reflaxe" / "go" / "GoCompiler.hx"
            compiler.parent.mkdir(parents=True)
            compiler.write_text(
                """
                // Dynamic Any GoStmt.GoRaw(\"ignored\")
                function lowerIoStdlibShimDecls() {
                    var label = \"Dynamic Any\";
                    function nestedEmitter() {
                        return GoExpr.GoRaw(\"nested\");
                    }
                    return GoStmt.GoRaw(\"if true {}\");
                }
                function typedHelper(value:Dynamic):Any {
                    var qualified = GoExpr.GoRaw(\"value\");
                    return GoRaw(\"bare constructor\");
                }
                """,
                encoding="utf-8",
            )
            runtime = root / "runtime" / "hxrt" / "value.go"
            runtime.parent.mkdir(parents=True)
            runtime.write_text(
                'package hxrt\nimport (\n r "reflect"\n u "unsafe"\n)\n'
                "func value(v any) { _ = r.ValueOf(v); _ = u.Pointer(nil) }\n",
                encoding="utf-8",
            )
            portable = root / "examples" / "demo" / "generated" / "portable" / "main.go"
            portable.parent.mkdir(parents=True)
            portable.write_text(
                'package main\nimport "reflect"\nfunc main() { _ = reflect.ValueOf(1) }\n',
                encoding="utf-8",
            )
            metal = root / "examples" / "demo" / "generated" / "metal" / "main.go"
            metal.parent.mkdir(parents=True)
            metal.write_text(
                'package main\nimport "reflect"\nfunc main() { _ = reflect.ValueOf(1) }\n',
                encoding="utf-8",
            )

            report = self.runner.build_report(self.runner.collect_findings(root))

        self.assertEqual(4, report["totals"]["go_raw"])
        self.assertEqual(1, report["totals"]["haxe_dynamic"])
        self.assertEqual(1, report["totals"]["haxe_any"])
        self.assertEqual(2, report["totals"]["go_unsafe"])
        self.assertEqual(6, report["totals"]["go_reflection"])
        self.assertEqual(1, report["totals"]["compiler_shim"])
        raw_contexts = {
            row["context"]
            for row in report["by_file"]
            if row["metric"] == "go_raw" and row["file"].endswith("GoCompiler.hx")
        }
        self.assertEqual({"lowerIoStdlibShimDecls", "typedHelper"}, raw_contexts)
        profiles = {row["profile"]: row["count"] for row in report["by_profile"]}
        self.assertEqual(2, profiles["portable"])
        self.assertEqual(2, profiles["metal"])
        self.assertGreater(profiles["shared"], 0)

    def test_ratchet_rejects_new_locations_and_increases_but_accepts_reductions(self) -> None:
        key = {
            "metric": "go_raw",
            "file": "src/reflaxe/go/GoCompiler.hx",
            "context": "typedHelper",
            "owner": "compiler_core",
            "capability": "typed_lowering",
            "profile": "shared",
            "surface": "compiler",
            "classification": "avoidable",
            "exception_id": "typed_ast_gap",
        }
        policy = {
            "schema_version": 1,
            "guarded_metrics": list(self.runner.GUARDED_METRICS),
            "exceptions": {
                "typed_ast_gap": {
                    "owner": "compiler",
                    "classification": "avoidable",
                    "what": "Raw Go AST construction.",
                    "why": "The typed AST does not yet model every required Go form.",
                    "how": "Reduce through typed AST migrations; do not raise the ceiling.",
                }
            },
            "limits": [{**key, "max_count": 2}],
        }

        reduced = self.runner.build_report([key])
        self.assertEqual([], self.runner.compare_report_to_policy(reduced, policy))

        increased = self.runner.build_report([key, key, key])
        increase_errors = self.runner.compare_report_to_policy(increased, policy)
        self.assertTrue(any("exceeds baseline" in error for error in increase_errors))

        new_key = {**key, "file": "src/reflaxe/go/NewEmitter.hx"}
        unexplained = self.runner.build_report([key, new_key])
        new_errors = self.runner.compare_report_to_policy(unexplained, policy)
        self.assertTrue(any("unexplained" in error for error in new_errors))

    def test_repository_policy_is_complete_and_current(self) -> None:
        policy = json.loads(POLICY.read_text(encoding="utf-8"))
        validation_errors = self.runner.validate_policy(policy)
        self.assertEqual([], validation_errors)

        report = self.runner.build_report(self.runner.collect_findings(ROOT))
        ratchet_errors = self.runner.compare_report_to_policy(report, policy)
        self.assertEqual([], ratchet_errors)
        self.assertEqual(set(self.runner.GUARDED_METRICS), set(report["totals"]))
        self.assertEqual(0, report["totals"]["go_unsafe"])

        rendered = json.dumps(report, sort_keys=True)
        self.assertNotIn(str(ROOT), rendered)
        self.assertFalse(any(Path(row["file"]).is_absolute() for row in report["by_file"]))

    def test_every_exception_has_owner_and_why_what_how_rationale(self) -> None:
        policy = json.loads(POLICY.read_text(encoding="utf-8"))
        self.assertTrue(policy["exceptions"])
        classifications: set[str] = set()
        for exception_id, exception in policy["exceptions"].items():
            with self.subTest(exception=exception_id):
                self.assertTrue(exception["owner"])
                self.assertIn(exception["classification"], {"required", "avoidable"})
                classifications.add(exception["classification"])
                for field in ("what", "why", "how"):
                    self.assertGreaterEqual(len(exception[field].strip()), 20)
        self.assertEqual({"required", "avoidable"}, classifications)

    def test_reports_cover_every_required_dimension(self) -> None:
        report = self.runner.build_report(self.runner.collect_findings(ROOT))
        for dimension in (
            "by_file",
            "by_owner",
            "by_capability",
            "by_profile",
            "by_surface",
            "by_classification",
        ):
            self.assertIn(dimension, report)
            self.assertTrue(report[dimension])

        markdown = self.runner.render_markdown(report)
        for heading in ("By metric", "By owner", "By capability", "By profile", "By surface", "By file"):
            self.assertIn(heading, markdown)

    def test_ci_and_docs_expose_the_ratchet(self) -> None:
        package = json.loads((ROOT / "package.json").read_text(encoding="utf-8"))
        scripts = package["scripts"]
        self.assertEqual("python3 test/run-compiler-debt-ratchet.py", scripts["test:compiler-debt"])
        self.assertIn("npm run test:compiler-debt", scripts["test:changed"])

        release_runner = (ROOT / "test" / "run-release-contracts.py").read_text(encoding="utf-8")
        self.assertIn("test/test_compiler_debt_ratchet.py", release_runner)
        workflow = (ROOT / ".github" / "workflows" / "ci-harness.yml").read_text(encoding="utf-8")
        self.assertIn("name: compiler-debt-report", workflow)
        self.assertIn("path: .cache/compiler-debt", workflow)
        release_checklist = (ROOT / "docs" / "release-readiness-checklist.md").read_text(encoding="utf-8")
        self.assertIn("npm run test:compiler-debt", release_checklist)

        policy_doc = POLICY_DOC.read_text(encoding="utf-8")
        for phrase in (
            "Counts are directional evidence",
            "required",
            "avoidable",
            "portable",
            "metal",
            ".cache/compiler-debt/report.json",
        ):
            self.assertIn(phrase, policy_doc)


if __name__ == "__main__":
    unittest.main()
