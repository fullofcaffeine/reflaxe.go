import unittest
from pathlib import Path


ROOT = Path(__file__).resolve().parents[1]
GO_COMPILER = ROOT / "src" / "reflaxe" / "go" / "GoCompiler.hx"
LOWERING_MODULE = ROOT / "src" / "reflaxe" / "go" / "compiler" / "GoLambdaIterableLowering.hx"
SOURCE_PLANNER = ROOT / "src" / "reflaxe" / "go" / "compiler" / "GoSourceOwnedStdlibPlanner.hx"
LAMBDA_OVERRIDE = ROOT / "std" / "go" / "_std" / "Lambda.hx"
STRUCTURAL_ITERATOR_SNAPSHOT = ROOT / "test" / "snapshot" / "core" / "structural_iterator_assignment" / "intended"
INLINE_EFFECT_SNAPSHOT = ROOT / "test" / "snapshot" / "core" / "inline_structural_iterator_effect" / "intended"
INLINE_CONCRETE_EFFECT_SNAPSHOT = (
    ROOT / "test" / "snapshot" / "core" / "inline_concrete_structural_iterator_effect" / "intended"
)
CONSTRUCTOR_ITERATOR_SNAPSHOT = ROOT / "test" / "snapshot" / "core" / "structural_iterator_constructor_argument" / "intended"
XML_OVERRIDE = ROOT / "std" / "go" / "_std" / "Xml.hx"
XML_PRINTER_SNAPSHOT = (
    ROOT / "test" / "snapshot" / "stdlib" / "xml_root_dom_basic" / "intended" / "module_haxe_xml_printer.go"
)


def function_block(source: str, name: str, next_name: str) -> str:
    start = source.index(f"function {name}")
    end = source.index(f"function {next_name}", start)
    return source[start:end]


