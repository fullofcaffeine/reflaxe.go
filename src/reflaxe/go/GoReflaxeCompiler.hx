package reflaxe.go;

#if (macro || reflaxe_runtime)
import haxe.io.Path;
import haxe.macro.Context;
import haxe.macro.Type;
import reflaxe.GenericCompiler;
import reflaxe.data.ClassFuncData;
import reflaxe.data.ClassVarData;
import reflaxe.data.EnumOptionData;
import reflaxe.go.compiler.GoHxrtFeatureAnalyzer;
import reflaxe.output.DataAndFileInfo;
import reflaxe.output.StringOrBytes;
import sys.FileSystem;
import sys.io.File;

class GoReflaxeCompiler extends GenericCompiler<Bool, Bool, Dynamic, Dynamic, Dynamic> {
	var allModules:Array<ModuleType> = [];
	var selectedClasses:Array<ClassType> = [];
	var selectedEnums:Array<EnumType> = [];
	var generatedFiles:Array<GoCompiler.GoGeneratedFile> = [];
	var profile:GoProfile = GoProfile.Portable;
	var goModuleName:String = "snapshot";
	var rawNativeMode:RawNativeMode = RawNativeMode.Interp;
	var emitLineDirectives:Bool = false;
	var compilationContext:Null<CompilationContext> = null;

	static inline final HXRT_DEFAULT_FEATURES_DEFINE = "reflaxe_go_hxrt_default_features";
	static inline final HXRT_FEATURES_DEFINE = "reflaxe_go_hxrt_features";
	static inline final HXRT_NO_FEATURE_INFER_DEFINE = "reflaxe_go_hxrt_no_feature_infer";

	public function new() {
		super();
	}

	override public function filterTypes(moduleTypes:Array<ModuleType>):Array<ModuleType> {
		allModules = moduleTypes.copy();
		return moduleTypes;
	}

	override public function onCompileStart():Void {
		profile = ProfileResolver.resolve();
		goModuleName = resolveGoModuleName();
		rawNativeMode = RawNativeModeResolver.resolve();
		emitLineDirectives = Context.defined("reflaxe_go_line_directives");
		compilationContext = new CompilationContext(profile, goModuleName, rawNativeMode, emitLineDirectives);
		selectedClasses = [];
		selectedEnums = [];
		generatedFiles = [];
	}

	override public function onCompileEnd():Void {
		var context = compilationContext == null ? new CompilationContext(profile, goModuleName, rawNativeMode, emitLineDirectives) : compilationContext;
		compilationContext = context;
		var compiler = new GoCompiler(context);
		if (selectedClasses.length == 0 && selectedEnums.length == 0) {
			generatedFiles = compiler.compileModule(allModules);
		} else {
			generatedFiles = compiler.compileSelectedTypes(selectedClasses, selectedEnums);
		}
	}

	override public function generateFilesManually():Void {
		if (output == null) {
			Context.fatalError("GoReflaxeCompiler output manager is not initialized", Context.currentPos());
			return;
		}

		for (file in generatedFiles) {
			output.saveFile(file.relativePath, file.contents);
		}

		output.saveFile("go.mod", buildGoMod(goModuleName));
		writeRuntime(output, compilationContext);
	}

	override public function onOutputComplete():Void {
		if (output == null || output.outputDir == null) {
			return;
		}

		if (Context.defined("go_no_build") || Context.defined("go_codegen_only")) {
			return;
		}

		var outDir = output.outputDir;
		var goModPath = Path.join([outDir, "go.mod"]);
		if (!FileSystem.exists(goModPath)) {
			return;
		}

		var goCmd = Context.definedValue("go_cmd");
		if (goCmd == null || StringTools.trim(goCmd) == "") {
			goCmd = "go";
		} else {
			goCmd = StringTools.trim(goCmd);
		}

		var args = ["build"];
		var binaryOutput = Context.definedValue("go_build_output");
		if (binaryOutput != null && StringTools.trim(binaryOutput) != "") {
			args.push("-o");
			args.push(StringTools.trim(binaryOutput));
		}
		args.push(".");

		var originalCwd = Sys.getCwd();
		var code = -1;
		var commandLabel = goCmd + " " + args.join(" ");
		try {
			Sys.setCwd(outDir);
			code = Sys.command(goCmd, args);
			Sys.setCwd(originalCwd);
		} catch (err:Dynamic) {
			Sys.setCwd(originalCwd);
			#if eval
			Context.warning("`" + commandLabel + "` failed with exception: " + Std.string(err), Context.currentPos());
			#end
			return;
		}

		if (code != 0) {
			#if eval
			Context.warning("`" + commandLabel + "` failed (exit " + code + ") for output: " + outDir, Context.currentPos());
			#end
		}
	}

