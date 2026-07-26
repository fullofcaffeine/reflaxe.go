#!/usr/bin/env python3

from __future__ import annotations

import json
import subprocess
import unittest
from pathlib import Path


ROOT = Path(__file__).resolve().parent.parent
POLICY_PATH = ROOT / "github-governance-policy.json"
DOC_PATH = ROOT / "docs" / "github-governance-policy.md"
EVIDENCE_PATH = ROOT / "docs" / "reviews" / "github-governance-vfp-4.10.md"
VERIFIER_PATH = ROOT / "scripts" / "security" / "verify-github-governance.py"
RELEASE_CONTRACT_RUNNER = ROOT / "test" / "run-release-contracts.py"


class GithubGovernancePolicyContract(unittest.TestCase):
    def policy(self) -> dict[str, object]:
        return json.loads(POLICY_PATH.read_text(encoding="utf-8"))

    def test_policy_declares_fail_closed_host_controls(self) -> None:
        policy = self.policy()

        self.assertEqual(1, policy["schema_version"])
        self.assertEqual("fullofcaffeine/reflaxe.go", policy["repository"])
        self.assertEqual("master", policy["default_branch"])
        self.assertEqual(
            {
                "enabled": True,
                "allowed_actions": "all",
                "sha_pinning_required": True,
                "default_workflow_permissions": "read",
                "can_approve_pull_request_reviews": False,
            },
            policy["actions"],
        )
        self.assertEqual(
            {
                "secret_scanning": "enabled",
                "secret_scanning_push_protection": "enabled",
                "dependabot_security_updates": "enabled",
            },
            policy["security_and_analysis"],
        )
        self.assertEqual(
            {
                "secret_scanning_non_provider_patterns",
                "secret_scanning_validity_checks",
            },
            {item["control"] for item in policy["plan_limitations"]},
        )
        for limitation in policy["plan_limitations"]:
            self.assertEqual("disabled", limitation["observed_status"])
            self.assertIn("GitHub Secret Protection", limitation["reason"])
        self.assertTrue(policy["vulnerability_alerts"])
        self.assertTrue(policy["immutable_releases"])

        rulesets = {item["name"]: item for item in policy["rulesets"]}
        branch = rulesets["Protect master"]
        self.assertEqual("branch", branch["target"])
        self.assertEqual(["~DEFAULT_BRANCH"], branch["include"])
        self.assertEqual(
            {
                "deletion",
                "non_fast_forward",
                "pull_request",
                "required_linear_history",
                "required_status_checks",
            },
            {rule["type"] for rule in branch["rules"]},
        )
        status_rule = next(
            rule for rule in branch["rules"] if rule["type"] == "required_status_checks"
        )
        contexts = {
            check["context"]
            for check in status_rule["parameters"]["required_status_checks"]
        }
        self.assertEqual(
            {
                "Quality Gate",
                "test:ci (ubuntu-latest, go 1.25.12)",
                "test:ci (ubuntu-latest, go 1.26.5)",
                "test:ci (macos-latest, go 1.26.5)",
                "CodeQL (go)",
                "CodeQL (python)",
                "gitleaks",
            },
            contexts,
        )

        tags = rulesets["Protect release tags"]
        self.assertEqual("tag", tags["target"])
        self.assertEqual(["refs/tags/v*"], tags["include"])
        self.assertEqual(
            {"deletion", "non_fast_forward"},
            {rule["type"] for rule in tags["rules"]},
        )

    def test_source_verifier_and_operator_documentation_are_present(self) -> None:
        proc = subprocess.run(
            ["python3", str(VERIFIER_PATH), "--mode", "source"],
            cwd=ROOT,
            text=True,
            capture_output=True,
        )
        self.assertEqual(0, proc.returncode, proc.stdout + proc.stderr)

        documentation = DOC_PATH.read_text(encoding="utf-8")
        for phrase in (
            "What this protects",
            "Why administrators can bypass",
            "How to verify",
            "immutable releases",
            "secret scanning",
            "Dependabot security updates",
            "github-governance-policy.json",
        ):
            self.assertIn(phrase, documentation)

        evidence = EVIDENCE_PATH.read_text(encoding="utf-8")
        for phrase in (
            "2026-07-26",
            "Protect master",
            "Protect release tags",
            "19778148",
            "19778147",
            "sha_pinning_required",
            "Plan limitations",
            "--mode live",
        ):
            self.assertIn(phrase, evidence)

    def test_contract_is_wired_into_the_release_gate(self) -> None:
        runner = RELEASE_CONTRACT_RUNNER.read_text(encoding="utf-8")
        self.assertIn(
            '["python3", "test/test_github_governance_policy.py"]',
            runner,
        )
        package = json.loads((ROOT / "package.json").read_text(encoding="utf-8"))
        self.assertEqual(
            "python3 scripts/security/verify-github-governance.py --mode source",
            package["scripts"]["security:github-governance"],
        )
        self.assertEqual(
            "python3 scripts/security/verify-github-governance.py --mode live",
            package["scripts"]["security:github-governance:live"],
        )


if __name__ == "__main__":
    unittest.main()
