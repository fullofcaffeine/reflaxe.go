package reflaxe.go;

#if (macro || reflaxe_runtime)
import haxe.io.Path;
import haxe.macro.Context;
import haxe.macro.Type;
import reflaxe.GenericCompiler;
import reflaxe.data.ClassFuncData;
import reflaxe.data.ClassVarData;
import reflaxe.data.EnumOptionData;
import reflaxe.go.analyze.GoProfileContractAnalyzer;
import reflaxe.go.analyze.GoProfileContractAnalyzer.PortableNativeScanMode;
import reflaxe.go.analyze.MetalLaneAnalyzer;
import reflaxe.go.compiler.GoBuildContext;
import reflaxe.go.compiler.GoHxrtFeatureAnalyzer;
import reflaxe.go.compiler.GoHxrtFeatureAnalyzer.GoHxrtFeatureInference;
import reflaxe.go.compiler.GoHxrtFeatureAnalyzer.GoHxrtFeatureReason;
import reflaxe.go.compiler.GoAutoLoweringModeTools;
import reflaxe.go.compiler.GoBuildContextResolver;
import reflaxe.go.compiler.GoPostBuildRunner;
import reflaxe.output.DataAndFileInfo;
import reflaxe.output.StringOrBytes;
import sys.FileSystem;
import sys.io.File;

private typedef RuntimeCopyPlan = {
	final fullCopy:Bool;
	final selectiveEnabled:Bool;
	final manualFeatures:Array<String>;
	final inferredFeatures:Array<String>;
	final features:Array<String>;
	final files:Array<String>;
	final reasons:Array<RuntimeFeatureReason>;
}

private typedef ContractReportSnapshot = {
	final schemaVersion:Int;
	final contract:String;
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
	final metalLaneModules:Array<String>;
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
	final metalFallbackViolationsByModule:Array<ContractFallbackModuleSummary>;
	final metalFallbackViolations:Array<ContractFallbackViolation>;
}

private typedef ContractLoweringDecision = {
	final feature:String;
	final kind:String;
	final outcome:String;
	final detail:String;
	final location:String;
	final module:String;
	final inMetalLane:Bool;
}

private typedef ContractDiagnosticEntry = {
	final code:String;
	final severity:String;
	final module:String;
	final location:String;
	final message:String;
}

private typedef ContractFallbackViolation = {
	final kind:String;
	final detail:String;
	final location:String;
	final module:String;
	final inMetalLane:Bool;
}

private typedef ContractFallbackModuleSummary = {
	final module:String;
	final inMetalLane:Bool;
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
	final mode:String;
	final selectiveEnabled:Bool;
	final fullCopy:Bool;
	final inferenceDisabled:Bool;
	final manualFeatures:Array<String>;
	final inferredFeatures:Array<String>;
	final selectedFeatures:Array<String>;
	final files:Array<String>;
	final reasons:Array<RuntimeFeatureReason>;
}

