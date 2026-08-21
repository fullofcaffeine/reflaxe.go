#!/usr/bin/env python3

from __future__ import annotations

import json
import hashlib
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


def sha256(data: bytes) -> str:
    return hashlib.sha256(data).hexdigest()


def write_interrupted_transaction(
    module_root: Path,
    *,
    package_dir: str,
    package_name: str,
    runtime_dir: str,
    live_bytes: bytes,
    declared_new_bytes: bytes,
) -> tuple[Path, bytes]:
    package_root = module_root / package_dir
    ownership_path = package_root / ".reflaxe-go-owned.json"
    old_manifest_bytes = ownership_path.read_bytes()
    old_manifest = json.loads(old_manifest_bytes)
    target_record = next(
        record
        for record in old_manifest["files"]
        if record["path"].startswith(f"{package_dir}/haxego_generated_")
    )
    target_path = module_root / target_record["path"]
    old_target_bytes = target_path.read_bytes()

    workspace = package_root / ".reflaxe-go-transaction"
    workspace.mkdir()
    marker = (
        "{\n"
        '  "schemaVersion": 1,\n'
        f'  "modulePath": {json.dumps(MODULE_PATH)},\n'
        f'  "packageDir": {json.dumps(package_dir)},\n'
        f'  "packageName": {json.dumps(package_name)},\n'
        f'  "runtimeDir": {json.dumps(runtime_dir)}\n'
        "}\n"
    ).encode()
    (workspace / "project.json").write_bytes(marker)
    (workspace / "old-ownership.json").write_bytes(old_manifest_bytes)

    for record in old_manifest["files"]:
        backup = workspace / "backup" / record["path"]
        backup.parent.mkdir(parents=True, exist_ok=True)
        backup.write_bytes((module_root / record["path"]).read_bytes())

    new_manifest = json.loads(old_manifest_bytes)
    for record in new_manifest["files"]:
        if record["path"] == target_record["path"]:
            record["sha256"] = sha256(declared_new_bytes)
            break
    new_manifest_bytes = (json.dumps(new_manifest, indent=2) + "\n").encode()
    (workspace / "new-ownership.json").write_bytes(new_manifest_bytes)
    target_path.write_bytes(live_bytes)
    journal = {
        "schemaVersion": 1,
        "projectSha256": sha256(marker),
        "oldManifestSha256": sha256(old_manifest_bytes),
        "newManifestSha256": sha256(new_manifest_bytes),
    }
    (package_root / ".reflaxe-go-transaction.json").write_text(
        json.dumps(journal), encoding="utf-8"
    )
    return target_path, old_target_bytes


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

            all_generated_go = sorted(module_root.glob("cmd/tool/*.go")) + sorted(
                module_root.glob("internal/haxe_hxrt/*.go")
            )
            formatted = subprocess.run(
                ["gofmt", "-d", *[str(path) for path in all_generated_go]],
                cwd=module_root,
                capture_output=True,
                text=True,
                timeout=60,
            )
            self.assertEqual(0, formatted.returncode, formatted.stderr)
            self.assertEqual("", formatted.stdout, formatted.stdout)

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

    def test_runtime_directory_rejects_caller_owned_content(self) -> None:
        with tempfile.TemporaryDirectory(prefix="haxe-go-package-output-") as raw:
            module_root = Path(raw)
            go_mod, go_sum = write_module(module_root)
            package_dir = module_root / "internal" / "service"
            package_dir.mkdir(parents=True)
            (package_dir / "service.go").write_text(
                "package service\n\nconst CallerOwned = true\n", encoding="utf-8"
            )
            runtime_dir = module_root / "internal" / "haxe_hxrt"
            runtime_dir.mkdir()
            caller_runtime = b"package hxrt\n\nconst CallerOwnedRuntime = true\n"
            (runtime_dir / "caller.go").write_bytes(caller_runtime)
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
            self.assertEqual(caller_runtime, (runtime_dir / "caller.go").read_bytes())
            self.assertFalse(generated_sources(package_dir))
            assert_module_files_unchanged(self, module_root, go_mod, go_sum)
            self.assertNotIn(str(module_root), output)

    def test_rerun_rejects_a_caller_modified_generated_file(self) -> None:
        with tempfile.TemporaryDirectory(prefix="haxe-go-package-output-") as raw:
            module_root = Path(raw)
            go_mod, go_sum = write_module(module_root)
            package_dir = module_root / "internal" / "service"
            package_dir.mkdir(parents=True)
            (package_dir / "service.go").write_text(
                "package service\n\nconst CallerOwned = true\n", encoding="utf-8"
            )
            write_manifest(
                module_root,
                package_dir="internal/service",
                package_name="service",
                runtime_dir="internal/haxe_hxrt",
                entrypoint={"kind": "caller-bridge", "symbol": "RunHaxeMain"},
            )

            first = run_compiler(module_root, "internal/service")
            self.assertEqual(0, first.returncode, first.stdout + first.stderr)
            generated = generated_sources(package_dir)
            self.assertTrue(generated)
            modified = b"package service\n\nconst CallerModified = true\n"
            generated[0].write_bytes(modified)

            repeated = run_compiler(module_root, "internal/service")
            output = repeated.stdout + repeated.stderr

            self.assertNotEqual(0, repeated.returncode, output)
            self.assertIn("GO-GENERATED-FILE-CONFLICT", output)
            self.assertEqual(modified, generated[0].read_bytes())
            assert_module_files_unchanged(self, module_root, go_mod, go_sum)
            self.assertNotIn(str(module_root), output)

    def test_rerun_rejects_a_missing_owned_generated_file(self) -> None:
        with tempfile.TemporaryDirectory(prefix="haxe-go-package-output-") as raw:
            module_root = Path(raw)
            go_mod, go_sum = write_module(module_root)
            package_dir = module_root / "internal" / "service"
            package_dir.mkdir(parents=True)
            caller = b"package service\n\nconst CallerOwned = true\n"
            (package_dir / "service.go").write_bytes(caller)
            write_manifest(
                module_root,
                package_dir="internal/service",
                package_name="service",
                runtime_dir="internal/haxe_hxrt",
                entrypoint={"kind": "caller-bridge", "symbol": "RunHaxeMain"},
            )
            first = run_compiler(module_root, "internal/service")
            self.assertEqual(0, first.returncode, first.stdout + first.stderr)
            missing = generated_sources(package_dir)[0]
            missing.unlink()

            repeated = run_compiler(module_root, "internal/service")
            output = repeated.stdout + repeated.stderr

            self.assertNotEqual(0, repeated.returncode, output)
            self.assertIn("GO-GENERATED-FILE-CONFLICT", output)
            self.assertFalse(missing.exists())
            self.assertEqual(caller, (package_dir / "service.go").read_bytes())
            assert_module_files_unchanged(self, module_root, go_mod, go_sum)
            self.assertNotIn(str(module_root), output)

    def test_ownership_record_rejects_case_aliased_paths(self) -> None:
        with tempfile.TemporaryDirectory(prefix="haxe-go-package-output-") as raw:
            module_root = Path(raw)
            go_mod, go_sum = write_module(module_root)
            package_dir = module_root / "internal" / "service"
            package_dir.mkdir(parents=True)
            caller = b"package service\n\nconst CallerOwned = true\n"
            (package_dir / "service.go").write_bytes(caller)
            write_manifest(
                module_root,
                package_dir="internal/service",
                package_name="service",
                runtime_dir="internal/haxe_hxrt",
                entrypoint={"kind": "caller-bridge", "symbol": "RunHaxeMain"},
            )
            first = run_compiler(module_root, "internal/service")
            self.assertEqual(0, first.returncode, first.stdout + first.stderr)
            ownership_path = package_dir / ".reflaxe-go-owned.json"
            ownership = json.loads(ownership_path.read_text(encoding="utf-8"))
            aliased = dict(ownership["files"][0])
            aliased["path"] = aliased["path"].upper()
            ownership["files"].append(aliased)
            ownership["files"].sort(key=lambda record: record["path"])
            tampered = (json.dumps(ownership, indent=2) + "\n").encode()
            ownership_path.write_bytes(tampered)

            repeated = run_compiler(module_root, "internal/service")
            output = repeated.stdout + repeated.stderr

            self.assertNotEqual(0, repeated.returncode, output)
            self.assertIn("GO-OUTPUT-PATH-004", output)
            self.assertEqual(tampered, ownership_path.read_bytes())
            self.assertEqual(caller, (package_dir / "service.go").read_bytes())
            assert_module_files_unchanged(self, module_root, go_mod, go_sum)
            self.assertNotIn(str(module_root), output)

    def test_case_alias_cannot_grant_ownership_to_an_exact_caller_path(self) -> None:
        with tempfile.TemporaryDirectory(prefix="haxe-go-package-output-") as raw:
            module_root = Path(raw)
            go_mod, go_sum = write_module(module_root)
            package_dir = module_root / "internal" / "service"
            package_dir.mkdir(parents=True)
            caller = b"package service\n\nconst CallerOwned = true\n"
            (package_dir / "service.go").write_bytes(caller)
            write_manifest(
                module_root,
                package_dir="internal/service",
                package_name="service",
                runtime_dir="internal/haxe_hxrt",
                entrypoint={"kind": "caller-bridge", "symbol": "RunHaxeMain"},
            )
            first = run_compiler(module_root, "internal/service")
            self.assertEqual(0, first.returncode, first.stdout + first.stderr)

            generated = generated_sources(package_dir)[0]
            alias = generated.with_name(generated.name[0].upper() + generated.name[1:])
            if alias.exists():
                self.skipTest("filesystem does not distinguish path case")
            old_bytes = generated.read_bytes()
            alias.write_bytes(old_bytes)

            ownership_path = package_dir / ".reflaxe-go-owned.json"
            ownership = json.loads(ownership_path.read_text(encoding="utf-8"))
            generated_relative = generated.relative_to(module_root).as_posix()
            alias_relative = alias.relative_to(module_root).as_posix()
            target_record = next(
                record
                for record in ownership["files"]
                if record["path"] == generated_relative
            )
            target_record["path"] = alias_relative
            ownership["files"].sort(key=lambda record: record["path"])
            ownership_path.write_text(
                json.dumps(ownership, indent=2) + "\n", encoding="utf-8"
            )
            (module_root / "Main.hx").write_text(
                'class Main { static function main():Void { Sys.println("changed"); } }\n',
                encoding="utf-8",
            )

            repeated = run_compiler(module_root, "internal/service")
            output = repeated.stdout + repeated.stderr

            self.assertNotEqual(0, repeated.returncode, output)
            self.assertIn("GO-GENERATED-FILE-CONFLICT", output)
            self.assertEqual(old_bytes, generated.read_bytes())
            self.assertEqual(old_bytes, alias.read_bytes())
            self.assertEqual(caller, (package_dir / "service.go").read_bytes())
            assert_module_files_unchanged(self, module_root, go_mod, go_sum)
            self.assertNotIn(str(module_root), output)

    def test_legacy_inventory_cannot_delete_a_caller_file(self) -> None:
        with tempfile.TemporaryDirectory(prefix="haxe-go-package-output-") as raw:
            module_root = Path(raw)
            go_mod, go_sum = write_module(module_root)
            package_dir = module_root / "internal" / "service"
            package_dir.mkdir(parents=True)
            (package_dir / "service.go").write_text(
                "package service\n\nconst CallerOwned = true\n", encoding="utf-8"
            )
            write_manifest(
                module_root,
                package_dir="internal/service",
                package_name="service",
                runtime_dir="internal/haxe_hxrt",
                entrypoint={"kind": "caller-bridge", "symbol": "RunHaxeMain"},
            )

            first = run_compiler(module_root, "internal/service")
            self.assertEqual(0, first.returncode, first.stdout + first.stderr)
            caller_path = package_dir / "stale_inventory_owner.go"
            caller_bytes = b"package service\n\nconst StaleInventoryOwner = true\n"
            caller_path.write_bytes(caller_bytes)
            metadata_path = module_root / "_GeneratedFiles.json"
            metadata = {
                "version": 1,
                "id": 1,
                "wasCached": False,
                "filesGenerated": [caller_path.relative_to(module_root).as_posix()],
            }
            metadata_path.write_text(json.dumps(metadata), encoding="utf-8")

            repeated = run_compiler(module_root, "internal/service")
            output = repeated.stdout + repeated.stderr

            self.assertEqual(0, repeated.returncode, output)
            self.assertTrue(caller_path.is_file(), output)
            self.assertEqual(caller_bytes, caller_path.read_bytes())
            assert_module_files_unchanged(self, module_root, go_mod, go_sum)

    def test_legacy_inventory_cannot_grant_compiler_main_ownership(self) -> None:
        with tempfile.TemporaryDirectory(prefix="haxe-go-package-output-") as raw:
            module_root = Path(raw)
            go_mod, go_sum = write_module(module_root)
            package_dir = module_root / "cmd" / "generated"
            package_dir.mkdir(parents=True)
            collision_path = package_dir / "haxego_generated_main.go"
            collision = b"package main\n\nconst CallerOwned = true\n"
            collision_path.write_bytes(collision)
            write_manifest(
                module_root,
                package_dir="cmd/generated",
                package_name="main",
                runtime_dir="internal/haxe_hxrt",
                entrypoint={"kind": "compiler-main"},
            )
            metadata = {
                "version": 1,
                "id": 1,
                "wasCached": False,
                "filesGenerated": [
                    collision_path.relative_to(module_root).as_posix()
                ],
            }
            (module_root / "_GeneratedFiles.json").write_text(
                json.dumps(metadata), encoding="utf-8"
            )

            completed = run_compiler(module_root, "cmd/generated")
            output = completed.stdout + completed.stderr

            self.assertNotEqual(0, completed.returncode, output)
            self.assertTrue(
                "GO-ENTRYPOINT-OWNERSHIP" in output
                or "GO-GENERATED-FILE-CONFLICT" in output,
                output,
            )
            self.assertEqual(collision, collision_path.read_bytes())
            assert_module_files_unchanged(self, module_root, go_mod, go_sum)
            self.assertNotIn(str(module_root), output)

    def test_owned_stale_generated_source_is_removed(self) -> None:
        with tempfile.TemporaryDirectory(prefix="haxe-go-package-output-") as raw:
            module_root = Path(raw)
            go_mod, go_sum = write_module(module_root)
            (module_root / "Main.hx").write_text(
                "class Main { static function main():Void { Helper.run(); } }\n",
                encoding="utf-8",
            )
            (module_root / "Helper.hx").write_text(
                'class Helper { public static function run():Void { Sys.println("helper"); } }\n',
                encoding="utf-8",
            )
            package_dir = module_root / "internal" / "service"
            package_dir.mkdir(parents=True)
            (package_dir / "service.go").write_text(
                "package service\n\nconst CallerOwned = true\n", encoding="utf-8"
            )
            write_manifest(
                module_root,
                package_dir="internal/service",
                package_name="service",
                runtime_dir="internal/haxe_hxrt",
                entrypoint={"kind": "caller-bridge", "symbol": "RunHaxeMain"},
            )

            first = run_compiler(module_root, "internal/service")
            self.assertEqual(0, first.returncode, first.stdout + first.stderr)
            first_sources = set(generated_sources(package_dir))
            self.assertGreaterEqual(len(first_sources), 2)

            (module_root / "Main.hx").write_text(
                'class Main { static function main():Void { Sys.println("main"); } }\n',
                encoding="utf-8",
            )
            (module_root / "Helper.hx").unlink()
            repeated = run_compiler(module_root, "internal/service")
            self.assertEqual(0, repeated.returncode, repeated.stdout + repeated.stderr)
            second_sources = set(generated_sources(package_dir))

            stale = first_sources - second_sources
            self.assertTrue(stale)
            self.assertTrue(all(not path.exists() for path in stale))
            self.assertEqual(
                b"package service\n\nconst CallerOwned = true\n",
                (package_dir / "service.go").read_bytes(),
            )
            assert_module_files_unchanged(self, module_root, go_mod, go_sum)

    def test_interrupted_install_rolls_back_before_the_next_generation(self) -> None:
        with tempfile.TemporaryDirectory(prefix="haxe-go-package-output-") as raw:
            module_root = Path(raw)
            go_mod, go_sum = write_module(module_root)
            package_dir = module_root / "internal" / "service"
            package_dir.mkdir(parents=True)
            caller = b"package service\n\nconst CallerOwned = true\n"
            (package_dir / "service.go").write_bytes(caller)
            write_manifest(
                module_root,
                package_dir="internal/service",
                package_name="service",
                runtime_dir="internal/haxe_hxrt",
                entrypoint={"kind": "caller-bridge", "symbol": "RunHaxeMain"},
            )
            first = run_compiler(module_root, "internal/service")
            self.assertEqual(0, first.returncode, first.stdout + first.stderr)

            declared_new = b"package service\n\nconst InterruptedNew = true\n"
            target, old_bytes = write_interrupted_transaction(
                module_root,
                package_dir="internal/service",
                package_name="service",
                runtime_dir="internal/haxe_hxrt",
                live_bytes=declared_new,
                declared_new_bytes=declared_new,
            )

            repeated = run_compiler(module_root, "internal/service")
            output = repeated.stdout + repeated.stderr

            self.assertEqual(0, repeated.returncode, output)
            self.assertEqual(old_bytes, target.read_bytes())
            self.assertEqual(caller, (package_dir / "service.go").read_bytes())
            self.assertFalse((package_dir / ".reflaxe-go-transaction").exists())
            self.assertFalse((package_dir / ".reflaxe-go-transaction.json").exists())
            assert_module_files_unchanged(self, module_root, go_mod, go_sum)
            self.assertNotIn(str(module_root), output)

    def test_committed_install_cleans_before_the_next_generation(self) -> None:
        with tempfile.TemporaryDirectory(prefix="haxe-go-package-output-") as raw:
            module_root = Path(raw)
            go_mod, go_sum = write_module(module_root)
            package_dir = module_root / "internal" / "service"
            package_dir.mkdir(parents=True)
            caller = b"package service\n\nconst CallerOwned = true\n"
            (package_dir / "service.go").write_bytes(caller)
            write_manifest(
                module_root,
                package_dir="internal/service",
                package_name="service",
                runtime_dir="internal/haxe_hxrt",
                entrypoint={"kind": "caller-bridge", "symbol": "RunHaxeMain"},
            )
            first = run_compiler(module_root, "internal/service")
            self.assertEqual(0, first.returncode, first.stdout + first.stderr)

            declared_new = b"package service\n\nconst CommittedNew = true\n"
            target, old_bytes = write_interrupted_transaction(
                module_root,
                package_dir="internal/service",
                package_name="service",
                runtime_dir="internal/haxe_hxrt",
                live_bytes=declared_new,
                declared_new_bytes=declared_new,
            )
            workspace = package_dir / ".reflaxe-go-transaction"
            (package_dir / ".reflaxe-go-owned.json").write_bytes(
                (workspace / "new-ownership.json").read_bytes()
            )

            repeated = run_compiler(module_root, "internal/service")
            output = repeated.stdout + repeated.stderr

            self.assertEqual(0, repeated.returncode, output)
            self.assertEqual(old_bytes, target.read_bytes())
            self.assertEqual(caller, (package_dir / "service.go").read_bytes())
            self.assertFalse(workspace.exists())
            self.assertFalse((package_dir / ".reflaxe-go-transaction.json").exists())
            assert_module_files_unchanged(self, module_root, go_mod, go_sum)
            self.assertNotIn(str(module_root), output)

    def test_interrupted_install_with_unknown_live_bytes_fails_closed(self) -> None:
        with tempfile.TemporaryDirectory(prefix="haxe-go-package-output-") as raw:
            module_root = Path(raw)
            go_mod, go_sum = write_module(module_root)
            package_dir = module_root / "internal" / "service"
            package_dir.mkdir(parents=True)
            caller = b"package service\n\nconst CallerOwned = true\n"
            (package_dir / "service.go").write_bytes(caller)
            write_manifest(
                module_root,
                package_dir="internal/service",
                package_name="service",
                runtime_dir="internal/haxe_hxrt",
                entrypoint={"kind": "caller-bridge", "symbol": "RunHaxeMain"},
            )
            first = run_compiler(module_root, "internal/service")
            self.assertEqual(0, first.returncode, first.stdout + first.stderr)

            unknown = b"package service\n\nconst UnknownThirdState = true\n"
            declared_new = b"package service\n\nconst InterruptedNew = true\n"
            target, _ = write_interrupted_transaction(
                module_root,
                package_dir="internal/service",
                package_name="service",
                runtime_dir="internal/haxe_hxrt",
                live_bytes=unknown,
                declared_new_bytes=declared_new,
            )

            repeated = run_compiler(module_root, "internal/service")
            output = repeated.stdout + repeated.stderr

            self.assertNotEqual(0, repeated.returncode, output)
            self.assertIn("GO-OUTPUT-TRANSACTION", output)
            self.assertEqual(unknown, target.read_bytes())
            self.assertEqual(caller, (package_dir / "service.go").read_bytes())
            self.assertTrue((package_dir / ".reflaxe-go-transaction").exists())
            self.assertTrue((package_dir / ".reflaxe-go-transaction.json").exists())
            assert_module_files_unchanged(self, module_root, go_mod, go_sum)
            self.assertNotIn(str(module_root), output)


if __name__ == "__main__":
    unittest.main()
