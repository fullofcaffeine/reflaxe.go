#!/usr/bin/env python3

from __future__ import annotations

import json
import unittest
from pathlib import Path


REPO_ROOT = Path(__file__).resolve().parent.parent
RUNNER = REPO_ROOT / "scripts" / "security" / "run-go-tooling-gates.py"
RUNTIME_FIXTURE = REPO_ROOT / "test" / "go_tooling" / "hxrt_gate_test.go"
FLAGSHIP_FIXTURE = REPO_ROOT / "test" / "go_tooling" / "flagship_gate_test.go"
THREAD_POOL_FIXTURE = REPO_ROOT / "test" / "go_tooling" / "thread_pool_gate_test.go"
CHANNEL_FIXTURE = REPO_ROOT / "test" / "go_tooling" / "channel_gate_test.go"
POLICY_DOC = REPO_ROOT / "docs" / "go-tooling-gates.md"
CI_HARNESS = REPO_ROOT / ".github" / "workflows" / "ci-harness.yml"
PACKAGE_JSON = REPO_ROOT / "package.json"
RELEASE_CHECKLIST = REPO_ROOT / "docs" / "release-readiness-checklist.md"
RELEASE_STATUS = REPO_ROOT / "scripts" / "release" / "check-release-state.sh"
RELEASE_CONTRACTS = REPO_ROOT / "test" / "run-release-contracts.py"


class GoToolingGateContractTest(unittest.TestCase):
    def test_runner_pins_tools_and_declares_every_release_scope(self) -> None:
        runner = RUNNER.read_text(encoding="utf-8")

        self.assertIn('STATICCHECK_VERSION = "v0.7.0"', runner)
        self.assertIn("runtime/hxrt", runner)
        self.assertIn("test/snapshot/stdlib/sys_thread_runtime_direct/intended", runner)
        self.assertIn("test/snapshot/go_native/channel_try_recv/intended", runner)
        for app in ("pulseforge", "fluxproxy"):
            for profile in ("portable", "metal"):
                self.assertIn(f"examples/{app}/generated/{profile}", runner)
        self.assertIn('"-race"', runner)
        self.assertIn('"-gcflags=all=-d=checkptr=2"', runner)
        self.assertIn('"vet", "-stdmethods=false"', runner)
        self.assertIn('"-checks=SA*"', runner)
        self.assertIn("GO_TOOLING_COMMAND_TIMEOUT_SECONDS", runner)
        self.assertIn("manifest.json", runner)
        self.assertIn("summary.md", runner)
        self.assertNotIn("shutil.rmtree(report_dir)", runner)

    def test_runtime_fixture_exercises_concurrency_and_reflection_paths(self) -> None:
        fixture = RUNTIME_FIXTURE.read_text(encoding="utf-8")

        self.assertIn("TestConcurrentThreadMessagesAndAtomics", fixture)
        self.assertIn("ThreadSpawn", fixture)
        self.assertIn("ThreadSendMessage", fixture)
        self.assertIn("AtomicInt", fixture)
        self.assertIn("AtomicObject", fixture)
        self.assertIn("sync.WaitGroup", fixture)
        self.assertIn("unsafe.Pointer", fixture)
        self.assertIn("TestCheckptrReflectionPaths", fixture)

    def test_flagship_fixture_runs_scripted_entrypoint(self) -> None:
        fixture = FLAGSHIP_FIXTURE.read_text(encoding="utf-8")

        self.assertIn("TestScriptedEntrypoint", fixture)
        self.assertIn('"--scripted"', fixture)
        self.assertIn("main()", fixture)

    def test_generated_concurrency_fixtures_cover_failure_and_race_paths(self) -> None:
        pools = THREAD_POOL_FIXTURE.read_text(encoding="utf-8")
        self.assertIn("submissions = 10_000", pools)
        self.assertIn("[]int{1, 2, 8}", pools)
        self.assertIn("ThreadWaitForAll", pools)
        self.assertIn("accepted tasks", pools)
        self.assertIn("TestGeneratedThreadPoolsReplaceWorkerAfterHaxeThrow", pools)
        self.assertIn("expected worker failure", pools)

        channels = CHANNEL_FIXTURE.read_text(encoding="utf-8")
        self.assertIn("tryRecv after close", channels)
        self.assertIn("send after close", channels)
        self.assertIn("double close", channels)
        self.assertIn("nil channel", channels)

    def test_ci_matrix_blocks_release_and_always_uploads_reports(self) -> None:
        workflow = CI_HARNESS.read_text(encoding="utf-8")

        self.assertIn("go-tooling:", workflow)
        self.assertIn('go: ["1.25.x", "1.26.x"]', workflow)
        self.assertIn("go-version: ${{ matrix.go }}", workflow)
        self.assertIn("python3 scripts/security/run-go-tooling-gates.py", workflow)
        self.assertIn("Upload Go tooling gate reports", workflow)
        self.assertIn("if: always()", workflow)
        self.assertIn("name: go-tooling-${{ matrix.go }}", workflow)
        self.assertIn("path: .cache/security/go-tooling", workflow)
        semantic_release = workflow.split("  semantic-release:", 1)[1]
        self.assertIn("- go-tooling", semantic_release)

    def test_commands_and_release_docs_expose_the_blocking_gate(self) -> None:
        scripts = json.loads(PACKAGE_JSON.read_text(encoding="utf-8"))["scripts"]
        self.assertEqual(
            scripts["security:go-tooling"],
            "python3 scripts/security/run-go-tooling-gates.py",
        )

        policy = POLICY_DOC.read_text(encoding="utf-8")
        self.assertIn("# Go Tooling Release Gates", policy)
        self.assertIn("npm run security:go-tooling", policy)
        self.assertIn("Staticcheck 2026.1 (`v0.7.0`)", policy)
        self.assertIn("`SA*`", policy)
        self.assertIn("`stdmethods`", policy)
        self.assertIn("no test retries", policy)
        self.assertIn("30-minute", policy)
        self.assertIn("six-minute", policy)

        checklist = RELEASE_CHECKLIST.read_text(encoding="utf-8")
        self.assertIn("npm run security:go-tooling", checklist)
        self.assertIn("race detector", checklist)
        self.assertIn("checkptr", checklist)

        release_status = RELEASE_STATUS.read_text(encoding="utf-8")
        self.assertIn('grep -Fq -- "$pattern" "$path"', release_status)
        self.assertIn('require_contains ".github/workflows/ci-harness.yml" "- go-tooling"', release_status)
        self.assertIn("release-blocking Go tooling gate", release_status)

        contracts = RELEASE_CONTRACTS.read_text(encoding="utf-8")
        self.assertIn("test/test_go_tooling_gate_contract.py", contracts)
        self.assertIn("test/test_go_tooling_gate_semantics.py", contracts)


if __name__ == "__main__":
    unittest.main()
