package reflaxe.go;

class CompilationContext {
	public final profile:GoProfile;
	public final goModuleName:String;
	public final runtimeImportPath:String;
	public final rawNativeMode:RawNativeMode;
	public final emitLineDirectives:Bool;

	public function new(profile:GoProfile, ?goModuleName:String, ?rawNativeMode:RawNativeMode, ?emitLineDirectives:Bool) {
		this.profile = profile;
		var moduleName = normalizeGoModuleName(goModuleName);
		this.goModuleName = moduleName;
		this.runtimeImportPath = moduleName + "/hxrt";
		this.rawNativeMode = rawNativeMode == null ? RawNativeMode.Interp : rawNativeMode;
		this.emitLineDirectives = emitLineDirectives == true;
	}

	static function normalizeGoModuleName(raw:Null<String>):String {
		if (raw == null) {
			return "snapshot";
		}

		var trimmed = StringTools.trim(raw);
		return trimmed == "" ? "snapshot" : trimmed;
	}
}
