#!/usr/bin/env python3

from __future__ import annotations

import unittest
from pathlib import Path


REPO_ROOT = Path(__file__).resolve().parent.parent


class ExampleCollectionPurityTest(unittest.TestCase):
    def test_pulse_pipeline_has_no_haxe_ds_imports(self) -> None:
        path = REPO_ROOT / "examples" / "pulseforge" / "app" / "core" / "PulsePipeline.hx"
        text = path.read_text(encoding="utf-8")
        self.assertNotIn("import haxe.ds.", text, "PulsePipeline should avoid haxe.ds.* imports in metal-purity migration slice")

    def test_flux_pipeline_has_no_haxe_ds_imports(self) -> None:
        path = REPO_ROOT / "examples" / "fluxproxy" / "app" / "core" / "FluxPipeline.hx"
        text = path.read_text(encoding="utf-8")
        self.assertNotIn("import haxe.ds.", text, "FluxPipeline should avoid haxe.ds.* imports in metal-purity migration slice")

    def test_storyboard_harness_has_no_haxe_ds_imports(self) -> None:
        path = REPO_ROOT / "examples" / "profile_storyboard" / "Harness.hx"
        text = path.read_text(encoding="utf-8")
        self.assertNotIn("import haxe.ds.", text, "profile_storyboard Harness should avoid haxe.ds.* imports in metal-purity migration slice")

    def test_storyboard_storycard_has_no_haxe_ds_imports(self) -> None:
        path = REPO_ROOT / "examples" / "profile_storyboard" / "domain" / "StoryCard.hx"
        text = path.read_text(encoding="utf-8")
        self.assertNotIn("import haxe.ds.", text, "profile_storyboard StoryCard should avoid haxe.ds.* imports in metal-purity migration slice")

    def test_tui_todo_item_has_no_haxe_ds_imports(self) -> None:
        path = REPO_ROOT / "examples" / "tui_todo" / "model" / "TodoItem.hx"
        text = path.read_text(encoding="utf-8")
        self.assertNotIn("import haxe.ds.", text, "tui_todo TodoItem should avoid haxe.ds.* imports in metal-purity migration slice")

    def test_tui_todo_store_has_no_haxe_ds_imports(self) -> None:
        path = REPO_ROOT / "examples" / "tui_todo" / "model" / "TodoStore.hx"
        text = path.read_text(encoding="utf-8")
        self.assertNotIn("import haxe.ds.", text, "tui_todo TodoStore should avoid haxe.ds.* imports in metal-purity migration slice")

    def test_tui_todo_app_has_no_haxe_ds_imports(self) -> None:
        path = REPO_ROOT / "examples" / "tui_todo" / "app" / "TodoApp.hx"
        text = path.read_text(encoding="utf-8")
        self.assertNotIn("import haxe.ds.", text, "tui_todo TodoApp should avoid haxe.ds.* imports in metal-purity migration slice")

    def test_tui_todo_harness_has_no_haxe_ds_imports(self) -> None:
        path = REPO_ROOT / "examples" / "tui_todo" / "Harness.hx"
        text = path.read_text(encoding="utf-8")
        self.assertNotIn("import haxe.ds.", text, "tui_todo Harness should avoid haxe.ds.* imports in metal-purity migration slice")

    def test_tui_todo_interactive_cli_has_no_haxe_ds_imports(self) -> None:
        path = REPO_ROOT / "examples" / "tui_todo" / "InteractiveCli.hx"
        text = path.read_text(encoding="utf-8")
        self.assertNotIn("import haxe.ds.", text, "tui_todo InteractiveCli should avoid haxe.ds.* imports in metal-purity migration slice")


if __name__ == "__main__":
    unittest.main()
