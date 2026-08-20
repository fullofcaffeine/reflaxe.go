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
    return env


def write_project(module_root: Path, *, with_go_sum: bool = True) -> tuple[bytes, bytes | None]:
    go_mod = b"// caller comment\n\nmodule   example.com/caller/project\n\ngo 1.24.0\n"
    go_sum = b"example.com/dependency v1.0.0 h1:caller-owned-bytes\n" if with_go_sum else None
    (module_root / "go.mod").write_bytes(go_mod)
    if go_sum is not None:
        (module_root / "go.sum").write_bytes(go_sum)
    (module_root / "Main.hx").write_text(
        'class Main { static function main():Void { trace("existing module"); } }\n',
        encoding="utf-8",
    )
    manifest = {
        "schemaVersion": 1,
        "mode": "existing-module",
        "moduleRoot": ".",
        "packageDir": ".",
        "packageName": "main",
        "runtimeDir": "hxrt",
        "entrypoint": {"kind": "compiler-main"},
        "build": {"kind": "none"},
    }
    (module_root / "reflaxe-go-project.json").write_text(
        json.dumps(manifest, indent=2) + "\n", encoding="utf-8"
    )
    return go_mod, go_sum


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
            f"go_output={module_root}",
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
    case: unittest.TestCase,
    module_root: Path,
    go_mod: bytes,
    go_sum: bytes | None,
) -> None:
    case.assertEqual(go_mod, (module_root / "go.mod").read_bytes())
    if go_sum is None:
        case.assertFalse((module_root / "go.sum").exists())
    else:
        case.assertEqual(go_sum, (module_root / "go.sum").read_bytes())


@unittest.skipUnless(shutil.which("haxe"), "requires Haxe")
class ExistingModulePreservationTest(unittest.TestCase):
    def test_success_preserves_present_module_files_byte_for_byte(self) -> None:
        with tempfile.TemporaryDirectory(prefix="haxe-go-existing-module-") as raw:
            module_root = Path(raw)
            go_mod, go_sum = write_project(module_root)

            completed = run_compiler(module_root)
            output = completed.stdout + completed.stderr

            self.assertEqual(0, completed.returncode, output)
            assert_module_files_unchanged(self, module_root, go_mod, go_sum)
            go_test = subprocess.run(
                ["go", "test", "./..."],
                cwd=module_root,
                capture_output=True,
                text=True,
                timeout=180,
            )
            self.assertEqual(
                0, go_test.returncode, go_test.stdout + go_test.stderr
            )
            assert_module_files_unchanged(self, module_root, go_mod, go_sum)
            self.assertTrue((module_root / "main.go").is_file())
            self.assertTrue((module_root / "hxrt").is_dir())
            metadata = json.loads(
                (module_root / "_GeneratedFiles.json").read_text(encoding="utf-8")
            )
            self.assertNotIn("go.mod", metadata["filesGenerated"])
            self.assertNotIn("go.sum", metadata["filesGenerated"])

    def test_success_keeps_an_absent_go_sum_absent(self) -> None:
        with tempfile.TemporaryDirectory(prefix="haxe-go-existing-module-") as raw:
            module_root = Path(raw)
            go_mod, go_sum = write_project(module_root, with_go_sum=False)

            completed = run_compiler(module_root)
            output = completed.stdout + completed.stderr

            self.assertEqual(0, completed.returncode, output)
            assert_module_files_unchanged(self, module_root, go_mod, go_sum)

    def test_reserved_module_file_alias_in_old_metadata_fails_before_generation(self) -> None:
        with tempfile.TemporaryDirectory(prefix="haxe-go-existing-module-") as raw:
            module_root = Path(raw)
            go_mod, go_sum = write_project(module_root)
            metadata = {
                "version": 1,
                "id": 1,
                "wasCached": False,
                "filesGenerated": ["GO.MOD"],
            }
            (module_root / "_GeneratedFiles.json").write_text(
                json.dumps(metadata), encoding="utf-8"
            )

            completed = run_compiler(module_root)
            output = completed.stdout + completed.stderr

            self.assertNotEqual(0, completed.returncode, output)
            self.assertIn("GO-EXISTING-MODULE-MUTATION", output)
            assert_module_files_unchanged(self, module_root, go_mod, go_sum)
            self.assertFalse((module_root / "main.go").exists())
            self.assertNotIn(str(module_root), output)

    def test_unknown_manifest_field_fails_before_generation(self) -> None:
        with tempfile.TemporaryDirectory(prefix="haxe-go-existing-module-") as raw:
            module_root = Path(raw)
            go_mod, go_sum = write_project(module_root)
            manifest_path = module_root / "reflaxe-go-project.json"
            manifest = json.loads(manifest_path.read_text(encoding="utf-8"))
            manifest["unexpected"] = True
            manifest_path.write_text(json.dumps(manifest), encoding="utf-8")

            completed = run_compiler(module_root)
            output = completed.stdout + completed.stderr

            self.assertNotEqual(0, completed.returncode, output)
            self.assertIn("GO-EXISTING-MODULE-MANIFEST", output)
            assert_module_files_unchanged(self, module_root, go_mod, go_sum)
            self.assertFalse((module_root / "main.go").exists())
            self.assertNotIn(str(module_root), output)

    def test_output_failure_still_preserves_module_files(self) -> None:
        with tempfile.TemporaryDirectory(prefix="haxe-go-existing-module-") as raw:
            module_root = Path(raw)
            outside = module_root / "outside"
            outside.mkdir()
            go_mod, go_sum = write_project(module_root)
            (module_root / "hxrt").symlink_to(outside, target_is_directory=True)

            completed = run_compiler(module_root)
            output = completed.stdout + completed.stderr

            self.assertNotEqual(0, completed.returncode, output)
            self.assertIn("Refused generated output", output)
            assert_module_files_unchanged(self, module_root, go_mod, go_sum)
            self.assertEqual([], list(outside.iterdir()))
            self.assertNotIn(str(module_root), output)


if __name__ == "__main__":
    unittest.main()
