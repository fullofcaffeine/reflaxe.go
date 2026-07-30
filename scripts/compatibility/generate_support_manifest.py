#!/usr/bin/env python3

from __future__ import annotations

import argparse
import hashlib
import json
import re
import sys
from collections import Counter
from pathlib import Path
from typing import Any


REPO_ROOT = Path(__file__).resolve().parents[2]
DEFAULT_SOURCE = REPO_ROOT / "docs" / "compatibility-support-source.json"
DEFAULT_TOOLCHAIN = REPO_ROOT / "docs" / "toolchain-policy.json"
DEFAULT_PORTABLE_INVENTORY = REPO_ROOT / "test" / "portable_stdlib_inventory.json"
DEFAULT_MANIFEST = REPO_ROOT / "docs" / "compatibility-support-manifest.json"
DEFAULT_MATRIX_DOC = REPO_ROOT / "docs" / "compatibility-support-matrix.md"
DEFAULT_RELEASE_DOC = REPO_ROOT / "docs" / "compatibility-release-status.md"
GENERATED_BY = "scripts/compatibility/generate_support_manifest.py"

ALLOWED_STATES = {
    "semantic-diff-supported",
    "compile-go-test-run-supported",
    "compile-only",
    "experimental",
    "compatibility-only",
    "excluded",
}
RELEASE_ELIGIBLE_STATES = {
    "semantic-diff-supported",
    "compile-go-test-run-supported",
}
ALLOWED_EVIDENCE_KINDS = {"path", "npm-script", "bead"}
ALLOWED_EVIDENCE_LEVELS = {
    "semantic-diff",
    "compile-go-test-run",
    "compile-only",
    "policy",
    "compatibility",
    "blocker",
}
STRONG_EVIDENCE_LEVELS = {"semantic-diff", "compile-go-test-run"}
PORTABLE_STATUS_MAP = {
    "semantic-diff": "semantic-diff-supported",
    "snapshot": "compile-go-test-run-supported",
    "compile-only": "compile-only",
    "unsupported": "excluded",
}
BEAD_ID = re.compile(r"^haxe_go-[A-Za-z0-9_.-]+$")


class ContractError(Exception):
    pass


def load_json(path: Path, label: str) -> dict[str, Any]:
    try:
        value = json.loads(path.read_text(encoding="utf-8"))
    except FileNotFoundError as error:
        raise ContractError(f"{label} not found: {display_path(path)}") from error
    except json.JSONDecodeError as error:
        raise ContractError(f"{label} is not valid JSON: {display_path(path)}: {error}") from error
    if not isinstance(value, dict):
        raise ContractError(f"{label} must be a JSON object")
    return value


def display_path(path: Path) -> str:
    try:
        return path.resolve().relative_to(REPO_ROOT).as_posix()
    except ValueError:
        return path.as_posix()


def sha256(path: Path) -> str:
    return hashlib.sha256(path.read_bytes()).hexdigest()


def require_object(owner: dict[str, Any], field: str, label: str) -> dict[str, Any]:
    value = owner.get(field)
    if not isinstance(value, dict):
        raise ContractError(f"{label} must declare object {field}")
    return value


def require_list(owner: dict[str, Any], field: str, label: str, allow_empty: bool = False) -> list[Any]:
    value = owner.get(field)
    if not isinstance(value, list) or (not allow_empty and not value):
        suffix = "an array" if allow_empty else "a non-empty array"
        raise ContractError(f"{label} must declare {field} as {suffix}")
    return value


def require_string(owner: dict[str, Any], field: str, label: str) -> str:
    value = owner.get(field)
    if not isinstance(value, str) or not value.strip():
        raise ContractError(f"{label} must declare {field}")
    return value


def require_bool(owner: dict[str, Any], field: str, label: str) -> bool:
    value = owner.get(field)
    if not isinstance(value, bool):
        raise ContractError(f"{label} must declare boolean {field}")
    return value


