package reflaxe.go;

#if macro
import haxe.io.Path;
import haxe.macro.Compiler;
import haxe.macro.Context;
import sys.FileSystem;
#end

class CompilerBootstrap {
  #if macro
  static var bootstrapped = false;

  public static function Start():Void {
    if (bootstrapped) {
      return;
    }
    bootstrapped = true;

    var root = findLibraryRoot();

    // Best-effort ordering: this macro is executed from extraParams.hxml very early.
    addClassPathIfExists(Path.join([root, "vendor", "reflaxe", "src"]));

    if (isGoBuild()) {
      addClassPathIfExists(Path.join([root, "std"]));
      addClassPathIfExists(Path.join([root, "std", "_std"]));
    }
  }

  static function isGoBuild():Bool {
    if (Context.defined("go")) {
      return true;
    }

    var goOutput = Context.definedValue("go_output");
    if (goOutput != null && goOutput != "") {
      return true;
    }

    var args = Sys.args();
    for (i in 0...args.length) {
      var arg = args[i];
      if (arg == "-D" && i + 1 < args.length) {
        var defineArg = args[i + 1];
        if (StringTools.startsWith(defineArg, "go_output")) {
          return true;
        }
      }
      if (StringTools.startsWith(arg, "-Dgo_output")) {
        return true;
      }
    }

    return false;
  }

  static function findLibraryRoot():String {
    var thisFile = Context.resolvePath("reflaxe/go/CompilerBootstrap.hx");
    var srcDir = Path.normalize(Path.directory(thisFile));
    return Path.normalize(Path.join([srcDir, "..", "..", ".."]));
  }

  static function addClassPathIfExists(path:String):Void {
    var normalized = Path.normalize(path);
    if (!FileSystem.exists(normalized)) {
      return;
    }
    Compiler.addClassPath(normalized);
  }
  #else
  public static function Start():Void {}
  #end
}
