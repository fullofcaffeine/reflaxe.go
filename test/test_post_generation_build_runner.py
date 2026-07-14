#!/usr/bin/env python3

from __future__ import annotations

import os
import shutil
import subprocess
import tempfile
import unittest
from pathlib import Path


ROOT = Path(__file__).resolve().parent.parent
FIXTURE = ROOT / "test" / "fixtures" / "post_generation_build_runner"
FAILURE_CASE = ROOT / "test" / "snapshot" / "negative" / "post_generation_go_build_failure"


class PostGenerationBuildRunnerTest(unittest.TestCase):
    def run_failure_case(self, *extra_args: str) -> subprocess.CompletedProcess[str]:
        env = os.environ.copy()
        env["HAXE_NO_SERVER"] = "1"
        try:
            return subprocess.run(
                ["haxe", "compile.hxml", *extra_args],
                cwd=FAILURE_CASE,
                env=env,
                capture_output=True,
                text=True,
                timeout=30,
            )
        finally:
            shutil.rmtree(FAILURE_CASE / "out", ignore_errors=True)

    def test_nonzero_backend_build_is_fatal_without_local_path_leaks(self) -> None:
        proc = self.run_failure_case()
        output = proc.stdout + proc.stderr

        self.assertNotEqual(0, proc.returncode, output)
        self.assertIn("fixture Go build stderr: exit 7", output)
        diagnostic = (
            "Post-generation Go build failed: "
            "`fail-go-command.sh build .` exited with status 7."
        )
        self.assertEqual(1, output.count(diagnostic), output)
        self.assertNotIn(str(ROOT), output)
        self.assertNotIn(str(FAILURE_CASE), output)

    def test_codegen_only_opt_outs_remain_explicit(self) -> None:
        for define in ("go_no_build", "go_codegen_only"):
            with self.subTest(define=define):
                proc = self.run_failure_case("-D", define)

                self.assertEqual(0, proc.returncode, proc.stdout + proc.stderr)
                self.assertNotIn("fixture Go build stderr", proc.stdout + proc.stderr)

    def test_successful_backend_build_remains_green(self) -> None:
        proc = self.run_failure_case("-D", "go_cmd=go")

        self.assertEqual(0, proc.returncode, proc.stdout + proc.stderr)
        self.assertNotIn("Post-generation Go build failed", proc.stdout + proc.stderr)

    def test_structured_failures_restore_cwd_and_redact_local_paths(self) -> None:
        with tempfile.TemporaryDirectory(prefix="haxe-go-post-build-") as workdir:
            env = os.environ.copy()
            env["HAXE_NO_SERVER"] = "1"
            env["GO_POST_BUILD_TEST_WORKDIR"] = workdir
            proc = subprocess.run(
                [
                    "haxe",
                    "-cp",
                    "src",
                    "-cp",
                    str(FIXTURE),
                    "--interp",
                    "-main",
                    "Main",
                ],
                cwd=ROOT,
                env=env,
                capture_output=True,
                text=True,
                timeout=30,
            )

        self.assertEqual(0, proc.returncode, proc.stdout + proc.stderr)
        self.assertEqual("post-generation build runner: OK\n", proc.stdout)
        self.assertNotIn(str(ROOT), proc.stdout + proc.stderr)
        self.assertNotIn(workdir, proc.stdout + proc.stderr)


if __name__ == "__main__":
    unittest.main()
