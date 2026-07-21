#!/usr/bin/env python3

from __future__ import annotations

import hashlib
import json
import re
import shutil
import subprocess
import sys
import tempfile
import unittest
from pathlib import Path


REPO_ROOT = Path(__file__).resolve().parent.parent
AGENTS = REPO_ROOT / "AGENTS.md"
CLAUDE = REPO_ROOT / "CLAUDE.md"
BEADS_README = REPO_ROOT / ".beads" / "README.md"
BEADS_CONFIG = REPO_ROOT / ".beads" / "config.yaml"
LEGACY_ARCHIVE = REPO_ROOT / ".beads" / "issues.jsonl"
RELEASE_CONTRACT_RUNNER = REPO_ROOT / "test" / "run-release-contracts.py"
HEALTH_SCRIPT = REPO_ROOT / "scripts" / "beads" / "check-health.sh"
HIERARCHY_DEADLOCK_GUARD = REPO_ROOT / "scripts" / "beads" / "check-hierarchy-deadlocks.py"
MIGRATION_INTEGRITY_RECORD = REPO_ROOT / "docs" / "reviews" / "gpt-5.6-pro" / "beads-v53-migration-integrity.md"
PUBLIC_CONTRACT_DECISION_LOG = (
    REPO_ROOT / "docs" / "reviews" / "gpt-5.6-pro" / "decision-log-vfp-6.3-public-contract.tsv"
)

LEGACY_ARCHIVE_SHA256 = "0e34e32cb1ac25fdc8592aea85aa5630ca31ab59076b3e33faa6611a4e51911c"


