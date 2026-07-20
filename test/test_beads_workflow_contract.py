#!/usr/bin/env python3

from __future__ import annotations

import hashlib
import json
import re
import shutil
import subprocess
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
