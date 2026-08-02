#!/usr/bin/env python3

from __future__ import annotations

import json
import os
import re
import subprocess
import tempfile
import unittest
from pathlib import Path


ROOT = Path(__file__).resolve().parent.parent
PACKAGE_PATH = ROOT / "package.json"
BOOTSTRAP_PATH = ROOT / "scripts" / "ci" / "setup-pinned-npm.sh"
RELEASE_CONTRACT_RUNNER = ROOT / "test" / "run-release-contracts.py"
WORKFLOW_DIR = ROOT / ".github" / "workflows"
WORKFLOW_SETUP_COUNTS = {
    # Quality, official inventory, tooling, both performance jobs, and release
    # all invoke npm and therefore must use the pinned bootstrap first.
    "ci-harness.yml": 6,
    "ci-quality.yml": 1,
    "security-static-analysis.yml": 1,
}


class PinnedNpmBootstrapContract(unittest.TestCase):
    def package_manager_version(self) -> str:
        package = json.loads(PACKAGE_PATH.read_text(encoding="utf-8"))
        package_manager = package["packageManager"]
        self.assertRegex(package_manager, r"^npm@\d+\.\d+\.\d+$")
        return package_manager.removeprefix("npm@")

    def run_with_fake_npm(self, reported_version: str) -> tuple[subprocess.CompletedProcess[str], str]:
        temp_dir = tempfile.TemporaryDirectory()
        self.addCleanup(temp_dir.cleanup)
        bin_dir = Path(temp_dir.name)
        invocation_log = bin_dir / "npm-invocations.txt"
        fake_npm = bin_dir / "npm"
        fake_npm.write_text(
            "#!/usr/bin/env bash\n"
            "printf '%s\\n' \"$*\" >>\"$FAKE_NPM_LOG\"\n"
            "if [[ \"${1:-}\" == \"--version\" ]]; then\n"
            "  printf '%s\\n' \"$FAKE_NPM_VERSION\"\n"
            "fi\n",
            encoding="utf-8",
        )
        fake_npm.chmod(0o755)
        env = os.environ.copy()
        env.update(
            {
                "FAKE_NPM_LOG": str(invocation_log),
                "FAKE_NPM_VERSION": reported_version,
                "PATH": f"{bin_dir}{os.pathsep}{env['PATH']}",
            }
        )
        proc = subprocess.run(
            ["bash", str(BOOTSTRAP_PATH)],
            cwd=ROOT,
            env=env,
            text=True,
            capture_output=True,
        )
        invocations = (
            invocation_log.read_text(encoding="utf-8")
            if invocation_log.exists()
            else ""
        )
        return proc, invocations

    def test_bootstrap_installs_and_verifies_the_declared_exact_version(self) -> None:
        expected_version = self.package_manager_version()
        proc, invocations = self.run_with_fake_npm(expected_version)

        self.assertEqual(0, proc.returncode, proc.stdout + proc.stderr)
        self.assertIn(
            "install --global --ignore-scripts --no-audit --no-fund "
            f"npm@{expected_version}",
            invocations,
        )
        self.assertIn("--version", invocations)
        self.assertIn(f"npm {expected_version}", proc.stdout)

    def test_bootstrap_fails_when_the_executable_version_does_not_match(self) -> None:
        expected_version = self.package_manager_version()
        proc, _ = self.run_with_fake_npm("0.0.0")

        self.assertNotEqual(0, proc.returncode)
        self.assertIn(
            f"expected npm {expected_version}, got 0.0.0",
            proc.stderr,
        )

    def test_every_ci_npm_job_bootstraps_before_invoking_npm(self) -> None:
        for workflow_name, expected_count in WORKFLOW_SETUP_COUNTS.items():
            workflow = (WORKFLOW_DIR / workflow_name).read_text(encoding="utf-8")
            lines = workflow.splitlines()
            setup_count = 0
            bootstrap_count = 0
            pinned_ready = False

            for line_number, line in enumerate(lines, 1):
                if re.fullmatch(r"^  [A-Za-z0-9_-]+:\s*$", line):
                    pinned_ready = False
                if "uses: actions/setup-node@" in line:
                    setup_count += 1
                    pinned_ready = False
                if "run: bash scripts/ci/setup-pinned-npm.sh" in line:
                    bootstrap_count += 1
                    pinned_ready = True
                if "run: npm " in line:
                    self.assertTrue(
                        pinned_ready,
                        f"{workflow_name}:{line_number} invokes npm before the "
                        "declared version is active",
                    )

            self.assertEqual(expected_count, setup_count, workflow_name)
            self.assertEqual(setup_count, bootstrap_count, workflow_name)

    def test_contract_is_wired_into_the_release_gate(self) -> None:
        runner = RELEASE_CONTRACT_RUNNER.read_text(encoding="utf-8")
        self.assertIn(
            '["python3", "test/test_pinned_npm_bootstrap.py"]',
            runner,
        )


if __name__ == "__main__":
    unittest.main()
