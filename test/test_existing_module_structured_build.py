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


def haxe_env() -> dict[str, str]:
    env = os.environ.copy()
    env["HAXE_NO_SERVER"] = "1"
    env["GOFLAGS"] = ""
    return env


def write_project(module_root: Path, build: dict[str, object]) -> None:
    (module_root / "go.mod").write_text(
        "module example.com/caller/project\n\ngo 1.24.0\n", encoding="utf-8"
    )
    (module_root / "Main.hx").write_text(
        'class Main { static function main():Void { Sys.println("bridge ran"); } }\n',
        encoding="utf-8",
    )
    package_dir = module_root / "cmd" / "tool"
    package_dir.mkdir(parents=True)
    (package_dir / "main.go").write_text(
        """package main

import "fmt"

var BuildLabel = "unset"

func main() {
	fmt.Println(BuildLabel)
	RunHaxeMain()
}
""",
        encoding="utf-8",
    )
    manifest = {
        "schemaVersion": 1,
        "mode": "existing-module",
        "moduleRoot": ".",
        "packageDir": "cmd/tool",
        "packageName": "main",
        "runtimeDir": "internal/haxe_hxrt",
        "entrypoint": {"kind": "caller-bridge", "symbol": "RunHaxeMain"},
        "build": build,
    }
    (module_root / "reflaxe-go-project.json").write_text(
        json.dumps(manifest, indent=2) + "\n", encoding="utf-8"
    )


def run_compiler(module_root: Path) -> subprocess.CompletedProcess[str]:
    return subprocess.run(
        [
            shutil.which("haxe") or "haxe",
            "-cp",
            str(module_root),
            "-cp",
            str(ROOT / "src"),
            "-cp",
            str(ROOT / "std"),
            "-cp",
            str(ROOT / "std" / "go" / "_std"),
            "-cp",
            str(ROOT / "vendor" / "reflaxe" / "src"),
            "-D",
            "reflaxe=4.0.0-beta",
            "-D",
            "reflaxe.go=0.0.0-development",
            "--macro",
            'nullSafety("reflaxe.go")',
            "--macro",
            "reflaxe.go.CompilerBootstrap.Start()",
            "--macro",
            "reflaxe.go.CompilerInit.Start()",
            "-D",
            f"go_output={module_root / 'cmd' / 'tool'}",
            "-D",
            f"reflaxe_go_project={module_root / 'reflaxe-go-project.json'}",
            "-D",
            "no-traces",
            "-D",
            "reflaxe.dont_output_metadata_id",
            "-main",
            "Main",
        ],
        cwd=module_root,
        env=haxe_env(),
        capture_output=True,
        text=True,
        timeout=240,
    )


def structured_build(**overrides: object) -> dict[str, object]:
    build: dict[str, object] = {
        "kind": "go-build",
        "packageTarget": "./cmd/tool",
        "output": "dist/tool",
        "tags": ["zeta", "gms_pure_go", "zeta"],
        "ldflags": [
            "-s",
            "-w",
            "-X",
            "main.BuildLabel=release candidate;$(touch should-not-exist)",
        ],
        "trimpath": True,
        "race": True,
        "arguments": ["-buildvcs=false"],
    }
    build.update(overrides)
    return build


