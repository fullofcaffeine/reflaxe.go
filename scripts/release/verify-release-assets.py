#!/usr/bin/env python3

"""Verify the complete local release-asset bundle before GitHub is mutated."""

from __future__ import annotations

import argparse
import hashlib
import json
from pathlib import Path
import re
import subprocess
import sys
import zipfile


VERSION_PATTERN = re.compile(
    r"(?:0|[1-9][0-9]*)\.(?:0|[1-9][0-9]*)\.(?:0|[1-9][0-9]*)"
)
SOURCE_SHA_PATTERN = re.compile(r"[0-9a-f]{40}")
SHA256_PATTERN = re.compile(r"sha256:[0-9a-f]{64}")
SOURCE_REPOSITORY = "git+https://github.com/fullofcaffeine/reflaxe.go"
ROOT = Path(__file__).resolve().parents[2]
HAXELIB_VERIFIER = ROOT / "scripts" / "release" / "verify-haxelib-artifact.py"
PACKAGE_MANIFEST = "reflaxe-package-manifest.json"


class AssetVerificationError(RuntimeError):
    """Raised when local release evidence is incomplete or contradictory."""


def fail(message: str) -> None:
    raise AssetVerificationError(message)


def sha256(path: Path) -> str:
    return hashlib.sha256(path.read_bytes()).hexdigest()


def load_json(path: Path, label: str) -> dict[str, object]:
    try:
        value = json.loads(path.read_text(encoding="utf-8"))
    except (OSError, UnicodeDecodeError, json.JSONDecodeError) as error:
        fail(f"cannot read {label}: {error}")
    if not isinstance(value, dict):
        fail(f"{label} must be a JSON object")
    return value


def require_exact_keys(
    value: dict[str, object], expected: set[str], label: str
) -> None:
    if set(value) != expected:
        fail(f"{label} fields must be exactly {sorted(expected)}")


def verify_identity(version: str, tag: str, source_sha: str) -> None:
    if VERSION_PATTERN.fullmatch(version) is None or version == "0.0.0":
        fail("release version must be canonical stable SemVer and not 0.0.0")
    if tag != f"v{version}":
        fail(f"release tag must be exactly v{version}")
    if SOURCE_SHA_PATTERN.fullmatch(source_sha) is None:
        fail("source SHA must be exactly 40 lowercase hexadecimal characters")


def regular_asset(root: Path, relative_name: str) -> Path:
    if not relative_name or Path(relative_name).name != relative_name:
        fail(f"release asset path must be one relative file name: {relative_name!r}")
    path = root / relative_name
    if path.is_symlink() or not path.is_file():
        fail(f"release asset must be a regular non-symlink file: {relative_name}")
    return path


def verify_content_manifest(
    path: Path, *, archive_name: str, version: str, tag: str, source_sha: str
) -> None:
    document = load_json(path, "content manifest")
    require_exact_keys(
        document,
        {
            "schemaVersion",
            "kind",
            "release",
            "artifact",
            "embeddedManifestSha256",
            "stagedReleaseIdentity",
            "stagedReleaseIdentitySha256",
            "reproducibility",
            "contents",
        },
        "content manifest",
    )
    if document.get("schemaVersion") != 1:
        fail("content manifest schemaVersion must be 1")
    if document.get("kind") != "haxe.go-haxelib-release-artifact":
        fail("content manifest kind is not the haxe.go release artifact contract")
    if document.get("release") != {
        "versionAuthority": "git-tag",
        "version": version,
        "tag": tag,
        "sourceCommit": source_sha,
    }:
        fail("content manifest release identity does not match the requested release")
    artifact = document.get("artifact")
    if not isinstance(artifact, dict):
        fail("content manifest artifact record is missing")
    archive = path.parent / archive_name
    if artifact != {
        "file": archive_name,
        "sha256": sha256(archive),
        "size": archive.stat().st_size,
    }:
        fail("content manifest does not describe the exact release ZIP")
    reproducibility = document.get("reproducibility")
    if not isinstance(reproducibility, dict):
        fail("content manifest reproducibility evidence is missing")
    if reproducibility != {
        "buildCount": 2,
        "byteIdentical": True,
        "embeddedManifestIdentical": True,
        "sourceExport": "git-archive",
    }:
        fail("content manifest does not prove two byte-identical builds")
    staged_identity = document.get("stagedReleaseIdentity")
    if not isinstance(staged_identity, dict):
        fail("content manifest staged release identity is missing")
    for field, expected in {
        "schema_version": 1,
        "kind": "haxe.go-release-metadata",
        "version_authority": "git-tag",
        "version": version,
        "tag": tag,
        "source_commit": source_sha,
    }.items():
        if staged_identity.get(field) != expected:
            fail(f"content manifest staged identity field {field} is inconsistent")

    try:
        with zipfile.ZipFile(archive) as package:
            embedded_bytes = package.read(PACKAGE_MANIFEST)
        embedded = json.loads(embedded_bytes)
    except (
        OSError,
        KeyError,
        UnicodeDecodeError,
        json.JSONDecodeError,
        zipfile.BadZipFile,
    ) as error:
        fail(f"cannot read the ZIP's embedded content manifest: {error}")
    embedded_digest = hashlib.sha256(embedded_bytes).hexdigest()
    if document.get("embeddedManifestSha256") != embedded_digest:
        fail("content manifest does not hash the ZIP's embedded content manifest")
    if document.get("contents") != embedded:
        fail("content manifest does not match the content manifest embedded in the ZIP")


