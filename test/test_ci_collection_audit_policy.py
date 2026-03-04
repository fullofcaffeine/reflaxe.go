#!/usr/bin/env python3

from __future__ import annotations

import unittest
from pathlib import Path


REPO_ROOT = Path(__file__).resolve().parent.parent


class CiCollectionAuditPolicyTest(unittest.TestCase):
    def test_ci_harness_defaults_enforce_full_scope_collection_audit(self) -> None:
        workflow = (REPO_ROOT / ".github" / "workflows" / "ci-harness.yml").read_text(encoding="utf-8")
        self.assertIn('GO_METAL_COLLECTION_AUDIT_ENFORCE: "1"', workflow)
        self.assertIn('GO_METAL_COLLECTION_AUDIT_MAX: "0"', workflow)


if __name__ == "__main__":
    unittest.main()
