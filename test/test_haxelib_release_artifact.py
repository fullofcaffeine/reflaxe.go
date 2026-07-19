#!/usr/bin/env python3

from __future__ import annotations

import hashlib
import json
from pathlib import Path
import re
import shutil
import subprocess
import sys
import tempfile
import unittest
import zipfile


ROOT = Path(__file__).resolve().parent.parent
BUILDER = ROOT / "scripts" / "release" / "build-haxelib-artifact.py"
VERIFIER = ROOT / "scripts" / "release" / "verify-haxelib-artifact.py"
POLICY_DOC = ROOT / "docs" / "release-version-policy.md"
VERSION = "0.999999.0"
TAG = f"v{VERSION}"
FIXED_ZIP_TIMESTAMP = (2000, 1, 1, 0, 0, 0)
POSIX_HOME_PATH = re.compile(r"(?<![A-Za-z0-9])/(?:Users|home)/[^\s\"'`]+")
WINDOWS_HOME_PATH = re.compile(r"(?i)(?<![A-Za-z0-9])[A-Z]:[\\/]+Users[\\/]+[^\s\"'`]+")


def sha256(path: Path) -> str:
    return hashlib.sha256(path.read_bytes()).hexdigest()


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


def run_builder(output_dir: Path, *extra: str) -> subprocess.CompletedProcess[str]:
    return subprocess.run(
        [
            sys.executable,
            str(BUILDER),
            "--version",
            VERSION,
            "--tag",
            TAG,
            "--source-sha",
            head_sha(),
            "--output-dir",
            str(output_dir),
            *extra,
        ],
        cwd=ROOT,
        capture_output=True,
        text=True,
        timeout=180,
    )


def run_verifier(
    archive: Path,
    *,
    version: str = VERSION,
    tag: str = TAG,
    source_sha: str | None = None,
) -> subprocess.CompletedProcess[str]:
    return subprocess.run(
        [
            sys.executable,
            str(VERIFIER),
            "--zip",
            str(archive),
            "--version",
            version,
            "--tag",
            tag,
            "--source-sha",
            source_sha or head_sha(),
        ],
        cwd=ROOT,
        capture_output=True,
        text=True,
        timeout=120,
    )


