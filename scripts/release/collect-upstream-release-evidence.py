#!/usr/bin/env python3

"""Turn GitHub job results and resolved tools into exact-SHA gate evidence.

What: Record the upstream results that authorize the release job to continue.
Why: The readiness collector must consume structured results rather than infer
that every check passed merely because it received a commit SHA.
How: Accept GitHub `needs.*.result` values plus resolved Haxe/Node versions,
validate them against policy, and expand matrix jobs into their exact lanes.
"""

from __future__ import annotations

import argparse
import json
import re
import sys
from pathlib import Path
from typing import Any


ROOT = Path(__file__).resolve().parents[2]
READINESS_POLICY = ROOT / "release" / "readiness-policy.json"
TOOLCHAIN_POLICY = ROOT / "docs" / "toolchain-policy.json"
SHA_RE = re.compile(r"[0-9a-f]{40}")
VERSION_RE = re.compile(
    r"(?:0|[1-9][0-9]*)\.(?:0|[1-9][0-9]*)\.(?:0|[1-9][0-9]*)"
)
RUNNER_IMAGE_RE = re.compile(r"[A-Za-z0-9._-]+")


class UpstreamEvidenceError(RuntimeError):
    """Workflow results or resolved tools cannot authorize a release."""


def fail(message: str) -> None:
    raise UpstreamEvidenceError(message)


def load_object(path: Path, label: str) -> dict[str, Any]:
    try:
        value = json.loads(path.read_text(encoding="utf-8"))
    except (OSError, UnicodeDecodeError, json.JSONDecodeError) as error:
        fail(f"cannot read {label}: {error}")
    if not isinstance(value, dict):
        fail(f"{label} must be a JSON object")
    return value


def normalized_version(value: str, label: str) -> str:
    normalized = value.removeprefix("v")
    if VERSION_RE.fullmatch(normalized) is None:
        fail(f"resolved {label} version must be exact X.Y.Z")
    return normalized


def collect(arguments: argparse.Namespace) -> dict[str, Any]:
    if SHA_RE.fullmatch(arguments.tested_sha) is None:
        fail("tested SHA must be exactly 40 lowercase hexadecimal characters")
    results = {
        "quality": arguments.quality_result,
        "gitleaks": arguments.gitleaks_result,
        "dependency-audit": arguments.dependency_audit_result,
        "go-tooling": arguments.go_tooling_result,
        "github-governance-live": arguments.github_governance_result,
    }
    failed = sorted(name for name, result in results.items() if result != "success")
    if failed:
        fail(f"upstream workflow gate did not succeed: {', '.join(failed)}")

    readiness = load_object(READINESS_POLICY, "readiness policy")
    toolchains = load_object(TOOLCHAIN_POLICY, "toolchain policy")
    haxe_version = normalized_version(arguments.haxe_version, "Haxe")
    node_version = normalized_version(arguments.node_version, "Node")
    if haxe_version != toolchains.get("haxe", {}).get("ci_selector"):
        fail("resolved Haxe version differs from toolchain policy")
    node_lines = toolchains.get("node", {}).get("supported_tooling_lines")
    if (
        not isinstance(node_lines, list)
        or node_version.split(".", 1)[0] not in node_lines
    ):
        fail("resolved Node version is outside supported tooling lines")
    go_versions = toolchains.get("go", {}).get("ci_versions")
    if not isinstance(go_versions, list) or not go_versions:
        fail("toolchain policy has no exact Go CI versions")
    if (
        RUNNER_IMAGE_RE.fullmatch(arguments.runner_image_os) is None
        or RUNNER_IMAGE_RE.fullmatch(arguments.runner_image_version) is None
    ):
        fail("exact hosted runner image OS and version are required")

    gate_ids = readiness.get("requiredSecurityGates")
    if not isinstance(gate_ids, list) or not gate_ids:
        fail("readiness policy has no required security gates")
    expected_ids = {
        "gitleaks",
        "github-governance-live",
        *(f"dependency-audit:go{version}" for version in go_versions),
        *(f"go-tooling:go{version}" for version in go_versions),
    }
    if set(gate_ids) != expected_ids:
        fail("required security gates do not match supported Go matrix")

    return {
        "schemaVersion": 1,
        "kind": "haxe.go-upstream-gate-evidence",
        "testedSha": arguments.tested_sha,
        "publicApi": {"result": "pass"},
        "platform": {
            "id": "linux-amd64",
            "os": "linux",
            "architecture": "amd64",
            "runnerImage": {
                "os": arguments.runner_image_os,
                "version": arguments.runner_image_version,
            },
        },
        "toolchains": {
            "haxe": {"resolved": [haxe_version]},
            "go": {"resolved": go_versions},
            "node": {"resolved": [node_version]},
        },
        "security": {
            "reachableVulnerabilities": [],
            "gates": [
                {
                    "id": gate_id,
                    "result": "pass",
                    "testedSha": arguments.tested_sha,
                }
                for gate_id in sorted(gate_ids)
            ],
        },
    }


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--tested-sha", required=True)
    parser.add_argument("--quality-result", required=True)
    parser.add_argument("--gitleaks-result", required=True)
    parser.add_argument("--dependency-audit-result", required=True)
    parser.add_argument("--go-tooling-result", required=True)
    parser.add_argument("--github-governance-result", required=True)
    parser.add_argument("--haxe-version", required=True)
    parser.add_argument("--node-version", required=True)
    parser.add_argument("--runner-image-os", required=True)
    parser.add_argument("--runner-image-version", required=True)
    parser.add_argument("--output", required=True, type=Path)
    return parser.parse_args()


def main() -> int:
    arguments = parse_args()
    try:
        evidence = collect(arguments)
        if arguments.output.exists():
            fail(f"upstream evidence output already exists: {arguments.output}")
        arguments.output.parent.mkdir(parents=True, exist_ok=True)
        arguments.output.write_text(
            json.dumps(evidence, indent=2, sort_keys=True) + "\n",
            encoding="utf-8",
        )
    except (UpstreamEvidenceError, OSError) as error:
        print(f"[upstream-release-evidence] ERROR: {error}", file=sys.stderr)
        return 2
    print("[upstream-release-evidence] exact-SHA workflow evidence: OK")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
