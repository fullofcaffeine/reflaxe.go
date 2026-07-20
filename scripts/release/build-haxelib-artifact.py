#!/usr/bin/env python3

"""Build a release Haxelib ZIP twice from one exact Git commit and emit evidence."""

from __future__ import annotations

import argparse
from dataclasses import dataclass
import hashlib
import io
import json
import os
from pathlib import Path, PurePosixPath
import re
import shutil
import subprocess
import sys
import tarfile
import tempfile
import zipfile


ROOT = Path(__file__).resolve().parents[2]
VERIFIER = ROOT / "scripts" / "release" / "verify-haxelib-artifact.py"
ASSET_VERIFIER = ROOT / "scripts" / "release" / "verify-release-assets.py"
PACKAGE_MANIFEST = "reflaxe-package-manifest.json"
SOURCE_REPOSITORY = "git+https://github.com/fullofcaffeine/reflaxe.go"
VERSION_PATTERN = re.compile(
    r"(?:0|[1-9][0-9]*)\.(?:0|[1-9][0-9]*)\.(?:0|[1-9][0-9]*)"
)
SOURCE_SHA_PATTERN = re.compile(r"[0-9a-f]{40}")


class ArtifactBuildError(RuntimeError):
    """Raised when exact-source artifact evidence cannot be produced safely."""


@dataclass(frozen=True)
class ReleaseIdentity:
    version: str
    tag: str
    source_sha: str
    release_note: str | None


@dataclass(frozen=True)
class CleanBuild:
    archive_path: Path
    archive_bytes: bytes
    embedded_manifest_bytes: bytes
    embedded_manifest: dict[str, object]
    release_identity_bytes: bytes
    release_identity: dict[str, object]


def sha256_bytes(value: bytes) -> str:
    return hashlib.sha256(value).hexdigest()


def run_checked(
    command: list[str], *, cwd: Path, environment: dict[str, str] | None = None
) -> subprocess.CompletedProcess[bytes]:
    process = subprocess.run(
        command,
        cwd=cwd,
        env=environment,
        capture_output=True,
        timeout=180,
    )
    if process.returncode != 0:
        stdout = process.stdout.decode("utf-8", errors="replace")
        stderr = process.stderr.decode("utf-8", errors="replace")
        raise ArtifactBuildError(
            f"command failed ({' '.join(command)}):\n{stdout}{stderr}"
        )
    return process


def validate_identity(
    version: str, tag: str, source_sha: str, output_dir: Path, release_note: str | None
) -> ReleaseIdentity:
    if VERSION_PATTERN.fullmatch(version) is None or version == "0.0.0":
        raise ArtifactBuildError(
            "release version must be canonical stable SemVer and not the 0.0.0 sentinel"
        )
    if tag != f"v{version}":
        raise ArtifactBuildError(f"release tag must be exactly v{version}")
    if SOURCE_SHA_PATTERN.fullmatch(source_sha) is None:
        raise ArtifactBuildError(
            "source SHA must be exactly 40 lowercase hexadecimal characters"
        )
    if release_note is not None and not release_note:
        raise ArtifactBuildError("release note must not be empty")
    if output_dir.exists():
        raise ArtifactBuildError(f"output directory must not already exist: {output_dir}")

    resolved = subprocess.run(
        ["git", "rev-parse", "--verify", f"{source_sha}^{{commit}}"],
        cwd=ROOT,
        capture_output=True,
        text=True,
    )
    if resolved.returncode != 0 or resolved.stdout.strip() != source_sha:
        raise ArtifactBuildError(
            f"source SHA does not resolve to the exact commit in this repository: {source_sha}"
        )

    tag_exists = subprocess.run(
        ["git", "show-ref", "--verify", "--quiet", f"refs/tags/{tag}"],
        cwd=ROOT,
    ).returncode == 0
    if tag_exists:
        tag_commit = run_checked(
            ["git", "rev-parse", "--verify", f"refs/tags/{tag}^{{commit}}"],
            cwd=ROOT,
        ).stdout.decode("ascii").strip()
        if tag_commit != source_sha:
            raise ArtifactBuildError(
                f"existing release tag {tag} resolves to {tag_commit}, not {source_sha}"
            )

    return ReleaseIdentity(version, tag, source_sha, release_note)


