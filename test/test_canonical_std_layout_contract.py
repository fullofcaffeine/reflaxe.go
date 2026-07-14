#!/usr/bin/env python3

from __future__ import annotations

import hashlib
import json
from pathlib import Path
import shutil
import subprocess
import sys
import tempfile
import unittest


ROOT = Path(__file__).resolve().parent.parent
FIXTURE_ROOT = ROOT / "test" / "fixtures" / "canonical_std_selection"
STATUS_PATH = ROOT / "test" / "canonical_std_layout_status.json"
sys.path.insert(0, str(ROOT / "scripts" / "ci"))

from canonical_stdlib_layout_check import (  # noqa: E402
    audit_package_layout,
    audit_source_layout,
)


def write_json(path: Path, value: object) -> None:
    path.parent.mkdir(parents=True, exist_ok=True)
    path.write_text(json.dumps(value, indent=2) + "\n", encoding="utf-8")


def write_canonical_source(root: Path) -> None:
    write_json(
        root / "haxelib.json",
        {
            "name": "reflaxe.go",
            "classPath": "src",
            "reflaxe": {
                "name": "Go",
                "abbv": "go",
                "stdPaths": ["std", "std/go/_std"],
            },
        },
    )
    hxml = root / "haxe_libraries" / "reflaxe.go.hxml"
    hxml.parent.mkdir(parents=True, exist_ok=True)
    hxml.write_text(
        "-cp ${SCOPE_DIR}/src/\n"
        "-cp ${SCOPE_DIR}/std/\n"
        "-cp ${SCOPE_DIR}/std/go/_std/\n"
        "-lib reflaxe\n"
        "--macro reflaxe.go.CompilerBootstrap.Start()\n"
        "--macro reflaxe.go.CompilerInit.Start()\n",
        encoding="utf-8",
    )
    (root / "haxe_libraries" / "reflaxe.hxml").write_text(
        "-cp ${SCOPE_DIR}/vendor/reflaxe/src/\n",
        encoding="utf-8",
    )
    override = root / "std" / "go" / "_std" / "Lambda.hx"
    override.parent.mkdir(parents=True, exist_ok=True)
    shutil.copyfile(FIXTURE_ROOT / "SourceProbe.hx.fixture", override)

    for support_root in ("haxe", "hxrt", "sys"):
        shutil.copytree(
            ROOT / "std" / support_root,
            root / "std" / support_root,
            dirs_exist_ok=True,
        )
    for facade in (ROOT / "std" / "go").glob("*.hx"):
        destination = root / "std" / "go" / facade.name
        destination.parent.mkdir(parents=True, exist_ok=True)
        shutil.copyfile(facade, destination)


def write_package_manifest(package_root: Path, source_root: Path) -> None:
    entries: list[dict[str, object]] = []
    for package_file in sorted(
        (
            path
            for path in package_root.rglob("*")
            if path.is_file() and path.name != "reflaxe-package-manifest.json"
        ),
        key=lambda path: path.relative_to(package_root).as_posix().encode("utf-8"),
    ):
        package_path = package_file.relative_to(package_root).as_posix()
        if package_path == "haxelib.json":
            source_file = source_root / "haxelib.json"
            kind = "metadata"
        elif package_path == "src/Lambda.cross.hx":
            source_file = source_root / "std" / "go" / "_std" / "Lambda.hx"
            kind = "stdlib-override"
        elif package_path.startswith("src/"):
            relative = package_file.relative_to(package_root / "src")
            std_source = source_root / "std" / relative
            class_path_source = source_root / "src" / relative
            if std_source.is_file():
                source_file = std_source
                kind = "stdlib"
            else:
                source_file = class_path_source
                kind = "class-path"
        elif package_path.startswith("runtime/"):
            source_file = source_root / package_path
            kind = "runtime"
        elif package_path.startswith("vendor/reflaxe/"):
            source_file = source_root / package_path
            kind = "vendored-reflaxe"
        else:
            raise AssertionError(f"synthetic package has no source mapping for {package_path}")
        assert source_file.is_file(), source_file
        entries.append(
            {
                "sourcePath": source_file.relative_to(source_root).as_posix(),
                "packagePath": package_path,
                "kind": kind,
                "sourceSha256": hashlib.sha256(source_file.read_bytes()).hexdigest(),
                "packageSha256": hashlib.sha256(package_file.read_bytes()).hexdigest(),
                "size": package_file.stat().st_size,
            }
        )
    write_json(
        package_root / "reflaxe-package-manifest.json",
        {
            "schemaVersion": 1,
            "format": "reflaxe.go-haxelib-package",
            "archive": {
                "compression": "stored",
                "fileMode": "0644",
                "ordering": "utf8-bytewise",
                "timestamp": "2000-01-01T00:00:00Z",
            },
            "classPath": "src",
            "entries": entries,
        },
    )


