#!/usr/bin/env python3

from __future__ import annotations

import unittest
from pathlib import Path


ROOT = Path(__file__).resolve().parents[1]
GO_COMPILER = ROOT / "src" / "reflaxe" / "go" / "GoCompiler.hx"
GO_TYPE_MAPPER = ROOT / "src" / "reflaxe" / "go" / "compiler" / "GoTypeMapper.hx"
HXRT_ANALYZER = (
    ROOT / "src" / "reflaxe" / "go" / "compiler" / "GoHxrtFeatureAnalyzer.hx"
)
HXRT_EQUALITY = ROOT / "runtime" / "hxrt" / "equality.go"
XML_OVERRIDE = ROOT / "std" / "go" / "_std" / "Xml.hx"
SNAPSHOT_MAIN = (
    ROOT / "test" / "snapshot" / "core" / "array_remove_insert" / "intended" / "main.go"
)


def function_block(source: str, name: str, next_name: str) -> str:
    start = source.index(f"function {name}")
    end = source.index(f"function {next_name}", start)
    return source[start:end]


def go_function_block(source: str, name: str, next_name: str) -> str:
    start = source.index(f"func {name}")
    end = source.index(f"func {next_name}", start)
    return source[start:end]


class ArrayRemoveInsertLoweringContract(unittest.TestCase):
    def test_array_mutations_use_structured_general_lowering(self) -> None:
        compiler = GO_COMPILER.read_text(encoding="utf-8")

        equality = function_block(
            compiler, "lowerArrayElementEqualityExpr", "lowerArrayRemoveExpr"
        )
        self.assertIn("hxrt.StringEqualStringPtr", equality)
        self.assertIn("hxrt.HaxeEqual", equality)

        remove = function_block(compiler, "lowerArrayRemoveExpr", "lowerArrayInsertExpr")
        self.assertIn("lowerArrayMutationSite", remove)
        self.assertIn("GoStmt.GoWhile", remove)
        self.assertIn('GoExpr.GoSelector(site.tempExpr, "Len")', remove)
        self.assertIn('GoExpr.GoSelector(site.tempExpr, "Get")', remove)
        self.assertIn('GoExpr.GoSelector(site.tempExpr, "RemoveAt")', remove)
        self.assertNotIn("site.writeBack", remove)

        insert = function_block(compiler, "lowerArrayInsertExpr", "lowerArrayInstanceCall")
        self.assertIn("lowerArrayMutationSite", insert)
        self.assertIn('GoExpr.GoSelector(site.tempExpr, "Insert")', insert)
        self.assertNotIn("site.writeBack", insert)

        for block in [equality, remove, insert]:
            for forbidden in ["Xml", "GoStmt.GoRaw", "reflect", "unsafe"]:
                self.assertNotIn(forbidden, block)

    def test_root_arrays_use_the_shared_carrier_while_native_slices_stay_typed(
        self,
    ) -> None:
        mapper = GO_TYPE_MAPPER.read_text(encoding="utf-8")
        storage = function_block(mapper, "arrayElementStorageGoType", "typeToGoType")
        self.assertIn("isNullablePrimitiveType(elementType)", storage)
        self.assertIn('"any"', storage)

        type_mapping = function_block(mapper, "typeToGoType", "isStringType")
        self.assertIn("nativeSliceElementType(type)", type_mapping)
        self.assertIn('return "[]" + arrayElementStorageGoType', type_mapping)
        self.assertIn('classType.name == "Array"', type_mapping)
        self.assertIn('"*hxrt.Array"', type_mapping)

        array_element = function_block(
            mapper, "arrayElementGoType", "scalarGoType"
        )
        self.assertIn("arrayElementStorageGoType(params[0]", array_element)

    def test_generic_equality_runtime_is_reference_safe(self) -> None:
        runtime = HXRT_EQUALITY.read_text(encoding="utf-8")
        start = runtime.index("func HaxeEqual")
        equality = runtime[start:]
        self.assertIn("haxeNumericValue", equality)
        self.assertIn("referenceEqual", equality)
        self.assertNotIn("DeepEqual", equality)

    def test_erased_equality_is_a_selective_runtime_feature(self) -> None:
        analyzer = HXRT_ANALYZER.read_text(encoding="utf-8")
        self.assertIn('var HxrtEquality = "equality";', analyzer)
        self.assertIn(
            "FEATURE_EQUALITY:GoHxrtFeatureId = GoHxrtFeatureId.HxrtEquality",
            analyzer,
        )
        self.assertIn('add(FEATURE_EQUALITY, "compiler_surface"', analyzer)
        self.assertIn('case FEATURE_EQUALITY:\n\t\t\t\t["equality.go"]', analyzer)
        self.assertIn('case FEATURE_ATOMIC_OBJECT:\n\t\t\t\t[FEATURE_EQUALITY]', analyzer)

    def test_generated_go_mutates_the_shared_carrier_in_place(self) -> None:
        generated = SNAPSHOT_MAIN.read_text(encoding="utf-8")
        insert_generic = go_function_block(generated, "insertGeneric", "main")
        remove_generic = go_function_block(
            generated, "removeGeneric", "showNullableInts"
        )

        self.assertIn("values := hxrt.NewArray", insert_generic)
        self.assertIn("values.Insert(", insert_generic)
        self.assertIn("values.Len()", insert_generic)
        self.assertIn("hxrt.HaxeEqual", remove_generic)
        self.assertIn("values.RemoveAt(", remove_generic)
        self.assertIn("nullableInts := hxrt.NewArray(nil, 1, nil)", generated)
        self.assertIn('["values"] = hxrt.NewArray', generated)
        self.assertNotIn('["values"] = hx_arr_', generated)
        self.assertNotIn("append(values", insert_generic)
        self.assertNotIn("copy(values", insert_generic)
        self.assertNotIn(".remove(", generated)
        self.assertNotIn(".insert(", generated)

    def test_xml_uses_upstream_array_mutations(self) -> None:
        xml = XML_OVERRIDE.read_text(encoding="utf-8")
        self.assertIn("if (children.remove(x))", xml)
        self.assertIn("x.parent.children.remove(x);", xml)
        self.assertIn("children.insert(pos, x);", xml)
        self.assertNotIn("var remaining = new Array<Xml>();", xml)
        self.assertNotIn("var inserted = new Array<Xml>();", xml)


if __name__ == "__main__":
    unittest.main()
