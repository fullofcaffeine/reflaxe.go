package reflaxe.go.macros;

#if macro
import haxe.io.Path;
import haxe.macro.Context;
import haxe.macro.Type;
import haxe.macro.TypedExprTools;
import reflaxe.go.GoProfile;
import reflaxe.go.ProfileResolver;
import sys.FileSystem;
import sys.io.File;
#end

class StrictModeEnforcer {
  #if macro
  static var initialized = false;

  public static function init():Void {
    if (initialized) {
      return;
    }
    initialized = true;

    if (!isGoBuild()) {
      return;
    }

    var projectRoot = normalizePath(Sys.getCwd());
    var allowFrameworkTypedInjections = ProfileResolver.resolve() == GoProfile.Metal;
    var preflightFindings = preflightScanForGoInjections(projectRoot);
    if (preflightFindings.length > 0) {
      Context.fatalError("StrictModeEnforcer: __go__ is not allowed in strict mode (" + preflightFindings[0] + ")", Context.currentPos());
    }
    Context.onAfterTyping(types -> enforce(types, projectRoot, allowFrameworkTypedInjections));
  }

  static function enforce(types:Array<ModuleType>, projectRoot:String, allowFrameworkTypedInjections:Bool):Void {
    for (moduleType in types) {
      switch (moduleType) {
        case TClassDecl(classRef):
          var classType = classRef.get();
          if (!isStrictProjectSource(classType.pos, projectRoot)) {
            continue;
          }
          enforceNoGoInjectionInClass(classType, projectRoot, allowFrameworkTypedInjections);
        case _:
      }
    }
  }

  static function enforceNoGoInjectionInClass(classType:ClassType, projectRoot:String, allowFrameworkTypedInjections:Bool):Void {
    var allFields = classType.fields.get().concat(classType.statics.get());
    for (field in allFields) {
      var expr = field.expr();
      if (expr == null) {
        continue;
      }
      scanForGoInjection(expr, projectRoot, allowFrameworkTypedInjections);
    }
  }

  static function scanForGoInjection(expr:TypedExpr, projectRoot:String, allowFrameworkTypedInjections:Bool):Void {
    if (isGoInjectionCall(expr)) {
      if (allowFrameworkTypedInjections && isFrameworkTypedInjectionExpr(expr.pos, projectRoot)) {
        TypedExprTools.iter(expr, e -> scanForGoInjection(e, projectRoot, allowFrameworkTypedInjections));
        return;
      }
      Context.error("StrictModeEnforcer: __go__ is not allowed in strict mode. "
        + "Prefer a typed wrapper or move target-specific interop into `std/`.", expr.pos);
    }

    TypedExprTools.iter(expr, e -> scanForGoInjection(e, projectRoot, allowFrameworkTypedInjections));
  }

  static function isGoInjectionCall(expr:TypedExpr):Bool {
    return switch (expr.expr) {
      case TCall(callTarget, _):
        switch (callTarget.expr) {
          case TIdent(name):
            name == "__go__";
          case TLocal(variable):
            variable.name == "__go__";
          case TField(_, fieldAccess):
            switch (fieldAccess) {
              case FInstance(_, _, classField) | FStatic(_, classField) | FAnon(classField) | FClosure(_, classField):
                classField.get().name == "__go__";
              case FEnum(_, enumField):
                enumField.name == "__go__";
              case FDynamic(name):
                name == "__go__";
            }
          case _:
            false;
        }
      case _:
        false;
    }
  }

  static function isStrictProjectSource(pos:haxe.macro.Expr.Position, projectRoot:String):Bool {
    var root = ensureTrailingSlash(projectRoot);
    var file = normalizePath(Context.getPosInfos(pos).file);
    if (file == null || file == "") {
      return false;
    }

    if (!Path.isAbsolute(file)) {
      file = normalizePath(Path.join([root, file]));
    }

    if (!StringTools.startsWith(file, root)) {
      return false;
    }

    if (file.indexOf("/src/reflaxe/") != -1 || file.indexOf("/std/") != -1) {
      return false;
    }

    return true;
  }

  static function isFrameworkTypedInjectionExpr(pos:haxe.macro.Expr.Position, projectRoot:String):Bool {
    var root = ensureTrailingSlash(projectRoot);
    var file = normalizePath(Context.getPosInfos(pos).file);
    if (file == null || file == "") {
      return false;
    }

    if (!Path.isAbsolute(file)) {
      file = normalizePath(Path.join([root, file]));
    }

    if (!StringTools.startsWith(file, root)) {
      return true;
    }

    return file.indexOf("/src/reflaxe/go/macros/") != -1 || file.indexOf("/std/go/metal/") != -1;
  }

  static function preflightScanForGoInjections(projectRoot:String):Array<String> {
    var root = ensureTrailingSlash(projectRoot);
    var files = new Array<String>();

    for (classPath in Context.getClassPath()) {
      var full = absolutePath(classPath);
      var fullWithSlash = ensureTrailingSlash(full);
      if (!StringTools.startsWith(full, root) && !StringTools.startsWith(fullWithSlash, root)) {
        continue;
      }
      if (full.indexOf("/src/reflaxe/") != -1 || full.indexOf("/std/") != -1) {
        continue;
      }
      collectHxFiles(full, files);
    }

    var findings = new Array<String>();
    for (path in files) {
      var content = File.getContent(path);
      if (StringTools.contains(content, "__go__(")) {
        findings.push(path);
      }
    }

    return findings;
  }

  static function absolutePath(path:String):String {
    if (Path.isAbsolute(path)) {
      return normalizePath(path);
    }
    return normalizePath(Path.join([Sys.getCwd(), path]));
  }

  static function collectHxFiles(path:String, out:Array<String>):Void {
    if (!FileSystem.exists(path)) {
      return;
    }
    if (FileSystem.isDirectory(path)) {
      for (entry in FileSystem.readDirectory(path)) {
        collectHxFiles(Path.join([path, entry]), out);
      }
      return;
    }
    if (StringTools.endsWith(path, ".hx")) {
      out.push(path);
    }
  }

  static function ensureTrailingSlash(path:String):String {
    var normalized = normalizePath(path);
    return StringTools.endsWith(normalized, "/") ? normalized : normalized + "/";
  }

  static function normalizePath(path:String):String {
    return Path.normalize(path).split("\\").join("/");
  }

  static function isGoBuild():Bool {
    var targetName = Context.definedValue("target.name");
    return targetName == "go" || Context.defined("go_output");
  }
  #else
  public static function init():Void {}
  #end
}