def reject_unknown_fields(owner: dict[str, Any], allowed: set[str], label: str) -> None:
    unknown = sorted(set(owner) - allowed)
    if unknown:
        raise ContractError(f"{label} has unknown fields: {unknown}")


def require_state(owner: dict[str, Any], label: str) -> str:
    if "state" not in owner:
        raise ContractError(f"{label} must declare state")
    state = owner["state"]
    if not isinstance(state, str) or state not in ALLOWED_STATES:
        raise ContractError(f"{label} has unknown evidence state: {state!r}")
    return state


def require_string_list(owner: dict[str, Any], field: str, label: str, allow_empty: bool = False) -> list[str]:
    values = require_list(owner, field, label, allow_empty=allow_empty)
    for value in values:
        if not isinstance(value, str) or not value.strip():
            raise ContractError(f"{label}.{field} contains an empty or non-string value")
    if len(set(values)) != len(values):
        raise ContractError(f"{label}.{field} contains duplicates")
    return values


def validate_source_header(source: dict[str, Any]) -> None:
    reject_unknown_fields(
        source,
        {
            "schema_version",
            "manifest_schema_version",
            "as_of",
            "release_claim",
            "evidence_states",
            "evidence",
            "platforms",
            "presets",
            "surfaces",
            "trust_assumptions",
            "known_blockers",
        },
        "compatibility source",
    )
    if source.get("schema_version") != 1:
        raise ContractError("compatibility source schema_version must be 1")
    if source.get("manifest_schema_version") != 1:
        raise ContractError("compatibility source manifest_schema_version must be 1")
    require_string(source, "as_of", "compatibility source")

    states = require_object(source, "evidence_states", "compatibility source")
    if set(states) != ALLOWED_STATES:
        missing = sorted(ALLOWED_STATES - set(states))
        unknown = sorted(set(states) - ALLOWED_STATES)
        raise ContractError(f"evidence_states must be closed; missing={missing}, unknown={unknown}")
    for state, policy in states.items():
        if not isinstance(policy, dict):
            raise ContractError(f"evidence state {state} must be an object")
        reject_unknown_fields(policy, {"meaning", "release_eligible"}, f"evidence state {state}")
        require_string(policy, "meaning", f"evidence state {state}")
        release_eligible = require_bool(policy, "release_eligible", f"evidence state {state}")
        if release_eligible != (state in RELEASE_ELIGIBLE_STATES):
            raise ContractError(f"evidence state {state} has inconsistent release_eligible")

    release_claim = require_object(source, "release_claim", "compatibility source")
    reject_unknown_fields(
        release_claim,
        {
            "lifecycle",
            "statement",
            "default_disposition",
            "admitted_preset",
            "admitted_platform",
            "forbidden_phrases",
        },
        "release_claim",
    )
    require_string(release_claim, "lifecycle", "release_claim")
    statement = require_string(release_claim, "statement", "release_claim")
    require_string(release_claim, "default_disposition", "release_claim")
    require_string(release_claim, "admitted_preset", "release_claim")
    require_string(release_claim, "admitted_platform", "release_claim")
    forbidden = require_string_list(release_claim, "forbidden_phrases", "release_claim")
    lowered_statement = statement.lower()
    for phrase in forbidden:
        if phrase.lower() in lowered_statement:
            raise ContractError(f"release claim contains forbidden phrase: {phrase}")


