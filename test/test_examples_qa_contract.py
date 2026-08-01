#!/usr/bin/env python3

from __future__ import annotations

import importlib.util
from pathlib import Path
import re
import sys
from unittest import mock
import unittest


REPO_ROOT = Path(__file__).resolve().parent.parent
RUN_EXAMPLES = REPO_ROOT / "test" / "run-examples.py"
EXAMPLES_ROOT = REPO_ROOT / "examples"
EXAMPLES_MANIFEST = EXAMPLES_ROOT / "qa-manifest.json"
PACKAGE_JSON = REPO_ROOT / "package.json"
RUN_CI = REPO_ROOT / "test" / "run-ci.py"
CI_HARNESS = REPO_ROOT / ".github" / "workflows" / "ci-harness.yml"
AGENTS = REPO_ROOT / "AGENTS.md"
DOC = REPO_ROOT / "docs" / "examples-qa-contract.md"
DOCS_INDEX = REPO_ROOT / "docs" / "index.md"


def load_run_examples_module():
    spec = importlib.util.spec_from_file_location("run_examples", RUN_EXAMPLES)
    if spec is None or spec.loader is None:
        raise RuntimeError("failed to load run-examples.py")
    module = importlib.util.module_from_spec(spec)
    sys.modules["run_examples"] = module
    spec.loader.exec_module(module)
    return module


