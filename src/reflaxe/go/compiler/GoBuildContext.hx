package reflaxe.go.compiler;

import reflaxe.go.GoProfile;
import reflaxe.go.RawNativeMode;

/**
	Why
	Define parsing in separate macro, lowering, runtime, and report stages would
	allow policy drift and accidental profile-shaped semantic branches.

	What
	Immutable build-level configuration for compatibility preset, native policy
	axes, strictness, optimizer/planner, runtime, and report settings.

	How
	`GoBuildContextResolver` constructs it once; boundary discovery returns a
	copy with deterministic module names while preserving policy provenance.
**/
class GoBuildContext {
	/** Compatibility selector retained for public define and report stability. */
	public final profile:GoProfile;

	public final policyPreset:GoPolicyPreset;
	public final semanticBoundarySource:GoSemanticBoundarySource;
	public final nativeAuthorityPolicy:GoNativeAuthorityPolicy;
	public final nativeAuthorityPolicySource:GoPolicyResolutionSource;
	public final nativeSpecializationPolicy:GoNativeSpecializationPolicy;
	public final nativeSpecializationPolicySource:GoPolicyResolutionSource;
	public final nativeFallbackPolicy:GoNativeFallbackPolicy;
	public final nativeFallbackPolicySource:GoPolicyResolutionSource;
	public final goModuleName:String;
	public final rawNativeMode:RawNativeMode;
	public final emitLineDirectives:Bool;
	public final strictExamples:Bool;
	public final strictUserBoundaryPolicy:String;
	public final strictUserBoundaries:Bool;
	public final metalFallbackAllowed:Bool;
	public final metalContractHardError:Bool;
	public final hxrtForceFullCopy:Bool;
	public final hxrtFeaturesDefinePresent:Bool;
	public final hxrtNoFeatureInfer:Bool;
	public final hxrtManualFeatures:Array<String>;
	public final contractReportEnabled:Bool;
	public final runtimePlanReportEnabled:Bool;
	public final optimizerPlanReportEnabled:Bool;
	public final typeUsageReportEnabled:Bool;
	public final surfaceContractReportEnabled:Bool;
	public final autoLoweringMode:GoAutoLoweringMode;
	public final optimizationPreset:String;
	public final portableStringFastpathEnabled:Bool;
	public final portableConcurrencyFastpathEnabled:Bool;
	public final nativeStackTraceEnabled:Bool;
	public final nativeBoundaryModules:Array<String>;

	/** Compatibility alias for report and macro consumers using the old name. */
	public final metalLaneModules:Array<String>;

