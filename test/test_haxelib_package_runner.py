#!/usr/bin/env python3

from __future__ import annotations

import hashlib
import json
import os
from pathlib import Path, PurePosixPath
import shutil
import subprocess
import sys
import tempfile
import unittest
import zipfile


ROOT = Path(__file__).resolve().parent.parent
RUNNER = ROOT / "Run.hx"
PACKAGE_SCRIPT = ROOT / "scripts" / "release" / "package-haxelib.sh"
ZIP_SCRIPT = ROOT / "scripts" / "release" / "deterministic-zip.py"
PACKAGE_MANIFEST = "reflaxe-package-manifest.json"
FIXED_ZIP_TIMESTAMP = (2000, 1, 1, 0, 0, 0)

sys.path.insert(0, str(ROOT / "scripts" / "ci"))

from canonical_stdlib_layout_check import (  # noqa: E402
    audit_package_layout,
    expected_packaged_cross_files,
)


def sha256(path: Path) -> str:
    return hashlib.sha256(path.read_bytes()).hexdigest()


def write_json(path: Path, value: object) -> None:
    path.parent.mkdir(parents=True, exist_ok=True)
    path.write_text(json.dumps(value, indent=2) + "\n", encoding="utf-8")


def write_text(path: Path, content: str) -> None:
    path.parent.mkdir(parents=True, exist_ok=True)
    path.write_text(content, encoding="utf-8")


def run_package_runner(
    source_root: Path, package_root: Path, *, clean: bool = True
) -> subprocess.CompletedProcess[str]:
    command = [
        shutil.which("haxe") or "haxe",
        "--run",
        "Run",
        "build",
        str(package_root),
        "--source-root",
        str(source_root),
    ]
    if clean:
        command.append("--clean")
    return subprocess.run(
        command,
        cwd=ROOT,
        capture_output=True,
        text=True,
        timeout=120,
    )


def write_synthetic_source(root: Path) -> None:
    write_json(
        root / "haxelib.json",
        {
            "name": "reflaxe.go",
            "url": "https://example.invalid/reflaxe.go",
            "license": "GPL-3.0-only",
            "tags": ["haxe", "go"],
            "description": "Synthetic package-runner fixture",
            "version": "0.0.0",
            "releasenote": "Synthetic fixture",
            "contributors": ["fixture"],
            "classPath": "src",
            "dependencies": {},
            "reflaxe": {
                "name": "Go",
                "abbv": "go",
                "stdPaths": ["std", "std/go/_std"],
            },
        },
    )
    write_text(root / "src" / "Main.hx", "class Main {}\n")
    write_text(root / "src" / ".cache" / "scratch.txt", "untracked classpath debris\n")
    write_text(root / "std" / "Support.hx", "class Support {}\n")
    write_text(
        root / "std" / "other" / "_std" / "NestedSupport.hx",
        "package other._std; class NestedSupport {}\n",
    )
    write_text(
        root / "std" / "go" / "_std" / "Override.hx",
        "class Override {}\n",
    )
    write_text(root / "runtime" / "hxrt" / "core.go", "package hxrt\n")
    write_text(
        root / "runtime" / "hxrt" / "core_test.go",
        "package hxrt\n\n// Repository-only runtime test.\n",
    )
    write_text(root / "runtime" / "hxrt" / "scratch.tmp", "untracked runtime debris\n")
    write_text(
        root / "vendor" / "reflaxe" / "src" / "reflaxe" / "ReflectCompiler.hx",
        "package reflaxe; class ReflectCompiler {}\n",
    )
    write_text(
        root / "vendor" / "reflaxe" / "src" / "reflaxe" / "scratch.txt",
        "untracked vendor debris\n",
    )
    write_text(root / "vendor" / "reflaxe" / "PATCHES.md", "# Fixture provenance\n")
    write_text(root / "vendor" / "reflaxe" / "LICENSE", "fixture vendor license\n")
    write_text(
        root / "vendor" / "reflaxe" / "FUTURE_MODIFICATIONS.md",
        "# Fixture future modifications\n",
    )
    write_json(
        root / "vendor" / "reflaxe" / "haxelib.json",
        {"name": "reflaxe", "version": "4.0.0-beta"},
    )
    write_text(root / "LICENSE", "fixture license\n")
    write_text(root / "LICENSING.md", "# Fixture licensing policy\n")
    write_text(root / "README.md", "# Fixture\n")
    write_text(root / "extraParams.hxml", "--macro fixture.Start()\n")
    write_json(root / "license-policy.json", {"fixture": True})
    write_text(
        root / "licenses" / "HAXE-GO-GENERATED-MIT.txt",
        "fixture haxe.go generated-output license\n",
    )
    write_text(
        root / "licenses" / "HAXE-STDLIB-MIT.txt",
        "fixture Haxe standard library license\n",
    )
    write_text(root / "scripts" / "dev" / "go-hx.sh", "#!/usr/bin/env bash\n")
    write_text(
        root / "scripts" / "dev" / "haxe_go_watch.py",
        "#!/usr/bin/env python3\n",
    )
    shutil.copyfile(RUNNER, root / "Run.hx")


