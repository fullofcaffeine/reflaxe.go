#!/usr/bin/env python3

from __future__ import annotations

import json
import unittest
from pathlib import Path


ROOT = Path(__file__).resolve().parent.parent
SOURCE_ROOT = ROOT / "src" / "reflaxe" / "go"
DOC = ROOT / "docs" / "haxe-4.3.7-modernization.md"
RELEASE_RUNNER = ROOT / "test" / "run-release-contracts.py"


class HaxeMacroLifecycleContractTest(unittest.TestCase):
    def test_macros_use_supported_phase_hooks(self) -> None:
        guard = (
            SOURCE_ROOT / "compiler" / "SiblingTargetConflictGuard.hx"
        ).read_text(encoding="utf-8")
        self.assertIn("Context.onAfterInitMacros(validate);", guard)

        for relative in [
            "macros/BoundaryEnforcer.hx",
            "macros/NativeAuthorityGate.hx",
            "macros/NativeBoundaryEnforcer.hx",
            "macros/StrictModeEnforcer.hx",
        ]:
            source = (SOURCE_ROOT / relative).read_text(encoding="utf-8")
            self.assertIn("Context.onAfterTyping", source, relative)

        combined = "\n".join(
            path.read_text(encoding="utf-8")
            for path in sorted(SOURCE_ROOT.rglob("*.hx"))
        )
        self.assertNotIn("onMacroContextReused", combined)
        self.assertNotIn("@:persistent static var initialized", combined)
        self.assertNotIn("@:persistent static var bootstrapped", combined)

    def test_lifecycle_policy_is_documented(self) -> None:
        source = DOC.read_text(encoding="utf-8")
        self.assertIn("## Macro lifecycle and compilation-server safety", source)
        self.assertIn("onAfterInitMacros", source)
        self.assertIn("onAfterTyping", source)
        self.assertIn("onMacroContextReused", source)
        self.assertIn("@:persistent", source)

    def test_lifecycle_gate_is_release_blocking(self) -> None:
        scripts = json.loads((ROOT / "package.json").read_text(encoding="utf-8"))["scripts"]
        command = "python3 test/test_haxe_macro_lifecycle.py"
        self.assertEqual(command, scripts.get("test:haxe-macro-lifecycle"))
        self.assertIn("npm run test:haxe-macro-lifecycle", scripts["test"])
        self.assertIn("npm run test:haxe-macro-lifecycle", scripts["test:changed"])
        self.assertIn(
            "test/test_haxe_macro_lifecycle.py",
            RELEASE_RUNNER.read_text(encoding="utf-8"),
        )


if __name__ == "__main__":
    unittest.main()
