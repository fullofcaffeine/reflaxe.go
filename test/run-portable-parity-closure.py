#!/usr/bin/env python3

from __future__ import annotations

import argparse
import json
from pathlib import Path
from typing import Any

ROOT = Path(__file__).resolve().parent.parent
INVENTORY_FILE = ROOT / "test" / "portable_stdlib_inventory.json"
ALLOWLIST_FILE = ROOT / "test" / "portable_allowlist.json"
CONFORMANCE_FILE = ROOT / "test" / "portable_conformance_tier1.json"
CACHE_ROOT = ROOT / "test" / ".test-cache"
SUMMARY_JSON = CACHE_ROOT / "portable_parity_closure_summary.json"
SUMMARY_MD = CACHE_ROOT / "portable_parity_closure_summary.md"

VALID_STATUS = {"unsupported", "compile-only", "snapshot", "semantic-diff"}


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(
        description="Generate portable parity closure summary and promotion queue from full-module inventory."
    )
    parser.add_argument("--list-blockers", action="store_true", help="Print blocker modules to stdout")
    return parser.parse_args()


def load_json(path: Path) -> dict[str, Any]:
    try:
        value = json.loads(path.read_text(encoding="utf-8"))
    except FileNotFoundError as exc:
        raise SystemExit(f"missing file: {path}") from exc
    except json.JSONDecodeError as exc:
        raise SystemExit(f"{path}: invalid JSON ({exc})") from exc
    if not isinstance(value, dict):
        raise SystemExit(f"{path}: root must be an object")
    return value


def extract_tier1_modules(allowlist: dict[str, Any]) -> list[str]:
    if allowlist.get("schema_version") != 1:
        raise SystemExit("portable_allowlist.json: schema_version must be 1")
    tiers = allowlist.get("tiers")
    if not isinstance(tiers, dict):
        raise SystemExit("portable_allowlist.json: tiers must be an object")
    tier1 = tiers.get("tier1")
    if not isinstance(tier1, list):
        raise SystemExit("portable_allowlist.json: tier1 must be an array")
    modules: list[str] = []
    for item in tier1:
        if not isinstance(item, str) or not item.strip():
            raise SystemExit(f"portable_allowlist.json: invalid tier1 module entry: {item!r}")
        modules.append(item)
    return modules


def extract_conformance_modules(conformance: dict[str, Any]) -> dict[str, list[str]]:
    if conformance.get("schema_version") != 1:
        raise SystemExit("portable_conformance_tier1.json: schema_version must be 1")
    modules = conformance.get("modules")
    if not isinstance(modules, dict):
        raise SystemExit("portable_conformance_tier1.json: modules must be an object")
    out: dict[str, list[str]] = {}
    for module, raw_cases in modules.items():
        if not isinstance(module, str) or not module.strip():
            raise SystemExit("portable_conformance_tier1.json: module keys must be non-empty strings")
        if not isinstance(raw_cases, list) or not raw_cases:
            raise SystemExit(f"portable_conformance_tier1.json: module `{module}` must map to a non-empty case array")
        cases: list[str] = []
        for case_id in raw_cases:
            if not isinstance(case_id, str) or not case_id.strip():
                raise SystemExit(f"portable_conformance_tier1.json: invalid case id for `{module}`: {case_id!r}")
            cases.append(case_id)
        out[module] = sorted(set(cases))
    return out


def next_promotion_step(status: str) -> str:
    if status == "unsupported":
        return "keep explicit exclusion policy, or promote only if the module becomes portable-eligible on Go"
    if status == "compile-only":
        return "promote to snapshot (add deterministic generated/runtime smoke contract)"
    if status == "snapshot":
        return "promote to semantic-diff (add interp-vs-go runtime parity fixture)"
    if status == "semantic-diff":
        return "none"
    return "unknown"


