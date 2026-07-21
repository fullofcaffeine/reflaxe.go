#!/usr/bin/env python3

from __future__ import annotations

import json
import shutil
import subprocess
import tempfile
import unittest
from pathlib import Path


ROOT = Path(__file__).resolve().parent.parent
GUARD = ROOT / "src" / "reflaxe" / "go" / "compiler" / "SiblingTargetConflictGuard.hx"
COMPILER_INIT = ROOT / "src" / "reflaxe" / "go" / "CompilerInit.hx"
RELEASE_CONTRACT_RUNNER = ROOT / "test" / "run-release-contracts.py"
CROSS_OVERRIDE_DOC = ROOT / "docs" / "cross-overrides-and-hardening.md"
START_HERE_DOC = ROOT / "docs" / "start-here.md"
DIAGNOSTIC_PREFIX = "Reflaxe.Go cannot compile with competing sibling targets: "


class SiblingTargetClasspathContractTest(unittest.TestCase):
    def test_guard_is_release_blocking_and_runs_before_compiler_setup(self) -> None:
        package = json.loads((ROOT / "package.json").read_text(encoding="utf-8"))
        scripts = package["scripts"]
        command = "python3 test/test_sibling_target_classpath_guard.py"
        self.assertEqual(command, scripts.get("test:sibling-target-classpaths"))
        self.assertIn("npm run test:sibling-target-classpaths", scripts["test"])
        self.assertIn("npm run test:sibling-target-classpaths", scripts["test:changed"])
        self.assertIn(
            "test/test_sibling_target_classpath_guard.py",
            RELEASE_CONTRACT_RUNNER.read_text(encoding="utf-8"),
        )

        self.assertTrue(GUARD.is_file())
        compiler_init = COMPILER_INIT.read_text(encoding="utf-8")
        guard_call = compiler_init.index("SiblingTargetConflictGuard.init();")
        compiler_setup = compiler_init.index("GoBuildContextResolver.resolve();")
        self.assertLess(guard_call, compiler_setup)

        guard = GUARD.read_text(encoding="utf-8")
        self.assertIn("Context.onAfterInitMacros(validate)", guard)
        self.assertNotRegex(guard, r"\b(?:Dynamic|Any|Reflect)\b")

    def test_documentation_explains_policy_and_sibling_precedent(self) -> None:
        document = CROSS_OVERRIDE_DOC.read_text(encoding="utf-8")
        for phrase in [
            "## Fail-fast contract",
            "source-checkout signal",
            "packaged signal",
            "reflaxe.rust at `85067736d0b929dfc67d6684d59b7e2bd3bae6ea`",
            "reflaxe.elixir at `17f1c66ae4c6bcae3c15cf694c16e63f27f2d9aa`",
        ]:
            self.assertIn(phrase, document)
        self.assertRegex(
            document,
            r"does not make\s+simultaneous multi-target compilation supported",
        )

        start_here = START_HERE_DOC.read_text(encoding="utf-8")
        for phrase in [
            "## One Reflaxe target per compilation",
            "Installing multiple Reflaxe target libraries is safe",
            "separate Haxe invocation",
            "`--next`",
        ]:
            self.assertIn(phrase, start_here)


