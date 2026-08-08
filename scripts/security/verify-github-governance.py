#!/usr/bin/env python3

"""Verify source-owned and live GitHub governance against one policy.

What: Check the repository's declared GitHub security and release controls.
Why: Workflow YAML cannot prove that mutable host settings are still active.
How: Source mode validates durable repository evidence; live mode additionally
queries GitHub and compares the effective settings and rulesets fail-closed.
"""

from __future__ import annotations

import argparse
import json
import re
import subprocess
import sys
from pathlib import Path
from typing import Any


ROOT = Path(__file__).resolve().parent.parent.parent
POLICY_PATH = ROOT / "github-governance-policy.json"
WORKFLOW_ROOT = ROOT / ".github" / "workflows"
DEPENDABOT_PATH = ROOT / ".github" / "dependabot.yml"
DOCUMENTATION_PATH = ROOT / "docs" / "github-governance-policy.md"
FULL_SHA_RE = re.compile(r"^[0-9a-f]{40}$")


class GovernanceError(RuntimeError):
    pass


def load_policy() -> dict[str, Any]:
    try:
        policy = json.loads(POLICY_PATH.read_text(encoding="utf-8"))
    except (OSError, json.JSONDecodeError) as exc:
        raise GovernanceError(f"cannot load {POLICY_PATH.name}: {exc}") from exc
    if policy.get("schema_version") != 1:
        raise GovernanceError("unsupported github-governance-policy schema")
    if not isinstance(policy.get("repository"), str) or "/" not in policy["repository"]:
        raise GovernanceError("policy repository must be owner/name")
    if not isinstance(policy.get("rulesets"), list) or not policy["rulesets"]:
        raise GovernanceError("policy must declare at least one ruleset")
    return policy


def verify_source(policy: dict[str, Any]) -> None:
    actions = policy.get("actions", {})
    expected_actions = {
        "enabled": True,
        "allowed_actions": "all",
        "sha_pinning_required": True,
        "default_workflow_permissions": "read",
        "can_approve_pull_request_reviews": False,
    }
    if actions != expected_actions:
        raise GovernanceError("actions policy is not the reviewed least-privilege shape")

    security = policy.get("security_and_analysis", {})
    expected_security_keys = {
        "secret_scanning",
        "secret_scanning_push_protection",
        "dependabot_security_updates",
    }
    if set(security) != expected_security_keys or set(security.values()) != {"enabled"}:
        raise GovernanceError("all declared security_and_analysis controls must be enabled")
    if policy.get("vulnerability_alerts") is not True:
        raise GovernanceError("Dependabot vulnerability alerts must be required")
    if policy.get("immutable_releases") is not True:
        raise GovernanceError("immutable releases must be required")
    limitations = policy.get("plan_limitations", [])
    expected_limitations = {
        "secret_scanning_non_provider_patterns",
        "secret_scanning_validity_checks",
    }
    if {item.get("control") for item in limitations} != expected_limitations:
        raise GovernanceError("GitHub Secret Protection plan limitations drifted")
    for limitation in limitations:
        if limitation.get("observed_status") != "disabled":
            raise GovernanceError("plan limitation must record its observed disabled state")
        if "GitHub Secret Protection" not in limitation.get("reason", ""):
            raise GovernanceError("plan limitation must explain the unavailable product")

    names = [ruleset.get("name") for ruleset in policy["rulesets"]]
    if len(names) != len(set(names)):
        raise GovernanceError("ruleset names must be unique")
    required_rulesets = {"Protect master", "Protect release tags"}
    if set(names) != required_rulesets:
        raise GovernanceError(f"rulesets must be exactly {sorted(required_rulesets)}")

    workflow_count = 0
    for path in sorted(WORKFLOW_ROOT.glob("*.yml")):
        workflow_count += 1
        text = path.read_text(encoding="utf-8")
        if "permissions:" not in text:
            raise GovernanceError(
                f"{path.relative_to(ROOT)} has no explicit GITHUB_TOKEN permissions"
            )
        for line_number, line in enumerate(text.splitlines(), 1):
            stripped = line.strip()
            if not stripped.startswith("uses:"):
                continue
            reference = stripped.split("uses:", 1)[1].split("#", 1)[0].strip()
            if reference.startswith("./"):
                continue
            if "@" not in reference or not FULL_SHA_RE.fullmatch(reference.rsplit("@", 1)[1]):
                raise GovernanceError(
                    f"{path.relative_to(ROOT)}:{line_number}: action is not pinned "
                    "to a full commit SHA"
                )
    if workflow_count == 0:
        raise GovernanceError("no GitHub Actions workflows found")

    dependabot = DEPENDABOT_PATH.read_text(encoding="utf-8")
    for ecosystem in ('package-ecosystem: "github-actions"', 'package-ecosystem: "npm"'):
        if ecosystem not in dependabot:
            raise GovernanceError(f"dependabot.yml is missing {ecosystem}")

    documentation = DOCUMENTATION_PATH.read_text(encoding="utf-8")
    if "github-governance-policy.json" not in documentation:
        raise GovernanceError("governance documentation does not name the policy authority")

    print(
        "[github-governance] source: OK "
        f"({workflow_count} workflows; {len(policy['rulesets'])} rulesets declared)"
    )


