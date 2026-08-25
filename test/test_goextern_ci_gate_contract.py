#!/usr/bin/env python3

from __future__ import annotations

import importlib.util
import sys
import unittest
from pathlib import Path


REPO_ROOT = Path(__file__).resolve().parent.parent
RUN_CI = REPO_ROOT / "test" / "run-ci.py"


def load_run_ci():
    spec = importlib.util.spec_from_file_location("run_ci", RUN_CI)
    module = importlib.util.module_from_spec(spec)
    assert spec.loader is not None
    sys.modules["run_ci"] = module
    spec.loader.exec_module(module)
    return module


class GoexternCiGateContractTest(unittest.TestCase):
    @classmethod
    def setUpClass(cls) -> None:
        cls.run_ci = load_run_ci()

    def test_goextern_confidence_stage_includes_the_runtime_tracer(self) -> None:
        self.assertEqual(
            self.run_ci.build_goextern_command(),
            ["npm", "run", "test:goextern"],
        )

    def test_mismatched_toolchain_runs_current_toolchain_smoke(self) -> None:
        decision = self.run_ci.resolve_goextern_fixture_stage(
            skip=False,
            current_release="1.25",
            target_release="1.23",
        )

        self.assertEqual(decision.kind, "smoke")
        self.assertIn("current toolchain smoke", decision.reason)
        self.assertEqual(decision.command, ["python3", "test/run-goextern-fixtures.py", "--smoke"])

    def test_matching_toolchain_runs_pinned_fixture_drift_check(self) -> None:
        decision = self.run_ci.resolve_goextern_fixture_stage(
            skip=False,
            current_release="1.23",
            target_release="1.23",
        )

        self.assertEqual(decision.kind, "fixtures")
        self.assertEqual(decision.command, ["python3", "test/run-goextern-fixtures.py"])


if __name__ == "__main__":
    unittest.main()
