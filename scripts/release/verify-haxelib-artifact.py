#!/usr/bin/env python3

"""Verify one Reflaxe.Go Haxelib ZIP against release identity and layout policy."""

from __future__ import annotations

import argparse
import hashlib
import json
from pathlib import Path, PurePosixPath
import re
import stat
import subprocess
import sys
import tempfile
import zipfile


ROOT = Path(__file__).resolve().parents[2]
ZIP_WRITER = ROOT / "scripts" / "release" / "deterministic-zip.py"
PACKAGE_MANIFEST = "reflaxe-package-manifest.json"
FIXED_TIMESTAMP = (2000, 1, 1, 0, 0, 0)
EXPECTED_ARCHIVE_POLICY = {
    "compression": "stored",
    "fileMode": "0644",
    "ordering": "utf8-bytewise",
    "timestamp": "2000-01-01T00:00:00Z",
}
ALLOWED_ROOT_FILES = {
    "LICENSE",
    "README.md",
    "Run.hx",
    "extraParams.hxml",
    "haxelib.json",
    PACKAGE_MANIFEST,
}
ALLOWED_ROOT_DIRECTORIES = {"runtime", "src", "vendor"}
FORBIDDEN_SEGMENTS = {
    ".cache",
    ".git",
    "__pycache__",
    "benches",
    "node_modules",
    "target",
    "test",
    "tests",
}
EXPECTED_KINDS = {
    "class-path",
    "metadata",
    "package-runner",
    "runtime",
    "stdlib",
    "stdlib-override",
    "vendored-reflaxe",
}
TEXT_SUFFIXES = {".go", ".hxml", ".hx", ".json", ".md", ".txt"}
VERSION_PATTERN = re.compile(
    r"(?:0|[1-9][0-9]*)\.(?:0|[1-9][0-9]*)\.(?:0|[1-9][0-9]*)"
)
SOURCE_SHA_PATTERN = re.compile(r"[0-9a-f]{40}")
SHA256_PATTERN = re.compile(r"[0-9a-f]{64}")
POSIX_HOME_PATH = re.compile(r"(?<![A-Za-z0-9])/(?:Users|home)/[^\s\"'`]+")
WINDOWS_HOME_PATH = re.compile(
    r"(?i)(?<![A-Za-z0-9])[A-Z]:[\\/]+Users[\\/]+[^\s\"'`]+"
)


class ArtifactVerificationError(RuntimeError):
    """Raised when a candidate archive cannot be trusted as a release artifact."""


def utf8_key(value: str) -> bytes:
    return value.encode("utf-8")


def sha256_bytes(value: bytes) -> str:
    return hashlib.sha256(value).hexdigest()


def safe_relative_path(value: object, label: str) -> str:
    if not isinstance(value, str) or not value or "\\" in value or "\0" in value:
        raise ArtifactVerificationError(f"{label} is not a safe relative path")
    path = PurePosixPath(value)
    if (
        path.is_absolute()
        or path.as_posix() != value
        or any(segment in {"", ".", ".."} for segment in path.parts)
        or (len(value) >= 2 and value[0].isalpha() and value[1] == ":")
    ):
        raise ArtifactVerificationError(f"{label} is not a safe relative path: {value!r}")
    return value


def validate_identity(version: str, tag: str, source_sha: str) -> None:
    if VERSION_PATTERN.fullmatch(version) is None or version == "0.0.0":
        raise ArtifactVerificationError("release version must be canonical non-sentinel SemVer")
    if tag != f"v{version}":
        raise ArtifactVerificationError(f"release tag must be exactly v{version}")
    if SOURCE_SHA_PATTERN.fullmatch(source_sha) is None:
        raise ArtifactVerificationError("source SHA must be 40 lowercase hexadecimal characters")


def read_archive(archive_path: Path) -> tuple[bytes, list[zipfile.ZipInfo], dict[str, bytes]]:
    try:
        archive_bytes = archive_path.read_bytes()
        with zipfile.ZipFile(archive_path) as archive:
            if archive.comment:
                raise ArtifactVerificationError("canonical ZIP must not have an archive comment")
            infos = archive.infolist()
            names = [info.filename for info in infos]
            for name in names:
                safe_relative_path(name, "archive member")
            if len(names) != len(set(names)):
                raise ArtifactVerificationError("canonical ZIP must not contain duplicate members")
            if names != sorted(names, key=utf8_key):
                raise ArtifactVerificationError("canonical ZIP members are not in UTF-8 byte order")

            files: dict[str, bytes] = {}
            for info in infos:
                if info.is_dir():
                    raise ArtifactVerificationError(
                        f"canonical ZIP must not contain directory entries: {info.filename}"
                    )
                unix_mode = (info.external_attr >> 16) & 0xFFFF
                if stat.S_IFMT(unix_mode) != stat.S_IFREG:
                    raise ArtifactVerificationError(
                        f"canonical ZIP member is not a regular file: {info.filename}"
                    )
                if unix_mode & 0o777 != 0o644:
                    raise ArtifactVerificationError(
                        f"canonical ZIP member mode must be 0644: {info.filename}"
                    )
                if info.date_time != FIXED_TIMESTAMP:
                    raise ArtifactVerificationError(
                        f"canonical ZIP timestamp differs: {info.filename}"
                    )
                if info.compress_type != zipfile.ZIP_STORED:
                    raise ArtifactVerificationError(
                        f"canonical ZIP compression must be stored: {info.filename}"
                    )
                if info.extra or info.comment:
                    raise ArtifactVerificationError(
                        f"canonical ZIP member metadata is not empty: {info.filename}"
                    )
                files[info.filename] = archive.read(info)
    except (OSError, zipfile.BadZipFile) as error:
        raise ArtifactVerificationError(f"cannot read Haxelib ZIP: {error}") from error
    return archive_bytes, infos, files


