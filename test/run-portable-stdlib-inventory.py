#!/usr/bin/env python3

from __future__ import annotations

import argparse
import json
from pathlib import Path
from typing import Any

ROOT = Path(__file__).resolve().parent.parent
CACHE_ROOT = ROOT / "test" / ".test-cache"
FULL_MODULES_FILE = ROOT / "test" / "upstream_std_modules_full.txt"
STRICT_SWEEP_MODULES_FILE = ROOT / "test" / "upstream_std_modules.txt"
INVENTORY_FILE = ROOT / "test" / "portable_stdlib_inventory.json"
PROMOTIONS_FILE = ROOT / "test" / "portable_parity_promotions.json"
SUMMARY_JSON = CACHE_ROOT / "portable_stdlib_inventory_summary.json"
SUMMARY_MD = CACHE_ROOT / "portable_stdlib_inventory_summary.md"

ALLOWED_STATUS = {"unsupported", "compile-only", "snapshot", "semantic-diff"}
ALLOWED_OWNER = {"unassigned", "mixed", "compiler_shim", "runtime_hxrt", "staged_std"}

SEMANTIC_DIFF_EXPLICIT = {
    "Array",
    "Class",
    "Date",
    "DateTools",
    "EReg",
    "Enum",
    "EnumValue",
    "IntIterator",
    "Lambda",
    "Math",
    "Reflect",
    "String",
    "StringBuf",
    "Std",
    "StringTools",
    "Type",
    "haxe.Exception",
    "haxe.Int32",
    "haxe.Int64",
    "haxe.Int64Helper",
    "haxe.Json",
    "haxe.PosInfos",
    "haxe.Serializer",
    "haxe.Unserializer",
    "haxe.atomic.AtomicBool",
    "haxe.atomic.AtomicInt",
    "haxe.atomic.AtomicObject",
    "haxe.ds.EnumValueMap",
    "haxe.ds.IntMap",
    "haxe.ds.List",
    "haxe.ds.Map",
    "haxe.ds.Option",
    "haxe.ds.ObjectMap",
    "haxe.ds.ReadOnlyArray",
    "haxe.ds.StringMap",
    "haxe.ds.Vector",
    "haxe.iterators.HashMapKeyValueIterator",
    "haxe.iterators.MapKeyValueIterator",
    "haxe.format.JsonParser",
    "haxe.format.JsonPrinter",
    "haxe.io.Bytes",
    "haxe.io.BytesBuffer",
    "haxe.io.BytesInput",
    "haxe.io.BytesOutput",
    "haxe.io.Input",
    "haxe.io.Output",
    "haxe.io.Path",
    "sys.FileSystem",
    "sys.Http",
    "sys.io.File",
    "sys.io.Process",
    "sys.net.Host",
    "sys.net.Socket",
}

SEMANTIC_DIFF_PREFIXES = (
    "haxe.crypto.",
    "haxe.xml.",
    "haxe.zip.",
)

OWNER_OVERRIDES = {
    "haxe.Json": "runtime_hxrt",
    "haxe.format.JsonParser": "runtime_hxrt",
    "haxe.format.JsonPrinter": "runtime_hxrt",
    "haxe.atomic.AtomicBool": "runtime_hxrt",
    "haxe.atomic.AtomicInt": "runtime_hxrt",
    "haxe.atomic.AtomicObject": "runtime_hxrt",
    "sys.io.File": "runtime_hxrt",
    "sys.io.Process": "runtime_hxrt",
    "sys.FileSystem": "runtime_hxrt",
    "EReg": "compiler_shim",
    "haxe.Serializer": "compiler_shim",
    "haxe.Unserializer": "compiler_shim",
    "sys.Http": "compiler_shim",
    "sys.net.Host": "compiler_shim",
    "sys.net.Socket": "compiler_shim",
    "haxe.io.Bytes": "compiler_shim",
    "haxe.io.BytesBuffer": "compiler_shim",
    "haxe.io.BytesInput": "compiler_shim",
    "haxe.io.BytesOutput": "compiler_shim",
    "haxe.io.Input": "compiler_shim",
    "haxe.io.Output": "compiler_shim",
}

