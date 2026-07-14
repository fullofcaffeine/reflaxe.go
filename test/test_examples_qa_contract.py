#!/usr/bin/env python3

from __future__ import annotations

import importlib.util
from pathlib import Path
import sys
import unittest


REPO_ROOT = Path(__file__).resolve().parent.parent
RUN_EXAMPLES = REPO_ROOT / "test" / "run-examples.py"
EXAMPLES_ROOT = REPO_ROOT / "examples"
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


if __name__ == "__main__":
    unittest.main()
