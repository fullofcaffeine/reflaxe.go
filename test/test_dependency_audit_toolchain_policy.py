#!/usr/bin/env python3

from __future__ import annotations

import unittest
from pathlib import Path


REPO_ROOT = Path(__file__).resolve().parent.parent
DEPENDENCY_AUDIT = REPO_ROOT / "scripts" / "security" / "run-dependency-audit.sh"
AUDIT_DOC = REPO_ROOT / "docs" / "security-dependency-audit.md"


class DependencyAuditToolchainPolicyTest(unittest.TestCase):
    def test_govulncheck_default_pin_is_not_mistaken_for_build_support(self) -> None:
        script = DEPENDENCY_AUDIT.read_text(encoding="utf-8")
        self.assertIn('govulncheck_version="${GOVULNCHECK_VERSION:-v1.6.0}"', script)
        self.assertIn("docs/toolchain-policy.md is authoritative", script)
        self.assertNotIn("Go 1.22/1.23 CI lanes", script)
        self.assertNotIn('govulncheck_version="${GOVULNCHECK_VERSION:-latest}"', script)

    def test_govulncheck_text_gate_preserves_reports_and_fails_findings(self) -> None:
        script = DEPENDENCY_AUDIT.read_text(encoding="utf-8")
        self.assertIn("env -u GITHUB_ACTIONS", script)
        self.assertIn("print_sanitized_govulncheck_report", script)
        self.assertIn("-show traces", script)
        self.assertIn("-format=text", script)
        self.assertIn("GOVULNCHECK_REPORT_DIR", script)
        self.assertIn("avoid GitHub problem matchers", script)
        self.assertIn("s#([[:alnum:]_./-]+\\.go):([0-9]+):([0-9]+):#\\1 line \\2 col \\3:#g", script)
        self.assertIn("reachable vulnerabilities are release-blocking", script)
        self.assertNotIn("continuing after classified audit annotation", script)
        self.assertNotIn("remain visible but non-blocking", script)
        self.assertIn("docs/security-dependency-audit.md", script)

    def test_dependency_audit_doc_explains_stdlib_reachability(self) -> None:
        doc = AUDIT_DOC.read_text(encoding="utf-8")
        self.assertIn("# Security Dependency Audit", doc)
        self.assertIn("Go standard-library vulnerability", doc)
        self.assertIn("SSL and network support", doc)
        self.assertIn("fails closed", doc)
        self.assertIn("toolchain vulnerability", doc)
        self.assertIn("project defect", doc)
        self.assertIn("uploaded", doc)
        self.assertIn("not a dependency install failure", doc)


if __name__ == "__main__":
    unittest.main()
