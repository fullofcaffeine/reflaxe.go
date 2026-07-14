#!/usr/bin/env python3

"""Audit licensing provenance and fail closed before a release is published."""

from __future__ import annotations

import argparse
from datetime import date
import hashlib
import json
from pathlib import Path, PurePosixPath
import re
import sys
from typing import NoReturn


REPOSITORY_ROOT = Path(__file__).resolve().parents[2]
POLICY_KIND = "haxe.go-license-policy"
SHA256_PATTERN = re.compile(r"[0-9a-f]{64}")
DATE_PATTERN = re.compile(r"[0-9]{4}-[0-9]{2}-[0-9]{2}")
APPROVAL_AUTHORITIES = {"project-copyright-owner", "qualified-legal-review"}
APPROVAL_FIELDS = {
    "decidedBy",
    "authority",
    "decisionDate",
    "decisionRecord",
    "scopeSha256",
}


class LicensePolicyError(RuntimeError):
    """Raised when policy facts, approval, or packaged notice bytes drift."""


def fail(message: str) -> NoReturn:
    raise LicensePolicyError(message)


def sha256_bytes(value: bytes) -> str:
    return hashlib.sha256(value).hexdigest()


def canonical_scope_digest(policy: dict[str, object]) -> str:
    try:
        scope = {
            key: policy[key]
            for key in (
                "shippedSourcePatterns",
                "components",
                "generatedOutputClasses",
                "releasePackage",
            )
        }
    except KeyError as error:
        fail(f"policy scope is missing {error.args[0]!r}")
    encoded = json.dumps(
        scope,
        ensure_ascii=False,
        separators=(",", ":"),
        sort_keys=True,
    ).encode("utf-8")
    return sha256_bytes(encoded)


def load_policy(path: Path) -> dict[str, object]:
    try:
        value = json.loads(path.read_text(encoding="utf-8"))
    except OSError as error:
        fail(f"cannot read license policy {path}: {error}")
    except json.JSONDecodeError as error:
        fail(f"license policy is not valid JSON: {error}")
    if not isinstance(value, dict):
        fail("license policy must contain a JSON object")
    return value


def safe_relative_path(value: object, label: str) -> str:
    if not isinstance(value, str) or not value:
        fail(f"{label} must be a non-empty repository-relative path")
    path = PurePosixPath(value)
    if (
        "\\" in value
        or path.is_absolute()
        or path.as_posix() != value
        or any(part in {"", ".", ".."} for part in path.parts)
    ):
        fail(f"{label} is not a safe repository-relative POSIX path: {value!r}")
    return value


def require_object(value: object, label: str) -> dict[str, object]:
    if not isinstance(value, dict):
        fail(f"{label} must be a JSON object")
    return value


def require_list(value: object, label: str) -> list[object]:
    if not isinstance(value, list):
        fail(f"{label} must be a JSON array")
    return value


def require_nonempty_string(value: object, label: str) -> str:
    if not isinstance(value, str) or not value.strip():
        fail(f"{label} must be a non-empty string")
    return value


def require_string_list(value: object, label: str) -> list[str]:
    items = require_list(value, label)
    result: list[str] = []
    for index, item in enumerate(items):
        result.append(require_nonempty_string(item, f"{label}[{index}]"))
    return result


def resolve_inside(root: Path, relative: str, label: str) -> Path:
    path = root.joinpath(*PurePosixPath(relative).parts)
    try:
        path.resolve(strict=False).relative_to(root.resolve())
    except ValueError:
        fail(f"{label} escapes the policy root: {relative}")
    return path


def validate_hash(value: object, label: str) -> str:
    if not isinstance(value, str) or SHA256_PATTERN.fullmatch(value) is None:
        fail(f"{label} must be exactly 64 lowercase hexadecimal characters")
    return value