class ExamplesQaContractTest(unittest.TestCase):
    def test_every_example_directory_is_visible_to_examples_harness(self) -> None:
        module = load_run_examples_module()
        cases = module.discover_cases()
        discovered_examples = {case.example for case in cases}
        example_dirs = {
            path.name
            for path in EXAMPLES_ROOT.iterdir()
            if path.is_dir() and not path.name.startswith(".") and (path / "README.md").exists()
        }

        self.assertEqual(example_dirs, discovered_examples)

    def test_every_discovered_example_lane_has_compile_run_and_generated_contracts(self) -> None:
        module = load_run_examples_module()
        cases = module.discover_cases()
        self.assertGreater(len(cases), 0)

        missing: list[str] = []
        for case in cases:
            required = [
                case.compile_hxml,
                case.compile_ci_hxml,
                case.expected_stdout,
                case.expected_ci_stdout,
                case.generated_dir,
            ]
            for path in required:
                if not path.exists():
                    missing.append(f"{case.case_id}: missing {path.relative_to(REPO_ROOT).as_posix()}")

        self.assertEqual([], missing)

    def test_example_tiers_and_claim_surfaces_are_consumed_by_the_harness(self) -> None:
        module = load_run_examples_module()
        metadata = module.load_example_metadata()
        manifest = EXAMPLES_MANIFEST.read_text(encoding="utf-8")
        self.assertIn('"flagship-application"', manifest)
        self.assertIn('"capability-showcase"', manifest)
        self.assertEqual({case.example for case in module.discover_cases()}, set(metadata))
        for case in module.discover_cases():
            self.assertIn(case.profile, case.metadata.profiles)
            for lane in ("default", "ci"):
                declaration = case.metadata.lanes[lane]
                self.assertTrue(declaration.product_surfaces)
                self.assertTrue(declaration.evidence_modes)

    def test_native_example_evidence_is_declared_per_executed_lane(self) -> None:
        module = load_run_examples_module()
        metadata = module.load_example_metadata()

        for lane in ("default", "ci"):
            self.assertNotIn("go-native-metal", metadata["incident_api"].lanes[lane].product_surfaces)
            self.assertNotIn("native-metal", metadata["incident_api"].lanes[lane].evidence_modes)
            self.assertIn("go-native-metal", metadata["interop_smoke"].lanes[lane].product_surfaces)
            self.assertIn("native-metal", metadata["interop_smoke"].lanes[lane].evidence_modes)

        self.assertNotIn(
            "go-native-metal", metadata["fluxproxy"].lanes["default"].product_surfaces
        )
        self.assertIn("go-native-metal", metadata["fluxproxy"].lanes["ci"].product_surfaces)

    def test_examples_harness_compiles_tests_runs_and_checks_expected_output(self) -> None:
        text = RUN_EXAMPLES.read_text(encoding="utf-8")

        self.assertIn('["haxe", case.compile_ci_hxml.name, "-D", "go_no_build"]', text)
        self.assertIn('["haxe", case.compile_hxml.name, "-D", "go_no_build"]', text)
        self.assertIn('["go", "test", "./..."]', text)
        self.assertIn('["go", "run", ".", *run_ci_args]', text)
        self.assertIn('["go", "run", ".", *run_args]', text)
        self.assertIn("compare_stdout(case.expected_ci_stdout", text)
        self.assertIn("compare_stdout(case.expected_stdout", text)
        self.assertIn("collect_tree_deltas(case.generated_dir, case.out_dir)", text)
        self.assertIn("load_example_metadata()", text)

    def test_changed_examples_unions_all_git_states_and_fails_closed(self) -> None:
        module = load_run_examples_module()
        all_examples = set(module.load_example_metadata())
        changed = {
            "examples/fluxproxy/Main.hx",
            "examples/incident_api/expected/portable.stdout",
        }
        with mock.patch.object(module, "collect_changed_paths", return_value=changed) as collect:
            self.assertEqual({"fluxproxy", "incident_api"}, module.changed_examples())
        collect.assert_called_once()

        with mock.patch.object(
            module,
            "collect_changed_paths",
            side_effect=module.GitChangeDiscoveryError("deliberate discovery failure"),
        ):
            self.assertEqual(all_examples, module.changed_examples())

        with mock.patch.object(
            module,
            "collect_changed_paths",
            return_value={"examples/qa-manifest.json"},
        ):
            self.assertEqual(all_examples, module.changed_examples())

    def test_example_builds_use_canonical_library_classpaths(self) -> None:
        module = load_run_examples_module()
        failures: list[str] = []
        for case in module.discover_cases():
            for path in (case.compile_hxml, case.compile_ci_hxml):
                text = path.read_text(encoding="utf-8")
                relative = path.relative_to(REPO_ROOT).as_posix()
                classpath_directives = [
                    line.strip()
                    for line in text.splitlines()
                    if line.strip().startswith(("-cp ", "-p ", "--class-path ", "-lib ", "--library "))
                ]
                if classpath_directives != ["-cp .", "-lib reflaxe.go"]:
                    failures.append(
                        f"{relative}: unexpected source classpath directives {classpath_directives!r}"
                    )
                if "CompilerBootstrap.Start" in text or "CompilerInit.Start" in text:
                    failures.append(f"{relative}: duplicates library-owned compiler macros")

        self.assertEqual([], failures)

    def test_go_native_example_modules_declare_their_boundary(self) -> None:
        missing: list[str] = []
        for path in EXAMPLES_ROOT.rglob("*.hx"):
            relative = path.relative_to(REPO_ROOT)
            if "generated" in relative.parts or any(
                part.startswith("out_") for part in relative.parts
            ):
                continue

            text = path.read_text(encoding="utf-8")
            owns_native_api = bool(
                re.search(r"^\s*import\s+go\.", text, flags=re.MULTILINE)
                or re.search(r"\bgo\.[A-Z][A-Za-z0-9_]*", text)
                or "@:go.import" in text
            )
            if owns_native_api and "@:goNative" not in text:
                missing.append(relative.as_posix())

        self.assertEqual(
            [],
            missing,
            "example modules that own typed Go APIs must declare @:goNative",
        )

    def test_ci_and_agent_docs_treat_examples_as_qa_contracts(self) -> None:
        package_json = PACKAGE_JSON.read_text(encoding="utf-8")
        run_ci = RUN_CI.read_text(encoding="utf-8")
        ci_harness = CI_HARNESS.read_text(encoding="utf-8")
        agents = AGENTS.read_text(encoding="utf-8")
        doc = DOC.read_text(encoding="utf-8")
        docs_index = DOCS_INDEX.read_text(encoding="utf-8")

        self.assertIn('"test:examples": "python3 test/run-examples.py"', package_json)
        self.assertIn('"examples:compile": "python3 test/run-examples.py --compile-only"', package_json)
        self.assertIn("Keep examples on full runs by default.", run_ci)
        self.assertIn("Run stable CI surface", ci_harness)
        self.assertIn("Examples are QA contracts", agents)
        self.assertIn("npm run test:examples", agents)
        self.assertIn("examples-qa-contract.md", docs_index)
        self.assertIn("What the harness checks", doc)
        self.assertIn("Every example profile lane must compile", doc)
        self.assertIn("expected/*.stdout", doc)
        self.assertIn("Modules that own typed Go APIs declare `@:goNative`", doc)


if __name__ == "__main__":
    unittest.main()
