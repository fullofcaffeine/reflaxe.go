package reflaxe.go;

#if (macro || reflaxe_runtime)
import haxe.io.Path;
import haxe.macro.Context;
import haxe.macro.Type;
import haxe.macro.TypedExprTools;
import reflaxe.GenericCompiler;
import reflaxe.data.ClassFuncData;
import reflaxe.data.ClassVarData;
import reflaxe.data.EnumOptionData;
import reflaxe.go.analyze.GoProfileContractAnalyzer;
import reflaxe.go.analyze.GoProfileContractAnalyzer.PortableNativeScanMode;
import reflaxe.go.analyze.GoNativeBoundaryAnalyzer;
import reflaxe.go.compiler.GoBuildContext;
import reflaxe.go.compiler.GoBuildRequest;
import reflaxe.go.compiler.GoCompilerDefine;
import reflaxe.go.compiler.GoHxrtFeatureAnalyzer;
import reflaxe.go.compiler.GoHxrtFeatureAnalyzer.GoHxrtFeatureReason;
import reflaxe.go.compiler.GoRuntimeCapabilityManifest;
import reflaxe.go.compiler.GoRuntimeCapabilityManifest.GoRuntimeCapabilityManifestSnapshot;
import reflaxe.go.compiler.GoRuntimeCapabilityManifest.GoRuntimeCapabilitySelection;
import reflaxe.go.compiler.GoAutoLoweringModeTools;
import reflaxe.go.compiler.GoBuildContextResolver;
import reflaxe.go.compiler.GoExistingModuleOutputPlan;
import reflaxe.go.compiler.GoExistingModuleOutputTransaction;
import reflaxe.go.compiler.GoGeneratedOutputBoundary;
import reflaxe.go.compiler.GoGeneratedOutputBoundary.GoOutputPathError;
import reflaxe.go.compiler.GoModuleFileGuard;
import reflaxe.go.compiler.GoOutputIdentity;
import reflaxe.go.compiler.GoOutputIdentity.GoGeneratedFileStyle;
import reflaxe.go.compiler.GoPostBuildRunner;
import reflaxe.go.compiler.GoProjectMode;
import reflaxe.go.compiler.GoProjectMode.GoEntrypointSymbol;
import reflaxe.go.compiler.GoProjectModeResolver;
import reflaxe.go.compiler.GoProjectModeError;
import reflaxe.go.compiler.GoSurfaceContractRegistry;
import reflaxe.go.compiler.GoSurfaceContractRegistry.GoSurfaceContractRegistrySnapshot;
import reflaxe.go.compiler.GoSurfaceContractRegistry.GoSurfaceImportRequirement;
import reflaxe.go.compiler.GoSurfacePlanner;
import reflaxe.go.compiler.GoSurfacePlanner.GoSurfacePlanDecision;
import reflaxe.go.compiler.GoTypeUsageLedger;
import reflaxe.go.compiler.GoTypeUsageLedger.GoImmutableList;
import reflaxe.go.compiler.GoTypeUsageLedger.GoTypeUsageLedgerSnapshot;
import reflaxe.output.DataAndFileInfo;
import reflaxe.output.StringOrBytes;
import sys.FileSystem;

private typedef ContractReportSnapshot = {
	final schemaVersion:Int;
	final contract:String;
	final policyPreset:String;
	final semanticBoundarySource:String;
	final nativeAuthorityPolicy:String;
	final nativeAuthorityPolicySource:String;
	final nativeSpecializationPolicy:String;
	final nativeSpecializationPolicySource:String;
	final nativeFallbackPolicy:String;
	final nativeFallbackPolicySource:String;
	final autoLoweringMode:String;
	final strictExamples:Bool;
	final strictUserBoundaryPolicy:String;
	final strictUserBoundaries:Bool;
	final metalFallbackAllowed:Bool;
	final metalContractHardError:Bool;
	final emitLineDirectives:Bool;
	final rawNativeMode:String;
	final hxrtSelectiveEnabled:Bool;
	final hxrtForceFullCopy:Bool;
	final hxrtNoFeatureInfer:Bool;
	final hxrtManualFeatures:Array<String>;
	final nativeBoundaryModules:Array<String>;
	final metalLaneModules:Array<String>;
	final nativeFallbackEventCount:Int;
	final nativeFallbackBoundaryEventCount:Int;
	final nativeFallbackNonBoundaryEventCount:Int;
	final metalFallbackViolationCount:Int;
	final metalFallbackLaneViolationCount:Int;
	final metalFallbackNonLaneViolationCount:Int;
	final portableNativeImportScanMode:String;
	final portableNativeImportHitCount:Int;
	final portableNativeImportTypedHitCount:Int;
	final portableNativeImportScannerHitCount:Int;
	final portableNativeImportHits:Array<String>;
	final portableNativeImportTypedHits:Array<String>;
	final portableNativeImportScannerHits:Array<String>;
	final contractDiagnosticCount:Int;
	final contractDiagnostics:Array<ContractDiagnosticEntry>;
	final loweringDecisionCount:Int;
	final loweringDecisionAttemptCount:Int;
	final loweringDecisionSuccessCount:Int;
	final loweringDecisionFallbackCount:Int;
	final loweringDecisions:Array<ContractLoweringDecision>;
	final nativeFallbackEventsByModule:Array<ContractFallbackModuleSummary>;
	final nativeFallbackEvents:Array<ContractFallbackEvent>;
	final metalFallbackViolationsByModule:Array<ContractFallbackModuleSummary>;
	final metalFallbackViolations:Array<ContractFallbackEvent>;
}

private typedef ContractLoweringDecision = {
	final feature:String;
	final kind:String;
	final outcome:String;
	final detail:String;
	final location:String;
	final module:String;
	final inNativeBoundary:Bool;
}

private typedef ContractDiagnosticEntry = {
	final code:String;
	final severity:String;
	final module:String;
	final location:String;
	final message:String;
}

private typedef ContractFallbackEvent = {
	final kind:String;
	final detail:String;
	final location:String;
	final module:String;
	final inNativeBoundary:Bool;
}

private typedef ContractFallbackModuleSummary = {
	final module:String;
	final inNativeBoundary:Bool;
	final count:Int;
}

private typedef RuntimeFeatureReason = {
	final feature:String;
	final sourceKind:String;
	final source:String;
}

private typedef RuntimePlanReportSnapshot = {
	final schemaVersion:Int;
	final contract:String;
	final policyPreset:String;
	final semanticBoundarySource:String;
	final mode:String;
	final selectiveEnabled:Bool;
	final fullCopy:Bool;
	final inferenceDisabled:Bool;
	final manualFeatures:Array<String>;
	final inferredFeatures:Array<String>;
	final selectedFeatures:Array<String>;
	final files:Array<String>;
	final reasons:Array<RuntimeFeatureReason>;
	final manifestAuthority:String;
	final capabilities:Array<GoRuntimeCapabilitySelection>;
	final surfacePlanAuthority:String;
	final surfacePlanDecisionCount:Int;
	final requiredSurfaceImports:Array<GoSurfaceImportRequirement>;
	final requiredSurfaceRuntimeFeatures:Array<String>;
	final surfacePlans:Array<GoSurfacePlanDecision>;
}

private typedef OptimizerPlanReportSnapshot = {
	final schemaVersion:Int;
	final contract:String;
	final policyPreset:String;
	final nativeSpecializationPolicy:String;
	final nativeSpecializationPolicySource:String;
	final autoLoweringMode:String;
	final optimizationPreset:String;
	final portableStringFastpathEnabled:Bool;
	final portableConcurrencyFastpathEnabled:Bool;
	final goAstPassSelectionSource:String;
	final goAstPasses:Array<String>;
	final goAstPassSelectionReasons:Array<OptimizerPassSelectionReason>;
	final stringInstanceTypedLowerings:Int;
	final stringInstanceLegacyLowerings:Int;
	final stringLengthFieldTypedLowerings:Int;
	final stringLengthFieldLegacyLowerings:Int;
	final portableConcurrencyTypedFastpathHits:Int;
	final portableConcurrencyTypedFastpathFallbacks:Int;
	final goCollectionsTypedLowerings:Int;
	final goCollectionsTypedFallbacks:Int;
	final goResultTypedLowerings:Int;
	final goResultTypedFallbacks:Int;
	final loweringFallbackBoundaryCount:Int;
	final loweringFallbackNonBoundaryCount:Int;
	final loweringFallbackLaneCount:Int;
	final loweringFallbackNonLaneCount:Int;
	final autoLoweringCapabilities:Array<OptimizerCapabilitySummary>;
	final surfacePlanAuthority:String;
	final surfacePlanDecisionCount:Int;
	final requiredSurfaceImports:Array<GoSurfaceImportRequirement>;
	final requiredSurfaceRuntimeFeatures:Array<String>;
	final surfacePlans:Array<GoSurfacePlanDecision>;
}

private typedef OptimizerPassSelectionReason = {
	final pass:String;
	final reason:String;
	final source:String;
}

private typedef OptimizerCapabilitySummary = {
	final id:String;
	final attempts:Int;
	final successes:Int;
	final fallbacks:Int;
	final fallbackReasonCounts:Array<OptimizerCapabilityFallbackReasonCount>;
}

private typedef OptimizerCapabilityFallbackReasonCount = {
	final kind:String;
	final count:Int;
}

private typedef OptimizerCapabilityAccumulator = {
	var attempts:Int;
	var successes:Int;
	var fallbacks:Int;
	var fallbackReasonCounts:Map<String, Int>;
}

/**
	Typed marker for Reflaxe's legacy per-node output hooks.

	Why
	haxe.go stages typed modules and emits them together, so these hooks always
	return `null`. Using `Dynamic` for their unused generic result made that
	implementation detail look like an untyped compiler boundary.

	What
	Names the deliberately unused output slot required by `GenericCompiler`.

	How
	All five generic hook result types use this marker while the real generated
	files continue to flow through `generateOutputIterator()`.
**/
enum GoReflaxeStagedOutput {
	StagedOutput;
}

