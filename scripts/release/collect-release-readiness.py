#!/usr/bin/env python3

"""Collect exact-SHA readiness evidence from governed repository authorities.

What: Produce the evidence document consumed by verify-release-readiness.py.
Why: Release automation should not hand-author claims, toolchain versions, asset
digests, exclusions, or licensing state.
How: Join the governed source files with the already-gated workflow SHA and the
independently verified local release-asset manifest.
"""

from __future__ import annotations

import argparse
import json
import re
import sys
from pathlib import Path
from typing import Any

from readiness_common import compatibility_surface


ROOT = Path(__file__).resolve().parents[2]
POLICY_PATH = ROOT / "release" / "readiness-policy.json"
COMPATIBILITY_PATH = ROOT / "docs" / "compatibility-support-manifest.json"
LICENSE_PATH = ROOT / "license-policy.json"
SHA_RE = re.compile(r"[0-9a-f]{40}")
DIGEST_RE = re.compile(r"sha256:[0-9a-f]{64}")


class CollectionError(RuntimeError):
    """The repository and workflow inputs cannot form trustworthy evidence."""


def fail(message: str) -> None:
    raise CollectionError(message)


def load_object(path: Path, label: str) -> dict[str, Any]:
    try:
        value = json.loads(path.read_text(encoding="utf-8"))
    except (OSError, UnicodeDecodeError, json.JSONDecodeError) as error:
        fail(f"cannot read {label}: {error}")
    if not isinstance(value, dict):
        fail(f"{label} must be a JSON object")
    return value


def require_object(value: Any, label: str) -> dict[str, Any]:
    if not isinstance(value, dict):
        fail(f"{label} must be an object")
    return value


def require_list(value: Any, label: str) -> list[Any]:
    if not isinstance(value, list):
        fail(f"{label} must be a list")
    return value


def role_for_asset(name: str, version: str) -> str:
    expected = {
        f"reflaxe.go-{version}.zip": "haxelib-zip",
        f"reflaxe.go-{version}.zip.sha256": "checksum",
        f"reflaxe.go-{version}.manifest.json": "content-manifest",
        f"reflaxe.go-{version}.provenance.json": "provenance",
    }
    role = expected.get(name)
    if role is None:
        fail(f"release asset has no readiness role: {name}")
    return role