def main() -> int:
    args = parse_args()

    inventory = load_json(INVENTORY_FILE)
    allowlist = load_json(ALLOWLIST_FILE)
    conformance = load_json(CONFORMANCE_FILE)

    modules = inventory.get("modules")
    if not isinstance(modules, list):
        raise SystemExit("portable_stdlib_inventory.json: modules must be an array")

    inventory_rows: list[dict[str, Any]] = []
    for entry in modules:
        if not isinstance(entry, dict):
            raise SystemExit("portable_stdlib_inventory.json: module entries must be objects")
        module = entry.get("module")
        status = entry.get("status")
        owner = entry.get("owner")
        blocker_issue = entry.get("blocker_issue")
        blocker_family = entry.get("blocker_family")
        closure_target = entry.get("closure_target")
        if not isinstance(module, str) or not module:
            raise SystemExit("portable_stdlib_inventory.json: module must be non-empty string")
        if status not in VALID_STATUS:
            raise SystemExit(f"portable_stdlib_inventory.json: invalid status for {module}: {status!r}")
        if not isinstance(owner, str) or not owner:
            raise SystemExit(f"portable_stdlib_inventory.json: invalid owner for {module}: {owner!r}")
        if status == "compile-only":
            if not isinstance(blocker_issue, str) or not blocker_issue.strip():
                raise SystemExit(f"portable_stdlib_inventory.json: compile-only module missing blocker_issue: {module}")
            if not isinstance(blocker_family, str) or not blocker_family.strip():
                raise SystemExit(f"portable_stdlib_inventory.json: compile-only module missing blocker_family: {module}")
            if not isinstance(closure_target, str) or not closure_target.strip():
                raise SystemExit(f"portable_stdlib_inventory.json: compile-only module missing closure_target: {module}")
        inventory_rows.append(entry)

    inventory_rows.sort(key=lambda item: str(item["module"]))

    status_counts: dict[str, int] = {key: 0 for key in sorted(VALID_STATUS)}
    owner_counts: dict[str, int] = {}
    blockers: list[dict[str, str]] = []
    for row in inventory_rows:
        status = str(row["status"])
        owner = str(row["owner"])
        status_counts[status] = status_counts.get(status, 0) + 1
        owner_counts[owner] = owner_counts.get(owner, 0) + 1
        if status != "semantic-diff":
            blocker = {
                "module": str(row["module"]),
                "status": status,
                "owner": owner,
                "next_step": next_promotion_step(status),
            }
            if status == "compile-only":
                blocker["blocker_issue"] = str(row["blocker_issue"])
                blocker["blocker_family"] = str(row["blocker_family"])
                blocker["closure_target"] = str(row["closure_target"])
            blockers.append(blocker)

    tier1_modules = sorted(extract_tier1_modules(allowlist))
    conformance_modules = extract_conformance_modules(conformance)
    conformance_keys = sorted(conformance_modules.keys())

    missing_tier1_conformance = sorted(set(tier1_modules) - set(conformance_keys))
    extra_tier1_conformance = sorted(set(conformance_keys) - set(tier1_modules))
    if missing_tier1_conformance or extra_tier1_conformance:
        problems: list[str] = []
        if missing_tier1_conformance:
            problems.append("missing mappings: " + ", ".join(missing_tier1_conformance))
        if extra_tier1_conformance:
            problems.append("unexpected mappings: " + ", ".join(extra_tier1_conformance))
        raise SystemExit("tier1 conformance map mismatch (" + "; ".join(problems) + ")")

    conformance_case_set: set[str] = set()
    for module in tier1_modules:
        for case_id in conformance_modules[module]:
            conformance_case_set.add(case_id)

    summary = {
        "schema_version": 1,
        "inventory_module_count": len(inventory_rows),
        "status_counts": status_counts,
        "owner_counts": dict(sorted(owner_counts.items())),
        "remaining_blocker_count": len(blockers),
        "remaining_blockers": blockers,
        "tier1_module_count": len(tier1_modules),
        "tier1_conformance_case_count": len(conformance_case_set),
        "tier1_modules": tier1_modules,
        "tier1_conformance_cases": sorted(conformance_case_set),
    }

    CACHE_ROOT.mkdir(parents=True, exist_ok=True)
    SUMMARY_JSON.write_text(json.dumps(summary, indent=2, sort_keys=True) + "\n", encoding="utf-8")

    md_lines = [
        "# Portable Parity Closure Summary",
        "",
        f"- inventory_module_count: `{summary['inventory_module_count']}`",
        f"- remaining_blocker_count: `{summary['remaining_blocker_count']}`",
        f"- tier1_module_count: `{summary['tier1_module_count']}`",
        f"- tier1_conformance_case_count: `{summary['tier1_conformance_case_count']}`",
        "",
        "## Status counts",
    ]
    for status in sorted(status_counts.keys()):
        md_lines.append(f"- `{status}`: `{status_counts[status]}`")

    md_lines.extend(["", "## Owner counts"])
    for owner in sorted(owner_counts.keys()):
        md_lines.append(f"- `{owner}`: `{owner_counts[owner]}`")

    md_lines.extend(["", "## Remaining blockers (non semantic-diff)"])
    if blockers:
        for blocker in blockers:
            if blocker["status"] == "compile-only":
                md_lines.append(
                    f"- `{blocker['module']}` ({blocker['status']}, owner `{blocker['owner']}`, "
                    f"issue `{blocker['blocker_issue']}`, target `{blocker['closure_target']}`) -> {blocker['next_step']}"
                )
            else:
                md_lines.append(
                    f"- `{blocker['module']}` ({blocker['status']}, owner `{blocker['owner']}`) -> {blocker['next_step']}"
                )
    else:
        md_lines.append("- none")

    md_lines.extend(["", "Artifacts:"])
    md_lines.append(f"- `{SUMMARY_JSON.relative_to(ROOT)}`")
    md_lines.append(f"- `{SUMMARY_MD.relative_to(ROOT)}`")
    SUMMARY_MD.write_text("\n".join(md_lines) + "\n", encoding="utf-8")

    if args.list_blockers:
        for blocker in blockers:
            if blocker["status"] == "compile-only":
                print(
                    f"{blocker['module']} [{blocker['status']}] "
                    f"[{blocker['blocker_issue']}, target {blocker['closure_target']}] -> {blocker['next_step']}"
                )
            else:
                print(f"{blocker['module']} [{blocker['status']}] -> {blocker['next_step']}")

    print(
        f"[PASS] portable parity closure summary generated "
        f"({len(inventory_rows)} modules, {len(blockers)} blockers)"
    )
    print(f"[PASS] summary: {SUMMARY_JSON.relative_to(ROOT)}")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