	public function new(profile:GoProfile, policyResolution:GoPolicyResolution, goModuleName:String, rawNativeMode:RawNativeMode, emitLineDirectives:Bool,
			strictExamples:Bool, strictUserBoundaryPolicy:String, strictUserBoundaries:Bool, metalFallbackAllowed:Bool, metalContractHardError:Bool,
			hxrtForceFullCopy:Bool, hxrtFeaturesDefinePresent:Bool, hxrtNoFeatureInfer:Bool, hxrtManualFeatures:Array<String>, contractReportEnabled:Bool,
			runtimePlanReportEnabled:Bool, optimizerPlanReportEnabled:Bool, typeUsageReportEnabled:Bool, surfaceContractReportEnabled:Bool,
			autoLoweringMode:GoAutoLoweringMode, optimizationPreset:String, portableStringFastpathEnabled:Bool, portableConcurrencyFastpathEnabled:Bool,
			nativeStackTraceEnabled:Bool, nativeBoundaryModules:Array<String>) {
		this.profile = profile;
		var resolvedPolicy = policyResolution == null ? legacyPolicyResolution(profile) : policyResolution;
		this.policyPreset = resolvedPolicy.preset;
		this.semanticBoundarySource = resolvedPolicy.semanticBoundarySource;
		this.nativeAuthorityPolicy = resolvedPolicy.nativeAuthority;
		this.nativeAuthorityPolicySource = resolvedPolicy.nativeAuthoritySource;
		this.nativeSpecializationPolicy = resolvedPolicy.nativeSpecialization;
		this.nativeSpecializationPolicySource = resolvedPolicy.nativeSpecializationSource;
		this.nativeFallbackPolicy = resolvedPolicy.nativeFallback;
		this.nativeFallbackPolicySource = resolvedPolicy.nativeFallbackSource;
		this.goModuleName = normalizeGoModuleName(goModuleName);
		this.rawNativeMode = rawNativeMode == null ? RawNativeMode.Interp : rawNativeMode;
		this.emitLineDirectives = emitLineDirectives == true;
		this.strictExamples = strictExamples == true;
		this.strictUserBoundaryPolicy = normalizeStrictUserBoundaryPolicy(strictUserBoundaryPolicy);
		this.strictUserBoundaries = strictUserBoundaries == true;
		this.metalFallbackAllowed = metalFallbackAllowed == true;
		this.metalContractHardError = metalContractHardError == true;
		this.hxrtForceFullCopy = hxrtForceFullCopy == true;
		this.hxrtFeaturesDefinePresent = hxrtFeaturesDefinePresent == true;
		this.hxrtNoFeatureInfer = hxrtNoFeatureInfer == true;
		this.hxrtManualFeatures = sortedUniqueLowercase(hxrtManualFeatures);
		this.contractReportEnabled = contractReportEnabled == true;
		this.runtimePlanReportEnabled = runtimePlanReportEnabled == true;
		this.optimizerPlanReportEnabled = optimizerPlanReportEnabled == true;
		this.typeUsageReportEnabled = typeUsageReportEnabled == true;
		this.surfaceContractReportEnabled = surfaceContractReportEnabled == true;
		this.autoLoweringMode = autoLoweringMode == null ? GoAutoLoweringMode.Off : autoLoweringMode;
		this.optimizationPreset = normalizeOptimizationPreset(optimizationPreset);
		this.portableStringFastpathEnabled = portableStringFastpathEnabled == true;
		this.portableConcurrencyFastpathEnabled = portableConcurrencyFastpathEnabled == true;
		this.nativeStackTraceEnabled = nativeStackTraceEnabled == true;
		this.nativeBoundaryModules = sortedUnique(nativeBoundaryModules);
		this.metalLaneModules = this.nativeBoundaryModules.copy();
	}

	/** Compatibility predicate. New behavior must branch on a typed policy axis. */
	public inline function isMetalContract():Bool {
		return usesMetalCompatibilityPreset();
	}

	public inline function usesMetalCompatibilityPreset():Bool {
		return policyPreset == GoPolicyPreset.MetalCompatibility;
	}

	public inline function hasExplicitNativeAuthority():Bool {
		return nativeAuthorityPolicy == GoNativeAuthorityPolicy.Explicit;
	}

	public inline function usesEagerNativeSpecialization():Bool {
		return nativeSpecializationPolicy == GoNativeSpecializationPolicy.Eager;
	}

	public inline function requiresNativeFallbackError():Bool {
		return nativeFallbackPolicy == GoNativeFallbackPolicy.Error;
	}

	public inline function isHxrtSelectiveEnabled():Bool {
		return hxrtFeaturesDefinePresent || hxrtNoFeatureInfer;
	}

	public function withNativeBoundaryModules(nativeBoundaryModules:Array<String>):GoBuildContext {
		return new GoBuildContext(profile, currentPolicyResolution(), goModuleName, rawNativeMode, emitLineDirectives, strictExamples,
			strictUserBoundaryPolicy, strictUserBoundaries, metalFallbackAllowed, metalContractHardError, hxrtForceFullCopy, hxrtFeaturesDefinePresent,
			hxrtNoFeatureInfer, hxrtManualFeatures, contractReportEnabled, runtimePlanReportEnabled, optimizerPlanReportEnabled, typeUsageReportEnabled,
			surfaceContractReportEnabled, autoLoweringMode, optimizationPreset, portableStringFastpathEnabled, portableConcurrencyFastpathEnabled,
			nativeStackTraceEnabled, nativeBoundaryModules);
	}

