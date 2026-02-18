package reflaxe.go;

#if macro
import haxe.io.Path;
import haxe.macro.Compiler;
import haxe.macro.Context;
import sys.FileSystem;
import sys.io.File;
#end

class BuildDetection {
  #if macro
  public static function isGoBuild():Bool {
    return isTargetBuild("go_output", "go");
  }

  public static function isTargetBuild(outputDefine:String, targetName:String):Bool {
    var outputValue = Context.definedValue(outputDefine);
    if (outputValue != null && outputValue != "") {
      return true;
    }
    if (Context.defined(outputDefine) || Context.defined(targetName)) {
      return true;
    }

    var configuredTargetName = Context.definedValue("target.name");
    if (configuredTargetName == targetName) {
      return true;
    }

    var config = Compiler.getConfiguration();
    if (config != null) {
      switch (config.platform) {
        #if (haxe >= version("5.0.0"))
        case CustomTarget(name) if (name == targetName):
          return true;
        #end
        case _:
      }
    }

    return hasDefineInCompilerArgs(outputDefine);
  }

  static function hasDefineInCompilerArgs(defineName:String):Bool {
    var config = Compiler.getConfiguration();
    if (config == null || config.args == null) {
      return false;
    }

    var args = config.args;
    if (argsContainDefine(args, defineName)) {
      return true;
    }

    var seen = new Map<String, Bool>();
    var cwd = normalizeDir(Sys.getCwd());
    for (arg in args) {
      if (StringTools.endsWith(arg, ".hxml")) {
        if (hxmlContainsDefineAny(resolveIncludeCandidates(cwd, cwd, arg), defineName, seen)) {
          return true;
        }
        continue;
      }
      if (StringTools.startsWith(arg, "@")) {
        var includePath = arg.substr(1);
        if (hxmlContainsDefineAny(resolveIncludeCandidates(cwd, cwd, includePath), defineName, seen)) {
          return true;
        }
      }
    }

    return false;
  }

  static function argsContainDefine(args:Array<String>, defineName:String):Bool {
    var i = 0;
    while (i < args.length) {
      var arg = args[i];
      if (arg == "-D" || arg == "--define") {
        if (i + 1 < args.length) {
          if (defineArgMatches(args[i + 1], defineName)) {
            return true;
          }
        }
        i += 2;
        continue;
      }

      if (StringTools.startsWith(arg, "-D")) {
        if (defineArgMatches(arg.substr(2), defineName)) {
          return true;
        }
      }

      if (StringTools.startsWith(arg, "--define=")) {
        if (defineArgMatches(arg.substr("--define=".length), defineName)) {
          return true;
        }
      }

      i += 1;
    }

    return false;
  }

  static function defineArgMatches(raw:String, defineName:String):Bool {
    if (raw == null || raw == "") {
      return false;
    }
    return raw == defineName || StringTools.startsWith(raw, defineName + "=");
  }

  static function hxmlContainsDefineAny(paths:Array<String>, defineName:String, seen:Map<String, Bool>):Bool {
    for (path in paths) {
      if (hxmlContainsDefine(path, defineName, seen)) {
        return true;
      }
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

    var tokens = parseHxmlTokens(normalizedPath);
    if (argsContainDefine(tokens, defineName)) {
      return true;
    }

    var parentDir = normalizeDir(Path.directory(normalizedPath));
    for (token in tokens) {
      var nestedInclude = nestedIncludePath(token);
      if (nestedInclude == null) {
        continue;
      }
      var nestedCandidates = resolveIncludeCandidates(parentDir, normalizeDir(Sys.getCwd()), nestedInclude);
      if (hxmlContainsDefineAny(nestedCandidates, defineName, seen)) {
        return true;
      }
    }

    return false;
  }

  static function parseHxmlTokens(hxmlPath:String):Array<String> {
    var content = File.getContent(hxmlPath);
    var tokens = new Array<String>();
    for (line in content.split("\n")) {
      var raw = StringTools.trim(line);
      if (raw.length == 0 || StringTools.startsWith(raw, "#")) {
        continue;
      }

      var commentIndex = raw.indexOf("#");
      if (commentIndex >= 0) {
        raw = StringTools.trim(raw.substr(0, commentIndex));
      }
      if (raw.length == 0) {
        continue;
      }

      for (token in tokenizeLine(raw)) {
        if (token.length > 0) {
          tokens.push(token);
        }
      }
    }
    return tokens;
  }

  static function tokenizeLine(line:String):Array<String> {
    var tokens = new Array<String>();
    var token = "";
    var inQuotes = false;

    var i = 0;
    while (i < line.length) {
      var ch = line.charAt(i);

      if (ch == "\"") {
        inQuotes = !inQuotes;
        i += 1;
        continue;
      }

      if (!inQuotes && isWhitespaceChar(ch)) {
        if (token.length > 0) {
          tokens.push(token);
          token = "";
        }
        i += 1;
        continue;
      }

      token += ch;
      i += 1;
    }

    if (token.length > 0) {
      tokens.push(token);
    }
    return tokens;
  }

  static function isWhitespaceChar(ch:String):Bool {
    return ch == " " || ch == "\t" || ch == "\r";
  }

  static function resolveIncludeCandidates(baseDir:String, rootDir:String, includePath:String):Array<String> {
    if (includePath == null || includePath == "") {
      return [];
    }
    if (Path.isAbsolute(includePath)) {
      return [Path.normalize(includePath)];
    }
    var first = Path.normalize(Path.join([baseDir, includePath]));
    var second = Path.normalize(Path.join([rootDir, includePath]));
    if (first == second) {
      return [first];
    }
    return [first, second];
  }

  static function nestedIncludePath(token:String):Null<String> {
    if (token == null || token == "") {
      return null;
    }
    if (StringTools.startsWith(token, "@")) {
      var nested = token.substr(1);
      return nested == "" ? null : nested;
    }
    if (StringTools.endsWith(token, ".hxml")) {
      return token;
    }
    return null;
  }

  static function normalizeDir(path:String):String {
    if (path == null || path == "") {
      return "";
    }
    return Path.normalize(path);
  }
  #else
  public static function isGoBuild():Bool {
    return false;
  }
  #end
}
