#!/usr/bin/env python3

from __future__ import annotations

import unittest
from pathlib import Path


REPO_ROOT = Path(__file__).resolve().parent.parent
DEPENDENCY_AUDIT = REPO_ROOT / "scripts" / "security" / "run-dependency-audit.sh"
AUDIT_DOC = REPO_ROOT / "docs" / "security-dependency-audit.md"


class DependencyAuditToolchainPolicyTest(unittest.TestCase):
    def test_govulncheck_default_is_pinned_to_go_123_compatible_release(self) -> None:
        script = DEPENDENCY_AUDIT.read_text(encoding="utf-8")
        self.assertIn('govulncheck_version="${GOVULNCHECK_VERSION:-v1.1.4}"', script)
        self.assertIn("v1.2.0+ currently requires Go 1.25+", script)
        self.assertNotIn('govulncheck_version="${GOVULNCHECK_VERSION:-latest}"', script)

    def test_govulncheck_annotations_are_repo_classified(self) -> None:
        script = DEPENDENCY_AUDIT.read_text(encoding="utf-8")
        self.assertIn("env -u GITHUB_ACTIONS", script)
        self.assertIn("print_classified_govulncheck_report", script)
        self.assertIn("avoid GitHub problem matchers", script)
        self.assertIn("s#([[:alnum:]_./-]+\\.go):([0-9]+):([0-9]+):#\\1 line \\2 col \\3:#g", script)
        self.assertIn("[deps][govulncheck-stdlib-reachability]", script)
        self.assertIn("classified audit evidence", script)
        self.assertIn("continuing after classified audit annotation", script)
        self.assertIn("docs/security-dependency-audit.md", script)

    def test_dependency_audit_doc_explains_stdlib_reachability(self) -> None:
        doc = AUDIT_DOC.read_text(encoding="utf-8")
        self.assertIn("# Security Dependency Audit", doc)
        self.assertIn("Go standard-library vulnerability", doc)
        self.assertIn("SSL and network support", doc)
        self.assertIn("classified audit evidence", doc)
        self.assertIn("not a dependency install failure", doc)


if __name__ == "__main__":
    unittest.main()
