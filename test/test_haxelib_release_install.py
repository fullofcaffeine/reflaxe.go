#!/usr/bin/env python3

from __future__ import annotations

import json
from pathlib import Path
import shutil
import subprocess
import sys
import tempfile
import unittest
import zipfile


ROOT = Path(__file__).resolve().parent.parent
ARTIFACT_BUILDER = ROOT / "scripts" / "release" / "build-haxelib-artifact.py"
SMOKE_RUNNER = ROOT / "scripts" / "ci" / "run-isolated-haxelib-smoke.py"
FIXTURE_ROOT = ROOT / "test" / "fixtures" / "haxelib_release_install"
POLICY_DOC = ROOT / "docs" / "release-version-policy.md"
PACKAGE_PATH = ROOT / "package.json"
RELEASE_RUNNER_PATH = ROOT / "test" / "run-release-contracts.py"
VERSION = "0.999996.0"
TAG = f"v{VERSION}"


def head_sha() -> str:
    return subprocess.run(
        ["git", "rev-parse", "HEAD"],
        cwd=ROOT,
        check=True,
        capture_output=True,
        text=True,
    ).stdout.strip()


def tracked_status() -> str:
    return subprocess.run(
        ["git", "status", "--porcelain=v1", "--untracked-files=no"],
        cwd=ROOT,
        check=True,
        capture_output=True,
        text=True,
    ).stdout


def run_smoke(
    archive: Path,
    fixture_root: Path,
    profile: str,
) -> subprocess.CompletedProcess[str]:
    return subprocess.run(
        [
            sys.executable,
            str(SMOKE_RUNNER),
            "--archive",
            str(archive),
            "--fixture-root",
            str(fixture_root),
            "--profile",
            profile,
        ],
        cwd=ROOT,
        capture_output=True,
        text=True,
        timeout=240,
    )


class HaxelibReleaseInstallContractTest(unittest.TestCase):
    maxDiff = None

    def test_release_install_matrix_is_documented_and_release_blocking(self) -> None:
        self.assertTrue(ARTIFACT_BUILDER.is_file())
        self.assertTrue(SMOKE_RUNNER.is_file())
        for profile in ("portable", "metal"):
            fixture = FIXTURE_ROOT / profile
            self.assertTrue((fixture / "Main.hx").is_file(), profile)
            self.assertTrue((fixture / "expected.stdout").is_file(), profile)
        portable_source = (FIXTURE_ROOT / "portable" / "Main.hx").read_text(
            encoding="utf-8"
        )
        self.assertIn("StringTools.replace", portable_source)
        self.assertIn("sys.thread.Mutex", portable_source)
        metal_source = (FIXTURE_ROOT / "metal" / "Main.hx").read_text(
            encoding="utf-8"
        )
        self.assertIn("import go.Fmt", metal_source)
        self.assertIn("Fmt.println", metal_source)

        package = json.loads(PACKAGE_PATH.read_text(encoding="utf-8"))
        scripts = package.get("scripts", {})
        contract = "python3 test/test_haxelib_release_install.py"
        self.assertEqual(contract, scripts.get("test:haxelib-release-install"))
        self.assertIn(
            '"test/test_haxelib_release_install.py"',
            RELEASE_RUNNER_PATH.read_text(encoding="utf-8"),
        )

        smoke_runner = SMOKE_RUNNER.read_text(encoding="utf-8")
        self.assertIn('parser.add_argument("--profile"', smoke_runner)
        self.assertIn('choices=("portable", "metal")', smoke_runner)

        policy = " ".join(POLICY_DOC.read_text(encoding="utf-8").split())
        for phrase in (
            "Each preset runs in a fresh local Haxelib repository",
            "canonical std override and an `hxrt` capability",
            "typed `go.*` facade",
            "test:haxelib-release-install",
        ):
            self.assertIn(phrase, policy)

    @unittest.skipUnless(
        all(shutil.which(tool) for tool in ("git", "haxe", "haxelib", "go")),
        "requires Git, Haxe, Haxelib, and Go",
    )
    def test_exact_release_zip_runs_portable_and_metal_in_clean_repositories(self) -> None:
        with tempfile.TemporaryDirectory(prefix="haxe-go-release-install-") as raw:
            temp = Path(raw)
            output_dir = temp / "artifacts"
            before = tracked_status()
            build = subprocess.run(
                [
                    sys.executable,
                    str(ARTIFACT_BUILDER),
                    "--version",
                    VERSION,
                    "--tag",
                    TAG,
                    "--source-sha",
                    head_sha(),
                    "--output-dir",
                    str(output_dir),
                ],
                cwd=ROOT,
                capture_output=True,
                text=True,
                timeout=300,
            )
            self.assertEqual(0, build.returncode, build.stdout + build.stderr)
            archive = output_dir / f"reflaxe.go-{VERSION}.zip"
            self.assertTrue(archive.is_file())

            outer_manifest = json.loads(
                (output_dir / f"reflaxe.go-{VERSION}.manifest.json").read_text(
                    encoding="utf-8"
                )
            )
            self.assertEqual(head_sha(), outer_manifest["release"]["sourceCommit"])
            self.assertTrue(outer_manifest["reproducibility"]["byteIdentical"])
            with zipfile.ZipFile(archive) as package:
                haxelib = json.loads(package.read("haxelib.json"))
            self.assertEqual(VERSION, haxelib.get("version"))
            self.assertEqual({}, haxelib.get("dependencies"))

            for profile in ("portable", "metal"):
                smoke = run_smoke(archive, FIXTURE_ROOT / profile, profile)
                self.assertEqual(
                    0,
                    smoke.returncode,
                    f"{profile} release install failed:\n{smoke.stdout}{smoke.stderr}",
                )
                evidence = json.loads(smoke.stdout)
                self.assertEqual(VERSION, evidence["package"]["version"])
                self.assertEqual(
                    {"name": profile, "profile": profile},
                    evidence.get("fixture"),
                )
                self.assertEqual(
                    {
                        "checkoutClasspathsAbsent": True,
                        "generatedPathLeaksAbsent": True,
                        "goRun": "pass",
                        "goTest": "pass",
                        "haxeCompile": "pass",
                        "haxelibInstall": "pass",
                        "stdout": "pass",
                    },
                    evidence.get("checks"),
                )
                self.assertNotIn(str(ROOT), smoke.stdout)
                self.assertNotIn(str(temp), smoke.stdout)

            invalid_fixture = temp / "invalid-fixture"
            invalid_fixture.mkdir()
            (invalid_fixture / "Main.hx").write_text(
                "class Main { static function main():Void { this is invalid } }\n",
                encoding="utf-8",
            )
            (invalid_fixture / "expected.stdout").write_text(
                "never\n",
                encoding="utf-8",
            )
            compile_failure = run_smoke(archive, invalid_fixture, "portable")
            self.assertNotEqual(0, compile_failure.returncode)
            self.assertIn("ERROR [compile]", compile_failure.stderr)
            self.assertNotIn("ERROR [package]", compile_failure.stderr)
            self.assertEqual(before, tracked_status(), "release install test mutated checkout")


if __name__ == "__main__":
    unittest.main()
