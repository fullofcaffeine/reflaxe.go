#!/usr/bin/env python3

"""Evaluate one exact release candidate or published release fail-closed.

What: Join release identity, admitted compatibility scope, API, security,
licensing, blocker, artifact, and hosted GitHub evidence into one verdict.
Why: Individually green checks can still describe different commits or a
broader public claim than the evidence proves.
How: Read a versioned policy and an immutable evidence document, require exact
SHA joins, and reject missing, contradictory, unsupported, or unowned truth.
"""

from __future__ import annotations

import argparse
import hashlib
import json
import re
import subprocess
import sys
from pathlib import Path
from typing import Any

from readiness_common import compatibility_surface


SHA_RE = re.compile(r"[0-9a-f]{40}")
DIGEST_RE = re.compile(r"sha256:[0-9a-f]{64}")
VERSION_RE = re.compile(r"(?:0|[1-9][0-9]*)\.(?:0|[1-9][0-9]*)\.(?:0|[1-9][0-9]*)")
OPEN_STATUSES = {"open", "in_progress", "blocked"}
EVIDENCE_FIELDS = {
    "schemaVersion",
    "kind",
    "phase",
    "release",
    "compatibility",
    "publicApi",
    "platform",
    "toolchains",
    "security",
    "licensing",
    "blockers",
    "artifacts",
    "github",
}
ROOT = Path(__file__).resolve().parents[2]


class ReadinessError(RuntimeError):
    """The evidence cannot support the requested release verdict."""


def fail(message: str) -> None:
    raise ReadinessError(message)


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


def unique_strings(value: Any, label: str) -> list[str]:
    items = require_list(value, label)
    if any(not isinstance(item, str) or not item for item in items):
        fail(f"{label} must contain non-empty strings")
    if len(items) != len(set(items)):
        fail(f"{label} must not contain duplicates")
    return items


def verify_policy(policy: dict[str, Any]) -> None:
    if policy.get("schemaVersion") != 2:
        fail("unsupported readiness policy schema")
    if policy.get("kind") != "haxe.go-release-readiness-policy":
        fail("readiness policy kind is invalid")
    repository = policy.get("repository")
    if not isinstance(repository, str) or "/" not in repository:
        fail("readiness policy repository must be owner/name")
    required = (
        "compatibility",
        "releaseLineAdmission",
        "toolchains",
        "requiredSecurityGates",
        "requiredArtifactRoles",
        "authorities",
    )
    for field in required:
        if field not in policy:
            fail(f"readiness policy is missing {field}")

    admission = require_object(
        policy.get("releaseLineAdmission"), "release-line admission policy"
    )
    if admission.get("mode") != "historical-beta-baseline":
        fail("release-line admission mode is unsupported")
    baseline_tag = admission.get("baselineTag")
    if not isinstance(baseline_tag, str) or re.fullmatch(
        r"v(?:0|[1-9][0-9]*)\.(?:0|[1-9][0-9]*)\.(?:0|[1-9][0-9]*)",
        baseline_tag,
    ) is None:
        fail("release-line admission baseline tag is invalid")
    baseline_sha = admission.get("baselineSourceSha")
    if not isinstance(baseline_sha, str) or SHA_RE.fullmatch(baseline_sha) is None:
        fail("release-line admission baseline SHA is invalid")
    if admission.get("routineReleaseProof") != "current-exact-sha-ci-and-authorities":
        fail("routine release proof is unsupported")
    required_triggers = {
        "compatibility-claim-or-admitted-scope",
        "security-or-trust-boundary",
        "public-api-or-abi",
        "licensing-or-distribution-rights",
        "release-policy-or-publication-authority",
        "stable-major-graduation",
    }
    if set(
        unique_strings(admission.get("freshReviewTriggers"), "fresh review triggers")
    ) != required_triggers:
        fail("fresh review triggers are incomplete or unknown")
    try:
        resolved = subprocess.run(
            ["git", "rev-parse", "--verify", f"{baseline_tag}^{{commit}}"],
            cwd=ROOT,
            text=True,
            capture_output=True,
            timeout=10,
        )
    except (OSError, subprocess.TimeoutExpired) as error:
        fail(f"cannot resolve release-line admission baseline tag: {error}")
    if resolved.returncode != 0 or resolved.stdout.strip() != baseline_sha:
        fail("release-line admission baseline tag and SHA differ")


