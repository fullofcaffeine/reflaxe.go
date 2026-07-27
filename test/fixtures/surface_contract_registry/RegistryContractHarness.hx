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

		final productionArray = GoSurfaceContractRegistry.defaultRegistry().snapshot(ledger(arrayShape(intShape())));
		assertEquals(5, productionArray.catalogCount, "production must include the three proven collection contracts");
		assertEquals(GoSurfaceDecisionOutcome.Admitted, productionArray.decisions.at(0).outcome,
			"fully typed portable Array must admit a shared slice-backed carrier");
		assertEquals(GoNativeRepresentation.GoSlice, productionArray.decisions.at(0).selectedRepresentation,
			"portable Array must select the semantic Go slice carrier");
		assertEquals(GoSurfaceFallbackRepresentation.HxrtArray, productionArray.decisions.at(0).fallbackRepresentation,
			"portable Array must retain its shared hxrt fallback");
		assertEquals(GoHxrtFeatureId.HxrtArray, productionArray.decisions.at(0).fallbackRuntimeRequirements.at(0),
			"Array fallback must report its exact runtime requirement");

		final nestedArray = GoSurfaceContractRegistry.defaultRegistry().snapshot(ledger(arrayShape(arrayShape(intShape()))));
		assertEquals(GoSurfaceDecisionOutcome.Admitted, nestedArray.decisions.at(0).outcome,
			"nested typed Array must remain eligible because each level preserves shared identity");

		final genericArray = GoSurfaceContractRegistry.defaultRegistry().snapshot(ledger(arrayShape(GoTypeShape.TypeParameter("T"))));
		assertEquals(GoSurfaceDecisionOutcome.Rejected, genericArray.decisions.at(0).outcome,
			"a named generic Array element has no proven concrete carrier and must fall back");
		assertTrue(genericArray.decisions.at(0).detail.indexOf("binding_has_proven_collection_carrier:element") >= 0,
			"generic Array fallback must report the opaque carrier boundary");

		final dynamicArray = GoSurfaceContractRegistry.defaultRegistry().snapshot(ledger(arrayShape(GoTypeShape.DynamicShape(null))));
		assertEquals(GoSurfaceDecisionOutcome.Rejected, dynamicArray.decisions.at(0).outcome,
			"Dynamic Array elements can hide unproven aliases and must fall back");
		assertEquals(GoSurfaceDecisionReason.EligibilityRejected, dynamicArray.decisions.at(0).reason,
			"Dynamic Array fallback must name the eligibility boundary");
		assertEquals(GoSurfaceFallbackRepresentation.HxrtArray, dynamicArray.decisions.at(0).fallbackRepresentation,
			"rejected Array must retain shared Haxe semantics");

		final unknownArray = GoSurfaceContractRegistry.defaultRegistry()
			.snapshot(ledger(arrayShape(GoTypeShape.UnknownShape(GoUnknownTypeShapeReason.Missing))));
		assertEquals(GoSurfaceDecisionOutcome.Rejected, unknownArray.decisions.at(0).outcome, "unresolved Array elements must fail closed");

		final typedefArray = GoSurfaceContractRegistry.defaultRegistry().snapshot(ledger(arrayShape(typedefShape("Main.HiddenDynamic"))));
		assertEquals(GoSurfaceDecisionOutcome.Rejected, typedefArray.decisions.at(0).outcome,
			"typedef storage is opaque in the ledger and must fail closed even without a visible Dynamic node");
		assertTrue(typedefArray.decisions.at(0).detail.indexOf("binding_has_proven_collection_carrier:element") >= 0,
			"typedef Array fallback must report the unproven carrier rule");

		final abstractArray = GoSurfaceContractRegistry.defaultRegistry().snapshot(ledger(arrayShape(abstractShape("Main.HiddenDynamicAbstract"))));
		assertEquals(GoSurfaceDecisionOutcome.Rejected, abstractArray.decisions.at(0).outcome,
			"user abstract storage is opaque in the ledger and must fail closed");

		final stringMap = GoSurfaceContractRegistry.defaultRegistry().snapshot(ledger(stringMapShape(intShape())));
		assertEquals(GoSurfaceDecisionOutcome.Admitted, stringMap.decisions.at(0).outcome,
			"StringMap with a typed value must admit a fixed-string-key map carrier");
		assertEquals(GoNativeRepresentation.GoMap, stringMap.decisions.at(0).selectedRepresentation, "StringMap must select the semantic Go map carrier");
		assertEquals(GoHxrtFeatureId.HxrtMapString, stringMap.decisions.at(0).fallbackRuntimeRequirements.at(0),
			"StringMap fallback must report the string-keyed runtime");

		final intMap = GoSurfaceContractRegistry.defaultRegistry().snapshot(ledger(intMapShape(arrayShape(intShape()))));
		assertEquals(GoSurfaceDecisionOutcome.Admitted, intMap.decisions.at(0).outcome,
			"IntMap with a nested typed value must admit a fixed-int-key map carrier");
		assertEquals(GoHxrtFeatureId.HxrtMapInt, intMap.decisions.at(0).fallbackRuntimeRequirements.at(0),
			"IntMap fallback must report the integer-keyed runtime");

		final dynamicStringMap = GoSurfaceContractRegistry.defaultRegistry().snapshot(ledger(stringMapShape(GoTypeShape.DynamicShape(null))));
		assertEquals(GoSurfaceDecisionOutcome.Rejected, dynamicStringMap.decisions.at(0).outcome, "Dynamic map values must fall back");
		assertEquals(GoSurfaceFallbackRepresentation.HxrtMap, dynamicStringMap.decisions.at(0).fallbackRepresentation,
			"rejected StringMap must retain portable map behavior");

		final nestedDynamicStringMap = GoSurfaceContractRegistry.defaultRegistry()
			.snapshot(ledger(stringMapShape(arrayShape(GoTypeShape.DynamicShape(null)))));
		assertEquals(GoSurfaceDecisionOutcome.Rejected, nestedDynamicStringMap.decisions.at(0).outcome,
			"Dynamic hidden inside a nested map value must fail recursively");

		final typedefStringMap = GoSurfaceContractRegistry.defaultRegistry().snapshot(ledger(stringMapShape(typedefShape("Main.HiddenDynamic"))));
		assertEquals(GoSurfaceDecisionOutcome.Rejected, typedefStringMap.decisions.at(0).outcome,
			"StringMap must not admit a typedef whose underlying storage is absent from the ledger");

		final dynamicIntMap = GoSurfaceContractRegistry.defaultRegistry().snapshot(ledger(intMapShape(GoTypeShape.DynamicShape(null))));
		assertEquals(GoSurfaceDecisionOutcome.Rejected, dynamicIntMap.decisions.at(0).outcome, "Dynamic IntMap values must fall back");
		assertEquals(GoSurfaceDecisionReason.EligibilityRejected, dynamicIntMap.decisions.at(0).reason,
			"Dynamic IntMap fallback must report its typed eligibility reason");
		assertEquals(GoSurfaceFallbackRepresentation.HxrtMap, dynamicIntMap.decisions.at(0).fallbackRepresentation,
			"rejected IntMap must retain portable map behavior");
		assertEquals(GoHxrtFeatureId.HxrtMapInt, dynamicIntMap.decisions.at(0).fallbackRuntimeRequirements.at(0),
			"rejected IntMap must report the integer-keyed runtime");

		final nestedDynamicIntMap = GoSurfaceContractRegistry.defaultRegistry().snapshot(ledger(intMapShape(arrayShape(GoTypeShape.DynamicShape(null)))));
		assertEquals(GoSurfaceDecisionOutcome.Rejected, nestedDynamicIntMap.decisions.at(0).outcome,
			"Dynamic hidden inside a nested IntMap value must fail recursively");

		final abstractIntMap = GoSurfaceContractRegistry.defaultRegistry().snapshot(ledger(intMapShape(abstractShape("Main.HiddenDynamicAbstract"))));
		assertEquals(GoSurfaceDecisionOutcome.Rejected, abstractIntMap.decisions.at(0).outcome,
			"IntMap must not admit an abstract whose underlying storage is absent from the ledger");

		final objectMap = GoSurfaceContractRegistry.defaultRegistry().snapshot(ledger(objectMapShape(classShape("Main.Box"), stringShape())));
		assertEquals(GoSurfaceDecisionOutcome.Rejected, objectMap.decisions.at(0).outcome,
			"ObjectMap must remain unadmitted until object identity has a native proof");
		assertEquals(GoSurfaceDecisionReason.ContractMissing, objectMap.decisions.at(0).reason,
			"unsafe object-key equality must be reported as a missing contract");

		final guardedObjectMap = GoSurfaceContractRegistry.create([objectMapContract()])
			.snapshot(ledger(objectMapShape(classShape("Main.Box"), stringShape())));
		assertEquals(GoSurfaceDecisionOutcome.Rejected, guardedObjectMap.decisions.at(0).outcome,
			"the fixed-key comparability rule must reject object-identity maps");
		assertEquals(GoSurfaceDecisionReason.EligibilityRejected, guardedObjectMap.decisions.at(0).reason,
			"object-key rejection must identify eligibility rather than silently stringify the key");
		assertTrue(guardedObjectMap.decisions.at(0).detail.indexOf("surface_has_fixed_go_comparable_map_key") >= 0,
			"the report must expose the typed fixed-key rejection rule");

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
		final nativeMap = GoSurfaceContractRegistry.defaultRegistry()
			.snapshot(ledger(GoTypeShape.Nominal(GoTypeUsageTargetKind.Class, "go.Map", GoImmutableList.fromArray([arrayShape(intShape()), intShape()]))));
		assertEquals(0, nativeMap.decisionCount, "target-native go.Map and its Std.string fallback must stay outside portable admission");

		final productionReport = GoSurfaceContractRegistry.defaultRegistry().snapshot(ledgerMany([
			arrayShape(intShape()),
			arrayShape(GoTypeShape.DynamicShape(null)),
			stringMapShape(intShape()),
			stringMapShape(GoTypeShape.DynamicShape(null)),
			intMapShape(arrayShape(intShape())),
			intMapShape(GoTypeShape.DynamicShape(null)),
			objectMapShape(classShape("Main.Box"), stringShape())
		]));
		final firstJson = GoSurfaceContractRegistry.renderJson(productionReport);
		final secondJson = GoSurfaceContractRegistry.renderJson(productionReport);
		assertEquals(firstJson, secondJson, "registry JSON must be byte deterministic");
		final markdown = GoSurfaceContractRegistry.renderMarkdown(productionReport);
		assertTrue(markdown.indexOf("Contract admitted") >= 0, "human report must explain the decision");
		assertTrue(markdown.indexOf("Eligibility rule rejected") >= 0, "human report must explain typed collection fallback");
		assertTrue(markdown.indexOf("Known portable surface has no admitted contract") >= 0, "human report must explain the ObjectMap identity boundary");
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

	static function objectMapContract():GoSurfaceContract {
		return {
			surfaceId: GoSurfaceId.HaxeObjectMap,
			contractVersion: 1,
			sourceContract: GoSourceContractKind.PortableHaxe,
			sourceSemanticsId: "test.haxe.object-map",
			sourceSemanticsVersion: 1,
			sourceSemantics: "Test-only object-key map contract that must fail typed fixed-key comparability.",
			eligibleShape: GoSurfaceTypePattern.NominalPattern(GoSurfaceNominalKind.Class, "haxe.ds.ObjectMap",
				GoImmutableList.fromArray([GoSurfaceTypePattern.Bind("key"), GoSurfaceTypePattern.Bind("value")])),
			eligibilityRules: GoImmutableList.fromArray([
				GoSurfaceEligibilityRule.NoUnknownShapes,
				GoSurfaceEligibilityRule.SurfaceHasFixedGoComparableMapKey,
				GoSurfaceEligibilityRule.BindingContainsNoDynamic("value")
			]),
			nativeRepresentation: GoNativeRepresentation.GoMap,
			nativeImports: GoImmutableList.fromArray([]),
			nativeRuntimeRequirements: GoImmutableList.fromArray([]),
			fallbackRepresentation: GoSurfaceFallbackRepresentation.HxrtMap,
			fallbackPolicy: GoSurfaceFallbackPolicy.ReasonedRuntimeRequirement,
			fallbackImports: GoImmutableList.fromArray([]),
			fallbackRuntimeRequirements: GoImmutableList.fromArray([GoHxrtFeatureId.HxrtMapObject]),
			noHxrtStatus: GoNoHxrtStatus.Conditional,
			proofs: GoImmutableList.fromArray([
				{
					proofId: "object-map-identity-semantic-diff",
					kind: GoSurfaceProofKind.SemanticDiff,
					fixturePath: "test/semantic_diff/ds_maps_list_contract"
				}
			]),
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
		return ledgerMany([shape]);
	}

	static function ledgerMany(shapes:Array<GoTypeShape>):GoTypeUsageLedgerSnapshot {
		final usages = [
			for (shape in shapes)
				{
					level: GoTypeUsageLevelId.VariableDeclaration,
					shape: shape
				}
		];
		final module:GoTypeUsageModuleEvidence = {
			module: "Main",
			kind: GoTypeUsageTargetKind.Class,
			location: "Main:1",
			typeUsages: GoImmutableList.fromArray(usages),
			memberUsages: GoImmutableList.fromArray([]),
			nativeImports: GoImmutableList.fromArray([])
		};
		return {
			schemaVersion: 1,
			source: "registry_test",
			scannerFallback: "none",
			moduleCount: 1,
			typeUsageCount: usages.length,
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

	static function stringMapShape(value:GoTypeShape):GoTypeShape {
		return GoTypeShape.Nominal(GoTypeUsageTargetKind.Class, "haxe.ds.StringMap", GoImmutableList.fromArray([value]));
	}

	static function intMapShape(value:GoTypeShape):GoTypeShape {
		return GoTypeShape.Nominal(GoTypeUsageTargetKind.Class, "haxe.ds.IntMap", GoImmutableList.fromArray([value]));
	}

	static function objectMapShape(key:GoTypeShape, value:GoTypeShape):GoTypeShape {
		return GoTypeShape.Nominal(GoTypeUsageTargetKind.Class, "haxe.ds.ObjectMap", GoImmutableList.fromArray([key, value]));
	}

	static function classShape(path:String):GoTypeShape {
		return GoTypeShape.Nominal(GoTypeUsageTargetKind.Class, path, GoImmutableList.fromArray([]));
	}

	static function typedefShape(path:String):GoTypeShape {
		return GoTypeShape.Nominal(GoTypeUsageTargetKind.Typedef, path, GoImmutableList.fromArray([]));
	}

	static function abstractShape(path:String):GoTypeShape {
		return GoTypeShape.Nominal(GoTypeUsageTargetKind.Abstract, path, GoImmutableList.fromArray([]));
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
