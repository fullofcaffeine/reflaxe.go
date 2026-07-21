#!/usr/bin/env python3

from __future__ import annotations

from collections import Counter
import json
import re
import unittest
from pathlib import Path


ROOT = Path(__file__).resolve().parent.parent
POLICY = ROOT / "test" / "haxe_dynamic_boundary_policy.json"
REFLAXE_COMPILER = ROOT / "src" / "reflaxe" / "go" / "GoReflaxeCompiler.hx"
STDLIB_PLANNER = (
    ROOT / "src" / "reflaxe" / "go" / "compiler" / "GoSourceOwnedStdlibPlanner.hx"
)
RELEASE_RUNNER = ROOT / "test" / "run-release-contracts.py"
DYNAMIC_TYPE_SITE = re.compile(r"(?::|<|,)\s*Dynamic\b")
BLOCK_COMMENT = re.compile(r"/\*.*?\*/", re.DOTALL)
LINE_COMMENT = re.compile(r"//.*?$", re.MULTILINE)


class HaxeDynamicBoundaryContractTest(unittest.TestCase):
    def test_legacy_reflaxe_hooks_use_a_typed_no_output_marker(self) -> None:
        source = REFLAXE_COMPILER.read_text(encoding="utf-8")
        self.assertIn("enum GoReflaxeStagedOutput", source)
        self.assertRegex(
            source,
            r"GenericCompiler<\s*GoReflaxeStagedOutput,\s*"
            r"GoReflaxeStagedOutput,\s*GoReflaxeStagedOutput,\s*"
            r"GoReflaxeStagedOutput,\s*GoReflaxeStagedOutput\s*>",
        )
        self.assertIn(
            "compileExpressionImpl(expr:TypedExpr, "
            "topLevel:Bool):Null<GoReflaxeStagedOutput>",
            source,
        )

    def test_known_string_throwing_macro_lookups_do_not_catch_dynamic(self) -> None:
        source = STDLIB_PLANNER.read_text(encoding="utf-8")
        self.assertNotIn("catch (_:Dynamic)", source)
        self.assertEqual(3, source.count("catch (_:String)"))

    def test_remaining_compiler_dynamic_types_match_reviewed_boundaries(self) -> None:
        policy = json.loads(POLICY.read_text(encoding="utf-8"))
        actual: Counter[str] = Counter()
        source_root = ROOT / "src" / "reflaxe" / "go"
        for path in sorted(source_root.rglob("*.hx")):
            source = path.read_text(encoding="utf-8")
            code_only = LINE_COMMENT.sub("", BLOCK_COMMENT.sub("", source))
            count = len(DYNAMIC_TYPE_SITE.findall(code_only))
            if count:
                actual[path.relative_to(ROOT).as_posix()] = count
        self.assertEqual(policy["allowedCompilerDynamicTypeSites"], dict(actual))

    def test_dynamic_boundary_gate_is_release_blocking(self) -> None:
        scripts = json.loads((ROOT / "package.json").read_text(encoding="utf-8"))["scripts"]
        command = "python3 test/test_haxe_dynamic_boundaries.py"
        self.assertEqual(command, scripts.get("test:haxe-dynamic-boundaries"))
        self.assertIn("npm run test:haxe-dynamic-boundaries", scripts["test"])
        self.assertIn("npm run test:haxe-dynamic-boundaries", scripts["test:changed"])
        self.assertIn(
            "test/test_haxe_dynamic_boundaries.py",
            RELEASE_RUNNER.read_text(encoding="utf-8"),
        )


if __name__ == "__main__":
    unittest.main()
