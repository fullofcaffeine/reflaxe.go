#!/usr/bin/env python3

from __future__ import annotations

import json
import os
import subprocess
import sys
import tempfile
import unittest
from pathlib import Path


REPO_ROOT = Path(__file__).resolve().parent.parent
RUNNER = REPO_ROOT / "scripts" / "security" / "run-go-tooling-gates.py"
EXPECTED_TARGETS = {
    "hxrt",
    "sys-thread-runtime",
    "go-channel-runtime",
    "pulseforge-portable",
    "pulseforge-metal",
    "fluxproxy-portable",
    "fluxproxy-metal",
}
EXPECTED_GATES = {"race", "checkptr", "vet", "staticcheck"}


class GoToolingGateSemanticsTest(unittest.TestCase):
    def run_with_fake_tools(
        self,
        *,
        go_fail_match: str = "",
        staticcheck_fail: bool = False,
        go_sleep_match: str = "",
        go_sleep_target: str = "",
        go_sleep_seconds: int = 5,
        timeout_seconds: int = 10,
    ) -> tuple[subprocess.CompletedProcess[str], Path, Path, Path]:
        temp_dir = tempfile.TemporaryDirectory()
        self.addCleanup(temp_dir.cleanup)
        root = Path(temp_dir.name)
        bin_dir = root / "bin"
        report_dir = root / "reports"
        go_log = root / "go.log"
        staticcheck_log = root / "staticcheck.log"
        bin_dir.mkdir()

        fake_go = bin_dir / "go"
        fake_go.write_text(
            """#!/usr/bin/env bash
set -euo pipefail
printf 'args=%s\\n' "$*" >>"$FAKE_GO_LOG"
if [[ "$1" == "version" ]]; then
  printf 'go version go1.26.6 linux/amd64\\n'
  exit 0
fi
if [[ -n "${FAKE_GO_SLEEP_MATCH:-}" && "$*" == *"$FAKE_GO_SLEEP_MATCH"* && ( -z "${FAKE_GO_SLEEP_TARGET:-}" || "${PWD##*/}" == "$FAKE_GO_SLEEP_TARGET" ) ]]; then
  exec sleep "$FAKE_GO_SLEEP_SECONDS"
fi
if [[ -n "${FAKE_GO_FAIL_MATCH:-}" && "$*" == *"$FAKE_GO_FAIL_MATCH"* ]]; then
  printf 'synthetic go tooling failure\\n' >&2
  exit 7
fi
printf 'synthetic go tooling success\\n'
""",
            encoding="utf-8",
        )
        fake_go.chmod(0o755)

        fake_staticcheck = bin_dir / "staticcheck"
        fake_staticcheck.write_text(
            """#!/usr/bin/env bash
set -euo pipefail
printf 'args=%s\\n' "$*" >>"$FAKE_STATICCHECK_LOG"
if [[ "$1" == "-version" ]]; then
  printf 'staticcheck 2026.1 (v0.7.0)\\n'
  exit 0
fi
if [[ "${FAKE_STATICCHECK_FAIL:-0}" == "1" ]]; then
  printf 'synthetic staticcheck finding (SA9999)\\n' >&2
  exit 1
fi
printf 'synthetic staticcheck success\\n'
""",
            encoding="utf-8",
        )
        fake_staticcheck.chmod(0o755)

        env = os.environ.copy()
        env.update(
            {
                "GO_BIN": str(fake_go),
                "STATICCHECK_BIN": str(fake_staticcheck),
                "GO_TOOLING_REPORT_DIR": str(report_dir),
                "GO_TOOLING_COMMAND_TIMEOUT_SECONDS": str(timeout_seconds),
                "FAKE_GO_LOG": str(go_log),
                "FAKE_STATICCHECK_LOG": str(staticcheck_log),
                "FAKE_GO_FAIL_MATCH": go_fail_match,
                "FAKE_GO_SLEEP_MATCH": go_sleep_match,
                "FAKE_GO_SLEEP_TARGET": go_sleep_target,
                "FAKE_GO_SLEEP_SECONDS": str(go_sleep_seconds),
                "FAKE_STATICCHECK_FAIL": "1" if staticcheck_fail else "0",
            }
        )
        proc = subprocess.run(
            ["python3", str(RUNNER)],
            cwd=REPO_ROOT,
            env=env,
            capture_output=True,
            text=True,
            timeout=30,
        )
        return proc, report_dir, go_log, staticcheck_log

    def test_clean_run_executes_every_gate_once_and_records_evidence(self) -> None:
        proc, report_dir, go_log, staticcheck_log = self.run_with_fake_tools()

        self.assertEqual(proc.returncode, 0, proc.stdout + proc.stderr)
        go_calls = go_log.read_text(encoding="utf-8")
        staticcheck_calls = staticcheck_log.read_text(encoding="utf-8")
        target_count = len(EXPECTED_TARGETS)
        self.assertEqual(go_calls.count("test -race"), target_count)
        self.assertEqual(go_calls.count("test -gcflags=all=-d=checkptr=2"), target_count)
        self.assertEqual(go_calls.count("vet -stdmethods=false"), target_count)
        self.assertEqual(staticcheck_calls.count("-checks=SA*"), target_count)

        manifest = json.loads((report_dir / "manifest.json").read_text(encoding="utf-8"))
        self.assertEqual(manifest["result"], "pass")
        expected_runs = {
            (target, gate)
            for target in EXPECTED_TARGETS
            for gate in EXPECTED_GATES
        }
        self.assertEqual(len(manifest["runs"]), len(expected_runs))
        self.assertEqual(
            {(run["target"], run["gate"]) for run in manifest["runs"]},
            expected_runs,
        )
        self.assertTrue(all(run["result"] == "pass" for run in manifest["runs"]))
        self.assertIn("PASS", (report_dir / "summary.md").read_text(encoding="utf-8"))
        self.assertEqual(len(list((report_dir / "reports").glob("*.txt"))), len(expected_runs))

    def test_race_failure_is_not_retried_or_downgraded(self) -> None:
        proc, report_dir, go_log, _ = self.run_with_fake_tools(go_fail_match="-race")

        self.assertEqual(proc.returncode, 1, proc.stdout + proc.stderr)
        self.assertNotIn("all gates passed", proc.stdout)
        self.assertEqual(
            go_log.read_text(encoding="utf-8").count("test -race"),
            len(EXPECTED_TARGETS),
        )
        manifest = json.loads((report_dir / "manifest.json").read_text(encoding="utf-8"))
        failures = [run for run in manifest["runs"] if run["result"] == "fail"]
        self.assertEqual(len(failures), len(EXPECTED_TARGETS))
        self.assertEqual({run["target"] for run in failures}, EXPECTED_TARGETS)
        self.assertTrue(all(run["gate"] == "race" for run in failures))
        self.assertIn(
            "synthetic go tooling failure",
            (report_dir / "reports" / "race-hxrt.txt").read_text(encoding="utf-8"),
        )

    def test_staticcheck_finding_blocks_the_gate(self) -> None:
        proc, report_dir, _, staticcheck_log = self.run_with_fake_tools(staticcheck_fail=True)

        self.assertEqual(proc.returncode, 1, proc.stdout + proc.stderr)
        self.assertEqual(
            staticcheck_log.read_text(encoding="utf-8").count("-checks=SA*"),
            len(EXPECTED_TARGETS),
        )
        manifest = json.loads((report_dir / "manifest.json").read_text(encoding="utf-8"))
        failures = [run for run in manifest["runs"] if run["result"] == "fail"]
        self.assertEqual(len(failures), len(EXPECTED_TARGETS))
        self.assertEqual({run["target"] for run in failures}, EXPECTED_TARGETS)
        self.assertTrue(all(run["gate"] == "staticcheck" for run in failures))

    def test_timeout_fails_closed_without_retry(self) -> None:
        proc, report_dir, go_log, _ = self.run_with_fake_tools(
            go_sleep_match="-race",
            go_sleep_target="hxrt",
            go_sleep_seconds=10,
            # Give process startup a generous allowance, then prove one real
            # gate command is stopped well before its ten-second fake workload.
            timeout_seconds=5,
        )

        self.assertEqual(proc.returncode, 1, proc.stdout + proc.stderr)
        self.assertEqual(
            go_log.read_text(encoding="utf-8").count("test -race"),
            len(EXPECTED_TARGETS),
        )
        manifest = json.loads((report_dir / "manifest.json").read_text(encoding="utf-8"))
        timeouts = [run for run in manifest["runs"] if run["result"] == "timeout"]
        self.assertEqual(len(timeouts), 1)
        self.assertEqual(timeouts[0]["target"], "hxrt")
        self.assertEqual(timeouts[0]["gate"], "race")

    def test_report_output_rejects_repository_root(self) -> None:
        env = os.environ.copy()
        env.update(
            {
                "GO_BIN": sys.executable,
                "GO_TOOLING_REPORT_DIR": str(REPO_ROOT),
            }
        )
        proc = subprocess.run(
            ["python3", str(RUNNER)],
            cwd=REPO_ROOT,
            env=env,
            capture_output=True,
            text=True,
            timeout=10,
        )

        self.assertEqual(proc.returncode, 1, proc.stdout + proc.stderr)
        self.assertIn("must be a dedicated output directory", proc.stderr)


if __name__ == "__main__":
    unittest.main()
