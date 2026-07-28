package reflaxe.go.compiler;

#if (macro || reflaxe_runtime)
import reflaxe.go.compiler.GoHxrtFeatureAnalyzer.GoHxrtFeatureId;
import reflaxe.go.compiler.GoSurfaceContractRegistry.GoNativeRepresentation;
import reflaxe.go.compiler.GoSurfaceContractRegistry.GoSurfaceContractRegistrySnapshot;
import reflaxe.go.compiler.GoSurfaceContractRegistry.GoSurfaceDecision;
import reflaxe.go.compiler.GoSurfaceContractRegistry.GoSurfaceDecisionOutcome;
import reflaxe.go.compiler.GoSurfaceContractRegistry.GoSurfaceDecisionReason;
import reflaxe.go.compiler.GoSurfaceContractRegistry.GoSurfaceId;
import reflaxe.go.compiler.GoSurfaceContractRegistry.GoSurfaceImportRequirement;
import reflaxe.go.compiler.GoTypeUsageLedger.GoImmutableList;
import reflaxe.go.compiler.GoTypeUsageLedger.GoTypeShape;
import reflaxe.go.compiler.GoTypeUsageLedger.GoTypeUsageLevelId;

/**
	Why
	Optimizer, import, and runtime code must agree on whether one observed
	portable surface uses its admitted Go carrier or its semantics-safe fallback.

	What
	Names the three possible planner outcomes: a registry-admitted native carrier,
	a registered fallback, or unchanged legacy lowering when no fallback contract
	exists.

	How
	`GoSurfacePlanner.plan(...)` is the only constructor of production decisions.
**/
enum abstract GoSurfacePlanSelection(String) to String {
	var Native = "native";
	var Fallback = "fallback";
	var Existing = "existing";
}

/**
	Why
	A selected representation without a stable reason would make reports
	insufficient to audit profile or define bypasses.

	What
	Explains whether the registry admitted an active carrier, rejected the shape,
	left an admitted carrier pending activation, or had no registered fallback.

	How
	The planner derives these values only from its immutable registry decision and
	the closed activated-carrier table below.
**/
enum abstract GoSurfacePlanReason(String) to String {
	var RegistryAdmitted = "registry_admitted";
	var RegistryRejected = "registry_rejected";
	var CarrierNotActivated = "carrier_not_activated";
	var NoRegisteredFallback = "no_registered_fallback";
}

/**
	One immutable portable-surface choice consumed by all downstream planners.

	Why
	The registry decision proves eligibility, but consumers also need the actual
	representation, fallback explanation, imports, and runtime cost selected for
	this compiler version.

	What
	Preserves the used typed shape and contract decision alongside the selected
	consequences.

	How
	The fields contain only typed closed values, stable strings, and deeply
	read-only lists. Locations come from the typed ledger and never contain
	absolute paths.
**/
typedef GoSurfacePlanDecision = {
	final module:String;
	final location:String;
	final usageLevel:GoTypeUsageLevelId;
	final surfaceId:GoSurfaceId;
	final usedType:GoTypeShape;
	final contractVersion:Int;
	final eligibilityOutcome:GoSurfaceDecisionOutcome;
	final eligibilityReason:GoSurfaceDecisionReason;
	final eligibilityDetail:String;
	final selection:GoSurfacePlanSelection;
	final selectionReason:GoSurfacePlanReason;
	final selectedRepresentation:Null<String>;
	final fallbackReason:Null<String>;
	final imports:GoImmutableList<GoSurfaceImportRequirement>;
	final runtimeRequirements:GoImmutableList<GoHxrtFeatureId>;
}

/**
	The immutable build-wide surface plan.

	Why
	Recomputing a choice in the optimizer, import collector, and runtime slicer
	would recreate the profile/define drift this authority is meant to remove.

	What
	Contains every sorted per-shape decision plus the deduplicated import and
	runtime consequences downstream consumers must consider.

	How
	`CompilationContext` receives exactly one snapshot before lowering and never
	replaces it.
**/
typedef GoSurfacePlanSnapshot = {
	final schemaVersion:Int;
	final authority:String;
	final nativeSpecializationPolicy:String;
	final nativeSpecializationPolicySource:String;
	final decisionCount:Int;
	final nativeCount:Int;
	final fallbackCount:Int;
	final existingCount:Int;
	final requiredImports:GoImmutableList<GoSurfaceImportRequirement>;
	final requiredRuntimeFeatures:GoImmutableList<GoHxrtFeatureId>;
	final decisions:GoImmutableList<GoSurfacePlanDecision>;
}

