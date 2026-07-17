#!/usr/bin/env python3

from __future__ import annotations

import unittest
from pathlib import Path


ROOT = Path(__file__).resolve().parents[1]
GO_COMPILER = ROOT / "src" / "reflaxe" / "go" / "GoCompiler.hx"
XML_OVERRIDE = ROOT / "std" / "go" / "_std" / "Xml.hx"
SNAPSHOT_MAIN = (
    ROOT / "test" / "snapshot" / "core" / "inline_throw_accessor_result_type" / "intended" / "main.go"
)


def function_block(source: str, name: str, next_name: str) -> str:
    start = source.index(f"function {name}")
    end = source.index(f"function {next_name}", start)
    return source[start:end]


def go_function_block(source: str, name: str, next_name: str) -> str:
    start = source.index(f"func {name}")
    end = source.index(f"func {next_name}", start)
    return source[start:end]


class InlineThrowAccessorResultTypeContract(unittest.TestCase):
    def test_throw_lowering_distinguishes_terminal_values_from_guard_continuations(self) -> None:
        compiler = GO_COMPILER.read_text(encoding="utf-8")

        typed_throw = function_block(compiler, "lowerThrowExprForType", "lowerExpectedThrowExpr")
        self.assertIn("valueStorageGoType(resultType)", typed_throw)
        self.assertIn("GoVarDecl(zeroName, resultTypeName", typed_throw)

        expected_throw = function_block(compiler, "lowerExpectedThrowExpr", "currentFunctionVarNameScope")
        self.assertIn("lowerThrowExprForType(value, resultType)", expected_throw)
        self.assertIn("case TMeta(_, inner) | TParenthesis(inner) | TCast(inner, _)", expected_throw)
        self.assertIn("case TBlock(exprs) if (exprs.length > 0)", expected_throw)
        self.assertIn("withoutThrowFallback", expected_throw)

        fallback = function_block(compiler, "nonVoidThrowFallbackReturnStmts", "lowerThrowExprForType")
        self.assertIn("throwFallbackSuppressionDepth > 0", fallback)
        self.assertIn("currentFunctionReturnType()", fallback)

        statement_block = function_block(compiler, "lowerBlock", "restoreBlockNonNullPrimitiveFacts")
        self.assertIn("index < exprs.length - 1", statement_block)
        self.assertIn("withoutThrowFallback", statement_block)

        expression_block = function_block(compiler, "lowerExprWithPrefix", "lowerTypeExpr")
        self.assertIn("withoutThrowFallback", expression_block)

        function_scope = function_block(compiler, "pushFunctionReturnType", "currentFunctionReturnType")
        self.assertIn("throwFallbackSuppressionDepthScopes", function_scope)
        self.assertIn("throwFallbackSuppressionDepth = 0", function_scope)
        self.assertIn("throwFallbackSuppressionDepthScopes.pop()", function_scope)

        for block in [typed_throw, expected_throw, fallback, statement_block, expression_block, function_scope]:
            for forbidden in ["Xml", "GoStmt.GoRaw", "reflect", "unsafe"]:
                self.assertNotIn(forbidden, block)

    def test_generated_fallbacks_use_immediate_types(self) -> None:
        main_go = SNAPSHOT_MAIN.read_text(encoding="utf-8")

        bool_context = go_function_block(main_go, "boolContext", "capture")
        generic_value = go_function_block(main_go, "genericValue", "intContext")
        int_context = go_function_block(main_go, "intContext", "main")
        nullable_context = go_function_block(main_go, "nullableContext", "stringContext")
        string_context = main_go[main_go.index("func stringContext") : main_go.index("type I_")]

        self.assertIn("var hx_throw_zero_", bool_context)
        self.assertIn(" bool", bool_context)
        self.assertIn("var hx_throw_zero_", generic_value)
        self.assertIn(" any", generic_value)
        for guard_context in [int_context, nullable_context, string_context]:
            self.assertNotIn(
                "hx_throw_zero",
                guard_context,
                "a guard throw with a later value tail must not return from the generated IIFE",
            )

    def test_xml_restores_upstream_inline_accessors(self) -> None:
        xml_source = XML_OVERRIDE.read_text(encoding="utf-8")
        for accessor in ["get_nodeName", "set_nodeName", "get_nodeValue", "set_nodeValue"]:
            self.assertIn(f"#if !cppia inline #end function {accessor}", xml_source)
        self.assertNotIn("non-inline throwing accessors", xml_source)


if __name__ == "__main__":
    unittest.main()