class GoReflaxeCompiler extends GenericCompiler<GoReflaxeStagedOutput, GoReflaxeStagedOutput, GoReflaxeStagedOutput, GoReflaxeStagedOutput,
	GoReflaxeStagedOutput> {
	var allModules:Array<ModuleType> = [];
	var selectedClasses:Array<ClassType> = [];
	var selectedEnums:Array<EnumType> = [];
	var generatedFiles:Array<GoCompiler.GoGeneratedFile> = [];
	var buildContext:Null<GoBuildContext> = null;
	var compilationContext:Null<CompilationContext> = null;
	var lastRuntimePlan:Null<GoRuntimeCapabilityManifestSnapshot> = null;
	var outputBoundary:Null<GoGeneratedOutputBoundary> = null;
	var existingModuleOutputPlan:Null<GoExistingModuleOutputPlan> = null;
	var projectMode:GoProjectMode = GoProjectMode.Standalone;
	var moduleFileGuard:Null<GoModuleFileGuard> = null;
	var typeUsageLedger:GoTypeUsageLedger = new GoTypeUsageLedger();
	var lastTypeUsageReport:GoTypeUsageLedgerSnapshot = GoTypeUsageLedger.emptySnapshot();
	var lastSurfaceContractReport:GoSurfaceContractRegistrySnapshot = GoSurfaceContractRegistry.emptySnapshot();

	public function new() {
		super();
	}

	override public function filterTypes(moduleTypes:Array<ModuleType>):Array<ModuleType> {
		allModules = moduleTypes.copy();
		return moduleTypes;
	}

	override public function onCompileStart():Void {
		compilationContext = null;
		selectedClasses = [];
		selectedEnums = [];
		generatedFiles = [];
		lastRuntimePlan = null;
		outputBoundary = null;
		existingModuleOutputPlan = null;
		projectMode = GoProjectMode.Standalone;
		moduleFileGuard = null;
		typeUsageLedger = new GoTypeUsageLedger();
		lastTypeUsageReport = GoTypeUsageLedger.emptySnapshot();
		lastSurfaceContractReport = GoSurfaceContractRegistry.emptySnapshot();
		try {
			projectMode = GoProjectModeResolver.resolve();
		} catch (error:GoProjectModeError) {
			Context.fatalError(error.message, Context.currentPos());
			return;
		} catch (error:GoOutputPathError) {
			Context.fatalError(error.message, Context.currentPos());
			return;
		}
		var resolvedBuildContext = GoBuildContextResolver.resolve();
		switch (projectMode) {
			case Standalone:
			case ExistingModule(project):
				resolvedBuildContext = resolvedBuildContext.withGoModuleName(project.modulePath);
				if (output == null) {
					Context.fatalError("GoReflaxeCompiler output manager is not initialized", Context.currentPos());
					return;
				}
				// `go_output` remains the package assertion. The managed writer is
				// rooted at the caller module so package and runtime paths share one
				// confined, module-relative namespace.
				output.setOutputDir(project.moduleRoot);
		}
		buildContext = resolvedBuildContext;
	}

	override public function onCompileEnd():Void {
		var resolvedBuildContext = effectiveBuildContext();
		var boundarySnapshot = GoNativeBoundaryAnalyzer.collect(allModules);
		resolvedBuildContext = resolvedBuildContext.withNativeBoundaryModules(boundarySnapshot.modules);
		buildContext = resolvedBuildContext;
		var outputIdentity = resolveOutputIdentity(resolvedBuildContext);
		var runtimeImportPath = outputIdentity.runtimeImportPath;
		var authoritySnapshot = typeUsageLedger.snapshot([], runtimeImportPath);
		var surfaceContractSnapshot = GoSurfaceContractRegistry.defaultRegistry().snapshot(authoritySnapshot);
		var surfacePlan = GoSurfacePlanner.plan(resolvedBuildContext, surfaceContractSnapshot);
		var context = CompilationContext.fromBuildContext(resolvedBuildContext, authoritySnapshot, surfaceContractSnapshot, surfacePlan, runtimeImportPath);
		compilationContext = context;
		lastSurfaceContractReport = surfaceContractSnapshot;
		var compiler = new GoCompiler(context, resolveSelectedMainIdentity(), outputIdentity);
		retainSuppliedPortableFacadeEnums();
		if (selectedClasses.length == 0 && selectedEnums.length == 0) {
			generatedFiles = compiler.compileModule(allModules);
		} else {
			generatedFiles = compiler.compileSelectedTypes(selectedClasses, selectedEnums);
		}
		lastTypeUsageReport = typeUsageLedger.snapshot(context.inferredHxrtFeatureReasons, context.runtimeImportPath);
	}

	function resolveOutputIdentity(buildContext:GoBuildContext):GoOutputIdentity {
		return switch (projectMode) {
			case Standalone:
				GoOutputIdentity.standalone(buildContext.goModuleName + "/hxrt");
			case ExistingModule(project):
				final entrySymbol = switch (project.entrypoint) {
					case CompilerMain: GoEntrypointSymbol.named("main");
					case CallerBridge(symbol): symbol;
				};
				new GoOutputIdentity({
					packageName: project.packageName,
					entrySymbol: entrySymbol,
					runtimeImportPath: project.runtimeImportPath(),
					fileStyle: GoGeneratedFileStyle.ExistingModuleFiles
				});
		};
	}

	/**
		Retains exact fixture- or package-supplied portable facade declarations.

		Why
		Reflaxe treats its own `reflaxe.*` compiler namespace as framework code, so
		its normal selected-type callbacks can omit externally supplied
		`reflaxe.std` declarations even when user expressions reference their
		constructors. Without exact retention, generated calls have no declaration.

		What
		Adds only `reflaxe.std.Option` and `reflaxe.std.Result` enum declarations
		that Haxe actually included in the typed module set.

		How
		Match exact enum paths, deduplicate against normal callback selection, and
		leave every other `reflaxe.*` module untouched. This does not bundle those
		source modules or infer authority from a namespace prefix.
	**/
	function retainSuppliedPortableFacadeEnums():Void {
		var selected = new Map<String, Bool>();
		for (enumType in selectedEnums) {
			selected.set(enumType.pack.concat([enumType.name]).join("."), true);
		}
		for (moduleType in Context.getAllModuleTypes()) {
			switch (moduleType) {
				case TEnumDecl(enumRef):
					var enumType = enumRef.get();
					var path = enumType.pack.concat([enumType.name]).join(".");
					if ((path == "reflaxe.std.Option" || path == "reflaxe.std.Result") && !selected.exists(path)) {
						selectedEnums.push(enumType);
						selected.set(path, true);
					}
				case _:
			}
		}
	}

	/**
		Why
		The Go backend must not infer its executable entrypoint from a class naming
		convention or from unrelated static methods named `main`. Thread-aware Haxe
		programs can also wrap the direct selected-main call in the typed main
		expression.

		What
		Resolves the exact Haxe-selected main class into the stable class/module
		identity consumed by Go lowering and file layout.

		How
		Uses Reflaxe's direct `getMainModule()` result first. If Haxe wrapped the main
		expression, walks only that authoritative expression for its referenced static
		`main`; it never scans arbitrary project classes for candidates.
	**/
	function resolveSelectedMainIdentity():GoCompiler.GoMainIdentity {
		var directMain = getMainModule();
		switch (directMain) {
			case TClassDecl(classRef):
				return mainIdentityForClass(classRef.get());
			case _:
		}

		var mainExpr = getMainExpr();
		var selectedClass:Null<ClassType> = null;
		if (mainExpr != null) {
			function visit(expr:TypedExpr):Void {
				if (selectedClass != null) {
					return;
				}
				switch (expr.expr) {
					case TField(_, FStatic(classRef, fieldRef)) if (fieldRef.get().name == "main"):
						selectedClass = classRef.get();
					case _:
				}
				TypedExprTools.iter(expr, visit);
			}
			visit(mainExpr);
		}

		if (selectedClass == null) {
			Context.fatalError("Haxe-selected main class could not be resolved from the typed main expression", Context.currentPos());
			return {className: "", moduleName: ""};
		}
		return mainIdentityForClass(selectedClass);
	}

	function mainIdentityForClass(classType:ClassType):GoCompiler.GoMainIdentity {
		var packageName = classType.pack.join(".");
		return {
			className: packageName == "" ? classType.name : packageName + "." + classType.name,
			moduleName: classType.module
		};
	}

	/**
		What: Establishes output confinement before Reflaxe starts any write or
		managed-file deletion.

		Why: The generic output manager accepts arbitrary strings and trusts its old
		generated-file inventory. Both sources must be proven safe before generation.

		How: Resolve the configured root, preflight metadata and extra-file keys, then
		run the normal Reflaxe lifecycle while translating only typed path failures to
		a stable, machine-path-free compiler diagnostic.
	**/
	override public function generateFiles():Void {
		if (output == null) {
			Context.fatalError("GoReflaxeCompiler output manager is not initialized", Context.currentPos());
			return;
		}
		var configuredRoot = output.outputDir;
		if (configuredRoot == null) {
			Context.fatalError("GoReflaxeCompiler output directory is not initialized", Context.currentPos());
			return;
		}

		try {
			switch (projectMode) {
				case Standalone:
					final boundary = new GoGeneratedOutputBoundary(configuredRoot);
					boundary.validateManagedFileMetadata();
					for (path in extraFiles.keys()) {
						boundary.validateDestination(path);
					}
					outputBoundary = boundary;
					super.generateFiles();
				case ExistingModule(project):
					moduleFileGuard = new GoModuleFileGuard(project.moduleRoot);
					final boundary = new GoGeneratedOutputBoundary(configuredRoot, ["go.mod", "go.sum"]);
					outputBoundary = boundary;
					final plan = new GoExistingModuleOutputPlan(project, boundary);
					existingModuleOutputPlan = plan;
					generateFilesManually();
					collectExistingModuleExtraFiles(plan);
					existingModuleOutputPlan = null;
					new GoExistingModuleOutputTransaction(project, boundary).commit(plan);
			}
			if (moduleFileGuard != null) {
				moduleFileGuard.verify();
			}
		} catch (error:GoOutputPathError) {
			if (moduleFileGuard != null) {
				try {
					moduleFileGuard.verify();
				} catch (verificationError:GoOutputPathError) {
					error = verificationError;
				}
			}
			Context.fatalError(error.message, Context.currentPos());
		}
	}

	/** Preserve Reflaxe's priority ordering while adding macro extra files to one plan. */
	function collectExistingModuleExtraFiles(plan:GoExistingModuleOutputPlan):Void {
		for (path in extraFiles.keys()) {
			final content = extraFiles.get(path);
			if (content == null) {
				continue;
			}
			final priorities:Array<Int> = [];
			for (priority in content.keys()) {
				final fragment = content.get(priority);
				if (fragment != null && StringTools.trim(fragment).length > 0) {
					priorities.push(priority);
				}
			}
			priorities.sort((left, right) -> left - right);
			final fragments:Array<String> = [];
			for (priority in priorities) {
				final fragment = content.get(priority);
				if (fragment != null && StringTools.trim(fragment).length > 0) {
					fragments.push(fragment);
				}
			}
			plan.add(path, fragments.join("\n\n"));
		}
	}

	override public function generateFilesManually():Void {
		if (output == null || outputBoundary == null) {
			Context.fatalError("GoReflaxeCompiler output boundary is not initialized", Context.currentPos());
			return;
		}

		for (file in generatedFiles) {
			saveGeneratedFile(generatedSourcePath(file.relativePath), file.contents);
		}

		var resolvedBuildContext = effectiveBuildContext();
		switch (projectMode) {
			case Standalone:
				saveGeneratedFile("go.mod", buildGoMod(resolvedBuildContext.goModuleName));
			case ExistingModule(_):
		}
		writeGeneratedLicenseMaterial();
		writeRuntime(compilationContext, resolvedBuildContext);
		emitBuildReports(compilationContext, resolvedBuildContext);
		switch (projectMode) {
			case Standalone:
			case ExistingModule(project):
				switch (project.build) {
					case NoBuild:
					case GoBuild(request):
						saveGeneratedFile(GoBuildRequest.REPORT_PATH, request.invocation().renderJson());
				}
		}
	}

	function generatedSourcePath(fileName:String):String {
		return switch (projectMode) {
			case Standalone: fileName;
			case ExistingModule(project): project.generatedSourcePath(fileName);
		};
	}

	/**
		What: Saves one compiler-generated artifact through the established boundary.

		Why: Code, module metadata, and reports must not call Reflaxe's permissive
		writer directly.

		How: Require the initialized manager and boundary, then pass the untrusted
		relative spelling to the typed validator immediately before the managed write.
	**/
	function saveGeneratedFile(path:String, content:StringOrBytes):Void {
		if (output == null || outputBoundary == null) {
			throw new GoOutputPathError(InvalidRoot, "the compiler output boundary is unavailable");
		}
		if (existingModuleOutputPlan != null) {
			existingModuleOutputPlan.add(path, content);
		} else {
			outputBoundary.saveFile(output, path, content);
		}
	}

	/**
		What: Copies one packaged runtime file through the established boundary.

		Why: Directory-entry names and feature-plan paths are output-path inputs even
		when the runtime source tree is repository controlled.

		How: Let the boundary validate and prepare the destination before it reads and
		saves the support file through Reflaxe's managed writer.
	**/
	function copyGeneratedFile(sourcePath:String, targetPath:String):Void {
		if (output == null || outputBoundary == null) {
			throw new GoOutputPathError(InvalidRoot, "the compiler output boundary is unavailable");
		}
		if (existingModuleOutputPlan != null) {
			try {
				existingModuleOutputPlan.addBytes(targetPath, sys.io.File.getBytes(sourcePath));
			} catch (error:GoOutputPathError) {
				throw error;
			} catch (_:haxe.Exception) {
				throw new GoOutputPathError(WriteFailed, "a compiler support file could not be read");
			}
		} else {
			outputBoundary.copyManagedFile(output, sourcePath, targetPath);
		}
	}

	override public function onOutputComplete():Void {
		if (output == null || output.outputDir == null) {
			return;
		}

		switch (projectMode) {
			case Standalone:
			case ExistingModule(project):
				if (moduleFileGuard != null) {
					try {
						moduleFileGuard.verify();
					} catch (error:GoOutputPathError) {
						Context.fatalError(error.message, Context.currentPos());
					}
				}
				switch (project.build) {
					case NoBuild:
						return;
					case GoBuild(request):
						final invocation = request.invocation();
						final result = GoPostBuildRunner.runGoverned(project.moduleRoot, invocation.command, invocation.arguments,
							invocation.environment.processEntries());
						if (moduleFileGuard != null) {
							try {
								moduleFileGuard.verify();
							} catch (error:GoOutputPathError) {
								Context.fatalError(error.message, Context.currentPos());
							}
						}
						switch (result) {
							case BuildSucceeded:
							case BuildFailed(_):
								Context.fatalError(GoPostBuildRunner.failureMessage(invocation.command, invocation.arguments, result), Context.currentPos());
						}
						return;
				}
		}

		if (Context.defined(GoCompilerDefine.DefineGoNoBuild) || Context.defined(GoCompilerDefine.DefineGoCodegenOnly)) {
			return;
		}

		var outDir = output.outputDir;
		var goModPath = Path.join([outDir, "go.mod"]);
		if (!FileSystem.exists(goModPath)) {
			return;
		}

		var goCmd = Context.definedValue(GoCompilerDefine.DefineGoCommand);
		if (goCmd == null || StringTools.trim(goCmd) == "") {
			goCmd = "go";
		} else {
			goCmd = StringTools.trim(goCmd);
		}

		var args = ["build"];
		var binaryOutput = Context.definedValue(GoCompilerDefine.DefineGoBuildOutput);
		if (binaryOutput != null && StringTools.trim(binaryOutput) != "") {
			args.push("-o");
			args.push(StringTools.trim(binaryOutput));
		}
		args.push(".");

		var result = GoPostBuildRunner.run(outDir, goCmd, args);
		switch (result) {
			case BuildSucceeded:
			case BuildFailed(_):
				Context.fatalError(GoPostBuildRunner.failureMessage(goCmd, args, result), Context.currentPos());
		}
	}

	public function generateOutputIterator():Iterator<DataAndFileInfo<StringOrBytes>> {
		var empty:Array<DataAndFileInfo<StringOrBytes>> = [];
		return empty.iterator();
	}

	public function compileClassImpl(classType:ClassType, varFields:Array<ClassVarData>, funcFields:Array<ClassFuncData>):Null<GoReflaxeStagedOutput> {
		var currentModule = getCurrentModule();
		if (currentModule != null) {
			typeUsageLedger.collect(currentModule, getTypeUsage());
		}
		selectedClasses.push(classType);
		return null;
	}

	public function compileEnumImpl(enumType:EnumType, options:Array<EnumOptionData>):Null<GoReflaxeStagedOutput> {
		var currentModule = getCurrentModule();
		if (currentModule != null) {
			typeUsageLedger.collect(currentModule, getTypeUsage());
		}
		selectedEnums.push(enumType);
		return null;
	}

	override public function compileTypedefImpl(typedefType:DefType):Null<GoReflaxeStagedOutput> {
		var currentModule = getCurrentModule();
		if (currentModule != null) {
			typeUsageLedger.collect(currentModule, getTypeUsage());
		}
		return null;
	}

	override public function compileAbstractImpl(abstractType:AbstractType):Null<GoReflaxeStagedOutput> {
		var currentModule = getCurrentModule();
		if (currentModule != null) {
			typeUsageLedger.collect(currentModule, getTypeUsage());
		}
		return null;
	}

	public function compileExpressionImpl(expr:TypedExpr, topLevel:Bool):Null<GoReflaxeStagedOutput> {
		return null;
	}

	function buildGoMod(moduleName:String):String {
		return ["module " + moduleName, "", "go 1.22", ""].join("\n");
	}

	/**
		What: Copies the two approved license texts into every generated Go module.

		Why: Generated files can combine user-owned code with project-owned runtime
		or support portions and Haxe standard-library-derived portions. Keeping the
		notices beside the generated module makes those permissions travel with the
		code without assigning a license to the user's own work.

		How: Resolve the immutable packaged license sources and copy their exact bytes
		through the same confined, managed output boundary as runtime support files.
	**/
	function writeGeneratedLicenseMaterial():Void {
		var libraryRoot = findLibraryRoot();
		copyRequiredGeneratedMaterial(libraryRoot, "licenses/HAXE-GO-GENERATED-MIT.txt", "LICENSES/HAXE-GO-GENERATED-MIT.txt");
		copyRequiredGeneratedMaterial(libraryRoot, "licenses/HAXE-STDLIB-MIT.txt", "LICENSES/HAXE-STDLIB-MIT.txt");
	}

	/**
		What: Copies one required generated-project legal artifact.

		Why: A missing packaged notice must fail compilation clearly instead of
		producing a generated tree whose redistribution terms are incomplete.

		How: Validate the trusted package source, then delegate the destination write
		to the confined generated-output boundary.
	**/
	function copyRequiredGeneratedMaterial(libraryRoot:String, sourceRelativePath:String, targetRelativePath:String):Void {
		var sourcePath = Path.join([libraryRoot, sourceRelativePath]);
		if (!FileSystem.exists(sourcePath) || FileSystem.isDirectory(sourcePath)) {
			Context.fatalError('Missing packaged generated-output license material: "' + sourceRelativePath + '"', Context.currentPos());
		}
		copyGeneratedFile(sourcePath, targetRelativePath);
	}

	function writeRuntime(context:Null<CompilationContext>, buildContext:GoBuildContext):Void {
		var runtimeSource = Path.join([findLibraryRoot(), "runtime", "hxrt"]);
		if (!FileSystem.exists(runtimeSource) || !FileSystem.isDirectory(runtimeSource)) {
			Context.fatalError("Missing packaged hxrt runtime directory", Context.currentPos());
		}

		var plan = resolveRuntimeCopyPlan(context, buildContext);
		lastRuntimePlan = plan;
		if (context != null) {
			context.selectedHxrtFeatures = [for (feature in plan.selectedFeatures) feature];
		}

		for (fileName in plan.files) {
			var sourcePath = Path.join([runtimeSource, fileName]);
			if (!FileSystem.exists(sourcePath) || FileSystem.isDirectory(sourcePath)) {
				Context.fatalError('Missing packaged hxrt file required by the feature plan: "' + fileName + '"', Context.currentPos());
			}
			var targetPath = switch (projectMode) {
				case Standalone: Path.join(["hxrt", fileName]);
				case ExistingModule(project): project.runtimePath(fileName);
			};
			copyGeneratedFile(sourcePath, targetPath);
		}
	}

	function findLibraryRoot():String {
		var thisFile = Context.resolvePath("reflaxe/go/GoReflaxeCompiler.hx");
		var srcDir = Path.normalize(Path.directory(thisFile));
		return Path.normalize(Path.join([srcDir, "..", "..", ".."]));
	}

	function effectiveBuildContext():GoBuildContext {
		if (buildContext != null) {
			return buildContext;
		}
		var resolved = GoBuildContextResolver.resolve();
		buildContext = resolved;
		return resolved;
	}

	function resolveRuntimeCopyPlan(context:Null<CompilationContext>, buildContext:GoBuildContext):GoRuntimeCapabilityManifestSnapshot {
		var inferredFeatures = context == null ? [] : context.inferredHxrtFeatures;
		var inferredReasons = context == null ? [] : context.inferredHxrtFeatureReasons;
		return GoRuntimeCapabilityManifest.build(buildContext, inferredFeatures, inferredReasons);
	}

	function emitBuildReports(context:Null<CompilationContext>, buildContext:GoBuildContext):Void {
		if (buildContext.contractReportEnabled) {
			var contractSnapshot = buildContractReportSnapshot(buildContext, context);
			saveGeneratedFile("profile_contract.json", renderContractReportJson(contractSnapshot));
			saveGeneratedFile("profile_contract.md", renderContractReportMarkdown(contractSnapshot));
		}

		if (buildContext.runtimePlanReportEnabled) {
			var runtimeSnapshot = buildRuntimePlanReportSnapshot(buildContext, context);
			saveGeneratedFile("hxrt_plan.json", renderRuntimePlanJson(runtimeSnapshot));
			saveGeneratedFile("hxrt_plan.md", renderRuntimePlanMarkdown(runtimeSnapshot));
		}

		if (buildContext.optimizerPlanReportEnabled) {
			var optimizerSnapshot = buildOptimizerPlanReportSnapshot(buildContext, context);
			saveGeneratedFile("optimizer_plan.json", renderOptimizerPlanJson(optimizerSnapshot));
			saveGeneratedFile("optimizer_plan.md", renderOptimizerPlanMarkdown(optimizerSnapshot));
		}

		if (buildContext.typeUsageReportEnabled) {
			saveGeneratedFile("type_usage.json", GoTypeUsageLedger.renderJson(lastTypeUsageReport));
			saveGeneratedFile("type_usage.md", GoTypeUsageLedger.renderMarkdown(lastTypeUsageReport));
		}

		if (buildContext.surfaceContractReportEnabled) {
			saveGeneratedFile("surface_contracts.json", GoSurfaceContractRegistry.renderJson(lastSurfaceContractReport));
			saveGeneratedFile("surface_contracts.md", GoSurfaceContractRegistry.renderMarkdown(lastSurfaceContractReport));
		}
	}

	function buildContractReportSnapshot(buildContext:GoBuildContext, context:Null<CompilationContext>):ContractReportSnapshot {
		var contractLabel = buildContext.profile == GoProfile.Metal ? "metal" : "portable";
		var manualFeatures = sortedUniqueStrings(buildContext.hxrtManualFeatures.copy());
		var boundaryModules = sortedUniqueStrings(buildContext.nativeBoundaryModules.copy());
		var contractDiagnostics = new Array<ContractDiagnosticEntry>();
		var portableNativeImportHits = new Array<String>();
		var portableNativeImportTypedHits = new Array<String>();
		var portableNativeImportScannerHits = new Array<String>();
		var nativePolicy = GoProfileContractAnalyzer.resolvePortableNativePolicyModeFromDefines();
		var nativeScanMode = GoProfileContractAnalyzer.resolvePortableNativeScanModeFromDefines();
		var nativeAllowPrefixes = GoProfileContractAnalyzer.resolvePortableNativeAllowPrefixesFromDefines();
		var analyzed = GoProfileContractAnalyzer.analyze(allModules, buildContext, Sys.getCwd(), nativePolicy, nativeScanMode, nativeAllowPrefixes);
		portableNativeImportHits = analyzed.portableNativeImportHits.copy();
		portableNativeImportTypedHits = analyzed.portableNativeImportTypedHits.copy();
		portableNativeImportScannerHits = analyzed.portableNativeImportScannerHits.copy();
		for (entry in analyzed.diagnostics) {
			if (entry == null) {
				continue;
			}
			contractDiagnostics.push({
				code: entry.code,
				severity: entry.severity,
				module: entry.module,
				location: entry.location,
				message: entry.message
			});
		}
		contractDiagnostics.sort(compareContractDiagnostics);
		var fallbackEvents = new Array<ContractFallbackEvent>();
		var loweringDecisions = new Array<ContractLoweringDecision>();
		var boundaryEventCount = 0;
		var nonBoundaryEventCount = 0;
		var boundaryCountsByModule = new Map<String, Int>();
		var nonBoundaryCountsByModule = new Map<String, Int>();
		var loweringAttemptCount = 0;
		var loweringSuccessCount = 0;
		var loweringFallbackCount = 0;
		if (context != null) {
			for (entry in context.loweringDecisionLedger) {
				if (entry == null) {
					continue;
				}
				switch (entry.outcome) {
					case "attempted":
						loweringAttemptCount++;
					case "succeeded":
						loweringSuccessCount++;
					case "fallback":
						loweringFallbackCount++;
					case _:
				}
				loweringDecisions.push({
					feature: entry.feature,
					kind: entry.kind,
					outcome: entry.outcome,
					detail: entry.detail,
					location: entry.location,
					module: entry.module,
					inNativeBoundary: entry.inNativeBoundary
				});
			}
			for (event in context.nativeFallbackEvents) {
				if (event == null) {
					continue;
				}
				if (event.inNativeBoundary) {
					boundaryEventCount++;
					incrementIntMap(boundaryCountsByModule, event.module);
				} else {
					nonBoundaryEventCount++;
					incrementIntMap(nonBoundaryCountsByModule, event.module);
				}
				fallbackEvents.push({
					kind: event.kind,
					detail: event.detail,
					location: event.location,
					module: event.module,
					inNativeBoundary: event.inNativeBoundary
				});
			}
		}
		loweringDecisions.sort(compareContractLoweringDecision);
		fallbackEvents.sort(compareContractFallbackEvents);
		var fallbackSummary = buildContractFallbackModuleSummary(boundaryCountsByModule, nonBoundaryCountsByModule);
		var legacyMetalFallbackViolations:Array<ContractFallbackEvent> = buildContext.usesMetalCompatibilityPreset() ? fallbackEvents.copy() : [];
		var legacyMetalFallbackSummary:Array<ContractFallbackModuleSummary> = buildContext.usesMetalCompatibilityPreset() ? fallbackSummary.copy() : [];
		var legacyMetalBoundaryViolationCount = buildContext.usesMetalCompatibilityPreset() ? boundaryEventCount : 0;
		var legacyMetalNonBoundaryViolationCount = buildContext.usesMetalCompatibilityPreset() ? nonBoundaryEventCount : 0;
		return {
			schemaVersion: 8,
			contract: contractLabel,
			policyPreset: buildContext.policyPreset.label(),
			semanticBoundarySource: buildContext.semanticBoundarySource.label(),
			nativeAuthorityPolicy: buildContext.nativeAuthorityPolicy.label(),
			nativeAuthorityPolicySource: buildContext.nativeAuthorityPolicySource.label(),
			nativeSpecializationPolicy: buildContext.nativeSpecializationPolicy.label(),
			nativeSpecializationPolicySource: buildContext.nativeSpecializationPolicySource.label(),
			nativeFallbackPolicy: buildContext.nativeFallbackPolicy.label(),
			nativeFallbackPolicySource: buildContext.nativeFallbackPolicySource.label(),
			autoLoweringMode: GoAutoLoweringModeTools.label(buildContext.autoLoweringMode),
			strictExamples: buildContext.strictExamples,
			strictUserBoundaryPolicy: buildContext.strictUserBoundaryPolicy,
			strictUserBoundaries: buildContext.strictUserBoundaries,
			metalFallbackAllowed: buildContext.metalFallbackAllowed,
			metalContractHardError: buildContext.metalContractHardError,
			emitLineDirectives: buildContext.emitLineDirectives,
			rawNativeMode: RawNativeModeResolver.label(buildContext.rawNativeMode),
			hxrtSelectiveEnabled: buildContext.isHxrtSelectiveEnabled(),
			hxrtForceFullCopy: buildContext.hxrtForceFullCopy,
			hxrtNoFeatureInfer: buildContext.hxrtNoFeatureInfer,
			hxrtManualFeatures: manualFeatures,
			nativeBoundaryModules: boundaryModules,
			metalLaneModules: boundaryModules.copy(),
			loweringDecisionCount: loweringDecisions.length,
			loweringDecisionAttemptCount: loweringAttemptCount,
			loweringDecisionSuccessCount: loweringSuccessCount,
			loweringDecisionFallbackCount: loweringFallbackCount,
			loweringDecisions: loweringDecisions,
			nativeFallbackEventCount: fallbackEvents.length,
			nativeFallbackBoundaryEventCount: boundaryEventCount,
			nativeFallbackNonBoundaryEventCount: nonBoundaryEventCount,
			metalFallbackViolationCount: legacyMetalFallbackViolations.length,
			metalFallbackLaneViolationCount: legacyMetalBoundaryViolationCount,
			metalFallbackNonLaneViolationCount: legacyMetalNonBoundaryViolationCount,
			portableNativeImportScanMode: portableNativeScanModeLabel(nativeScanMode),
			portableNativeImportHitCount: portableNativeImportHits.length,
			portableNativeImportTypedHitCount: portableNativeImportTypedHits.length,
			portableNativeImportScannerHitCount: portableNativeImportScannerHits.length,
			portableNativeImportHits: portableNativeImportHits,
			portableNativeImportTypedHits: portableNativeImportTypedHits,
			portableNativeImportScannerHits: portableNativeImportScannerHits,
			contractDiagnosticCount: contractDiagnostics.length,
			contractDiagnostics: contractDiagnostics,
			nativeFallbackEventsByModule: fallbackSummary,
			nativeFallbackEvents: fallbackEvents,
			metalFallbackViolationsByModule: legacyMetalFallbackSummary,
			metalFallbackViolations: legacyMetalFallbackViolations
		};
	}

	function buildRuntimePlanReportSnapshot(buildContext:GoBuildContext, context:Null<CompilationContext>):RuntimePlanReportSnapshot {
		var contractLabel = buildContext.profile == GoProfile.Metal ? "metal" : "portable";
		var plan = lastRuntimePlan == null ? resolveRuntimeCopyPlan(context, buildContext) : lastRuntimePlan;
		var selectedFeatures = [for (feature in plan.selectedFeatures) cast(feature, String)];
		var manualFeatures = [for (feature in plan.manualFeatures) cast(feature, String)];
		var inferredFeatures = [for (feature in plan.inferredFeatures) cast(feature, String)];
		var files = [for (fileName in plan.files) fileName];
		var reasons = [
			for (reason in plan.reasons)
				{
					feature: reason.feature,
					sourceKind: reason.sourceKind,
					source: reason.source
				}
		];
		var surfacePlan = context == null ? GoSurfacePlanner.emptySnapshot() : context.surfacePlan;
		return {
			schemaVersion: 4,
			contract: contractLabel,
			policyPreset: buildContext.policyPreset.label(),
			semanticBoundarySource: buildContext.semanticBoundarySource.label(),
			mode: plan.fullCopy ? "full_copy" : "selective",
			selectiveEnabled: plan.selectiveEnabled,
			fullCopy: plan.fullCopy,
			inferenceDisabled: buildContext.hxrtNoFeatureInfer,
			manualFeatures: manualFeatures,
			inferredFeatures: inferredFeatures,
			selectedFeatures: selectedFeatures,
			files: files,
			reasons: reasons,
			manifestAuthority: plan.authority,
			capabilities: [for (capability in plan.capabilities) capability],
			surfacePlanAuthority: surfacePlan.authority,
			surfacePlanDecisionCount: surfacePlan.decisionCount,
			requiredSurfaceImports: [for (requirement in surfacePlan.requiredImports) requirement],
			requiredSurfaceRuntimeFeatures: [for (feature in surfacePlan.requiredRuntimeFeatures) feature],
			surfacePlans: [for (decision in surfacePlan.decisions) decision]
		};
	}

	function buildOptimizerPlanReportSnapshot(buildContext:GoBuildContext, context:Null<CompilationContext>):OptimizerPlanReportSnapshot {
		var contractLabel = buildContext.profile == GoProfile.Metal ? "metal" : "portable";
		var goAstPasses:Array<String> = [];
		var goAstPassSelectionSource = "planner";
		var goAstPassSelectionReasons:Array<OptimizerPassSelectionReason> = [];
		var loweringFallbackBoundaryCount = 0;
		var loweringFallbackNonBoundaryCount = 0;
		var autoLoweringCapabilities:Array<OptimizerCapabilitySummary> = [];
		var surfacePlan = context == null ? GoSurfacePlanner.emptySnapshot() : context.surfacePlan;
		if (context != null) {
			goAstPasses = context.appliedGoAstPassNames.copy();
			goAstPassSelectionSource = context.selectedGoAstPassSource;
			for (entry in context.selectedGoAstPassReasons) {
				if (entry == null) {
					continue;
				}
				goAstPassSelectionReasons.push({
					pass: entry.pass,
					reason: entry.reason,
					source: entry.source
				});
			}
			for (entry in context.loweringDecisionLedger) {
				if (entry == null || entry.outcome != "fallback") {
					continue;
				}
				if (entry.inNativeBoundary) {
					loweringFallbackBoundaryCount++;
				} else {
					loweringFallbackNonBoundaryCount++;
				}
			}
			autoLoweringCapabilities = buildOptimizerCapabilitySummaries(context);
		}
		return {
			schemaVersion: 7,
			contract: contractLabel,
			policyPreset: buildContext.policyPreset.label(),
			nativeSpecializationPolicy: buildContext.nativeSpecializationPolicy.label(),
			nativeSpecializationPolicySource: buildContext.nativeSpecializationPolicySource.label(),
			autoLoweringMode: GoAutoLoweringModeTools.label(buildContext.autoLoweringMode),
			optimizationPreset: buildContext.optimizationPreset,
			portableStringFastpathEnabled: buildContext.portableStringFastpathEnabled,
			portableConcurrencyFastpathEnabled: buildContext.portableConcurrencyFastpathEnabled,
			goAstPassSelectionSource: goAstPassSelectionSource,
			goAstPasses: goAstPasses,
			goAstPassSelectionReasons: goAstPassSelectionReasons,
			stringInstanceTypedLowerings: context == null ? 0 : context.optimizerStringInstanceTypedLowerings,
			stringInstanceLegacyLowerings: context == null ? 0 : context.optimizerStringInstanceLegacyLowerings,
			stringLengthFieldTypedLowerings: context == null ? 0 : context.optimizerStringLengthFieldTypedLowerings,
			stringLengthFieldLegacyLowerings: context == null ? 0 : context.optimizerStringLengthFieldLegacyLowerings,
			portableConcurrencyTypedFastpathHits: context == null ? 0 : context.optimizerPortableConcurrencyTypedFastpathHits,
			portableConcurrencyTypedFastpathFallbacks: context == null ? 0 : context.optimizerPortableConcurrencyTypedFastpathFallbacks,
			goCollectionsTypedLowerings: context == null ? 0 : context.optimizerGoCollectionsTypedLowerings,
			goCollectionsTypedFallbacks: context == null ? 0 : context.optimizerGoCollectionsTypedFallbacks,
			goResultTypedLowerings: context == null ? 0 : context.optimizerGoResultTypedLowerings,
			goResultTypedFallbacks: context == null ? 0 : context.optimizerGoResultTypedFallbacks,
			loweringFallbackBoundaryCount: loweringFallbackBoundaryCount,
			loweringFallbackNonBoundaryCount: loweringFallbackNonBoundaryCount,
			loweringFallbackLaneCount: loweringFallbackBoundaryCount,
			loweringFallbackNonLaneCount: loweringFallbackNonBoundaryCount,
			autoLoweringCapabilities: autoLoweringCapabilities,
			surfacePlanAuthority: surfacePlan.authority,
			surfacePlanDecisionCount: surfacePlan.decisionCount,
			requiredSurfaceImports: [for (requirement in surfacePlan.requiredImports) requirement],
			requiredSurfaceRuntimeFeatures: [for (feature in surfacePlan.requiredRuntimeFeatures) feature],
			surfacePlans: [for (decision in surfacePlan.decisions) decision]
		};
	}

	function buildOptimizerCapabilitySummaries(context:CompilationContext):Array<OptimizerCapabilitySummary> {
		var byCapability:Map<String, OptimizerCapabilityAccumulator> = [];
		for (entry in context.loweringDecisionLedger) {
			if (entry == null) {
				continue;
			}
			var capability = entry.feature == null ? "" : StringTools.trim(entry.feature);
			if (capability == "") {
				continue;
			}
			var accumulator = byCapability.get(capability);
			if (accumulator == null) {
				accumulator = {
					attempts: 0,
					successes: 0,
					fallbacks: 0,
					fallbackReasonCounts: []
				};
				byCapability.set(capability, accumulator);
			}
			switch (entry.outcome) {
				case "attempted":
					accumulator.attempts++;
				case "succeeded":
					accumulator.successes++;
				case "fallback":
					accumulator.fallbacks++;
					var kind = entry.kind == null ? "" : StringTools.trim(entry.kind);
					if (kind != "") {
						incrementIntMap(accumulator.fallbackReasonCounts, kind);
					}
				case _:
			}
		}

		var stringSuccesses = context.optimizerStringInstanceTypedLowerings + context.optimizerStringLengthFieldTypedLowerings;
		var stringFallbacks = context.optimizerStringInstanceLegacyLowerings + context.optimizerStringLengthFieldLegacyLowerings;
		var stringAttempts = stringSuccesses + stringFallbacks;
		if (stringAttempts > 0) {
			var stringCapability = "go.string.typed";
			var accumulator = byCapability.get(stringCapability);
			if (accumulator == null) {
				accumulator = {
					attempts: 0,
					successes: 0,
					fallbacks: 0,
					fallbackReasonCounts: []
				};
				byCapability.set(stringCapability, accumulator);
			}
			accumulator.attempts += stringAttempts;
			accumulator.successes += stringSuccesses;
			accumulator.fallbacks += stringFallbacks;
			if (stringFallbacks > 0) {
				incrementIntMapBy(accumulator.fallbackReasonCounts, "optimizer_preset_disabled", stringFallbacks);
			}
		}

		var capabilityIds = sortedUniqueStrings([for (id in byCapability.keys()) id]);
		var summaries:Array<OptimizerCapabilitySummary> = [];
		for (capability in capabilityIds) {
			var accumulator = byCapability.get(capability);
			if (accumulator == null) {
				continue;
			}
			var kinds = sortedUniqueStrings([for (kind in accumulator.fallbackReasonCounts.keys()) kind]);
			var reasonCounts:Array<OptimizerCapabilityFallbackReasonCount> = [];
			for (kind in kinds) {
				reasonCounts.push({
					kind: kind,
					count: accumulator.fallbackReasonCounts.get(kind)
				});
			}
			summaries.push({
				id: capability,
				attempts: accumulator.attempts,
				successes: accumulator.successes,
				fallbacks: accumulator.fallbacks,
				fallbackReasonCounts: reasonCounts
			});
		}
		return summaries;
	}

	static function compareContractFallbackEvents(a:ContractFallbackEvent, b:ContractFallbackEvent):Int {
		var moduleOrder = Reflect.compare(a.module, b.module);
		if (moduleOrder != 0) {
			return moduleOrder;
		}
		var laneOrder = Reflect.compare(a.inNativeBoundary ? 1 : 0, b.inNativeBoundary ? 1 : 0);
		if (laneOrder != 0) {
			return laneOrder;
		}
		var kindOrder = Reflect.compare(a.kind, b.kind);
		if (kindOrder != 0) {
			return kindOrder;
		}
		var locationOrder = Reflect.compare(a.location, b.location);
		if (locationOrder != 0) {
			return locationOrder;
		}
		return Reflect.compare(a.detail, b.detail);
	}

	static function compareContractLoweringDecision(a:ContractLoweringDecision, b:ContractLoweringDecision):Int {
		var moduleOrder = Reflect.compare(a.module, b.module);
		if (moduleOrder != 0) {
			return moduleOrder;
		}
		var laneOrder = Reflect.compare(a.inNativeBoundary ? 1 : 0, b.inNativeBoundary ? 1 : 0);
		if (laneOrder != 0) {
			return laneOrder;
		}
		var featureOrder = Reflect.compare(a.feature, b.feature);
		if (featureOrder != 0) {
			return featureOrder;
		}
		var kindOrder = Reflect.compare(a.kind, b.kind);
		if (kindOrder != 0) {
			return kindOrder;
		}
		var outcomeOrder = Reflect.compare(a.outcome, b.outcome);
		if (outcomeOrder != 0) {
			return outcomeOrder;
		}
		var locationOrder = Reflect.compare(a.location, b.location);
		if (locationOrder != 0) {
			return locationOrder;
		}
		return Reflect.compare(a.detail, b.detail);
	}

	static function compareContractDiagnostics(a:ContractDiagnosticEntry, b:ContractDiagnosticEntry):Int {
		var moduleOrder = Reflect.compare(a.module, b.module);
		if (moduleOrder != 0) {
			return moduleOrder;
		}
		var severityOrder = Reflect.compare(a.severity, b.severity);
		if (severityOrder != 0) {
			return severityOrder;
		}
		var codeOrder = Reflect.compare(a.code, b.code);
		if (codeOrder != 0) {
			return codeOrder;
		}
		var locationOrder = Reflect.compare(a.location, b.location);
		if (locationOrder != 0) {
			return locationOrder;
		}
		return Reflect.compare(a.message, b.message);
	}

	static function buildContractFallbackModuleSummary(boundaryCountsByModule:Map<String, Int>,
			nonBoundaryCountsByModule:Map<String, Int>):Array<ContractFallbackModuleSummary> {
		var summary = new Array<ContractFallbackModuleSummary>();
		for (moduleName in boundaryCountsByModule.keys()) {
			summary.push({
				module: moduleName,
				inNativeBoundary: true,
				count: boundaryCountsByModule.get(moduleName)
			});
		}
		for (moduleName in nonBoundaryCountsByModule.keys()) {
			summary.push({
				module: moduleName,
				inNativeBoundary: false,
				count: nonBoundaryCountsByModule.get(moduleName)
			});
		}
		summary.sort(compareContractFallbackModuleSummary);
		return summary;
	}

	static function compareContractFallbackModuleSummary(a:ContractFallbackModuleSummary, b:ContractFallbackModuleSummary):Int {
		var moduleOrder = Reflect.compare(a.module, b.module);
		if (moduleOrder != 0) {
			return moduleOrder;
		}
		var laneOrder = Reflect.compare(a.inNativeBoundary ? 1 : 0, b.inNativeBoundary ? 1 : 0);
		if (laneOrder != 0) {
			return laneOrder;
		}
		return Reflect.compare(a.count, b.count);
	}

	static function renderContractReportJson(snapshot:ContractReportSnapshot):String {
		var lines:Array<String> = [];
		lines.push("{");
		lines.push('\t"schemaVersion": ' + snapshot.schemaVersion + ",");
		lines.push('\t"contract": "' + jsonEscape(snapshot.contract) + '",');
		lines.push('\t"policyPreset": "' + jsonEscape(snapshot.policyPreset) + '",');
		lines.push('\t"semanticBoundarySource": "' + jsonEscape(snapshot.semanticBoundarySource) + '",');
		lines.push('\t"nativeAuthorityPolicy": "' + jsonEscape(snapshot.nativeAuthorityPolicy) + '",');
		lines.push('\t"nativeAuthorityPolicySource": "' + jsonEscape(snapshot.nativeAuthorityPolicySource) + '",');
		lines.push('\t"nativeSpecializationPolicy": "' + jsonEscape(snapshot.nativeSpecializationPolicy) + '",');
		lines.push('\t"nativeSpecializationPolicySource": "' + jsonEscape(snapshot.nativeSpecializationPolicySource) + '",');
		lines.push('\t"nativeFallbackPolicy": "' + jsonEscape(snapshot.nativeFallbackPolicy) + '",');
		lines.push('\t"nativeFallbackPolicySource": "' + jsonEscape(snapshot.nativeFallbackPolicySource) + '",');
		lines.push('\t"autoLoweringMode": "' + jsonEscape(snapshot.autoLoweringMode) + '",');
		lines.push('\t"strictExamples": ' + boolString(snapshot.strictExamples) + ",");
		lines.push('\t"strictUserBoundaryPolicy": "' + jsonEscape(snapshot.strictUserBoundaryPolicy) + '",');
		lines.push('\t"strictUserBoundaries": ' + boolString(snapshot.strictUserBoundaries) + ",");
		lines.push('\t"metalFallbackAllowed": ' + boolString(snapshot.metalFallbackAllowed) + ",");
		lines.push('\t"metalContractHardError": ' + boolString(snapshot.metalContractHardError) + ",");
		lines.push('\t"emitLineDirectives": ' + boolString(snapshot.emitLineDirectives) + ",");
		lines.push('\t"rawNativeMode": "' + jsonEscape(snapshot.rawNativeMode) + '",');
		lines.push('\t"hxrtSelectiveEnabled": ' + boolString(snapshot.hxrtSelectiveEnabled) + ",");
		lines.push('\t"hxrtForceFullCopy": ' + boolString(snapshot.hxrtForceFullCopy) + ",");
		lines.push('\t"hxrtNoFeatureInfer": ' + boolString(snapshot.hxrtNoFeatureInfer) + ",");
		lines.push('\t"loweringDecisionCount": ' + snapshot.loweringDecisionCount + ",");
		lines.push('\t"loweringDecisionAttemptCount": ' + snapshot.loweringDecisionAttemptCount + ",");
		lines.push('\t"loweringDecisionSuccessCount": ' + snapshot.loweringDecisionSuccessCount + ",");
		lines.push('\t"loweringDecisionFallbackCount": ' + snapshot.loweringDecisionFallbackCount + ",");
		lines.push('\t"nativeFallbackEventCount": ' + snapshot.nativeFallbackEventCount + ",");
		lines.push('\t"nativeFallbackBoundaryEventCount": ' + snapshot.nativeFallbackBoundaryEventCount + ",");
		lines.push('\t"nativeFallbackNonBoundaryEventCount": ' + snapshot.nativeFallbackNonBoundaryEventCount + ",");
		lines.push('\t"metalFallbackViolationCount": ' + snapshot.metalFallbackViolationCount + ",");
		lines.push('\t"metalFallbackLaneViolationCount": ' + snapshot.metalFallbackLaneViolationCount + ",");
		lines.push('\t"metalFallbackNonLaneViolationCount": ' + snapshot.metalFallbackNonLaneViolationCount + ",");
		lines.push('\t"portableNativeImportScanMode": "' + jsonEscape(snapshot.portableNativeImportScanMode) + '",');
		lines.push('\t"portableNativeImportHitCount": ' + snapshot.portableNativeImportHitCount + ",");
		lines.push('\t"portableNativeImportTypedHitCount": ' + snapshot.portableNativeImportTypedHitCount + ",");
		lines.push('\t"portableNativeImportScannerHitCount": ' + snapshot.portableNativeImportScannerHitCount + ",");
		lines.push('\t"contractDiagnosticCount": ' + snapshot.contractDiagnosticCount + ",");
		lines.push('\t"hxrtManualFeatures": [');
		appendJsonStringArray(lines, snapshot.hxrtManualFeatures, 2);
		lines.push("\t],");
		lines.push('\t"nativeBoundaryModules": [');
		appendJsonStringArray(lines, snapshot.nativeBoundaryModules, 2);
		lines.push("\t],");
		lines.push('\t"metalLaneModules": [');
		appendJsonStringArray(lines, snapshot.metalLaneModules, 2);
		lines.push("\t],");
		lines.push('\t"portableNativeImportHits": [');
		appendJsonStringArray(lines, snapshot.portableNativeImportHits, 2);
		lines.push("\t],");
		lines.push('\t"portableNativeImportTypedHits": [');
		appendJsonStringArray(lines, snapshot.portableNativeImportTypedHits, 2);
		lines.push("\t],");
		lines.push('\t"portableNativeImportScannerHits": [');
		appendJsonStringArray(lines, snapshot.portableNativeImportScannerHits, 2);
		lines.push("\t],");
		lines.push('\t"contractDiagnostics": [');
		appendJsonContractDiagnosticArray(lines, snapshot.contractDiagnostics, 2);
		lines.push("\t],");
		lines.push('\t"loweringDecisions": [');
		appendJsonContractLoweringDecisionArray(lines, snapshot.loweringDecisions, 2);
		lines.push("\t],");
		lines.push('\t"nativeFallbackEventsByModule": [');
		appendJsonContractFallbackSummaryArray(lines, snapshot.nativeFallbackEventsByModule, 2);
		lines.push("\t],");
		lines.push('\t"nativeFallbackEvents": [');
		appendJsonContractFallbackArray(lines, snapshot.nativeFallbackEvents, 2);
		lines.push("\t],");
		lines.push('\t"metalFallbackViolationsByModule": [');
		appendJsonContractFallbackSummaryArray(lines, snapshot.metalFallbackViolationsByModule, 2);
		lines.push("\t],");
		lines.push('\t"metalFallbackViolations": [');
		appendJsonContractFallbackArray(lines, snapshot.metalFallbackViolations, 2);
		lines.push("\t]");
		lines.push("}");
		return lines.join("\n") + "\n";
	}

	static function renderContractReportMarkdown(snapshot:ContractReportSnapshot):String {
		var lines:Array<String> = [];
		lines.push("# Contract Report");
		lines.push("");
		lines.push("- schema version: `" + snapshot.schemaVersion + "`");
		lines.push("- contract: `" + snapshot.contract + "`");
		lines.push("- policy preset: `" + snapshot.policyPreset + "`");
		lines.push("- semantic boundary source: `" + snapshot.semanticBoundarySource + "`");
		lines.push("- native authority policy: `"
			+ snapshot.nativeAuthorityPolicy
			+ "` (source `"
			+ snapshot.nativeAuthorityPolicySource
			+ "`)");
		lines.push("- native specialization policy: `" + snapshot.nativeSpecializationPolicy + "` (source `" + snapshot.nativeSpecializationPolicySource +
			"`)");
		lines.push("- native fallback policy: `"
			+ snapshot.nativeFallbackPolicy
			+ "` (source `"
			+ snapshot.nativeFallbackPolicySource
			+ "`)");
		lines.push("- auto lowering mode: `" + snapshot.autoLoweringMode + "`");
		lines.push("- strict examples: `" + boolLabel(snapshot.strictExamples) + "`");
		lines.push("- strict user boundary policy: `" + snapshot.strictUserBoundaryPolicy + "`");
		lines.push("- strict user boundaries: `" + boolLabel(snapshot.strictUserBoundaries) + "`");
		lines.push("- metal fallback allowed: `" + boolLabel(snapshot.metalFallbackAllowed) + "`");
		lines.push("- metal contract hard error: `" + boolLabel(snapshot.metalContractHardError) + "`");
		lines.push("- emit line directives: `" + boolLabel(snapshot.emitLineDirectives) + "`");
		lines.push("- raw native mode: `" + snapshot.rawNativeMode + "`");
		lines.push("- hxrt selective enabled: `" + boolLabel(snapshot.hxrtSelectiveEnabled) + "`");
		lines.push("- hxrt force full copy: `" + boolLabel(snapshot.hxrtForceFullCopy) + "`");
		lines.push("- hxrt no feature infer: `" + boolLabel(snapshot.hxrtNoFeatureInfer) + "`");
		lines.push("- native fallback events: `" + snapshot.nativeFallbackEventCount + "`");
		lines.push("- native fallback boundary events: `" + snapshot.nativeFallbackBoundaryEventCount + "`");
		lines.push("- native fallback non-boundary events: `" + snapshot.nativeFallbackNonBoundaryEventCount + "`");
		lines.push("- metal fallback violations: `" + snapshot.metalFallbackViolationCount + "`");
		lines.push("- metal fallback lane violations: `" + snapshot.metalFallbackLaneViolationCount + "`");
		lines.push("- metal fallback non-lane violations: `" + snapshot.metalFallbackNonLaneViolationCount + "`");
		lines.push("- portable native import scan mode: `" + snapshot.portableNativeImportScanMode + "`");
		lines.push("- portable native import hits: `" + snapshot.portableNativeImportHitCount + "`");
		lines.push("- portable native import typed hits: `" + snapshot.portableNativeImportTypedHitCount + "`");
		lines.push("- portable native import scanner hits: `" + snapshot.portableNativeImportScannerHitCount + "`");
		lines.push("- contract diagnostics: `" + snapshot.contractDiagnosticCount + "`");
		lines.push("- lowering decisions: `" + snapshot.loweringDecisionCount + "` (attempts `" + snapshot.loweringDecisionAttemptCount + "`, success `"
			+ snapshot.loweringDecisionSuccessCount + "`, fallback `" + snapshot.loweringDecisionFallbackCount + "`)");
		lines.push("");
		lines.push("## hxrt manual features");
		if (snapshot.hxrtManualFeatures.length == 0) {
			lines.push("- none");
		} else {
			for (feature in snapshot.hxrtManualFeatures) {
				lines.push("- `" + feature + "`");
			}
		}
		lines.push("");
		lines.push("## native boundary modules");
		if (snapshot.nativeBoundaryModules.length == 0) {
			lines.push("- none");
		} else {
			for (moduleName in snapshot.nativeBoundaryModules) {
				lines.push("- `" + moduleName + "`");
			}
		}
		lines.push("");
		lines.push("## metal lane modules");
		if (snapshot.metalLaneModules.length == 0) {
			lines.push("- none");
		} else {
			for (moduleName in snapshot.metalLaneModules) {
				lines.push("- `" + moduleName + "`");
			}
		}
		lines.push("");
		lines.push("## portable native import hits");
		if (snapshot.portableNativeImportHits.length == 0) {
			lines.push("- none");
		} else {
			for (moduleName in snapshot.portableNativeImportHits) {
				lines.push("- `" + moduleName + "`");
			}
		}
		lines.push("");
		lines.push("## portable native import typed hits");
		if (snapshot.portableNativeImportTypedHits.length == 0) {
			lines.push("- none");
		} else {
			for (moduleName in snapshot.portableNativeImportTypedHits) {
				lines.push("- `" + moduleName + "`");
			}
		}
		lines.push("");
		lines.push("## portable native import scanner hits");
		if (snapshot.portableNativeImportScannerHits.length == 0) {
			lines.push("- none");
		} else {
			for (moduleName in snapshot.portableNativeImportScannerHits) {
				lines.push("- `" + moduleName + "`");
			}
		}
		lines.push("");
		lines.push("## contract diagnostics");
		if (snapshot.contractDiagnostics.length == 0) {
			lines.push("- none");
		} else {
			for (entry in snapshot.contractDiagnostics) {
				lines.push("- `"
					+ entry.module
					+ "` | `"
					+ entry.code
					+ "` | `"
					+ entry.severity
					+ "` | `"
					+ entry.location
					+ "` | "
					+ entry.message);
			}
		}
		lines.push("");
		lines.push("## lowering decisions");
		if (snapshot.loweringDecisions.length == 0) {
			lines.push("- none");
		} else {
			for (entry in snapshot.loweringDecisions) {
				var laneLabel = entry.inNativeBoundary ? "native-boundary" : "non-boundary";
				lines.push("- `" + entry.module + "` (" + laneLabel + ") | `" + entry.feature + "` | `" + entry.kind + "` | `" + entry.outcome + "` | `"
					+ entry.location + "` | " + entry.detail);
			}
		}
		lines.push("");
		lines.push("## native fallback event summary by module");
		if (snapshot.nativeFallbackEventsByModule.length == 0) {
			lines.push("- none");
		} else {
			for (entry in snapshot.nativeFallbackEventsByModule) {
				var boundaryLabel = entry.inNativeBoundary ? "native-boundary" : "non-boundary";
				lines.push("- `" + entry.module + "` (" + boundaryLabel + "): `" + entry.count + "`");
			}
		}
		lines.push("");
		lines.push("## native fallback events");
		if (snapshot.nativeFallbackEvents.length == 0) {
			lines.push("- none");
		} else {
			for (entry in snapshot.nativeFallbackEvents) {
				var boundaryLabel = entry.inNativeBoundary ? "native-boundary" : "non-boundary";
				lines.push("- `" + entry.module + "` (" + boundaryLabel + ") | `" + entry.kind + "` | `" + entry.location + "` | " + entry.detail);
			}
		}
		lines.push("");
		lines.push("## metal fallback violation summary by module");
		if (snapshot.metalFallbackViolationsByModule.length == 0) {
			lines.push("- none");
		} else {
			for (entry in snapshot.metalFallbackViolationsByModule) {
				var laneLabel = entry.inNativeBoundary ? "lane" : "non-lane";
				lines.push("- `" + entry.module + "` (" + laneLabel + "): `" + entry.count + "`");
			}
		}
		lines.push("");
		lines.push("## metal fallback violations");
		if (snapshot.metalFallbackViolations.length == 0) {
			lines.push("- none");
		} else {
			for (entry in snapshot.metalFallbackViolations) {
				var laneLabel = entry.inNativeBoundary ? "lane" : "non-lane";
				lines.push("- `" + entry.module + "` (" + laneLabel + ") | `" + entry.kind + "` | `" + entry.location + "` | " + entry.detail);
			}
		}
		lines.push("");
		return lines.join("\n");
	}

	static function renderRuntimePlanJson(snapshot:RuntimePlanReportSnapshot):String {
		var lines:Array<String> = [];
		lines.push("{");
		lines.push('\t"schemaVersion": ' + snapshot.schemaVersion + ",");
		lines.push('\t"contract": "' + jsonEscape(snapshot.contract) + '",');
		lines.push('\t"policyPreset": "' + jsonEscape(snapshot.policyPreset) + '",');
		lines.push('\t"semanticBoundarySource": "' + jsonEscape(snapshot.semanticBoundarySource) + '",');
		lines.push('\t"mode": "' + jsonEscape(snapshot.mode) + '",');
		lines.push('\t"selectiveEnabled": ' + boolString(snapshot.selectiveEnabled) + ",");
		lines.push('\t"fullCopy": ' + boolString(snapshot.fullCopy) + ",");
		lines.push('\t"inferenceDisabled": ' + boolString(snapshot.inferenceDisabled) + ",");
		lines.push('\t"manifestAuthority": "' + jsonEscape(snapshot.manifestAuthority) + '",');
		lines.push('\t"manualFeatures": [');
		appendJsonStringArray(lines, snapshot.manualFeatures, 2);
		lines.push("\t],");
		lines.push('\t"inferredFeatures": [');
		appendJsonStringArray(lines, snapshot.inferredFeatures, 2);
		lines.push("\t],");
		lines.push('\t"selectedFeatures": [');
		appendJsonStringArray(lines, snapshot.selectedFeatures, 2);
		lines.push("\t],");
		lines.push('\t"files": [');
		appendJsonStringArray(lines, snapshot.files, 2);
		lines.push("\t],");
		lines.push('\t"capabilities": [');
		appendJsonRuntimeCapabilities(lines, snapshot.capabilities, 2);
		lines.push("\t],");
		lines.push('\t"surfacePlanAuthority": "' + jsonEscape(snapshot.surfacePlanAuthority) + '",');
		lines.push('\t"surfacePlanDecisionCount": ' + snapshot.surfacePlanDecisionCount + ",");
		lines.push('\t"requiredSurfaceImports": [');
		appendJsonSurfaceImportArray(lines, snapshot.requiredSurfaceImports, 2);
		lines.push("\t],");
		lines.push('\t"requiredSurfaceRuntimeFeatures": [');
		appendJsonStringArray(lines, snapshot.requiredSurfaceRuntimeFeatures, 2);
		lines.push("\t],");
		lines.push('\t"surfacePlans": [');
		appendJsonSurfacePlanArray(lines, snapshot.surfacePlans, 2);
		lines.push("\t],");
		lines.push('\t"reasons": [');
		appendJsonRuntimeReasons(lines, snapshot.reasons, 2);
		lines.push("\t]");
		lines.push("}");
		return lines.join("\n") + "\n";
	}

	static function renderRuntimePlanMarkdown(snapshot:RuntimePlanReportSnapshot):String {
		var lines:Array<String> = [];
		lines.push("# Runtime Plan Report");
		lines.push("");
		lines.push("- schema version: `" + snapshot.schemaVersion + "`");
		lines.push("- contract: `" + snapshot.contract + "`");
		lines.push("- policy preset: `" + snapshot.policyPreset + "`");
		lines.push("- semantic boundary source: `" + snapshot.semanticBoundarySource + "`");
		lines.push("- mode: `" + snapshot.mode + "`");
		lines.push("- selective enabled: `" + boolLabel(snapshot.selectiveEnabled) + "`");
		lines.push("- full copy: `" + boolLabel(snapshot.fullCopy) + "`");
		lines.push("- inference disabled: `" + boolLabel(snapshot.inferenceDisabled) + "`");
		lines.push("- manifest authority: `" + snapshot.manifestAuthority + "`");
		lines.push("- surface plan authority: `" + snapshot.surfacePlanAuthority + "`");
		lines.push("- surface plan decisions: `" + snapshot.surfacePlanDecisionCount + "`");
		lines.push("");
		lines.push("## manual features");
		if (snapshot.manualFeatures.length == 0) {
			lines.push("- none");
		} else {
			for (feature in snapshot.manualFeatures) {
				lines.push("- `" + feature + "`");
			}
		}
		lines.push("");
		lines.push("## inferred features");
		if (snapshot.inferredFeatures.length == 0) {
			lines.push("- none");
		} else {
			for (feature in snapshot.inferredFeatures) {
				lines.push("- `" + feature + "`");
			}
		}
		lines.push("");
		lines.push("## selected features");
		if (snapshot.selectedFeatures.length == 0) {
			lines.push("- none");
		} else {
			for (feature in snapshot.selectedFeatures) {
				lines.push("- `" + feature + "`");
			}
		}
		lines.push("");
		lines.push("## runtime files");
		if (snapshot.files.length == 0) {
			lines.push("- full copy (`runtime/hxrt/**`, excluding footprint-explicit diagnostic/capability files unless their typed use or define is enabled)");
		} else {
			for (fileName in snapshot.files) {
				lines.push("- `" + fileName + "`");
			}
		}
		lines.push("");
		lines.push("## capability manifest");
		for (capability in snapshot.capabilities) {
			var capabilityFiles = [for (fileName in capability.files) fileName];
			lines.push("- `" + capability.id + "` -> `" + capabilityFiles.join(", ") + "`");
			for (reason in capability.reasons) {
				lines.push("  - `" + reason.sourceKind + "` (`" + reason.source + "`)");
			}
		}
		lines.push("");
		lines.push("## selection reasons");
		if (snapshot.reasons.length == 0) {
			lines.push("- none");
		} else {
			for (reason in snapshot.reasons) {
				lines.push("- `" + reason.feature + "` <- `" + reason.sourceKind + "` (`" + reason.source + "`)");
			}
		}
		lines.push("");
		appendSurfacePlanMarkdown(lines, snapshot.requiredSurfaceImports, snapshot.requiredSurfaceRuntimeFeatures, snapshot.surfacePlans);
		lines.push("");
		return lines.join("\n");
	}

	static function renderOptimizerPlanJson(snapshot:OptimizerPlanReportSnapshot):String {
		var lines:Array<String> = [];
		lines.push("{");
		lines.push('\t"schemaVersion": ' + snapshot.schemaVersion + ",");
		lines.push('\t"contract": "' + jsonEscape(snapshot.contract) + '",');
		lines.push('\t"policyPreset": "' + jsonEscape(snapshot.policyPreset) + '",');
		lines.push('\t"nativeSpecializationPolicy": "' + jsonEscape(snapshot.nativeSpecializationPolicy) + '",');
		lines.push('\t"nativeSpecializationPolicySource": "' + jsonEscape(snapshot.nativeSpecializationPolicySource) + '",');
		lines.push('\t"autoLoweringMode": "' + jsonEscape(snapshot.autoLoweringMode) + '",');
		lines.push('\t"optimizationPreset": "' + jsonEscape(snapshot.optimizationPreset) + '",');
		lines.push('\t"portableStringFastpathEnabled": ' + boolString(snapshot.portableStringFastpathEnabled) + ",");
		lines.push('\t"portableConcurrencyFastpathEnabled": ' + boolString(snapshot.portableConcurrencyFastpathEnabled) + ",");
		lines.push('\t"stringInstanceTypedLowerings": ' + snapshot.stringInstanceTypedLowerings + ",");
		lines.push('\t"stringInstanceLegacyLowerings": ' + snapshot.stringInstanceLegacyLowerings + ",");
		lines.push('\t"stringLengthFieldTypedLowerings": ' + snapshot.stringLengthFieldTypedLowerings + ",");
		lines.push('\t"stringLengthFieldLegacyLowerings": ' + snapshot.stringLengthFieldLegacyLowerings + ",");
		lines.push('\t"portableConcurrencyTypedFastpathHits": ' + snapshot.portableConcurrencyTypedFastpathHits + ",");
		lines.push('\t"portableConcurrencyTypedFastpathFallbacks": ' + snapshot.portableConcurrencyTypedFastpathFallbacks + ",");
		lines.push('\t"goCollectionsTypedLowerings": ' + snapshot.goCollectionsTypedLowerings + ",");
		lines.push('\t"goCollectionsTypedFallbacks": ' + snapshot.goCollectionsTypedFallbacks + ",");
		lines.push('\t"goResultTypedLowerings": ' + snapshot.goResultTypedLowerings + ",");
		lines.push('\t"goResultTypedFallbacks": ' + snapshot.goResultTypedFallbacks + ",");
		lines.push('\t"loweringFallbackBoundaryCount": ' + snapshot.loweringFallbackBoundaryCount + ",");
		lines.push('\t"loweringFallbackNonBoundaryCount": ' + snapshot.loweringFallbackNonBoundaryCount + ",");
		lines.push('\t"loweringFallbackLaneCount": ' + snapshot.loweringFallbackLaneCount + ",");
		lines.push('\t"loweringFallbackNonLaneCount": ' + snapshot.loweringFallbackNonLaneCount + ",");
		lines.push('\t"autoLoweringCapabilities": [');
		appendJsonOptimizerCapabilitySummaries(lines, snapshot.autoLoweringCapabilities, 2);
		lines.push("\t],");
		lines.push('\t"surfacePlanAuthority": "' + jsonEscape(snapshot.surfacePlanAuthority) + '",');
		lines.push('\t"surfacePlanDecisionCount": ' + snapshot.surfacePlanDecisionCount + ",");
		lines.push('\t"requiredSurfaceImports": [');
		appendJsonSurfaceImportArray(lines, snapshot.requiredSurfaceImports, 2);
		lines.push("\t],");
		lines.push('\t"requiredSurfaceRuntimeFeatures": [');
		appendJsonStringArray(lines, snapshot.requiredSurfaceRuntimeFeatures, 2);
		lines.push("\t],");
		lines.push('\t"surfacePlans": [');
		appendJsonSurfacePlanArray(lines, snapshot.surfacePlans, 2);
		lines.push("\t],");
		lines.push('\t"goAstPassSelectionSource": "' + jsonEscape(snapshot.goAstPassSelectionSource) + '",');
		lines.push('\t"goAstPasses": [');
		appendJsonStringArray(lines, snapshot.goAstPasses, 2);
		lines.push("\t],");
		lines.push('\t"goAstPassSelectionReasons": [');
		appendJsonOptimizerPassSelectionReasons(lines, snapshot.goAstPassSelectionReasons, 2);
		lines.push("\t]");
		lines.push("}");
		return lines.join("\n") + "\n";
	}

	static function renderOptimizerPlanMarkdown(snapshot:OptimizerPlanReportSnapshot):String {
		var lines:Array<String> = [];
		lines.push("# Optimizer Plan Report");
		lines.push("");
		lines.push("- schema version: `" + snapshot.schemaVersion + "`");
		lines.push("- contract: `" + snapshot.contract + "`");
		lines.push("- policy preset: `" + snapshot.policyPreset + "`");
		lines.push("- native specialization policy: `" + snapshot.nativeSpecializationPolicy + "` (source `" + snapshot.nativeSpecializationPolicySource +
			"`)");
		lines.push("- auto lowering mode: `" + snapshot.autoLoweringMode + "`");
		lines.push("- optimization preset: `" + snapshot.optimizationPreset + "`");
		lines.push("- portable string fastpath enabled: `" + boolLabel(snapshot.portableStringFastpathEnabled) + "`");
		lines.push("- portable concurrency fastpath enabled: `" + boolLabel(snapshot.portableConcurrencyFastpathEnabled) + "`");
		lines.push("- string instance typed lowerings: `" + snapshot.stringInstanceTypedLowerings + "`");
		lines.push("- string instance legacy lowerings: `" + snapshot.stringInstanceLegacyLowerings + "`");
		lines.push("- string length field typed lowerings: `" + snapshot.stringLengthFieldTypedLowerings + "`");
		lines.push("- string length field legacy lowerings: `" + snapshot.stringLengthFieldLegacyLowerings + "`");
		lines.push("- portable concurrency typed fastpath hits: `" + snapshot.portableConcurrencyTypedFastpathHits + "`");
		lines.push("- portable concurrency typed fastpath fallbacks: `" + snapshot.portableConcurrencyTypedFastpathFallbacks + "`");
		lines.push("- go collections typed lowerings: `" + snapshot.goCollectionsTypedLowerings + "`");
		lines.push("- go collections typed fallbacks: `" + snapshot.goCollectionsTypedFallbacks + "`");
		lines.push("- go result typed lowerings: `" + snapshot.goResultTypedLowerings + "`");
		lines.push("- go result typed fallbacks: `" + snapshot.goResultTypedFallbacks + "`");
		lines.push("- lowering fallback boundary count: `" + snapshot.loweringFallbackBoundaryCount + "`");
		lines.push("- lowering fallback non-boundary count: `" + snapshot.loweringFallbackNonBoundaryCount + "`");
		lines.push("- lowering fallback lane count: `" + snapshot.loweringFallbackLaneCount + "`");
		lines.push("- lowering fallback non-lane count: `" + snapshot.loweringFallbackNonLaneCount + "`");
		lines.push("- surface plan authority: `" + snapshot.surfacePlanAuthority + "`");
		lines.push("- surface plan decisions: `" + snapshot.surfacePlanDecisionCount + "`");
		lines.push("- go ast pass selection source: `" + snapshot.goAstPassSelectionSource + "`");
		lines.push("");
		lines.push("## auto lowering capabilities");
		if (snapshot.autoLoweringCapabilities.length == 0) {
			lines.push("- none");
		} else {
			for (capability in snapshot.autoLoweringCapabilities) {
				lines.push("- `" + capability.id + "` | attempts `" + capability.attempts + "` | success `" + capability.successes + "` | fallback `"
					+ capability.fallbacks + "`");
				if (capability.fallbackReasonCounts.length == 0) {
					lines.push("  fallback reasons: none");
				} else {
					var parts = new Array<String>();
					for (reason in capability.fallbackReasonCounts) {
						parts.push(reason.kind + "=" + reason.count);
					}
					lines.push("  fallback reasons: " + parts.join(", "));
				}
			}
		}
		lines.push("");
		lines.push("## go ast passes");
		if (snapshot.goAstPasses.length == 0) {
			lines.push("- none");
		} else {
			for (passName in snapshot.goAstPasses) {
				lines.push("- `" + passName + "`");
			}
		}
		lines.push("");
		lines.push("## go ast pass selection reasons");
		if (snapshot.goAstPassSelectionReasons.length == 0) {
			lines.push("- none");
		} else {
			for (entry in snapshot.goAstPassSelectionReasons) {
				lines.push("- `" + entry.pass + "` | `" + entry.source + "` | " + entry.reason);
			}
		}
		lines.push("");
		appendSurfacePlanMarkdown(lines, snapshot.requiredSurfaceImports, snapshot.requiredSurfaceRuntimeFeatures, snapshot.surfacePlans);
		lines.push("");
		return lines.join("\n");
	}

	static function appendJsonStringArray(lines:Array<String>, values:Array<String>, indentLevel:Int):Void {
		var indent = [for (_ in 0...indentLevel) "\t"].join("");
		for (index in 0...values.length) {
			var suffix = index == values.length - 1 ? "" : ",";
			lines.push(indent + '"' + jsonEscape(values[index]) + '"' + suffix);
		}
	}

	static function appendJsonSurfaceImportArray(lines:Array<String>, imports:Array<GoSurfaceImportRequirement>, indentLevel:Int):Void {
		var indent = [for (_ in 0...indentLevel) "\t"].join("");
		for (index in 0...imports.length) {
			var requirement = imports[index];
			var suffix = index == imports.length - 1 ? "" : ",";
			lines.push(indent + '{"path":"' + jsonEscape(requirement.path) + '","reason":"' + jsonEscape(requirement.reason) + '"}' + suffix);
		}
	}

	static function appendJsonImmutableSurfaceImportArray(lines:Array<String>, imports:GoImmutableList<GoSurfaceImportRequirement>, indentLevel:Int):Void {
		appendJsonSurfaceImportArray(lines, [for (requirement in imports) requirement], indentLevel);
	}

	static function appendJsonSurfacePlanArray(lines:Array<String>, decisions:Array<GoSurfacePlanDecision>, indentLevel:Int):Void {
		var indent = [for (_ in 0...indentLevel) "\t"].join("");
		for (index in 0...decisions.length) {
			var decision = decisions[index];
			var suffix = index == decisions.length - 1 ? "" : ",";
			lines.push(indent + "{");
			lines.push(indent + '\t"module": "' + jsonEscape(decision.module) + '",');
			lines.push(indent + '\t"location": "' + jsonEscape(decision.location) + '",');
			lines.push(indent + '\t"usageLevel": "' + jsonEscape(decision.usageLevel) + '",');
			lines.push(indent + '\t"usedType": ' + GoTypeUsageLedger.renderShapeJson(decision.usedType) + ",");
			lines.push(indent + '\t"contract": {');
			lines.push(indent + '\t\t"surfaceId": "' + jsonEscape(decision.surfaceId) + '",');
			lines.push(indent + '\t\t"version": ' + decision.contractVersion);
			lines.push(indent + "\t},");
			lines.push(indent + '\t"eligibility": {');
			lines.push(indent + '\t\t"outcome": "' + jsonEscape(decision.eligibilityOutcome) + '",');
			lines.push(indent + '\t\t"reason": "' + jsonEscape(decision.eligibilityReason) + '",');
			lines.push(indent + '\t\t"detail": "' + jsonEscape(decision.eligibilityDetail) + '"');
			lines.push(indent + "\t},");
			lines.push(indent + '\t"selection": "' + jsonEscape(decision.selection) + '",');
			lines.push(indent + '\t"selectionReason": "' + jsonEscape(decision.selectionReason) + '",');
			lines.push(indent + '\t"selectedRepresentation": ' + nullableJsonString(decision.selectedRepresentation) + ",");
			lines.push(indent + '\t"fallbackReason": ' + nullableJsonString(decision.fallbackReason) + ",");
			lines.push(indent + '\t"imports": [');
			appendJsonImmutableSurfaceImportArray(lines, decision.imports, indentLevel + 2);
			lines.push(indent + "\t],");
			lines.push(indent + '\t"runtimeRequirements": [');
			appendJsonStringArray(lines, [for (feature in decision.runtimeRequirements) feature], indentLevel + 2);
			lines.push(indent + "\t],");
			lines.push(indent + '\t"noHxrtStatus": ' + nullableJsonString(decision.noHxrtStatus) + ",");
			lines.push(indent + '\t"selectedNoHxrtEligible": ' + boolString(decision.selectedNoHxrtEligible));
			lines.push(indent + "}" + suffix);
		}
	}

	static function appendSurfacePlanMarkdown(lines:Array<String>, requiredImports:Array<GoSurfaceImportRequirement>, requiredRuntimeFeatures:Array<String>,
			decisions:Array<GoSurfacePlanDecision>):Void {
		lines.push("## portable surface plan consequences");
		lines.push("");
		lines.push("- required imports: `"
			+ (requiredImports.length == 0 ? "none" : [for (requirement in requiredImports) requirement.path].join(", "))
			+ "`");
		lines.push("- required runtime features: `" + (requiredRuntimeFeatures.length == 0 ? "none" : requiredRuntimeFeatures.join(", ")) + "`");
		lines.push("");
		lines.push("## portable surface decisions");
		lines.push("");
		if (decisions.length == 0) {
			lines.push("- none");
			return;
		}
		for (decision in decisions) {
			var fallback = decision.fallbackReason == null ? "none" : decision.fallbackReason;
			var representation = decision.selectedRepresentation == null ? "none" : decision.selectedRepresentation;
			var imports = [
				for (requirement in decision.imports)
					requirement.path + " (" + requirement.reason + ")"
			];
			var runtimeRequirements = [for (feature in decision.runtimeRequirements) feature];
			lines.push("- `"
				+ decision.module
				+ "` | location `"
				+ decision.location
				+ "` | usage `"
				+ decision.usageLevel
				+ "` | `"
				+ decision.surfaceId
				+ "` v"
				+ decision.contractVersion
				+ " | used type `"
				+ GoTypeUsageLedger.renderShapeJson(decision.usedType)
				+ "` | eligibility `"
				+ decision.eligibilityOutcome
				+ ":"
				+ decision.eligibilityReason
				+ "` | eligibility detail `"
				+ decision.eligibilityDetail
				+ "` | selection `"
				+ decision.selection
				+ ":"
				+ decision.selectionReason
				+ "` | representation `"
				+ representation
				+ "` | fallback `"
				+ fallback
				+ "` | imports `"
				+ (imports.length == 0 ? "none" : imports.join(", "))
				+ "` | runtime `"
				+ (runtimeRequirements.length == 0 ? "none" : runtimeRequirements.join(", "))
				+ "` | no-hxrt contract `"
				+ (decision.noHxrtStatus == null ? "none" : decision.noHxrtStatus)
				+ "` | selected no-hxrt eligible `"
				+ boolLabel(decision.selectedNoHxrtEligible)
				+ "`");
		}
	}

	static function appendJsonRuntimeCapabilities(lines:Array<String>, capabilities:Array<GoRuntimeCapabilitySelection>, indentLevel:Int):Void {
		var indent = [for (_ in 0...indentLevel) "\t"].join("");
		for (index in 0...capabilities.length) {
			var capability = capabilities[index];
			var suffix = index == capabilities.length - 1 ? "" : ",";
			lines.push(indent + "{");
			lines.push(indent + '\t"id": "' + jsonEscape(capability.id) + '",');
			lines.push(indent + '\t"files": [');
			appendJsonStringArray(lines, [for (fileName in capability.files) fileName], indentLevel + 2);
			lines.push(indent + "\t],");
			lines.push(indent + '\t"reasons": [');
			appendJsonRuntimeReasons(lines, [
				for (reason in capability.reasons)
					{
						feature: reason.feature,
						sourceKind: reason.sourceKind,
						source: reason.source
					}
			], indentLevel + 2);
			lines.push(indent + "\t]");
			lines.push(indent + "}" + suffix);
		}
	}

	static function nullableJsonString(value:Null<String>):String {
		return value == null ? "null" : '"' + jsonEscape(value) + '"';
	}

	static function appendJsonContractFallbackArray(lines:Array<String>, events:Array<ContractFallbackEvent>, indentLevel:Int):Void {
		var indent = [for (_ in 0...indentLevel) "\t"].join("");
		for (index in 0...events.length) {
			var event = events[index];
			var suffix = index == events.length - 1 ? "" : ",";
			lines.push(indent + "{");
			lines.push(indent + '\t"module": "' + jsonEscape(event.module) + '",');
			lines.push(indent + '\t"inNativeBoundary": ' + boolString(event.inNativeBoundary) + ",");
			lines.push(indent + '\t"inMetalLane": ' + boolString(event.inNativeBoundary) + ",");
			lines.push(indent + '\t"kind": "' + jsonEscape(event.kind) + '",');
			lines.push(indent + '\t"location": "' + jsonEscape(event.location) + '",');
			lines.push(indent + '\t"detail": "' + jsonEscape(event.detail) + '"');
			lines.push(indent + "}" + suffix);
		}
	}

	static function appendJsonContractLoweringDecisionArray(lines:Array<String>, decisions:Array<ContractLoweringDecision>, indentLevel:Int):Void {
		var indent = [for (_ in 0...indentLevel) "\t"].join("");
		for (index in 0...decisions.length) {
			var entry = decisions[index];
			var suffix = index == decisions.length - 1 ? "" : ",";
			lines.push(indent + "{");
			lines.push(indent + '\t"module": "' + jsonEscape(entry.module) + '",');
			lines.push(indent + '\t"inNativeBoundary": ' + boolString(entry.inNativeBoundary) + ",");
			lines.push(indent + '\t"inMetalLane": ' + boolString(entry.inNativeBoundary) + ",");
			lines.push(indent + '\t"feature": "' + jsonEscape(entry.feature) + '",');
			lines.push(indent + '\t"kind": "' + jsonEscape(entry.kind) + '",');
			lines.push(indent + '\t"outcome": "' + jsonEscape(entry.outcome) + '",');
			lines.push(indent + '\t"location": "' + jsonEscape(entry.location) + '",');
			lines.push(indent + '\t"detail": "' + jsonEscape(entry.detail) + '"');
			lines.push(indent + "}" + suffix);
		}
	}

	static function appendJsonContractDiagnosticArray(lines:Array<String>, diagnostics:Array<ContractDiagnosticEntry>, indentLevel:Int):Void {
		var indent = [for (_ in 0...indentLevel) "\t"].join("");
		for (index in 0...diagnostics.length) {
			var entry = diagnostics[index];
			var suffix = index == diagnostics.length - 1 ? "" : ",";
			lines.push(indent + "{");
			lines.push(indent + '\t"module": "' + jsonEscape(entry.module) + '",');
			lines.push(indent + '\t"code": "' + jsonEscape(entry.code) + '",');
			lines.push(indent + '\t"severity": "' + jsonEscape(entry.severity) + '",');
			lines.push(indent + '\t"location": "' + jsonEscape(entry.location) + '",');
			lines.push(indent + '\t"message": "' + jsonEscape(entry.message) + '"');
			lines.push(indent + "}" + suffix);
		}
	}

	static function appendJsonContractFallbackSummaryArray(lines:Array<String>, summary:Array<ContractFallbackModuleSummary>, indentLevel:Int):Void {
		var indent = [for (_ in 0...indentLevel) "\t"].join("");
		for (index in 0...summary.length) {
			var entry = summary[index];
			var suffix = index == summary.length - 1 ? "" : ",";
			lines.push(indent + "{");
			lines.push(indent + '\t"module": "' + jsonEscape(entry.module) + '",');
			lines.push(indent + '\t"inNativeBoundary": ' + boolString(entry.inNativeBoundary) + ",");
			lines.push(indent + '\t"inMetalLane": ' + boolString(entry.inNativeBoundary) + ",");
			lines.push(indent + '\t"count": ' + entry.count);
			lines.push(indent + "}" + suffix);
		}
	}

	static function appendJsonRuntimeReasons(lines:Array<String>, reasons:Array<RuntimeFeatureReason>, indentLevel:Int):Void {
		var indent = [for (_ in 0...indentLevel) "\t"].join("");
		for (index in 0...reasons.length) {
			var reason = reasons[index];
			var suffix = index == reasons.length - 1 ? "" : ",";
			lines.push(indent + "{");
			lines.push(indent + '\t"feature": "' + jsonEscape(reason.feature) + '",');
			lines.push(indent + '\t"sourceKind": "' + jsonEscape(reason.sourceKind) + '",');
			lines.push(indent + '\t"source": "' + jsonEscape(reason.source) + '"');
			lines.push(indent + "}" + suffix);
		}
	}

	static function appendJsonOptimizerPassSelectionReasons(lines:Array<String>, reasons:Array<OptimizerPassSelectionReason>, indentLevel:Int):Void {
		var indent = [for (_ in 0...indentLevel) "\t"].join("");
		for (index in 0...reasons.length) {
			var entry = reasons[index];
			var suffix = index == reasons.length - 1 ? "" : ",";
			lines.push(indent + "{");
			lines.push(indent + '\t"pass": "' + jsonEscape(entry.pass) + '",');
			lines.push(indent + '\t"source": "' + jsonEscape(entry.source) + '",');
			lines.push(indent + '\t"reason": "' + jsonEscape(entry.reason) + '"');
			lines.push(indent + "}" + suffix);
		}
	}

	static function appendJsonOptimizerCapabilitySummaries(lines:Array<String>, capabilities:Array<OptimizerCapabilitySummary>, indentLevel:Int):Void {
		var indent = [for (_ in 0...indentLevel) "\t"].join("");
		for (index in 0...capabilities.length) {
			var entry = capabilities[index];
			var suffix = index == capabilities.length - 1 ? "" : ",";
			lines.push(indent + "{");
			lines.push(indent + '\t"id": "' + jsonEscape(entry.id) + '",');
			lines.push(indent + '\t"attempts": ' + entry.attempts + ",");
			lines.push(indent + '\t"successes": ' + entry.successes + ",");
			lines.push(indent + '\t"fallbacks": ' + entry.fallbacks + ",");
			lines.push(indent + '\t"fallbackReasonCounts": [');
			appendJsonOptimizerCapabilityFallbackReasonCounts(lines, entry.fallbackReasonCounts, indentLevel + 2);
			lines.push(indent + "\t]");
			lines.push(indent + "}" + suffix);
		}
	}

	static function appendJsonOptimizerCapabilityFallbackReasonCounts(lines:Array<String>, reasons:Array<OptimizerCapabilityFallbackReasonCount>,
			indentLevel:Int):Void {
		var indent = [for (_ in 0...indentLevel) "\t"].join("");
		for (index in 0...reasons.length) {
			var entry = reasons[index];
			var suffix = index == reasons.length - 1 ? "" : ",";
			lines.push(indent + "{");
			lines.push(indent + '\t"kind": "' + jsonEscape(entry.kind) + '",');
			lines.push(indent + '\t"count": ' + entry.count);
			lines.push(indent + "}" + suffix);
		}
	}

	static function boolString(value:Bool):String {
		return value ? "true" : "false";
	}

	static function boolLabel(value:Bool):String {
		return value ? "yes" : "no";
	}

	static function portableNativeScanModeLabel(mode:PortableNativeScanMode):String {
		return switch (mode) {
			case Typed:
				"typed";
			case Scanner:
				"scanner";
			case Hybrid:
				"hybrid";
		};
	}

	static function jsonEscape(value:String):String {
		var escaped = value == null ? "" : value;
		escaped = StringTools.replace(escaped, "\\", "\\\\");
		escaped = StringTools.replace(escaped, '"', '\\"');
		escaped = StringTools.replace(escaped, "\n", "\\n");
		escaped = StringTools.replace(escaped, "\r", "\\r");
		escaped = StringTools.replace(escaped, "\t", "\\t");
		return escaped;
	}

	static function sortedUniqueStrings(values:Array<String>):Array<String> {
		var out:Array<String> = [];
		for (value in values) {
			if (value == null) {
				continue;
			}
			var normalized = StringTools.trim(value);
			if (normalized == "" || out.indexOf(normalized) != -1) {
				continue;
			}
			out.push(normalized);
		}
		out.sort((a, b) -> a < b ? -1 : (a > b ? 1 : 0));
		return out;
	}

	static function incrementIntMap(map:Map<String, Int>, key:String):Void {
		incrementIntMapBy(map, key, 1);
	}

	static function incrementIntMapBy(map:Map<String, Int>, key:String, amount:Int):Void {
		var existing = map.exists(key) ? map.get(key) : 0;
		map.set(key, existing + amount);
	}
}
#else
class GoReflaxeCompiler {
	public function new() {}
}
#end
