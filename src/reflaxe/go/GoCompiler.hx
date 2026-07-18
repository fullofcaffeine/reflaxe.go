package reflaxe.go;

#if macro
import haxe.macro.Context;
import haxe.macro.Expr;
import haxe.macro.Expr.Binop;
import haxe.macro.Expr.Unop;
import haxe.macro.PositionTools;
import haxe.macro.Type;
import reflaxe.go.analyze.GoProfileContractAnalyzer;
import reflaxe.go.compiler.GoAutoLoweringMode;
import reflaxe.go.compiler.GoExprOperatorOps;
import reflaxe.go.compiler.GoHxrtFeatureAnalyzer;
import reflaxe.go.compiler.GoHxrtFeatureAnalyzer.GoHxrtFeatureInference;
import reflaxe.go.compiler.GoLambdaIterableLowering;
import reflaxe.go.compiler.GoNativeTypeEligibility;
import reflaxe.go.compiler.GoNativeTypeEligibility.GoNativeEligibilityRole;
import reflaxe.go.compiler.GoNativeTypeEligibility.GoNativeTypeEligibilityResult;
import reflaxe.go.compiler.GoSourceModuleRegistry;
import reflaxe.go.compiler.GoSourceOwnedStdlibPlanner;
import reflaxe.go.compiler.GoStdlibShimClassifier;
import reflaxe.go.compiler.GoStdlibOwnership;
import reflaxe.go.compiler.GoTestAstFixtureEmitter;
import reflaxe.go.compiler.GoTypeMapper;
import reflaxe.go.compiler.emit.GoTypeReflectionEmitter;
import reflaxe.go.compiler.emit.GoRttiMetadataEmitter;
import reflaxe.go.compiler.emit.GoSerializationSourceBridgeEmitter;
import reflaxe.go.ast.GoAST.GoDecl;
import reflaxe.go.ast.GoAST.GoExpr;
import reflaxe.go.ast.GoAST.GoFile;
import reflaxe.go.ast.GoAST.GoInterfaceMethod;
import reflaxe.go.ast.GoAST.GoParam;
import reflaxe.go.ast.GoAST.GoSelectCase;
import reflaxe.go.ast.GoAST.GoSelectClause;
import reflaxe.go.ast.GoAST.GoStmt;
import reflaxe.go.ast.GoAST.GoSwitchCase;
import reflaxe.go.ast.GoAST.GoTypeSwitchCase;
import reflaxe.go.ast.GoASTPrinter;
import reflaxe.go.ast.GoASTTransformer;
import reflaxe.go.ast.GoBinaryOperator;
import reflaxe.go.ast.GoBuiltinType;
import reflaxe.go.ast.GoCompositeElement;
import reflaxe.go.ast.GoImportPath;
import reflaxe.go.ast.GoSimpleStmt;
import reflaxe.go.ast.GoType;
import reflaxe.go.ast.GoUnaryOperator;
import reflaxe.go.naming.GoNaming;
#end

typedef GoGeneratedFile = {
	final relativePath:String;
	final contents:String;
}

/**
	Why
	Go entrypoint naming, support placement, and lifecycle transforms must all use
	the same Haxe-selected class instead of independently assuming root-level
	`Main`.

	What
	Carries the selected Haxe class and source-module identities into Go lowering.

	How
	`GoReflaxeCompiler` derives both values from the authoritative typed main
	expression before constructing `GoCompiler`.
**/
typedef GoMainIdentity = {
	final className:String;
	final moduleName:String;
}

#if macro
private typedef LoweredExpr = {
	final expr:GoExpr;
	final isStringLike:Bool;
}

private typedef LoweredExprWithPrefix = {
	final prefix:Array<GoStmt>;
	final expr:GoExpr;
	final isStringLike:Bool;
}

private typedef ArrayMethodCall = {
	final target:TypedExpr;
	final methodName:String;
}

private typedef GoChanMethodCall = {
	final target:TypedExpr;
	final methodName:String;
	final elementType:Type;
}

private typedef GoSliceMethodCall = {
	final target:TypedExpr;
	final methodName:String;
	final elementType:Type;
}

private typedef GoMapMethodCall = {
	final target:TypedExpr;
	final methodName:String;
	final keyType:Type;
	final valueType:Type;
}

private typedef NativeMapTypePair = {
	final keyGoType:String;
	final valueGoType:String;
}

private typedef TypeReflectionClassMetadata = {
	final goTypeName:String;
	final haxeTypeName:String;
	final constructorSymbol:String;
	final constructible:Bool;
	final superHaxeTypeName:Null<String>;
	final staticFieldNames:Array<String>;
	final instanceFieldNames:Array<String>;
}

private typedef RttiClassMetadata = {
	final haxeTypeName:String;
	final rttiSymbol:Null<String>;
	final metaSymbol:Null<String>;
}

private typedef TypeReflectionEnumConstructorMetadata = {
	final name:String;
	final index:Int;
	final symbol:String;
	final arity:Int;
}

private typedef TypeReflectionEnumMetadata = {
	final goTypeName:String;
	final haxeTypeName:String;
	final constructors:Array<TypeReflectionEnumConstructorMetadata>;
}

private typedef GoResultMethodCall = {
	final target:TypedExpr;
	final methodName:String;
	final elementType:Type;
}

private typedef FunctionInfo = {
	final defaults:Array<Null<TypedExpr>>;
}

private typedef ConstructorBodyLowering = {
	final superArgs:Null<Array<TypedExpr>>;
	final body:Array<GoStmt>;
}

private typedef ReturnRedirect = {
	final flagName:String;
	final valueName:Null<String>;
	final valueType:Null<Type>;
}

private typedef LoopBreakTarget = {
	var label:Null<String>;
}

/**
	What:
	Represents the lowered write site for mutating a shared Haxe Array carrier or
	a native slice-shaped collection through a target expression.

	Why:
	Shared carriers need receiver capture for once-only evaluation, while raw slice
	fields may still need read -> mutate temp -> write-back when `append` replaces
	their header. Treating both as an unexamined Go lvalue would duplicate effects
	or lose a native header update.

	How:
	`lowerArrayMutationSite()` computes prefix statements, exposes the carrier/slice
	expression that later method lowering mutates, and returns a write-back closure
	for the native slice cases that must store a replacement header.
**/
private typedef ArrayMutationSite = {
	final prefix:Array<GoStmt>;
	final tempExpr:GoExpr;
	final sliceType:String;
	final writeBack:GoExpr->Array<GoStmt>;
}
#end

class GoCompiler {
	#if macro
	final compilationContext:CompilationContext;
	final mainIdentity:GoMainIdentity;
	final staticFunctionInfos:Map<String, FunctionInfo>;
	final localFunctionScopes:Array<Map<String, FunctionInfo>>;
	final localLambdaAliasScopes:Array<Map<String, String>>;
	final localRestIteratorScopes:Array<Array<String>>;
	final localArrayStorageOverrides:Map<Int, Type>;
	final requiredStdlibShimGroups:Map<String, Bool>;
	final requiredNativeChanElementTypes:Map<String, Bool>;
	final requiredNativeSliceElementTypes:Map<String, Bool>;
	final requiredNativeMapTypePairs:Map<String, NativeMapTypePair>;
	final requiredNativeResultElementTypes:Map<String, Bool>;
	final externImportPaths:Map<String, Bool>;
	final externImportPackages:Map<String, String>;
	final usedExternClassPaths:Map<String, Bool>;
	final sourceModuleRegistry:GoSourceModuleRegistry;
	final sourceOwnedStdlibPlanner:GoSourceOwnedStdlibPlanner;
	final lambdaIterableLowering:GoLambdaIterableLowering;
	final functionVarNameScopes:Array<Map<Int, String>>;
	final functionVarNameCountScopes:Array<Map<String, Int>>;
	final optionalPrimitiveParamScopes:Array<Map<Int, String>>;
	final nonNullPrimitiveLocalScopes:Array<Map<Int, String>>;
	final narrowedPrimitiveStorageScopes:Array<Map<Int, String>>;
	final localNeverReassignedScopes:Array<Map<Int, Bool>>;
	final functionReturnTypeScopes:Array<Type>;
	final returnRedirectScopes:Array<Null<ReturnRedirect>>;
	final constructorReturnScopes:Array<Bool>;
	final loopBreakTargetScopes:Array<LoopBreakTarget>;
	var switchDepth:Int;
	var throwFallbackSuppressionDepth:Int;
	final throwFallbackSuppressionDepthScopes:Array<Int>;
	var cachedVoidType:Null<Type>;
	var requiresEqualitySurface:Bool;
	var requiresReflectFieldsShim:Bool;
	var projectClasses:Array<ClassType>;
	var projectEnums:Array<EnumType>;
	final availableClassesByName:Map<String, ClassType>;
	final pendingRequiredClassesByName:Map<String, ClassType>;
	final availableEnumsByName:Map<String, EnumType>;
	final pendingRequiredEnumsByName:Map<String, EnumType>;
	final requiredSourceOwnedClassNames:Map<String, Bool>;
	final requiredNominalClassTypeNames:Map<String, Bool>;
	var globalLeafReceiverTypes:Map<String, Bool>;
	var tempVarCounter:Int;
	var requiresTypeValueSupport:Bool;
	#end

	public function new(?compilationContext:CompilationContext, ?mainIdentity:GoMainIdentity) {
		#if macro
		this.compilationContext = compilationContext == null ? new CompilationContext(GoProfile.Portable, "snapshot") : compilationContext;
		this.mainIdentity = if (mainIdentity == null || mainIdentity.className == "" || mainIdentity.moduleName == "") {
			Context.fatalError("GoCompiler requires the Haxe-selected main class identity", Context.currentPos());
			{className: "", moduleName: ""};
		} else {
			mainIdentity;
		};
		staticFunctionInfos = new Map<String, FunctionInfo>();
		localFunctionScopes = [];
		localLambdaAliasScopes = [];
		localRestIteratorScopes = [];
		localArrayStorageOverrides = new Map<Int, Type>();
		requiredStdlibShimGroups = new Map<String, Bool>();
		requiredNativeChanElementTypes = new Map<String, Bool>();
		requiredNativeSliceElementTypes = new Map<String, Bool>();
		requiredNativeMapTypePairs = new Map<String, NativeMapTypePair>();
		requiredNativeResultElementTypes = new Map<String, Bool>();
		externImportPaths = new Map<String, Bool>();
		externImportPackages = new Map<String, String>();
		usedExternClassPaths = new Map<String, Bool>();
		functionVarNameScopes = [];
		functionVarNameCountScopes = [];
		optionalPrimitiveParamScopes = [];
		nonNullPrimitiveLocalScopes = [];
		narrowedPrimitiveStorageScopes = [];
		localNeverReassignedScopes = [];
		functionReturnTypeScopes = [];
		returnRedirectScopes = [];
		constructorReturnScopes = [];
		loopBreakTargetScopes = [];
		switchDepth = 0;
		throwFallbackSuppressionDepth = 0;
		throwFallbackSuppressionDepthScopes = [];
		cachedVoidType = null;
		requiresEqualitySurface = false;
		requiresReflectFieldsShim = false;
		projectClasses = [];
		projectEnums = [];
		availableClassesByName = new Map<String, ClassType>();
		pendingRequiredClassesByName = new Map<String, ClassType>();
		availableEnumsByName = new Map<String, EnumType>();
		pendingRequiredEnumsByName = new Map<String, EnumType>();
		requiredSourceOwnedClassNames = new Map<String, Bool>();
		requiredNominalClassTypeNames = new Map<String, Bool>();
		sourceModuleRegistry = new GoSourceModuleRegistry(normalizeModuleLabel, normalizeSourcePath, sourceModuleToFilePath);
		sourceOwnedStdlibPlanner = new GoSourceOwnedStdlibPlanner({
			availableClassesByName: availableClassesByName,
			pendingRequiredClassesByName: pendingRequiredClassesByName,
			availableEnumsByName: availableEnumsByName,
			pendingRequiredEnumsByName: pendingRequiredEnumsByName,
			requiredSourceOwnedClassNames: requiredSourceOwnedClassNames,
			isCompilerOwnedAuthority: GoStdlibOwnership.isCompilerOwnedAuthority,
			fullClassName: fullClassName,
			fullEnumName: fullEnumName,
			requireStdlibShimGroup: requireStdlibShimGroup
		});
		lambdaIterableLowering = new GoLambdaIterableLowering({
			lowerExpr: lowerExpr,
			lowerToStatements: lowerToStatements,
			freshTempName: freshTempName,
			isArrayType: isArrayType,
			isHaxeArrayType: isHaxeArrayType,
			arrayElementType: arrayElementType,
			arrayElementGoType: arrayElementGoType,
			haxeDsListElementType: haxeDsListElementType,
			scalarGoType: scalarGoType,
			functionResultGoType: function(type:Type):String return isNullablePrimitiveType(type) ? "any" : scalarGoType(type),
			lowerNullableAwareTypeAssertExpr: lowerNullableAwareTypeAssertExpr,
			interfaceFieldName: interfaceFieldName,
			noteSourceOwnedStdlibUsage: noteSourceOwnedStdlibUsage,
			localVarName: localVarName,
			lookupLocalLambdaAlias: lookupLocalLambdaAlias
		});
		globalLeafReceiverTypes = new Map<String, Bool>();
		tempVarCounter = 0;
		requiresTypeValueSupport = false;
		#end
	}

	#if macro
	public function compileModule(types:Array<ModuleType>):Array<GoGeneratedFile> {
		sourceOwnedStdlibPlanner.cacheAvailableClasses(collectAllClasses(types));
		sourceOwnedStdlibPlanner.cacheAvailableEnums(collectAllEnums(types));
		return compileResolvedTypes(collectProjectClasses(types), collectProjectEnums(types));
	}

	public function compileSelectedTypes(classes:Array<ClassType>, enums:Array<EnumType>):Array<GoGeneratedFile> {
		sourceOwnedStdlibPlanner.cacheAvailableClasses(classes);
		sourceOwnedStdlibPlanner.cacheAvailableEnums(enums);
		return compileResolvedTypes(normalizeProjectClasses(classes), normalizeProjectEnums(enums));
	}

	function compileResolvedTypes(classes:Array<ClassType>, enums:Array<EnumType>):Array<GoGeneratedFile> {
		projectClasses = classes.copy();
		projectEnums = enums.copy();
		sourceOwnedStdlibPlanner.cacheAvailableClasses(classes);
		sourceOwnedStdlibPlanner.cacheAvailableEnums(enums);
		clearClassMap(pendingRequiredClassesByName);
		clearEnumMap(pendingRequiredEnumsByName);
		clearBoolMap(requiredSourceOwnedClassNames);
		clearBoolMap(requiredNominalClassTypeNames);
		sourceModuleRegistry.rebuild(classes, enums);
		globalLeafReceiverTypes = buildGlobalLeafReceiverTypes(projectClasses);
		syncCompilationContextLeafReceivers();
		clearBoolMap(compilationContext.leafReturningFunctions);
		requiresEqualitySurface = false;
		requiresReflectFieldsShim = false;
		resetRequiredNativeChanElementTypes();
		resetRequiredNativeSliceElementTypes();
		resetRequiredNativeMapTypePairs();
		resetRequiredNativeResultElementTypes();
		resetExternImportPaths();
		buildStaticFunctionInfoTable(classes);
		requiresTypeValueSupport = false;
		var moduleDecls = new Map<String, Array<GoDecl>>();
		for (enumType in enums) {
			appendModuleDecls(moduleDecls, enumType.module, lowerEnumDecls(enumType));
		}
		clearClassMap(pendingRequiredClassesByName);
		var queuedClassNames = new Map<String, Bool>();
		var classQueue = classes.copy();
		for (classType in classes) {
			queuedClassNames.set(fullClassName(classType), true);
		}
		drainPendingClassQueue(moduleDecls, classQueue, queuedClassNames, projectClasses);
		var queuedEnumNames = new Map<String, Bool>();
		for (enumType in enums) {
			queuedEnumNames.set(fullEnumName(enumType), true);
		}
		var enumQueue = new Array<EnumType>();
		for (requiredName in pendingRequiredEnumsByName.keys()) {
			if (queuedEnumNames.exists(requiredName)) {
				continue;
			}
			var requiredEnum = pendingRequiredEnumsByName.get(requiredName);
			if (requiredEnum == null) {
				continue;
			}
			queuedEnumNames.set(requiredName, true);
			enumQueue.push(requiredEnum);
			projectEnums.push(requiredEnum);
		}
		clearEnumMap(pendingRequiredEnumsByName);
		drainPendingEnumQueue(moduleDecls, enumQueue, queuedEnumNames, projectEnums);
		drainPendingClassQueue(moduleDecls, classQueue, queuedClassNames, projectClasses);
		drainPendingEnumQueue(moduleDecls, enumQueue, queuedEnumNames, projectEnums);

		// Type reflection shims emitted with stdlib symbols rely on the same runtime
		// class/enum marker structs used by TTypeExpr lowering.
		if (requiredStdlibShimGroups.exists("stdlib_symbols")) {
			requiresTypeValueSupport = true;
		}

		var preludeDecls = new Array<GoDecl>();
		if (requiresTypeValueSupport) {
			preludeDecls = preludeDecls.concat(lowerTypeValueDecls());
		}
		var supportDecls = new Array<GoDecl>();
		supportDecls = supportDecls.concat(lowerStdlibShimDecls());
		supportDecls = supportDecls.concat(lowerTestAstStmtDecls());
		populateLeafReturningFunctions(moduleDecls, preludeDecls, supportDecls);
		var requiredShimGroups = sortedRequiredStdlibShimGroups();
		compilationContext.requiredStdlibShimGroups = requiredShimGroups.copy();
		var inferredRuntimeFeatures = inferRuntimeFeatures(requiredShimGroups);
		compilationContext.inferredHxrtFeatures = inferredRuntimeFeatures.features;
		compilationContext.inferredHxrtFeatureReasons = inferredRuntimeFeatures.reasons;
		if (inferredRuntimeFeatures.features.indexOf(GoHxrtFeatureAnalyzer.FEATURE_THREAD) >= 0) {
			prependPortableThreadDrain(moduleDecls);
		}

		var supportImports = buildSupportImports();
		var testAstCase = Context.definedValue("reflaxe_go_test_ast_stmt_case");
		if (testAstCase != null && testAstCase != "") {
			supportImports = GoTestAstFixtureEmitter.imports(testAstCase).concat(supportImports);
		}
		var moduleImports = buildModuleImports();

		var moduleNames = [for (moduleName in moduleDecls.keys()) moduleName];
		moduleNames.sort(function(a, b) return Reflect.compare(a, b));

		var generated = new Array<GoGeneratedFile>();
		var usedFileNames = new Map<String, Int>();
		for (moduleName in moduleNames) {
			var moduleFileDecls = moduleDecls.get(moduleName);
			if (moduleFileDecls == null || moduleFileDecls.length == 0) {
				continue;
			}

			var fileDecls = moduleFileDecls.copy();
			var candidateImports = moduleImports.copy();
			var relativePath = moduleFilePath(moduleName, usedFileNames);
			if (moduleName == mainIdentity.moduleName) {
				fileDecls = preludeDecls.concat(fileDecls).concat(supportDecls);
				candidateImports = supportImports.concat(candidateImports);
				preludeDecls = [];
				supportDecls = [];
			}
			generated.push(renderGeneratedFile(relativePath, fileDecls, candidateImports));
		}

		if (preludeDecls.length > 0 || supportDecls.length > 0) {
			generated.push(renderGeneratedFile(nextGoFileName("support", usedFileNames), preludeDecls.concat(supportDecls), supportImports));
		}

		if (generatedUsesSharedArraySurface(generated)) {
			inferredRuntimeFeatures = GoHxrtFeatureAnalyzer.expandWithReasons(inferredRuntimeFeatures.features.concat([GoHxrtFeatureAnalyzer.FEATURE_ARRAY]),
				inferredRuntimeFeatures.reasons.concat([
					{
						feature: GoHxrtFeatureAnalyzer.FEATURE_ARRAY,
						sourceKind: "generated_surface",
						source: "hxrt.Array"
					}
				]));
		}
		compilationContext.inferredHxrtFeatures = inferredRuntimeFeatures.features;
		compilationContext.inferredHxrtFeatureReasons = inferredRuntimeFeatures.reasons;

		return generated;
	}

	/**
		What: Detects whether final generated Go references the selectively packaged
		portable Array carrier.
		Why: The Haxe typer visits unused core Array signatures even in programs that
		emit no arrays, so class-usage inference would copy `array.go` into every build.
		How: Inspect only final printer output for the closed carrier constructor/type
		markers, after DCE and source-owned planning have selected actual declarations.
	**/
	function generatedUsesSharedArraySurface(files:Array<GoGeneratedFile>):Bool {
		for (file in files) {
			if (file.contents.indexOf("*hxrt.Array") != -1
				|| file.contents.indexOf("hxrt.NewArray") != -1
				|| file.contents.indexOf("hxrt.ArrayFromValues") != -1) {
				return true;
			}
		}
		return false;
	}

	/**
		Why
		Go exits when `main` returns even if goroutines are still running, while
		portable Haxe threads are foreground threads that finish before shutdown.

		What
		Adds one deferred runtime drain to the generated Go entrypoint, but only
		when runtime feature inference proves the program uses `sys.thread`.

		How
		Rewrites the already-lowered `main` declaration through the typed Go AST.
		The runtime counter includes nested portable threads and deliberately excludes
		detached goroutines launched through `go.Go.spawn`.
	**/
	function prependPortableThreadDrain(moduleDecls:Map<String, Array<GoDecl>>):Void {
		var mainDecls = moduleDecls.get(mainIdentity.moduleName);
		if (mainDecls == null) {
			return;
		}
		for (index in 0...mainDecls.length) {
			switch (mainDecls[index]) {
				case GoDecl.GoFuncDecl("main", null, params, results, body):
					mainDecls[index] = GoDecl.GoFuncDecl("main", null, params, results,
						[GoStmt.GoDeferStmt(GoExpr.GoCall(GoExpr.GoIdent("hxrt.ThreadWaitForAll"), []))].concat(body));
					return;
				case _:
			}
		}
	}

	function appendModuleDecls(bucket:Map<String, Array<GoDecl>>, moduleName:String, decls:Array<GoDecl>):Void {
		if (decls.length == 0) {
			return;
		}
		var key = moduleName == null || moduleName == "" ? mainIdentity.moduleName : moduleName;
		if (GoStdlibOwnership.isCompilerOwnedModule(key)) {
			return;
		}
		var existing = bucket.get(key);
		if (existing == null) {
			existing = [];
			bucket.set(key, existing);
		}
		for (decl in decls) {
			existing.push(decl);
		}
	}

	function moduleFilePath(moduleName:String, usedFileNames:Map<String, Int>):String {
		if (moduleName == mainIdentity.moduleName) {
			return nextGoFileName("main", usedFileNames);
		}
		var sanitized = sanitizeFileToken(moduleName).toLowerCase();
		if (sanitized == "") {
			sanitized = "module";
		}
		return nextGoFileName("module_" + sanitized, usedFileNames);
	}

	function nextGoFileName(base:String, usedFileNames:Map<String, Int>):String {
		var key = base.toLowerCase();
		var count = usedFileNames.exists(key) ? usedFileNames.get(key) : 0;
		usedFileNames.set(key, count + 1);
		if (count == 0) {
			return base + ".go";
		}
		return base + "_" + count + ".go";
	}

	function sanitizeFileToken(value:String):String {
		var token = ~/[^A-Za-z0-9]+/g.replace(value, "_");
		while (StringTools.startsWith(token, "_")) {
			token = token.substr(1);
		}
		while (StringTools.endsWith(token, "_")) {
			token = token.substr(0, token.length - 1);
		}
		return token;
	}

	function buildSupportImports():Array<String> {
		var imports = [compilationContext.runtimeImportPath];
		if (requiredStdlibShimGroups.exists("stdlib_symbols")) {
			imports.push("reflect");
			imports.push("strings");
		}
		if (requiredStdlibShimGroups.exists("go_result")) {
			imports.push("errors");
		}
		if (requiredStdlibShimGroups.exists("go_concurrency")) {
			imports.push("reflect");
		}
		return imports;
	}

	function buildModuleImports():Array<String> {
		var imports = [compilationContext.runtimeImportPath];
		for (path in externImportPaths.keys()) {
			imports.push(path);
		}
		return imports;
	}

	function renderGeneratedFile(relativePath:String, decls:Array<GoDecl>, candidateImports:Array<String>):GoGeneratedFile {
		var file:GoFile = {
			packageName: "main",
			imports: [for (path in candidateImports) GoImportPath.parse(path)],
			decls: decls
		};
		var transformed = GoASTTransformer.transform(file, compilationContext);
		var filtered = filterImportsByUsage(transformed);
		return {
			relativePath: relativePath,
			contents: GoASTPrinter.printFile(filtered)
		};
	}

	function filterImportsByUsage(file:GoFile):GoFile {
		var dedup = new Map<String, Bool>();
		var filtered = new Array<GoImportPath>();
		for (path in file.imports) {
			if (path == null) {
				continue;
			}
			var trimmed = path.value();
			if (dedup.exists(trimmed)) {
				continue;
			}
			var alias = importAliasForPath(trimmed);
			if (alias == "" || fileUsesImportAlias(file, alias)) {
				dedup.set(trimmed, true);
				filtered.push(trimmed);
			}
		}
		filtered.sort(function(a, b) return Reflect.compare(a, b));
		return {
			packageName: file.packageName,
			imports: filtered,
			decls: file.decls
		};
	}

	function importAliasForPath(path:String):String {
		var alias = externImportPackages.exists(path) ? externImportPackages.get(path) : null;
		if (alias == null || alias == "") {
			var segments = [for (segment in path.split("/")) StringTools.trim(segment)];
			var index = segments.length - 1;
			while (index >= 0 && segments[index] == "") {
				index--;
			}
			alias = index >= 0 ? segments[index] : "";
		}
		return alias == null || alias == "" ? "" : normalizeIdent(alias);
	}

	function fileUsesImportAlias(file:GoFile, alias:String):Bool {
		for (decl in file.decls) {
			if (declUsesImportAlias(decl, alias)) {
				return true;
			}
		}
		return false;
	}

	function declUsesImportAlias(decl:GoDecl, alias:String):Bool {
		return switch (decl) {
			case GoInterfaceDecl(_, methods):
				var used = false;
				for (method in methods) {
					for (param in method.params) {
						if (typeNameUsesImportAlias(param.typeName, alias)) {
							used = true;
							break;
						}
					}
					if (used) {
						break;
					}
					for (result in method.results) {
						if (typeNameUsesImportAlias(result, alias)) {
							used = true;
							break;
						}
					}
					if (used) {
						break;
					}
				}
				used;
			case GoStructDecl(_, fields):
				var used = false;
				for (field in fields) {
					if (typeNameUsesImportAlias(field.typeName, alias)) {
						used = true;
						break;
					}
				}
				used;
			case GoGlobalVarDecl(_, typeName, value): typeNameUsesImportAlias(typeName, alias) || (value != null && exprUsesImportAlias(value, alias));
			case GoFuncDecl(_, receiver, params, results, body):
				if (receiver != null && typeNameUsesImportAlias(receiver.typeName, alias)) {
					true;
				} else {
					var used = false;
					for (param in params) {
						if (typeNameUsesImportAlias(param.typeName, alias)) {
							used = true;
							break;
						}
					}
					if (!used) {
						for (result in results) {
							if (typeNameUsesImportAlias(result, alias)) {
								used = true;
								break;
							}
						}
					}
					if (!used) {
						for (stmt in body) {
							if (stmtUsesImportAlias(stmt, alias)) {
								used = true;
								break;
							}
						}
					}
					used;
				}
		};
	}

	function stmtUsesImportAlias(stmt:GoStmt, alias:String):Bool {
		return switch (stmt) {
			case GoVarDecl(_, typeName, value, _): typeName != null && typeNameUsesImportAlias(typeName,
					alias) || (value != null && exprUsesImportAlias(value, alias));
			case GoMultiAssign(_, value, _): exprUsesImportAlias(value, alias);
			case GoAssign(left, right, _): exprUsesImportAlias(left, alias) || exprUsesImportAlias(right, alias);
			case GoIncDec(target, _): exprUsesImportAlias(target, alias);
			case GoExprStmt(expr):
				exprUsesImportAlias(expr, alias);
			case GoGoStmt(call):
				exprUsesImportAlias(call, alias);
			case GoDeferStmt(call):
				exprUsesImportAlias(call, alias);
			case GoSendStmt(channel, value): exprUsesImportAlias(channel, alias) || exprUsesImportAlias(value, alias);
			case GoRaw(code):
				rawCodeUsesImportAlias(code, alias);
			case GoWhile(cond, body):
				if (exprUsesImportAlias(cond, alias)) {
					true;
				} else {
					var used = false;
					for (child in body) {
						if (stmtUsesImportAlias(child, alias)) {
							used = true;
							break;
						}
					}
					used;
				}
			case GoFor(initializer, condition, post, body):
				(initializer != null && simpleStmtUsesImportAlias(initializer, alias))
				|| (condition != null && exprUsesImportAlias(condition, alias))
				|| (post != null && simpleStmtUsesImportAlias(post, alias))
				|| stmtListUsesImportAlias(body, alias);
			case GoLabeled(_, child):
				stmtUsesImportAlias(child, alias);
			case GoRangeStmt(_, _, source, _, body):
				if (exprUsesImportAlias(source, alias)) {
					true;
				} else {
					var used = false;
					for (child in body) {
						if (stmtUsesImportAlias(child, alias)) {
							used = true;
							break;
						}
					}
					used;
				}
			case GoIf(cond, thenBody, elseBody):
				if (exprUsesImportAlias(cond, alias)) {
					true;
				} else {
					var used = false;
					for (child in thenBody) {
						if (stmtUsesImportAlias(child, alias)) {
							used = true;
							break;
						}
					}
					if (!used && elseBody != null) {
						for (child in elseBody) {
							if (stmtUsesImportAlias(child, alias)) {
								used = true;
								break;
							}
						}
					}
					used;
				}
			case GoSwitch(value, cases, defaultBody):
				if (exprUsesImportAlias(value, alias)) {
					true;
				} else {
					var used = false;
					for (switchCase in cases) {
						for (caseValue in switchCase.values) {
							if (exprUsesImportAlias(caseValue, alias)) {
								used = true;
								break;
							}
						}
						if (used) {
							break;
						}
						for (child in switchCase.body) {
							if (stmtUsesImportAlias(child, alias)) {
								used = true;
								break;
							}
						}
						if (used) {
							break;
						}
					}
					if (!used && defaultBody != null) {
						for (child in defaultBody) {
							if (stmtUsesImportAlias(child, alias)) {
								used = true;
								break;
							}
						}
					}
					used;
				}
			case GoTypeSwitch(value, _, cases, defaultBody):
				if (exprUsesImportAlias(value, alias)) {
					true;
				} else {
					var used = false;
					for (typeCase in cases) {
						if (typeNameUsesImportAlias(typeCase.typeName, alias)) {
							used = true;
							break;
						}
						for (child in typeCase.body) {
							if (stmtUsesImportAlias(child, alias)) {
								used = true;
								break;
							}
						}
						if (used) {
							break;
						}
					}
					if (!used && defaultBody != null) {
						for (child in defaultBody) {
							if (stmtUsesImportAlias(child, alias)) {
								used = true;
								break;
							}
						}
					}
					used;
				}
			case GoSelect(cases):
				var used = false;
				for (selectCase in cases) {
					if (selectClauseUsesImportAlias(selectCase.clause, alias)) {
						used = true;
						break;
					}
					for (child in selectCase.body) {
						if (stmtUsesImportAlias(child, alias)) {
							used = true;
							break;
						}
					}
					if (used) {
						break;
					}
				}
				used;
			case GoBreak(_), GoContinue:
				false;
			case GoReturn(expr): expr != null && exprUsesImportAlias(expr, alias);
		};
	}

	function stmtListUsesImportAlias(stmts:Array<GoStmt>, alias:String):Bool {
		for (stmt in stmts) {
			if (stmtUsesImportAlias(stmt, alias)) {
				return true;
			}
		}
		return false;
	}

	function simpleStmtUsesImportAlias(stmt:GoSimpleStmt, alias:String):Bool {
		return switch (stmt) {
			case GoSimpleDeclare(_, value): exprUsesImportAlias(value, alias);
			case GoSimpleAssign(left, right, _): exprUsesImportAlias(left, alias) || exprUsesImportAlias(right, alias);
			case GoSimpleIncDec(target, _): exprUsesImportAlias(target, alias);
			case GoSimpleExpr(expr): exprUsesImportAlias(expr, alias);
			case GoSimpleSend(channel, value): exprUsesImportAlias(channel, alias) || exprUsesImportAlias(value, alias);
		};
	}

	function selectClauseUsesImportAlias(clause:GoSelectClause, alias:String):Bool {
		return switch (clause) {
			case GoSelectSend(channel, value): exprUsesImportAlias(channel, alias) || exprUsesImportAlias(value, alias);
			case GoSelectRecv(recv):
				exprUsesImportAlias(recv, alias);
			case GoSelectRecvAssign(target, recv, _): exprUsesImportAlias(target, alias) || exprUsesImportAlias(recv, alias);
			case GoSelectRecvAssignOk(target, okTarget, recv, _): exprUsesImportAlias(target,
					alias) || exprUsesImportAlias(okTarget, alias) || exprUsesImportAlias(recv, alias);
			case GoSelectDefault:
				false;
		};
	}

	function exprUsesImportAlias(expr:GoExpr, alias:String):Bool {
		return switch (expr) {
			case GoIdent(name): name == alias || rawCodeUsesImportAlias(name, alias);
			case GoIntLiteral(_), GoFloatLiteral(_), GoBoolLiteral(_), GoStringLiteral(_), GoNil:
				false;
			case GoSelector(target, _):
				exprUsesImportAlias(target, alias);
			case GoIndex(target, index): exprUsesImportAlias(target, alias) || exprUsesImportAlias(index, alias);
			case GoSlice(target, start, end): exprUsesImportAlias(target,
					alias) || (start != null && exprUsesImportAlias(start, alias)) || (end != null && exprUsesImportAlias(end, alias));
			case GoCompositeLiteral(typeName, elements):
				if (typeNameUsesImportAlias(typeName, alias)) {
					true;
				} else {
					var used = false;
					for (element in elements) {
						if (compositeElementUsesImportAlias(element, alias)) {
							used = true;
							break;
						}
					}
					used;
				}
			case GoMakeSlice(elementType, length, capacity): typeNameUsesImportAlias(elementType,
					alias) || exprUsesImportAlias(length, alias) || (capacity != null && exprUsesImportAlias(capacity, alias));
			case GoFuncLiteral(params, results, body):
				var used = false;
				for (param in params) {
					if (typeNameUsesImportAlias(param.typeName, alias)) {
						used = true;
						break;
					}
				}
				if (!used) {
					for (result in results) {
						if (typeNameUsesImportAlias(result, alias)) {
							used = true;
							break;
						}
					}
				}
				if (!used) {
					for (stmt in body) {
						if (stmtUsesImportAlias(stmt, alias)) {
							used = true;
							break;
						}
					}
				}
				used;
			case GoRaw(code):
				rawCodeUsesImportAlias(code, alias);
			case GoTypeAssert(inner, typeName): exprUsesImportAlias(inner, alias) || typeNameUsesImportAlias(typeName, alias);
			case GoRecvExpr(channel):
				exprUsesImportAlias(channel, alias);
			case GoUnary(_, inner):
				exprUsesImportAlias(inner, alias);
			case GoBinary(_, left, right): exprUsesImportAlias(left, alias) || exprUsesImportAlias(right, alias);
			case GoCall(callee, args):
				if (exprUsesImportAlias(callee, alias)) {
					true;
				} else {
					var used = false;
					for (arg in args) {
						if (exprUsesImportAlias(arg, alias)) {
							used = true;
							break;
						}
					}
					used;
				}
		};
	}

	function compositeElementUsesImportAlias(element:GoCompositeElement, alias:String):Bool {
		return switch (element) {
			case GoCompositeValue(value), GoCompositeField(_, value): exprUsesImportAlias(value, alias);
			case GoCompositeKeyValue(key, value): exprUsesImportAlias(key, alias) || exprUsesImportAlias(value, alias);
		};
	}

	function typeNameUsesImportAlias(typeName:GoType, alias:String):Bool {
		if (typeName == null) {
			return false;
		}
		return typeName.usesPackage(alias);
	}

	function rawCodeUsesImportAlias(code:String, alias:String):Bool {
		if (code == null || code == "") {
			return false;
		}
		return new EReg("\\b" + EReg.escape(alias) + "\\s*\\.", "").match(code);
	}

	function collectProjectClasses(types:Array<ModuleType>):Array<ClassType> {
		var collected = new Array<ClassType>();
		for (moduleType in types) {
			switch (moduleType) {
				case TClassDecl(classRef):
					collected.push(classRef.get());
				case _:
			}
		}
		return normalizeProjectClasses(collected);
	}

	function collectAllClasses(types:Array<ModuleType>):Array<ClassType> {
		var collected = new Array<ClassType>();
		for (moduleType in types) {
			switch (moduleType) {
				case TClassDecl(classRef):
					collected.push(classRef.get());
				case _:
			}
		}
		return collected;
	}

	function normalizeProjectClasses(classes:Array<ClassType>):Array<ClassType> {
		var dedup = new Map<String, ClassType>();
		for (classType in classes) {
			if (!isProjectClass(classType)) {
				continue;
			}
			var className = fullClassName(classType);
			if (!dedup.exists(className)) {
				dedup.set(className, classType);
			}
		}

		var normalized = [for (classType in dedup) classType];
		normalized.sort(function(a, b) return Reflect.compare(fullClassName(a), fullClassName(b)));
		ensureSelectedMainClass(normalized);
		return normalized;
	}

	/**
		Why
		Manual DCE may pass only selected project types into this compiler, so a stale
		or missing selected-main identity must fail before codegen can silently emit a
		library-shaped Go package.

		What
		Confirms that the Haxe-selected entry class is present in the normalized
		project class set.

		How
		Matches both class and module identities; no other static `main` method can
		satisfy this contract.
	**/
	function ensureSelectedMainClass(classes:Array<ClassType>):Void {
		for (classType in classes) {
			if (isSelectedMainClass(classType)) {
				return;
			}
		}
		Context.fatalError('Haxe-selected main class "' + mainIdentity.className + '" was not found among project modules', Context.currentPos());
	}

	function isSelectedMainClass(classType:ClassType):Bool {
		return classType.module == mainIdentity.moduleName && fullClassName(classType) == mainIdentity.className;
	}

	function collectProjectEnums(types:Array<ModuleType>):Array<EnumType> {
		var collected = new Array<EnumType>();
		for (moduleType in types) {
			switch (moduleType) {
				case TEnumDecl(enumRef):
					collected.push(enumRef.get());
				case _:
			}
		}
		return normalizeProjectEnums(collected);
	}

	function collectAllEnums(types:Array<ModuleType>):Array<EnumType> {
		var collected = new Array<EnumType>();
		for (moduleType in types) {
			switch (moduleType) {
				case TEnumDecl(enumRef):
					collected.push(enumRef.get());
				case _:
			}
		}
		return collected;
	}

	function normalizeProjectEnums(enums:Array<EnumType>):Array<EnumType> {
		var dedup = new Map<String, EnumType>();
		for (enumType in enums) {
			if (!isProjectEnum(enumType)) {
				continue;
			}
			var enumName = fullEnumName(enumType);
			if (!dedup.exists(enumName)) {
				dedup.set(enumName, enumType);
			}
		}
		var normalized = [for (enumType in dedup) enumType];
		normalized.sort(function(a, b) return Reflect.compare(fullEnumName(a), fullEnumName(b)));
		return normalized;
	}

	function isProjectClass(classType:ClassType):Bool {
		if (isRequiredStdlibClass(classType) || requiredSourceOwnedClassNames.exists(fullClassName(classType))) {
			return true;
		}

		if (classType.isExtern) {
			return false;
		}

		var moduleName = classType.module;
		if (StringTools.startsWith(moduleName, "haxe.")
			|| StringTools.startsWith(moduleName, "sys.")
			|| StringTools.startsWith(moduleName, "StdTypes")
			|| StringTools.startsWith(moduleName, "reflaxe.go")) {
			return false;
		}

		var file = Std.string(PositionTools.toLocation(classType.pos).file);
		if (file == null) {
			return false;
		}
		if (StringTools.contains(file, "/std/") || StringTools.contains(file, "/vendor/")) {
			return false;
		}

		return true;
	}

	function isProjectEnum(enumType:EnumType):Bool {
		if (isRequiredStdlibEnum(enumType)) {
			return true;
		}

		if (enumType.isExtern) {
			return false;
		}

		var moduleName = enumType.module;
		if (StringTools.startsWith(moduleName, "haxe.")
			|| StringTools.startsWith(moduleName, "sys.")
			|| StringTools.startsWith(moduleName, "StdTypes")
			|| StringTools.startsWith(moduleName, "reflaxe.go")) {
			return false;
		}

		var file = Std.string(PositionTools.toLocation(enumType.pos).file);
		if (file == null) {
			return false;
		}
		if (StringTools.contains(file, "/std/") || StringTools.contains(file, "/vendor/")) {
			return false;
		}

		return true;
	}

	function isRequiredStdlibClass(classType:ClassType):Bool {
		var pack = classType.pack.join(".");
		return (pack == "haxe" && classType.name == "Int64Helper")
			|| (pack == "haxe" && classType.name == "Json")
			|| (pack == "haxe" && classType.name == "IMap")
			|| (pack == "haxe._EnumFlags" && classType.name == "EnumFlags_Impl_")
			|| (pack == "haxe.format" && (classType.name == "JsonParser" || classType.name == "JsonPrinter"))
			|| (pack == "haxe._Int64" && (classType.name == "Int64_Impl_" || classType.name == "___Int64"))
			|| (pack == "haxe._Int32" && classType.name == "Int32_Impl_")
			|| (pack == "haxe.ds" && classType.name == "HashMap")
			|| (pack == "haxe.ds._HashMap" && classType.name == "HashMapData")
			|| (pack == "haxe.iterators" && classType.name == "HashMapKeyValueIterator")
			|| (pack == "" && classType.name == "Lambda")
			|| (pack == "go" && (classType.name == "Go" || classType.name == "Chan" || classType.name == "Select"));
	}

	function isRequiredStdlibEnum(enumType:EnumType):Bool {
		return false;
	}

	function fullClassName(classType:ClassType):String {
		return classType.pack.length == 0 ? classType.name : classType.pack.join(".") + "." + classType.name;
	}

	function fullEnumName(enumType:EnumType):String {
		return enumType.pack.length == 0 ? enumType.name : enumType.pack.join(".") + "." + enumType.name;
	}

	function goRawQuotedString(value:String):String {
		var escaped = StringTools.replace(value, "\\", "\\\\");
		escaped = StringTools.replace(escaped, "\"", "\\\"");
		escaped = StringTools.replace(escaped, "\n", "\\n");
		escaped = StringTools.replace(escaped, "\r", "\\r");
		escaped = StringTools.replace(escaped, "\t", "\\t");
		return "\"" + escaped + "\"";
	}

	function goStringArrayCarrierLiteral(values:Array<String>):String {
		if (values.length == 0) {
			return "hxrt.NewArray()";
		}
		var entries = [for (value in values) "hxrt.StringFromLiteral(" + goRawQuotedString(value) + ")"];
		return "hxrt.NewArray(" + entries.join(", ") + ")";
	}

	/**
		What: materialize the backend-owned `haxe.Resource.content` table.

		Why: `haxe.Resource` methods can come from source-owned std inclusion, but the
		actual resource payloads are exposed to targets through compiler resources
		(`Context.getResources()` / `__resources__()`), not reusable Haxe source. If we
		do nothing, generated Go has the helper methods but an empty content table.

		How: sort resource names for deterministic output and place the existing
		`{name,data,str}` records in the shared Array carrier, storing every payload in
		the `data` field as base64 so both text and binary resources flow through the
		stdlib `getString` / `getBytes` decode paths unchanged.
	**/
	function haxeResourceContentLiteral():GoExpr {
		var resources = Context.getResources();
		var names = [for (name in resources.keys()) name];
		names.sort(function(a, b) return Reflect.compare(a, b));
		if (names.length == 0) {
			return GoExpr.GoCall(GoExpr.GoIdent("hxrt.NewArray"), []);
		}

		var entries = new Array<String>();
		for (name in names) {
			var bytes = resources.get(name);
			var encoded = bytes == null ? "" : haxe.crypto.Base64.encode(bytes);
			entries.push('map[string]any{"name": hxrt.StringFromLiteral(' + goRawQuotedString(name) + '), "data": hxrt.StringFromLiteral('
				+ goRawQuotedString(encoded) + '), "str": nil}');
		}
		return GoExpr.GoRaw("hxrt.NewArray(" + entries.join(", ") + ")");
	}

	function classHasInstanceLayout(classType:ClassType):Bool {
		var instanceDataCount = 0;
		var instanceMethodCount = 0;
		for (field in classType.fields.get()) {
			switch (field.kind) {
				case FVar(_, _):
					instanceDataCount++;
				case FMethod(MethDynamic):
					instanceDataCount++;
				case FMethod(_):
					if (field.name != "new" && unwrapFunction(field.expr()) != null) {
						instanceMethodCount++;
					}
			}
		}
		var hasCtor = false;
		if (classType.constructor != null) {
			hasCtor = unwrapFunction(classType.constructor.get().expr()) != null;
		}
		return projectSuperClass(classType) != null || instanceDataCount > 0 || instanceMethodCount > 0 || hasCtor;
	}

	function collectClassStaticFieldNames(classType:ClassType):Array<String> {
		var names = new Array<String>();
		for (field in classType.statics.get()) {
			switch (field.kind) {
				case FVar(_, _):
					if (field.name != "__init__") {
						names.push(field.name);
					}
				case FMethod(_):
					names.push(field.name);
			}
		}
		names.sort(function(a, b) return Reflect.compare(a, b));
		return names;
	}

	function appendClassInstanceFieldNames(classType:ClassType, names:Map<String, Bool>, out:Array<String>):Void {
		if (classType.superClass != null) {
			appendClassInstanceFieldNames(classType.superClass.t.get(), names, out);
		}
		for (field in classType.fields.get()) {
			var include = switch (field.kind) {
				case FVar(_, _):
					true;
				case FMethod(_):
					field.name != "new";
			}
			if (include && !names.exists(field.name)) {
				names.set(field.name, true);
				out.push(field.name);
			}
		}
	}

	function collectClassInstanceFieldNames(classType:ClassType):Array<String> {
		var names = new Map<String, Bool>();
		var out = new Array<String>();
		appendClassInstanceFieldNames(classType, names, out);
		out.sort(function(a, b) return Reflect.compare(a, b));
		return out;
	}

	function typeReflectionClassMetadata():Array<TypeReflectionClassMetadata> {
		var entries = new Array<TypeReflectionClassMetadata>();
		for (classType in projectClasses) {
			if (classType.isExtern || classType.isInterface) {
				continue;
			}
			switch (classType.kind) {
				case KTypeParameter(_):
					continue;
				case _:
			}
			var constructible = classHasInstanceLayout(classType);
			var superName:Null<String> = null;
			if (classType.superClass != null) {
				superName = fullClassName(classType.superClass.t.get());
			}
			entries.push({
				goTypeName: classTypeName(classType),
				haxeTypeName: fullClassName(classType),
				constructorSymbol: constructible ? constructorSymbol(classType) : "",
				constructible: constructible,
				superHaxeTypeName: superName,
				staticFieldNames: collectClassStaticFieldNames(classType),
				instanceFieldNames: constructible ? collectClassInstanceFieldNames(classType) : []
			});
		}
		entries.sort(function(a, b) return Reflect.compare(a.goTypeName, b.goTypeName));
		return entries;
	}

	function rttiClassMetadata():Array<RttiClassMetadata> {
		var entries = new Array<RttiClassMetadata>();
		for (classType in projectClasses) {
			if (classType.isExtern || classType.isInterface) {
				continue;
			}
			switch (classType.kind) {
				case KTypeParameter(_):
					continue;
				case _:
			}
			var rttiSymbol:Null<String> = null;
			var metaSymbol:Null<String> = null;
			for (field in classType.statics.get()) {
				switch (field.kind) {
					case FVar(_, _):
						if (field.name == "__rtti") {
							rttiSymbol = staticSymbol(classType, field.name);
						} else if (field.name == "__meta__") {
							metaSymbol = staticSymbol(classType, field.name);
						}
					case _:
				}
			}
			if (rttiSymbol != null || metaSymbol != null) {
				entries.push({
					haxeTypeName: fullClassName(classType),
					rttiSymbol: rttiSymbol,
					metaSymbol: metaSymbol
				});
			}
		}
		entries.sort(function(a, b) return Reflect.compare(a.haxeTypeName, b.haxeTypeName));
		return entries;
	}

	function typeReflectionEnumMetadata():Array<TypeReflectionEnumMetadata> {
		var entries = new Array<TypeReflectionEnumMetadata>();
		for (enumType in projectEnums) {
			if (enumType.isExtern) {
				continue;
			}
			var constructors = new Array<TypeReflectionEnumConstructorMetadata>();
			for (constructor in enumType.constructs) {
				constructors.push({
					name: constructor.name,
					index: constructor.index,
					symbol: enumConstructorSymbol(enumType, constructor.name),
					arity: enumConstructorArgs(constructor.type).length
				});
			}
			constructors.sort(function(a, b) return a.index - b.index);
			entries.push({
				goTypeName: enumTypeName(enumType),
				haxeTypeName: fullEnumName(enumType),
				constructors: constructors
			});
		}
		entries.sort(function(a, b) return Reflect.compare(a.goTypeName, b.goTypeName));
		return entries;
	}

	/**
		What: Resolve the superclass whose layout participates in generated Go.

		Why: Go needs an explicit embedded-base selector for Haxe nominal upcasts.
		A user class can extend staged std even when manual DCE did not initially put
		that std class in the project queue, so selection order must not erase the
		inheritance link.

		How: Promote typed source-backed std superclasses through the planner, then
		apply the existing project, required-source, and compiler-owned embedding
		authorities.
	**/
	function projectSuperClass(classType:ClassType):Null<ClassType> {
		if (classType.superClass == null) {
			return null;
		}
		if (GoStdlibOwnership.isCompilerOwnedAuthority(fullClassName(classType))) {
			// A compiler-owned carrier controls its own concrete Go layout. Promoting a
			// typed Haxe superclass here would claim an embedding that the synthetic
			// declaration does not actually emit and would force invalid virtual calls.
			return null;
		}
		var superType = classType.superClass.t.get();
		var superName = fullClassName(superType);
		if (sourceOwnedStdlibPlanner.hasLoadedSourceOwnedStdlibClass(superName)
			|| sourceOwnedStdlibPlanner.requireSourceOwnedStdlibSuperclass(superType)) {
			requireSourceOwnedStdlibClass(superName);
			return superType;
		}
		return (isProjectClass(superType)
			|| requiredSourceOwnedClassNames.exists(superName)
			|| GoStdlibOwnership.isEmbeddableCompilerOwnedSuper(superName)) ? superType : null;
	}

	function buildGlobalLeafReceiverTypes(classes:Array<ClassType>):Map<String, Bool> {
		var hasSubclass = new Map<String, Bool>();
		for (classType in classes) {
			if (classType.isExtern || classType.isInterface) {
				continue;
			}
			var superClass = projectSuperClass(classType);
			if (superClass != null) {
				hasSubclass.set(fullClassName(superClass), true);
			}
		}

		var out = new Map<String, Bool>();
		for (classType in classes) {
			if (classType.isExtern || classType.isInterface) {
				continue;
			}
			var className = fullClassName(classType);
			if (!hasSubclass.exists(className)) {
				out.set("*" + classTypeName(classType), true);
			}
		}
		return out;
	}

	function syncCompilationContextLeafReceivers():Void {
		clearBoolMap(compilationContext.leafReceiverTypes);
		for (typeName in globalLeafReceiverTypes.keys()) {
			compilationContext.leafReceiverTypes.set(typeName, true);
		}
	}

	function populateLeafReturningFunctions(moduleDecls:Map<String, Array<GoDecl>>, preludeDecls:Array<GoDecl>, supportDecls:Array<GoDecl>):Void {
		clearBoolMap(compilationContext.leafReturningFunctions);
		for (decl in preludeDecls) {
			noteLeafReturningFunctionDecl(decl);
		}
		for (decl in supportDecls) {
			noteLeafReturningFunctionDecl(decl);
		}
		for (moduleName in moduleDecls.keys()) {
			var decls = moduleDecls.get(moduleName);
			if (decls == null) {
				continue;
			}
			for (decl in decls) {
				noteLeafReturningFunctionDecl(decl);
			}
		}
	}

	function noteLeafReturningFunctionDecl(decl:GoDecl):Void {
		switch (decl) {
			case GoDecl.GoFuncDecl(name, receiver, _, results, _):
				if (receiver == null && results.length == 1 && globalLeafReceiverTypes.exists(results[0].render())) {
					compilationContext.leafReturningFunctions.set(name, true);
				}
			case _:
		}
	}

	function clearBoolMap(map:Map<String, Bool>):Void {
		var keys = [for (key in map.keys()) key];
		for (key in keys) {
			map.remove(key);
		}
	}

	function clearStringMap(map:Map<String, String>):Void {
		var keys = [for (key in map.keys()) key];
		for (key in keys) {
			map.remove(key);
		}
	}

	function clearClassMap(map:Map<String, ClassType>):Void {
		var keys = [for (key in map.keys()) key];
		for (key in keys) {
			map.remove(key);
		}
	}

	function clearEnumMap(map:Map<String, EnumType>):Void {
		var keys = [for (key in map.keys()) key];
		for (key in keys) {
			map.remove(key);
		}
	}

	function interfaceSymbol(classType:ClassType):String {
		return "I_" + classTypeName(classType);
	}

	function buildStaticFunctionInfoTable(classes:Array<ClassType>):Void {
		for (classType in classes) {
			var fields = classType.statics.get();
			for (field in fields) {
				var func = unwrapFunction(field.expr());
				if (func != null) {
					staticFunctionInfos.set(staticSymbol(classType, field.name), buildFunctionInfo(func));
				}
			}
		}
	}

	function lowerEnums(enums:Array<EnumType>):Array<GoDecl> {
		var decls = new Array<GoDecl>();
		for (enumType in enums) {
			decls = decls.concat(lowerEnumDecls(enumType));
		}
		return decls;
	}

	function lowerEnumDecls(enumType:EnumType):Array<GoDecl> {
		if (GoStdlibOwnership.isCompilerOwnedAuthority(fullEnumName(enumType))) {
			return [];
		}
		var decls = new Array<GoDecl>();
		var enumName = enumTypeName(enumType);
		decls.push(GoDecl.GoStructDecl(enumName, [{name: "tag", typeName: "int"}, {name: "params", typeName: "[]any"}]));

		var constructors = [for (field in enumType.constructs) field];
		constructors.sort(function(a, b) return a.index - b.index);

		for (constructor in constructors) {
			var symbol = enumConstructorSymbol(enumType, constructor.name);
			var ctorArgs = enumConstructorArgs(constructor.type);
			var enumLiteral = GoExpr.GoUnary(GoUnaryOperator.AddressOf, GoExpr.GoCompositeLiteral(GoType.named(enumName), [
				GoCompositeElement.GoCompositeField("tag", GoExpr.GoIntLiteral(constructor.index))
			]));
			if (ctorArgs.length == 0) {
				decls.push(GoDecl.GoGlobalVarDecl(symbol, "*" + enumName, enumLiteral));
			} else {
				var params = new Array<GoParam>();
				var payloadExprs = new Array<GoExpr>();
				for (index in 0...ctorArgs.length) {
					var arg = ctorArgs[index];
					var argName = normalizeIdent(arg.name == "" ? ("arg" + index) : arg.name);
					params.push({
						name: argName,
						typeName: nilCapablePrimitiveParamType(arg.t)
					});
					payloadExprs.push(GoExpr.GoIdent(argName));
				}

				decls.push(GoDecl.GoFuncDecl(symbol, null, params, ["*" + enumName], [
					GoStmt.GoVarDecl("enumValue", null, enumLiteral, true),
					GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("enumValue"), "params"),
						GoExpr.GoCompositeLiteral(GoType.slice(GoType.builtin(GoBuiltinType.AnyType)),
							[for (payload in payloadExprs) GoCompositeElement.GoCompositeValue(payload)])),
					GoStmt.GoReturn(GoExpr.GoIdent("enumValue"))
				]));
			}
		}
		return decls;
	}

	function enumConstructorArgs(type:Type):Array<{name:String, opt:Bool, t:Type}> {
		var followed = Context.follow(type);
		return switch (followed) {
			case TFun(args, _):
				args;
			case _:
				[];
		};
	}

	function lowerClasses(classes:Array<ClassType>):Array<GoDecl> {
		var decls = new Array<GoDecl>();
		for (classType in classes) {
			decls = decls.concat(lowerClassDecls(classType));
		}
		return decls;
	}

	function drainPendingClassQueue(moduleDecls:Map<String, Array<GoDecl>>, classQueue:Array<ClassType>, queuedClassNames:Map<String, Bool>,
			projectClasses:Array<ClassType>):Void {
		for (requiredName in pendingRequiredClassesByName.keys()) {
			if (queuedClassNames.exists(requiredName)) {
				continue;
			}
			var requiredClass = pendingRequiredClassesByName.get(requiredName);
			if (requiredClass == null) {
				continue;
			}
			queuedClassNames.set(requiredName, true);
			classQueue.push(requiredClass);
			projectClasses.push(requiredClass);
		}
		clearClassMap(pendingRequiredClassesByName);
		while (classQueue.length > 0) {
			var classType = classQueue.shift();
			appendModuleDecls(moduleDecls, classType.module, lowerClassDecls(classType));
			for (requiredName in pendingRequiredClassesByName.keys()) {
				if (queuedClassNames.exists(requiredName)) {
					continue;
				}
				var requiredClass = pendingRequiredClassesByName.get(requiredName);
				if (requiredClass == null) {
					continue;
				}
				queuedClassNames.set(requiredName, true);
				classQueue.push(requiredClass);
				projectClasses.push(requiredClass);
			}
			clearClassMap(pendingRequiredClassesByName);
		}
	}

	function drainPendingEnumQueue(moduleDecls:Map<String, Array<GoDecl>>, enumQueue:Array<EnumType>, queuedEnumNames:Map<String, Bool>,
			projectEnums:Array<EnumType>):Void {
		for (requiredName in pendingRequiredEnumsByName.keys()) {
			if (queuedEnumNames.exists(requiredName)) {
				continue;
			}
			var requiredEnum = pendingRequiredEnumsByName.get(requiredName);
			if (requiredEnum == null) {
				continue;
			}
			queuedEnumNames.set(requiredName, true);
			enumQueue.push(requiredEnum);
			projectEnums.push(requiredEnum);
		}
		clearEnumMap(pendingRequiredEnumsByName);
		while (enumQueue.length > 0) {
			var enumType = enumQueue.shift();
			appendModuleDecls(moduleDecls, enumType.module, lowerEnumDecls(enumType));
			for (requiredName in pendingRequiredEnumsByName.keys()) {
				if (queuedEnumNames.exists(requiredName)) {
					continue;
				}
				var requiredEnum = pendingRequiredEnumsByName.get(requiredName);
				if (requiredEnum == null) {
					continue;
				}
				queuedEnumNames.set(requiredName, true);
				enumQueue.push(requiredEnum);
				projectEnums.push(requiredEnum);
			}
			clearEnumMap(pendingRequiredEnumsByName);
		}
	}

	function lowerTypeValueDecls():Array<GoDecl> {
		return [
			GoDecl.GoStructDecl("hxrt__TypeClassValue", [{name: "name", typeName: "*string"}]),
			GoDecl.GoStructDecl("hxrt__TypeEnumValue", [{name: "name", typeName: "*string"}])
		];
	}

	function lowerStdlibShimDecls():Array<GoDecl> {
		var decls = new Array<GoDecl>();
		if (requiredStdlibShimGroups.exists("stdlib_symbols")) {
			decls = decls.concat(lowerStdlibSymbolShimDecls());
		}
		decls = decls.concat(lowerSerializationSourceBridgeShimDecls());
		if (requiredStdlibShimGroups.exists("go_concurrency")) {
			decls = decls.concat(lowerGoConcurrencyShimDecls());
		}
		if (requiredStdlibShimGroups.exists("go_collections")) {
			decls = decls.concat(lowerTypedGoCollectionShimDecls());
		}
		if (requiredStdlibShimGroups.exists("go_result")) {
			decls = decls.concat(lowerGoResultShimDecls());
		}
		return decls;
	}

	function lowerGoConcurrencyShimDecls():Array<GoDecl> {
		var spawnBody = shouldScopeDetachedGoroutineIdentity() ? [
			GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("hxrt.ThreadSpawnDetached"), [GoExpr.GoIdent("fn")]))
		] : [GoStmt.GoGoStmt(GoExpr.GoCall(GoExpr.GoIdent("fn"), []))];
		var decls = [
			GoDecl.GoFuncDecl("go__concurrency_makeChan", null, [{name: "buffer", typeName: "int"}], ["any"], [
				GoStmt.GoIf(GoExpr.GoBinary(">", GoExpr.GoIdent("buffer"), GoExpr.GoIntLiteral(0)), [GoStmt.GoReturn(GoExpr.GoRaw("make(chan any, buffer)"))],
					null),
				GoStmt.GoReturn(GoExpr.GoRaw("make(chan any)"))
			]),
			GoDecl.GoFuncDecl("go__concurrency_send", null, [
				{
					name: "channel",
					typeName: "any"
				},
				{name: "value", typeName: "any"}
			], [], [
				GoStmt.GoRaw("chanValue := reflect.ValueOf(channel)"),
				GoStmt.GoRaw("if !chanValue.IsValid() || chanValue.Kind() != reflect.Chan {"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("sendValue := reflect.ValueOf(value)"),
				GoStmt.GoRaw("if !sendValue.IsValid() {"),
				GoStmt.GoRaw("\tsendValue = reflect.Zero(chanValue.Type().Elem())"),
				GoStmt.GoRaw("} else if !sendValue.Type().AssignableTo(chanValue.Type().Elem()) {"),
				GoStmt.GoRaw("\tif sendValue.Type().ConvertibleTo(chanValue.Type().Elem()) {"),
				GoStmt.GoRaw("\t\tsendValue = sendValue.Convert(chanValue.Type().Elem())"),
				GoStmt.GoRaw("\t} else {"),
				GoStmt.GoRaw("\t\treturn"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("chanValue.Send(sendValue)")
			]),
			GoDecl.GoFuncDecl("go__concurrency_trySend", null, [
				{
					name: "channel",
					typeName: "any"
				},
				{name: "value", typeName: "any"}
			], ["bool"], [
				GoStmt.GoRaw("chanValue := reflect.ValueOf(channel)"),
				GoStmt.GoRaw("if !chanValue.IsValid() || chanValue.Kind() != reflect.Chan {"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("sendValue := reflect.ValueOf(value)"),
				GoStmt.GoRaw("if !sendValue.IsValid() {"),
				GoStmt.GoRaw("\tsendValue = reflect.Zero(chanValue.Type().Elem())"),
				GoStmt.GoRaw("} else if !sendValue.Type().AssignableTo(chanValue.Type().Elem()) {"),
				GoStmt.GoRaw("\tif sendValue.Type().ConvertibleTo(chanValue.Type().Elem()) {"),
				GoStmt.GoRaw("\t\tsendValue = sendValue.Convert(chanValue.Type().Elem())"),
				GoStmt.GoRaw("\t} else {"),
				GoStmt.GoRaw("\t\treturn false"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("cases := []reflect.SelectCase{"),
				GoStmt.GoRaw("\t{Dir: reflect.SelectSend, Chan: chanValue, Send: sendValue},"),
				GoStmt.GoRaw("\t{Dir: reflect.SelectDefault},"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("chosen, _, _ := reflect.Select(cases)"),
				GoStmt.GoRaw("return chosen == 0")
			]),
			GoDecl.GoFuncDecl("go__concurrency_recv", null, [
				{
					name: "channel",
					typeName: "any"
				}
			], ["any"], [
				GoStmt.GoRaw("chanValue := reflect.ValueOf(channel)"),
				GoStmt.GoRaw("if !chanValue.IsValid() || chanValue.Kind() != reflect.Chan {"),
				GoStmt.GoRaw("\treturn nil"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("recvValue, _ := chanValue.Recv()"),
				GoStmt.GoRaw("if !recvValue.IsValid() {"),
				GoStmt.GoRaw("\treturn nil"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("return recvValue.Interface()")
			]),
			GoDecl.GoFuncDecl("go__concurrency_recvOr", null, [
				{
					name: "channel",
					typeName: "any"
				},
				{name: "defaultValue", typeName: "any"}
			], ["any"], [
				GoStmt.GoRaw("chanValue := reflect.ValueOf(channel)"),
				GoStmt.GoRaw("if !chanValue.IsValid() || chanValue.Kind() != reflect.Chan {"),
				GoStmt.GoRaw("\treturn defaultValue"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("cases := []reflect.SelectCase{"),
				GoStmt.GoRaw("\t{Dir: reflect.SelectRecv, Chan: chanValue},"),
				GoStmt.GoRaw("\t{Dir: reflect.SelectDefault},"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("chosen, recvValue, received := reflect.Select(cases)"),
				GoStmt.GoRaw("if chosen == 0 {"),
				GoStmt.GoRaw("\tif !received {"),
				GoStmt.GoRaw("\t\treturn defaultValue"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn recvValue.Interface()"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("return defaultValue")
			]),
			GoDecl.GoFuncDecl("go__concurrency_tryRecv", null, [
				{
					name: "channel",
					typeName: "any"
				}
			], ["*go___Result"], [
				GoStmt.GoRaw("chanValue := reflect.ValueOf(channel)"),
				GoStmt.GoRaw("if !chanValue.IsValid() || chanValue.Kind() != reflect.Chan {"),
				GoStmt.GoRaw("\treturn New_go___Result(nil, New_go___Error(hxrt.StringFromLiteral(\"empty\")))"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("cases := []reflect.SelectCase{"),
				GoStmt.GoRaw("\t{Dir: reflect.SelectRecv, Chan: chanValue},"),
				GoStmt.GoRaw("\t{Dir: reflect.SelectDefault},"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("chosen, recvValue, received := reflect.Select(cases)"),
				GoStmt.GoRaw("if chosen == 0 {"),
				GoStmt.GoRaw("\tif !received {"),
				GoStmt.GoRaw("\t\treturn New_go___Result(nil, New_go___Error(hxrt.StringFromLiteral(\"closed\")))"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn New_go___Result(recvValue.Interface(), nil)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("return New_go___Result(nil, New_go___Error(hxrt.StringFromLiteral(\"empty\")))")
			]),
			GoDecl.GoFuncDecl("go__concurrency_close", null, [
				{
					name: "channel",
					typeName: "any"
				}
			], [], [
				GoStmt.GoRaw("chanValue := reflect.ValueOf(channel)"),
				GoStmt.GoRaw("if !chanValue.IsValid() || chanValue.Kind() != reflect.Chan {"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("chanValue.Close()")
			]),
			GoDecl.GoFuncDecl("go__concurrency_spawn", null, [
				{
					name: "fn",
					typeName: "func()"
				}
			], [], spawnBody)
		];
		if (useTypedGoConcurrencySpecialization()) {
			decls = decls.concat(lowerTypedGoConcurrencyShimDecls());
		}
		return decls;
	}

	/**
		Why
		`go.Go.spawn` must remain a bare native goroutine unless the program can
		actually request portable thread identity. When `sys.thread` is present,
		callbacks can call `Thread.current()` or `Tls`, so the compiler-owned spawn
		boundary must release that lazily-created state on return or panic.

		What
		Detects whether the finalized project class graph contains a portable
		`sys.thread` surface.

		How
		Uses the same class-path condition as HXRT thread feature inference. The
		generated wrapper still does not join or recover the goroutine.
	**/
	function shouldScopeDetachedGoroutineIdentity():Bool {
		for (classType in projectClasses) {
			if (StringTools.startsWith(fullClassName(classType), "sys.thread.")) {
				return true;
			}
		}
		return false;
	}

	function lowerTypedGoConcurrencyShimDecls():Array<GoDecl> {
		var elementTypes = [for (elementType in requiredNativeChanElementTypes.keys()) elementType];
		if (elementTypes.length == 0) {
			return [];
		}

		elementTypes.sort(function(a, b) return Reflect.compare(a, b));

		var chanTypeName = GoNaming.typeSymbol(["go"], "Chan");
		var chanCtorName = GoNaming.constructorSymbol(["go"], "Chan");
		var chanPointerType = "*" + chanTypeName;
		var decls = new Array<GoDecl>();

		for (elementType in elementTypes) {
			var chanType = "chan " + elementType;
			var makeName = nativeChanShimName("go__concurrency_makeChan", elementType);
			var setBufferName = nativeChanShimName("go__concurrency_setBuffer", elementType);
			var newChanName = nativeChanShimName("go__concurrency_newChan", elementType);
			var sendName = nativeChanShimName("go__concurrency_send", elementType);
			var trySendName = nativeChanShimName("go__concurrency_trySend", elementType);
			var recvName = nativeChanShimName("go__concurrency_recv", elementType);
			var recvOrName = nativeChanShimName("go__concurrency_recvOr", elementType);
			var tryRecvName = nativeChanShimName("go__concurrency_tryRecv", elementType);
			var closeName = nativeChanShimName("go__concurrency_close", elementType);

			decls.push(GoDecl.GoFuncDecl(makeName, null, [{name: "buffer", typeName: "int"}], ["any"], [
				GoStmt.GoIf(GoExpr.GoBinary(">", GoExpr.GoIdent("buffer"), GoExpr.GoIntLiteral(0)),
					[GoStmt.GoReturn(GoExpr.GoRaw("make(" + chanType + ", buffer)"))], null),
				GoStmt.GoReturn(GoExpr.GoRaw("make(" + chanType + ")"))
			]));

			decls.push(GoDecl.GoFuncDecl(setBufferName, null, [{name: "channel", typeName: chanPointerType}, {name: "buffer", typeName: "int"}], [], [
				GoStmt.GoIf(GoExpr.GoBinary("==", GoExpr.GoIdent("channel"), GoExpr.GoNil), [GoStmt.GoReturn(null)], null),
				GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("channel"), "__hx_native"),
					GoExpr.GoCall(GoExpr.GoIdent(makeName), [GoExpr.GoIdent("buffer")]))
			]));

			decls.push(GoDecl.GoFuncDecl(newChanName, null, [{name: "buffer", typeName: "int"}], [chanPointerType], [
				GoStmt.GoVarDecl("channel", null, GoExpr.GoCall(GoExpr.GoIdent(chanCtorName), []), true),
				GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent(setBufferName), [GoExpr.GoIdent("channel"), GoExpr.GoIdent("buffer")])),
				GoStmt.GoReturn(GoExpr.GoIdent("channel"))
			]));

			decls.push(GoDecl.GoFuncDecl(sendName, null, [{name: "channel", typeName: "any"}, {name: "value", typeName: elementType}], [], [
				GoStmt.GoSendStmt(GoExpr.GoTypeAssert(GoExpr.GoIdent("channel"), chanType), GoExpr.GoIdent("value"))
			]));

			decls.push(GoDecl.GoFuncDecl(trySendName, null, [{name: "channel", typeName: "any"}, {name: "value", typeName: elementType}], ["bool"], [
				GoStmt.GoSelect([
					{
						clause: GoSelectClause.GoSelectSend(GoExpr.GoTypeAssert(GoExpr.GoIdent("channel"), chanType), GoExpr.GoIdent("value")),
						body: [GoStmt.GoReturn(GoExpr.GoBoolLiteral(true))]
					},
					{
						clause: GoSelectClause.GoSelectDefault,
						body: [GoStmt.GoReturn(GoExpr.GoBoolLiteral(false))]
					}
				])
			]));

			decls.push(GoDecl.GoFuncDecl(recvName, null, [{name: "channel", typeName: "any"}], [elementType], [
				GoStmt.GoReturn(GoExpr.GoRecvExpr(GoExpr.GoTypeAssert(GoExpr.GoIdent("channel"), chanType)))
			]));

			decls.push(GoDecl.GoFuncDecl(recvOrName, null, [
				{name: "channel", typeName: "any"},
				{name: "defaultValue", typeName: elementType}
			], [elementType], [
				GoStmt.GoSelect([
					{
						clause: GoSelectClause.GoSelectRecvAssignOk(GoExpr.GoIdent("value"), GoExpr.GoIdent("received"),
							GoExpr.GoRecvExpr(GoExpr.GoTypeAssert(GoExpr.GoIdent("channel"), chanType)), true),
						body: [
							GoStmt.GoIf(GoExpr.GoUnary("!", GoExpr.GoIdent("received")), [GoStmt.GoReturn(GoExpr.GoIdent("defaultValue"))], null),
							GoStmt.GoReturn(GoExpr.GoIdent("value"))
						]
					},
					{
						clause: GoSelectClause.GoSelectDefault,
						body: [GoStmt.GoReturn(GoExpr.GoIdent("defaultValue"))]
					}
				])
			]));

			decls.push(GoDecl.GoFuncDecl(tryRecvName, null, [{name: "channel", typeName: "any"}], ["*go___Result"], [
				GoStmt.GoSelect([
					{
						clause: GoSelectClause.GoSelectRecvAssignOk(GoExpr.GoIdent("value"), GoExpr.GoIdent("received"),
							GoExpr.GoRecvExpr(GoExpr.GoTypeAssert(GoExpr.GoIdent("channel"), chanType)), true),
						body: [
							GoStmt.GoIf(GoExpr.GoUnary("!", GoExpr.GoIdent("received")), [
								GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("New_go___Result"), [
									GoExpr.GoNil,
									GoExpr.GoCall(GoExpr.GoIdent("New_go___Error"), [
										GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoStringLiteral("closed")])
									])
								]))
							], null),
							GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("New_go___Result"), [GoExpr.GoIdent("value"), GoExpr.GoNil]))
						]
					},
					{
						clause: GoSelectClause.GoSelectDefault,
						body: [
							GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("New_go___Result"), [
								GoExpr.GoNil,
								GoExpr.GoCall(GoExpr.GoIdent("New_go___Error"), [
									GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoStringLiteral("empty")])
								])
							]))
						]
					}
				])
			]));

			decls.push(GoDecl.GoFuncDecl(closeName, null, [{name: "channel", typeName: "any"}], [], [
				GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("close"), [GoExpr.GoTypeAssert(GoExpr.GoIdent("channel"), chanType)]))
			]));
		}
		return decls;
	}

	function lowerTypedGoCollectionShimDecls():Array<GoDecl> {
		if (!useTypedGoCollectionsSpecialization()) {
			return [];
		}

		var decls = new Array<GoDecl>();
		var sliceElementTypes = [for (elementType in requiredNativeSliceElementTypes.keys()) elementType];
		sliceElementTypes.sort(function(a, b) return Reflect.compare(a, b));
		var sliceTypeName = GoNaming.typeSymbol(["go"], "Slice");
		var slicePointerType = "*" + sliceTypeName;
		for (elementType in sliceElementTypes) {
			var pushName = nativeSliceShimName("go__slice_push", elementType);
			var setName = nativeSliceShimName("go__slice_set", elementType);
			var getName = nativeSliceShimName("go__slice_get", elementType);
			var lengthName = nativeSliceShimName("go__slice_length", elementType);
			var toArrayName = nativeSliceShimName("go__slice_toArray", elementType);

			decls.push(GoDecl.GoFuncDecl(pushName, null, [
				{
					name: "slice",
					typeName: slicePointerType
				},
				{name: "value", typeName: elementType}
			], [], [
				GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("slice"), "data"),
					GoExpr.GoCall(GoExpr.GoIdent("append"), [GoExpr.GoSelector(GoExpr.GoIdent("slice"), "data"), GoExpr.GoIdent("value")]))
			]));

			decls.push(GoDecl.GoFuncDecl(setName, null, [
				{
					name: "slice",
					typeName: slicePointerType
				},
				{name: "index", typeName: "int"},
				{name: "value", typeName: elementType}
			], [], [
				GoStmt.GoAssign(GoExpr.GoIndex(GoExpr.GoSelector(GoExpr.GoIdent("slice"), "data"), GoExpr.GoIdent("index")), GoExpr.GoIdent("value"))
			]));

			decls.push(GoDecl.GoFuncDecl(getName, null, [
				{
					name: "slice",
					typeName: slicePointerType
				},
				{name: "index", typeName: "int"}
			], [elementType], [
				GoStmt.GoVarDecl("raw", "any", GoExpr.GoIndex(GoExpr.GoSelector(GoExpr.GoIdent("slice"), "data"), GoExpr.GoIdent("index")), true),
				GoStmt.GoIf(GoExpr.GoBinary("==", GoExpr.GoIdent("raw"), GoExpr.GoNil), [
					GoStmt.GoVarDecl("zero", elementType, null, false),
					GoStmt.GoReturn(GoExpr.GoIdent("zero"))
				], null),
				GoStmt.GoReturn(GoExpr.GoTypeAssert(GoExpr.GoIdent("raw"), elementType))
			]));

			decls.push(GoDecl.GoFuncDecl(lengthName, null, [
				{
					name: "slice",
					typeName: slicePointerType
				}
			], ["int"], [
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("len"), [GoExpr.GoSelector(GoExpr.GoIdent("slice"), "data")]))
			]));

			decls.push(GoDecl.GoFuncDecl(toArrayName, null, [
				{
					name: "slice",
					typeName: slicePointerType
				}
			], ["[]" + elementType], [
				GoStmt.GoVarDecl("raw", "[]any", GoExpr.GoSelector(GoExpr.GoIdent("slice"), "data"), true),
				GoStmt.GoVarDecl("out", "[]" + elementType,
					GoExpr.GoMakeSlice(elementType, GoExpr.GoCall(GoExpr.GoIdent("len"), [GoExpr.GoIdent("raw")]), null), true),
				GoStmt.GoRangeStmt("idx", "value", GoExpr.GoIdent("raw"), true, [
					GoStmt.GoIf(GoExpr.GoBinary("==", GoExpr.GoIdent("value"), GoExpr.GoNil), [GoStmt.GoContinue], null),
					GoStmt.GoAssign(GoExpr.GoIndex(GoExpr.GoIdent("out"), GoExpr.GoIdent("idx")), GoExpr.GoTypeAssert(GoExpr.GoIdent("value"), elementType))
				]),
				GoStmt.GoReturn(GoExpr.GoIdent("out"))
			]));
		}

		var mapSignatures = [for (signature in requiredNativeMapTypePairs.keys()) signature];
		mapSignatures.sort(function(a, b) return Reflect.compare(a, b));
		var mapTypeName = GoNaming.typeSymbol(["go"], "Map");
		var mapPointerType = "*" + mapTypeName;
		for (signature in mapSignatures) {
			var pair = requiredNativeMapTypePairs.get(signature);
			if (pair == null) {
				continue;
			}
			var keyType = pair.keyGoType;
			var valueType = pair.valueGoType;
			var setName = nativeMapShimName("go__map_set", keyType, valueType);
			var getName = nativeMapShimName("go__map_get", keyType, valueType);
			var existsName = nativeMapShimName("go__map_exists", keyType, valueType);

			decls.push(GoDecl.GoFuncDecl(setName, null, [
				{
					name: "mapValue",
					typeName: mapPointerType
				},
				{name: "key", typeName: keyType},
				{name: "value", typeName: valueType}
			], [], [
				GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoSelector(GoExpr.GoIdent("mapValue"), "inner"), "set"), [
					GoExpr.GoCall(GoExpr.GoIdent("hxrt.StdString"), [GoExpr.GoCall(GoExpr.GoIdent("any"), [GoExpr.GoIdent("key")])]),
					GoExpr.GoIdent("value")
				]))
			]));

			decls.push(GoDecl.GoFuncDecl(getName, null, [
				{
					name: "mapValue",
					typeName: mapPointerType
				},
				{name: "key", typeName: keyType}
			], ["any"], [
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoSelector(GoExpr.GoIdent("mapValue"), "inner"), "get"), [
					GoExpr.GoCall(GoExpr.GoIdent("hxrt.StdString"), [GoExpr.GoCall(GoExpr.GoIdent("any"), [GoExpr.GoIdent("key")])])
				]))
			]));

			decls.push(GoDecl.GoFuncDecl(existsName, null, [
				{
					name: "mapValue",
					typeName: mapPointerType
				},
				{name: "key", typeName: keyType}
			], ["bool"], [
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoSelector(GoExpr.GoIdent("mapValue"), "inner"), "exists"), [
					GoExpr.GoCall(GoExpr.GoIdent("hxrt.StdString"), [GoExpr.GoCall(GoExpr.GoIdent("any"), [GoExpr.GoIdent("key")])])
				]))
			]));
		}

		return decls;
	}

	function lowerGoResultShimDecls():Array<GoDecl> {
		var resultTypeName = GoNaming.typeSymbol(["go"], "Result");
		var resultPointerType = "*" + resultTypeName;
		var decls = [
			GoDecl.GoFuncDecl("go__result_fromValueError", null, [{name: "value", typeName: "any"}, {name: "err", typeName: "error"}], [resultPointerType], [
				GoStmt.GoIf(GoExpr.GoBinary("!=", GoExpr.GoIdent("err"), GoExpr.GoNil), [
					GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("New_go___Result"), [
						GoExpr.GoNil,
						GoExpr.GoCall(GoExpr.GoIdent("New_go___Error"), [
							GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("err"), "Error"), [])])
						])
					]))
				], null),
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("New_go___Result"), [GoExpr.GoIdent("value"), GoExpr.GoNil]))
			])
		];
		return decls.concat(lowerTypedGoResultShimDecls());
	}

	function lowerTypedGoResultShimDecls():Array<GoDecl> {
		if (!useTypedGoResultSpecialization()) {
			return [];
		}

		var elementTypes = [for (elementType in requiredNativeResultElementTypes.keys()) elementType];
		if (elementTypes.length == 0) {
			return [];
		}

		elementTypes.sort(function(a, b) return Reflect.compare(a, b));
		var resultTypeName = GoNaming.typeSymbol(["go"], "Result");
		var resultPointerType = "*" + resultTypeName;
		var decls = new Array<GoDecl>();

		for (elementType in elementTypes) {
			var okName = nativeResultShimName("go__result_ok", elementType);
			var failureName = nativeResultShimName("go__result_failure", elementType);
			var valueErrorName = nativeResultShimName("go__result_valueError", elementType);
			var isOkName = nativeResultShimName("go__result_isOk", elementType);
			var isErrName = nativeResultShimName("go__result_isErr", elementType);
			var unwrapName = nativeResultShimName("go__result_unwrap", elementType);
			var errorName = nativeResultShimName("go__result_error", elementType);

			decls.push(GoDecl.GoFuncDecl(okName, null, [{name: "value", typeName: elementType}], [resultPointerType], [
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("New_go___Result"), [GoExpr.GoIdent("value"), GoExpr.GoNil]))
			]));

			decls.push(GoDecl.GoFuncDecl(failureName, null, [{name: "message", typeName: "*string"}], [resultPointerType], [
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("New_go___Result"), [
					GoExpr.GoNil,
					GoExpr.GoCall(GoExpr.GoIdent("New_go___Error"), [GoExpr.GoIdent("message")])
				]))
			]));

			decls.push(GoDecl.GoFuncDecl(valueErrorName, null, [{name: "result", typeName: resultPointerType}], [elementType, "error"], [
				GoStmt.GoRaw("var zero " + elementType),
				GoStmt.GoRaw("if result == nil {"),
				GoStmt.GoRaw("return zero, errors.New(\"nil go.Result\")"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if result.errorValue != nil {"),
				GoStmt.GoRaw("return zero, errors.New(*hxrt.StdString(result.errorValue.message))"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if result.value == nil {"),
				GoStmt.GoRaw("return zero, nil"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("return result.value.(" + elementType + "), nil")
			]));

			decls.push(GoDecl.GoFuncDecl(isOkName, null, [{name: "result", typeName: resultPointerType}], ["bool"], [
				GoStmt.GoRaw("_, err := " + valueErrorName + "(result)"),
				GoStmt.GoReturn(GoExpr.GoBinary("==", GoExpr.GoIdent("err"), GoExpr.GoNil))
			]));

			decls.push(GoDecl.GoFuncDecl(isErrName, null, [{name: "result", typeName: resultPointerType}], ["bool"], [
				GoStmt.GoRaw("_, err := " + valueErrorName + "(result)"),
				GoStmt.GoReturn(GoExpr.GoBinary("!=", GoExpr.GoIdent("err"), GoExpr.GoNil))
			]));

			decls.push(GoDecl.GoFuncDecl(unwrapName, null, [{name: "result", typeName: resultPointerType}], [elementType], [
				GoStmt.GoRaw("value, err := " + valueErrorName + "(result)"),
				GoStmt.GoRaw("if err != nil {"),
				GoStmt.GoRaw("hxrt.Throw(hxrt.StringFromLiteral(err.Error()))"),
				GoStmt.GoRaw("var zero " + elementType),
				GoStmt.GoRaw("return zero"),
				GoStmt.GoRaw("}"),
				GoStmt.GoReturn(GoExpr.GoIdent("value"))
			]));

			decls.push(GoDecl.GoFuncDecl(errorName, null, [{name: "result", typeName: resultPointerType}], ["*string"], [
				GoStmt.GoRaw("_, err := " + valueErrorName + "(result)"),
				GoStmt.GoIf(GoExpr.GoBinary("==", GoExpr.GoIdent("err"), GoExpr.GoNil), [GoStmt.GoReturn(GoExpr.GoNil)], null),
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("err"), "Error"), [])]))
			]));
		}

		return decls;
	}

	function lowerStdlibSymbolShimDecls():Array<GoDecl> {
		var decls = [
			GoDecl.GoStructDecl("Std", []),
			GoDecl.GoStructDecl("Type", []),
			GoDecl.GoStructDecl("Reflect", []),
			GoDecl.GoFuncDecl("Reflect_compare", null, [
				{
					name: "a",
					typeName: "any"
				},
				{name: "b", typeName: "any"}
			], ["int"], [
				GoStmt.GoRaw("toFloat := func(value any) (float64, bool) {"),
				GoStmt.GoRaw("\tswitch v := value.(type) {"),
				GoStmt.GoRaw("\tcase int:"),
				GoStmt.GoRaw("\t\treturn float64(v), true"),
				GoStmt.GoRaw("\tcase int8:"),
				GoStmt.GoRaw("\t\treturn float64(v), true"),
				GoStmt.GoRaw("\tcase int16:"),
				GoStmt.GoRaw("\t\treturn float64(v), true"),
				GoStmt.GoRaw("\tcase int32:"),
				GoStmt.GoRaw("\t\treturn float64(v), true"),
				GoStmt.GoRaw("\tcase int64:"),
				GoStmt.GoRaw("\t\treturn float64(v), true"),
				GoStmt.GoRaw("\tcase uint:"),
				GoStmt.GoRaw("\t\treturn float64(v), true"),
				GoStmt.GoRaw("\tcase uint8:"),
				GoStmt.GoRaw("\t\treturn float64(v), true"),
				GoStmt.GoRaw("\tcase uint16:"),
				GoStmt.GoRaw("\t\treturn float64(v), true"),
				GoStmt.GoRaw("\tcase uint32:"),
				GoStmt.GoRaw("\t\treturn float64(v), true"),
				GoStmt.GoRaw("\tcase uint64:"),
				GoStmt.GoRaw("\t\treturn float64(v), true"),
				GoStmt.GoRaw("\tcase float32:"),
				GoStmt.GoRaw("\t\treturn float64(v), true"),
				GoStmt.GoRaw("\tcase float64:"),
				GoStmt.GoRaw("\t\treturn v, true"),
				GoStmt.GoRaw("\tdefault:"),
				GoStmt.GoRaw("\t\treturn 0, false"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if af, ok := toFloat(a); ok {"),
				GoStmt.GoRaw("\tif bf, okB := toFloat(b); okB {"),
				GoStmt.GoRaw("\t\tif af < bf {"),
				GoStmt.GoRaw("\t\t\treturn -1"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t\tif af > bf {"),
				GoStmt.GoRaw("\t\t\treturn 1"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t\treturn 0"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("aStr := *hxrt.StdString(a)"),
				GoStmt.GoRaw("bStr := *hxrt.StdString(b)"),
				GoStmt.GoRaw("if aStr < bStr {"),
				GoStmt.GoRaw("\treturn -1"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if aStr > bStr {"),
				GoStmt.GoRaw("\treturn 1"),
				GoStmt.GoRaw("}"),
				GoStmt.GoReturn(GoExpr.GoIntLiteral(0))
			]),
			GoDecl.GoFuncDecl("Reflect_compareMethods", null, [
				{
					name: "a",
					typeName: "any"
				},
				{name: "b", typeName: "any"}
			], ["bool"], [
				GoStmt.GoRaw("if a == nil || b == nil {"),
				GoStmt.GoRaw("\treturn a == nil && b == nil"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("av := reflect.ValueOf(a)"),
				GoStmt.GoRaw("bv := reflect.ValueOf(b)"),
				GoStmt.GoRaw("if !av.IsValid() || !bv.IsValid() {"),
				GoStmt.GoRaw("\treturn !av.IsValid() && !bv.IsValid()"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if av.Kind() == reflect.Func && bv.Kind() == reflect.Func {"),
				GoStmt.GoRaw("\tif av.IsNil() || bv.IsNil() {"),
				GoStmt.GoRaw("\t\treturn av.IsNil() && bv.IsNil()"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn av.Pointer() == bv.Pointer()"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("return reflect.DeepEqual(a, b)")
			]),
			GoDecl.GoFuncDecl("Reflect_field", null, [
				{
					name: "obj",
					typeName: "any"
				},
				{name: "field", typeName: "*string"}
			], ["any"], [
				GoStmt.GoRaw("if obj == nil {"),
				GoStmt.GoRaw("\treturn nil"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("key := *hxrt.StdString(field)"),
				GoStmt.GoRaw("if metadataValue, ok := hxrt_typeClassMetadataField(obj, key); ok {"),
				GoStmt.GoRaw("\treturn metadataValue"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("switch value := obj.(type) {"),
				GoStmt.GoRaw("case map[string]any:"),
				GoStmt.GoRaw("\treturn value[key]"),
				GoStmt.GoRaw("case map[any]any:"),
				GoStmt.GoRaw("\treturn value[key]"),
				GoStmt.GoRaw("case *map[string]any:"),
				GoStmt.GoRaw("\tif value == nil {"),
				GoStmt.GoRaw("\t\treturn nil"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn (*value)[key]"),
				GoStmt.GoRaw("case *map[any]any:"),
				GoStmt.GoRaw("\tif value == nil {"),
				GoStmt.GoRaw("\t\treturn nil"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn (*value)[key]"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("rv := reflect.ValueOf(obj)"),
				GoStmt.GoRaw("if !rv.IsValid() {"),
				GoStmt.GoRaw("\treturn nil"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if rv.Kind() == reflect.Pointer {"),
				GoStmt.GoRaw("\tif rv.IsNil() {"),
				GoStmt.GoRaw("\t\treturn nil"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\trv = rv.Elem()"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if rv.Kind() == reflect.Struct {"),
				GoStmt.GoRaw("\tif fieldValue := rv.FieldByName(key); fieldValue.IsValid() && fieldValue.CanInterface() {"),
				GoStmt.GoRaw("\t\treturn fieldValue.Interface()"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("method := reflect.ValueOf(obj).MethodByName(key)"),
				GoStmt.GoRaw("if method.IsValid() {"),
				GoStmt.GoRaw("\treturn method.Interface()"),
				GoStmt.GoRaw("}"),
				GoStmt.GoReturn(GoExpr.GoNil)
			]),
			GoDecl.GoFuncDecl("Reflect_hasField", null, [
				{
					name: "obj",
					typeName: "any"
				},
				{name: "field", typeName: "*string"}
			], ["bool"], [
				GoStmt.GoRaw("if obj == nil {"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("key := *hxrt.StdString(field)"),
				GoStmt.GoRaw("if _, ok := hxrt_typeClassMetadataField(obj, key); ok {"),
				GoStmt.GoRaw("\treturn true"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("switch value := obj.(type) {"),
				GoStmt.GoRaw("case map[string]any:"),
				GoStmt.GoRaw("\t_, ok := value[key]"),
				GoStmt.GoRaw("\treturn ok"),
				GoStmt.GoRaw("case map[any]any:"),
				GoStmt.GoRaw("\t_, ok := value[key]"),
				GoStmt.GoRaw("\treturn ok"),
				GoStmt.GoRaw("case *map[string]any:"),
				GoStmt.GoRaw("\tif value == nil {"),
				GoStmt.GoRaw("\t\treturn false"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\t_, ok := (*value)[key]"),
				GoStmt.GoRaw("\treturn ok"),
				GoStmt.GoRaw("case *map[any]any:"),
				GoStmt.GoRaw("\tif value == nil {"),
				GoStmt.GoRaw("\t\treturn false"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\t_, ok := (*value)[key]"),
				GoStmt.GoRaw("\treturn ok"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("rv := reflect.ValueOf(obj)"),
				GoStmt.GoRaw("if !rv.IsValid() {"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if rv.Kind() == reflect.Pointer {"),
				GoStmt.GoRaw("\tif rv.IsNil() {"),
				GoStmt.GoRaw("\t\treturn false"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\trv = rv.Elem()"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if rv.Kind() == reflect.Struct {"),
				GoStmt.GoRaw("\tif rv.FieldByName(key).IsValid() {"),
				GoStmt.GoRaw("\t\treturn true"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("return reflect.ValueOf(obj).MethodByName(key).IsValid()")
			]),
			GoDecl.GoFuncDecl("Reflect_setField", null, [
				{
					name: "obj",
					typeName: "any"
				},
				{name: "field", typeName: "*string"},
				{name: "value", typeName: "any"}
			], [], [
				GoStmt.GoRaw("if obj == nil {"),
				GoStmt.GoRaw("\thxrt.Throw(hxrt.StringFromLiteral(\"Null Access\"))"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("key := *hxrt.StdString(field)"),
				GoStmt.GoRaw("switch target := obj.(type) {"),
				GoStmt.GoRaw("case map[string]any:"),
				GoStmt.GoRaw("\ttarget[key] = value"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("case map[any]any:"),
				GoStmt.GoRaw("\ttarget[key] = value"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("case *map[string]any:"),
				GoStmt.GoRaw("\tif target == nil {"),
				GoStmt.GoRaw("\t\thxrt.Throw(hxrt.StringFromLiteral(\"Null Access\"))"),
				GoStmt.GoRaw("\t\treturn"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\t(*target)[key] = value"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("case *map[any]any:"),
				GoStmt.GoRaw("\tif target == nil {"),
				GoStmt.GoRaw("\t\thxrt.Throw(hxrt.StringFromLiteral(\"Null Access\"))"),
				GoStmt.GoRaw("\t\treturn"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\t(*target)[key] = value"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("rv := reflect.ValueOf(obj)"),
				GoStmt.GoRaw("if !rv.IsValid() || rv.Kind() != reflect.Pointer {"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if rv.IsNil() {"),
				GoStmt.GoRaw("\thxrt.Throw(hxrt.StringFromLiteral(\"Null Access\"))"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("rv = rv.Elem()"),
				GoStmt.GoRaw("if rv.Kind() != reflect.Struct {"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("fieldValue := rv.FieldByName(key)"),
				GoStmt.GoRaw("if !fieldValue.IsValid() || !fieldValue.CanSet() {"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if value == nil {"),
				GoStmt.GoRaw("\tfieldValue.Set(reflect.Zero(fieldValue.Type()))"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("incoming := reflect.ValueOf(value)"),
				GoStmt.GoRaw("if incoming.Type().AssignableTo(fieldValue.Type()) {"),
				GoStmt.GoRaw("\tfieldValue.Set(incoming)"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if incoming.Type().ConvertibleTo(fieldValue.Type()) {"),
				GoStmt.GoRaw("\tfieldValue.Set(incoming.Convert(fieldValue.Type()))"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if fieldValue.Kind() == reflect.Interface {"),
				GoStmt.GoRaw("\tfieldValue.Set(incoming)"),
				GoStmt.GoRaw("}")
			]),
			GoDecl.GoStructDecl("haxe__ds__Option",
				[
					{
						name: "tag",
						typeName: "int"
					},
					{name: "params", typeName: "[]any"}
				]),
			GoDecl.GoGlobalVarDecl("haxe__ds__Option_None", "*haxe__ds__Option", GoExpr.GoRaw("&haxe__ds__Option{tag: 1, params: []any{}}")),
			GoDecl.GoFuncDecl("haxe__ds__Option_Some", null, [
				{
					name: "value",
					typeName: "any"
				}
			], ["*haxe__ds__Option"],
				[GoStmt.GoReturn(GoExpr.GoRaw("&haxe__ds__Option{tag: 0, params: []any{value}}"))]),
		];
		if (requiresReflectFieldsShim) {
			decls.push(reflectFieldsShimDecl());
		}
		decls = decls.concat(lowerTypeReflectionShimDecls());
		return decls;
	}

	function lowerSerializationSourceBridgeShimDecls():Array<GoDecl> {
		return GoSerializationSourceBridgeEmitter.emit(requiredSourceOwnedClassNames.exists("haxe.Serializer"),
			requiredSourceOwnedClassNames.exists("haxe.Unserializer"));
	}

	function reflectFieldsShimDecl():GoDecl {
		return GoDecl.GoFuncDecl("Reflect_fields", null, [
			{
				name: "obj",
				typeName: "any"
			}
		], ["*hxrt.Array"], [
			GoStmt.GoRaw("if obj == nil {"),
			GoStmt.GoRaw("\treturn hxrt.NewArray()"),
			GoStmt.GoRaw("}"),
			GoStmt.GoRaw("switch value := obj.(type) {"),
			GoStmt.GoRaw("case map[string]any:"),
			GoStmt.GoRaw("\tkeys := hxrt.NewArray()"),
			GoStmt.GoRaw("\tfor key := range value {"),
			GoStmt.GoRaw("\t\tkeys.Push(hxrt.StringFromLiteral(key))"),
			GoStmt.GoRaw("\t}"),
			GoStmt.GoRaw("\treturn keys"),
			GoStmt.GoRaw("case map[any]any:"),
			GoStmt.GoRaw("\tkeys := hxrt.NewArray()"),
			GoStmt.GoRaw("\tfor key := range value {"),
			GoStmt.GoRaw("\t\tkeys.Push(hxrt.StdString(key))"),
			GoStmt.GoRaw("\t}"),
			GoStmt.GoRaw("\treturn keys"),
			GoStmt.GoRaw("case *map[string]any:"),
			GoStmt.GoRaw("\tif value == nil {"),
			GoStmt.GoRaw("\t\treturn hxrt.NewArray()"),
			GoStmt.GoRaw("\t}"),
			GoStmt.GoRaw("\tkeys := hxrt.NewArray()"),
			GoStmt.GoRaw("\tfor key := range *value {"),
			GoStmt.GoRaw("\t\tkeys.Push(hxrt.StringFromLiteral(key))"),
			GoStmt.GoRaw("\t}"),
			GoStmt.GoRaw("\treturn keys"),
			GoStmt.GoRaw("case *map[any]any:"),
			GoStmt.GoRaw("\tif value == nil {"),
			GoStmt.GoRaw("\t\treturn hxrt.NewArray()"),
			GoStmt.GoRaw("\t}"),
			GoStmt.GoRaw("\tkeys := hxrt.NewArray()"),
			GoStmt.GoRaw("\tfor key := range *value {"),
			GoStmt.GoRaw("\t\tkeys.Push(hxrt.StdString(key))"),
			GoStmt.GoRaw("\t}"),
			GoStmt.GoRaw("\treturn keys"),
			GoStmt.GoRaw("}"),
			GoStmt.GoRaw("rv := reflect.ValueOf(obj)"),
			GoStmt.GoRaw("if !rv.IsValid() {"),
			GoStmt.GoRaw("\treturn hxrt.NewArray()"),
			GoStmt.GoRaw("}"),
			GoStmt.GoRaw("if rv.Kind() == reflect.Pointer {"),
			GoStmt.GoRaw("\tif rv.IsNil() {"),
			GoStmt.GoRaw("\t\treturn hxrt.NewArray()"),
			GoStmt.GoRaw("\t}"),
			GoStmt.GoRaw("\trv = rv.Elem()"),
			GoStmt.GoRaw("}"),
			GoStmt.GoRaw("if rv.Kind() != reflect.Struct {"),
			GoStmt.GoRaw("\treturn hxrt.NewArray()"),
			GoStmt.GoRaw("}"),
			GoStmt.GoRaw("rt := rv.Type()"),
			GoStmt.GoRaw("keys := hxrt.NewArray()"),
			GoStmt.GoRaw("for i := 0; i < rv.NumField(); i++ {"),
			GoStmt.GoRaw("\tfield := rt.Field(i)"),
			GoStmt.GoRaw("\tif field.PkgPath != \"\" {"),
			GoStmt.GoRaw("\t\tcontinue"),
			GoStmt.GoRaw("\t}"),
			GoStmt.GoRaw("\tkeys.Push(hxrt.StringFromLiteral(field.Name))"),
			GoStmt.GoRaw("}"),
			GoStmt.GoRaw("return keys")
		]);
	}

	function lowerTypeReflectionShimDecls():Array<GoDecl> {
		return GoTypeReflectionEmitter.emit(typeReflectionClassMetadata(), typeReflectionEnumMetadata(), goRawQuotedString, goStringArrayCarrierLiteral)
			.concat(GoRttiMetadataEmitter.emit(rttiClassMetadata(), goRawQuotedString));
	}

	function lowerClassDecls(classType:ClassType):Array<GoDecl> {
		if (GoStdlibOwnership.isCompilerOwnedAuthority(fullClassName(classType))) {
			return [];
		}
		if (classType.isInterface) {
			return lowerInterfaceDecls(classType);
		}

		var decls = new Array<GoDecl>();
		var typeName = classTypeName(classType);
		var superClass = projectSuperClass(classType);
		var directHaxeExceptionSuper = directHaxeExceptionSuperClass(classType);

		var instanceDataFields = new Array<GoParam>();
		var instanceMethods = new Array<{name:String, func:TFunc, fieldType:Type}>();
		for (field in classType.fields.get()) {
			switch (field.kind) {
				case FVar(_, _):
					instanceDataFields.push({
						name: normalizeIdent(field.name),
						typeName: scalarGoType(field.type)
					});
				case FMethod(MethDynamic):
					// Haxe dynamic methods are mutable per-instance function values, not
					// virtual methods in the generated Go interface.
					instanceDataFields.push({
						name: normalizeIdent(field.name),
						typeName: scalarGoType(field.type)
					});
				case FMethod(_):
					if (field.name != "new") {
						var methodFunc = unwrapFunction(field.expr());
						if (methodFunc != null) {
							instanceMethods.push({name: field.name, func: methodFunc, fieldType: field.type});
						}
					}
			}
		}
		if (directHaxeExceptionSuper && !hasStructField(instanceDataFields, "__hx_exception")) {
			instanceDataFields.push({
				name: "__hx_exception",
				typeName: "*hxrt.ExceptionValue"
			});
		}

		var ctorRef = classType.constructor;
		var ctorFunc:Null<TFunc> = null;
		if (ctorRef != null) {
			ctorFunc = unwrapFunction(ctorRef.get().expr());
		}

		var hasInstanceLayout = superClass != null || instanceDataFields.length > 0 || instanceMethods.length > 0 || ctorFunc != null;
		var dispatchMethods = hasInstanceLayout ? collectDispatchMethods(classType) : [];
		if (hasInstanceLayout) {
			var instanceFields = new Array<GoParam>();
			if (superClass != null) {
				instanceFields.push({
					name: "",
					typeName: "*" + classTypeName(superClass)
				});
			}
			instanceFields.push({
				name: "__hx_this",
				typeName: interfaceSymbol(classType)
			});
			instanceFields = instanceFields.concat(instanceDataFields);

			var interfaceMethods = new Array<GoInterfaceMethod>();
			for (method in dispatchMethods) {
				interfaceMethods.push({
					name: method.name,
					params: lowerFunctionParams(method.func, typedFunctionArgs(method.fieldType)),
					results: lowerFunctionResults(method.func.t)
				});
			}
			decls.push(GoDecl.GoInterfaceDecl(interfaceSymbol(classType), interfaceMethods));
			decls.push(GoDecl.GoStructDecl(typeName, instanceFields));
			decls.push(lowerConstructorDecl(classType, ctorFunc, ctorRef == null ? null : ctorRef.get().type, superClass));
			if (directHaxeExceptionSuper) {
				decls.push(lowerHaxeExceptionCarrierDecl(classType));
			}
		} else if (requiredNominalClassTypeNames.exists(fullClassName(classType))) {
			// What: retain the nominal Go type for a static-only Haxe class.
			// Why: Haxe still permits that class in type positions (for example
			// `var probe:Sys = null`) even though it has no instance layout to lower.
			// How: emit only the empty carrier; no constructor or instance interface is
			// invented for a class whose staged source declares only static members.
			decls.push(GoDecl.GoStructDecl(typeName, []));
		}

		for (method in instanceMethods) {
			decls.push(lowerInstanceMethodDecl(classType, method.name, method.func, method.fieldType));
		}
		if (hasPortableToString(dispatchMethods)) {
			decls.push(lowerGoStringerAdapterDecl(classType));
		}
		var staticFields = classType.statics.get().copy();
		staticFields.sort(function(a, b) return Reflect.compare(a.name, b.name));
		for (field in staticFields) {
			var symbol = staticSymbol(classType, field.name);
			switch (field.kind) {
				case FVar(_, _):
					if (field.name == "__init__") {
						continue;
					}
					var loweredValue = if (fullClassName(classType) == "haxe.Resource" && field.name == "content") {
						haxeResourceContentLiteral();
					} else {
						var valueExpr = field.expr();
						valueExpr == null ? null : materializeExprWithPrefix(lowerExprWithExpectedUpcast(valueExpr, field.type), field.type).expr;
					}
					decls.push(GoDecl.GoGlobalVarDecl(symbol, scalarGoType(field.type), loweredValue));
				case FMethod(_):
					var func = unwrapFunction(field.expr());
					if (func != null) {
						decls.push(lowerFunctionDecl(symbol, func, null, classType.module, field.type));
					}
			}
		}

		return decls;
	}

	function lowerInterfaceDecls(classType:ClassType):Array<GoDecl> {
		var methods = new Array<GoInterfaceMethod>();
		var seen = new Map<String, Bool>();
		for (field in classType.fields.get()) {
			switch (field.kind) {
				case FMethod(_):
					if (field.name == "new") {
						continue;
					}
					var method = lowerInterfaceMethod(classType, field);
					if (method != null && !seen.exists(method.name)) {
						seen.set(method.name, true);
						methods.push(method);
					}
				case _:
			}
		}
		return [GoDecl.GoInterfaceDecl(classTypeName(classType), methods)];
	}

	function lowerInterfaceMethod(classType:ClassType, field:ClassField):Null<GoInterfaceMethod> {
		var methodName = interfaceFieldName(classType, field);
		var followed = Context.follow(field.type);
		return switch (followed) {
			case TFun(args, returnType):
				{
					name: methodName,
					params: lowerTypedFunArgs(args),
					results: lowerFunctionResults(returnType)
				};
			case _:
				var methodFunc = unwrapFunction(field.expr());
				if (methodFunc == null) {
					null;
				} else {
					{
						name: methodName,
						params: lowerFunctionParams(methodFunc, typedFunctionArgs(field.type)),
						results: lowerFunctionResults(methodFunc.t)
					};
				}
		};
	}

	function lowerTypedFunArgs(args:Array<{name:String, opt:Bool, t:Type}>):Array<GoParam> {
		var out = new Array<GoParam>();
		var used = new Map<String, Int>();
		for (index in 0...args.length) {
			var arg = args[index];
			var rawName = arg.name;
			if (rawName == null || rawName == "") {
				rawName = "arg" + index;
			}
			var baseName = normalizeIdent(rawName);
			var count = used.exists(baseName) ? used.get(baseName) : 0;
			used.set(baseName, count + 1);
			var finalName = count == 0 ? baseName : baseName + "_" + count;
			out.push({
				name: finalName,
				typeName: nilCapablePrimitiveParamType(arg.t)
			});
		}
		return out;
	}

	function lowerFunctionDecl(name:String, func:TFunc, receiver:Null<GoParam>, ?sourceModule:String, ?functionType:Type):GoDecl {
		pushFunctionVarNameScope();
		var params = lowerFunctionParams(func, typedFunctionArgs(functionType == null ? func.t : functionType));
		var results = lowerFunctionResults(func.t);
		pushFunctionReturnType(func.t);
		var body = lowerFunctionBody(func.expr);
		prependLineDirective(body, func.expr.pos, sourceModule);
		popFunctionReturnType();
		popFunctionVarNameScope();
		return GoDecl.GoFuncDecl(name, receiver, params, results, body);
	}

	/**
		What: Rebind every generated superclass dispatch carrier to the receiver
		being constructed.

		Why: Each generated ancestor owns a separate `__hx_this` field. Rebinding
		only the direct superclass leaves deeper carriers pointing at an intermediate
		object, so a leaf override is lost after a deep nominal upcast or from inside
		an inherited base-method closure.

		How: Walk the `projectSuperClass` embedding path recursively; that path ends
		before native/extern ancestors, and this helper additionally stops at
		compiler-owned authority. Emit typed assignments deepest-first after the
		direct superclass constructor returns and before the subclass body.
	**/
	function lowerAncestorDispatchCarrierRebindings(superClass:ClassType, carrier:GoExpr):Array<GoStmt> {
		if (GoStdlibOwnership.isCompilerOwnedAuthority(fullClassName(superClass))) {
			return [];
		}

		var out = new Array<GoStmt>();
		var ancestor = projectSuperClass(superClass);
		if (ancestor != null) {
			var ancestorCarrier = GoExpr.GoSelector(carrier, classTypeName(ancestor));
			out = out.concat(lowerAncestorDispatchCarrierRebindings(ancestor, ancestorCarrier));
		}
		out.push(GoStmt.GoAssign(GoExpr.GoSelector(carrier, "__hx_this"), GoExpr.GoIdent("self")));
		return out;
	}

	function lowerConstructorDecl(classType:ClassType, ctorFunc:Null<TFunc>, ctorType:Null<Type>, superClass:Null<ClassType>):GoDecl {
		pushFunctionVarNameScope();
		var typeName = classTypeName(classType);
		var params = ctorFunc == null ? [] : lowerFunctionParams(ctorFunc, typedFunctionArgs(ctorType == null ? ctorFunc.t : ctorType));
		var body = new Array<GoStmt>();
		body.push(GoStmt.GoVarDecl("self", null, GoExpr.GoUnary(GoUnaryOperator.AddressOf, GoExpr.GoCompositeLiteral(GoType.named(typeName), [])), true));

		var loweredCtorBody:ConstructorBodyLowering = {
			superArgs: null,
			body: []
		};
		if (ctorFunc != null) {
			pushConstructorReturnScope();
			loweredCtorBody = lowerConstructorBody(ctorFunc.expr);
			popConstructorReturnScope();
		}

		if (superClass != null) {
			var superTypeName = classTypeName(superClass);
			var superCarrier = GoExpr.GoSelector(GoExpr.GoIdent("self"), superTypeName);
			var superCtorArgs = loweredCtorBody.superArgs == null ? [] : [for (arg in loweredCtorBody.superArgs) lowerExpr(arg).expr];
			body.push(GoStmt.GoAssign(superCarrier, GoExpr.GoCall(GoExpr.GoIdent(constructorSymbol(superClass)), superCtorArgs)));
			body = body.concat(lowerAncestorDispatchCarrierRebindings(superClass, superCarrier));
		} else if (directHaxeExceptionSuperClass(classType)) {
			var exceptionCtorArgs = loweredCtorBody.superArgs == null ? [] : [for (arg in loweredCtorBody.superArgs) lowerExpr(arg).expr];
			while (exceptionCtorArgs.length < 3) {
				exceptionCtorArgs.push(GoExpr.GoNil);
			}
			body.push(GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("self"), "__hx_exception"),
				GoExpr.GoCall(GoExpr.GoIdent("hxrt.BindException"), [GoExpr.GoIdent("self")].concat(exceptionCtorArgs))));
		}
		body.push(GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("self"), "__hx_this"), GoExpr.GoIdent("self")));
		body = body.concat(lowerDynamicMethodInitializers(classType));
		if (ctorFunc != null) {
			prependLineDirective(loweredCtorBody.body, ctorFunc.expr.pos, classType.module);
		}
		body = body.concat(loweredCtorBody.body);
		body.push(GoStmt.GoReturn(GoExpr.GoIdent("self")));
		popFunctionVarNameScope();
		return GoDecl.GoFuncDecl(constructorSymbol(classType), null, params, ["*" + typeName], body);
	}

	/**
		What: Initialize each Haxe `dynamic function` as a function-valued instance field.

		Why: Dynamic methods can be reassigned independently on every object (the callback
		surface on `haxe.http.HttpBase` is a canonical example). A Go interface method is
		shared by the type and its selector cannot appear on the left side of assignment.

		How: Keep the typed source function expression, lower it with the declared field
		type, and assign it after the concrete receiver is wired but before the user
		constructor body runs. Inherited dynamic fields are initialized by the superclass
		constructor and remain ordinary promoted Go fields.
	**/
	function lowerDynamicMethodInitializers(classType:ClassType):Array<GoStmt> {
		var out = new Array<GoStmt>();
		for (field in classType.fields.get()) {
			switch (field.kind) {
				case FMethod(MethDynamic):
					var value = field.expr();
					if (value != null) {
						var lowered = lowerExprWithExpectedUpcast(value, field.type);
						out = out.concat(lowered.prefix);
						out.push(GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("self"), normalizeIdent(field.name)), lowered.expr));
					}
				case _:
			}
		}
		return out;
	}

	function lowerInstanceMethodDecl(classType:ClassType, fieldName:String, func:TFunc, fieldType:Type):GoDecl {
		return lowerFunctionDecl(normalizeIdent(fieldName), func, {
			name: "self",
			typeName: "*" + classTypeName(classType)
		}, classType.module, fieldType);
	}

	/**
		What: Detect the ordinary Haxe `toString():String` object protocol.

		Why: Generated Haxe methods intentionally use their source spelling, so
		`toString` is unexported to the separate `hxrt` Go package. Without a small
		generated adapter, `Std.string` falls back to Go's struct dump instead of the
		portable Haxe result.

		How: Recognize the method only from typed dispatch information. Inherited and
		overridden implementations therefore follow the same `__hx_this` virtual
		dispatch path as an ordinary Haxe call.
	**/
	function hasPortableToString(methods:Array<{name:String, func:TFunc, fieldType:Type}>):Bool {
		for (method in methods) {
			if (method.name == "toString" && method.func.args.length == 0 && isStringType(method.func.t)) {
				return true;
			}
		}
		return false;
	}

	/**
		What: Expose Haxe `toString()` through Go's standard `fmt.Stringer` shape.

		Why: `hxrt.StdString` receives erased values and cannot call unexported Haxe
		methods across a Go package boundary. The adapter keeps object policy in the
		source method while making that policy visible to the runtime formatter.

		How: Emit `String() string` as a typed AST declaration and delegate through
		`__hx_this.toString()` so subclass overrides remain authoritative.
	**/
	function lowerGoStringerAdapterDecl(classType:ClassType):GoDecl {
		var receiver = GoExpr.GoIdent("self");
		var virtualReceiver = GoExpr.GoSelector(receiver, "__hx_this");
		var portableString = GoExpr.GoCall(GoExpr.GoSelector(virtualReceiver, "toString"), []);
		return GoDecl.GoFuncDecl("String", {
			name: "self",
			typeName: "*" + classTypeName(classType)
		}, [], ["string"], [GoStmt.GoReturn(GoExpr.GoUnary("*", portableString))]);
	}

	function lowerHaxeExceptionCarrierDecl(classType:ClassType):GoDecl {
		return GoDecl.GoFuncDecl("HxExceptionValue", {
			name: "self",
			typeName: "*" + classTypeName(classType)
		}, [], ["*hxrt.ExceptionValue"],
			[GoStmt.GoReturn(GoExpr.GoSelector(GoExpr.GoIdent("self"), "__hx_exception"))]);
	}

	function prependLineDirective(body:Array<GoStmt>, pos:haxe.macro.Expr.Position, sourceModule:Null<String>):Void {
		var directive = lineDirectiveStmt(pos, sourceModule);
		if (directive == null) {
			return;
		}
		body.unshift(directive);
	}

	function lineDirectiveStmt(pos:haxe.macro.Expr.Position, sourceModule:Null<String>):Null<GoStmt> {
		if (!compilationContext.emitLineDirectives || sourceModule == null || sourceModule == "") {
			return null;
		}
		var line = 1;
		var location = PositionTools.toLocation(pos);
		if (location != null && location.range != null && location.range.start != null && location.range.start.line > 0) {
			line = location.range.start.line;
		}
		return GoStmt.GoRaw("//line " + sourceModuleToFilePath(sourceModule) + ":" + line);
	}

	function sourceModuleToFilePath(moduleName:String):String {
		return StringTools.replace(moduleName, ".", "/") + ".hx";
	}

	function hasStructField(fields:Array<GoParam>, fieldName:String):Bool {
		for (field in fields) {
			if (field.name == fieldName) {
				return true;
			}
		}
		return false;
	}

	function lowerConstructorBody(expr:TypedExpr):ConstructorBodyLowering {
		pushLocalScope();
		var bodyExprs:Array<TypedExpr> = switch (expr.expr) {
			case TBlock(exprs): exprs;
			case _:
				[expr];
		};

		var startIndex = 0;
		var superArgs:Null<Array<TypedExpr>> = null;
		if (bodyExprs.length > 0) {
			var extracted = extractSuperCtorArgs(bodyExprs[0]);
			if (extracted != null) {
				superArgs = extracted;
				startIndex = 1;
			}
		}

		var out = new Array<GoStmt>();
		for (index in startIndex...bodyExprs.length) {
			out = out.concat(lowerToStatements(bodyExprs[index]));
		}
		popLocalScope();

		return {
			superArgs: superArgs,
			body: out
		};
	}

	function extractSuperCtorArgs(expr:TypedExpr):Null<Array<TypedExpr>> {
		return switch (expr.expr) {
			case TCall(callee, args):
				isSuperCtorCall(callee) ? args : null;
			case TMeta(_, inner):
				extractSuperCtorArgs(inner);
			case TParenthesis(inner):
				extractSuperCtorArgs(inner);
			case TCast(inner, _):
				extractSuperCtorArgs(inner);
			case _:
				null;
		};
	}

	function collectDispatchMethods(classType:ClassType):Array<{name:String, func:TFunc, fieldType:Type}> {
		var orderedNames = new Array<String>();
		var methods = new Map<String, {func:TFunc, fieldType:Type}>();

		function collect(current:ClassType):Void {
			var superClass = projectSuperClass(current);
			if (superClass != null) {
				collect(superClass);
			}

			for (field in current.fields.get()) {
				switch (field.kind) {
					case FMethod(MethDynamic):
						// Mutable function fields do not participate in virtual dispatch.
						null;
					case FMethod(_):
						if (field.name == "new") {
							continue;
						}
						var methodFunc = unwrapFunction(field.expr());
						if (methodFunc == null) {
							continue;
						}
						var methodName = normalizeIdent(field.name);
						if (!methods.exists(methodName)) {
							orderedNames.push(methodName);
						}
						methods.set(methodName, {func: methodFunc, fieldType: field.type});
					case _:
				}
			}
		}

		collect(classType);

		var out = new Array<{name:String, func:TFunc, fieldType:Type}>();
		for (name in orderedNames) {
			var method = methods.get(name);
			out.push({
				name: name,
				func: method.func,
				fieldType: method.fieldType
			});
		}
		return out;
	}

	function unwrapFunction(expr:Null<TypedExpr>):Null<TFunc> {
		if (expr == null) {
			return null;
		}

		return switch (expr.expr) {
			case TFunction(func):
				func;
			case TMeta(_, inner):
				unwrapFunction(inner);
			case TParenthesis(inner):
				unwrapFunction(inner);
			case TCast(inner, _):
				unwrapFunction(inner);
			case _:
				null;
		};
	}

	function typedFunctionArgs(functionType:Type):Array<{name:String, opt:Bool, t:Type}> {
		return switch (Context.follow(functionType)) {
			case TFun(args, _):
				args;
			case _:
				[];
		};
	}

	function lowerFunctionParams(func:TFunc, ?typedArgs:Array<{name:String, opt:Bool, t:Type}>):Array<GoParam> {
		if (typedArgs == null) {
			typedArgs = typedFunctionArgs(func.t);
		}
		var params = new Array<GoParam>();
		for (index in 0...func.args.length) {
			var arg = func.args[index];
			var typedArg = index < typedArgs.length ? typedArgs[index] : null;
			var isOptionalPrimitive = isOptionalPrimitiveFunctionArg(arg, typedArg);
			var isNilCapablePrimitive = isOptionalPrimitive || isNullablePrimitiveParamType(arg.v.t, typedArg);
			registerOptionalPrimitiveParam(arg.v, isNilCapablePrimitive);
			params.push({
				name: localVarName(arg.v),
				typeName: isNilCapablePrimitive ? "any" : scalarGoType(arg.v.t)
			});
		}
		return params;
	}

	function buildFunctionInfo(func:TFunc):FunctionInfo {
		return {
			defaults: [for (arg in func.args) arg.value]
		};
	}

	function resolveConstructorInfo(classType:ClassType):Null<FunctionInfo> {
		if (classType.constructor == null) {
			return null;
		}
		var ctorExpr = classType.constructor.get().expr();
		var ctorFunc = unwrapFunction(ctorExpr);
		return ctorFunc == null ? null : buildFunctionInfo(ctorFunc);
	}

	/**
		What: Resolves one constructor parameter as emitted by the shared Go class.
		Why: Haxe has already validated concrete generic arguments, but a generated
		generic constructor keeps erased method signatures such as `next():any` inside
		the structural carrier. Adapting to the applied source type would violate that
		hidden Go ABI even though both outer values are `map[string]any`.
		How: Read the requested argument from the constructor's declared function type
		before class-generic substitution.
	**/
	function emittedConstructorParamType(classType:ClassType, index:Int):Null<Type> {
		if (classType.constructor == null) {
			return null;
		}
		return callParamType(classType.constructor.get().type, index);
	}

	/**
		What: Lowers one explicit or default constructor argument for its parameter type.
		Why: Constructors previously bypassed structural and prefix-aware expected-type
		coercion even though ordinary calls used it.
		How: Reuse `lowerExprWithExpectedUpcast` and materialize any ordered prefix into
		one expression suitable for the generated Go constructor call.
	**/
	function lowerConstructorArg(classType:ClassType, arg:TypedExpr, index:Int):GoExpr {
		var paramType = emittedConstructorParamType(classType, index);
		if (paramType == null) {
			return lowerExpr(arg).expr;
		}
		return materializeExprWithPrefix(lowerExprWithExpectedUpcast(arg, paramType), paramType).expr;
	}

	function lowerFunctionResults(returnType:Type):Array<GoType> {
		if (isVoidType(returnType)) {
			return [];
		}
		// A public `Null<Int|Float|Bool>` result must retain the nil state in Go.
		// `Context.follow()` erases the Haxe `Null` wrapper inside scalarGoType(), so
		// select nil-capable storage before applying the ordinary scalar mapping.
		return [
			GoType.parse(isNullablePrimitiveType(returnType) ? "any" : scalarGoType(returnType))
		];
	}

	function lowerToStatements(expr:TypedExpr):Array<GoStmt> {
		return switch (expr.expr) {
			case TBlock(exprs):
				lowerBlock(exprs);
			case TMeta(_, inner):
				lowerToStatements(inner);
			case TParenthesis(inner):
				lowerToStatements(inner);
			case TCast(inner, _):
				lowerToStatements(inner);
			case TVar(variable, value):
				var variableName = localVarName(variable);
				var restIteratorCtorArg = restIteratorCtorArg(value);
				if (restIteratorCtorArg != null) {
					registerRestIterator(variableName);
					return lowerRestIteratorCtor(variableName, restIteratorCtorArg);
				}

				var functionValue = unwrapFunction(value);
				if (functionValue != null) {
					registerLocalFunction(variableName, functionValue);
				}
				var lambdaAlias = value == null ? null : lambdaIterableLowering.functionAliasName(value);
				if (lambdaAlias != null) {
					registerLocalLambdaAlias(variableName, lambdaAlias);
				}

				var lowered = value == null ? null : lowerExprWithExpectedUpcast(value, variable.t);
				var prefix = lowered == null ? [] : lowered.prefix;
				var loweredValue = lowered == null ? null : lowered.expr;
				if (value != null && loweredValue != null) {
					var valueKnownNonNullPrimitive = nonNullPrimitiveExprGoType(value) != null;
					loweredValue = coerceAnyExprToType(loweredValue, value.t, variable.t, !isSharedArrayElementExpr(value) && ((exprBackedByAny(value)
						&& !valueKnownNonNullPrimitive)
						|| shouldForceAnyCoerce(value.t, variable.t)));
				}
				var storageOverride = localArrayStorageOverrides.exists(variable.id) ? localArrayStorageOverrides.get(variable.id) : null;
				var goType = storageOverride == null ? valueStorageGoType(variable.t) : typeToGoType(storageOverride);
				var narrowedStorageGoType = value == null ? null : nonNullPrimitiveExprGoType(value);
				// Keep storage narrowing local to immutable-after-declaration vars; reassigned
				// nullable primitives still need nil-capable storage for later writes.
				if (narrowedStorageGoType != null && isNullablePrimitiveType(variable.t) && localNeverReassigned(variable)) {
					goType = narrowedStorageGoType;
					registerNarrowedPrimitiveStorage(variable, narrowedStorageGoType);
				}
				var useShort = loweredValue != null && !isNilExpr(loweredValue) && goType != "any" && !isInterfaceType(variable.t);
				var decl = GoStmt.GoVarDecl(variableName, goType, loweredValue, useShort);
				var consume = GoStmt.GoAssign(GoExpr.GoIdent("_"), GoExpr.GoIdent(variableName));

				if (prefix.length > 0) {
					prefix.push(decl);
					prefix.push(consume);
					prefix;
				} else {
					[decl, consume];
				}
			case TBinop(op, left, right):
				switch (op) {
					case OpAssign:
						var indexedArrayAssign = lowerHaxeArrayIndexAssignStatements(left, right);
						if (indexedArrayAssign != null) {
							indexedArrayAssign;
						} else {
							var loweredRight = lowerExprWithExpectedUpcast(right, assignmentStorageType(left));
							var lengthAssignStmts = lowerArrayLengthAssign(left, loweredRight.expr);
							var assignStmts = if (lengthAssignStmts != null) {
								lengthAssignStmts;
							} else {
								[GoStmt.GoAssign(lowerLValue(left), loweredRight.expr)];
							};
							if (loweredRight.prefix.length > 0) {
								loweredRight.prefix.concat(assignStmts);
							} else {
								assignStmts;
							}
						}
					case OpAssignOp(assignOp):
						var sharedArrayAssign = lowerHaxeArrayIndexAssignOpExpr(left, right, assignOp, expr.pos);
						if (sharedArrayAssign != null) {
							exprStatement(sharedArrayAssign.expr);
						} else {
							var loweredRight = lowerExprWithPrefix(right);
							var stringAppendFromSharedArray = assignOp == OpAdd
								&& (isStringType(left.t) || isStringType(right.t))
								&& isSharedArrayElementExpr(right);
							var rightExpr = stringAppendFromSharedArray ? lowerSharedArrayElementStorageExpr(right) : upcastIfNeeded(loweredRight.expr,
								right.t, left.t, right);
							if (!stringAppendFromSharedArray && isSharedArrayElementExpr(right)) {
								rightExpr = coerceStoredArrayElementExpr(rightExpr, left.t);
							}
							var targetExpr = lowerLValue(left);
							var assignExpr = lowerAssignOpExpr(assignOp, targetExpr, rightExpr, left.t, right.t, expr.pos, stringAppendFromSharedArray);
							var assignStmt = GoStmt.GoAssign(targetExpr, assignExpr);
							if (loweredRight.prefix.length > 0) {
								loweredRight.prefix.concat([assignStmt]);
							} else {
								[assignStmt];
							}
						}
					case _:
						exprStatement(lowerExpr(expr).expr);
				}
			case TIf(condition, thenBranch, elseBranch):
				var facts = conditionNonNullFacts(condition);
				var thenBody = lowerWithNonNullPrimitiveFacts(facts.thenFacts, function() return lowerToStatements(thenBranch));
				var elseBody = elseBranch == null ? null : lowerWithNonNullPrimitiveFacts(facts.elseFacts, function() return lowerToStatements(elseBranch));
					[GoStmt.GoIf(lowerExpr(condition).expr, thenBody, elseBody)];
			case TWhile(condition, body, normalWhile):
				if (normalWhile) {
					[lowerLoopStmt(lowerExpr(condition).expr, body)];
				} else {
					var firstPassVar = freshTempName("hx_do_first");
					var loweredCondition = lowerExpr(condition).expr;
					var loopCondition = GoExpr.GoBinary("||", GoExpr.GoIdent(firstPassVar), loweredCondition);
					var target:LoopBreakTarget = {label: null};
					loopBreakTargetScopes.push(target);
					var loopBody = [GoStmt.GoAssign(GoExpr.GoIdent(firstPassVar), GoExpr.GoBoolLiteral(false))].concat(lowerToStatements(body));
					loopBreakTargetScopes.pop();
					var loopStmt:GoStmt = GoStmt.GoWhile(loopCondition, loopBody);
					if (target.label != null) {
						loopStmt = GoStmt.GoLabeled(target.label, loopStmt);
					}
					[GoStmt.GoVarDecl(firstPassVar, null, GoExpr.GoBoolLiteral(true), true), loopStmt];
				}
			case TBreak:
				[GoStmt.GoBreak(switchDepth > 0 ? currentLoopBreakLabel() : null)];
			case TContinue:
				[GoStmt.GoContinue];
			case TUnop(op, postFix, value):
				switch (op) {
					case OpIncrement:
						var sharedArrayUnit = lowerHaxeArrayIndexUnitExpr(value, op, postFix, expr.pos);
						if (sharedArrayUnit != null) {
							exprStatement(sharedArrayUnit.expr);
						} else {
							var target = lowerLValue(value);
							[
								GoStmt.GoAssign(target, unitStepExpr(target, GoBinaryOperator.Add, value.t, expr.pos))
							];
						}
					case OpDecrement:
						var sharedArrayUnit = lowerHaxeArrayIndexUnitExpr(value, op, postFix, expr.pos);
						if (sharedArrayUnit != null) {
							exprStatement(sharedArrayUnit.expr);
						} else {
							var target = lowerLValue(value);
							[
								GoStmt.GoAssign(target, unitStepExpr(target, GoBinaryOperator.Subtract, value.t, expr.pos))
							];
						}
					case _:
						exprStatement(lowerExpr(expr).expr);
				}
			case TSwitch(value, cases, defaultExpr):
				[lowerSwitchStmt(value, cases, defaultExpr)];
			case TThrow(value):
				var loweredValue = lowerExprWithPrefix(value);
				loweredValue.prefix.concat([
					GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("hxrt.Throw"), [loweredValue.expr]))
				]).concat(nonVoidThrowFallbackReturnStmts());
			case TTry(tryExpr, catches):
				lowerTryCatchStmt(tryExpr, catches);
			case TReturn(value):
				var redirect = currentReturnRedirect();
				if (redirect != null) {
					var redirected = new Array<GoStmt>();
					if (value != null) {
						var loweredReturn = redirect.valueName != null
							&& redirect.valueType != null ? lowerExprWithExpectedUpcast(value, redirect.valueType) : lowerExprWithPrefix(value);
						var returnExpr = loweredReturn.expr;
						redirected = redirected.concat(loweredReturn.prefix);
						if (redirect.valueName != null && redirect.valueType != null) {
							redirected.push(GoStmt.GoAssign(GoExpr.GoIdent(redirect.valueName), returnExpr));
						} else {
							redirected.push(GoStmt.GoExprStmt(returnExpr));
						}
					}
					redirected.push(GoStmt.GoAssign(GoExpr.GoIdent(redirect.flagName), GoExpr.GoBoolLiteral(true)));
					redirected.push(GoStmt.GoReturn(null));
					redirected;
				} else if (value == null) {
					[
						GoStmt.GoReturn(inConstructorReturnScope() && currentFunctionReturnType() == null ? GoExpr.GoIdent("self") : null)
					];
				} else {
					var expectedReturnType = currentFunctionReturnType();
					var loweredReturn = expectedReturnType == null ? lowerExprWithPrefix(value) : lowerExprWithExpectedUpcast(value, expectedReturnType);
					var returnExpr = loweredReturn.expr;
					var returnStmt = GoStmt.GoReturn(returnExpr);
					if (loweredReturn.prefix.length > 0) {
						loweredReturn.prefix.concat([returnStmt]);
					} else {
						[returnStmt];
					}
				}
			case TCall(callee, args):
				if (isSuperCtorCall(callee)) {
					[];
				} else {
					var arrayCall = asArrayMethodCall(callee);
					if (arrayCall != null && arrayCall.methodName == "push") {
						var site = lowerArrayMutationSite(arrayCall.target);
						var shouldMaskToByte = isBytesBufferStorageArray(arrayCall.target);
						var pushArgs = usesSharedArrayCarrier(arrayCall.target) ? [] : [site.tempExpr];
						var body = site.prefix.copy();
						for (arg in args) {
							var loweredArg = lowerExprWithPrefix(arg);
							body = body.concat(loweredArg.prefix);
							var appendValue = loweredArg.expr;
							if (shouldMaskToByte) {
								appendValue = GoExpr.GoBinary("&", appendValue, GoExpr.GoIntLiteral(255));
							}
							pushArgs.push(appendValue);
						}
						if (usesSharedArrayCarrier(arrayCall.target)) {
							body.concat([
								GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoSelector(site.tempExpr, "Push"), pushArgs))
							]);
						} else {
							body.concat([
								GoStmt.GoAssign(site.tempExpr, GoExpr.GoCall(GoExpr.GoIdent("append"), pushArgs))
							]).concat(site.writeBack(site.tempExpr));
						}
					} else if (arrayCall != null && arrayCall.methodName == "pop") {
						var site = lowerArrayMutationSite(arrayCall.target);
						if (usesSharedArrayCarrier(arrayCall.target)) {
							site.prefix.concat([GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoSelector(site.tempExpr, "Pop"), []))]);
						} else {
							var lenExpr = GoExpr.GoCall(GoExpr.GoIdent("len"), [site.tempExpr]);
							site.prefix.concat([
								GoStmt.GoIf(GoExpr.GoBinary(">", lenExpr, GoExpr.GoIntLiteral(0)), [
									GoStmt.GoAssign(site.tempExpr, GoExpr.GoSlice(site.tempExpr, null, GoExpr.GoBinary("-", lenExpr, GoExpr.GoIntLiteral(1))))
								], null)
							]).concat(site.writeBack(site.tempExpr));
						}
					} else {
						exprStatement(lowerCall(callee, args, expr.t).expr);
					}
				}
			case _:
				exprStatement(lowerExpr(expr).expr);
		};
	}

	function exprStatement(expr:GoExpr):Array<GoStmt> {
		return switch (expr) {
			case GoExpr.GoCall(_, _):
				[GoStmt.GoExprStmt(expr)];
			case _:
				if (isNilExpr(expr)) {
					[];
				} else {
					[GoStmt.GoAssign(GoExpr.GoIdent("_"), expr)];
				}
		};
	}

	function fieldAccessName(access:FieldAccess):Null<String> {
		return switch (access) {
			case FInstance(_, _, field):
				field.get().name;
			case FAnon(field):
				field.get().name;
			case FDynamic(name):
				name;
			case _:
				null;
		};
	}

	function isBytesBufferStorageArray(target:TypedExpr):Bool {
		return switch (target.expr) {
			case TField(receiver, access):
				var fieldName = fieldAccessName(access);
				if (fieldName != "b") {
					false;
				} else {
					switch (Context.follow(receiver.t)) {
						case TInst(classRef, _): var classType = classRef.get(); classType.pack.join(".") == "haxe.io" && classType.name == "BytesBuffer";
						case _:
							false;
					}
				}
			case TMeta(_, inner):
				isBytesBufferStorageArray(inner);
			case TParenthesis(inner):
				isBytesBufferStorageArray(inner);
			case TCast(inner, _):
				isBytesBufferStorageArray(inner);
			case _:
				false;
		};
	}

	/**
		What: Lowers one indexed write through the shared portable Array carrier.
		Why: Direct Go indexing cannot grow sparse Haxe arrays and cannot update one
		shared header across aliases.
		How: Evaluate receiver, index, and value once in source order, then call the
		carrier's typed boundary method that owns null-filled growth.
	**/
	function lowerHaxeArrayIndexAssignStatements(left:TypedExpr, right:TypedExpr):Null<Array<GoStmt>> {
		return switch (left.expr) {
			case TArray(target, index) if (usesSharedArrayCarrier(target)):
				var loweredTarget = lowerExprWithPrefix(target);
				var loweredIndex = lowerExprWithPrefix(index);
				var loweredRight = lowerExprWithExpectedUpcast(right, left.t);
				var targetName = freshTempName("hx_array_target");
				var indexName = freshTempName("hx_array_index");
				var targetExpr = GoExpr.GoIdent(targetName);
				loweredTarget.prefix.concat([GoStmt.GoVarDecl(targetName, "*hxrt.Array", loweredTarget.expr, true)])
					.concat(loweredIndex.prefix)
					.concat([GoStmt.GoVarDecl(indexName, "int", loweredIndex.expr, true)])
					.concat(loweredRight.prefix)
					.concat([
						GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoSelector(targetExpr, "Set"), [GoExpr.GoIdent(indexName), loweredRight.expr]))
					]);
			case TMeta(_, inner) | TParenthesis(inner) | TCast(inner, _):
				lowerHaxeArrayIndexAssignStatements(inner, right);
			case _:
				null;
		};
	}

	function lowerHaxeArrayIndexAssignExpr(left:TypedExpr, right:TypedExpr):Null<LoweredExpr> {
		return switch (left.expr) {
			case TArray(target, index) if (usesSharedArrayCarrier(target)):
				var loweredTarget = lowerExprWithPrefix(target);
				var loweredIndex = lowerExprWithPrefix(index);
				var loweredRight = lowerExprWithExpectedUpcast(right, left.t);
				var targetName = freshTempName("hx_array_target");
				var indexName = freshTempName("hx_array_index");
				var assigned = GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent(targetName), "Set"), [GoExpr.GoIdent(indexName), loweredRight.expr]);
				var result = coerceAnyExprToType(assigned, left.t, left.t, true);
				var prefix = loweredTarget.prefix.concat([GoStmt.GoVarDecl(targetName, "*hxrt.Array", loweredTarget.expr, true)])
					.concat(loweredIndex.prefix)
					.concat([GoStmt.GoVarDecl(indexName, "int", loweredIndex.expr, true)])
					.concat(loweredRight.prefix);
				{
					expr: GoExpr.GoCall(GoExpr.GoFuncLiteral([], [typeToGoType(left.t)], prefix.concat([GoStmt.GoReturn(result)])), []),
					isStringLike: isStringType(left.t)
				};
			case TMeta(_, inner) | TParenthesis(inner) | TCast(inner, _):
				lowerHaxeArrayIndexAssignExpr(inner, right);
			case _:
				null;
		};
	}

	/**
		What: Lowers a compound indexed mutation through the shared Array carrier.
		Why: A carrier pointer is not a Go indexable lvalue, and expanding the source
		operation with repeated receiver/index expressions would duplicate side effects.
		How: Capture receiver, index, and the current typed element in source order,
		then evaluate the right side once, compute the assigned value, store it through
		`Set`, and return that same value for expression-form assignments.
	**/
	function lowerHaxeArrayIndexAssignOpExpr(left:TypedExpr, right:TypedExpr, op:Binop, sourcePos:haxe.macro.Expr.Position):Null<LoweredExpr> {
		return switch (left.expr) {
			case TArray(target, index) if (usesSharedArrayCarrier(target)):
				var loweredTarget = lowerExprWithPrefix(target);
				var loweredIndex = lowerExprWithPrefix(index);
				var loweredRight = lowerExprWithPrefix(right);
				var targetName = freshTempName("hx_array_target");
				var indexName = freshTempName("hx_array_index");
				var currentName = freshTempName("hx_array_current");
				var assignedName = freshTempName("hx_array_assigned");
				var targetExpr = GoExpr.GoIdent(targetName);
				var indexExpr = GoExpr.GoIdent(indexName);
				var currentExpr = coerceStoredArrayElementExpr(GoExpr.GoCall(GoExpr.GoSelector(targetExpr, "Get"), [indexExpr]), left.t);
				var stringAppendFromSharedArray = op == OpAdd
					&& (isStringType(left.t) || isStringType(right.t))
					&& isSharedArrayElementExpr(right);
				var rightExpr = stringAppendFromSharedArray ? lowerSharedArrayElementStorageExpr(right) : upcastIfNeeded(loweredRight.expr, right.t, left.t,
					right);
				if (!stringAppendFromSharedArray && isSharedArrayElementExpr(right)) {
					rightExpr = coerceStoredArrayElementExpr(rightExpr, left.t);
				}
				var assignedExpr = lowerAssignOpExpr(op, GoExpr.GoIdent(currentName), rightExpr, left.t, right.t, sourcePos, stringAppendFromSharedArray);
				var resultGoType = valueStorageGoType(left.t);
				var body = loweredTarget.prefix.concat([GoStmt.GoVarDecl(targetName, "*hxrt.Array", loweredTarget.expr, true)])
					.concat(loweredIndex.prefix)
					.concat([
						GoStmt.GoVarDecl(indexName, "int", loweredIndex.expr, true),
						GoStmt.GoVarDecl(currentName, resultGoType, currentExpr, false)
					])
					.concat(loweredRight.prefix)
					.concat([
						GoStmt.GoVarDecl(assignedName, resultGoType, assignedExpr, false),
						GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoSelector(targetExpr, "Set"), [indexExpr, GoExpr.GoIdent(assignedName)])),
						GoStmt.GoReturn(GoExpr.GoIdent(assignedName))
					]);
				{
					expr: GoExpr.GoCall(GoExpr.GoFuncLiteral([], [resultGoType], body), []),
					isStringLike: isStringType(left.t)
				};
			case TMeta(_, inner) | TParenthesis(inner) | TCast(inner, _):
				lowerHaxeArrayIndexAssignOpExpr(inner, right, op, sourcePos);
			case _:
				null;
		};
	}

	/**
		What: Lowers prefix/postfix increment and decrement on one shared Array slot.
		Why: Unary mutation needs the same once-only receiver/index capture as compound
		assignment, plus the correct old-vs-new expression result.
		How: Read the typed element once, compute and store the next value through
		`Set`, then return the captured old value for postfix or the new value for prefix.
	**/
	function lowerHaxeArrayIndexUnitExpr(value:TypedExpr, op:Unop, postFix:Bool, sourcePos:haxe.macro.Expr.Position):Null<LoweredExpr> {
		return switch (value.expr) {
			case TArray(target, index) if (usesSharedArrayCarrier(target)):
				var loweredTarget = lowerExprWithPrefix(target);
				var loweredIndex = lowerExprWithPrefix(index);
				var targetName = freshTempName("hx_array_target");
				var indexName = freshTempName("hx_array_index");
				var currentName = freshTempName("hx_array_current");
				var nextName = freshTempName("hx_array_next");
				var targetExpr = GoExpr.GoIdent(targetName);
				var indexExpr = GoExpr.GoIdent(indexName);
				var currentExpr = coerceStoredArrayElementExpr(GoExpr.GoCall(GoExpr.GoSelector(targetExpr, "Get"), [indexExpr]), value.t);
				var opSymbol = op == OpIncrement ? GoBinaryOperator.Add : GoBinaryOperator.Subtract;
				var resultGoType = valueStorageGoType(value.t);
				var body = loweredTarget.prefix.concat([GoStmt.GoVarDecl(targetName, "*hxrt.Array", loweredTarget.expr, true)])
					.concat(loweredIndex.prefix)
					.concat([
						GoStmt.GoVarDecl(indexName, "int", loweredIndex.expr, true),
						GoStmt.GoVarDecl(currentName, resultGoType, currentExpr, false),
						GoStmt.GoVarDecl(nextName, resultGoType, unitStepExpr(GoExpr.GoIdent(currentName), opSymbol, value.t, sourcePos), false),
						GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoSelector(targetExpr, "Set"), [indexExpr, GoExpr.GoIdent(nextName)])),
						GoStmt.GoReturn(GoExpr.GoIdent(postFix ? currentName : nextName))
					]);
				{
					expr: GoExpr.GoCall(GoExpr.GoFuncLiteral([], [resultGoType], body), []),
					isStringLike: false
				};
			case TMeta(_, inner) | TParenthesis(inner) | TCast(inner, _):
				lowerHaxeArrayIndexUnitExpr(inner, op, postFix, sourcePos);
			case _:
				null;
		};
	}

	function lowerArrayLengthAssign(left:TypedExpr, rightExpr:GoExpr):Null<Array<GoStmt>> {
		return switch (left.expr) {
			case TField(target, access):
				var fieldName = fieldAccessName(access);
				if (fieldName != "length" || !isArrayType(target.t)) {
					null;
				} else {
					var targetExpr = lowerExpr(target).expr;
					if (usesSharedArrayCarrier(target)) {
						return [
							GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoSelector(targetExpr, "SetLength"), [rightExpr]))
						];
					}
					var desiredLenName = freshTempName("hx_len");
					var zeroName = freshTempName("hx_zero");
					var desiredLen = GoExpr.GoIdent(desiredLenName);
					var currentLen = GoExpr.GoCall(GoExpr.GoIdent("len"), [targetExpr]);
					[
						GoStmt.GoVarDecl(desiredLenName, "int", rightExpr, true),
						GoStmt.GoIf(GoExpr.GoBinary("<", desiredLen, GoExpr.GoIntLiteral(0)), [GoStmt.GoAssign(desiredLen, GoExpr.GoIntLiteral(0))], null),
						GoStmt.GoIf(GoExpr.GoBinary("<=", desiredLen, currentLen),
							[GoStmt.GoAssign(targetExpr, GoExpr.GoSlice(targetExpr, null, desiredLen))], [
								GoStmt.GoVarDecl(zeroName, arrayElementGoType(target.t), null, false),
								GoStmt.GoWhile(GoExpr.GoBinary("<", GoExpr.GoCall(GoExpr.GoIdent("len"), [targetExpr]), desiredLen), [
									GoStmt.GoAssign(targetExpr, GoExpr.GoCall(GoExpr.GoIdent("append"), [targetExpr, GoExpr.GoIdent(zeroName)]))
								])
							])
					];
				}
			case _:
				null;
		};
	}

	function lowerFunctionBody(expr:TypedExpr):Array<GoStmt> {
		pushLocalScope();
		pushReturnRedirectMask();
		var out = lowerToStatements(expr);
		popReturnRedirect();
		popLocalScope();
		return out;
	}

	function lowerSwitchStmt(value:TypedExpr, cases:Array<{values:Array<TypedExpr>, expr:TypedExpr}>, defaultExpr:Null<TypedExpr>):GoStmt {
		var stringSwitch = isStringType(value.t);
		var loweredCases = new Array<GoSwitchCase>();
		for (caseEntry in cases) {
			loweredCases.push({
				values: [
					for (caseValue in caseEntry.values)
						stringSwitch ? lowerStringComparableExpr(caseValue) : lowerExpr(caseValue).expr
				],
				body: lowerInSwitchContext(function() return lowerToStatements(caseEntry.expr))
			});
		}

		return GoStmt.GoSwitch(stringSwitch ? lowerStringComparableExpr(value) : lowerExpr(value).expr, loweredCases,
			defaultExpr == null ? null : lowerInSwitchContext(function() return lowerToStatements(defaultExpr)));
	}

	function lowerSwitchExpr(value:TypedExpr, cases:Array<{values:Array<TypedExpr>, expr:TypedExpr}>, defaultExpr:Null<TypedExpr>,
			resultType:Type):LoweredExprWithPrefix {
		var temp = freshTempName("hx_switch");
		var stringSwitch = isStringType(value.t);
		var loweredCases = new Array<GoSwitchCase>();

		for (caseEntry in cases) {
			var loweredCase = lowerInSwitchContext(function() return lowerExprWithPrefix(caseEntry.expr));
			var caseBody = loweredCase.prefix.concat([GoStmt.GoAssign(GoExpr.GoIdent(temp), loweredCase.expr)]);
			loweredCases.push({
				values: [
					for (caseValue in caseEntry.values)
						stringSwitch ? lowerStringComparableExpr(caseValue) : lowerExpr(caseValue).expr
				],
				body: caseBody
			});
		}

		var defaultBody:Null<Array<GoStmt>> = null;
		if (defaultExpr != null) {
			var loweredDefault = lowerInSwitchContext(function() return lowerExprWithPrefix(defaultExpr));
			defaultBody = loweredDefault.prefix.concat([GoStmt.GoAssign(GoExpr.GoIdent(temp), loweredDefault.expr)]);
		}

		return {
			prefix: [
				GoStmt.GoVarDecl(temp, valueStorageGoType(resultType), null, false),
				GoStmt.GoSwitch(stringSwitch ? lowerStringComparableExpr(value) : lowerExpr(value).expr, loweredCases, defaultBody)
			],
			expr: GoExpr.GoIdent(temp),
			isStringLike: isStringType(resultType)
		};
	}

	function lowerStringComparableExpr(expr:TypedExpr):GoExpr {
		return GoExpr.GoUnary("*", GoExpr.GoCall(GoExpr.GoIdent("hxrt.StdString"), [lowerExpr(expr).expr]));
	}

	function lowerLoopStmt(condition:GoExpr, body:TypedExpr):GoStmt {
		var target:LoopBreakTarget = {label: null};
		loopBreakTargetScopes.push(target);
		var bodyStmts = lowerToStatements(body);
		loopBreakTargetScopes.pop();
		var loopStmt:GoStmt = GoStmt.GoWhile(condition, bodyStmts);
		if (target.label != null) {
			return GoStmt.GoLabeled(target.label, loopStmt);
		}
		return loopStmt;
	}

	function lowerInSwitchContext<T>(lower:Void->T):T {
		switchDepth++;
		var out = lower();
		switchDepth--;
		return out;
	}

	function currentLoopBreakLabel():Null<String> {
		if (loopBreakTargetScopes.length == 0) {
			return null;
		}
		var target = loopBreakTargetScopes[loopBreakTargetScopes.length - 1];
		if (target.label == null) {
			target.label = freshTempName("hx_loop");
		}
		return target.label;
	}

	function lowerIfExpr(condition:TypedExpr, thenBranch:TypedExpr, elseBranch:Null<TypedExpr>, resultType:Type):LoweredExprWithPrefix {
		var elseExpr = elseBranch;
		if (elseExpr == null) {
			Context.fatalError("If-expression requires an else branch", condition.pos);
		}

		var loweredCondition = lowerExprWithPrefix(condition);
		var facts = conditionNonNullFacts(condition);
		var loweredThen = lowerWithNonNullPrimitiveFacts(facts.thenFacts, function() return lowerExprWithExpectedUpcast(thenBranch, resultType));
		var loweredElse = lowerWithNonNullPrimitiveFacts(facts.elseFacts, function() return lowerExprWithExpectedUpcast(elseExpr, resultType));
		var temp = freshTempName("hx_if");
		var loweredThenValue = loweredThen.expr;
		var loweredElseValue = loweredElse.expr;
		var thenKnownNonNullPrimitive = nonNullPrimitiveExprGoTypeWithFacts(thenBranch, facts.thenFacts) != null;
		var elseKnownNonNullPrimitive = nonNullPrimitiveExprGoTypeWithFacts(elseExpr, facts.elseFacts) != null;
		loweredThenValue = coerceAnyExprToType(loweredThenValue, thenBranch.t, resultType, !thenKnownNonNullPrimitive && (exprBackedByAny(thenBranch)
			|| shouldForceAnyCoerce(thenBranch.t, resultType)));
		loweredElseValue = coerceAnyExprToType(loweredElseValue, elseExpr.t, resultType, !elseKnownNonNullPrimitive && (exprBackedByAny(elseExpr)
			|| shouldForceAnyCoerce(elseExpr.t, resultType)));

		var prefix = [GoStmt.GoVarDecl(temp, valueStorageGoType(resultType), null, false)].concat(loweredCondition.prefix);

		prefix.push(GoStmt.GoIf(loweredCondition.expr, loweredThen.prefix.concat([GoStmt.GoAssign(GoExpr.GoIdent(temp), loweredThenValue)]),
			loweredElse.prefix.concat([GoStmt.GoAssign(GoExpr.GoIdent(temp), loweredElseValue)])));

		return {
			prefix: prefix,
			expr: GoExpr.GoIdent(temp),
			isStringLike: isStringType(resultType)
		};
	}

	function lowerTryCatchStmt(tryExpr:TypedExpr, catches:Array<{v:TVar, expr:TypedExpr}>):Array<GoStmt> {
		var redirect:Null<ReturnRedirect> = null;
		var redirectPrefix = new Array<GoStmt>();
		var outerReturnType = currentFunctionReturnType();
		if (outerReturnType != null && tryCatchContainsReturn(tryExpr, catches)) {
			var flagName = freshTempName("hx_try_return");
			var valueName:Null<String> = null;
			var valueType:Null<Type> = null;

			redirectPrefix.push(GoStmt.GoVarDecl(flagName, "bool", GoExpr.GoBoolLiteral(false), true));
			if (!isVoidType(outerReturnType)) {
				valueName = freshTempName("hx_try_value");
				valueType = outerReturnType;
				redirectPrefix.push(GoStmt.GoVarDecl(valueName, typeToGoType(outerReturnType), null, false));
			}

			redirect = {
				flagName: flagName,
				valueName: valueName,
				valueType: valueType
			};
		}

		if (catches.length == 0) {
			var tryBody = lowerStatementsWithReturnType(tryExpr, voidType(), redirect);
			var out = redirectPrefix.concat([GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoFuncLiteral([], [], tryBody), []))]);
			if (redirect != null) {
				out.push(tryReturnRedirectStmt(redirect));
			}
			return out;
		}

		var caughtName = freshTempName("hx_caught");
		var typeBindingName = freshTempName("hx_typed");
		var typedCases = new Array<GoTypeSwitchCase>();
		var dynamicBody:Null<Array<GoStmt>> = null;

		for (index in 0...catches.length) {
			var catchEntry = catches[index];
			var catchVarName = localVarName(catchEntry.v);
			var catchType = typeToGoType(catchEntry.v.t);
			var catchExprBody = lowerStatementsWithReturnType(catchEntry.expr, voidType(), redirect);
			var haxeExceptionCatch = isHaxeExceptionType(catchEntry.v.t);
			var dynamicCatch = isDynamicCatchType(catchEntry.v.t) || haxeExceptionCatch || catchType == "any";

			if (dynamicCatch) {
				if (index != catches.length - 1) {
					Context.fatalError("Dynamic catch must be the final catch clause", catchEntry.expr.pos);
				}
				var dynamicValueExpr = haxeExceptionCatch ? GoExpr.GoCall(GoExpr.GoIdent("hxrt.ExceptionCaught"),
					[GoExpr.GoIdent(caughtName)]) : GoExpr.GoIdent(caughtName);
				var dynamicValueType = haxeExceptionCatch ? "*hxrt.ExceptionValue" : "any";
				dynamicBody = [
					GoStmt.GoVarDecl(catchVarName, dynamicValueType, dynamicValueExpr, true),
					GoStmt.GoAssign(GoExpr.GoIdent("_"), GoExpr.GoIdent(catchVarName))
				].concat(catchExprBody);
			} else {
				typedCases.push({
					typeName: catchType,
					body: [
						GoStmt.GoVarDecl(catchVarName, catchType, GoExpr.GoIdent(typeBindingName), true),
						GoStmt.GoAssign(GoExpr.GoIdent("_"), GoExpr.GoIdent(catchVarName))
					].concat(catchExprBody)
				});
			}
		}

		if (dynamicBody == null) {
			dynamicBody = [
				GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("hxrt.Throw"), [GoExpr.GoIdent(caughtName)]))
			];
		}

		var catchBody:Array<GoStmt> = if (typedCases.length == 0) {
			dynamicBody;
		} else {
			[
				GoStmt.GoTypeSwitch(GoExpr.GoIdent(caughtName), typeBindingName, typedCases, dynamicBody)
			];
		};

		var out = redirectPrefix.concat([
			GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("hxrt.TryCatch"), [
				GoExpr.GoFuncLiteral([], [], lowerStatementsWithReturnType(tryExpr, voidType(), redirect)),
				GoExpr.GoFuncLiteral([
					{
						name: caughtName,
						typeName: "any"
					}
				], [], catchBody)
			]))
		]);
		if (redirect != null) {
			out.push(tryReturnRedirectStmt(redirect));
		}
		return out;
	}

	function tryReturnRedirectStmt(redirect:ReturnRedirect):GoStmt {
		var returnStmt = redirect.valueName == null ? GoStmt.GoReturn(null) : GoStmt.GoReturn(GoExpr.GoIdent(redirect.valueName));
		return GoStmt.GoIf(GoExpr.GoIdent(redirect.flagName), [returnStmt], null);
	}

	function tryCatchContainsReturn(tryExpr:TypedExpr, catches:Array<{v:TVar, expr:TypedExpr}>):Bool {
		if (exprContainsReturn(tryExpr)) {
			return true;
		}
		for (catchEntry in catches) {
			if (exprContainsReturn(catchEntry.expr)) {
				return true;
			}
		}
		return false;
	}

	function exprContainsReturn(expr:TypedExpr):Bool {
		return switch (expr.expr) {
			case TReturn(_):
				true;
			case TFunction(_):
				false;
			case TMeta(_, inner):
				exprContainsReturn(inner);
			case TParenthesis(inner):
				exprContainsReturn(inner);
			case TCast(inner, _):
				exprContainsReturn(inner);
			case TEnumIndex(inner):
				exprContainsReturn(inner);
			case TEnumParameter(target, _, _):
				exprContainsReturn(target);
			case TVar(_, value): value != null && exprContainsReturn(value);
			case TArray(target, index): exprContainsReturn(target) || exprContainsReturn(index);
			case TBinop(_, left, right): exprContainsReturn(left) || exprContainsReturn(right);
			case TUnop(_, _, value):
				exprContainsReturn(value);
			case TField(target, _):
				exprContainsReturn(target);
			case TNew(_, _, args):
				var hasReturn = false;
				for (arg in args) {
					if (exprContainsReturn(arg)) {
						hasReturn = true;
						break;
					}
				}
				hasReturn;
			case TCall(callee, args):
				if (exprContainsReturn(callee)) {
					true;
				} else {
					var hasReturn = false;
					for (arg in args) {
						if (exprContainsReturn(arg)) {
							hasReturn = true;
							break;
						}
					}
					hasReturn;
				}
			case TObjectDecl(fields):
				var hasReturn = false;
				for (field in fields) {
					if (exprContainsReturn(field.expr)) {
						hasReturn = true;
						break;
					}
				}
				hasReturn;
			case TArrayDecl(values):
				var hasReturn = false;
				for (value in values) {
					if (exprContainsReturn(value)) {
						hasReturn = true;
						break;
					}
				}
				hasReturn;
			case TBlock(exprs):
				var hasReturn = false;
				for (inner in exprs) {
					if (exprContainsReturn(inner)) {
						hasReturn = true;
						break;
					}
				}
				hasReturn;
			case TIf(condition, thenBranch, elseBranch): exprContainsReturn(condition) || exprContainsReturn(thenBranch) || (elseBranch != null
					&& exprContainsReturn(elseBranch));
			case TSwitch(value, cases, defaultExpr):
				if (exprContainsReturn(value)) {
					true;
				} else {
					var hasReturn = false;
					for (caseEntry in cases) {
						for (caseValue in caseEntry.values) {
							if (exprContainsReturn(caseValue)) {
								hasReturn = true;
								break;
							}
						}
						if (hasReturn || exprContainsReturn(caseEntry.expr)) {
							hasReturn = true;
							break;
						}
					}
					hasReturn || (defaultExpr != null && exprContainsReturn(defaultExpr))
					;
				}
			case TWhile(condition, body, _): exprContainsReturn(condition) || exprContainsReturn(body);
			case TFor(_, iterator, body): exprContainsReturn(iterator) || exprContainsReturn(body);
			case TTry(innerTry, innerCatches):
				if (exprContainsReturn(innerTry)) {
					true;
				} else {
					var hasReturn = false;
					for (catchEntry in innerCatches) {
						if (exprContainsReturn(catchEntry.expr)) {
							hasReturn = true;
							break;
						}
					}
					hasReturn;
				}
			case TThrow(value):
				exprContainsReturn(value);
			case TConst(_):
				false;
			case TLocal(_):
				false;
			case TIdent(_):
				false;
			case TTypeExpr(_):
				false;
			case TBreak:
				false;
			case TContinue:
				false;
		};
	}

	function lowerTryCatchExpr(tryExpr:TypedExpr, catches:Array<{v:TVar, expr:TypedExpr}>, resultType:Type):LoweredExprWithPrefix {
		var temp = freshTempName("hx_try");
		var loweredTry = lowerExprWithExpectedUpcast(tryExpr, resultType);
		var loweredTryValue = loweredTry.expr;
		var tempExpr = GoExpr.GoIdent(temp);

		if (catches.length == 0) {
			return {
				prefix: [
					GoStmt.GoVarDecl(temp, valueStorageGoType(resultType), null, false),
					GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoFuncLiteral([], [], loweredTry.prefix.concat([GoStmt.GoAssign(tempExpr, loweredTryValue)])), []))
				],
				expr: tempExpr,
				isStringLike: isStringType(resultType)
			};
		}

		var caughtName = freshTempName("hx_caught");
		var typeBindingName = freshTempName("hx_typed");
		var typedCases = new Array<GoTypeSwitchCase>();
		var dynamicBody:Null<Array<GoStmt>> = null;

		for (index in 0...catches.length) {
			var catchEntry = catches[index];
			var catchVarName = localVarName(catchEntry.v);
			var catchType = typeToGoType(catchEntry.v.t);
			var loweredCatch = lowerExprWithExpectedUpcast(catchEntry.expr, resultType);
			var loweredCatchValue = loweredCatch.expr;
			var catchExprBody = loweredCatch.prefix.concat([GoStmt.GoAssign(tempExpr, loweredCatchValue)]);
			var haxeExceptionCatch = isHaxeExceptionType(catchEntry.v.t);
			var dynamicCatch = isDynamicCatchType(catchEntry.v.t) || haxeExceptionCatch || catchType == "any";

			if (dynamicCatch) {
				if (index != catches.length - 1) {
					Context.fatalError("Dynamic catch must be the final catch clause", catchEntry.expr.pos);
				}
				var dynamicValueExpr = haxeExceptionCatch ? GoExpr.GoCall(GoExpr.GoIdent("hxrt.ExceptionCaught"),
					[GoExpr.GoIdent(caughtName)]) : GoExpr.GoIdent(caughtName);
				var dynamicValueType = haxeExceptionCatch ? "*hxrt.ExceptionValue" : "any";
				dynamicBody = [
					GoStmt.GoVarDecl(catchVarName, dynamicValueType, dynamicValueExpr, true),
					GoStmt.GoAssign(GoExpr.GoIdent("_"), GoExpr.GoIdent(catchVarName))
				].concat(catchExprBody);
			} else {
				typedCases.push({
					typeName: catchType,
					body: [
						GoStmt.GoVarDecl(catchVarName, catchType, GoExpr.GoIdent(typeBindingName), true),
						GoStmt.GoAssign(GoExpr.GoIdent("_"), GoExpr.GoIdent(catchVarName))
					].concat(catchExprBody)
				});
			}
		}

		if (dynamicBody == null) {
			dynamicBody = [
				GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("hxrt.Throw"), [GoExpr.GoIdent(caughtName)]))
			];
		}

		var catchBody:Array<GoStmt> = if (typedCases.length == 0) {
			dynamicBody;
		} else {
			[
				GoStmt.GoTypeSwitch(GoExpr.GoIdent(caughtName), typeBindingName, typedCases, dynamicBody)
			];
		};

		return {
			prefix: [
				GoStmt.GoVarDecl(temp, valueStorageGoType(resultType), null, false),
				GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("hxrt.TryCatch"), [
					GoExpr.GoFuncLiteral([], [], loweredTry.prefix.concat([GoStmt.GoAssign(tempExpr, loweredTryValue)])),
					GoExpr.GoFuncLiteral([
						{
							name: caughtName,
							typeName: "any"
						}
					], [], catchBody)
				]))
			],
			expr: tempExpr,
			isStringLike: isStringType(resultType)
		};
	}

	function lowerObjectDeclExpr(fields:Array<{name:String, expr:TypedExpr}>):LoweredExprWithPrefix {
		var temp = freshTempName("hx_obj");
		var objectType = GoType.map(GoType.builtin(GoBuiltinType.StringType), GoType.builtin(GoBuiltinType.AnyType));
		var prefix = [
			GoStmt.GoVarDecl(temp, objectType, GoExpr.GoCompositeLiteral(objectType, []), true)
		];
		var targetExpr = GoExpr.GoIdent(temp);

		for (field in fields) {
			var loweredValue = lowerExprWithPrefix(field.expr);
			prefix = prefix.concat(loweredValue.prefix);
			prefix.push(GoStmt.GoAssign(GoExpr.GoIndex(targetExpr, GoExpr.GoStringLiteral(field.name)), loweredValue.expr));
		}

		return {
			prefix: prefix,
			expr: targetExpr,
			isStringLike: false
		};
	}

	/**
		What: Lowers a statement block while retaining local and continuation facts.
		Why: A throw inside any non-final statement has later generated code, so adding
		a synthetic function return there can use the wrong type after inline expansion.
		How: Suppress throw fallbacks only before the final statement; terminal throws
		still receive the enclosing function's required Go zero return.
	**/
	function lowerBlock(exprs:Array<TypedExpr>):Array<GoStmt> {
		pushLocalScope();
		var out = new Array<GoStmt>();
		var appliedNonNullFacts = new Map<Int, Null<String>>();
		for (index in 0...exprs.length) {
			var inner = exprs[index];
			registerBlockLocalReassignmentInfo(inner, exprs, index);
			var loweredInner = index < exprs.length - 1 ? withoutThrowFallback(function() return lowerToStatements(inner)) : lowerToStatements(inner);
			out = out.concat(loweredInner);
			var continuingFacts = continuingNonNullPrimitiveFactsAfterGuard(inner, exprs, index);
			for (fact in continuingFacts) {
				var scope = currentNonNullPrimitiveLocalScope();
				if (scope != null) {
					if (!appliedNonNullFacts.exists(fact.variable.id)) {
						appliedNonNullFacts.set(fact.variable.id, scope.exists(fact.variable.id) ? scope.get(fact.variable.id) : null);
					}
					scope.set(fact.variable.id, fact.goType);
				}
			}
		}
		restoreBlockNonNullPrimitiveFacts(appliedNonNullFacts);
		popLocalScope();
		return out;
	}

	function restoreBlockNonNullPrimitiveFacts(previous:Map<Int, Null<String>>):Void {
		var scope = currentNonNullPrimitiveLocalScope();
		if (scope == null) {
			return;
		}
		for (variableId => old in previous) {
			if (old == null) {
				scope.remove(variableId);
			} else {
				scope.set(variableId, old);
			}
		}
	}

	function continuingNonNullPrimitiveFactsAfterGuard(expr:TypedExpr, block:Array<TypedExpr>, index:Int):Array<{variable:TVar, goType:String}> {
		return switch (expr.expr) {
			case TIf(condition, thenBranch, elseBranch):
				var facts = conditionNonNullFacts(condition);
				var continuingFacts = new Array<{variable:TVar, goType:String}>();
				if (exprAlwaysTerminates(thenBranch)) {
					continuingFacts = facts.elseFacts;
				} else if (elseBranch != null && exprAlwaysTerminates(elseBranch)) {
					continuingFacts = facts.thenFacts;
				}
				[
					for (fact in continuingFacts)
						if (!blockAssignsToVariableAfter(block, index, fact.variable)) fact
				];
			case TMeta(_, inner) | TParenthesis(inner) | TCast(inner, _):
				continuingNonNullPrimitiveFactsAfterGuard(inner, block, index);
			case _:
				[];
		};
	}

	function exprAlwaysTerminates(expr:TypedExpr):Bool {
		return switch (expr.expr) {
			case TReturn(_) | TThrow(_):
				true;
			case TBlock(exprs): exprs.length > 0 && exprAlwaysTerminates(exprs[exprs.length - 1]);
			case TIf(_, thenBranch, elseBranch): elseBranch != null && exprAlwaysTerminates(thenBranch) && exprAlwaysTerminates(elseBranch);
			case TMeta(_, inner) | TParenthesis(inner) | TCast(inner, _):
				exprAlwaysTerminates(inner);
			case _:
				false;
		};
	}

	function pushLocalScope():Void {
		localFunctionScopes.push(new Map<String, FunctionInfo>());
		localLambdaAliasScopes.push(new Map<String, String>());
		localRestIteratorScopes.push([]);
		localNeverReassignedScopes.push(new Map<Int, Bool>());
	}

	function popLocalScope():Void {
		if (localFunctionScopes.length > 0) {
			localFunctionScopes.pop();
		}
		if (localLambdaAliasScopes.length > 0) {
			localLambdaAliasScopes.pop();
		}
		if (localRestIteratorScopes.length > 0) {
			localRestIteratorScopes.pop();
		}
		if (localNeverReassignedScopes.length > 0) {
			localNeverReassignedScopes.pop();
		}
	}

	function pushFunctionVarNameScope():Void {
		functionVarNameScopes.push(new Map<Int, String>());
		functionVarNameCountScopes.push(new Map<String, Int>());
		optionalPrimitiveParamScopes.push(new Map<Int, String>());
		nonNullPrimitiveLocalScopes.push(new Map<Int, String>());
		narrowedPrimitiveStorageScopes.push(new Map<Int, String>());
	}

	function popFunctionVarNameScope():Void {
		if (functionVarNameScopes.length > 0) {
			functionVarNameScopes.pop();
		}
		if (functionVarNameCountScopes.length > 0) {
			functionVarNameCountScopes.pop();
		}
		if (optionalPrimitiveParamScopes.length > 0) {
			optionalPrimitiveParamScopes.pop();
		}
		if (nonNullPrimitiveLocalScopes.length > 0) {
			nonNullPrimitiveLocalScopes.pop();
		}
		if (narrowedPrimitiveStorageScopes.length > 0) {
			narrowedPrimitiveStorageScopes.pop();
		}
	}

	function currentLocalNeverReassignedScope():Null<Map<Int, Bool>> {
		if (localNeverReassignedScopes.length == 0) {
			return null;
		}
		return localNeverReassignedScopes[localNeverReassignedScopes.length - 1];
	}

	function registerBlockLocalReassignmentInfo(expr:TypedExpr, block:Array<TypedExpr>, index:Int):Void {
		switch (expr.expr) {
			case TVar(variable, _):
				var scope = currentLocalNeverReassignedScope();
				if (scope != null) {
					scope.set(variable.id, !blockAssignsToVariableAfter(block, index, variable));
				}
			case _:
		}
	}

	function blockAssignsToVariableAfter(block:Array<TypedExpr>, index:Int, variable:TVar):Bool {
		var next = index + 1;
		while (next < block.length) {
			if (exprAssignsToVariable(block[next], variable)) {
				return true;
			}
			next++;
		}
		return false;
	}

	function exprAssignsToVariable(expr:TypedExpr, variable:TVar):Bool {
		return switch (expr.expr) {
			case TBinop(OpAssign, left, right) | TBinop(OpAssignOp(_), left, right): exprTargetsVariable(left,
					variable) || exprAssignsToVariable(right, variable);
			case TVar(_, value): value != null && exprAssignsToVariable(value, variable);
			case TBlock(exprs):
				var found = false;
				for (inner in exprs) {
					if (exprAssignsToVariable(inner, variable)) {
						found = true;
						break;
					}
				}
				found;
			case TIf(condition, thenBranch, elseBranch): exprAssignsToVariable(condition,
					variable) || exprAssignsToVariable(thenBranch, variable) || (elseBranch != null
					&& exprAssignsToVariable(elseBranch, variable));
			case TSwitch(value, cases, defaultExpr):
				if (exprAssignsToVariable(value, variable)) {
					true;
				} else {
					var found = false;
					for (caseEntry in cases) {
						for (caseValue in caseEntry.values) {
							if (exprAssignsToVariable(caseValue, variable)) {
								found = true;
								break;
							}
						}
						if (!found && exprAssignsToVariable(caseEntry.expr, variable)) {
							found = true;
						}
						if (found) {
							break;
						}
					}
					found || (defaultExpr != null && exprAssignsToVariable(defaultExpr, variable))
					;
				}
			case TTry(tryExpr, catches):
				if (exprAssignsToVariable(tryExpr, variable)) {
					true;
				} else {
					var found = false;
					for (catchEntry in catches) {
						if (exprAssignsToVariable(catchEntry.expr, variable)) {
							found = true;
							break;
						}
					}
					found;
				}
			case TWhile(condition, body, _): exprAssignsToVariable(condition, variable) || exprAssignsToVariable(body, variable);
			case TFor(_, iterator, body): exprAssignsToVariable(iterator, variable) || exprAssignsToVariable(body, variable);
			case TParenthesis(inner) | TMeta(_, inner) | TCast(inner, _):
				exprAssignsToVariable(inner, variable);
			case TArray(target, index): exprAssignsToVariable(target, variable) || exprAssignsToVariable(index, variable);
			case TField(target, _):
				exprAssignsToVariable(target, variable);
			case TCall(callee, args):
				if (exprAssignsToVariable(callee, variable)) {
					true;
				} else {
					var found = false;
					for (arg in args) {
						if (exprAssignsToVariable(arg, variable)) {
							found = true;
							break;
						}
					}
					found;
				}
			case TUnop(_, _, value) | TThrow(value) | TReturn(value): value != null && exprAssignsToVariable(value, variable);
			case TFunction(func):
				exprAssignsToVariable(func.expr, variable);
			case _:
				false;
		};
	}

	function exprTargetsVariable(expr:TypedExpr, variable:TVar):Bool {
		return switch (expr.expr) {
			case TLocal(target):
				target.id == variable.id;
			case TParenthesis(inner) | TMeta(_, inner) | TCast(inner, _):
				exprTargetsVariable(inner, variable);
			case _:
				false;
		};
	}

	/**
		What: Starts a generated function's return and throw-fallback scope.
		Why: Continuation-aware suppression in an outer statement must not remove the
		terminal fallback required by a nested function literal.
		How: Save the outer suppression depth, reset it for the nested function, then
		restore it when that function's return scope closes.
	**/
	function pushFunctionReturnType(returnType:Type):Void {
		functionReturnTypeScopes.push(returnType);
		throwFallbackSuppressionDepthScopes.push(throwFallbackSuppressionDepth);
		throwFallbackSuppressionDepth = 0;
	}

	function popFunctionReturnType():Void {
		if (functionReturnTypeScopes.length > 0) {
			functionReturnTypeScopes.pop();
		}
		if (throwFallbackSuppressionDepthScopes.length > 0) {
			throwFallbackSuppressionDepth = throwFallbackSuppressionDepthScopes.pop();
		}
	}

	function currentFunctionReturnType():Null<Type> {
		if (functionReturnTypeScopes.length == 0) {
			return null;
		}
		return functionReturnTypeScopes[functionReturnTypeScopes.length - 1];
	}

	/**
		What: Lowers statements that are guaranteed to have a later continuation.
		Why: A guard throw before a value block's terminal expression does not need a
		synthetic Go return; the generated tail already satisfies the function or IIFE.
		How: Suppress only fallback returns while lowering the pre-tail statement.
	**/
	function withoutThrowFallback<T>(lower:Void->T):T {
		throwFallbackSuppressionDepth++;
		var out = lower();
		throwFallbackSuppressionDepth--;
		return out;
	}

	function pushConstructorReturnScope():Void {
		constructorReturnScopes.push(true);
	}

	function popConstructorReturnScope():Void {
		if (constructorReturnScopes.length > 0) {
			constructorReturnScopes.pop();
		}
	}

	function inConstructorReturnScope():Bool {
		return constructorReturnScopes.length > 0;
	}

	function pushReturnRedirectMask():Void {
		returnRedirectScopes.push(null);
	}

	function pushReturnRedirect(redirect:ReturnRedirect):Void {
		returnRedirectScopes.push(redirect);
	}

	function popReturnRedirect():Void {
		if (returnRedirectScopes.length > 0) {
			returnRedirectScopes.pop();
		}
	}

	function currentReturnRedirect():Null<ReturnRedirect> {
		if (returnRedirectScopes.length == 0) {
			return null;
		}
		return returnRedirectScopes[returnRedirectScopes.length - 1];
	}

	function lowerStatementsWithReturnType(expr:TypedExpr, returnType:Type, ?redirect:Null<ReturnRedirect>):Array<GoStmt> {
		pushFunctionReturnType(returnType);
		if (redirect != null) {
			pushReturnRedirect(redirect);
		}
		var lowered = lowerToStatements(expr);
		if (redirect != null) {
			popReturnRedirect();
		}
		popFunctionReturnType();
		return lowered;
	}

	function voidType():Type {
		if (cachedVoidType == null) {
			cachedVoidType = Context.getType("Void");
		}
		return cachedVoidType;
	}

	function nonVoidThrowFallbackReturnStmts():Array<GoStmt> {
		if (throwFallbackSuppressionDepth > 0) {
			return [];
		}
		var returnType = currentFunctionReturnType();
		if (returnType == null || isVoidType(returnType)) {
			return [];
		}

		var zeroName = freshTempName("hx_throw_zero");
		var returnTypeName = valueStorageGoType(returnType);
		return [
			GoStmt.GoVarDecl(zeroName, returnTypeName, null, false),
			GoStmt.GoReturn(GoExpr.GoIdent(zeroName))
		];
	}

	/**
		What: Lowers a throwing value expression for its immediate expected result type.
		Why: Haxe inline expansion can give the `TThrow` node a surrounding comparison
		or interpolation type even though the enclosing accessor branch must still
		produce the accessor's own result type for Go's unreachable fallback return.
		How: Emit the shared runtime throw, then a typed zero return using the caller's
		expected storage type solely to satisfy Go's static return rules.
	**/
	function lowerThrowExprForType(value:TypedExpr, resultType:Type):LoweredExpr {
		var loweredValue = lowerExprWithPrefix(value);
		var throwBody = loweredValue.prefix.concat([
			GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("hxrt.Throw"), [loweredValue.expr]))
		]);
		if (isVoidType(resultType)) {
			return {
				expr: GoExpr.GoCall(GoExpr.GoFuncLiteral([], [], throwBody), []),
				isStringLike: false
			};
		}
		var resultTypeName = valueStorageGoType(resultType);
		var zeroName = freshTempName("hx_throw_zero");
		return {
			expr: GoExpr.GoCall(GoExpr.GoFuncLiteral([], [resultTypeName], throwBody.concat([
				GoStmt.GoVarDecl(zeroName, resultTypeName, null, false),
				GoStmt.GoReturn(GoExpr.GoIdent(zeroName))
			])), []),
			isStringLike: isStringType(resultType)
		};
	}

	/**
		What: Lowers a wrapper or block chain whose terminal value is a throw.
		Why: Inline accessor guards commonly represent their throwing branch as a
		one-expression `TBlock`, so recognizing only a bare `TThrow` still loses the
		immediate expected type.
		How: Recurse through compile-only wrappers and the final block expression,
		retaining every preceding block statement in source order; decline normal tails.
	**/
	function lowerExpectedThrowExpr(expr:TypedExpr, resultType:Type):Null<LoweredExprWithPrefix> {
		return switch (expr.expr) {
			case TThrow(value):
				var loweredThrow = lowerThrowExprForType(value, resultType);
				{
					prefix: [],
					expr: loweredThrow.expr,
					isStringLike: loweredThrow.isStringLike
				};
			case TMeta(_, inner) | TParenthesis(inner) | TCast(inner, _):
				lowerExpectedThrowExpr(inner, resultType);
			case TBlock(exprs) if (exprs.length > 0):
				var loweredTail = lowerExpectedThrowExpr(exprs[exprs.length - 1], resultType);
				if (loweredTail == null) {
					null;
				} else {
					var prefix = new Array<GoStmt>();
					for (index in 0...exprs.length - 1) {
						prefix = prefix.concat(withoutThrowFallback(function() return lowerToStatements(exprs[index])));
					}
					{
						prefix: prefix.concat(loweredTail.prefix),
						expr: loweredTail.expr,
						isStringLike: loweredTail.isStringLike
					};
				}
			case _: null;
		};
	}

	function currentFunctionVarNameScope():Null<Map<Int, String>> {
		if (functionVarNameScopes.length == 0) {
			return null;
		}
		return functionVarNameScopes[functionVarNameScopes.length - 1];
	}

	function currentFunctionVarNameCountScope():Null<Map<String, Int>> {
		if (functionVarNameCountScopes.length == 0) {
			return null;
		}
		return functionVarNameCountScopes[functionVarNameCountScopes.length - 1];
	}

	function localVarName(variable:TVar):String {
		var index = functionVarNameScopes.length - 1;
		while (index >= 0) {
			var scope = functionVarNameScopes[index];
			if (scope.exists(variable.id)) {
				return scope.get(variable.id);
			}
			index--;
		}

		var base = normalizeIdent(variable.name);
		var currentScope = currentFunctionVarNameScope();
		var countScope = currentFunctionVarNameCountScope();
		if (currentScope == null || countScope == null) {
			return base;
		}

		var next = countScope.exists(base) ? countScope.get(base) : 0;
		countScope.set(base, next + 1);
		var assigned = next == 0 ? base : base + "_" + next;
		currentScope.set(variable.id, assigned);
		return assigned;
	}

	function isOptionalPrimitiveFunctionArg(arg:{v:TVar, value:Null<TypedExpr>}, typedArg:Null<{name:String, opt:Bool, t:Type}>):Bool {
		if (typedArg == null || !typedArg.opt) {
			return false;
		}
		if (!isGoNilDefaultValue(arg.value)) {
			return false;
		}
		var goType = typeToGoType(arg.v.t);
		return goType == "int" || goType == "float64" || goType == "bool";
	}

	function isNullablePrimitiveParamType(variableType:Type, typedArg:Null<{name:String, opt:Bool, t:Type}>):Bool {
		if (typedArg != null && typedArg.opt) {
			return false;
		}
		return isNullablePrimitiveType(variableType) || (typedArg != null && !typedArg.opt && isNullablePrimitiveType(typedArg.t));
	}

	function nilCapablePrimitiveParamType(type:Type):String {
		return isNullablePrimitiveType(type) ? "any" : scalarGoType(type);
	}

	/**
		What:
		Returns the Go storage type for an enum constructor payload.

		Why:
		`TEnumParameter` can appear with a narrowed expression type, but the
		constructor declaration is the source of truth for whether the payload was
		`Null<Int>`, `Null<Float>`, or `Null<Bool>`. Those payloads must stay `any`
		until a null guard proves they contain the primitive value.

		How:
		Look up the constructor argument type by index and use `any` only for
		declared nullable primitives. For generic payload constructors, prefer the
		concrete expression type so `Either<String, Int>` still extracts `*string`
		and `int` from the enum payload array.
	**/
	function enumPayloadStorageGoType(constructor:EnumField, index:Int, fallback:Type):String {
		var ctorArgs = enumConstructorArgs(constructor.type);
		var payloadType = index >= 0 && index < ctorArgs.length ? ctorArgs[index].t : fallback;
		if (isNullablePrimitiveType(payloadType)) {
			return "any";
		}
		var fallbackType = scalarGoType(fallback);
		return fallbackType != "any" ? fallbackType : scalarGoType(payloadType);
	}

	function isGoNilDefaultValue(expr:Null<TypedExpr>):Bool {
		if (expr == null) {
			return false;
		}
		return switch (expr.expr) {
			case TConst(TNull):
				true;
			case TMeta(_, inner):
				isGoNilDefaultValue(inner);
			case TParenthesis(inner):
				isGoNilDefaultValue(inner);
			case TCast(inner, _):
				isGoNilDefaultValue(inner);
			case _:
				false;
		};
	}

	function registerOptionalPrimitiveParam(variable:TVar, isOptionalPrimitive:Bool):Void {
		if (!isOptionalPrimitive) {
			return;
		}
		var goType = typeToGoType(variable.t);
		if (goType != "int" && goType != "float64" && goType != "bool") {
			return;
		}
		var scope = currentOptionalPrimitiveParamScope();
		if (scope == null) {
			return;
		}
		scope.set(variable.id, goType);
	}

	function currentOptionalPrimitiveParamScope():Null<Map<Int, String>> {
		if (optionalPrimitiveParamScopes.length == 0) {
			return null;
		}
		return optionalPrimitiveParamScopes[optionalPrimitiveParamScopes.length - 1];
	}

	function isRegisteredOptionalPrimitiveParam(variable:TVar):Bool {
		return registeredOptionalPrimitiveParamGoType(variable) != null;
	}

	function registeredOptionalPrimitiveParamGoType(variable:TVar):Null<String> {
		var index = optionalPrimitiveParamScopes.length - 1;
		while (index >= 0) {
			var scope = optionalPrimitiveParamScopes[index];
			if (scope.exists(variable.id)) {
				return scope.get(variable.id);
			}
			index--;
		}
		return null;
	}

	function nullablePrimitiveValueGoType(type:Type):Null<String> {
		if (isNullableIntType(type) || isIntType(type) || isHaxeInt32Type(type)) {
			return "int";
		}
		if (isNullableFloatType(type) || isFloatType(type)) {
			return "float64";
		}
		if (isNullableBoolType(type) || isBoolType(type)) {
			return "bool";
		}
		return null;
	}

	function primitiveNilCapableLocalGoType(variable:TVar):Null<String> {
		var optionalGoType = registeredOptionalPrimitiveParamGoType(variable);
		if (optionalGoType != null) {
			return optionalGoType;
		}
		return isNullablePrimitiveType(variable.t) ? nullablePrimitiveValueGoType(variable.t) : null;
	}

	function currentNonNullPrimitiveLocalScope():Null<Map<Int, String>> {
		if (nonNullPrimitiveLocalScopes.length == 0) {
			return null;
		}
		return nonNullPrimitiveLocalScopes[nonNullPrimitiveLocalScopes.length - 1];
	}

	function currentNarrowedPrimitiveStorageScope():Null<Map<Int, String>> {
		if (narrowedPrimitiveStorageScopes.length == 0) {
			return null;
		}
		return narrowedPrimitiveStorageScopes[narrowedPrimitiveStorageScopes.length - 1];
	}

	function registerNonNullPrimitiveLocal(variable:TVar, goType:String):Void {
		var scope = currentNonNullPrimitiveLocalScope();
		if (scope != null) {
			scope.set(variable.id, goType);
		}
	}

	function registerNarrowedPrimitiveStorage(variable:TVar, goType:String):Void {
		var scope = currentNarrowedPrimitiveStorageScope();
		if (scope != null) {
			scope.set(variable.id, goType);
		}
		registerNonNullPrimitiveLocal(variable, goType);
	}

	function registeredNonNullPrimitiveLocalGoType(variable:TVar):Null<String> {
		var index = nonNullPrimitiveLocalScopes.length - 1;
		while (index >= 0) {
			var scope = nonNullPrimitiveLocalScopes[index];
			if (scope.exists(variable.id)) {
				return scope.get(variable.id);
			}
			index--;
		}
		return null;
	}

	function registeredNarrowedPrimitiveStorageGoType(variable:TVar):Null<String> {
		var index = narrowedPrimitiveStorageScopes.length - 1;
		while (index >= 0) {
			var scope = narrowedPrimitiveStorageScopes[index];
			if (scope.exists(variable.id)) {
				return scope.get(variable.id);
			}
			index--;
		}
		return null;
	}

	function localNeverReassigned(variable:TVar):Bool {
		var index = localNeverReassignedScopes.length - 1;
		while (index >= 0) {
			var scope = localNeverReassignedScopes[index];
			if (scope.exists(variable.id)) {
				return scope.get(variable.id);
			}
			index--;
		}
		return false;
	}

	function nonNullPrimitiveExprGoType(expr:TypedExpr):Null<String> {
		return switch (expr.expr) {
			case TLocal(variable):
				var narrowedStorage = registeredNarrowedPrimitiveStorageGoType(variable);
				if (narrowedStorage != null) {
					narrowedStorage;
				} else {
					registeredNonNullPrimitiveLocalGoType(variable);
				}
			case TMeta(_, inner) | TParenthesis(inner) | TCast(inner, _):
				nonNullPrimitiveExprGoType(inner);
			case _:
				null;
		};
	}

	function nonNullPrimitiveExprGoTypeWithFacts(expr:TypedExpr, facts:Array<{variable:TVar, goType:String}>):Null<String> {
		var current = nonNullPrimitiveExprGoType(expr);
		if (current != null) {
			return current;
		}
		var factExpr = nullablePrimitiveLocalFact(expr);
		if (factExpr == null) {
			return null;
		}
		for (fact in facts) {
			if (fact.variable.id == factExpr.variable.id) {
				return fact.goType;
			}
		}
		return null;
	}

	function nullablePrimitiveLocalFact(expr:TypedExpr):Null<{variable:TVar, goType:String}> {
		return switch (expr.expr) {
			case TLocal(variable):
				var goType = primitiveNilCapableLocalGoType(variable);
				goType == null ? null : {variable: variable, goType: goType};
			case TMeta(_, inner) | TParenthesis(inner) | TCast(inner, _):
				nullablePrimitiveLocalFact(inner);
			case _:
				null;
		};
	}

	function conditionNonNullFacts(condition:TypedExpr):{thenFacts:Array<{variable:TVar, goType:String}>, elseFacts:Array<{variable:TVar, goType:String}>} {
		var empty = {thenFacts: [], elseFacts: []};
		return switch (condition.expr) {
			case TBinop(op, left, right):
				var leftFact = nullablePrimitiveLocalFact(left);
				var rightFact = nullablePrimitiveLocalFact(right);
				var fact = null;
				if (leftFact != null && isNullLiteralExpr(right)) {
					fact = leftFact;
				} else if (rightFact != null && isNullLiteralExpr(left)) {
					fact = rightFact;
				}
				if (fact == null) {
					empty;
				} else {
					switch (op) {
						case OpNotEq:
							{thenFacts: [fact], elseFacts: []};
						case OpEq:
							{thenFacts: [], elseFacts: [fact]};
						case _:
							empty;
					}
				}
			case TParenthesis(inner) | TMeta(_, inner) | TCast(inner, _):
				conditionNonNullFacts(inner);
			case _:
				empty;
		};
	}

	function lowerWithNonNullPrimitiveFacts<T>(facts:Array<{variable:TVar, goType:String}>, lower:Void->T):T {
		var scope = currentNonNullPrimitiveLocalScope();
		if (scope == null || facts.length == 0) {
			return lower();
		}
		var previous = new Map<Int, Null<String>>();
		for (fact in facts) {
			previous.set(fact.variable.id, scope.exists(fact.variable.id) ? scope.get(fact.variable.id) : null);
			scope.set(fact.variable.id, fact.goType);
		}
		var result = lower();
		for (fact in facts) {
			var old = previous.get(fact.variable.id);
			if (old == null) {
				scope.remove(fact.variable.id);
			} else {
				scope.set(fact.variable.id, old);
			}
		}
		return result;
	}

	function registerLocalFunction(name:String, func:TFunc):Void {
		var scope = currentLocalScope();
		if (scope == null) {
			return;
		}
		scope.set(name, buildFunctionInfo(func));
	}

	function currentLocalScope():Null<Map<String, FunctionInfo>> {
		if (localFunctionScopes.length == 0) {
			return null;
		}
		return localFunctionScopes[localFunctionScopes.length - 1];
	}

	function registerLocalLambdaAlias(name:String, alias:String):Void {
		var scope = currentLocalLambdaAliasScope();
		if (scope == null) {
			return;
		}
		scope.set(name, alias);
	}

	function currentLocalLambdaAliasScope():Null<Map<String, String>> {
		if (localLambdaAliasScopes.length == 0) {
			return null;
		}
		return localLambdaAliasScopes[localLambdaAliasScopes.length - 1];
	}

	function lookupLocalLambdaAlias(name:String):Null<String> {
		var index = localLambdaAliasScopes.length - 1;
		while (index >= 0) {
			var scope = localLambdaAliasScopes[index];
			if (scope.exists(name)) {
				return scope.get(name);
			}
			index--;
		}
		return null;
	}

	function registerRestIterator(name:String):Void {
		var scope = currentRestIteratorScope();
		if (scope == null) {
			return;
		}
		scope.push(name);
	}

	function currentRestIteratorScope():Null<Array<String>> {
		if (localRestIteratorScopes.length == 0) {
			return null;
		}
		return localRestIteratorScopes[localRestIteratorScopes.length - 1];
	}

	function isRegisteredRestIterator(name:String):Bool {
		var index = localRestIteratorScopes.length - 1;
		while (index >= 0) {
			var scope = localRestIteratorScopes[index];
			for (registered in scope) {
				if (registered == name) {
					return true;
				}
			}
			index--;
		}
		return false;
	}

	function resolveImplicitRestIteratorTarget():Null<String> {
		var index = localRestIteratorScopes.length - 1;
		while (index >= 0) {
			var scope = localRestIteratorScopes[index];
			if (scope.length > 0) {
				return scope[scope.length - 1];
			}
			index--;
		}
		return null;
	}

	function resolveFunctionInfo(callee:TypedExpr):Null<FunctionInfo> {
		return switch (callee.expr) {
			case TField(_, FStatic(classRef, field)):
				var classType = classRef.get();
				var resolved = field.get();
				var symbol = staticSymbol(classType, resolved.name);
				if (staticFunctionInfos.exists(symbol)) {
					staticFunctionInfos.get(symbol);
				} else if (GoStdlibOwnership.isCompilerOwnedAuthority(fullClassName(classType))) {
					null;
				} else {
					// Source-owned std classes can be queued after the initial static-info pass;
					// resolve their defaults lazily without changing compiler-owned vararg shims.
					var func = unwrapFunction(resolved.expr());
					if (func == null) {
						null;
					} else {
						var info = buildFunctionInfo(func);
						staticFunctionInfos.set(symbol, info);
						info;
					}
				}
			case TField(_, FInstance(_, _, field)) | TField(_, FAnon(field)) | TField(_, FClosure(_, field)):
				var func = unwrapFunction(field.get().expr());
				func == null ? null : buildFunctionInfo(func);
			case TLocal(variable):
				lookupLocalFunction(localVarName(variable));
			case TMeta(_, inner):
				resolveFunctionInfo(inner);
			case TParenthesis(inner):
				resolveFunctionInfo(inner);
			case TCast(inner, _):
				resolveFunctionInfo(inner);
			case _:
				null;
		};
	}

	function shouldApplySourceDefaultArgPadding(callee:TypedExpr):Bool {
		return switch (callee.expr) {
			case TField(_, FStatic(_, _)):
				true;
			case TField(_, FInstance(_, _, _)):
				true;
			case TField(target, FAnon(field)) | TField(target, FClosure(_, field)):
				true;
			case TLocal(_):
				true;
			case TMeta(_, inner):
				shouldApplySourceDefaultArgPadding(inner);
			case TParenthesis(inner):
				shouldApplySourceDefaultArgPadding(inner);
			case TCast(inner, _):
				shouldApplySourceDefaultArgPadding(inner);
			case _:
				false;
		};
	}

	function lookupLocalFunction(name:String):Null<FunctionInfo> {
		var index = localFunctionScopes.length - 1;
		while (index >= 0) {
			var scope = localFunctionScopes[index];
			if (scope.exists(name)) {
				return scope.get(name);
			}
			index--;
		}
		return null;
	}

	function resolveRestIteratorTargetName(target:TypedExpr):Null<String> {
		return switch (target.expr) {
			case TLocal(variable):
				var name = localVarName(variable);
				isRegisteredRestIterator(name) ? name : null;
			case TConst(TThis):
				resolveImplicitRestIteratorTarget();
			case TMeta(_, inner):
				resolveRestIteratorTargetName(inner);
			case TParenthesis(inner):
				resolveRestIteratorTargetName(inner);
			case TCast(inner, _):
				resolveRestIteratorTargetName(inner);
			case _:
				null;
		};
	}

	function restIteratorFieldName(targetName:Null<String>, fieldName:String):Null<String> {
		if (targetName == null) {
			return null;
		}
		return switch (fieldName) {
			case "args":
				targetName + "_args";
			case "current":
				targetName + "_current";
			case _:
				null;
		};
	}

	function lowerLValue(expr:TypedExpr):GoExpr {
		return switch (expr.expr) {
			case TLocal(variable):
				GoExpr.GoIdent(localVarName(variable));
			case TArray(target, index):
				GoExpr.GoIndex(lowerExpr(target).expr, lowerExpr(index).expr);
			case TField(target, access):
				switch (access) {
					case FAnon(field) if (isAnonymousObjectType(target.t)):
						GoExpr.GoIndex(lowerExpr(target).expr, GoExpr.GoStringLiteral(field.get().name));
					case FDynamic(name) if (isAnonymousObjectType(target.t)):
						GoExpr.GoIndex(lowerExpr(target).expr, GoExpr.GoStringLiteral(name));
					case _:
						lowerField(target, access).expr;
				}
			case TParenthesis(inner):
				lowerLValue(inner);
			case TCast(inner, _):
				lowerLValue(inner);
			case _:
				unsupportedExpr(expr, "Unsupported assignment target");
				GoExpr.GoNil;
		};
	}

	/** Resolves a local's representation-only array type for assignments. */
	function assignmentStorageType(expr:TypedExpr):Type {
		return switch (expr.expr) {
			case TLocal(variable):
				localArrayStorageOverrides.exists(variable.id) ? localArrayStorageOverrides.get(variable.id) : expr.t;
			case TMeta(_, inner) | TParenthesis(inner) | TCast(inner, _):
				assignmentStorageType(inner);
			case _:
				expr.t;
		};
	}

	/**
		What:
		Builds a safe mutation plan for shared Arrays and raw native slices.

		Why:
		The shared carrier owns header updates internally. Raw Go slices still require
		explicit write-back when `append` replaces a field or computed lvalue header.

		How:
		Keep plain locals direct. Capture other receivers once; retain the original
		raw-slice write-back paths, while shared carriers use a no-op write-back.
	**/
	function lowerArrayMutationSite(target:TypedExpr):ArrayMutationSite {
		var shared = usesSharedArrayCarrier(target);
		var sliceType = typeToGoType(target.t);
		var tempName = freshTempName("hx_arr");
		var tempExpr = GoExpr.GoIdent(tempName);

		return switch (target.expr) {
			case TLocal(_):
				{
					prefix: [],
					tempExpr: lowerLValue(target),
					sliceType: sliceType,
					writeBack: function(_value:GoExpr):Array<GoStmt> {
						return [];
					}
				};
			case TField(parent, FAnon(field)) if (!shared && isAnonymousObjectType(parent.t)):
				var objectName = freshTempName("hx_obj");
				var objectExpr = GoExpr.GoIdent(objectName);
				var fieldName = field.get().name;
				{
					prefix: [
						GoStmt.GoVarDecl(objectName, typeToGoType(parent.t), lowerExpr(parent).expr, true),
						GoStmt.GoVarDecl(tempName, sliceType, lowerAnonymousFieldRead(objectExpr, fieldName, target.t), true)
					],
					tempExpr: tempExpr,
					sliceType: sliceType,
					writeBack: function(value:GoExpr):Array<GoStmt> {
						return [
							GoStmt.GoAssign(GoExpr.GoIndex(objectExpr, GoExpr.GoStringLiteral(fieldName)), value)
						];
					}
				};
			case TField(parent, FDynamic(name)) if (!shared && isAnonymousObjectType(parent.t)):
				var objectName = freshTempName("hx_obj");
				var objectExpr = GoExpr.GoIdent(objectName);
				{
					prefix: [
						GoStmt.GoVarDecl(objectName, typeToGoType(parent.t), lowerExpr(parent).expr, true),
						GoStmt.GoVarDecl(tempName, sliceType, lowerAnonymousFieldRead(objectExpr, name, target.t), true)
					],
					tempExpr: tempExpr,
					sliceType: sliceType,
					writeBack: function(value:GoExpr):Array<GoStmt> {
						return [GoStmt.GoAssign(GoExpr.GoIndex(objectExpr, GoExpr.GoStringLiteral(name)), value)];
					}
				};
			case TParenthesis(inner):
				lowerArrayMutationSite(inner);
			case TMeta(_, inner):
				lowerArrayMutationSite(inner);
			case TCast(inner, _):
				lowerArrayMutationSite(inner);
			case _:
				var targetExpr:GoExpr = shared ? GoExpr.GoNil : lowerLValue(target);
				{
					prefix: [GoStmt.GoVarDecl(tempName, sliceType, lowerExpr(target).expr, true)],
					tempExpr: tempExpr,
					sliceType: sliceType,
					writeBack: function(value:GoExpr):Array<GoStmt> {
						return shared ? [] : [GoStmt.GoAssign(targetExpr, value)];
					}
				};
		};
	}

	function lowerExpr(expr:TypedExpr):LoweredExpr {
		return switch (expr.expr) {
			case TConst(constant):
				lowerConst(constant);
			case TArrayDecl(values):
				{
					expr: isHaxeArrayType(expr.t) ? GoExpr.GoCall(GoExpr.GoIdent("hxrt.NewArray"),
						[for (value in values) lowerExpr(value).expr]) : GoExpr.GoCompositeLiteral(GoType.slice(arrayElementGoType(expr.t)),
							[for (value in values) GoCompositeElement.GoCompositeValue(lowerExpr(value).expr)]),
					isStringLike: false
				};
			case TObjectDecl(fields):
				materializeExprWithPrefix(lowerObjectDeclExpr(fields), expr.t);
			case TBlock(exprs):
				var restPacked = lowerRestPackBlock(exprs);
				if (restPacked != null) {
					{expr: restPacked, isStringLike: false};
				} else {
					materializeExprWithPrefix(lowerExprWithPrefix(expr), expr.t);
				}
			case TArray(target, index):
				{
					expr: usesSharedArrayCarrier(target) ? GoExpr.GoCall(GoExpr.GoSelector(lowerExpr(target).expr, "Get"),
						[lowerExpr(index).expr]) : GoExpr.GoIndex(lowerExpr(target).expr, lowerExpr(index).expr),
					isStringLike: isStringType(expr.t)
				};
			case TEnumIndex(inner):
				{
					expr: GoExpr.GoSelector(lowerExpr(inner).expr, "tag"),
					isStringLike: false
				};
			case TEnumParameter(target, constructor, index):
				var payload = GoExpr.GoIndex(GoExpr.GoSelector(lowerExpr(target).expr, "params"), GoExpr.GoIntLiteral(index));
				var payloadType = enumPayloadStorageGoType(constructor, index, expr.t);
				{
					expr: payloadType == "any" ? payload : GoExpr.GoTypeAssert(payload, payloadType),
					isStringLike: isStringType(expr.t)
				};
			case TNew(classRef, _, args):
				var classType = classRef.get();
				noteSourceOwnedStdlibUsage(classType);
				var loweredArgs = [
					for (index in 0...args.length)
						lowerConstructorArg(classType, args[index], index)
				];
				var ctorInfo = resolveConstructorInfo(classType);
				if (ctorInfo != null
					&& !GoStdlibOwnership.isCompilerOwnedAuthority(fullClassName(classType))
					&& loweredArgs.length < ctorInfo.defaults.length) {
					for (i in loweredArgs.length...ctorInfo.defaults.length) {
						var defaultValue = ctorInfo.defaults[i];
						if (defaultValue == null) {
							Context.fatalError("Missing required constructor argument at position " + i, expr.pos);
						}
						loweredArgs.push(lowerConstructorArg(classType, defaultValue, i));
					}
				}
				if (isHaxeValueExceptionClass(classType)) {
					while (loweredArgs.length < 3) {
						loweredArgs.push(GoExpr.GoNil);
					}
					{
						expr: GoExpr.GoCall(GoExpr.GoIdent("hxrt.NewValueException"), loweredArgs),
						isStringLike: false
					};
				} else if (useTypedGoConcurrencySpecialization() && isGoChanClass(classType)) {
					noteLoweringAttempt("go.concurrency.typed", "go_chan_new", expr.pos, "Attempt typed go.Chan constructor specialization.");
					var elementEligibility = goChanElementEligibility(expr.t, "Could not resolve go.Chan element type for constructor specialization.");
					if (elementEligibility.eligible) {
						var elementGoType = elementEligibility.goType;
						requireStdlibShimGroup("go_concurrency");
						registerNativeChanElementGoType(elementGoType);
						noteProvenConcurrencyFastpathHit(expr.pos);
						noteLoweringSuccess("go.concurrency.typed", "go_chan_new", expr.pos,
							'Applied typed go.Chan constructor specialization (element type: ' + elementGoType + ").");
						{
							expr: GoExpr.GoCall(GoExpr.GoIdent(nativeChanShimName("go__concurrency_newChan", elementGoType)), [GoExpr.GoIntLiteral(0)]),
							isStringLike: false
						};
					} else {
						noteProvenConcurrencyFastpathFallback(expr.pos);
						noteLoweringFallback("go.concurrency.typed", "go_chan_new_unmorphable", expr.pos,
							withEligibilityReason("Could not monomorphize go.Chan element type for constructor specialization.", elementEligibility));
						{
							expr: GoExpr.GoCall(GoExpr.GoIdent(constructorSymbol(classType)), loweredArgs),
							isStringLike: false
						};
					}
				} else if (classType.pack.length == 0 && classType.name == "Array") {
					{
						expr: GoExpr.GoCall(GoExpr.GoIdent("hxrt.NewArray"), []),
						isStringLike: false
					};
				} else {
					{
						expr: GoExpr.GoCall(GoExpr.GoIdent(constructorSymbol(classType)), loweredArgs),
						isStringLike: false
					};
				}
			case TFunction(func):
				pushFunctionVarNameScope();
				var loweredParams = lowerFunctionParams(func);
				var loweredResults = lowerFunctionResults(func.t);
				pushFunctionReturnType(func.t);
				var loweredBody = lowerFunctionBody(func.expr);
				popFunctionReturnType();
				popFunctionVarNameScope();
				{
					expr: GoExpr.GoFuncLiteral(loweredParams, loweredResults, loweredBody),
					isStringLike: false
				};
			case TLocal(variable):
				var localExpr:GoExpr = GoExpr.GoIdent(localVarName(variable));
				var narrowedStorageGoType = registeredNarrowedPrimitiveStorageGoType(variable);
				var arrayStorageOverride = localArrayStorageOverrides.exists(variable.id) ? localArrayStorageOverrides.get(variable.id) : null;
				var variableGoType = arrayStorageOverride != null ? typeToGoType(arrayStorageOverride) : (narrowedStorageGoType == null ? valueStorageGoType(variable.t) : narrowedStorageGoType);
				var exprGoType = valueStorageGoType(expr.t);
				var optionalPrimitiveGoType = registeredOptionalPrimitiveParamGoType(variable);
				var nonNullPrimitiveGoType = registeredNonNullPrimitiveLocalGoType(variable);
				if (narrowedStorageGoType != null) {
					localExpr = GoExpr.GoIdent(localVarName(variable));
				} else if (nonNullPrimitiveGoType != null && variableGoType == "any") {
					localExpr = GoExpr.GoTypeAssert(localExpr, nonNullPrimitiveGoType);
				} else if (optionalPrimitiveGoType != null) {
					localExpr = GoExpr.GoTypeAssert(localExpr, optionalPrimitiveGoType);
				} else if (variableGoType == "any" && exprGoType != "any") {
					localExpr = GoExpr.GoTypeAssert(localExpr, exprGoType);
				}
				{
					expr: localExpr,
					isStringLike: isStringType(expr.t)
				};
			case TIdent(name):
				{
					expr: GoExpr.GoIdent(name),
					isStringLike: isStringType(expr.t)
				};
			case TParenthesis(inner):
				lowerExpr(inner);
			case TMeta(_, inner):
				lowerExpr(inner);
			case TCast(inner, _):
				var loweredInner = lowerExpr(inner);
				var innerGoType = typeToGoType(inner.t);
				var castGoType = typeToGoType(expr.t);
				var castExpr = loweredInner.expr;
				// Array-like casts inside inline abstracts expose their underlying storage.
				// Convert only when an assignment, call, or return supplies a real expected
				// type; converting here turns every Vector index/length access into a copy.
				var storageTransparentArrayCast = isArrayType(inner.t) && isArrayType(expr.t);
				if (!storageTransparentArrayCast && innerGoType != castGoType) {
					if (castGoType != "any" && innerGoType == "any") {
						castExpr = lowerNullableAwareTypeAssertExpr(castExpr, expr.t);
					} else if (castGoType == "any" && innerGoType != "any") {
						castExpr = GoExpr.GoCall(GoExpr.GoIdent("any"), [castExpr]);
					} else if (isInterfaceType(inner.t) && !isInterfaceType(expr.t)) {
						// What: retain Haxe's proven interface-to-concrete nominal cast in Go.
						// Why: multi-type abstracts such as Map store IMap but inline calls through
						// their selected concrete implementation; dropping the cast leaves an
						// interface receiver with a concrete-only method selector.
						// How: use the ordinary Go type assertion for the concrete target type.
						castExpr = GoExpr.GoTypeAssert(castExpr, castGoType);
					}
				}
				{
					expr: castExpr,
					isStringLike: isStringType(expr.t)
				};
			case TIf(condition, thenBranch, elseBranch):
				materializeExprWithPrefix(lowerIfExpr(condition, thenBranch, elseBranch, expr.t), expr.t);
			case TSwitch(value, cases, defaultExpr):
				materializeExprWithPrefix(lowerSwitchExpr(value, cases, defaultExpr, expr.t), expr.t);
			case TTry(tryExpr, catches):
				materializeExprWithPrefix(lowerTryCatchExpr(tryExpr, catches, expr.t), expr.t);
			case TField(target, access):
				lowerField(target, access);
			case TCall(callee, args):
				var injected = lowerTargetCodeInjectionExpr(expr);
				injected != null ? injected : lowerCall(callee, args, expr.t);
			case TThrow(value):
				lowerThrowExprForType(value, expr.t);
			case TTypeExpr(moduleType):
				lowerTypeExpr(moduleType);
			case TBinop(op, left, right):
				switch (op) {
					case OpAssign:
						var indexedArrayAssign = lowerHaxeArrayIndexAssignExpr(left, right);
						if (indexedArrayAssign != null) {
							indexedArrayAssign;
						} else {
							var targetExpr = lowerLValue(left);
							var loweredRight = lowerExprWithExpectedUpcast(right, assignmentStorageType(left));
							var rightExpr = loweredRight.expr;
							{
								expr: GoExpr.GoCall(GoExpr.GoFuncLiteral([], [typeToGoType(left.t)],
									loweredRight.prefix.concat([GoStmt.GoAssign(targetExpr, rightExpr), GoStmt.GoReturn(targetExpr)])),
									[]),
								isStringLike: isStringType(left.t)
							};
						}
					case OpAssignOp(assignOp):
						var sharedArrayAssign = lowerHaxeArrayIndexAssignOpExpr(left, right, assignOp, expr.pos);
						if (sharedArrayAssign != null) {
							sharedArrayAssign;
						} else {
							var targetExpr = lowerLValue(left);
							var loweredRight = lowerExprWithPrefix(right);
							var stringAppendFromSharedArray = assignOp == OpAdd
								&& (isStringType(left.t) || isStringType(right.t))
								&& isSharedArrayElementExpr(right);
							var rightExpr = stringAppendFromSharedArray ? lowerSharedArrayElementStorageExpr(right) : upcastIfNeeded(loweredRight.expr,
								right.t, left.t, right);
							if (!stringAppendFromSharedArray && isSharedArrayElementExpr(right)) {
								rightExpr = coerceStoredArrayElementExpr(rightExpr, left.t);
							}
							var assignExpr = lowerAssignOpExpr(assignOp, targetExpr, rightExpr, left.t, right.t, expr.pos, stringAppendFromSharedArray);
							{
								expr: GoExpr.GoCall(GoExpr.GoFuncLiteral([], [typeToGoType(left.t)],
									loweredRight.prefix.concat([GoStmt.GoAssign(targetExpr, assignExpr), GoStmt.GoReturn(targetExpr)])),
									[]),
								isStringLike: isStringType(left.t)
							};
						}
					case _:
						lowerBinop(op, left, right, expr.t);
				}
			case TUnop(op, postFix, value):
				var sharedArrayUnit = (op == OpIncrement || op == OpDecrement) ? lowerHaxeArrayIndexUnitExpr(value, op, postFix, expr.pos) : null;
				if (sharedArrayUnit != null) {
					return sharedArrayUnit;
				}
				if (postFix) {
					return switch (op) {
						case OpIncrement, OpDecrement:
							var target = lowerLValue(value);
							var temp = freshTempName("hx_post");
							var opSymbol = op == OpIncrement ? GoBinaryOperator.Add : GoBinaryOperator.Subtract;
							{
								expr: GoExpr.GoCall(GoExpr.GoFuncLiteral([], [typeToGoType(value.t)], [
									GoStmt.GoVarDecl(temp, null, target, true),
									GoStmt.GoAssign(target, unitStepExpr(target, opSymbol, value.t, expr.pos)),
									GoStmt.GoReturn(GoExpr.GoIdent(temp))
								]), []),
								isStringLike: isStringType(expr.t)
							};
						case _:
							unsupportedExpr(expr, "Unsupported postfix unary operator");
					};
				}
				return switch (op) {
					case OpIncrement, OpDecrement:
						var target = lowerLValue(value);
						var opSymbol = op == OpIncrement ? GoBinaryOperator.Add : GoBinaryOperator.Subtract;
						{
							expr: GoExpr.GoCall(GoExpr.GoFuncLiteral([], [typeToGoType(value.t)], [
								GoStmt.GoAssign(target, unitStepExpr(target, opSymbol, value.t, expr.pos)),
								GoStmt.GoReturn(target)
							]), []),
							isStringLike: isStringType(expr.t)
						};
					case _:
						var loweredValue = lowerExpr(value).expr;
						var unaryExpr = if (isInt32SemanticType(expr.t, expr.pos) && (op == OpNeg || op == OpNegBits)) {
							wrapInt32Expr(GoExpr.GoUnary(unopSymbol(op), GoExpr.GoCall(GoExpr.GoIdent("int32"), [loweredValue])));
						} else {
							GoExpr.GoUnary(unopSymbol(op), loweredValue);
						}
						{
							expr: unaryExpr,
							isStringLike: isStringType(expr.t)
						};
				};
			case _:
				unsupportedExpr(expr, "Unsupported expression");
		};
	}

	/**
		What
		Lower the framework raw Go escape hatch `__go__`.

		Why
		`CompilerInit` already declares `targetCodeInjectionName: "__go__"`, and sibling
		backends (`reflaxe.rust`, `reflaxe.ocaml`) treat raw target injection as an explicit
		backend responsibility. Before this hook existed, `haxe.go` policy scanners could
		recognize `__go__`, but the manual lowering pipeline never turned it into raw Go,
		so allowed callsites still emitted unresolved `__go__(...)` into generated output.

		How
		Recognize typed `__go__` calls, require a constant first string argument, and expand
		`{0}`, `{1}`, ... placeholders by lowering only the referenced Haxe expressions and
		printing them with `GoASTPrinter.printExprForInjection`. If the template has no
		placeholders, treat it as literal raw Go text.
	**/
	function lowerTargetCodeInjectionExpr(expr:TypedExpr):Null<LoweredExpr> {
		if (!GoProfileContractAnalyzer.isGoInjectionCall(expr)) {
			return null;
		}

		var arguments = switch (expr.expr) {
			case TCall(_, args): args;
			case _: null;
		};
		if (arguments == null) {
			return null;
		}

		if (arguments.length == 0) {
			Context.error("__go__ requires at least one constant String argument.", expr.pos);
			return {
				expr: GoExpr.GoNil,
				isStringLike: isStringType(expr.t)
			};
		}

		var injectionString = switch (arguments[0].expr) {
			case TConst(TString(value)): value;
			case _:
				Context.error("__go__ first parameter must be a constant String.", arguments[0].pos);
				null;
		};
		if (injectionString == null) {
			return {
				expr: GoExpr.GoNil,
				isStringLike: isStringType(expr.t)
			};
		}

		var cachedArgs:Array<Null<GoExpr>> = [for (_ in 1...arguments.length) null];
		function getRenderedArg(index:Int):Null<String> {
			if (index < 0 || index >= cachedArgs.length) {
				return null;
			}
			if (cachedArgs[index] == null) {
				cachedArgs[index] = lowerExpr(arguments[index + 1]).expr;
			}
			return GoASTPrinter.printExprForInjection(cachedArgs[index]);
		}

		var rendered = new StringBuf();
		var lastMatchPosition:Null<{pos:Int, len:Int}> = null;
		~/{(\d+)}/g.map(injectionString, function(ereg) {
			var lastPos = lastMatchPosition == null ? 0 : lastMatchPosition.pos + lastMatchPosition.len;
			lastMatchPosition = ereg.matchedPos();
			if (lastMatchPosition.pos != lastPos) {
				rendered.add(injectionString.substring(lastPos, lastMatchPosition.pos));
			}

			var expressionIndex = Std.parseInt(ereg.matched(1));
			if (expressionIndex != null) {
				var compiled = getRenderedArg(expressionIndex);
				if (compiled != null) {
					rendered.add(compiled);
				}
			}
			return "";
		});

		if (lastMatchPosition == null) {
			rendered.add(injectionString);
		} else {
			rendered.add(injectionString.substring(lastMatchPosition.pos + lastMatchPosition.len));
		}

		return {
			expr: GoExpr.GoRaw(rendered.toString()),
			isStringLike: isStringType(expr.t)
		};
	}

	function lowerTargetCodeInjectionCall(callee:TypedExpr, args:Array<TypedExpr>, returnType:Type):Null<LoweredExpr> {
		if (!isTargetCodeInjectionCallee(callee)) {
			return null;
		}

		if (args.length == 0) {
			Context.error("__go__ requires at least one constant String argument.", callee.pos);
			return {
				expr: GoExpr.GoNil,
				isStringLike: isStringType(returnType)
			};
		}

		var injectionString = switch (args[0].expr) {
			case TConst(TString(value)): value;
			case _:
				Context.error("__go__ first parameter must be a constant String.", args[0].pos);
				null;
		};
		if (injectionString == null) {
			return {
				expr: GoExpr.GoNil,
				isStringLike: isStringType(returnType)
			};
		}

		var cachedArgs:Array<Null<GoExpr>> = [for (_ in 1...args.length) null];
		function getRenderedArg(index:Int):Null<String> {
			if (index < 0 || index >= cachedArgs.length) {
				return null;
			}
			if (cachedArgs[index] == null) {
				cachedArgs[index] = lowerExpr(args[index + 1]).expr;
			}
			return GoASTPrinter.printExprForInjection(cachedArgs[index]);
		}

		var rendered = new StringBuf();
		var lastMatchPosition:Null<{pos:Int, len:Int}> = null;
		~/{(\d+)}/g.map(injectionString, function(ereg) {
			var lastPos = lastMatchPosition == null ? 0 : lastMatchPosition.pos + lastMatchPosition.len;
			lastMatchPosition = ereg.matchedPos();
			if (lastMatchPosition.pos != lastPos) {
				rendered.add(injectionString.substring(lastPos, lastMatchPosition.pos));
			}

			var expressionIndex = Std.parseInt(ereg.matched(1));
			if (expressionIndex != null) {
				var compiled = getRenderedArg(expressionIndex);
				if (compiled != null) {
					rendered.add(compiled);
				}
			}
			return "";
		});

		if (lastMatchPosition == null) {
			rendered.add(injectionString);
		} else {
			rendered.add(injectionString.substring(lastMatchPosition.pos + lastMatchPosition.len));
		}

		return {
			expr: GoExpr.GoRaw(rendered.toString()),
			isStringLike: isStringType(returnType)
		};
	}

	function isTargetCodeInjectionCallee(callee:TypedExpr):Bool {
		return switch (callee.expr) {
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
		};
	}

	function materializeExprWithPrefix(lowered:LoweredExprWithPrefix, resultType:Type):LoweredExpr {
		if (lowered.prefix.length == 0) {
			return {
				expr: lowered.expr,
				isStringLike: lowered.isStringLike
			};
		}

		if (isVoidType(resultType)) {
			return {
				expr: GoExpr.GoCall(GoExpr.GoFuncLiteral([], [], lowered.prefix), []),
				isStringLike: false
			};
		}

		return {
			expr: GoExpr.GoCall(GoExpr.GoFuncLiteral([], [valueStorageGoType(resultType)], lowered.prefix.concat([GoStmt.GoReturn(lowered.expr)])), []),
			isStringLike: lowered.isStringLike
		};
	}

	function lowerExprWithPrefix(expr:TypedExpr):LoweredExprWithPrefix {
		return switch (expr.expr) {
			case TBlock(exprs):
				var storageOverride = vectorBlockStorageOverride(expr.t, exprs);
				if (storageOverride != null) {
					localArrayStorageOverrides.set(storageOverride.variable.id, storageOverride.storageType);
				}
				var result:LoweredExprWithPrefix = if (exprs.length == 0) {
					{prefix: [], expr: GoExpr.GoNil, isStringLike: false};
				} else {
					var prefix = new Array<GoStmt>();
					for (index in 0...exprs.length - 1) {
						prefix = prefix.concat(withoutThrowFallback(function() return lowerToStatements(exprs[index])));
					}
					var tail = lowerExprWithPrefix(exprs[exprs.length - 1]);
					{
						prefix: prefix.concat(tail.prefix),
						expr: tail.expr,
						isStringLike: tail.isStringLike
					};
				};
				if (storageOverride != null) {
					localArrayStorageOverrides.remove(storageOverride.variable.id);
				}
				result;
			case TSwitch(value, cases, defaultExpr):
				lowerSwitchExpr(value, cases, defaultExpr, expr.t);
			case TIf(condition, thenBranch, elseBranch):
				lowerIfExpr(condition, thenBranch, elseBranch, expr.t);
			case TTry(tryExpr, catches):
				lowerTryCatchExpr(tryExpr, catches, expr.t);
			case TObjectDecl(fields):
				lowerObjectDeclExpr(fields);
			case TArray(target, index):
				var loweredTarget = lowerExprWithPrefix(target);
				var loweredIndex = lowerExprWithPrefix(index);
				{
					prefix: loweredTarget.prefix.concat(loweredIndex.prefix),
					expr: usesSharedArrayCarrier(target) ? GoExpr.GoCall(GoExpr.GoSelector(loweredTarget.expr, "Get"),
						[loweredIndex.expr]) : GoExpr.GoIndex(loweredTarget.expr, loweredIndex.expr),
					isStringLike: isStringType(expr.t)
				};
			case TUnop(op, postFix, value):
				var sharedArrayUnit = (op == OpIncrement || op == OpDecrement) ? lowerHaxeArrayIndexUnitExpr(value, op, postFix, expr.pos) : null;
				if (sharedArrayUnit != null) {
					{
						prefix: [],
						expr: sharedArrayUnit.expr,
						isStringLike: sharedArrayUnit.isStringLike
					};
				} else if (postFix) {
					switch (op) {
						case OpIncrement, OpDecrement:
							var target = lowerLValue(value);
							var temp = freshTempName("hx_post");
							var opSymbol = op == OpIncrement ? GoBinaryOperator.Add : GoBinaryOperator.Subtract;
							{
								prefix: [
									GoStmt.GoVarDecl(temp, null, target, true),
									GoStmt.GoAssign(target, unitStepExpr(target, opSymbol, value.t, expr.pos))
								],
								expr: GoExpr.GoIdent(temp),
								isStringLike: false
							};
						case _:
							Context.fatalError("Unsupported postfix unary operator :: " + Std.string(expr.expr), expr.pos);
							{
								prefix: [],
								expr: GoExpr.GoNil,
								isStringLike: false
							};
					}
				} else {
					var lowered = lowerExpr(expr);
					{
						prefix: [],
						expr: lowered.expr,
						isStringLike: lowered.isStringLike
					};
				}
			case TMeta(_, inner):
				lowerExprWithPrefix(inner);
			case TParenthesis(inner):
				lowerExprWithPrefix(inner);
			case TCast(inner, _):
				var lowered = lowerExpr(expr);
				{
					prefix: [],
					expr: lowered.expr,
					isStringLike: lowered.isStringLike
				};
			case _:
				var lowered = lowerExpr(expr);
				{
					prefix: [],
					expr: lowered.expr,
					isStringLike: lowered.isStringLike
				};
		};
	}

	function lowerTypeExpr(moduleType:ModuleType):LoweredExpr {
		requiresTypeValueSupport = true;
		var markerType = moduleTypeRepresentsEnumValue(moduleType) ? "hxrt__TypeEnumValue" : "hxrt__TypeClassValue";
		return {
			expr: GoExpr.GoUnary(GoUnaryOperator.AddressOf, GoExpr.GoCompositeLiteral(GoType.named(markerType), [
				GoCompositeElement.GoCompositeField("name",
					GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoStringLiteral(moduleTypeDisplayName(moduleType))]))
			])),
			isStringLike: false
		};
	}

	function lowerTestAstStmtDecls():Array<GoDecl> {
		var testCase = Context.definedValue("reflaxe_go_test_ast_stmt_case");
		if (testCase == null || testCase == "") {
			return [];
		}

		var emitted = GoTestAstFixtureEmitter.emit(testCase);
		if (emitted == null) {
			Context.fatalError('Unknown AST statement test case "' + testCase + '"', Context.currentPos());
			return [];
		}

		return emitted;
	}

	function moduleTypeRepresentsEnumValue(moduleType:ModuleType):Bool {
		return switch (moduleType) {
			case TEnumDecl(_):
				true;
			case TTypeDecl(typeRef):
				switch (Context.follow(TType(typeRef, []))) {
					case TEnum(_, _):
						true;
					case _:
						false;
				}
			case _:
				false;
		};
	}

	function moduleTypeDisplayName(moduleType:ModuleType):String {
		return switch (moduleType) {
			case TClassDecl(classRef):
				fullClassName(classRef.get());
			case TEnumDecl(enumRef):
				fullEnumName(enumRef.get());
			case TTypeDecl(typeRef):
				var typeDef = typeRef.get();
				typeDef.pack.length == 0 ? typeDef.name : typeDef.pack.join(".") + "." + typeDef.name;
			case TAbstract(abstractRef):
				var abstractType = abstractRef.get();
				abstractType.pack.length == 0 ? abstractType.name : abstractType.pack.join(".") + "." + abstractType.name;
		};
	}

	function lowerConst(constant:TConstant):LoweredExpr {
		return switch (constant) {
			case TNull:
				{expr: GoExpr.GoNil, isStringLike: true};
			case TInt(value):
				{expr: GoExpr.GoIntLiteral(value), isStringLike: false};
			case TFloat(value):
				{expr: GoExpr.GoFloatLiteral(value), isStringLike: false};
			case TString(value):
				{
					expr: GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoStringLiteral(value)]),
					isStringLike: true
				};
			case TBool(value):
				{expr: GoExpr.GoBoolLiteral(value), isStringLike: false};
			case TThis:
				{expr: GoExpr.GoIdent("self"), isStringLike: false};
			case TSuper:
				{expr: GoExpr.GoIdent("self"), isStringLike: false};
		};
	}

	function lowerField(target:TypedExpr, access:FieldAccess):LoweredExpr {
		return switch (access) {
			case FStatic(classRef, field):
				var resolved = field.get();
				var classType = classRef.get();
				noteStaticStdlibFieldUsage(classType, resolved.name, resolved.pos);
				if (classType.isExtern) {
					var externPackage = externClassPackageName(classType);
					if (externPackage != null) {
						noteExternImportPath(classType, externPackage);
						return {
							expr: GoExpr.GoSelector(GoExpr.GoIdent(externPackage), externFieldName(resolved)),
							isStringLike: isStringType(resolved.type)
						};
					}
				}
				{
					expr: GoExpr.GoIdent(staticSymbol(classType, resolved.name)),
					isStringLike: isStringType(resolved.type)
				};
			case FInstance(classRef, _, field):
				var resolved = field.get();
				var classType = classRef.get();
				noteSourceOwnedStdlibUsage(classType);
				var loweredTarget = lowerExpr(target).expr;
				if (isSharedArrayElementExpr(target)) {
					loweredTarget = coerceStoredArrayElementExpr(loweredTarget, target.t);
				}
				var staticInterfaceSelector = interfaceSelectorForStaticReceiver(target.t, resolved.name);

				if (isSuperTarget(target) && isMethodField(resolved)) {
					var baseSelector = GoExpr.GoSelector(GoExpr.GoIdent("self"), classTypeName(classType));
					return {
						expr: GoExpr.GoSelector(baseSelector, normalizeIdent(resolved.name)),
						isStringLike: isStringType(resolved.type)
					};
				}

				if (staticInterfaceSelector != null) {
					return {
						expr: GoExpr.GoSelector(loweredTarget, staticInterfaceSelector),
						isStringLike: isStringType(resolved.type)
					};
				}

				if (classType.isInterface) {
					return {
						expr: GoExpr.GoSelector(loweredTarget, interfaceFieldName(classType, resolved)),
						isStringLike: isStringType(resolved.type)
					};
				}

				if (classType.pack.length == 0 && classType.name == "String" && resolved.name == "length") {
					var lengthHelper = if (useStringFastpath()) {
						compilationContext.optimizerStringLengthFieldTypedLowerings++;
						"hxrt.StringLengthStringPtr";
					} else {
						compilationContext.optimizerStringLengthFieldLegacyLowerings++;
						"hxrt.StringLength";
					};
					return {
						expr: GoExpr.GoCall(GoExpr.GoIdent(lengthHelper), [loweredTarget]),
						isStringLike: false
					};
				}

				if (isHaxeExceptionFamilyType(target.t) && resolved.name == "message") {
					return {
						expr: GoExpr.GoCall(GoExpr.GoIdent("hxrt.ExceptionMessage"), [loweredTarget]),
						isStringLike: true
					};
				}

				if (isHaxeValueExceptionType(target.t) && resolved.name == "value") {
					return {
						expr: GoExpr.GoSelector(loweredTarget, "Value"),
						isStringLike: isStringType(resolved.type)
					};
				}

				var restTargetName = resolveRestIteratorTargetName(target);
				var restFieldName = restIteratorFieldName(restTargetName, resolved.name);
				if (restFieldName != null) {
					return {
						expr: GoExpr.GoIdent(restFieldName),
						isStringLike: isStringType(resolved.type)
					};
				}
				if (resolved.name == "length" && isArrayType(target.t)) {
					{
						expr: usesSharedArrayCarrier(target) ? GoExpr.GoCall(GoExpr.GoSelector(loweredTarget, "Len"),
							[]) : GoExpr.GoCall(GoExpr.GoIdent("len"), [loweredTarget]),
						isStringLike: false
					};
				} else if (classType.isInterface) {
					{
						expr: GoExpr.GoSelector(loweredTarget, interfaceFieldName(classType, resolved)),
						isStringLike: isStringType(resolved.type)
					};
				} else if (classType.isExtern) {
					var externPackage = externClassPackageName(classType);
					if (externPackage != null) {
						noteExternImportPath(classType, externPackage);
					}
					{
						expr: GoExpr.GoSelector(loweredTarget, externFieldName(resolved)),
						isStringLike: isStringType(resolved.type)
					};
				} else if (shouldUseVirtualDispatch(classType, resolved, target.t)) {
					{
						expr: GoExpr.GoSelector(GoExpr.GoSelector(loweredTarget, "__hx_this"), normalizeIdent(resolved.name)),
						isStringLike: isStringType(resolved.type)
					};
				} else {
					{
						expr: GoExpr.GoSelector(loweredTarget, normalizeIdent(resolved.name)),
						isStringLike: isStringType(resolved.type)
					};
				}
			case FAnon(field):
				var resolved = field.get();
				var loweredTarget = lowerExpr(target).expr;
				if (isSharedArrayElementExpr(target)) {
					loweredTarget = coerceStoredArrayElementExpr(loweredTarget, target.t);
				}
				if (isAnonymousObjectType(target.t)) {
					return {
						expr: lowerAnonymousFieldRead(loweredTarget, resolved.name, resolved.type),
						isStringLike: isStringType(resolved.type)
					};
				}
				if (isHaxeExceptionFamilyType(target.t) && resolved.name == "message") {
					return {
						expr: GoExpr.GoCall(GoExpr.GoIdent("hxrt.ExceptionMessage"), [loweredTarget]),
						isStringLike: true
					};
				}
				if (isHaxeValueExceptionType(target.t) && resolved.name == "value") {
					return {
						expr: GoExpr.GoSelector(loweredTarget, "Value"),
						isStringLike: isStringType(resolved.type)
					};
				}
				var restTargetName = resolveRestIteratorTargetName(target);
				var restFieldName = restIteratorFieldName(restTargetName, resolved.name);
				if (restFieldName != null) {
					return {
						expr: GoExpr.GoIdent(restFieldName),
						isStringLike: isStringType(resolved.type)
					};
				}
				if (resolved.name == "length" && isArrayType(target.t)) {
					{
						expr: usesSharedArrayCarrier(target) ? GoExpr.GoCall(GoExpr.GoSelector(loweredTarget, "Len"),
							[]) : GoExpr.GoCall(GoExpr.GoIdent("len"), [loweredTarget]),
						isStringLike: false
					};
				} else {
					{
						expr: GoExpr.GoSelector(loweredTarget, normalizeIdent(resolved.name)),
						isStringLike: isStringType(resolved.type)
					};
				}
			case FDynamic(name):
				var loweredTarget = lowerExpr(target).expr;
				if (isSharedArrayElementExpr(target) && typeToGoType(target.t) != "any") {
					loweredTarget = coerceStoredArrayElementExpr(loweredTarget, target.t);
				}
				if (isAnonymousObjectType(target.t)) {
					return {
						expr: GoExpr.GoIndex(loweredTarget, GoExpr.GoStringLiteral(name)),
						isStringLike: false
					};
				}
				var dynamicExpr = if (name == "length" && isArrayType(target.t)) {
					usesSharedArrayCarrier(target) ? GoExpr.GoCall(GoExpr.GoSelector(loweredTarget, "Len"),
						[]) : GoExpr.GoCall(GoExpr.GoIdent("len"), [loweredTarget]);
				} else {
					GoExpr.GoSelector(loweredTarget, normalizeIdent(name));
				};
				{
					expr: dynamicExpr,
					isStringLike: false
				};
			case FClosure(_, field):
				var resolved = field.get();
				var loweredTarget = lowerExpr(target).expr;
				if (isSharedArrayElementExpr(target)) {
					loweredTarget = coerceStoredArrayElementExpr(loweredTarget, target.t);
				}
				{
					expr: GoExpr.GoSelector(loweredTarget, normalizeIdent(resolved.name)),
					isStringLike: isStringType(resolved.type)
				};
			case FEnum(enumRef, field):
				{
					expr: GoExpr.GoIdent(enumConstructorSymbol(enumRef.get(), field.name)),
					isStringLike: false
				};
		};
	}

	function lowerAnonymousFieldRead(targetExpr:GoExpr, fieldName:String, fieldType:Type):GoExpr {
		var fieldExpr = GoExpr.GoIndex(targetExpr, GoExpr.GoStringLiteral(fieldName));
		var fieldGoType = scalarGoType(fieldType);
		if (fieldGoType == "any") {
			return fieldExpr;
		}

		var objectName = freshTempName("hx_obj");
		var valueName = freshTempName("hx_field");
		var zeroName = freshTempName("hx_zero");

		return GoExpr.GoCall(GoExpr.GoFuncLiteral([{name: objectName, typeName: "map[string]any"}], [fieldGoType], [
			GoStmt.GoVarDecl(valueName, "any", GoExpr.GoIndex(GoExpr.GoIdent(objectName), GoExpr.GoStringLiteral(fieldName)), true),
			GoStmt.GoIf(GoExpr.GoBinary("==", GoExpr.GoIdent(valueName), GoExpr.GoNil), [
				GoStmt.GoVarDecl(zeroName, fieldGoType, null, false),
				GoStmt.GoReturn(GoExpr.GoIdent(zeroName))
			], null),
			GoStmt.GoReturn(GoExpr.GoTypeAssert(GoExpr.GoIdent(valueName), fieldGoType))
		]), [targetExpr]);
	}

	function lowerNilSafeTypeAssertExpr(expr:GoExpr, assertedType:String):GoExpr {
		var valueName = freshTempName("hx_value");
		var zeroName = freshTempName("hx_zero");
		return GoExpr.GoCall(GoExpr.GoFuncLiteral([{name: valueName, typeName: "any"}], [assertedType], [
			GoStmt.GoIf(GoExpr.GoBinary("==", GoExpr.GoIdent(valueName), GoExpr.GoNil), [
				GoStmt.GoVarDecl(zeroName, assertedType, null, false),
				GoStmt.GoReturn(GoExpr.GoIdent(zeroName))
			], null),
			GoStmt.GoReturn(GoExpr.GoTypeAssert(GoExpr.GoIdent(valueName), assertedType))
		]), [expr]);
	}

	function lowerNilPreservingTypeAssertExpr(expr:GoExpr, assertedType:String):GoExpr {
		var valueName = freshTempName("hx_value");
		return GoExpr.GoCall(GoExpr.GoFuncLiteral([{name: valueName, typeName: "any"}], ["any"], [
			GoStmt.GoIf(GoExpr.GoBinary("==", GoExpr.GoIdent(valueName), GoExpr.GoNil), [GoStmt.GoReturn(GoExpr.GoNil)], null),
			GoStmt.GoReturn(GoExpr.GoTypeAssert(GoExpr.GoIdent(valueName), assertedType))
		]), [expr]);
	}

	function lowerNullableAwareTypeAssertExpr(expr:GoExpr, returnType:Type):GoExpr {
		var expectedType = typeToGoType(returnType);
		if (expectedType == "any") {
			return expr;
		}
		if (shouldPreserveNilInTypeAssert(returnType, expectedType)) {
			return lowerNilPreservingTypeAssertExpr(expr, expectedType);
		}
		return lowerNilSafeTypeAssertExpr(expr, expectedType);
	}

	function shouldPreserveNilInTypeAssert(returnType:Type, expectedType:String):Bool {
		if (!isNullablePrimitiveType(returnType)) {
			return false;
		}
		return expectedType == "int" || expectedType == "float64" || expectedType == "bool";
	}

	function lowerCall(callee:TypedExpr, args:Array<TypedExpr>, returnType:Type):LoweredExpr {
		var injectedCall = lowerTargetCodeInjectionCall(callee, args, returnType);
		if (injectedCall != null) {
			return injectedCall;
		}

		if (isStaticCall(callee, "Reflect", [], "fields")) {
			requireStdlibShimGroup("stdlib_symbols");
			requiresReflectFieldsShim = true;
		}

		var nativeSliceBoundaryCall = lowerNativeSliceBoundaryCall(callee, args, returnType);
		if (nativeSliceBoundaryCall != null) {
			return nativeSliceBoundaryCall;
		}

		var arrayInstanceCall = lowerArrayInstanceCall(callee, args, returnType);
		if (arrayInstanceCall != null) {
			return arrayInstanceCall;
		}

		var restAbstractCall = lowerRestAbstractCall(callee, args, returnType);
		if (restAbstractCall != null) {
			return restAbstractCall;
		}

		var stringInstanceCall = lowerStringInstanceCall(callee, args, returnType);
		if (stringInstanceCall != null) {
			return stringInstanceCall;
		}

		var externReceiverCall = lowerExternReceiverCall(callee, args, returnType);
		if (externReceiverCall != null) {
			return externReceiverCall;
		}

		var lambdaSourceCall = lowerLambdaSourceCallAdapter(callee, args, returnType);
		if (lambdaSourceCall != null) {
			return lambdaSourceCall;
		}

		var lambdaFunctionValueCall = lowerLambdaFunctionValueCall(callee, args, returnType);
		if (lambdaFunctionValueCall != null) {
			return lambdaFunctionValueCall;
		}

		var dsSortHelperCall = lowerDsSortHelperCall(callee, args, returnType);
		if (dsSortHelperCall != null) {
			return dsSortHelperCall;
		}

		var nativeChanCall = lowerTypedGoChanCall(callee, args, returnType);
		if (nativeChanCall != null) {
			return nativeChanCall;
		}

		var nativeSliceCall = lowerTypedGoSliceCall(callee, args, returnType);
		if (nativeSliceCall != null) {
			return nativeSliceCall;
		}

		var nativeMapCall = lowerTypedGoMapCall(callee, args, returnType);
		if (nativeMapCall != null) {
			return nativeMapCall;
		}

		var nativeResultCall = lowerTypedGoResultCall(callee, args, returnType);
		if (nativeResultCall != null) {
			return nativeResultCall;
		}

		if (isStaticCall(callee, "Go", ["go"], "__chanMake")) {
			requireStdlibShimGroup("go_concurrency");
			var buffer = args.length > 0 ? lowerExpr(args[0]).expr : GoExpr.GoIntLiteral(0);
			return {
				expr: GoExpr.GoCall(GoExpr.GoIdent("go__concurrency_makeChan"), [buffer]),
				isStringLike: false
			};
		}

		if (isStaticCall(callee, "Go", ["go"], "__chanSend")) {
			requireStdlibShimGroup("go_concurrency");
			var channel = args.length > 0 ? lowerExpr(args[0]).expr : GoExpr.GoNil;
			var value = args.length > 1 ? lowerExpr(args[1]).expr : GoExpr.GoNil;
			return {
				expr: GoExpr.GoCall(GoExpr.GoIdent("go__concurrency_send"), [channel, value]),
				isStringLike: false
			};
		}

		if (isStaticCall(callee, "Go", ["go"], "__chanRecv")) {
			requireStdlibShimGroup("go_concurrency");
			var channel = args.length > 0 ? lowerExpr(args[0]).expr : GoExpr.GoNil;
			var rawRecv = GoExpr.GoCall(GoExpr.GoIdent("go__concurrency_recv"), [channel]);
			return {
				expr: lowerNullableAwareTypeAssertExpr(rawRecv, returnType),
				isStringLike: isStringType(returnType)
			};
		}

		if (isStaticCall(callee, "Go", ["go"], "__chanTrySend")) {
			requireStdlibShimGroup("go_concurrency");
			var channel = args.length > 0 ? lowerExpr(args[0]).expr : GoExpr.GoNil;
			var value = args.length > 1 ? lowerExpr(args[1]).expr : GoExpr.GoNil;
			return {
				expr: GoExpr.GoCall(GoExpr.GoIdent("go__concurrency_trySend"), [channel, value]),
				isStringLike: false
			};
		}

		if (isStaticCall(callee, "Go", ["go"], "__chanRecvOr")) {
			requireStdlibShimGroup("go_concurrency");
			var channel = args.length > 0 ? lowerExpr(args[0]).expr : GoExpr.GoNil;
			var defaultValue = args.length > 1 ? lowerExpr(args[1]).expr : GoExpr.GoNil;
			var rawRecvOr = GoExpr.GoCall(GoExpr.GoIdent("go__concurrency_recvOr"), [channel, defaultValue]);
			return {
				expr: lowerNullableAwareTypeAssertExpr(rawRecvOr, returnType),
				isStringLike: isStringType(returnType)
			};
		}

		if (isStaticCall(callee, "Go", ["go"], "__chanTryRecv")) {
			requireStdlibShimGroup("go_concurrency");
			var channel = args.length > 0 ? lowerExpr(args[0]).expr : GoExpr.GoNil;
			return {
				expr: GoExpr.GoCall(GoExpr.GoIdent("go__concurrency_tryRecv"), [channel]),
				isStringLike: false
			};
		}

		if (isStaticCall(callee, "Go", ["go"], "__chanClose")) {
			requireStdlibShimGroup("go_concurrency");
			var channel = args.length > 0 ? lowerExpr(args[0]).expr : GoExpr.GoNil;
			return {
				expr: GoExpr.GoCall(GoExpr.GoIdent("go__concurrency_close"), [channel]),
				isStringLike: false
			};
		}

		if (isStaticCall(callee, "Go", ["go"], "__goSpawn")) {
			requireStdlibShimGroup("go_concurrency");
			var fn = args.length > 0 ? lowerExpr(args[0]).expr : GoExpr.GoFuncLiteral([], [], []);
			return {
				expr: GoExpr.GoCall(GoExpr.GoIdent("go__concurrency_spawn"), [fn]),
				isStringLike: false
			};
		}

		if (isStaticCall(callee, "Log", ["haxe"], "trace")) {
			var arg = args.length > 0 ? lowerExpr(args[0]).expr : GoExpr.GoNil;
			return {
				expr: GoExpr.GoCall(GoExpr.GoIdent("hxrt.Println"), [arg]),
				isStringLike: false
			};
		}

		if (isStaticCall(callee, "Std", [], "string")) {
			var arg = args.length > 0 ? lowerExpr(args[0]).expr : GoExpr.GoNil;
			return {
				expr: GoExpr.GoCall(GoExpr.GoIdent("hxrt.StdString"), [arg]),
				isStringLike: true
			};
		}

		if (isStaticCall(callee, "Std", [], "parseInt")) {
			var arg = args.length > 0 ? lowerExpr(args[0]).expr : GoExpr.GoNil;
			return {
				expr: GoExpr.GoCall(GoExpr.GoIdent("hxrt.StdParseInt"), [arg]),
				isStringLike: false
			};
		}

		if (isStaticCall(callee, "String", [], "fromCharCode")) {
			var code = args.length > 0 ? lowerExpr(args[0]).expr : GoExpr.GoIntLiteral(0);
			return {
				expr: GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromCharCode"), [code]),
				isStringLike: true
			};
		}

		if (isStaticCall(callee, "Std", [], "isOfType")) {
			return lowerStdIsOfTypeCall(args);
		}

		if (isStaticCall(callee, "Exception", ["haxe"], "caught")) {
			var arg = args.length > 0 ? lowerExpr(args[0]).expr : GoExpr.GoNil;
			return {
				expr: GoExpr.GoCall(GoExpr.GoIdent("hxrt.ExceptionCaught"), [arg]),
				isStringLike: false
			};
		}

		if (isStaticCall(callee, "Exception", ["haxe"], "thrown")) {
			var arg = args.length > 0 ? lowerExpr(args[0]).expr : GoExpr.GoNil;
			return {
				expr: GoExpr.GoCall(GoExpr.GoIdent("hxrt.ExceptionThrown"), [arg]),
				isStringLike: false
			};
		}

		var exceptionMessageTarget = asHaxeExceptionMessageGetterTarget(callee);
		if (exceptionMessageTarget != null && args.length == 0) {
			return {
				expr: GoExpr.GoCall(GoExpr.GoIdent("hxrt.ExceptionMessage"), [lowerExpr(exceptionMessageTarget).expr]),
				isStringLike: true
			};
		}

		var exceptionToStringTarget = asHaxeExceptionToStringTarget(callee);
		if (exceptionToStringTarget != null && args.length == 0) {
			return {
				expr: GoExpr.GoCall(GoExpr.GoIdent("hxrt.ExceptionMessage"), [lowerExpr(exceptionToStringTarget).expr]),
				isStringLike: true
			};
		}

		var valueExceptionUnwrapTarget = asHaxeValueExceptionUnwrapTarget(callee);
		if (valueExceptionUnwrapTarget != null && args.length == 0) {
			return {
				expr: GoExpr.GoCall(GoExpr.GoIdent("hxrt.ExceptionThrown"), [lowerExpr(valueExceptionUnwrapTarget).expr]),
				isStringLike: false
			};
		}

		var loweredArgs = new Array<GoExpr>();
		for (index in 0...args.length) {
			var arg = args[index];
			var paramType = callParamType(callee.t, index);
			var emittedParamType = emittedCallParamType(callee, index);
			var nullablePrimitiveArg = paramType != null
				&& isNullablePrimitiveType(paramType) ? lowerNullablePrimitiveCallArgExpr(arg) : null;
			var loweredArg = nullablePrimitiveArg == null ? lowerCallArgExpr(arg) : nullablePrimitiveArg;
			if (paramType != null) {
				loweredArg = upcastIfNeeded(loweredArg, arg.t, paramType, arg);
				if (!isNullablePrimitiveType(paramType)) {
					var argKnownNonNullPrimitive = nonNullPrimitiveExprGoType(arg) != null;
					if (!argKnownNonNullPrimitive && isSharedArrayElementExpr(arg)) {
						loweredArg = coerceStoredArrayElementExpr(loweredArg, paramType);
					} else {
						loweredArg = coerceAnyExprToType(loweredArg, arg.t, paramType, !argKnownNonNullPrimitive && (exprBackedByAny(arg)
							|| shouldForceAnyCoerce(arg.t, paramType)));
					}
				}
			}
			if (emittedParamType != null) {
				loweredArg = adaptErasedFunctionCallArg(loweredArg, arg.t, emittedParamType);
			}
			loweredArg = normalizeExternCallArg(callee, loweredArg, paramType, returnType);
			loweredArgs.push(loweredArg);
		}
		var functionInfo = resolveFunctionInfo(callee);
		if (functionInfo != null && shouldApplySourceDefaultArgPadding(callee) && loweredArgs.length < functionInfo.defaults.length) {
			for (i in loweredArgs.length...functionInfo.defaults.length) {
				var defaultValue = functionInfo.defaults[i];
				if (defaultValue == null) {
					Context.fatalError("Missing required argument at position " + i, callee.pos);
				}
				loweredArgs.push(lowerExpr(defaultValue).expr);
			}
		}

		var callExpr:GoExpr = GoExpr.GoCall(lowerExpr(callee).expr, loweredArgs);
		if (isExternValueErrorCall(callee, returnType)) {
			requireStdlibShimGroup("go_result");
			return {
				expr: GoExpr.GoCall(GoExpr.GoIdent("go__result_fromValueError"), [callExpr]),
				isStringLike: false
			};
		}
		var tupleReturnExpr = lowerExternTupleReturnExpr(callee, returnType, callExpr);
		if (tupleReturnExpr != null) {
			return {
				expr: tupleReturnExpr,
				isStringLike: false
			};
		}
		if (shouldAssertGenericCallResult(callee, returnType)) {
			callExpr = lowerNullableAwareTypeAssertExpr(callExpr, returnType);
		}
		callExpr = normalizeExternStringCallResult(callee, returnType, callExpr);

		return {
			expr: callExpr,
			isStringLike: isStringType(returnType)
		};
	}

	/**
		What
		Adapts direct `haxe.ds.ArraySort` and `haxe.ds.ListSort` calls to Go-compatible
		entrypoints.

		Why
		These upstream helper modules are source-owned, but their generic public
		entrypoints currently erase to `[]any` / `func(any, any) int` on `haxe.go`.
		Without a call-site bridge, direct typed calls fail even though the underlying
		sort implementations are valid.

		How
		For a shared Haxe Array, call the staged helper with the carrier directly. For
		native slice-shaped collections, box to `[]any`, invoke the helper, and copy
		sorted values back. For `ListSort`, adapt the comparator to erased `any`
		parameters and type-assert the returned head to the expected node type.
	**/
	function lowerDsSortHelperCall(callee:TypedExpr, args:Array<TypedExpr>, returnType:Type):Null<LoweredExpr> {
		if (isStaticCall(callee, "ArraySort", ["haxe", "ds"], "sort")) {
			if (args.length != 2) {
				Context.fatalError("haxe.ds.ArraySort.sort expects exactly 2 arguments", callee.pos);
			}
			noteSourceOwnedStdlibStaticCall(callee);
			var arrayType = args[0].t;
			if (!isArrayType(arrayType)) {
				return null;
			}
			var sliceExpr = lowerExpr(args[0]).expr;
			var comparatorExpr = lowerExpr(args[1]).expr;
			if (isHaxeArrayType(arrayType)) {
				return {
					expr: GoExpr.GoCall(GoExpr.GoIdent("haxe__ds__ArraySort_sort"), [sliceExpr, lowerTypedComparatorToAny(comparatorExpr, arrayType)]),
					isStringLike: false
				};
			}
			var sliceType = typeToGoType(arrayType);
			var rawSliceName = freshTempName("hx_sort_raw");
			var sourceName = freshTempName("hx_sort_src");
			var body = new Array<GoStmt>();
			body.push(GoStmt.GoVarDecl(rawSliceName, "[]any", lowerTypedArrayToAnyCoerce(GoExpr.GoIdent(sourceName), arrayType), true));
			body.push(GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("haxe__ds__ArraySort_sort"), [
				GoExpr.GoIdent(rawSliceName),
				lowerTypedComparatorToAny(comparatorExpr, arrayType)
			])));
			body = body.concat(lowerAnyArrayCopyBack(GoExpr.GoIdent(rawSliceName), GoExpr.GoIdent(sourceName), arrayType));
			return {
				expr: GoExpr.GoCall(GoExpr.GoFuncLiteral([{name: sourceName, typeName: sliceType}], [], body), [sliceExpr]),
				isStringLike: false
			};
		}

		if (isStaticCall(callee, "ListSort", ["haxe", "ds"], "sort")
			|| isStaticCall(callee, "ListSort", ["haxe", "ds"], "sortSingleLinked")) {
			if (args.length != 2) {
				Context.fatalError("haxe.ds.ListSort direct calls expect exactly 2 arguments", callee.pos);
			}
			noteSourceOwnedStdlibStaticCall(callee);
			var loweredTarget = lowerExpr(args[0]).expr;
			var loweredComparator = lowerExpr(args[1]).expr;
			var rawCall = GoExpr.GoCall(lowerExpr(callee).expr, [loweredTarget, lowerTypedComparatorToAny(loweredComparator, args[0].t)]);
			return {
				expr: lowerNullableAwareTypeAssertExpr(rawCall, returnType),
				isStringLike: false
			};
		}

		return null;
	}

	function noteSourceOwnedStdlibStaticCall(callee:TypedExpr):Void {
		switch (callee.expr) {
			case TField(_, FStatic(classRef, _)):
				noteSourceOwnedStdlibUsage(classRef.get());
			case TMeta(_, inner):
				noteSourceOwnedStdlibStaticCall(inner);
			case TParenthesis(inner):
				noteSourceOwnedStdlibStaticCall(inner);
			case TCast(inner, _):
				noteSourceOwnedStdlibStaticCall(inner);
			case _:
		}
	}

	function lowerTypedArrayToAnyCoerce(typedSliceExpr:GoExpr, sourceArrayType:Type):GoExpr {
		if (isHaxeArrayType(sourceArrayType)) {
			return GoExpr.GoCall(GoExpr.GoSelector(typedSliceExpr, "Values"), []);
		}
		if (!isArrayType(sourceArrayType) || arrayElementGoType(sourceArrayType) == "any") {
			return typedSliceExpr;
		}
		return lowerTypedSliceToAnyByGoType(typedSliceExpr, arrayElementGoType(sourceArrayType));
	}

	/** Boxes a typed native slice into the carrier's deliberately erased storage. */
	function lowerTypedSliceToAnyByGoType(typedSliceExpr:GoExpr, elementGoType:String):GoExpr {
		if (elementGoType == "any") {
			return typedSliceExpr;
		}
		var sourceName = freshTempName("hx_sort_src");
		var itemName = freshTempName("hx_sort_item");
		var outName = freshTempName("hx_sort_out");
		var sourceType = "[]" + elementGoType;
		return GoExpr.GoCall(GoExpr.GoFuncLiteral([{name: sourceName, typeName: sourceType}], ["[]any"], [
			GoStmt.GoVarDecl(outName, "[]any",
				GoExpr.GoMakeSlice("any", GoExpr.GoIntLiteral(0), GoExpr.GoCall(GoExpr.GoIdent("len"), [GoExpr.GoIdent(sourceName)])), true),
			GoStmt.GoRangeStmt(null, itemName, GoExpr.GoIdent(sourceName), true, [
				GoStmt.GoAssign(GoExpr.GoIdent(outName), GoExpr.GoCall(GoExpr.GoIdent("append"), [GoExpr.GoIdent(outName), GoExpr.GoIdent(itemName)]))
			]),
			GoStmt.GoReturn(GoExpr.GoIdent(outName))
		]), [typedSliceExpr]);
	}

	/**
		What: Lowers the two explicit portable/native slice copy operations.
		Why: A Go slice and a Haxe Array have incompatible identity and growth
		contracts, so neither direction may be disguised as a no-op cast.
		How: Reuse the typed expected-value bridge, which copies values and restores
		the statically known element type at the boundary.
	**/
	function lowerNativeSliceBoundaryCall(callee:TypedExpr, args:Array<TypedExpr>, returnType:Type):Null<LoweredExpr> {
		if (isStaticCall(callee, "NativeSlice", ["go"], "fromArray")) {
			if (args.length != 1) {
				Context.fatalError("go.NativeSlice.fromArray expects exactly one Array", callee.pos);
			}
			return materializeExprWithPrefix(lowerExprWithExpectedUpcast(args[0], returnType), returnType);
		}
		if (isStaticCall(callee, "NativeSlice", ["go"], "append")) {
			if (args.length != 2) {
				Context.fatalError("go.NativeSlice.append expects one slice and one value", callee.pos);
			}
			var loweredTarget = lowerExprWithPrefix(args[0]);
			var loweredValue = lowerExprWithPrefix(args[1]);
			var targetName = freshTempName("hx_native_slice");
			var appended = GoExpr.GoCall(GoExpr.GoIdent("append"), [GoExpr.GoIdent(targetName), loweredValue.expr]);
			return materializeExprWithPrefix({
				prefix: loweredTarget.prefix.concat([GoStmt.GoVarDecl(targetName, typeToGoType(args[0].t), loweredTarget.expr, true)])
					.concat(loweredValue.prefix),
				expr: appended,
				isStringLike: false
			}, returnType);
		}

		return switch (callee.expr) {
			case TField(target, FInstance(classRef, _, field)):
				var classType = classRef.get();
				if (classType.pack.join(".") == "go" && classType.name == "NativeSlice" && field.get().name == "toArray") {
					if (args.length != 0) {
						Context.fatalError("go.NativeSlice.toArray expects no arguments", callee.pos);
					}
					var lowered = lowerExprWithPrefix(target);
					var adapted = upcastIfNeeded(lowered.expr, target.t, returnType, target);
					materializeExprWithPrefix({prefix: lowered.prefix, expr: adapted, isStringLike: false}, returnType);
				} else {
					null;
				}
			case TMeta(_, inner) | TParenthesis(inner) | TCast(inner, _):
				lowerNativeSliceBoundaryCall(inner, args, returnType);
			case _:
				null;
		};
	}

	function lowerAnyArrayCopyBack(rawSliceExpr:GoExpr, targetSliceExpr:GoExpr, targetArrayType:Type):Array<GoStmt> {
		if (isHaxeArrayType(targetArrayType)) {
			return [];
		}
		var targetElementType = arrayElementType(targetArrayType);
		var targetElementGoType = arrayElementGoType(targetArrayType);
		if (targetElementType == null || targetElementGoType == "any") {
			return [];
		}
		var rawName = freshTempName("hx_sort_raw");
		var targetName = freshTempName("hx_sort_dst");
		var indexName = freshTempName("hx_sort_i");
		var itemName = freshTempName("hx_sort_item");
		var convertedItemExpr = lowerNullableAwareTypeAssertExpr(GoExpr.GoIdent(itemName), targetElementType);
		return [
			GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoFuncLiteral([
				{name: rawName, typeName: "[]any"},
				{name: targetName, typeName: typeToGoType(targetArrayType)}
			], [], [
				GoStmt.GoRangeStmt(indexName, itemName, GoExpr.GoIdent(rawName), true, [
					GoStmt.GoAssign(GoExpr.GoIndex(GoExpr.GoIdent(targetName), GoExpr.GoIdent(indexName)), convertedItemExpr)
				])
			]), [rawSliceExpr, targetSliceExpr]))
		];
	}

	function lowerTypedComparatorToAny(comparatorExpr:GoExpr, sourceType:Type):GoExpr {
		var targetType = arrayElementType(sourceType);
		if (targetType == null) {
			targetType = sourceType;
		}
		var targetGoType = scalarGoType(targetType);
		if (targetGoType == "any") {
			return comparatorExpr;
		}
		var leftName = freshTempName("hx_cmp_left");
		var rightName = freshTempName("hx_cmp_right");
		var leftExpr = lowerNullableAwareTypeAssertExpr(GoExpr.GoIdent(leftName), targetType);
		var rightExpr = lowerNullableAwareTypeAssertExpr(GoExpr.GoIdent(rightName), targetType);
		return GoExpr.GoFuncLiteral([{name: leftName, typeName: "any"}, {name: rightName, typeName: "any"}], ["int"],
			[GoStmt.GoReturn(GoExpr.GoCall(comparatorExpr, [leftExpr, rightExpr]))]);
	}

	function lowerTypedGoChanCall(callee:TypedExpr, args:Array<TypedExpr>, returnType:Type):Null<LoweredExpr> {
		if (!useTypedGoConcurrencySpecialization()) {
			return null;
		}

		if (isStaticCall(callee, "Go", ["go"], "newChan")) {
			noteLoweringAttempt("go.concurrency.typed", "go_chan_new", callee.pos, "Attempt typed go.Go.newChan specialization.");
			var elementEligibility = goChanElementEligibility(returnType, "Could not resolve go.Go.newChan return type for native specialization.");
			if (!elementEligibility.eligible) {
				noteProvenConcurrencyFastpathFallback(callee.pos);
				noteLoweringFallback("go.concurrency.typed", "go_chan_new_unmorphable", callee.pos,
					withEligibilityReason("Could not monomorphize go.Go.newChan return type for native specialization.", elementEligibility));
				return null;
			}
			var elementGoType = elementEligibility.goType;
			if (elementGoType == null) {
				noteProvenConcurrencyFastpathFallback(callee.pos);
				noteLoweringFallback("go.concurrency.typed", "go_chan_new_unmorphable", callee.pos,
					"Could not monomorphize go.Go.newChan return type for native specialization.");
				return null;
			}
			requireStdlibShimGroup("go_concurrency");
			registerNativeChanElementGoType(elementGoType);
			noteProvenConcurrencyFastpathHit(callee.pos);
			noteLoweringSuccess("go.concurrency.typed", "go_chan_new", callee.pos,
				'Applied typed go.Go.newChan specialization (element type: ' + elementGoType + ").");
			var buffer = args.length > 0 ? lowerExpr(args[0]).expr : GoExpr.GoIntLiteral(0);
			return {
				expr: GoExpr.GoCall(GoExpr.GoIdent(nativeChanShimName("go__concurrency_newChan", elementGoType)), [buffer]),
				isStringLike: false
			};
		}

		var methodCall = asGoChanMethodCall(callee);
		if (methodCall == null) {
			return null;
		}
		var methodKind = "go_chan_method_" + methodCall.methodName;
		noteLoweringAttempt("go.concurrency.typed", methodKind, callee.pos, "Attempt typed go.Chan method specialization.");

		var elementEligibility = nativeTypeEligibility(methodCall.elementType, GoNativeEligibilityRole.ChanElement,
			"Could not resolve go.Chan method element type for native specialization.");
		if (!elementEligibility.eligible) {
			noteProvenConcurrencyFastpathFallback(callee.pos);
			noteLoweringFallback("go.concurrency.typed", "go_chan_method_unmorphable", callee.pos,
				withEligibilityReason("Could not monomorphize go.Chan method call for native specialization.", elementEligibility));
			return null;
		}
		var elementGoType = elementEligibility.goType;
		if (elementGoType == null) {
			noteProvenConcurrencyFastpathFallback(callee.pos);
			noteLoweringFallback("go.concurrency.typed", "go_chan_method_unmorphable", callee.pos,
				"Could not monomorphize go.Chan method call for native specialization.");
			return null;
		}

		requireStdlibShimGroup("go_concurrency");
		registerNativeChanElementGoType(elementGoType);
		noteProvenConcurrencyFastpathHit(callee.pos);
		noteLoweringSuccess("go.concurrency.typed", methodKind, callee.pos,
			'Applied typed go.Chan method specialization (element type: ' + elementGoType + ").");

		var channel = lowerExpr(methodCall.target).expr;
		var channelNative = GoExpr.GoSelector(channel, "__hx_native");

		switch (methodCall.methodName) {
			case "send":
				var value = args.length > 0 ? lowerExpr(args[0]).expr : GoExpr.GoNil;
				if (args.length > 0 && elementGoType != "any" && exprBackedByAny(args[0])) {
					value = GoExpr.GoTypeAssert(value, elementGoType);
				}
				return {
					expr: GoExpr.GoCall(GoExpr.GoIdent(nativeChanShimName("go__concurrency_send", elementGoType)), [channelNative, value]),
					isStringLike: false
				};
			case "trySend":
				var value = args.length > 0 ? lowerExpr(args[0]).expr : GoExpr.GoNil;
				if (args.length > 0 && elementGoType != "any" && exprBackedByAny(args[0])) {
					value = GoExpr.GoTypeAssert(value, elementGoType);
				}
				return {
					expr: GoExpr.GoCall(GoExpr.GoIdent(nativeChanShimName("go__concurrency_trySend", elementGoType)), [channelNative, value]),
					isStringLike: false
				};
			case "recv":
				return {
					expr: GoExpr.GoCall(GoExpr.GoIdent(nativeChanShimName("go__concurrency_recv", elementGoType)), [channelNative]),
					isStringLike: isStringType(returnType)
				};
			case "recvOr":
				var defaultValue = args.length > 0 ? lowerExpr(args[0]).expr : GoExpr.GoNil;
				if (args.length > 0 && elementGoType != "any" && exprBackedByAny(args[0])) {
					defaultValue = GoExpr.GoTypeAssert(defaultValue, elementGoType);
				}
				return {
					expr: GoExpr.GoCall(GoExpr.GoIdent(nativeChanShimName("go__concurrency_recvOr", elementGoType)), [channelNative, defaultValue]),
					isStringLike: isStringType(returnType)
				};
			case "tryRecv":
				return {
					expr: GoExpr.GoCall(GoExpr.GoIdent(nativeChanShimName("go__concurrency_tryRecv", elementGoType)), [channelNative]),
					isStringLike: false
				};
			case "close":
				return {
					expr: GoExpr.GoCall(GoExpr.GoIdent(nativeChanShimName("go__concurrency_close", elementGoType)), [channelNative]),
					isStringLike: false
				};
			case "__hx_setBuffer":
				var buffer = args.length > 0 ? lowerExpr(args[0]).expr : GoExpr.GoIntLiteral(0);
				return {
					expr: GoExpr.GoCall(GoExpr.GoIdent(nativeChanShimName("go__concurrency_setBuffer", elementGoType)), [channel, buffer]),
					isStringLike: false
				};
			case _:
				return null;
		}
	}

	function lowerTypedGoSliceCall(callee:TypedExpr, args:Array<TypedExpr>, returnType:Type):Null<LoweredExpr> {
		if (!useTypedGoCollectionsSpecialization()) {
			return null;
		}

		var methodCall = asGoSliceMethodCall(callee);
		if (methodCall == null) {
			return null;
		}
		var methodKind = "go_slice_method_" + methodCall.methodName;
		noteLoweringAttempt("go.collections.typed", methodKind, callee.pos, "Attempt typed go.Slice method specialization.");

		var elementEligibility = goSliceElementEligibility(methodCall.target.t, "Could not resolve go.Slice element type for native specialization.");
		if (!elementEligibility.eligible) {
			noteLoweringFallback("go.collections.typed", "go_slice_method_unmorphable", callee.pos,
				withEligibilityReason("Could not monomorphize go.Slice element type for native specialization.", elementEligibility));
			return null;
		}
		var elementGoType = elementEligibility.goType;
		if (elementGoType == null) {
			noteLoweringFallback("go.collections.typed", "go_slice_method_unmorphable", callee.pos,
				"Could not monomorphize go.Slice element type for native specialization.");
			return null;
		}

		requireStdlibShimGroup("go_collections");
		registerNativeSliceElementGoType(elementGoType);
		noteLoweringSuccess("go.collections.typed", methodKind, callee.pos,
			'Applied typed go.Slice method specialization (element type: ' + elementGoType + ").");
		var sliceExpr = lowerExpr(methodCall.target).expr;

		switch (methodCall.methodName) {
			case "push":
				var value = args.length > 0 ? lowerExpr(args[0]).expr : GoExpr.GoNil;
				return {
					expr: GoExpr.GoCall(GoExpr.GoIdent(nativeSliceShimName("go__slice_push", elementGoType)), [sliceExpr, value]),
					isStringLike: false
				};
			case "set":
				var index = args.length > 0 ? lowerExpr(args[0]).expr : GoExpr.GoIntLiteral(0);
				var value = args.length > 1 ? lowerExpr(args[1]).expr : GoExpr.GoNil;
				return {
					expr: GoExpr.GoCall(GoExpr.GoIdent(nativeSliceShimName("go__slice_set", elementGoType)), [sliceExpr, index, value]),
					isStringLike: false
				};
			case "get":
				var index = args.length > 0 ? lowerExpr(args[0]).expr : GoExpr.GoIntLiteral(0);
				return {
					expr: GoExpr.GoCall(GoExpr.GoIdent(nativeSliceShimName("go__slice_get", elementGoType)), [sliceExpr, index]),
					isStringLike: isStringType(returnType)
				};
			case "get_length":
				return {
					expr: GoExpr.GoCall(GoExpr.GoIdent(nativeSliceShimName("go__slice_length", elementGoType)), [sliceExpr]),
					isStringLike: false
				};
			case "toArray":
				var typedSlice = GoExpr.GoCall(GoExpr.GoIdent(nativeSliceShimName("go__slice_toArray", elementGoType)), [sliceExpr]);
				return {
					expr: GoExpr.GoCall(GoExpr.GoIdent("hxrt.ArrayFromValues"), [lowerTypedSliceToAnyByGoType(typedSlice, elementGoType)]),
					isStringLike: false
				};
			case _:
				return null;
		}
	}

	function lowerTypedGoMapCall(callee:TypedExpr, args:Array<TypedExpr>, returnType:Type):Null<LoweredExpr> {
		if (!useTypedGoCollectionsSpecialization()) {
			return null;
		}

		var methodCall = asGoMapMethodCall(callee);
		if (methodCall == null) {
			return null;
		}
		var methodKind = "go_map_method_" + methodCall.methodName;
		noteLoweringAttempt("go.collections.typed", methodKind, callee.pos, "Attempt typed go.Map method specialization.");

		var pair = goMapTypePair(methodCall.target.t);
		if (pair == null) {
			noteLoweringFallback("go.collections.typed", "go_map_method_unmorphable", callee.pos,
				"Could not monomorphize go.Map key/value types for native specialization.");
			return null;
		}
		var keyEligibility = nativeTypeEligibility(pair.keyType, GoNativeEligibilityRole.MapKey,
			"Could not resolve go.Map key type for native specialization.");
		if (!keyEligibility.eligible) {
			noteLoweringFallback("go.collections.typed", "go_map_method_unmorphable", callee.pos,
				withEligibilityReason("Could not monomorphize go.Map key/value types for native specialization.", keyEligibility));
			return null;
		}
		var valueEligibility = nativeTypeEligibility(pair.valueType, GoNativeEligibilityRole.MapValue,
			"Could not resolve go.Map value type for native specialization.");
		if (!valueEligibility.eligible) {
			noteLoweringFallback("go.collections.typed", "go_map_method_unmorphable", callee.pos,
				withEligibilityReason("Could not monomorphize go.Map key/value types for native specialization.", valueEligibility));
			return null;
		}
		var keyGoType = keyEligibility.goType;
		var valueGoType = valueEligibility.goType;
		if (keyGoType == null || valueGoType == null) {
			noteLoweringFallback("go.collections.typed", "go_map_method_unmorphable", callee.pos,
				"Could not monomorphize go.Map key/value types for native specialization.");
			return null;
		}

		requireStdlibShimGroup("go_collections");
		registerNativeMapTypePair(keyGoType, valueGoType);
		noteLoweringSuccess("go.collections.typed", methodKind, callee.pos,
			'Applied typed go.Map method specialization (key: '
			+ keyGoType
			+ ", value: "
			+ valueGoType
			+ ").");
		var mapExpr = lowerExpr(methodCall.target).expr;

		switch (methodCall.methodName) {
			case "set":
				var keyExpr = args.length > 0 ? lowerExpr(args[0]).expr : GoExpr.GoNil;
				var valueExpr = args.length > 1 ? lowerExpr(args[1]).expr : GoExpr.GoNil;
				return {
					expr: GoExpr.GoCall(GoExpr.GoIdent(nativeMapShimName("go__map_set", keyGoType, valueGoType)), [mapExpr, keyExpr, valueExpr]),
					isStringLike: false
				};
			case "get":
				var keyExpr = args.length > 0 ? lowerExpr(args[0]).expr : GoExpr.GoNil;
				var rawCall = GoExpr.GoCall(GoExpr.GoIdent(nativeMapShimName("go__map_get", keyGoType, valueGoType)), [mapExpr, keyExpr]);
				return {
					expr: lowerNullableAwareTypeAssertExpr(rawCall, returnType),
					isStringLike: isStringType(returnType)
				};
			case "exists":
				var keyExpr = args.length > 0 ? lowerExpr(args[0]).expr : GoExpr.GoNil;
				return {
					expr: GoExpr.GoCall(GoExpr.GoIdent(nativeMapShimName("go__map_exists", keyGoType, valueGoType)), [mapExpr, keyExpr]),
					isStringLike: false
				};
			case _:
				return null;
		}
	}

	function lowerTypedGoResultCall(callee:TypedExpr, args:Array<TypedExpr>, returnType:Type):Null<LoweredExpr> {
		if (!useTypedGoResultSpecialization()) {
			return null;
		}

		var returnElementEligibility = goResultElementEligibility(returnType, "Could not resolve go.Result<T> return type for native specialization.");

		if (isStaticCall(callee, "Result", ["go"], "ok") || isStaticCall(callee, "Go", ["go"], "ok")) {
			noteLoweringAttempt("go.result.typed", "go_result_static_ok", callee.pos, "Attempt typed go.Result.ok specialization.");
			if (!returnElementEligibility.eligible || returnElementEligibility.goType == null) {
				noteLoweringFallback("go.result.typed", "go_result_static_ok_unmorphable", callee.pos,
					withEligibilityReason("Could not monomorphize go.Result<T>.ok return type for native specialization.", returnElementEligibility));
				return null;
			}
			var returnElementGoType = returnElementEligibility.goType;
			requireStdlibShimGroup("go_result");
			registerNativeResultElementGoType(returnElementGoType);
			noteLoweringSuccess("go.result.typed", "go_result_static_ok", callee.pos,
				'Applied typed go.Result.ok specialization (element type: ' + returnElementGoType + ").");
			var value = args.length > 0 ? lowerExpr(args[0]).expr : GoExpr.GoNil;
			return {
				expr: GoExpr.GoCall(GoExpr.GoIdent(nativeResultShimName("go__result_ok", returnElementGoType)), [value]),
				isStringLike: false
			};
		}

		if (isStaticCall(callee, "Result", ["go"], "failure") || isStaticCall(callee, "Go", ["go"], "fail")) {
			noteLoweringAttempt("go.result.typed", "go_result_static_failure", callee.pos, "Attempt typed go.Result.failure specialization.");
			if (!returnElementEligibility.eligible || returnElementEligibility.goType == null) {
				noteLoweringFallback("go.result.typed", "go_result_static_failure_unmorphable", callee.pos,
					withEligibilityReason("Could not monomorphize go.Result<T>.failure return type for native specialization.", returnElementEligibility));
				return null;
			}
			var returnElementGoType = returnElementEligibility.goType;
			requireStdlibShimGroup("go_result");
			registerNativeResultElementGoType(returnElementGoType);
			noteLoweringSuccess("go.result.typed", "go_result_static_failure", callee.pos,
				'Applied typed go.Result.failure specialization (element type: ' + returnElementGoType + ").");
			var message = args.length > 0 ? lowerExpr(args[0]).expr : GoExpr.GoNil;
			return {
				expr: GoExpr.GoCall(GoExpr.GoIdent(nativeResultShimName("go__result_failure", returnElementGoType)), [message]),
				isStringLike: false
			};
		}

		var methodCall = asGoResultMethodCall(callee);
		if (methodCall == null) {
			return null;
		}
		var methodKind = "go_result_method_" + methodCall.methodName;
		noteLoweringAttempt("go.result.typed", methodKind, callee.pos, "Attempt typed go.Result method specialization.");

		var receiverEligibility = goResultElementEligibility(methodCall.target.t,
			"Could not resolve go.Result<T> method receiver type for native specialization.");
		if (!receiverEligibility.eligible || receiverEligibility.goType == null) {
			noteLoweringFallback("go.result.typed", "go_result_method_unmorphable", callee.pos,
				withEligibilityReason("Could not monomorphize go.Result<T> method receiver for native specialization.", receiverEligibility));
			return null;
		}
		var elementGoType = receiverEligibility.goType;

		requireStdlibShimGroup("go_result");
		registerNativeResultElementGoType(elementGoType);
		noteLoweringSuccess("go.result.typed", methodKind, callee.pos, 'Applied typed go.Result method specialization (element type: ' + elementGoType + ").");
		var resultExpr = lowerExpr(methodCall.target).expr;

		switch (methodCall.methodName) {
			case "isOk":
				return {
					expr: GoExpr.GoCall(GoExpr.GoIdent(nativeResultShimName("go__result_isOk", elementGoType)), [resultExpr]),
					isStringLike: false
				};
			case "isErr":
				return {
					expr: GoExpr.GoCall(GoExpr.GoIdent(nativeResultShimName("go__result_isErr", elementGoType)), [resultExpr]),
					isStringLike: false
				};
			case "unwrap":
				return {
					expr: GoExpr.GoCall(GoExpr.GoIdent(nativeResultShimName("go__result_unwrap", elementGoType)), [resultExpr]),
					isStringLike: isStringType(returnType)
				};
			case "error":
				return {
					expr: GoExpr.GoCall(GoExpr.GoIdent(nativeResultShimName("go__result_error", elementGoType)), [resultExpr]),
					isStringLike: true
				};
			case _:
				return null;
		}
	}

	function lowerExternReceiverCall(callee:TypedExpr, args:Array<TypedExpr>, returnType:Type):Null<LoweredExpr> {
		return switch (callee.expr) {
			case TField(_, FStatic(classRef, fieldRef)):
				var classType = classRef.get();
				var field = fieldRef.get();
				if (!classType.isExtern || !hasExternReceiverMeta(field)) {
					null;
				} else if (args.length == 0) {
					Context.fatalError("Extern @:go.receiver call requires a receiver argument", callee.pos);
					null;
				} else {
					var externPackage = externClassPackageName(classType);
					if (externPackage != null) {
						noteExternImportPath(classType, externPackage);
					}
					var loweredReceiver = lowerExpr(args[0]).expr;
					var loweredArgs = new Array<GoExpr>();
					for (index in 1...args.length) {
						var arg = args[index];
						var loweredArg = lowerCallArgExpr(arg);
						var paramType = callParamType(callee.t, index);
						if (paramType != null) {
							loweredArg = upcastIfNeeded(loweredArg, arg.t, paramType, arg);
						}
						loweredArg = normalizeExternCallArg(callee, loweredArg, paramType, returnType);
						loweredArgs.push(loweredArg);
					}

					var callExpr = GoExpr.GoCall(GoExpr.GoSelector(loweredReceiver, externFieldName(field)), loweredArgs);
					if (isExternValueErrorCall(callee, returnType)) {
						requireStdlibShimGroup("go_result");
						callExpr = GoExpr.GoCall(GoExpr.GoIdent("go__result_fromValueError"), [callExpr]);
						return {
							expr: callExpr,
							isStringLike: false
						};
					}
					var tupleReturnExpr = lowerExternTupleReturnExpr(callee, returnType, callExpr);
					if (tupleReturnExpr != null) {
						return {
							expr: tupleReturnExpr,
							isStringLike: false
						};
					}
					if (shouldAssertGenericCallResult(callee, returnType)) {
						callExpr = lowerNullableAwareTypeAssertExpr(callExpr, returnType);
					}
					callExpr = normalizeExternStringCallResult(callee, returnType, callExpr);

					{
						expr: callExpr,
						isStringLike: isStringType(returnType)
					};
				}
			case TMeta(_, inner):
				lowerExternReceiverCall(inner, args, returnType);
			case TParenthesis(inner):
				lowerExternReceiverCall(inner, args, returnType);
			case TCast(inner, _):
				lowerExternReceiverCall(inner, args, returnType);
			case _:
				null;
		};
	}

	function lowerStringInstanceCall(callee:TypedExpr, args:Array<TypedExpr>, returnType:Type):Null<LoweredExpr> {
		return switch (callee.expr) {
			case TField(target, FInstance(classRef, _, field)):
				var classType = classRef.get();
				var resolvedField = field.get();
				if (classType.pack.length != 0 || classType.name != "String") {
					null;
				} else {
					var loweredTarget = lowerExpr(target).expr;
					if (isSharedArrayElementExpr(target)) {
						loweredTarget = coerceStoredArrayElementExpr(loweredTarget, target.t);
					}
					var useTypedHelpers = useStringFastpath();
					switch (resolvedField.name) {
						case "charAt":
							var indexExpr = args.length > 0 ? lowerExpr(args[0]).expr : GoExpr.GoIntLiteral(0);
							var helper = if (useTypedHelpers) {
								compilationContext.optimizerStringInstanceTypedLowerings++;
								"hxrt.StringCharAtStringPtr";
							} else {
								compilationContext.optimizerStringInstanceLegacyLowerings++;
								"hxrt.StringCharAt";
							};
							{
								expr: GoExpr.GoCall(GoExpr.GoIdent(helper), [loweredTarget, indexExpr]),
								isStringLike: true
							};
						case "charCodeAt":
							var indexExpr = args.length > 0 ? lowerExpr(args[0]).expr : GoExpr.GoIntLiteral(0);
							var useIntReturn = isIntType(returnType) && !isNullableIntType(returnType);
							var helper = if (useTypedHelpers) {
								compilationContext.optimizerStringInstanceTypedLowerings++;
								useIntReturn ? "hxrt.StringCharCodeAtStringPtr" : "hxrt.StringCharCodeAtAnyStringPtr";
							} else {
								compilationContext.optimizerStringInstanceLegacyLowerings++;
								useIntReturn ? "hxrt.StringCharCodeAt" : "hxrt.StringCharCodeAtAny";
							};
							{
								expr: GoExpr.GoCall(GoExpr.GoIdent(helper), [loweredTarget, indexExpr]),
								isStringLike: false
							};
						case "substring":
							var startExpr = args.length > 0 ? lowerExpr(args[0]).expr : GoExpr.GoIntLiteral(0);
							var lengthHelper = useTypedHelpers ? "hxrt.StringLengthStringPtr" : "hxrt.StringLength";
							var substringHelper = useTypedHelpers ? "hxrt.StringSubstringStringPtr" : "hxrt.StringSubstring";
							if (useTypedHelpers) {
								compilationContext.optimizerStringInstanceTypedLowerings++;
							} else {
								compilationContext.optimizerStringInstanceLegacyLowerings++;
							}
							var endExpr = args.length > 1 ? lowerExpr(args[1]).expr : GoExpr.GoCall(GoExpr.GoIdent(lengthHelper), [loweredTarget]);
							{
								expr: GoExpr.GoCall(GoExpr.GoIdent(substringHelper), [loweredTarget, startExpr, endExpr]),
								isStringLike: true
							};
						case "substr":
							var posExpr = args.length > 0 ? lowerExpr(args[0]).expr : GoExpr.GoIntLiteral(0);
							var lenExpr = args.length > 1 ? lowerExpr(args[1]).expr : GoExpr.GoIntLiteral(0);
							var hasLenExpr = GoExpr.GoBoolLiteral(args.length > 1);
							var helper = if (useTypedHelpers) {
								compilationContext.optimizerStringInstanceTypedLowerings++;
								"hxrt.StringSubstrStringPtr";
							} else {
								compilationContext.optimizerStringInstanceLegacyLowerings++;
								"hxrt.StringSubstr";
							};
							{
								expr: GoExpr.GoCall(GoExpr.GoIdent(helper), [loweredTarget, posExpr, lenExpr, hasLenExpr]),
								isStringLike: true
							};
						case "lastIndexOf":
							var searchExpr = args.length > 0 ? lowerExpr(args[0]).expr : GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"),
								[GoExpr.GoStringLiteral("")]);
							var startExpr = args.length > 1 ? lowerExpr(args[1]).expr : GoExpr.GoIntLiteral(0);
							var hasStartExpr = GoExpr.GoBoolLiteral(args.length > 1);
							var helper = if (useTypedHelpers) {
								compilationContext.optimizerStringInstanceTypedLowerings++;
								"hxrt.StringLastIndexOfStringPtr";
							} else {
								compilationContext.optimizerStringInstanceLegacyLowerings++;
								"hxrt.StringLastIndexOf";
							};
							{
								expr: GoExpr.GoCall(GoExpr.GoIdent(helper), [loweredTarget, searchExpr, startExpr, hasStartExpr]),
								isStringLike: false
							};
						case "split":
							var delimiterExpr = args.length > 0 ? lowerExpr(args[0]).expr : GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"),
								[GoExpr.GoStringLiteral("")]);
							var helper = if (useTypedHelpers) {
								compilationContext.optimizerStringInstanceTypedLowerings++;
								"hxrt.StringSplitStringPtr";
							} else {
								compilationContext.optimizerStringInstanceLegacyLowerings++;
								"hxrt.StringSplit";
							};
							var nativeParts = GoExpr.GoCall(GoExpr.GoIdent(helper), [loweredTarget, delimiterExpr]);
							{
								expr: GoExpr.GoCall(GoExpr.GoIdent("hxrt.ArrayFromValues"), [lowerTypedSliceToAnyByGoType(nativeParts, "*string")]),
								isStringLike: false
							};
						case "toLowerCase":
							var helper = if (useTypedHelpers) {
								compilationContext.optimizerStringInstanceTypedLowerings++;
								"hxrt.StringToLowerCaseStringPtr";
							} else {
								compilationContext.optimizerStringInstanceLegacyLowerings++;
								"hxrt.StringToLowerCase";
							};
							{
								expr: GoExpr.GoCall(GoExpr.GoIdent(helper), [loweredTarget]),
								isStringLike: true
							};
						case "toUpperCase":
							var helper = if (useTypedHelpers) {
								compilationContext.optimizerStringInstanceTypedLowerings++;
								"hxrt.StringToUpperCaseStringPtr";
							} else {
								compilationContext.optimizerStringInstanceLegacyLowerings++;
								"hxrt.StringToUpperCase";
							};
							{
								expr: GoExpr.GoCall(GoExpr.GoIdent(helper), [loweredTarget]),
								isStringLike: true
							};
						case _:
							null;
					}
				}
			case TMeta(_, inner):
				lowerStringInstanceCall(inner, args, returnType);
			case TParenthesis(inner):
				lowerStringInstanceCall(inner, args, returnType);
			case TCast(inner, _):
				lowerStringInstanceCall(inner, args, returnType);
			case _:
				null;
		};
	}

	function lowerCallArgExpr(expr:TypedExpr):GoExpr {
		if (!isVoidType(expr.t)) {
			return lowerExpr(expr).expr;
		}

		var lowered = lowerExprWithPrefix(expr);
		var body = lowered.prefix.copy();
		if (!isNilExpr(lowered.expr)) {
			body.push(GoStmt.GoExprStmt(lowered.expr));
		}
		body.push(GoStmt.GoReturn(GoExpr.GoNil));
		return GoExpr.GoCall(GoExpr.GoFuncLiteral([], ["any"], body), []);
	}

	function lowerNullablePrimitiveCallArgExpr(expr:TypedExpr):Null<GoExpr> {
		return switch (expr.expr) {
			case TLocal(variable):
				(registeredOptionalPrimitiveParamGoType(variable) != null
					|| valueStorageGoType(variable.t) == "any") ? GoExpr.GoIdent(localVarName(variable)) : null;
			case TMeta(_, inner):
				lowerNullablePrimitiveCallArgExpr(inner);
			case TParenthesis(inner):
				lowerNullablePrimitiveCallArgExpr(inner);
			case TCast(inner, _):
				lowerNullablePrimitiveCallArgExpr(inner);
			case _:
				null;
		};
	}

	function haxeDsListElementType(type:Type):Null<Type> {
		var followed = Context.follow(type);
		return switch (followed) {
			case TInst(classRef, params):
				var classType = classRef.get();
				if (classType.pack.join(".") == "haxe.ds" && classType.name == "List" && params.length == 1) {
					params[0];
				} else {
					null;
				}
			case TAbstract(abstractRef, params):
				var abstractType = abstractRef.get();
				if (abstractType.pack.length == 0 && abstractType.name == "Null" && params.length == 1) {
					haxeDsListElementType(params[0]);
				} else {
					null;
				}
			case TMono(ref):
				var resolved = ref.get();
				resolved == null ? null : haxeDsListElementType(resolved);
			case _:
				null;
		};
	}

	function arrayElementType(type:Type):Null<Type> {
		var nativeSliceElement = nativeSliceElementType(type);
		if (nativeSliceElement != null) {
			return nativeSliceElement;
		}
		var restElement = restElementType(type);
		if (restElement != null) {
			return restElement;
		}
		var vectorElement = vectorElementType(type);
		if (vectorElement != null) {
			return vectorElement;
		}
		var readOnlyElement = readOnlyArrayElementType(type);
		if (readOnlyElement != null) {
			return readOnlyElement;
		}
		return switch (Context.follow(type)) {
			case TInst(classRef, params):
				var classType = classRef.get();
				if (classType.pack.length == 0 && classType.name == "Array" && params.length == 1) {
					params[0];
				} else {
					null;
				}
			case _:
				null;
		};
	}

	function lowerLambdaFunctionValueCall(callee:TypedExpr, args:Array<TypedExpr>, returnType:Type):Null<LoweredExpr> {
		if (args.length < 2) {
			return null;
		}
		var alias = lambdaIterableLowering.functionAliasName(callee);
		if (alias == null) {
			return null;
		}

		// Function-value Lambda.map alias call path.
		if (alias == "map") {
			if (args.length != 2) {
				Context.fatalError("Lambda.map expects exactly 2 arguments", callee.pos);
			}
			var calleeExpr = lowerExpr(callee).expr;
			var dynamicSourceExpr = lambdaIterableLowering.dynamicIterableSource(args[0]);
			var mapperExpr = lowerExpr(args[1]).expr;
			var adaptedMapperExpr = lambdaIterableLowering.mapperAnyAdapter(mapperExpr, args[1].t);
			var mappedAnyExpr = GoExpr.GoCall(calleeExpr, [dynamicSourceExpr, adaptedMapperExpr]);
			return {
				expr: mappedAnyExpr,
				isStringLike: false
			};
		}

		// Function-value Lambda.fold alias call path.
		if (alias == "fold") {
			if (args.length != 3) {
				Context.fatalError("Lambda.fold expects exactly 3 arguments", callee.pos);
			}
			var calleeExpr = lowerExpr(callee).expr;
			var dynamicSourceExpr = lambdaIterableLowering.dynamicIterableSource(args[0]);
			var folderExpr = lowerExpr(args[1]).expr;
			var initExpr = lowerExpr(args[2]).expr;
			var adaptedFolderExpr = lambdaIterableLowering.folderAnyAdapter(folderExpr, args[1].t);
			var foldedAnyExpr = GoExpr.GoCall(calleeExpr, [dynamicSourceExpr, adaptedFolderExpr, initExpr]);
			return {
				expr: lowerNullableAwareTypeAssertExpr(foldedAnyExpr, returnType),
				isStringLike: false
			};
		}

		return null;
	}

	/**
		What
		- Adapts direct calls to the staged `Lambda` source entrypoints.

		Why
		- Go slices, the staged list carrier, and concrete iterator classes do not
		  implement the erased `map[string]any` shape generated for Haxe's structural
		  `Iterable<T>` parameter. Typed callbacks and generic results have the same
		  Go invariance mismatch. Calling the source functions without a bridge would
		  therefore reject valid portable Haxe programs.

		How
		- Wrap the input in the manual iterator protocol and adapt only callback
		  parameters. Array-producing staged functions already return the shared
		  Haxe Array carrier, so this bridge preserves that result without an extra
		  conversion. The staged Haxe function still owns every loop, comparison,
		  early exit, allocation, and algorithmic decision.
	**/
	function lowerLambdaSourceCallAdapter(callee:TypedExpr, args:Array<TypedExpr>, returnType:Type):Null<LoweredExpr> {
		if (isStaticCall(callee, "Lambda", [], "array") || lambdaIterableLowering.isGeneratedCall(callee, "array")) {
			if (args.length != 1) {
				Context.fatalError("Lambda.array expects exactly 1 argument", callee.pos);
			}
			var arrayExpr = GoExpr.GoCall(lowerExpr(callee).expr, [lambdaIterableLowering.dynamicIterableSource(args[0])]);
			return {
				expr: arrayExpr,
				isStringLike: false
			};
		}

		if (isStaticCall(callee, "Lambda", [], "list") || lambdaIterableLowering.isGeneratedCall(callee, "list")) {
			if (args.length != 1) {
				Context.fatalError("Lambda.list expects exactly 1 argument", callee.pos);
			}
			return {
				expr: GoExpr.GoCall(lowerExpr(callee).expr, [lambdaIterableLowering.dynamicIterableSource(args[0])]),
				isStringLike: false
			};
		}

		if (isStaticCall(callee, "Lambda", [], "mapi") || lambdaIterableLowering.isGeneratedCall(callee, "mapi")) {
			if (args.length != 2) {
				Context.fatalError("Lambda.mapi expects exactly 2 arguments", callee.pos);
			}
			var mapperExpr = lambdaIterableLowering.indexedMapperAnyAdapter(lowerExpr(args[1]).expr, args[1].t);
			var mappedExpr = GoExpr.GoCall(lowerExpr(callee).expr, [lambdaIterableLowering.dynamicIterableSource(args[0]), mapperExpr]);
			return {
				expr: mappedExpr,
				isStringLike: false
			};
		}

		if (isStaticCall(callee, "Lambda", [], "flatten") || lambdaIterableLowering.isGeneratedCall(callee, "flatten")) {
			if (args.length != 1) {
				Context.fatalError("Lambda.flatten expects exactly 1 argument", callee.pos);
			}
			requireSourceOwnedStdlibModule("Lambda");
			var flattenedExpr = GoExpr.GoCall(lowerExpr(callee).expr, [lambdaIterableLowering.dynamicNestedIterableSource(args[0])]);
			return {
				expr: flattenedExpr,
				isStringLike: false
			};
		}

		if (isStaticCall(callee, "Lambda", [], "flatMap") || lambdaIterableLowering.isGeneratedCall(callee, "flatMap")) {
			if (args.length != 2) {
				Context.fatalError("Lambda.flatMap expects exactly 2 arguments", callee.pos);
			}
			requireSourceOwnedStdlibModule("Lambda");
			var mapperExpr = lambdaIterableLowering.iterableMapperAnyAdapter(lowerExpr(args[1]).expr, args[1].t);
			var mappedExpr = GoExpr.GoCall(lowerExpr(callee).expr, [lambdaIterableLowering.dynamicIterableSource(args[0]), mapperExpr]);
			return {
				expr: mappedExpr,
				isStringLike: false
			};
		}

		if (isStaticCall(callee, "Lambda", [], "count") || lambdaIterableLowering.isGeneratedCall(callee, "count")) {
			if (args.length < 1 || args.length > 2) {
				Context.fatalError("Lambda.count expects 1 or 2 arguments", callee.pos);
			}
			var predicateExpr = GoExpr.GoNil;
			if (args.length == 2 && !isNullLiteralExpr(args[1])) {
				predicateExpr = lambdaIterableLowering.predicateAnyAdapter(lowerExpr(args[1]).expr, args[1].t);
			}
			return {
				expr: GoExpr.GoCall(lowerExpr(callee).expr, [lambdaIterableLowering.dynamicIterableSource(args[0]), predicateExpr]),
				isStringLike: false
			};
		}

		if (isStaticCall(callee, "Lambda", [], "empty") || lambdaIterableLowering.isGeneratedCall(callee, "empty")) {
			if (args.length != 1) {
				Context.fatalError("Lambda.empty expects exactly 1 argument", callee.pos);
			}
			return {
				expr: GoExpr.GoCall(lowerExpr(callee).expr, [lambdaIterableLowering.dynamicIterableSource(args[0])]),
				isStringLike: false
			};
		}

		if (isStaticCall(callee, "Lambda", [], "exists") || lambdaIterableLowering.isGeneratedCall(callee, "exists")) {
			if (args.length != 2) {
				Context.fatalError("Lambda.exists expects exactly 2 arguments", callee.pos);
			}
			var predicateExpr = lambdaIterableLowering.predicateAnyAdapter(lowerExpr(args[1]).expr, args[1].t);
			return {
				expr: GoExpr.GoCall(lowerExpr(callee).expr, [lambdaIterableLowering.dynamicIterableSource(args[0]), predicateExpr]),
				isStringLike: false
			};
		}

		if (isStaticCall(callee, "Lambda", [], "foreach") || lambdaIterableLowering.isGeneratedCall(callee, "foreach")) {
			if (args.length != 2) {
				Context.fatalError("Lambda.foreach expects exactly 2 arguments", callee.pos);
			}
			var predicateExpr = lambdaIterableLowering.predicateAnyAdapter(lowerExpr(args[1]).expr, args[1].t);
			return {
				expr: GoExpr.GoCall(lowerExpr(callee).expr, [lambdaIterableLowering.dynamicIterableSource(args[0]), predicateExpr]),
				isStringLike: false
			};
		}

		if (isStaticCall(callee, "Lambda", [], "has") || lambdaIterableLowering.isGeneratedCall(callee, "has")) {
			if (args.length != 2) {
				Context.fatalError("Lambda.has expects exactly 2 arguments", callee.pos);
			}
			return {
				expr: GoExpr.GoCall(lowerExpr(callee).expr, [lambdaIterableLowering.dynamicIterableSource(args[0]), lowerExpr(args[1]).expr]),
				isStringLike: false
			};
		}

		if (isStaticCall(callee, "Lambda", [], "iter") || lambdaIterableLowering.isGeneratedCall(callee, "iter")) {
			if (args.length != 2) {
				Context.fatalError("Lambda.iter expects exactly 2 arguments", callee.pos);
			}
			var consumerExpr = lambdaIterableLowering.consumerAnyAdapter(lowerExpr(args[1]).expr, args[1].t);
			return {
				expr: GoExpr.GoCall(lowerExpr(callee).expr, [lambdaIterableLowering.dynamicIterableSource(args[0]), consumerExpr]),
				isStringLike: false
			};
		}

		if (isStaticCall(callee, "Lambda", [], "filter") || lambdaIterableLowering.isGeneratedCall(callee, "filter")) {
			if (args.length != 2) {
				Context.fatalError("Lambda.filter expects exactly 2 arguments", callee.pos);
			}
			var predicateExpr = lambdaIterableLowering.predicateAnyAdapter(lowerExpr(args[1]).expr, args[1].t);
			var filteredExpr = GoExpr.GoCall(lowerExpr(callee).expr, [lambdaIterableLowering.dynamicIterableSource(args[0]), predicateExpr]);
			return {
				expr: filteredExpr,
				isStringLike: false
			};
		}

		if (isStaticCall(callee, "Lambda", [], "map") || lambdaIterableLowering.isGeneratedCall(callee, "map")) {
			if (args.length != 2) {
				Context.fatalError("Lambda.map expects exactly 2 arguments", callee.pos);
			}
			var mapperExpr = lambdaIterableLowering.mapperAnyAdapter(lowerExpr(args[1]).expr, args[1].t);
			var mappedExpr = GoExpr.GoCall(lowerExpr(callee).expr, [lambdaIterableLowering.dynamicIterableSource(args[0]), mapperExpr]);
			return {
				expr: mappedExpr,
				isStringLike: false
			};
		}

		if (isStaticCall(callee, "Lambda", [], "fold") || lambdaIterableLowering.isGeneratedCall(callee, "fold")) {
			if (args.length != 3) {
				Context.fatalError("Lambda.fold expects exactly 3 arguments", callee.pos);
			}
			var folderExpr = lambdaIterableLowering.folderAnyAdapter(lowerExpr(args[1]).expr, args[1].t);
			var foldedExpr = GoExpr.GoCall(lowerExpr(callee).expr, [
				lambdaIterableLowering.dynamicIterableSource(args[0]),
				folderExpr,
				lowerExpr(args[2]).expr
			]);
			return {
				expr: lowerNullableAwareTypeAssertExpr(foldedExpr, returnType),
				isStringLike: false
			};
		}

		if (isStaticCall(callee, "Lambda", [], "foldi") || lambdaIterableLowering.isGeneratedCall(callee, "foldi")) {
			if (args.length != 3) {
				Context.fatalError("Lambda.foldi expects exactly 3 arguments", callee.pos);
			}
			var folderExpr = lambdaIterableLowering.indexedFolderAnyAdapter(lowerExpr(args[1]).expr, args[1].t);
			var foldedExpr = GoExpr.GoCall(lowerExpr(callee).expr, [
				lambdaIterableLowering.dynamicIterableSource(args[0]),
				folderExpr,
				lowerExpr(args[2]).expr
			]);
			return {
				expr: lowerNullableAwareTypeAssertExpr(foldedExpr, returnType),
				isStringLike: false
			};
		}

		if (isStaticCall(callee, "Lambda", [], "indexOf") || lambdaIterableLowering.isGeneratedCall(callee, "indexOf")) {
			if (args.length != 2) {
				Context.fatalError("Lambda.indexOf expects exactly 2 arguments", callee.pos);
			}
			return {
				expr: GoExpr.GoCall(lowerExpr(callee).expr, [lambdaIterableLowering.dynamicIterableSource(args[0]), lowerExpr(args[1]).expr]),
				isStringLike: false
			};
		}

		if (isStaticCall(callee, "Lambda", [], "find") || lambdaIterableLowering.isGeneratedCall(callee, "find")) {
			if (args.length != 2) {
				Context.fatalError("Lambda.find expects exactly 2 arguments", callee.pos);
			}
			var predicateExpr = lambdaIterableLowering.predicateAnyAdapter(lowerExpr(args[1]).expr, args[1].t);
			var foundExpr = GoExpr.GoCall(lowerExpr(callee).expr, [lambdaIterableLowering.dynamicIterableSource(args[0]), predicateExpr]);
			return {
				expr: lowerNullableAwareTypeAssertExpr(foundExpr, returnType),
				isStringLike: false
			};
		}

		if (isStaticCall(callee, "Lambda", [], "findIndex") || lambdaIterableLowering.isGeneratedCall(callee, "findIndex")) {
			if (args.length != 2) {
				Context.fatalError("Lambda.findIndex expects exactly 2 arguments", callee.pos);
			}
			var predicateExpr = lambdaIterableLowering.predicateAnyAdapter(lowerExpr(args[1]).expr, args[1].t);
			return {
				expr: GoExpr.GoCall(lowerExpr(callee).expr, [lambdaIterableLowering.dynamicIterableSource(args[0]), predicateExpr]),
				isStringLike: false
			};
		}

		if (isStaticCall(callee, "Lambda", [], "concat") || lambdaIterableLowering.isGeneratedCall(callee, "concat")) {
			if (args.length != 2) {
				Context.fatalError("Lambda.concat expects exactly 2 arguments", callee.pos);
			}
			var concatExpr = GoExpr.GoCall(lowerExpr(callee).expr, [
				lambdaIterableLowering.dynamicIterableSource(args[0]),
				lambdaIterableLowering.dynamicIterableSource(args[1])
			]);
			return {
				expr: concatExpr,
				isStringLike: false
			};
		}

		return null;
	}

	function lowerStdIsOfTypeCall(args:Array<TypedExpr>):LoweredExpr {
		if (args.length != 2) {
			Context.fatalError("Std.isOfType expects exactly 2 arguments", Context.currentPos());
		}

		var targetType = stdIsOfTypeTargetType(args[1]);
		if (targetType == null) {
			Context.fatalError("Std.isOfType requires a type literal as the second argument", args[1].pos);
		}

		var loweredValue = lowerExprWithPrefix(args[0]);
		var loweredCheck = lowerStdIsOfTypeExpr(loweredValue.expr, args[0], targetType);
		if (loweredValue.prefix.length == 0) {
			return {
				expr: loweredCheck,
				isStringLike: false
			};
		}

		return {
			expr: GoExpr.GoCall(GoExpr.GoFuncLiteral([], ["bool"], loweredValue.prefix.concat([GoStmt.GoReturn(loweredCheck)])), []),
			isStringLike: false
		};
	}

	function stdIsOfTypeTargetType(expr:TypedExpr):Null<Type> {
		return switch (expr.expr) {
			case TTypeExpr(moduleType):
				switch (moduleType) {
					case TClassDecl(classRef):
						TInst(classRef, []);
					case TEnumDecl(enumRef):
						TEnum(enumRef, []);
					case TTypeDecl(typeRef):
						Context.follow(TType(typeRef, []));
					case TAbstract(abstractRef):
						TAbstract(abstractRef, []);
					case _:
						null;
				}
			case TMeta(_, inner):
				stdIsOfTypeTargetType(inner);
			case TParenthesis(inner):
				stdIsOfTypeTargetType(inner);
			case TCast(inner, _):
				stdIsOfTypeTargetType(inner);
			case _:
				null;
		};
	}

	function lowerStdIsOfTypeExpr(valueExpr:GoExpr, valueTypedExpr:TypedExpr, targetType:Type):GoExpr {
		if (isDynamicCatchType(targetType)) {
			if (isNullLiteralExpr(valueTypedExpr)) {
				return GoExpr.GoBoolLiteral(false);
			}
			if (isDefinitelyNonNullableType(valueTypedExpr.t)) {
				return GoExpr.GoBoolLiteral(true);
			}
			return GoExpr.GoBinary("!=", valueExpr, GoExpr.GoNil);
		}

		if (isStdClassMetaType(targetType)) {
			if (isStdClassMetaType(valueTypedExpr.t)) {
				return GoExpr.GoBinary("!=", valueExpr, GoExpr.GoNil);
			}
			if (!isAnyLikeType(valueTypedExpr.t)) {
				return GoExpr.GoBoolLiteral(false);
			}
			return stdIsOfTypeTypeSwitch(valueExpr, ["*hxrt__TypeClassValue"]);
		}

		if (isStdEnumMetaType(targetType)) {
			if (isStdEnumMetaType(valueTypedExpr.t)) {
				return GoExpr.GoBinary("!=", valueExpr, GoExpr.GoNil);
			}
			if (!isAnyLikeType(valueTypedExpr.t)) {
				return GoExpr.GoBoolLiteral(false);
			}
			return stdIsOfTypeTypeSwitch(valueExpr, ["*hxrt__TypeEnumValue"]);
		}

		if (isBoolType(targetType)) {
			if (isBoolType(valueTypedExpr.t)) {
				return GoExpr.GoBoolLiteral(true);
			}
			if (!isAnyLikeType(valueTypedExpr.t)) {
				return GoExpr.GoBoolLiteral(false);
			}
			return stdIsOfTypeTypeSwitch(valueExpr, ["bool"]);
		}

		if (isIntType(targetType)) {
			if (isIntType(valueTypedExpr.t)) {
				return GoExpr.GoBoolLiteral(true);
			}
			if (isFloatType(valueTypedExpr.t) || isBoolType(valueTypedExpr.t) || isStringType(valueTypedExpr.t)) {
				return GoExpr.GoBoolLiteral(false);
			}
			if (!isAnyLikeType(valueTypedExpr.t)) {
				return GoExpr.GoBoolLiteral(false);
			}
			return stdIsOfTypeTypeSwitch(valueExpr, [
				"int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64", "uintptr"
			]);
		}

		if (isFloatType(targetType)) {
			if (isIntType(valueTypedExpr.t) || isFloatType(valueTypedExpr.t)) {
				return GoExpr.GoBoolLiteral(true);
			}
			if (isBoolType(valueTypedExpr.t) || isStringType(valueTypedExpr.t)) {
				return GoExpr.GoBoolLiteral(false);
			}
			if (!isAnyLikeType(valueTypedExpr.t)) {
				return GoExpr.GoBoolLiteral(false);
			}
			return stdIsOfTypeTypeSwitch(valueExpr, [
				"int",
				"int8",
				"int16",
				"int32",
				"int64",
				"uint",
				"uint8",
				"uint16",
				"uint32",
				"uint64",
				"uintptr",
				"float32",
				"float64"
			]);
		}

		if (isStringType(targetType)) {
			if (isStringType(valueTypedExpr.t)) {
				return GoExpr.GoBinary("!=", valueExpr, GoExpr.GoNil);
			}
			if (!isAnyLikeType(valueTypedExpr.t)) {
				return GoExpr.GoBoolLiteral(false);
			}
			return stdIsOfTypeTypeSwitch(valueExpr, ["*string", "string"]);
		}

		if (isArrayType(targetType)) {
			if (isArrayType(valueTypedExpr.t)) {
				return GoExpr.GoBinary("!=", valueExpr, GoExpr.GoNil);
			}
			if (isAnyLikeType(valueTypedExpr.t)) {
				return stdIsOfTypeTypeSwitch(valueExpr, stdIsOfTypeArrayTypeNames());
			}
			return GoExpr.GoBoolLiteral(false);
		}

		var targetClass = classFromType(targetType);
		if (targetClass != null) {
			return stdIsOfTypeClassExpr(valueExpr, valueTypedExpr.t, targetClass);
		}

		var targetEnum = switch (Context.follow(targetType)) {
			case TEnum(enumRef, _):
				enumRef.get();
			case _:
				null;
		};
		if (targetEnum != null) {
			return stdIsOfTypeEnumExpr(valueExpr, valueTypedExpr.t, targetEnum);
		}

		// For unresolved runtime-value abstract targets (for example @:runtimeValue @:coreType),
		// align with Haxe runtime behavior by returning false instead of hard-failing compilation.
		return stdIsOfTypeUnknownTargetExpr(valueExpr, valueTypedExpr.t, targetType);
	}

	function stdIsOfTypeUnknownTargetExpr(valueExpr:GoExpr, valueType:Type, targetType:Type):GoExpr {
		var targetGoType = typeToGoType(targetType);
		if (targetGoType == "any") {
			return GoExpr.GoBoolLiteral(false);
		}

		if (!isAnyLikeType(valueType)) {
			return GoExpr.GoBoolLiteral(false);
		}

		return stdIsOfTypeTypeSwitch(valueExpr, [targetGoType]);
	}

	function stdIsOfTypeClassExpr(valueExpr:GoExpr, valueType:Type, targetClass:ClassType):GoExpr {
		if (isHaxeExceptionClass(targetClass)) {
			if (isHaxeExceptionFamilyType(valueType)) {
				return GoExpr.GoBinary("!=", valueExpr, GoExpr.GoNil);
			}
			if (!isAnyLikeType(valueType)) {
				return GoExpr.GoBoolLiteral(false);
			}
			return stdIsOfTypeTypeSwitch(valueExpr, ["*hxrt.ExceptionValue", "hxrt.ExceptionCarrier"]);
		}

		var valueClass = classFromType(valueType);
		if (valueClass != null) {
			if (inheritancePath(valueClass, targetClass) != null) {
				return GoExpr.GoBinary("!=", valueExpr, GoExpr.GoNil);
			}

			if (inheritancePath(targetClass, valueClass) != null) {
				var valueTypeName = "*" + classTypeName(valueClass);
				var targetPointerType = "*" + classTypeName(targetClass);
				return GoExpr.GoCall(GoExpr.GoFuncLiteral([{name: "hx_value", typeName: valueTypeName}], ["bool"], [
					GoStmt.GoIf(GoExpr.GoBinary("==", GoExpr.GoIdent("hx_value"), GoExpr.GoNil), [GoStmt.GoReturn(GoExpr.GoBoolLiteral(false))], null),
					GoStmt.GoRaw("_, ok := hx_value.__hx_this.(" + targetPointerType + ")"),
					GoStmt.GoReturn(GoExpr.GoIdent("ok"))
				]), [valueExpr]);
			}

			return GoExpr.GoBoolLiteral(false);
		}

		if (!isAnyLikeType(valueType)) {
			return GoExpr.GoBoolLiteral(false);
		}

		return stdIsOfTypeTypeSwitch(valueExpr, stdIsOfTypeClassTypeNames(targetClass));
	}

	function stdIsOfTypeEnumExpr(valueExpr:GoExpr, valueType:Type, targetEnum:EnumType):GoExpr {
		var valueEnum = switch (Context.follow(valueType)) {
			case TEnum(enumRef, _):
				enumRef.get();
			case _:
				null;
		};

		if (valueEnum != null) {
			return fullEnumName(valueEnum) == fullEnumName(targetEnum) ? GoExpr.GoBinary("!=", valueExpr, GoExpr.GoNil) : GoExpr.GoBoolLiteral(false);
		}

		if (!isAnyLikeType(valueType)) {
			return GoExpr.GoBoolLiteral(false);
		}

		return stdIsOfTypeTypeSwitch(valueExpr, ["*" + enumTypeName(targetEnum)]);
	}

	function stdIsOfTypeTypeSwitch(valueExpr:GoExpr, typeNames:Array<String>):GoExpr {
		return GoExpr.GoCall(GoExpr.GoFuncLiteral([{name: "hx_value", typeName: "any"}], ["bool"], [
			GoStmt.GoTypeSwitch(GoExpr.GoIdent("hx_value"), null, [
				for (typeName in typeNames)
					{
						typeName: typeName,
						body: [GoStmt.GoReturn(GoExpr.GoBoolLiteral(true))]
					}
			], [GoStmt.GoReturn(GoExpr.GoBoolLiteral(false))])
		]), [GoExpr.GoCall(GoExpr.GoIdent("any"), [valueExpr])]);
	}

	function stdIsOfTypeClassTypeNames(targetClass:ClassType):Array<String> {
		var seen = new Map<String, Bool>();
		var out = new Array<String>();
		for (candidate in projectClasses) {
			if (!hasInstanceLayout(candidate)) {
				continue;
			}
			if (inheritancePath(candidate, targetClass) != null) {
				var typeName = "*" + classTypeName(candidate);
				if (!seen.exists(typeName)) {
					seen.set(typeName, true);
					out.push(typeName);
				}
			}
		}

		var targetTypeName = "*" + classTypeName(targetClass);
		if (hasInstanceLayout(targetClass) && !seen.exists(targetTypeName)) {
			out.push(targetTypeName);
		}

		var targetPack = targetClass.pack.join(".");
		if (targetPack == "" && targetClass.name == "Class") {
			var classMarker = "*hxrt__TypeClassValue";
			if (!seen.exists(classMarker)) {
				seen.set(classMarker, true);
				out.push(classMarker);
			}
		}
		if (targetPack == "" && targetClass.name == "Enum") {
			var enumMarker = "*hxrt__TypeEnumValue";
			if (!seen.exists(enumMarker)) {
				seen.set(enumMarker, true);
				out.push(enumMarker);
			}
		}

		out.sort(function(a, b) return Reflect.compare(a, b));
		return out;
	}

	function stdIsOfTypeArrayTypeNames():Array<String> {
		return ["*hxrt.Array"];
	}

	function hasInstanceLayout(classType:ClassType):Bool {
		if (classType.isInterface) {
			return false;
		}

		if (projectSuperClass(classType) != null) {
			return true;
		}

		for (field in classType.fields.get()) {
			switch (field.kind) {
				case FVar(_, _):
					return true;
				case FMethod(_):
					if (field.name != "new") {
						return true;
					}
			}
		}

		return classType.constructor != null;
	}

	function isStaticCall(callee:TypedExpr, className:String, classPack:Array<String>, fieldName:String):Bool {
		return switch (callee.expr) {
			case TField(_, FStatic(classRef, field)): var classType = classRef.get(); classType.name == className && classType.pack.join(".") == classPack.join(".") && field.get()
					.name == fieldName;
			case _:
				false;
		};
	}

	function useStringFastpath():Bool {
		return compilationContext.buildContext.portableStringFastpathEnabled;
	}

	function useProvenConcurrencyFastpath():Bool {
		return !compilationContext.buildContext.usesEagerNativeSpecialization()
			&& compilationContext.buildContext.portableConcurrencyFastpathEnabled;
	}

	function useProvenAutoLoweringPlannerMode():Bool {
		return !compilationContext.buildContext.usesEagerNativeSpecialization()
			&& compilationContext.buildContext.autoLoweringMode != GoAutoLoweringMode.Off;
	}

	function useTypedGoConcurrencySpecialization():Bool {
		return compilationContext.buildContext.usesEagerNativeSpecialization() || useProvenConcurrencyFastpath();
	}

	function useTypedGoCollectionsSpecialization():Bool {
		return compilationContext.buildContext.usesEagerNativeSpecialization() || useProvenAutoLoweringPlannerMode();
	}

	function useTypedGoResultSpecialization():Bool {
		return compilationContext.buildContext.usesEagerNativeSpecialization() || useProvenAutoLoweringPlannerMode();
	}

	function noteProvenConcurrencyFastpathHit(pos:haxe.macro.Expr.Position):Void {
		if (!useProvenConcurrencyFastpath()) {
			return;
		}
		if (isFrameworkInternalPos(pos)) {
			return;
		}
		compilationContext.optimizerPortableConcurrencyTypedFastpathHits++;
	}

	function noteProvenConcurrencyFastpathFallback(pos:haxe.macro.Expr.Position):Void {
		if (!useProvenConcurrencyFastpath()) {
			return;
		}
		if (isFrameworkInternalPos(pos)) {
			return;
		}
		compilationContext.optimizerPortableConcurrencyTypedFastpathFallbacks++;
	}

	function noteLoweringAttempt(feature:String, kind:String, pos:haxe.macro.Expr.Position, detail:String):Void {
		noteLoweringDecision(feature, kind, "attempted", pos, detail);
	}

	function noteLoweringSuccess(feature:String, kind:String, pos:haxe.macro.Expr.Position, detail:String):Void {
		noteLoweringDecision(feature, kind, "succeeded", pos, detail);
		noteOptimizerTypedLoweringSuccess(feature);
	}

	function noteLoweringFallback(feature:String, kind:String, pos:haxe.macro.Expr.Position, detail:String):Void {
		if (shouldSuppressProvenInternalFallbackReport(feature, pos)) {
			return;
		}
		noteLoweringDecision(feature, kind, "fallback", pos, detail);
		noteOptimizerTypedLoweringFallback(feature);
		noteNativeFallback(kind, pos, detail);
	}

	function noteOptimizerTypedLoweringSuccess(feature:String):Void {
		switch (feature) {
			case "go.collections.typed":
				compilationContext.optimizerGoCollectionsTypedLowerings++;
			case "go.result.typed":
				compilationContext.optimizerGoResultTypedLowerings++;
			case _:
		}
	}

	function noteOptimizerTypedLoweringFallback(feature:String):Void {
		switch (feature) {
			case "go.collections.typed":
				compilationContext.optimizerGoCollectionsTypedFallbacks++;
			case "go.result.typed":
				compilationContext.optimizerGoResultTypedFallbacks++;
			case _:
		}
	}

	function shouldSuppressProvenInternalFallbackReport(feature:String, pos:haxe.macro.Expr.Position):Bool {
		if (!useProvenConcurrencyFastpath()) {
			return false;
		}
		if (feature != "go.concurrency.typed") {
			return false;
		}
		return isFrameworkInternalPos(pos);
	}

	function noteLoweringDecision(feature:String, kind:String, outcome:String, pos:haxe.macro.Expr.Position, detail:String):Void {
		var metadata = loweringDecisionMetadata(pos);
		compilationContext.loweringDecisionLedger.push({
			feature: feature,
			kind: kind,
			outcome: outcome,
			detail: detail,
			location: metadata.location,
			module: metadata.moduleName,
			inNativeBoundary: metadata.inNativeBoundary
		});
	}

	function noteNativeFallback(kind:String, pos:haxe.macro.Expr.Position, detail:String):Void {
		var metadata = loweringDecisionMetadata(pos);
		var violation = {
			kind: kind,
			detail: detail,
			location: metadata.location,
			module: metadata.moduleName,
			inNativeBoundary: metadata.inNativeBoundary
		};
		compilationContext.nativeFallbackEvents.push(violation);
		var hardError = compilationContext.buildContext.requiresNativeFallbackError() && !isFrameworkInternalPos(pos);
		if (hardError) {
			Context.error("Native specialization fallback is not allowed by `-D reflaxe_go_native_fallback=error`: "
				+ detail
				+ " Use `-D reflaxe_go_native_fallback=allow` to permit the semantics-safe fallback. "
				+ "The legacy `-D reflaxe_go_metal_allow_fallback` alias remains available for metal compatibility builds.",
				pos);
		}
	}

	function loweringDecisionMetadata(pos:haxe.macro.Expr.Position):{moduleName:String, inNativeBoundary:Bool, location:String} {
		var moduleName = sourceModuleRegistry.sourceModuleForPos(pos);
		var inNativeBoundary = compilationContext.buildContext.nativeBoundaryModules.indexOf(moduleName) != -1;
		var location = fallbackLocationLabel(pos, moduleName);
		return {
			moduleName: moduleName,
			inNativeBoundary: inNativeBoundary,
			location: location
		};
	}

	function fallbackLocationLabel(pos:haxe.macro.Expr.Position, moduleName:String):String {
		var line = 1;
		var location = PositionTools.toLocation(pos);
		if (location != null && location.range != null && location.range.start != null && location.range.start.line > 0) {
			line = location.range.start.line;
		}
		return normalizeModuleLabel(moduleName) + ":" + line;
	}

	function isFrameworkInternalPos(pos:haxe.macro.Expr.Position):Bool {
		var file = normalizeSourcePath(Context.getPosInfos(pos).file);
		return file.indexOf("/std/") != -1
			|| file.indexOf("/src/go/") != -1
			|| file.indexOf("/src/reflaxe/") != -1
			|| StringTools.startsWith(file, "std/")
			|| StringTools.startsWith(file, "src/go/")
			|| StringTools.startsWith(file, "src/reflaxe/");
	}

	static function normalizeSourcePath(value:String):String {
		return value == null ? "" : value.split("\\").join("/");
	}

	static function normalizeModuleLabel(value:Null<String>):String {
		if (value == null) {
			return "<unknown>";
		}
		var trimmed = StringTools.trim(value);
		return trimmed == "" ? "<unknown>" : trimmed;
	}

	function isGoChanClass(classType:ClassType):Bool {
		return classType.pack.join(".") == "go" && classType.name == "Chan";
	}

	function isGoSliceClass(classType:ClassType):Bool {
		return classType.pack.join(".") == "go" && classType.name == "Slice";
	}

	function isGoMapClass(classType:ClassType):Bool {
		return classType.pack.join(".") == "go" && classType.name == "Map";
	}

	function isGoResultClass(classType:ClassType):Bool {
		return classType.pack.join(".") == "go" && classType.name == "Result";
	}

	function goChanElementType(type:Type):Null<Type> {
		var followed = Context.follow(type);
		return switch (followed) {
			case TInst(classRef, params):
				var classType = classRef.get();
				if (isGoChanClass(classType) && params.length == 1) {
					params[0];
				} else {
					null;
				}
			case TAbstract(abstractRef, params):
				var abstractType = abstractRef.get();
				if (abstractType.pack.length == 0 && abstractType.name == "Null" && params.length == 1) {
					goChanElementType(params[0]);
				} else {
					null;
				}
			case TMono(ref):
				var resolved = ref.get();
				resolved == null ? null : goChanElementType(resolved);
			case _:
				null;
		};
	}

	function goSliceElementType(type:Type):Null<Type> {
		var followed = Context.follow(type);
		return switch (followed) {
			case TInst(classRef, params):
				var classType = classRef.get();
				if (isGoSliceClass(classType) && params.length == 1) {
					params[0];
				} else {
					null;
				}
			case TAbstract(abstractRef, params):
				var abstractType = abstractRef.get();
				if (abstractType.pack.length == 0 && abstractType.name == "Null" && params.length == 1) {
					goSliceElementType(params[0]);
				} else {
					null;
				}
			case TMono(ref):
				var resolved = ref.get();
				resolved == null ? null : goSliceElementType(resolved);
			case _:
				null;
		};
	}

	function goMapTypePair(type:Type):Null<{keyType:Type, valueType:Type}> {
		var followed = Context.follow(type);
		return switch (followed) {
			case TInst(classRef, params):
				var classType = classRef.get();
				if (isGoMapClass(classType) && params.length == 2) {
					{
						keyType: params[0],
						valueType: params[1]
					};
				} else {
					null;
				}
			case TAbstract(abstractRef, params):
				var abstractType = abstractRef.get();
				if (abstractType.pack.length == 0 && abstractType.name == "Null" && params.length == 1) {
					goMapTypePair(params[0]);
				} else {
					null;
				}
			case TMono(ref):
				var resolved = ref.get();
				resolved == null ? null : goMapTypePair(resolved);
			case _:
				null;
		};
	}

	function goResultElementType(type:Type):Null<Type> {
		var followed = Context.follow(type);
		return switch (followed) {
			case TInst(classRef, params):
				var classType = classRef.get();
				if (isGoResultClass(classType) && params.length == 1) {
					params[0];
				} else {
					null;
				}
			case TAbstract(abstractRef, params):
				var abstractType = abstractRef.get();
				if (abstractType.pack.length == 0 && abstractType.name == "Null" && params.length == 1) {
					goResultElementType(params[0]);
				} else {
					null;
				}
			case TMono(ref):
				var resolved = ref.get();
				resolved == null ? null : goResultElementType(resolved);
			case _:
				null;
		};
	}

	function isMonomorphizableNativeElementType(elementGoType:String):Bool {
		return elementGoType != null && elementGoType != "" && elementGoType != "any";
	}

	function isMonomorphizableNativeChanElementType(elementGoType:String):Bool {
		return isMonomorphizableNativeElementType(elementGoType);
	}

	function goChanElementEligibility(type:Type, missingMessage:String):GoNativeTypeEligibilityResult {
		return nativeTypeEligibility(goChanElementType(type), GoNativeEligibilityRole.ChanElement, missingMessage);
	}

	function goChanElementGoType(type:Type):Null<String> {
		var eligibility = goChanElementEligibility(type, "Could not resolve go.Chan element type for native specialization.");
		return eligibility.eligible ? eligibility.goType : null;
	}

	function goSliceElementEligibility(type:Type, missingMessage:String):GoNativeTypeEligibilityResult {
		return nativeTypeEligibility(goSliceElementType(type), GoNativeEligibilityRole.SliceElement, missingMessage);
	}

	function goSliceElementGoType(type:Type):Null<String> {
		var eligibility = goSliceElementEligibility(type, "Could not resolve go.Slice element type for native specialization.");
		return eligibility.eligible ? eligibility.goType : null;
	}

	function goMapTypePairGoTypes(type:Type):Null<NativeMapTypePair> {
		var pair = goMapTypePair(type);
		if (pair == null) {
			return null;
		}
		var keyEligibility = nativeTypeEligibility(pair.keyType, GoNativeEligibilityRole.MapKey,
			"Could not resolve go.Map key type for native specialization.");
		if (!keyEligibility.eligible || keyEligibility.goType == null) {
			return null;
		}
		var valueEligibility = nativeTypeEligibility(pair.valueType, GoNativeEligibilityRole.MapValue,
			"Could not resolve go.Map value type for native specialization.");
		if (!valueEligibility.eligible || valueEligibility.goType == null) {
			return null;
		}
		return {
			keyGoType: keyEligibility.goType,
			valueGoType: valueEligibility.goType
		};
	}

	function goResultElementEligibility(type:Type, missingMessage:String):GoNativeTypeEligibilityResult {
		return nativeTypeEligibility(goResultElementType(type), GoNativeEligibilityRole.ResultElement, missingMessage);
	}

	function goResultElementGoType(type:Type):Null<String> {
		var eligibility = goResultElementEligibility(type, "Could not resolve go.Result<T> element type for native specialization.");
		return eligibility.eligible ? eligibility.goType : null;
	}

	function nativeTypeEligibility(type:Null<Type>, role:GoNativeEligibilityRole, missingMessage:String):GoNativeTypeEligibilityResult {
		if (type == null) {
			return {
				eligible: false,
				goType: null,
				reasonCode: "missing_type",
				reason: missingMessage
			};
		}
		return GoNativeTypeEligibility.resolve(type, role, classTypeName, enumTypeName);
	}

	function withEligibilityReason(base:String, eligibility:GoNativeTypeEligibilityResult):String {
		var reason = eligibility.reason;
		if (reason == null || StringTools.trim(reason) == "") {
			return base;
		}
		var prefix = StringTools.endsWith(base, ".") ? base.substr(0, base.length - 1) : base;
		return prefix + ": " + reason;
	}

	function registerNativeChanElementGoType(elementGoType:String):Void {
		if (!useTypedGoConcurrencySpecialization()) {
			return;
		}
		if (!isMonomorphizableNativeChanElementType(elementGoType)) {
			return;
		}
		requiredNativeChanElementTypes.set(elementGoType, true);
	}

	function registerNativeSliceElementGoType(elementGoType:String):Void {
		if (!useTypedGoCollectionsSpecialization()) {
			return;
		}
		if (!isMonomorphizableNativeElementType(elementGoType)) {
			return;
		}
		requiredNativeSliceElementTypes.set(elementGoType, true);
	}

	function registerNativeMapTypePair(keyGoType:String, valueGoType:String):Void {
		if (!useTypedGoCollectionsSpecialization()) {
			return;
		}
		if (!isMonomorphizableNativeElementType(keyGoType) || !isMonomorphizableNativeElementType(valueGoType)) {
			return;
		}
		var signature = nativeMapTypeSignature(keyGoType, valueGoType);
		requiredNativeMapTypePairs.set(signature, {
			keyGoType: keyGoType,
			valueGoType: valueGoType
		});
	}

	function registerNativeResultElementGoType(elementGoType:String):Void {
		if (!useTypedGoResultSpecialization()) {
			return;
		}
		if (!isMonomorphizableNativeElementType(elementGoType)) {
			return;
		}
		requiredNativeResultElementTypes.set(elementGoType, true);
	}

	function nativeTypeHash(value:String):String {
		var hash = 0x811C9DC5;
		for (index in 0...value.length) {
			hash ^= value.charCodeAt(index);
			hash *= 0x01000193;
		}
		return StringTools.hex(hash, 8).toLowerCase();
	}

	function nativeTypeSuffix(typeKey:String):String {
		var normalized = GoNaming.normalizeIdent(typeKey);
		if (normalized == "" || normalized == "hx_tmp") {
			normalized = "t";
		}
		return normalized + "_" + nativeTypeHash(typeKey);
	}

	function nativeMapTypeSignature(keyGoType:String, valueGoType:String):String {
		return keyGoType + "__" + valueGoType;
	}

	function nativeChanShimName(base:String, elementGoType:String):String {
		return base + "__" + nativeTypeSuffix(elementGoType);
	}

	function nativeSliceShimName(base:String, elementGoType:String):String {
		return base + "__" + nativeTypeSuffix(elementGoType);
	}

	function nativeMapShimName(base:String, keyGoType:String, valueGoType:String):String {
		return base + "__" + nativeTypeSuffix(nativeMapTypeSignature(keyGoType, valueGoType));
	}

	function nativeResultShimName(base:String, elementGoType:String):String {
		return base + "__" + nativeTypeSuffix(elementGoType);
	}

	function shouldAssertGenericCallResult(callee:TypedExpr, returnType:Type):Bool {
		if (typeToGoType(returnType) == "any") {
			return false;
		}

		return switch (callee.expr) {
			case TField(_, FInstance(classRef, _, field)):
				var classType = classRef.get();
				var pack = classType.pack.join(".");
				var fieldName = field.get().name;
				if (classType.params.length > 0) {
					true;
				} else if (pack == "haxe.ds") {
					if ((classType.name == "IntMap" || classType.name == "StringMap" || classType.name == "ObjectMap" || classType.name == "EnumValueMap")
						&& fieldName == "get") {
						true;
					} else if (classType.name == "List" && (fieldName == "pop" || fieldName == "first" || fieldName == "last")) {
						true;
					} else {
						false;
					}
				} else if (pack == "go"
					&& classType.name == "Chan"
					&& (fieldName == "recv" || fieldName == "recvOr" || fieldName == "tryRecv")) {
					true;
				} else {
					false;
				}
			case _:
				false;
		};
	}

	/**
		What: Adapts imported Go string results to haxe.go's pointer-backed string carrier.
		Why: Ordinary Go APIs return non-nullable string values, but framework externs
		may explicitly return `Null<String>` as an already pointer-backed value. Passing
		that nullable result through `Std.string` would turn native nil into `"null"`.
		How: Preserve explicit nullable-string results unchanged and normalize only the
		non-nullable imported Go string surface.
	**/
	function normalizeExternStringCallResult(callee:TypedExpr, returnType:Type, callExpr:GoExpr):GoExpr {
		if (!isStringType(returnType)) {
			return callExpr;
		}
		if (!isGoImportExternCall(callee)) {
			return callExpr;
		}
		var nullableInner = nullableInnerType(returnType);
		if (nullableInner != null && isStringType(nullableInner)) {
			return callExpr;
		}
		return GoExpr.GoCall(GoExpr.GoIdent("hxrt.StdString"), [callExpr]);
	}

	function normalizeExternCallArg(callee:TypedExpr, argExpr:GoExpr, paramType:Null<Type>, returnType:Type):GoExpr {
		if (paramType == null || (!isExternValueErrorCall(callee, returnType) && !isExternTupleReturnCall(callee, returnType))) {
			return argExpr;
		}
		if (isStringType(paramType)) {
			return GoExpr.GoUnary("*", GoExpr.GoCall(GoExpr.GoIdent("hxrt.StdString"), [argExpr]));
		}
		return argExpr;
	}

	function lowerExternTupleReturnExpr(callee:TypedExpr, returnType:Type, callExpr:GoExpr):Null<GoExpr> {
		if (!isExternTupleReturnCall(callee, returnType)) {
			return null;
		}
		var carrier = tupleReturnCarrierClass(returnType);
		if (carrier == null) {
			Context.fatalError("@:go.tupleReturn extern calls must return a concrete Haxe carrier class", callee.pos);
			return null;
		}
		var constructorArgs = tupleReturnCarrierConstructorArgTypes(carrier, returnType, callee.pos);
		if (constructorArgs == null || constructorArgs.length == 0) {
			Context.fatalError("@:go.tupleReturn carrier must have a constructor with one parameter per Go result value", callee.pos);
			return null;
		}

		var tempNames = [for (_ in constructorArgs) freshTempName("hx_tuple")];
		var body = new Array<GoStmt>();
		body.push(GoStmt.GoMultiAssign(tempNames, callExpr, true));

		var loweredArgs = new Array<GoExpr>();
		for (index in 0...constructorArgs.length) {
			loweredArgs.push(coerceExternTupleReturnValue(GoExpr.GoIdent(tempNames[index]), constructorArgs[index]));
		}
		body.push(GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent(constructorSymbol(carrier)), loweredArgs)));
		return GoExpr.GoCall(GoExpr.GoFuncLiteral([], [typeToGoType(returnType)], body), []);
	}

	function tupleReturnCarrierClass(returnType:Type):Null<ClassType> {
		return switch (Context.follow(returnType)) {
			case TInst(classRef, _):
				classRef.get();
			case _:
				null;
		};
	}

	function tupleReturnCarrierConstructorArgTypes(carrier:ClassType, returnType:Type, pos:haxe.macro.Expr.Position):Null<Array<Type>> {
		if (carrier.constructor == null) {
			return [];
		}
		var ctor = carrier.constructor.get();
		return switch (Context.follow(ctor.type)) {
			case TFun(args, _):
				[for (arg in args) arg.t];
			case _:
				Context.fatalError("@:go.tupleReturn carrier constructor must be a function", pos);
				null;
		};
	}

	function coerceExternTupleReturnValue(value:GoExpr, targetType:Type):GoExpr {
		if (isStringType(targetType)) {
			return GoExpr.GoCall(GoExpr.GoIdent("hxrt.StdString"), [value]);
		}
		if (isGoErrorType(targetType)) {
			return GoExpr.GoCall(GoExpr.GoFuncLiteral([{name: "err", typeName: "error"}], ["*go___Error"], [
				GoStmt.GoIf(GoExpr.GoBinary("==", GoExpr.GoIdent("err"), GoExpr.GoNil), [GoStmt.GoReturn(GoExpr.GoNil)], null),
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("New_go___Error"), [
					GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("err"), "Error"), [])])
				]))
			]), [value]);
		}
		return value;
	}

	function isExternValueErrorCall(callee:TypedExpr, returnType:Type):Bool {
		if (goResultElementType(returnType) == null) {
			return false;
		}
		return switch (callee.expr) {
			case TField(_, FStatic(classRef, fieldRef)): var classType = classRef.get(); classType.isExtern && externClassImportPath(classType) != null && hasExternValueErrorMeta(fieldRef.get());
			case TField(_, FInstance(classRef, _, fieldRef)): var classType = classRef.get(); classType.isExtern && externClassImportPath(classType) != null && hasExternValueErrorMeta(fieldRef.get());
			case TMeta(_, inner):
				isExternValueErrorCall(inner, returnType);
			case TParenthesis(inner):
				isExternValueErrorCall(inner, returnType);
			case TCast(inner, _):
				isExternValueErrorCall(inner, returnType);
			case _:
				false;
		};
	}

	function isExternTupleReturnCall(callee:TypedExpr, returnType:Type):Bool {
		if (tupleReturnCarrierClass(returnType) == null) {
			return false;
		}
		return switch (callee.expr) {
			case TField(_, FStatic(classRef, fieldRef)): var classType = classRef.get(); classType.isExtern && externClassImportPath(classType) != null && hasExternTupleReturnMeta(fieldRef.get());
			case TField(_, FInstance(classRef, _, fieldRef)): var classType = classRef.get(); classType.isExtern && externClassImportPath(classType) != null && hasExternTupleReturnMeta(fieldRef.get());
			case TMeta(_, inner):
				isExternTupleReturnCall(inner, returnType);
			case TParenthesis(inner):
				isExternTupleReturnCall(inner, returnType);
			case TCast(inner, _):
				isExternTupleReturnCall(inner, returnType);
			case _:
				false;
		};
	}

	function isGoImportExternCall(callee:TypedExpr):Bool {
		return switch (callee.expr) {
			case TField(_, FStatic(classRef, _)): var classType = classRef.get(); classType.isExtern && externClassImportPath(classType) != null;
			case TField(_, FInstance(classRef, _, _)): var classType = classRef.get(); classType.isExtern && externClassImportPath(classType) != null;
			case TMeta(_, inner):
				isGoImportExternCall(inner);
			case TParenthesis(inner):
				isGoImportExternCall(inner);
			case TCast(inner, _):
				isGoImportExternCall(inner);
			case _:
				false;
		};
	}

	function isSuperCtorCall(callee:TypedExpr):Bool {
		return switch (callee.expr) {
			case TConst(TSuper):
				true;
			case TMeta(_, inner):
				isSuperCtorCall(inner);
			case TParenthesis(inner):
				isSuperCtorCall(inner);
			case TCast(inner, _):
				isSuperCtorCall(inner);
			case _:
				false;
		};
	}

	function isSuperTarget(target:TypedExpr):Bool {
		return switch (target.expr) {
			case TConst(TSuper):
				true;
			case TMeta(_, inner):
				isSuperTarget(inner);
			case TParenthesis(inner):
				isSuperTarget(inner);
			case TCast(inner, _):
				isSuperTarget(inner);
			case _:
				false;
		};
	}

	function isMethodField(field:ClassField):Bool {
		return switch (field.kind) {
			case FMethod(MethDynamic):
				false;
			case FMethod(_):
				true;
			case _:
				false;
		};
	}

	/**
		What: Decide whether an instance access must use the generated Haxe virtual
		receiver.

		Why: A field inherited from source-owned std can be resolved against its base
		declaration even when a future concrete receiver has an individually approved
		compiler authority. Such a synthetic carrier may not have `__hx_this`, so
		classifying only the declaring class could emit invalid Go.

		How: Give concrete compiler-owned receiver authority precedence, then apply the
		ordinary source-class, method, and global leaf checks.
	**/
	function shouldUseVirtualDispatch(classType:ClassType, field:ClassField, receiverType:Type):Bool {
		switch (Context.follow(receiverType)) {
			case TInst(receiverRef, _):
				if (GoStdlibOwnership.isCompilerOwnedAuthority(fullClassName(receiverRef.get()))) {
					return false;
				}
			case _:
		}
		if (!isProjectClass(classType)) {
			return false;
		}
		var className = fullClassName(classType);
		if (GoStdlibOwnership.isCompilerOwnedAuthority(className)) {
			return false;
		}
		if (!isMethodField(field)) {
			return false;
		}
		if (globalLeafReceiverTypes.exists("*" + classTypeName(classType))) {
			return false;
		}
		return field.name != "new";
	}

	function callParamType(calleeType:Type, index:Int):Null<Type> {
		var followed = Context.follow(calleeType);
		return switch (followed) {
			case TFun(args, _): index >= 0 && index < args.length ? args[index].t : null;
			case _:
				null;
		};
	}

	/**
		What
		- Finds the parameter type used by the emitted class method declaration before
		  Haxe substitutes concrete receiver type arguments at the call site.

		Why
		- Haxe permits `List<Int>.filter(Int -> Bool)`, while the shared Go method is
		  emitted as `filter(func(any) bool)`. Go function types are invariant, so the
		  otherwise valid typed callback needs a small representation bridge.

		How
		- Read the original instance or static field signature and leave anonymous or
		  dynamic calls alone because they do not have a shared nominal declaration.
	**/
	function emittedCallParamType(callee:TypedExpr, index:Int):Null<Type> {
		return switch (callee.expr) {
			case TField(_, FInstance(_, _, fieldRef)) | TField(_, FStatic(_, fieldRef)):
				callParamType(fieldRef.get().type, index);
			case TMeta(_, inner) | TParenthesis(inner) | TCast(inner, _):
				emittedCallParamType(inner, index);
			case _:
				null;
		};
	}

	/**
		What
		- Adapts a concrete Haxe function value to the erased function signature used
		  by a shared generated Go method.

		Why
		- A value such as `func(int) bool` cannot be passed directly where Go expects
		  `func(any) bool`, even though Haxe has already proved the generic call safe.

		How
		- Wrap only differing function signatures with matching arity, assert erased
		  inputs back to their concrete Haxe types, and keep the erased return value.
		  Ordinary values and already-compatible function types pass through unchanged.
	**/
	function adaptErasedFunctionCallArg(expr:GoExpr, concreteType:Type, emittedType:Type):GoExpr {
		if (typeToGoType(concreteType) == typeToGoType(emittedType)) {
			return expr;
		}

		return switch ([Context.follow(concreteType), Context.follow(emittedType)]) {
			case [TFun(concreteArgs, concreteReturn), TFun(emittedArgs, emittedReturn)] if (concreteArgs.length == emittedArgs.length):
				var params = new Array<GoParam>();
				var callArgs = new Array<GoExpr>();
				for (index in 0...emittedArgs.length) {
					var name = freshTempName("hx_erased_callback_arg");
					params.push({name: name, typeName: GoType.parse(typeToGoType(emittedArgs[index].t))});
					var callArg:GoExpr = GoExpr.GoIdent(name);
					if (typeToGoType(emittedArgs[index].t) == "any" && typeToGoType(concreteArgs[index].t) != "any") {
						callArg = lowerNullableAwareTypeAssertExpr(callArg, concreteArgs[index].t);
					}
					callArgs.push(callArg);
				}

				var callbackCall = GoExpr.GoCall(expr, callArgs);
				var adapter:GoExpr;
				if (isVoidType(emittedReturn)) {
					adapter = GoExpr.GoFuncLiteral(params, [], [GoStmt.GoExprStmt(callbackCall)]);
				} else {
					var returnExpr = callbackCall;
					if (typeToGoType(concreteReturn) == "any" && typeToGoType(emittedReturn) != "any") {
						returnExpr = lowerNullableAwareTypeAssertExpr(returnExpr, emittedReturn);
					}
					adapter = GoExpr.GoFuncLiteral(params, [GoType.parse(typeToGoType(emittedReturn))], [GoStmt.GoReturn(returnExpr)]);
				}
				adapter;
			case _:
				expr;
		};
	}

	/**
		What: Lowers one Haxe binary operation with representation-aware operands.
		Why: Go's operators do not preserve every Haxe rule: erased generic values can
		contain pointer-backed strings or non-comparable carriers, and direct interface
		equality can therefore return the wrong answer or panic.
		How: Keep statically typed operations native, select the established string/null/
		numeric paths, and route only erased equality through `hxrt.HaxeEqual`.
	**/
	function lowerBinop(op:Binop, left:TypedExpr, right:TypedExpr, resultType:Type):LoweredExpr {
		var leftLowered = lowerExpr(left);
		var rightLowered = lowerExpr(right);
		var leftStoredArrayElement = isSharedArrayElementExpr(left);
		var rightStoredArrayElement = isSharedArrayElementExpr(right);
		var stringMode = leftLowered.isStringLike || rightLowered.isStringLike || isStringType(left.t) || isStringType(right.t);
		var leftIsNull = isNullLiteralExpr(left) || isNilExpr(leftLowered.expr);
		var rightIsNull = isNullLiteralExpr(right) || isNilExpr(rightLowered.expr);
		var nullComparison = leftIsNull || rightIsNull;
		var impossiblePrimitiveNullComparison = nullComparison
			&& ((leftIsNull && isDefinitelyNonNullableType(right.t)) || (rightIsNull && isDefinitelyNonNullableType(left.t)));
		var optionalPrimitiveLocalNullComparison = nullComparison
			&& ((leftIsNull && isOptionalPrimitiveLocalExpr(right)) || (rightIsNull && isOptionalPrimitiveLocalExpr(left)));
		if (optionalPrimitiveLocalNullComparison) {
			if (leftIsNull) {
				var rawRight = lowerOptionalPrimitiveNilComparisonExpr(right);
				if (rawRight != null) {
					rightLowered = {expr: rawRight, isStringLike: rightLowered.isStringLike};
				}
			}
			if (rightIsNull) {
				var rawLeft = lowerOptionalPrimitiveNilComparisonExpr(left);
				if (rawLeft != null) {
					leftLowered = {expr: rawLeft, isStringLike: leftLowered.isStringLike};
				}
			}
		}
		var coerceNullableOperands = switch (op) {
			case OpEq | OpNotEq:
				false;
			case _:
				!nullComparison;
		};
		var leftExprForOperator = leftStoredArrayElement
			&& coerceNullableOperands ? coerceStoredArrayElementExpr(leftLowered.expr,
				left.t) : (coerceNullableOperands ? coerceNullablePrimitiveOperandForUse(leftLowered.expr, left) : leftLowered.expr);
		var rightExprForOperator = rightStoredArrayElement
			&& coerceNullableOperands ? coerceStoredArrayElementExpr(rightLowered.expr,
				right.t) : (coerceNullableOperands ? coerceNullablePrimitiveOperandForUse(rightLowered.expr, right) : rightLowered.expr);
		var useStringEquality = stringMode && (!nullComparison || isStringType(left.t) || isStringType(right.t));
		var typedStringOps = isStringType(left.t) && isStringType(right.t) && !leftStoredArrayElement && !rightStoredArrayElement;
		var storedArrayEquality = (op == OpEq || op == OpNotEq) && (leftStoredArrayElement || rightStoredArrayElement);
		var erasedEquality = (op == OpEq || op == OpNotEq) && !nullComparison && (isAnyLikeType(left.t) || isAnyLikeType(right.t));
		var anyNullComparison = nullComparison && (isAnyLikeType(left.t) || isAnyLikeType(right.t));
		var floatMode = isFloatType(left.t) || isFloatType(right.t) || isFloatType(resultType) || isNullableFloatType(left.t)
			|| isNullableFloatType(right.t) || isNullableFloatType(resultType);
		var int32Mode = isInt32SemanticType(left.t, left.pos)
			|| isInt32SemanticType(right.t, right.pos)
			|| isInt32SemanticType(resultType, left.pos);
		if (floatMode) {
			int32Mode = false;
		}

		return switch (op) {
			case OpAdd if (stringMode):
				var leftStringExpr = leftStoredArrayElement ? lowerSharedArrayElementStorageExpr(left) : leftLowered.expr;
				var rightStringExpr = rightStoredArrayElement ? lowerSharedArrayElementStorageExpr(right) : rightLowered.expr;
				{
					expr: GoExpr.GoCall(GoExpr.GoIdent(typedStringOps ? "hxrt.StringConcatStringPtr" : "hxrt.StringConcatAny"),
						[leftStringExpr, rightStringExpr]),
					isStringLike: true
				};
			case OpEq if (useStringEquality):
				{
					expr: GoExpr.GoCall(GoExpr.GoIdent(typedStringOps ? "hxrt.StringEqualStringPtr" : "hxrt.StringEqualAny"),
						[leftLowered.expr, rightLowered.expr]),
					isStringLike: false
				};
			case OpNotEq if (useStringEquality):
				{
					expr: GoExpr.GoUnary("!",
						GoExpr.GoCall(GoExpr.GoIdent(typedStringOps ? "hxrt.StringEqualStringPtr" : "hxrt.StringEqualAny"),
							[leftLowered.expr, rightLowered.expr])),
					isStringLike: false
				};
			case OpEq if (impossiblePrimitiveNullComparison):
				{
					expr: GoExpr.GoBoolLiteral(false),
					isStringLike: false
				};
			case OpNotEq if (impossiblePrimitiveNullComparison):
				{
					expr: GoExpr.GoBoolLiteral(true),
					isStringLike: false
				};
			case OpEq if (anyNullComparison):
				if (leftIsNull && rightIsNull) {
					{
						expr: GoExpr.GoBoolLiteral(true),
						isStringLike: false
					};
				} else {
					var targetExpr = leftIsNull ? rightLowered.expr : leftLowered.expr;
					{
						expr: GoExpr.GoCall(GoExpr.GoIdent("hxrt.AnyEqualsNull"), [targetExpr]),
						isStringLike: false
					};
				}
			case OpNotEq if (anyNullComparison):
				if (leftIsNull && rightIsNull) {
					{
						expr: GoExpr.GoBoolLiteral(false),
						isStringLike: false
					};
				} else {
					var targetExpr = leftIsNull ? rightLowered.expr : leftLowered.expr;
					{
						expr: GoExpr.GoUnary("!", GoExpr.GoCall(GoExpr.GoIdent("hxrt.AnyEqualsNull"), [targetExpr])),
						isStringLike: false
					};
				}
			case OpEq if (erasedEquality):
				requiresEqualitySurface = true;
				{
					expr: GoExpr.GoCall(GoExpr.GoIdent("hxrt.HaxeEqual"), [leftLowered.expr, rightLowered.expr]),
					isStringLike: false
				};
			case OpNotEq if (erasedEquality):
				requiresEqualitySurface = true;
				{
					expr: GoExpr.GoUnary("!", GoExpr.GoCall(GoExpr.GoIdent("hxrt.HaxeEqual"), [leftLowered.expr, rightLowered.expr])),
					isStringLike: false
				};
			case OpEq if (storedArrayEquality):
				requiresEqualitySurface = true;
				{
					expr: GoExpr.GoCall(GoExpr.GoIdent("hxrt.HaxeEqual"), [leftLowered.expr, rightLowered.expr]),
					isStringLike: false
				};
			case OpNotEq if (storedArrayEquality):
				requiresEqualitySurface = true;
				{
					expr: GoExpr.GoUnary("!", GoExpr.GoCall(GoExpr.GoIdent("hxrt.HaxeEqual"), [leftLowered.expr, rightLowered.expr])),
					isStringLike: false
				};
			case OpAdd | OpSub | OpMult | OpDiv if (floatMode):
				{
					expr: GoExpr.GoBinary(binopSymbol(op), floatOperandExpr(leftExprForOperator, left.t, left),
						floatOperandExpr(rightExprForOperator, right.t, right)),
					isStringLike: false
				};
			case OpMod if (floatMode):
				{
					expr: GoExpr.GoCall(GoExpr.GoIdent("hxrt.FloatMod"), [
						floatOperandExpr(leftExprForOperator, left.t, left),
						floatOperandExpr(rightExprForOperator, right.t, right)
					]),
					isStringLike: false
				};
			case OpUShr if (int32Mode):
				var int32Left = coerceNullableIntOperandExpr(leftExprForOperator, left.t, left);
				var int32Right = coerceNullableIntOperandExpr(rightExprForOperator, right.t, right);
				{
					expr: lowerHaxeInt32BinopExpr(op, int32Left, int32Right),
					isStringLike: false
				};
			case OpAdd | OpSub | OpMult | OpMod | OpAnd | OpOr | OpXor | OpShl | OpShr if (int32Mode):
				var int32Left = coerceNullableIntOperandExpr(leftExprForOperator, left.t, left);
				var int32Right = coerceNullableIntOperandExpr(rightExprForOperator, right.t, right);
				{
					expr: lowerHaxeInt32BinopExpr(op, int32Left, int32Right),
					isStringLike: false
				};
			case OpUShr:
				var ushrInner = GoExpr.GoBinary(">>", GoExpr.GoCall(GoExpr.GoIdent("uint32"), [leftLowered.expr]),
					GoExpr.GoCall(GoExpr.GoIdent("uint"), [rightLowered.expr]));
				var ushrCast = scalarGoType(resultType) == "int32" ? "int32" : "int";
				{
					expr: GoExpr.GoCall(GoExpr.GoIdent(ushrCast), [ushrInner]),
					isStringLike: false
				};
			case _:
				{
					expr: GoExpr.GoBinary(binopSymbol(op), leftExprForOperator, rightExprForOperator),
					isStringLike: isStringType(resultType)
				};
		};
	}

	function lowerAssignOpExpr(op:Binop, leftExpr:GoExpr, rightExpr:GoExpr, leftType:Type, rightType:Type, ?sourcePos:haxe.macro.Expr.Position,
			?rightUsesErasedArrayStorage:Bool = false):GoExpr {
		if (op == OpAssign) {
			return rightExpr;
		}
		if (op == OpAdd && (isStringType(leftType) || isStringType(rightType))) {
			var typedStringOps = isStringType(leftType) && isStringType(rightType) && !rightUsesErasedArrayStorage;
			return GoExpr.GoCall(GoExpr.GoIdent(typedStringOps ? "hxrt.StringConcatStringPtr" : "hxrt.StringConcatAny"), [leftExpr, rightExpr]);
		}
		if ((isInt32SemanticType(leftType, sourcePos) || isInt32SemanticType(rightType, sourcePos))
			&& !isFloatType(leftType)
			&& !isFloatType(rightType)
			&& !isNullableFloatType(leftType)
			&& !isNullableFloatType(rightType)
			&& (op == OpAdd || op == OpSub || op == OpMult || op == OpMod || op == OpAnd || op == OpOr || op == OpXor || op == OpShl || op == OpShr
				|| op == OpUShr)) {
			var int32Left = coerceNullableIntOperandExpr(leftExpr, leftType);
			var int32Right = coerceNullableIntOperandExpr(rightExpr, rightType);
			return lowerHaxeInt32BinopExpr(op, int32Left, int32Right);
		}
		if ((op == OpAdd || op == OpSub || op == OpMult || op == OpDiv)
			&& (isFloatType(leftType) || isNullableFloatType(leftType) || isFloatType(rightType) || isNullableFloatType(rightType))) {
			return GoExpr.GoBinary(binopSymbol(op), floatOperandExpr(leftExpr, leftType), floatOperandExpr(rightExpr, rightType));
		}
		if (op == OpMod
			&& (isFloatType(leftType) || isFloatType(rightType) || isNullableFloatType(leftType) || isNullableFloatType(rightType))) {
			return GoExpr.GoCall(GoExpr.GoIdent("hxrt.FloatMod"), [floatOperandExpr(leftExpr, leftType), floatOperandExpr(rightExpr, rightType)]);
		}
		if (op == OpUShr) {
			var ushrInner = GoExpr.GoBinary(">>", GoExpr.GoCall(GoExpr.GoIdent("uint32"), [leftExpr]), GoExpr.GoCall(GoExpr.GoIdent("uint"), [rightExpr]));
			var ushrCast = scalarGoType(leftType) == "int32" ? "int32" : "int";
			return GoExpr.GoCall(GoExpr.GoIdent(ushrCast), [ushrInner]);
		}
		return GoExpr.GoBinary(binopSymbol(op), leftExpr, rightExpr);
	}

	function lowerHaxeInt32BinopExpr(op:Binop, leftExpr:GoExpr, rightExpr:GoExpr):GoExpr {
		return GoExprOperatorOps.lowerHaxeInt32BinopExpr(op, leftExpr, rightExpr);
	}

	function coerceNullableIntOperandExpr(expr:GoExpr, operandType:Type, ?operand:TypedExpr):GoExpr {
		if (!isNullableIntType(operandType)) {
			return expr;
		}
		if (operand != null && (exprUsesNarrowedPrimitiveStorage(operand) || nonNullPrimitiveExprGoType(operand) != null)) {
			return expr;
		}
		return GoExpr.GoCall(GoExpr.GoIdent("hxrt.IntFromNullableAny"), [expr]);
	}

	function coerceNullablePrimitiveOperandForUse(expr:GoExpr, operand:TypedExpr):GoExpr {
		if (!isNullablePrimitiveType(operand.t)) {
			return expr;
		}
		if (isNullableIntType(operand.t)) {
			return coerceNullableIntOperandExpr(expr, operand.t, operand);
		}
		if (isNullableFloatType(operand.t)) {
			return coerceNullableFloatOperandExpr(expr, operand.t, operand);
		}
		if (isNullableBoolType(operand.t)) {
			if (exprUsesNarrowedPrimitiveStorage(operand) || nonNullPrimitiveExprGoType(operand) != null) {
				return expr;
			}
			return lowerNilSafeTypeAssertExpr(expr, "bool");
		}
		return expr;
	}

	function coerceNullableFloatOperandExpr(expr:GoExpr, operandType:Type, ?operand:TypedExpr):GoExpr {
		if (!isNullableFloatType(operandType)) {
			return expr;
		}
		if (operand != null && (exprUsesNarrowedPrimitiveStorage(operand) || nonNullPrimitiveExprGoType(operand) != null)) {
			return expr;
		}
		return lowerNilSafeTypeAssertExpr(expr, "float64");
	}

	function shouldForceAnyCoerce(fromType:Type, toType:Type):Bool {
		if (!isNullablePrimitiveType(fromType) || isNullablePrimitiveType(toType)) {
			return false;
		}
		return typeToGoType(toType) != "any";
	}

	function coerceAnyExprToType(expr:GoExpr, fromType:Type, toType:Type, ?fromAnyOverride:Bool = false):GoExpr {
		var fromGoType = typeToGoType(fromType);
		var toGoType = typeToGoType(toType);
		if ((!fromAnyOverride && fromGoType != "any") || toGoType == "any") {
			return expr;
		}
		// A nullable primitive destination uses `any` storage so Go nil can survive.
		// Non-null primitive destinations still use nil-safe coercion below.
		if (isNullablePrimitiveType(toType)) {
			return expr;
		}
		if (isIntType(toType) || isHaxeInt32Type(toType)) {
			return GoExpr.GoCall(GoExpr.GoIdent("hxrt.IntFromNullableAny"), [expr]);
		}
		if (isFloatType(toType) || isNullableFloatType(toType)) {
			return lowerNilSafeTypeAssertExpr(expr, "float64");
		}
		if (isBoolType(toType) || isNullableBoolType(toType)) {
			return lowerNilSafeTypeAssertExpr(expr, "bool");
		}
		if (isStringType(toType)) {
			return GoExpr.GoCall(GoExpr.GoIdent("hxrt.StdString"), [expr]);
		}
		return expr;
	}

	/**
		What: Recovers one statically known value from the portable Array's erased,
		nil-capable element storage.
		Why: Generic `any -> String` conversion would turn a stored nil into the text
		`"null"`, and leaving class/interface elements erased would produce invalid Go
		selectors. Array reads need assertions, not general Dynamic conversion.
		How: Preserve nullable primitives as `any`, use the existing numeric/bool
		coercions when a concrete scalar is required, and nil-safely assert every other
		concrete Go type.
	**/
	function coerceStoredArrayElementExpr(expr:GoExpr, targetType:Type):GoExpr {
		if (isNullablePrimitiveType(targetType)) {
			return expr;
		}
		if (isIntType(targetType) || isHaxeInt32Type(targetType)) {
			return GoExpr.GoCall(GoExpr.GoIdent("hxrt.IntFromNullableAny"), [expr]);
		}
		if (isFloatType(targetType)) {
			return lowerNilSafeTypeAssertExpr(expr, "float64");
		}
		if (isBoolType(targetType)) {
			return lowerNilSafeTypeAssertExpr(expr, "bool");
		}
		if (typeToGoType(targetType) == "any") {
			return expr;
		}
		return lowerNullableAwareTypeAssertExpr(expr, targetType);
	}

	function wrapInt32Expr(expr:GoExpr):GoExpr {
		return GoExprOperatorOps.wrapInt32Expr(expr);
	}

	function floatOperandExpr(expr:GoExpr, operandType:Type, ?operand:TypedExpr):GoExpr {
		if (isNullableFloatType(operandType)) {
			return coerceNullableFloatOperandExpr(expr, operandType, operand);
		}
		return GoExprOperatorOps.floatOperandExpr(expr, isFloatType(operandType));
	}

	function unitStepExpr(target:GoExpr, opSymbol:GoBinaryOperator, valueType:Type, ?sourcePos:haxe.macro.Expr.Position):GoExpr {
		return GoExprOperatorOps.unitStepExpr(target, opSymbol, isInt32SemanticType(valueType, sourcePos));
	}

	function binopSymbol(op:Binop):GoBinaryOperator {
		return GoExprOperatorOps.binopSymbol(op);
	}

	function unopSymbol(op:Unop):GoUnaryOperator {
		return GoExprOperatorOps.unopSymbol(op);
	}

	function classFromType(type:Type):Null<ClassType> {
		var followed = Context.follow(type);
		return switch (followed) {
			case TInst(classRef, _):
				classRef.get();
			case _:
				null;
		};
	}

	/**
		What
		- Resolves the Go selector declared by an interface receiver's field metadata.

		Why
		- Haxe may retain the concrete implementation in `FInstance` while the
		  receiver expression is statically typed as an interface. Selecting the
		  concrete method name bypasses bridges such as `IMap.setIMap` and produces
		  invalid Go.

		How
		- Inspect the receiver's static class, require an interface, and look up the
		  matching interface field before applying `@:go.name`/`@:native` metadata.
	**/
	function interfaceSelectorForStaticReceiver(receiverType:Type, fieldName:String):Null<String> {
		var receiverClass = classFromType(receiverType);
		if (receiverClass == null || !receiverClass.isInterface) {
			return null;
		}
		for (field in receiverClass.fields.get()) {
			if (field.name == fieldName) {
				return interfaceFieldName(receiverClass, field);
			}
		}
		return null;
	}

	function inheritancePath(fromClass:ClassType, toClass:ClassType):Null<Array<ClassType>> {
		if (fullClassName(fromClass) == fullClassName(toClass)) {
			return [];
		}

		var path = new Array<ClassType>();
		var current = fromClass;
		while (true) {
			var parent = projectSuperClass(current);
			if (parent == null) {
				return null;
			}
			path.push(parent);
			if (fullClassName(parent) == fullClassName(toClass)) {
				return path;
			}
			current = parent;
		}
		return null;
	}

	/**
		What: Lowers a value for a known assignment/return/branch target type.
		Why: A direct array iterator adapter must replace the original inline
		expression and its prefixes; adapting only the final Go expression can leave
		behind erased `ArrayIterator` construction state or duplicate cursor locals.
		How: Ask the iterable owner for the pre-lowering array-cursor plan first,
		then fall back to ordinary expression lowering plus nominal or concrete-class
		structural upcasting.
	**/
	function lowerExprWithExpectedUpcast(source:TypedExpr, targetType:Type):LoweredExprWithPrefix {
		var expectedThrow = lowerExpectedThrowExpr(source, targetType);
		if (expectedThrow != null) {
			return expectedThrow;
		}
		var directNativeLiteral = lowerArrayLiteralForExpectedStorage(source, targetType);
		if (directNativeLiteral != null) {
			return directNativeLiteral;
		}
		var nativeArrayIterator = lambdaIterableLowering.nativeArrayStructuralIteratorCoerce(source, targetType);
		if (nativeArrayIterator != null) {
			return nativeArrayIterator;
		}
		var inlineConcreteIterator = lambdaIterableLowering.inlineConcreteStructuralIteratorCoerce(source, targetType);
		if (inlineConcreteIterator != null) {
			return inlineConcreteIterator;
		}
		var lowered = lowerExprWithPrefix(source);
		var adapted = upcastIfNeeded(lowered.expr, source.t, targetType, source);
		if (isSharedArrayElementExpr(source)) {
			adapted = coerceStoredArrayElementExpr(adapted, targetType);
		}
		return {
			prefix: lowered.prefix,
			expr: adapted,
			isStringLike: lowered.isStringLike
		};
	}

	/**
		What: Lowers a fresh Array literal directly into an expected native slice view.
		Why: Inline abstracts such as `Vector<T>` construct their raw storage through
		a temporary root-Array-typed node; materializing a shared carrier only to copy
		it immediately adds runtime footprint without any observable alias.
		How: Apply this only to a literal with a statically raw array-like destination.
		Every non-literal conversion retains the explicit copying bridge.
	**/
	function lowerArrayLiteralForExpectedStorage(source:TypedExpr, targetType:Type):Null<LoweredExprWithPrefix> {
		if (!isArrayType(targetType) || isHaxeArrayType(targetType)) {
			return null;
		}
		return switch (source.expr) {
			case TArrayDecl(values):
				{
					prefix: [],
					expr: GoExpr.GoCompositeLiteral(GoType.slice(arrayElementGoType(targetType)),
						[for (value in values) GoCompositeElement.GoCompositeValue(lowerExpr(value).expr)]),
					isStringLike: false
				};
			case TMeta(_, inner) | TParenthesis(inner):
				lowerArrayLiteralForExpectedStorage(inner, targetType);
			case _:
				null;
		};
	}

	/**
		What: Adapts one already-lowered value to its Haxe-proven target type.
		Why: Go does not consider a generated class pointer assignable to Haxe's
		anonymous iterator map, while ordinary class inheritance still needs its
		embedded-base selector path.
		How: Prefer source-aware shared-Array/native-slice and inline concrete-tail
		plans when typed
		source is available, then the general concrete iterator adapter, then the
		existing nominal upcast; leave unrelated values unchanged.
	**/
	function upcastIfNeeded(expr:GoExpr, fromType:Type, toType:Type, ?sourceTypedExpr:TypedExpr):GoExpr {
		var arrayRepresentation = adaptArrayRepresentation(expr, fromType, toType, sourceTypedExpr);
		if (arrayRepresentation != null) {
			return arrayRepresentation;
		}
		if (sourceTypedExpr != null) {
			var nativeArrayIterator = lambdaIterableLowering.nativeArrayStructuralIteratorCoerce(sourceTypedExpr, toType);
			if (nativeArrayIterator != null) {
				return materializeExprWithPrefix(nativeArrayIterator, toType).expr;
			}
			var inlineConcreteIterator = lambdaIterableLowering.inlineConcreteStructuralIteratorCoerce(sourceTypedExpr, toType);
			if (inlineConcreteIterator != null) {
				return materializeExprWithPrefix(inlineConcreteIterator, toType).expr;
			}
		}
		var structuralIterator = lambdaIterableLowering.structuralIteratorCoerce(expr, fromType, toType);
		if (structuralIterator != null) {
			return structuralIterator;
		}

		var fromClass = classFromType(fromType);
		var toClass = classFromType(toType);
		if (fromClass == null || toClass == null) {
			return expr;
		}
		var path = inheritancePath(fromClass, toClass);
		if (path == null || path.length == 0) {
			return expr;
		}

		var out = expr;
		for (classType in path) {
			out = GoExpr.GoSelector(out, classTypeName(classType));
		}
		return out;
	}

	/**
		What: Converts explicitly different portable/native array representations at
		a typed assignment, argument, return, or inline-abstract boundary.
		Why: Root Haxe Array and ReadOnlyArray now have shared identity, while
		fixed, Vector, Rest, and explicit native
		array-like APIs intentionally remain Go slices; allowing Go inference to pick
		the source representation produces invalid or semantically hidden conversions.
		How: Copy a shared carrier's erased values into the target typed slice, or copy
		a typed slice into a new shared carrier. Equal representations pass through.
	**/
	function adaptArrayRepresentation(expr:GoExpr, fromType:Type, toType:Type, ?sourceTypedExpr:TypedExpr):Null<GoExpr> {
		if (!isArrayType(fromType) || !isArrayType(toType)) {
			return null;
		}
		var sourceStorageType = sourceTypedExpr == null ? fromType : arrayStorageType(sourceTypedExpr);
		var fromShared = sourceTypedExpr != null && isRestPackExpr(sourceTypedExpr) ? false : isHaxeArrayType(sourceStorageType);
		var toShared = isHaxeArrayType(toType);
		if (fromShared == toShared) {
			return null;
		}
		if (fromShared) {
			var valuesMethod = arrayElementGoType(toType) == "any" ? "ValuesCopy" : "Values";
			return lambdaIterableLowering.anyArrayCoerce(GoExpr.GoCall(GoExpr.GoSelector(expr, valuesMethod), []), toType);
		}
		return GoExpr.GoCall(GoExpr.GoIdent("hxrt.ArrayFromValues"), [lowerTypedArrayToAnyCoerce(expr, sourceStorageType)]);
	}

	/** Recognizes the typer-generated raw slice block used to pack rest arguments. */
	function isRestPackExpr(expr:TypedExpr):Bool {
		return switch (expr.expr) {
			case TBlock(exprs):
				var found = false;
				for (entry in exprs) {
					switch (entry.expr) {
						case TBinop(OpAssign, _, right):
							switch (right.expr) {
								case TArrayDecl(_): found = true;
								case _:
							}
						case _:
					}
				}
				found;
			case TMeta(_, inner) | TParenthesis(inner) | TCast(inner, _):
				isRestPackExpr(inner);
			case _:
				false;
		};
	}

	function typeToGoType(type:Type):String {
		return GoTypeMapper.typeToGoType(type, classTypeNameForMappedType, enumTypeName);
	}

	function valueStorageGoType(type:Type):String {
		// Nullable primitive expression temps must be able to hold Go nil; keep the
		// broader type mapper unchanged so signatures and eligibility stay stable.
		return isNullablePrimitiveType(type) ? "any" : typeToGoType(type);
	}

	function isStringType(type:Type):Bool {
		return GoTypeMapper.isStringType(type);
	}

	function isGoErrorType(type:Type):Bool {
		var inner = nullableInnerType(type);
		if (inner != null) {
			return isGoErrorType(inner);
		}
		return switch (Context.follow(type)) {
			case TInst(classRef, _): var classType = classRef.get(); classType.pack.length == 1 && classType.pack[0] == "go" && classType.name == "Error";
			case _:
				false;
		};
	}

	function nullableInnerType(type:Type):Null<Type> {
		return switch (type) {
			case TAbstract(abstractRef, params): var abstractType = abstractRef.get(); abstractType.pack.length == 0 && abstractType.name == "Null" && params.length == 1 ? params[0] : null;
			case TMono(ref):
				var resolved = ref.get();
				resolved == null ? null : nullableInnerType(resolved);
			case TType(_, _):
				var followed = haxe.macro.TypeTools.follow(type, true);
				followed != type ? nullableInnerType(followed) : null;
			case TLazy(f):
				nullableInnerType(f());
			case _:
				null;
		};
	}

	function isInterfaceType(type:Type):Bool {
		return GoTypeMapper.isInterfaceType(type);
	}

	function isBoolType(type:Type):Bool {
		return GoTypeMapper.isBoolType(type);
	}

	function isIntType(type:Type):Bool {
		return GoTypeMapper.isIntType(type);
	}

	function isNullableIntType(type:Type):Bool {
		return GoTypeMapper.isNullableIntType(type);
	}

	function isNullableFloatType(type:Type):Bool {
		return GoTypeMapper.isNullableFloatType(type);
	}

	function isNullableBoolType(type:Type):Bool {
		return GoTypeMapper.isNullableBoolType(type);
	}

	function isNullablePrimitiveType(type:Type):Bool {
		return GoTypeMapper.isNullablePrimitiveType(type);
	}

	function isHaxeInt32Type(type:Type):Bool {
		return GoTypeMapper.isHaxeInt32Type(type);
	}

	function isInt32SemanticType(type:Type, ?sourcePos:haxe.macro.Expr.Position):Bool {
		return GoTypeMapper.isInt32SemanticType(type);
	}

	function isInt32StdlibPosition(pos:Null<haxe.macro.Expr.Position>):Bool {
		if (pos == null) {
			return false;
		}

		var file = Std.string(PositionTools.toLocation(pos).file);
		if (file == null || file == "") {
			return false;
		}

		var normalized = StringTools.replace(file, "\\", "/");
		return StringTools.contains(normalized, "/std/haxe/Int32.hx")
			|| StringTools.contains(normalized, "/std/haxe/Int64.hx")
			|| StringTools.contains(normalized, "/std/haxe/Int64Helper.hx");
	}

	function isFloatType(type:Type):Bool {
		return GoTypeMapper.isFloatType(type);
	}

	function isAnyLikeType(type:Type):Bool {
		return isDynamicCatchType(type) || typeToGoType(type) == "any";
	}

	function isStdClassMetaType(type:Type):Bool {
		return GoTypeMapper.isStdClassMetaType(type);
	}

	function isStdEnumMetaType(type:Type):Bool {
		return GoTypeMapper.isStdEnumMetaType(type);
	}

	function isDefinitelyNonNullableType(type:Type):Bool {
		return GoTypeMapper.isDefinitelyNonNullableType(type);
	}

	function isNullLiteralExpr(expr:TypedExpr):Bool {
		return switch (expr.expr) {
			case TConst(TNull):
				true;
			case TMeta(_, inner):
				isNullLiteralExpr(inner);
			case TParenthesis(inner):
				isNullLiteralExpr(inner);
			case TCast(inner, _):
				isNullLiteralExpr(inner);
			case _:
				false;
		};
	}

	function isOptionalPrimitiveLocalExpr(expr:TypedExpr):Bool {
		return switch (expr.expr) {
			case TLocal(variable): isRegisteredOptionalPrimitiveParam(variable) || valueStorageGoType(variable.t) == "any";
			case TMeta(_, inner):
				isOptionalPrimitiveLocalExpr(inner);
			case TParenthesis(inner):
				isOptionalPrimitiveLocalExpr(inner);
			case TCast(inner, _):
				isOptionalPrimitiveLocalExpr(inner);
			case _:
				false;
		};
	}

	function lowerOptionalPrimitiveNilComparisonExpr(expr:TypedExpr):Null<GoExpr> {
		return switch (expr.expr) {
			case TLocal(variable):
				(isRegisteredOptionalPrimitiveParam(variable)
					|| valueStorageGoType(variable.t) == "any") ? GoExpr.GoIdent(localVarName(variable)) : null;
			case TMeta(_, inner):
				lowerOptionalPrimitiveNilComparisonExpr(inner);
			case TParenthesis(inner):
				lowerOptionalPrimitiveNilComparisonExpr(inner);
			case TCast(inner, _):
				lowerOptionalPrimitiveNilComparisonExpr(inner);
			case _:
				null;
		};
	}

	function exprBackedByAny(expr:TypedExpr):Bool {
		return switch (expr.expr) {
			case TLocal(variable): registeredNarrowedPrimitiveStorageGoType(variable) == null && registeredNonNullPrimitiveLocalGoType(variable) == null && valueStorageGoType(variable.t) == "any";
			case TArray(target, _): usesSharedArrayCarrier(target);
			case TMeta(_, inner):
				exprBackedByAny(inner);
			case TParenthesis(inner):
				exprBackedByAny(inner);
			case TCast(inner, _):
				exprBackedByAny(inner);
			case _:
				false;
		};
	}

	/**
		What: Recognizes a value read from the shared Array's erased element storage.
		Why: Haxe inlines iterator and rest-argument reads inside expression blocks; if
		the terminal indexed read is hidden by that block, Go receives `any` where the
		static Haxe element type is required.
		How: Peel compile-time wrappers and expression blocks only through their final
		value, leaving unrelated statements and native-slice reads unchanged.
	**/
	function isSharedArrayElementExpr(expr:TypedExpr):Bool {
		return switch (expr.expr) {
			case TArray(target, _):
				usesSharedArrayCarrier(target);
			case TBlock(exprs) if (exprs.length > 0):
				isSharedArrayElementExpr(exprs[exprs.length - 1]);
			case TMeta(_, inner) | TParenthesis(inner) | TCast(inner, _):
				isSharedArrayElementExpr(inner);
			case _:
				false;
		};
	}

	/**
		What: Recovers the carrier's erased value for a shared-Array indexed read.
		Why: Haxe may wrap a non-string element in a String-typed cast while typing
		string concatenation. Lowering that wrapper first would assert that an Int or
		Bool already is `*string`, rather than letting `StringConcatAny` apply Haxe's
		ordinary string conversion.
		How: Peel compile-time wrappers, preserve expression-block statements in an
		`any`-returning closure, and lower the terminal carrier read without a type
		assertion. This helper is used only after shared-element recognition succeeds.
	**/
	function lowerSharedArrayElementStorageExpr(expr:TypedExpr):GoExpr {
		return switch (expr.expr) {
			case TArray(target, index) if (usesSharedArrayCarrier(target)):
				GoExpr.GoCall(GoExpr.GoSelector(lowerExpr(target).expr, "Get"), [lowerExpr(index).expr]);
			case TBlock(exprs) if (exprs.length > 0):
				var prefix = new Array<GoStmt>();
				for (index in 0...exprs.length - 1) {
					prefix = prefix.concat(withoutThrowFallback(function() return lowerToStatements(exprs[index])));
				}
				var value = lowerSharedArrayElementStorageExpr(exprs[exprs.length - 1]);
				prefix.length == 0 ? value : GoExpr.GoCall(GoExpr.GoFuncLiteral([], ["any"], prefix.concat([GoStmt.GoReturn(value)])), []);
			case TMeta(_, inner) | TParenthesis(inner) | TCast(inner, _):
				lowerSharedArrayElementStorageExpr(inner);
			case _:
				lowerExpr(expr).expr;
		};
	}

	function exprUsesNarrowedPrimitiveStorage(expr:TypedExpr):Bool {
		return switch (expr.expr) {
			case TLocal(variable):
				registeredNarrowedPrimitiveStorageGoType(variable) != null;
			case TMeta(_, inner) | TParenthesis(inner) | TCast(inner, _):
				exprUsesNarrowedPrimitiveStorage(inner);
			case _:
				false;
		};
	}

	function isAnonymousObjectType(type:Type):Bool {
		return GoTypeMapper.isAnonymousObjectType(type);
	}

	function isArrayType(type:Type):Bool {
		return GoTypeMapper.isArrayType(type);
	}

	function isHaxeArrayType(type:Type):Bool {
		return GoTypeMapper.isHaxeArrayType(type);
	}

	function arrayFieldDeclaredType(access:FieldAccess):Null<Type> {
		return switch (access) {
			case FInstance(_, _, field):
				field.get().type;
			case FStatic(_, field):
				field.get().type;
			case FAnon(field):
				field.get().type;
			case FClosure(_, field):
				field.get().type;
			case _:
				null;
		};
	}

	/** What: Resolves the type of the carrier actually held by an array expression.
		Why: Vector and other inline views expose Array-typed nodes while retaining a
		raw slice, so the outer expression type alone can request the wrong operations.
		How: Peel only array-to-array compile-time wrappers and trust the innermost
		storage-bearing expression. **/
	function arrayStorageType(expr:TypedExpr):Type {
		return switch (expr.expr) {
			case TBlock(exprs) if (exprs.length > 0):
				arrayStorageType(exprs[exprs.length - 1]);
			case TMeta(_, inner) | TParenthesis(inner):
				arrayStorageType(inner);
			case TCast(inner, _) if (isArrayType(inner.t)):
				arrayStorageType(inner);
			case TLocal(variable):
				localArrayStorageOverrides.exists(variable.id) ? localArrayStorageOverrides.get(variable.id) : variable.t;
			case TField(_, access): var declared = arrayFieldDeclaredType(access); declared != null && isArrayType(declared) ? declared : expr.t;
			case _:
				expr.t;
		};
	}

	/**
		What: Finds the root-Array temporary used by an inlined `Vector<T>` constructor.
		Why: Haxe exposes the temporary as `Array<T>` even though the enclosing abstract
		contract retains fixed native-slice storage on this target.
		How: Require a Vector result, a locally declared Array temporary, and that same
		local as the block's terminal value before applying the representation override.
	**/
	function vectorBlockStorageOverride(blockType:Type, exprs:Array<TypedExpr>):Null<{variable:TVar, storageType:Type}> {
		if (vectorElementType(blockType) == null || exprs.length == 0) {
			return null;
		}
		var terminal = terminalLocalVariable(exprs[exprs.length - 1]);
		if (terminal == null || !isHaxeArrayType(terminal.t)) {
			return null;
		}
		for (entry in exprs) {
			switch (entry.expr) {
				case TVar(variable, _) if (variable.id == terminal.id):
					return {variable: variable, storageType: blockType};
				case _:
			}
		}
		return null;
	}

	function terminalLocalVariable(expr:TypedExpr):Null<TVar> {
		return switch (expr.expr) {
			case TLocal(variable): variable;
			case TMeta(_, inner) | TParenthesis(inner) | TCast(inner, _): terminalLocalVariable(inner);
			case _: null;
		};
	}

	/** True when an array expression's actual storage is the shared carrier. **/
	function usesSharedArrayCarrier(expr:TypedExpr):Bool {
		if (isBytesBufferStorageArray(expr)) {
			return false;
		}
		return isHaxeArrayType(arrayStorageType(expr));
	}

	function arrayElementGoType(type:Type):String {
		return GoTypeMapper.arrayElementGoType(type, classTypeNameForMappedType, enumTypeName);
	}

	function scalarGoType(type:Type):String {
		return GoTypeMapper.scalarGoType(type, classTypeNameForMappedType, enumTypeName);
	}

	function goFunctionType(args:Array<{name:String, opt:Bool, t:Type}>, returnType:Type):String {
		return GoTypeMapper.goFunctionType(args, returnType, classTypeNameForMappedType, enumTypeName);
	}

	/**
		What: record a Haxe class whose nominal type is rendered into generated Go.

		Why: static-only classes have no instance layout, but Haxe still permits them
		in type positions. Source-owned std classes may also have been removed by DCE,
		so observing the type must enqueue their staged declaration before emission.

		How: mark the nominal carrier, route known staged std dependencies through the
		existing source-owned planner, then return the ordinary normalized Go name.
	**/
	function classTypeNameForMappedType(classType:ClassType):String {
		requiredNominalClassTypeNames.set(fullClassName(classType), true);
		sourceOwnedStdlibPlanner.requireTypedSourceOwnedStdlibClass(classType);
		noteSourceOwnedStdlibUsage(classType);
		return classTypeName(classType);
	}

	function restElementType(type:Type):Null<Type> {
		return GoTypeMapper.restElementType(type);
	}

	function nativeSliceElementType(type:Type):Null<Type> {
		return GoTypeMapper.nativeSliceElementType(type);
	}

	function vectorElementType(type:Type):Null<Type> {
		return GoTypeMapper.vectorElementType(type);
	}

	function readOnlyArrayElementType(type:Type):Null<Type> {
		return GoTypeMapper.readOnlyArrayElementType(type);
	}

	function restIteratorCtorArg(expr:Null<TypedExpr>):Null<TypedExpr> {
		if (expr == null) {
			return null;
		}

		return switch (expr.expr) {
			case TNew(classRef, _, args):
				var classType = classRef.get();
				if (classType.pack.join(".") == "haxe.iterators" && classType.name == "RestIterator" && args.length == 1) {
					args[0];
				} else {
					null;
				}
			case TMeta(_, inner):
				restIteratorCtorArg(inner);
			case TParenthesis(inner):
				restIteratorCtorArg(inner);
			case TCast(inner, _):
				restIteratorCtorArg(inner);
			case _:
				null;
		};
	}

	function lowerRestIteratorCtor(variableName:String, argsExpr:TypedExpr):Array<GoStmt> {
		return [
			GoStmt.GoVarDecl(variableName + "_args", typeToGoType(argsExpr.t), lowerExpr(argsExpr).expr, true),
			GoStmt.GoVarDecl(variableName + "_current", "int", GoExpr.GoIntLiteral(0), true)
		];
	}

	function lowerRestPackBlock(exprs:Array<TypedExpr>):Null<GoExpr> {
		for (expr in exprs) {
			switch (expr.expr) {
				case TBinop(OpAssign, _, right):
					switch (right.expr) {
						case TArrayDecl(values):
							return GoExpr.GoCompositeLiteral(GoType.slice(arrayElementGoType(right.t)),
								[for (value in values) GoCompositeElement.GoCompositeValue(lowerExpr(value).expr)]);
						case _:
					}
				case _:
			}
		}
		return null;
	}

	function isVoidType(type:Type):Bool {
		return GoTypeMapper.isVoidType(type);
	}

	function isDynamicCatchType(type:Type):Bool {
		return GoTypeMapper.isDynamicCatchType(type);
	}

	function isTypeParameterClass(classType:ClassType):Bool {
		return GoTypeMapper.isTypeParameterClass(classType);
	}

	function isHaxeExceptionClass(classType:ClassType):Bool {
		return GoTypeMapper.isHaxeExceptionClass(classType);
	}

	function isHaxeExceptionFamilyClass(classType:ClassType):Bool {
		if (isHaxeExceptionClass(classType) || isHaxeValueExceptionClass(classType)) {
			return true;
		}
		var cursor = classType.superClass;
		while (cursor != null) {
			var superType = cursor.t.get();
			if (isHaxeExceptionClass(superType) || isHaxeValueExceptionClass(superType)) {
				return true;
			}
			cursor = superType.superClass;
		}
		return false;
	}

	function isHaxeValueExceptionClass(classType:ClassType):Bool {
		return classType.pack.join(".") == "haxe" && classType.name == "ValueException";
	}

	function isHaxeExceptionType(type:Type):Bool {
		return GoTypeMapper.isHaxeExceptionType(type);
	}

	function isHaxeExceptionFamilyType(type:Type):Bool {
		return switch (Context.follow(type)) {
			case TInst(classRef, _):
				isHaxeExceptionFamilyClass(classRef.get());
			case _:
				false;
		};
	}

	function isHaxeValueExceptionType(type:Type):Bool {
		return switch (Context.follow(type)) {
			case TInst(classRef, _):
				isHaxeValueExceptionClass(classRef.get());
			case _:
				false;
		};
	}

	function directHaxeExceptionSuperClass(classType:ClassType):Bool {
		if (isHaxeValueExceptionClass(classType)) {
			return false;
		}
		return switch (classType.superClass) {
			case null:
				false;
			case superRef: var superType = superRef.t.get(); isHaxeExceptionClass(superType) || isHaxeValueExceptionClass(superType);
		};
	}

	function isNilExpr(expr:GoExpr):Bool {
		return switch (expr) {
			case GoNil: true;
			case _: false;
		};
	}

	function requireStdlibShimGroup(group:String):Void {
		requiredStdlibShimGroups.set(group, true);
	}

	function noteStaticStdlibFieldUsage(classType:ClassType, fieldName:String, pos:Position):Void {
		if (classType.pack.length == 0 && classType.name == "Sys" && fieldName == "cpuTime") {
			Context.error("Sys.cpuTime is unsupported on haxe.go: Go's standard library does not expose portable process CPU time. "
				+ "Use an explicit Go-native module/API boundary for platform-specific process accounting.",
				pos);
		}
		sourceOwnedStdlibPlanner.noteSourceOwnedStdlibUsage(classType);
	}

	function noteSourceOwnedStdlibUsage(classType:ClassType):Void {
		sourceOwnedStdlibPlanner.noteSourceOwnedStdlibUsage(classType);
	}

	function requireSourceOwnedStdlibClass(className:String):Void {
		sourceOwnedStdlibPlanner.requireSourceOwnedStdlibClass(className);
	}

	/**
		What: enqueue every class/enum emitted from a source-owned stdlib module.

		Why: some upstream std modules, such as `haxe.Template`, rely on private
		module-local enums or helper types that `Context.getType()` cannot resolve by
		name. Requiring only the public class leaves those companions out of the Go
		output and breaks direct module usage.

		How: resolve the module once through `Context.getModule(...)`, cache any class
		or enum declarations it contains, and enqueue them through the existing
		source-owned class/enum pending maps.
	**/
	function requireSourceOwnedStdlibModule(moduleName:String):Void {
		sourceOwnedStdlibPlanner.requireSourceOwnedStdlibModule(moduleName);
	}

	function requireSourceOwnedStdlibEnum(enumName:String):Void {
		sourceOwnedStdlibPlanner.requireSourceOwnedStdlibEnum(enumName);
	}

	function hasLoadedSourceOwnedStdlibClass(className:String):Bool {
		return sourceOwnedStdlibPlanner.hasLoadedSourceOwnedStdlibClass(className);
	}

	function noteStdlibClass(classType:ClassType):Void {
		for (group in GoStdlibShimClassifier.requiredGroupsForClass(classType)) {
			requireStdlibShimGroup(group);
		}
	}

	function noteStdlibEnum(enumType:EnumType):Void {
		switch (fullEnumName(enumType)) {
			case "haxe.ds.Either":
				requireSourceOwnedStdlibEnum("haxe.ds.Either");
			case "sys.thread.NextEventTime":
				requireSourceOwnedStdlibEnum("sys.thread.NextEventTime");
			case _:
		}
		for (group in GoStdlibShimClassifier.requiredGroupsForEnum(enumType)) {
			requireStdlibShimGroup(group);
		}
	}

	function sortedRequiredStdlibShimGroups():Array<String> {
		var groups = [for (group in requiredStdlibShimGroups.keys()) group];
		groups.sort(Reflect.compare);
		return groups;
	}

	function inferRuntimeFeatures(requiredShimGroups:Array<String>):GoHxrtFeatureInference {
		var classPathSet = new Map<String, Bool>();
		for (classType in projectClasses)
			classPathSet.set(fullClassName(classType), true);
		for (classPath in usedExternClassPaths.keys())
			classPathSet.set(classPath, true);
		var classPaths = [for (classPath in classPathSet.keys()) classPath];
		classPaths.sort(Reflect.compare);
		var enumPaths = [for (enumType in projectEnums) fullEnumName(enumType)];
		enumPaths.sort(Reflect.compare);
		return GoHxrtFeatureAnalyzer.inferWithReasons(classPaths, enumPaths, requiredShimGroups, requiresEqualitySurface);
	}

	function resetExternImportPaths():Void {
		for (path in externImportPaths.keys()) {
			externImportPaths.remove(path);
		}
		for (path in externImportPackages.keys()) {
			externImportPackages.remove(path);
		}
		for (classPath in usedExternClassPaths.keys()) {
			usedExternClassPaths.remove(classPath);
		}
	}

	function resetRequiredNativeChanElementTypes():Void {
		for (elementType in requiredNativeChanElementTypes.keys()) {
			requiredNativeChanElementTypes.remove(elementType);
		}
	}

	function resetRequiredNativeSliceElementTypes():Void {
		for (elementType in requiredNativeSliceElementTypes.keys()) {
			requiredNativeSliceElementTypes.remove(elementType);
		}
	}

	function resetRequiredNativeMapTypePairs():Void {
		for (signature in requiredNativeMapTypePairs.keys()) {
			requiredNativeMapTypePairs.remove(signature);
		}
	}

	function resetRequiredNativeResultElementTypes():Void {
		for (elementType in requiredNativeResultElementTypes.keys()) {
			requiredNativeResultElementTypes.remove(elementType);
		}
	}

	/**
		What: Record one typed extern import and the class whose lowered use required it.

		Why: staged inline methods can disappear from emitted class output while their
		typed runtime extern call remains. Selective runtime inference must still see
		that surviving authority or it can omit the runtime file that defines the call.

		How: retain the existing import-path/package mapping and a deduplicated Haxe
		class path consumed by `GoHxrtFeatureAnalyzer` after lowering completes.
	**/
	function noteExternImportPath(classType:ClassType, packageName:String):Void {
		var path = externClassImportPath(classType);
		if (path == null || path == "") {
			return;
		}
		externImportPaths.set(path, true);
		usedExternClassPaths.set(fullClassName(classType), true);
		if (packageName != null && packageName != "") {
			externImportPackages.set(path, packageName);
		}
	}

	function normalizeMetaName(name:String):String {
		return StringTools.startsWith(name, ":") ? name.substr(1) : name;
	}

	function metaNameEquals(actual:String, expected:String):Bool {
		return normalizeMetaName(actual) == normalizeMetaName(expected);
	}

	function unwrapMetaExpr(expr:Expr):Expr {
		return switch (expr.expr) {
			case EParenthesis(inner):
				unwrapMetaExpr(inner);
			case EMeta(_, inner):
				unwrapMetaExpr(inner);
			case _:
				expr;
		};
	}

	function readConstStringExpr(expr:Expr):Null<String> {
		return switch (unwrapMetaExpr(expr).expr) {
			case EConst(CString(value, _)):
				value;
			case _:
				null;
		};
	}

	function readMetadataString(meta:MetaAccess, names:Array<String>):Null<String> {
		if (meta == null) {
			return null;
		}
		for (entry in meta.get()) {
			var matches = false;
			for (name in names) {
				if (metaNameEquals(entry.name, name)) {
					matches = true;
					break;
				}
			}
			if (!matches) {
				continue;
			}
			if (entry.params == null || entry.params.length == 0) {
				Context.fatalError("@" + normalizeMetaName(entry.name) + " requires a string parameter", entry.pos);
				return null;
			}
			var value = readConstStringExpr(entry.params[0]);
			if (value == null) {
				Context.fatalError("@" + normalizeMetaName(entry.name) + " requires a compile-time string parameter", entry.pos);
				return null;
			}
			return StringTools.trim(value);
		}
		return null;
	}

	function hasMetadata(meta:MetaAccess, names:Array<String>):Bool {
		if (meta == null) {
			return false;
		}
		for (entry in meta.get()) {
			for (name in names) {
				if (metaNameEquals(entry.name, name)) {
					return true;
				}
			}
		}
		return false;
	}

	function externClassImportPath(classType:ClassType):Null<String> {
		var value = readMetadataString(classType.meta, ["go.import"]);
		if (value == null || value == "") {
			return null;
		}
		// Allow staged stdlib externs to bind to the active runtime import path without
		// hardcoding a module prefix like "snapshot/hxrt".
		if (value == "hxrt") {
			return compilationContext.runtimeImportPath;
		}
		return value;
	}

	function externClassPackageName(classType:ClassType):Null<String> {
		var importPath = externClassImportPath(classType);
		if (importPath == null) {
			return null;
		}

		var packageName = readMetadataString(classType.meta, ["go.package", "go.pkg"]);
		if (packageName == null || packageName == "") {
			var segments = [for (segment in importPath.split("/")) StringTools.trim(segment)];
			var index = segments.length - 1;
			while (index >= 0 && segments[index] == "") {
				index--;
			}
			packageName = index >= 0 ? segments[index] : "";
		}
		if (packageName == null || packageName == "") {
			Context.fatalError('Unable to infer Go package identifier from @:go.import("' + importPath + '")', classType.pos);
			return null;
		}

		return normalizeIdent(packageName);
	}

	function externClassTypeName(classType:ClassType):String {
		var typeName = readMetadataString(classType.meta, ["go.name", "native"]);
		return typeName == null || typeName == "" ? classType.name : typeName;
	}

	function externFieldName(field:ClassField):String {
		var mapped = readMetadataString(field.meta, ["go.name", "native"]);
		return mapped == null || mapped == "" ? field.name : mapped;
	}

	function interfaceFieldName(classType:ClassType, field:ClassField):String {
		var mapped = readMetadataString(field.meta, ["go.name", "native"]);
		if (mapped != null && mapped != "") {
			return normalizeIdent(mapped);
		}
		return normalizeIdent(field.name);
	}

	/**
		What: Compares two elements for `Array.remove` using Haxe `==` semantics.
		Why: Go pointer equality is wrong for Haxe strings, while erased, nullable,
		and non-comparable carriers cannot safely use Go's `==` operator.
		How: Keep typed comparable values native, compare strings by contents, and
		delegate only interface-backed or non-comparable shapes to the narrow runtime
		helper.
	**/
	function lowerArrayElementEqualityExpr(left:GoExpr, right:GoExpr, elementType:Type, elementGoType:String):GoExpr {
		if (isStringType(elementType)) {
			return GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringEqualStringPtr"), [left, right]);
		}
		if (elementGoType == "any"
			|| StringTools.startsWith(elementGoType, "[]")
			|| StringTools.startsWith(elementGoType, "map[")
			|| StringTools.startsWith(elementGoType, "func(")) {
			requiresEqualitySurface = true;
			return GoExpr.GoCall(GoExpr.GoIdent("hxrt.HaxeEqual"), [left, right]);
		}
		return GoExpr.GoBinary("==", left, right);
	}

	/**
		What: Lowers `Array.remove(value)` as a first-match shared-carrier mutation.
		Why: The portable carrier deliberately exposes only representation primitives;
		Haxe equality and first-match policy still need typed compiler lowering.
		How: Evaluate the receiver and value once, range to the first Haxe-equal
		element, then ask the carrier to remove that known slot in place.
	**/
	function lowerArrayRemoveExpr(target:TypedExpr, args:Array<TypedExpr>):LoweredExpr {
		if (args.length != 1) {
			Context.fatalError("Array.remove expects exactly one value", target.pos);
		}
		var site = lowerArrayMutationSite(target);
		var loweredValue = lowerExprWithPrefix(args[0]);
		var valueName = freshTempName("hx_remove_value");
		var indexName = freshTempName("hx_remove_index");
		var elementName = freshTempName("hx_remove_element");
		requiresEqualitySurface = true;
		var indexExpr = GoExpr.GoIdent(indexName);
		var foundBody = [
			GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoSelector(site.tempExpr, "RemoveAt"), [indexExpr])),
			GoStmt.GoReturn(GoExpr.GoBoolLiteral(true))
		];
		var body = site.prefix.concat(loweredValue.prefix).concat([
			GoStmt.GoVarDecl(valueName, "any", loweredValue.expr, false),
			GoStmt.GoVarDecl(indexName, "int", GoExpr.GoIntLiteral(0), true),
			GoStmt.GoWhile(GoExpr.GoBinary("<", indexExpr, GoExpr.GoCall(GoExpr.GoSelector(site.tempExpr, "Len"), [])), [
				GoStmt.GoVarDecl(elementName, "any", GoExpr.GoCall(GoExpr.GoSelector(site.tempExpr, "Get"), [indexExpr]), true),
				GoStmt.GoIf(GoExpr.GoCall(GoExpr.GoIdent("hxrt.HaxeEqual"), [GoExpr.GoIdent(elementName), GoExpr.GoIdent(valueName)]), foundBody, null),
				GoStmt.GoAssign(indexExpr, GoExpr.GoBinary("+", indexExpr, GoExpr.GoIntLiteral(1)))
			]),
			GoStmt.GoReturn(GoExpr.GoBoolLiteral(false))
		]);
		return {
			expr: GoExpr.GoCall(GoExpr.GoFuncLiteral([], ["bool"], body), []),
			isStringLike: false
		};
	}

	/**
		What: Lowers `Array.insert(position, value)` through the shared carrier.
		Why: Portable Haxe clamps oversized positions and resolves negative positions
		from the end, behavior owned by the carrier rather than a copied slice header.
		How: Evaluate receiver and arguments once in order, then pass the typed position
		and erased value to the carrier's in-place insertion operation.
	**/
	function lowerArrayInsertExpr(target:TypedExpr, args:Array<TypedExpr>):LoweredExpr {
		if (args.length != 2) {
			Context.fatalError("Array.insert expects a position and value", target.pos);
		}
		var site = lowerArrayMutationSite(target);
		var loweredPosition = lowerExprWithPrefix(args[0]);
		var loweredValue = lowerExprWithPrefix(args[1]);
		var positionName = freshTempName("hx_insert_position");
		var valueName = freshTempName("hx_insert_value");
		var body = site.prefix.concat(loweredPosition.prefix)
			.concat([GoStmt.GoVarDecl(positionName, "int", loweredPosition.expr, true)])
			.concat(loweredValue.prefix)
			.concat([
				GoStmt.GoVarDecl(valueName, "any", loweredValue.expr, false),
				GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoSelector(site.tempExpr, "Insert"), [GoExpr.GoIdent(positionName), GoExpr.GoIdent(valueName)]))
			]);
		return {
			expr: GoExpr.GoCall(GoExpr.GoFuncLiteral([], [], body), []),
			isStringLike: false
		};
	}

	function lowerArrayInstanceCall(callee:TypedExpr, args:Array<TypedExpr>, returnType:Type):Null<LoweredExpr> {
		var methodCall = asArrayMethodCall(callee);
		if (methodCall == null || !isArrayType(methodCall.target.t)) {
			return null;
		}

		return switch (methodCall.methodName) {
			case "remove":
				lowerArrayRemoveExpr(methodCall.target, args);
			case "insert":
				lowerArrayInsertExpr(methodCall.target, args);
			case "copy" if (args.length == 0):
				var shared = usesSharedArrayCarrier(methodCall.target);
				var copied = shared ? GoExpr.GoCall(GoExpr.GoSelector(lowerExpr(methodCall.target).expr, "Copy"),
					[]) : cloneArrayExpr(lowerExpr(methodCall.target).expr, methodCall.target.t);
				if (!shared && isHaxeArrayType(returnType)) {
					copied = GoExpr.GoCall(GoExpr.GoIdent("hxrt.ArrayFromValues"),
						[lowerTypedSliceToAnyByGoType(copied, arrayElementGoType(methodCall.target.t))]);
				}
				{
					expr: copied,
					isStringLike: false
				};
			case "join":
				var delimiterExpr = args.length > 0 ? lowerExpr(args[0]).expr : GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"),
					[GoExpr.GoStringLiteral("")]);
				var targetExpr = lowerExpr(methodCall.target).expr;
				var sourceExpr = usesSharedArrayCarrier(methodCall.target) ? GoExpr.GoCall(GoExpr.GoSelector(targetExpr, "Values"),
					[]) : lowerTypedArrayToAnyCoerce(targetExpr, methodCall.target.t);
				{
					expr: GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringJoinAny"), [sourceExpr, delimiterExpr]),
					isStringLike: true
				};
			case "push":
				var site = lowerArrayMutationSite(methodCall.target);
				var shared = usesSharedArrayCarrier(methodCall.target);
				var pushArgs = shared ? [] : [site.tempExpr];
				var shouldMaskToByte = isBytesBufferStorageArray(methodCall.target);
				var prefix = site.prefix.copy();
				for (arg in args) {
					var loweredArg = lowerExprWithPrefix(arg);
					prefix = prefix.concat(loweredArg.prefix);
					var appendValue = loweredArg.expr;
					if (shouldMaskToByte) {
						appendValue = GoExpr.GoBinary("&", appendValue, GoExpr.GoIntLiteral(255));
					}
					pushArgs.push(appendValue);
				}
				var body = if (shared) {
					prefix.concat([
						GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoSelector(site.tempExpr, "Push"), pushArgs))
					]);
				} else {
					prefix.concat([
						GoStmt.GoAssign(site.tempExpr, GoExpr.GoCall(GoExpr.GoIdent("append"), pushArgs))
					]).concat(site.writeBack(site.tempExpr)).concat([GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("len"), [site.tempExpr]))]);
				};
				{
					expr: GoExpr.GoCall(GoExpr.GoFuncLiteral([], ["int"], body), []),
					isStringLike: false
				};
			case "pop" if (args.length == 0):
				var site = lowerArrayMutationSite(methodCall.target);
				var shared = usesSharedArrayCarrier(methodCall.target);
				var resultType = shared ? valueStorageGoType(returnType) : typeToGoType(returnType);
				var body = if (shared) {
					var popped = coerceStoredArrayElementExpr(GoExpr.GoCall(GoExpr.GoSelector(site.tempExpr, "Pop"), []), returnType);
					site.prefix.concat([GoStmt.GoReturn(popped)]);
				} else {
					var lenName = freshTempName("hx_len");
					var valueName = freshTempName("hx_value");
					var zeroName = freshTempName("hx_zero");
					site.prefix.concat([
						GoStmt.GoVarDecl(lenName, "int", GoExpr.GoCall(GoExpr.GoIdent("len"), [site.tempExpr]), true),
						GoStmt.GoIf(GoExpr.GoBinary("==", GoExpr.GoIdent(lenName), GoExpr.GoIntLiteral(0)), [
							GoStmt.GoVarDecl(zeroName, resultType, null, false),
							GoStmt.GoReturn(GoExpr.GoIdent(zeroName))
						],
							null),
						GoStmt.GoVarDecl(valueName, resultType,
							GoExpr.GoIndex(site.tempExpr, GoExpr.GoBinary("-", GoExpr.GoIdent(lenName), GoExpr.GoIntLiteral(1))), true),
						GoStmt.GoAssign(site.tempExpr,
							GoExpr.GoSlice(site.tempExpr, null, GoExpr.GoBinary("-", GoExpr.GoIdent(lenName), GoExpr.GoIntLiteral(1))))
					]).concat(site.writeBack(site.tempExpr)).concat([GoStmt.GoReturn(GoExpr.GoIdent(valueName))]);
				};
				{
					expr: GoExpr.GoCall(GoExpr.GoFuncLiteral([], [resultType], body), []),
					isStringLike: isStringType(returnType)
				};
			case _:
				null;
		};
	}

	function lowerRestAbstractCall(callee:TypedExpr, args:Array<TypedExpr>, returnType:Type):Null<LoweredExpr> {
		var isAppend = isStaticCall(callee, "Rest_Impl_", ["haxe", "_Rest"], "append");
		var isPrepend = isStaticCall(callee, "Rest_Impl_", ["haxe", "_Rest"], "prepend");
		if (!isAppend && !isPrepend) {
			return null;
		}

		if (args.length < 2) {
			Context.fatalError("haxe.Rest helper lowering requires source and value arguments", callee.pos);
			return null;
		}

		var sourceExpr = lowerExpr(args[0]).expr;
		var valueExpr = lowerExpr(args[1]).expr;
		var elementType = restElementType(args[0].t);
		var elementGoType = elementType == null ? arrayElementGoType(returnType) : scalarGoType(elementType);
		var sliceType = "[]" + elementGoType;

		return {
			expr: isAppend ? appendClonedArrayExpr(sourceExpr, valueExpr, sliceType,
				elementGoType) : prependClonedArrayExpr(sourceExpr, valueExpr, sliceType, elementGoType),
			isStringLike: false
		};
	}

	function hasExternReceiverMeta(field:ClassField):Bool {
		return hasMetadata(field.meta, ["go.receiver"]);
	}

	function hasExternValueErrorMeta(field:ClassField):Bool {
		return hasMetadata(field.meta, ["go.valueError", "go.value_error"]);
	}

	function hasExternTupleReturnMeta(field:ClassField):Bool {
		return hasMetadata(field.meta, ["go.tupleReturn", "go.tuple_return"]);
	}

	function classTypeName(classType:ClassType):String {
		if (classType.isExtern) {
			var packageName = externClassPackageName(classType);
			if (packageName != null) {
				noteExternImportPath(classType, packageName);
				return packageName + "." + externClassTypeName(classType);
			}
		}
		noteStdlibClass(classType);
		return GoNaming.typeSymbol(classType.pack, classType.name);
	}

	function enumTypeName(enumType:EnumType):String {
		noteStdlibEnum(enumType);
		return GoNaming.typeSymbol(enumType.pack, enumType.name);
	}

	function constructorSymbol(classType:ClassType):String {
		noteStdlibClass(classType);
		return GoNaming.constructorSymbol(classType.pack, classType.name);
	}

	function enumConstructorSymbol(enumType:EnumType, fieldName:String):String {
		return enumTypeName(enumType) + "_" + normalizeIdent(fieldName);
	}

	function staticSymbol(classType:ClassType, fieldName:String):String {
		noteStdlibClass(classType);
		return GoNaming.staticSymbol(classType.pack, classType.name, fieldName, isSelectedMainClass(classType));
	}

	function normalizeIdent(name:String):String {
		return GoNaming.normalizeIdent(name);
	}

	function cloneArrayExpr(sourceExpr:GoExpr, sourceType:Type):GoExpr {
		var elementGoType = arrayElementGoType(sourceType);
		var sliceType = "[]" + elementGoType;
		return GoExpr.GoCall(GoExpr.GoFuncLiteral([{name: "src", typeName: sliceType}], [sliceType], [
			GoStmt.GoRaw("out := append(" + sliceType + "{}, src...)"),
			GoStmt.GoReturn(GoExpr.GoIdent("out"))
		]), [sourceExpr]);
	}

	function appendClonedArrayExpr(sourceExpr:GoExpr, valueExpr:GoExpr, sliceType:String, elementGoType:String):GoExpr {
		return GoExpr.GoCall(GoExpr.GoFuncLiteral([{name: "src", typeName: sliceType}, {name: "value", typeName: elementGoType}], [sliceType], [
			GoStmt.GoRaw("out := append(" + sliceType + "{}, src...)"),
			GoStmt.GoRaw("out = append(out, value)"),
			GoStmt.GoReturn(GoExpr.GoIdent("out"))
		]), [sourceExpr, valueExpr]);
	}

	function prependClonedArrayExpr(sourceExpr:GoExpr, valueExpr:GoExpr, sliceType:String, elementGoType:String):GoExpr {
		return GoExpr.GoCall(GoExpr.GoFuncLiteral([{name: "src", typeName: sliceType}, {name: "value", typeName: elementGoType}], [sliceType],
			[GoStmt.GoRaw("return append(" + sliceType + "{value}, src...)")]),
			[sourceExpr, valueExpr]);
	}

	function freshTempName(prefix:String):String {
		tempVarCounter++;
		return prefix + "_" + tempVarCounter;
	}

	function asArrayMethodCall(callee:TypedExpr):Null<ArrayMethodCall> {
		return switch (callee.expr) {
			case TField(target, FInstance(classRef, _, field)):
				var classType = classRef.get();
				if (classType.pack.length == 0 && classType.name == "Array") {
					{target: target, methodName: field.get().name};
				} else {
					null;
				}
			case TField(target, FAnon(field)):
				{target: target, methodName: field.get().name};
			case TField(target, FDynamic(name)):
				{target: target, methodName: name};
			case _:
				null;
		};
	}

	function asGoChanMethodCall(callee:TypedExpr):Null<GoChanMethodCall> {
		return switch (callee.expr) {
			case TField(target, FInstance(classRef, _, field)):
				var classType = classRef.get();
				var elementType = goChanElementType(target.t);
				if (isGoChanClass(classType) && elementType != null) {
					{
						target: target,
						methodName: field.get().name,
						elementType: elementType
					};
				} else {
					null;
				}
			case TMeta(_, inner):
				asGoChanMethodCall(inner);
			case TParenthesis(inner):
				asGoChanMethodCall(inner);
			case TCast(inner, _):
				asGoChanMethodCall(inner);
			case _:
				null;
		};
	}

	function asGoSliceMethodCall(callee:TypedExpr):Null<GoSliceMethodCall> {
		return switch (callee.expr) {
			case TField(target, FInstance(classRef, _, field)):
				var classType = classRef.get();
				var elementType = goSliceElementType(target.t);
				if (isGoSliceClass(classType) && elementType != null) {
					{
						target: target,
						methodName: field.get().name,
						elementType: elementType
					};
				} else {
					null;
				}
			case TMeta(_, inner):
				asGoSliceMethodCall(inner);
			case TParenthesis(inner):
				asGoSliceMethodCall(inner);
			case TCast(inner, _):
				asGoSliceMethodCall(inner);
			case _:
				null;
		};
	}

	function asGoMapMethodCall(callee:TypedExpr):Null<GoMapMethodCall> {
		return switch (callee.expr) {
			case TField(target, FInstance(classRef, _, field)):
				var classType = classRef.get();
				var pair = goMapTypePair(target.t);
				if (isGoMapClass(classType) && pair != null) {
					{
						target: target,
						methodName: field.get().name,
						keyType: pair.keyType,
						valueType: pair.valueType
					};
				} else {
					null;
				}
			case TMeta(_, inner):
				asGoMapMethodCall(inner);
			case TParenthesis(inner):
				asGoMapMethodCall(inner);
			case TCast(inner, _):
				asGoMapMethodCall(inner);
			case _:
				null;
		};
	}

	function asGoResultMethodCall(callee:TypedExpr):Null<GoResultMethodCall> {
		return switch (callee.expr) {
			case TField(target, FInstance(classRef, _, field)):
				var classType = classRef.get();
				var elementType = goResultElementType(target.t);
				if (isGoResultClass(classType) && elementType != null) {
					{
						target: target,
						methodName: field.get().name,
						elementType: elementType
					};
				} else {
					null;
				}
			case TMeta(_, inner):
				asGoResultMethodCall(inner);
			case TParenthesis(inner):
				asGoResultMethodCall(inner);
			case TCast(inner, _):
				asGoResultMethodCall(inner);
			case _:
				null;
		};
	}

	function asHaxeExceptionMessageGetterTarget(callee:TypedExpr):Null<TypedExpr> {
		return switch (callee.expr) {
			case TField(target, FInstance(classRef, _, field)):
				var classType = classRef.get();
				if (isHaxeExceptionFamilyClass(classType) && field.get().name == "get_message") {
					target;
				} else {
					null;
				}
			case _:
				null;
		};
	}

	function asHaxeExceptionToStringTarget(callee:TypedExpr):Null<TypedExpr> {
		return switch (callee.expr) {
			case TField(target, FInstance(classRef, _, field)):
				var classType = classRef.get();
				if ((isHaxeExceptionClass(classType) || isHaxeValueExceptionClass(classType)) && field.get().name == "toString") {
					target;
				} else {
					null;
				}
			case TField(target, FAnon(field)) | TField(target, FClosure(_, field)):
				if ((isHaxeExceptionType(target.t) || isHaxeValueExceptionType(target.t)) && field.get().name == "toString") {
					target;
				} else {
					null;
				}
			case TMeta(_, inner):
				asHaxeExceptionToStringTarget(inner);
			case TParenthesis(inner):
				asHaxeExceptionToStringTarget(inner);
			case TCast(inner, _):
				asHaxeExceptionToStringTarget(inner);
			case _:
				null;
		};
	}

	function asHaxeValueExceptionUnwrapTarget(callee:TypedExpr):Null<TypedExpr> {
		return switch (callee.expr) {
			case TField(target, FInstance(_, _, field)) | TField(target, FAnon(field)) | TField(target, FClosure(_, field)):
				if (isHaxeValueExceptionType(target.t) && field.get().name == "unwrap") {
					target;
				} else {
					null;
				}
			case TMeta(_, inner):
				asHaxeValueExceptionUnwrapTarget(inner);
			case TParenthesis(inner):
				asHaxeValueExceptionUnwrapTarget(inner);
			case TCast(inner, _):
				asHaxeValueExceptionUnwrapTarget(inner);
			case _:
				null;
		};
	}

	function unsupportedExpr(expr:TypedExpr, message:String):LoweredExpr {
		Context.fatalError(message + " :: " + Std.string(expr.expr), expr.pos);
		return {expr: GoExpr.GoNil, isStringLike: false};
	}
	#end
}
