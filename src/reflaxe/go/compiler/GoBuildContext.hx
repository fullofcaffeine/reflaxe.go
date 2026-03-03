package reflaxe.go.compiler;

import reflaxe.go.GoProfile;
import reflaxe.go.RawNativeMode;

/**
	Immutable build-level configuration resolved once per compile.

	Why:
	- Keep contract and capability defaults centralized.
	- Avoid define-parsing drift between CompilerInit and output/runtime stages.
**/
class GoBuildContext {
	public final profile:GoProfile;
	public final goModuleName:String;
	public final rawNativeMode:RawNativeMode;
	public final emitLineDirectives:Bool;
	public final strictExamples:Bool;
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
	public final autoLoweringMode:GoAutoLoweringMode;
	public final optimizationPreset:String;
	public final portableStringFastpathEnabled:Bool;
	public final portableConcurrencyFastpathEnabled:Bool;
	public final metalLaneModules:Array<String>;

	public function new(profile:GoProfile, goModuleName:String, rawNativeMode:RawNativeMode, emitLineDirectives:Bool, strictExamples:Bool,
			strictUserBoundaries:Bool, metalFallbackAllowed:Bool, metalContractHardError:Bool, hxrtForceFullCopy:Bool, hxrtFeaturesDefinePresent:Bool,
			hxrtNoFeatureInfer:Bool, hxrtManualFeatures:Array<String>, contractReportEnabled:Bool, runtimePlanReportEnabled:Bool,
			optimizerPlanReportEnabled:Bool, autoLoweringMode:GoAutoLoweringMode, optimizationPreset:String, portableStringFastpathEnabled:Bool,
			portableConcurrencyFastpathEnabled:Bool, metalLaneModules:Array<String>) {
		this.profile = profile;
		this.goModuleName = normalizeGoModuleName(goModuleName);
		this.rawNativeMode = rawNativeMode == null ? RawNativeMode.Interp : rawNativeMode;
		this.emitLineDirectives = emitLineDirectives == true;
		this.strictExamples = strictExamples == true;
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
		this.autoLoweringMode = autoLoweringMode == null ? GoAutoLoweringMode.Off : autoLoweringMode;
		this.optimizationPreset = normalizeOptimizationPreset(optimizationPreset);
		this.portableStringFastpathEnabled = portableStringFastpathEnabled == true;
		this.portableConcurrencyFastpathEnabled = portableConcurrencyFastpathEnabled == true;
		this.metalLaneModules = sortedUnique(metalLaneModules);
	}

	public inline function isMetalContract():Bool {
		return profile == GoProfile.Metal;
	}

	public inline function isHxrtSelectiveEnabled():Bool {
		return hxrtFeaturesDefinePresent || hxrtNoFeatureInfer;
	}

	public function withMetalLaneModules(metalLaneModules:Array<String>):GoBuildContext {
		return new GoBuildContext(profile, goModuleName, rawNativeMode, emitLineDirectives, strictExamples, strictUserBoundaries, metalFallbackAllowed,
			metalContractHardError, hxrtForceFullCopy, hxrtFeaturesDefinePresent, hxrtNoFeatureInfer, hxrtManualFeatures, contractReportEnabled,
			runtimePlanReportEnabled, optimizerPlanReportEnabled, autoLoweringMode, optimizationPreset, portableStringFastpathEnabled,
			portableConcurrencyFastpathEnabled, metalLaneModules);
	}

	public static function legacyDefaults(profile:GoProfile, ?goModuleName:String, ?rawNativeMode:RawNativeMode, ?emitLineDirectives:Bool):GoBuildContext {
		return new GoBuildContext(profile, normalizeGoModuleName(goModuleName), rawNativeMode == null ? RawNativeMode.Interp : rawNativeMode,
			emitLineDirectives == true, false, false, false, false, false, false, false, [], false, false, false, GoAutoLoweringMode.Off, "portable_fast",
			true, true, []);
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
}
