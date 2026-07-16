import unittest
from pathlib import Path


ROOT = Path(__file__).resolve().parents[1]
GO_COMPILER = ROOT / "src" / "reflaxe" / "go" / "GoCompiler.hx"
LOWERING_MODULE = ROOT / "src" / "reflaxe" / "go" / "compiler" / "GoLambdaIterableLowering.hx"
LAMBDA_OVERRIDE = ROOT / "std" / "go" / "_std" / "Lambda.hx"


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
        for method in ["count", "empty", "exists", "has", "filter", "map", "fold", "iter"]:
            self.assertIn(
                f"public static function {method}",
                source,
                f"Lambda.{method} should lower from staged Haxe source",
            )
        for method in ["mapi", "flatten", "flatMap"]:
            self.assertIn(
                f"public static inline function {method}",
                source,
                f"Lambda.{method} should preserve upstream inline representation until full carrier parity",
            )

        compiler_source = GO_COMPILER.read_text(encoding="utf-8")
        self.assertNotIn(
            "function lowerLambdaStaticCall",
            compiler_source,
            "GoCompiler must not reimplement public Lambda algorithms",
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
            "consumerAnyAdapter",
            "folderAnyAdapter",
            "anyArrayCoerce",
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


if __name__ == "__main__":
    unittest.main()
