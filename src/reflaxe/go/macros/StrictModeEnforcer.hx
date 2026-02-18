package reflaxe.go.macros;

#if macro
import haxe.io.Path;
import haxe.macro.Context;
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

    var findings = findGoInjections();
    if (findings.length > 0) {
      Context.fatalError('StrictModeEnforcer: __go__ is not allowed in strict mode (' + findings[0] + ')', Context.currentPos());
    }
  }

  static function findGoInjections():Array<String> {
    var files = new Array<String>();
    var cwd = absolutePath(Sys.getCwd());

    for (classPath in Context.getClassPath()) {
      var full = absolutePath(classPath);
      if (!StringTools.startsWith(full, cwd)) {
        continue;
      }
      collectHxFiles(full, files);
    }

    var hits = new Array<String>();
    for (filePath in files) {
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
