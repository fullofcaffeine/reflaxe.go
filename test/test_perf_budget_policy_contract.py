#!/usr/bin/env python3

from __future__ import annotations

import unittest
from pathlib import Path


REPO_ROOT = Path(__file__).resolve().parent.parent
POLICY_DOC = REPO_ROOT / "docs" / "performance-budget-policy.md"
KNOWN_GAPS_DOC = REPO_ROOT / "docs" / "known-gaps.md"
RELEASE_CHECKLIST_DOC = REPO_ROOT / "docs" / "release-readiness-checklist.md"
CI_HARNESS = REPO_ROOT / ".github" / "workflows" / "ci-harness.yml"


class PerfBudgetPolicyContractTest(unittest.TestCase):
    def test_policy_doc_defines_budget_terms_and_current_decisions(self) -> None:
        policy = POLICY_DOC.read_text(encoding="utf-8")
        self.assertIn("# Performance Budget Policy", policy)
        self.assertIn("## Terms", policy)
        self.assertIn("soft warning", policy)
        self.assertIn("hard gate", policy)
        self.assertIn("baseline", policy)
        self.assertIn("portable-vs-metal delta", policy)
        self.assertIn("pure-Go baseline", policy)
        self.assertIn("## Current Release Policy", policy)
        self.assertIn("Portable regressions stay warning-only by default", policy)
        self.assertIn("Metal microbench regressions are release-blocking in CI", policy)
        self.assertIn("Flagship app metal regressions stay warning-only by default", policy)
        self.assertIn("HXRT selective runtime drift is release-blocking in CI", policy)
        self.assertIn("Do not update perf baselines just to make warnings disappear", policy)

    def test_policy_doc_lists_harnesses_and_enforcement_knobs(self) -> None:
        policy = POLICY_DOC.read_text(encoding="utf-8")
        for command in (
            "npm run test:perf:go",
            "npm run test:perf:apps",
            "npm run test:perf:hxrt-selective",
        ):
            self.assertIn(command, policy)
        for define in (
            "GO_PERF_ENFORCE_METAL_BUDGET",
            "GO_PERF_ENFORCE_DELTA_BUDGET",
            "GO_APP_PERF_ENFORCE_METAL_BUDGET",
            "GO_APP_PERF_ENFORCE_DELTA_BUDGET",
            "GO_HXRT_SLICE_ENFORCE",
        ):
            self.assertIn(define, policy)

    def test_release_docs_link_perf_policy(self) -> None:
        known_gaps = KNOWN_GAPS_DOC.read_text(encoding="utf-8")
        checklist = RELEASE_CHECKLIST_DOC.read_text(encoding="utf-8")
        self.assertIn("docs/performance-budget-policy.md", known_gaps)
        self.assertIn("docs/performance-budget-policy.md", checklist)
        self.assertIn("Performance budget policy", checklist)

    def test_ci_enforcement_posture_is_explicit(self) -> None:
        workflow = CI_HARNESS.read_text(encoding="utf-8")
        self.assertIn("GO_PERF_ENFORCE_METAL_BUDGET: \"1\"", workflow)
        self.assertIn("GO_PERF_ENFORCE_DELTA_BUDGET: \"0\"", workflow)
        self.assertIn("GO_HXRT_SLICE_ENFORCE: \"1\"", workflow)
        self.assertIn("GO_APP_PERF_ENFORCE_METAL_BUDGET: \"0\"", workflow)
        self.assertIn("GO_APP_PERF_ENFORCE_DELTA_BUDGET: \"0\"", workflow)


if __name__ == "__main__":
    unittest.main()
