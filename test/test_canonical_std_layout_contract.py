#!/usr/bin/env python3

from __future__ import annotations

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
        "--macro reflaxe.go.CompilerBootstrap.Start()\n"
        "--macro reflaxe.go.CompilerInit.Start()\n",
        encoding="utf-8",
    )
    override = root / "std" / "go" / "_std" / "Lambda.hx"
    override.parent.mkdir(parents=True, exist_ok=True)
    shutil.copyfile(FIXTURE_ROOT / "SourceProbe.hx.fixture", override)


def write_canonical_package(package_root: Path, source_root: Path) -> None:
    write_json(
        package_root / "haxelib.json",
        {"name": "reflaxe.go", "classPath": "src", "version": "0.0.0"},
    )
    packaged_override = package_root / "src" / "Lambda.cross.hx"
    packaged_override.parent.mkdir(parents=True, exist_ok=True)
    shutil.copyfile(FIXTURE_ROOT / "PackageProbe.hx.fixture", packaged_override)

    # The mapping assertion needs the ordinary source authority as its input.
    assert (source_root / "std" / "go" / "_std" / "Lambda.hx").is_file()


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
        self.assertEqual(status.get("schemaVersion"), 1)

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
                ),
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
            self.assertIn("absolute-path-leak", codes)

    def test_package_contract_rejects_unflattened_or_path_leaking_layout(self) -> None:
        with tempfile.TemporaryDirectory(prefix="haxe-go-canonical-std-bad-package-") as raw:
            temp = Path(raw)
            source_root = temp / "source"
            package_root = temp / "package"
            write_canonical_source(source_root)
            write_canonical_package(package_root, source_root)
            (package_root / "std" / "go" / "_std").mkdir(parents=True)
            packaged_override = package_root / "src" / "Lambda.cross.hx"
            packaged_override.write_text(
                packaged_override.read_text(encoding="utf-8")
                + '\nclass LocalPathLeak { static final value = "C:\\\\Users\\\\example\\\\checkout"; }\n',
                encoding="utf-8",
            )

            codes = {violation.code for violation in audit_package_layout(package_root, source_root)}
            self.assertIn("package-unflattened-std", codes)
            self.assertIn("absolute-path-leak", codes)


@unittest.skipUnless(shutil.which("haxe") and shutil.which("go"), "requires Haxe and Go")
class CanonicalStdSelectionBehaviorTest(unittest.TestCase):
    maxDiff = None

    def run_haxe_go(
        self,
        *,
        app: Path,
        upstream: Path,
        target_root: Path,
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
            "-cp",
            str(target_root),
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
            write_canonical_package(package_root, source_root)
            shutil.copytree(ROOT / "src", package_root / "src", dirs_exist_ok=True)
            shutil.copytree(ROOT / "vendor" / "reflaxe", package_root / "vendor" / "reflaxe")
            shutil.copytree(ROOT / "runtime", package_root / "runtime")
            self.assertEqual([], audit_source_layout(source_root))
            self.assertEqual([], audit_package_layout(package_root, source_root))

            source_output = temp / "out-source"
            package_output = temp / "out-package"
            source_stdout = self.run_haxe_go(
                app=app,
                upstream=upstream,
                target_root=source_root / "std" / "go" / "_std",
                compiler_src=ROOT / "src",
                reflaxe_src=ROOT / "vendor" / "reflaxe" / "src",
                output=source_output,
            )
            package_stdout = self.run_haxe_go(
                app=app,
                upstream=upstream,
                target_root=package_root / "src",
                compiler_src=package_root / "src",
                reflaxe_src=package_root / "vendor" / "reflaxe" / "src",
                output=package_output,
            )

            self.assertEqual("source-override\n", source_stdout)
            self.assertEqual("package-override\n", package_stdout)

            for output in (source_output, package_output):
                for path in output.rglob("*"):
                    if not path.is_file() or path.suffix not in {".go", ".json", ".mod", ".sum", ".txt"}:
                        continue
                    content = path.read_text(encoding="utf-8", errors="replace")
                    self.assertNotIn(str(ROOT), content, f"checkout path leaked into {path.name}")
                    self.assertNotIn(str(temp), content, f"temporary path leaked into {path.name}")


if __name__ == "__main__":
    unittest.main()