PROMOTION_LEVEL_KEYS = ("snapshot", "semantic_diff")


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(
        description="Validate and generate portable stdlib inventory (portable-eligible Haxe 4.3.7 modules)."
    )
    parser.add_argument(
        "--update",
        action="store_true",
        help="Write generated inventory to test/portable_stdlib_inventory.json.",
    )
    return parser.parse_args()


def load_modules(path: Path) -> list[str]:
    modules: list[str] = []
    for raw in path.read_text(encoding="utf-8").splitlines():
        line = raw.strip()
        if not line or line.startswith("#"):
            continue
        modules.append(line)
    return modules


def load_promotions(path: Path, valid_modules: set[str], auto_semantic_modules: set[str]) -> dict[str, set[str]]:
    try:
        raw = json.loads(path.read_text(encoding="utf-8"))
    except FileNotFoundError as exc:
        raise SystemExit(
            "portable parity promotions file missing. "
            "Expected: test/portable_parity_promotions.json"
        ) from exc
    except json.JSONDecodeError as exc:
        raise SystemExit(f"{path.name}: invalid JSON ({exc})") from exc

    if not isinstance(raw, dict):
        raise SystemExit(f"{path.name}: root must be an object")
    if raw.get("schema_version") != 1:
        raise SystemExit(f"{path.name}: schema_version must be 1")

    levels = raw.get("levels")
    if not isinstance(levels, dict):
        raise SystemExit(f"{path.name}: levels must be an object")

    unknown_keys = sorted(set(levels.keys()) - set(PROMOTION_LEVEL_KEYS))
    if unknown_keys:
        raise SystemExit(
            f"{path.name}: unknown levels keys: {', '.join(unknown_keys)} "
            f"(allowed: {', '.join(PROMOTION_LEVEL_KEYS)})"
        )

    normalized: dict[str, set[str]] = {}
    for key in PROMOTION_LEVEL_KEYS:
        raw_modules = levels.get(key, [])
        if not isinstance(raw_modules, list):
            raise SystemExit(f"{path.name}: levels.{key} must be an array")
        if raw_modules != sorted(raw_modules):
            raise SystemExit(f"{path.name}: levels.{key} must be sorted")

        seen: set[str] = set()
        modules: set[str] = set()
        for module in raw_modules:
            if not isinstance(module, str) or not module.strip():
                raise SystemExit(f"{path.name}: levels.{key} contains invalid module entry: {module!r}")
            if module in seen:
                raise SystemExit(f"{path.name}: levels.{key} has duplicate module `{module}`")
            seen.add(module)
            if module not in valid_modules:
                raise SystemExit(f"{path.name}: levels.{key} module `{module}` not found in full module list")
            modules.add(module)
        normalized[key] = modules

    overlap = normalized["snapshot"] & normalized["semantic_diff"]
    if overlap:
        raise SystemExit(
            f"{path.name}: modules cannot be in both snapshot and semantic_diff levels: "
            + ", ".join(sorted(overlap))
        )

    redundant_snapshot = normalized["snapshot"] & auto_semantic_modules
    if redundant_snapshot:
        raise SystemExit(
            f"{path.name}: snapshot promotions cannot include modules already auto-promoted to semantic-diff: "
            + ", ".join(sorted(redundant_snapshot))
        )

    return normalized


def is_semantic_diff_module(module: str) -> bool:
    if module in SEMANTIC_DIFF_EXPLICIT:
        return True
    for prefix in SEMANTIC_DIFF_PREFIXES:
        if module.startswith(prefix):
            return True
    return False


def select_owner(module: str, status: str, in_strict_sweep: bool) -> str:
    if module in OWNER_OVERRIDES:
        return OWNER_OVERRIDES[module]
    if module.startswith(SEMANTIC_DIFF_PREFIXES):
        return "compiler_shim"
    if status == "compile-only" or in_strict_sweep:
        return "mixed"
    return "unassigned"


