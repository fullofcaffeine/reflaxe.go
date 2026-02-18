package reflaxe.go;

#if (macro || reflaxe_runtime)
import haxe.io.Path;
import haxe.macro.Context;
import haxe.macro.Type;
import reflaxe.GenericCompiler;
import reflaxe.data.ClassFuncData;
import reflaxe.data.ClassVarData;
import reflaxe.data.EnumOptionData;
import reflaxe.output.DataAndFileInfo;
import reflaxe.output.StringOrBytes;
import sys.FileSystem;
import sys.io.File;

class GoReflaxeCompiler extends GenericCompiler<Bool, Bool, Dynamic, Dynamic, Dynamic> {
  var allModules:Array<ModuleType> = [];
  var modulesByName:Map<String, Array<ModuleType>> = [];
  var selectedModuleNames:Map<String, Bool> = [];
  var generatedFiles:Array<GoCompiler.GoGeneratedFile> = [];
  var profile:GoProfile = GoProfile.Portable;

  public function new() {
    super();
  }

  override public function filterTypes(moduleTypes:Array<ModuleType>):Array<ModuleType> {
    allModules = moduleTypes.copy();
    modulesByName = [];
    for (moduleType in moduleTypes) {
      var moduleName = moduleNameOf(moduleType);
      if (moduleName != null && moduleName != "") {
        var list = modulesByName.get(moduleName);
        if (list == null) {
          list = [];
          modulesByName.set(moduleName, list);
        }
        list.push(moduleType);
      }
    }
    return moduleTypes;
  }

  override public function onCompileStart():Void {
    profile = ProfileResolver.resolve();
    selectedModuleNames = [];
    generatedFiles = [];
  }

  override public function onCompileEnd():Void {
    var selectedModules = collectSelectedModules();
    var compiler = new GoCompiler(new CompilationContext(profile));
    generatedFiles = compiler.compileModule(selectedModules);
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

  public function compileClassImpl(classType:ClassType, varFields:Array<ClassVarData>, funcFields:Array<ClassFuncData>):Null<Bool> {
    markModuleSelected(classType.module);
    return null;
  }

  public function compileEnumImpl(enumType:EnumType, options:Array<EnumOptionData>):Null<Bool> {
    markModuleSelected(enumType.module);
    return null;
  }

  public function compileExpressionImpl(expr:TypedExpr, topLevel:Bool):Null<Dynamic> {
    return null;
  }

  function collectSelectedModules():Array<ModuleType> {
    var selected = new Array<ModuleType>();
    var names = [for (name in selectedModuleNames.keys()) name];
    names.sort(function(a, b) return Reflect.compare(a, b));
    for (name in names) {
      var moduleTypeList = modulesByName.get(name);
      if (moduleTypeList != null) {
        for (moduleType in moduleTypeList) {
          selected.push(moduleType);
        }
      }
    }
    if (selected.length > 0) {
      return selected;
    }
    return allModules;
  }

  function markModuleSelected(moduleName:String):Void {
    if (moduleName == null || moduleName == "") {
      return;
    }
    if (modulesByName.exists(moduleName)) {
      selectedModuleNames.set(moduleName, true);
    }
  }

  function moduleNameOf(moduleType:ModuleType):String {
    return switch (moduleType) {
      case TClassDecl(classRef): classRef.get().module;
      case TEnumDecl(enumRef): enumRef.get().module;
      case TTypeDecl(defRef): defRef.get().module;
      case TAbstract(abstractRef): abstractRef.get().module;
    };
  }

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
