#!/usr/bin/env python3

from __future__ import annotations

import json
from pathlib import Path
from typing import Any

ROOT = Path(__file__).resolve().parent.parent
CONTRACT_SCHEMA = 8
OPTIMIZER_SCHEMA = 6
RUNTIME_SCHEMA = 2


def load_json(path: Path) -> dict[str, Any]:
    if not path.exists():
        raise SystemExit(f"missing required artifact: {path}")
    data = json.loads(path.read_text(encoding="utf-8"))
    if not isinstance(data, dict):
        raise SystemExit(f"invalid json root (expected object): {path}")
    return data


def require_keys(obj: dict[str, Any], keys: list[str], label: str) -> None:
    missing = [key for key in keys if key not in obj]
    if missing:
        raise SystemExit(f"{label}: missing required keys: {', '.join(missing)}")


def require_schema(obj: dict[str, Any], expected: int, label: str) -> None:
    value = obj.get("schemaVersion")
    if value != expected:
        raise SystemExit(f"{label}: schemaVersion expected {expected}, got {value!r}")


def require_nonempty_list(value: Any, label: str) -> list[Any]:
    if not isinstance(value, list) or len(value) == 0:
        raise SystemExit(f"{label}: expected non-empty array")
    return value


def require_int(value: Any, label: str) -> int:
    if not isinstance(value, int):
        raise SystemExit(f"{label}: expected int, got {type(value).__name__}")
    return value


def require_counter_eq(obj: dict[str, Any], key: str, expected: int, label: str) -> None:
    value = require_int(obj.get(key), f"{label}.{key}")
    if value != expected:
        raise SystemExit(f"{label}.{key}: expected {expected}, got {value}")


def require_counter_gt(obj: dict[str, Any], key: str, threshold: int, label: str) -> None:
    value = require_int(obj.get(key), f"{label}.{key}")
    if value <= threshold:
        raise SystemExit(f"{label}.{key}: expected > {threshold}, got {value}")


def require_optimizer_capabilities(value: Any, label: str) -> None:
    if not isinstance(value, list):
        raise SystemExit(f"{label}: expected array")
    for index, entry in enumerate(value):
        if not isinstance(entry, dict):
            raise SystemExit(f"{label}[{index}]: expected object")
        for key in ("id", "attempts", "successes", "fallbacks", "fallbackReasonCounts"):
            if key not in entry:
                raise SystemExit(f"{label}[{index}]: missing key `{key}`")
        if not isinstance(entry["id"], str) or entry["id"].strip() == "":
            raise SystemExit(f"{label}[{index}].id: expected non-empty string")
        require_int(entry["attempts"], f"{label}[{index}].attempts")
        require_int(entry["successes"], f"{label}[{index}].successes")
        require_int(entry["fallbacks"], f"{label}[{index}].fallbacks")
        reasons = entry["fallbackReasonCounts"]
        if not isinstance(reasons, list):
            raise SystemExit(f"{label}[{index}].fallbackReasonCounts: expected array")
        for reason_index, reason in enumerate(reasons):
            if not isinstance(reason, dict):
                raise SystemExit(f"{label}[{index}].fallbackReasonCounts[{reason_index}]: expected object")
            if "kind" not in reason or "count" not in reason:
                raise SystemExit(
                    f"{label}[{index}].fallbackReasonCounts[{reason_index}]: missing `kind` or `count`"
                )
            if not isinstance(reason["kind"], str) or reason["kind"].strip() == "":
                raise SystemExit(f"{label}[{index}].fallbackReasonCounts[{reason_index}].kind: expected string")
            require_int(reason["count"], f"{label}[{index}].fallbackReasonCounts[{reason_index}].count")


def require_capability(capabilities: Any, capability_id: str, label: str) -> dict[str, Any]:
    if not isinstance(capabilities, list):
        raise SystemExit(f"{label}: expected capabilities array")
    for entry in capabilities:
        if isinstance(entry, dict) and entry.get("id") == capability_id:
            return entry
    raise SystemExit(f"{label}: missing capability `{capability_id}`")