def validate_evidence(source: dict[str, Any], package: dict[str, Any]) -> dict[str, dict[str, Any]]:
    entries = require_list(source, "evidence", "compatibility source")
    result: dict[str, dict[str, Any]] = {}
    package_scripts = package.get("scripts", {})
    if not isinstance(package_scripts, dict):
        raise ContractError("package.json scripts must be an object")

    for raw_entry in entries:
        if not isinstance(raw_entry, dict):
            raise ContractError("evidence entry must be an object")
        evidence_id = require_string(raw_entry, "id", "evidence entry")
        if evidence_id in result:
            raise ContractError(f"duplicate evidence id: {evidence_id}")
        kind = require_string(raw_entry, "kind", f"evidence {evidence_id}")
        level = require_string(raw_entry, "level", f"evidence {evidence_id}")
        if kind not in ALLOWED_EVIDENCE_KINDS:
            raise ContractError(f"evidence {evidence_id} has unknown kind: {kind}")
        if level not in ALLOWED_EVIDENCE_LEVELS:
            raise ContractError(f"evidence {evidence_id} has unknown level: {level}")
        field_by_kind = {"path": "path", "npm-script": "script", "bead": "bead"}
        reject_unknown_fields(
            raw_entry,
            {"id", "kind", "level", field_by_kind[kind]},
            f"evidence {evidence_id}",
        )
        if kind == "path":
            raw_path = require_string(raw_entry, "path", f"evidence {evidence_id}")
            path = Path(raw_path)
            if path.is_absolute() or ".." in path.parts:
                raise ContractError(f"evidence {evidence_id} path must be repository-relative")
            if not (REPO_ROOT / path).exists():
                raise ContractError(f"evidence path does not exist: {raw_path}")
        elif kind == "npm-script":
            script = require_string(raw_entry, "script", f"evidence {evidence_id}")
            if script not in package_scripts:
                raise ContractError(f"evidence npm script does not exist: {script}")
        else:
            bead = require_string(raw_entry, "bead", f"evidence {evidence_id}")
            if BEAD_ID.fullmatch(bead) is None:
                raise ContractError(f"evidence {evidence_id} has invalid Bead id: {bead}")
        result[evidence_id] = raw_entry
    return result


def validate_evidence_references(
    owner: dict[str, Any],
    label: str,
    evidence: dict[str, dict[str, Any]],
) -> list[str]:
    evidence_ids = require_string_list(owner, "evidence_ids", label)
    for evidence_id in evidence_ids:
        if evidence_id not in evidence:
            raise ContractError(f"{label} references unknown evidence: {evidence_id}")
    return evidence_ids


def validate_platforms_or_presets(
    source: dict[str, Any],
    field: str,
    evidence: dict[str, dict[str, Any]],
) -> list[dict[str, Any]]:
    entries = require_list(source, field, "compatibility source")
    seen: set[str] = set()
    for raw_entry in entries:
        if not isinstance(raw_entry, dict):
            raise ContractError(f"{field} entry must be an object")
        entry_id = require_string(raw_entry, "id", f"{field} entry")
        label = f"{field} {entry_id}"
        common_fields = {
            "id",
            "state",
            "release_admitted",
            "qualification",
            "evidence_ids",
            "exclusions",
        }
        specific_fields = {"os", "architecture"} if field == "platforms" else {"selector"}
        reject_unknown_fields(raw_entry, common_fields | specific_fields, label)
        if entry_id in seen:
            raise ContractError(f"duplicate {field} id: {entry_id}")
        seen.add(entry_id)
        state = require_state(raw_entry, label)
        release_admitted = require_bool(raw_entry, "release_admitted", label)
        require_string(raw_entry, "qualification", label)
        require_string_list(raw_entry, "exclusions", label, allow_empty=True)
        evidence_ids = validate_evidence_references(raw_entry, label, evidence)
        levels = {evidence[evidence_id]["level"] for evidence_id in evidence_ids}
        if state == "semantic-diff-supported" and "semantic-diff" not in levels:
            raise ContractError(f"{label} lacks semantic-diff evidence")
        if state == "compile-go-test-run-supported" and not (levels & STRONG_EVIDENCE_LEVELS):
            raise ContractError(f"{label} lacks compile-go-test-run evidence")
        if release_admitted and not (levels & STRONG_EVIDENCE_LEVELS):
            raise ContractError(f"{label} lacks executable release evidence")
        if field == "platforms":
            require_string(raw_entry, "os", label)
            require_string(raw_entry, "architecture", label)
        else:
            require_string(raw_entry, "selector", label)
        if release_admitted and state not in RELEASE_ELIGIBLE_STATES:
            raise ContractError(f"{label} cannot be release admitted with state {state}")
    return entries