private typedef OptimizerPlanReportSnapshot = {
	final schemaVersion:Int;
	final contract:String;
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
	final loweringFallbackLaneCount:Int;
	final loweringFallbackNonLaneCount:Int;
	final autoLoweringCapabilities:Array<OptimizerCapabilitySummary>;
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

class GoReflaxeCompiler extends GenericCompiler<Bool, Bool, Dynamic, Dynamic, Dynamic> {
	var allModules:Array<ModuleType> = [];
	var selectedClasses:Array<ClassType> = [];
	var selectedEnums:Array<EnumType> = [];
	var generatedFiles:Array<GoCompiler.GoGeneratedFile> = [];
	var buildContext:Null<GoBuildContext> = null;
	var compilationContext:Null<CompilationContext> = null;
	var lastRuntimePlan:Null<RuntimeCopyPlan> = null;

	public function new() {
		super();
	}

	override public function filterTypes(moduleTypes:Array<ModuleType>):Array<ModuleType> {
		allModules = moduleTypes.copy();
		return moduleTypes;
	}

	override public function onCompileStart():Void {
		buildContext = GoBuildContextResolver.resolve();
		compilationContext = CompilationContext.fromBuildContext(buildContext);
		selectedClasses = [];
		selectedEnums = [];
		generatedFiles = [];
		lastRuntimePlan = null;
	}

	override public function onCompileEnd():Void {
		var resolvedBuildContext = effectiveBuildContext();
		var laneSnapshot = MetalLaneAnalyzer.collect(allModules);
		resolvedBuildContext = resolvedBuildContext.withMetalLaneModules(laneSnapshot.modules);
		buildContext = resolvedBuildContext;
		var context = CompilationContext.fromBuildContext(resolvedBuildContext);
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

		var resolvedBuildContext = effectiveBuildContext();
		output.saveFile("go.mod", buildGoMod(resolvedBuildContext.goModuleName));
		writeRuntime(output, compilationContext, resolvedBuildContext);
		emitBuildReports(output, compilationContext, resolvedBuildContext);
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

	function buildGoMod(moduleName:String):String {
		return ["module " + moduleName, "", "go 1.22", ""].join("\n");
	}

	function writeRuntime(outputManager:reflaxe.output.OutputManager, context:Null<CompilationContext>, buildContext:GoBuildContext):Void {
		var runtimeSource = Path.join([findLibraryRoot(), "runtime", "hxrt"]);
		if (!FileSystem.exists(runtimeSource) || !FileSystem.isDirectory(runtimeSource)) {
			Context.fatalError('Missing runtime directory at "' + runtimeSource + '"', Context.currentPos());
		}

		var plan = resolveRuntimeCopyPlan(context, buildContext);
		lastRuntimePlan = plan;
		if (context != null) {
			context.selectedHxrtFeatures = plan.features.copy();
		}

		if (plan.fullCopy) {
			writeRuntimeDir(outputManager, runtimeSource, "hxrt", buildContext);
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

	function writeRuntimeDir(outputManager:reflaxe.output.OutputManager, sourceDir:String, targetDir:String, buildContext:GoBuildContext):Void {
		for (entry in FileSystem.readDirectory(sourceDir)) {
			var sourcePath = Path.join([sourceDir, entry]);
			var targetPath = Path.join([targetDir, entry]);

			if (FileSystem.isDirectory(sourcePath)) {
				writeRuntimeDir(outputManager, sourcePath, targetPath, buildContext);
			} else if (shouldCopyRuntimeFileInFullMode(entry, buildContext)) {
				outputManager.saveFile(targetPath, File.getContent(sourcePath));
			}
		}
	}

	function shouldCopyRuntimeFileInFullMode(fileName:String, buildContext:GoBuildContext):Bool {
		return switch (fileName) {
			case "stack.go":
				// Native stack capture pulls in Go runtime frame machinery and must remain
				// footprint-explicit even when users request the broad hxrt runtime bundle.
				buildContext.nativeStackTraceEnabled;
			case _:
				true;
		};
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

	function resolveRuntimeCopyPlan(context:Null<CompilationContext>, buildContext:GoBuildContext):RuntimeCopyPlan {
		var forceFullCopy = buildContext.hxrtForceFullCopy;
		var selectiveEnabled = buildContext.isHxrtSelectiveEnabled();
		var manualFeatures = sortedUniqueStrings(buildContext.hxrtManualFeatures.copy());
		var inferredFeatures = new Array<String>();
		var inferredReasons = new Array<RuntimeFeatureReason>();
		if (!buildContext.hxrtNoFeatureInfer && context != null) {
			for (feature in context.inferredHxrtFeatures) {
				if (feature != null && StringTools.trim(feature) != "" && inferredFeatures.indexOf(feature) == -1) {
					inferredFeatures.push(feature);
				}
			}
			for (entry in context.inferredHxrtFeatureReasons) {
				if (entry == null) {
					continue;
				}
				inferredReasons.push({
					feature: entry.feature,
					sourceKind: entry.sourceKind,
					source: entry.source
				});
			}
		}
		inferredFeatures = sortedUniqueStrings(inferredFeatures);
		if (buildContext.nativeStackTraceEnabled && inferredFeatures.indexOf(GoHxrtFeatureAnalyzer.FEATURE_STACK) == -1) {
			inferredFeatures.push(GoHxrtFeatureAnalyzer.FEATURE_STACK);
			inferredFeatures = sortedUniqueStrings(inferredFeatures);
			inferredReasons.push({
				feature: GoHxrtFeatureAnalyzer.FEATURE_STACK,
				sourceKind: "define",
				source: GoBuildContextResolver.NATIVE_STACK_TRACE_DEFINE
			});
		}

		if (forceFullCopy || !selectiveEnabled) {
			return {
				fullCopy: true,
				selectiveEnabled: selectiveEnabled,
				manualFeatures: manualFeatures,
				inferredFeatures: inferredFeatures,
				features: [],
				files: [],
				reasons: []
			};
		}

		var selection = buildRuntimeFeatureSelection(manualFeatures, inferredFeatures, inferredReasons);
		var expanded = sortedUniqueStrings(selection.features.copy());
		var files = GoHxrtFeatureAnalyzer.filesForFeatures(expanded);
		files = sortedUniqueStrings(files);
		return {
			fullCopy: false,
			selectiveEnabled: true,
			manualFeatures: manualFeatures,
			inferredFeatures: inferredFeatures,
			features: expanded,
			files: files,
			reasons: selection.reasons
		};
	}

	function emitBuildReports(outputManager:reflaxe.output.OutputManager, context:Null<CompilationContext>, buildContext:GoBuildContext):Void {
		if (buildContext.contractReportEnabled) {
			var contractSnapshot = buildContractReportSnapshot(buildContext, context);
			outputManager.saveFile("profile_contract.json", renderContractReportJson(contractSnapshot));
			outputManager.saveFile("profile_contract.md", renderContractReportMarkdown(contractSnapshot));
		}

		if (buildContext.runtimePlanReportEnabled) {
			var runtimeSnapshot = buildRuntimePlanReportSnapshot(buildContext, context);
			outputManager.saveFile("hxrt_plan.json", renderRuntimePlanJson(runtimeSnapshot));
			outputManager.saveFile("hxrt_plan.md", renderRuntimePlanMarkdown(runtimeSnapshot));
		}

		if (buildContext.optimizerPlanReportEnabled) {
			var optimizerSnapshot = buildOptimizerPlanReportSnapshot(buildContext, context);
			outputManager.saveFile("optimizer_plan.json", renderOptimizerPlanJson(optimizerSnapshot));
			outputManager.saveFile("optimizer_plan.md", renderOptimizerPlanMarkdown(optimizerSnapshot));
		}
	}

	function buildContractReportSnapshot(buildContext:GoBuildContext, context:Null<CompilationContext>):ContractReportSnapshot {
		var contractLabel = buildContext.profile == GoProfile.Metal ? "metal" : "portable";
		var manualFeatures = sortedUniqueStrings(buildContext.hxrtManualFeatures.copy());
		var laneModules = sortedUniqueStrings(buildContext.metalLaneModules.copy());
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
		var fallbackViolations = new Array<ContractFallbackViolation>();
		var loweringDecisions = new Array<ContractLoweringDecision>();
		var laneViolationCount = 0;
		var nonLaneViolationCount = 0;
		var laneCountsByModule = new Map<String, Int>();
		var nonLaneCountsByModule = new Map<String, Int>();
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
					inMetalLane: entry.inMetalLane
				});
			}
			for (violation in context.metalFallbackViolations) {
				if (violation == null) {
					continue;
				}
				if (violation.inMetalLane) {
					laneViolationCount++;
					incrementIntMap(laneCountsByModule, violation.module);
				} else {
					nonLaneViolationCount++;
					incrementIntMap(nonLaneCountsByModule, violation.module);
				}
				fallbackViolations.push({
					kind: violation.kind,
					detail: violation.detail,
					location: violation.location,
					module: violation.module,
					inMetalLane: violation.inMetalLane
				});
			}
		}
		loweringDecisions.sort(compareContractLoweringDecision);
		fallbackViolations.sort(compareContractFallbackViolations);
		var fallbackSummary = buildContractFallbackModuleSummary(laneCountsByModule, nonLaneCountsByModule);
		return {
			schemaVersion: 7,
			contract: contractLabel,
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
			metalLaneModules: laneModules,
			loweringDecisionCount: loweringDecisions.length,
			loweringDecisionAttemptCount: loweringAttemptCount,
			loweringDecisionSuccessCount: loweringSuccessCount,
			loweringDecisionFallbackCount: loweringFallbackCount,
			loweringDecisions: loweringDecisions,
			metalFallbackViolationCount: fallbackViolations.length,
			metalFallbackLaneViolationCount: laneViolationCount,
			metalFallbackNonLaneViolationCount: nonLaneViolationCount,
			portableNativeImportScanMode: portableNativeScanModeLabel(nativeScanMode),
			portableNativeImportHitCount: portableNativeImportHits.length,
			portableNativeImportTypedHitCount: portableNativeImportTypedHits.length,
			portableNativeImportScannerHitCount: portableNativeImportScannerHits.length,
			portableNativeImportHits: portableNativeImportHits,
			portableNativeImportTypedHits: portableNativeImportTypedHits,
			portableNativeImportScannerHits: portableNativeImportScannerHits,
			contractDiagnosticCount: contractDiagnostics.length,
			contractDiagnostics: contractDiagnostics,
			metalFallbackViolationsByModule: fallbackSummary,
			metalFallbackViolations: fallbackViolations
		};
	}

	function buildRuntimePlanReportSnapshot(buildContext:GoBuildContext, context:Null<CompilationContext>):RuntimePlanReportSnapshot {
		var contractLabel = buildContext.profile == GoProfile.Metal ? "metal" : "portable";
		var plan = lastRuntimePlan == null ? resolveRuntimeCopyPlan(context, buildContext) : lastRuntimePlan;
		var selectedFeatures = sortedUniqueStrings(plan.features.copy());
		var manualFeatures = sortedUniqueStrings(plan.manualFeatures.copy());
		var inferredFeatures = sortedUniqueStrings(plan.inferredFeatures.copy());
		var files = sortedUniqueStrings(plan.files.copy());
		return {
			schemaVersion: 1,
			contract: contractLabel,
			mode: plan.fullCopy ? "full_copy" : "selective",
			selectiveEnabled: plan.selectiveEnabled,
			fullCopy: plan.fullCopy,
			inferenceDisabled: buildContext.hxrtNoFeatureInfer,
			manualFeatures: manualFeatures,
			inferredFeatures: inferredFeatures,
			selectedFeatures: selectedFeatures,
			files: files,
			reasons: plan.reasons.copy()
		};
	}

	function buildOptimizerPlanReportSnapshot(buildContext:GoBuildContext, context:Null<CompilationContext>):OptimizerPlanReportSnapshot {
		var contractLabel = buildContext.profile == GoProfile.Metal ? "metal" : "portable";
		var goAstPasses:Array<String> = [];
		var goAstPassSelectionSource = "planner";
		var goAstPassSelectionReasons:Array<OptimizerPassSelectionReason> = [];
		var loweringFallbackLaneCount = 0;
		var loweringFallbackNonLaneCount = 0;
		var autoLoweringCapabilities:Array<OptimizerCapabilitySummary> = [];
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
				if (entry.inMetalLane) {
					loweringFallbackLaneCount++;
				} else {
					loweringFallbackNonLaneCount++;
				}
			}
			autoLoweringCapabilities = buildOptimizerCapabilitySummaries(context);
		}
		return {
			schemaVersion: 5,
			contract: contractLabel,
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
			loweringFallbackLaneCount: loweringFallbackLaneCount,
			loweringFallbackNonLaneCount: loweringFallbackNonLaneCount,
			autoLoweringCapabilities: autoLoweringCapabilities
		};
	}

	function buildRuntimeFeatureSelection(manualFeatures:Array<String>, inferredFeatures:Array<String>,
			inferredReasons:Array<RuntimeFeatureReason>):GoHxrtFeatureInference {
		var selected = inferredFeatures.copy();
		var reasons:Array<GoHxrtFeatureReason> = [];
		for (entry in inferredReasons) {
			if (entry == null) {
				continue;
			}
			reasons.push({
				feature: entry.feature,
				sourceKind: entry.sourceKind,
				source: entry.source
			});
		}
		for (feature in manualFeatures) {
			if (feature == null || !GoHxrtFeatureAnalyzer.isKnownFeature(feature)) {
				continue;
			}
			if (selected.indexOf(feature) == -1) {
				selected.push(feature);
			}
			reasons.push({
				feature: feature,
				sourceKind: "manual_define",
				source: GoBuildContextResolver.HXRT_FEATURES_DEFINE
			});
		}
		return GoHxrtFeatureAnalyzer.expandWithReasons(selected, reasons);
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

	static function compareContractFallbackViolations(a:ContractFallbackViolation, b:ContractFallbackViolation):Int {
		var moduleOrder = Reflect.compare(a.module, b.module);
		if (moduleOrder != 0) {
			return moduleOrder;
		}
		var laneOrder = Reflect.compare(a.inMetalLane ? 1 : 0, b.inMetalLane ? 1 : 0);
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
		var laneOrder = Reflect.compare(a.inMetalLane ? 1 : 0, b.inMetalLane ? 1 : 0);
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

	static function buildContractFallbackModuleSummary(laneCountsByModule:Map<String, Int>,
			nonLaneCountsByModule:Map<String, Int>):Array<ContractFallbackModuleSummary> {
		var summary = new Array<ContractFallbackModuleSummary>();
		for (moduleName in laneCountsByModule.keys()) {
			summary.push({
				module: moduleName,
				inMetalLane: true,
				count: laneCountsByModule.get(moduleName)
			});
		}
		for (moduleName in nonLaneCountsByModule.keys()) {
			summary.push({
				module: moduleName,
				inMetalLane: false,
				count: nonLaneCountsByModule.get(moduleName)
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
		var laneOrder = Reflect.compare(a.inMetalLane ? 1 : 0, b.inMetalLane ? 1 : 0);
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
				var laneLabel = entry.inMetalLane ? "lane" : "non-lane";
				lines.push("- `" + entry.module + "` (" + laneLabel + ") | `" + entry.feature + "` | `" + entry.kind + "` | `" + entry.outcome + "` | `"
					+ entry.location + "` | " + entry.detail);
			}
		}
		lines.push("");
		lines.push("## metal fallback violation summary by module");
		if (snapshot.metalFallbackViolationsByModule.length == 0) {
			lines.push("- none");
		} else {
			for (entry in snapshot.metalFallbackViolationsByModule) {
				var laneLabel = entry.inMetalLane ? "lane" : "non-lane";
				lines.push("- `" + entry.module + "` (" + laneLabel + "): `" + entry.count + "`");
			}
		}
		lines.push("");
		lines.push("## metal fallback violations");
		if (snapshot.metalFallbackViolations.length == 0) {
			lines.push("- none");
		} else {
			for (entry in snapshot.metalFallbackViolations) {
				var laneLabel = entry.inMetalLane ? "lane" : "non-lane";
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
		lines.push('\t"mode": "' + jsonEscape(snapshot.mode) + '",');
		lines.push('\t"selectiveEnabled": ' + boolString(snapshot.selectiveEnabled) + ",");
		lines.push('\t"fullCopy": ' + boolString(snapshot.fullCopy) + ",");
		lines.push('\t"inferenceDisabled": ' + boolString(snapshot.inferenceDisabled) + ",");
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
		lines.push("- mode: `" + snapshot.mode + "`");
		lines.push("- selective enabled: `" + boolLabel(snapshot.selectiveEnabled) + "`");
		lines.push("- full copy: `" + boolLabel(snapshot.fullCopy) + "`");
		lines.push("- inference disabled: `" + boolLabel(snapshot.inferenceDisabled) + "`");
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
			lines.push("- full copy (`runtime/hxrt/**`, excluding footprint-explicit diagnostic files unless their define is enabled)");
		} else {
			for (fileName in snapshot.files) {
				lines.push("- `" + fileName + "`");
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
		return lines.join("\n");
	}

	static function renderOptimizerPlanJson(snapshot:OptimizerPlanReportSnapshot):String {
		var lines:Array<String> = [];
		lines.push("{");
		lines.push('\t"schemaVersion": ' + snapshot.schemaVersion + ",");
		lines.push('\t"contract": "' + jsonEscape(snapshot.contract) + '",');
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
		lines.push('\t"loweringFallbackLaneCount": ' + snapshot.loweringFallbackLaneCount + ",");
		lines.push('\t"loweringFallbackNonLaneCount": ' + snapshot.loweringFallbackNonLaneCount + ",");
		lines.push('\t"autoLoweringCapabilities": [');
		appendJsonOptimizerCapabilitySummaries(lines, snapshot.autoLoweringCapabilities, 2);
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
		lines.push("- lowering fallback lane count: `" + snapshot.loweringFallbackLaneCount + "`");
		lines.push("- lowering fallback non-lane count: `" + snapshot.loweringFallbackNonLaneCount + "`");
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
		return lines.join("\n");
	}

	static function appendJsonStringArray(lines:Array<String>, values:Array<String>, indentLevel:Int):Void {
		var indent = [for (_ in 0...indentLevel) "\t"].join("");
		for (index in 0...values.length) {
			var suffix = index == values.length - 1 ? "" : ",";
			lines.push(indent + '"' + jsonEscape(values[index]) + '"' + suffix);
		}
	}

	static function appendJsonContractFallbackArray(lines:Array<String>, violations:Array<ContractFallbackViolation>, indentLevel:Int):Void {
		var indent = [for (_ in 0...indentLevel) "\t"].join("");
		for (index in 0...violations.length) {
			var violation = violations[index];
			var suffix = index == violations.length - 1 ? "" : ",";
			lines.push(indent + "{");
			lines.push(indent + '\t"module": "' + jsonEscape(violation.module) + '",');
			lines.push(indent + '\t"inMetalLane": ' + boolString(violation.inMetalLane) + ",");
			lines.push(indent + '\t"kind": "' + jsonEscape(violation.kind) + '",');
			lines.push(indent + '\t"location": "' + jsonEscape(violation.location) + '",');
			lines.push(indent + '\t"detail": "' + jsonEscape(violation.detail) + '"');
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
			lines.push(indent + '\t"inMetalLane": ' + boolString(entry.inMetalLane) + ",");
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
			lines.push(indent + '\t"inMetalLane": ' + boolString(entry.inMetalLane) + ",");
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