	/** Compatibility alias; new code should use `withNativeBoundaryModules`. */
	public function withMetalLaneModules(metalLaneModules:Array<String>):GoBuildContext {
		return withNativeBoundaryModules(metalLaneModules);
	}

	public static function legacyDefaults(profile:GoProfile, ?goModuleName:String, ?rawNativeMode:RawNativeMode, ?emitLineDirectives:Bool):GoBuildContext {
		return new GoBuildContext(profile, legacyPolicyResolution(profile), normalizeGoModuleName(goModuleName),
			rawNativeMode == null ? RawNativeMode.Interp : rawNativeMode, emitLineDirectives == true, false, "auto", false, false, false, false, false, false,
			[], false, false, false, false, false, GoAutoLoweringMode.Off, "portable_fast", true, true, false, []);
	}

	function currentPolicyResolution():GoPolicyResolution {
		return {
			preset: policyPreset,
			semanticBoundarySource: semanticBoundarySource,
			nativeAuthority: nativeAuthorityPolicy,
			nativeAuthoritySource: nativeAuthorityPolicySource,
			nativeSpecialization: nativeSpecializationPolicy,
			nativeSpecializationSource: nativeSpecializationPolicySource,
			nativeFallback: nativeFallbackPolicy,
			nativeFallbackSource: nativeFallbackPolicySource
		};
	}

	static function legacyPolicyResolution(profile:GoProfile):GoPolicyResolution {
		var preset = GoPolicyPreset.fromLegacyProfile(profile);
		return {
			preset: preset,
			semanticBoundarySource: GoSemanticBoundarySource.TypedApiOrModule,
			nativeAuthority: GoNativeAuthorityPolicy.defaultFor(preset),
			nativeAuthoritySource: GoPolicyResolutionSource.PolicyPreset,
			nativeSpecialization: GoNativeSpecializationPolicy.defaultFor(preset),
			nativeSpecializationSource: GoPolicyResolutionSource.PolicyPreset,
			nativeFallback: GoNativeFallbackPolicy.defaultFor(preset),
			nativeFallbackSource: GoPolicyResolutionSource.PolicyPreset
		};
	}

	static function normalizeGoModuleName(raw:Null<String>):String {
		if (raw == null) {
			return "snapshot";
		}
		var trimmed = StringTools.trim(raw);
		return trimmed == "" ? "snapshot" : trimmed;
	}

	static function sortedUniqueLowercase(values:Null<Array<String>>):Array<String> {
		var out:Array<String> = [];
		if (values != null) {
			for (value in values) {
				var normalized = value == null ? "" : StringTools.trim(value).toLowerCase();
				if (normalized != "" && out.indexOf(normalized) == -1) {
					out.push(normalized);
				}
			}
		}
		out.sort((a, b) -> a < b ? -1 : (a > b ? 1 : 0));
		return out;
	}

	static function sortedUnique(values:Null<Array<String>>):Array<String> {
		var out:Array<String> = [];
		if (values != null) {
			for (value in values) {
				var normalized = value == null ? "" : StringTools.trim(value);
				if (normalized != "" && out.indexOf(normalized) == -1) {
					out.push(normalized);
				}
			}
		}
		out.sort((a, b) -> a < b ? -1 : (a > b ? 1 : 0));
		return out;
	}

	static function normalizeOptimizationPreset(raw:Null<String>):String {
		if (raw == null) {
			return "portable_fast";
		}
		var trimmed = StringTools.trim(raw).toLowerCase();
		return trimmed == "" ? "portable_fast" : trimmed;
	}

	static function normalizeStrictUserBoundaryPolicy(raw:Null<String>):String {
		if (raw == null) {
			return "auto";
		}
		var trimmed = StringTools.trim(raw).toLowerCase();
		return trimmed == "" ? "auto" : trimmed;
	}
}
