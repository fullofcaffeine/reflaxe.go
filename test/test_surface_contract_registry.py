#!/usr/bin/env python3

from __future__ import annotations

import json
import base64
import hashlib
import re
import shutil
import subprocess
import unittest
from pathlib import Path
from typing import Any


ROOT = Path(__file__).resolve().parent.parent
FIXTURE = ROOT / "test" / "fixtures" / "surface_contract_registry"
REGISTRY = ROOT / "src" / "reflaxe" / "go" / "compiler" / "GoSurfaceContractRegistry.hx"
PLANNER = ROOT / "src" / "reflaxe" / "go" / "compiler" / "GoSurfacePlanner.hx"
TYPE_LEDGER = ROOT / "src" / "reflaxe" / "go" / "compiler" / "GoTypeUsageLedger.hx"
HXRT_ANALYZER = ROOT / "src" / "reflaxe" / "go" / "compiler" / "GoHxrtFeatureAnalyzer.hx"
CONTEXT = ROOT / "src" / "reflaxe" / "go" / "CompilationContext.hx"
COMPILER = ROOT / "src" / "reflaxe" / "go" / "GoReflaxeCompiler.hx"
LOWERING_COMPILER = ROOT / "src" / "reflaxe" / "go" / "GoCompiler.hx"
DEFINE = ROOT / "src" / "reflaxe" / "go" / "compiler" / "GoCompilerDefine.hx"
SCHEMA = ROOT / "docs" / "schemas" / "surface-contracts-v2.schema.json"
LEGACY_SCHEMA = ROOT / "docs" / "schemas" / "surface-contracts-v1.schema.json"
LEGACY_SCHEMA_SHA256 = (
    "d5c0aa66702849d97b81a3990b7e8e5e5b3e7ba2afd178fa0e2e76756631b114"
)
DOC = ROOT / "docs" / "surface-contract-registry.md"
PLANNER_DOC = ROOT / "docs" / "surface-planner.md"
COLLECTION_DOC = ROOT / "docs" / "portable-collection-contract.md"
VALUE_SURFACE_DOC = ROOT / "docs" / "portable-string-bytes-iterator-contract.md"
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

        self.assertEqual(portable["schemaVersion"], 2)
        self.assertEqual(portable["registryVersion"], 1)
        self.assertEqual(
            portable["authority"],
            "typed_usage_plus_versioned_surface_contract",
        )
        self.assertEqual(
            portable["profileAdmission"],
            "forbidden",
        )
        self.assertEqual(portable["catalogCount"], 8)
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
        string_decisions = [
            entry
            for entry in portable["decisions"]
            if entry["surfaceId"] == "haxe.String"
        ]
        bytes_decisions = [
            entry
            for entry in portable["decisions"]
            if entry["surfaceId"] == "haxe.io.Bytes"
        ]
        iterator_decisions = [
            entry
            for entry in portable["decisions"]
            if entry["surfaceId"] == "haxe.Iterator"
        ]
        function_decisions = [
            entry
            for entry in portable["decisions"]
            if entry["surfaceId"] == "haxe.Function"
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
        self.assertTrue(string_decisions)
        self.assertTrue(
            all(
                entry["outcome"] == "admitted"
                and entry["selectedRepresentation"] == "go_string"
                and entry["runtimeRequirements"] == ["string"]
                for entry in string_decisions
            )
        )
        self.assertTrue(bytes_decisions)
        self.assertTrue(
            all(
                entry["outcome"] == "admitted"
                and entry["selectedRepresentation"] == "go_byte_slice"
                and entry["runtimeRequirements"] == ["bytes"]
                for entry in bytes_decisions
            )
        )
        self.assertTrue(
            any(
                entry["outcome"] == "admitted"
                and entry["selectedRepresentation"] == "go_iterator"
                for entry in iterator_decisions
            )
        )
        self.assertTrue(
            any(
                entry["outcome"] == "rejected"
                and entry["reason"] == "eligibility_rejected"
                and "binding_contains_no_dynamic:element"
                in entry["detail"]
                for entry in iterator_decisions
            )
        )
        self.assertTrue(function_decisions)
        self.assertTrue(
            all(
                entry["outcome"] == "rejected"
                and entry["reason"] == "contract_missing"
                for entry in function_decisions
            )
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
                "haxe.String",
                "haxe.io.Bytes",
                "haxe.ds.StringMap",
                "haxe.ds.IntMap",
                "haxe.Iterator",
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
        for (
            surface_id,
            native_representation,
            runtime_requirement,
            fallback_representation,
            no_hxrt_status,
        ) in (
            (
                "haxe.String",
                "go_string",
                ["string"],
                "hxrt_string",
                "ineligible",
            ),
            (
                "haxe.io.Bytes",
                "go_byte_slice",
                ["bytes"],
                "hxrt_bytes",
                "ineligible",
            ),
            (
                "haxe.Iterator",
                "go_iterator",
                [],
                "hxrt_iterator",
                "eligible",
            ),
        ):
            contract = contracts[surface_id]
            self.assertEqual(contract["sourceContract"], "portable_haxe")
            self.assertEqual(
                contract["nativeRepresentation"],
                native_representation,
            )
            self.assertEqual(
                contract["nativeRuntimeRequirements"],
                runtime_requirement,
            )
            self.assertEqual(
                contract["fallbackRepresentation"],
                fallback_representation,
            )
            self.assertEqual(contract["noHxrtStatus"], no_hxrt_status)
            self.assertIn(
                "semantic_diff",
                {proof["kind"] for proof in contract["proofs"]},
            )
        self.assertEqual(
            contracts["haxe.String"]["fallbackRuntimeRequirements"],
            ["string"],
        )
        self.assertEqual(
            contracts["haxe.io.Bytes"]["fallbackRuntimeRequirements"],
            ["bytes"],
        )
        self.assertIn(
            "bytes-normalization-semantic-diff",
            {
                proof["proofId"]
                for proof in contracts["haxe.io.Bytes"]["proofs"]
            },
        )
        self.assertEqual(
            contracts["haxe.Iterator"]["fallbackRuntimeRequirements"],
            ["core"],
        )
        self.assertIn(
            "binding_has_proven_collection_carrier:element",
            contracts["haxe.Iterator"]["eligibilityRules"],
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
        bytes_go = (
            FIXTURE / "out-portable" / "module_haxe_io_bytes.go"
        ).read_text(encoding="utf-8")
        main_go = (FIXTURE / "out-portable" / "main.go").read_text(encoding="utf-8")
        type_usage = json.loads(
            (FIXTURE / "out-portable" / "type_usage.json").read_text(
                encoding="utf-8"
            )
        )
        self.assertIn("type reflaxe__std__Option struct", option_go)
        self.assertIn("params []any", option_go)
        self.assertIn("type reflaxe__std__Result struct", result_go)
        self.assertIn("params []any", result_go)
        self.assertNotIn("go___Result", result_go)
        self.assertIn("*hxrt.Array", main_go)
        self.assertNotIn("var values []int", main_go)
        self.assertIn("b []int", bytes_go)
        self.assertIn("__hx_raw *hxrt.ByteView", bytes_go)
        self.assertIn("iterator map[string]any", main_go)
        self.assertIn('"hasNext"] = func() bool', main_go)
        self.assertIn('"next"] = func() int', main_go)
        observed_shapes = [
            usage["shape"]
            for module in type_usage["modules"]
            for usage in module["typeUsages"]
        ]
        self.assertTrue(
            any(
                shape_contains(
                    shape,
                    kind="typedef",
                    path_fragment="Main.IteratorAlias",
                )
                for shape in observed_shapes
            ),
            "a user-owned Iterator typedef must remain nominal/opaque in the "
            "real typed ledger",
        )
        observed_ids = {entry["surfaceId"] for entry in portable["decisions"]}
        self.assertTrue(
            {
                "haxe.Array",
                "haxe.String",
                "haxe.io.Bytes",
                "haxe.Iterator",
                "haxe.Function",
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

    def test_optimizer_import_and_runtime_plans_consume_one_registry_decision(self) -> None:
        portable, _, _ = self.compile_report("portable")
        portable_optimizer = json.loads(
            (FIXTURE / "out-portable" / "optimizer_plan.json").read_text(
                encoding="utf-8"
            )
        )
        portable_runtime = json.loads(
            (FIXTURE / "out-portable" / "hxrt_plan.json").read_text(
                encoding="utf-8"
            )
        )
        metal, _, _ = self.compile_report("metal")
        metal_optimizer = json.loads(
            (FIXTURE / "out-metal" / "optimizer_plan.json").read_text(
                encoding="utf-8"
            )
        )
        metal_runtime = json.loads(
            (FIXTURE / "out-metal" / "hxrt_plan.json").read_text(
                encoding="utf-8"
            )
        )

        self.assertEqual(portable, metal)
        self.assertEqual(
            portable_optimizer["nativeSpecializationPolicy"],
            "eager",
            "the portable lane must exercise the explicit eager override",
        )
        self.assertEqual(
            metal_optimizer["nativeSpecializationPolicy"],
            "proven",
            "the metal lane must exercise the explicit proven override",
        )
        self.assertEqual(
            portable_optimizer["surfacePlans"],
            metal_optimizer["surfacePlans"],
            "policy presets and specialization eagerness must not alter registry admission",
        )
        self.assertEqual(
            portable_optimizer["surfacePlans"],
            portable_runtime["surfacePlans"],
            "optimizer and runtime reports must consume the same immutable plan",
        )
        self.assertEqual(
            metal_optimizer["surfacePlans"],
            metal_runtime["surfacePlans"],
        )
        self.assertEqual(
            portable_optimizer["requiredSurfaceImports"],
            portable_runtime["requiredSurfaceImports"],
        )
        self.assertEqual(
            portable_optimizer["requiredSurfaceRuntimeFeatures"],
            portable_runtime["requiredSurfaceRuntimeFeatures"],
        )

        plans = portable_optimizer["surfacePlans"]
        self.assertTrue(plans)
        for entry in plans:
            self.assertEqual(
                set(entry),
                {
                    "module",
                    "location",
                    "usageLevel",
                    "usedType",
                    "contract",
                    "eligibility",
                    "selection",
                    "selectionReason",
                    "selectedRepresentation",
                    "fallbackReason",
                    "imports",
                    "runtimeRequirements",
                },
            )
            self.assertEqual(
                set(entry["contract"]),
                {"surfaceId", "version"},
            )
            self.assertEqual(
                set(entry["eligibility"]),
                {"outcome", "reason", "detail"},
            )
            self.assertIsInstance(entry["usedType"], dict)
            self.assertIn(entry["selection"], {"native", "fallback", "existing"})
            self.assertIsInstance(entry["imports"], list)
            self.assertIsInstance(entry["runtimeRequirements"], list)

        def matching(
            surface_id: str,
            *,
            eligibility: str | None = None,
            selection: str | None = None,
        ) -> list[dict[str, Any]]:
            return [
                entry
                for entry in plans
                if entry["contract"]["surfaceId"] == surface_id
                and (
                    eligibility is None
                    or entry["eligibility"]["outcome"] == eligibility
                )
                and (selection is None or entry["selection"] == selection)
            ]

        self.assertTrue(matching("haxe.String", selection="native"))
        self.assertTrue(matching("haxe.io.Bytes", selection="native"))
        self.assertTrue(
            matching(
                "haxe.Iterator",
                eligibility="admitted",
                selection="fallback",
            )
        )
        self.assertTrue(
            matching(
                "haxe.Iterator",
                eligibility="rejected",
                selection="fallback",
            )
        )
        admitted_array_fallbacks = matching(
            "haxe.Array",
            eligibility="admitted",
            selection="fallback",
        )
        self.assertTrue(admitted_array_fallbacks)
        self.assertTrue(
            all(
                entry["selectionReason"] == "carrier_not_activated"
                and entry["fallbackReason"]
                for entry in admitted_array_fallbacks
            )
        )
        rejected_array_fallbacks = matching(
            "haxe.Array",
            eligibility="rejected",
            selection="fallback",
        )
        self.assertTrue(rejected_array_fallbacks)
        self.assertTrue(
            all(
                entry["selectionReason"] == "registry_rejected"
                and entry["eligibility"]["reason"] in entry["fallbackReason"]
                for entry in rejected_array_fallbacks
            )
        )
        self.assertTrue(
            matching("reflaxe.std.Option", eligibility="admitted", selection="fallback")
        )
        self.assertTrue(
            matching("reflaxe.std.Result", eligibility="admitted", selection="fallback")
        )
        function_existing = matching(
            "haxe.Function",
            eligibility="rejected",
            selection="existing",
        )
        self.assertTrue(function_existing)
        self.assertTrue(
            all(
                entry["selectionReason"] == "no_registered_fallback"
                for entry in function_existing
            )
        )

        self.assertEqual(portable_optimizer["requiredSurfaceImports"], [])
        self.assertEqual(
            portable_optimizer["requiredSurfaceRuntimeFeatures"],
            ["array", "bytes", "core", "string"],
        )
        for value in flatten_strings(portable_optimizer):
            self.assertNotIn(str(ROOT), value)
            self.assertFalse(value.startswith("/"), value)
        optimizer_markdown = (
            FIXTURE / "out-portable" / "optimizer_plan.md"
        ).read_text(encoding="utf-8")
        for required_evidence in (
            "used type `",
            "location `",
            "eligibility detail `",
            "representation `",
            "imports `",
            "runtime `",
        ):
            self.assertIn(required_evidence, optimizer_markdown)

        planner_source = PLANNER.read_text(encoding="utf-8")
        context_source = CONTEXT.read_text(encoding="utf-8")
        compiler_source = COMPILER.read_text(encoding="utf-8")
        lowering_source = LOWERING_COMPILER.read_text(encoding="utf-8")
        self.assertNotIn(":Dynamic", planner_source)
        self.assertNotIn("<Dynamic>", planner_source)
        self.assertNotIn(":Any", planner_source)
        self.assertNotIn("<Any>", planner_source)
        self.assertNotIn("GoProfile", planner_source)
        self.assertNotIn("usesMetal", planner_source)
        self.assertIn(
            "public final surfacePlan:GoSurfacePlanSnapshot",
            context_source,
        )
        self.assertIn("GoSurfacePlanner.plan", compiler_source)
        self.assertIn("compilationContext.surfacePlan", lowering_source)
        self.assertIn("moduleDeclsUseSharedArraySurface", lowering_source)
        self.assertIn("lowered_go_ast:hxrt.Array", lowering_source)
        self.assertNotIn("generatedUsesSharedArraySurface", lowering_source)
        self.assertNotIn("file.contents.indexOf", lowering_source)

    def test_schema_docs_and_proof_paths_are_governed(self) -> None:
        schema = json.loads(SCHEMA.read_text(encoding="utf-8"))
        legacy_schema = json.loads(LEGACY_SCHEMA.read_text(encoding="utf-8"))
        self.assertEqual(
            hashlib.sha256(LEGACY_SCHEMA.read_bytes()).hexdigest(),
            LEGACY_SCHEMA_SHA256,
            "published schema v1 is immutable; add a new schema version instead",
        )
        self.assertEqual(
            schema["$id"],
            "urn:reflaxe-go:schema:surface-contracts:v2",
        )
        self.assertEqual(schema["properties"]["schemaVersion"]["const"], 2)
        self.assertEqual(schema["properties"]["registryVersion"]["const"], 1)
        self.assertEqual(
            legacy_schema["$id"],
            "urn:reflaxe-go:schema:surface-contracts:v1",
        )
        self.assertEqual(
            legacy_schema["properties"]["schemaVersion"]["const"],
            1,
        )
        self.assertNotIn(
            "haxe.Iterator",
            legacy_schema["$defs"]["surfaceId"]["enum"],
        )
        self.assertEqual(
            {
                entry["properties"]["kind"]["const"]
                for entry in legacy_schema["$defs"]["pattern"]["oneOf"]
            },
            {"bind", "nominal", "function"},
        )
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
        planner_doc = PLANNER_DOC.read_text(encoding="utf-8")
        collection_doc = COLLECTION_DOC.read_text(encoding="utf-8")
        value_surface_doc = VALUE_SURFACE_DOC.read_text(encoding="utf-8")
        defines_doc = DEFINES_DOC.read_text(encoding="utf-8")
        self.assertIn("What it is", doc)
        self.assertIn("Why it exists", doc)
        self.assertIn("How it works", doc)
        self.assertIn("Production admits eight", doc)
        self.assertNotIn("intentionally empty production catalog", doc)
        self.assertIn("not native `go.Result<T>`", doc)
        self.assertIn("planner](surface-planner.md) now consumes", doc)
        self.assertIn("shared, slice-backed carrier", collection_doc)
        self.assertIn("ObjectMap", collection_doc)
        self.assertIn("Std.string", collection_doc)
        self.assertIn("not semantic evidence", collection_doc)
        self.assertIn("nullable pointer-backed Go string", value_surface_doc)
        self.assertIn("shared data/view carrier", value_surface_doc)
        self.assertIn("exact `hasNext`/`next` protocol", value_surface_doc)
        self.assertIn("does not mean Go `range`", value_surface_doc)
        self.assertIn("haxe_go-vfp.7.11", value_surface_doc)
        self.assertIn("fast paths also require the selected String carrier", value_surface_doc)
        self.assertIn("haxe.rust", value_surface_doc)
        self.assertIn("haxe.elixir", value_surface_doc)
        self.assertIn("haxe.ruby", value_surface_doc)
        self.assertIn("Genes", value_surface_doc)
        self.assertIn("Profiles cannot admit a surface", doc)
        self.assertIn("haxe.rust", doc)
        self.assertIn("haxe.ruby", doc)
        self.assertIn("haxe.elixir.codex", doc)
        self.assertIn("Genes", doc)
        self.assertIn("reflaxe_go_surface_contract_report", defines_doc)
        self.assertIn("surface_contracts.json", defines_doc)
        self.assertIn("One authority, three consumers", planner_doc)
        self.assertIn("profile name, optimization define, or", planner_doc)
        self.assertIn("go.Slice", planner_doc)
        self.assertIn("optimizer_plan.json` schema v7", planner_doc)
        self.assertIn("hxrt_plan.json` schema v3", planner_doc)

        populated = self.populated_contract_report()
        self.validate_schema(populated)
        populated_report = json.loads(populated)
        self.assertEqual(populated_report["catalogCount"], 8)
        self.assertEqual(populated_report["admittedCount"], 6)
        self.assertEqual(populated_report["rejectedCount"], 6)
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
        self.assertIn(
            ("haxe.Iterator", "admitted", "contract_admitted"),
            populated_decisions,
        )
        self.assertIn(
            ("haxe.Iterator", "rejected", "eligibility_rejected"),
            populated_decisions,
        )
        self.assertIn(
            ("haxe.Function", "rejected", "contract_missing"),
            populated_decisions,
        )
        pattern_kinds = {
            entry["properties"]["kind"]["const"]
            for entry in schema["$defs"]["pattern"]["oneOf"]
        }
        self.assertEqual(
            pattern_kinds,
            {"bind", "nominal", "function", "anonymous"},
        )
        self.assertFalse(schema["$defs"]["patternField"]["additionalProperties"])

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
