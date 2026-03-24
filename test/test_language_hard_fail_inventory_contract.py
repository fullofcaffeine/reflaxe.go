#!/usr/bin/env python3

from __future__ import annotations

import unittest
from pathlib import Path


REPO_ROOT = Path(__file__).resolve().parent.parent


class LanguageHardFailInventoryContractTest(unittest.TestCase):
    def test_known_gaps_frames_remaining_lowering_guards_as_invariants(self) -> None:
        known_gaps = (REPO_ROOT / 'docs' / 'known-gaps.md').read_text(encoding='utf-8')
        self.assertIn('These remaining lowering guards are invariant checks, not open supported-language gaps.', known_gaps)
        self.assertIn('No currently supported Haxe source construct is expected to hit them in normal typed lowering.', known_gaps)

    def test_feature_matrix_marks_hard_fail_inventory_as_invariant_only(self) -> None:
        feature_matrix = (REPO_ROOT / 'docs' / 'feature-support-matrix.md').read_text(encoding='utf-8')
        self.assertIn(
            'The remaining inventory is invariant-only: parser/front-end rejection or node-family closure proof, not active supported-language holes.',
            feature_matrix,
        )
        self.assertIn('negative/non_lvalue_assignment_invariant', feature_matrix)
        self.assertIn('negative/postfix_non_inc_dec_invariant', feature_matrix)
        self.assertIn('semantic-diff/type_expr_contract', feature_matrix)
        self.assertIn('std_is_of_type_runtime_core_abstract_contract', feature_matrix)


if __name__ == '__main__':
    unittest.main()
