#!/usr/bin/env python3

"""Create the canonical, metadata-normalized Reflaxe.Go Haxelib ZIP."""

from __future__ import annotations

import argparse
import hashlib
import json
import os
from pathlib import Path, PurePosixPath
import stat
import sys
import tempfile
import zipfile


MANIFEST_NAME = "reflaxe-package-manifest.json"
EXPECTED_ARCHIVE_POLICY = {
    "compression": "stored",
    "fileMode": "0644",
    "ordering": "utf8-bytewise",
    "timestamp": "2000-01-01T00:00:00Z",
}
FIXED_TIMESTAMP = (2000, 1, 1, 0, 0, 0)
NORMALIZED_MODE = 0o644


class ArchiveError(RuntimeError):
    """Raised when input cannot produce the canonical package archive."""


def utf8_key(value: str) -> bytes:
    return value.encode("utf-8")


def validate_member_name(name: str) -> None:
    path = PurePosixPath(name)
    if (
        not name
        or "\0" in name
        or "\\" in name
        or path.is_absolute()
        or name.endswith("/")
        or path.as_posix() != name
        or any(segment in {"", ".", ".."} for segment in path.parts)
        or (len(name) >= 2 and name[0].isalpha() and name[1] == ":")
    ):
        raise ArchiveError(f"unsafe archive member path: {name!r}")


def collect_files(root: Path) -> list[tuple[str, Path]]:
    files: list[tuple[str, Path]] = []

    def visit(directory: Path, segments: tuple[str, ...]) -> None:
        with os.scandir(directory) as iterator:
            children = sorted(iterator, key=lambda entry: utf8_key(entry.name))
        for child in children:
            relative_segments = (*segments, child.name)
            relative = "/".join(relative_segments)
            if child.is_symlink():
                raise ArchiveError(f"symbolic links are not allowed in package archives: {relative}")
            if child.is_dir(follow_symlinks=False):
                visit(Path(child.path), relative_segments)
            elif child.is_file(follow_symlinks=False):
                validate_member_name(relative)
                files.append((relative, Path(child.path)))
            else:
                raise ArchiveError(f"special files are not allowed in package archives: {relative}")

    visit(root, ())
    files.sort(key=lambda item: utf8_key(item[0]))
    names = [name for name, _ in files]
    if len(names) != len(set(names)):
        raise ArchiveError("package archive contains duplicate member paths")
    return files


def validate_manifest(root: Path, files: list[tuple[str, Path]]) -> None:
    path = root / MANIFEST_NAME
    try:
        manifest = json.loads(path.read_text(encoding="utf-8"))
    except FileNotFoundError as error:
        raise ArchiveError(f"package manifest is missing: {MANIFEST_NAME}") from error
    except json.JSONDecodeError as error:
        raise ArchiveError(f"package manifest is not valid JSON: {error}") from error
    if (
        not isinstance(manifest, dict)
        or manifest.get("schemaVersion") != 1
        or manifest.get("format") != "reflaxe.go-haxelib-package"
        or manifest.get("archive") != EXPECTED_ARCHIVE_POLICY
    ):
        raise ArchiveError("package manifest archive policy does not match the deterministic ZIP writer")

    entries = manifest.get("entries")
    if not isinstance(entries, list):
        raise ArchiveError("package manifest entries must be an array")
    files_by_name = dict(files)
    package_paths: list[str] = []
    for index, entry in enumerate(entries):
        if not isinstance(entry, dict):
            raise ArchiveError(f"package manifest entry {index} is not an object")
        package_path = entry.get("packagePath")
        if not isinstance(package_path, str):
            raise ArchiveError(f"package manifest entry {index} has no package path")
        validate_member_name(package_path)
        if package_path == MANIFEST_NAME:
            raise ArchiveError("package manifest cannot contain a self-hash entry")
        package_paths.append(package_path)
        package_file = files_by_name.get(package_path)
        if package_file is None:
            raise ArchiveError(f"package manifest entry is missing from the stage: {package_path}")
        content = package_file.read_bytes()
        expected_hash = entry.get("packageSha256")
        if not isinstance(expected_hash, str) or hashlib.sha256(content).hexdigest() != expected_hash:
            raise ArchiveError(f"package hash differs from manifest: {package_path}")
        expected_size = entry.get("size")
        if not isinstance(expected_size, int) or isinstance(expected_size, bool) or len(content) != expected_size:
            raise ArchiveError(f"package size differs from manifest: {package_path}")

    if package_paths != sorted(package_paths, key=utf8_key):
        raise ArchiveError("package manifest entries are not sorted by UTF-8 path")
    if len(package_paths) != len(set(package_paths)):
        raise ArchiveError("package manifest contains duplicate package paths")
    actual_paths = set(files_by_name) - {MANIFEST_NAME}
    if set(package_paths) != actual_paths:
        raise ArchiveError("package manifest does not cover the staged file set exactly")


def create_deterministic_zip(source: Path, output: Path) -> None:
    root = source.resolve(strict=True)
    if not root.is_dir():
        raise ArchiveError(f"ZIP source is not a directory: {source}")
    output = output.resolve()
    if output == root or root in output.parents:
        raise ArchiveError("ZIP output must be outside the staged package directory")

    files = collect_files(root)
    validate_manifest(root, files)
    output.parent.mkdir(parents=True, exist_ok=True)
    temporary_path: Path | None = None
    try:
        with tempfile.NamedTemporaryFile(
            prefix=f".{output.name}.", suffix=".tmp", dir=output.parent, delete=False
        ) as temporary:
            temporary_path = Path(temporary.name)
        with zipfile.ZipFile(temporary_path, mode="w", compression=zipfile.ZIP_STORED) as archive:
            for name, path in files:
                info = zipfile.ZipInfo(filename=name, date_time=FIXED_TIMESTAMP)
                info.create_system = 3
                info.compress_type = zipfile.ZIP_STORED
                info.external_attr = (stat.S_IFREG | NORMALIZED_MODE) << 16
                info.internal_attr = 0
                info.extra = b""
                info.comment = b""
                archive.writestr(info, path.read_bytes())
        os.replace(temporary_path, output)
        temporary_path = None
    finally:
        if temporary_path is not None:
            temporary_path.unlink(missing_ok=True)


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("source", type=Path, help="staged package directory")
    parser.add_argument("output", type=Path, help="output ZIP path")
    return parser.parse_args()


def main() -> int:
    args = parse_args()
    try:
        create_deterministic_zip(args.source, args.output)
    except (ArchiveError, OSError) as error:
        print(f"[deterministic-zip] ERROR: {error}", file=sys.stderr)
        return 2
    print(f"[deterministic-zip] wrote {args.output}")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
