#!/usr/bin/env python3

from __future__ import annotations

import os
import subprocess
import tempfile
import unittest
from pathlib import Path


REPO_ROOT = Path(__file__).resolve().parent.parent
AUDIT_SCRIPT = REPO_ROOT / "scripts" / "security" / "run-dependency-audit.sh"
CI_HARNESS = REPO_ROOT / ".github" / "workflows" / "ci-harness.yml"
SECURITY_WORKFLOW = REPO_ROOT / ".github" / "workflows" / "security-static-analysis.yml"


class DependencyAuditGateSemanticsTest(unittest.TestCase):
    def run_with_fake_govulncheck(
        self,
        *,
        output: str,
        exit_code: int,
        extra_env: dict[str, str] | None = None,
    ) -> tuple[subprocess.CompletedProcess[str], Path]:
        temp_dir = tempfile.TemporaryDirectory()
        self.addCleanup(temp_dir.cleanup)
        root = Path(temp_dir.name)
        bin_dir = root / "bin"
        report_dir = root / "reports"
        bin_dir.mkdir()

        fake_npm = bin_dir / "npm"
        fake_npm.write_text("#!/usr/bin/env bash\nexit 0\n", encoding="utf-8")
        fake_npm.chmod(0o755)

        fake_govulncheck = bin_dir / "govulncheck"
        fake_govulncheck.write_text(
            "#!/usr/bin/env bash\nprintf '%s\\n' \"$FAKE_GOVULN_OUTPUT\"\nexit \"$FAKE_GOVULN_EXIT\"\n",
            encoding="utf-8",
        )
        fake_govulncheck.chmod(0o755)

        env = os.environ.copy()
        env.update(
            {
                "CI": "true",
                "FAKE_GOVULN_OUTPUT": output,
                "FAKE_GOVULN_EXIT": str(exit_code),
                "GOVULNCHECK_BIN": str(fake_govulncheck),
                "GOVULNCHECK_REPORT_DIR": str(report_dir),
                "PATH": f"{bin_dir}{os.pathsep}{env['PATH']}",
            }
        )
        env.update(extra_env or {})
        proc = subprocess.run(
            ["bash", str(AUDIT_SCRIPT)],
            cwd=REPO_ROOT,
            env=env,
            capture_output=True,
            text=True,
            timeout=30,
        )
        return proc, report_dir

    def test_zero_finding_scan_passes_and_records_clean_report(self) -> None:
        proc, report_dir = self.run_with_fake_govulncheck(
            output="No vulnerabilities found.",
            exit_code=0,
        )

        self.assertEqual(proc.returncode, 0, proc.stdout + proc.stderr)
        self.assertIn("No vulnerabilities found.", (report_dir / "govulncheck.txt").read_text(encoding="utf-8"))
        metadata = (report_dir / "metadata.txt").read_text(encoding="utf-8")
        self.assertIn("result=clean", metadata)
        self.assertIn("govulncheck_exit=0", metadata)
        self.assertIn("govulncheck_version=v1.6.0", metadata)

    def test_reachable_finding_fails_ci_and_preserves_trace(self) -> None:
        finding = "\n".join(
            [
                "Vulnerability #1: GO-TEST-0001",
                "Found in: stdlib@go1.test",
                "Trace: hxrt.Test calls vulnerable.Symbol",
                "Your code is affected by 1 vulnerability.",
            ]
        )
        proc, report_dir = self.run_with_fake_govulncheck(output=finding, exit_code=3)

        self.assertEqual(proc.returncode, 1, proc.stdout + proc.stderr)
        self.assertIn("reachable vulnerabilities are release-blocking", proc.stdout + proc.stderr)
        self.assertNotIn("dependency audit passed", proc.stdout)
        self.assertIn("GO-TEST-0001", (report_dir / "govulncheck.txt").read_text(encoding="utf-8"))
        metadata = (report_dir / "metadata.txt").read_text(encoding="utf-8")
        self.assertIn("result=reachable_vulnerabilities", metadata)
        self.assertIn("govulncheck_exit=3", metadata)

    def test_tool_failure_fails_ci_without_becoming_a_clean_scan(self) -> None:
        proc, report_dir = self.run_with_fake_govulncheck(
            output="govulncheck: loading packages: synthetic failure",
            exit_code=2,
        )

        self.assertEqual(proc.returncode, 2, proc.stdout + proc.stderr)
        self.assertNotIn("dependency audit passed", proc.stdout)
        metadata = (report_dir / "metadata.txt").read_text(encoding="utf-8")
        self.assertIn("result=tool_error", metadata)
        self.assertIn("govulncheck_exit=2", metadata)

    def test_ci_rejects_explicit_govulncheck_skip(self) -> None:
        proc, report_dir = self.run_with_fake_govulncheck(
            output="should not run",
            exit_code=0,
            extra_env={"SKIP_GOVULNCHECK": "1"},
        )

        self.assertEqual(proc.returncode, 1, proc.stdout + proc.stderr)
        self.assertIn("SKIP_GOVULNCHECK=1 is not permitted in CI", proc.stderr)
        metadata = (report_dir / "metadata.txt").read_text(encoding="utf-8")
        self.assertIn("result=ci_skip_rejected", metadata)

    def test_ci_uploads_dependency_audit_reports_even_on_failure(self) -> None:
        for workflow_path in (CI_HARNESS, SECURITY_WORKFLOW):
            workflow = workflow_path.read_text(encoding="utf-8")
            with self.subTest(workflow=workflow_path.name):
                self.assertIn('go: ["1.25.x", "1.26.x"]', workflow)
                self.assertIn("go-version: ${{ matrix.go }}", workflow)
                self.assertIn("Upload dependency audit reports", workflow)
                self.assertIn("if: always()", workflow)
                self.assertIn("path: .cache/security", workflow)
                self.assertIn("if-no-files-found: warn", workflow)
                self.assertIn("name: dependency-audit-${{ matrix.go }}", workflow)


if __name__ == "__main__":
    unittest.main()