def verify_release_line_ancestry(
    policy: dict[str, Any], tested_sha: str, *, mode: str
) -> None:
    """Require a real release candidate to remain on the admitted beta history."""

    if mode != "live":
        return
    admission = require_object(
        policy.get("releaseLineAdmission"), "release-line admission policy"
    )
    baseline_sha = admission.get("baselineSourceSha")
    try:
        ancestry = subprocess.run(
            ["git", "merge-base", "--is-ancestor", baseline_sha, tested_sha],
            cwd=ROOT,
            text=True,
            capture_output=True,
            timeout=10,
        )
    except (OSError, subprocess.TimeoutExpired) as error:
        fail(f"cannot verify release-line ancestry: {error}")
    if ancestry.returncode == 1:
        fail("release SHA does not descend from the historical beta baseline")
    if ancestry.returncode != 0:
        detail = (ancestry.stderr or ancestry.stdout).strip()
        fail(f"cannot verify release-line ancestry: {detail}")


def verify_repository_authorities(
    policy: dict[str, Any], evidence: dict[str, Any]
) -> None:
    compatibility = load_object(
        ROOT / "docs" / "compatibility-support-manifest.json",
        "compatibility authority",
    )
    claim = require_object(
        compatibility.get("release_claim"), "compatibility release claim"
    )
    expected = require_object(policy["compatibility"], "compatibility policy")
    if expected.get("lifecycle") != claim.get("lifecycle"):
        fail("readiness lifecycle drifted from compatibility authority")
    if expected.get("claim") != claim.get("statement"):
        fail("readiness claim drifted from compatibility authority")
    authority_scopes = {
        f"preset:{claim.get('admitted_preset')}",
        f"platform:{claim.get('admitted_platform')}",
    }
    if set(expected.get("admittedScopes", [])) != authority_scopes:
        fail("readiness admitted scopes drifted from compatibility authority")
    known_owner_ids = {
        blocker.get("id")
        for blocker in require_list(
            compatibility.get("known_blockers"), "known compatibility blockers"
        )
        if isinstance(blocker, dict)
    }
    exclusion_owners = set(
        require_object(
            expected.get("requiredExclusions"), "required exclusions"
        ).values()
    )
    blocker_scopes = require_object(
        expected.get("blockerScopes"), "compatibility blocker scopes"
    )
    blocker_ids = set(blocker_scopes)
    if not known_owner_ids.issubset(blocker_ids) or not exclusion_owners.issubset(
        blocker_ids
    ):
        fail(
            "readiness blocker ownership differs from compatibility authority"
        )
    # A release review may govern an already-admitted scope without pretending
    # that the review is a permanent compatibility exclusion. Such additional
    # owners disappear automatically when the tracker records them as closed,
    # but they may never invent or govern a scope outside the compatibility
    # authority's admitted product/platform boundary.
    for blocker_id in blocker_ids - known_owner_ids:
        if blocker_scopes[blocker_id] not in authority_scopes:
            fail(
                "additional readiness blocker does not govern an admitted scope: "
                f"{blocker_id}"
            )
    for scope, blocker_id in require_object(
        expected.get("requiredExclusions"), "required exclusions"
    ).items():
        if blocker_scopes.get(blocker_id) != scope:
            fail(f"excluded scope does not match blocker authority: {scope}")

    toolchain = load_object(
        ROOT / "docs" / "toolchain-policy.json", "toolchain authority"
    )
    authority_toolchains = {
        "haxe": [toolchain.get("haxe", {}).get("ci_selector")],
        "go": toolchain.get("go", {}).get("ci_versions"),
        "node": [toolchain.get("node", {}).get("ci_selector")],
    }
    if policy.get("toolchains") != authority_toolchains:
        fail("readiness toolchains drifted from toolchain authority")

    licensing = load_object(ROOT / "license-policy.json", "licensing authority")
    actual_license = require_object(
        evidence.get("licensing"), "licensing evidence"
    )
    approval = require_object(
        licensing.get("approval"), "licensing authority approval"
    )
    if actual_license.get("status") != licensing.get("status"):
        fail("licensing approval status differs from licensing authority")
    if actual_license.get("scopeSha256") != approval.get("scopeSha256"):
        fail("licensing evidence scope differs from licensing authority")
    if actual_license.get("unresolvedQuestions") != licensing.get(
        "unresolvedQuestions"
    ):
        fail("unresolved licensing questions differ from licensing authority")


