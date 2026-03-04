package reflaxe.go.compiler;

import reflaxe.go.GoProfile;
#if (macro || reflaxe_runtime)
import haxe.macro.Context;
import reflaxe.go.ProfileResolver;
import reflaxe.go.RawNativeMode;
import reflaxe.go.RawNativeModeResolver;

class GoBuildContextResolver {
	public static inline final STRICT_EXAMPLES_DEFINE = "reflaxe_go_strict_examples";
	public static inline final STRICT_DEFINE = "reflaxe_go_strict";
	public static inline final STRICT_POLICY_DEFINE = "reflaxe_go_strict_policy";
	public static inline final METAL_ALLOW_FALLBACK_DEFINE = "reflaxe_go_metal_allow_fallback";
	public static inline final AUTO_MODE_DEFINE = "reflaxe_go_auto";
	public static inline final GO_MODULE_DEFINE = "go_module";
	public static inline final LINE_DIRECTIVES_DEFINE = "reflaxe_go_line_directives";
	public static inline final HXRT_DEFAULT_FEATURES_DEFINE = "reflaxe_go_hxrt_default_features";
	public static inline final HXRT_FEATURES_DEFINE = "reflaxe_go_hxrt_features";
	public static inline final HXRT_NO_FEATURE_INFER_DEFINE = "reflaxe_go_hxrt_no_feature_infer";
	public static inline final CONTRACT_REPORT_DEFINE = "reflaxe_go_contract_report";
	public static inline final RUNTIME_PLAN_REPORT_DEFINE = "reflaxe_go_runtime_plan_report";
	public static inline final OPTIMIZER_PLAN_REPORT_DEFINE = "reflaxe_go_optimizer_plan_report";
	public static inline final OPT_PRESET_DEFINE = "reflaxe_go_opt";
	public static inline final OPT_PORTABLE_CONCURRENCY_FASTPATH_DEFINE = "reflaxe_go_opt_go_concurrency_fastpath";

	public static function resolve():GoBuildContext {
		var profile = ProfileResolver.resolve();
		var metalFallbackAllowed = Context.defined(METAL_ALLOW_FALLBACK_DEFINE);
		var strictPolicy = parseStrictUserBoundaryPolicy(Context.definedValue(STRICT_POLICY_DEFINE));
		var strictUserBoundaries = resolveStrictUserBoundaries(profile, strictPolicy, Context.defined(STRICT_DEFINE));
		var autoLoweringMode = parseAutoLoweringMode(Context.definedValue(AUTO_MODE_DEFINE));
		var optimizationPreset = parseOptimizationPreset(Context.definedValue(OPT_PRESET_DEFINE));
		var portableStringFastpathEnabled = optimizationPreset == "portable_fast";
		var portableConcurrencyFastpathEnabled = parseBoolDefine(OPT_PORTABLE_CONCURRENCY_FASTPATH_DEFINE, portableStringFastpathEnabled);
		return new GoBuildContext(profile, normalizeGoModuleName(Context.definedValue(GO_MODULE_DEFINE)), RawNativeModeResolver.resolve(),
			Context.defined(LINE_DIRECTIVES_DEFINE), Context.defined(STRICT_EXAMPLES_DEFINE), strictPolicy, strictUserBoundaries,
			metalFallbackAllowed, profile == GoProfile.Metal
			&& !metalFallbackAllowed, Context.defined(HXRT_DEFAULT_FEATURES_DEFINE), Context.defined(HXRT_FEATURES_DEFINE),
			Context.defined(HXRT_NO_FEATURE_INFER_DEFINE), parseManualHxrtFeatures(Context.definedValue(HXRT_FEATURES_DEFINE)),
			Context.defined(CONTRACT_REPORT_DEFINE), Context.defined(RUNTIME_PLAN_REPORT_DEFINE), parseBoolDefine(OPTIMIZER_PLAN_REPORT_DEFINE, false),
			autoLoweringMode, optimizationPreset, portableStringFastpathEnabled, portableConcurrencyFastpathEnabled, []);
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

	static function resolveStrictUserBoundaries(profile:GoProfile, strictPolicy:String, strictDefineEnabled:Bool):Bool {
		if (strictDefineEnabled) {
			return true;
		}
		return switch (strictPolicy) {
			case "on":
				true;
			case "off":
				false;
			case _:
				profile == GoProfile.Metal;
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
