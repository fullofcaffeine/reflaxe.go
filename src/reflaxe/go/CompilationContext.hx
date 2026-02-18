package reflaxe.go;

class CompilationContext {
  public final profile:GoProfile;
  public final goModuleName:String;
  public final runtimeImportPath:String;

  public function new(profile:GoProfile, ?goModuleName:String) {
    this.profile = profile;
    var moduleName = normalizeGoModuleName(goModuleName);
    this.goModuleName = moduleName;
    this.runtimeImportPath = moduleName + "/hxrt";
  }

  static function normalizeGoModuleName(raw:Null<String>):String {
    if (raw == null) {
      return "snapshot";
    }

    var trimmed = StringTools.trim(raw);
    return trimmed == "" ? "snapshot" : trimmed;
  }
}