def safe_archive_member(value: str) -> PurePosixPath:
    normalized = value.rstrip("/")
    path = PurePosixPath(normalized)
    if (
        not normalized
        or "\\" in normalized
        or path.is_absolute()
        or path.as_posix() != normalized
        or any(segment in {"", ".", ".."} for segment in path.parts)
    ):
        raise ArtifactBuildError(f"git archive contains an unsafe member: {value!r}")
    return path


def export_commit(source_sha: str, destination: Path) -> None:
    archive_process = run_checked(
        ["git", "archive", "--format=tar", source_sha], cwd=ROOT
    )
    destination.mkdir(parents=True)
    try:
        with tarfile.open(fileobj=io.BytesIO(archive_process.stdout), mode="r:") as archive:
            for member in sorted(archive.getmembers(), key=lambda item: item.name.encode("utf-8")):
                relative = safe_archive_member(member.name)
                output = destination.joinpath(*relative.parts)
                if member.isdir():
                    output.mkdir(parents=True, exist_ok=True)
                    continue
                if not member.isfile():
                    raise ArtifactBuildError(
                        f"git archive member is not a regular file or directory: {member.name}"
                    )
                source = archive.extractfile(member)
                if source is None:
                    raise ArtifactBuildError(
                        f"git archive member could not be read: {member.name}"
                    )
                output.parent.mkdir(parents=True, exist_ok=True)
                output.write_bytes(source.read())
    except tarfile.TarError as error:
        raise ArtifactBuildError(f"cannot extract exact Git source archive: {error}") from error


def deterministic_environment() -> dict[str, str]:
    environment = dict(os.environ)
    environment.update({"LC_ALL": "C", "PYTHONHASHSEED": "0", "TZ": "UTC"})
    return environment


def parse_json_bytes(value: bytes, label: str) -> dict[str, object]:
    try:
        parsed = json.loads(value)
    except (json.JSONDecodeError, UnicodeDecodeError) as error:
        raise ArtifactBuildError(f"{label} is not valid JSON") from error
    if not isinstance(parsed, dict):
        raise ArtifactBuildError(f"{label} must contain a JSON object")
    return parsed


def validate_staged_release_identity(
    metadata_root: Path,
    staged: dict[str, object],
    identity: ReleaseIdentity,
) -> None:
    expected_fields = {
        "schema_version": 1,
        "kind": "haxe.go-release-metadata",
        "version_authority": "git-tag",
        "version": identity.version,
        "tag": identity.tag,
        "source_commit": identity.source_sha,
    }
    for field, expected in expected_fields.items():
        if staged.get(field) != expected:
            raise ArtifactBuildError(
                f"staged release identity field {field!r} does not match {expected!r}"
            )
    files = staged.get("files")
    if not isinstance(files, dict) or set(files) != {"haxelib.json", "package.json"}:
        raise ArtifactBuildError(
            "staged release identity must hash haxelib.json and package.json exactly"
        )
    for name in sorted(files):
        record = files[name]
        if not isinstance(record, dict):
            raise ArtifactBuildError(f"staged release identity record is invalid: {name}")
        expected_hash = record.get("sha256")
        actual_hash = sha256_bytes((metadata_root / name).read_bytes())
        if expected_hash != actual_hash:
            raise ArtifactBuildError(
                f"staged release identity hash differs for {name}"
            )


def verify_archive(archive: Path, identity: ReleaseIdentity) -> None:
    run_checked(
        [
            sys.executable,
            str(VERIFIER),
            "--zip",
            str(archive),
            "--version",
            identity.version,
            "--tag",
            identity.tag,
            "--source-sha",
            identity.source_sha,
        ],
        cwd=ROOT,
        environment=deterministic_environment(),
    )