def verify_release(evidence: dict[str, Any]) -> tuple[str, str, str]:
    release = require_object(evidence.get("release"), "release identity")
    version = release.get("version")
    tag = release.get("tag")
    tested_sha = release.get("testedSha")
    if not isinstance(version, str) or VERSION_RE.fullmatch(version) is None:
        fail("release version is not canonical SemVer")
    if version == "0.0.0":
        fail("release version cannot be the development sentinel")
    if tag != f"v{version}":
        fail("release tag does not match version")
    if release.get("manifestTag") != tag:
        fail("manifest tag does not match release tag")
    if not isinstance(tested_sha, str) or SHA_RE.fullmatch(tested_sha) is None:
        fail("tested SHA must be exactly 40 lowercase hexadecimal characters")
    if release.get("sourceSha") != tested_sha:
        fail("tested SHA does not match source SHA")
    return version, tag, tested_sha


def verify_compatibility(
    policy: dict[str, Any], evidence: dict[str, Any]
) -> set[str]:
    expected = require_object(policy["compatibility"], "compatibility policy")
    actual = require_object(evidence.get("compatibility"), "compatibility evidence")
    if actual.get("lifecycle") != expected.get("lifecycle"):
        fail("compatibility lifecycle differs from policy")
    if actual.get("claim") != expected.get("claim"):
        fail("public compatibility claim differs from governed claim")

    admitted = set(unique_strings(actual.get("admittedScopes"), "admitted scopes"))
    expected_admitted = set(
        unique_strings(expected.get("admittedScopes"), "policy admitted scopes")
    )
    if admitted != expected_admitted:
        fail("admitted scopes differ from readiness policy")
    evidenced = set(
        unique_strings(actual.get("evidencedScopes"), "evidenced scopes")
    )
    if not admitted.issubset(evidenced):
        fail("compatibility claim exceeds evidence")
    manifest = load_object(
        ROOT / "docs" / "compatibility-support-manifest.json",
        "compatibility authority",
    )
    surface_digest, operation_count, symbol_count = compatibility_surface(manifest)
    if actual.get("surfaceSha256") != surface_digest:
        fail("operation/member surface digest differs from compatibility authority")
    if actual.get("admittedOperationCount") != operation_count:
        fail("operation/member surface count differs from compatibility authority")
    if actual.get("admittedSymbolCount") != symbol_count:
        fail("operation/member symbol count differs from compatibility authority")

    required_exclusions = require_object(
        expected.get("requiredExclusions"), "required exclusions"
    )
    exclusions = require_list(actual.get("exclusions"), "compatibility exclusions")
    by_scope: dict[str, dict[str, Any]] = {}
    for raw in exclusions:
        exclusion = require_object(raw, "compatibility exclusion")
        scope = exclusion.get("scope")
        owner = exclusion.get("owner")
        if not isinstance(scope, str) or not scope:
            fail("exclusion scope must be a non-empty string")
        if scope in by_scope:
            fail(f"duplicate exclusion scope: {scope}")
        if not isinstance(owner, str) or not owner:
            fail(f"unowned exclusion: {scope}")
        if exclusion.get("advertisedAsSupported") is not False:
            fail(f"excluded scope advertised as supported: {scope}")
        by_scope[scope] = exclusion
    if set(by_scope) != set(required_exclusions):
        fail("roadmap exclusions differ from readiness policy")
    for scope, owner in required_exclusions.items():
        if by_scope[scope].get("owner") != owner:
            fail(f"exclusion owner differs from policy: {scope}")
    return admitted


