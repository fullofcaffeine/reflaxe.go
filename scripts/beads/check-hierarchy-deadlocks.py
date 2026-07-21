#!/usr/bin/env python3
"""Reject active blocking edges that point through the issue hierarchy.

A parent-child edge is organizational. A separate blocking edge between an
ancestor and one of its descendants creates a feedback loop once a blocked
parent's state is inherited by its descendants. Sibling-to-sibling ordering is
valid and is intentionally not rejected.
"""

from __future__ import annotations

import argparse
import json
import subprocess
import sys
from pathlib import Path
from typing import Any


BLOCKING_TYPES = {"blocks", "conditional-blocks", "waits-for"}
INACTIVE_STATUSES = {"closed", "pinned", "tombstone"}


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(
        description="Find active blocking dependencies between ancestors and descendants."
    )
    parser.add_argument(
        "--input",
        type=Path,
        help="Read a bd JSONL export instead of exporting the current database.",
    )
    return parser.parse_args()


def load_records(input_path: Path | None) -> list[dict[str, Any]]:
    if input_path is None:
        completed = subprocess.run(
            ["bd", "export"],
            check=False,
            capture_output=True,
            text=True,
        )
        if completed.returncode != 0:
            detail = completed.stderr.strip() or completed.stdout.strip()
            raise RuntimeError(f"bd export failed: {detail}")
        raw = completed.stdout
    else:
        raw = input_path.read_text(encoding="utf-8")

    records: list[dict[str, Any]] = []
    for line_number, line in enumerate(raw.splitlines(), start=1):
        if not line.strip():
            continue
        try:
            record = json.loads(line)
        except json.JSONDecodeError as error:
            raise ValueError(f"invalid JSON on line {line_number}: {error}") from error
        if not isinstance(record, dict):
            raise ValueError(f"line {line_number} is not a JSON object")
        if record.get("_type") != "memory":
            records.append(record)
    return records


def dependency_rows(records: list[dict[str, Any]]) -> list[tuple[str, str, str]]:
    rows: set[tuple[str, str, str]] = set()
    for record in records:
        for dependency in record.get("dependencies", []):
            if not isinstance(dependency, dict):
                raise ValueError(f"issue {record.get('id', '<unknown>')} has a malformed dependency")
            issue_id = dependency.get("issue_id")
            depends_on_id = dependency.get("depends_on_id")
            dependency_type = dependency.get("type")
            if not all(isinstance(value, str) and value for value in (issue_id, depends_on_id, dependency_type)):
                raise ValueError(f"issue {record.get('id', '<unknown>')} has an incomplete dependency")
            rows.add((issue_id, depends_on_id, dependency_type))
    return sorted(rows)


def build_parent_map(
    issue_ids: set[str], rows: list[tuple[str, str, str]]
) -> dict[str, str]:
    parents: dict[str, str] = {}
    for child_id, parent_id, dependency_type in rows:
        if dependency_type != "parent-child":
            continue
        if child_id not in issue_ids or parent_id not in issue_ids:
            raise ValueError(f"parent-child edge references a missing issue: {child_id} -> {parent_id}")
        previous = parents.get(child_id)
        if previous is not None and previous != parent_id:
            raise ValueError(f"issue has multiple structural parents: {child_id} -> {previous}, {parent_id}")
        parents[child_id] = parent_id

    for issue_id in issue_ids:
        seen: set[str] = set()
        current = issue_id
        while current in parents:
            if current in seen:
                raise ValueError(f"parent-child cycle includes {current}")
            seen.add(current)
            current = parents[current]
    return parents


def is_ancestor(candidate: str, issue_id: str, parents: dict[str, str]) -> bool:
    current = issue_id
    while current in parents:
        current = parents[current]
        if current == candidate:
            return True
    return False


def main() -> int:
    args = parse_args()
    try:
        records = load_records(args.input)
        statuses = {
            record["id"]: str(record.get("status", "open"))
            for record in records
            if isinstance(record.get("id"), str)
        }
        rows = dependency_rows(records)
        parents = build_parent_map(set(statuses), rows)
    except (OSError, RuntimeError, ValueError) as error:
        print(f"[beads-hierarchy] error: {error}", file=sys.stderr)
        return 2

    findings: list[tuple[bool, str, str, str, str]] = []
    for issue_id, depends_on_id, dependency_type in rows:
        if dependency_type not in BLOCKING_TYPES:
            continue
        if is_ancestor(issue_id, depends_on_id, parents):
            relationship = "parent-depends-on-descendant"
        elif is_ancestor(depends_on_id, issue_id, parents):
            relationship = "child-depends-on-ancestor"
        else:
            continue

        active = (
            statuses.get(issue_id, "closed") not in INACTIVE_STATUSES
            and statuses.get(depends_on_id, "closed") not in INACTIVE_STATUSES
        )
        findings.append((active, relationship, issue_id, depends_on_id, dependency_type))

    active_findings = [finding for finding in findings if finding[0]]
    inactive_findings = [finding for finding in findings if not finding[0]]
    print(f"[beads-hierarchy] active={len(active_findings)} inactive={len(inactive_findings)}")

    for _active, relationship, issue_id, depends_on_id, dependency_type in inactive_findings:
        print(
            f"[beads-hierarchy] inactive {relationship}: "
            f"{issue_id} -> {depends_on_id} ({dependency_type})"
        )
    for _active, relationship, issue_id, depends_on_id, dependency_type in active_findings:
        print(
            f"[beads-hierarchy] {relationship}: "
            f"{issue_id} -> {depends_on_id} ({dependency_type})",
            file=sys.stderr,
        )

    return 1 if active_findings else 0


if __name__ == "__main__":
    raise SystemExit(main())
