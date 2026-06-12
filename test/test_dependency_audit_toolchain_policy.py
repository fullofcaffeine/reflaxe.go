#!/usr/bin/env python3

from __future__ import annotations

import unittest
from pathlib import Path


REPO_ROOT = Path(__file__).resolve().parent.parent
DEPENDENCY_AUDIT = REPO_ROOT / "scripts" / "security" / "run-dependency-audit.sh"


class DependencyAuditToolchainPolicyTest(unittest.TestCase):
    def test_govulncheck_default_is_pinned_to_go_123_compatible_release(self) -> None:
        script = DEPENDENCY_AUDIT.read_text(encoding="utf-8")
        self.assertIn('govulncheck_version="${GOVULNCHECK_VERSION:-v1.1.4}"', script)
        self.assertIn("v1.2.0+ currently requires Go 1.25+", script)
        self.assertNotIn('govulncheck_version="${GOVULNCHECK_VERSION:-latest}"', script)


if __name__ == "__main__":
    unittest.main()