def build_once(workspace: Path, identity: ReleaseIdentity, index: int) -> CleanBuild:
    build_root = workspace / f"build-{index}"
    source_root = build_root / "source"
    metadata_root = build_root / "metadata"
    archive_path = build_root / "reflaxe.go.zip"
    export_commit(identity.source_sha, source_root)

    stager = source_root / "scripts" / "release" / "stage-release-metadata.py"
    package_script = source_root / "scripts" / "release" / "package-haxelib.sh"
    for required in (stager, package_script):
        if not required.is_file():
            raise ArtifactBuildError(
                f"exact source commit lacks required release builder: {required.relative_to(source_root)}"
            )

    stage_command = [
        sys.executable,
        str(stager),
        "--version",
        identity.version,
        "--source-sha",
        identity.source_sha,
        "--output-dir",
        str(metadata_root),
    ]
    if identity.release_note is not None:
        stage_command.extend(["--release-note", identity.release_note])
    run_checked(
        stage_command,
        cwd=source_root,
        environment=deterministic_environment(),
    )
    shutil.copyfile(metadata_root / "haxelib.json", source_root / "haxelib.json")
    shutil.copyfile(metadata_root / "package.json", source_root / "package.json")

    run_checked(
        ["bash", str(package_script), str(archive_path), str(source_root)],
        cwd=source_root,
        environment=deterministic_environment(),
    )
    verify_archive(archive_path, identity)

    archive_bytes = archive_path.read_bytes()
    try:
        with zipfile.ZipFile(archive_path) as archive:
            embedded_manifest_bytes = archive.read(PACKAGE_MANIFEST)
    except (zipfile.BadZipFile, KeyError) as error:
        raise ArtifactBuildError(
            "built archive does not contain its package manifest"
        ) from error
    release_identity_bytes = (metadata_root / "release-identity.json").read_bytes()
    staged_release_identity = parse_json_bytes(
        release_identity_bytes, "staged release identity"
    )
    validate_staged_release_identity(metadata_root, staged_release_identity, identity)
    return CleanBuild(
        archive_path=archive_path,
        archive_bytes=archive_bytes,
        embedded_manifest_bytes=embedded_manifest_bytes,
        embedded_manifest=parse_json_bytes(
            embedded_manifest_bytes, "embedded package manifest"
        ),
        release_identity_bytes=release_identity_bytes,
        release_identity=staged_release_identity,
    )


