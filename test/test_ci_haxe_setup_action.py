#!/usr/bin/env python3

from __future__ import annotations

import unittest
from pathlib import Path


REPO_ROOT = Path(__file__).resolve().parent.parent
ACTION_PATH = REPO_ROOT / ".github" / "actions" / "setup-haxe-linux" / "action.yml"
WORKFLOW_PATHS = [
    REPO_ROOT / ".github" / "workflows" / "ci-quality.yml",
    REPO_ROOT / ".github" / "workflows" / "ci-harness.yml",
]
HAXE_ACTION_REF = (
    "uses: krdlab/setup-haxe@d93667502be3b4f31a94a3308a74388f2e178a8d"
    " # v2.1.0"
)


class CiHaxeSetupActionTest(unittest.TestCase):
    def test_linux_haxe_setup_is_centralized_in_composite_action(self) -> None:
        action = ACTION_PATH.read_text(encoding="utf-8")
        self.assertIn(HAXE_ACTION_REF, action)
        self.assertIn("bash scripts/ci/setup-haxe-linux-fallback.sh", action)
        self.assertIn("warmup_attempts=", action)
        self.assertIn("SETUP_HAXE_OUTCOME", action)

        for path in WORKFLOW_PATHS:
            workflow = path.read_text(encoding="utf-8")
            self.assertIn("uses: ./.github/actions/setup-haxe-linux", workflow)
            self.assertNotIn("bash scripts/ci/setup-haxe-linux-fallback.sh", workflow)
            self.assertNotIn("id: setup_haxe_linux", workflow)

        harness = WORKFLOW_PATHS[1].read_text(encoding="utf-8")
        self.assertNotIn("uses: krdlab/setup-haxe@", harness)

    def test_ci_harness_uses_shared_haxe_setup_for_each_haxe_job(self) -> None:
        workflow = (REPO_ROOT / ".github" / "workflows" / "ci-harness.yml").read_text(encoding="utf-8")
        # Quality, official inventory, both performance jobs, and release all
        # compile Haxe on Linux and must share the hardened setup action.
        self.assertEqual(workflow.count("uses: ./.github/actions/setup-haxe-linux"), 5)

    def test_ci_quality_keeps_macos_setup_separate(self) -> None:
        workflow = (REPO_ROOT / ".github" / "workflows" / "ci-quality.yml").read_text(encoding="utf-8")
        self.assertEqual(workflow.count("uses: ./.github/actions/setup-haxe-linux"), 1)
        self.assertIn("Setup Haxe (macOS)", workflow)
        self.assertIn(HAXE_ACTION_REF, workflow)
        self.assertIn("haxe-version: ${{ env.HAXE_VERSION }}", workflow)
        self.assertNotIn("brew install haxe neko", workflow)


if __name__ == "__main__":
    unittest.main()