/**
	The single portable-representation planner shared by optimizer, imports, and
	runtime selection.

	Why
	A compatibility preset or eager specialization policy may choose defaults for
	explicit native APIs, but neither is semantic evidence for portable Haxe
	representations. Only a typed registry decision can admit those.

	What
	Combines `GoBuildContext` policy provenance with every immutable registry
	decision and selects the carrier implemented by this compiler version.

	How
	Already-implemented String and Bytes carriers may be selected after admission.
	Iterator, Array, map, Option, and Result remain on their named fallbacks until
	their distinct native representations and promotion gates land in `.7.7`.
**/
class GoSurfacePlanner {
	public static inline final SCHEMA_VERSION = 1;
	public static inline final AUTHORITY = "go_build_context_plus_typed_registry_decision";

	public static function plan(buildContext:GoBuildContext, registry:GoSurfaceContractRegistrySnapshot):GoSurfacePlanSnapshot {
		var source = registry == null ? GoSurfaceContractRegistry.emptySnapshot() : registry;
		var decisions = new Array<GoSurfacePlanDecision>();
		var requiredImportsByKey = new Map<String, GoSurfaceImportRequirement>();
		var requiredRuntimeById = new Map<String, GoHxrtFeatureId>();
		var nativeCount = 0;
		var fallbackCount = 0;
		var existingCount = 0;

		for (registryDecision in source.decisions) {
			var decision = select(registryDecision);
			decisions.push(decision);
			switch (decision.selection) {
				case GoSurfacePlanSelection.Native:
					nativeCount++;
				case GoSurfacePlanSelection.Fallback:
					fallbackCount++;
				case GoSurfacePlanSelection.Existing:
					existingCount++;
			}
			for (requirement in decision.imports) {
				requiredImportsByKey.set(importKey(requirement), {
					path: requirement.path,
					reason: requirement.reason
				});
			}
			for (feature in decision.runtimeRequirements) {
				requiredRuntimeById.set(feature, feature);
			}
		}

		decisions.sort((a, b) -> compareStrings(decisionKey(a), decisionKey(b)));
		var requiredImports = [for (requirement in requiredImportsByKey) requirement];
		requiredImports.sort((a, b) -> compareStrings(importKey(a), importKey(b)));
		var requiredRuntimeFeatures = [for (feature in requiredRuntimeById) feature];
		requiredRuntimeFeatures.sort((a, b) -> compareStrings(a, b));

		return {
			schemaVersion: SCHEMA_VERSION,
			authority: AUTHORITY,
			nativeSpecializationPolicy: buildContext.nativeSpecializationPolicy.label(),
			nativeSpecializationPolicySource: buildContext.nativeSpecializationPolicySource.label(),
			decisionCount: decisions.length,
			nativeCount: nativeCount,
			fallbackCount: fallbackCount,
			existingCount: existingCount,
			requiredImports: GoImmutableList.fromArray(requiredImports),
			requiredRuntimeFeatures: GoImmutableList.fromArray(requiredRuntimeFeatures),
			decisions: GoImmutableList.fromArray(decisions)
		};
	}

	/** Empty authority used only before a real build context and typed registry exist. */
	public static function emptySnapshot():GoSurfacePlanSnapshot {
		return {
			schemaVersion: SCHEMA_VERSION,
			authority: AUTHORITY,
			nativeSpecializationPolicy: "",
			nativeSpecializationPolicySource: "",
			decisionCount: 0,
			nativeCount: 0,
			fallbackCount: 0,
			existingCount: 0,
			requiredImports: GoImmutableList.fromArray([]),
			requiredRuntimeFeatures: GoImmutableList.fromArray([]),
			decisions: GoImmutableList.fromArray([])
		};
	}

	/**
		Answers whether every observed use of one shape-invariant surface selected
		the same admitted native carrier.

		This narrow query currently gates String fast paths. Future generic carrier
		lowerings must match the exact `usedType` decision instead of widening this
		build-wide predicate to parameterized surfaces. With no observed decision,
		there is no governed source use to veto the compiler's established String
		helper policy, so the universal predicate remains true.
	**/
	public static function allObservedUsesSelectNative(snapshot:GoSurfacePlanSnapshot, surfaceId:GoSurfaceId, representation:GoNativeRepresentation):Bool {
		if (snapshot == null) {
			return false;
		}
		for (decision in snapshot.decisions) {
			if (decision.surfaceId != surfaceId) {
				continue;
			}
			if (decision.selection != GoSurfacePlanSelection.Native || decision.selectedRepresentation != representation) {
				return false;
			}
		}
		return true;
	}