def write_canonical_package(package_root: Path, source_root: Path) -> None:
    write_json(
        package_root / "haxelib.json",
        {"name": "reflaxe.go", "classPath": "src", "version": "0.0.0"},
    )
    packaged_override = package_root / "src" / "Lambda.cross.hx"
    packaged_override.parent.mkdir(parents=True, exist_ok=True)
    shutil.copyfile(FIXTURE_ROOT / "PackageProbe.hx.fixture", packaged_override)

    canonical_root = source_root / "std" / "go" / "_std"
    for source in (source_root / "std").rglob("*.hx"):
        if source.is_relative_to(canonical_root):
            continue
        destination = package_root / "src" / source.relative_to(source_root / "std")
        destination.parent.mkdir(parents=True, exist_ok=True)
        shutil.copyfile(source, destination)

    # The mapping assertion needs the ordinary source authority as its input.
    assert (source_root / "std" / "go" / "_std" / "Lambda.hx").is_file()
    write_package_manifest(package_root, source_root)


class CanonicalStdLayoutAuditTest(unittest.TestCase):
    def test_full_changed_and_ci_entrypoints_run_this_contract(self) -> None:
        package = json.loads((ROOT / "package.json").read_text(encoding="utf-8"))
        scripts = package.get("scripts", {})
        contract_command = "python3 test/test_canonical_std_layout_contract.py"
        self.assertEqual(contract_command, scripts.get("test:canonical-std-layout"))
        self.assertIn("npm run test:canonical-std-layout", scripts.get("test", ""))
        self.assertIn("npm run test:canonical-std-layout", scripts.get("test:changed", ""))

        ci_runner = (ROOT / "test" / "run-ci.py").read_text(encoding="utf-8")
        self.assertIn("build_canonical_std_layout_command", ci_runner)
        self.assertIn("Canonical std layout contract stage", ci_runner)

    def test_current_source_layout_is_explicitly_red_until_migrated(self) -> None:
        status = json.loads(STATUS_PATH.read_text(encoding="utf-8"))
        self.assertEqual(status.get("schemaVersion"), 2)

        violations = audit_source_layout(ROOT)
        codes = {violation.code for violation in violations}
        state = status.get("sourceLayout")

        if state == "expected-red":
            self.assertTrue(
                violations,
                "canonical source layout unexpectedly became green; flip sourceLayout to required-green",
            )
            allowed = set(status.get("expectedViolationCodes", []))
            self.assertFalse(
                codes - allowed,
                "canonical source layout gained unapproved violations: " + ", ".join(sorted(codes - allowed)),
            )
            print(
                "[canonical-std] EXPECTED RED: "
                + "; ".join(violation.render() for violation in violations)
            )
            return

        self.assertEqual(state, "required-green")
        self.assertEqual([], violations, "canonical source layout must be green")

    def test_synthetic_canonical_source_and_package_layouts_pass(self) -> None:
        with tempfile.TemporaryDirectory(prefix="haxe-go-canonical-std-layout-") as raw:
            temp = Path(raw)
            source_root = temp / "source"
            package_root = temp / "package"
            write_canonical_source(source_root)
            write_canonical_package(package_root, source_root)

            self.assertEqual([], audit_source_layout(source_root))
            self.assertEqual([], audit_package_layout(package_root, source_root))

    def test_source_contract_rejects_cross_files_and_wrong_precedence(self) -> None:
        with tempfile.TemporaryDirectory(prefix="haxe-go-canonical-std-bad-source-") as raw:
            source_root = Path(raw)
            write_canonical_source(source_root)
            (source_root / "std" / "Legacy.cross.hx").write_text("class Legacy {}\n", encoding="utf-8")
            hxml = source_root / "haxe_libraries" / "reflaxe.go.hxml"
            hxml.write_text(
                hxml.read_text(encoding="utf-8").replace(
                    "-cp ${SCOPE_DIR}/std/\n-cp ${SCOPE_DIR}/std/go/_std/\n",
                    "-cp ${SCOPE_DIR}/std/go/_std/\n-cp ${SCOPE_DIR}/std/\n",
                ).replace("-lib reflaxe\n", ""),
                encoding="utf-8",
            )
            override = source_root / "std" / "go" / "_std" / "Lambda.hx"
            rejected_path = "/" + "Users/example/checkout"
            override.write_text(
                override.read_text(encoding="utf-8") + f"\n// rejected fixture path: {rejected_path}\n",
                encoding="utf-8",
            )

            codes = {violation.code for violation in audit_source_layout(source_root)}
            self.assertIn("source-cross-files", codes)
            self.assertIn("source-classpath-precedence", codes)
            self.assertIn("source-vendored-reflaxe-pretyping", codes)
            self.assertIn("absolute-path-leak", codes)

    def test_source_contract_rejects_legacy_support_root_and_classpath(self) -> None:
        with tempfile.TemporaryDirectory(prefix="haxe-go-canonical-std-legacy-root-") as raw:
            source_root = Path(raw)
            write_canonical_source(source_root)
            legacy_support = source_root / "std" / "_std" / "LegacySupport.hx"
            legacy_support.parent.mkdir(parents=True)
            legacy_support.write_text("class LegacySupport {}\n", encoding="utf-8")

            codes = {violation.code for violation in audit_source_layout(source_root)}
            self.assertIn("source-legacy-support-root", codes)

        with tempfile.TemporaryDirectory(prefix="haxe-go-canonical-std-legacy-classpath-") as raw:
            source_root = Path(raw)
            write_canonical_source(source_root)
            hxml = source_root / "haxe_libraries" / "reflaxe.go.hxml"
            hxml.write_text(
                hxml.read_text(encoding="utf-8").replace(
                    "-cp ${SCOPE_DIR}/std/go/_std/\n",
                    "-cp ${SCOPE_DIR}/std/_std/\n-cp ${SCOPE_DIR}/std/go/_std/\n",
                ),
                encoding="utf-8",
            )

            codes = {violation.code for violation in audit_source_layout(source_root)}
            self.assertIn("source-legacy-support-classpath", codes)

    def test_live_bootstrap_avoids_reflective_classpath_surgery(self) -> None:
        bootstrap = (ROOT / "src" / "reflaxe" / "go" / "CompilerBootstrap.hx").read_text(
            encoding="utf-8"
        )
        self.assertNotIn("Compiler.getConfiguration", bootstrap)
        self.assertNotIn("injectClassPathsFirst", bootstrap)
        self.assertNotRegex(bootstrap, r"\b(?:Dynamic|Reflect)\b")
        self.assertIn("Compiler.addClassPath(vendoredReflaxe)", bootstrap)
        self.assertIn("Context.fatalError", bootstrap)

    def test_package_contract_rejects_unflattened_or_path_leaking_layout(self) -> None:
        with tempfile.TemporaryDirectory(prefix="haxe-go-canonical-std-bad-package-") as raw:
            temp = Path(raw)
            source_root = temp / "source"
            package_root = temp / "package"
            write_canonical_source(source_root)
            write_canonical_package(package_root, source_root)
            (package_root / "src" / "hxrt" / "string" / "GoStringRuntime.hx").unlink()
            (package_root / "std" / "go" / "_std").mkdir(parents=True)
            packaged_override = package_root / "src" / "Lambda.cross.hx"
            packaged_override.write_text(
                packaged_override.read_text(encoding="utf-8")
                + '\nclass LocalPathLeak { static final value = "C:\\\\Users\\\\example\\\\checkout"; }\n',
                encoding="utf-8",
            )

            codes = {violation.code for violation in audit_package_layout(package_root, source_root)}
            self.assertIn("package-unflattened-std", codes)
            self.assertIn("package-ordinary-support-mapping", codes)
            self.assertIn("absolute-path-leak", codes)