def verify_toolchains(policy: dict[str, Any], evidence: dict[str, Any]) -> None:
    expected = require_object(policy["toolchains"], "toolchain policy")
    actual = require_object(evidence.get("toolchains"), "toolchain evidence")
    if set(actual) != set(expected):
        fail("supported toolchain families differ from policy")
    for family in ("haxe", "go", "node"):
        record = require_object(actual.get(family), f"evidenced {family} toolchains")
        if set(record) != {"resolved"}:
            fail(f"evidenced {family} toolchain fields must be exactly ['resolved']")
        resolved = unique_strings(
            record.get("resolved"), f"resolved {family} toolchains"
        )
        supported = unique_strings(expected.get(family), f"policy {family} toolchains")
        if family in {"haxe", "go"}:
            if resolved != supported:
                fail(f"supported toolchain evidence differs for {family}")
            continue
        if len(resolved) != 1 or VERSION_RE.fullmatch(resolved[0]) is None:
            fail("resolved Node toolchain must be one exact semantic version")
        if resolved[0].split(".", 1)[0] not in supported:
            fail("supported toolchain evidence differs for node")


def verify_public_api(evidence: dict[str, Any], tested_sha: str) -> None:
    public_api = require_object(evidence.get("publicApi"), "public API evidence")
    if public_api.get("result") != "pass":
        fail("public API policy did not pass")
    if public_api.get("testedSha") != tested_sha:
        fail("public API tested SHA differs from release SHA")


def verify_platform(evidence: dict[str, Any]) -> None:
    platform = require_object(evidence.get("platform"), "release platform evidence")
    if set(platform) != {"id", "os", "architecture", "runnerImage"}:
        fail("release platform evidence fields are incomplete or unknown")
    compatibility = load_object(
        ROOT / "docs" / "compatibility-support-manifest.json",
        "compatibility authority",
    )
    admitted = [
        item
        for item in require_list(
            compatibility.get("platforms"), "compatibility platforms"
        )
        if isinstance(item, dict) and item.get("release_admitted") is True
    ]
    if len(admitted) != 1:
        fail("compatibility authority must admit exactly one release platform")
    expected = admitted[0]
    if (
        platform.get("id") != expected.get("id")
        or platform.get("os") != expected.get("os")
        or platform.get("architecture") != expected.get("architecture")
    ):
        fail("release platform differs from admitted compatibility platform")
    image = require_object(platform.get("runnerImage"), "hosted runner image evidence")
    if set(image) != {"os", "version"}:
        fail("hosted runner image evidence fields are incomplete or unknown")
    for field in ("os", "version"):
        value = image.get(field)
        if not isinstance(value, str) or re.fullmatch(r"[A-Za-z0-9._-]+", value) is None:
            fail("exact hosted runner image OS and version are required")


def verify_security(
    policy: dict[str, Any], evidence: dict[str, Any], tested_sha: str
) -> None:
    security = require_object(evidence.get("security"), "security evidence")
    vulnerabilities = unique_strings(
        security.get("reachableVulnerabilities"), "reachable vulnerabilities"
    )
    if vulnerabilities:
        fail("reachable vulnerabilities are release-blocking")
    required = set(
        unique_strings(policy["requiredSecurityGates"], "required security gates")
    )
    raw_gates = require_list(security.get("gates"), "security gates")
    gates: dict[str, dict[str, Any]] = {}
    for raw in raw_gates:
        gate = require_object(raw, "security gate")
        gate_id = gate.get("id")
        if not isinstance(gate_id, str) or not gate_id or gate_id in gates:
            fail("security gate IDs must be unique non-empty strings")
        gates[gate_id] = gate
    if set(gates) != required:
        fail("required security gate set is incomplete or contains unknown gates")
    for gate_id, gate in gates.items():
        if gate.get("result") != "pass":
            fail(f"required security gate did not pass: {gate_id}")
        if gate.get("testedSha") != tested_sha:
            fail(f"gate tested SHA differs from release SHA: {gate_id}")