class BeadsWorkflowContractTest(unittest.TestCase):
    def read(self, path: Path) -> str:
        return path.read_text(encoding="utf-8")

    def test_agent_docs_use_current_dolt_commands(self) -> None:
        for path in [AGENTS, CLAUDE, BEADS_README]:
            text = self.read(path)
            self.assertNotRegex(text, r"(?m)^\s*bd sync(?:\s|$)", str(path))
            self.assertIn("bd prime", text, str(path))
            self.assertIn("bd update <id> --claim", text, str(path))
            self.assertIn("bd dolt pull", text, str(path))
            self.assertIn("bd dolt push", text, str(path))

        agents = self.read(AGENTS)
        self.assertNotIn("bd update <id> --status in_progress", agents)

    def test_tracker_readme_explains_storage_provenance_and_recovery(self) -> None:
        text = self.read(BEADS_README)
        for heading in ["## What it is", "## Why it exists", "## How it works", "## Recovery"]:
            self.assertIn(heading, text)

        required_phrases = [
            "refs/dolt/data",
            "operational source of truth",
            "canonical legacy provenance archive",
            "578 live legacy issues",
            "haxe.go-dsn",
            LEGACY_ARCHIVE_SHA256,
            "export.auto: false",
            "bd bootstrap --yes",
            "bd backup init",
            "bd backup restore",
            "chmod 700 .beads",
        ]
        for phrase in required_phrases:
            self.assertIn(phrase, text)

    def test_beads_config_protects_legacy_archive_and_remote_history(self) -> None:
        text = self.read(BEADS_CONFIG)
        required_settings = [
            'sync.remote: "git+https://github.com/fullofcaffeine/reflaxe.go.git"',
            'federation.remote: "git+https://github.com/fullofcaffeine/reflaxe.go.git"',
            "export.auto: false",
            "export.git-add: false",
            "sync.require_confirmation_on_mass_delete: true",
        ]
        for setting in required_settings:
            self.assertIn(setting, text)

    def test_legacy_archive_identity_and_tombstone_are_immutable(self) -> None:
        raw = LEGACY_ARCHIVE.read_bytes()
        self.assertEqual(hashlib.sha256(raw).hexdigest(), LEGACY_ARCHIVE_SHA256)

        records = [json.loads(line) for line in raw.decode("utf-8").splitlines() if line.strip()]
        ids = [record["id"] for record in records]
        self.assertEqual(len(records), 579)
        self.assertEqual(len(set(ids)), 579)

        tombstones = [record for record in records if record.get("status") == "tombstone"]
        self.assertEqual([record["id"] for record in tombstones], ["haxe.go-dsn"])
        self.assertEqual(sum(record.get("status") == "closed" for record in records), 578)

    def test_contract_runs_in_the_release_contract_surface(self) -> None:
        runner = self.read(RELEASE_CONTRACT_RUNNER)
        self.assertIn('test/test_beads_workflow_contract.py', runner)

    def test_health_script_uses_supported_read_only_checks(self) -> None:
        self.assertTrue(HEALTH_SCRIPT.exists(), "tracker health script must exist")
        self.assertNotEqual(HEALTH_SCRIPT.stat().st_mode & 0o111, 0, "tracker health script must be executable")

        script = self.read(HEALTH_SCRIPT)
        required_fragments = [
            "bd config validate",
            "bd dep cycles",
            "bd lint",
            "bd orphans",
            "bd vc status --json",
            "git ls-remote",
            "bd export --include-memories",
            "--session-close",
            "--verify-remote",
        ]
        for fragment in required_fragments:
            self.assertIn(fragment, script)

        readme = self.read(BEADS_README)
        self.assertIn("scripts/beads/check-health.sh", readme)
        self.assertIn("bd doctor", readme)
        self.assertIn("embedded mode", readme)
        self.assertIn("bd preflight --check", readme)

    def test_health_script_help_is_side_effect_free(self) -> None:
        if not HEALTH_SCRIPT.exists():
            self.fail("tracker health script must exist")
        completed = subprocess.run([str(HEALTH_SCRIPT), "--help"], cwd=REPO_ROOT, capture_output=True, text=True)
        self.assertEqual(completed.returncode, 0, completed.stderr)
        self.assertIn("--session-close", completed.stdout)
        self.assertIn("--verify-remote", completed.stdout)

    def run_hierarchy_guard(self, records: list[dict[str, object]]) -> subprocess.CompletedProcess[str]:
        with tempfile.TemporaryDirectory(prefix="haxe-go-beads-deadlock-test-") as raw:
            fixture = Path(raw) / "issues.jsonl"
            fixture.write_text(
                "".join(json.dumps(record) + "\n" for record in records),
                encoding="utf-8",
            )
            return subprocess.run(
                [sys.executable, str(HIERARCHY_DEADLOCK_GUARD), "--input", str(fixture)],
                cwd=REPO_ROOT,
                capture_output=True,
                text=True,
            )

    def test_hierarchy_deadlock_guard_rejects_active_ancestor_blocking_edges(self) -> None:
        records = [
            {
                "id": "fixture-root",
                "title": "Root",
                "status": "open",
                "dependencies": [
                    {
                        "issue_id": "fixture-root",
                        "depends_on_id": "fixture-root.1",
                        "type": "blocks",
                    }
                ],
            },
            {
                "id": "fixture-root.1",
                "title": "Child",
                "status": "open",
                "dependencies": [
                    {
                        "issue_id": "fixture-root.1",
                        "depends_on_id": "fixture-root",
                        "type": "parent-child",
                    }
                ],
            },
            {
                "id": "fixture-root.2",
                "title": "Other child",
                "status": "open",
                "dependencies": [
                    {
                        "issue_id": "fixture-root.2",
                        "depends_on_id": "fixture-root",
                        "type": "parent-child",
                    },
                    {
                        "issue_id": "fixture-root.2",
                        "depends_on_id": "fixture-root",
                        "type": "conditional-blocks",
                    },
                ],
            },
        ]

        completed = self.run_hierarchy_guard(records)
        self.assertNotEqual(completed.returncode, 0)
        self.assertIn("parent-depends-on-descendant", completed.stderr)
        self.assertIn("fixture-root -> fixture-root.1", completed.stderr)
        self.assertIn("child-depends-on-ancestor", completed.stderr)
        self.assertIn("fixture-root.2 -> fixture-root", completed.stderr)

    def test_hierarchy_deadlock_guard_allows_sibling_ordering_and_inactive_history(self) -> None:
        records = [
            {"id": "fixture-root", "title": "Root", "status": "open"},
            {
                "id": "fixture-root.1",
                "title": "First child",
                "status": "open",
                "dependencies": [
                    {
                        "issue_id": "fixture-root.1",
                        "depends_on_id": "fixture-root",
                        "type": "parent-child",
                    },
                    {
                        "issue_id": "fixture-root.1",
                        "depends_on_id": "fixture-root.2",
                        "type": "blocks",
                    },
                ],
            },
            {
                "id": "fixture-root.2",
                "title": "Sibling",
                "status": "open",
                "dependencies": [
                    {
                        "issue_id": "fixture-root.2",
                        "depends_on_id": "fixture-root",
                        "type": "parent-child",
                    }
                ],
            },
            {
                "id": "fixture-closed",
                "title": "Closed parent",
                "status": "closed",
                "dependencies": [
                    {
                        "issue_id": "fixture-closed",
                        "depends_on_id": "fixture-closed.1",
                        "type": "blocks",
                    }
                ],
            },
            {
                "id": "fixture-closed.1",
                "title": "Closed child",
                "status": "closed",
                "dependencies": [
                    {
                        "issue_id": "fixture-closed.1",
                        "depends_on_id": "fixture-closed",
                        "type": "parent-child",
                    }
                ],
            },
        ]

        completed = self.run_hierarchy_guard(records)
        self.assertEqual(completed.returncode, 0, completed.stderr)
        self.assertIn("active=0", completed.stdout)
        self.assertIn("inactive=1", completed.stdout)

    def test_health_script_runs_hierarchy_deadlock_guard(self) -> None:
        script = self.read(HEALTH_SCRIPT)
        self.assertIn("scripts/beads/check-hierarchy-deadlocks.py", script)

    def test_hierarchy_and_readiness_model_is_documented(self) -> None:
        readme = self.read(BEADS_README)
        for phrase in [
            "## Hierarchy and readiness",
            "parent-child",
            "ancestor",
            "descendant",
            "sibling",
            "scripts/beads/check-hierarchy-deadlocks.py",
        ]:
            self.assertIn(phrase, readme)

        migration = self.read(MIGRATION_INTEGRITY_RECORD)
        self.assertNotIn(
            "The new client treats structural parent-child links as affecting readiness",
            migration,
        )
        self.assertIn("feedback loop", migration)
        self.assertIn("b5365b7dbbf7609b1cbe3d80776f76702c14c2d5c246d82852ad5d2cfbbbadab", migration)

        decision_log = self.read(PUBLIC_CONTRACT_DECISION_LOG)
        self.assertIn("Initial hypothesis; disproved by haxe_go-7hiq", decision_log)
        self.assertIn("Correct the ready-list diagnosis", decision_log)

    def test_documented_bd_commands_exist_when_bd_is_installed(self) -> None:
        bd = shutil.which("bd")
        if bd is None:
            self.skipTest("bd is not installed in this environment")

        commands = [
            [bd, "prime", "--help"],
            [bd, "dolt", "pull", "--help"],
            [bd, "dolt", "push", "--help"],
            [bd, "bootstrap", "--help"],
            [bd, "backup", "restore", "--help"],
        ]
        for command in commands:
            completed = subprocess.run(command, cwd=REPO_ROOT, capture_output=True, text=True)
            self.assertEqual(completed.returncode, 0, " ".join(command))


if __name__ == "__main__":
    unittest.main()
