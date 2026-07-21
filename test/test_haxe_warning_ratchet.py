#!/usr/bin/env python3

from __future__ import annotations

from collections import Counter
import json
import re
import shutil
import subprocess
import tempfile
import unittest
from pathlib import Path


ROOT = Path(__file__).resolve().parent.parent
POLICY_PATH = ROOT / "test" / "haxe_warning_policy.json"
PACKAGE_PATH = ROOT / "package.json"
RELEASE_RUNNER = ROOT / "test" / "run-release-contracts.py"
SOURCE_LIBRARY_HXML = ROOT / "haxe_libraries" / "reflaxe.go.hxml"
PACKAGE_EXTRA_PARAMS = ROOT / "extraParams.hxml"
NULL_SAFETY_MACRO = '--macro nullSafety("reflaxe.go")'
WARNING_PATTERN = re.compile(
    r"^(.*?):\d+: characters \d+-\d+ : Warning : \((W[^)]+)\)",
    re.MULTILINE,
)


def normalize_warning_path(raw_path: str) -> str:
    normalized = raw_path.replace("\\", "/")
    root = ROOT.as_posix().rstrip("/") + "/"
    if normalized.startswith(root):
        return normalized[len(root) :]

    marker = "/std/"
    if marker in normalized:
        return "haxe-stdlib/" + normalized.rsplit(marker, 1)[1]
    return "external/" + Path(normalized).name


class HaxeWarningRatchetContractTest(unittest.TestCase):
    def test_compiler_package_enables_scoped_null_safety(self) -> None:
        self.assertIn(
            NULL_SAFETY_MACRO,
            SOURCE_LIBRARY_HXML.read_text(encoding="utf-8").splitlines(),
        )
        self.assertIn(
            NULL_SAFETY_MACRO,
            PACKAGE_EXTRA_PARAMS.read_text(encoding="utf-8").splitlines(),
        )

    def test_warning_gate_is_release_blocking(self) -> None:
        package = json.loads(PACKAGE_PATH.read_text(encoding="utf-8"))
        scripts = package["scripts"]
        command = "python3 test/test_haxe_warning_ratchet.py"

        self.assertEqual(command, scripts.get("test:haxe-warnings"))
        self.assertIn("npm run test:haxe-warnings", scripts["test"])
        self.assertIn("npm run test:haxe-warnings", scripts["test:changed"])
        self.assertIn(
            "test/test_haxe_warning_ratchet.py",
            RELEASE_RUNNER.read_text(encoding="utf-8"),
        )

    @unittest.skipUnless(shutil.which("haxe"), "requires Haxe")
    def test_project_warnings_do_not_exceed_the_reviewed_policy(self) -> None:
        policy = json.loads(POLICY_PATH.read_text(encoding="utf-8"))
        version = subprocess.run(
            [shutil.which("haxe") or "haxe", "--version"],
            capture_output=True,
            text=True,
            check=True,
        ).stdout.strip()
        self.assertEqual(policy["haxeVersion"], version)

        with tempfile.TemporaryDirectory(prefix="haxe-go-warning-ratchet-") as raw:
            temp = Path(raw)
            app = temp / "app"
            app.mkdir()
            (app / "Main.hx").write_text(
                'class Main { static function main():Void { trace("warning ratchet"); } }\n',
                encoding="utf-8",
            )

            command = [
                shutil.which("haxe") or "haxe",
                "-cp",
                str(app),
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
                "-w",
                policy["warningOptions"],
                "--macro",
                'nullSafety("reflaxe.go")',
                "--macro",
                "reflaxe.go.CompilerBootstrap.Start()",
                "--macro",
                "reflaxe.go.CompilerInit.Start()",
                "-D",
                f"go_output={temp / 'out'}",
                "-D",
                "go_no_build",
                "-D",
                "no-traces",
                "-D",
                "reflaxe.dont_output_metadata_id",
                "-main",
                "Main",
            ]
            completed = subprocess.run(
                command,
                cwd=app,
                capture_output=True,
                text=True,
                timeout=180,
            )

        output = completed.stdout + completed.stderr
        self.assertEqual(0, completed.returncode, "warning probe failed:\n" + output)
        current = Counter(
            (normalize_warning_path(path), warning)
            for path, warning in WARNING_PATTERN.findall(output)
        )

        project_count = sum(
            count
            for (path, _warning), count in current.items()
            if path.startswith("src/")
        )
        self.assertLessEqual(
            project_count,
            policy["projectMaxCount"],
            "project-owned Haxe warnings exceed the reviewed ceiling:\n"
            + "\n".join(
                f"  {path} {warning}: {count}"
                for (path, warning), count in sorted(current.items())
                if path.startswith("src/")
            ),
        )

        limits = {
            (entry["file"], entry["warning"]): entry["maxCount"]
            for entry in policy["externalLimits"]
        }
        unexpected = {
            key: count
            for key, count in current.items()
            if not key[0].startswith("src/")
            and (key not in limits or count > limits[key])
        }
        self.assertEqual(
            {},
            unexpected,
            "external Haxe warnings exceed the reviewed per-file ceilings",
        )


if __name__ == "__main__":
    unittest.main()