def verify_licensing(evidence: dict[str, Any]) -> None:
    licensing = require_object(evidence.get("licensing"), "licensing evidence")
    if licensing.get("status") != "approved":
        fail("licensing approval is unresolved")
    digest = licensing.get("scopeSha256")
    if not isinstance(digest, str) or re.fullmatch(r"[0-9a-f]{64}", digest) is None:
        fail("licensing approval scope digest is invalid")
    unresolved = unique_strings(
        licensing.get("unresolvedQuestions"), "unresolved licensing questions"
    )
    if unresolved:
        fail("unresolved licensing questions remain")


def verify_blockers(
    policy: dict[str, Any],
    evidence: dict[str, Any],
    admitted: set[str],
    tested_sha: str,
) -> None:
    authority = require_object(evidence.get("blockers"), "blocker evidence")
    if set(authority) != {"schemaVersion", "kind", "tracker", "records"}:
        fail("blocker evidence fields are incomplete or unknown")
    if authority.get("schemaVersion") != 1:
        fail("blocker evidence schema is unsupported")
    if authority.get("kind") != "haxe.go-release-blocker-evidence":
        fail("blocker evidence kind is invalid")
    tracker = require_object(authority.get("tracker"), "blocker tracker evidence")
    if set(tracker) != {"ref", "commit", "observedAt"}:
        fail("blocker tracker evidence fields are incomplete or unknown")
    if tracker.get("ref") != "refs/dolt/data":
        fail("blocker tracker evidence must identify refs/dolt/data")
    commit = tracker.get("commit")
    if not isinstance(commit, str) or SHA_RE.fullmatch(commit) is None:
        fail("blocker tracker evidence commit is invalid")
    if not isinstance(tracker.get("observedAt"), str) or not tracker["observedAt"]:
        fail("blocker tracker evidence observation date is missing")
    blockers = require_list(authority.get("records"), "blocker records")
    admission_policy = require_object(
        policy.get("releaseLineAdmission"), "release-line admission policy"
    )
    admission_owner = admission_policy.get("owner")
    admission_scope = admission_policy.get("scope")
    if not isinstance(admission_owner, str) or not admission_owner:
        fail("release-line admission owner is invalid")
    if not isinstance(admission_scope, str) or admission_scope not in admitted:
        fail("release-line admission scope is not release-admitted")
    seen: set[str] = set()
    for raw in blockers:
        blocker = require_object(raw, "blocker")
        blocker_id = blocker.get("id")
        priority = blocker.get("priority")
        status = blocker.get("status")
        scopes = set(unique_strings(blocker.get("scopes"), "blocker scopes"))
        if not isinstance(blocker_id, str) or not blocker_id or blocker_id in seen:
            fail("blocker IDs must be unique non-empty strings")
        seen.add(blocker_id)
        if not isinstance(priority, int) or isinstance(priority, bool) or priority < 0:
            fail(f"blocker priority is invalid: {blocker_id}")
        if not isinstance(status, str) or not status:
            fail(f"blocker status is invalid: {blocker_id}")
        if priority <= 1 and status in OPEN_STATUSES and scopes.intersection(admitted):
            fail(f"applicable unresolved P0/P1 blocker: {blocker_id}")
        expected_fields = {"id", "priority", "status", "scopes"}
        if blocker_id == admission_owner:
            expected_fields.add("admission")
        if set(blocker) != expected_fields:
            fail(f"blocker evidence fields are incomplete or unknown: {blocker_id}")
    compatibility = require_object(policy["compatibility"], "compatibility policy")
    expected_scopes = require_object(
        compatibility.get("blockerScopes"), "compatibility blocker scopes"
    )
    if seen != set(expected_scopes):
        fail("blocker evidence does not account for every governed roadmap owner")
    actual_by_id = {
        blocker.get("id"): blocker
        for blocker in blockers
        if isinstance(blocker, dict)
    }
    for blocker_id, scope in expected_scopes.items():
        if actual_by_id[blocker_id].get("scopes") != [scope]:
            fail(f"blocker evidence scope differs for {blocker_id}")
    if expected_scopes.get(admission_owner) != admission_scope:
        fail("release-line admission owner does not govern its release scope")

    final_record = require_object(
        actual_by_id[admission_owner].get("admission"), "final admission record"
    )
    if set(final_record) != {"schemaVersion", "oracleReview", "localDisposition"}:
        fail("final admission record fields are incomplete or unknown")
    if final_record.get("schemaVersion") != 1:
        fail("final admission record schema is unsupported")
    oracle = require_object(final_record.get("oracleReview"), "Oracle review admission")
    if set(oracle) != {
        "verdict",
        "reviewedSourceSha",
        "requestId",
        "frozenPacketSha256",
    }:
        fail("Oracle review admission fields are incomplete or unknown")
    oracle_expectations = {
        "verdict": admission_policy.get("oracleVerdict"),
        "reviewedSourceSha": admission_policy.get("oracleReviewedSourceSha"),
        "requestId": admission_policy.get("oracleRequestId"),
        "frozenPacketSha256": admission_policy.get("frozenPacketSha256"),
    }
    oracle_messages = {
        "verdict": "Oracle verdict",
        "reviewedSourceSha": "Oracle reviewed SHA",
        "requestId": "Oracle request",
        "frozenPacketSha256": "frozen packet",
    }
    for field, expected_value in oracle_expectations.items():
        if oracle.get(field) != expected_value:
            fail(f"{oracle_messages[field]} differs from final admission policy")

    disposition = require_object(
        final_record.get("localDisposition"), "local admission disposition"
    )
    if set(disposition) != {
        "verdict",
        "reviewedSourceSha",
        "dispositionSha256",
        "processorModel",
        "processorReasoningLevel",
    }:
        fail("local admission disposition fields are incomplete or unknown")
    if disposition.get("verdict") != admission_policy.get("localVerdict"):
        fail("local admission verdict differs from final admission policy")
    if disposition.get("reviewedSourceSha") != admission_policy.get(
        "baselineSourceSha"
    ):
        fail("local admission baseline SHA differs from release-line policy")
    disposition_digest = disposition.get("dispositionSha256")
    if (
        not isinstance(disposition_digest, str)
        or re.fullmatch(r"[0-9a-f]{64}", disposition_digest) is None
        or disposition_digest == "0" * 64
    ):
        fail("local admission disposition digest is invalid")
    disposition_relative = admission_policy.get("dispositionPath")
    if not isinstance(disposition_relative, str) or not disposition_relative:
        fail("final admission disposition path is invalid")
    disposition_path = (ROOT / disposition_relative).resolve()
    if (
        not disposition_path.is_relative_to(ROOT.resolve())
        or not disposition_path.is_file()
    ):
        fail("final admission disposition authority is unavailable")
    authority_digest = hashlib.sha256(disposition_path.read_bytes()).hexdigest()
    if disposition_digest != authority_digest:
        fail("local admission disposition digest differs from source authority")
    if disposition.get("processorModel") != admission_policy.get("processorModel"):
        fail("local admission processor model differs from final admission policy")
    if disposition.get("processorReasoningLevel") != admission_policy.get(
        "processorReasoningLevel"
    ):
        fail("local admission reasoning level differs from final admission policy")


