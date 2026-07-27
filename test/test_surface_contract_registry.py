#!/usr/bin/env python3

from __future__ import annotations

import json
import base64
import re
import shutil
import subprocess
import unittest
from pathlib import Path
from typing import Any


ROOT = Path(__file__).resolve().parent.parent
FIXTURE = ROOT / "test" / "fixtures" / "surface_contract_registry"
REGISTRY = ROOT / "src" / "reflaxe" / "go" / "compiler" / "GoSurfaceContractRegistry.hx"
TYPE_LEDGER = ROOT / "src" / "reflaxe" / "go" / "compiler" / "GoTypeUsageLedger.hx"
HXRT_ANALYZER = ROOT / "src" / "reflaxe" / "go" / "compiler" / "GoHxrtFeatureAnalyzer.hx"
CONTEXT = ROOT / "src" / "reflaxe" / "go" / "CompilationContext.hx"
COMPILER = ROOT / "src" / "reflaxe" / "go" / "GoReflaxeCompiler.hx"
DEFINE = ROOT / "src" / "reflaxe" / "go" / "compiler" / "GoCompilerDefine.hx"
SCHEMA = ROOT / "docs" / "schemas" / "surface-contracts-v1.schema.json"
DOC = ROOT / "docs" / "surface-contract-registry.md"
COLLECTION_DOC = ROOT / "docs" / "portable-collection-contract.md"
DEFINES_DOC = ROOT / "docs" / "defines-reference.md"
NATIVE_MAP = ROOT / "src" / "go" / "Map.hx"
SCHEMA_VALIDATOR = ROOT / "test" / "validate_surface_contract_schema.mjs"


def flatten_strings(value: Any) -> list[str]:
    if isinstance(value, str):
        return [value]
    if isinstance(value, list):
        return [item for entry in value for item in flatten_strings(entry)]
    if isinstance(value, dict):
        return [item for entry in value.values() for item in flatten_strings(entry)]
    return []

def shape_contains(shape: dict[str, Any], *, kind: str, path_fragment: str) -> bool:
    if shape.get("kind") == kind and path_fragment in shape.get("path", ""):
        return True
    for parameter in shape.get("parameters", []):
        if shape_contains(parameter, kind=kind, path_fragment=path_fragment):
            return True
    for argument in shape.get("arguments", []):
        if shape_contains(argument["shape"], kind=kind, path_fragment=path_fragment):
            return True
    return_type = shape.get("returnType")
    if isinstance(return_type, dict) and shape_contains(
        return_type,
        kind=kind,
        path_fragment=path_fragment,
    ):
        return True
    for field in shape.get("fields", []):
        if shape_contains(field["shape"], kind=kind, path_fragment=path_fragment):
            return True
    return False


