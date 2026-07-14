#!/usr/bin/env python3

"""Stage release metadata without modifying the tested source checkout."""

from __future__ import annotations

import argparse
import hashlib
import json
import re
import sys
from pathlib import Path


ROOT = Path(__file__).resolve().parents[2]
PACKAGE_PATH = ROOT / "package.json"
HAXELIB_PATH = ROOT / "haxelib.json"
DEVELOPMENT_VERSION = "0.0.0"
VERSION_PATTERN = re.compile(
    r"(?:0|[1-9][0-9]*)\.(?:0|[1-9][0-9]*)\.(?:0|[1-9][0-9]*)"
)
SOURCE_SHA_PATTERN = re.compile(r"[0-9a-f]{40}")


class StagingError(RuntimeError):
    """Raised when release metadata cannot be staged safely."""


def sha256(path: Path) -> str:
    return hashlib.sha256(path.read_bytes()).hexdigest()


def load_json(path: Path) -> dict[str, object]:
    try:
        value = json.loads(path.read_text(encoding="utf-8"))
    except (OSError, json.JSONDecodeError) as exc:
        raise StagingError(f"cannot read {path.name}: {exc}") from exc
    if not isinstance(value, dict):
        raise StagingError(f"{path.name} must contain a JSON object")
    return value


def write_json(path: Path, value: dict[str, object]) -> None:
    path.write_text(json.dumps(value, indent=2) + "\n", encoding="utf-8")


def validate_sources(
    package: dict[str, object],
    haxelib: dict[str, object],
) -> None:
    if package.get("version") != DEVELOPMENT_VERSION:
        raise StagingError(
            f"package.json must use the {DEVELOPMENT_VERSION} development sentinel"
        )
    if haxelib.get("version") != DEVELOPMENT_VERSION:
        raise StagingError(
            f"haxelib.json must use the {DEVELOPMENT_VERSION} development sentinel"
        )
    if package.get("name") != haxelib.get("name"):
        raise StagingError("package.json and haxelib.json package names differ")
    if package.get("license") != haxelib.get("license"):
        raise StagingError("package.json and haxelib.json licenses differ")


def parse_arguments() -> argparse.Namespace:
    parser = argparse.ArgumentParser(
        description=(
            "Write versioned package and Haxelib metadata into a new staging "
            "directory while leaving the tested checkout unchanged."
        )
    )
    parser.add_argument("--version", required=True)
    parser.add_argument("--source-sha", required=True)
    parser.add_argument("--output-dir", required=True, type=Path)
    parser.add_argument("--release-note")
    return parser.parse_args()


def stage(arguments: argparse.Namespace) -> Path:
    version = arguments.version
    source_sha = arguments.source_sha
    output_dir: Path = arguments.output_dir

    if VERSION_PATTERN.fullmatch(version) is None:
        raise StagingError(
            f"version {version!r} must be canonical stable SemVer X.Y.Z"
        )
    if version == DEVELOPMENT_VERSION:
        raise StagingError(
            f"the {DEVELOPMENT_VERSION} development sentinel cannot be staged as a release"
        )
    if SOURCE_SHA_PATTERN.fullmatch(source_sha) is None:
        raise StagingError("source SHA must be exactly 40 lowercase hexadecimal characters")
    if output_dir.exists():
        raise StagingError(f"output directory already exists: {output_dir}")

    package = load_json(PACKAGE_PATH)
    haxelib = load_json(HAXELIB_PATH)
    validate_sources(package, haxelib)

    staged_package = dict(package)
    staged_haxelib = dict(haxelib)
    staged_package["version"] = version
    staged_haxelib["version"] = version
    staged_haxelib["releasenote"] = (
        arguments.release_note
        if arguments.release_note is not None
        else f"Release v{version} from tested commit {source_sha}"
    )
    if not staged_haxelib["releasenote"]:
        raise StagingError("release note must not be empty")

    output_dir.mkdir(parents=True)
    package_output = output_dir / "package.json"
    haxelib_output = output_dir / "haxelib.json"
    write_json(package_output, staged_package)
    write_json(haxelib_output, staged_haxelib)

    identity = {
        "schema_version": 1,
        "kind": "haxe.go-release-metadata",
        "version_authority": "git-tag",
        "version": version,
        "tag": f"v{version}",
        "source_commit": source_sha,
        "files": {
            "haxelib.json": {"sha256": sha256(haxelib_output)},
            "package.json": {"sha256": sha256(package_output)},
        },
    }
    identity_output = output_dir / "release-identity.json"
    write_json(identity_output, identity)
    return identity_output


def main() -> int:
    try:
        identity_output = stage(parse_arguments())
    except (StagingError, OSError) as exc:
        print(f"[release-metadata] error: {exc}", file=sys.stderr)
        return 1

    print(f"[release-metadata] staged release identity: {identity_output}")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
