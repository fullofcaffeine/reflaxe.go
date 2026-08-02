#!/usr/bin/env python3

from __future__ import annotations

import importlib.util
import json
from pathlib import Path
import subprocess
import sys
import unittest


ROOT = Path(__file__).resolve().parent.parent
PLANNER = ROOT / "test" / "run-test-plan.py"
STRATEGY = ROOT / "test" / "testing-strategy.json"


def load_planner():
    spec = importlib.util.spec_from_file_location("run_test_plan_contract", PLANNER)
    if spec is None or spec.loader is None:
        raise RuntimeError("failed to load run-test-plan.py")
    module = importlib.util.module_from_spec(spec)
    sys.modules[spec.name] = module
    spec.loader.exec_module(module)
    return module


class AffectedTestPlanTest(unittest.TestCase):
    def setUp(self) -> None:
        self.strategy = json.loads(STRATEGY.read_text(encoding="utf-8"))

    def test_runtime_change_selects_runtime_and_independent_tracer_surfaces(self) -> None:
        module = load_planner()
        plan = module.build_plan(self.strategy, {"runtime/hxrt/socket.go"})
        self.assertFalse(plan["fullExpansion"])
        self.assertIn("runtime-stdlib", plan["selectedSurfaces"])
        self.assertIn("runtime-stdlib", plan["selectedOwners"])
        self.assertIn("portable-tracer-smoke", plan["selectedOwners"])

    def test_watch_tool_and_scaffold_select_the_managed_dev_contract(self) -> None:
        module = load_planner()
        for changed in (
            {"scripts/dev/haxe_go_watch.py"},
            {"templates/basic/scripts/dev/watch.py"},
            {"templates/basic/package.json"},
        ):
            plan = module.build_plan(self.strategy, changed)
            self.assertIn("diagnostics-tooling", plan["selectedOwners"])
            self.assertIn("npm run test:dev-watch", plan["commands"])

    def test_unknown_change_expands_to_every_surface_and_owner(self) -> None:
        module = load_planner()
        plan = module.build_plan(self.strategy, {"mystery/new-format.xyz"})
        self.assertTrue(plan["fullExpansion"])
        self.assertEqual(
            {item["id"] for item in self.strategy["productSurfaces"]},
            set(plan["selectedSurfaces"]),
        )
        self.assertEqual(
            {item["id"] for item in self.strategy["testOwners"]},
            set(plan["selectedOwners"]),
        )

    def test_explain_json_names_paths_owners_surfaces_and_commands(self) -> None:
        proc = subprocess.run(
            [
                sys.executable,
                str(PLANNER),
                "--json",
                "--changed-file",
                "examples/interop_smoke/Main.hx",
            ],
            cwd=ROOT,
            capture_output=True,
            text=True,
            check=True,
        )
        plan = json.loads(proc.stdout)
        self.assertEqual(["examples/interop_smoke/Main.hx"], plan["changedFiles"])
        self.assertIn("package-examples", plan["selectedOwners"])
        self.assertIn("package-downstream-examples", plan["selectedSurfaces"])
        self.assertTrue(plan["commands"])
        self.assertTrue(plan["reasons"])

    def test_git_discovery_failure_conservatively_expands_the_plan(self) -> None:
        proc = subprocess.run(
            [
                sys.executable,
                str(PLANNER),
                "--json",
                "--base",
                "definitely-not-a-valid-git-revision",
            ],
            cwd=ROOT,
            capture_output=True,
            text=True,
            check=True,
        )
        plan = json.loads(proc.stdout)
        self.assertTrue(plan["fullExpansion"])
        self.assertIn("Git change discovery failed", " ".join(plan["fullExpansionReasons"]))
        self.assertEqual(
            {item["id"] for item in self.strategy["testOwners"]},
            set(plan["selectedOwners"]),
        )


if __name__ == "__main__":
    unittest.main()