def validate_surfaces(
    source: dict[str, Any],
    evidence: dict[str, dict[str, Any]],
) -> list[dict[str, Any]]:
    surfaces = require_list(source, "surfaces", "compatibility source")
    seen_surface_ids: set[str] = set()
    seen_symbols: set[str] = set()
    for raw_surface in surfaces:
        if not isinstance(raw_surface, dict):
            raise ContractError("surface entry must be an object")
        surface_id = require_string(raw_surface, "id", "surface entry")
        label = f"surface {surface_id}"
        reject_unknown_fields(raw_surface, {"id", "kind", "scope", "operations"}, label)
        if surface_id in seen_surface_ids:
            raise ContractError(f"duplicate surface id: {surface_id}")
        seen_surface_ids.add(surface_id)
        require_string(raw_surface, "kind", label)
        require_string(raw_surface, "scope", label)
        operations = raw_surface.get("operations")
        if not isinstance(operations, list) or not operations:
            raise ContractError(f"{label} must declare at least one operation/member")
        seen_operation_ids: set[str] = set()
        for raw_operation in operations:
            if not isinstance(raw_operation, dict):
                raise ContractError(f"{label} operation must be an object")
            operation_id = require_string(raw_operation, "id", f"{label} operation")
            operation_label = f"{label} operation {operation_id}"
            reject_unknown_fields(
                raw_operation,
                {
                    "id",
                    "symbols",
                    "granularity",
                    "state",
                    "release_admitted",
                    "qualification",
                    "evidence_ids",
                    "exclusions",
                    "blockers",
                },
                operation_label,
            )
            if operation_id in seen_operation_ids:
                raise ContractError(f"duplicate operation id in {label}: {operation_id}")
            seen_operation_ids.add(operation_id)
            symbols = require_string_list(raw_operation, "symbols", operation_label)
            for symbol in symbols:
                if symbol in seen_symbols:
                    raise ContractError(f"duplicate operation/member symbol: {symbol}")
                seen_symbols.add(symbol)
            granularity = require_string(raw_operation, "granularity", operation_label)
            if granularity not in {"operation", "member", "surface"}:
                raise ContractError(f"{operation_label} has unknown granularity: {granularity}")
            state = require_state(raw_operation, operation_label)
            release_admitted = require_bool(raw_operation, "release_admitted", operation_label)
            require_string(raw_operation, "qualification", operation_label)
            require_string_list(raw_operation, "exclusions", operation_label, allow_empty=True)
            blockers = require_string_list(raw_operation, "blockers", operation_label, allow_empty=True)
            for blocker in blockers:
                if BEAD_ID.fullmatch(blocker) is None:
                    raise ContractError(f"{operation_label} has invalid blocker Bead id: {blocker}")
            evidence_ids = validate_evidence_references(raw_operation, operation_label, evidence)
            levels = {evidence[evidence_id]["level"] for evidence_id in evidence_ids}
            if state == "semantic-diff-supported" and "semantic-diff" not in levels:
                raise ContractError(f"{operation_label} lacks semantic-diff evidence")
            if state == "compile-go-test-run-supported" and not (levels & STRONG_EVIDENCE_LEVELS):
                raise ContractError(f"{operation_label} lacks compile-go-test-run evidence")
            if release_admitted:
                if state not in RELEASE_ELIGIBLE_STATES:
                    raise ContractError(f"{operation_label} cannot be release admitted with state {state}")
                if not (levels & STRONG_EVIDENCE_LEVELS):
                    raise ContractError(f"{operation_label} lacks executable release evidence")
            elif state in {"experimental", "compatibility-only", "excluded"}:
                if not blockers and not raw_operation["exclusions"]:
                    raise ContractError(f"{operation_label} must declare a blocker or exclusion")
    return surfaces


