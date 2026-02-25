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
	public static inline final METAL_ALLOW_FALLBACK_DEFINE = "reflaxe_go_metal_allow_fallback";
	public static inline final GO_MODULE_DEFINE = "go_module";
	public static inline final LINE_DIRECTIVES_DEFINE = "reflaxe_go_line_directives";
	public static inline final HXRT_DEFAULT_FEATURES_DEFINE = "reflaxe_go_hxrt_default_features";
	public static inline final HXRT_FEATURES_DEFINE = "reflaxe_go_hxrt_features";
	public static inline final HXRT_NO_FEATURE_INFER_DEFINE = "reflaxe_go_hxrt_no_feature_infer";
	public static inline final CONTRACT_REPORT_DEFINE = "reflaxe_go_contract_report";
	public static inline final RUNTIME_PLAN_REPORT_DEFINE = "reflaxe_go_runtime_plan_report";

	public static function resolve():GoBuildContext {
		var profile = ProfileResolver.resolve();
		var metalFallbackAllowed = Context.defined(METAL_ALLOW_FALLBACK_DEFINE);
		return new GoBuildContext(profile, normalizeGoModuleName(Context.definedValue(GO_MODULE_DEFINE)), RawNativeModeResolver.resolve(),
			Context.defined(LINE_DIRECTIVES_DEFINE), Context.defined(STRICT_EXAMPLES_DEFINE), Context.defined(STRICT_DEFINE) || (profile == GoProfile.Metal
				&& !metalFallbackAllowed), metalFallbackAllowed, profile == GoProfile.Metal && !metalFallbackAllowed,
			Context.defined(HXRT_DEFAULT_FEATURES_DEFINE), Context.defined(HXRT_FEATURES_DEFINE),
			Context.defined(HXRT_NO_FEATURE_INFER_DEFINE), parseManualHxrtFeatures(Context.definedValue(HXRT_FEATURES_DEFINE)),
			Context.defined(CONTRACT_REPORT_DEFINE), Context.defined(RUNTIME_PLAN_REPORT_DEFINE), []);
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
}
#else
class GoBuildContextResolver {
	public static function resolve():GoBuildContext {
		return GoBuildContext.legacyDefaults(GoProfile.Portable);
	}
}
#end
