#if macro
import haxe.crypto.Base64;
import haxe.io.Bytes;
import reflaxe.go.compiler.GoHxrtFeatureAnalyzer.GoHxrtFeatureId;
import reflaxe.go.compiler.GoSurfaceContractRegistry;
import reflaxe.go.compiler.GoSurfaceContractRegistry.GoSurfaceContractRegistryException;
import reflaxe.go.compiler.GoSurfaceContractRegistry.GoFamilySyncExpectation;
import reflaxe.go.compiler.GoSurfaceContractRegistry.GoNativeRepresentation;
import reflaxe.go.compiler.GoSurfaceContractRegistry.GoNoHxrtStatus;
import reflaxe.go.compiler.GoSurfaceContractRegistry.GoSourceContractKind;
import reflaxe.go.compiler.GoSurfaceContractRegistry.GoSurfaceContract;
import reflaxe.go.compiler.GoSurfaceContractRegistry.GoSurfaceDecisionOutcome;
import reflaxe.go.compiler.GoSurfaceContractRegistry.GoSurfaceDecisionReason;
import reflaxe.go.compiler.GoSurfaceContractRegistry.GoSurfaceEligibilityRule;
import reflaxe.go.compiler.GoSurfaceContractRegistry.GoSurfaceFallbackPolicy;
import reflaxe.go.compiler.GoSurfaceContractRegistry.GoSurfaceFallbackRepresentation;
import reflaxe.go.compiler.GoSurfaceContractRegistry.GoSurfaceId;
import reflaxe.go.compiler.GoSurfaceContractRegistry.GoSurfaceProof;
import reflaxe.go.compiler.GoSurfaceContractRegistry.GoSurfaceProofKind;
import reflaxe.go.compiler.GoSurfaceContractRegistry.GoSurfaceNominalKind;
import reflaxe.go.compiler.GoSurfaceContractRegistry.GoSurfaceTypePattern;
import reflaxe.go.compiler.GoSurfaceContractRegistry.GoSurfaceValidationCode;
import reflaxe.go.compiler.GoTypeUsageLedger.GoImmutableList;
import reflaxe.go.compiler.GoTypeUsageLedger.GoTypeShape;
import reflaxe.go.compiler.GoTypeUsageLedger.GoTypeUsageEvidence;
import reflaxe.go.compiler.GoTypeUsageLedger.GoTypeUsageLedgerSnapshot;
import reflaxe.go.compiler.GoTypeUsageLedger.GoTypeUsageLevelId;
import reflaxe.go.compiler.GoTypeUsageLedger.GoTypeUsageModuleEvidence;
import reflaxe.go.compiler.GoTypeUsageLedger.GoTypeUsageTargetKind;
import reflaxe.go.compiler.GoTypeUsageLedger.GoUnknownTypeShapeReason;