def asset_map(assets: Any, label: str, *, with_roles: bool) -> dict[str, str]:
    result: dict[str, str] = {}
    for raw in require_list(assets, label):
        asset = require_object(raw, f"{label} record")
        name = asset.get("name")
        digest = asset.get("digest")
        if not isinstance(name, str) or not name or name in result:
            fail(f"{label} names must be unique non-empty strings")
        if not isinstance(digest, str) or DIGEST_RE.fullmatch(digest) is None:
            fail(f"{label} digest is invalid: {name}")
        if with_roles and (not isinstance(asset.get("role"), str) or not asset["role"]):
            fail(f"{label} role is invalid: {name}")
        result[name] = digest
    return result


def verify_artifacts(
    policy: dict[str, Any],
    evidence: dict[str, Any],
    version: str,
    tested_sha: str,
) -> dict[str, str]:
    artifacts = require_object(evidence.get("artifacts"), "artifact evidence")
    if artifacts.get("verifiedForSha") != tested_sha:
        fail("artifact verified SHA differs from release SHA")
    raw_assets = require_list(artifacts.get("assets"), "release assets")
    assets = asset_map(raw_assets, "release assets", with_roles=True)
    roles = [require_object(item, "release asset").get("role") for item in raw_assets]
    required_roles = unique_strings(
        policy["requiredArtifactRoles"], "required artifact roles"
    )
    if sorted(roles) != sorted(required_roles):
        fail("required artifact roles are incomplete or duplicated")
    expected_names = {
        "haxelib-zip": f"reflaxe.go-{version}.zip",
        "checksum": f"reflaxe.go-{version}.zip.sha256",
        "content-manifest": f"reflaxe.go-{version}.manifest.json",
        "provenance": f"reflaxe.go-{version}.provenance.json",
    }
    for raw in raw_assets:
        asset = require_object(raw, "release asset")
        if asset.get("name") != expected_names.get(asset.get("role")):
            fail(f"release asset name does not match its role: {asset.get('role')}")
    subjects = set(
        unique_strings(
            artifacts.get("provenanceSubjects"), "provenance subjects"
        )
    )
    expected_subjects = set(assets) - {expected_names["provenance"]}
    if subjects != expected_subjects:
        fail("provenance subjects do not cover the exact non-provenance assets")
    return assets