@unittest.skipUnless(shutil.which("haxe") and shutil.which("go"), "requires Haxe and Go")
class CanonicalStdSelectionBehaviorTest(unittest.TestCase):
    maxDiff = None

    def test_bootstrap_diagnoses_missing_reflaxe_and_accepts_flattened_package(self) -> None:
        with tempfile.TemporaryDirectory(prefix="haxe-go-bootstrap-layout-") as raw:
            package_root = Path(raw)
            compiler_src = package_root / "src"
            app = package_root / "app"
            bootstrap = compiler_src / "reflaxe" / "go" / "CompilerBootstrap.hx"
            bootstrap.parent.mkdir(parents=True)
            app.mkdir()
            shutil.copyfile(ROOT / "src" / "reflaxe" / "go" / "CompilerBootstrap.hx", bootstrap)
            (app / "Main.hx").write_text(
                "class Main { static function main():Void {} }\n",
                encoding="utf-8",
            )
            command = [
                shutil.which("haxe") or "haxe",
                "-cp",
                str(app),
                "-cp",
                str(compiler_src),
                "--macro",
                "reflaxe.go.CompilerBootstrap.Start()",
                "-main",
                "Main",
                "--interp",
            ]

            missing = subprocess.run(
                command,
                cwd=app,
                capture_output=True,
                text=True,
                timeout=120,
            )
            self.assertNotEqual(0, missing.returncode)
            self.assertIn("could not resolve its vendored Reflaxe framework", missing.stderr)

            flattened_reflaxe = compiler_src / "reflaxe" / "ReflectCompiler.hx"
            flattened_reflaxe.write_text(
                "package reflaxe; class ReflectCompiler {}\n",
                encoding="utf-8",
            )
            packaged = subprocess.run(
                command,
                cwd=app,
                capture_output=True,
                text=True,
                timeout=120,
            )
            self.assertEqual(
                0,
                packaged.returncode,
                "flattened package bootstrap failed:\n" + packaged.stdout + packaged.stderr,
            )

    def test_non_go_bootstrap_resolves_vendored_reflaxe(self) -> None:
        with tempfile.TemporaryDirectory(prefix="haxe-go-non-go-bootstrap-") as raw:
            app = Path(raw)
            (app / "Main.hx").write_text(
                "class Main { static function main():Void {} }\n",
                encoding="utf-8",
            )
            (app / "VendoredReflaxeProbe.hx").write_text(
                "#if macro\n"
                "class VendoredReflaxeProbe {\n"
                "  public static function run():Void {\n"
                '    haxe.macro.Context.resolvePath("reflaxe/ReflectCompiler.hx");\n'
                "  }\n"
                "}\n"
                "#end\n",
                encoding="utf-8",
            )
            process = subprocess.run(
                [
                    shutil.which("haxe") or "haxe",
                    "-cp",
                    str(app),
                    "-cp",
                    str(ROOT / "src"),
                    "--macro",
                    "reflaxe.go.CompilerBootstrap.Start()",
                    "--macro",
                    "VendoredReflaxeProbe.run()",
                    "-main",
                    "Main",
                    "--interp",
                ],
                cwd=app,
                capture_output=True,
                text=True,
                timeout=120,
            )
            self.assertEqual(
                0,
                process.returncode,
                "non-Go bootstrap failed:\n" + process.stdout + process.stderr,
            )

    def run_haxe_go(
        self,
        *,
        app: Path,
        upstream: Path,
        target_roots: list[Path],
        compiler_src: Path,
        reflaxe_src: Path,
        output: Path,
    ) -> str:
        command = [
            shutil.which("haxe") or "haxe",
            "-cp",
            str(app),
            "-cp",
            str(upstream),
        ]
        for target_root in target_roots:
            command.extend(["-cp", str(target_root)])
        command.extend(
            [
                "-cp",
                str(compiler_src),
                "-cp",
                str(reflaxe_src),
                "--macro",
                "reflaxe.go.CompilerBootstrap.Start()",
                "--macro",
                "reflaxe.go.CompilerInit.Start()",
                "-D",
                f"go_output={output}",
                "-D",
                "reflaxe.dont_output_metadata_id",
                "-main",
                "Main",
            ]
        )
        compile_process = subprocess.run(
            command,
            cwd=app,
            capture_output=True,
            text=True,
            timeout=120,
        )
        self.assertEqual(
            0,
            compile_process.returncode,
            "Haxe compile failed:\n" + compile_process.stdout + compile_process.stderr,
        )

        go_test = subprocess.run(
            [shutil.which("go") or "go", "test", "./..."],
            cwd=output,
            capture_output=True,
            text=True,
            timeout=120,
        )
        self.assertEqual(0, go_test.returncode, "go test failed:\n" + go_test.stdout + go_test.stderr)

        go_run = subprocess.run(
            [shutil.which("go") or "go", "run", "."],
            cwd=output,
            capture_output=True,
            text=True,
            timeout=120,
        )
        self.assertEqual(0, go_run.returncode, "go run failed:\n" + go_run.stdout + go_run.stderr)
        return go_run.stdout

    def test_source_and_staged_package_overrides_win_by_runtime_behavior(self) -> None:
        with tempfile.TemporaryDirectory(prefix="haxe-go-canonical-std-selection-") as raw:
            temp = Path(raw)
            app = temp / "app"
            upstream = temp / "upstream"
            source_root = temp / "source"
            package_root = temp / "package"
            app.mkdir()
            upstream.mkdir()
            shutil.copyfile(FIXTURE_ROOT / "Main.hx", app / "Main.hx")
            shutil.copyfile(FIXTURE_ROOT / "UpstreamProbe.hx.fixture", upstream / "Lambda.hx")
            write_canonical_source(source_root)
            shutil.copytree(ROOT / "src", source_root / "src")
            shutil.copytree(ROOT / "vendor" / "reflaxe", source_root / "vendor" / "reflaxe")
            shutil.copytree(ROOT / "runtime", source_root / "runtime")
            write_canonical_package(package_root, source_root)
            shutil.copytree(source_root / "src", package_root / "src", dirs_exist_ok=True)
            shutil.copytree(source_root / "vendor" / "reflaxe", package_root / "vendor" / "reflaxe")
            shutil.copytree(source_root / "runtime", package_root / "runtime")
            write_package_manifest(package_root, source_root)
            self.assertEqual([], audit_source_layout(source_root))
            self.assertEqual([], audit_package_layout(package_root, source_root))

            source_output = temp / "out-source"
            package_output = temp / "out-package"
            source_stdout = self.run_haxe_go(
                app=app,
                upstream=upstream,
                target_roots=[source_root / "std", source_root / "std" / "go" / "_std"],
                compiler_src=ROOT / "src",
                reflaxe_src=ROOT / "vendor" / "reflaxe" / "src",
                output=source_output,
            )
            package_stdout = self.run_haxe_go(
                app=app,
                upstream=upstream,
                target_roots=[package_root / "src"],
                compiler_src=package_root / "src",
                reflaxe_src=package_root / "vendor" / "reflaxe" / "src",
                output=package_output,
            )

            self.assertEqual("source-override\n65\nOK\nsupport-ok\n", source_stdout)
            self.assertEqual("package-override\n65\nOK\nsupport-ok\n", package_stdout)

            for output in (source_output, package_output):
                for path in output.rglob("*"):
                    if not path.is_file() or path.suffix not in {".go", ".json", ".mod", ".sum", ".txt"}:
                        continue
                    content = path.read_text(encoding="utf-8", errors="replace")
                    self.assertNotIn(str(ROOT), content, f"checkout path leaked into {path.name}")
                    self.assertNotIn(str(temp), content, f"temporary path leaked into {path.name}")


if __name__ == "__main__":
    unittest.main()