class RegistryContractHarness {
	public static function run():Void {
		final valid = arrayContract();
		assertHasIssue([valid, valid], GoSurfaceValidationCode.DuplicateSurface);
		assertHasIssue([arrayContract(cast "unknown.surface")], GoSurfaceValidationCode.UnknownSurface);
		assertHasIssue([arrayContract(null, GoImmutableList.fromArray([]))], GoSurfaceValidationCode.MissingSemanticProof);
		assertHasIssue([arrayContract(null, null, GoImmutableList.fromArray([cast "not_a_feature"]))], GoSurfaceValidationCode.UnknownRuntimeRequirement);
		assertHasIssue([arrayContract(null, null, null, true)], GoSurfaceValidationCode.InvalidShape);
		assertHasIssue([arrayContract(null, null, null, false, true)], GoSurfaceValidationCode.InvalidShape);
		assertHasIssue([arrayContract(null, null, null, false, false, true)], GoSurfaceValidationCode.InvalidShape);
		assertHasIssue([arrayContract(null, null, null, false, false, false, false, "binding_colon")], GoSurfaceValidationCode.InvalidShape);
		assertHasIssue([arrayContract(null, null, null, false, false, false, true)], GoSurfaceValidationCode.InvalidShape);
		assertCreateRejects([arrayContract(null, null, null, true)], GoSurfaceValidationCode.InvalidShape);
		assertCreateRejects([arrayContract(null, null, null, false, true)], GoSurfaceValidationCode.InvalidShape);
		assertCreateRejects(null, GoSurfaceValidationCode.MalformedContract);
		assertHasIssue([arrayContract(null, null, null, false, false, false, false, "pattern_element")], GoSurfaceValidationCode.InvalidShape);
		assertHasIssue([arrayContract(null, null, null, false, false, false, false, "rule_list")], GoSurfaceValidationCode.UnknownEligibilityRule);
		assertHasIssue([arrayContract(null, null, null, false, false, false, false, "rule_element")], GoSurfaceValidationCode.UnknownEligibilityRule);
		assertHasIssue([arrayContract(null, null, null, false, false, false, false, "import_list")], GoSurfaceValidationCode.InvalidImportRequirement);
		assertHasIssue([arrayContract(null, null, null, false, false, false, false, "import_element")], GoSurfaceValidationCode.InvalidImportRequirement);
		assertHasIssue([arrayContract(null, null, null, false, false, false, false, "runtime_list")], GoSurfaceValidationCode.UnknownRuntimeRequirement);
		assertHasIssue([arrayContract(null, null, null, false, false, false, false, "proof_list")], GoSurfaceValidationCode.MissingProof);
		assertHasIssue([arrayContract(null, null, null, false, false, false, false, "proof_element")], GoSurfaceValidationCode.MissingProof);
		assertHasIssue([
			arrayContract(null, GoImmutableList.fromArray([
				{
					proofId: "windows-absolute",
					kind: GoSurfaceProofKind.SemanticDiff,
					fixturePath: "C:/workspace/proof"
				}
			]))
		], GoSurfaceValidationCode.UnsafeProofPath);
		assertHasIssue([
			arrayContract(null, GoImmutableList.fromArray([
				{
					proofId: "windows-drive-relative",
					kind: GoSurfaceProofKind.SemanticDiff,
					fixturePath: "C:fixture"
				}
			]))
		], GoSurfaceValidationCode.UnsafeProofPath);

		final source = [valid];
		final registry = GoSurfaceContractRegistry.create(source);
		source.resize(0);
		final admitted = registry.snapshot(ledger(arrayShape(intShape())));
		assertEquals(1, admitted.catalogCount, "registry must deep-copy its input catalog");
		assertEquals(1, admitted.decisionCount, "one known Array shape should produce one decision");
		assertEquals(GoSurfaceDecisionOutcome.Admitted, admitted.decisions.at(0).outcome, "valid proven Array contract should admit");
		assertEquals(GoSurfaceDecisionReason.ContractAdmitted, admitted.decisions.at(0).reason, "admission reason should be explicit");
		assertEquals(GoNativeRepresentation.GoSlice, admitted.decisions.at(0).selectedRepresentation, "native representation must come from the contract");
		assertEquals(GoSurfaceFallbackRepresentation.HxrtArray, admitted.decisions.at(0)
			.fallbackRepresentation, "fallback representation must remain visible");

		final rejected = GoSurfaceContractRegistry.create([valid]).snapshot(ledger(arrayShape(GoTypeShape.DynamicShape(null))));
		assertEquals(GoSurfaceDecisionOutcome.Rejected, rejected.decisions.at(0).outcome, "Dynamic element shape must fail closed");
		assertEquals(GoSurfaceDecisionReason.EligibilityRejected, rejected.decisions.at(0).reason, "shape rejection reason should be stable");

		final missing = GoSurfaceContractRegistry.defaultRegistry().snapshot(ledger(arrayShape(intShape())));
		assertEquals(2, missing.catalogCount, "production must admit only the proven Option and Result facade contracts");
		assertEquals(GoSurfaceDecisionOutcome.Rejected, missing.decisions.at(0).outcome, "known but unregistered surface must reject");
		assertEquals(GoSurfaceDecisionReason.ContractMissing, missing.decisions.at(0).reason, "missing contract must be explained");

		final option = GoSurfaceContractRegistry.defaultRegistry().snapshot(ledger(optionShape(intShape())));
		assertEquals(GoSurfaceDecisionOutcome.Admitted, option.decisions.at(0).outcome, "fully typed portable Option must be admitted");
		assertEquals(GoNativeRepresentation.GoOption, option.decisions.at(0)
			.selectedRepresentation, "portable Option must select the typed Go option carrier");
		assertEquals(GoSurfaceFallbackRepresentation.PortableOption, option.decisions.at(0).fallbackRepresentation,
			"portable Option must retain its tagged fallback");
		assertEquals(0, option.decisions.at(0).runtimeRequirements.length, "typed Option carrier must not require hxrt");

		final genericOption = GoSurfaceContractRegistry.defaultRegistry().snapshot(ledger(optionShape(GoTypeShape.TypeParameter("T"))));
		assertEquals(GoSurfaceDecisionOutcome.Admitted, genericOption.decisions.at(0).outcome,
			"a named generic Option parameter is typed and must remain eligible");

		final dynamicOption = GoSurfaceContractRegistry.defaultRegistry().snapshot(ledger(optionShape(GoTypeShape.DynamicShape(null))));
		assertEquals(GoSurfaceDecisionOutcome.Rejected, dynamicOption.decisions.at(0).outcome, "Dynamic Option payload must fall back");
		assertEquals(GoSurfaceDecisionReason.EligibilityRejected, dynamicOption.decisions.at(0).reason,
			"Dynamic Option fallback must name the eligibility boundary");
		assertEquals(GoSurfaceFallbackRepresentation.PortableOption, dynamicOption.decisions.at(0).fallbackRepresentation,
			"rejected Option must retain source semantics through its fallback");

		final result = GoSurfaceContractRegistry.defaultRegistry().snapshot(ledger(resultShape(intShape(), stringShape())));
		assertEquals(GoSurfaceDecisionOutcome.Admitted, result.decisions.at(0).outcome, "fully typed portable Result must be admitted");
		assertEquals(GoNativeRepresentation.GoResult, result.decisions.at(0).selectedRepresentation,
			"portable Result must select a typed two-parameter Go result carrier");
		assertEquals(GoSurfaceFallbackRepresentation.PortableResult, result.decisions.at(0).fallbackRepresentation,
			"portable Result must retain its tagged fallback");
		assertEquals(0, result.decisions.at(0).runtimeRequirements.length, "typed Result carrier must not require hxrt");

		final genericResult = GoSurfaceContractRegistry.defaultRegistry()
			.snapshot(ledger(resultShape(GoTypeShape.TypeParameter("T"), GoTypeShape.TypeParameter("E"))));
		assertEquals(GoSurfaceDecisionOutcome.Admitted, genericResult.decisions.at(0).outcome,
			"named generic Result parameters are typed and must preserve both T and E");

		final dynamicError = GoSurfaceContractRegistry.defaultRegistry().snapshot(ledger(resultShape(intShape(), GoTypeShape.DynamicShape(null))));
		assertEquals(GoSurfaceDecisionOutcome.Rejected, dynamicError.decisions.at(0).outcome, "Dynamic Result error payload must fall back");
		assertEquals(GoSurfaceFallbackRepresentation.PortableResult, dynamicError.decisions.at(0).fallbackRepresentation,
			"rejected Result must retain its typed-error source contract");

		final unknownError = GoSurfaceContractRegistry.defaultRegistry()
			.snapshot(ledger(resultShape(intShape(), GoTypeShape.UnknownShape(GoUnknownTypeShapeReason.Missing))));
		assertEquals(GoSurfaceDecisionOutcome.Rejected, unknownError.decisions.at(0).outcome, "unresolved Result error payload must fall back");

		final nativeResult = GoSurfaceContractRegistry.defaultRegistry()
			.snapshot(ledger(GoTypeShape.Nominal(GoTypeUsageTargetKind.Class, "go.Result", GoImmutableList.fromArray([intShape()]))));
		assertEquals(0, nativeResult.decisionCount, "native go.Result must stay outside the portable registry");

		final firstJson = GoSurfaceContractRegistry.renderJson(admitted);
		final secondJson = GoSurfaceContractRegistry.renderJson(admitted);
		assertEquals(firstJson, secondJson, "registry JSON must be byte deterministic");
		final markdown = GoSurfaceContractRegistry.renderMarkdown(admitted);
		assertTrue(markdown.indexOf("Contract admitted") >= 0, "human report must explain the decision");
		assertTrue(markdown.indexOf("variable_declaration") >= 0, "human report must identify the usage level");
		assertTrue(markdown.indexOf('"kind":"class"') >= 0, "human report must identify the complete typed shape");
		Sys.println("SURFACE_REGISTRY_JSON=" + Base64.encode(Bytes.ofString(firstJson)));
		Sys.println("surface registry macro contract passed");
	}