def validate_trust_and_blockers(source: dict[str, Any]) -> None:
    trust = require_list(source, "trust_assumptions", "compatibility source")
    seen_trust: set[str] = set()
    for entry in trust:
        if not isinstance(entry, dict):
            raise ContractError("trust assumption must be an object")
        entry_id = require_string(entry, "id", "trust assumption")
        reject_unknown_fields(entry, {"id", "statement", "release_required"}, f"trust assumption {entry_id}")
        if entry_id in seen_trust:
            raise ContractError(f"duplicate trust assumption: {entry_id}")
        seen_trust.add(entry_id)
        require_string(entry, "statement", f"trust assumption {entry_id}")
        require_bool(entry, "release_required", f"trust assumption {entry_id}")

    blockers = require_list(source, "known_blockers", "compatibility source")
    seen_blockers: set[str] = set()
    for entry in blockers:
        if not isinstance(entry, dict):
            raise ContractError("known blocker must be an object")
        blocker_id = require_string(entry, "id", "known blocker")
        reject_unknown_fields(entry, {"id", "scope", "release_blocker_for"}, f"known blocker {blocker_id}")
        if BEAD_ID.fullmatch(blocker_id) is None:
            raise ContractError(f"known blocker has invalid Bead id: {blocker_id}")
        if blocker_id in seen_blockers:
            raise ContractError(f"duplicate known blocker: {blocker_id}")
        seen_blockers.add(blocker_id)
        require_string(entry, "scope", f"known blocker {blocker_id}")
        require_string(entry, "release_blocker_for", f"known blocker {blocker_id}")

    operation_blockers = {
        blocker
        for surface in source["surfaces"]
        for operation in surface["operations"]
        for blocker in operation["blockers"]
    }
    missing = sorted(operation_blockers - seen_blockers)
    if missing:
        raise ContractError(f"operation blockers missing from known_blockers: {missing}")


def validate_release_claim(source: dict[str, Any]) -> None:
    claim = source["release_claim"]
    platforms = {entry["id"]: entry for entry in source["platforms"]}
    presets = {entry["id"]: entry for entry in source["presets"]}
    platform_id = claim["admitted_platform"]
    preset_id = claim["admitted_preset"]
    if platform_id not in platforms or not platforms[platform_id]["release_admitted"]:
        raise ContractError(f"release claim admitted_platform is missing or not admitted: {platform_id}")
    if preset_id not in presets or not presets[preset_id]["release_admitted"]:
        raise ContractError(f"release claim admitted_preset is missing or not admitted: {preset_id}")


def derive_portable_inventory(portable: dict[str, Any]) -> dict[str, Any]:
    if portable.get("schema_version") != 1:
        raise ContractError("portable stdlib inventory schema_version must be 1")
    raw_modules = portable.get("modules")
    if not isinstance(raw_modules, list) or not raw_modules:
        raise ContractError("portable stdlib inventory must declare modules")
    modules: list[dict[str, Any]] = []
    seen_modules: set[str] = set()
    for raw_module in raw_modules:
        if not isinstance(raw_module, dict):
            raise ContractError("portable stdlib module entry must be an object")
        module = require_string(raw_module, "module", "portable stdlib module")
        if module in seen_modules:
            raise ContractError(f"duplicate portable stdlib module: {module}")
        seen_modules.add(module)
        source_status = require_string(raw_module, "status", f"portable stdlib module {module}")
        if source_status not in PORTABLE_STATUS_MAP:
            raise ContractError(f"portable stdlib module {module} has unknown status: {source_status}")
        coverage_evidence = raw_module.get("coverage_evidence", [])
        if not isinstance(coverage_evidence, list):
            raise ContractError(f"portable stdlib module {module}.coverage_evidence must be an array")
        modules.append(
            {
                "module": module,
                "source_evidence_tier": source_status,
                "state": PORTABLE_STATUS_MAP[source_status],
                "admission": "module-evidence-only",
                "release_admitted": False,
                "coverage_evidence": coverage_evidence,
                "notes": raw_module.get("notes", ""),
            }
        )
    counts = Counter(entry["state"] for entry in modules)
    return {
        "source": "test/portable_stdlib_inventory.json",
        "rule": "Module evidence never admits every member. Release admission comes only from the operation/member inventory in surfaces.",
        "status_counts": {state: counts[state] for state in sorted(counts)},
        "modules": modules,
    }


