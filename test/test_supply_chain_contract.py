#!/usr/bin/env python3

from __future__ import annotations

import json
import re
import unittest
from pathlib import Path


ROOT = Path(__file__).resolve().parent.parent
WORKFLOW_ROOT = ROOT / ".github"
ACTION_LOCK = ROOT / ".github" / "actions-lock.json"
DEPENDABOT = ROOT / ".github" / "dependabot.yml"
PACKAGE_JSON = ROOT / "package.json"
PACKAGE_LOCK = ROOT / "package-lock.json"
CI_HARNESS = ROOT / ".github" / "workflows" / "ci-harness.yml"
DEPENDENCY_AUDIT = ROOT / "scripts" / "security" / "run-dependency-audit.sh"
SUPPLY_CHAIN_RUNNER = ROOT / "scripts" / "security" / "verify-supply-chain.py"
SUPPLY_CHAIN_DOC = ROOT / "docs" / "supply-chain-policy.md"
OPERATIONAL_AUDIT_REVIEW = ROOT / "docs" / "reviews" / "npm-operational-dependency-audit-vfp-4.12.md"
RELEASE_STATUS = ROOT / "scripts" / "release" / "check-release-state.sh"
RELEASE_CONTRACTS = ROOT / "test" / "run-release-contracts.py"


EXPECTED_ACTIONS = {
    "actions/checkout": ("v6.0.3", "df4cb1c069e1874edd31b4311f1884172cec0e10"),
    "actions/dependency-review-action": ("v5.0.0", "a1d282b36b6f3519aa1f3fc636f609c47dddb294"),
    "actions/download-artifact": ("v8.0.1", "3e5f45b2cfb9172054b4087a40e8e0b5a5461e7c"),
    "actions/setup-go": ("v6.5.0", "924ae3a1cded613372ab5595356fb5720e22ba16"),
    "actions/setup-node": ("v6.5.0", "249970729cb0ef3589644e2896645e5dc5ba9c38"),
    "actions/upload-artifact": ("v7.0.1", "043fb46d1a93c77aae656e7c1c64a875d1fc6a0a"),
    "github/codeql-action": ("v4.37.0", "99df26d4f13ea111d4ec1a7dddef6063f76b97e9"),
    "gitleaks/gitleaks-action": ("v3.0.0", "e0c47f4f8be36e29cdc102c57e68cb5cbf0e8d1e"),
    "krdlab/setup-haxe": ("v2.1.0", "d93667502be3b4f31a94a3308a74388f2e178a8d"),
    "softprops/action-gh-release": ("v3.0.2", "3d0d9888cb7fd7b750713d6e236d1fcb99157228"),
}

USES_RE = re.compile(r"^\s*uses:\s*([^\s#]+)(?:\s+#\s*(\S+))?\s*$")

EXPECTED_OPERATIONAL_AUDIT_BASELINE = {
    "@sigstore/core": "moderate",
    "@sigstore/sign": "moderate",
    "@sigstore/verify": "moderate",
    "brace-expansion": "moderate",
    "diff": "low",
    "glob": "high",
    "handlebars": "critical",
    "ip-address": "moderate",
    "js-yaml": "moderate",
    "libnpmdiff": "high",
    "libnpmpublish": "high",
    "lodash-es": "high",
    "minimatch": "high",
    "npm": "high",
    "pacote": "high",
    "picomatch": "high",
    "sigstore": "high",
    "socks": "moderate",
    "tar": "high",
}