@unittest.skipUnless(shutil.which("haxe"), "requires Haxe")
class SiblingTargetClasspathBehaviorTest(unittest.TestCase):
    maxDiff = None

    def run_compile(
        self,
        app: Path,
        *,
        extra_classpaths: tuple[Path, ...] = (),
        defines: tuple[str, ...] = (),
        late_macros: tuple[str, ...] = (),
        go_build: bool = True,
    ) -> subprocess.CompletedProcess[str]:
        command = [shutil.which("haxe") or "haxe", "-cp", str(app)]
        for classpath in extra_classpaths:
            command.extend(["-cp", str(classpath)])
        command.extend(["-cp", str(ROOT / "src")])
        if go_build:
            command.extend(
                [
                    "-cp",
                    str(ROOT / "std"),
                    "-cp",
                    str(ROOT / "std" / "go" / "_std"),
                ]
            )
        command.extend(
            [
                "-cp",
                str(ROOT / "vendor" / "reflaxe" / "src"),
                "--macro",
                "reflaxe.go.CompilerBootstrap.Start()",
                "--macro",
                "reflaxe.go.CompilerInit.Start()",
            ]
        )
        for macro in late_macros:
            command.extend(["--macro", macro])
        for define in defines:
            command.extend(["-D", define])
        if go_build:
            command.extend(
                [
                    "-D",
                    f"go_output={app / 'out'}",
                    "-D",
                    "reflaxe.dont_output_metadata_id",
                ]
            )
        command.extend(["-main", "Main"])
        if not go_build:
            command.append("--interp")
        return subprocess.run(
            command,
            cwd=app,
            capture_output=True,
            text=True,
            timeout=120,
        )

    def make_app(self, temp: Path) -> Path:
        app = temp / "app"
        app.mkdir()
        (app / "Main.hx").write_text(
            "class Main { static function main():Void {} }\n",
            encoding="utf-8",
        )
        return app

    def test_source_and_packaged_signals_fail_with_sorted_target_names(self) -> None:
        with tempfile.TemporaryDirectory(prefix="haxe-go-sibling-source-") as raw:
            temp = Path(raw)
            app = self.make_app(temp)
            rust_std = temp / "sibling" / "std" / "rust" / "_std"
            rust_std.mkdir(parents=True)
            packaged_ruby = temp / "reflaxe.ruby" / "package" / "src"
            packaged_ruby.mkdir(parents=True)
            (packaged_ruby / "StringTools.cross.hx").write_text(
                "class StringTools {}\n",
                encoding="utf-8",
            )

            completed = self.run_compile(
                app,
                extra_classpaths=(rust_std, packaged_ruby),
                defines=("reflaxe.ruby=fixture",),
            )

            self.assertNotEqual(0, completed.returncode)
            diagnostic = completed.stdout + completed.stderr
            self.assertIn(
                DIAGNOSTIC_PREFIX + "reflaxe.ruby, reflaxe.rust.",
                diagnostic,
            )
            self.assertNotIn(str(temp), diagnostic)

    def test_late_sibling_activation_is_caught_after_init_macros(self) -> None:
        with tempfile.TemporaryDirectory(prefix="haxe-go-sibling-late-") as raw:
            temp = Path(raw)
            app = self.make_app(temp)
            (app / "LateSiblingActivation.hx").write_text(
                "#if macro\n"
                "class LateSiblingActivation {\n"
                "  public static function run():Void {\n"
                '    haxe.macro.Compiler.define("reflaxe.elixir", "fixture");\n'
                "  }\n"
                "}\n"
                "#end\n",
                encoding="utf-8",
            )

            completed = self.run_compile(
                app,
                late_macros=("LateSiblingActivation.run()",),
            )

            self.assertNotEqual(0, completed.returncode)
            diagnostic = completed.stdout + completed.stderr
            self.assertIn(DIAGNOSTIC_PREFIX + "reflaxe.elixir.", diagnostic)

    def test_harmless_classpaths_leave_single_target_go_build_unchanged(self) -> None:
        with tempfile.TemporaryDirectory(prefix="haxe-go-sibling-harmless-") as raw:
            temp = Path(raw)
            app = self.make_app(temp)
            harmless = temp / "libraries" / "rust_helpers" / "src"
            harmless.mkdir(parents=True)
            (harmless / "Helper.hx").write_text("class Helper {}\n", encoding="utf-8")

            completed = self.run_compile(app, extra_classpaths=(harmless,))

            self.assertEqual(
                0,
                completed.returncode,
                "single-target Go compile failed:\n" + completed.stdout + completed.stderr,
            )

    def test_non_go_init_ignores_sibling_target_signals(self) -> None:
        with tempfile.TemporaryDirectory(prefix="haxe-go-sibling-non-go-") as raw:
            temp = Path(raw)
            app = self.make_app(temp)
            rust_std = temp / "std" / "rust" / "_std"
            rust_std.mkdir(parents=True)

            completed = self.run_compile(
                app,
                extra_classpaths=(rust_std,),
                defines=("reflaxe.rust=fixture",),
                go_build=False,
            )

            self.assertEqual(
                0,
                completed.returncode,
                "non-Go macro use failed:\n" + completed.stdout + completed.stderr,
            )


if __name__ == "__main__":
    unittest.main()