def verify_haxelib_archive(
    archive: Path, *, version: str, tag: str, source_sha: str
) -> None:
    try:
        process = subprocess.run(
            [
                sys.executable,
                str(HAXELIB_VERIFIER),
                "--zip",
                str(archive),
                "--version",
                version,
                "--tag",
                tag,
                "--source-sha",
                source_sha,
            ],
            cwd=ROOT,
            capture_output=True,
            text=True,
            timeout=120,
        )
    except (OSError, subprocess.TimeoutExpired) as error:
        fail(f"cannot run the independent Haxelib ZIP verifier: {error}")
    if process.returncode != 0:
        detail = (process.stdout + process.stderr).strip()
        fail(f"independent Haxelib ZIP verification failed: {detail}")


def verify_provenance(
    path: Path,
    *,
    subject_paths: dict[str, Path],
    version: str,
    tag: str,
    source_sha: str,
) -> None:
    statement = load_json(path, "provenance statement")
    require_exact_keys(
        statement, {"_type", "subject", "predicateType", "predicate"}, "provenance"
    )
    if statement["_type"] != "https://in-toto.io/Statement/v1":
        fail("provenance must use the in-toto Statement v1 type")
    if statement["predicateType"] != "https://slsa.dev/provenance/v1":
        fail("provenance must use the SLSA provenance v1 predicate")
    subjects = statement.get("subject")
    if not isinstance(subjects, list):
        fail("provenance subjects must be an array")
    actual_subjects: dict[str, str] = {}
    for subject in subjects:
        if not isinstance(subject, dict):
            fail("every provenance subject must be an object")
        require_exact_keys(subject, {"name", "digest"}, "provenance subject")
        name = subject.get("name")
        digest = subject.get("digest")
        if not isinstance(name, str) or not isinstance(digest, dict):
            fail("provenance subject name and digest are invalid")
        require_exact_keys(digest, {"sha256"}, f"provenance subject {name} digest")
        if name in actual_subjects:
            fail(f"provenance contains duplicate subject {name}")
        value = digest.get("sha256")
        if not isinstance(value, str):
            fail(f"provenance subject {name} SHA-256 is invalid")
        actual_subjects[name] = value
    expected_subjects = {name: sha256(asset) for name, asset in subject_paths.items()}
    if actual_subjects != expected_subjects:
        fail("provenance subjects do not hash the ZIP, checksum, and content manifest")

    predicate = statement.get("predicate")
    if not isinstance(predicate, dict):
        fail("provenance predicate must be an object")
    require_exact_keys(predicate, {"buildDefinition", "runDetails"}, "provenance predicate")
    definition = predicate.get("buildDefinition")
    if not isinstance(definition, dict):
        fail("provenance buildDefinition must be an object")
    expected_builder = (
        "https://github.com/fullofcaffeine/reflaxe.go/blob/"
        f"{source_sha}/scripts/release/build-haxelib-artifact.py"
    )
    expected_build_type = (
        "https://github.com/fullofcaffeine/reflaxe.go/blob/"
        f"{source_sha}/docs/release-version-policy.md#deterministic-release-asset-build"
    )
    if definition.get("buildType") != expected_build_type:
        fail("provenance build type does not identify the exact source contract")
    if definition.get("externalParameters") != {
        "sourceCommit": source_sha,
        "tag": tag,
        "version": version,
    }:
        fail("provenance external parameters do not match release identity")
    if definition.get("internalParameters") != {
        "buildCount": 2,
        "sourceExport": "git-archive",
    }:
        fail("provenance internal build parameters are not the approved deterministic build")
    if definition.get("resolvedDependencies") != [
        {
            "uri": f"{SOURCE_REPOSITORY}@refs/tags/{tag}",
            "digest": {"gitCommit": source_sha},
        }
    ]:
        fail("provenance source dependency does not bind the exact tag and commit")
    if predicate.get("runDetails") != {
        "builder": {"id": expected_builder},
        "metadata": {"invocationId": f"{tag}@{source_sha}"},
    }:
        fail("provenance run details do not identify the approved builder invocation")