def collect(arguments: argparse.Namespace) -> dict[str, Any]:
    if SHA_RE.fullmatch(arguments.tested_sha) is None:
        fail("tested SHA must be exactly 40 lowercase hexadecimal characters")
    if arguments.tag != f"v{arguments.version}":
        fail("release tag does not match release version")

    policy = load_object(POLICY_PATH, "readiness policy")
    compatibility = load_object(COMPATIBILITY_PATH, "compatibility authority")
    licensing = load_object(LICENSE_PATH, "licensing authority")
    upstream = load_object(arguments.upstream_evidence, "upstream gate evidence")
    if upstream.get("schemaVersion") != 1:
        fail("upstream gate evidence schema is unsupported")
    if upstream.get("kind") != "haxe.go-upstream-gate-evidence":
        fail("upstream gate evidence kind is invalid")
    if upstream.get("testedSha") != arguments.tested_sha:
        fail("upstream evidence SHA does not match the tested release SHA")
    asset_manifest = load_object(arguments.assets, "release asset manifest")
    if asset_manifest.get("schemaVersion") != 1:
        fail("release asset manifest schema is unsupported")
    if asset_manifest.get("tag") != arguments.tag:
        fail("release asset manifest tag differs from candidate tag")
    if asset_manifest.get("sourceSha") != arguments.tested_sha:
        fail("release asset manifest SHA differs from tested SHA")

    assets: list[dict[str, str]] = []
    for raw in require_list(asset_manifest.get("assets"), "release assets"):
        asset = require_object(raw, "release asset")
        name = asset.get("name")
        digest = asset.get("digest")
        if not isinstance(name, str) or not name:
            fail("release asset name is invalid")
        if not isinstance(digest, str) or DIGEST_RE.fullmatch(digest) is None:
            fail(f"release asset digest is invalid: {name}")
        assets.append(
            {
                "role": role_for_asset(name, arguments.version),
                "name": name,
                "digest": digest,
            }
        )

    compatibility_policy = require_object(
        policy.get("compatibility"), "readiness compatibility policy"
    )
    exclusions = [
        {
            "scope": scope,
            "owner": owner,
            "advertisedAsSupported": False,
        }
        for scope, owner in sorted(
            require_object(
                compatibility_policy.get("requiredExclusions"),
                "required exclusions",
            ).items()
        )
    ]
    blockers = load_object(arguments.blocker_evidence, "blocker evidence")
    claim = require_object(
        compatibility.get("release_claim"), "compatibility release claim"
    )
    approval = require_object(licensing.get("approval"), "licensing approval")
    artifact_names = {asset["name"] for asset in assets}
    provenance_name = f"reflaxe.go-{arguments.version}.provenance.json"
    surface_digest, operation_count, symbol_count = compatibility_surface(compatibility)
    public_api = require_object(upstream.get("publicApi"), "upstream public API evidence")
    platform = require_object(upstream.get("platform"), "upstream platform evidence")
    toolchains = require_object(upstream.get("toolchains"), "upstream toolchain evidence")
    security = require_object(upstream.get("security"), "upstream security evidence")

    return {
        "schemaVersion": 1,
        "kind": "haxe.go-release-readiness-evidence",
        "phase": arguments.phase,
        "release": {
            "version": arguments.version,
            "tag": arguments.tag,
            "testedSha": arguments.tested_sha,
            "sourceSha": arguments.tested_sha,
            "manifestTag": asset_manifest["tag"],
        },
        "compatibility": {
            "lifecycle": claim.get("lifecycle"),
            "claim": claim.get("statement"),
            "admittedScopes": compatibility_policy.get("admittedScopes"),
            "evidencedScopes": compatibility_policy.get("admittedScopes"),
            "surfaceSha256": surface_digest,
            "admittedOperationCount": operation_count,
            "admittedSymbolCount": symbol_count,
            "exclusions": exclusions,
        },
        "publicApi": {
            "result": public_api.get("result"),
            "testedSha": arguments.tested_sha,
        },
        "platform": platform,
        "toolchains": toolchains,
        "security": security,
        "licensing": {
            "status": licensing.get("status"),
            "scopeSha256": approval.get("scopeSha256"),
            "unresolvedQuestions": licensing.get("unresolvedQuestions"),
        },
        "blockers": blockers,
        "artifacts": {
            "verifiedForSha": arguments.tested_sha,
            "assets": assets,
            "provenanceSubjects": sorted(artifact_names - {provenance_name}),
        },
        "github": None,
    }


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--phase", required=True, choices=("candidate", "published"))
    parser.add_argument("--version", required=True)
    parser.add_argument("--tag", required=True)
    parser.add_argument("--tested-sha", required=True)
    parser.add_argument("--upstream-evidence", required=True, type=Path)
    parser.add_argument("--blocker-evidence", required=True, type=Path)
    parser.add_argument("--assets", required=True, type=Path)
    parser.add_argument("--output", required=True, type=Path)
    return parser.parse_args()


def main() -> int:
    arguments = parse_args()
    try:
        evidence = collect(arguments)
        if arguments.output.exists():
            fail(f"evidence output already exists: {arguments.output}")
        arguments.output.parent.mkdir(parents=True, exist_ok=True)
        arguments.output.write_text(
            json.dumps(evidence, indent=2, sort_keys=True) + "\n",
            encoding="utf-8",
        )
    except (CollectionError, OSError) as error:
        print(f"[release-readiness-collect] ERROR: {error}", file=sys.stderr)
        return 2
    print(f"[release-readiness-collect] wrote {arguments.phase} evidence")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
