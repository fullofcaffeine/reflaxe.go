package reflaxe.go.compiler;

import reflaxe.go.GoProfile;
#if (macro || reflaxe_runtime)
import haxe.macro.Context;
import reflaxe.go.ProfileResolver;
import reflaxe.go.RawNativeMode;
import reflaxe.go.RawNativeModeResolver;

class GoBuildContextResolver {
	public static inline final STRICT_EXAMPLES_DEFINE:GoCompilerDefine = GoCompilerDefine.DefineStrictExamples;
	public static inline final STRICT_DEFINE:GoCompilerDefine = GoCompilerDefine.DefineStrict;
	public static inline final STRICT_POLICY_DEFINE:GoCompilerDefine = GoCompilerDefine.DefineStrictPolicy;
	public static inline final METAL_ALLOW_FALLBACK_DEFINE:GoCompilerDefine = GoCompilerDefine.DefineMetalAllowFallback;
	public static inline final NATIVE_AUTHORITY_DEFINE:GoCompilerDefine = GoCompilerDefine.DefineNativeAuthority;
	public static inline final NATIVE_SPECIALIZATION_DEFINE:GoCompilerDefine = GoCompilerDefine.DefineNativeSpecialization;
	public static inline final NATIVE_FALLBACK_DEFINE:GoCompilerDefine = GoCompilerDefine.DefineNativeFallback;
	public static inline final AUTO_MODE_DEFINE:GoCompilerDefine = GoCompilerDefine.DefineAutoMode;
	public static inline final GO_MODULE_DEFINE:GoCompilerDefine = GoCompilerDefine.DefineGoModule;
	public static inline final LINE_DIRECTIVES_DEFINE:GoCompilerDefine = GoCompilerDefine.DefineLineDirectives;
	public static inline final HXRT_DEFAULT_FEATURES_DEFINE:GoCompilerDefine = GoCompilerDefine.DefineHxrtDefaultFeatures;
	public static inline final HXRT_FEATURES_DEFINE:GoCompilerDefine = GoCompilerDefine.DefineHxrtFeatures;
	public static inline final HXRT_NO_FEATURE_INFER_DEFINE:GoCompilerDefine = GoCompilerDefine.DefineHxrtNoFeatureInfer;
	public static inline final CONTRACT_REPORT_DEFINE:GoCompilerDefine = GoCompilerDefine.DefineContractReport;
	public static inline final RUNTIME_PLAN_REPORT_DEFINE:GoCompilerDefine = GoCompilerDefine.DefineRuntimePlanReport;
	public static inline final OPTIMIZER_PLAN_REPORT_DEFINE:GoCompilerDefine = GoCompilerDefine.DefineOptimizerPlanReport;
	public static inline final OPT_PRESET_DEFINE:GoCompilerDefine = GoCompilerDefine.DefineOptimizationPreset;
	public static inline final OPT_PORTABLE_CONCURRENCY_FASTPATH_DEFINE:GoCompilerDefine = GoCompilerDefine.DefinePortableConcurrencyFastpath;
	public static inline final NATIVE_STACK_TRACE_DEFINE:GoCompilerDefine = GoCompilerDefine.DefineNativeStackTrace;

	public static function resolve():GoBuildContext {
		var profile = ProfileResolver.resolve();
		var metalFallbackAllowed = Context.defined(METAL_ALLOW_FALLBACK_DEFINE);
		var policyResolution = resolvePolicy(profile, metalFallbackAllowed);
		var strictPolicy = parseStrictUserBoundaryPolicy(Context.definedValue(STRICT_POLICY_DEFINE));
		var strictUserBoundaries = resolveStrictUserBoundaries(policyResolution.preset, strictPolicy, Context.defined(STRICT_DEFINE));
		var autoLoweringMode = parseAutoLoweringMode(Context.definedValue(AUTO_MODE_DEFINE));
		var optimizationPreset = parseOptimizationPreset(Context.definedValue(OPT_PRESET_DEFINE));
		var portableStringFastpathEnabled = optimizationPreset == "portable_fast";
		var portableConcurrencyFastpathEnabled = parseBoolDefine(OPT_PORTABLE_CONCURRENCY_FASTPATH_DEFINE, portableStringFastpathEnabled);
		return new GoBuildContext(profile, policyResolution, normalizeGoModuleName(Context.definedValue(GO_MODULE_DEFINE)), RawNativeModeResolver.resolve(),
			Context.defined(LINE_DIRECTIVES_DEFINE), Context.defined(STRICT_EXAMPLES_DEFINE), strictPolicy, strictUserBoundaries, metalFallbackAllowed,
			policyResolution.preset == GoPolicyPreset.MetalCompatibility
			&& policyResolution.nativeFallback == GoNativeFallbackPolicy.Error,
			Context.defined(HXRT_DEFAULT_FEATURES_DEFINE), Context.defined(HXRT_FEATURES_DEFINE), Context.defined(HXRT_NO_FEATURE_INFER_DEFINE),
			parseManualHxrtFeatures(Context.definedValue(HXRT_FEATURES_DEFINE)), Context.defined(CONTRACT_REPORT_DEFINE),
			Context.defined(RUNTIME_PLAN_REPORT_DEFINE), parseBoolDefine(OPTIMIZER_PLAN_REPORT_DEFINE, false), autoLoweringMode, optimizationPreset,
			portableStringFastpathEnabled, portableConcurrencyFastpathEnabled, Context.defined(NATIVE_STACK_TRACE_DEFINE), []);
	}

