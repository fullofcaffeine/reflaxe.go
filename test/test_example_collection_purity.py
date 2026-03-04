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


if __name__ == "__main__":
    unittest.main()