	static function arrayContract(?surfaceId:GoSurfaceId, ?proofs:GoImmutableList<GoSurfaceProof>,
			?nativeRuntimeRequirements:GoImmutableList<GoHxrtFeatureId>, ?nullShape:Bool, ?unknownShape:Bool, ?emptyBinding:Bool, ?unknownTargetKind:Bool,
			?malformedCollection:String):GoSurfaceContract {
		final bindingName = emptyBinding == true ? "" : malformedCollection == "binding_colon" ? "element:unsafe" : "element";
		final targetKind:GoSurfaceNominalKind = unknownTargetKind == true ? cast "bogus" : GoSurfaceNominalKind.Class;
		final patternParameters = malformedCollection == "pattern_element" ? GoImmutableList.fromArray([null]) : GoImmutableList.fromArray([GoSurfaceTypePattern.Bind(bindingName)]);
		final eligibleShape:GoSurfaceTypePattern = nullShape == true ? null : unknownShape == true ? cast "bogus" : GoSurfaceTypePattern.NominalPattern(targetKind,
			"Array", patternParameters);
		final eligibilityRules = malformedCollection == "rule_list" ? null : malformedCollection == "rule_element" ? GoImmutableList.fromArray([null]) : GoImmutableList.fromArray([
			GoSurfaceEligibilityRule.NoUnknownShapes,
			GoSurfaceEligibilityRule.BindingContainsNoDynamic("element")
		]);
		final nativeImports = malformedCollection == "import_list" ? null : malformedCollection == "import_element" ? GoImmutableList.fromArray([null]) : GoImmutableList.fromArray([]);
		final runtimeRequirements = malformedCollection == "runtime_list" ? null : nativeRuntimeRequirements == null ? GoImmutableList.fromArray([]) : nativeRuntimeRequirements;
		final resolvedProofs = malformedCollection == "proof_list" ? null : malformedCollection == "proof_element" ? GoImmutableList.fromArray([null]) : proofs == null ? GoImmutableList.fromArray([
			{
				proofId: "array-identity-semantic-diff",
				kind: GoSurfaceProofKind.SemanticDiff,
				fixturePath: "test/semantic_diff/array_identity_contract"
			}
		]) : proofs;
		return {
			surfaceId: surfaceId == null ? GoSurfaceId.HaxeArray : surfaceId,
			contractVersion: 1,
			sourceContract: GoSourceContractKind.PortableHaxe,
			sourceSemanticsId: "haxe.array",
			sourceSemanticsVersion: 1,
			sourceSemantics: "ordered" + String.fromCharCode(1) + " mutable Haxe Array with alias-visible mutation",
			eligibleShape: eligibleShape,
			eligibilityRules: eligibilityRules,
			nativeRepresentation: GoNativeRepresentation.GoSlice,
			nativeImports: nativeImports,
			nativeRuntimeRequirements: runtimeRequirements,
			fallbackRepresentation: GoSurfaceFallbackRepresentation.HxrtArray,
			fallbackPolicy: GoSurfaceFallbackPolicy.ReasonedRuntimeRequirement,
			fallbackImports: GoImmutableList.fromArray([]),
			fallbackRuntimeRequirements: GoImmutableList.fromArray([GoHxrtFeatureId.HxrtArray]),
			noHxrtStatus: GoNoHxrtStatus.Conditional,
			proofs: resolvedProofs,
			familySyncExpectation: GoFamilySyncExpectation.TargetLocal,
			familyContractId: "",
			familyContractVersion: 0
		};
	}

