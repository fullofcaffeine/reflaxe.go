#!/usr/bin/env python3

from __future__ import annotations

import argparse
import json
import subprocess
from pathlib import Path
from typing import Any

ROOT = Path(__file__).resolve().parent.parent
ALLOWLIST_FILE = ROOT / "test" / "portable_allowlist.json"
PLAN_FILE = ROOT / "test" / "portable_conformance_tier1.json"
SEMANTIC_CORE_ROOT = ROOT / "test" / "semantic_diff"
SEMANTIC_LANES_ROOT = ROOT / "test" / "semantic_diff_lanes"
CACHE_ROOT = ROOT / "test" / ".test-cache"
SUMMARY_JSON = CACHE_ROOT / "portable_conformance_tier1_summary.json"
SUMMARY_MD = CACHE_ROOT / "portable_conformance_tier1_summary.md"


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(
        description="Run deterministic Tier1 portable conformance suite mapped from portable allowlist modules."
    )
    parser.add_argument("--timeout", type=int, default=120, help="Timeout passed to semantic diff runner")
    parser.add_argument("--list", action="store_true", help="List module->case mapping and selected case set")
    parser.add_argument("--module", action="append", default=[], help="Run only selected tier1 module(s)")
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


def semantic_root_for_suite(suite: str) -> Path:
    if suite == "core":
        return SEMANTIC_CORE_ROOT
    if suite == "lanes":
        return SEMANTIC_LANES_ROOT
    raise SystemExit(f"unsupported suite in plan: {suite}")


def discover_cases(semantic_root: Path) -> set[str]:
    case_ids: set[str] = set()
    if not semantic_root.exists():
        return case_ids
    for case_dir in semantic_root.iterdir():
        if not case_dir.is_dir():
            continue
        if (case_dir / "Main.hx").exists():
            case_ids.add(case_dir.name)
    return case_ids


def parse_tier1_modules(allowlist: dict[str, Any]) -> list[str]:
    if allowlist.get("schema_version") != 1:
        raise SystemExit("portable_allowlist.json: schema_version must be 1")
    tiers = allowlist.get("tiers")
    if not isinstance(tiers, dict):
        raise SystemExit("portable_allowlist.json: tiers must be an object")
    tier1 = tiers.get("tier1")
    if not isinstance(tier1, list):
        raise SystemExit("portable_allowlist.json: tier1 must be an array")
    modules: list[str] = []
    for module in tier1:
        if not isinstance(module, str) or not module.strip():
            raise SystemExit(f"portable_allowlist.json: invalid tier1 module entry: {module!r}")
        modules.append(module)
    if modules != sorted(modules):
        raise SystemExit("portable_allowlist.json: tier1 modules must be sorted for deterministic conformance")
    return modules


def parse_plan(plan: dict[str, Any]) -> tuple[str, str, dict[str, list[str]]]:
    if plan.get("schema_version") != 1:
        raise SystemExit("portable_conformance_tier1.json: schema_version must be 1")
    suite = plan.get("suite")
    if not isinstance(suite, str) or not suite:
        raise SystemExit("portable_conformance_tier1.json: suite must be a non-empty string")
    tier = plan.get("tier")
    if not isinstance(tier, str) or not tier:
        raise SystemExit("portable_conformance_tier1.json: tier must be a non-empty string")
    raw_modules = plan.get("modules")
    if not isinstance(raw_modules, dict):
        raise SystemExit("portable_conformance_tier1.json: modules must be an object")

    mapping: dict[str, list[str]] = {}
    for module, raw_cases in raw_modules.items():
        if not isinstance(module, str) or not module.strip():
            raise SystemExit("portable_conformance_tier1.json: module keys must be non-empty strings")
        if not isinstance(raw_cases, list) or not raw_cases:
            raise SystemExit(f"portable_conformance_tier1.json: module `{module}` must map to a non-empty case array")
        cases: list[str] = []
        seen: set[str] = set()
        for case_id in raw_cases:
            if not isinstance(case_id, str) or not case_id.strip():
                raise SystemExit(f"portable_conformance_tier1.json: invalid case id for `{module}`: {case_id!r}")
            if case_id in seen:
                raise SystemExit(f"portable_conformance_tier1.json: duplicate case `{case_id}` in module `{module}`")
            seen.add(case_id)
            cases.append(case_id)
        if cases != sorted(cases):
            raise SystemExit(f"portable_conformance_tier1.json: cases for `{module}` must be sorted")
        mapping[module] = cases

    return suite, tier, mapping