class SurfaceContractRegistryTest(unittest.TestCase):
    def run_macro_contract(self) -> str:
        proc = subprocess.run(
            [
                "haxe",
                "-cp",
                "src",
                "-cp",
                str(FIXTURE),
                "-lib",
                "reflaxe",
                "--macro",
                "RegistryContractHarness.run()",
                "--no-output",
            ],
            cwd=ROOT,
            capture_output=True,
            text=True,
            check=False,
        )
        self.assertEqual(proc.returncode, 0, proc.stdout + proc.stderr)
        self.assertIn("surface registry macro contract passed", proc.stdout + proc.stderr)
        return proc.stdout + proc.stderr

    def validate_schema(self, report: bytes) -> None:
        proc = subprocess.run(
            ["node", str(SCHEMA_VALIDATOR), str(SCHEMA), "-"],
            cwd=ROOT,
            input=report,
            capture_output=True,
            check=False,
        )
        self.assertEqual(proc.returncode, 0, proc.stdout.decode() + proc.stderr.decode())

    def populated_contract_report(self) -> bytes:
        output = self.run_macro_contract()
        marker = "SURFACE_REGISTRY_JSON="
        encoded = next(line[len(marker) :] for line in output.splitlines() if line.startswith(marker))
        return base64.b64decode(encoded)

    def compile_report(self, profile: str) -> tuple[dict[str, Any], bytes, bytes]:
        out = FIXTURE / f"out-{profile}"
        shutil.rmtree(out, ignore_errors=True)
        self.addCleanup(shutil.rmtree, out, ignore_errors=True)
        proc = subprocess.run(
            ["haxe", f"compile.{profile}.hxml"],
            cwd=FIXTURE,
            capture_output=True,
            text=True,
            check=False,
        )
        self.assertEqual(proc.returncode, 0, proc.stdout + proc.stderr)
        json_path = out / "surface_contracts.json"
        markdown_path = out / "surface_contracts.md"
        self.assertTrue(json_path.is_file())
        self.assertTrue(markdown_path.is_file())
        return (
            json.loads(json_path.read_text(encoding="utf-8")),
            json_path.read_bytes(),
            markdown_path.read_bytes(),
        )

    def test_typed_schema_validation_and_fail_closed_evaluation(self) -> None:
        self.populated_contract_report()
        source = REGISTRY.read_text(encoding="utf-8")
        self.assertNotIn(":Dynamic", source)
        self.assertNotIn("<Dynamic>", source)
        self.assertNotIn(":Any", source)
        self.assertNotIn("<Any>", source)
        self.assertIn("GoImmutableList<GoSurfaceContract>", source)
        self.assertIn("GoImmutableList<GoSurfaceDecision>", source)
        self.assertIn("GoSurfaceTypePattern", source)
        self.assertIn("GoSurfaceValidationCode", source)

    def test_report_is_deterministic_and_profile_independent(self) -> None:
        portable, portable_json, portable_markdown = self.compile_report("portable")
        portable_again, portable_json_again, portable_markdown_again = self.compile_report("portable")
        metal, metal_json, metal_markdown = self.compile_report("metal")

        self.assertEqual(portable, portable_again)
        self.assertEqual(portable_json, portable_json_again)
        self.assertEqual(portable_markdown, portable_markdown_again)
        self.assertEqual(portable, metal)
        self.assertEqual(portable_json, metal_json)
        self.assertEqual(portable_markdown, metal_markdown)

        self.assertEqual(portable["schemaVersion"], 1)
        self.assertEqual(portable["registryVersion"], 1)
        self.assertEqual(
            portable["authority"],
            "typed_usage_plus_versioned_surface_contract",
        )
        self.assertEqual(
            portable["profileAdmission"],
            "forbidden",
        )
        self.assertEqual(portable["catalogCount"], 5)
        self.assertGreater(portable["decisionCount"], 0)
        self.assertEqual(
            portable["decisionCount"],
            portable["admittedCount"] + portable["rejectedCount"],
        )
        option_decisions = [
            entry
            for entry in portable["decisions"]
            if entry["surfaceId"] == "reflaxe.std.Option"
        ]
        result_decisions = [
            entry
            for entry in portable["decisions"]
            if entry["surfaceId"] == "reflaxe.std.Result"
        ]
        array_decisions = [
            entry
            for entry in portable["decisions"]
            if entry["surfaceId"] == "haxe.Array"
        ]
        self.assertTrue(
            any(
                entry["outcome"] == "admitted"
                and entry["selectedRepresentation"] == "go_option"
                for entry in option_decisions
            )
        )
        self.assertTrue(
            any(
                entry["outcome"] == "admitted"
                and entry["selectedRepresentation"] == "go_result"
                for entry in result_decisions
            )
        )
        self.assertTrue(
            any(entry["reason"] == "eligibility_rejected" for entry in option_decisions)
        )
        self.assertTrue(
            any(entry["reason"] == "eligibility_rejected" for entry in result_decisions)
        )
        self.assertTrue(
            any(
                entry["outcome"] == "admitted"
                and entry["selectedRepresentation"] == "go_slice"
                for entry in array_decisions
            )
        )
        self.assertTrue(
            any(entry["reason"] == "eligibility_rejected" for entry in array_decisions)
        )
        hidden_alias_decisions = [
            entry
            for entry in array_decisions
            if shape_contains(
                entry["shape"],
                kind="typedef",
                path_fragment="HiddenDynamic",
            )
            or shape_contains(
                entry["shape"],
                kind="abstract",
                path_fragment="HiddenDynamicAbstract",
            )
        ]
        self.assertTrue(hidden_alias_decisions)
        self.assertTrue(
            all(entry["outcome"] == "rejected" for entry in hidden_alias_decisions)
        )
        self.assertTrue(
            all(
                entry["reason"] == "eligibility_rejected"
                and "binding_has_proven_collection_carrier:element"
                in entry["detail"]
                for entry in hidden_alias_decisions
            )
        )
        self.assertTrue(
            all(entry["shape"]["parameters"] for entry in option_decisions)
        )
        self.assertTrue(
            all(entry["shape"]["parameters"] for entry in result_decisions)
        )
        contracts = {
            entry["surfaceId"]: entry for entry in portable["contracts"]
        }
        self.assertEqual(
            set(contracts),
            {
                "haxe.Array",
                "haxe.ds.StringMap",
                "haxe.ds.IntMap",
                "reflaxe.std.Option",
                "reflaxe.std.Result",
            },
        )
        for surface_id, fallback_runtime in (
            ("haxe.Array", ["array"]),
            ("haxe.ds.StringMap", ["map_string"]),
            ("haxe.ds.IntMap", ["map_int"]),
        ):
            contract = contracts[surface_id]
            self.assertEqual(contract["sourceContract"], "portable_haxe")
            carrier_binding = "element" if surface_id == "haxe.Array" else "value"
            self.assertIn(
                f"binding_has_proven_collection_carrier:{carrier_binding}",
                contract["eligibilityRules"],
            )
            if surface_id == "haxe.Array":
                self.assertNotIn(
                    "surface_has_fixed_go_comparable_map_key",
                    contract["eligibilityRules"],
                )
            else:
                self.assertIn(
                    "surface_has_fixed_go_comparable_map_key",
                    contract["eligibilityRules"],
                )
            self.assertEqual(contract["nativeRuntimeRequirements"], [])
            self.assertEqual(
                contract["fallbackRuntimeRequirements"],
                fallback_runtime,
            )
            self.assertEqual(contract["noHxrtStatus"], "conditional")
            self.assertIn(
                "portable-collections-semantic-diff",
                {proof["proofId"] for proof in contract["proofs"]},
            )
            if surface_id in {"haxe.ds.StringMap", "haxe.ds.IntMap"}:
                self.assertNotIn(
                    "generated_shape",
                    {proof["kind"] for proof in contract["proofs"]},
                    "map contracts must not claim an unasserted generated shape",
                )
        for surface_id, native_representation, fallback_representation in (
            ("reflaxe.std.Option", "go_option", "portable_option"),
            ("reflaxe.std.Result", "go_result", "portable_result"),
        ):
            contract = contracts[surface_id]
            self.assertEqual(
                contract["sourceContract"],
                "portable_family_facade",
            )
            self.assertEqual(
                contract["nativeRepresentation"],
                native_representation,
            )
            self.assertEqual(
                contract["fallbackRepresentation"],
                fallback_representation,
            )
            self.assertEqual(contract["nativeRuntimeRequirements"], [])
            self.assertEqual(contract["fallbackRuntimeRequirements"], [])
            self.assertEqual(contract["noHxrtStatus"], "eligible")
            self.assertEqual(contract["familySyncExpectation"], "target_local")
            self.assertEqual(contract["familyContractId"], "")
            self.assertEqual(contract["familyContractVersion"], 0)
            self.assertIn(
                "generated_shape",
                {proof["kind"] for proof in contract["proofs"]},
            )

        option_go = (
            FIXTURE / "out-portable" / "module_reflaxe_std_option.go"
        ).read_text(encoding="utf-8")
        result_go = (
            FIXTURE / "out-portable" / "module_reflaxe_std_result.go"
        ).read_text(encoding="utf-8")
        main_go = (FIXTURE / "out-portable" / "main.go").read_text(encoding="utf-8")
        self.assertIn("type reflaxe__std__Option struct", option_go)
        self.assertIn("params []any", option_go)
        self.assertIn("type reflaxe__std__Result struct", result_go)
        self.assertIn("params []any", result_go)
        self.assertNotIn("go___Result", result_go)
        self.assertIn("*hxrt.Array", main_go)
        self.assertNotIn("var values []int", main_go)
        rejected = [
            entry
            for entry in portable["decisions"]
            if entry["surfaceId"] in {"haxe.String", "haxe.io.Bytes"}
        ]
        self.assertTrue(rejected)
        self.assertTrue(all(entry["outcome"] == "rejected" for entry in rejected))
        self.assertTrue(
            all(entry["reason"] == "contract_missing" for entry in rejected)
        )
        observed_ids = {entry["surfaceId"] for entry in portable["decisions"]}
        self.assertTrue(
            {
                "haxe.Array",
                "haxe.String",
                "haxe.io.Bytes",
            }.issubset(observed_ids)
        )

        for value in flatten_strings(portable):
            self.assertNotIn(str(ROOT), value)
            self.assertFalse(value.startswith("/"), value)
            self.assertNotIn("metal", value.lower())
        self.validate_schema(portable_json)

    def test_registry_is_one_final_authority_without_profile_bypass(self) -> None:
        registry_source = REGISTRY.read_text(encoding="utf-8")
        context_source = CONTEXT.read_text(encoding="utf-8")
        compiler_source = COMPILER.read_text(encoding="utf-8")
        define_source = DEFINE.read_text(encoding="utf-8")

        self.assertNotIn("GoProfile", registry_source)
        self.assertNotIn("buildContext.profile", registry_source)
        self.assertNotIn("usesMetal", registry_source)
        self.assertNotIn("Metal", registry_source)
        self.assertNotIn("Std.string", registry_source)
        self.assertNotIn('"go.Map"', registry_source)
        self.assertIn("Std.string(key)", NATIVE_MAP.read_text(encoding="utf-8"))
        self.assertIn(
            "public final surfaceContractRegistry:GoSurfaceContractRegistrySnapshot",
            context_source,
        )
        self.assertNotIn("context.surfaceContractRegistry =", compiler_source)
        self.assertIn(
            'DefineSurfaceContractReport = "reflaxe_go_surface_contract_report"',
            define_source,
        )

    def test_schema_docs_and_proof_paths_are_governed(self) -> None:
        schema = json.loads(SCHEMA.read_text(encoding="utf-8"))
        self.assertEqual(
            schema["$id"],
            "urn:reflaxe-go:schema:surface-contracts:v1",
        )
        self.assertEqual(schema["properties"]["schemaVersion"]["const"], 1)
        self.assertEqual(schema["properties"]["registryVersion"]["const"], 1)
        self.assertEqual(
            schema["properties"]["profileAdmission"]["const"],
            "forbidden",
        )
        self.assertFalse(schema["additionalProperties"])
        self.assertFalse(schema["$defs"]["contract"]["additionalProperties"])
        self.assertFalse(schema["$defs"]["decision"]["additionalProperties"])
        self.assertEqual(
            schema["$defs"]["decisionReason"]["enum"],
            [
                "contract_admitted",
                "contract_missing",
                "shape_mismatch",
                "eligibility_rejected",
            ],
        )
        simple_eligibility_values = schema["$defs"]["contract"]["properties"][
            "eligibilityRules"
        ]["items"]["anyOf"][0]["enum"]
        self.assertIn(
            "surface_has_fixed_go_comparable_map_key",
            simple_eligibility_values,
        )
        binding_rule_pattern = schema["$defs"]["contract"]["properties"][
            "eligibilityRules"
        ]["items"]["anyOf"][1]["pattern"]
        self.assertIn("has_proven_collection_carrier", binding_rule_pattern)

        doc = DOC.read_text(encoding="utf-8")
        collection_doc = COLLECTION_DOC.read_text(encoding="utf-8")
        defines_doc = DEFINES_DOC.read_text(encoding="utf-8")
        self.assertIn("What it is", doc)
        self.assertIn("Why it exists", doc)
        self.assertIn("How it works", doc)
        self.assertIn("Production admits five", doc)
        self.assertNotIn("intentionally empty production catalog", doc)
        self.assertIn("not native `go.Result<T>`", doc)
        self.assertIn("planner does not consume", doc)
        self.assertIn("shared, slice-backed carrier", collection_doc)
        self.assertIn("ObjectMap", collection_doc)
        self.assertIn("Std.string", collection_doc)
        self.assertIn("not semantic evidence", collection_doc)
        self.assertIn("Profiles cannot admit a surface", doc)
        self.assertIn("haxe.rust", doc)
        self.assertIn("haxe.ruby", doc)
        self.assertIn("haxe.elixir.codex", doc)
        self.assertIn("Genes", doc)
        self.assertIn("reflaxe_go_surface_contract_report", defines_doc)
        self.assertIn("surface_contracts.json", defines_doc)

        populated = self.populated_contract_report()
        self.validate_schema(populated)
        populated_report = json.loads(populated)
        self.assertEqual(populated_report["catalogCount"], 5)
        self.assertEqual(populated_report["admittedCount"], 3)
        self.assertEqual(populated_report["rejectedCount"], 4)
        populated_decisions = {
            (entry["surfaceId"], entry["outcome"], entry["reason"])
            for entry in populated_report["decisions"]
        }
        self.assertIn(
            ("haxe.ds.StringMap", "admitted", "contract_admitted"),
            populated_decisions,
        )
        self.assertIn(
            ("haxe.ds.StringMap", "rejected", "eligibility_rejected"),
            populated_decisions,
        )
        self.assertIn(
            ("haxe.ds.IntMap", "rejected", "eligibility_rejected"),
            populated_decisions,
        )
        self.assertIn(
            ("haxe.ds.ObjectMap", "rejected", "contract_missing"),
            populated_decisions,
        )

        registry_source = REGISTRY.read_text(encoding="utf-8")
        ledger_source = TYPE_LEDGER.read_text(encoding="utf-8")
        hxrt_source = HXRT_ANALYZER.read_text(encoding="utf-8")
        self.assertNotIn(
            "does not enable any of them in the production catalog",
            registry_source,
        )

        def enum_values(source: str, name: str) -> list[str]:
            body = re.search(
                rf"enum abstract {name}\(String\).*?\{{(.*?)\n\}}",
                source,
                re.DOTALL,
            )
            self.assertIsNotNone(body, name)
            return re.findall(r'\bvar \w+ = "([^"]+)";', body.group(1))

        vocabularies = {
            "GoSurfaceId": "surfaceId",
            "GoSourceContractKind": "sourceContract",
            "GoNativeRepresentation": "nativeRepresentation",
            "GoSurfaceFallbackRepresentation": "fallbackRepresentation",
            "GoSurfaceFallbackPolicy": "fallbackPolicy",
            "GoNoHxrtStatus": "noHxrtStatus",
            "GoFamilySyncExpectation": "familySyncExpectation",
            "GoSurfaceProofKind": "proofKind",
            "GoSurfaceDecisionOutcome": "decisionOutcome",
            "GoSurfaceDecisionReason": "decisionReason",
            "GoSurfaceNominalKind": "nominalTargetKind",
        }
        for abstract_name, schema_name in vocabularies.items():
            self.assertEqual(
                enum_values(registry_source, abstract_name),
                schema["$defs"][schema_name]["enum"],
                f"{abstract_name} and schema {schema_name} must stay synchronized",
            )
        self.assertEqual(
            enum_values(ledger_source, "GoTypeUsageLevelId"),
            schema["$defs"]["usageLevel"]["enum"],
        )
        self.assertEqual(
            enum_values(hxrt_source, "GoHxrtFeatureId"),
            schema["$defs"]["runtimeRequirement"]["enum"],
        )

        for contract in populated_report["contracts"]:
            for proof in contract["proofs"]:
                self.assertFalse(Path(proof["fixturePath"]).is_absolute())
                self.assertTrue((ROOT / proof["fixturePath"]).exists())


if __name__ == "__main__":
    unittest.main()
