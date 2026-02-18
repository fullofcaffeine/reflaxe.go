package reflaxe.go;

#if macro
import haxe.macro.Context;
#end

class CompilerInit {
  #if macro
  static var initialized = false;

  public static function Start():Void {
    if (initialized || !isGoBuild()) {
      return;
    }
    initialized = true;

    var _ = ProfileResolver.resolve();

    Context.fatalError(
      "reflaxe.go scaffold is active but Go codegen is not implemented yet (Milestone 0).",
      Context.currentPos()
    );
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