def validate_layout(files: dict[str, bytes]) -> None:
    names = set(files)
    for required in (
        "haxelib.json",
        PACKAGE_MANIFEST,
        "runtime/hxrt/core.go",
        "src/reflaxe/go/CompilerInit.hx",
        "vendor/reflaxe/src/reflaxe/ReflectCompiler.hx",
    ):
        if required not in names:
            raise ArtifactVerificationError(f"required package member is missing: {required}")

    for name in names:
        parts = PurePosixPath(name).parts
        if len(parts) == 1:
            if name not in ALLOWED_ROOT_FILES:
                raise ArtifactVerificationError(f"unexpected top-level package member: {name}")
        elif parts[0] not in ALLOWED_ROOT_DIRECTORIES:
            raise ArtifactVerificationError(f"unexpected package root: {parts[0]}")
        if any(segment in FORBIDDEN_SEGMENTS for segment in parts) or ".DS_Store" in parts:
            raise ArtifactVerificationError(f"development debris is not allowed: {name}")


def parse_json(files: dict[str, bytes], name: str) -> dict[str, object]:
    try:
        value = json.loads(files[name])
    except (KeyError, json.JSONDecodeError, UnicodeDecodeError) as error:
        raise ArtifactVerificationError(f"package member is not valid JSON: {name}") from error
    if not isinstance(value, dict):
        raise ArtifactVerificationError(f"package JSON member must contain an object: {name}")
    return value


def validate_haxelib(
    files: dict[str, bytes], version: str, tag: str, source_sha: str
) -> None:
    haxelib = parse_json(files, "haxelib.json")
    if haxelib.get("version") != version:
        raise ArtifactVerificationError(
            f"packaged Haxelib version {haxelib.get('version')!r} does not match {version}"
        )
    if haxelib.get("classPath") != "src" or "reflaxe" in haxelib:
        raise ArtifactVerificationError(
            "packaged Haxelib metadata must use classPath=src and omit source-only reflaxe metadata"
        )
    release_note = haxelib.get("releasenote")
    if (
        not isinstance(release_note, str)
        or tag not in release_note
        or source_sha not in release_note
    ):
        raise ArtifactVerificationError(
            "packaged Haxelib release note does not bind the tag and tested source commit"
        )