def verify_github(
    evidence: dict[str, Any],
    phase: str,
    tag: str,
    tested_sha: str,
    assets: dict[str, str],
) -> None:
    github = evidence.get("github")
    if phase == "candidate":
        if github is not None:
            fail("candidate evidence must not pretend hosted GitHub state exists")
        return
    if phase != "published":
        fail("evidence phase must be candidate or published")
    hosted = require_object(github, "hosted GitHub evidence")
    if hosted.get("apiAuthoritative") is not True:
        fail("GitHub API state must be authoritative")
    if hosted.get("tag") != tag:
        fail("hosted release tag differs from release identity")
    if hosted.get("targetCommit") != tested_sha:
        fail("hosted release target differs from tested SHA")
    if hosted.get("draft") is not False:
        fail("published release remains a draft")
    if hosted.get("immutable") is not True:
        fail("published release is not immutable")
    hosted_assets = asset_map(hosted.get("assets"), "hosted assets", with_roles=False)
    if hosted_assets != assets:
        fail("hosted assets do not exactly match verified local assets")


def run_gh_api(repository: str, endpoint: str) -> Any:
    command = [
        "gh",
        "api",
        "-H",
        "Accept: application/vnd.github+json",
        "-H",
        "X-GitHub-Api-Version: 2022-11-28",
        f"repos/{repository}/{endpoint}",
    ]
    try:
        process = subprocess.run(
            command,
            cwd=ROOT,
            text=True,
            capture_output=True,
            timeout=30,
        )
    except (OSError, subprocess.TimeoutExpired) as error:
        fail(f"GitHub API query failed: {error}")
    if process.returncode != 0:
        detail = (process.stderr or process.stdout).strip()
        fail(f"GitHub API query failed for {endpoint}: {detail}")
    try:
        return json.loads(process.stdout)
    except json.JSONDecodeError:
        fail(f"GitHub API returned invalid JSON for {endpoint}")


