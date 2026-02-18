package reflaxe.go.macros;

#if macro
import haxe.io.Path;
import haxe.macro.Context;
import sys.FileSystem;
import sys.io.File;
#end

class BoundaryEnforcer {
  #if macro
  static var initialized = false;

  public static function init():Void {
    if (initialized) {
      return;
    }
    initialized = true;

    var findings = findGoInjections(true);
    if (findings.length > 0) {
      Context.fatalError('BoundaryEnforcer: __go__ is not allowed in strict examples (' + findings[0] + ')', Context.currentPos());
    }
  }

  static function findGoInjections(limitToExamples:Bool):Array<String> {
    var out = new Array<String>();
    var cwd = absolutePath(Sys.getCwd());

    for (classPath in Context.getClassPath()) {
      var full = absolutePath(classPath);
      if (!StringTools.startsWith(full, cwd)) {
        continue;
      }

      if (limitToExamples) {
        var normalized = full.split('\\').join('/');
        if (!StringTools.contains(normalized, "/test/") && !StringTools.contains(normalized, "/example") && !StringTools.contains(normalized, "/examples")) {
          continue;
        }
      }

      collectHxFiles(full, out);
    }

    var hits = new Array<String>();
    for (filePath in out) {
      var content = File.getContent(filePath);
      if (StringTools.contains(content, "__go__(")) {
        hits.push(filePath);
      }
    }
    return hits;
  }

  static function absolutePath(path:String):String {
    if (Path.isAbsolute(path)) {
      return Path.normalize(path);
    }
    return Path.normalize(Path.join([Sys.getCwd(), path]));
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
  #else
  public static function init():Void {}
  #end
}
