#!/usr/bin/env python3

from __future__ import annotations

import unittest
from pathlib import Path


REPO_ROOT = Path(__file__).resolve().parent.parent
POLICY_DOC = REPO_ROOT / "docs" / "performance-budget-policy.md"
TRIAGE_DOC = REPO_ROOT / "docs" / "perf-warning-triage.md"
KNOWN_GAPS_DOC = REPO_ROOT / "docs" / "known-gaps.md"
RELEASE_CHECKLIST_DOC = REPO_ROOT / "docs" / "release-readiness-checklist.md"
DOC_INDEX = REPO_ROOT / "docs" / "index.md"
CI_HARNESS = REPO_ROOT / ".github" / "workflows" / "ci-harness.yml"
GO_PERF_SCRIPT = REPO_ROOT / "scripts" / "ci" / "perf-go-profiles.sh"
APP_PERF_SCRIPT = REPO_ROOT / "scripts" / "ci" / "perf-apps.sh"


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
        self.assertIn("[soft-budget-signal]", policy)
        self.assertIn("[not-enforced]", policy)
        self.assertIn("docs/perf-warning-triage.md", policy)

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
        index = DOC_INDEX.read_text(encoding="utf-8")
        self.assertIn("docs/performance-budget-policy.md", known_gaps)
        self.assertIn("docs/performance-budget-policy.md", checklist)
        self.assertIn("Performance budget policy", checklist)
        self.assertIn("perf-warning-triage.md", index)

    def test_ci_enforcement_posture_is_explicit(self) -> None:
        workflow = CI_HARNESS.read_text(encoding="utf-8")
        self.assertIn("GO_PERF_ENFORCE_METAL_BUDGET: \"1\"", workflow)
        self.assertIn("GO_PERF_ENFORCE_DELTA_BUDGET: \"0\"", workflow)
        self.assertIn("GO_HXRT_SLICE_ENFORCE: \"1\"", workflow)
        self.assertIn("GO_APP_PERF_ENFORCE_METAL_BUDGET: \"0\"", workflow)
        self.assertIn("GO_APP_PERF_ENFORCE_DELTA_BUDGET: \"0\"", workflow)

    def test_perf_ci_annotations_explain_warning_only_policy(self) -> None:
        scripts = GO_PERF_SCRIPT.read_text(encoding="utf-8") + "\n" + APP_PERF_SCRIPT.read_text(encoding="utf-8")
        self.assertIn("soft-budget-signal", scripts)
        self.assertIn("warning-only; see docs/performance-budget-policy.md", scripts)
        self.assertIn("not-enforced", scripts)
        self.assertIn("hard gate is disabled", scripts)

    def test_perf_warning_triage_records_current_decision(self) -> None:
        triage = TRIAGE_DOC.read_text(encoding="utf-8")
        self.assertIn("# Perf Warning Triage", triage)
        self.assertIn("June 13, 2026", triage)
        self.assertIn("Do not update baselines", triage)
        self.assertIn("Do not promote app perf warnings to hard gates", triage)
        self.assertIn("Go profile microbench", triage)
        self.assertIn("warnings=0", triage)
        self.assertIn("FluxProxy", triage)
        self.assertIn("multi-run startup variance", triage)
        self.assertIn("haxe.go-nhh2.4", triage)


if __name__ == "__main__":
    unittest.main()
