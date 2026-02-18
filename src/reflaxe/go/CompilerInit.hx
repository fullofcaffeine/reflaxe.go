package reflaxe.go;

#if macro
import haxe.macro.Context;
#end

class CompilerInit {
  #if macro
  static var initialized = false;

  public static function Start():Void {
    if (initialized) {
      return;
    }
    initialized = true;

    Context.onAfterTyping(function(_) {
      if (!isGoBuild()) {
        return;
      }

      var outputDir = resolveOutputDir();
      var _ = ProfileResolver.resolve();
      var compiler = new GoCompiler();
      var files = compiler.compileModule();

      GoOutputIterator.writeFiles(outputDir, files);
      GoOutputIterator.writeGoMod(outputDir, "snapshot");
      GoOutputIterator.copyRuntime(outputDir);
    });
  }

  static function isGoBuild():Bool {
    var goOutput = Context.definedValue("go_output");
    if (goOutput != null && goOutput != "") {
      return true;
    }

    if (Context.defined("go")) {
      return true;
    }

    var args = Sys.args();
    for (i in 0...args.length) {
      var arg = args[i];
      if (arg == "-D" && i + 1 < args.length && StringTools.startsWith(args[i + 1], "go_output")) {
        return true;
      }
      if (StringTools.startsWith(arg, "-Dgo_output")) {
        return true;
      }
    }

    return false;
  }

  static function resolveOutputDir():String {
    var outputDir = Context.definedValue("go_output");
    if (outputDir == null || outputDir == "") {
      Context.fatalError("Missing required define -D go_output=<dir>", Context.currentPos());
    }
    return outputDir;
  }
  #else
  public static function Start():Void {}
  #end
}