def validate_plan_against_tier(
    tier_modules: list[str],
    suite: str,
    tier: str,
    mapping: dict[str, list[str]],
    discovered_cases: set[str],
) -> None:
    if tier != "tier1":
        raise SystemExit(f"portable_conformance_tier1.json: tier must be `tier1` (found {tier!r})")

    tier_set = set(tier_modules)
    mapping_set = set(mapping.keys())
    if tier_set != mapping_set:
        missing = sorted(tier_set - mapping_set)
        extra = sorted(mapping_set - tier_set)
        problems: list[str] = []
        if missing:
            problems.append(f"missing mappings: {', '.join(missing)}")
        if extra:
            problems.append(f"unexpected mappings: {', '.join(extra)}")
        raise SystemExit("portable conformance mapping mismatch with tier1 allowlist (" + "; ".join(problems) + ")")

    if suite not in {"core", "lanes"}:
        raise SystemExit(f"portable_conformance_tier1.json: unsupported suite `{suite}`")

    for module in tier_modules:
        cases = mapping[module]
        for case_id in cases:
            if case_id not in discovered_cases:
                raise SystemExit(
                    f"portable_conformance_tier1.json: case `{case_id}` for module `{module}` "
                    f"does not exist in suite `{suite}`"
                )


def selected_modules_from_args(tier_modules: list[str], raw_requested: list[str]) -> list[str]:
    if not raw_requested:
        return list(tier_modules)

    requested: list[str] = []
    seen: set[str] = set()
    for item in raw_requested:
        module = item.strip()
        if not module:
            continue
        if module in seen:
            continue
        seen.add(module)
        requested.append(module)

    unknown = sorted(module for module in requested if module not in set(tier_modules))
    if unknown:
        raise SystemExit("unknown --module value(s): " + ", ".join(unknown))

    # Preserve allowlist ordering for deterministic output.
    requested_set = set(requested)
    return [module for module in tier_modules if module in requested_set]


def selected_cases_for_modules(modules: list[str], mapping: dict[str, list[str]]) -> list[str]:
    cases: set[str] = set()
    for module in modules:
        for case_id in mapping[module]:
            cases.add(case_id)
    return sorted(cases)


def write_summary(
    suite: str,
    tier: str,
    modules: list[str],
    mapping: dict[str, list[str]],
    cases: list[str],
) -> None:
    CACHE_ROOT.mkdir(parents=True, exist_ok=True)
    payload = {
        "schema_version": 1,
        "suite": suite,
        "tier": tier,
        "module_count": len(modules),
        "case_count": len(cases),
        "modules": modules,
        "cases": cases,
        "module_case_map": {module: mapping[module] for module in modules},
    }
    SUMMARY_JSON.write_text(json.dumps(payload, indent=2, sort_keys=True) + "\n", encoding="utf-8")

    lines = [
        "# Portable Tier1 Conformance Summary",
        "",
        f"- suite: `{suite}`",
        f"- tier: `{tier}`",
        f"- module_count: `{len(modules)}`",
        f"- case_count: `{len(cases)}`",
        "",
        "## Modules",
    ]
    for module in modules:
        lines.append(f"- `{module}` -> `{', '.join(mapping[module])}`")
    lines.append("")
    lines.append("## Cases")
    for case_id in cases:
        lines.append(f"- `{case_id}`")
    lines.append("")
    lines.append("Artifacts:")
    lines.append(f"- `{SUMMARY_JSON.relative_to(ROOT)}`")
    lines.append(f"- `{SUMMARY_MD.relative_to(ROOT)}`")
    SUMMARY_MD.write_text("\n".join(lines) + "\n", encoding="utf-8")


def build_runner_command(suite: str, timeout: int, cases: list[str]) -> list[str]:
    cmd = ["python3", "test/run-semantic-diff.py", "--suite", suite, "--timeout", str(timeout)]
    for case_id in cases:
        cmd.extend(["--case", case_id])
    return cmd


def main() -> int:
    args = parse_args()

    allowlist = load_json(ALLOWLIST_FILE)
    plan = load_json(PLAN_FILE)
    tier_modules = parse_tier1_modules(allowlist)
    suite, tier, mapping = parse_plan(plan)
    discovered_cases = discover_cases(semantic_root_for_suite(suite))
    validate_plan_against_tier(tier_modules, suite, tier, mapping, discovered_cases)

    selected_modules = selected_modules_from_args(tier_modules, args.module)
    selected_cases = selected_cases_for_modules(selected_modules, mapping)

    if not selected_modules:
        print("No tier1 modules selected")
        return 0
    if not selected_cases:
        print("No conformance cases selected")
        return 0

    write_summary(suite, tier, selected_modules, mapping, selected_cases)

    if args.list:
        print(f"suite={suite} tier={tier}")
        for module in selected_modules:
            print(f"{module}: {', '.join(mapping[module])}")
        print("")
        print("cases:")
        for case_id in selected_cases:
            print(case_id)
        print("")
        print(f"summary: {SUMMARY_JSON.relative_to(ROOT)}")
        return 0

    cmd = build_runner_command(suite, args.timeout, selected_cases)
    print("$ " + " ".join(cmd))
    proc = subprocess.run(cmd, cwd=ROOT)
    if proc.returncode != 0:
        return proc.returncode

    print(f"[PASS] tier1 portable conformance validated ({len(selected_modules)} modules, {len(selected_cases)} cases)")
    print(f"[PASS] summary: {SUMMARY_JSON.relative_to(ROOT)}")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