	/**
		Why
		The compatibility preset must be only a default bundle; explicit axes need
		deterministic precedence and provenance.

		What
		Resolves native authority, specialization, and fallback as typed values.

		How
		Canonical axis defines win, the legacy metal fallback define is honored
		when non-contradictory, and preset defaults fill every remaining value.
	**/
	static function resolvePolicy(profile:GoProfile, legacyMetalFallbackAllowed:Bool):GoPolicyResolution {
		var preset = GoPolicyPreset.fromLegacyProfile(profile);
		var authorityRaw = Context.definedValue(NATIVE_AUTHORITY_DEFINE);
		var authority = authorityRaw == null ? GoNativeAuthorityPolicy.defaultFor(preset) : parseNativeAuthorityPolicy(authorityRaw);
		var authoritySource = authorityRaw == null ? GoPolicyResolutionSource.PolicyPreset : GoPolicyResolutionSource.NativeAuthorityDefine;

		var specializationRaw = Context.definedValue(NATIVE_SPECIALIZATION_DEFINE);
		var specialization = specializationRaw == null ? GoNativeSpecializationPolicy.defaultFor(preset) : parseNativeSpecializationPolicy(specializationRaw);
		var specializationSource = specializationRaw == null ? GoPolicyResolutionSource.PolicyPreset : GoPolicyResolutionSource.NativeSpecializationDefine;

		var fallbackRaw = Context.definedValue(NATIVE_FALLBACK_DEFINE);
		var fallback = fallbackRaw == null ? GoNativeFallbackPolicy.defaultFor(preset) : parseNativeFallbackPolicy(fallbackRaw);
		var fallbackSource = fallbackRaw == null ? GoPolicyResolutionSource.PolicyPreset : GoPolicyResolutionSource.NativeFallbackDefine;
		if (legacyMetalFallbackAllowed) {
			if (fallbackRaw != null && fallback != GoNativeFallbackPolicy.Allow) {
				Context.fatalError('Conflicting native fallback settings: `-D ' + METAL_ALLOW_FALLBACK_DEFINE + '` requires `-D ' + NATIVE_FALLBACK_DEFINE
					+ '=allow`, but the canonical axis selected `error`.',
					Context.currentPos());
			}
			if (fallbackRaw == null) {
				fallback = GoNativeFallbackPolicy.Allow;
				fallbackSource = GoPolicyResolutionSource.LegacyMetalFallbackDefine;
			}
		}

		return {
			preset: preset,
			semanticBoundarySource: GoSemanticBoundarySource.TypedApiOrModule,
			nativeAuthority: authority,
			nativeAuthoritySource: authoritySource,
			nativeSpecialization: specialization,
			nativeSpecializationSource: specializationSource,
			nativeFallback: fallback,
			nativeFallbackSource: fallbackSource
		};
	}

	static function parseNativeAuthorityPolicy(raw:String):GoNativeAuthorityPolicy {
		var normalized = StringTools.trim(raw).toLowerCase();
		return switch (normalized) {
			case "guarded":
				GoNativeAuthorityPolicy.Guarded;
			case "explicit":
				GoNativeAuthorityPolicy.Explicit;
			case _:
				Context.fatalError('Unknown `' + NATIVE_AUTHORITY_DEFINE + '` value "' + raw + '" (expected: guarded, explicit)', Context.currentPos());
				GoNativeAuthorityPolicy.Guarded;
		};
	}

	static function parseNativeSpecializationPolicy(raw:String):GoNativeSpecializationPolicy {
		var normalized = StringTools.trim(raw).toLowerCase();
		return switch (normalized) {
			case "proven":
				GoNativeSpecializationPolicy.Proven;
			case "eager":
				GoNativeSpecializationPolicy.Eager;
			case _:
				Context.fatalError('Unknown `' + NATIVE_SPECIALIZATION_DEFINE + '` value "' + raw + '" (expected: proven, eager)', Context.currentPos());
				GoNativeSpecializationPolicy.Proven;
		};
	}

