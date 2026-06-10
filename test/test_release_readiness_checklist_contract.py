#!/usr/bin/env python3

from __future__ import annotations

import unittest
from pathlib import Path


REPO_ROOT = Path(__file__).resolve().parent.parent


class ReleaseReadinessChecklistContractTest(unittest.TestCase):
    def test_checklist_doc_exists_with_required_sections(self) -> None:
        checklist = (REPO_ROOT / "docs" / "release-readiness-checklist.md").read_text(encoding="utf-8")
        self.assertIn("# Release Readiness Checklist", checklist)
        self.assertIn("## Required GA gates", checklist)
        self.assertIn("## Reproducible command set", checklist)
        self.assertIn("## Pass criteria", checklist)

    def test_checklist_includes_canonical_commands(self) -> None:
        checklist = (REPO_ROOT / "docs" / "release-readiness-checklist.md").read_text(encoding="utf-8")
        self.assertIn("python3 test/run-ci.py", checklist)
        self.assertIn("npm run test:stdlib:governance", checklist)
        self.assertIn("npm run test:release-contracts", checklist)
        self.assertIn("npm run release:status", checklist)
        self.assertIn("npm run test:perf:go", checklist)
        self.assertIn("npm run test:perf:hxrt-selective", checklist)
        self.assertIn("npm run test:perf:apps", checklist)

    def test_checklist_uses_current_parity_closure_wording(self) -> None:
        checklist = (REPO_ROOT / "docs" / "release-readiness-checklist.md").read_text(encoding="utf-8")
        self.assertNotIn("explicit blocker-backed remaining modules", checklist)
        self.assertIn("0 actionable blockers", checklist)
        self.assertIn("policy-locked", checklist)

    def test_checklist_links_to_release_visibility_and_run_ci(self) -> None:
        checklist = (REPO_ROOT / "docs" / "release-readiness-checklist.md").read_text(encoding="utf-8")
        self.assertIn("docs/release-visibility.md", checklist)
        self.assertIn("test/run-ci.py", checklist)

    def test_start_here_references_checklist(self) -> None:
        start_here = (REPO_ROOT / "docs" / "start-here.md").read_text(encoding="utf-8")
        self.assertIn("docs/release-readiness-checklist.md", start_here)

    def test_ci_runner_has_release_contracts_stage(self) -> None:
        ci_runner = (REPO_ROOT / "test" / "run-ci.py").read_text(encoding="utf-8")
        self.assertIn("==> Release contracts stage", ci_runner)
        self.assertIn("def build_release_contracts_command", ci_runner)
        self.assertIn("test/run-release-contracts.py", ci_runner)


if __name__ == "__main__":
    unittest.main()
