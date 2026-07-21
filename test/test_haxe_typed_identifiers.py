#!/usr/bin/env python3

from __future__ import annotations

import json
import re
import unittest
from pathlib import Path


ROOT = Path(__file__).resolve().parent.parent
DEFINE_TYPE = ROOT / "src" / "reflaxe" / "go" / "compiler" / "GoCompilerDefine.hx"
METADATA_TYPE = ROOT / "src" / "reflaxe" / "go" / "compiler" / "GoMetadataName.hx"
COMPILER_BOOTSTRAP = ROOT / "src" / "reflaxe" / "go" / "CompilerBootstrap.hx"
DIAGNOSTICS = ROOT / "src" / "reflaxe" / "go" / "analyze" / "GoProfileContractAnalyzer.hx"
HXRT_FEATURES = ROOT / "src" / "reflaxe" / "go" / "compiler" / "GoHxrtFeatureAnalyzer.hx"
RELEASE_RUNNER = ROOT / "test" / "run-release-contracts.py"


class HaxeTypedIdentifierContractTest(unittest.TestCase):
    def test_typed_identifier_gate_is_release_blocking(self) -> None:
        scripts = json.loads((ROOT / "package.json").read_text(encoding="utf-8"))["scripts"]
        command = "python3 test/test_haxe_typed_identifiers.py"

        self.assertEqual(command, scripts.get("test:haxe-typed-identifiers"))
        self.assertIn("npm run test:haxe-typed-identifiers", scripts["test"])
        self.assertIn("npm run test:haxe-typed-identifiers", scripts["test:changed"])
        self.assertIn(
            "test/test_haxe_typed_identifiers.py",
            RELEASE_RUNNER.read_text(encoding="utf-8"),
        )

    def test_compiler_defines_are_closed_typed_values_at_macro_calls(self) -> None:
        source = DEFINE_TYPE.read_text(encoding="utf-8")
        self.assertIn("enum abstract GoCompilerDefine(String) to String", source)
        self.assertNotIn("from String", source)
        for value in [
            "go_output",
            "target.name",
            "reflaxe_go_profile",
            "reflaxe_go_native_authority",
            "reflaxe_go_hxrt_features",
            "go_no_build",
        ]:
            self.assertIn(f'= "{value}"', source)

        bootstrap = COMPILER_BOOTSTRAP.read_text(encoding="utf-8")
        self.assertIn('static inline final REFLAXE_DEFINE:String = "reflaxe";', bootstrap)
        self.assertNotIn("import reflaxe.go.compiler.GoCompilerDefine", bootstrap)

        literal_call = re.compile(
            r'(?:Context\.(?:defined|definedValue)|MacroCompiler\.define)\(\s*"'
        )
        offenders: list[str] = []
        for path in sorted((ROOT / "src" / "reflaxe" / "go").rglob("*.hx")):
            if path == DEFINE_TYPE:
                continue
            for line_number, line in enumerate(
                path.read_text(encoding="utf-8").splitlines(), start=1
            ):
                if literal_call.search(line):
                    offenders.append(f"{path.relative_to(ROOT)}:{line_number}")
        self.assertEqual([], offenders, "raw compiler-define literals: " + ", ".join(offenders))

    def test_metadata_and_diagnostic_identifiers_are_typed(self) -> None:
        metadata = METADATA_TYPE.read_text(encoding="utf-8")
        self.assertIn("enum abstract GoMetadataName(String) to String", metadata)
        self.assertNotIn("from String", metadata)
        for value in [
            "goNative",
            "goMetal",
            "goAllowRaw",
            "go.import",
            "go.name",
            "go.receiver",
            "go.tupleReturn",
        ]:
            self.assertIn(f'= "{value}"', metadata)

        diagnostics = DIAGNOSTICS.read_text(encoding="utf-8")
        self.assertIn("enum abstract GoContractDiagnosticCode(String) to String", diagnostics)
        self.assertIn("enum abstract GoContractDiagnosticSeverity(String) to String", diagnostics)
        self.assertIn("var code:GoContractDiagnosticCode;", diagnostics)
        self.assertIn("var severity:GoContractDiagnosticSeverity;", diagnostics)

    def test_runtime_surface_ids_are_typed_at_the_registry(self) -> None:
        source = HXRT_FEATURES.read_text(encoding="utf-8")
        self.assertIn("enum abstract GoHxrtFeatureId(String) to String", source)
        self.assertNotRegex(
            source,
            r"public static inline final FEATURE_[A-Z0-9_]+\s*=\s*\"",
        )


if __name__ == "__main__":
    unittest.main()
