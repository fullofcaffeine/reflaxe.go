#!/usr/bin/env python3

from __future__ import annotations

import unittest
from pathlib import Path


REPO_ROOT = Path(__file__).resolve().parent.parent
WORKFLOW_DIR = REPO_ROOT / ".github" / "workflows"
HAXE_ACTION = REPO_ROOT / ".github" / "actions" / "setup-haxe-linux" / "action.yml"


class GithubActionsNode24PolicyTest(unittest.TestCase):
    def test_workflows_opt_into_node24_action_runtime(self) -> None:
        for workflow in sorted(WORKFLOW_DIR.glob("*.yml")):
            with self.subTest(workflow=workflow.name):
                text = workflow.read_text(encoding="utf-8")
                self.assertIn("FORCE_JAVASCRIPT_ACTIONS_TO_NODE24: \"true\"", text)

    def test_workflows_do_not_pin_known_node20_action_majors(self) -> None:
        blocked_refs = (
            "actions/checkout@v4",
            "actions/setup-node@v4",
            "actions/setup-go@v5",
            "actions/upload-artifact@v4",
            "actions/download-artifact@v4",
            "gitleaks/gitleaks-action@v2",
            "actions/dependency-review-action@v4",
            "github/codeql-action/init@v3",
            "github/codeql-action/analyze@v3",
            "softprops/action-gh-release@v2",
        )
        for workflow in sorted(WORKFLOW_DIR.glob("*.yml")):
            text = workflow.read_text(encoding="utf-8")
            for blocked_ref in blocked_refs:
                with self.subTest(workflow=workflow.name, blocked_ref=blocked_ref):
                    self.assertNotIn(f"uses: {blocked_ref}", text)

    def test_haxe_setup_wrapper_uses_node24_ready_upstream_action(self) -> None:
        text = HAXE_ACTION.read_text(encoding="utf-8")
        self.assertIn("uses: krdlab/setup-haxe@v2", text)
        self.assertNotIn("uses: krdlab/setup-haxe@v1", text)


if __name__ == "__main__":
    unittest.main()
