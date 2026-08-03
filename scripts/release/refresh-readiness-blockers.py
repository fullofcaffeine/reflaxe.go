#!/usr/bin/env python3

"""Collect release blocker evidence from the remote Beads tracker.

What: Export the exact priority/status of compatibility-manifest blockers.
Why: A stale local or hand-maintained copy could incorrectly permit a release.
How: Read scoped blocker IDs from readiness policy, initialize an isolated
Beads client against the configured remote, and bind its records to one stable
`refs/dolt/data` commit for the rest of the release workflow.
"""

from __future__ import annotations

import argparse
import json
import re
import subprocess
import sys
import tempfile
from pathlib import Path
from typing import Any


ROOT = Path(__file__).resolve().parents[2]
POLICY_PATH = ROOT / "release" / "readiness-policy.json"
SHA_RE = re.compile(r"[0-9a-f]{40}")


class BlockerSnapshotError(RuntimeError):
    """The tracker cannot produce governed release blocker evidence."""


def fail(message: str) -> None:
    raise BlockerSnapshotError(message)


def load_object(path: Path, label: str) -> dict[str, Any]:
    try:
        value = json.loads(path.read_text(encoding="utf-8"))
    except (OSError, UnicodeDecodeError, json.JSONDecodeError) as error:
        fail(f"cannot read {label}: {error}")
    if not isinstance(value, dict):
        fail(f"{label} must be a JSON object")
    return value


def run_text(command: list[str], label: str, *, cwd: Path = ROOT) -> str:
    try:
        process = subprocess.run(
            command,
            cwd=cwd,
            text=True,
            capture_output=True,
            timeout=30,
        )
    except (OSError, subprocess.TimeoutExpired) as error:
        fail(f"cannot query {label}: {error}")
    if process.returncode != 0:
        fail(f"cannot query {label}: {(process.stderr or process.stdout).strip()}")
    return process.stdout


def run_json(command: list[str], label: str, *, cwd: Path = ROOT) -> Any:
    output = run_text(command, label, cwd=cwd)
    try:
        return json.loads(output)
    except json.JSONDecodeError:
        fail(f"{label} returned invalid JSON")


def remote_tracker_commit(reference: str) -> str:
    try:
        process = subprocess.run(
            ["git", "ls-remote", "origin", reference],
            cwd=ROOT,
            text=True,
            capture_output=True,
            timeout=30,
        )
    except (OSError, subprocess.TimeoutExpired) as error:
        fail(f"cannot query remote tracker ref: {error}")
    fields = process.stdout.strip().split()
    if process.returncode != 0 or len(fields) != 2 or fields[1] != reference:
        fail(f"remote tracker ref is missing: {reference}")
    if SHA_RE.fullmatch(fields[0]) is None:
        fail("remote tracker ref did not resolve to a full commit")
    return fields[0]


def build_snapshot(policy: dict[str, Any], observed_at: str) -> dict[str, Any]:
    compatibility = policy.get("compatibility")
    if not isinstance(compatibility, dict):
        fail("readiness policy compatibility section is missing")
    scopes = compatibility.get("blockerScopes")
    if not isinstance(scopes, dict) or not scopes:
        fail("readiness policy blockerScopes must be a non-empty object")
    final_admission = policy.get("finalAdmission")
    if not isinstance(final_admission, dict):
        fail("readiness policy finalAdmission must be an object")
    admission_owner = final_admission.get("owner")
    if not isinstance(admission_owner, str) or not admission_owner:
        fail("readiness policy finalAdmission owner is invalid")
    reference = "refs/dolt/data"
    before = remote_tracker_commit(reference)
    remote = run_text(
        ["bd", "config", "get", "sync.remote"], "Beads remote configuration"
    ).strip()
    prefix = run_text(
        ["bd", "config", "get", "issue_prefix"], "Beads issue prefix"
    ).strip()
    if not remote or not prefix:
        fail("Beads remote configuration or issue prefix is empty")
    blockers: list[dict[str, Any]] = []
    with tempfile.TemporaryDirectory(prefix="haxe-go-release-blockers-") as raw:
        checkout = Path(raw) / "remote"
        checkout.mkdir()
        run_text(["git", "init", "-q"], "temporary Git initialization", cwd=checkout)
        run_text(
            [
                "bd",
                "init",
                "--prefix",
                prefix,
                "--remote",
                remote,
                "--skip-agents",
                "--skip-hooks",
                "--non-interactive",
                "--quiet",
            ],
            "remote Beads initialization",
            cwd=checkout,
        )
        for blocker_id, scope in scopes.items():
            if not isinstance(blocker_id, str) or not isinstance(scope, str):
                fail("blockerScopes entries must map strings to strings")
            raw_issue = run_json(
                ["bd", "-C", str(checkout), "show", blocker_id, "--json"],
                f"remote Bead {blocker_id}",
            )
            if (
                not isinstance(raw_issue, list)
                or len(raw_issue) != 1
                or not isinstance(raw_issue[0], dict)
            ):
                fail(f"remote Bead {blocker_id} did not return one issue")
            issue = raw_issue[0]
            priority = issue.get("priority")
            status = issue.get("status")
            if not isinstance(priority, int) or isinstance(priority, bool):
                fail(f"remote Bead {blocker_id} priority is invalid")
            if not isinstance(status, str) or not status:
                fail(f"remote Bead {blocker_id} status is invalid")
            record: dict[str, Any] = {
                "id": blocker_id,
                "priority": priority,
                "status": status,
                "scopes": [scope],
            }
            if blocker_id == admission_owner:
                metadata = issue.get("metadata")
                if not isinstance(metadata, dict):
                    fail(f"remote Bead {blocker_id} has no final admission metadata")
                admission = metadata.get("releaseAdmission")
                if not isinstance(admission, dict):
                    fail(f"remote Bead {blocker_id} has no releaseAdmission record")
                record["admission"] = admission
            blockers.append(record)
    after = remote_tracker_commit(reference)
    if after != before:
        fail("remote tracker advanced while blocker evidence was collected")
    return {
        "schemaVersion": 1,
        "kind": "haxe.go-release-blocker-evidence",
        "tracker": {
            "ref": reference,
            "commit": before,
            "observedAt": observed_at,
        },
        "records": blockers,
    }


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--observed-at", required=True)
    parser.add_argument("--output", required=True, type=Path)
    return parser.parse_args()


def main() -> int:
    arguments = parse_args()
    try:
        policy = load_object(POLICY_PATH, "readiness policy")
        expected = build_snapshot(policy, arguments.observed_at)
        rendered = json.dumps(expected, indent=2) + "\n"
        if arguments.output.exists():
            fail(f"blocker evidence output already exists: {arguments.output}")
        arguments.output.parent.mkdir(parents=True, exist_ok=True)
        arguments.output.write_text(rendered, encoding="utf-8")
    except (BlockerSnapshotError, OSError, KeyError, TypeError) as error:
        print(f"[release-blockers] ERROR: {error}", file=sys.stderr)
        return 2
    print("[release-blockers] remote Beads blocker evidence: OK")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