def build_manifest(
    source_path: Path,
    toolchain_path: Path,
    portable_path: Path,
) -> dict[str, Any]:
    source = load_json(source_path, "compatibility support source")
    toolchain = load_json(toolchain_path, "toolchain policy")
    portable = load_json(portable_path, "portable stdlib inventory")
    package = load_json(REPO_ROOT / "package.json", "package.json")

    validate_source_header(source)
    evidence = validate_evidence(source, package)
    platforms = validate_platforms_or_presets(source, "platforms", evidence)
    presets = validate_platforms_or_presets(source, "presets", evidence)
    surfaces = validate_surfaces(source, evidence)
    validate_trust_and_blockers(source)
    validate_release_claim(source)

    return {
        "schema_version": source["manifest_schema_version"],
        "kind": "haxe.go-compatibility-support-manifest",
        "as_of": source["as_of"],
        "generated_by": GENERATED_BY,
        "inputs": {
            "compatibility_source": display_path(source_path),
            "compatibility_source_sha256": sha256(source_path),
            "toolchain_policy": display_path(toolchain_path),
            "toolchain_policy_sha256": sha256(toolchain_path),
            "portable_stdlib_inventory": display_path(portable_path),
            "portable_stdlib_inventory_sha256": sha256(portable_path),
        },
        "release_claim": source["release_claim"],
        "evidence_states": source["evidence_states"],
        "toolchains": toolchain,
        "platforms": platforms,
        "presets": presets,
        "evidence": source["evidence"],
        "surfaces": surfaces,
        "portable_stdlib": derive_portable_inventory(portable),
        "trust_assumptions": source["trust_assumptions"],
        "known_blockers": source["known_blockers"],
    }


def markdown_escape(value: Any) -> str:
    return str(value).replace("|", "\\|").replace("\n", " ")


def yes_no(value: bool) -> str:
    return "yes" if value else "no"


