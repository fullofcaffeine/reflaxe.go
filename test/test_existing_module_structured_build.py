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


def haxe_env(**overrides: str) -> dict[str, str]:
    env = os.environ.copy()
    env["HAXE_NO_SERVER"] = "1"
    env.update(overrides)
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


def run_compiler(
    module_root: Path, **environment: str
) -> subprocess.CompletedProcess[str]:
    environment.setdefault("GOCACHE", str(module_root / "go-cache"))
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
        env=haxe_env(**environment),
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
        "environment": [{"name": "GOCACHE", "source": "inherit"}],
    }
    build.update(overrides)
    return build


@unittest.skipUnless(shutil.which("haxe") and shutil.which("go"), "requires Haxe and Go")
class ExistingModuleStructuredBuildTest(unittest.TestCase):
    def test_build_request_records_exact_arguments_and_runs_without_shell_expansion(self) -> None:
        with tempfile.TemporaryDirectory(prefix="haxe-go-structured-build-") as raw:
            module_root = Path(raw)
            (module_root / "dist").mkdir()
            go_cache = module_root / "go-cache"
            build = structured_build(
                environment=[
                    {"name": "CGO_ENABLED", "source": "literal", "value": "1"},
                    {"name": "GOCACHE", "source": "inherit"},
                    {"name": "GOPROXY", "source": "inherit"},
                ]
            )
            write_project(module_root, build)
            go_mod = (module_root / "go.mod").read_bytes()

            completed = run_compiler(
                module_root,
                GOCACHE=str(go_cache),
                GOPROXY="https://build-user:build-password@example.invalid",
                GOFLAGS="-ambient-flags-must-not-run",
                GOWORK=str(module_root / "ambient-secret-workspace"),
                GOTOOLCHAIN="auto",
                CLOUD_ACCESS_TOKEN="ambient-secret-token",
            )
            output = completed.stdout + completed.stderr

            self.assertEqual(0, completed.returncode, output)
            report_path = module_root / "reflaxe_go_build.json"
            first_report = report_path.read_bytes()
            self.assertEqual(
                {
                    "schemaVersion": 2,
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
                    "environment": [
                        {
                            "name": "CGO_ENABLED",
                            "source": "literal",
                            "value": "1",
                        },
                        {
                            "name": "GOCACHE",
                            "source": "inherit",
                            "value": "<path>",
                        },
                        {
                            "name": "GOENV",
                            "source": "compiler",
                            "value": "off",
                        },
                        {
                            "name": "GOPROXY",
                            "source": "inherit",
                            "value": "<redacted>",
                        },
                        {
                            "name": "GOTOOLCHAIN",
                            "source": "compiler",
                            "value": "local",
                        },
                        {
                            "name": "GOWORK",
                            "source": "compiler",
                            "value": "off",
                        },
                    ],
                },
                json.loads(first_report),
            )
            report_text = first_report.decode("utf-8")
            self.assertNotIn("build-password", report_text)
            self.assertNotIn(str(module_root), report_text)
            self.assertNotIn("CLOUD_ACCESS_TOKEN", report_text)
            self.assertNotIn("ambient-secret-token", report_text)
            self.assertNotIn("GOFLAGS", report_text)
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

            repeated = run_compiler(
                module_root,
                GOCACHE=str(go_cache),
                GOPROXY="https://build-user:build-password@example.invalid",
                GOFLAGS="-ambient-flags-must-not-run",
                GOWORK=str(module_root / "ambient-secret-workspace"),
                GOTOOLCHAIN="auto",
                CLOUD_ACCESS_TOKEN="ambient-secret-token",
            )
            self.assertEqual(0, repeated.returncode, repeated.stdout + repeated.stderr)
            self.assertEqual(first_report, report_path.read_bytes())
            self.assertEqual(go_mod, (module_root / "go.mod").read_bytes())
            self.assertFalse((module_root / "go.sum").exists())

    def test_unknown_and_forbidden_environment_names_fail_before_writes(self) -> None:
        for name in ("CLOUD_ACCESS_TOKEN", "GOFLAGS"):
            with self.subTest(name=name), tempfile.TemporaryDirectory(
                prefix="haxe-go-structured-build-"
            ) as raw:
                module_root = Path(raw)
                write_project(
                    module_root,
                    structured_build(
                        environment=[
                            {"name": name, "source": "literal", "value": "secret"}
                        ]
                    ),
                )

                completed = run_compiler(module_root)
                output = completed.stdout + completed.stderr

                self.assertNotEqual(0, completed.returncode, output)
                self.assertIn("GO-BUILD-ENVIRONMENT", output)
                self.assertIn(name, output)
                self.assertNotIn("secret", output)
                self.assertFalse((module_root / "haxego_generated_main.go").exists())
                self.assertFalse((module_root / "reflaxe_go_build.json").exists())
                self.assertNotIn(str(module_root), output)

    def test_incompatible_cgo_configuration_fails_before_writes(self) -> None:
        with tempfile.TemporaryDirectory(prefix="haxe-go-structured-build-") as raw:
            module_root = Path(raw)
            write_project(
                module_root,
                structured_build(
                    race=False,
                    environment=[
                        {"name": "CGO_ENABLED", "source": "literal", "value": "0"},
                        {"name": "CC", "source": "literal", "value": "clang"},
                    ],
                ),
            )

            completed = run_compiler(module_root)
            output = completed.stdout + completed.stderr

            self.assertNotEqual(0, completed.returncode, output)
            self.assertIn("GO-BUILD-ENVIRONMENT", output)
            self.assertIn("CGO_ENABLED", output)
            self.assertIn("CC", output)
            self.assertFalse((module_root / "haxego_generated_main.go").exists())
            self.assertFalse((module_root / "reflaxe_go_build.json").exists())
            self.assertNotIn(str(module_root), output)

    def test_cgo_disabled_build_uses_the_governed_environment(self) -> None:
        with tempfile.TemporaryDirectory(prefix="haxe-go-structured-build-") as raw:
            module_root = Path(raw)
            (module_root / "dist").mkdir()
            go_cache = module_root / "go-cache"
            write_project(
                module_root,
                structured_build(
                    race=False,
                    environment=[
                        {"name": "CGO_ENABLED", "source": "literal", "value": "0"},
                        {"name": "GOCACHE", "source": "inherit"},
                    ],
                ),
            )

            completed = run_compiler(
                module_root,
                GOCACHE=str(go_cache),
                GOFLAGS="-ambient-flags-must-not-run",
                CLOUD_ACCESS_TOKEN="ambient-secret-token",
            )
            output = completed.stdout + completed.stderr

            self.assertEqual(0, completed.returncode, output)
            self.assertTrue((module_root / "dist" / "tool").is_file())
            report = json.loads((module_root / "reflaxe_go_build.json").read_text())
            self.assertIn(
                {"name": "CGO_ENABLED", "source": "literal", "value": "0"},
                report["environment"],
            )

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
