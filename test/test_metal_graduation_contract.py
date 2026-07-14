#!/usr/bin/env python3

from __future__ import annotations

import unittest
from pathlib import Path


REPO_ROOT = Path(__file__).resolve().parent.parent


class MetalGraduationContractTest(unittest.TestCase):
    def test_profiles_doc_no_longer_marks_metal_experimental(self) -> None:
        profiles_doc = (REPO_ROOT / "docs" / "profiles.md").read_text(encoding="utf-8")
        self.assertNotIn("`metal` (experimental)", profiles_doc)

    def test_known_gaps_doc_no_longer_marks_metal_experimental(self) -> None:
        known_gaps_doc = (REPO_ROOT / "docs" / "known-gaps.md").read_text(encoding="utf-8")
        self.assertNotIn("`metal` is still marked experimental.", known_gaps_doc)

    def test_ci_has_explicit_metal_fallback_diagnostics_stage(self) -> None:
        ci_runner = (REPO_ROOT / "test" / "run-ci.py").read_text(encoding="utf-8")
        self.assertIn("==> Metal fallback diagnostics stage", ci_runner)
        self.assertIn("def build_metal_fallback_diagnostics_command", ci_runner)
        self.assertIn("core/report_artifacts_lane_fallback", ci_runner)

    def test_profile_guide_has_compatibility_migration_and_deprecation_gate(self) -> None:
        semantics_guide = (REPO_ROOT / "docs" / "profile-semantics-guide.md").read_text(encoding="utf-8")
        self.assertIn("## Migrating an existing metal build", semantics_guide)
        self.assertIn("There is no requirement to migrate today", semantics_guide)
        self.assertIn("## Review boundary", semantics_guide)
        self.assertIn("genuine independent deep review", semantics_guide)
        self.assertIn("SemVer migration plan", semantics_guide)


if __name__ == "__main__":
    unittest.main()