	static function parseNativeFallbackPolicy(raw:String):GoNativeFallbackPolicy {
		var normalized = StringTools.trim(raw).toLowerCase();
		return switch (normalized) {
			case "allow":
				GoNativeFallbackPolicy.Allow;
			case "error":
				GoNativeFallbackPolicy.Error;
			case _:
				Context.fatalError('Unknown `' + NATIVE_FALLBACK_DEFINE + '` value "' + raw + '" (expected: allow, error)', Context.currentPos());
				GoNativeFallbackPolicy.Allow;
		};
	}

	static function parseManualHxrtFeatures(raw:Null<String>):Array<String> {
		if (raw == null) {
			return [];
		}
		var tokens = [for (part in raw.split(",")) StringTools.trim(part).toLowerCase()];
		var out:Array<String> = [];
		for (token in tokens) {
			if (token == "") {
				continue;
			}
			if (!GoHxrtFeatureAnalyzer.isKnownFeature(token)) {
				var expected = GoHxrtFeatureAnalyzer.knownFeatures().join(", ");
				Context.fatalError('Unknown `' + HXRT_FEATURES_DEFINE + '` feature "' + token + '" (expected one of: ' + expected + ")", Context.currentPos());
			}
			if (!out.contains(token)) {
				out.push(token);
			}
		}
		return out;
	}

	static function normalizeGoModuleName(raw:Null<String>):String {
		if (raw == null) {
			return "snapshot";
		}
		var trimmed = StringTools.trim(raw);
		return trimmed == "" ? "snapshot" : trimmed;
	}

	static function parseOptimizationPreset(raw:Null<String>):String {
		if (raw == null) {
			return "portable_fast";
		}
		var normalized = StringTools.trim(raw).toLowerCase();
		if (normalized == "") {
			return "portable_fast";
		}
		return switch (normalized) {
			case "portable_fast", "none":
				normalized;
			case _:
				Context.fatalError('Unknown `' + OPT_PRESET_DEFINE + '` preset "' + raw + '" (expected: portable_fast, none)', Context.currentPos());
				"portable_fast";
		};
	}

	static function parseStrictUserBoundaryPolicy(raw:Null<String>):String {
		if (raw == null) {
			return "auto";
		}
		var normalized = StringTools.trim(raw).toLowerCase();
		if (normalized == "") {
			return "auto";
		}
		return switch (normalized) {
			case "auto", "on", "off":
				normalized;
			case _:
				Context.fatalError('Unknown `' + STRICT_POLICY_DEFINE + '` value "' + raw + '" (expected: auto, on, off)', Context.currentPos());
				"auto";
		};
	}

	static function resolveStrictUserBoundaries(preset:GoPolicyPreset, strictPolicy:String, strictDefineEnabled:Bool):Bool {
		if (strictDefineEnabled) {
			return true;
		}
		return switch (strictPolicy) {
			case "on":
				true;
			case "off":
				false;
			case _:
				preset == GoPolicyPreset.MetalCompatibility;
		};
	}

	static function parseAutoLoweringMode(raw:Null<String>):GoAutoLoweringMode {
		if (raw == null) {
			return GoAutoLoweringMode.Off;
		}
		var normalized = StringTools.trim(raw).toLowerCase();
		if (normalized == "") {
			return GoAutoLoweringMode.Auto;
		}
		return switch (normalized) {
			case "off":
				GoAutoLoweringMode.Off;
			case "auto":
				GoAutoLoweringMode.Auto;
			case "auto_strict":
				GoAutoLoweringMode.AutoStrict;
			case _:
				Context.fatalError('Unknown `' + AUTO_MODE_DEFINE + '` mode "' + raw + '" (expected: off, auto, auto_strict)', Context.currentPos());
				GoAutoLoweringMode.Off;
		};
	}

	static function parseBoolDefine(defineName:String, defaultValue:Bool):Bool {
		var raw = Context.definedValue(defineName);
		if (raw == null) {
			return defaultValue;
		}
		var normalized = StringTools.trim(raw).toLowerCase();
		if (normalized == "") {
			return true;
		}
		return switch (normalized) {
			case "1", "true", "yes", "on":
				true;
			case "0", "false", "no", "off":
				false;
			case _:
				Context.fatalError('Unknown boolean value for `-D ' + defineName + '=' + raw + '` (expected one of: 1,true,yes,on,0,false,no,off)',
					Context.currentPos());
				defaultValue;
		};
	}
}
#else
class GoBuildContextResolver {
	public static function resolve():GoBuildContext {
		return GoBuildContext.legacyDefaults(GoProfile.Portable);
	}
}
#end