def run_gh(
    policy: dict[str, Any],
    endpoint: str,
    *,
    expect_json: bool = True,
) -> Any:
    command = [
        "gh",
        "api",
        "-H",
        "Accept: application/vnd.github+json",
        "-H",
        f"X-GitHub-Api-Version: {policy['api_version']}",
        endpoint,
    ]
    proc = subprocess.run(command, cwd=ROOT, text=True, capture_output=True)
    if proc.returncode != 0:
        detail = (proc.stderr or proc.stdout).strip()
        if re.search(
            r"(?:\b403\b|\b404\b|resource not accessible|requires? admin)",
            detail,
            flags=re.IGNORECASE,
        ):
            raise GovernanceError(
                "governance reader lacks required Administration: read-only access "
                f"to {endpoint}; verify the app permission and this-repository-only "
                f"installation. GitHub said: {detail}"
            )
        raise GovernanceError(f"GitHub API {endpoint} failed: {detail}")
    if not expect_json:
        return None
    try:
        return json.loads(proc.stdout)
    except json.JSONDecodeError as exc:
        raise GovernanceError(f"GitHub API {endpoint} returned invalid JSON") from exc


def require_subset(expected: Any, actual: Any, path: str) -> None:
    if isinstance(expected, dict):
        if not isinstance(actual, dict):
            raise GovernanceError(f"{path} is not an object")
        for key, value in expected.items():
            if key not in actual:
                raise GovernanceError(f"{path}.{key} is missing")
            require_subset(value, actual[key], f"{path}.{key}")
        return
    if isinstance(expected, list):
        if not isinstance(actual, list):
            raise GovernanceError(f"{path} is not a list")
        if all(isinstance(item, dict) and "context" in item for item in expected):
            expected_by_context = {item["context"]: item for item in expected}
            actual_by_context = {item.get("context"): item for item in actual}
            if set(expected_by_context) != set(actual_by_context):
                raise GovernanceError(f"{path} status-check contexts drifted")
            for context, value in expected_by_context.items():
                require_subset(value, actual_by_context[context], f"{path}[{context}]")
            return
        if expected != actual:
            raise GovernanceError(f"{path} drifted: expected {expected!r}, got {actual!r}")
        return
    if expected != actual:
        raise GovernanceError(f"{path} drifted: expected {expected!r}, got {actual!r}")


def verify_live(policy: dict[str, Any]) -> None:
    repository = policy["repository"]
    repo_state = run_gh(policy, f"repos/{repository}")
    if repo_state.get("default_branch") != policy["default_branch"]:
        raise GovernanceError("live default branch differs from policy")
    live_security = repo_state.get("security_and_analysis", {})
    for name, expected_status in policy["security_and_analysis"].items():
        actual_status = live_security.get(name, {}).get("status")
        if actual_status != expected_status:
            raise GovernanceError(
                f"security_and_analysis.{name}: expected {expected_status}, "
                f"got {actual_status}"
            )
    for limitation in policy["plan_limitations"]:
        name = limitation["control"]
        actual_status = live_security.get(name, {}).get("status")
        if actual_status != limitation["observed_status"]:
            raise GovernanceError(
                f"plan limitation {name}: expected {limitation['observed_status']}, "
                f"got {actual_status}; update the policy if the feature became available"
            )

    actions = run_gh(policy, f"repos/{repository}/actions/permissions")
    for name in ("enabled", "allowed_actions", "sha_pinning_required"):
        require_subset(policy["actions"][name], actions.get(name), f"actions.{name}")
    workflow_permissions = run_gh(
        policy, f"repos/{repository}/actions/permissions/workflow"
    )
    for name in ("default_workflow_permissions", "can_approve_pull_request_reviews"):
        require_subset(
            policy["actions"][name],
            workflow_permissions.get(name),
            f"actions.{name}",
        )

    immutable = run_gh(policy, f"repos/{repository}/immutable-releases")
    if immutable.get("enabled") is not policy["immutable_releases"]:
        raise GovernanceError("immutable release setting differs from policy")
    run_gh(policy, f"repos/{repository}/vulnerability-alerts", expect_json=False)

    bypass = policy["bypass_actor"]
    actor = run_gh(policy, f"users/{bypass['login']}")
    expected_bypass = {
        "actor_id": actor["id"],
        "actor_type": bypass["type"],
        "bypass_mode": bypass["mode"],
    }
    summaries = run_gh(policy, f"repos/{repository}/rulesets")
    live_by_name = {item["name"]: item for item in summaries}
    if set(live_by_name) != {item["name"] for item in policy["rulesets"]}:
        raise GovernanceError("live ruleset names differ from policy")

    for expected in policy["rulesets"]:
        ruleset_id = live_by_name[expected["name"]]["id"]
        actual = run_gh(policy, f"repos/{repository}/rulesets/{ruleset_id}")
        require_subset(expected["target"], actual.get("target"), f"{expected['name']}.target")
        require_subset(
            expected["enforcement"],
            actual.get("enforcement"),
            f"{expected['name']}.enforcement",
        )
        require_subset(
            expected["include"],
            actual.get("conditions", {}).get("ref_name", {}).get("include"),
            f"{expected['name']}.include",
        )
        if expected_bypass not in actual.get("bypass_actors", []):
            raise GovernanceError(f"{expected['name']} lacks reviewed owner bypass")
        actual_rules = {rule["type"]: rule for rule in actual.get("rules", [])}
        if set(actual_rules) != {rule["type"] for rule in expected["rules"]}:
            raise GovernanceError(f"{expected['name']} rule types differ from policy")
        for expected_rule in expected["rules"]:
            require_subset(
                expected_rule,
                actual_rules[expected_rule["type"]],
                f"{expected['name']}.{expected_rule['type']}",
            )

    print(
        "[github-governance] live: OK "
        f"({repository}; {len(policy['rulesets'])} active rulesets)"
    )


def main() -> int:
    parser = argparse.ArgumentParser()
    parser.add_argument("--mode", choices=("source", "live"), default="source")
    args = parser.parse_args()
    try:
        policy = load_policy()
        verify_source(policy)
        if args.mode == "live":
            verify_live(policy)
    except (GovernanceError, OSError) as exc:
        print(f"[github-governance] error: {exc}", file=sys.stderr)
        return 1
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
