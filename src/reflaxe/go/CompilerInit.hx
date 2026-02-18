package reflaxe.go;

#if macro
import haxe.macro.Context;
import reflaxe.BaseCompiler.BaseCompilerFileOutputType;
import reflaxe.ReflectCompiler;
import reflaxe.go.macros.BoundaryEnforcer;
import reflaxe.go.macros.StrictModeEnforcer;
#end

class CompilerInit {
  #if macro
  static var initialized = false;

  public static function Start():Void {
    if (!isGoBuild()) {
      return;
    }

    if (initialized) {
      return;
    }
    initialized = true;

    var profile = ProfileResolver.resolve();
    if (Context.defined("reflaxe_go_strict_examples")) {
      BoundaryEnforcer.init();
    }
    if (Context.defined("reflaxe_go_strict") || profile == GoProfile.Metal) {
      StrictModeEnforcer.init();
    }

    ReflectCompiler.Start();
    ReflectCompiler.AddCompiler(new GoReflaxeCompiler(), {
      outputDirDefineName: "go_output",
      fileOutputType: Manual,
      fileOutputExtension: ".go",
      targetCodeInjectionName: "__go__",
      expressionPreprocessors: [],
      ignoreBodilessFunctions: false,
      ignoreExterns: true
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

  #else
  public static function Start():Void {}
  #end
}