def validate_component_inventory(root: Path, policy: dict[str, object]) -> None:
    components = require_list(policy.get("components"), "components")
    if not components:
        fail("components must not be empty")
    seen_ids: set[str] = set()
    component_sources: dict[str, str] = {}
    for index, raw_component in enumerate(components):
        component = require_object(raw_component, f"components[{index}]")
        component_id = require_nonempty_string(
            component.get("id"), f"components[{index}].id"
        )
        if component_id in seen_ids:
            fail(f"component id is duplicated: {component_id}")
        seen_ids.add(component_id)
        require_nonempty_string(
            component.get("provenance"), f"component {component_id} provenance"
        )
        licenses = require_string_list(
            component.get("declaredLicenses"),
            f"component {component_id} declaredLicenses",
        )
        if not licenses:
            fail(f"component {component_id} must declare at least one observed license")
        treatment = require_nonempty_string(
            component.get("generatedOutputTreatment"),
            f"component {component_id} generatedOutputTreatment",
        )
        if treatment == "approved":
            fail(
                f"component {component_id} uses the ambiguous treatment 'approved'; "
                "record the concrete treatment instead"
            )

        patterns = require_string_list(
            component.get("sourcePatterns"),
            f"component {component_id} sourcePatterns",
        )
        if not patterns:
            fail(f"component {component_id} must cover at least one source pattern")
        for pattern in patterns:
            safe_relative_path(pattern, f"component {component_id} source pattern")
            matches = sorted(path for path in root.glob(pattern) if path.is_file())
            if not matches:
                fail(
                    f"component {component_id} source pattern matches no files: {pattern}"
                )
            for match in matches:
                try:
                    match.resolve().relative_to(root.resolve())
                except ValueError:
                    fail(
                        f"component {component_id} source pattern escapes the root: {pattern}"
                    )
                relative = match.relative_to(root).as_posix()
                previous = component_sources.get(relative)
                if previous is not None and previous != component_id:
                    fail(
                        f"shipped source is assigned to multiple components: "
                        f"{relative} ({previous}, {component_id})"
                    )
                component_sources[relative] = component_id

        evidence = component.get("licenseEvidence", [])
        for evidence_index, raw_record in enumerate(
            require_list(evidence, f"component {component_id} licenseEvidence")
        ):
            record = require_object(
                raw_record,
                f"component {component_id} licenseEvidence[{evidence_index}]",
            )
            relative = safe_relative_path(
                record.get("path"),
                f"component {component_id} license evidence path",
            )
            expected = validate_hash(
                record.get("sha256"),
                f"component {component_id} license evidence sha256",
            )
            source = resolve_inside(root, relative, "license evidence")
            if not source.is_file():
                fail(f"required license material is missing: {relative}")
            actual = sha256_bytes(source.read_bytes())
            if actual != expected:
                fail(
                    f"license material hash differs for {relative}: "
                    f"expected {expected}, got {actual}"
                )

    shipped_patterns = require_string_list(
        policy.get("shippedSourcePatterns"), "shippedSourcePatterns"
    )
    if not shipped_patterns:
        fail("shippedSourcePatterns must not be empty")
    shipped_sources: set[str] = set()
    for pattern in shipped_patterns:
        safe_relative_path(pattern, "shipped source pattern")
        matches = sorted(path for path in root.glob(pattern) if path.is_file())
        if not matches:
            fail(f"shipped source pattern matches no files: {pattern}")
        for match in matches:
            try:
                match.resolve().relative_to(root.resolve())
            except ValueError:
                fail(f"shipped source pattern escapes the root: {pattern}")
            shipped_sources.add(match.relative_to(root).as_posix())

    unclassified = sorted(shipped_sources - set(component_sources))
    if unclassified:
        fail(f"shipped source is not assigned to a component: {unclassified[0]}")
    non_shipped = sorted(set(component_sources) - shipped_sources)
    if non_shipped:
        fail(f"component classifies a source that is not shipped: {non_shipped[0]}")


