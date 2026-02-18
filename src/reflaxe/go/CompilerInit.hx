package reflaxe.go;

#if macro
import haxe.macro.Context;
import reflaxe.go.macros.BoundaryEnforcer;
import reflaxe.go.macros.StrictModeEnforcer;
#end

class CompilerInit {
  #if macro
  static var initialized = false;

  public static function Start():Void {
    if (initialized) {
      return;
    }
    initialized = true;

    if (isGoBuild()) {
      var profile = ProfileResolver.resolve();
      if (Context.defined("reflaxe_go_strict_examples")) {
        BoundaryEnforcer.init();
      }
      if (Context.defined("reflaxe_go_strict") || profile == GoProfile.Metal) {
        StrictModeEnforcer.init();
      }
    }

    Context.onAfterTyping(function(types) {
      if (!isGoBuild()) {
        return;
      }

      var outputDir = resolveOutputDir();
      var profile = ProfileResolver.resolve();
      var compiler = new GoCompiler(new CompilationContext(profile));
      var files = compiler.compileModule(types);

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
