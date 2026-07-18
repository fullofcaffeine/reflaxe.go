package reflaxe.go;

import reflaxe.go.compiler.GoBuildContext;
import reflaxe.go.compiler.GoHxrtFeatureAnalyzer.GoHxrtFeatureReason;

/**
	Why
	A failed typed representation can be observable under either compatibility
	preset, so the ledger must not describe it as a metal-only violation.

	What
	Records one native specialization fallback with stable source attribution.

	How
	Lowering appends the event before applying fallback policy; reports then split
	it by explicit native boundary without changing the fallback decision.
**/
typedef NativeFallbackEvent = {
	var kind:String;
	var detail:String;
	var location:String;
	var module:String;
	var inNativeBoundary:Bool;
}

/**
	Why
	Reports and optimizer gates need deterministic evidence for every attempted
	typed lowering, including whether source owned explicit native authority.

	What
	Stores one attempted, successful, or fallback lowering decision.

	How
	The compiler resolves the source module and `@:goNative` membership at the
	call site, then report renderers preserve canonical and compatibility fields.
**/
typedef LoweringDecisionLedgerEntry = {
	var feature:String;
	var kind:String;
	var outcome:String;
	var detail:String;
	var location:String;
	var module:String;
	var inNativeBoundary:Bool;
}

typedef GoAstPassSelectionReasonEntry = {
	var pass:String;
	var reason:String;
	var source:String;
}

class CompilationContext {
	public final buildContext:GoBuildContext;
	public final profile:GoProfile;
	public final goModuleName:String;
	public final runtimeImportPath:String;
	public final rawNativeMode:RawNativeMode;
	public final emitLineDirectives:Bool;
	public final leafReceiverTypes:Map<String, Bool>;
	public final leafReturningFunctions:Map<String, Bool>;
	public var inferredHxrtFeatures:Array<String>;
	public var inferredHxrtFeatureReasons:Array<GoHxrtFeatureReason>;
	public var selectedHxrtFeatures:Array<String>;
	public var requiredStdlibShimGroups:Array<String>;
	public var nativeFallbackEvents:Array<NativeFallbackEvent>;
	public var loweringDecisionLedger:Array<LoweringDecisionLedgerEntry>;
	public var appliedGoAstPassNames:Array<String>;
	public var selectedGoAstPassSource:String;
	public var selectedGoAstPassReasons:Array<GoAstPassSelectionReasonEntry>;
	public var optimizerStringInstanceTypedLowerings:Int;
	public var optimizerStringInstanceLegacyLowerings:Int;
	public var optimizerStringLengthFieldTypedLowerings:Int;
	public var optimizerStringLengthFieldLegacyLowerings:Int;
	public var optimizerPortableConcurrencyTypedFastpathHits:Int;
	public var optimizerPortableConcurrencyTypedFastpathFallbacks:Int;
	public var optimizerGoCollectionsTypedLowerings:Int;
	public var optimizerGoCollectionsTypedFallbacks:Int;
	public var optimizerGoResultTypedLowerings:Int;
	public var optimizerGoResultTypedFallbacks:Int;

	public function new(profile:GoProfile, ?goModuleName:String, ?rawNativeMode:RawNativeMode, ?emitLineDirectives:Bool, ?buildContext:GoBuildContext) {
		this.profile = profile;
		var moduleName = normalizeGoModuleName(goModuleName);
		this.goModuleName = moduleName;
		this.runtimeImportPath = moduleName + "/hxrt";
		this.rawNativeMode = rawNativeMode == null ? RawNativeMode.Interp : rawNativeMode;
		this.emitLineDirectives = emitLineDirectives == true;
		this.buildContext = buildContext == null ? GoBuildContext.legacyDefaults(profile, moduleName, this.rawNativeMode,
			this.emitLineDirectives) : buildContext;
		this.leafReceiverTypes = new Map<String, Bool>();
		this.leafReturningFunctions = new Map<String, Bool>();
		this.inferredHxrtFeatures = [];
		this.inferredHxrtFeatureReasons = [];
		this.selectedHxrtFeatures = [];
		this.requiredStdlibShimGroups = [];
		this.nativeFallbackEvents = [];
		this.loweringDecisionLedger = [];
		this.appliedGoAstPassNames = [];
		this.selectedGoAstPassSource = "legacy_lean_bundle";
		this.selectedGoAstPassReasons = [];
		this.optimizerStringInstanceTypedLowerings = 0;
		this.optimizerStringInstanceLegacyLowerings = 0;
		this.optimizerStringLengthFieldTypedLowerings = 0;
		this.optimizerStringLengthFieldLegacyLowerings = 0;
		this.optimizerPortableConcurrencyTypedFastpathHits = 0;
		this.optimizerPortableConcurrencyTypedFastpathFallbacks = 0;
		this.optimizerGoCollectionsTypedLowerings = 0;
		this.optimizerGoCollectionsTypedFallbacks = 0;
		this.optimizerGoResultTypedLowerings = 0;
		this.optimizerGoResultTypedFallbacks = 0;
	}

	public static function fromBuildContext(buildContext:GoBuildContext):CompilationContext {
		return new CompilationContext(buildContext.profile, buildContext.goModuleName, buildContext.rawNativeMode, buildContext.emitLineDirectives,
			buildContext);
	}

	static function normalizeGoModuleName(raw:Null<String>):String {
		if (raw == null) {
			return "snapshot";
		}

		var trimmed = StringTools.trim(raw);
		return trimmed == "" ? "snapshot" : trimmed;
	}
}