def module_notes(module: str, status: str, in_strict_sweep: bool) -> str:
    if status == "semantic-diff":
        return (
            "Covered by semantic-diff/runtime contracts in test/semantic_diff and "
            "documented in docs/feature-support-matrix.md."
        )
    if status == "snapshot":
        return "Covered by snapshot-level deterministic generated-code/runtime smoke contracts."
    if status == "compile-only":
        if in_strict_sweep:
            return (
                "Covered by strict upstream stdlib sweep compile/go-test checks "
                "(test/upstream_std_modules.txt)."
            )
        return (
            "Covered by full portable-eligible upstream stdlib sweep compile checks "
            "(test/upstream_std_modules_full.txt); runtime parity contracts are not yet promoted."
        )
    return "Portable-eligible module inventoried; parity promotion is pending."


def module_evidence(status: str, in_full_sweep: bool, in_strict_sweep: bool) -> list[str]:
    evidence: list[str] = []
    if in_full_sweep:
        evidence.append("upstream_sweep:full_compile")
    if in_strict_sweep:
        evidence.append("upstream_sweep:strict_go_test")
    if status == "snapshot":
        evidence.append("snapshot")
    if status == "semantic-diff":
        evidence.append("semantic_diff")
    return evidence


def build_inventory(
    full_modules: list[str], strict_sweep_modules: set[str], promotions: dict[str, set[str]]
) -> dict[str, Any]:
    snapshot_promotions = promotions["snapshot"]
    semantic_promotions = promotions["semantic_diff"]
    modules_payload: list[dict[str, Any]] = []
    for module in sorted(full_modules):
        in_full_sweep = True
        in_strict_sweep = module in strict_sweep_modules
        status = "compile-only"
        if is_semantic_diff_module(module) or module in semantic_promotions:
            status = "semantic-diff"
        elif module in snapshot_promotions:
            status = "snapshot"

        owner = select_owner(module, status, in_strict_sweep)
        entry = {
            "module": module,
            "portable_eligible": True,
            "status": status,
            "owner": owner,
            "in_full_sweep": in_full_sweep,
            "in_strict_sweep": in_strict_sweep,
            "coverage_evidence": module_evidence(status, in_full_sweep, in_strict_sweep),
            "notes": module_notes(module, status, in_strict_sweep),
        }
        modules_payload.append(entry)

    return {
        "schema_version": 1,
        "baseline": {
            "haxe_version": "4.3.7",
            "module_source": "test/upstream_std_modules_full.txt",
            "portable_scope": "portable-eligible modules only; target-specific namespaces excluded",
            "excluded_prefix_examples": ["cpp.*", "java.*", "cs.*", "hl.*", "lua.*", "php.*", "python.*", "js.*"],
        },
        "generated_by": "test/run-portable-stdlib-inventory.py",
        "modules": modules_payload,
    }


def validate_inventory_schema(inventory: dict[str, Any], full_modules: list[str]) -> None:
    if inventory.get("schema_version") != 1:
        raise SystemExit("portable_stdlib_inventory.json: schema_version must be 1")

    modules = inventory.get("modules")
    if not isinstance(modules, list):
        raise SystemExit("portable_stdlib_inventory.json: modules must be an array")

    seen: set[str] = set()
    ordered_modules: list[str] = []
    for entry in modules:
        if not isinstance(entry, dict):
            raise SystemExit("portable_stdlib_inventory.json: module entries must be objects")

        module = entry.get("module")
        status = entry.get("status")
        owner = entry.get("owner")
        portable_eligible = entry.get("portable_eligible")

        if not isinstance(module, str) or not module:
            raise SystemExit("portable_stdlib_inventory.json: module must be a non-empty string")
        if module in seen:
            raise SystemExit(f"portable_stdlib_inventory.json: duplicate module entry: {module}")
        seen.add(module)
        ordered_modules.append(module)

        if status not in ALLOWED_STATUS:
            raise SystemExit(f"portable_stdlib_inventory.json: invalid status for {module}: {status!r}")
        if owner not in ALLOWED_OWNER:
            raise SystemExit(f"portable_stdlib_inventory.json: invalid owner for {module}: {owner!r}")
        if portable_eligible is not True:
            raise SystemExit(f"portable_stdlib_inventory.json: portable_eligible must be true for {module}")

    expected = sorted(full_modules)
    if sorted(seen) != expected:
        missing = sorted(set(expected) - set(seen))
        extra = sorted(set(seen) - set(expected))
        details: list[str] = []
        if missing:
            details.append(f"missing={missing[:10]}")
        if extra:
            details.append(f"extra={extra[:10]}")
        raise SystemExit(
            "portable_stdlib_inventory.json: module set must match test/upstream_std_modules_full.txt "
            + "; ".join(details)
        )

    if ordered_modules != sorted(ordered_modules):
        raise SystemExit("portable_stdlib_inventory.json: modules must be sorted by module name")