def render_matrix(manifest: dict[str, Any]) -> str:
    claim = manifest["release_claim"]
    lines = [
        "<!-- generated; edit compatibility-support-source.json and run npm run compatibility:generate -->",
        "# Compatibility and Support Matrix",
        "",
        claim["statement"],
        "",
        "This is the human rendering of",
        "[`compatibility-support-manifest.json`](compatibility-support-manifest.json).",
        "The JSON manifest is authoritative. A module-level green result is evidence inventory,",
        "not blanket admission of every member or error path.",
        "",
        f"Default disposition: {claim['default_disposition']}",
        "",
        "## Evidence states",
        "",
        "| State | Release eligible | Meaning |",
        "| --- | --- | --- |",
    ]
    for state, policy in manifest["evidence_states"].items():
        lines.append(f"| `{state}` | {yes_no(policy['release_eligible'])} | {markdown_escape(policy['meaning'])} |")

    toolchain = manifest["toolchains"]
    lines.extend(
        [
            "",
            "## Toolchains",
            "",
            "| Role | Admitted value |",
            "| --- | --- |",
            f"| Haxe compiler | `{', '.join(toolchain['haxe']['supported_versions'])}` |",
            f"| Generated Go language floor | `{toolchain['go']['generated_language_floor']}` |",
            f"| Go build lines | `{', '.join(toolchain['go']['supported_build_lines'])}`; latest patch required |",
            f"| Node tooling line | `{', '.join(toolchain['node']['supported_tooling_lines'])}` |",
            "",
            "## Platforms and architectures",
            "",
            "| ID | OS | Architecture | State | Release admitted | Qualification |",
            "| --- | --- | --- | --- | --- | --- |",
        ]
    )
    for entry in manifest["platforms"]:
        lines.append(
            f"| `{entry['id']}` | `{entry['os']}` | `{entry['architecture']}` | `{entry['state']}` | "
            f"{yes_no(entry['release_admitted'])} | {markdown_escape(entry['qualification'])} |"
        )

    lines.extend(
        [
            "",
            "## Compatibility presets",
            "",
            "| Preset | Selector | State | Release admitted | Qualification |",
            "| --- | --- | --- | --- | --- |",
        ]
    )
    for entry in manifest["presets"]:
        lines.append(
            f"| `{entry['id']}` | `-D {entry['selector']}` | `{entry['state']}` | "
            f"{yes_no(entry['release_admitted'])} | {markdown_escape(entry['qualification'])} |"
        )

    lines.extend(["", "## Operation/member admission and native-surface inventory", ""])
    for surface in manifest["surfaces"]:
        lines.extend(
            [
                f"### `{surface['id']}`",
                "",
                surface["scope"],
                "",
                "| Symbol | Granularity | State | Release admitted | Evidence | Qualification |",
                "| --- | --- | --- | --- | --- | --- |",
            ]
        )
        for operation in surface["operations"]:
            evidence = ", ".join(f"`{item}`" for item in operation["evidence_ids"])
            for symbol in operation["symbols"]:
                lines.append(
                    f"| `{markdown_escape(symbol)}` | `{operation['granularity']}` | `{operation['state']}` | "
                    f"{yes_no(operation['release_admitted'])} | {evidence} | {markdown_escape(operation['qualification'])} |"
                )

    counts = manifest["portable_stdlib"]["status_counts"]
    lines.extend(
        [
            "",
            "## Portable stdlib module evidence",
            "",
            manifest["portable_stdlib"]["rule"],
            "",
            "| Derived state | Module count |",
            "| --- | ---: |",
        ]
    )
    for state, count in counts.items():
        lines.append(f"| `{state}` | {count} |")

    lines.extend(["", "## Trust assumptions", ""])
    for entry in manifest["trust_assumptions"]:
        lines.append(f"- `{entry['id']}`: {entry['statement']}")

    lines.extend(["", "## Known blockers", "", "| Bead | Scope | Blocks |", "| --- | --- | --- |"])
    for entry in manifest["known_blockers"]:
        lines.append(f"| `{entry['id']}` | {markdown_escape(entry['scope'])} | {markdown_escape(entry['release_blocker_for'])} |")
    lines.append("")
    return "\n".join(lines)


