#!/usr/bin/env python3

from __future__ import annotations

import json
import os
from pathlib import Path
import shutil
import subprocess
import tempfile
import unittest


ROOT = Path(__file__).resolve().parent.parent
FIXTURE = ROOT / "test" / "fixtures" / "generated_output_confinement"
COMPILER_FIXTURE = ROOT / "test" / "snapshot" / "core" / "hello_trace"


def haxe_env() -> dict[str, str]:
    env = os.environ.copy()
    env["HAXE_NO_SERVER"] = "1"
    return env


def run_compiler(output_root: Path) -> subprocess.CompletedProcess[str]:
    return subprocess.run(
        [
            shutil.which("haxe") or "haxe",
            "-cp",
            str(COMPILER_FIXTURE),
            "-lib",
            "reflaxe.go",
            "-D",
            f"go_output={output_root}",
            "-D",
            "go_no_build",
            "-D",
            "reflaxe.dont_output_metadata_id",
            "-D",
            "no-traces",
            "-D",
            "no_traces",
            "-main",
            "Main",
        ],
        cwd=ROOT,
        env=haxe_env(),
        capture_output=True,
        text=True,
        timeout=60,
    )


@unittest.skipUnless(shutil.which("haxe"), "requires Haxe")
class GeneratedOutputConfinementTest(unittest.TestCase):
    def test_typed_boundary_rejects_cross_host_and_symlink_escapes(self) -> None:
        with tempfile.TemporaryDirectory(prefix="haxe-go-output-boundary-") as raw:
            temp = Path(raw)
            output_root = temp / "generated"
            outside = temp / "outside"
            output_root.mkdir()
            outside.mkdir()
            (output_root / "inside").mkdir()
            (outside / "existing.go").write_text("sentinel\n", encoding="utf-8")
            (output_root / "escape-dir").symlink_to(outside, target_is_directory=True)
            (output_root / "escape-file.go").symlink_to(outside / "existing.go")
            (output_root / "broken-file.go").symlink_to(outside / "new.go")
            (output_root / "inside-link").symlink_to(
                output_root / "inside", target_is_directory=True
            )

            env = haxe_env()
            env["GO_OUTPUT_CONFINEMENT_ROOT"] = str(output_root)
            env["GO_OUTPUT_CONFINEMENT_OUTSIDE"] = str(outside)
            proc = subprocess.run(
                [
                    shutil.which("haxe") or "haxe",
                    "-cp",
                    "src",
                    "-cp",
                    "vendor/reflaxe/src",
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

            output = proc.stdout + proc.stderr
            self.assertEqual(0, proc.returncode, output)
            self.assertEqual("generated-output confinement: OK\n", proc.stdout)
            self.assertEqual("sentinel\n", (outside / "existing.go").read_text(encoding="utf-8"))
            self.assertFalse((outside / "new.go").exists())
            self.assertFalse((outside / "escaped.go").exists())
            self.assertNotIn(str(temp), output)

    def test_compiler_rejects_runtime_copy_through_external_symlink(self) -> None:
        with tempfile.TemporaryDirectory(prefix="haxe-go-output-runtime-") as raw:
            temp = Path(raw)
            output_root = temp / "generated"
            outside = temp / "outside"
            output_root.mkdir()
            outside.mkdir()
            (output_root / "hxrt").symlink_to(outside, target_is_directory=True)

            proc = run_compiler(output_root)
            output = proc.stdout + proc.stderr

            self.assertNotEqual(0, proc.returncode, output)
            self.assertIn("Refused generated output", output)
            self.assertEqual([], list(outside.iterdir()))
            self.assertNotIn(str(temp), output)

    def test_compiler_rejects_a_symlink_as_the_configured_output_root(self) -> None:
        with tempfile.TemporaryDirectory(prefix="haxe-go-output-root-link-") as raw:
            temp = Path(raw)
            outside = temp / "outside"
            output_root = temp / "generated"
            outside.mkdir()
            output_root.symlink_to(outside, target_is_directory=True)

            proc = run_compiler(output_root)
            output = proc.stdout + proc.stderr

            self.assertNotEqual(0, proc.returncode, output)
            self.assertIn("Refused generated output", output)
            self.assertEqual([], list(outside.iterdir()))
            self.assertNotIn(str(temp), output)

    def test_compiler_rejects_poisoned_managed_file_metadata_before_deletion(self) -> None:
        with tempfile.TemporaryDirectory(prefix="haxe-go-output-metadata-") as raw:
            temp = Path(raw)
            output_root = temp / "generated"
            outside = temp / "outside"
            output_root.mkdir()
            outside.mkdir()
            sentinel = outside / "stale.go"
            sentinel.write_text("must survive\n", encoding="utf-8")
            metadata = {
                "version": 1,
                "id": 1,
                "wasCached": False,
                "filesGenerated": ["../outside/stale.go"],
            }
            (output_root / "_GeneratedFiles.json").write_text(
                json.dumps(metadata), encoding="utf-8"
            )

            proc = run_compiler(output_root)
            output = proc.stdout + proc.stderr

            self.assertNotEqual(0, proc.returncode, output)
            self.assertIn("Refused generated output", output)
            self.assertEqual("must survive\n", sentinel.read_text(encoding="utf-8"))
            self.assertFalse((output_root / "main.go").exists())
            self.assertNotIn(str(temp), output)

    def test_all_go_owned_writers_route_through_the_boundary(self) -> None:
        compiler = (ROOT / "src" / "reflaxe" / "go" / "GoReflaxeCompiler.hx").read_text(
            encoding="utf-8"
        )
        iterator = (ROOT / "src" / "reflaxe" / "go" / "GoOutputIterator.hx").read_text(
            encoding="utf-8"
        )
        boundary = (
            ROOT
            / "src"
            / "reflaxe"
            / "go"
            / "compiler"
            / "GoGeneratedOutputBoundary.hx"
        )

        self.assertTrue(boundary.is_file())
        self.assertIn("GoGeneratedOutputBoundary", compiler)
        self.assertIn("validateManagedFileMetadata", compiler)
        self.assertIn("new GoExistingModuleOutputPlan", compiler)
        self.assertIn("new GoExistingModuleOutputTransaction", compiler)
        self.assertNotIn("outputManager.saveFile(", compiler)
        self.assertNotIn("output.saveFile(", compiler)
        self.assertIn("GoGeneratedOutputBoundary", iterator)
        self.assertNotIn("File.saveContent(", iterator)
        self.assertNotIn("File.copy(", iterator)
        inspector = (
            ROOT
            / "src"
            / "reflaxe"
            / "go"
            / "compiler"
            / "GoPackageDirectoryInspector.hx"
        ).read_text(encoding="utf-8")
        self.assertIn("GoExistingModuleOwnership", inspector)
        self.assertNotIn("_GeneratedFiles.json", inspector)


if __name__ == "__main__":
    unittest.main()