def build_artifacts(identity: ReleaseIdentity, output_dir: Path) -> dict[str, Path]:
    build_count = 2
    with tempfile.TemporaryDirectory(prefix="haxe-go-release-build-") as raw:
        workspace = Path(raw)
        builds = [build_once(workspace, identity, index + 1) for index in range(build_count)]
        first, second = builds
        if first.archive_bytes != second.archive_bytes:
            raise ArtifactBuildError(
                "two clean Git-archive builds produced different ZIP bytes "
                f"({sha256_bytes(first.archive_bytes)} != {sha256_bytes(second.archive_bytes)})"
            )
        if first.embedded_manifest_bytes != second.embedded_manifest_bytes:
            raise ArtifactBuildError(
                "two clean builds produced different embedded content manifests"
            )
        if first.release_identity_bytes != second.release_identity_bytes:
            raise ArtifactBuildError(
                "two clean builds produced different staged release identities"
            )

        archive_name = f"reflaxe.go-{identity.version}.zip"
        manifest_name = f"reflaxe.go-{identity.version}.manifest.json"
        checksum_name = f"{archive_name}.sha256"
        provenance_name = f"reflaxe.go-{identity.version}.provenance.json"
        archive_digest = sha256_bytes(first.archive_bytes)
        artifact_manifest = {
            "schemaVersion": 1,
            "kind": "haxe.go-haxelib-release-artifact",
            "release": {
                "versionAuthority": "git-tag",
                "version": identity.version,
                "tag": identity.tag,
                "sourceCommit": identity.source_sha,
            },
            "artifact": {
                "file": archive_name,
                "sha256": archive_digest,
                "size": len(first.archive_bytes),
            },
            "embeddedManifestSha256": sha256_bytes(
                first.embedded_manifest_bytes
            ),
            "stagedReleaseIdentity": first.release_identity,
            "stagedReleaseIdentitySha256": sha256_bytes(
                first.release_identity_bytes
            ),
            "reproducibility": {
                "buildCount": build_count,
                "byteIdentical": True,
                "embeddedManifestIdentical": True,
                "sourceExport": "git-archive",
            },
            "contents": first.embedded_manifest,
        }
        manifest_bytes = (
            json.dumps(artifact_manifest, indent=2, sort_keys=True) + "\n"
        ).encode("utf-8")
        checksum_bytes = f"{archive_digest}  {archive_name}\n".encode("ascii")
        builder_uri = (
            "https://github.com/fullofcaffeine/reflaxe.go/blob/"
            f"{identity.source_sha}/scripts/release/build-haxelib-artifact.py"
        )
        provenance = {
            "_type": "https://in-toto.io/Statement/v1",
            "subject": [
                {
                    "name": name,
                    "digest": {"sha256": sha256_bytes(contents)},
                }
                for name, contents in (
                    (archive_name, first.archive_bytes),
                    (checksum_name, checksum_bytes),
                    (manifest_name, manifest_bytes),
                )
            ],
            "predicateType": "https://slsa.dev/provenance/v1",
            "predicate": {
                "buildDefinition": {
                    "buildType": (
                        "https://github.com/fullofcaffeine/reflaxe.go/blob/"
                        f"{identity.source_sha}/docs/release-version-policy.md"
                        "#deterministic-release-asset-build"
                    ),
                    "externalParameters": {
                        "sourceCommit": identity.source_sha,
                        "tag": identity.tag,
                        "version": identity.version,
                    },
                    "internalParameters": {
                        "buildCount": build_count,
                        "sourceExport": "git-archive",
                    },
                    "resolvedDependencies": [
                        {
                            "uri": f"{SOURCE_REPOSITORY}@refs/tags/{identity.tag}",
                            "digest": {"gitCommit": identity.source_sha},
                        }
                    ],
                },
                "runDetails": {
                    "builder": {"id": builder_uri},
                    "metadata": {
                        "invocationId": f"{identity.tag}@{identity.source_sha}"
                    },
                },
            },
        }
        provenance_bytes = (
            json.dumps(provenance, indent=2, sort_keys=True) + "\n"
        ).encode("utf-8")
        hosted_assets = (
            (archive_name, first.archive_bytes),
            (checksum_name, checksum_bytes),
            (manifest_name, manifest_bytes),
            (provenance_name, provenance_bytes),
        )
        asset_manifest = {
            "schemaVersion": 1,
            "tag": identity.tag,
            "sourceSha": identity.source_sha,
            "assets": [
                {
                    "name": name,
                    "path": name,
                    "size": len(contents),
                    "digest": f"sha256:{sha256_bytes(contents)}",
                }
                for name, contents in hosted_assets
            ],
        }
        asset_manifest_bytes = (
            json.dumps(asset_manifest, indent=2, sort_keys=True) + "\n"
        ).encode("utf-8")

        output_dir.mkdir(parents=True)
        archive_output = output_dir / archive_name
        manifest_output = output_dir / manifest_name
        checksum_output = output_dir / checksum_name
        provenance_output = output_dir / provenance_name
        asset_manifest_output = output_dir / "release-assets.json"
        try:
            archive_output.write_bytes(first.archive_bytes)
            manifest_output.write_bytes(manifest_bytes)
            checksum_output.write_bytes(checksum_bytes)
            provenance_output.write_bytes(provenance_bytes)
            asset_manifest_output.write_bytes(asset_manifest_bytes)
            verify_archive(archive_output, identity)
            run_checked(
                [
                    sys.executable,
                    str(ASSET_VERIFIER),
                    "--assets",
                    str(asset_manifest_output),
                    "--version",
                    identity.version,
                    "--tag",
                    identity.tag,
                    "--source-sha",
                    identity.source_sha,
                ],
                cwd=ROOT,
                environment=deterministic_environment(),
            )
        except Exception:
            shutil.rmtree(output_dir, ignore_errors=True)
            raise
        return {
            "archive": archive_output,
            "checksum": checksum_output,
            "manifest": manifest_output,
            "provenance": provenance_output,
            "asset_manifest": asset_manifest_output,
        }


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--version", required=True)
    parser.add_argument("--tag", required=True)
    parser.add_argument("--source-sha", required=True)
    parser.add_argument("--output-dir", required=True, type=Path)
    parser.add_argument("--release-note")
    return parser.parse_args()


def main() -> int:
    args = parse_args()
    try:
        identity = validate_identity(
            args.version,
            args.tag,
            args.source_sha,
            args.output_dir,
            args.release_note,
        )
        outputs = build_artifacts(identity, args.output_dir)
    except (ArtifactBuildError, OSError, ValueError) as error:
        print(f"[haxelib-build] ERROR: {error}", file=sys.stderr)
        return 2
    print(
        json.dumps(
            {name: str(path) for name, path in sorted(outputs.items())},
            sort_keys=True,
        )
    )
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
