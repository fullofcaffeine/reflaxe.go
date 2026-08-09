#!/usr/bin/env python3

from __future__ import annotations

import unittest
from pathlib import Path


REPO_ROOT = Path(__file__).resolve().parent.parent
RELEASE_CLAIM = (
    "Haxe.Go is a pre-1.0 beta for pinned, application-qualified portable workloads "
    "on the admitted toolchain, platform, and operation/member surface."
)


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
        self.assertIn("npm run compatibility:verify", checklist)
        self.assertIn("npm run release:status", checklist)
        self.assertIn("npm run release:readiness", checklist)
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
        self.assertIn("docs/compatibility-support-manifest.json", checklist)
        self.assertIn("docs/compatibility-release-status.md", checklist)
        self.assertIn("release/readiness-policy.json", checklist)
        self.assertIn("candidate", checklist)
        self.assertIn("published", checklist)
        self.assertIn("GitHub API", checklist)
        self.assertIn("test/run-ci.py", checklist)

    def test_public_release_docs_use_the_bounded_compatibility_claim(self) -> None:
        readme = (REPO_ROOT / "README.md").read_text(encoding="utf-8")
        checklist = (REPO_ROOT / "docs" / "release-readiness-checklist.md").read_text(encoding="utf-8")
        visibility = (REPO_ROOT / "docs" / "release-visibility.md").read_text(encoding="utf-8")

        self.assertIn(RELEASE_CLAIM, readme)
        self.assertIn("docs/compatibility-release-status.md", readme)
        self.assertIn("compatibility-support-manifest.json", visibility)
        self.assertIn("compatibility-release-status.md", visibility)
        for document in (readme, checklist, visibility):
            self.assertNotIn("beta-stable", document.lower())

    def test_start_here_references_checklist(self) -> None:
        start_here = (REPO_ROOT / "docs" / "start-here.md").read_text(encoding="utf-8")
        self.assertIn("docs/release-readiness-checklist.md", start_here)

    def test_ci_runner_has_release_contracts_stage(self) -> None:
        ci_runner = (REPO_ROOT / "test" / "run-ci.py").read_text(encoding="utf-8")
        self.assertIn("==> Release contracts stage", ci_runner)
        self.assertIn("def build_release_contracts_command", ci_runner)
        self.assertIn("test/run-release-contracts.py", ci_runner)

    def test_docs_separate_beta_admission_from_routine_release_proof(self) -> None:
        checklist = (REPO_ROOT / "docs" / "release-readiness-checklist.md").read_text(encoding="utf-8")
        evidence = (REPO_ROOT / "docs" / "release-readiness-evidence.md").read_text(encoding="utf-8")
        version_policy = (REPO_ROOT / "docs" / "release-version-policy.md").read_text(encoding="utf-8")
        for document in (checklist, evidence, version_policy):
            self.assertIn("historical beta baseline", document)
            self.assertIn("routine release", document)
        self.assertIn("v0.54.1", checklist)
        self.assertIn("v0.54.2", checklist)
        self.assertIn("v0.54.3", checklist)
        self.assertIn("unpublished", checklist)
        self.assertIn("fresh review", version_policy)


if __name__ == "__main__":
    unittest.main()
