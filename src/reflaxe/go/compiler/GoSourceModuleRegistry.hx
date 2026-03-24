package reflaxe.go.compiler;

#if macro
import haxe.macro.Context;
import haxe.macro.PositionTools;
import haxe.macro.Type;

/**
	Registry for mapping typed source positions back to Haxe module labels.

	Why:
	- `GoCompiler` uses source-module identity for generated file attribution,
	  line-directive rendering, and report metadata.
	- That bookkeeping is orchestration logic, not stdlib lowering logic.

	What:
	- Caches source-file and suffix-to-module lookups for classes and enums.
	- Resolves a typed expression position back to a normalized module label.

	How:
	- Build a deterministic lookup table once per compile from the selected typed
	  classes/enums, then query it wherever module attribution is needed.
**/
class GoSourceModuleRegistry {
	final sourceFileToModule:Map<String, String>;
	final sourceModuleBySuffix:Map<String, String>;
	var sourceModuleSuffixes:Array<String>;
	final normalizeModuleLabel:Null<String>->String;
	final normalizeSourcePath:String->String;
	final sourceModuleToFilePath:String->String;

	public function new(normalizeModuleLabel:Null<String>->String, normalizeSourcePath:String->String, sourceModuleToFilePath:String->String) {
		sourceFileToModule = [];
		sourceModuleBySuffix = [];
		sourceModuleSuffixes = [];
		this.normalizeModuleLabel = normalizeModuleLabel;
		this.normalizeSourcePath = normalizeSourcePath;
		this.sourceModuleToFilePath = sourceModuleToFilePath;
	}

	public function rebuild(classes:Array<ClassType>, enums:Array<EnumType>):Void {
		clearStringMap(sourceFileToModule);
		clearStringMap(sourceModuleBySuffix);
		sourceModuleSuffixes = [];

		for (classType in classes) {
			registerSourceModule(classType.module, classType.pos);
		}
		for (enumType in enums) {
			registerSourceModule(enumType.module, enumType.pos);
		}
		sourceModuleSuffixes.sort(compareSuffixBySpecificity);
	}

	public function sourceModuleForPos(pos:haxe.macro.Expr.Position):String {
		var sourcePath = normalizeSourcePath(Context.getPosInfos(pos).file);
		if (sourcePath != "" && sourceFileToModule.exists(sourcePath)) {
			return sourceFileToModule.get(sourcePath);
		}

		for (suffix in sourceModuleSuffixes) {
			if (pathEndsWithSuffix(sourcePath, suffix)) {
				return sourceModuleBySuffix.get(suffix);
			}
		}
		return "<unknown>";
	}

	function registerSourceModule(moduleName:Null<String>, pos:haxe.macro.Expr.Position):Void {
		var normalizedModule = normalizeModuleLabel(moduleName);
		if (normalizedModule == "<unknown>") {
			return;
		}

		var location = PositionTools.toLocation(pos);
		var sourcePath = location == null ? "" : normalizeSourcePath(Std.string(location.file));
		if (sourcePath != "" && !sourceFileToModule.exists(sourcePath)) {
			sourceFileToModule.set(sourcePath, normalizedModule);
		}

		var suffix = sourceModuleToFilePath(normalizedModule);
		if (suffix != "" && !sourceModuleBySuffix.exists(suffix)) {
			sourceModuleBySuffix.set(suffix, normalizedModule);
			sourceModuleSuffixes.push(suffix);
		}
	}

	static function clearStringMap(map:Map<String, String>):Void {
		var keys = [for (key in map.keys()) key];
		for (key in keys) {
			map.remove(key);
		}
	}

	static function pathEndsWithSuffix(path:String, suffix:String):Bool {
		if (path == suffix) {
			return true;
		}
		return path != null && suffix != null && path.length > suffix.length && StringTools.endsWith(path, "/" + suffix);
	}

	static function compareSuffixBySpecificity(a:String, b:String):Int {
		if (a.length == b.length) {
			return Reflect.compare(a, b);
		}
		return b.length - a.length;
	}
}
#end