class HaxelibReleaseArtifactContractTest(unittest.TestCase):
    maxDiff = None

    def test_release_artifact_entrypoints_are_release_blocking(self) -> None:
        self.assertTrue(BUILDER.is_file())
        self.assertTrue(VERIFIER.is_file())
        package = json.loads((ROOT / "package.json").read_text(encoding="utf-8"))
        scripts = package.get("scripts", {})
        self.assertEqual(
            "python3 scripts/release/build-haxelib-artifact.py",
            scripts.get("release:build-haxelib"),
        )
        self.assertEqual(
            "python3 scripts/release/verify-haxelib-artifact.py",
            scripts.get("release:verify-haxelib"),
        )
        contract = "python3 test/test_haxelib_release_artifact.py"
        self.assertEqual(contract, scripts.get("test:haxelib-artifact"))
        release_contracts = (ROOT / "test" / "run-release-contracts.py").read_text(
            encoding="utf-8"
        )
        self.assertIn('"test/test_haxelib_release_artifact.py"', release_contracts)

        policy = POLICY_DOC.read_text(encoding="utf-8")
        for phrase in (
            "release:build-haxelib",
            "two independent `git archive` exports",
            "reflaxe.go-<version>.zip.sha256",
            "existing tag",
        ):
            self.assertIn(phrase, policy)

        builder = BUILDER.read_text(encoding="utf-8")
        self.assertIn("git", builder)
        self.assertIn("archive", builder)
        self.assertIn("build_count", builder)
        self.assertNotIn("shutil.copytree(ROOT", builder)

    @unittest.skipUnless(
        shutil.which("git") and shutil.which("haxe"),
        "requires Git and Haxe",
    )
    def test_two_clean_exact_commit_builds_emit_verified_release_evidence(self) -> None:
        with tempfile.TemporaryDirectory(prefix="haxe-go-release-artifact-") as raw:
            temp = Path(raw)
            output_dir = temp / "artifacts"
            before = tracked_status()
            process = run_builder(output_dir)
            after = tracked_status()
            self.assertEqual(0, process.returncode, process.stdout + process.stderr)
            self.assertEqual(before, after, "artifact build mutated the tracked checkout")

            archive_name = f"reflaxe.go-{VERSION}.zip"
            manifest_name = f"reflaxe.go-{VERSION}.manifest.json"
            checksum_name = f"{archive_name}.sha256"
            self.assertEqual(
                {archive_name, manifest_name, checksum_name},
                {path.name for path in output_dir.iterdir()},
            )
            archive = output_dir / archive_name
            manifest_path = output_dir / manifest_name
            checksum_path = output_dir / checksum_name

            manifest = json.loads(manifest_path.read_text(encoding="utf-8"))
            self.assertEqual(1, manifest.get("schemaVersion"))
            self.assertEqual("haxe.go-haxelib-release-artifact", manifest.get("kind"))
            self.assertEqual(
                {
                    "sourceCommit": head_sha(),
                    "tag": TAG,
                    "version": VERSION,
                    "versionAuthority": "git-tag",
                },
                manifest.get("release"),
            )
            self.assertEqual(
                {
                    "buildCount": 2,
                    "byteIdentical": True,
                    "embeddedManifestIdentical": True,
                    "sourceExport": "git-archive",
                },
                manifest.get("reproducibility"),
            )
            self.assertEqual(archive_name, manifest["artifact"]["file"])
            self.assertEqual(archive.stat().st_size, manifest["artifact"]["size"])
            self.assertEqual(sha256(archive), manifest["artifact"]["sha256"])
            self.assertRegex(
                manifest.get("embeddedManifestSha256", ""),
                r"^[0-9a-f]{64}$",
            )
            self.assertEqual(TAG, manifest["stagedReleaseIdentity"]["tag"])
            self.assertEqual(
                head_sha(),
                manifest["stagedReleaseIdentity"]["source_commit"],
            )
            self.assertRegex(
                manifest.get("stagedReleaseIdentitySha256", ""),
                r"^[0-9a-f]{64}$",
            )
            self.assertEqual(394, len(manifest["contents"]["entries"]))
            packaged_sources = {
                entry["sourcePath"] for entry in manifest["contents"]["entries"]
            }
            self.assertTrue(
                {
                    "runtime/hxrt/reflect.go",
                    "src/reflaxe/go/compiler/emit/GoGeneratedFieldMetadataEmitter.hx",
                    "src/reflaxe/go/compiler/emit/GoReflectMetadataEmitter.hx",
                    "std/go/_std/Reflect.hx",
                    "std/go/_std/Std.hx",
                    "std/go/_std/haxe/Log.hx",
                    "std/hxrt/reflect/NativeReflect.hx",
                    "std/hxrt/reflect/ReflectFieldLookup.hx",
                    "std/reflaxe/go/internal/CompilerReflect.hx",
                }.issubset(packaged_sources)
            )

            manifest_text = manifest_path.read_text(encoding="utf-8")
            self.assertNotIn(str(ROOT), manifest_text)
            self.assertNotIn(str(temp), manifest_text)
            self.assertIsNone(POSIX_HOME_PATH.search(manifest_text))
            self.assertIsNone(WINDOWS_HOME_PATH.search(manifest_text))

            self.assertEqual(
                f"{sha256(archive)}  {archive_name}\n",
                checksum_path.read_text(encoding="utf-8"),
            )

            with zipfile.ZipFile(archive) as package:
                infos = package.infolist()
                names = [info.filename for info in infos]
                self.assertEqual(sorted(names, key=lambda value: value.encode("utf-8")), names)
                self.assertEqual(len(names), len(set(names)))
                self.assertNotIn("std", {name.split("/", 1)[0] for name in names})
                self.assertFalse(any("__pycache__" in name or "node_modules" in name for name in names))
                for info in infos:
                    self.assertEqual(FIXED_ZIP_TIMESTAMP, info.date_time)
                    self.assertEqual(0o644, (info.external_attr >> 16) & 0o777)
                    self.assertEqual(zipfile.ZIP_STORED, info.compress_type)

                haxelib = json.loads(package.read("haxelib.json"))
                self.assertEqual(VERSION, haxelib.get("version"))
                self.assertIn(TAG, haxelib.get("releasenote", ""))
                self.assertIn(head_sha(), haxelib.get("releasenote", ""))
                self.assertNotIn("reflaxe", haxelib)
                embedded = json.loads(package.read("reflaxe-package-manifest.json"))
                self.assertEqual(manifest["contents"], embedded)

                cross_entries = [
                    entry
                    for entry in embedded["entries"]
                    if entry["packagePath"].endswith(".cross.hx")
                ]
                self.assertEqual(107, len(cross_entries))
                self.assertTrue(
                    all(
                        entry["kind"] == "stdlib-override"
                        and entry["sourcePath"].startswith("std/go/_std/")
                        for entry in cross_entries
                    )
                )

                for info in infos:
                    if Path(info.filename).suffix.lower() not in {
                        ".go",
                        ".hxml",
                        ".hx",
                        ".json",
                        ".md",
                    }:
                        continue
                    text = package.read(info).decode("utf-8", errors="replace")
                    if info.filename.startswith("vendor/"):
                        continue
                    self.assertIsNone(POSIX_HOME_PATH.search(text), info.filename)
                    self.assertIsNone(WINDOWS_HOME_PATH.search(text), info.filename)

            verify = run_verifier(archive)
            self.assertEqual(0, verify.returncode, verify.stdout + verify.stderr)
            summary = json.loads(verify.stdout)
            self.assertEqual(sha256(archive), summary["sha256"])
            self.assertEqual(395, summary["entries"])

            wrong_version = run_verifier(
                archive,
                version="0.999998.0",
                tag="v0.999998.0",
            )
            self.assertNotEqual(0, wrong_version.returncode)
            self.assertIn("version", wrong_version.stderr)

            tampered = temp / "tampered.zip"
            shutil.copyfile(archive, tampered)
            with zipfile.ZipFile(tampered, mode="a") as package:
                package.writestr("private/secret.txt", "not declared\n")
            rejected = run_verifier(tampered)
            self.assertNotEqual(0, rejected.returncode)
            self.assertRegex(rejected.stderr, r"canonical|unexpected|manifest")

    @unittest.skipUnless(shutil.which("git"), "requires Git")
    def test_builder_rejects_identity_and_output_drift_before_writing(self) -> None:
        with tempfile.TemporaryDirectory(prefix="haxe-go-release-artifact-reject-") as raw:
            temp = Path(raw)
            wrong_tag = subprocess.run(
                [
                    sys.executable,
                    str(BUILDER),
                    "--version",
                    VERSION,
                    "--tag",
                    "v0.1.0",
                    "--source-sha",
                    head_sha(),
                    "--output-dir",
                    str(temp / "wrong-tag"),
                ],
                cwd=ROOT,
                capture_output=True,
                text=True,
            )
            self.assertNotEqual(0, wrong_tag.returncode)
            self.assertIn("tag must be exactly", wrong_tag.stderr)
            self.assertFalse((temp / "wrong-tag").exists())

            invalid_sha = subprocess.run(
                [
                    sys.executable,
                    str(BUILDER),
                    "--version",
                    VERSION,
                    "--tag",
                    TAG,
                    "--source-sha",
                    "a" * 40,
                    "--output-dir",
                    str(temp / "invalid-sha"),
                ],
                cwd=ROOT,
                capture_output=True,
                text=True,
            )
            self.assertNotEqual(0, invalid_sha.returncode)
            self.assertIn("does not resolve to the exact commit", invalid_sha.stderr)
            self.assertFalse((temp / "invalid-sha").exists())

            existing = temp / "existing"
            existing.mkdir()
            output_drift = run_builder(existing)
            self.assertNotEqual(0, output_drift.returncode)
            self.assertIn("must not already exist", output_drift.stderr)

            historical_tags = subprocess.run(
                ["git", "tag", "--merged", "HEAD", "--list", "v*"],
                cwd=ROOT,
                check=True,
                capture_output=True,
                text=True,
            ).stdout.splitlines()
            for historical_tag in historical_tags:
                if (
                    re.fullmatch(
                        r"v(?:0|[1-9][0-9]*)\.(?:0|[1-9][0-9]*)\."
                        r"(?:0|[1-9][0-9]*)",
                        historical_tag,
                    )
                    is None
                ):
                    continue
                historical_sha = subprocess.run(
                    ["git", "rev-parse", f"{historical_tag}^{{commit}}"],
                    cwd=ROOT,
                    check=True,
                    capture_output=True,
                    text=True,
                ).stdout.strip()
                if historical_sha == head_sha():
                    continue
                mismatched_tag = subprocess.run(
                    [
                        sys.executable,
                        str(BUILDER),
                        "--version",
                        historical_tag.removeprefix("v"),
                        "--tag",
                        historical_tag,
                        "--source-sha",
                        head_sha(),
                        "--output-dir",
                        str(temp / "existing-tag-mismatch"),
                    ],
                    cwd=ROOT,
                    capture_output=True,
                    text=True,
                )
                self.assertNotEqual(0, mismatched_tag.returncode)
                self.assertIn("existing release tag", mismatched_tag.stderr)
                self.assertFalse((temp / "existing-tag-mismatch").exists())
                break


if __name__ == "__main__":
    unittest.main()
