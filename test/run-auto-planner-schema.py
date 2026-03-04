#!/usr/bin/env python3

from __future__ import annotations

import json
from pathlib import Path
from typing import Any

ROOT = Path(__file__).resolve().parent.parent


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
    runtime_path = ROOT / "test/snapshot/core/report_artifacts_basic/intended/hxrt_plan.json"

    contract = load_json(contract_path)
    require_schema(contract, 7, str(contract_path))
    require_keys(
        contract,
        [
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
        ],
        str(contract_path),
    )

    optimizer = load_json(optimizer_path)
    require_schema(optimizer, 5, str(optimizer_path))
    require_keys(
        optimizer,
        [
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
    require_schema(optimizer_legacy, 5, str(optimizer_legacy_path))
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
    require_schema(optimizer_typed, 5, str(optimizer_typed_path))
    require_optimizer_capabilities(
        optimizer_typed.get("autoLoweringCapabilities"),
        str(optimizer_typed_path) + ".autoLoweringCapabilities",
    )
    require_counter_gt(optimizer_typed, "goCollectionsTypedLowerings", 0, str(optimizer_typed_path))
    require_counter_eq(optimizer_typed, "goCollectionsTypedFallbacks", 0, str(optimizer_typed_path))
    require_counter_gt(optimizer_typed, "goResultTypedLowerings", 0, str(optimizer_typed_path))
    require_counter_eq(optimizer_typed, "goResultTypedFallbacks", 0, str(optimizer_typed_path))
    require_counter_eq(optimizer_typed, "loweringFallbackNonLaneCount", 0, str(optimizer_typed_path))

    optimizer_fallback = load_json(optimizer_fallback_path)
    require_schema(optimizer_fallback, 5, str(optimizer_fallback_path))
    require_optimizer_capabilities(
        optimizer_fallback.get("autoLoweringCapabilities"),
        str(optimizer_fallback_path) + ".autoLoweringCapabilities",
    )
    require_counter_eq(optimizer_fallback, "goCollectionsTypedLowerings", 0, str(optimizer_fallback_path))
    require_counter_gt(optimizer_fallback, "goCollectionsTypedFallbacks", 0, str(optimizer_fallback_path))
    require_counter_eq(optimizer_fallback, "goResultTypedLowerings", 0, str(optimizer_fallback_path))
    require_counter_gt(optimizer_fallback, "goResultTypedFallbacks", 0, str(optimizer_fallback_path))
    require_counter_gt(optimizer_fallback, "loweringFallbackNonLaneCount", 0, str(optimizer_fallback_path))

    runtime = load_json(runtime_path)
    require_schema(runtime, 1, str(runtime_path))
    require_keys(
        runtime,
        [
            "contract",
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

    print("[PASS] auto planner report schema gate (contract/runtime/optimizer artifacts)")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