def require_fallback_reason(capability: dict[str, Any], kind: str, expected: int, label: str) -> None:
    reasons = capability.get("fallbackReasonCounts")
    if not isinstance(reasons, list):
        raise SystemExit(f"{label}.fallbackReasonCounts: expected array")
    for reason in reasons:
        if isinstance(reason, dict) and reason.get("kind") == kind:
            count = require_int(reason.get("count"), f"{label}.fallbackReasonCounts.{kind}")
            if count != expected:
                raise SystemExit(f"{label}.fallbackReasonCounts.{kind}: expected {expected}, got {count}")
            return
    raise SystemExit(f"{label}.fallbackReasonCounts: missing reason `{kind}`")


def require_reason_entries(entries: list[Any], expected_source_prefix: str, label: str) -> None:
    for index, entry in enumerate(entries):
        if not isinstance(entry, dict):
            raise SystemExit(f"{label}: reason entry #{index} is not an object")
        for key in ("pass", "source", "reason"):
            if key not in entry:
                raise SystemExit(f"{label}: reason entry #{index} missing key `{key}`")
            if not isinstance(entry[key], str) or entry[key].strip() == "":
                raise SystemExit(f"{label}: reason entry #{index} has invalid `{key}`")
        source_value = entry["source"]
        if not source_value.startswith(expected_source_prefix):
            raise SystemExit(
                f"{label}: reason entry #{index} source expected prefix `{expected_source_prefix}`, got `{source_value}`"
            )


def require_runtime_reason_entries(entries: Any, label: str) -> None:
    if not isinstance(entries, list) or len(entries) == 0:
        raise SystemExit(f"{label}: expected non-empty reasons array")
    seen_kinds: set[str] = set()
    for index, entry in enumerate(entries):
        if not isinstance(entry, dict):
            raise SystemExit(f"{label}[{index}]: expected object")
        for key in ("feature", "sourceKind", "source"):
            if key not in entry:
                raise SystemExit(f"{label}[{index}]: missing key `{key}`")
            if not isinstance(entry[key], str) or entry[key].strip() == "":
                raise SystemExit(f"{label}[{index}].{key}: expected non-empty string")
        seen_kinds.add(entry["sourceKind"])
    required_kinds = {"baseline", "class_usage", "dependency_edge", "manual_define"}
    missing = sorted(required_kinds - seen_kinds)
    if missing:
        raise SystemExit(f"{label}: missing expected sourceKind values: {', '.join(missing)}")


