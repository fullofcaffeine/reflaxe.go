import unittest
from pathlib import Path


ROOT = Path(__file__).resolve().parents[1]
GO_COMPILER = ROOT / "src" / "reflaxe" / "go" / "GoCompiler.hx"
LOWERING_MODULE = ROOT / "src" / "reflaxe" / "go" / "compiler" / "GoLambdaIterableLowering.hx"


class LambdaIterableLoweringOwnershipContract(unittest.TestCase):
    def test_lambda_iterable_adapter_logic_has_dedicated_owner(self):
        compiler_source = GO_COMPILER.read_text()
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
