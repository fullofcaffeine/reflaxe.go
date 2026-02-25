package reflaxe.go;

import reflaxe.go.compiler.GoBuildContext;

typedef MetalFallbackViolation = {
	var kind:String;
	var detail:String;
	var location:String;
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
	public var selectedHxrtFeatures:Array<String>;
	public var requiredStdlibShimGroups:Array<String>;
	public var requiresIoHelperSurface:Bool;
	public var metalFallbackViolations:Array<MetalFallbackViolation>;

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
		this.selectedHxrtFeatures = [];
		this.requiredStdlibShimGroups = [];
		this.requiresIoHelperSurface = false;
		this.metalFallbackViolations = [];
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