def main() -> int:
    contract_path = ROOT / "test/snapshot/core/report_artifacts_auto_mode/intended/profile_contract.json"
    optimizer_path = ROOT / "test/snapshot/core/report_artifacts_auto_mode/intended/optimizer_plan.json"
    optimizer_legacy_path = (
        ROOT / "test/snapshot/core/optimizer_plan_pass_selection_legacy_granular/intended/optimizer_plan.json"
    )
    optimizer_typed_path = (
        ROOT / "test/snapshot/core/report_artifacts_auto_collections_result/intended/optimizer_plan.json"
    )
    optimizer_fallback_path = (
        ROOT / "test/snapshot/core/optimizer_plan_auto_collections_result_fallback/intended/optimizer_plan.json"
    )
    optimizer_string_fastpath_path = (
        ROOT / "test/snapshot/core/optimizer_plan_string_fastpath_enabled/intended/optimizer_plan.json"
    )
    optimizer_string_legacy_path = (
        ROOT / "test/snapshot/core/optimizer_plan_string_fastpath_disabled/intended/optimizer_plan.json"
    )
    runtime_path = ROOT / "test/snapshot/core/report_artifacts_runtime_reason_provenance/intended/hxrt_plan.json"

    contract = load_json(contract_path)
    require_schema(contract, CONTRACT_SCHEMA, str(contract_path))
    require_keys(
        contract,
        [
            "policyPreset",
            "semanticBoundarySource",
            "nativeAuthorityPolicy",
            "nativeAuthorityPolicySource",
            "nativeSpecializationPolicy",
            "nativeSpecializationPolicySource",
            "nativeFallbackPolicy",
            "nativeFallbackPolicySource",
            "autoLoweringMode",
            "portableNativeImportScanMode",
            "portableNativeImportHitCount",
            "portableNativeImportTypedHitCount",
            "portableNativeImportScannerHitCount",
            "portableNativeImportHits",
            "portableNativeImportTypedHits",
            "portableNativeImportScannerHits",
            "loweringDecisionCount",
            "loweringDecisionAttemptCount",
            "loweringDecisionSuccessCount",
            "loweringDecisionFallbackCount",
            "loweringDecisions",
            "nativeBoundaryModules",
            "nativeFallbackEventCount",
            "nativeFallbackBoundaryEventCount",
            "nativeFallbackNonBoundaryEventCount",
            "nativeFallbackEventsByModule",
            "nativeFallbackEvents",
        ],
        str(contract_path),
    )

    optimizer = load_json(optimizer_path)
    require_schema(optimizer, OPTIMIZER_SCHEMA, str(optimizer_path))
    require_keys(
        optimizer,
        [
            "policyPreset",
            "nativeSpecializationPolicy",
            "nativeSpecializationPolicySource",
            "autoLoweringMode",
            "optimizationPreset",
            "goAstPassSelectionSource",
            "goAstPasses",
            "goAstPassSelectionReasons",
            "goCollectionsTypedLowerings",
            "goCollectionsTypedFallbacks",
            "goResultTypedLowerings",
            "goResultTypedFallbacks",
            "loweringFallbackLaneCount",
            "loweringFallbackNonLaneCount",
            "loweringFallbackBoundaryCount",
            "loweringFallbackNonBoundaryCount",
            "autoLoweringCapabilities",
        ],
        str(optimizer_path),
    )
    optimizer_source = optimizer["goAstPassSelectionSource"]
    if not isinstance(optimizer_source, str) or not optimizer_source.startswith("planner"):
        raise SystemExit(
            f"{optimizer_path}: goAstPassSelectionSource expected prefix `planner`, got `{optimizer_source}`"
        )
    planner_reasons = require_nonempty_list(optimizer["goAstPassSelectionReasons"], str(optimizer_path))
    require_reason_entries(planner_reasons, "planner", str(optimizer_path))
    require_optimizer_capabilities(optimizer["autoLoweringCapabilities"], str(optimizer_path) + ".autoLoweringCapabilities")

    optimizer_legacy = load_json(optimizer_legacy_path)
    require_schema(optimizer_legacy, OPTIMIZER_SCHEMA, str(optimizer_legacy_path))
    if optimizer_legacy.get("goAstPassSelectionSource") != "legacy_granular_bundle":
        raise SystemExit(
            f"{optimizer_legacy_path}: goAstPassSelectionSource expected `legacy_granular_bundle`, got `{optimizer_legacy.get('goAstPassSelectionSource')}`"
        )
    legacy_reasons = require_nonempty_list(optimizer_legacy.get("goAstPassSelectionReasons"), str(optimizer_legacy_path))
    require_reason_entries(legacy_reasons, "legacy_granular_bundle", str(optimizer_legacy_path))
    require_optimizer_capabilities(
        optimizer_legacy.get("autoLoweringCapabilities"),
        str(optimizer_legacy_path) + ".autoLoweringCapabilities",
    )

    optimizer_typed = load_json(optimizer_typed_path)
    require_schema(optimizer_typed, OPTIMIZER_SCHEMA, str(optimizer_typed_path))
    require_optimizer_capabilities(
        optimizer_typed.get("autoLoweringCapabilities"),
        str(optimizer_typed_path) + ".autoLoweringCapabilities",
    )
    typed_capabilities = optimizer_typed["autoLoweringCapabilities"]
    typed_collections = require_capability(
        typed_capabilities, "go.collections.typed", str(optimizer_typed_path) + ".autoLoweringCapabilities"
    )
    typed_result = require_capability(
        typed_capabilities, "go.result.typed", str(optimizer_typed_path) + ".autoLoweringCapabilities"
    )
    require_counter_gt(typed_collections, "attempts", 0, str(optimizer_typed_path) + ".go.collections.typed")
    require_counter_gt(typed_collections, "successes", 0, str(optimizer_typed_path) + ".go.collections.typed")
    require_counter_eq(typed_collections, "fallbacks", 0, str(optimizer_typed_path) + ".go.collections.typed")
    if typed_collections.get("fallbackReasonCounts") != []:
        raise SystemExit(
            f"{optimizer_typed_path}.go.collections.typed.fallbackReasonCounts: expected empty for typed-success contract"
        )
    require_counter_gt(typed_result, "attempts", 0, str(optimizer_typed_path) + ".go.result.typed")
    require_counter_gt(typed_result, "successes", 0, str(optimizer_typed_path) + ".go.result.typed")
    require_counter_eq(typed_result, "fallbacks", 0, str(optimizer_typed_path) + ".go.result.typed")
    if typed_result.get("fallbackReasonCounts") != []:
        raise SystemExit(f"{optimizer_typed_path}.go.result.typed.fallbackReasonCounts: expected empty for typed-success contract")
    require_counter_gt(optimizer_typed, "goCollectionsTypedLowerings", 0, str(optimizer_typed_path))
    require_counter_eq(optimizer_typed, "goCollectionsTypedFallbacks", 0, str(optimizer_typed_path))
    require_counter_gt(optimizer_typed, "goResultTypedLowerings", 0, str(optimizer_typed_path))
    require_counter_eq(optimizer_typed, "goResultTypedFallbacks", 0, str(optimizer_typed_path))
    require_counter_eq(optimizer_typed, "loweringFallbackNonLaneCount", 0, str(optimizer_typed_path))

    optimizer_fallback = load_json(optimizer_fallback_path)
    require_schema(optimizer_fallback, OPTIMIZER_SCHEMA, str(optimizer_fallback_path))
    require_optimizer_capabilities(
        optimizer_fallback.get("autoLoweringCapabilities"),
        str(optimizer_fallback_path) + ".autoLoweringCapabilities",
    )
    fallback_capabilities = optimizer_fallback["autoLoweringCapabilities"]
    fallback_collections = require_capability(
        fallback_capabilities, "go.collections.typed", str(optimizer_fallback_path) + ".autoLoweringCapabilities"
    )
    fallback_result = require_capability(
        fallback_capabilities, "go.result.typed", str(optimizer_fallback_path) + ".autoLoweringCapabilities"
    )
    require_counter_gt(fallback_collections, "attempts", 0, str(optimizer_fallback_path) + ".go.collections.typed")
    require_counter_gt(fallback_collections, "fallbacks", 0, str(optimizer_fallback_path) + ".go.collections.typed")
    if not fallback_collections.get("fallbackReasonCounts"):
        raise SystemExit(
            f"{optimizer_fallback_path}.go.collections.typed.fallbackReasonCounts: expected non-empty fallback reasons"
        )
    require_counter_gt(fallback_result, "attempts", 0, str(optimizer_fallback_path) + ".go.result.typed")
    require_counter_gt(fallback_result, "fallbacks", 0, str(optimizer_fallback_path) + ".go.result.typed")
    if not fallback_result.get("fallbackReasonCounts"):
        raise SystemExit(f"{optimizer_fallback_path}.go.result.typed.fallbackReasonCounts: expected non-empty fallback reasons")
    require_counter_eq(optimizer_fallback, "goCollectionsTypedLowerings", 0, str(optimizer_fallback_path))
    require_counter_gt(optimizer_fallback, "goCollectionsTypedFallbacks", 0, str(optimizer_fallback_path))
    require_counter_eq(optimizer_fallback, "goResultTypedLowerings", 0, str(optimizer_fallback_path))
    require_counter_gt(optimizer_fallback, "goResultTypedFallbacks", 0, str(optimizer_fallback_path))
    require_counter_gt(optimizer_fallback, "loweringFallbackNonLaneCount", 0, str(optimizer_fallback_path))

    optimizer_string_fastpath = load_json(optimizer_string_fastpath_path)
    require_schema(optimizer_string_fastpath, OPTIMIZER_SCHEMA, str(optimizer_string_fastpath_path))
    require_optimizer_capabilities(
        optimizer_string_fastpath.get("autoLoweringCapabilities"),
        str(optimizer_string_fastpath_path) + ".autoLoweringCapabilities",
    )
    string_fastpath_capability = require_capability(
        optimizer_string_fastpath["autoLoweringCapabilities"],
        "go.string.typed",
        str(optimizer_string_fastpath_path) + ".autoLoweringCapabilities",
    )
    string_fastpath_attempts = (
        optimizer_string_fastpath["stringInstanceTypedLowerings"]
        + optimizer_string_fastpath["stringLengthFieldTypedLowerings"]
    )
    require_counter_eq(
        string_fastpath_capability,
        "attempts",
        string_fastpath_attempts,
        str(optimizer_string_fastpath_path) + ".go.string.typed",
    )
    require_counter_eq(
        string_fastpath_capability,
        "successes",
        string_fastpath_attempts,
        str(optimizer_string_fastpath_path) + ".go.string.typed",
    )
    require_counter_eq(string_fastpath_capability, "fallbacks", 0, str(optimizer_string_fastpath_path) + ".go.string.typed")
    if string_fastpath_capability.get("fallbackReasonCounts") != []:
        raise SystemExit(f"{optimizer_string_fastpath_path}.go.string.typed.fallbackReasonCounts: expected empty")

    optimizer_string_legacy = load_json(optimizer_string_legacy_path)
    require_schema(optimizer_string_legacy, OPTIMIZER_SCHEMA, str(optimizer_string_legacy_path))
    require_optimizer_capabilities(
        optimizer_string_legacy.get("autoLoweringCapabilities"),
        str(optimizer_string_legacy_path) + ".autoLoweringCapabilities",
    )
    string_legacy_capability = require_capability(
        optimizer_string_legacy["autoLoweringCapabilities"],
        "go.string.typed",
        str(optimizer_string_legacy_path) + ".autoLoweringCapabilities",
    )
    string_legacy_attempts = (
        optimizer_string_legacy["stringInstanceLegacyLowerings"]
        + optimizer_string_legacy["stringLengthFieldLegacyLowerings"]
    )
    require_counter_eq(
        string_legacy_capability,
        "attempts",
        string_legacy_attempts,
        str(optimizer_string_legacy_path) + ".go.string.typed",
    )
    require_counter_eq(string_legacy_capability, "successes", 0, str(optimizer_string_legacy_path) + ".go.string.typed")
    require_counter_eq(
        string_legacy_capability,
        "fallbacks",
        string_legacy_attempts,
        str(optimizer_string_legacy_path) + ".go.string.typed",
    )
    require_fallback_reason(
        string_legacy_capability,
        "optimizer_preset_disabled",
        string_legacy_attempts,
        str(optimizer_string_legacy_path) + ".go.string.typed",
    )

    runtime = load_json(runtime_path)
    require_schema(runtime, RUNTIME_SCHEMA, str(runtime_path))
    require_keys(
        runtime,
        [
            "contract",
            "policyPreset",
            "semanticBoundarySource",
            "mode",
            "selectiveEnabled",
            "fullCopy",
            "inferenceDisabled",
            "manualFeatures",
            "inferredFeatures",
            "selectedFeatures",
            "files",
            "reasons",
        ],
        str(runtime_path),
    )
    require_runtime_reason_entries(runtime["reasons"], str(runtime_path) + ".reasons")

    print("[PASS] auto planner report schema gate (contract/runtime/optimizer artifacts)")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