class LambdaIterableLoweringOwnershipContract(unittest.TestCase):
    def test_public_lambda_algorithms_are_staged_haxe(self):
        self.assertTrue(
            LAMBDA_OVERRIDE.exists(),
            "Lambda algorithms should be owned by the canonical staged Haxe stdlib",
        )
        source = LAMBDA_OVERRIDE.read_text(encoding="utf-8")
        for heading in ["What", "Why", "How"]:
            self.assertIn(heading, source, f"Lambda override should document {heading}")
        self.assertIn("@:dce", source, "unused staged Lambda helpers must not bloat generated Go")
        self.assertIn(
            '@:ifFeature("Lambda.flatten", "Lambda.flatMap")',
            source,
            "the private nested carrier should be retained only for the two helpers that use it",
        )
        for method in [
            "array",
            "list",
            "map",
            "mapi",
            "flatten",
            "flatMap",
            "has",
            "exists",
            "foreach",
            "iter",
            "filter",
            "fold",
            "foldi",
            "count",
            "empty",
            "indexOf",
            "find",
            "findIndex",
            "concat",
        ]:
            self.assertIn(
                f"public static function {method}",
                source,
                f"Lambda.{method} should lower from staged Haxe source",
            )

        compiler_source = GO_COMPILER.read_text(encoding="utf-8")
        self.assertNotIn(
            "function lowerLambdaStaticCall",
            compiler_source,
            "GoCompiler must not reimplement public Lambda algorithms",
        )

        self.assertIn(
            "private class LambdaGoIterableCarrier",
            source,
            "flatMap's nominal representation carrier should remain private staged source",
        )
        carrier_class_index = source.index("private class LambdaGoIterableCarrier")
        carrier_source = source[source.rfind("/**", 0, carrier_class_index) :]
        for heading in ["What", "Why", "How"]:
            self.assertIn(heading, carrier_source, f"flatMap carrier should document {heading}")
        for forbidden in ["Dynamic", "Any", "__go__", "GoInjection", "for (", "while ("]:
            self.assertNotIn(
                forbidden,
                carrier_source,
                f"flatMap carrier must stay a representation-only wrapper without {forbidden}",
            )
        self.assertIn("return source.iterator();", carrier_source)
        self.assertGreaterEqual(
            compiler_source.count('requireSourceOwnedStdlibModule("Lambda")'),
            2,
            "flatten and flatMap must retain the private companion without affecting other Lambda calls",
        )

        adapter = function_block(
            compiler_source,
            "lowerLambdaSourceCallAdapter",
            "lowerStdIsOfTypeCall",
        )
        for forbidden in ["GoStmt.GoRaw", "reflect.DeepEqual", ".items", 'GoIdent("len")']:
            self.assertNotIn(
                forbidden,
                adapter,
                f"Lambda adapter must not own algorithmic behavior through {forbidden}",
            )
        for bridge in [
            "dynamicIterableSource",
            "predicateAnyAdapter",
            "mapperAnyAdapter",
            "indexedMapperAnyAdapter",
            "iterableMapperAnyAdapter",
            "consumerAnyAdapter",
            "folderAnyAdapter",
            "indexedFolderAnyAdapter",
            "anyArrayCoerce",
            "dynamicNestedIterableSource",
        ]:
            self.assertIn(bridge, adapter)

    def test_sort_calls_keep_only_erased_generic_adaptation(self):
        compiler_source = GO_COMPILER.read_text(encoding="utf-8")
        adapter = function_block(
            compiler_source,
            "lowerDsSortHelperCall",
            "noteSourceOwnedStdlibStaticCall",
        )
        for algorithm in ["doMerge", "rotate", "upper", "lower(a", "gcd"]:
            self.assertNotIn(algorithm, adapter)
        for bridge in [
            "lowerTypedArrayToAnyCoerce",
            "lowerTypedComparatorToAny",
            "lowerAnyArrayCopyBack",
            "lowerNullableAwareTypeAssertExpr",
        ]:
            self.assertIn(bridge, adapter)

    def test_lambda_iterable_adapter_logic_has_dedicated_owner(self):
        compiler_source = GO_COMPILER.read_text(encoding="utf-8")
        self.assertTrue(LOWERING_MODULE.exists(), "Lambda/Iterable lowering policy should live outside GoCompiler.hx")
        lowering_source = LOWERING_MODULE.read_text(encoding="utf-8")
        self.assertGreaterEqual(
            lowering_source.count('GoIdent("Lambda_goIterableCarrierAdapter")'),
            2,
            "flatten elements and flatMap callback results should share the staged nominal carrier",
        )
        for symbol in [
            "function tryLambdaSourcePlan",
            "function lowerLambdaManualIteratorProtocolSource",
            "function lowerLambdaDynamicIterableSource",
            "function lowerLambdaPredicateAnyAdapter",
            "function lowerLambdaMapperAnyAdapter",
            "function lowerLambdaConsumerAnyAdapter",
            "function lowerLambdaFolderAnyAdapter",
            "function lowerLambdaAnyArrayCoerce",
        ]:
            self.assertNotIn(symbol, compiler_source)

    def test_structural_iterator_assignments_reuse_the_typed_iterable_adapter(self):
        compiler_source = GO_COMPILER.read_text(encoding="utf-8")
        lowering_source = LOWERING_MODULE.read_text(encoding="utf-8")
        planner_source = SOURCE_PLANNER.read_text(encoding="utf-8")

        upcast = function_block(compiler_source, "upcastIfNeeded", "typeToGoType")
        self.assertIn(
            "lambdaIterableLowering.structuralIteratorCoerce",
            upcast,
            "declarations, returns, assignments, and arguments should share one structural coercion hook",
        )

        adapter = function_block(lowering_source, "structuralIteratorCoerce", "trySourcePlan")
        for required in [
            'GoStringLiteral("hasNext")',
            'GoStringLiteral("next")',
            'GoSelector(GoExpr.GoIdent(sourceName), "__hx_this")',
            "lowerNullableAwareTypeAssertExpr",
        ]:
            self.assertIn(required, adapter)
        for forbidden in [
            "reflect",
            "unsafe",
            "GoStmt.GoRaw",
            "ArrayIterator",
            "MapKeyValueIterator",
            "Xml",
        ]:
            self.assertNotIn(
                forbidden,
                adapter,
                f"the structural adapter must be typed and class-agnostic, not coupled through {forbidden}",
            )
        native_array_adapter = function_block(
            lowering_source,
            "nativeArrayStructuralIteratorCoerce",
            "structuralIteratorCoerce",
        )
        self.assertIn("nativeArrayCursorMap", native_array_adapter)
        for forbidden in ["lowerTypedArrayToAnyCoerce", "[]any", "reflect", "unsafe"]:
            self.assertNotIn(
                forbidden,
                native_array_adapter,
                "the native array cursor must retain the live typed slice instead of boxing a copy",
            )

        self.assertNotIn(
            'case "haxe.iterators.ArrayIterator", "haxe.iterators.ArrayKeyValueIterator":',
            planner_source,
            "a direct structural ArrayIterator conversion should use the live typed array plan without staging an erased class",
        )

    def test_structural_iterator_generated_shape_stays_private_and_typed(self):
        main_go = (STRUCTURAL_ITERATOR_SNAPSHOT / "main.go").read_text(encoding="utf-8")
        for required in [
            'map[string]any {',
            '["hasNext"] = func() bool',
            '["next"] = func()',
            ".__hx_this.hasNext()",
            ".__hx_this.next()",
            ".(*string)",
        ]:
            self.assertIn(required, main_go)

        self.assertNotIn("New_haxe__iterators__ArrayIterator", main_go)
        self.assertNotIn("any(hx_structural_array", main_go)
        capture_index = main_go.index(" := arrayValues")
        mutation_index = main_go.index("arrayValues[0] = 8")
        consume_index = main_go.index('hxrt.StringFromLiteral("array=")')
        self.assertLess(capture_index, mutation_index)
        self.assertLess(mutation_index, consume_index)
        self.assertIn("func (self *SnapshotGenericIterator) hasNext() bool", main_go)
        self.assertIn("func (self *SnapshotGenericIterator) next() any", main_go)
        for forbidden in ["reflect.", "unsafe.", ") HasNext(", ") Next("]:
            self.assertNotIn(forbidden, main_go)

    def test_effectful_inline_iterator_prefix_is_preserved_once(self):
        compiler_source = GO_COMPILER.read_text(encoding="utf-8")
        lowering_source = LOWERING_MODULE.read_text(encoding="utf-8")
        inline_go = (INLINE_EFFECT_SNAPSHOT / "main.go").read_text(encoding="utf-8")
        xml_source = XML_OVERRIDE.read_text(encoding="utf-8")
        xml_printer_go = XML_PRINTER_SNAPSHOT.read_text(encoding="utf-8")

        cursor_planner = function_block(lowering_source, "nativeArrayCursorPlan", "nativeArrayCursorMap")
        self.assertIn("setup.concat(tailPlan.setup)", cursor_planner)
        self.assertIn("foldTrailingArrayAliases", cursor_planner)

        native_adapter = function_block(
            lowering_source,
            "nativeArrayStructuralIteratorCoerce",
            "structuralIteratorCoerce",
        )
        self.assertIn("lowerToStatements(setupExpr)", native_adapter)
        self.assertIn("prefix: prefix", native_adapter)

        upcast = function_block(compiler_source, "upcastIfNeeded", "typeToGoType")
        self.assertIn(
            "materializeExprWithPrefix(nativeArrayIterator, toType)",
            upcast,
            "expression-only call arguments must keep an inline iterator's ordered prefix",
        )

        main_go = inline_go[inline_go.index("func main()") :]
        effect = "effectCount = int(int32((effectCount + 1)))"
        self.assertEqual(main_go.count(effect), 1)
        effect_index = main_go.index(effect)
        capture_index = main_go.index(" := values", effect_index)
        mutation_index = main_go.index("values[0] = 9", capture_index)
        self.assertLess(effect_index, capture_index)
        self.assertLess(capture_index, mutation_index)
        argument_effect = "argumentEffectCount = int(int32((argumentEffectCount + 1)))"
        self.assertEqual(main_go.count(argument_effect), 1)
        for forbidden in ["ArrayIterator", "reflect.", "unsafe.", "[]any(hx_structural_array"]:
            self.assertNotIn(forbidden, main_go)

        self.assertIn("public inline function iterator():Iterator<Xml>", xml_source)
        self.assertGreaterEqual(xml_printer_go.count("value.ensureElementType()"), 3)
        self.assertIn("value.children", xml_printer_go)
        self.assertNotIn("value.iterator()", xml_printer_go)

    def test_effectful_inline_concrete_tail_reuses_the_structural_adapter(self):
        compiler_source = GO_COMPILER.read_text(encoding="utf-8")
        lowering_source = LOWERING_MODULE.read_text(encoding="utf-8")
        main_go = (INLINE_CONCRETE_EFFECT_SNAPSHOT / "main.go").read_text(encoding="utf-8")

        validation_plan = function_block(
            lowering_source,
            "concreteStructuralIteratorPlan",
            "unwrapStructuralSourceExpr",
        )
        for required in [
            "structuralIteratorShape(toType)",
            'concreteIteratorMethod(sourceClass, sourceParams, "hasNext")',
            'concreteIteratorMethod(sourceClass, sourceParams, "next")',
        ]:
            self.assertIn(required, validation_plan)

        tail_plan = function_block(
            lowering_source,
            "inlineConcreteIteratorTailPlan",
            "inlineConcreteStructuralIteratorCoerce",
        )
        self.assertIn("exprs.slice(0, exprs.length - 1).concat(tailPlan.setup)", tail_plan)
        self.assertIn("concreteStructuralIteratorPlan(source.t, toType)", tail_plan)

        inline_adapter = function_block(
            lowering_source,
            "inlineConcreteStructuralIteratorCoerce",
            "structuralIteratorCoerce",
        )
        self.assertIn("lowerToStatements(setupExpr)", inline_adapter)
        self.assertIn("structuralIteratorCoerce(loweredTail.expr, tailPlan.sourceExpr.t, toType)", inline_adapter)
        for forbidden in ["nativeArrayCursorMap", "reflect", "unsafe", "GoStmt.GoRaw"]:
            self.assertNotIn(forbidden, inline_adapter)

        expected_upcast = function_block(compiler_source, "lowerExprWithExpectedUpcast", "upcastIfNeeded")
        self.assertIn("inlineConcreteStructuralIteratorCoerce", expected_upcast)
        self.assertLess(
            expected_upcast.index("nativeArrayStructuralIteratorCoerce"),
            expected_upcast.index("inlineConcreteStructuralIteratorCoerce"),
            "Array must retain its specialized live-slice plan before general concrete-tail recovery",
        )
        source_aware_upcast = function_block(compiler_source, "upcastIfNeeded", "typeToGoType")
        self.assertIn("materializeExprWithPrefix(inlineConcreteIterator, toType)", source_aware_upcast)
        self.assertLess(
            source_aware_upcast.index("nativeArrayStructuralIteratorCoerce"),
            source_aware_upcast.index("inlineConcreteStructuralIteratorCoerce"),
            "expression contexts must preserve the same Array-first ownership order",
        )

        main_start = main_go.index("func main()")
        main_end = main_go.index("func note", main_start)
        main_body = main_go[main_start:main_end]
        generic_effect = main_body.index('StringFromLiteral("generic:effect")')
        generic_constructor = main_body.index("New_SnapshotInlineGenericIterator", generic_effect)
        virtual_effect = main_body.index('StringFromLiteral("virtual:effect")', generic_constructor)
        virtual_constructor = main_body.index("New_SnapshotInlineSpecializedStringIterator", virtual_effect)
        self.assertLess(generic_effect, generic_constructor)
        self.assertLess(virtual_effect, virtual_constructor)
        self.assertEqual(main_body.count('StringFromLiteral("generic:effect")'), 1)
        self.assertEqual(main_body.count('StringFromLiteral("virtual:effect")'), 1)
        self.assertIn("any(hx_structural_iterator_", main_body)
        self.assertIn(".__hx_this.next()", main_body)
        for forbidden in ["ArrayIterator", "reflect.", "unsafe.", ") HasNext(", ") Next("]:
            self.assertNotIn(forbidden, main_body)

    def test_constructor_arguments_reuse_expected_type_coercion_and_emitted_abi(self):
        compiler_source = GO_COMPILER.read_text(encoding="utf-8")
        main_go = (CONSTRUCTOR_ITERATOR_SNAPSHOT / "main.go").read_text(encoding="utf-8")

        emitted_type = function_block(compiler_source, "emittedConstructorParamType", "lowerConstructorArg")
        self.assertIn("classType.constructor.get().type", emitted_type)
        self.assertNotIn(
            "applyTypeParameters",
            emitted_type,
            "the adapter closure shape must match the erased generated constructor ABI",
        )

        constructor_arg = function_block(compiler_source, "lowerConstructorArg", "lowerFunctionResults")
        self.assertIn("lowerExprWithExpectedUpcast", constructor_arg)
        self.assertIn("materializeExprWithPrefix", constructor_arg)
        self.assertGreaterEqual(compiler_source.count("lowerConstructorArg(classType"), 3)

        constructor_start = main_go.index("arrayConsumer := New_SnapshotIntConsumer")
        mutation_index = main_go.index("arrayValues[0] = 9", constructor_start)
        constructor_call = main_go[constructor_start:mutation_index]
        before_index = constructor_call.index('StringFromLiteral("before")')
        effect_index = constructor_call.index('StringFromLiteral("iterator")')
        after_index = constructor_call.index('StringFromLiteral("after")')
        capture_index = constructor_call.index(" := arrayValues")
        self.assertLess(before_index, effect_index)
        self.assertLess(effect_index, capture_index)
        self.assertLess(capture_index, after_index)

        self.assertIn("New_SnapshotGenericConsumer(func(hx_structural_iterator_", main_go)
        self.assertIn('["next"] = func() any', main_go)
        self.assertIn("any(hx_structural_iterator_", main_go)
        self.assertIn(".__hx_this.next()", main_go)
        for forbidden in ["ArrayIterator", "reflect.", "unsafe.", ") HasNext(", ") Next("]:
            self.assertNotIn(forbidden, main_go)


if __name__ == "__main__":
    unittest.main()