def validate_generated_output(policy: dict[str, object]) -> None:
    outputs = require_list(
        policy.get("generatedOutputClasses"), "generatedOutputClasses"
    )
    if not outputs:
        fail("generatedOutputClasses must not be empty")
    seen_ids: set[str] = set()
    for index, raw_output in enumerate(outputs):
        output = require_object(raw_output, f"generatedOutputClasses[{index}]")
        output_id = require_nonempty_string(
            output.get("id"), f"generatedOutputClasses[{index}].id"
        )
        if output_id in seen_ids:
            fail(f"generated-output class id is duplicated: {output_id}")
        seen_ids.add(output_id)
        require_nonempty_string(output.get("origin"), f"output {output_id} origin")
        require_nonempty_string(
            output.get("licenseTreatment"),
            f"output {output_id} licenseTreatment",
        )
        artifacts = output.get("requiredArtifacts")
        if artifacts is not None:
            for artifact in require_string_list(
                artifacts, f"output {output_id} requiredArtifacts"
            ):
                safe_relative_path(artifact, f"output {output_id} required artifact")


def validate_release_material(
    root: Path,
    policy_path: Path,
    policy: dict[str, object],
    package_root: Path | None,
) -> None:
    release_package = require_object(
        policy.get("releasePackage"), "releasePackage"
    )
    required = require_list(
        release_package.get("requiredFiles"), "releasePackage.requiredFiles"
    )
    if not required:
        fail("releasePackage.requiredFiles must not be empty")
    seen_sources: set[str] = set()
    seen_packages: set[str] = set()
    for index, raw_record in enumerate(required):
        record = require_object(raw_record, f"releasePackage.requiredFiles[{index}]")
        source_relative = safe_relative_path(
            record.get("sourcePath"),
            f"releasePackage.requiredFiles[{index}].sourcePath",
        )
        package_relative = safe_relative_path(
            record.get("packagePath"),
            f"releasePackage.requiredFiles[{index}].packagePath",
        )
        if source_relative in seen_sources:
            fail(f"release license source is duplicated: {source_relative}")
        if package_relative in seen_packages:
            fail(f"release package license path is duplicated: {package_relative}")
        seen_sources.add(source_relative)
        seen_packages.add(package_relative)

        source = resolve_inside(root, source_relative, "release license source")
        if not source.is_file():
            fail(f"required release license material is missing: {source_relative}")
        expected_value = record.get("sha256")
        if expected_value is None:
            if source.resolve() != policy_path.resolve():
                fail(
                    "only the self-referential license policy may use a null sha256: "
                    f"{source_relative}"
                )
            expected = sha256_bytes(source.read_bytes())
        else:
            expected = validate_hash(
                expected_value,
                f"releasePackage.requiredFiles[{index}].sha256",
            )
            actual = sha256_bytes(source.read_bytes())
            if actual != expected:
                fail(
                    f"license material hash differs for {source_relative}: "
                    f"expected {expected}, got {actual}"
                )

        if package_root is not None:
            packaged = resolve_inside(
                package_root.resolve(), package_relative, "packaged license material"
            )
            if not packaged.is_file():
                fail(
                    f"required package license material is missing: {package_relative}"
                )
            packaged_hash = sha256_bytes(packaged.read_bytes())
            if packaged_hash != expected:
                fail(
                    f"packaged license material hash differs for {package_relative}: "
                    f"expected {expected}, got {packaged_hash}"
                )