class SupplyChainContractTest(unittest.TestCase):
    def test_npm_lock_is_committed_and_matches_declared_root_metadata(self) -> None:
        package = json.loads(PACKAGE_JSON.read_text(encoding="utf-8"))
        lock = json.loads(PACKAGE_LOCK.read_text(encoding="utf-8"))
        root = lock["packages"][""]

        self.assertEqual(lock["lockfileVersion"], 3)
        self.assertEqual(lock["name"], package["name"])
        self.assertEqual(lock["version"], package["version"])
        self.assertEqual(root["name"], package["name"])
        self.assertEqual(root["version"], package["version"])
        self.assertEqual(root.get("dependencies", {}), package.get("dependencies", {}))
        self.assertEqual(root.get("devDependencies", {}), package.get("devDependencies", {}))

    def test_ci_uses_clean_locked_installs_only(self) -> None:
        npm_install = re.compile(r"\bnpm\s+(?:install|i)\b")
        install_lines: list[str] = []
        ci_lines: list[str] = []
        for path in sorted((WORKFLOW_ROOT / "workflows").glob("*.yml")):
            for line in path.read_text(encoding="utf-8").splitlines():
                if npm_install.search(line):
                    install_lines.append(f"{path.name}: {line.strip()}")
                if re.search(r"\bnpm\s+ci\b", line):
                    ci_lines.append(f"{path.name}: {line.strip()}")

        self.assertEqual(install_lines, [])
        self.assertGreaterEqual(len(ci_lines), 5)

        dependency_audit = DEPENDENCY_AUDIT.read_text(encoding="utf-8")
        self.assertIn('cp package-lock.json "$npm_audit_tmp_dir/package-lock.json"', dependency_audit)
        self.assertIn("npm ci --ignore-scripts --include=dev", dependency_audit)
        self.assertIn("npm audit --include=dev --audit-level=high", dependency_audit)
        self.assertNotIn("--omit=dev", dependency_audit)
        self.assertNotIn("npm install --ignore-scripts --package-lock-only", dependency_audit)

    def test_every_external_action_is_manifested_and_commit_pinned(self) -> None:
        lock = json.loads(ACTION_LOCK.read_text(encoding="utf-8"))
        locked = {
            entry["repository"]: (entry["version"], entry["commit"])
            for entry in lock["actions"]
        }
        self.assertEqual(locked, EXPECTED_ACTIONS)

        seen: set[str] = set()
        for path in sorted(WORKFLOW_ROOT.rglob("*.yml")):
            for line_number, line in enumerate(path.read_text(encoding="utf-8").splitlines(), 1):
                match = USES_RE.match(line)
                if match is None:
                    continue
                use, comment = match.groups()
                if use.startswith("./"):
                    continue
                action, commit = use.rsplit("@", 1)
                repository = "/".join(action.split("/")[:2])
                self.assertRegex(commit, r"^[0-9a-f]{40}$", f"{path}:{line_number}")
                self.assertIn(repository, locked, f"{path}:{line_number}")
                expected_version, expected_commit = locked[repository]
                self.assertEqual(commit, expected_commit, f"{path}:{line_number}")
                self.assertEqual(comment, expected_version, f"{path}:{line_number}")
                seen.add(repository)

        self.assertEqual(seen, set(locked))

    def test_dependabot_tracks_action_and_npm_updates(self) -> None:
        dependabot = DEPENDABOT.read_text(encoding="utf-8")
        self.assertIn('package-ecosystem: "github-actions"', dependabot)
        self.assertIn('package-ecosystem: "npm"', dependabot)
        self.assertGreaterEqual(dependabot.count('interval: "weekly"'), 2)

    def test_operational_dependency_audit_records_complete_baseline_triage(self) -> None:
        review = OPERATIONAL_AUDIT_REVIEW.read_text(encoding="utf-8")
        rows = re.findall(
            r"^\| `([^`]+)` \| `(critical|high|moderate|low)` \|",
            review,
            flags=re.MULTILINE,
        )

        self.assertEqual(dict(rows), EXPECTED_OPERATIONAL_AUDIT_BASELINE)
        self.assertEqual(len(rows), 19)
        self.assertIn("configured execution path", review)
        self.assertIn("installed but inactive", review)
        self.assertIn("npm audit --include=dev --audit-level=high", review)
        self.assertIn("found 0 vulnerabilities", review)

    def test_supply_chain_gate_is_release_blocking_and_documented(self) -> None:
        package = json.loads(PACKAGE_JSON.read_text(encoding="utf-8"))
        self.assertEqual(
            package["scripts"]["security:supply-chain"],
            "python3 scripts/security/verify-supply-chain.py",
        )
        self.assertEqual(
            package["scripts"]["verify:vendor-reflaxe"],
            "python3 scripts/vendor/verify-reflaxe-provenance.py",
        )
        self.assertEqual(
            package["scripts"]["verify:vendor-reflaxe:reconstruct"],
            "python3 scripts/vendor/verify-reflaxe-provenance.py --reconstruct",
        )

        runner = SUPPLY_CHAIN_RUNNER.read_text(encoding="utf-8")
        self.assertIn("actions-lock.json", runner)
        self.assertIn("package-lock.json", runner)
        self.assertIn("verify-reflaxe-provenance.py", runner)

        workflow = CI_HARNESS.read_text(encoding="utf-8")
        self.assertIn("Verify locked dependencies and vendored provenance", workflow)
        self.assertIn("npm run security:supply-chain", workflow)

        release_status = RELEASE_STATUS.read_text(encoding="utf-8")
        self.assertIn("scripts/security/verify-supply-chain.py", release_status)
        self.assertIn("supply-chain provenance: OK", release_status)

        contracts = RELEASE_CONTRACTS.read_text(encoding="utf-8")
        self.assertIn("test/test_supply_chain_contract.py", contracts)
        self.assertIn("test/test_vendor_reflaxe_provenance.py", contracts)

        policy = SUPPLY_CHAIN_DOC.read_text(encoding="utf-8")
        self.assertIn("# Supply-Chain Policy", policy)
        self.assertIn("npm ci", policy)
        self.assertIn("actions-lock.json", policy)
        self.assertIn("Dependabot", policy)
        self.assertIn("vendor Reflaxe", policy)
        self.assertIn("npm run security:supply-chain", policy)


if __name__ == "__main__":
    unittest.main()