def validate_package_manifest(files: dict[str, bytes]) -> dict[str, object]:
    manifest = parse_json(files, PACKAGE_MANIFEST)
    if (
        manifest.get("schemaVersion") != 1
        or manifest.get("format") != "reflaxe.go-haxelib-package"
        or manifest.get("classPath") != "src"
        or manifest.get("archive") != EXPECTED_ARCHIVE_POLICY
    ):
        raise ArtifactVerificationError("embedded package manifest header is not canonical")
    entries = manifest.get("entries")
    if not isinstance(entries, list):
        raise ArtifactVerificationError("embedded package manifest entries must be an array")

    package_paths: list[str] = []
    for index, raw_entry in enumerate(entries):
        if not isinstance(raw_entry, dict):
            raise ArtifactVerificationError(f"embedded manifest entry {index} is not an object")
        source_path = safe_relative_path(raw_entry.get("sourcePath"), "manifest source path")
        package_path = safe_relative_path(raw_entry.get("packagePath"), "manifest package path")
        package_paths.append(package_path)
        kind = raw_entry.get("kind")
        if kind not in EXPECTED_KINDS:
            raise ArtifactVerificationError(f"embedded manifest kind is unknown: {package_path}")
        source_digest = raw_entry.get("sourceSha256")
        package_digest = raw_entry.get("packageSha256")
        size = raw_entry.get("size")
        if not isinstance(source_digest, str) or SHA256_PATTERN.fullmatch(source_digest) is None:
            raise ArtifactVerificationError(f"embedded source hash is invalid: {source_path}")
        content = files.get(package_path)
        if content is None:
            raise ArtifactVerificationError(f"embedded manifest member is missing: {package_path}")
        if not isinstance(package_digest, str) or sha256_bytes(content) != package_digest:
            raise ArtifactVerificationError(f"embedded manifest package hash differs: {package_path}")
        if not isinstance(size, int) or isinstance(size, bool) or len(content) != size:
            raise ArtifactVerificationError(f"embedded manifest package size differs: {package_path}")

        is_cross = package_path.endswith(".cross.hx")
        if is_cross:
            if kind != "stdlib-override" or not source_path.startswith("std/go/_std/"):
                raise ArtifactVerificationError(
                    f"generated .cross.hx lacks canonical override provenance: {package_path}"
                )
            source_relative = PurePosixPath(source_path).relative_to("std/go/_std")
            expected = (PurePosixPath("src") / source_relative.with_suffix(".cross.hx")).as_posix()
            if package_path != expected:
                raise ArtifactVerificationError(
                    f"generated .cross.hx path does not match its source: {package_path}"
                )
        elif kind == "stdlib-override":
            raise ArtifactVerificationError(
                f"stdlib override was not generated as .cross.hx: {package_path}"
            )
        if kind == "stdlib" and (not package_path.endswith(".hx") or is_cross):
            raise ArtifactVerificationError(
                f"ordinary std/support source changed extension: {package_path}"
            )

    if package_paths != sorted(package_paths, key=utf8_key):
        raise ArtifactVerificationError("embedded manifest entries are not sorted")
    if len(package_paths) != len(set(package_paths)):
        raise ArtifactVerificationError("embedded manifest package paths are not unique")
    if set(package_paths) != set(files) - {PACKAGE_MANIFEST}:
        raise ArtifactVerificationError(
            "embedded manifest does not cover the archive member set exactly"
        )
    for package_path in package_paths:
        if package_path.endswith(".cross.hx"):
            plain = package_path.removesuffix(".cross.hx") + ".hx"
            if plain in files:
                raise ArtifactVerificationError(
                    f"generated override has a plain package counterpart: {plain}"
                )
    return manifest


def validate_no_local_paths(files: dict[str, bytes]) -> None:
    for name, content in files.items():
        if PurePosixPath(name).suffix.lower() not in TEXT_SUFFIXES:
            continue
        text = content.decode("utf-8", errors="replace")
        if name.startswith("vendor/"):
            continue
        if POSIX_HOME_PATH.search(text) or WINDOWS_HOME_PATH.search(text):
            raise ArtifactVerificationError(
                f"package member contains an absolute machine-local path: {name}"
            )


def validate_canonical_bytes(archive_path: Path, original: bytes, files: dict[str, bytes]) -> None:
    with tempfile.TemporaryDirectory(prefix="haxe-go-verify-zip-") as raw:
        temp = Path(raw)
        stage = temp / "stage"
        for name in sorted(files, key=utf8_key):
            destination = stage.joinpath(*PurePosixPath(name).parts)
            destination.parent.mkdir(parents=True, exist_ok=True)
            destination.write_bytes(files[name])
        rebuilt = temp / "canonical.zip"
        process = subprocess.run(
            [sys.executable, str(ZIP_WRITER), str(stage), str(rebuilt)],
            cwd=ROOT,
            capture_output=True,
            text=True,
            timeout=120,
        )
        if process.returncode != 0:
            raise ArtifactVerificationError(
                "cannot rebuild canonical ZIP: " + process.stdout + process.stderr
            )
        if rebuilt.read_bytes() != original:
            raise ArtifactVerificationError(
                f"archive is not in the canonical ZIP representation: {archive_path.name}"
            )


def verify_artifact(
    archive_path: Path, version: str, tag: str, source_sha: str
) -> dict[str, object]:
    validate_identity(version, tag, source_sha)
    archive_bytes, infos, files = read_archive(archive_path)
    validate_layout(files)
    validate_haxelib(files, version, tag, source_sha)
    manifest = validate_package_manifest(files)
    validate_no_local_paths(files)
    validate_canonical_bytes(archive_path, archive_bytes, files)
    return {
        "entries": len(infos),
        "packageManifestSha256": sha256_bytes(files[PACKAGE_MANIFEST]),
        "sha256": sha256_bytes(archive_bytes),
        "size": len(archive_bytes),
        "stdlibOverrides": sum(
            1
            for entry in manifest["entries"]
            if isinstance(entry, dict) and entry.get("kind") == "stdlib-override"
        ),
    }


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--zip", required=True, type=Path)
    parser.add_argument("--version", required=True)
    parser.add_argument("--tag", required=True)
    parser.add_argument("--source-sha", required=True)
    return parser.parse_args()


def main() -> int:
    args = parse_args()
    try:
        summary = verify_artifact(args.zip, args.version, args.tag, args.source_sha)
    except (ArtifactVerificationError, OSError, ValueError) as error:
        print(f"[haxelib-artifact] ERROR: {error}", file=sys.stderr)
        return 2
    print(json.dumps(summary, sort_keys=True))
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