	static function assertHasIssue(contracts:Array<GoSurfaceContract>, expected:GoSurfaceValidationCode):Void {
		final issues = GoSurfaceContractRegistry.validate(contracts);
		for (issue in issues) {
			if (issue.code == expected) {
				return;
			}
		}
		throw 'Expected validation issue ${expected}';
	}

	static function assertCreateRejects(contracts:Array<GoSurfaceContract>, expected:GoSurfaceValidationCode):Void {
		try {
			GoSurfaceContractRegistry.create(contracts);
		} catch (error:GoSurfaceContractRegistryException) {
			for (issue in error.issues) {
				if (issue.code == expected) {
					return;
				}
			}
			throw 'Expected create() issue ${expected}';
		}
		throw 'Expected create() to reject ${expected}';
	}

	static function ledger(shape:GoTypeShape):GoTypeUsageLedgerSnapshot {
		final usage:GoTypeUsageEvidence = {
			level: GoTypeUsageLevelId.VariableDeclaration,
			shape: shape
		};
		final module:GoTypeUsageModuleEvidence = {
			module: "Main",
			kind: GoTypeUsageTargetKind.Class,
			location: "Main:1",
			typeUsages: GoImmutableList.fromArray([usage]),
			memberUsages: GoImmutableList.fromArray([]),
			nativeImports: GoImmutableList.fromArray([])
		};
		return {
			schemaVersion: 1,
			source: "registry_test",
			scannerFallback: "none",
			moduleCount: 1,
			typeUsageCount: 1,
			memberUsageCount: 0,
			nativeImportCount: 0,
			capabilityCount: 0,
			modules: GoImmutableList.fromArray([module]),
			capabilities: GoImmutableList.fromArray([])
		};
	}

