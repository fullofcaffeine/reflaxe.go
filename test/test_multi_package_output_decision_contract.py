#!/usr/bin/env python3

from __future__ import annotations

import unittest
from pathlib import Path


REPO_ROOT = Path(__file__).resolve().parent.parent


class MultiPackageOutputDecisionContractTest(unittest.TestCase):
    def test_decision_record_is_scoped_to_current_bead(self) -> None:
        decision_doc = (REPO_ROOT / "docs" / "multi-package-output-evaluation.md").read_text(encoding="utf-8")
        self.assertIn("haxe.go-l3qt.4", decision_doc)
        self.assertIn("Decision: defer multi-package output for production GA.", decision_doc)

    def test_decision_record_has_reopen_boundary_conditions(self) -> None:
        decision_doc = (REPO_ROOT / "docs" / "multi-package-output-evaluation.md").read_text(encoding="utf-8")
        self.assertIn("## Boundary conditions that re-open this work", decision_doc)
        self.assertIn("import-graph planner", decision_doc)
        self.assertIn("cycle-breaking strategy", decision_doc)
        self.assertIn("single-package compatibility mode", decision_doc)

    def test_decision_record_has_measurable_production_reopen_triggers(self) -> None:
        decision_doc = (REPO_ROOT / "docs" / "multi-package-output-evaluation.md").read_text(encoding="utf-8")
        self.assertIn("## Measurable production reopen triggers", decision_doc)
        self.assertIn("Generated file size", decision_doc)
        self.assertIn("Go compile time", decision_doc)
        self.assertIn("Package-private boundary needs", decision_doc)
        self.assertIn("Go tooling limit", decision_doc)
        self.assertIn("Trigger evidence", decision_doc)

    def test_release_checklist_links_to_multi_package_reopen_triggers(self) -> None:
        checklist = (REPO_ROOT / "docs" / "release-readiness-checklist.md").read_text(encoding="utf-8")
        self.assertIn("docs/multi-package-output-evaluation.md#measurable-production-reopen-triggers", checklist)

    def test_known_gaps_marks_multi_package_as_non_blocking_with_boundaries(self) -> None:
        known_gaps_doc = (REPO_ROOT / "docs" / "known-gaps.md").read_text(encoding="utf-8")
        self.assertIn("non-blocking for production GA", known_gaps_doc)
        self.assertIn("docs/multi-package-output-evaluation.md", known_gaps_doc)
        self.assertIn("docs/multi-package-output-evaluation.md#measurable-production-reopen-triggers", known_gaps_doc)


if __name__ == "__main__":
    unittest.main()