	static function select(decision:GoSurfaceDecision):GoSurfacePlanDecision {
		if (decision.outcome == GoSurfaceDecisionOutcome.Admitted
			&& decision.selectedRepresentation != null
			&& isActivatedCarrier(decision.surfaceId, decision.selectedRepresentation)) {
			return copyDecision(decision, GoSurfacePlanSelection.Native, GoSurfacePlanReason.RegistryAdmitted, decision.selectedRepresentation, null,
				decision.nativeImports, decision.runtimeRequirements);
		}

		if (decision.fallbackRepresentation != null) {
			var reason = decision.outcome == GoSurfaceDecisionOutcome.Admitted ? GoSurfacePlanReason.CarrierNotActivated : GoSurfacePlanReason.RegistryRejected;
			var fallbackReason = decision.outcome == GoSurfaceDecisionOutcome.Admitted ? "The native carrier is admitted but not activated; this compiler keeps the semantics-safe fallback until its independent promotion gate lands." : decision.reason
				+ ": "
				+ decision.detail;
			return copyDecision(decision, GoSurfacePlanSelection.Fallback, reason, decision.fallbackRepresentation, fallbackReason, decision.fallbackImports,
				decision.fallbackRuntimeRequirements);
		}

		return copyDecision(decision, GoSurfacePlanSelection.Existing, GoSurfacePlanReason.NoRegisteredFallback, null,
			decision.reason + ": " + decision.detail, GoImmutableList.fromArray([]), GoImmutableList.fromArray([]));
	}

	/**
		Closed activation table for representations already emitted and proved by
		generated-shape fixtures.

		Adding a carrier here is a default behavior change. `.7.7` must first land
		its implementation, paired rollback fixture, semantics, runtime, and
		performance evidence.
	**/
	static function isActivatedCarrier(surfaceId:GoSurfaceId, representation:GoNativeRepresentation):Bool {
		return switch (surfaceId) {
			case GoSurfaceId.HaxeString:
				representation == GoNativeRepresentation.GoString;
			case GoSurfaceId.HaxeBytes:
				representation == GoNativeRepresentation.GoBytes;
			case GoSurfaceId.HaxeArray | GoSurfaceId.HaxeStringMap | GoSurfaceId.HaxeIntMap | GoSurfaceId.HaxeObjectMap | GoSurfaceId.HaxeIterator | GoSurfaceId.PortableOption | GoSurfaceId.PortableResult | GoSurfaceId.HaxeFunction:
				false;
			case _:
				false;
		};
	}

	static function copyDecision(decision:GoSurfaceDecision, selection:GoSurfacePlanSelection, selectionReason:GoSurfacePlanReason,
			selectedRepresentation:Null<String>, fallbackReason:Null<String>, imports:GoImmutableList<GoSurfaceImportRequirement>,
			runtimeRequirements:GoImmutableList<GoHxrtFeatureId>):GoSurfacePlanDecision {
		return {
			module: decision.module,
			location: decision.location,
			usageLevel: decision.usageLevel,
			surfaceId: decision.surfaceId,
			usedType: decision.shape,
			contractVersion: decision.contractVersion,
			eligibilityOutcome: decision.outcome,
			eligibilityReason: decision.reason,
			eligibilityDetail: decision.detail,
			selection: selection,
			selectionReason: selectionReason,
			selectedRepresentation: selectedRepresentation,
			fallbackReason: fallbackReason,
			imports: copyImports(imports),
			runtimeRequirements: copyRuntimeRequirements(runtimeRequirements)
		};
	}

	static function copyImports(imports:GoImmutableList<GoSurfaceImportRequirement>):GoImmutableList<GoSurfaceImportRequirement> {
		var copied:Array<GoSurfaceImportRequirement> = [
			for (requirement in imports)
				{
					path: requirement.path,
					reason: requirement.reason
				}
		];
		copied.sort((a, b) -> compareStrings(importKey(a), importKey(b)));
		return GoImmutableList.fromArray(copied);
	}

	static function copyRuntimeRequirements(requirements:GoImmutableList<GoHxrtFeatureId>):GoImmutableList<GoHxrtFeatureId> {
		var copied = [for (feature in requirements) feature];
		copied.sort((a, b) -> compareStrings(a, b));
		return GoImmutableList.fromArray(copied);
	}

	static function decisionKey(decision:GoSurfacePlanDecision):String {
		return decision.module + "\n" + decision.location + "\n" + decision.usageLevel + "\n" + decision.surfaceId + "\n"
			+ GoTypeUsageLedger.renderShapeJson(decision.usedType);
	}

	static function importKey(requirement:GoSurfaceImportRequirement):String {
		return requirement.path + "\n" + requirement.reason;
	}

	static function compareStrings(a:String, b:String):Int {
		return a < b ? -1 : (a > b ? 1 : 0);
	}
}
#end
