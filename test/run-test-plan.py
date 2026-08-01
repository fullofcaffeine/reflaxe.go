#!/usr/bin/env python3

"""Explain affected test ownership without skipping the current full gates."""

from __future__ import annotations

import argparse
import fnmatch
import json
import os
from pathlib import Path
from typing import Any

from git_changes import GitChangeDiscoveryError, collect_changed_paths


ROOT = Path(__file__).resolve().parent.parent
STRATEGY_PATH = ROOT / "test" / "testing-strategy.json"


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(
        description="Explain semantic-owner and product-surface affected test selection"
    )
    parser.add_argument("--base", default="", help="Optional Git merge-base/ref to compare with HEAD")
    parser.add_argument(
        "--changed-file",
        action="append",
        default=[],
        help="Use an explicit changed path instead of Git discovery (repeatable)",
    )
    parser.add_argument("--json", action="store_true", help="Emit machine-readable JSON")
    return parser.parse_args()


def load_strategy() -> dict[str, Any]:
    return json.loads(STRATEGY_PATH.read_text(encoding="utf-8"))


def path_matches(path: str, patterns: list[str]) -> bool:
    return any(path == pattern or fnmatch.fnmatchcase(path, pattern) for pattern in patterns)


def owner_matches(owner: dict[str, Any], path: str) -> bool:
    if not path_matches(path, owner.get("patterns", [])):
        return False
    return not path_matches(path, owner.get("excludePatterns", []))


def build_plan(
    strategy: dict[str, Any],
    changed_files: set[str],
    discovery_failure: str = "",
) -> dict[str, Any]:
    owners = strategy["testOwners"]
    selected_ids = {owner["id"] for owner in owners if owner.get("alwaysRun", False)}
    reasons: dict[str, list[str]] = {
        owner["id"]: ["always-run sentinel"]
        for owner in owners
        if owner.get("alwaysRun", False)
    }
    matched_paths: set[str] = set()
    full_expansion_reasons: list[str] = [discovery_failure] if discovery_failure else []

    for path in sorted(changed_files):
        matched = [owner for owner in owners if owner_matches(owner, path)]
        if not matched:
            full_expansion_reasons.append(f"unknown ownership: {path}")
            continue
        matched_paths.add(path)
        for owner in matched:
            selected_ids.add(owner["id"])
            reasons.setdefault(owner["id"], []).append(f"matched {path}")
            if owner.get("expandsToFull", False):
                full_expansion_reasons.append(f"cross-cutting owner {owner['id']}: {path}")

    full_expansion = bool(full_expansion_reasons)
    if full_expansion and strategy["affectedSelection"].get("unknownPathsExpandToFull", False):
        selected_ids = {owner["id"] for owner in owners}
        for owner in owners:
            reasons.setdefault(owner["id"], []).append("conservative full expansion")

    selected_owners = [owner for owner in owners if owner["id"] in selected_ids]
    selected_surfaces = sorted(
        {surface for owner in selected_owners for surface in owner.get("surfaces", [])}
    )
    commands = sorted(
        {command for owner in selected_owners for command in owner.get("commands", [])}
    )

    return {
        "schemaVersion": 1,
        "mode": strategy["affectedSelection"]["mode"],
        "changedFiles": sorted(changed_files),
        "matchedFiles": sorted(matched_paths),
        "selectedOwners": [owner["id"] for owner in selected_owners],
        "selectedSurfaces": selected_surfaces,
        "commands": commands,
        "reasons": {owner_id: reasons[owner_id] for owner_id in sorted(reasons)},
        "fullExpansion": full_expansion,
        "fullExpansionReasons": full_expansion_reasons,
        "executionPolicy": "observation only; npm run test:ci remains the complete required backstop",
    }


def default_base() -> str:
    configured = os.environ.get("TEST_PLAN_BASE_REF", "").strip()
    if configured:
        return configured
    github_base = os.environ.get("GITHUB_BASE_REF", "").strip()
    if github_base:
        return f"origin/{github_base}"
    if os.environ.get("GITHUB_ACTIONS", "").lower() == "true":
        return "HEAD^"
    return ""


def render_human(plan: dict[str, Any]) -> str:
    lines = [
        "Affected test plan (observation only)",
        f"Changed files: {len(plan['changedFiles'])}",
        f"Selected owners: {', '.join(plan['selectedOwners']) or '<none>'}",
        f"Selected surfaces: {', '.join(plan['selectedSurfaces']) or '<none>'}",
        f"Full expansion: {'yes' if plan['fullExpansion'] else 'no'}",
    ]
    if plan["fullExpansionReasons"]:
        lines.append("Expansion reasons:")
        lines.extend(f"  - {reason}" for reason in plan["fullExpansionReasons"])
    lines.append("Commands represented:")
    lines.extend(f"  - {command}" for command in plan["commands"])
    lines.append("No command is skipped from the required CI graph by this plan.")
    return "\n".join(lines)


def main() -> int:
    args = parse_args()
    changed_files = set(args.changed_file)
    discovery_failure = ""
    if not changed_files:
        try:
            changed_files = collect_changed_paths(ROOT, base=args.base or default_base())
        except GitChangeDiscoveryError as error:
            discovery_failure = str(error)
    plan = build_plan(load_strategy(), changed_files, discovery_failure)
    if args.json:
        print(json.dumps(plan, indent=2, sort_keys=True))
    else:
        print(render_human(plan))
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
