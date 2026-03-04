#!/usr/bin/env python3

from __future__ import annotations

import json
import subprocess
import sys
import tempfile
import textwrap
import unittest
from pathlib import Path


REPO_ROOT = Path(__file__).resolve().parent.parent
SCRIPT = REPO_ROOT / "test" / "run-metal-example-boundary.py"


def write_file(path: Path, content: str) -> None:
    path.parent.mkdir(parents=True, exist_ok=True)
    path.write_text(textwrap.dedent(content).strip() + "\n", encoding="utf-8")


class MetalExampleBoundaryTest(unittest.TestCase):
    def run_checker(self, repo_root: Path, *args: str) -> subprocess.CompletedProcess[str]:
        cmd = [sys.executable, str(SCRIPT), "--root", str(repo_root), *args]
        return subprocess.run(cmd, capture_output=True, text=True, check=False)

    def test_boundary_enforce_fails_for_boundary_module_import(self) -> None:
        with tempfile.TemporaryDirectory() as tmp:
            root = Path(tmp)
            write_file(
                root / "examples" / "demo" / "profile" / "MetalRuntime.hx",
                """
                package profile;
                import haxe.ds.List;
                class MetalRuntime {}
                """,
            )

            result = self.run_checker(root)
            self.assertEqual(result.returncode, 1, result.stdout + result.stderr)
            self.assertIn("Metal example boundary check failed", result.stdout)

    def test_boundary_scope_ignores_non_boundary_modules(self) -> None:
        with tempfile.TemporaryDirectory() as tmp:
            root = Path(tmp)
            write_file(
                root / "examples" / "demo" / "app" / "core" / "Core.hx",
                """
                package app.core;
                import haxe.ds.StringMap;
                class Core {}
                """,
            )

            result = self.run_checker(root)
            self.assertEqual(result.returncode, 0, result.stdout + result.stderr)
            self.assertIn("passed", result.stdout)

    def test_full_scope_audit_reports_violations_without_failing(self) -> None:
        with tempfile.TemporaryDirectory() as tmp:
            root = Path(tmp)
            report = root / "audit.json"
            write_file(
                root / "examples" / "zeta" / "app" / "core" / "Core.hx",
                """
                package app.core;
                import haxe.ds.StringMap;
                class Core {}
                """,
            )
            write_file(
                root / "examples" / "alpha" / "profile" / "MetalRuntime.hx",
                """
                package profile;
                import haxe.ds.IntMap;
                class MetalRuntime {}
                """,
            )

            result = self.run_checker(root, "--scope", "full", "--mode", "audit", "--report", str(report))
            self.assertEqual(result.returncode, 0, result.stdout + result.stderr)
            self.assertIn("Metal example boundary audit report", result.stdout)
            self.assertTrue(report.exists(), "expected audit report to be written")

            data = json.loads(report.read_text(encoding="utf-8"))
            self.assertEqual(data["scope"], "full")
            self.assertEqual(data["mode"], "audit")
            self.assertEqual(len(data["violations"]), 2)
            self.assertEqual(
                [entry["path"] for entry in data["violations"]],
                [
                    "examples/alpha/profile/MetalRuntime.hx",
                    "examples/zeta/app/core/Core.hx",
                ],
            )

    def test_audit_threshold_can_fail_when_enabled(self) -> None:
        with tempfile.TemporaryDirectory() as tmp:
            root = Path(tmp)
            write_file(
                root / "examples" / "demo" / "profile" / "MetalRuntime.hx",
                """
                package profile;
                import haxe.ds.List;
                class MetalRuntime {}
                """,
            )

            result = self.run_checker(root, "--mode", "audit", "--max-violations", "0")
            self.assertEqual(result.returncode, 1, result.stdout + result.stderr)
            self.assertIn("violation threshold exceeded", result.stdout)


if __name__ == "__main__":
    unittest.main()
