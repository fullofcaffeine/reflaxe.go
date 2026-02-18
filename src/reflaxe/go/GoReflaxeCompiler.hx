package reflaxe.go;

#if (macro || reflaxe_runtime)
import haxe.io.Path;
import haxe.macro.Context;
import haxe.macro.Type;
import reflaxe.BaseCompiler;
import reflaxe.data.ClassFuncData;
import reflaxe.data.ClassVarData;
import reflaxe.data.EnumOptionData;
import reflaxe.output.DataAndFileInfo;
import reflaxe.output.StringOrBytes;
import sys.FileSystem;
import sys.io.File;

class GoReflaxeCompiler extends BaseCompiler {
  var typedModules:Array<ModuleType> = [];
  var generatedFiles:Array<GoCompiler.GoGeneratedFile> = [];
  var profile:GoProfile = GoProfile.Portable;

  public function new() {
    super();
  }

  override public function filterTypes(moduleTypes:Array<ModuleType>):Array<ModuleType> {
    typedModules = moduleTypes.copy();
    return moduleTypes;
  }

  override public function onCompileStart():Void {
    profile = ProfileResolver.resolve();
    generatedFiles = [];
  }

  override public function onCompileEnd():Void {
    var compiler = new GoCompiler(new CompilationContext(profile));
    generatedFiles = compiler.compileModule(typedModules);
  }

  override public function generateFilesManually():Void {
    if (output == null) {
      Context.fatalError("GoReflaxeCompiler output manager is not initialized", Context.currentPos());
      return;
    }

    for (file in generatedFiles) {
      output.saveFile(file.relativePath, file.contents);
    }

    output.saveFile("go.mod", buildGoMod(resolveGoModuleName()));
    writeRuntime(output);
  }

  public function generateOutputIterator():Iterator<DataAndFileInfo<StringOrBytes>> {
    var empty:Array<DataAndFileInfo<StringOrBytes>> = [];
    return empty.iterator();
  }

  public function compileClass(classType:ClassType, varFields:Array<ClassVarData>, funcFields:Array<ClassFuncData>):Void {}

  public function compileEnum(enumType:EnumType, options:Array<EnumOptionData>):Void {}

  public function compileTypedef(classType:DefType):Void {}

  public function compileAbstract(classType:AbstractType):Void {}

  function resolveGoModuleName():String {
    var moduleName = Context.definedValue("go_module");
    if (moduleName == null || StringTools.trim(moduleName) == "") {
      return "snapshot";
    }
    return StringTools.trim(moduleName);
  }

  function buildGoMod(moduleName:String):String {
    return [
      "module " + moduleName,
      "",
      "go 1.22",
      ""
    ].join("\n");
  }

  function writeRuntime(outputManager:reflaxe.output.OutputManager):Void {
    var runtimeSource = Path.join([findLibraryRoot(), "runtime", "hxrt"]);
    if (!FileSystem.exists(runtimeSource) || !FileSystem.isDirectory(runtimeSource)) {
      Context.fatalError('Missing runtime directory at "' + runtimeSource + '"', Context.currentPos());
    }

    writeRuntimeDir(outputManager, runtimeSource, "hxrt");
  }

  function writeRuntimeDir(outputManager:reflaxe.output.OutputManager, sourceDir:String, targetDir:String):Void {
    for (entry in FileSystem.readDirectory(sourceDir)) {
      var sourcePath = Path.join([sourceDir, entry]);
      var targetPath = Path.join([targetDir, entry]);

      if (FileSystem.isDirectory(sourcePath)) {
        writeRuntimeDir(outputManager, sourcePath, targetPath);
      } else {
        outputManager.saveFile(targetPath, File.getContent(sourcePath));
      }
    }
  }

  function findLibraryRoot():String {
    var thisFile = Context.resolvePath("reflaxe/go/GoReflaxeCompiler.hx");
    var srcDir = Path.normalize(Path.directory(thisFile));
    return Path.normalize(Path.join([srcDir, "..", "..", ".."]));
  }
}
#else
class GoReflaxeCompiler {
  public function new() {}
}
#end