	public function generateOutputIterator():Iterator<DataAndFileInfo<StringOrBytes>> {
		var empty:Array<DataAndFileInfo<StringOrBytes>> = [];
		return empty.iterator();
	}

	public function compileClassImpl(classType:ClassType, varFields:Array<ClassVarData>, funcFields:Array<ClassFuncData>):Null<Bool> {
		selectedClasses.push(classType);
		return null;
	}

	public function compileEnumImpl(enumType:EnumType, options:Array<EnumOptionData>):Null<Bool> {
		selectedEnums.push(enumType);
		return null;
	}

	public function compileExpressionImpl(expr:TypedExpr, topLevel:Bool):Null<Dynamic> {
		return null;
	}

	function resolveGoModuleName():String {
		var moduleName = Context.definedValue("go_module");
		if (moduleName == null || StringTools.trim(moduleName) == "") {
			return "snapshot";
		}
		return StringTools.trim(moduleName);
	}

	function buildGoMod(moduleName:String):String {
		return ["module " + moduleName, "", "go 1.22", ""].join("\n");
	}

	function writeRuntime(outputManager:reflaxe.output.OutputManager, context:Null<CompilationContext>):Void {
		var runtimeSource = Path.join([findLibraryRoot(), "runtime", "hxrt"]);
		if (!FileSystem.exists(runtimeSource) || !FileSystem.isDirectory(runtimeSource)) {
			Context.fatalError('Missing runtime directory at "' + runtimeSource + '"', Context.currentPos());
		}

		var plan = resolveRuntimeCopyPlan(context);
		if (context != null) {
			context.selectedHxrtFeatures = plan.features.copy();
		}

		if (plan.fullCopy) {
			writeRuntimeDir(outputManager, runtimeSource, "hxrt");
			return;
		}

		for (fileName in plan.files) {
			var sourcePath = Path.join([runtimeSource, fileName]);
			if (!FileSystem.exists(sourcePath) || FileSystem.isDirectory(sourcePath)) {
				Context.fatalError('Missing runtime file for feature plan: "' + sourcePath + '"', Context.currentPos());
			}
			var targetPath = Path.join(["hxrt", fileName]);
			outputManager.saveFile(targetPath, File.getContent(sourcePath));
		}
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

	function resolveRuntimeCopyPlan(context:Null<CompilationContext>):{fullCopy:Bool, features:Array<String>, files:Array<String>} {
		var forceFullCopy = Context.defined(HXRT_DEFAULT_FEATURES_DEFINE);
		var manualFeaturesDefined = Context.defined(HXRT_FEATURES_DEFINE);
		var noFeatureInfer = Context.defined(HXRT_NO_FEATURE_INFER_DEFINE);
		var selectiveEnabled = manualFeaturesDefined || noFeatureInfer;

		if (forceFullCopy || !selectiveEnabled) {
			return {fullCopy: true, features: [], files: []};
		}

		var manualFeatures = parseManualHxrtFeatures(Context.definedValue(HXRT_FEATURES_DEFINE));
		var selected = manualFeatures.copy();
		if (!noFeatureInfer && context != null) {
			for (feature in context.inferredHxrtFeatures) {
				if (!selected.contains(feature)) {
					selected.push(feature);
				}
			}
		}

		var expanded = GoHxrtFeatureAnalyzer.expandWithDependencies(selected);
		return {
			fullCopy: false,
			features: expanded,
			files: GoHxrtFeatureAnalyzer.filesForFeatures(expanded)
		};
	}

	function parseManualHxrtFeatures(raw:Null<String>):Array<String> {
		if (raw == null) {
			return [];
		}
		var tokens = [for (part in raw.split(",")) StringTools.trim(part).toLowerCase()];
		var out:Array<String> = [];
		for (token in tokens) {
			if (token == "") {
				continue;
			}
			if (!GoHxrtFeatureAnalyzer.isKnownFeature(token)) {
				var expected = GoHxrtFeatureAnalyzer.knownFeatures().join(", ");
				Context.fatalError('Unknown `' + HXRT_FEATURES_DEFINE + '` feature "' + token + '" (expected one of: ' + expected + ")", Context.currentPos());
			}
			if (!out.contains(token)) {
				out.push(token);
			}
		}
		return out;
	}
}
#else
class GoReflaxeCompiler {
	public function new() {}
}
#end