@unittest.skipUnless(shutil.which("haxe") and shutil.which("go"), "requires Haxe and Go")
class ExistingModuleStructuredBuildTest(unittest.TestCase):
    def test_build_request_records_exact_arguments_and_runs_without_shell_expansion(self) -> None:
        with tempfile.TemporaryDirectory(prefix="haxe-go-structured-build-") as raw:
            module_root = Path(raw)
            (module_root / "dist").mkdir()
            write_project(module_root, structured_build())
            go_mod = (module_root / "go.mod").read_bytes()

            completed = run_compiler(module_root)
            output = completed.stdout + completed.stderr

            self.assertEqual(0, completed.returncode, output)
            report_path = module_root / "reflaxe_go_build.json"
            first_report = report_path.read_bytes()
            self.assertEqual(
                {
                    "schemaVersion": 1,
                    "contract": "reflaxe.go/build-invocation",
                    "workingDirectory": ".",
                    "command": "go",
                    "arguments": [
                        "build",
                        "-trimpath",
                        "-race",
                        "-tags=gms_pure_go,zeta",
                        "-ldflags=-s -w -X 'main.BuildLabel=release candidate;$(touch should-not-exist)'",
                        "-buildvcs=false",
                        "-mod=readonly",
                        "-o",
                        "dist/tool",
                        "./cmd/tool",
                    ],
                },
                json.loads(first_report),
            )
            self.assertFalse((module_root / "should-not-exist").exists())
            self.assertEqual(go_mod, (module_root / "go.mod").read_bytes())
            self.assertFalse((module_root / "go.sum").exists())

            binary = module_root / "dist" / "tool"
            run = subprocess.run(
                [str(binary)],
                cwd=module_root,
                capture_output=True,
                text=True,
                timeout=60,
            )
            self.assertEqual(0, run.returncode, run.stdout + run.stderr)
            self.assertEqual(
                "release candidate;$(touch should-not-exist)\nbridge ran\n",
                run.stdout,
            )
            self.assertFalse((module_root / "should-not-exist").exists())

            repeated = run_compiler(module_root)
            self.assertEqual(0, repeated.returncode, repeated.stdout + repeated.stderr)
            self.assertEqual(first_report, report_path.read_bytes())
            self.assertEqual(go_mod, (module_root / "go.mod").read_bytes())
            self.assertFalse((module_root / "go.sum").exists())

    def test_unapproved_argument_fails_before_generated_writes(self) -> None:
        with tempfile.TemporaryDirectory(prefix="haxe-go-structured-build-") as raw:
            module_root = Path(raw)
            write_project(
                module_root,
                structured_build(arguments=["-toolexec=sh -c touch escaped"]),
            )

            completed = run_compiler(module_root)
            output = completed.stdout + completed.stderr

            self.assertNotEqual(0, completed.returncode, output)
            self.assertIn("GO-BUILD-ARGUMENT", output)
            self.assertFalse((module_root / "haxego_generated_main.go").exists())
            self.assertFalse((module_root / "reflaxe_go_build.json").exists())
            self.assertFalse((module_root / "escaped").exists())
            self.assertNotIn(str(module_root), output)

    def test_output_escape_fails_before_generated_writes(self) -> None:
        with tempfile.TemporaryDirectory(prefix="haxe-go-structured-build-") as raw:
            module_root = Path(raw)
            write_project(module_root, structured_build(output="../outside/tool"))

            completed = run_compiler(module_root)
            output = completed.stdout + completed.stderr

            self.assertNotEqual(0, completed.returncode, output)
            self.assertIn("GO-BUILD-OUTPUT", output)
            self.assertFalse((module_root / "haxego_generated_main.go").exists())
            self.assertFalse((module_root / "reflaxe_go_build.json").exists())
            self.assertNotIn(str(module_root), output)

    def test_unquotable_linker_argument_fails_before_generated_writes(self) -> None:
        with tempfile.TemporaryDirectory(prefix="haxe-go-structured-build-") as raw:
            module_root = Path(raw)
            write_project(
                module_root,
                structured_build(ldflags=["-X", "main.BuildLabel=both'and\"quotes"]),
            )

            completed = run_compiler(module_root)
            output = completed.stdout + completed.stderr

            self.assertNotEqual(0, completed.returncode, output)
            self.assertIn("GO-BUILD-LDFLAG", output)
            self.assertFalse((module_root / "haxego_generated_main.go").exists())
            self.assertFalse((module_root / "reflaxe_go_build.json").exists())
            self.assertNotIn(str(module_root), output)

    def test_caller_owned_build_report_fails_before_generated_writes(self) -> None:
        with tempfile.TemporaryDirectory(prefix="haxe-go-structured-build-") as raw:
            module_root = Path(raw)
            write_project(module_root, structured_build())
            caller_report = b"caller-owned report\n"
            report_path = module_root / "reflaxe_go_build.json"
            report_path.write_bytes(caller_report)

            completed = run_compiler(module_root)
            output = completed.stdout + completed.stderr

            self.assertNotEqual(0, completed.returncode, output)
            self.assertIn("GO-GENERATED-FILE-CONFLICT", output)
            self.assertEqual(caller_report, report_path.read_bytes())
            self.assertFalse((module_root / "haxego_generated_main.go").exists())
            self.assertNotIn(str(module_root), output)


if __name__ == "__main__":
    unittest.main()
