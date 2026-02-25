#!/usr/bin/env python3

from __future__ import annotations

import json
from pathlib import Path

ROOT = Path(__file__).resolve().parent.parent
ALLOWLIST_FILE = ROOT / "test" / "portable_allowlist.json"
INVENTORY_FILE = ROOT / "test" / "portable_stdlib_inventory.json"
CACHE_ROOT = ROOT / "test" / ".test-cache"
SUMMARY_JSON = CACHE_ROOT / "portable_allowlist_summary.json"
SUMMARY_MD = CACHE_ROOT / "portable_allowlist_summary.md"


def load_json(path: Path) -> object:
    try:
        return json.loads(path.read_text(encoding="utf-8"))
    except json.JSONDecodeError as exc:
        raise SystemExit(f"{path}: invalid JSON ({exc})") from exc


def to_excluded_prefixes(raw: object) -> list[str]:
    prefixes: list[str] = []
    if not isinstance(raw, list):
        return prefixes
    for item in raw:
        if not isinstance(item, str):
            continue
        trimmed = item.strip()
        if not trimmed:
            continue
        if trimmed.endswith(".*"):
            prefixes.append(trimmed[:-1])
        else:
            prefixes.append(trimmed)
    return prefixes


def main() -> int:
    if not ALLOWLIST_FILE.exists():
        raise SystemExit(f"missing allowlist file: {ALLOWLIST_FILE}")
    if not INVENTORY_FILE.exists():
        raise SystemExit(f"missing inventory file: {INVENTORY_FILE}")

    allowlist = load_json(ALLOWLIST_FILE)
    inventory = load_json(INVENTORY_FILE)

    if not isinstance(allowlist, dict):
        raise SystemExit("portable_allowlist.json: root must be an object")
    if allowlist.get("schema_version") != 1:
        raise SystemExit("portable_allowlist.json: schema_version must be 1")
    tiers = allowlist.get("tiers")
    if not isinstance(tiers, dict):
        raise SystemExit("portable_allowlist.json: tiers must be an object")

    if not isinstance(inventory, dict):
        raise SystemExit("portable_stdlib_inventory.json: root must be an object")
    baseline = inventory.get("baseline")
    if not isinstance(baseline, dict):
        raise SystemExit("portable_stdlib_inventory.json: baseline must be an object")
    excluded_prefixes = to_excluded_prefixes(baseline.get("excluded_prefix_examples"))

    modules = inventory.get("modules")
    if not isinstance(modules, list):
        raise SystemExit("portable_stdlib_inventory.json: modules must be an array")
    inventory_by_module: dict[str, dict[str, object]] = {}
    for entry in modules:
        if not isinstance(entry, dict):
            continue
        module = entry.get("module")
        if isinstance(module, str) and module:
            inventory_by_module[module] = entry

    all_modules: list[str] = []
    for tier_name in sorted(tiers.keys()):
        tier_values = tiers[tier_name]
        if not isinstance(tier_values, list):
            raise SystemExit(f"portable_allowlist.json: tier `{tier_name}` must be an array")
        if tier_values != sorted(tier_values):
            raise SystemExit(f"portable_allowlist.json: tier `{tier_name}` modules must be sorted")
        for module in tier_values:
            if not isinstance(module, str) or not module.strip():
                raise SystemExit(f"portable_allowlist.json: tier `{tier_name}` has invalid module entry: {module!r}")
            all_modules.append(module)

            if any(module.startswith(prefix) for prefix in excluded_prefixes):
                raise SystemExit(
                    f"portable_allowlist.json: module `{module}` is target-native by excluded-prefix rule ({', '.join(excluded_prefixes)})"
                )

            inv_entry = inventory_by_module.get(module)
            if inv_entry is None:
                raise SystemExit(f"portable_allowlist.json: module `{module}` is not present in portable inventory")
            if inv_entry.get("portable_eligible") is not True:
                raise SystemExit(f"portable_allowlist.json: module `{module}` is not marked portable_eligible in inventory")

    if len(all_modules) != len(set(all_modules)):
        seen: set[str] = set()
        duplicates: list[str] = []
        for module in all_modules:
            if module in seen and module not in duplicates:
                duplicates.append(module)
            seen.add(module)
        raise SystemExit(f"portable_allowlist.json: duplicate modules across tiers: {', '.join(sorted(duplicates))}")

    CACHE_ROOT.mkdir(parents=True, exist_ok=True)
    summary = {
        "schema_version": 1,
        "tiers": {tier: len(values) for tier, values in sorted(tiers.items()) if isinstance(values, list)},
        "total_modules": len(all_modules),
        "excluded_prefixes": excluded_prefixes,
    }
    SUMMARY_JSON.write_text(json.dumps(summary, indent=2, sort_keys=True) + "\n", encoding="utf-8")
    md_lines = [
        "# Portable Allowlist Summary",
        "",
        f"- total modules: **{len(all_modules)}**",
    ]
    for tier, count in sorted(summary["tiers"].items()):
        md_lines.append(f"- {tier}: **{count}**")
    SUMMARY_MD.write_text("\n".join(md_lines) + "\n", encoding="utf-8")

    print(f"[PASS] portable allowlist validated ({len(all_modules)} modules across {len(tiers)} tiers)")
    print(f"[PASS] summary: {SUMMARY_JSON.relative_to(ROOT)}")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