def package_files(root: Path) -> list[str]:
    return sorted(
        (path.relative_to(root).as_posix() for path in root.rglob("*") if path.is_file()),
        key=lambda value: value.encode("utf-8"),
    )


class HaxelibPackageRunnerContractTest(unittest.TestCase):
    maxDiff = None

    def test_runner_and_release_entrypoints_are_typed_and_wired(self) -> None:
        self.assertTrue(RUNNER.is_file())
        runner = RUNNER.read_text(encoding="utf-8")
        self.assertIn("class PackageBuildConfig", runner)
        self.assertIn("typedef SourceHaxelibManifest", runner)
        self.assertIn("typedef PackageManifestEntry", runner)
        self.assertNotRegex(runner, r"\b(?:Dynamic|Any|Reflect)\b")

        self.assertTrue(PACKAGE_SCRIPT.is_file())
        self.assertTrue(ZIP_SCRIPT.is_file())
        package = json.loads((ROOT / "package.json").read_text(encoding="utf-8"))
        scripts = package.get("scripts", {})
        self.assertEqual(
            "bash scripts/release/package-haxelib.sh",
            scripts.get("package:haxelib"),
        )
        contract_command = "python3 test/test_haxelib_package_runner.py"
        self.assertEqual(contract_command, scripts.get("test:haxelib-package"))
        self.assertIn("npm run test:haxelib-package", scripts.get("test", ""))
        self.assertIn("npm run test:haxelib-package", scripts.get("test:changed", ""))
        release_contracts = (ROOT / "test" / "run-release-contracts.py").read_text(
            encoding="utf-8"
        )
        self.assertIn('"test/test_haxelib_package_runner.py"', release_contracts)

    @unittest.skipUnless(shutil.which("haxe"), "requires Haxe")
    def test_declared_std_paths_control_extension_mapping(self) -> None:
        with tempfile.TemporaryDirectory(prefix="haxe-go-package-runner-fixture-") as raw:
            temp = Path(raw)
            source_root = temp / "source"
            package_root = temp / "package"
            write_synthetic_source(source_root)

            process = run_package_runner(source_root, package_root)
            self.assertEqual(
                0,
                process.returncode,
                "package runner failed:\n" + process.stdout + process.stderr,
            )

            self.assertTrue((package_root / "src" / "Main.hx").is_file())
            self.assertTrue((package_root / "src" / "Support.hx").is_file())
            self.assertTrue(
                (package_root / "src" / "other" / "_std" / "NestedSupport.hx").is_file(),
                "an _std directory nested only inside an ordinary declared root stays ordinary",
            )
            self.assertFalse(
                (package_root / "src" / "other" / "_std" / "NestedSupport.cross.hx").exists()
            )
            self.assertTrue((package_root / "src" / "Override.cross.hx").is_file())
            self.assertFalse((package_root / "src" / "Override.hx").exists())
            self.assertFalse((package_root / "src" / "go" / "_std").exists())
            self.assertFalse((package_root / "std").exists())
            self.assertFalse((package_root / "src" / ".cache").exists())
            self.assertFalse((package_root / "runtime" / "hxrt" / "scratch.tmp").exists())
            self.assertFalse((package_root / "runtime" / "hxrt" / "core_test.go").exists())
            self.assertFalse(
                (
                    package_root
                    / "vendor"
                    / "reflaxe"
                    / "src"
                    / "reflaxe"
                    / "scratch.txt"
                ).exists()
            )

            packaged_haxelib = json.loads(
                (package_root / "haxelib.json").read_text(encoding="utf-8")
            )
            self.assertEqual("src", packaged_haxelib.get("classPath"))
            self.assertNotIn("reflaxe", packaged_haxelib)

    @unittest.skipUnless(shutil.which("haxe"), "requires Haxe")
    def test_runner_rejects_unsafe_paths_and_mapping_collisions(self) -> None:
        with tempfile.TemporaryDirectory(prefix="haxe-go-package-runner-reject-") as raw:
            temp = Path(raw)
            source_root = temp / "source"
            write_synthetic_source(source_root)
            manifest_path = source_root / "haxelib.json"
            manifest = json.loads(manifest_path.read_text(encoding="utf-8"))
            manifest["reflaxe"]["stdPaths"] = ["std", "../outside/_std"]
            write_json(manifest_path, manifest)

            unsafe = run_package_runner(source_root, temp / "unsafe-package")
            self.assertNotEqual(0, unsafe.returncode)
            self.assertIn("safe repository-relative path", unsafe.stderr)

            write_synthetic_source(source_root)
            write_text(source_root / "src" / "Support.hx", "class Collision {}\n")
            collision = run_package_runner(source_root, temp / "collision-package")
            self.assertNotEqual(0, collision.returncode)
            self.assertIn("package path collision", collision.stderr)
            self.assertIn("src/Support.hx", collision.stderr)

    @unittest.skipUnless(shutil.which("haxe"), "requires Haxe")
    def test_runner_rejects_cross_host_path_forms_and_output_symlink_escape(self) -> None:
        unsafe_paths = [
            "/private/outside",
            "//server/share/outside",
            "C:/outside",
            "C:outside",
            r"C:\outside",
            r"\\server\share\outside",
            "std//nested",
            "std/../outside",
            "NUL",
            "std:stream",
        ]
        for unsafe_path in unsafe_paths:
            with self.subTest(path=unsafe_path), tempfile.TemporaryDirectory(
                prefix="haxe-go-package-path-reject-"
            ) as raw:
                temp = Path(raw)
                source_root = temp / "source"
                write_synthetic_source(source_root)
                manifest_path = source_root / "haxelib.json"
                manifest = json.loads(manifest_path.read_text(encoding="utf-8"))
                manifest["reflaxe"]["stdPaths"] = [unsafe_path]
                write_json(manifest_path, manifest)

                proc = run_package_runner(source_root, temp / "package")
                self.assertNotEqual(0, proc.returncode, proc.stdout + proc.stderr)
                self.assertIn("safe repository-relative path", proc.stderr)
                self.assertNotIn(str(temp), proc.stdout + proc.stderr)

        with tempfile.TemporaryDirectory(prefix="haxe-go-package-symlink-reject-") as raw:
            temp = Path(raw)
            source_root = temp / "source"
            package_root = temp / "package"
            outside = temp / "outside"
            write_synthetic_source(source_root)
            package_root.mkdir()
            outside.mkdir()
            sentinel = outside / "must-survive.txt"
            sentinel.write_text("sentinel\n", encoding="utf-8")
            (package_root / "escape").symlink_to(outside, target_is_directory=True)

            proc = run_package_runner(source_root, package_root)

            self.assertNotEqual(0, proc.returncode, proc.stdout + proc.stderr)
            self.assertIn("package output contains a path outside its canonical root", proc.stderr)
            self.assertEqual("sentinel\n", sentinel.read_text(encoding="utf-8"))
            self.assertNotIn(str(temp), proc.stdout + proc.stderr)

        with tempfile.TemporaryDirectory(prefix="haxe-go-package-root-link-reject-") as raw:
            temp = Path(raw)
            source_root = temp / "source"
            package_root = temp / "package"
            outside = temp / "outside"
            write_synthetic_source(source_root)
            outside.mkdir()
            sentinel = outside / "must-survive.txt"
            sentinel.write_text("sentinel\n", encoding="utf-8")
            package_root.symlink_to(outside, target_is_directory=True)

            proc = run_package_runner(source_root, package_root)

            self.assertNotEqual(0, proc.returncode, proc.stdout + proc.stderr)
            self.assertIn("package output contains a path outside its canonical root", proc.stderr)
            self.assertEqual("sentinel\n", sentinel.read_text(encoding="utf-8"))
            self.assertNotIn(str(temp), proc.stdout + proc.stderr)

        with tempfile.TemporaryDirectory(prefix="haxe-go-package-source-link-reject-") as raw:
            temp = Path(raw)
            source_root = temp / "source"
            package_root = temp / "package"
            outside = temp / "outside.go"
            write_synthetic_source(source_root)
            outside.write_text("package outside\n", encoding="utf-8")
            runtime_file = source_root / "runtime" / "hxrt" / "core.go"
            runtime_file.unlink()
            runtime_file.symlink_to(outside)

            proc = run_package_runner(source_root, package_root)

            self.assertNotEqual(0, proc.returncode, proc.stdout + proc.stderr)
            self.assertIn("package source trees must not contain symbolic links", proc.stderr)
            self.assertEqual("package outside\n", outside.read_text(encoding="utf-8"))
            self.assertFalse((package_root / "runtime" / "hxrt" / "core.go").exists())
            self.assertNotIn(str(temp), proc.stdout + proc.stderr)

        with tempfile.TemporaryDirectory(prefix="haxe-go-package-existing-reject-") as raw:
            temp = Path(raw)
            source_root = temp / "source"
            package_root = temp / "package"
            write_synthetic_source(source_root)
            package_root.mkdir()

            proc = run_package_runner(source_root, package_root, clean=False)

            self.assertNotEqual(0, proc.returncode, proc.stdout + proc.stderr)
            self.assertIn("pass --clean to replace it", proc.stderr)
            self.assertNotIn(str(temp), proc.stdout + proc.stderr)

    @unittest.skipUnless(shutil.which("haxe"), "requires Haxe")
    def test_canonical_audit_rejects_tampered_package_manifest(self) -> None:
        with tempfile.TemporaryDirectory(prefix="haxe-go-package-manifest-tamper-") as raw:
            temp = Path(raw)
            source_root = temp / "source"
            package_root = temp / "package"
            write_synthetic_source(source_root)
            process = run_package_runner(source_root, package_root)
            self.assertEqual(0, process.returncode, process.stdout + process.stderr)
            self.assertEqual([], audit_package_layout(package_root, source_root))

            manifest_path = package_root / PACKAGE_MANIFEST
            manifest = json.loads(manifest_path.read_text(encoding="utf-8"))
            manifest["entries"][0]["packageSha256"] = "0" * 64
            write_json(manifest_path, manifest)

            violations = audit_package_layout(package_root, source_root)
            self.assertIn("package-map-manifest", {item.code for item in violations})

            archive = temp / "tampered.zip"
            zip_process = subprocess.run(
                [sys.executable, str(ZIP_SCRIPT), str(package_root), str(archive)],
                cwd=ROOT,
                capture_output=True,
                text=True,
                timeout=120,
            )
            self.assertNotEqual(0, zip_process.returncode)
            self.assertIn("package hash differs", zip_process.stderr)
            self.assertFalse(archive.exists())

    @unittest.skipUnless(shutil.which("haxe"), "requires Haxe")
    def test_live_package_manifest_and_archive_are_reproducible(self) -> None:
        with tempfile.TemporaryDirectory(prefix="haxe-go-package-runner-live-") as raw:
            temp = Path(raw)
            first_root = temp / "package-a"
            second_root = temp / "package-b"
            first = run_package_runner(ROOT, first_root)
            second = run_package_runner(ROOT, second_root)
            self.assertEqual(0, first.returncode, first.stdout + first.stderr)
            self.assertEqual(0, second.returncode, second.stdout + second.stderr)

            self.assertEqual([], audit_package_layout(first_root, ROOT))
            first_manifest_path = first_root / PACKAGE_MANIFEST
            second_manifest_path = second_root / PACKAGE_MANIFEST
            self.assertEqual(first_manifest_path.read_bytes(), second_manifest_path.read_bytes())

            manifest = json.loads(first_manifest_path.read_text(encoding="utf-8"))
            self.assertEqual(1, manifest.get("schemaVersion"))
            self.assertEqual("reflaxe.go-haxelib-package", manifest.get("format"))
            self.assertEqual(
                {
                    "compression": "stored",
                    "fileMode": "0644",
                    "ordering": "utf8-bytewise",
                    "timestamp": "2000-01-01T00:00:00Z",
                },
                manifest.get("archive"),
            )
            self.assertEqual("src", manifest.get("classPath"))

            entries = manifest.get("entries")
            self.assertIsInstance(entries, list)
            package_paths = [entry["packagePath"] for entry in entries]
            self.assertEqual(
                sorted(package_paths, key=lambda value: value.encode("utf-8")),
                package_paths,
            )
            self.assertEqual(len(package_paths), len(set(package_paths)))
            self.assertEqual(
                sorted(package_paths + [PACKAGE_MANIFEST], key=lambda value: value.encode("utf-8")),
                package_files(first_root),
            )

            for entry in entries:
                source_path = PurePosixPath(entry["sourcePath"])
                package_path = PurePosixPath(entry["packagePath"])
                self.assertFalse(source_path.is_absolute())
                self.assertFalse(package_path.is_absolute())
                self.assertNotIn("..", source_path.parts)
                self.assertNotIn("..", package_path.parts)
                source_file = ROOT.joinpath(*source_path.parts)
                package_file = first_root.joinpath(*package_path.parts)
                self.assertTrue(source_file.is_file(), entry)
                self.assertTrue(package_file.is_file(), entry)
                self.assertEqual(sha256(source_file), entry["sourceSha256"])
                self.assertEqual(sha256(package_file), entry["packageSha256"])
                self.assertEqual(package_file.stat().st_size, entry["size"])

            expected_cross = expected_packaged_cross_files(ROOT)
            actual_cross = {
                path
                for path in package_paths
                if path.endswith(".cross.hx")
            }
            self.assertEqual(expected_cross, actual_cross)
            by_source = {entry["sourcePath"]: entry for entry in entries}
            for source in sorted((ROOT / "std" / "go" / "_std").rglob("*.hx")):
                source_relative = source.relative_to(ROOT).as_posix()
                expected_path = (
                    Path("src")
                    / source.relative_to(ROOT / "std" / "go" / "_std").with_suffix(
                        ".cross.hx"
                    )
                ).as_posix()
                self.assertEqual(expected_path, by_source[source_relative]["packagePath"])
                self.assertEqual("stdlib-override", by_source[source_relative]["kind"])

            first_tree_hashes = {
                path: sha256(first_root / path) for path in package_files(first_root)
            }
            second_tree_hashes = {
                path: sha256(second_root / path) for path in package_files(second_root)
            }
            self.assertEqual(first_tree_hashes, second_tree_hashes)

            os.utime(first_root / "src" / "go" / "Go.hx", (1_800_000_000, 1_800_000_000))
            first_zip = temp / "first.zip"
            second_zip = temp / "second.zip"
            for package_root, output in ((first_root, first_zip), (second_root, second_zip)):
                process = subprocess.run(
                    [sys.executable, str(ZIP_SCRIPT), str(package_root), str(output)],
                    cwd=ROOT,
                    capture_output=True,
                    text=True,
                    timeout=120,
                )
                self.assertEqual(0, process.returncode, process.stdout + process.stderr)
            self.assertEqual(first_zip.read_bytes(), second_zip.read_bytes())

            with zipfile.ZipFile(first_zip) as archive:
                infos = archive.infolist()
                names = [info.filename for info in infos]
                self.assertEqual(
                    sorted(names, key=lambda value: value.encode("utf-8")),
                    names,
                )
                self.assertEqual(package_files(first_root), names)
                for info in infos:
                    self.assertEqual(FIXED_ZIP_TIMESTAMP, info.date_time)
                    self.assertEqual(0o644, (info.external_attr >> 16) & 0o777)
                    self.assertEqual(zipfile.ZIP_STORED, info.compress_type)


if __name__ == "__main__":
    unittest.main()