def resolve_live_tag_commit(repository: str, tag: str) -> str:
    reference = run_gh_api(repository, f"git/ref/tags/{tag}")
    current = require_object(reference, "GitHub tag reference").get("object")
    for _ in range(4):
        current_object = require_object(current, "GitHub tag object")
        object_type = current_object.get("type")
        object_sha = current_object.get("sha")
        if not isinstance(object_sha, str) or SHA_RE.fullmatch(object_sha) is None:
            fail("GitHub tag object SHA is invalid")
        if object_type == "commit":
            return object_sha
        if object_type != "tag":
            fail(f"GitHub tag resolves to unsupported object type: {object_type}")
        tag_object = run_gh_api(repository, f"git/tags/{object_sha}")
        current = require_object(tag_object, "annotated GitHub tag").get("object")
    fail("GitHub tag annotation chain is too deep")


def load_live_github(policy: dict[str, Any], evidence: dict[str, Any]) -> dict[str, Any]:
    release_identity = require_object(evidence.get("release"), "release identity")
    tag = release_identity.get("tag")
    if not isinstance(tag, str) or not tag:
        fail("release tag is required before querying GitHub")
    repository = policy["repository"]
    release = run_gh_api(repository, f"releases/tags/{tag}")
    raw_assets = require_list(release.get("assets"), "GitHub release assets")
    assets: list[dict[str, str]] = []
    for raw in raw_assets:
        asset = require_object(raw, "GitHub release asset")
        name = asset.get("name")
        digest = asset.get("digest")
        if not isinstance(name, str) or not isinstance(digest, str):
            fail("GitHub release asset is missing API name or digest")
        assets.append({"name": name, "digest": digest})
    return {
        "apiAuthoritative": True,
        "tag": release.get("tag_name"),
        "targetCommit": resolve_live_tag_commit(repository, tag),
        "draft": release.get("draft"),
        "immutable": release.get("immutable"),
        "assets": assets,
    }


def evaluate(
    policy: dict[str, Any], evidence: dict[str, Any], *, mode: str
) -> str:
    verify_policy(policy)
    if set(evidence) != EVIDENCE_FIELDS:
        fail(f"evidence fields must be exactly {sorted(EVIDENCE_FIELDS)}")
    if evidence.get("schemaVersion") != 1:
        fail("unsupported readiness evidence schema")
    if evidence.get("kind") != "haxe.go-release-readiness-evidence":
        fail("readiness evidence kind is invalid")
    if mode not in {"fixture", "live"}:
        fail("readiness mode must be fixture or live")
    phase = evidence.get("phase")
    if not isinstance(phase, str):
        fail("evidence phase must be candidate or published")

    version, tag, tested_sha = verify_release(evidence)
    verify_release_line_ancestry(policy, tested_sha, mode=mode)
    verify_repository_authorities(policy, evidence)
    admitted = verify_compatibility(policy, evidence)
    verify_public_api(evidence, tested_sha)
    verify_platform(evidence)
    verify_toolchains(policy, evidence)
    verify_security(policy, evidence, tested_sha)
    verify_licensing(evidence)
    verify_blockers(policy, evidence, admitted, tested_sha)
    assets = verify_artifacts(policy, evidence, version, tested_sha)
    if mode == "live" and phase == "published":
        evidence = dict(evidence)
        evidence["github"] = load_live_github(policy, evidence)
    verify_github(evidence, phase, tag, tested_sha, assets)
    return phase


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--policy", required=True, type=Path)
    parser.add_argument("--evidence", required=True, type=Path)
    parser.add_argument(
        "--mode",
        required=True,
        choices=("fixture", "live"),
        help=(
            "fixture uses supplied hosted state for deterministic tests; live "
            "replaces published hosted state with GitHub API responses"
        ),
    )
    return parser.parse_args()


def main() -> int:
    arguments = parse_args()
    try:
        phase = evaluate(
            load_object(arguments.policy, "readiness policy"),
            load_object(arguments.evidence, "readiness evidence"),
            mode=arguments.mode,
        )
    except ReadinessError as error:
        print(f"[release-readiness] ERROR: {error}", file=sys.stderr)
        return 2
    print(f"[release-readiness] {phase} release evidence: READY")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