def render_release_status(manifest: dict[str, Any]) -> str:
    claim = manifest["release_claim"]
    admitted_platform = next(entry for entry in manifest["platforms"] if entry["id"] == claim["admitted_platform"])
    admitted_preset = next(entry for entry in manifest["presets"] if entry["id"] == claim["admitted_preset"])
    admitted_symbols = [
        symbol
        for surface in manifest["surfaces"]
        for operation in surface["operations"]
        if operation["release_admitted"]
        for symbol in operation["symbols"]
    ]
    excluded_operations = [
        (surface["id"], operation)
        for surface in manifest["surfaces"]
        for operation in surface["operations"]
        if not operation["release_admitted"]
    ]
    lines = [
        "<!-- generated; edit compatibility-support-source.json and run npm run compatibility:generate -->",
        "# Compatibility Release Status",
        "",
        claim["statement"],
        "",
        "This status is generated from the same governed source as the machine manifest and",
        "must be used as the compatibility paragraph in release notes.",
        "",
        "## Admitted scope",
        "",
        f"- preset: `{admitted_preset['id']}` (`{admitted_preset['state']}`);",
        f"- platform: `{admitted_platform['os']}/{admitted_platform['architecture']}`;",
        f"- named operations/members: {len(admitted_symbols)};",
        "- toolchains: exact Haxe version and latest patched supported Go/Node lines from `toolchain-policy.json`;",
        "- trust: reviewed application source, locked tooling, application-controlled local file/process boundaries, and application-controlled, pre-resolved numeric TCP endpoints.",
        "",
        "No module-level inventory row expands this scope. Unlisted operations and error paths",
        "take the default excluded disposition.",
        "",
        "## Not admitted by this release scope",
        "",
    ]
    for surface_id, operation in excluded_operations:
        blockers = f"; blockers: {', '.join(operation['blockers'])}" if operation["blockers"] else ""
        symbols = ", ".join(f"`{symbol}`" for symbol in operation["symbols"])
        lines.append(f"- `{surface_id}` / {symbols}: `{operation['state']}`{blockers}.")
    lines.extend(
        [
            "- non-canonical operating-system/architecture combinations and moving runner identities;",
            "- a stable 1.x compatibility claim or published validated beta-baseline artifact.",
            "",
            "## Known blockers",
            "",
        ]
    )
    for blocker in manifest["known_blockers"]:
        lines.append(
            f"- `{blocker['id']}`: {markdown_escape(blocker['scope'])} "
            f"(blocks {markdown_escape(blocker['release_blocker_for'])})."
        )
    lines.extend(
        [
            "",
            "## Machine authority",
            "",
            "- `docs/compatibility-support-manifest.json`",
            "- public SemVer boundary: `docs/public-contract.md`",
            "- lifecycle and stable admission: `docs/semver-lifecycle-policy.md`",
            "- verify with `npm run compatibility:verify`",
            "",
        ]
    )
    rendered = "\n".join(lines)
    for phrase in claim["forbidden_phrases"]:
        if phrase.lower() in rendered.lower():
            raise ContractError(f"generated release status contains forbidden phrase: {phrase}")
    return rendered


def json_text(value: dict[str, Any]) -> str:
    return json.dumps(value, indent=2, ensure_ascii=False) + "\n"


def check_or_write(path: Path, content: str, check: bool, stale: list[str]) -> None:
    if check:
        try:
            current = path.read_text(encoding="utf-8")
        except FileNotFoundError:
            stale.append(f"missing generated artifact: {display_path(path)}")
            return
        if current != content:
            stale.append(f"stale generated artifact: {display_path(path)}")
        return
    path.parent.mkdir(parents=True, exist_ok=True)
    path.write_text(content, encoding="utf-8")


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(description="Generate and verify Haxe.Go compatibility support artifacts")
    parser.add_argument("--source", type=Path, default=DEFAULT_SOURCE)
    parser.add_argument("--toolchain", type=Path, default=DEFAULT_TOOLCHAIN)
    parser.add_argument("--portable-inventory", type=Path, default=DEFAULT_PORTABLE_INVENTORY)
    parser.add_argument("--manifest", type=Path, default=DEFAULT_MANIFEST)
    parser.add_argument("--matrix-doc", type=Path, default=DEFAULT_MATRIX_DOC)
    parser.add_argument("--release-doc", type=Path, default=DEFAULT_RELEASE_DOC)
    parser.add_argument("--check", action="store_true", help="Fail unless committed outputs are current")
    return parser.parse_args()


def main() -> int:
    args = parse_args()
    try:
        manifest = build_manifest(args.source, args.toolchain, args.portable_inventory)
        outputs = [
            (args.manifest, json_text(manifest)),
            (args.matrix_doc, render_matrix(manifest)),
            (args.release_doc, render_release_status(manifest)),
        ]
        stale: list[str] = []
        for path, content in outputs:
            check_or_write(path, content, args.check, stale)
        if stale:
            raise ContractError("; ".join(stale) + "; run npm run compatibility:generate")
    except ContractError as error:
        print(f"[compatibility-support] error: {error}", file=sys.stderr)
        return 1
    verb = "verified" if args.check else "generated"
    print(f"[compatibility-support] {verb}: {display_path(args.manifest)}")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
