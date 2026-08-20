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
MODULE_PATH = "example.com/caller/project"


def haxe_env() -> dict[str, str]:
    env = os.environ.copy()
    env["HAXE_NO_SERVER"] = "1"
    return env


def write_module(module_root: Path) -> tuple[bytes, bytes]:
    go_mod = f"// caller comment\n\nmodule {MODULE_PATH}\n\ngo 1.24.0\n".encode()
    go_sum = b"example.com/dependency v1.0.0 h1:caller-owned-bytes\n"
    (module_root / "go.mod").write_bytes(go_mod)
    (module_root / "go.sum").write_bytes(go_sum)
    (module_root / "Main.hx").write_text(
        'class Main { static function main():Void { Sys.println("bridge ran"); } }\n',
        encoding="utf-8",
    )
    return go_mod, go_sum


def write_manifest(
    module_root: Path,
    *,
    package_dir: str,
    package_name: str,
    runtime_dir: str,
    entrypoint: dict[str, str],
) -> Path:
    manifest = {
        "schemaVersion": 1,
        "mode": "existing-module",
        "moduleRoot": ".",
        "packageDir": package_dir,
        "packageName": package_name,
        "runtimeDir": runtime_dir,
        "entrypoint": entrypoint,
        "build": {"kind": "none"},
    }
    path = module_root / "reflaxe-go-project.json"
    path.write_text(json.dumps(manifest, indent=2) + "\n", encoding="utf-8")
    return path


def run_compiler(module_root: Path, package_dir: str) -> subprocess.CompletedProcess[str]:
    output_dir = module_root / package_dir
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
            f"go_output={output_dir}",
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
        timeout=180,
    )


def assert_module_files_unchanged(
    case: unittest.TestCase, module_root: Path, go_mod: bytes, go_sum: bytes
) -> None:
    case.assertEqual(go_mod, (module_root / "go.mod").read_bytes())
    case.assertEqual(go_sum, (module_root / "go.sum").read_bytes())


def generated_sources(package_dir: Path) -> list[Path]:
    return sorted(package_dir.glob("haxego_generated_*.go"))