def validate_approval(policy: dict[str, object], digest: str) -> None:
    approval = require_object(policy.get("approval"), "approval")
    if set(approval) != APPROVAL_FIELDS:
        fail(
            "approval must contain exactly: "
            + ", ".join(sorted(APPROVAL_FIELDS))
        )
    decided_by = require_nonempty_string(approval.get("decidedBy"), "approval.decidedBy")
    if decided_by.lower() in {"tbd", "unknown", "n/a"}:
        fail("approval.decidedBy must identify an accountable reviewer")
    authority = require_nonempty_string(approval.get("authority"), "approval.authority")
    if authority not in APPROVAL_AUTHORITIES:
        fail(
            "approval.authority must be project-copyright-owner or "
            "qualified-legal-review"
        )
    decision_date = require_nonempty_string(
        approval.get("decisionDate"), "approval.decisionDate"
    )
    if DATE_PATTERN.fullmatch(decision_date) is None:
        fail("approval.decisionDate must use YYYY-MM-DD")
    try:
        date.fromisoformat(decision_date)
    except ValueError:
        fail("approval.decisionDate is not a valid calendar date")
    require_nonempty_string(approval.get("decisionRecord"), "approval.decisionRecord")
    approved_digest = validate_hash(
        approval.get("scopeSha256"), "approval.scopeSha256"
    )
    if approved_digest != digest:
        fail(
            f"approval scope digest is stale: expected current scope {digest}, "
            f"got {approved_digest}"
        )

    questions = require_list(policy.get("unresolvedQuestions"), "unresolvedQuestions")
    if questions:
        fail("approved policy still contains unresolvedQuestions")
    for component in require_list(policy.get("components"), "components"):
        record = require_object(component, "component")
        if record.get("generatedOutputTreatment") == "unresolved":
            fail(
                f"approved policy leaves component {record.get('id')} output treatment unresolved"
            )
    for output in require_list(
        policy.get("generatedOutputClasses"), "generatedOutputClasses"
    ):
        record = require_object(output, "generated-output class")
        output_id = record.get("id")
        if record.get("licenseTreatment") == "unresolved":
            fail(f"approved policy leaves output {output_id} treatment unresolved")
        if not isinstance(record.get("requiredArtifacts"), list):
            fail(
                f"approved policy must give output {output_id} an explicit "
                "requiredArtifacts array"
            )


def validate_policy(
    root: Path,
    policy_path: Path,
    policy: dict[str, object],
    package_root: Path | None,
) -> str:
    if policy.get("schemaVersion") != 1:
        fail("license policy schemaVersion must be 1")
    if policy.get("kind") != POLICY_KIND:
        fail(f"license policy kind must be {POLICY_KIND!r}")
    status = policy.get("status")
    if status not in {"unresolved", "approved"}:
        fail("license policy status must be unresolved or approved")
    approval = require_object(policy.get("approval"), "approval")
    if set(approval) != APPROVAL_FIELDS:
        fail(
            "approval must contain exactly: "
            + ", ".join(sorted(APPROVAL_FIELDS))
        )
    if status == "unresolved" and any(value is not None for value in approval.values()):
        fail("an unresolved policy must not carry partial or apparent approval")

    validate_component_inventory(root, policy)
    validate_generated_output(policy)
    validate_release_material(root, policy_path, policy, package_root)
    questions = require_list(policy.get("unresolvedQuestions"), "unresolvedQuestions")
    for index, question in enumerate(questions):
        require_nonempty_string(question, f"unresolvedQuestions[{index}]")
    digest = canonical_scope_digest(policy)
    if status == "approved":
        validate_approval(policy, digest)
    return digest


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument(
        "--root",
        type=Path,
        default=REPOSITORY_ROOT,
        help="source tree containing license-policy.json",
    )
    parser.add_argument(
        "--policy",
        type=Path,
        help="policy path (defaults to <root>/license-policy.json)",
    )
    parser.add_argument(
        "--package-root",
        type=Path,
        help="optional staged package whose required license bytes are audited",
    )
    parser.add_argument("--mode", choices=("audit", "release"), default="audit")
    return parser.parse_args()


def main() -> int:
    args = parse_args()
    root = args.root.resolve()
    policy_path = (
        args.policy.resolve()
        if args.policy is not None
        else root / "license-policy.json"
    )
    try:
        if not root.is_dir():
            fail(f"policy root is not a directory: {root}")
        if args.package_root is not None and not args.package_root.is_dir():
            fail(f"package root is not a directory: {args.package_root}")
        policy = load_policy(policy_path)
        digest = validate_policy(root, policy_path, policy, args.package_root)
        status = policy.get("status")
        print("[license-policy] inventory and license material: OK")
        print(f"[license-policy] scope sha256: {digest}")
        if args.mode == "release":
            if status != "approved":
                fail("license policy is unresolved; public release is blocked")
            print("[license-policy] approved release policy: OK")
        else:
            print(f"[license-policy] policy status: {status}")
    except LicensePolicyError as error:
        print(f"[license-policy] error: {error}", file=sys.stderr)
        return 1
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