def build_summary(inventory: dict[str, Any]) -> dict[str, Any]:
    modules = inventory["modules"]
    status_counts: dict[str, int] = {status: 0 for status in sorted(ALLOWED_STATUS)}
    owner_counts: dict[str, int] = {owner: 0 for owner in sorted(ALLOWED_OWNER)}
    full_sweep_count = 0
    strict_sweep_count = 0
    for entry in modules:
        status_counts[entry["status"]] += 1
        owner_counts[entry["owner"]] += 1
        if entry.get("in_full_sweep"):
            full_sweep_count += 1
        if entry["in_strict_sweep"]:
            strict_sweep_count += 1

    return {
        "schema_version": 1,
        "total_modules": len(modules),
        "full_sweep_modules": full_sweep_count,
        "strict_sweep_modules": strict_sweep_count,
        "status_counts": status_counts,
        "owner_counts": owner_counts,
    }


def write_summary(summary: dict[str, Any]) -> None:
    CACHE_ROOT.mkdir(parents=True, exist_ok=True)
    SUMMARY_JSON.write_text(json.dumps(summary, indent=2, sort_keys=True) + "\n", encoding="utf-8")

    lines = [
        "# Portable Stdlib Inventory Summary",
        "",
        f"- total_modules: `{summary['total_modules']}`",
        f"- full_sweep_modules: `{summary['full_sweep_modules']}`",
        f"- strict_sweep_modules: `{summary['strict_sweep_modules']}`",
        "",
        "## Status counts",
    ]
    for status, count in summary["status_counts"].items():
        lines.append(f"- `{status}`: `{count}`")

    lines.append("")
    lines.append("## Owner counts")
    for owner, count in summary["owner_counts"].items():
        lines.append(f"- `{owner}`: `{count}`")

    lines.append("")
    lines.append("Artifacts:")
    lines.append(f"- `{SUMMARY_JSON.relative_to(ROOT)}`")
    lines.append(f"- `{INVENTORY_FILE.relative_to(ROOT)}`")
    SUMMARY_MD.write_text("\n".join(lines) + "\n", encoding="utf-8")


def main() -> int:
    args = parse_args()

    full_modules = load_modules(FULL_MODULES_FILE)
    strict_sweep_modules = set(load_modules(STRICT_SWEEP_MODULES_FILE))
    full_module_set = set(full_modules)
    auto_semantic_modules = {module for module in full_modules if is_semantic_diff_module(module)}
    promotions = load_promotions(PROMOTIONS_FILE, full_module_set, auto_semantic_modules)
    generated = build_inventory(full_modules, strict_sweep_modules, promotions)

    if args.update:
        INVENTORY_FILE.write_text(json.dumps(generated, indent=2, sort_keys=True) + "\n", encoding="utf-8")
        print(f"[PASS] wrote {INVENTORY_FILE.relative_to(ROOT)}")

    if not INVENTORY_FILE.exists():
        raise SystemExit(
            "portable stdlib inventory file missing. Run: python3 test/run-portable-stdlib-inventory.py --update"
        )

    existing = json.loads(INVENTORY_FILE.read_text(encoding="utf-8"))
    validate_inventory_schema(existing, full_modules)

    if existing != generated:
        raise SystemExit(
            "portable stdlib inventory drift detected. "
            "Run: python3 test/run-portable-stdlib-inventory.py --update"
        )

    summary = build_summary(existing)
    write_summary(summary)
    print("[PASS] portable stdlib inventory validated")
    print(f"[PASS] summary: {SUMMARY_JSON.relative_to(ROOT)}")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
