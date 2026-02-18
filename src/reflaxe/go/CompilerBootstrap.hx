package reflaxe.go;

#if macro
import haxe.io.Path;
import haxe.macro.Compiler;
import haxe.macro.Context;
import sys.FileSystem;
import sys.io.File;
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
    var vendoredReflaxe = Path.normalize(Path.join([root, "vendor", "reflaxe", "src"]));
    var standardLibrary = Path.normalize(Path.join([root, "std"]));
    var stagedStd = Path.normalize(Path.join([root, "std", "_std"]));

    if (isGoBuild()) {
      injectClassPathsFirst(filterExistingPaths([stagedStd, standardLibrary, vendoredReflaxe]));
      return;
    }

    injectClassPathsFirst(filterExistingPaths([vendoredReflaxe]));
  }

  static function filterExistingPaths(paths:Array<String>):Array<String> {
    var out = new Array<String>();
    for (path in paths) {
      var normalized = Path.normalize(path);
      if (!FileSystem.exists(normalized) || !FileSystem.isDirectory(normalized)) {
        continue;
      }
      out.push(normalized);
    }
    return out;
  }

  static function injectClassPathsFirst(paths:Array<String>):Void {
    if (paths == null || paths.length == 0) {
      return;
    }

    var config = Compiler.getConfiguration();
    if (config == null) {
      for (path in paths) {
        Compiler.addClassPath(path);
      }
      return;
    }

    var classPathField = "classPath";
    var existingDynamic:Dynamic = null;
    if (Reflect.hasField(config, "classPath")) {
      existingDynamic = Reflect.field(config, "classPath");
    } else if (Reflect.hasField(config, "classPaths")) {
      classPathField = "classPaths";
      existingDynamic = Reflect.field(config, "classPaths");
    }

    if (existingDynamic == null || !Std.isOfType(existingDynamic, Array)) {
      for (path in paths) {
        Compiler.addClassPath(path);
      }
      return;
    }

    var existing:Array<String> = cast existingDynamic;
    var injected = new Map<String, Bool>();
    var dedupedPaths = new Array<String>();
    for (path in paths) {
      var normalized = Path.normalize(path);
      if (injected.exists(normalized)) {
        continue;
      }
      injected.set(normalized, true);
      dedupedPaths.push(path);
    }

    var keep = new Array<String>();
    for (path in existing) {
      var normalized = Path.normalize(path);
      if (!injected.exists(normalized)) {
        keep.push(path);
      }
    }

    Reflect.setField(config, classPathField, dedupedPaths.concat(keep));
  }

  static function argsContainDefine(args:Array<String>, defineName:String):Bool {
    var i = 0;
    while (i < args.length) {
      var arg = args[i];
      if (arg == "-D" || arg == "--define") {
        if (i + 1 < args.length) {
          var defineArg = args[i + 1];
          if (defineArg == defineName || StringTools.startsWith(defineArg, defineName + "=")) {
            return true;
          }
        }
        i += 2;
        continue;
      }

      if (StringTools.startsWith(arg, "-D" + defineName)) {
        return true;
      }

      i += 1;
    }

    return false;
  }

  static function hxmlContainsDefine(hxmlPath:String, defineName:String, seen:Map<String, Bool>):Bool {
    var normalizedPath = Path.normalize(hxmlPath);
    if (seen.exists(normalizedPath)) {
      return false;
    }
    seen.set(normalizedPath, true);

    if (!FileSystem.exists(normalizedPath)) {
      return false;
    }

    var content = File.getContent(normalizedPath);
    var args:Array<String> = [];
    for (line in content.split("\n")) {
      var raw = StringTools.trim(line);
      if (raw.length == 0) {
        continue;
      }
      if (StringTools.startsWith(raw, "#")) {
        continue;
      }

      var commentIndex = raw.indexOf("#");
      if (commentIndex >= 0) {
        raw = StringTools.trim(raw.substr(0, commentIndex));
      }
      if (raw.length == 0) {
        continue;
      }

      for (token in raw.split(" ")) {
        var trimmed = StringTools.trim(token);
        if (trimmed.length > 0) {
          args.push(trimmed);
        }
      }
    }

    if (argsContainDefine(args, defineName)) {
      return true;
    }

    for (arg in args) {
      if (StringTools.startsWith(arg, "@")) {
        var nested = arg.substr(1);
        if (hxmlContainsDefine(nested, defineName, seen)) {
          return true;
        }
      }
    }

    return false;
  }

  static function hasDefineInArgs(defineName:String):Bool {
    var config = Compiler.getConfiguration();
    if (config == null) {
      return false;
    }

    var args = config.args;
    if (args == null) {
      return false;
    }
    if (argsContainDefine(args, defineName)) {
      return true;
    }

    var seen = new Map<String, Bool>();
    for (arg in args) {
      if (StringTools.endsWith(arg, ".hxml") && hxmlContainsDefine(arg, defineName, seen)) {
        return true;
      }
      if (StringTools.startsWith(arg, "@")) {
        var nested = arg.substr(1);
        if (hxmlContainsDefine(nested, defineName, seen)) {
          return true;
        }
      }
    }

    return false;
  }

  static function isGoBuild():Bool {
    var targetName = Context.definedValue("target.name");
    if (targetName == "go") {
      return true;
    }
    if (Context.defined("go_output")) {
      return true;
    }

    var config = Compiler.getConfiguration();
    if (config != null) {
      switch (config.platform) {
        #if (haxe >= version("5.0.0"))
        case CustomTarget("go"):
          return true;
        #end
        case _:
      }
    }

    return hasDefineInArgs("go_output");
  }

  static function findLibraryRoot():String {
    var thisFile = Context.resolvePath("reflaxe/go/CompilerBootstrap.hx");
    var srcDir = Path.normalize(Path.directory(thisFile));
    return Path.normalize(Path.join([srcDir, "..", "..", ".."]));
  }
  #else
  public static function Start():Void {}
  #end
}
