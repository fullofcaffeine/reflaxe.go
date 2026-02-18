package reflaxe.go;

#if macro
import haxe.io.Path;
import haxe.macro.Context;
import reflaxe.go.GoCompiler.GoGeneratedFile;
import sys.FileSystem;
import sys.io.File;
#end

class GoOutputIterator {
  public function new() {}

  #if macro
  public static function writeFiles(outputDir:String, files:Array<GoGeneratedFile>):Void {
    ensureDirectory(outputDir);
    for (file in files) {
      var path = Path.join([outputDir, file.relativePath]);
      ensureDirectory(Path.directory(path));
      File.saveContent(path, file.contents);
    }
  }

  public static function writeGoMod(outputDir:String, moduleName:String):Void {
    var goMod = [
      "module " + moduleName,
      "",
      "go 1.22",
      ""
    ].join("\n");
    var path = Path.join([outputDir, "go.mod"]);
    ensureDirectory(Path.directory(path));
    File.saveContent(path, goMod);
  }

  public static function copyRuntime(outputDir:String):Void {
    var runtimeSource = Path.join([findLibraryRoot(), "runtime", "hxrt"]);
    if (!FileSystem.exists(runtimeSource) || !FileSystem.isDirectory(runtimeSource)) {
      Context.fatalError('Missing runtime directory at "' + runtimeSource + '"', Context.currentPos());
    }

    var runtimeTarget = Path.join([outputDir, "hxrt"]);
    copyDirectory(runtimeSource, runtimeTarget);
  }

  static function copyDirectory(source:String, target:String):Void {
    ensureDirectory(target);
    for (entry in FileSystem.readDirectory(source)) {
      var sourcePath = Path.join([source, entry]);
      var targetPath = Path.join([target, entry]);
      if (FileSystem.isDirectory(sourcePath)) {
        copyDirectory(sourcePath, targetPath);
      } else {
        ensureDirectory(Path.directory(targetPath));
        File.copy(sourcePath, targetPath);
      }
    }
  }

  static function ensureDirectory(path:String):Void {
    if (path == null || path == "") {
      return;
    }
    if (!FileSystem.exists(path)) {
      FileSystem.createDirectory(path);
    }
  }

  static function findLibraryRoot():String {
    var thisFile = Context.resolvePath("reflaxe/go/GoOutputIterator.hx");
    var srcDir = Path.normalize(Path.directory(thisFile));
    return Path.normalize(Path.join([srcDir, "..", "..", ".."]));
  }
  #end
}