def verify_bundle(
    asset_manifest_path: Path, *, version: str, tag: str, source_sha: str
) -> dict[str, object]:
    verify_identity(version, tag, source_sha)
    if asset_manifest_path.is_symlink() or not asset_manifest_path.is_file():
        fail("release asset manifest must be a regular non-symlink file")
    document = load_json(asset_manifest_path, "release asset manifest")
    require_exact_keys(
        document,
        {"schemaVersion", "tag", "sourceSha", "assets"},
        "release asset manifest",
    )
    if document["schemaVersion"] != 1:
        fail("release asset manifest schemaVersion must be 1")
    if document["tag"] != tag or document["sourceSha"] != source_sha:
        fail("release asset manifest identity does not match the requested release")

    archive_name = f"reflaxe.go-{version}.zip"
    checksum_name = f"{archive_name}.sha256"
    manifest_name = f"reflaxe.go-{version}.manifest.json"
    provenance_name = f"reflaxe.go-{version}.provenance.json"
    expected_names = {archive_name, checksum_name, manifest_name, provenance_name}
    raw_assets = document.get("assets")
    if not isinstance(raw_assets, list) or len(raw_assets) != 4:
        fail("release asset manifest must contain exactly four hosted assets")
    root = asset_manifest_path.parent
    paths: dict[str, Path] = {}
    for record in raw_assets:
        if not isinstance(record, dict):
            fail("every release asset record must be an object")
        require_exact_keys(
            record,
            {"name", "path", "size", "digest"},
            "release asset record",
        )
        name = record.get("name")
        relative_path = record.get("path")
        if not isinstance(name, str) or relative_path != name:
            fail("release asset name and path must be the same safe file name")
        if name in paths:
            fail(f"release asset manifest contains duplicate name {name}")
        path = regular_asset(root, name)
        size = record.get("size")
        digest = record.get("digest")
        if not isinstance(size, int) or isinstance(size, bool) or size < 0:
            fail(f"release asset {name} size must be a non-negative integer")
        if not isinstance(digest, str) or SHA256_PATTERN.fullmatch(digest) is None:
            fail(f"release asset {name} digest must use lowercase sha256:<64 hex>")
        if size != path.stat().st_size:
            fail(f"release asset {name} size does not match local bytes")
        if digest != f"sha256:{sha256(path)}":
            fail(f"release asset {name} digest does not match local bytes")
        paths[name] = path
    if set(paths) != expected_names:
        fail(f"release assets must be exactly {sorted(expected_names)}")

    expected_checksum = f"{sha256(paths[archive_name])}  {archive_name}\n"
    try:
        actual_checksum = paths[checksum_name].read_text(encoding="ascii")
    except (OSError, UnicodeDecodeError) as error:
        fail(f"cannot read checksum sidecar: {error}")
    if actual_checksum != expected_checksum:
        fail("checksum sidecar does not hash the exact release ZIP")
    verify_haxelib_archive(
        paths[archive_name], version=version, tag=tag, source_sha=source_sha
    )
    verify_content_manifest(
        paths[manifest_name],
        archive_name=archive_name,
        version=version,
        tag=tag,
        source_sha=source_sha,
    )
    verify_provenance(
        paths[provenance_name],
        subject_paths={
            archive_name: paths[archive_name],
            checksum_name: paths[checksum_name],
            manifest_name: paths[manifest_name],
        },
        version=version,
        tag=tag,
        source_sha=source_sha,
    )
    return {"schemaVersion": 1, "tag": tag, "sourceSha": source_sha, "assets": 4}


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--assets", required=True, type=Path)
    parser.add_argument("--version", required=True)
    parser.add_argument("--tag", required=True)
    parser.add_argument("--source-sha", required=True)
    return parser.parse_args()


def main() -> int:
    args = parse_args()
    try:
        result = verify_bundle(
            args.assets,
            version=args.version,
            tag=args.tag,
            source_sha=args.source_sha,
        )
    except (AssetVerificationError, OSError) as error:
        print(f"[release-assets] ERROR: {error}", file=sys.stderr)
        return 2
    print(json.dumps(result, sort_keys=True))
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