	static function arrayShape(element:GoTypeShape):GoTypeShape {
		return GoTypeShape.Nominal(GoTypeUsageTargetKind.Class, "Array", GoImmutableList.fromArray([element]));
	}

	static function optionShape(element:GoTypeShape):GoTypeShape {
		return GoTypeShape.Nominal(GoTypeUsageTargetKind.Enum, "reflaxe.std.Option", GoImmutableList.fromArray([element]));
	}

	static function resultShape(value:GoTypeShape, error:GoTypeShape):GoTypeShape {
		return GoTypeShape.Nominal(GoTypeUsageTargetKind.Enum, "reflaxe.std.Result", GoImmutableList.fromArray([value, error]));
	}

	static function intShape():GoTypeShape {
		return GoTypeShape.Nominal(GoTypeUsageTargetKind.Abstract, "StdTypes.Int", GoImmutableList.fromArray([]));
	}

	static function stringShape():GoTypeShape {
		return GoTypeShape.Nominal(GoTypeUsageTargetKind.Class, "String", GoImmutableList.fromArray([]));
	}

	static function assertTrue(condition:Bool, message:String):Void {
		if (!condition) {
			throw message;
		}
	}

	static function assertEquals<T>(expected:T, actual:T, message:String):Void {
		if (expected != actual) {
			throw '$message (expected ${expected}, got ${actual})';
		}
	}
}
#end