@unittest.skipUnless(shutil.which("haxe") and shutil.which("go"), "requires Haxe and Go")
class ExistingModulePackageOutputTest(unittest.TestCase):
    def test_caller_bridge_joins_nested_main_package_and_runs(self) -> None:
        with tempfile.TemporaryDirectory(prefix="haxe-go-package-output-") as raw:
            module_root = Path(raw)
            go_mod, go_sum = write_module(module_root)
            package_dir = module_root / "cmd" / "tool"
            package_dir.mkdir(parents=True)
            caller_main = b"package main\n\nfunc main() { RunHaxeMain() }\n"
            caller_helper = b"package main\n\nconst callerOwned = true\n"
            (package_dir / "main.go").write_bytes(caller_main)
            (package_dir / "helper.go").write_bytes(caller_helper)
            write_manifest(
                module_root,
                package_dir="cmd/tool",
                package_name="main",
                runtime_dir="internal/haxe_hxrt",
                entrypoint={"kind": "caller-bridge", "symbol": "RunHaxeMain"},
            )

            completed = run_compiler(module_root, "cmd/tool")
            output = completed.stdout + completed.stderr

            self.assertEqual(0, completed.returncode, output)
            assert_module_files_unchanged(self, module_root, go_mod, go_sum)
            self.assertEqual(caller_main, (package_dir / "main.go").read_bytes())
            self.assertEqual(caller_helper, (package_dir / "helper.go").read_bytes())
            sources = generated_sources(package_dir)
            self.assertTrue(sources)
            generated_text = "\n".join(path.read_text(encoding="utf-8") for path in sources)
            self.assertTrue(all(path.read_text().startswith("package main\n") for path in sources))
            self.assertIn("func RunHaxeMain()", generated_text)
            self.assertNotIn("func main()", generated_text)
            self.assertIn(f'"{MODULE_PATH}/internal/haxe_hxrt"', generated_text)
            self.assertTrue((module_root / "internal" / "haxe_hxrt").is_dir())

            repeated = run_compiler(module_root, "cmd/tool")
            self.assertEqual(0, repeated.returncode, repeated.stdout + repeated.stderr)
            self.assertEqual(caller_main, (package_dir / "main.go").read_bytes())
            self.assertEqual(caller_helper, (package_dir / "helper.go").read_bytes())

            run = subprocess.run(
                ["go", "run", "./cmd/tool"],
                cwd=module_root,
                capture_output=True,
                text=True,
                timeout=180,
            )
            self.assertEqual(0, run.returncode, run.stdout + run.stderr)
            self.assertEqual("bridge ran\n", run.stdout)
            assert_module_files_unchanged(self, module_root, go_mod, go_sum)

    def test_caller_bridge_joins_non_main_package(self) -> None:
        with tempfile.TemporaryDirectory(prefix="haxe-go-package-output-") as raw:
            module_root = Path(raw)
            go_mod, go_sum = write_module(module_root)
            package_dir = module_root / "internal" / "service"
            package_dir.mkdir(parents=True)
            caller_source = b"package service\n\nconst CallerOwned = true\n"
            (package_dir / "service.go").write_bytes(caller_source)
            (package_dir / "service_test.go").write_text(
                "package service_test\n", encoding="utf-8"
            )
            write_manifest(
                module_root,
                package_dir="internal/service",
                package_name="service",
                runtime_dir="internal/haxe_hxrt",
                entrypoint={"kind": "caller-bridge", "symbol": "RunHaxeMain"},
            )

            completed = run_compiler(module_root, "internal/service")
            output = completed.stdout + completed.stderr

            self.assertEqual(0, completed.returncode, output)
            assert_module_files_unchanged(self, module_root, go_mod, go_sum)
            self.assertEqual(caller_source, (package_dir / "service.go").read_bytes())
            sources = generated_sources(package_dir)
            self.assertTrue(sources)
            self.assertTrue(all(path.read_text().startswith("package service\n") for path in sources))
            go_test = subprocess.run(
                ["go", "test", "./..."],
                cwd=module_root,
                capture_output=True,
                text=True,
                timeout=180,
            )
            self.assertEqual(0, go_test.returncode, go_test.stdout + go_test.stderr)

    def test_caller_bridge_keeps_portable_thread_drain_on_renamed_entry(self) -> None:
        with tempfile.TemporaryDirectory(prefix="haxe-go-package-output-") as raw:
            module_root = Path(raw)
            go_mod, go_sum = write_module(module_root)
            (module_root / "Main.hx").write_text(
                """import sys.thread.Lock;
import sys.thread.Thread;

class Main {
    static function main():Void {
        var delay = new Lock();
        Thread.create(function() {
            delay.wait(0.02);
            Sys.println("bridge worker");
        });
    }
}
""",
                encoding="utf-8",
            )
            package_dir = module_root / "cmd" / "tool"
            package_dir.mkdir(parents=True)
            (package_dir / "main.go").write_text(
                "package main\n\nfunc main() { RunHaxeMain() }\n", encoding="utf-8"
            )
            write_manifest(
                module_root,
                package_dir="cmd/tool",
                package_name="main",
                runtime_dir="internal/haxe_hxrt",
                entrypoint={"kind": "caller-bridge", "symbol": "RunHaxeMain"},
            )

            completed = run_compiler(module_root, "cmd/tool")
            self.assertEqual(
                0, completed.returncode, completed.stdout + completed.stderr
            )
            generated_text = "\n".join(
                path.read_text(encoding="utf-8")
                for path in generated_sources(package_dir)
            )
            self.assertIn("func RunHaxeMain()", generated_text)
            self.assertIn("defer hxrt.ThreadWaitForAll()", generated_text)
            run = subprocess.run(
                ["go", "run", "./cmd/tool"],
                cwd=module_root,
                capture_output=True,
                text=True,
                timeout=180,
            )
            self.assertEqual(0, run.returncode, run.stdout + run.stderr)
            self.assertEqual("bridge worker\n", run.stdout)
            assert_module_files_unchanged(self, module_root, go_mod, go_sum)

    def test_caller_bridge_does_not_collide_with_a_main_class_static(self) -> None:
        with tempfile.TemporaryDirectory(prefix="haxe-go-package-output-") as raw:
            module_root = Path(raw)
            go_mod, go_sum = write_module(module_root)
            (module_root / "Main.hx").write_text(
                """class Main {
    static function RunHaxeMain():Void {
        Sys.println("helper ran");
    }

    static function main():Void {
        RunHaxeMain();
    }
}
""",
                encoding="utf-8",
            )
            package_dir = module_root / "cmd" / "tool"
            package_dir.mkdir(parents=True)
            (package_dir / "main.go").write_text(
                "package main\n\nfunc main() { RunHaxeMain() }\n", encoding="utf-8"
            )
            write_manifest(
                module_root,
                package_dir="cmd/tool",
                package_name="main",
                runtime_dir="internal/haxe_hxrt",
                entrypoint={"kind": "caller-bridge", "symbol": "RunHaxeMain"},
            )

            completed = run_compiler(module_root, "cmd/tool")
            self.assertEqual(
                0, completed.returncode, completed.stdout + completed.stderr
            )
            run = subprocess.run(
                ["go", "run", "./cmd/tool"],
                cwd=module_root,
                capture_output=True,
                text=True,
                timeout=180,
            )
            self.assertEqual(0, run.returncode, run.stdout + run.stderr)
            self.assertEqual("helper ran\n", run.stdout)
            assert_module_files_unchanged(self, module_root, go_mod, go_sum)

    def test_invalid_package_identifier_fails_before_writes(self) -> None:
        with tempfile.TemporaryDirectory(prefix="haxe-go-package-output-") as raw:
            module_root = Path(raw)
            go_mod, go_sum = write_module(module_root)
            write_manifest(
                module_root,
                package_dir="internal/service",
                package_name="not-valid",
                runtime_dir="internal/haxe_hxrt",
                entrypoint={"kind": "caller-bridge", "symbol": "RunHaxeMain"},
            )

            completed = run_compiler(module_root, "internal/service")
            output = completed.stdout + completed.stderr

            self.assertNotEqual(0, completed.returncode, output)
            self.assertIn("GO-PACKAGE-NAME", output)
            self.assertFalse((module_root / "internal" / "service").exists())
            assert_module_files_unchanged(self, module_root, go_mod, go_sum)

    def test_package_directory_symlink_fails_before_writes(self) -> None:
        with tempfile.TemporaryDirectory(prefix="haxe-go-package-output-") as raw:
            module_root = Path(raw)
            outside = module_root / "outside"
            outside.mkdir()
            go_mod, go_sum = write_module(module_root)
            (module_root / "linked").symlink_to(outside, target_is_directory=True)
            write_manifest(
                module_root,
                package_dir="linked",
                package_name="service",
                runtime_dir="internal/haxe_hxrt",
                entrypoint={"kind": "caller-bridge", "symbol": "RunHaxeMain"},
            )

            completed = run_compiler(module_root, "linked")
            output = completed.stdout + completed.stderr

            self.assertNotEqual(0, completed.returncode, output)
            self.assertIn("GO-PACKAGE-DIR", output)
            self.assertEqual([], list(outside.iterdir()))
            assert_module_files_unchanged(self, module_root, go_mod, go_sum)
            self.assertNotIn(str(module_root), output)

    def test_package_and_runtime_case_alias_fails_before_writes(self) -> None:
        with tempfile.TemporaryDirectory(prefix="haxe-go-package-output-") as raw:
            module_root = Path(raw)
            go_mod, go_sum = write_module(module_root)
            write_manifest(
                module_root,
                package_dir="internal/service",
                package_name="service",
                runtime_dir="Internal/Service",
                entrypoint={"kind": "caller-bridge", "symbol": "RunHaxeMain"},
            )

            completed = run_compiler(module_root, "internal/service")
            output = completed.stdout + completed.stderr

            self.assertNotEqual(0, completed.returncode, output)
            self.assertIn("GO-PACKAGE-DIR", output)
            self.assertFalse((module_root / "internal").exists())
            self.assertFalse((module_root / "Internal").exists())
            assert_module_files_unchanged(self, module_root, go_mod, go_sum)

    def test_compiler_main_can_own_an_empty_nested_command_package(self) -> None:
        with tempfile.TemporaryDirectory(prefix="haxe-go-package-output-") as raw:
            module_root = Path(raw)
            go_mod, go_sum = write_module(module_root)
            write_manifest(
                module_root,
                package_dir="cmd/generated",
                package_name="main",
                runtime_dir="internal/haxe_hxrt",
                entrypoint={"kind": "compiler-main"},
            )

            completed = run_compiler(module_root, "cmd/generated")
            output = completed.stdout + completed.stderr

            self.assertEqual(0, completed.returncode, output)
            assert_module_files_unchanged(self, module_root, go_mod, go_sum)
            generated_text = "\n".join(
                path.read_text(encoding="utf-8")
                for path in generated_sources(module_root / "cmd" / "generated")
            )
            self.assertIn("func main()", generated_text)
            repeated = run_compiler(module_root, "cmd/generated")
            self.assertEqual(0, repeated.returncode, repeated.stdout + repeated.stderr)
            run = subprocess.run(
                ["go", "run", "./cmd/generated"],
                cwd=module_root,
                capture_output=True,
                text=True,
                timeout=180,
            )
            self.assertEqual(0, run.returncode, run.stdout + run.stderr)
            self.assertEqual("bridge ran\n", run.stdout)

    def test_package_mismatch_fails_before_writes(self) -> None:
        with tempfile.TemporaryDirectory(prefix="haxe-go-package-output-") as raw:
            module_root = Path(raw)
            go_mod, go_sum = write_module(module_root)
            package_dir = module_root / "internal" / "service"
            package_dir.mkdir(parents=True)
            caller_source = b"package different\n\nconst CallerOwned = true\n"
            (package_dir / "service.go").write_bytes(caller_source)
            write_manifest(
                module_root,
                package_dir="internal/service",
                package_name="service",
                runtime_dir="internal/haxe_hxrt",
                entrypoint={"kind": "caller-bridge", "symbol": "RunHaxeMain"},
            )

            completed = run_compiler(module_root, "internal/service")
            output = completed.stdout + completed.stderr

            self.assertNotEqual(0, completed.returncode, output)
            self.assertIn("GO-PACKAGE-MISMATCH", output)
            self.assertEqual(caller_source, (package_dir / "service.go").read_bytes())
            self.assertFalse(generated_sources(package_dir))
            assert_module_files_unchanged(self, module_root, go_mod, go_sum)
            self.assertNotIn(str(module_root), output)

    def test_compiler_main_rejects_a_mixed_owner_directory(self) -> None:
        with tempfile.TemporaryDirectory(prefix="haxe-go-package-output-") as raw:
            module_root = Path(raw)
            go_mod, go_sum = write_module(module_root)
            package_dir = module_root / "cmd" / "generated"
            package_dir.mkdir(parents=True)
            caller_source = b"package main\n\nconst CallerOwned = true\n"
            (package_dir / "helper.go").write_bytes(caller_source)
            write_manifest(
                module_root,
                package_dir="cmd/generated",
                package_name="main",
                runtime_dir="internal/haxe_hxrt",
                entrypoint={"kind": "compiler-main"},
            )

            completed = run_compiler(module_root, "cmd/generated")
            output = completed.stdout + completed.stderr

            self.assertNotEqual(0, completed.returncode, output)
            self.assertIn("GO-ENTRYPOINT-OWNERSHIP", output)
            self.assertEqual(caller_source, (package_dir / "helper.go").read_bytes())
            self.assertFalse(generated_sources(package_dir))
            assert_module_files_unchanged(self, module_root, go_mod, go_sum)

    def test_generated_filename_collision_preserves_caller_file(self) -> None:
        with tempfile.TemporaryDirectory(prefix="haxe-go-package-output-") as raw:
            module_root = Path(raw)
            go_mod, go_sum = write_module(module_root)
            package_dir = module_root / "internal" / "service"
            package_dir.mkdir(parents=True)
            caller_source = b"package service\n\nconst CallerOwned = true\n"
            collision = b"package service\n\nconst NotCompilerOwned = true\n"
            (package_dir / "service.go").write_bytes(caller_source)
            (package_dir / "haxego_generated_main.go").write_bytes(collision)
            write_manifest(
                module_root,
                package_dir="internal/service",
                package_name="service",
                runtime_dir="internal/haxe_hxrt",
                entrypoint={"kind": "caller-bridge", "symbol": "RunHaxeMain"},
            )

            completed = run_compiler(module_root, "internal/service")
            output = completed.stdout + completed.stderr

            self.assertNotEqual(0, completed.returncode, output)
            self.assertIn("GO-GENERATED-FILE-CONFLICT", output)
            self.assertEqual(collision, (package_dir / "haxego_generated_main.go").read_bytes())
            assert_module_files_unchanged(self, module_root, go_mod, go_sum)


if __name__ == "__main__":
    unittest.main()
