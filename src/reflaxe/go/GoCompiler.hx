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
import reflaxe.go.compiler.GoMetalTypeEligibility;
import reflaxe.go.compiler.GoMetalTypeEligibility.GoMetalEligibilityRole;
import reflaxe.go.compiler.GoMetalTypeEligibility.GoMetalTypeEligibilityResult;
import reflaxe.go.compiler.GoSourceModuleRegistry;
import reflaxe.go.compiler.GoSourceOwnedStdlibPlanner;
import reflaxe.go.compiler.GoStdlibShimClassifier;
import reflaxe.go.compiler.GoStdlibOwnership;
import reflaxe.go.compiler.GoTestAstFixtureEmitter;
import reflaxe.go.compiler.GoTypeMapper;
import reflaxe.go.compiler.emit.GoTypeReflectionEmitter;
import reflaxe.go.compiler.emit.GoRttiMetadataEmitter;
import reflaxe.go.compiler.emit.GoRegexSerializerEmitter;
import reflaxe.go.compiler.emit.GoNetSocketEmitter;
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
import reflaxe.go.naming.GoNaming;
#end

typedef GoGeneratedFile = {
	final relativePath:String;
	final contents:String;
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

private typedef MetalMapTypePair = {
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

/**
	What:
	Represents the lowered write site for mutating an `Array` through a target
	expression.

	Why:
	Some targets, especially anonymous-object fields lowered onto the current
	`map[string]any` carrier, cannot be mutated safely with a raw `append(...)`
	against the original lvalue. We need a read -> mutate temp -> write-back plan
	that still evaluates the target expression only once. Plain local array
	variables do not need that extra plan because assigning the local directly is
	already single-evaluation-safe.

	How:
	`lowerArrayMutationSite()` computes any required prefix statements, exposes the
	slice expression that later push/pop code mutates, and returns a write-back
	closure when the final slice must be stored back into another carrier.
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
	final staticFunctionInfos:Map<String, FunctionInfo>;
	final localFunctionScopes:Array<Map<String, FunctionInfo>>;
	final localLambdaAliasScopes:Array<Map<String, String>>;
	final localRestIteratorScopes:Array<Array<String>>;
	final requiredStdlibShimGroups:Map<String, Bool>;
	final requiredMetalChanElementTypes:Map<String, Bool>;
	final requiredMetalSliceElementTypes:Map<String, Bool>;
	final requiredMetalMapTypePairs:Map<String, MetalMapTypePair>;
	final requiredMetalResultElementTypes:Map<String, Bool>;
	final externImportPaths:Map<String, Bool>;
	final externImportPackages:Map<String, String>;
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
	var cachedVoidType:Null<Type>;
	var requiresIoHelperSurface:Bool;
	var requiresSysCommandSurface:Bool;
	var requiresIoStringInputSurface:Bool;
	var requiresIoBufferInputSurface:Bool;
	var requiresIoEofStringSurface:Bool;
	var requiresUdpSocketSurface:Bool;
	var requiresReflectFieldsShim:Bool;
	var projectClasses:Array<ClassType>;
	var projectEnums:Array<EnumType>;
	final availableClassesByName:Map<String, ClassType>;
	final pendingRequiredClassesByName:Map<String, ClassType>;
	final availableEnumsByName:Map<String, EnumType>;
	final pendingRequiredEnumsByName:Map<String, EnumType>;
	final requiredSourceOwnedClassNames:Map<String, Bool>;
	var globalLeafReceiverTypes:Map<String, Bool>;
	var tempVarCounter:Int;
	var requiresTypeValueSupport:Bool;
	#end

	public function new(?compilationContext:CompilationContext) {
		#if macro
		this.compilationContext = compilationContext == null ? new CompilationContext(GoProfile.Portable, "snapshot") : compilationContext;
		staticFunctionInfos = new Map<String, FunctionInfo>();
		localFunctionScopes = [];
		localLambdaAliasScopes = [];
		localRestIteratorScopes = [];
		requiredStdlibShimGroups = new Map<String, Bool>();
		requiredMetalChanElementTypes = new Map<String, Bool>();
		requiredMetalSliceElementTypes = new Map<String, Bool>();
		requiredMetalMapTypePairs = new Map<String, MetalMapTypePair>();
		requiredMetalResultElementTypes = new Map<String, Bool>();
		externImportPaths = new Map<String, Bool>();
		externImportPackages = new Map<String, String>();
		functionVarNameScopes = [];
		functionVarNameCountScopes = [];
		optionalPrimitiveParamScopes = [];
		nonNullPrimitiveLocalScopes = [];
		narrowedPrimitiveStorageScopes = [];
		localNeverReassignedScopes = [];
		functionReturnTypeScopes = [];
		returnRedirectScopes = [];
		cachedVoidType = null;
		requiresIoHelperSurface = false;
		requiresSysCommandSurface = false;
		requiresIoStringInputSurface = false;
		requiresIoBufferInputSurface = false;
		requiresIoEofStringSurface = false;
		requiresUdpSocketSurface = false;
		requiresReflectFieldsShim = false;
		projectClasses = [];
		projectEnums = [];
		availableClassesByName = new Map<String, ClassType>();
		pendingRequiredClassesByName = new Map<String, ClassType>();
		availableEnumsByName = new Map<String, EnumType>();
		pendingRequiredEnumsByName = new Map<String, EnumType>();
		requiredSourceOwnedClassNames = new Map<String, Bool>();
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
			requireStdlibShimGroup: requireStdlibShimGroup,
			markIoSourceOwnedHelperSurfaceRequired: function() requiresIoHelperSurface = true
		});
		lambdaIterableLowering = new GoLambdaIterableLowering({
			lowerExpr: lowerExpr,
			freshTempName: freshTempName,
			isArrayType: isArrayType,
			arrayElementType: arrayElementType,
			arrayElementGoType: arrayElementGoType,
			haxeDsListElementType: haxeDsListElementType,
			scalarGoType: scalarGoType,
			requireStdlibShimGroup: requireStdlibShimGroup,
			lowerNullableAwareTypeAssertExpr: lowerNullableAwareTypeAssertExpr,
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
		sourceModuleRegistry.rebuild(classes, enums);
		globalLeafReceiverTypes = buildGlobalLeafReceiverTypes(projectClasses);
		syncCompilationContextLeafReceivers();
		clearBoolMap(compilationContext.leafReturningFunctions);
		requiresIoHelperSurface = false;
		requiresSysCommandSurface = false;
		requiresIoStringInputSurface = false;
		requiresIoBufferInputSurface = false;
		requiresIoEofStringSurface = false;
		requiresUdpSocketSurface = false;
		requiresReflectFieldsShim = false;
		resetRequiredMetalChanElementTypes();
		resetRequiredMetalSliceElementTypes();
		resetRequiredMetalMapTypePairs();
		resetRequiredMetalResultElementTypes();
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
		applyStdlibShimGroupDependencies();
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
		compilationContext.requiresIoHelperSurface = requiresIoHelperSurface;
		var inferredRuntimeFeatures = inferRuntimeFeatures(requiredShimGroups);
		compilationContext.inferredHxrtFeatures = inferredRuntimeFeatures.features;
		compilationContext.inferredHxrtFeatureReasons = inferredRuntimeFeatures.reasons;

		var supportImports = buildSupportImports();
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
			if (moduleName == "Main") {
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

		return generated;
	}

	function appendModuleDecls(bucket:Map<String, Array<GoDecl>>, moduleName:String, decls:Array<GoDecl>):Void {
		if (decls.length == 0) {
			return;
		}
		var key = moduleName == null || moduleName == "" ? "Main" : moduleName;
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
		if (moduleName == "Main") {
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
		if (requiredStdlibShimGroups.exists("http")) {
			imports.push("bytes");
			imports.push("io");
			imports.push("net");
			imports.push("net/http");
			imports.push("net/url");
			imports.push("strings");
			imports.push("time");
		}
		if (requiredStdlibShimGroups.exists("io") && requiresIoHelperSurface) {
			imports.push("math");
		}
		if (requiredStdlibShimGroups.exists("io") && compilationContext.rawNativeMode == RawNativeMode.Utf16LE) {
			imports.push("unicode/utf16");
		}
		if (requiredStdlibShimGroups.exists("stdlib_symbols")) {
			imports.push("bytes");
			imports.push("compress/zlib");
			imports.push("crypto/md5");
			imports.push("crypto/sha1");
			imports.push("crypto/sha256");
			imports.push("encoding/base64");
			imports.push("encoding/hex");
			imports.push("encoding/xml");
			imports.push("io");
			imports.push("math");
			imports.push("path/filepath");
			imports.push("reflect");
			imports.push("strings");
			imports.push("time");
		}
		if (requiredStdlibShimGroups.exists("template_support")) {
			imports.push("reflect");
		}
		if (requiredStdlibShimGroups.exists("filesystem")) {
			imports.push("os");
			imports.push("path/filepath");
		}
		if (requiredStdlibShimGroups.exists("regex_serializer")) {
			imports.push("encoding/base64");
			imports.push("math");
			imports.push("reflect");
			imports.push("regexp");
			imports.push("sort");
			imports.push("strconv");
			imports.push("strings");
			imports.push("time");
			imports.push("unsafe");
		}
		if (requiredStdlibShimGroups.exists("net_socket")) {
			imports.push("bufio");
			imports.push("net");
			imports.push("os");
			imports.push("strconv");
			imports.push("strings");
			if (requiresUdpSocketSurface) {
				imports.push("syscall");
			}
			imports.push("time");
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
			imports: candidateImports,
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
		var filtered = new Array<String>();
		for (path in file.imports) {
			if (path == null) {
				continue;
			}
			var trimmed = StringTools.trim(path);
			if (trimmed == "") {
				continue;
			}
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
			case GoAssign(left, right): exprUsesImportAlias(left, alias) || exprUsesImportAlias(right, alias);
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
			case GoBreak, GoContinue:
				false;
			case GoReturn(expr): expr != null && exprUsesImportAlias(expr, alias);
		};
	}

	function selectClauseUsesImportAlias(clause:GoSelectClause, alias:String):Bool {
		return switch (clause) {
			case GoSelectSend(channel, value): exprUsesImportAlias(channel, alias) || exprUsesImportAlias(value, alias);
			case GoSelectRecv(recv):
				exprUsesImportAlias(recv, alias);
			case GoSelectRecvAssign(target, recv, _): exprUsesImportAlias(target, alias) || exprUsesImportAlias(recv, alias);
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
			case GoArrayLiteral(elementType, elements):
				if (typeNameUsesImportAlias(elementType, alias)) {
					true;
				} else {
					var used = false;
					for (element in elements) {
						if (exprUsesImportAlias(element, alias)) {
							used = true;
							break;
						}
					}
					used;
				}
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

	function typeNameUsesImportAlias(typeName:String, alias:String):Bool {
		if (typeName == null || typeName == "") {
			return false;
		}
		return rawCodeUsesImportAlias(typeName, alias);
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
		ensureMainClass(normalized);
		return normalized;
	}

	function ensureMainClass(classes:Array<ClassType>):Void {
		for (classType in classes) {
			if (fullClassName(classType) == "Main") {
				return;
			}
		}
		Context.fatalError("Main class was not found among project modules", Context.currentPos());
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
		if (isRequiredStdlibClass(classType)) {
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

	function goStringPointerArrayLiteral(values:Array<String>):String {
		if (values.length == 0) {
			return "[]*string{}";
		}
		var entries = [for (value in values) "hxrt.StringFromLiteral(" + goRawQuotedString(value) + ")"];
		return "[]*string{" + entries.join(", ") + "}";
	}

	/**
		What: materialize the backend-owned `haxe.Resource.content` table.

		Why: `haxe.Resource` methods can come from source-owned std inclusion, but the
		actual resource payloads are exposed to targets through compiler resources
		(`Context.getResources()` / `__resources__()`), not reusable Haxe source. If we
		do nothing, generated Go has the helper methods but an empty content table.

		How: sort resource names for deterministic output and emit the existing
		`Array<{name,data,str}>` shape as `[]map[string]any`, storing every payload in the
		`data` field as base64 so both text and binary resources flow through the std
		`getString` / `getBytes` decode paths unchanged.
	**/
	function haxeResourceContentLiteral():GoExpr {
		var resources = Context.getResources();
		var names = [for (name in resources.keys()) name];
		names.sort(function(a, b) return Reflect.compare(a, b));
		if (names.length == 0) {
			return GoExpr.GoRaw("[]map[string]any{}");
		}

		var entries = new Array<String>();
		for (name in names) {
			var bytes = resources.get(name);
			var encoded = bytes == null ? "" : haxe.crypto.Base64.encode(bytes);
			entries.push('map[string]any{"name": hxrt.StringFromLiteral(' + goRawQuotedString(name) + '), "data": hxrt.StringFromLiteral('
				+ goRawQuotedString(encoded) + '), "str": nil}');
		}
		return GoExpr.GoRaw("[]map[string]any{" + entries.join(", ") + "}");
	}

	function classHasInstanceLayout(classType:ClassType):Bool {
		var instanceDataCount = 0;
		var instanceMethodCount = 0;
		for (field in classType.fields.get()) {
			switch (field.kind) {
				case FVar(_, _):
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

	function serializerClassMetadata():Array<{goTypeName:String, haxeTypeName:String}> {
		var entries = new Array<{goTypeName:String, haxeTypeName:String}>();
		for (classType in projectClasses) {
			if (classType.isExtern || classType.isInterface) {
				continue;
			}
			switch (classType.kind) {
				case KTypeParameter(_):
					continue;
				case _:
			}
			if (!classHasInstanceLayout(classType)) {
				continue;
			}
			entries.push({goTypeName: classTypeName(classType), haxeTypeName: fullClassName(classType)});
		}
		entries.sort(function(a, b) return Reflect.compare(a.goTypeName, b.goTypeName));
		return entries;
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

	function serializerEnumMetadata():Array<{goTypeName:String, haxeTypeName:String, constructors:Array<String>}> {
		var entries = new Array<{goTypeName:String, haxeTypeName:String, constructors:Array<String>}>();
		for (enumType in projectEnums) {
			if (enumType.isExtern) {
				continue;
			}
			var constructors = [for (field in enumType.constructs) field];
			constructors.sort(function(a, b) return a.index - b.index);
			entries.push({
				goTypeName: enumTypeName(enumType),
				haxeTypeName: fullEnumName(enumType),
				constructors: [for (constructor in constructors) constructor.name]
			});
		}
		entries.sort(function(a, b) return Reflect.compare(a.goTypeName, b.goTypeName));
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

	function projectSuperClass(classType:ClassType):Null<ClassType> {
		if (classType.superClass == null) {
			return null;
		}
		var superType = classType.superClass.t.get();
		var superName = fullClassName(superType);
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
				if (receiver == null && results.length == 1 && globalLeafReceiverTypes.exists(results[0])) {
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
			if (ctorArgs.length == 0) {
				decls.push(GoDecl.GoGlobalVarDecl(symbol, "*" + enumName, GoExpr.GoRaw("&" + enumName + "{tag: " + constructor.index + "}")));
			} else {
				var params = new Array<GoParam>();
				var payloadExprs = new Array<GoExpr>();
				for (index in 0...ctorArgs.length) {
					var arg = ctorArgs[index];
					var argName = normalizeIdent(arg.name == "" ? ("arg" + index) : arg.name);
					params.push({
						name: argName,
						typeName: scalarGoType(arg.t)
					});
					payloadExprs.push(GoExpr.GoIdent(argName));
				}

				decls.push(GoDecl.GoFuncDecl(symbol, null, params, ["*" + enumName], [
					GoStmt.GoVarDecl("enumValue", null, GoExpr.GoRaw("&" + enumName + "{tag: " + constructor.index + "}"), true),
					GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("enumValue"), "params"), GoExpr.GoArrayLiteral("any", payloadExprs)),
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
		applyStdlibShimGroupDependencies();
		if (requiredStdlibShimGroups.exists("io")) {
			decls = decls.concat(lowerIoStdlibShimDecls());
		}
		if (requiredStdlibShimGroups.exists("ds")) {
			decls = decls.concat(lowerDsStdlibShimDecls());
		}
		if (requiredStdlibShimGroups.exists("http")) {
			decls = decls.concat(lowerHttpStdlibShimDecls());
		}
		if (requiredStdlibShimGroups.exists("sys")) {
			decls = decls.concat(lowerSysStdlibShimDecls());
		}
		if (requiredStdlibShimGroups.exists("filesystem")) {
			decls = decls.concat(lowerFileSystemShimDecls());
		}
		if (requiredStdlibShimGroups.exists("stdlib_symbols")) {
			decls = decls.concat(lowerStdlibSymbolShimDecls());
		}
		if (requiredStdlibShimGroups.exists("template_support")) {
			decls = decls.concat(lowerTemplateSupportShimDecls());
		}
		if (requiredStdlibShimGroups.exists("regex_serializer")) {
			decls = decls.concat(lowerRegexSerializerShimDecls());
		}
		if (requiredStdlibShimGroups.exists("net_socket")) {
			decls = decls.concat(lowerNetSocketShimDecls());
		}
		if (requiredStdlibShimGroups.exists("atomic")) {
			decls = decls.concat(lowerAtomicStdlibShimDecls());
		}
		if (requiredStdlibShimGroups.exists("go_concurrency")) {
			decls = decls.concat(lowerGoConcurrencyShimDecls());
		}
		if (requiredStdlibShimGroups.exists("go_collections")) {
			decls = decls.concat(lowerMetalGoCollectionShimDecls());
		}
		if (requiredStdlibShimGroups.exists("go_result")) {
			decls = decls.concat(lowerGoResultShimDecls());
		}
		return decls;
	}

	function applyStdlibShimGroupDependencies():Void {
		if (requiredStdlibShimGroups.exists("stdlib_symbols")) {
			// Symbol shims include crypto/xml/zip helpers that depend on haxe.io.Bytes.
			requireStdlibShimGroup("io");
		}
		if (requiredStdlibShimGroups.exists("filesystem")) {
			// FileSystem stat/date fields rely on Date symbol declarations.
			requireStdlibShimGroup("stdlib_symbols");
			requireStdlibShimGroup("io");
		}
		if (requiredStdlibShimGroups.exists("sys")) {
			// sys.io.File handles and byte helpers build on haxe.io surfaces and
			// reuse the stdlib-symbol raw-byte bridge.
			requireStdlibShimGroup("stdlib_symbols");
			requireStdlibShimGroup("io");
			requireIoSourceOwnedHelperSurface();
		}
		if (requiredStdlibShimGroups.exists("http")) {
			// Http request shims expose and consume haxe.io.Bytes payloads.
			requireStdlibShimGroup("io");
			requireStdlibShimGroup("ds");
		}
		if (requiredStdlibShimGroups.exists("ds")) {
			// DS shim declarations expose haxe.IMap, which is staged std.
			requireSourceOwnedStdlibModule("haxe.Constraints");
		}
		if (requiredStdlibShimGroups.exists("regex_serializer")) {
			// Serializer token support includes haxe.ds.List/StringMap/IntMap/ObjectMap families.
			requireStdlibShimGroup("ds");
		}
	}

	function lowerIoStdlibShimDecls():Array<GoDecl> {
		var rawNativeUtf16 = compilationContext.rawNativeMode == RawNativeMode.Utf16LE;
		var decls = [
			GoDecl.GoStructDecl("haxe__io__Encoding", [{name: "tag", typeName: "int"}]),
			GoDecl.GoInterfaceDecl("haxe__io__Input", [
				{
					name: "get_bigEndian",
					params: [],
					results: ["bool"]
				},
				{
					name: "set_bigEndian",
					params: [{name: "e", typeName: "bool"}],
					results: ["bool"]
				},
				{
					name: "readByte",
					params: [],
					results: ["int"]
				},
				{
					name: "readBytes",
					params: [
						{name: "buf", typeName: "*haxe__io__Bytes"},
						{name: "pos", typeName: "int"},
						{name: "len", typeName: "int"}
					],
					results: ["int"]
				},
				{
					name: "close",
					params: [],
					results: []
				},
				{
					name: "readAll",
					params: [{name: "bufsize", typeName: "...int"}],
					results: ["*haxe__io__Bytes"]
				},
				{
					name: "readFullBytes",
					params: [
						{name: "s", typeName: "*haxe__io__Bytes"},
						{name: "pos", typeName: "int"},
						{name: "len", typeName: "int"}
					],
					results: []
				},
				{
					name: "read",
					params: [{name: "nbytes", typeName: "int"}],
					results: ["*haxe__io__Bytes"]
				},
				{
					name: "readUntil",
					params: [{name: "end", typeName: "int"}],
					results: ["*string"]
				},
				{
					name: "readLine",
					params: [],
					results: ["*string"]
				},
				{
					name: "readFloat",
					params: [],
					results: ["float64"]
				},
				{
					name: "readDouble",
					params: [],
					results: ["float64"]
				},
				{
					name: "readInt8",
					params: [],
					results: ["int"]
				},
				{
					name: "readInt16",
					params: [],
					results: ["int"]
				},
				{
					name: "readUInt16",
					params: [],
					results: ["int"]
				},
				{
					name: "readInt24",
					params: [],
					results: ["int"]
				},
				{
					name: "readUInt24",
					params: [],
					results: ["int"]
				},
				{
					name: "readInt32",
					params: [],
					results: ["int"]
				},
				{
					name: "readString",
					params: [
						{name: "len", typeName: "int"},
						{name: "encoding", typeName: "...*haxe__io__Encoding"}
					],
					results: ["*string"]
				}
			]),
			GoDecl.GoInterfaceDecl("haxe__io__Output",
				[
					{
						name: "get_bigEndian",
						params: [],
						results: ["bool"]
					},
					{
						name: "set_bigEndian",
						params: [{name: "e", typeName: "bool"}],
						results: ["bool"]
					},
					{
						name: "writeByte",
						params: [{name: "c", typeName: "int"}],
						results: []
					},
					{
						name: "writeBytes",
						params: [
							{name: "s", typeName: "*haxe__io__Bytes"},
							{name: "pos", typeName: "int"},
							{name: "len", typeName: "int"}
						],
						results: ["int"]
					},
					{
						name: "flush",
						params: [],
						results: []
					},
					{
						name: "close",
						params: [],
						results: []
					},
					{
						name: "write",
						params: [{name: "s", typeName: "*haxe__io__Bytes"}],
						results: []
					},
					{
						name: "writeFullBytes",
						params: [
							{name: "s", typeName: "*haxe__io__Bytes"},
							{name: "pos", typeName: "int"},
							{name: "len", typeName: "int"}
						],
						results: []
					},
					{
						name: "writeFloat",
						params: [{name: "x", typeName: "float64"}],
						results: []
					},
					{
						name: "writeDouble",
						params: [{name: "x", typeName: "float64"}],
						results: []
					},
					{
						name: "writeInt8",
						params: [{name: "x", typeName: "int"}],
						results: []
					},
					{
						name: "writeInt16",
						params: [{name: "x", typeName: "int"}],
						results: []
					},
					{
						name: "writeUInt16",
						params: [{name: "x", typeName: "int"}],
						results: []
					},
					{
						name: "writeInt24",
						params: [{name: "x", typeName: "int"}],
						results: []
					},
					{
						name: "writeUInt24",
						params: [{name: "x", typeName: "int"}],
						results: []
					},
					{
						name: "writeInt32",
						params: [{name: "x", typeName: "int"}],
						results: []
					},
					{
						name: "prepare",
						params: [{name: "nbytes", typeName: "int"}],
						results: []
					},
					{
						name: "writeInput",
						params: [{name: "i", typeName: "haxe__io__Input"}, {name: "bufsize", typeName: "...int"}],
						results: []
					},
					{
						name: "writeString",
						params: [
							{name: "s", typeName: "*string"},
							{name: "encoding", typeName: "...*haxe__io__Encoding"}
						],
						results: []
					}
				]),
			GoDecl.GoStructDecl("haxe__io__Eof", []),
			GoDecl.GoStructDecl("haxe__io__Error", [{name: "tag", typeName: "int"}, {name: "params", typeName: "[]any"}]),
			GoDecl.GoStructDecl("haxe__io__Bytes", [
				{
					name: "b",
					typeName: "[]int"
				},
				{name: "length", typeName: "int"},
				{name: "__hx_raw", typeName: "[]byte"},
				{name: "__hx_rawValid", typeName: "bool"}
			]),
			GoDecl.GoStructDecl("haxe__io__BytesBuffer", [{name: "b", typeName: "[]int"}]),
			GoDecl.GoStructDecl("haxe__io__BytesInput",
				[
					{name: "bigEndian", typeName: "bool"},
					{name: "b", typeName: "[]int"},
					{name: "pos", typeName: "int"},
					{name: "len", typeName: "int"},
					{name: "totlen", typeName: "int"}
				]),
			GoDecl.GoStructDecl("haxe__io__StringInput", [{name: "haxe__io__BytesInput", typeName: "*haxe__io__BytesInput"}]),
			GoDecl.GoStructDecl("haxe__io__BufferInput", [
				{name: "bigEndian", typeName: "bool"},
				{name: "i", typeName: "haxe__io__Input"},
				{name: "buf", typeName: "*haxe__io__Bytes"},
				{name: "available", typeName: "int"},
				{name: "pos", typeName: "int"}
			]),
			GoDecl.GoStructDecl("haxe__io__BytesOutput", [
				{name: "bigEndian", typeName: "bool"},
				{name: "b", typeName: "*haxe__io__BytesBuffer"}
			]),
			GoDecl.GoFuncDecl("New_haxe__io__Input", null, [], ["haxe__io__Input"], [
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("New_haxe__io__BytesInput"), [GoExpr.GoRaw("&haxe__io__Bytes{b: []int{}, length: 0}")]))
			]),
			GoDecl.GoFuncDecl("New_haxe__io__StringInput", null, [
				{
					name: "s",
					typeName: "*string"
				}
			], ["*haxe__io__StringInput"], [
				GoStmt.GoReturn(GoExpr.GoRaw("&haxe__io__StringInput{haxe__io__BytesInput: New_haxe__io__BytesInput(haxe__io__Bytes_ofString(s))}"))
			]),
			GoDecl.GoFuncDecl("New_haxe__io__BufferInput", null, [
				{
					name: "i",
					typeName: "haxe__io__Input"
				},
				{name: "buf", typeName: "*haxe__io__Bytes"},
				{name: "rest", typeName: "...int"}
			], ["*haxe__io__BufferInput"],
				[
					GoStmt.GoVarDecl("resolvedPos", null, GoExpr.GoIntLiteral(0), true),
					GoStmt.GoVarDecl("resolvedAvailable", null, GoExpr.GoIntLiteral(0), true),
					GoStmt.GoIf(GoExpr.GoBinary(">", GoExpr.GoCall(GoExpr.GoIdent("len"), [GoExpr.GoIdent("rest")]), GoExpr.GoIntLiteral(0)), [
						GoStmt.GoAssign(GoExpr.GoIdent("resolvedPos"), GoExpr.GoIndex(GoExpr.GoIdent("rest"), GoExpr.GoIntLiteral(0)))
					], null),
					GoStmt.GoIf(GoExpr.GoBinary(">", GoExpr.GoCall(GoExpr.GoIdent("len"), [GoExpr.GoIdent("rest")]), GoExpr.GoIntLiteral(1)), [
						GoStmt.GoAssign(GoExpr.GoIdent("resolvedAvailable"), GoExpr.GoIndex(GoExpr.GoIdent("rest"), GoExpr.GoIntLiteral(1)))
					],
						null),
					GoStmt.GoReturn(GoExpr.GoRaw("&haxe__io__BufferInput{i: i, buf: buf, pos: resolvedPos, available: resolvedAvailable}"))
				]),
			GoDecl.GoFuncDecl("New_haxe__io__Output", null, [], ["haxe__io__Output"],
				[GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("New_haxe__io__BytesOutput"), []))]),
			GoDecl.GoFuncDecl("New_haxe__io__Eof", null, [], ["*haxe__io__Eof"], [GoStmt.GoReturn(GoExpr.GoRaw("&haxe__io__Eof{}"))]),
			GoDecl.GoFuncDecl("String", {
				name: "self",
				typeName: "*haxe__io__Eof"
			}, [], ["string"],
				[
					GoStmt.GoAssign(GoExpr.GoIdent("_"), GoExpr.GoIdent("self")),
					GoStmt.GoReturn(GoExpr.GoStringLiteral("Eof"))
				]),
			GoDecl.GoGlobalVarDecl("haxe__io__Encoding_UTF8", "*haxe__io__Encoding", GoExpr.GoRaw("&haxe__io__Encoding{tag: 0}")),
			GoDecl.GoGlobalVarDecl("haxe__io__Encoding_RawNative", "*haxe__io__Encoding", GoExpr.GoRaw("&haxe__io__Encoding{tag: 1}")),
			GoDecl.GoFuncDecl("String", {
				name: "self",
				typeName: "*haxe__io__Encoding"
			}, [], ["string"], [
				GoStmt.GoRaw("if self == nil {"),
				GoStmt.GoRaw("\treturn \"null\""),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("switch self.tag {"),
				GoStmt.GoRaw("case 0:"),
				GoStmt.GoRaw("\treturn \"UTF8\""),
				GoStmt.GoRaw("case 1:"),
				GoStmt.GoRaw("\treturn \"RawNative\""),
				GoStmt.GoRaw("default:"),
				GoStmt.GoRaw("\treturn \"Encoding\""),
				GoStmt.GoRaw("}")
			]),
			GoDecl.GoFuncDecl("toString", {
				name: "self",
				typeName: "*haxe__io__Encoding"
			}, [], ["*string"], [
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"),
					[GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("self"), "String"), [])]))
			]),
			GoDecl.GoFuncDecl("haxe__io__resolveEncodingTag", null, [
				{
					name: "encoding",
					typeName: "...*haxe__io__Encoding"
				}
			], ["int"], [
				GoStmt.GoRaw("if len(encoding) == 0 || encoding[0] == nil {"),
				GoStmt.GoRaw("\treturn 0"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("return encoding[0].tag")
			]),
			// Keep RawNative string conversion in the compiler: these helpers co-own
			// raw-native mode switching and the __hx_raw cache contract that
			// downstream raw-byte consumers rely on after Bytes mutation.
			GoDecl.GoFuncDecl("haxe__io__bytes_fromStringRawNativeUTF16LE", null, [
				{
					name: "value",
					typeName: "*string"
				}
			], ["*haxe__io__Bytes"], rawNativeUtf16 ? [
				GoStmt.GoRaw("runes := []rune(*hxrt.StdString(value))"),
				GoStmt.GoRaw("units := utf16.Encode(runes)"),
				GoStmt.GoRaw("raw := make([]byte, len(units)*2)"),
				GoStmt.GoRaw("for i := 0; i < len(units); i++ {"),
				GoStmt.GoRaw("\tunit := units[i]"),
				GoStmt.GoRaw("\traw[i*2] = byte(unit)"),
				GoStmt.GoRaw("\traw[i*2+1] = byte(unit >> 8)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("converted := make([]int, len(raw))"),
				GoStmt.GoRaw("for i := 0; i < len(raw); i++ {"),
				GoStmt.GoRaw("\tconverted[i] = int(raw[i])"),
				GoStmt.GoRaw("}"),
				GoStmt.GoReturn(GoExpr.GoRaw("&haxe__io__Bytes{b: converted, length: len(converted), __hx_raw: raw, __hx_rawValid: true}"))
			] : [
				GoStmt.GoRaw("raw := []byte(*hxrt.StdString(value))"),
				GoStmt.GoRaw("converted := make([]int, len(raw))"),
				GoStmt.GoRaw("for i := 0; i < len(raw); i++ {"),
				GoStmt.GoRaw("\tconverted[i] = int(raw[i])"),
				GoStmt.GoRaw("}"),
				GoStmt.GoReturn(GoExpr.GoRaw("&haxe__io__Bytes{b: converted, length: len(converted), __hx_raw: raw, __hx_rawValid: true}"))
			]),
			GoDecl.GoFuncDecl("haxe__io__bytes_toStringRawNativeUTF16LE", null, [
				{
					name: "value",
					typeName: "[]int"
				}
			], ["*string"], rawNativeUtf16 ? [
				GoStmt.GoRaw("if len(value) == 0 {"),
				GoStmt.GoRaw("\treturn hxrt.StringFromLiteral(\"\")"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("limit := len(value)"),
				GoStmt.GoRaw("if (limit & 1) == 1 {"),
				GoStmt.GoRaw("\tlimit--"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("units := make([]uint16, limit/2)"),
				GoStmt.GoRaw("for i := 0; i < len(units); i++ {"),
				GoStmt.GoRaw("\tlow := uint16(value[i*2] & 0xFF)"),
				GoStmt.GoRaw("\thigh := uint16(value[i*2+1] & 0xFF)"),
				GoStmt.GoRaw("\tunits[i] = low | (high << 8)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("return hxrt.StringFromLiteral(string(utf16.Decode(units)))")
			] : [
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.BytesToString"), [GoExpr.GoIdent("value")]))
			]),
			GoDecl.GoFuncDecl("toString", {
				name: "self",
				typeName: "*haxe__io__Eof"
			}, [], ["*string"],
				[
					GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoStringLiteral("Eof")]))
				]),
			GoDecl.GoGlobalVarDecl("haxe__io__Error_Blocked", "*haxe__io__Error", GoExpr.GoRaw("&haxe__io__Error{tag: 0}")),
			GoDecl.GoGlobalVarDecl("haxe__io__Error_Overflow", "*haxe__io__Error", GoExpr.GoRaw("&haxe__io__Error{tag: 1}")),
			GoDecl.GoGlobalVarDecl("haxe__io__Error_OutsideBounds", "*haxe__io__Error", GoExpr.GoRaw("&haxe__io__Error{tag: 2}")),
			GoDecl.GoFuncDecl("haxe__io__Error_Custom", null, [
				{
					name: "e",
					typeName: "any"
				}
			], ["*haxe__io__Error"],
				[GoStmt.GoReturn(GoExpr.GoRaw("&haxe__io__Error{tag: 3, params: []any{e}}"))]),
			GoDecl.GoFuncDecl("String", {
				name: "self",
				typeName: "*haxe__io__Error"
			}, [], ["string"], [
				GoStmt.GoRaw("if self == nil {"),
				GoStmt.GoRaw("\treturn \"null\""),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("switch self.tag {"),
				GoStmt.GoRaw("case 0:"),
				GoStmt.GoRaw("\treturn \"Blocked\""),
				GoStmt.GoRaw("case 1:"),
				GoStmt.GoRaw("\treturn \"Overflow\""),
				GoStmt.GoRaw("case 2:"),
				GoStmt.GoRaw("\treturn \"OutsideBounds\""),
				GoStmt.GoRaw("case 3:"),
				GoStmt.GoRaw("\tif len(self.params) == 0 {"),
				GoStmt.GoRaw("\t\treturn \"Custom(null)\""),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn \"Custom(\" + *hxrt.StdString(self.params[0]) + \")\""),
				GoStmt.GoRaw("default:"),
				GoStmt.GoRaw("\treturn \"Error\""),
				GoStmt.GoRaw("}")
			]),
			GoDecl.GoFuncDecl("toString", {
				name: "self",
				typeName: "*haxe__io__Error"
			}, [], ["*string"], [
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"),
					[GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("self"), "String"), [])]))
			]),
			GoDecl.GoFuncDecl("New_haxe__io__Bytes", null, [
				{
					name: "length",
					typeName: "int"
				},
				{name: "b", typeName: "[]int"}
			], ["*haxe__io__Bytes"], [
				GoStmt.GoIf(GoExpr.GoBinary("==", GoExpr.GoIdent("b"), GoExpr.GoNil),
					[GoStmt.GoAssign(GoExpr.GoIdent("b"), GoExpr.GoRaw("make([]int, length)"))], null),
				GoStmt.GoReturn(GoExpr.GoRaw("&haxe__io__Bytes{b: b, length: len(b)}"))
			]),
			GoDecl.GoFuncDecl("haxe__io__Bytes_alloc", null, [
				{
					name: "length",
					typeName: "int"
				}
			], ["*haxe__io__Bytes"], [
				GoStmt.GoReturn(GoExpr.GoRaw("&haxe__io__Bytes{b: make([]int, length), length: length}"))
			]),
			GoDecl.GoFuncDecl("haxe__io__Bytes_ofString", null, [
				{
					name: "value",
					typeName: "*string"
				},
				{name: "encoding", typeName: "...*haxe__io__Encoding"}
			], ["*haxe__io__Bytes"], [
				GoStmt.GoRaw("if haxe__io__resolveEncodingTag(encoding...) == 1 {"),
				GoStmt.GoRaw("\treturn haxe__io__bytes_fromStringRawNativeUTF16LE(value)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("raw := []byte(*hxrt.StdString(value))"),
				GoStmt.GoRaw("converted := make([]int, len(raw))"),
				GoStmt.GoRaw("for i := 0; i < len(raw); i++ {"),
				GoStmt.GoRaw("\tconverted[i] = int(raw[i])"),
				GoStmt.GoRaw("}"),
				GoStmt.GoReturn(GoExpr.GoRaw("&haxe__io__Bytes{b: converted, length: len(converted), __hx_raw: raw, __hx_rawValid: true}"))
			]),
			GoDecl.GoFuncDecl("haxe__io__Bytes_ofData", null, [
				{
					name: "b",
					typeName: "[]int"
				}
			], ["*haxe__io__Bytes"], [
				GoStmt.GoRaw("if b == nil {"),
				GoStmt.GoRaw("\treturn &haxe__io__Bytes{b: []int{}, length: 0}"),
				GoStmt.GoRaw("}"),
				GoStmt.GoReturn(GoExpr.GoRaw("&haxe__io__Bytes{b: b, length: len(b)}"))
			]),
			GoDecl.GoFuncDecl("haxe__io__Bytes_ofHex", null, [
				{
					name: "s",
					typeName: "*string"
				}
			], ["*haxe__io__Bytes"], [
				GoStmt.GoVarDecl("decoded", "[]int", GoExpr.GoCall(GoExpr.GoIdent("hxrt.BytesOfHex"), [GoExpr.GoIdent("s")]), true),
				GoStmt.GoReturn(GoExpr.GoRaw("&haxe__io__Bytes{b: decoded, length: len(decoded)}"))
			]),
			GoDecl.GoFuncDecl("toString", {
				name: "self",
				typeName: "*haxe__io__Bytes"
			}, [], ["*string"], [
				GoStmt.GoIf(GoExpr.GoBinary("==", GoExpr.GoIdent("self"), GoExpr.GoNil), [
					GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoStringLiteral("")]))
				],
					null),
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.BytesToString"), [GoExpr.GoSelector(GoExpr.GoIdent("self"), "b")]))
			]),
			GoDecl.GoFuncDecl("toHex", {
				name: "self",
				typeName: "*haxe__io__Bytes"
			}, [], ["*string"], [
				GoStmt.GoRaw("if self == nil || self.length == 0 {"),
				GoStmt.GoRaw("\treturn hxrt.StringFromLiteral(\"\")"),
				GoStmt.GoRaw("}"),
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.BytesToHex"), [
					GoExpr.GoSelector(GoExpr.GoIdent("self"), "b"),
					GoExpr.GoSelector(GoExpr.GoIdent("self"), "length")
				]))
			]),
			GoDecl.GoFuncDecl("getData", {
				name: "self",
				typeName: "*haxe__io__Bytes"
			}, [], ["[]int"], [
				GoStmt.GoIf(GoExpr.GoBinary("==", GoExpr.GoIdent("self"), GoExpr.GoNil), [GoStmt.GoReturn(GoExpr.GoRaw("[]int{}"))], null),
				GoStmt.GoReturn(GoExpr.GoSelector(GoExpr.GoIdent("self"), "b"))
			]),
			GoDecl.GoFuncDecl("getString", {
				name: "self",
				typeName: "*haxe__io__Bytes"
			}, [
				{name: "pos", typeName: "int"},
				{name: "len", typeName: "int"},
				{name: "encoding", typeName: "...*haxe__io__Encoding"}
			], ["*string"], [
				GoStmt.GoRaw("if self == nil || pos < 0 || len < 0 || pos+len > self.length {"),
				GoStmt.GoRaw("\thxrt.Throw(haxe__io__Error_OutsideBounds)"),
				GoStmt.GoRaw("\treturn hxrt.StringFromLiteral(\"\")"),
				GoStmt.GoRaw("}"),
				GoStmt.GoVarDecl("slice", null, GoExpr.GoRaw("self.b[pos:pos+len]"), true),
				GoStmt.GoRaw("if haxe__io__resolveEncodingTag(encoding...) == 1 {"),
				GoStmt.GoRaw("\treturn haxe__io__bytes_toStringRawNativeUTF16LE(slice)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.BytesToString"), [GoExpr.GoIdent("slice")]))
			]),
			GoDecl.GoFuncDecl("readString", {
				name: "self",
				typeName: "*haxe__io__Bytes"
			}, [{name: "pos", typeName: "int"}, {name: "len", typeName: "int"}], ["*string"],
				[
					GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("self"), "getString"), [GoExpr.GoIdent("pos"), GoExpr.GoIdent("len")]))
				]),
			GoDecl.GoFuncDecl("get", {
				name: "self",
				typeName: "*haxe__io__Bytes"
			}, [{name: "pos", typeName: "int"}], ["int"], [
				GoStmt.GoReturn(GoExpr.GoIndex(GoExpr.GoSelector(GoExpr.GoIdent("self"), "b"), GoExpr.GoIdent("pos")))
			]),
			GoDecl.GoFuncDecl("set", {name: "self", typeName: "*haxe__io__Bytes"}, [{name: "pos", typeName: "int"}, {name: "value", typeName: "int"}], [], [
				GoStmt.GoAssign(GoExpr.GoIndex(GoExpr.GoSelector(GoExpr.GoIdent("self"), "b"), GoExpr.GoIdent("pos")), GoExpr.GoIdent("value")),
				GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("self"), "__hx_rawValid"), GoExpr.GoBoolLiteral(false))
			]),
			GoDecl.GoFuncDecl("blit", {
				name: "self",
				typeName: "*haxe__io__Bytes"
			}, [
				{name: "pos", typeName: "int"},
				{name: "src", typeName: "*haxe__io__Bytes"},
				{name: "srcpos", typeName: "int"},
				{name: "len", typeName: "int"}
			], [], [
				GoStmt.GoRaw("if self == nil || src == nil || pos < 0 || srcpos < 0 || len < 0 || pos+len > self.length || srcpos+len > src.length {"),
				GoStmt.GoRaw("\thxrt.Throw(haxe__io__Error_OutsideBounds)"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if len == 0 {"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if self == src && pos > srcpos {"),
				GoStmt.GoRaw("\tfor i := len - 1; i >= 0; i-- {"),
				GoStmt.GoRaw("\t\tself.b[pos+i] = src.b[srcpos+i]"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("} else {"),
				GoStmt.GoRaw("\tfor i := 0; i < len; i++ {"),
				GoStmt.GoRaw("\t\tself.b[pos+i] = src.b[srcpos+i]"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("self.__hx_rawValid = false")
			]),
			GoDecl.GoFuncDecl("fill", {
				name: "self",
				typeName: "*haxe__io__Bytes"
			}, [
				{name: "pos", typeName: "int"},
				{name: "len", typeName: "int"},
				{name: "value", typeName: "int"}
			], [], [
				GoStmt.GoRaw("if self == nil || pos < 0 || len < 0 || pos+len > self.length {"),
				GoStmt.GoRaw("\thxrt.Throw(haxe__io__Error_OutsideBounds)"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("masked := value & 255"),
				GoStmt.GoRaw("for i := 0; i < len; i++ {"),
				GoStmt.GoRaw("\tself.b[pos+i] = masked"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("self.__hx_rawValid = false")
			]),
			GoDecl.GoFuncDecl("sub", {
				name: "self",
				typeName: "*haxe__io__Bytes"
			}, [{name: "pos", typeName: "int"}, {name: "len", typeName: "int"}],
				["*haxe__io__Bytes"], [
					GoStmt.GoRaw("if self == nil || pos < 0 || len < 0 || pos+len > self.length {"),
					GoStmt.GoRaw("\thxrt.Throw(haxe__io__Error_OutsideBounds)"),
					GoStmt.GoRaw("\treturn &haxe__io__Bytes{b: []int{}, length: 0}"),
					GoStmt.GoRaw("}"),
					GoStmt.GoRaw("if len == 0 {"),
					GoStmt.GoRaw("\treturn &haxe__io__Bytes{b: []int{}, length: 0}"),
					GoStmt.GoRaw("}"),
					GoStmt.GoRaw("copied := make([]int, len)"),
					GoStmt.GoRaw("copy(copied, self.b[pos:pos+len])"),
					GoStmt.GoRaw("return &haxe__io__Bytes{b: copied, length: len}")
				]),
			GoDecl.GoFuncDecl("compare", {
				name: "self",
				typeName: "*haxe__io__Bytes"
			}, [{name: "other", typeName: "*haxe__io__Bytes"}], ["int"],
				[
					GoStmt.GoRaw("if self == nil && other == nil {"),
					GoStmt.GoRaw("\treturn 0"),
					GoStmt.GoRaw("}"),
					GoStmt.GoRaw("if self == nil {"),
					GoStmt.GoRaw("\treturn -1"),
					GoStmt.GoRaw("}"),
					GoStmt.GoRaw("if other == nil {"),
					GoStmt.GoRaw("\treturn 1"),
					GoStmt.GoRaw("}"),
					GoStmt.GoRaw("limit := self.length"),
					GoStmt.GoRaw("if other.length < limit {"),
					GoStmt.GoRaw("\tlimit = other.length"),
					GoStmt.GoRaw("}"),
					GoStmt.GoRaw("for i := 0; i < limit; i++ {"),
					GoStmt.GoRaw("\tif self.b[i] < other.b[i] {"),
					GoStmt.GoRaw("\t\treturn -1"),
					GoStmt.GoRaw("\t}"),
					GoStmt.GoRaw("\tif self.b[i] > other.b[i] {"),
					GoStmt.GoRaw("\t\treturn 1"),
					GoStmt.GoRaw("\t}"),
					GoStmt.GoRaw("}"),
					GoStmt.GoRaw("if self.length < other.length {"),
					GoStmt.GoRaw("\treturn -1"),
					GoStmt.GoRaw("}"),
					GoStmt.GoRaw("if self.length > other.length {"),
					GoStmt.GoRaw("\treturn 1"),
					GoStmt.GoRaw("}"),
					GoStmt.GoRaw("return 0")
				]),
			GoDecl.GoFuncDecl("New_haxe__io__BytesBuffer", null, [], ["*haxe__io__BytesBuffer"],
				[GoStmt.GoReturn(GoExpr.GoRaw("&haxe__io__BytesBuffer{b: []int{}}"))]),
			GoDecl.GoFuncDecl("addByte", {
				name: "self",
				typeName: "*haxe__io__BytesBuffer"
			}, [{name: "value", typeName: "int"}], [], [
				GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("self"), "b"),
					GoExpr.GoCall(GoExpr.GoIdent("hxrt.BytesBufferAddByte"), [GoExpr.GoSelector(GoExpr.GoIdent("self"), "b"), GoExpr.GoIdent("value")]))
			]),
			GoDecl.GoFuncDecl("add", {
				name: "self",
				typeName: "*haxe__io__BytesBuffer"
			}, [{name: "src", typeName: "*haxe__io__Bytes"}], [], [
				GoStmt.GoIf(GoExpr.GoBinary("==", GoExpr.GoIdent("src"), GoExpr.GoNil), [GoStmt.GoReturn(null)], null),
				GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("self"), "b"), GoExpr.GoCall(GoExpr.GoIdent("hxrt.BytesBufferAdd"), [
					GoExpr.GoSelector(GoExpr.GoIdent("self"), "b"),
					GoExpr.GoSelector(GoExpr.GoIdent("src"), "b")
				]))
			]),
			GoDecl.GoFuncDecl("addBytes", {
				name: "self",
				typeName: "*haxe__io__BytesBuffer"
			}, [
				{name: "src", typeName: "*haxe__io__Bytes"},
				{name: "pos", typeName: "int"},
				{name: "len", typeName: "int"}
			], [], [
				GoStmt.GoRaw("if src == nil || pos < 0 || len < 0 || pos+len > src.length {"),
				GoStmt.GoRaw("\thxrt.Throw(haxe__io__Error_OutsideBounds)"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if len == 0 {"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("self"), "b"), GoExpr.GoCall(GoExpr.GoIdent("hxrt.BytesBufferAddSlice"), [
					GoExpr.GoSelector(GoExpr.GoIdent("self"), "b"),
					GoExpr.GoSelector(GoExpr.GoIdent("src"), "b"),
					GoExpr.GoIdent("pos"),
					GoExpr.GoIdent("len")
				]))
			]),
			GoDecl.GoFuncDecl("addString", {
				name: "self",
				typeName: "*haxe__io__BytesBuffer"
			}, [
				{name: "value", typeName: "*string"},
				{name: "encoding", typeName: "...*haxe__io__Encoding"}
			], [], [
				GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("self"), "add"), [
					GoExpr.GoCall(GoExpr.GoIdent("haxe__io__Bytes_ofString"), [GoExpr.GoIdent("value")])
				]))
			]),
			GoDecl.GoFuncDecl("getBytes", {
				name: "self",
				typeName: "*haxe__io__BytesBuffer"
			}, [], ["*haxe__io__Bytes"], [
				GoStmt.GoVarDecl("copied", null, GoExpr.GoCall(GoExpr.GoIdent("hxrt.BytesClone"), [GoExpr.GoSelector(GoExpr.GoIdent("self"), "b")]), true),
				GoStmt.GoReturn(GoExpr.GoRaw("&haxe__io__Bytes{b: copied, length: len(copied)}"))
			]),
			GoDecl.GoFuncDecl("get_length", {
				name: "self",
				typeName: "*haxe__io__BytesBuffer"
			}, [], ["int"], [
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.BytesBufferLength"), [GoExpr.GoSelector(GoExpr.GoIdent("self"), "b")]))
			]),
			GoDecl.GoFuncDecl("haxe__io__input_isEof", null, [
				{
					name: "value",
					typeName: "any"
				}
			], ["bool"], [
				GoStmt.GoRaw("_, ok := value.(*haxe__io__Eof)"),
				GoStmt.GoReturn(GoExpr.GoIdent("ok"))
			]),
			GoDecl.GoFuncDecl("haxe__io__input_readAll", null, [
				{
					name: "self",
					typeName: "haxe__io__Input"
				},
				{name: "bufsize", typeName: "...int"}
			], ["*haxe__io__Bytes"], [
				GoStmt.GoRaw("resolved := 1 << 14"),
				GoStmt.GoRaw("if len(bufsize) > 0 {"),
				GoStmt.GoRaw("\tresolved = bufsize[0]"),
				GoStmt.GoRaw("}"),
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("haxe__io__GoIoHelpers_inputReadAll"), [GoExpr.GoIdent("self"), GoExpr.GoIdent("resolved")]))
			]),
			GoDecl.GoFuncDecl("haxe__io__input_readFullBytes", null, [
				{
					name: "self",
					typeName: "haxe__io__Input"
				},
				{name: "s", typeName: "*haxe__io__Bytes"},
				{name: "pos", typeName: "int"},
				{name: "len", typeName: "int"}
			], [], [
				GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("haxe__io__GoIoHelpers_inputReadFullBytes"), [
					GoExpr.GoIdent("self"),
					GoExpr.GoIdent("s"),
					GoExpr.GoIdent("pos"),
					GoExpr.GoIdent("len")
				]))
			]),
			GoDecl.GoFuncDecl("haxe__io__input_read", null, [
				{
					name: "self",
					typeName: "haxe__io__Input"
				},
				{name: "nbytes", typeName: "int"}
			], ["*haxe__io__Bytes"], [
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("haxe__io__GoIoHelpers_inputRead"), [GoExpr.GoIdent("self"), GoExpr.GoIdent("nbytes")]))
			]),
			GoDecl.GoFuncDecl("haxe__io__input_readUntil", null, [
				{
					name: "self",
					typeName: "haxe__io__Input"
				},
				{name: "end", typeName: "int"}
			], ["*string"], [
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("haxe__io__GoIoHelpers_inputReadUntil"), [GoExpr.GoIdent("self"), GoExpr.GoIdent("end")]))
			]),
			GoDecl.GoFuncDecl("haxe__io__input_readLine", null, [
				{
					name: "self",
					typeName: "haxe__io__Input"
				}
			], ["*string"], [
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("haxe__io__GoIoHelpers_inputReadLine"), [GoExpr.GoIdent("self")]))
			]),
			GoDecl.GoFuncDecl("haxe__io__input_readFloat", null, [
				{
					name: "self",
					typeName: "haxe__io__Input"
				}
			], ["float64"], [
				GoStmt.GoRaw("bits := uint32(self.readInt32())"),
				GoStmt.GoRaw("return float64(math.Float32frombits(bits))")
			]),
			GoDecl.GoFuncDecl("haxe__io__input_readDouble", null, [
				{
					name: "self",
					typeName: "haxe__io__Input"
				}
			], ["float64"], [
				GoStmt.GoRaw("i1 := self.readInt32()"),
				GoStmt.GoRaw("i2 := self.readInt32()"),
				GoStmt.GoRaw("if self.get_bigEndian() {"),
				GoStmt.GoRaw("\treturn math.Float64frombits((uint64(uint32(i1)) << 32) | uint64(uint32(i2)))"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("return math.Float64frombits((uint64(uint32(i2)) << 32) | uint64(uint32(i1)))")
			]),
			GoDecl.GoFuncDecl("haxe__io__input_readInt8", null, [
				{
					name: "self",
					typeName: "haxe__io__Input"
				}
			], ["int"], [
				GoStmt.GoRaw("n := self.readByte()"),
				GoStmt.GoRaw("if n >= 128 {"),
				GoStmt.GoRaw("\treturn n - 256"),
				GoStmt.GoRaw("}"),
				GoStmt.GoReturn(GoExpr.GoIdent("n"))
			]),
			GoDecl.GoFuncDecl("haxe__io__input_readInt16", null, [
				{
					name: "self",
					typeName: "haxe__io__Input"
				}
			], ["int"], [
				GoStmt.GoRaw("ch1 := self.readByte()"),
				GoStmt.GoRaw("ch2 := self.readByte()"),
				GoStmt.GoRaw("n := 0"),
				GoStmt.GoRaw("if self.get_bigEndian() {"),
				GoStmt.GoRaw("\tn = ch2 | (ch1 << 8)"),
				GoStmt.GoRaw("} else {"),
				GoStmt.GoRaw("\tn = ch1 | (ch2 << 8)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if (n & 0x8000) != 0 {"),
				GoStmt.GoRaw("\treturn n - 0x10000"),
				GoStmt.GoRaw("}"),
				GoStmt.GoReturn(GoExpr.GoIdent("n"))
			]),
			GoDecl.GoFuncDecl("haxe__io__input_readUInt16", null, [
				{
					name: "self",
					typeName: "haxe__io__Input"
				}
			], ["int"], [
				GoStmt.GoRaw("ch1 := self.readByte()"),
				GoStmt.GoRaw("ch2 := self.readByte()"),
				GoStmt.GoRaw("if self.get_bigEndian() {"),
				GoStmt.GoRaw("\treturn ch2 | (ch1 << 8)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("return ch1 | (ch2 << 8)")
			]),
			GoDecl.GoFuncDecl("haxe__io__input_readInt24", null, [
				{
					name: "self",
					typeName: "haxe__io__Input"
				}
			], ["int"], [
				GoStmt.GoRaw("ch1 := self.readByte()"),
				GoStmt.GoRaw("ch2 := self.readByte()"),
				GoStmt.GoRaw("ch3 := self.readByte()"),
				GoStmt.GoRaw("n := 0"),
				GoStmt.GoRaw("if self.get_bigEndian() {"),
				GoStmt.GoRaw("\tn = ch3 | (ch2 << 8) | (ch1 << 16)"),
				GoStmt.GoRaw("} else {"),
				GoStmt.GoRaw("\tn = ch1 | (ch2 << 8) | (ch3 << 16)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if (n & 0x800000) != 0 {"),
				GoStmt.GoRaw("\treturn n - 0x1000000"),
				GoStmt.GoRaw("}"),
				GoStmt.GoReturn(GoExpr.GoIdent("n"))
			]),
			GoDecl.GoFuncDecl("haxe__io__input_readUInt24", null, [
				{
					name: "self",
					typeName: "haxe__io__Input"
				}
			], ["int"], [
				GoStmt.GoRaw("ch1 := self.readByte()"),
				GoStmt.GoRaw("ch2 := self.readByte()"),
				GoStmt.GoRaw("ch3 := self.readByte()"),
				GoStmt.GoRaw("if self.get_bigEndian() {"),
				GoStmt.GoRaw("\treturn ch3 | (ch2 << 8) | (ch1 << 16)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("return ch1 | (ch2 << 8) | (ch3 << 16)")
			]),
			GoDecl.GoFuncDecl("haxe__io__input_readInt32", null, [
				{
					name: "self",
					typeName: "haxe__io__Input"
				}
			], ["int"], [
				GoStmt.GoRaw("ch1 := self.readByte()"),
				GoStmt.GoRaw("ch2 := self.readByte()"),
				GoStmt.GoRaw("ch3 := self.readByte()"),
				GoStmt.GoRaw("ch4 := self.readByte()"),
				GoStmt.GoRaw("if self.get_bigEndian() {"),
				GoStmt.GoRaw("\treturn ch4 | (ch3 << 8) | (ch2 << 16) | (ch1 << 24)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("return ch1 | (ch2 << 8) | (ch3 << 16) | (ch4 << 24)")
			]),
			GoDecl.GoFuncDecl("haxe__io__input_readString", null, [
				{
					name: "self",
					typeName: "haxe__io__Input"
				},
				{name: "len", typeName: "int"},
				{name: "encoding", typeName: "...*haxe__io__Encoding"}
			], ["*string"], [
				GoStmt.GoRaw("b := haxe__io__Bytes_alloc(len)"),
				GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("haxe__io__input_readFullBytes"),
					[
						GoExpr.GoIdent("self"),
						GoExpr.GoIdent("b"),
						GoExpr.GoIntLiteral(0),
						GoExpr.GoIdent("len")
					])),
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("b"), "getString"),
					[GoExpr.GoIntLiteral(0), GoExpr.GoIdent("len"), GoExpr.GoRaw("encoding...")]))
			]),
			GoDecl.GoFuncDecl("haxe__io__output_write", null, [
				{
					name: "self",
					typeName: "haxe__io__Output"
				},
				{name: "s", typeName: "*haxe__io__Bytes"}
			], [], [
				GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("haxe__io__GoIoHelpers_outputWrite"), [GoExpr.GoIdent("self"), GoExpr.GoIdent("s")]))
			]),
			GoDecl.GoFuncDecl("haxe__io__output_writeFullBytes", null, [
				{
					name: "self",
					typeName: "haxe__io__Output"
				},
				{name: "s", typeName: "*haxe__io__Bytes"},
				{name: "pos", typeName: "int"},
				{name: "len", typeName: "int"}
			], [], [
				GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("haxe__io__GoIoHelpers_outputWriteFullBytes"), [
					GoExpr.GoIdent("self"),
					GoExpr.GoIdent("s"),
					GoExpr.GoIdent("pos"),
					GoExpr.GoIdent("len")
				]))
			]),
			GoDecl.GoFuncDecl("haxe__io__output_writeFloat", null, [
				{
					name: "self",
					typeName: "haxe__io__Output"
				},
				{name: "x", typeName: "float64"}
			], [],
				[GoStmt.GoRaw("self.writeInt32(int(math.Float32bits(float32(x))))")]),
			GoDecl.GoFuncDecl("haxe__io__output_writeDouble", null, [
				{
					name: "self",
					typeName: "haxe__io__Output"
				},
				{name: "x", typeName: "float64"}
			], [], [
				GoStmt.GoRaw("bits := math.Float64bits(x)"),
				GoStmt.GoRaw("low := int(uint32(bits))"),
				GoStmt.GoRaw("high := int(uint32(bits >> 32))"),
				GoStmt.GoRaw("if self.get_bigEndian() {"),
				GoStmt.GoRaw("\tself.writeInt32(high)"),
				GoStmt.GoRaw("\tself.writeInt32(low)"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("self.writeInt32(low)"),
				GoStmt.GoRaw("self.writeInt32(high)")
			]),
			GoDecl.GoFuncDecl("haxe__io__output_writeInt8", null, [
				{
					name: "self",
					typeName: "haxe__io__Output"
				},
				{name: "x", typeName: "int"}
			], [], [
				GoStmt.GoRaw("if x < -0x80 || x >= 0x80 {"),
				GoStmt.GoRaw("\thxrt.Throw(haxe__io__Error_Overflow)"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("self.writeByte(x & 0xFF)")
			]),
			GoDecl.GoFuncDecl("haxe__io__output_writeInt16", null, [
				{
					name: "self",
					typeName: "haxe__io__Output"
				},
				{name: "x", typeName: "int"}
			], [], [
				GoStmt.GoRaw("if x < -0x8000 || x >= 0x8000 {"),
				GoStmt.GoRaw("\thxrt.Throw(haxe__io__Error_Overflow)"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("self.writeUInt16(x & 0xFFFF)")
			]),
			GoDecl.GoFuncDecl("haxe__io__output_writeUInt16", null, [
				{
					name: "self",
					typeName: "haxe__io__Output"
				},
				{name: "x", typeName: "int"}
			], [], [
				GoStmt.GoRaw("if x < 0 || x >= 0x10000 {"),
				GoStmt.GoRaw("\thxrt.Throw(haxe__io__Error_Overflow)"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if self.get_bigEndian() {"),
				GoStmt.GoRaw("\tself.writeByte(x >> 8)"),
				GoStmt.GoRaw("\tself.writeByte(x & 0xFF)"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("self.writeByte(x & 0xFF)"),
				GoStmt.GoRaw("self.writeByte(x >> 8)")
			]),
			GoDecl.GoFuncDecl("haxe__io__output_writeInt24", null, [
				{
					name: "self",
					typeName: "haxe__io__Output"
				},
				{name: "x", typeName: "int"}
			], [], [
				GoStmt.GoRaw("if x < -0x800000 || x >= 0x800000 {"),
				GoStmt.GoRaw("\thxrt.Throw(haxe__io__Error_Overflow)"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("self.writeUInt24(x & 0xFFFFFF)")
			]),
			GoDecl.GoFuncDecl("haxe__io__output_writeUInt24", null, [
				{
					name: "self",
					typeName: "haxe__io__Output"
				},
				{name: "x", typeName: "int"}
			], [], [
				GoStmt.GoRaw("if x < 0 || x >= 0x1000000 {"),
				GoStmt.GoRaw("\thxrt.Throw(haxe__io__Error_Overflow)"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if self.get_bigEndian() {"),
				GoStmt.GoRaw("\tself.writeByte(x >> 16)"),
				GoStmt.GoRaw("\tself.writeByte((x >> 8) & 0xFF)"),
				GoStmt.GoRaw("\tself.writeByte(x & 0xFF)"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("self.writeByte(x & 0xFF)"),
				GoStmt.GoRaw("self.writeByte((x >> 8) & 0xFF)"),
				GoStmt.GoRaw("self.writeByte(x >> 16)")
			]),
			GoDecl.GoFuncDecl("haxe__io__output_writeInt32", null, [
				{
					name: "self",
					typeName: "haxe__io__Output"
				},
				{name: "x", typeName: "int"}
			], [], [
				GoStmt.GoRaw("if self.get_bigEndian() {"),
				GoStmt.GoRaw("\tself.writeByte(int(uint(x) >> 24))"),
				GoStmt.GoRaw("\tself.writeByte((x >> 16) & 0xFF)"),
				GoStmt.GoRaw("\tself.writeByte((x >> 8) & 0xFF)"),
				GoStmt.GoRaw("\tself.writeByte(x & 0xFF)"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("self.writeByte(x & 0xFF)"),
				GoStmt.GoRaw("self.writeByte((x >> 8) & 0xFF)"),
				GoStmt.GoRaw("self.writeByte((x >> 16) & 0xFF)"),
				GoStmt.GoRaw("self.writeByte(int(uint(x) >> 24))")
			]),
			GoDecl.GoFuncDecl("haxe__io__output_writeInput", null, [
				{
					name: "self",
					typeName: "haxe__io__Output"
				},
				{name: "i", typeName: "haxe__io__Input"},
				{name: "bufsize", typeName: "...int"}
			], [], [
				GoStmt.GoRaw("resolved := 4096"),
				GoStmt.GoRaw("if len(bufsize) > 0 {"),
				GoStmt.GoRaw("\tresolved = bufsize[0]"),
				GoStmt.GoRaw("}"),
				GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("haxe__io__GoIoHelpers_outputWriteInput"),
					[GoExpr.GoIdent("self"), GoExpr.GoIdent("i"), GoExpr.GoIdent("resolved")]))
			]),
			GoDecl.GoFuncDecl("haxe__io__output_writeString", null, [
				{
					name: "self",
					typeName: "haxe__io__Output"
				},
				{name: "s", typeName: "*string"},
				{name: "encoding", typeName: "...*haxe__io__Encoding"}
			], [], [
				GoStmt.GoRaw("var resolved *haxe__io__Encoding"),
				GoStmt.GoRaw("if len(encoding) > 0 {"),
				GoStmt.GoRaw("\tresolved = encoding[0]"),
				GoStmt.GoRaw("}"),
				GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("haxe__io__GoIoHelpers_outputWriteString"),
					[GoExpr.GoIdent("self"), GoExpr.GoIdent("s"), GoExpr.GoIdent("resolved")]))
			]),
			GoDecl.GoFuncDecl("New_haxe__io__BytesInput", null, [
				{
					name: "b",
					typeName: "*haxe__io__Bytes"
				},
				{name: "opts", typeName: "...int"}
			], ["*haxe__io__BytesInput"], [
				GoStmt.GoRaw("if b == nil {"),
				GoStmt.GoRaw("\thxrt.Throw(haxe__io__Error_OutsideBounds)"),
				GoStmt.GoRaw("\treturn &haxe__io__BytesInput{}"),
				GoStmt.GoRaw("}"),
				GoStmt.GoVarDecl("start", null, GoExpr.GoIntLiteral(0), true),
				GoStmt.GoIf(GoExpr.GoBinary(">", GoExpr.GoCall(GoExpr.GoIdent("len"), [GoExpr.GoIdent("opts")]), GoExpr.GoIntLiteral(0)), [
					GoStmt.GoAssign(GoExpr.GoIdent("start"), GoExpr.GoIndex(GoExpr.GoIdent("opts"), GoExpr.GoIntLiteral(0)))
				],
					null),
				GoStmt.GoVarDecl("sliceLen", null, GoExpr.GoBinary("-", GoExpr.GoSelector(GoExpr.GoIdent("b"), "length"), GoExpr.GoIdent("start")), true),
				GoStmt.GoIf(GoExpr.GoBinary(">", GoExpr.GoCall(GoExpr.GoIdent("len"), [GoExpr.GoIdent("opts")]), GoExpr.GoIntLiteral(1)), [
					GoStmt.GoAssign(GoExpr.GoIdent("sliceLen"), GoExpr.GoIndex(GoExpr.GoIdent("opts"), GoExpr.GoIntLiteral(1)))
				],
					null),
				GoStmt.GoRaw("if start < 0 || sliceLen < 0 || start+sliceLen > b.length {"),
				GoStmt.GoRaw("\thxrt.Throw(haxe__io__Error_OutsideBounds)"),
				GoStmt.GoRaw("\treturn &haxe__io__BytesInput{}"),
				GoStmt.GoRaw("}"),
				GoStmt.GoReturn(GoExpr.GoRaw("&haxe__io__BytesInput{b: b.b, pos: start, len: sliceLen, totlen: sliceLen}"))
			]),
			GoDecl.GoFuncDecl("get_position", {
				name: "self",
				typeName: "*haxe__io__BytesInput"
			}, [], ["int"],
				[GoStmt.GoReturn(GoExpr.GoSelector(GoExpr.GoIdent("self"), "pos"))]),
			GoDecl.GoFuncDecl("set_position", {
				name: "self",
				typeName: "*haxe__io__BytesInput"
			}, [{name: "p", typeName: "int"}], ["int"], [
				GoStmt.GoIf(GoExpr.GoBinary("<", GoExpr.GoIdent("p"), GoExpr.GoIntLiteral(0)), [GoStmt.GoAssign(GoExpr.GoIdent("p"), GoExpr.GoIntLiteral(0))],
					[
						GoStmt.GoIf(GoExpr.GoBinary(">", GoExpr.GoIdent("p"), GoExpr.GoSelector(GoExpr.GoIdent("self"), "totlen")), [
							GoStmt.GoAssign(GoExpr.GoIdent("p"), GoExpr.GoSelector(GoExpr.GoIdent("self"), "totlen"))
						],
							null)
					]),
				GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("self"), "len"),
					GoExpr.GoBinary("-", GoExpr.GoSelector(GoExpr.GoIdent("self"), "totlen"), GoExpr.GoIdent("p"))),
				GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("self"), "pos"), GoExpr.GoIdent("p")),
				GoStmt.GoReturn(GoExpr.GoIdent("p"))
			]),
			GoDecl.GoFuncDecl("get_length", {
				name: "self",
				typeName: "*haxe__io__BytesInput"
			}, [], ["int"],
				[GoStmt.GoReturn(GoExpr.GoSelector(GoExpr.GoIdent("self"), "totlen"))]),
			GoDecl.GoFuncDecl("readByte", {
				name: "self",
				typeName: "*haxe__io__BytesInput"
			}, [], ["int"], [
				GoStmt.GoRaw("if self == nil || self.len == 0 {"),
				GoStmt.GoRaw("\thxrt.Throw(&haxe__io__Eof{})"),
				GoStmt.GoRaw("\treturn 0"),
				GoStmt.GoRaw("}"),
				GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("self"), "len"),
					GoExpr.GoBinary("-", GoExpr.GoSelector(GoExpr.GoIdent("self"), "len"), GoExpr.GoIntLiteral(1))),
				GoStmt.GoVarDecl("value", null,
					GoExpr.GoIndex(GoExpr.GoSelector(GoExpr.GoIdent("self"), "b"), GoExpr.GoSelector(GoExpr.GoIdent("self"), "pos")), true),
				GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("self"), "pos"),
					GoExpr.GoBinary("+", GoExpr.GoSelector(GoExpr.GoIdent("self"), "pos"), GoExpr.GoIntLiteral(1))),
				GoStmt.GoReturn(GoExpr.GoIdent("value"))
			]),
			GoDecl.GoFuncDecl("readBytes", {
				name: "self",
				typeName: "*haxe__io__BytesInput"
			}, [
				{name: "buf", typeName: "*haxe__io__Bytes"},
				{name: "pos", typeName: "int"},
				{name: "len", typeName: "int"}
			], ["int"], [
				GoStmt.GoRaw("if buf == nil || pos < 0 || len < 0 || pos+len > buf.length {"),
				GoStmt.GoRaw("\thxrt.Throw(haxe__io__Error_OutsideBounds)"),
				GoStmt.GoRaw("\treturn 0"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if len > 0 && (self == nil || self.len == 0) {"),
				GoStmt.GoRaw("\thxrt.Throw(&haxe__io__Eof{})"),
				GoStmt.GoRaw("\treturn 0"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if self == nil {"),
				GoStmt.GoRaw("\treturn 0"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if self.len < len {"),
				GoStmt.GoRaw("\tlen = self.len"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("for i := 0; i < len; i++ {"),
				GoStmt.GoRaw("\tbuf.b[pos+i] = self.b[self.pos+i]"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("self.pos += len"),
				GoStmt.GoRaw("self.len -= len"),
				GoStmt.GoReturn(GoExpr.GoIdent("len"))
			]),
			GoDecl.GoFuncDecl("get_bigEndian", {
				name: "self",
				typeName: "*haxe__io__BytesInput"
			}, [], ["bool"], [
				GoStmt.GoIf(GoExpr.GoBinary("==", GoExpr.GoIdent("self"), GoExpr.GoNil), [GoStmt.GoReturn(GoExpr.GoBoolLiteral(false))], null),
				GoStmt.GoReturn(GoExpr.GoSelector(GoExpr.GoIdent("self"), "bigEndian"))
			]),
			GoDecl.GoFuncDecl("set_bigEndian", {
				name: "self",
				typeName: "*haxe__io__BytesInput"
			}, [{name: "e", typeName: "bool"}], ["bool"], [
				GoStmt.GoIf(GoExpr.GoBinary("!=", GoExpr.GoIdent("self"), GoExpr.GoNil), [
					GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("self"), "bigEndian"), GoExpr.GoIdent("e"))
				], null),
				GoStmt.GoReturn(GoExpr.GoIdent("e"))
			]),
			GoDecl.GoFuncDecl("close", {
				name: "self",
				typeName: "*haxe__io__BytesInput"
			}, [], [],
				[GoStmt.GoAssign(GoExpr.GoIdent("_"), GoExpr.GoIdent("self"))]),
			GoDecl.GoFuncDecl("readAll", {
				name: "self",
				typeName: "*haxe__io__BytesInput"
			},
				[{name: "bufsize", typeName: "...int"}], ["*haxe__io__Bytes"], [GoStmt.GoRaw("return haxe__io__input_readAll(self, bufsize...)")]),
			GoDecl.GoFuncDecl("readFullBytes", {
				name: "self",
				typeName: "*haxe__io__BytesInput"
			}, [
				{name: "s", typeName: "*haxe__io__Bytes"},
				{name: "pos", typeName: "int"},
				{name: "len", typeName: "int"}
			], [],
				[GoStmt.GoRaw("haxe__io__input_readFullBytes(self, s, pos, len)")]),
			GoDecl.GoFuncDecl("read", {
				name: "self",
				typeName: "*haxe__io__BytesInput"
			},
				[{name: "nbytes", typeName: "int"}], ["*haxe__io__Bytes"], [GoStmt.GoRaw("return haxe__io__input_read(self, nbytes)")]),
			GoDecl.GoFuncDecl("readUntil", {
				name: "self",
				typeName: "*haxe__io__BytesInput"
			},
				[{name: "end", typeName: "int"}], ["*string"], [GoStmt.GoRaw("return haxe__io__input_readUntil(self, end)")]),
			GoDecl.GoFuncDecl("readLine", {
				name: "self",
				typeName: "*haxe__io__BytesInput"
			}, [], ["*string"],
				[GoStmt.GoRaw("return haxe__io__input_readLine(self)")]),
			GoDecl.GoFuncDecl("readFloat", {
				name: "self",
				typeName: "*haxe__io__BytesInput"
			}, [], ["float64"],
				[GoStmt.GoRaw("return haxe__io__input_readFloat(self)")]),
			GoDecl.GoFuncDecl("readDouble", {
				name: "self",
				typeName: "*haxe__io__BytesInput"
			}, [], ["float64"],
				[GoStmt.GoRaw("return haxe__io__input_readDouble(self)")]),
			GoDecl.GoFuncDecl("readInt8", {
				name: "self",
				typeName: "*haxe__io__BytesInput"
			}, [], ["int"],
				[GoStmt.GoRaw("return haxe__io__input_readInt8(self)")]),
			GoDecl.GoFuncDecl("readInt16", {
				name: "self",
				typeName: "*haxe__io__BytesInput"
			}, [], ["int"],
				[GoStmt.GoRaw("return haxe__io__input_readInt16(self)")]),
			GoDecl.GoFuncDecl("readUInt16", {
				name: "self",
				typeName: "*haxe__io__BytesInput"
			}, [], ["int"],
				[GoStmt.GoRaw("return haxe__io__input_readUInt16(self)")]),
			GoDecl.GoFuncDecl("readInt24", {
				name: "self",
				typeName: "*haxe__io__BytesInput"
			}, [], ["int"],
				[GoStmt.GoRaw("return haxe__io__input_readInt24(self)")]),
			GoDecl.GoFuncDecl("readUInt24", {
				name: "self",
				typeName: "*haxe__io__BytesInput"
			}, [], ["int"],
				[GoStmt.GoRaw("return haxe__io__input_readUInt24(self)")]),
			GoDecl.GoFuncDecl("readInt32", {
				name: "self",
				typeName: "*haxe__io__BytesInput"
			}, [], ["int"],
				[GoStmt.GoRaw("return haxe__io__input_readInt32(self)")]),
			GoDecl.GoFuncDecl("readString", {
				name: "self",
				typeName: "*haxe__io__BytesInput"
			}, [
				{name: "len", typeName: "int"},
				{name: "encoding", typeName: "...*haxe__io__Encoding"}
			], ["*string"],
				[GoStmt.GoRaw("return haxe__io__input_readString(self, len, encoding...)")]),
			GoDecl.GoFuncDecl("get_bigEndian", {
				name: "self",
				typeName: "*haxe__io__StringInput"
			}, [], ["bool"],
				[GoStmt.GoRaw("return self.haxe__io__BytesInput.get_bigEndian()")]),
			GoDecl.GoFuncDecl("set_bigEndian", {
				name: "self",
				typeName: "*haxe__io__StringInput"
			},
				[{name: "e", typeName: "bool"}], ["bool"], [GoStmt.GoRaw("return self.haxe__io__BytesInput.set_bigEndian(e)")]),
			GoDecl.GoFuncDecl("close", {
				name: "self",
				typeName: "*haxe__io__StringInput"
			}, [], [], [GoStmt.GoRaw("self.haxe__io__BytesInput.close()")]),
			GoDecl.GoFuncDecl("readByte", {
				name: "self",
				typeName: "*haxe__io__StringInput"
			}, [], ["int"],
				[GoStmt.GoRaw("return self.haxe__io__BytesInput.readByte()")]),
			GoDecl.GoFuncDecl("readBytes", {
				name: "self",
				typeName: "*haxe__io__StringInput"
			}, [
				{name: "buf", typeName: "*haxe__io__Bytes"},
				{name: "pos", typeName: "int"},
				{name: "len", typeName: "int"}
			], ["int"],
				[GoStmt.GoRaw("return self.haxe__io__BytesInput.readBytes(buf, pos, len)")]),
			GoDecl.GoFuncDecl("readAll", {
				name: "self",
				typeName: "*haxe__io__StringInput"
			},
				[{name: "bufsize", typeName: "...int"}], ["*haxe__io__Bytes"], [GoStmt.GoRaw("return haxe__io__input_readAll(self, bufsize...)")]),
			GoDecl.GoFuncDecl("readFullBytes", {
				name: "self",
				typeName: "*haxe__io__StringInput"
			}, [
				{name: "s", typeName: "*haxe__io__Bytes"},
				{name: "pos", typeName: "int"},
				{name: "len", typeName: "int"}
			], [],
				[GoStmt.GoRaw("haxe__io__input_readFullBytes(self, s, pos, len)")]),
			GoDecl.GoFuncDecl("read", {
				name: "self",
				typeName: "*haxe__io__StringInput"
			},
				[{name: "nbytes", typeName: "int"}], ["*haxe__io__Bytes"], [GoStmt.GoRaw("return haxe__io__input_read(self, nbytes)")]),
			GoDecl.GoFuncDecl("readUntil", {
				name: "self",
				typeName: "*haxe__io__StringInput"
			},
				[{name: "end", typeName: "int"}], ["*string"], [GoStmt.GoRaw("return haxe__io__input_readUntil(self, end)")]),
			GoDecl.GoFuncDecl("readLine", {
				name: "self",
				typeName: "*haxe__io__StringInput"
			}, [], ["*string"],
				[GoStmt.GoRaw("return haxe__io__input_readLine(self)")]),
			GoDecl.GoFuncDecl("readFloat", {
				name: "self",
				typeName: "*haxe__io__StringInput"
			}, [], ["float64"],
				[GoStmt.GoRaw("return haxe__io__input_readFloat(self)")]),
			GoDecl.GoFuncDecl("readDouble", {
				name: "self",
				typeName: "*haxe__io__StringInput"
			}, [], ["float64"],
				[GoStmt.GoRaw("return haxe__io__input_readDouble(self)")]),
			GoDecl.GoFuncDecl("readInt8", {
				name: "self",
				typeName: "*haxe__io__StringInput"
			}, [], ["int"],
				[GoStmt.GoRaw("return haxe__io__input_readInt8(self)")]),
			GoDecl.GoFuncDecl("readInt16", {
				name: "self",
				typeName: "*haxe__io__StringInput"
			}, [], ["int"],
				[GoStmt.GoRaw("return haxe__io__input_readInt16(self)")]),
			GoDecl.GoFuncDecl("readUInt16", {
				name: "self",
				typeName: "*haxe__io__StringInput"
			}, [], ["int"],
				[GoStmt.GoRaw("return haxe__io__input_readUInt16(self)")]),
			GoDecl.GoFuncDecl("readInt24", {
				name: "self",
				typeName: "*haxe__io__StringInput"
			}, [], ["int"],
				[GoStmt.GoRaw("return haxe__io__input_readInt24(self)")]),
			GoDecl.GoFuncDecl("readUInt24", {
				name: "self",
				typeName: "*haxe__io__StringInput"
			}, [], ["int"],
				[GoStmt.GoRaw("return haxe__io__input_readUInt24(self)")]),
			GoDecl.GoFuncDecl("readInt32", {
				name: "self",
				typeName: "*haxe__io__StringInput"
			}, [], ["int"],
				[GoStmt.GoRaw("return haxe__io__input_readInt32(self)")]),
			GoDecl.GoFuncDecl("readString", {
				name: "self",
				typeName: "*haxe__io__StringInput"
			}, [
				{name: "len", typeName: "int"},
				{name: "encoding", typeName: "...*haxe__io__Encoding"}
			], ["*string"],
				[GoStmt.GoRaw("return haxe__io__input_readString(self, len, encoding...)")]),
			GoDecl.GoFuncDecl("get_bigEndian", {
				name: "self",
				typeName: "*haxe__io__BufferInput"
			}, [], ["bool"], [
				GoStmt.GoIf(GoExpr.GoBinary("==", GoExpr.GoIdent("self"), GoExpr.GoNil), [GoStmt.GoReturn(GoExpr.GoBoolLiteral(false))], null),
				GoStmt.GoReturn(GoExpr.GoSelector(GoExpr.GoIdent("self"), "bigEndian"))
			]),
			GoDecl.GoFuncDecl("set_bigEndian", {
				name: "self",
				typeName: "*haxe__io__BufferInput"
			}, [{name: "e", typeName: "bool"}], ["bool"], [
				GoStmt.GoIf(GoExpr.GoBinary("!=", GoExpr.GoIdent("self"), GoExpr.GoNil), [
					GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("self"), "bigEndian"), GoExpr.GoIdent("e"))
				], null),
				GoStmt.GoReturn(GoExpr.GoIdent("e"))
			]),
			GoDecl.GoFuncDecl("close", {
				name: "self",
				typeName: "*haxe__io__BufferInput"
			}, [], [],
				[GoStmt.GoAssign(GoExpr.GoIdent("_"), GoExpr.GoIdent("self"))]),
			GoDecl.GoFuncDecl("refill", {
				name: "self",
				typeName: "*haxe__io__BufferInput"
			}, [], [], [
				GoStmt.GoRaw("if self == nil || self.buf == nil || self.i == nil {"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if self.pos > 0 {"),
				GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("self.buf"), "blit"),
					[
						GoExpr.GoIntLiteral(0),
						GoExpr.GoIdent("self.buf"),
						GoExpr.GoSelector(GoExpr.GoIdent("self"), "pos"),
						GoExpr.GoSelector(GoExpr.GoIdent("self"), "available")
					])),
				GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("self"), "pos"), GoExpr.GoIntLiteral(0)),
				GoStmt.GoRaw("}"),
				GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("self"), "available"),
					GoExpr.GoBinary("+", GoExpr.GoSelector(GoExpr.GoIdent("self"), "available"),
						GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("self"), "i.readBytes"), [
							GoExpr.GoSelector(GoExpr.GoIdent("self"), "buf"),
							GoExpr.GoSelector(GoExpr.GoIdent("self"), "available"),
							GoExpr.GoBinary("-", GoExpr.GoSelector(GoExpr.GoSelector(GoExpr.GoIdent("self"), "buf"), "length"),
								GoExpr.GoSelector(GoExpr.GoIdent("self"), "available"))
						])))
			]),
			GoDecl.GoFuncDecl("readByte", {
				name: "self",
				typeName: "*haxe__io__BufferInput"
			}, [], ["int"], [
				GoStmt.GoRaw("if self.available == 0 {"),
				GoStmt.GoRaw("\tself.refill()"),
				GoStmt.GoRaw("}"),
				GoStmt.GoVarDecl("c", null,
					GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoSelector(GoExpr.GoIdent("self"), "buf"), "get"),
						[GoExpr.GoSelector(GoExpr.GoIdent("self"), "pos")]),
					true),
				GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("self"), "pos"),
					GoExpr.GoBinary("+", GoExpr.GoSelector(GoExpr.GoIdent("self"), "pos"), GoExpr.GoIntLiteral(1))),
				GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("self"), "available"),
					GoExpr.GoBinary("-", GoExpr.GoSelector(GoExpr.GoIdent("self"), "available"), GoExpr.GoIntLiteral(1))),
				GoStmt.GoReturn(GoExpr.GoIdent("c"))
			]),
			GoDecl.GoFuncDecl("readBytes", {
				name: "self",
				typeName: "*haxe__io__BufferInput"
			}, [
				{name: "buf", typeName: "*haxe__io__Bytes"},
				{name: "pos", typeName: "int"},
				{name: "len", typeName: "int"}
			], ["int"], [
				GoStmt.GoRaw("if self.available == 0 {"),
				GoStmt.GoRaw("\tself.refill()"),
				GoStmt.GoRaw("}"),
				GoStmt.GoVarDecl("size", null, GoExpr.GoIdent("len"), true),
				GoStmt.GoIf(GoExpr.GoBinary(">", GoExpr.GoIdent("len"), GoExpr.GoSelector(GoExpr.GoIdent("self"), "available")), [
					GoStmt.GoAssign(GoExpr.GoIdent("size"), GoExpr.GoSelector(GoExpr.GoIdent("self"), "available"))
				], null),
				GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("buf"), "blit"),
					[
						GoExpr.GoIdent("pos"),
						GoExpr.GoSelector(GoExpr.GoIdent("self"), "buf"),
						GoExpr.GoSelector(GoExpr.GoIdent("self"), "pos"),
						GoExpr.GoIdent("size")
					])),
				GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("self"), "pos"),
					GoExpr.GoBinary("+", GoExpr.GoSelector(GoExpr.GoIdent("self"), "pos"), GoExpr.GoIdent("size"))),
				GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("self"), "available"),
					GoExpr.GoBinary("-", GoExpr.GoSelector(GoExpr.GoIdent("self"), "available"), GoExpr.GoIdent("size"))),
				GoStmt.GoReturn(GoExpr.GoIdent("size"))
			]),
			GoDecl.GoFuncDecl("readAll", {
				name: "self",
				typeName: "*haxe__io__BufferInput"
			},
				[{name: "bufsize", typeName: "...int"}], ["*haxe__io__Bytes"], [GoStmt.GoRaw("return haxe__io__input_readAll(self, bufsize...)")]),
			GoDecl.GoFuncDecl("readFullBytes", {
				name: "self",
				typeName: "*haxe__io__BufferInput"
			}, [
				{name: "s", typeName: "*haxe__io__Bytes"},
				{name: "pos", typeName: "int"},
				{name: "len", typeName: "int"}
			], [],
				[GoStmt.GoRaw("haxe__io__input_readFullBytes(self, s, pos, len)")]),
			GoDecl.GoFuncDecl("read", {
				name: "self",
				typeName: "*haxe__io__BufferInput"
			},
				[{name: "nbytes", typeName: "int"}], ["*haxe__io__Bytes"], [GoStmt.GoRaw("return haxe__io__input_read(self, nbytes)")]),
			GoDecl.GoFuncDecl("readUntil", {
				name: "self",
				typeName: "*haxe__io__BufferInput"
			},
				[{name: "end", typeName: "int"}], ["*string"], [GoStmt.GoRaw("return haxe__io__input_readUntil(self, end)")]),
			GoDecl.GoFuncDecl("readLine", {
				name: "self",
				typeName: "*haxe__io__BufferInput"
			}, [], ["*string"],
				[GoStmt.GoRaw("return haxe__io__input_readLine(self)")]),
			GoDecl.GoFuncDecl("readFloat", {
				name: "self",
				typeName: "*haxe__io__BufferInput"
			}, [], ["float64"],
				[GoStmt.GoRaw("return haxe__io__input_readFloat(self)")]),
			GoDecl.GoFuncDecl("readDouble", {
				name: "self",
				typeName: "*haxe__io__BufferInput"
			}, [], ["float64"],
				[GoStmt.GoRaw("return haxe__io__input_readDouble(self)")]),
			GoDecl.GoFuncDecl("readInt8", {
				name: "self",
				typeName: "*haxe__io__BufferInput"
			}, [], ["int"],
				[GoStmt.GoRaw("return haxe__io__input_readInt8(self)")]),
			GoDecl.GoFuncDecl("readInt16", {
				name: "self",
				typeName: "*haxe__io__BufferInput"
			}, [], ["int"],
				[GoStmt.GoRaw("return haxe__io__input_readInt16(self)")]),
			GoDecl.GoFuncDecl("readUInt16", {
				name: "self",
				typeName: "*haxe__io__BufferInput"
			}, [], ["int"],
				[GoStmt.GoRaw("return haxe__io__input_readUInt16(self)")]),
			GoDecl.GoFuncDecl("readInt24", {
				name: "self",
				typeName: "*haxe__io__BufferInput"
			}, [], ["int"],
				[GoStmt.GoRaw("return haxe__io__input_readInt24(self)")]),
			GoDecl.GoFuncDecl("readUInt24", {
				name: "self",
				typeName: "*haxe__io__BufferInput"
			}, [], ["int"],
				[GoStmt.GoRaw("return haxe__io__input_readUInt24(self)")]),
			GoDecl.GoFuncDecl("readInt32", {
				name: "self",
				typeName: "*haxe__io__BufferInput"
			}, [], ["int"],
				[GoStmt.GoRaw("return haxe__io__input_readInt32(self)")]),
			GoDecl.GoFuncDecl("readString", {
				name: "self",
				typeName: "*haxe__io__BufferInput"
			}, [
				{name: "len", typeName: "int"},
				{name: "encoding", typeName: "...*haxe__io__Encoding"}
			],
				["*string"], [GoStmt.GoRaw("return haxe__io__input_readString(self, len, encoding...)")]),
			GoDecl.GoFuncDecl("New_haxe__io__BytesOutput", null, [], ["*haxe__io__BytesOutput"], [
				GoStmt.GoReturn(GoExpr.GoRaw("&haxe__io__BytesOutput{b: &haxe__io__BytesBuffer{b: []int{}}}"))
			]),
			GoDecl.GoFuncDecl("get_length", {
				name: "self",
				typeName: "*haxe__io__BytesOutput"
			}, [], ["int"], [
				GoStmt.GoIf(GoExpr.GoRaw("self == nil || self.b == nil"), [GoStmt.GoReturn(GoExpr.GoIntLiteral(0))], null),
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoSelector(GoExpr.GoIdent("self"), "b"), "get_length"), []))
			]),
			GoDecl.GoFuncDecl("writeByte", {
				name: "self",
				typeName: "*haxe__io__BytesOutput"
			}, [{name: "c", typeName: "int"}], [], [
				GoStmt.GoIf(GoExpr.GoRaw("self == nil || self.b == nil"), [GoStmt.GoReturn(null)], null),
				GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoSelector(GoExpr.GoIdent("self"), "b"), "addByte"), [GoExpr.GoIdent("c")]))
			]),
			GoDecl.GoFuncDecl("writeBytes", {
				name: "self",
				typeName: "*haxe__io__BytesOutput"
			}, [
				{name: "buf", typeName: "*haxe__io__Bytes"},
				{name: "pos", typeName: "int"},
				{name: "len", typeName: "int"}
			], ["int"], [
				GoStmt.GoRaw("if buf == nil || pos < 0 || len < 0 || pos+len > buf.length {"),
				GoStmt.GoRaw("\thxrt.Throw(haxe__io__Error_OutsideBounds)"),
				GoStmt.GoRaw("\treturn 0"),
				GoStmt.GoRaw("}"),
				GoStmt.GoIf(GoExpr.GoRaw("self == nil || self.b == nil"), [GoStmt.GoReturn(GoExpr.GoIntLiteral(0))], null),
				GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoSelector(GoExpr.GoIdent("self"), "b"), "addBytes"),
					[GoExpr.GoIdent("buf"), GoExpr.GoIdent("pos"), GoExpr.GoIdent("len")])),
				GoStmt.GoReturn(GoExpr.GoIdent("len"))
			]),
			GoDecl.GoFuncDecl("get_bigEndian", {
				name: "self",
				typeName: "*haxe__io__BytesOutput"
			}, [], ["bool"], [
				GoStmt.GoIf(GoExpr.GoBinary("==", GoExpr.GoIdent("self"), GoExpr.GoNil), [GoStmt.GoReturn(GoExpr.GoBoolLiteral(false))], null),
				GoStmt.GoReturn(GoExpr.GoSelector(GoExpr.GoIdent("self"), "bigEndian"))
			]),
			GoDecl.GoFuncDecl("set_bigEndian", {
				name: "self",
				typeName: "*haxe__io__BytesOutput"
			}, [{name: "e", typeName: "bool"}], ["bool"], [
				GoStmt.GoIf(GoExpr.GoBinary("!=", GoExpr.GoIdent("self"), GoExpr.GoNil), [
					GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("self"), "bigEndian"), GoExpr.GoIdent("e"))
				], null),
				GoStmt.GoReturn(GoExpr.GoIdent("e"))
			]),
			GoDecl.GoFuncDecl("flush", {
				name: "self",
				typeName: "*haxe__io__BytesOutput"
			}, [], [],
				[GoStmt.GoAssign(GoExpr.GoIdent("_"), GoExpr.GoIdent("self"))]),
			GoDecl.GoFuncDecl("close", {
				name: "self",
				typeName: "*haxe__io__BytesOutput"
			}, [], [],
				[GoStmt.GoAssign(GoExpr.GoIdent("_"), GoExpr.GoIdent("self"))]),
			GoDecl.GoFuncDecl("write", {
				name: "self",
				typeName: "*haxe__io__BytesOutput"
			},
				[{name: "s", typeName: "*haxe__io__Bytes"}], [], [GoStmt.GoRaw("haxe__io__output_write(self, s)")]),
			GoDecl.GoFuncDecl("writeFullBytes", {
				name: "self",
				typeName: "*haxe__io__BytesOutput"
			}, [
				{name: "s", typeName: "*haxe__io__Bytes"},
				{name: "pos", typeName: "int"},
				{name: "len", typeName: "int"}
			], [],
				[GoStmt.GoRaw("haxe__io__output_writeFullBytes(self, s, pos, len)")]),
			GoDecl.GoFuncDecl("writeFloat", {
				name: "self",
				typeName: "*haxe__io__BytesOutput"
			},
				[{name: "x", typeName: "float64"}], [], [GoStmt.GoRaw("haxe__io__output_writeFloat(self, x)")]),
			GoDecl.GoFuncDecl("writeDouble", {
				name: "self",
				typeName: "*haxe__io__BytesOutput"
			},
				[{name: "x", typeName: "float64"}], [], [GoStmt.GoRaw("haxe__io__output_writeDouble(self, x)")]),
			GoDecl.GoFuncDecl("writeInt8", {
				name: "self",
				typeName: "*haxe__io__BytesOutput"
			}, [{name: "x", typeName: "int"}],
				[], [GoStmt.GoRaw("haxe__io__output_writeInt8(self, x)")]),
			GoDecl.GoFuncDecl("writeInt16", {
				name: "self",
				typeName: "*haxe__io__BytesOutput"
			}, [{name: "x", typeName: "int"}],
				[], [GoStmt.GoRaw("haxe__io__output_writeInt16(self, x)")]),
			GoDecl.GoFuncDecl("writeUInt16", {
				name: "self",
				typeName: "*haxe__io__BytesOutput"
			}, [{name: "x", typeName: "int"}],
				[], [GoStmt.GoRaw("haxe__io__output_writeUInt16(self, x)")]),
			GoDecl.GoFuncDecl("writeInt24", {
				name: "self",
				typeName: "*haxe__io__BytesOutput"
			}, [{name: "x", typeName: "int"}],
				[], [GoStmt.GoRaw("haxe__io__output_writeInt24(self, x)")]),
			GoDecl.GoFuncDecl("writeUInt24", {
				name: "self",
				typeName: "*haxe__io__BytesOutput"
			}, [{name: "x", typeName: "int"}],
				[], [GoStmt.GoRaw("haxe__io__output_writeUInt24(self, x)")]),
			GoDecl.GoFuncDecl("writeInt32", {
				name: "self",
				typeName: "*haxe__io__BytesOutput"
			}, [{name: "x", typeName: "int"}], [],
				[GoStmt.GoRaw("haxe__io__output_writeInt32(self, x)")]),
			GoDecl.GoFuncDecl("prepare", {
				name: "self",
				typeName: "*haxe__io__BytesOutput"
			}, [{name: "nbytes", typeName: "int"}], [], [
				GoStmt.GoAssign(GoExpr.GoIdent("_"), GoExpr.GoIdent("self")),
				GoStmt.GoAssign(GoExpr.GoIdent("_"), GoExpr.GoIdent("nbytes"))
			]),
			GoDecl.GoFuncDecl("writeInput", {
				name: "self",
				typeName: "*haxe__io__BytesOutput"
			},
				[{name: "i", typeName: "haxe__io__Input"}, {name: "bufsize", typeName: "...int"}], [],
				[GoStmt.GoRaw("haxe__io__output_writeInput(self, i, bufsize...)")]),
			GoDecl.GoFuncDecl("writeString", {
				name: "self",
				typeName: "*haxe__io__BytesOutput"
			}, [
				{name: "s", typeName: "*string"},
				{name: "encoding", typeName: "...*haxe__io__Encoding"}
			], [],
				[GoStmt.GoRaw("haxe__io__output_writeString(self, s, encoding...)")]),
			GoDecl.GoFuncDecl("getBytes", {
				name: "self",
				typeName: "*haxe__io__BytesOutput"
			}, [], ["*haxe__io__Bytes"], [
				GoStmt.GoIf(GoExpr.GoRaw("self == nil || self.b == nil"), [GoStmt.GoReturn(GoExpr.GoRaw("&haxe__io__Bytes{b: []int{}, length: 0}"))], null),
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoSelector(GoExpr.GoIdent("self"), "b"), "getBytes"), []))
			])
		];
		decls = trimUnusedIoDirectSurface(decls);
		if (!requiresIoHelperSurface) {
			decls = trimIoShimToCoreSurface(decls);
		}
		return decls;
	}

	function trimUnusedIoDirectSurface(decls:Array<GoDecl>):Array<GoDecl> {
		var out = new Array<GoDecl>();
		for (decl in decls) {
			switch (decl) {
				case GoDecl.GoStructDecl(name, _):
					if (!requiresIoStringInputSurface && name == "haxe__io__StringInput") {
						continue;
					}
					if (!requiresIoBufferInputSurface && name == "haxe__io__BufferInput") {
						continue;
					}
					out.push(decl);
				case GoDecl.GoFuncDecl(name, receiver, _, _, _):
					if (!requiresIoStringInputSurface && receiver == null && name == "New_haxe__io__StringInput") {
						continue;
					}
					if (!requiresIoBufferInputSurface && receiver == null && name == "New_haxe__io__BufferInput") {
						continue;
					}
					if (!requiresIoStringInputSurface && receiver != null && receiver.typeName == "*haxe__io__StringInput") {
						continue;
					}
					if (!requiresIoBufferInputSurface && receiver != null && receiver.typeName == "*haxe__io__BufferInput") {
						continue;
					}
					if (!requiresIoEofStringSurface && receiver != null && receiver.typeName == "*haxe__io__Eof" && name == "String") {
						continue;
					}
					out.push(decl);
				case _:
					out.push(decl);
			}
		}
		return out;
	}

	function trimIoShimToCoreSurface(decls:Array<GoDecl>):Array<GoDecl> {
		var out = new Array<GoDecl>();
		for (decl in decls) {
			switch (decl) {
				case GoDecl.GoInterfaceDecl(name, methods):
					if (name == "haxe__io__Input") {
						out.push(GoDecl.GoInterfaceDecl(name, [for (method in methods) if (!isIoInputHelperMethodName(method.name)) method]));
					} else if (name == "haxe__io__Output") {
						out.push(GoDecl.GoInterfaceDecl(name, [for (method in methods) if (!isIoOutputHelperMethodName(method.name)) method]));
					} else {
						out.push(decl);
					}
				case GoDecl.GoFuncDecl(name, receiver, _, _, _):
					if (receiver == null && isIoHelperFunctionDecl(name)) {
						continue;
					}
					if (receiver != null && isIoInputHelperReceiverType(receiver.typeName) && isIoInputHelperMethodName(name)) {
						continue;
					}
					if (receiver != null && isIoOutputHelperReceiverType(receiver.typeName) && isIoOutputHelperMethodName(name)) {
						continue;
					}
					out.push(decl);
				case _:
					out.push(decl);
			}
		}
		return out;
	}

	function isIoInputHelperMethodName(name:String):Bool {
		return switch (name) {
			case "readAll", "readFullBytes", "read", "readUntil", "readLine", "readFloat", "readDouble", "readInt8", "readInt16", "readUInt16", "readInt24",
				"readUInt24", "readInt32", "readString":
				true;
			case _:
				false;
		};
	}

	function isIoInputHelperReceiverType(typeName:String):Bool {
		return switch (typeName) {
			case "*haxe__io__BytesInput", "*haxe__io__StringInput", "*haxe__io__BufferInput":
				true;
			case _:
				false;
		};
	}

	function isIoOutputHelperMethodName(name:String):Bool {
		return switch (name) {
			case "write", "writeFullBytes", "writeFloat", "writeDouble", "writeInt8", "writeInt16", "writeUInt16", "writeInt24", "writeUInt24", "writeInt32",
				"prepare", "writeInput", "writeString":
				true;
			case _:
				false;
		};
	}

	function isIoOutputHelperReceiverType(typeName:String):Bool {
		return switch (typeName) {
			case "*haxe__io__BytesOutput":
				true;
			case _:
				false;
		};
	}

	function isIoHelperFunctionDecl(name:String):Bool {
		return switch (name) {
			case "haxe__io__input_isEof", "haxe__io__input_readAll", "haxe__io__input_readFullBytes", "haxe__io__input_read", "haxe__io__input_readUntil",
				"haxe__io__input_readLine", "haxe__io__input_readFloat", "haxe__io__input_readDouble", "haxe__io__input_readInt8",
				"haxe__io__input_readInt16", "haxe__io__input_readUInt16", "haxe__io__input_readInt24", "haxe__io__input_readUInt24",
				"haxe__io__input_readInt32", "haxe__io__input_readString", "haxe__io__output_write", "haxe__io__output_writeFullBytes",
				"haxe__io__output_writeFloat", "haxe__io__output_writeDouble", "haxe__io__output_writeInt8", "haxe__io__output_writeInt16",
				"haxe__io__output_writeUInt16", "haxe__io__output_writeInt24", "haxe__io__output_writeUInt24", "haxe__io__output_writeInt32",
				"haxe__io__output_writeInput", "haxe__io__output_writeString":
				true;
			case _:
				false;
		};
	}

	function lowerAtomicStdlibShimDecls():Array<GoDecl> {
		return [
			GoDecl.GoFuncDecl("haxe__atomic___AtomicInt__AtomicInt_Impl___new", null, [{name: "value", typeName: "int"}], ["any"], [
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.AtomicIntNew"), [GoExpr.GoIdent("value")]))
			]),
			GoDecl.GoFuncDecl("haxe__atomic___AtomicInt__AtomicInt_Impl__add", null, [
				{
					name: "atom",
					typeName: "any"
				},
				{name: "value", typeName: "int"}
			], ["int"], [
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.AtomicIntAdd"), [GoExpr.GoIdent("atom"), GoExpr.GoIdent("value")]))
			]),
			GoDecl.GoFuncDecl("haxe__atomic___AtomicInt__AtomicInt_Impl__sub", null, [
				{
					name: "atom",
					typeName: "any"
				},
				{name: "value", typeName: "int"}
			], ["int"], [
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.AtomicIntSub"), [GoExpr.GoIdent("atom"), GoExpr.GoIdent("value")]))
			]),
			GoDecl.GoFuncDecl("haxe__atomic___AtomicInt__AtomicInt_Impl__and", null, [
				{
					name: "atom",
					typeName: "any"
				},
				{name: "value", typeName: "int"}
			], ["int"], [
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.AtomicIntAnd"), [GoExpr.GoIdent("atom"), GoExpr.GoIdent("value")]))
			]),
			GoDecl.GoFuncDecl("haxe__atomic___AtomicInt__AtomicInt_Impl__or", null, [
				{
					name: "atom",
					typeName: "any"
				},
				{name: "value", typeName: "int"}
			], ["int"], [
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.AtomicIntOr"), [GoExpr.GoIdent("atom"), GoExpr.GoIdent("value")]))
			]),
			GoDecl.GoFuncDecl("haxe__atomic___AtomicInt__AtomicInt_Impl__xor", null, [
				{
					name: "atom",
					typeName: "any"
				},
				{name: "value", typeName: "int"}
			], ["int"], [
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.AtomicIntXor"), [GoExpr.GoIdent("atom"), GoExpr.GoIdent("value")]))
			]),
			GoDecl.GoFuncDecl("haxe__atomic___AtomicInt__AtomicInt_Impl__compareExchange", null, [
				{
					name: "atom",
					typeName: "any"
				},
				{name: "expected", typeName: "int"},
				{name: "replacement", typeName: "int"}
			], ["int"], [
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.AtomicIntCompareExchange"), [
					GoExpr.GoIdent("atom"),
					GoExpr.GoIdent("expected"),
					GoExpr.GoIdent("replacement")
				]))
			]),
			GoDecl.GoFuncDecl("haxe__atomic___AtomicInt__AtomicInt_Impl__exchange", null, [
				{
					name: "atom",
					typeName: "any"
				},
				{name: "value", typeName: "int"}
			], ["int"], [
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.AtomicIntExchange"), [GoExpr.GoIdent("atom"), GoExpr.GoIdent("value")]))
			]),
			GoDecl.GoFuncDecl("haxe__atomic___AtomicInt__AtomicInt_Impl__load", null, [
				{
					name: "atom",
					typeName: "any"
				}
			], ["int"], [
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.AtomicIntLoad"), [GoExpr.GoIdent("atom")]))
			]),
			GoDecl.GoFuncDecl("haxe__atomic___AtomicInt__AtomicInt_Impl__store", null, [
				{
					name: "atom",
					typeName: "any"
				},
				{name: "value", typeName: "int"}
			], ["int"], [
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.AtomicIntStore"), [GoExpr.GoIdent("atom"), GoExpr.GoIdent("value")]))
			]),
			GoDecl.GoFuncDecl("haxe__atomic___AtomicObject__AtomicObject_Impl___new", null, [
				{
					name: "value",
					typeName: "any"
				}
			], ["any"], [
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.AtomicObjectNew"), [GoExpr.GoIdent("value")]))
			]),
			GoDecl.GoFuncDecl("haxe__atomic___AtomicObject__AtomicObject_Impl__load", null, [
				{
					name: "atom",
					typeName: "any"
				}
			], ["any"], [
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.AtomicObjectLoad"), [GoExpr.GoIdent("atom")]))
			]),
			GoDecl.GoFuncDecl("haxe__atomic___AtomicObject__AtomicObject_Impl__store", null, [
				{
					name: "atom",
					typeName: "any"
				},
				{name: "value", typeName: "any"}
			], ["any"], [
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.AtomicObjectStore"), [GoExpr.GoIdent("atom"), GoExpr.GoIdent("value")]))
			]),
			GoDecl.GoFuncDecl("haxe__atomic___AtomicObject__AtomicObject_Impl__exchange", null, [
				{
					name: "atom",
					typeName: "any"
				},
				{name: "value", typeName: "any"}
			], ["any"], [
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.AtomicObjectExchange"), [GoExpr.GoIdent("atom"), GoExpr.GoIdent("value")]))
			]),
			GoDecl.GoFuncDecl("haxe__atomic___AtomicObject__AtomicObject_Impl__compareExchange", null, [
				{
					name: "atom",
					typeName: "any"
				},
				{name: "expected", typeName: "any"},
				{name: "replacement", typeName: "any"}
			], ["any"], [
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.AtomicObjectCompareExchange"), [
					GoExpr.GoIdent("atom"),
					GoExpr.GoIdent("expected"),
					GoExpr.GoIdent("replacement")
				]))
			])
		];
	}

	function lowerGoConcurrencyShimDecls():Array<GoDecl> {
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
				GoStmt.GoRaw("chosen, recvValue, _ := reflect.Select(cases)"),
				GoStmt.GoRaw("if chosen == 0 {"),
				GoStmt.GoRaw("\tif !recvValue.IsValid() {"),
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
				GoStmt.GoRaw("chosen, recvValue, _ := reflect.Select(cases)"),
				GoStmt.GoRaw("if chosen == 0 {"),
				GoStmt.GoRaw("\tif !recvValue.IsValid() {"),
				GoStmt.GoRaw("\t\treturn New_go___Result(nil, nil)"),
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
			], [], [GoStmt.GoGoStmt(GoExpr.GoCall(GoExpr.GoIdent("fn"), []))])
		];
		if (useTypedGoConcurrencySpecialization()) {
			decls = decls.concat(lowerMetalGoConcurrencyShimDecls());
		}
		return decls;
	}

	function lowerMetalGoConcurrencyShimDecls():Array<GoDecl> {
		var elementTypes = [for (elementType in requiredMetalChanElementTypes.keys()) elementType];
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
			var makeName = metalChanShimName("go__concurrency_makeChan", elementType);
			var setBufferName = metalChanShimName("go__concurrency_setBuffer", elementType);
			var newChanName = metalChanShimName("go__concurrency_newChan", elementType);
			var sendName = metalChanShimName("go__concurrency_send", elementType);
			var trySendName = metalChanShimName("go__concurrency_trySend", elementType);
			var recvName = metalChanShimName("go__concurrency_recv", elementType);
			var recvOrName = metalChanShimName("go__concurrency_recvOr", elementType);
			var tryRecvName = metalChanShimName("go__concurrency_tryRecv", elementType);
			var closeName = metalChanShimName("go__concurrency_close", elementType);

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
						clause: GoSelectClause.GoSelectRecvAssign(GoExpr.GoIdent("value"),
							GoExpr.GoRecvExpr(GoExpr.GoTypeAssert(GoExpr.GoIdent("channel"), chanType)), true),
						body: [GoStmt.GoReturn(GoExpr.GoIdent("value"))]
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
						clause: GoSelectClause.GoSelectRecvAssign(GoExpr.GoIdent("value"),
							GoExpr.GoRecvExpr(GoExpr.GoTypeAssert(GoExpr.GoIdent("channel"), chanType)), true),
						body: [
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

	function lowerMetalGoCollectionShimDecls():Array<GoDecl> {
		if (!useTypedGoCollectionsSpecialization()) {
			return [];
		}

		var decls = new Array<GoDecl>();
		var sliceElementTypes = [for (elementType in requiredMetalSliceElementTypes.keys()) elementType];
		sliceElementTypes.sort(function(a, b) return Reflect.compare(a, b));
		var sliceTypeName = GoNaming.typeSymbol(["go"], "Slice");
		var slicePointerType = "*" + sliceTypeName;
		for (elementType in sliceElementTypes) {
			var pushName = metalSliceShimName("go__slice_push", elementType);
			var setName = metalSliceShimName("go__slice_set", elementType);
			var getName = metalSliceShimName("go__slice_get", elementType);
			var lengthName = metalSliceShimName("go__slice_length", elementType);
			var toArrayName = metalSliceShimName("go__slice_toArray", elementType);

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
				GoStmt.GoVarDecl("out", "[]" + elementType, GoExpr.GoRaw("make([]" + elementType + ", len(raw))"), true),
				GoStmt.GoRangeStmt("idx", "value", GoExpr.GoIdent("raw"), true, [
					GoStmt.GoIf(GoExpr.GoBinary("==", GoExpr.GoIdent("value"), GoExpr.GoNil), [GoStmt.GoContinue], null),
					GoStmt.GoAssign(GoExpr.GoIndex(GoExpr.GoIdent("out"), GoExpr.GoIdent("idx")), GoExpr.GoTypeAssert(GoExpr.GoIdent("value"), elementType))
				]),
				GoStmt.GoReturn(GoExpr.GoIdent("out"))
			]));
		}

		var mapSignatures = [for (signature in requiredMetalMapTypePairs.keys()) signature];
		mapSignatures.sort(function(a, b) return Reflect.compare(a, b));
		var mapTypeName = GoNaming.typeSymbol(["go"], "Map");
		var mapPointerType = "*" + mapTypeName;
		for (signature in mapSignatures) {
			var pair = requiredMetalMapTypePairs.get(signature);
			if (pair == null) {
				continue;
			}
			var keyType = pair.keyGoType;
			var valueType = pair.valueGoType;
			var setName = metalMapShimName("go__map_set", keyType, valueType);
			var getName = metalMapShimName("go__map_get", keyType, valueType);
			var existsName = metalMapShimName("go__map_exists", keyType, valueType);

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
		return decls.concat(lowerMetalGoResultShimDecls());
	}

	function lowerMetalGoResultShimDecls():Array<GoDecl> {
		if (!useTypedGoResultSpecialization()) {
			return [];
		}

		var elementTypes = [for (elementType in requiredMetalResultElementTypes.keys()) elementType];
		if (elementTypes.length == 0) {
			return [];
		}

		elementTypes.sort(function(a, b) return Reflect.compare(a, b));
		var resultTypeName = GoNaming.typeSymbol(["go"], "Result");
		var resultPointerType = "*" + resultTypeName;
		var decls = new Array<GoDecl>();

		for (elementType in elementTypes) {
			var okName = metalResultShimName("go__result_ok", elementType);
			var failureName = metalResultShimName("go__result_failure", elementType);
			var valueErrorName = metalResultShimName("go__result_valueError", elementType);
			var isOkName = metalResultShimName("go__result_isOk", elementType);
			var isErrName = metalResultShimName("go__result_isErr", elementType);
			var unwrapName = metalResultShimName("go__result_unwrap", elementType);
			var errorName = metalResultShimName("go__result_error", elementType);

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

	function lowerDsStdlibShimDecls():Array<GoDecl> {
		var decls = [
			GoDecl.GoStructDecl("haxe__ds__IntMap", [{name: "h", typeName: "map[int]any"}]),
			GoDecl.GoStructDecl("haxe__ds__StringMap", [{name: "h", typeName: "map[string]any"}]),
			GoDecl.GoStructDecl("haxe__ds__ObjectMap", [{name: "h", typeName: "map[any]any"}]),
			GoDecl.GoStructDecl("haxe__ds__EnumValueMap", [{name: "h", typeName: "map[any]any"}]),
			GoDecl.GoStructDecl("haxe__ds__List", [{name: "items", typeName: "[]any"}, {name: "length", typeName: "int"}]),
			GoDecl.GoFuncDecl("New_haxe__ds__IntMap", null, [], ["*haxe__ds__IntMap"], [GoStmt.GoReturn(GoExpr.GoRaw("&haxe__ds__IntMap{h: map[int]any{}}"))]),
			GoDecl.GoFuncDecl("set", {
				name: "self",
				typeName: "*haxe__ds__IntMap"
			}, [{name: "key", typeName: "any"}, {name: "value", typeName: "any"}], [], [
				GoStmt.GoVarDecl("resolvedKey", "int", GoExpr.GoCall(GoExpr.GoIdent("hxrt.IntFromNullableAny"), [GoExpr.GoIdent("key")]), true),
				GoStmt.GoAssign(GoExpr.GoIndex(GoExpr.GoSelector(GoExpr.GoIdent("self"), "h"), GoExpr.GoIdent("resolvedKey")), GoExpr.GoIdent("value"))
			]),
			GoDecl.GoFuncDecl("get", {
				name: "self",
				typeName: "*haxe__ds__IntMap"
			}, [{name: "key", typeName: "any"}], ["any"], [
				GoStmt.GoVarDecl("resolvedKey", "int", GoExpr.GoCall(GoExpr.GoIdent("hxrt.IntFromNullableAny"), [GoExpr.GoIdent("key")]), true),
				GoStmt.GoVarDecl("value", null, GoExpr.GoIndex(GoExpr.GoSelector(GoExpr.GoIdent("self"), "h"), GoExpr.GoIdent("resolvedKey")), true),
				GoStmt.GoReturn(GoExpr.GoIdent("value"))
			]),
			GoDecl.GoFuncDecl("exists", {
				name: "self",
				typeName: "*haxe__ds__IntMap"
			}, [{name: "key", typeName: "any"}], ["bool"], [
				GoStmt.GoVarDecl("resolvedKey", "int", GoExpr.GoCall(GoExpr.GoIdent("hxrt.IntFromNullableAny"), [GoExpr.GoIdent("key")]), true),
				GoStmt.GoRaw("_, ok := self.h[resolvedKey]"),
				GoStmt.GoReturn(GoExpr.GoIdent("ok"))
			]),
			GoDecl.GoFuncDecl("remove", {
				name: "self",
				typeName: "*haxe__ds__IntMap"
			}, [{name: "key", typeName: "any"}], ["bool"], [
				GoStmt.GoVarDecl("resolvedKey", "int", GoExpr.GoCall(GoExpr.GoIdent("hxrt.IntFromNullableAny"), [GoExpr.GoIdent("key")]), true),
				GoStmt.GoRaw("_, ok := self.h[resolvedKey]"),
				GoStmt.GoRaw("delete(self.h, resolvedKey)"),
				GoStmt.GoReturn(GoExpr.GoIdent("ok"))
			]),
			GoDecl.GoFuncDecl("keys", {
				name: "self",
				typeName: "*haxe__ds__IntMap"
			}, [], ["map[string]any"], [
				GoStmt.GoRaw("keys := make([]int, 0, len(self.h))"),
				GoStmt.GoRaw("for key := range self.h {"),
				GoStmt.GoRaw("\tkeys = append(keys, key)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("index := 0"),
				GoStmt.GoRaw("iter := map[string]any{}"),
				GoStmt.GoRaw("iter[\"hasNext\"] = func() bool { return index < len(keys) }"),
				GoStmt.GoRaw("iter[\"next\"] = func() int { key := keys[index]; index++; return key }"),
				GoStmt.GoReturn(GoExpr.GoIdent("iter"))
			]),
			GoDecl.GoFuncDecl("iterator", {
				name: "self",
				typeName: "*haxe__ds__IntMap"
			}, [], ["map[string]any"], [
				GoStmt.GoRaw("keys := make([]int, 0, len(self.h))"),
				GoStmt.GoRaw("for key := range self.h {"),
				GoStmt.GoRaw("\tkeys = append(keys, key)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("index := 0"),
				GoStmt.GoRaw("iter := map[string]any{}"),
				GoStmt.GoRaw("iter[\"hasNext\"] = func() bool { return index < len(keys) }"),
				GoStmt.GoRaw("iter[\"next\"] = func() any { key := keys[index]; index++; return self.h[key] }"),
				GoStmt.GoReturn(GoExpr.GoIdent("iter"))
			]),
			GoDecl.GoFuncDecl("keyValueIterator", {
				name: "self",
				typeName: "*haxe__ds__IntMap"
			}, [], ["map[string]any"], [
				GoStmt.GoRaw("keys := make([]int, 0, len(self.h))"),
				GoStmt.GoRaw("for key := range self.h {"),
				GoStmt.GoRaw("\tkeys = append(keys, key)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("index := 0"),
				GoStmt.GoRaw("iter := map[string]any{}"),
				GoStmt.GoRaw("iter[\"hasNext\"] = func() bool { return index < len(keys) }"),
				GoStmt.GoRaw("iter[\"next\"] = func() map[string]any { key := keys[index]; index++; return map[string]any{\"key\": key, \"value\": self.h[key]} }"),
				GoStmt.GoReturn(GoExpr.GoIdent("iter"))
			]),
			GoDecl.GoFuncDecl("copyIMap", {
				name: "self",
				typeName: "*haxe__ds__IntMap"
			}, [], ["haxe__IMap"], [
				GoStmt.GoRaw("copied := New_haxe__ds__IntMap()"),
				GoStmt.GoRaw("for key, value := range self.h {"),
				GoStmt.GoRaw("\tcopied.h[key] = value"),
				GoStmt.GoRaw("}"),
				GoStmt.GoReturn(GoExpr.GoIdent("copied"))
			]),
			GoDecl.GoFuncDecl("toString", {
				name: "self",
				typeName: "*haxe__ds__IntMap"
			}, [], ["*string"], [
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoStringLiteral("{}")]))
			]),
			GoDecl.GoFuncDecl("clear", {
				name: "self",
				typeName: "*haxe__ds__IntMap"
			},
				[], [], [GoStmt.GoRaw("self.h = map[int]any{}")]),
			GoDecl.GoFuncDecl("New_haxe__ds__StringMap", null, [], ["*haxe__ds__StringMap"],
				[GoStmt.GoReturn(GoExpr.GoRaw("&haxe__ds__StringMap{h: map[string]any{}}"))]),
			GoDecl.GoFuncDecl("set", {
				name: "self",
				typeName: "*haxe__ds__StringMap"
			}, [{name: "key", typeName: "any"}, {name: "value", typeName: "any"}], [], [
				GoStmt.GoAssign(GoExpr.GoRaw("self.h[*hxrt.StdString(key)]"), GoExpr.GoIdent("value"))
			]),
			GoDecl.GoFuncDecl("get", {name: "self", typeName: "*haxe__ds__StringMap"}, [{name: "key", typeName: "any"}], ["any"], [
				GoStmt.GoRaw("value := self.h[*hxrt.StdString(key)]"),
				GoStmt.GoReturn(GoExpr.GoIdent("value"))
			]),
			GoDecl.GoFuncDecl("exists", {name: "self", typeName: "*haxe__ds__StringMap"}, [{name: "key", typeName: "any"}], ["bool"], [
				GoStmt.GoRaw("_, ok := self.h[*hxrt.StdString(key)]"),
				GoStmt.GoReturn(GoExpr.GoIdent("ok"))
			]),
			GoDecl.GoFuncDecl("remove", {name: "self", typeName: "*haxe__ds__StringMap"}, [{name: "key", typeName: "any"}], ["bool"], [
				GoStmt.GoRaw("_, ok := self.h[*hxrt.StdString(key)]"),
				GoStmt.GoRaw("delete(self.h, *hxrt.StdString(key))"),
				GoStmt.GoReturn(GoExpr.GoIdent("ok"))
			]),
			GoDecl.GoFuncDecl("keys", {
				name: "self",
				typeName: "*haxe__ds__StringMap"
			}, [], ["map[string]any"], [
				GoStmt.GoRaw("keys := make([]string, 0, len(self.h))"),
				GoStmt.GoRaw("for key := range self.h {"),
				GoStmt.GoRaw("\tkeys = append(keys, key)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("index := 0"),
				GoStmt.GoRaw("iter := map[string]any{}"),
				GoStmt.GoRaw("iter[\"hasNext\"] = func() bool { return index < len(keys) }"),
				GoStmt.GoRaw("iter[\"next\"] = func() *string { key := keys[index]; index++; return hxrt.StringFromLiteral(key) }"),
				GoStmt.GoReturn(GoExpr.GoIdent("iter"))
			]),
			GoDecl.GoFuncDecl("iterator", {
				name: "self",
				typeName: "*haxe__ds__StringMap"
			}, [], ["map[string]any"], [
				GoStmt.GoRaw("keys := make([]string, 0, len(self.h))"),
				GoStmt.GoRaw("for key := range self.h {"),
				GoStmt.GoRaw("\tkeys = append(keys, key)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("index := 0"),
				GoStmt.GoRaw("iter := map[string]any{}"),
				GoStmt.GoRaw("iter[\"hasNext\"] = func() bool { return index < len(keys) }"),
				GoStmt.GoRaw("iter[\"next\"] = func() any { key := keys[index]; index++; return self.h[key] }"),
				GoStmt.GoReturn(GoExpr.GoIdent("iter"))
			]),
			GoDecl.GoFuncDecl("keyValueIterator", {
				name: "self",
				typeName: "*haxe__ds__StringMap"
			}, [], ["map[string]any"], [
				GoStmt.GoRaw("keys := make([]string, 0, len(self.h))"),
				GoStmt.GoRaw("for key := range self.h {"),
				GoStmt.GoRaw("\tkeys = append(keys, key)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("index := 0"),
				GoStmt.GoRaw("iter := map[string]any{}"),
				GoStmt.GoRaw("iter[\"hasNext\"] = func() bool { return index < len(keys) }"),
				GoStmt.GoRaw("iter[\"next\"] = func() map[string]any { key := keys[index]; index++; return map[string]any{\"key\": hxrt.StringFromLiteral(key), \"value\": self.h[key]} }"),
				GoStmt.GoReturn(GoExpr.GoIdent("iter"))
			]),
			GoDecl.GoFuncDecl("copyIMap", {
				name: "self",
				typeName: "*haxe__ds__StringMap"
			}, [], ["haxe__IMap"], [
				GoStmt.GoRaw("copied := New_haxe__ds__StringMap()"),
				GoStmt.GoRaw("for key, value := range self.h {"),
				GoStmt.GoRaw("\tcopied.h[key] = value"),
				GoStmt.GoRaw("}"),
				GoStmt.GoReturn(GoExpr.GoIdent("copied"))
			]),
			GoDecl.GoFuncDecl("toString", {
				name: "self",
				typeName: "*haxe__ds__StringMap"
			}, [], ["*string"], [
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoStringLiteral("{}")]))
			]),
			GoDecl.GoFuncDecl("clear", {
				name: "self",
				typeName: "*haxe__ds__StringMap"
			},
				[], [], [GoStmt.GoRaw("self.h = map[string]any{}")]),
			GoDecl.GoFuncDecl("New_haxe__ds__ObjectMap", null, [], ["*haxe__ds__ObjectMap"],
				[GoStmt.GoReturn(GoExpr.GoRaw("&haxe__ds__ObjectMap{h: map[any]any{}}"))]),
			GoDecl.GoFuncDecl("set", {
				name: "self",
				typeName: "*haxe__ds__ObjectMap"
			}, [{name: "key", typeName: "any"}, {name: "value", typeName: "any"}], [], [
				GoStmt.GoAssign(GoExpr.GoIndex(GoExpr.GoSelector(GoExpr.GoIdent("self"), "h"), GoExpr.GoIdent("key")), GoExpr.GoIdent("value"))
			]),
			GoDecl.GoFuncDecl("get", {
				name: "self",
				typeName: "*haxe__ds__ObjectMap"
			}, [{name: "key", typeName: "any"}], ["any"],
				[
					GoStmt.GoReturn(GoExpr.GoIndex(GoExpr.GoSelector(GoExpr.GoIdent("self"), "h"), GoExpr.GoIdent("key")))
				]),
			GoDecl.GoFuncDecl("exists", {name: "self", typeName: "*haxe__ds__ObjectMap"}, [{name: "key", typeName: "any"}], ["bool"],
				[GoStmt.GoRaw("_, ok := self.h[key]"), GoStmt.GoReturn(GoExpr.GoIdent("ok"))]),
			GoDecl.GoFuncDecl("remove", {name: "self", typeName: "*haxe__ds__ObjectMap"}, [{name: "key", typeName: "any"}], ["bool"], [
				GoStmt.GoRaw("_, ok := self.h[key]"),
				GoStmt.GoRaw("delete(self.h, key)"),
				GoStmt.GoReturn(GoExpr.GoIdent("ok"))
			]),
			GoDecl.GoFuncDecl("keys", {
				name: "self",
				typeName: "*haxe__ds__ObjectMap"
			}, [], ["map[string]any"], [
				GoStmt.GoRaw("keys := make([]any, 0, len(self.h))"),
				GoStmt.GoRaw("for key := range self.h {"),
				GoStmt.GoRaw("\tkeys = append(keys, key)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("index := 0"),
				GoStmt.GoRaw("iter := map[string]any{}"),
				GoStmt.GoRaw("iter[\"hasNext\"] = func() bool { return index < len(keys) }"),
				GoStmt.GoRaw("iter[\"next\"] = func() any { key := keys[index]; index++; return key }"),
				GoStmt.GoReturn(GoExpr.GoIdent("iter"))
			]),
			GoDecl.GoFuncDecl("iterator", {
				name: "self",
				typeName: "*haxe__ds__ObjectMap"
			}, [], ["map[string]any"], [
				GoStmt.GoRaw("keys := make([]any, 0, len(self.h))"),
				GoStmt.GoRaw("for key := range self.h {"),
				GoStmt.GoRaw("\tkeys = append(keys, key)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("index := 0"),
				GoStmt.GoRaw("iter := map[string]any{}"),
				GoStmt.GoRaw("iter[\"hasNext\"] = func() bool { return index < len(keys) }"),
				GoStmt.GoRaw("iter[\"next\"] = func() any { key := keys[index]; index++; return self.h[key] }"),
				GoStmt.GoReturn(GoExpr.GoIdent("iter"))
			]),
			GoDecl.GoFuncDecl("keyValueIterator", {
				name: "self",
				typeName: "*haxe__ds__ObjectMap"
			}, [], ["map[string]any"], [
				GoStmt.GoRaw("keys := make([]any, 0, len(self.h))"),
				GoStmt.GoRaw("for key := range self.h {"),
				GoStmt.GoRaw("\tkeys = append(keys, key)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("index := 0"),
				GoStmt.GoRaw("iter := map[string]any{}"),
				GoStmt.GoRaw("iter[\"hasNext\"] = func() bool { return index < len(keys) }"),
				GoStmt.GoRaw("iter[\"next\"] = func() map[string]any { key := keys[index]; index++; return map[string]any{\"key\": key, \"value\": self.h[key]} }"),
				GoStmt.GoReturn(GoExpr.GoIdent("iter"))
			]),
			GoDecl.GoFuncDecl("copyIMap", {
				name: "self",
				typeName: "*haxe__ds__ObjectMap"
			}, [], ["haxe__IMap"], [
				GoStmt.GoRaw("copied := New_haxe__ds__ObjectMap()"),
				GoStmt.GoRaw("for key, value := range self.h {"),
				GoStmt.GoRaw("\tcopied.h[key] = value"),
				GoStmt.GoRaw("}"),
				GoStmt.GoReturn(GoExpr.GoIdent("copied"))
			]),
			GoDecl.GoFuncDecl("toString", {
				name: "self",
				typeName: "*haxe__ds__ObjectMap"
			}, [], ["*string"], [
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoStringLiteral("{}")]))
			]),
			GoDecl.GoFuncDecl("clear", {
				name: "self",
				typeName: "*haxe__ds__ObjectMap"
			},
				[], [], [GoStmt.GoRaw("self.h = map[any]any{}")]),
			GoDecl.GoFuncDecl("New_haxe__ds__EnumValueMap", null, [], ["*haxe__ds__EnumValueMap"],
				[GoStmt.GoReturn(GoExpr.GoRaw("&haxe__ds__EnumValueMap{h: map[any]any{}}"))]),
			GoDecl.GoFuncDecl("set", {
				name: "self",
				typeName: "*haxe__ds__EnumValueMap"
			}, [{name: "key", typeName: "any"}, {name: "value", typeName: "any"}], [], [
				GoStmt.GoAssign(GoExpr.GoIndex(GoExpr.GoSelector(GoExpr.GoIdent("self"), "h"), GoExpr.GoIdent("key")), GoExpr.GoIdent("value"))
			]),
			GoDecl.GoFuncDecl("get", {
				name: "self",
				typeName: "*haxe__ds__EnumValueMap"
			}, [{name: "key", typeName: "any"}], ["any"],
				[
					GoStmt.GoReturn(GoExpr.GoIndex(GoExpr.GoSelector(GoExpr.GoIdent("self"), "h"), GoExpr.GoIdent("key")))
				]),
			GoDecl.GoFuncDecl("exists", {name: "self", typeName: "*haxe__ds__EnumValueMap"}, [{name: "key", typeName: "any"}], ["bool"],
				[GoStmt.GoRaw("_, ok := self.h[key]"), GoStmt.GoReturn(GoExpr.GoIdent("ok"))]),
			GoDecl.GoFuncDecl("remove", {name: "self", typeName: "*haxe__ds__EnumValueMap"}, [{name: "key", typeName: "any"}], ["bool"], [
				GoStmt.GoRaw("_, ok := self.h[key]"),
				GoStmt.GoRaw("delete(self.h, key)"),
				GoStmt.GoReturn(GoExpr.GoIdent("ok"))
			]),
			GoDecl.GoFuncDecl("keys", {
				name: "self",
				typeName: "*haxe__ds__EnumValueMap"
			}, [], ["map[string]any"], [
				GoStmt.GoRaw("keys := make([]any, 0, len(self.h))"),
				GoStmt.GoRaw("for key := range self.h {"),
				GoStmt.GoRaw("\tkeys = append(keys, key)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("index := 0"),
				GoStmt.GoRaw("iter := map[string]any{}"),
				GoStmt.GoRaw("iter[\"hasNext\"] = func() bool { return index < len(keys) }"),
				GoStmt.GoRaw("iter[\"next\"] = func() any { key := keys[index]; index++; return key }"),
				GoStmt.GoReturn(GoExpr.GoIdent("iter"))
			]),
			GoDecl.GoFuncDecl("iterator", {
				name: "self",
				typeName: "*haxe__ds__EnumValueMap"
			}, [], ["map[string]any"], [
				GoStmt.GoRaw("keys := make([]any, 0, len(self.h))"),
				GoStmt.GoRaw("for key := range self.h {"),
				GoStmt.GoRaw("\tkeys = append(keys, key)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("index := 0"),
				GoStmt.GoRaw("iter := map[string]any{}"),
				GoStmt.GoRaw("iter[\"hasNext\"] = func() bool { return index < len(keys) }"),
				GoStmt.GoRaw("iter[\"next\"] = func() any { key := keys[index]; index++; return self.h[key] }"),
				GoStmt.GoReturn(GoExpr.GoIdent("iter"))
			]),
			GoDecl.GoFuncDecl("keyValueIterator", {
				name: "self",
				typeName: "*haxe__ds__EnumValueMap"
			}, [], ["map[string]any"], [
				GoStmt.GoRaw("keys := make([]any, 0, len(self.h))"),
				GoStmt.GoRaw("for key := range self.h {"),
				GoStmt.GoRaw("\tkeys = append(keys, key)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("index := 0"),
				GoStmt.GoRaw("iter := map[string]any{}"),
				GoStmt.GoRaw("iter[\"hasNext\"] = func() bool { return index < len(keys) }"),
				GoStmt.GoRaw("iter[\"next\"] = func() map[string]any { key := keys[index]; index++; return map[string]any{\"key\": key, \"value\": self.h[key]} }"),
				GoStmt.GoReturn(GoExpr.GoIdent("iter"))
			]),
			GoDecl.GoFuncDecl("copyIMap", {
				name: "self",
				typeName: "*haxe__ds__EnumValueMap"
			}, [], ["haxe__IMap"], [
				GoStmt.GoRaw("copied := New_haxe__ds__EnumValueMap()"),
				GoStmt.GoRaw("for key, value := range self.h {"),
				GoStmt.GoRaw("\tcopied.h[key] = value"),
				GoStmt.GoRaw("}"),
				GoStmt.GoReturn(GoExpr.GoIdent("copied"))
			]),
			GoDecl.GoFuncDecl("toString", {
				name: "self",
				typeName: "*haxe__ds__EnumValueMap"
			}, [], ["*string"], [
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoStringLiteral("{}")]))
			]),
			GoDecl.GoFuncDecl("clear", {
				name: "self",
				typeName: "*haxe__ds__EnumValueMap"
			},
				[], [], [GoStmt.GoRaw("self.h = map[any]any{}")]),
			GoDecl.GoFuncDecl("New_haxe__ds__List", null, [], ["*haxe__ds__List"],
				[GoStmt.GoReturn(GoExpr.GoRaw("&haxe__ds__List{items: []any{}, length: 0}"))]),
			GoDecl.GoFuncDecl("add", {
				name: "self",
				typeName: "*haxe__ds__List"
			}, [{name: "item", typeName: "any"}], [], [
				GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("self"), "items"),
					GoExpr.GoCall(GoExpr.GoIdent("append"), [GoExpr.GoSelector(GoExpr.GoIdent("self"), "items"), GoExpr.GoIdent("item")])),
				GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("self"), "length"),
					GoExpr.GoCall(GoExpr.GoIdent("len"), [GoExpr.GoSelector(GoExpr.GoIdent("self"), "items")]))
			]),
			GoDecl.GoFuncDecl("push", {
				name: "self",
				typeName: "*haxe__ds__List"
			}, [{name: "item", typeName: "any"}], [], [
				GoStmt.GoRaw("self.items = append([]any{item}, self.items...)"),
				GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("self"), "length"),
					GoExpr.GoCall(GoExpr.GoIdent("len"), [GoExpr.GoSelector(GoExpr.GoIdent("self"), "items")]))
			]),
			GoDecl.GoFuncDecl("pop", {
				name: "self",
				typeName: "*haxe__ds__List"
			}, [], ["any"], [
				GoStmt.GoIf(GoExpr.GoBinary("==", GoExpr.GoCall(GoExpr.GoIdent("len"), [GoExpr.GoSelector(GoExpr.GoIdent("self"), "items")]),
					GoExpr.GoIntLiteral(0)),
					[GoStmt.GoReturn(GoExpr.GoNil)], null),
				GoStmt.GoVarDecl("head", null, GoExpr.GoIndex(GoExpr.GoSelector(GoExpr.GoIdent("self"), "items"), GoExpr.GoIntLiteral(0)), true),
				GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("self"), "items"),
					GoExpr.GoSlice(GoExpr.GoSelector(GoExpr.GoIdent("self"), "items"), GoExpr.GoIntLiteral(1), null)),
				GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("self"), "length"),
					GoExpr.GoCall(GoExpr.GoIdent("len"), [GoExpr.GoSelector(GoExpr.GoIdent("self"), "items")])),
				GoStmt.GoReturn(GoExpr.GoIdent("head"))
			]),
			GoDecl.GoFuncDecl("first", {
				name: "self",
				typeName: "*haxe__ds__List"
			}, [], ["any"], [
				GoStmt.GoIf(GoExpr.GoBinary("==", GoExpr.GoCall(GoExpr.GoIdent("len"), [GoExpr.GoSelector(GoExpr.GoIdent("self"), "items")]),
					GoExpr.GoIntLiteral(0)),
					[GoStmt.GoReturn(GoExpr.GoNil)], null),
				GoStmt.GoReturn(GoExpr.GoIndex(GoExpr.GoSelector(GoExpr.GoIdent("self"), "items"), GoExpr.GoIntLiteral(0)))
			]),
			GoDecl.GoFuncDecl("last", {
				name: "self",
				typeName: "*haxe__ds__List"
			}, [], ["any"], [
				GoStmt.GoVarDecl("size", null, GoExpr.GoCall(GoExpr.GoIdent("len"), [GoExpr.GoSelector(GoExpr.GoIdent("self"), "items")]), true),
				GoStmt.GoIf(GoExpr.GoBinary("==", GoExpr.GoIdent("size"), GoExpr.GoIntLiteral(0)), [GoStmt.GoReturn(GoExpr.GoNil)], null),
				GoStmt.GoReturn(GoExpr.GoIndex(GoExpr.GoSelector(GoExpr.GoIdent("self"), "items"),
					GoExpr.GoBinary("-", GoExpr.GoIdent("size"), GoExpr.GoIntLiteral(1))))
			])
		];
		return decls;
	}

	function lowerHttpStdlibShimDecls():Array<GoDecl> {
		return [
			GoDecl.GoStructDecl("hxrt__http__Pair", [{name: "name", typeName: "*string"}, {name: "value", typeName: "*string"}]),
			GoDecl.GoStructDecl("hxrt__http__FileUpload", [
				{name: "param", typeName: "*string"},
				{name: "filename", typeName: "*string"},
				{name: "size", typeName: "int"},
				{name: "mimeType", typeName: "*string"},
				{name: "fileRef", typeName: "any"}
			]),
			GoDecl.GoGlobalVarDecl("sys__Http_PROXY", "any", GoExpr.GoNil),
			GoDecl.GoStructDecl("sys__Http", [
				{name: "url", typeName: "*string"},
				{name: "responseAsString", typeName: "*string"},
				{name: "responseBytes", typeName: "*haxe__io__Bytes"},
				{name: "postData", typeName: "*string"},
				{name: "postBytes", typeName: "*haxe__io__Bytes"},
				{name: "headers", typeName: "[]hxrt__http__Pair"},
				{name: "params", typeName: "[]hxrt__http__Pair"},
				{name: "onData", typeName: "func(*string)"},
				{name: "onBytes", typeName: "func(*haxe__io__Bytes)"},
				{name: "onError", typeName: "func(*string)"},
				{name: "onStatus", typeName: "func(int)"},
				{name: "noShutdown", typeName: "bool"},
				{name: "cnxTimeout", typeName: "float64"},
				{name: "responseHeaders", typeName: "*haxe__ds__StringMap"},
				{name: "responseHeadersSameKey", typeName: "map[string][]*string"},
				{name: "fileUpload", typeName: "*hxrt__http__FileUpload"}
			]),
			GoDecl.GoFuncDecl("New_sys__Http", null, [{name: "url", typeName: "*string"}], ["*sys__Http"], [
				GoStmt.GoVarDecl("self", null,
					GoExpr.GoRaw("&sys__Http{url: url, headers: []hxrt__http__Pair{}, params: []hxrt__http__Pair{}, cnxTimeout: 10, responseHeaders: New_haxe__ds__StringMap(), responseHeadersSameKey: map[string][]*string{}}"),
					true),
				GoStmt.GoRaw("self.onData = func(data *string) {}"),
				GoStmt.GoRaw("self.onBytes = func(data *haxe__io__Bytes) {}"),
				GoStmt.GoRaw("self.onError = func(msg *string) {}"),
				GoStmt.GoRaw("self.onStatus = func(status int) {}"),
				GoStmt.GoReturn(GoExpr.GoIdent("self"))
			]),
			GoDecl.GoFuncDecl("setHeader", {
				name: "self",
				typeName: "*sys__Http"
			}, [{name: "name", typeName: "*string"}, {name: "value", typeName: "*string"}],
				[], [
					GoStmt.GoIf(GoExpr.GoBinary("==", GoExpr.GoIdent("self"), GoExpr.GoNil), [GoStmt.GoReturn(null)], null),
					GoStmt.GoRaw("for i := 0; i < len(self.headers); i++ {"),
					GoStmt.GoRaw("\tif *hxrt.StdString(self.headers[i].name) == *hxrt.StdString(name) {"),
					GoStmt.GoRaw("\t\tself.headers[i] = hxrt__http__Pair{name: name, value: value}"),
					GoStmt.GoRaw("\t\treturn"),
					GoStmt.GoRaw("\t}"),
					GoStmt.GoRaw("}"),
					GoStmt.GoRaw("self.headers = append(self.headers, hxrt__http__Pair{name: name, value: value})")
				]),
			GoDecl.GoFuncDecl("addHeader", {
				name: "self",
				typeName: "*sys__Http"
			},
				[{name: "header", typeName: "*string"}, {name: "value", typeName: "*string"}], [], [
					GoStmt.GoIf(GoExpr.GoBinary("==", GoExpr.GoIdent("self"), GoExpr.GoNil), [GoStmt.GoReturn(null)], null),
					GoStmt.GoRaw("self.headers = append(self.headers, hxrt__http__Pair{name: header, value: value})")
				]),
			GoDecl.GoFuncDecl("setParameter", {
				name: "self",
				typeName: "*sys__Http"
			}, [{name: "name", typeName: "*string"}, {name: "value", typeName: "*string"}],
				[], [
					GoStmt.GoIf(GoExpr.GoBinary("==", GoExpr.GoIdent("self"), GoExpr.GoNil), [GoStmt.GoReturn(null)], null),
					GoStmt.GoRaw("for i := 0; i < len(self.params); i++ {"),
					GoStmt.GoRaw("\tif *hxrt.StdString(self.params[i].name) == *hxrt.StdString(name) {"),
					GoStmt.GoRaw("\t\tself.params[i] = hxrt__http__Pair{name: name, value: value}"),
					GoStmt.GoRaw("\t\treturn"),
					GoStmt.GoRaw("\t}"),
					GoStmt.GoRaw("}"),
					GoStmt.GoRaw("self.params = append(self.params, hxrt__http__Pair{name: name, value: value})")
				]),
			GoDecl.GoFuncDecl("addParameter", {
				name: "self",
				typeName: "*sys__Http"
			}, [{name: "name", typeName: "*string"}, {name: "value", typeName: "*string"}],
				[], [
					GoStmt.GoIf(GoExpr.GoBinary("==", GoExpr.GoIdent("self"), GoExpr.GoNil), [GoStmt.GoReturn(null)], null),
					GoStmt.GoRaw("self.params = append(self.params, hxrt__http__Pair{name: name, value: value})")
				]),
			GoDecl.GoFuncDecl("setPostData", {
				name: "self",
				typeName: "*sys__Http"
			}, [{name: "data", typeName: "*string"}], [], [
				GoStmt.GoIf(GoExpr.GoBinary("==", GoExpr.GoIdent("self"), GoExpr.GoNil), [GoStmt.GoReturn(null)], null),
				GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("self"), "postData"), GoExpr.GoIdent("data")),
				GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("self"), "postBytes"), GoExpr.GoNil)
			]),
			GoDecl.GoFuncDecl("setPostBytes", {
				name: "self",
				typeName: "*sys__Http"
			}, [{name: "data", typeName: "*haxe__io__Bytes"}], [], [
				GoStmt.GoIf(GoExpr.GoBinary("==", GoExpr.GoIdent("self"), GoExpr.GoNil), [GoStmt.GoReturn(null)], null),
				GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("self"), "postBytes"), GoExpr.GoIdent("data")),
				GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("self"), "postData"), GoExpr.GoNil)
			]),
			GoDecl.GoFuncDecl("fileTransfer", {
				name: "self",
				typeName: "*sys__Http"
			}, [
				{name: "argname", typeName: "*string"},
				{name: "filename", typeName: "*string"},
				{name: "file", typeName: "any"},
				{name: "size", typeName: "int"},
				{name: "mimeType", typeName: "...*string"}
			], [], [
				GoStmt.GoIf(GoExpr.GoBinary("==", GoExpr.GoIdent("self"), GoExpr.GoNil), [GoStmt.GoReturn(null)], null),
				GoStmt.GoVarDecl("resolvedMime", null,
					GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoStringLiteral("application/octet-stream")]), true),
				GoStmt.GoIf(GoExpr.GoBinary(">", GoExpr.GoCall(GoExpr.GoIdent("len"), [GoExpr.GoIdent("mimeType")]), GoExpr.GoIntLiteral(0)), [
					GoStmt.GoIf(GoExpr.GoBinary("!=", GoExpr.GoIndex(GoExpr.GoIdent("mimeType"), GoExpr.GoIntLiteral(0)), GoExpr.GoNil), [
						GoStmt.GoAssign(GoExpr.GoIdent("resolvedMime"), GoExpr.GoIndex(GoExpr.GoIdent("mimeType"), GoExpr.GoIntLiteral(0)))
					],
						null)
				],
					null),
				GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("self"), "fileUpload"),
					GoExpr.GoRaw("&hxrt__http__FileUpload{param: argname, filename: filename, size: size, mimeType: resolvedMime, fileRef: file}"))
			]),
			GoDecl.GoFuncDecl("fileTransfert", {
				name: "self",
				typeName: "*sys__Http"
			}, [
				{name: "argname", typeName: "*string"},
				{name: "filename", typeName: "*string"},
				{name: "file", typeName: "any"},
				{name: "size", typeName: "int"},
				{name: "mimeType", typeName: "...*string"}
			],
				[], [GoStmt.GoRaw("self.fileTransfer(argname, filename, file, size, mimeType...)")]),
			GoDecl.GoFuncDecl("getResponseHeaderValues", {name: "self", typeName: "*sys__Http"}, [{name: "key", typeName: "*string"}], ["[]*string"], [
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("sys__GoHttpHelpers_getResponseHeaderValues"), [GoExpr.GoIdent("self"), GoExpr.GoIdent("key")]))
			]),
			GoDecl.GoFuncDecl("get_responseData", {
				name: "self",
				typeName: "*sys__Http"
			}, [], ["*string"], [
				GoStmt.GoIf(GoExpr.GoBinary("==", GoExpr.GoIdent("self"), GoExpr.GoNil), [
					GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoStringLiteral("")]))
				], null),
				GoStmt.GoIf(GoExpr.GoRaw("self.responseAsString == nil && self.responseBytes != nil"), [
					GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("self"), "responseAsString"),
						GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoSelector(GoExpr.GoIdent("self"), "responseBytes"), "toString"), []))
				], null),
				GoStmt.GoReturn(GoExpr.GoSelector(GoExpr.GoIdent("self"), "responseAsString"))
			]),
			GoDecl.GoFuncDecl("customRequest", {
				name: "self",
				typeName: "*sys__Http"
			}, [
				{name: "post", typeName: "bool"},
				{name: "api", typeName: "any"},
				{name: "rest", typeName: "...any"}
			], [], [
				GoStmt.GoVarDecl("socketOverride", "any", GoExpr.GoNil, false),
				GoStmt.GoVarDecl("methodOverride", "*string", GoExpr.GoNil, false),
				GoStmt.GoRaw("if len(rest) >= 1 {"),
				GoStmt.GoRaw("\tswitch candidate := rest[0].(type) {"),
				GoStmt.GoRaw("\tcase string:"),
				GoStmt.GoRaw("\t\tif len(rest) == 1 {"),
				GoStmt.GoRaw("\t\t\tmethodOverride = hxrt.StringFromLiteral(candidate)"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\tcase *string:"),
				GoStmt.GoRaw("\t\tif len(rest) == 1 {"),
				GoStmt.GoRaw("\t\t\tmethodOverride = candidate"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\tdefault:"),
				GoStmt.GoRaw("\t\tsocketOverride = candidate"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if len(rest) >= 2 {"),
				GoStmt.GoRaw("\tswitch candidate := rest[1].(type) {"),
				GoStmt.GoRaw("\tcase *string:"),
				GoStmt.GoRaw("\t\tmethodOverride = candidate"),
				GoStmt.GoRaw("\tcase string:"),
				GoStmt.GoRaw("\t\tmethodOverride = hxrt.StringFromLiteral(candidate)"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("}"),
				GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("self"), "hxrt__http__requestWith"), [
					GoExpr.GoIdent("post"),
					GoExpr.GoIdent("methodOverride"),
					GoExpr.GoIdent("api"),
					GoExpr.GoIdent("socketOverride")
				]))
			]),
			GoDecl.GoFuncDecl("request", {
				name: "self",
				typeName: "*sys__Http"
			}, [{name: "post", typeName: "...bool"}], [], [
				GoStmt.GoIf(GoExpr.GoBinary("==", GoExpr.GoIdent("self"), GoExpr.GoNil), [GoStmt.GoReturn(null)], null),
				GoStmt.GoVarDecl("isPost", "bool", GoExpr.GoBoolLiteral(false), true),
				GoStmt.GoIf(GoExpr.GoBinary(">", GoExpr.GoCall(GoExpr.GoIdent("len"), [GoExpr.GoIdent("post")]), GoExpr.GoIntLiteral(0)), [
					GoStmt.GoAssign(GoExpr.GoIdent("isPost"), GoExpr.GoIndex(GoExpr.GoIdent("post"), GoExpr.GoIntLiteral(0)))
				],
					null),
				GoStmt.GoIf(GoExpr.GoRaw("self.postData != nil || self.postBytes != nil || self.fileUpload != nil"),
					[GoStmt.GoAssign(GoExpr.GoIdent("isPost"), GoExpr.GoBoolLiteral(true))], null),
				GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("self"), "hxrt__http__requestWith"),
					[GoExpr.GoIdent("isPost"), GoExpr.GoNil, GoExpr.GoNil, GoExpr.GoNil]))
			]),
			GoDecl.GoFuncDecl("hxrt__http__requestWith", {
				name: "self",
				typeName: "*sys__Http"
			}, [
				{name: "post", typeName: "bool"},
				{name: "methodOverride", typeName: "*string"},
				{name: "api", typeName: "any"},
				{name: "sock", typeName: "any"}
			], [], [
				GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("self"), "responseAsString"), GoExpr.GoNil),
				GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("self"), "responseBytes"), GoExpr.GoNil),
				GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("self"), "responseHeaders"), GoExpr.GoCall(GoExpr.GoIdent("New_haxe__ds__StringMap"), [])),
				GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("self"), "responseHeadersSameKey"), GoExpr.GoRaw("map[string][]*string{}")),
				GoStmt.GoVarDecl("rawUrl", null, GoExpr.GoRaw("*hxrt.StdString(self.url)"), true),
				GoStmt.GoRaw("parsedURL, err := url.Parse(rawUrl)"),
				GoStmt.GoIf(GoExpr.GoRaw("err != nil || parsedURL == nil"), [
					GoStmt.GoIf(GoExpr.GoRaw("self.onError != nil"), [
						GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("self"), "onError"),
							[
								GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoStringLiteral("Invalid URL")])
							]))
					],
						null),
					GoStmt.GoReturn(null)
				],
					null),
				GoStmt.GoRaw("query := parsedURL.Query()"),
				GoStmt.GoRaw("for _, param := range self.params {"),
				GoStmt.GoRaw("\tquery.Set(*hxrt.StdString(param.name), *hxrt.StdString(param.value))"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("var bodyReader io.Reader = nil"),
				GoStmt.GoVarDecl("contentTypeOverride", "*string", GoExpr.GoNil, false),
				GoStmt.GoRaw("if post {"),
				GoStmt.GoRaw("\tif self.fileUpload != nil {"),
				GoStmt.GoRaw("\t\tmultipartPayload := \"\""),
				GoStmt.GoRaw("\t\tfor _, param := range self.params {"),
				GoStmt.GoRaw("\t\t\tmultipartPayload += \"--hxrt-go-boundary\\r\\n\""),
				GoStmt.GoRaw("\t\t\tmultipartPayload += \"Content-Disposition: form-data; name=\\\"\" + *hxrt.StdString(param.name) + \"\\\"\\r\\n\\r\\n\""),
				GoStmt.GoRaw("\t\t\tmultipartPayload += *hxrt.StdString(param.value) + \"\\r\\n\""),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t\tmultipartPayload += \"--hxrt-go-boundary\\r\\n\""),
				GoStmt.GoRaw("\t\tmultipartPayload += \"Content-Disposition: form-data; name=\\\"\" + *hxrt.StdString(self.fileUpload.param) + \"\\\"; filename=\\\"\" + *hxrt.StdString(self.fileUpload.filename) + \"\\\"\\r\\n\""),
				GoStmt.GoRaw("\t\tmultipartPayload += \"Content-Type: \" + *hxrt.StdString(self.fileUpload.mimeType) + \"\\r\\n\\r\\n\""),
				GoStmt.GoRaw("\t\tmultipartPayload += \"[uploaded-bytes=\" + *hxrt.StdString(self.fileUpload.size) + \"]\\r\\n\""),
				GoStmt.GoRaw("\t\tmultipartPayload += \"--hxrt-go-boundary--\\r\\n\""),
				GoStmt.GoRaw("\t\tbodyReader = strings.NewReader(multipartPayload)"),
				GoStmt.GoRaw("\t\tcontentTypeOverride = hxrt.StringFromLiteral(\"multipart/form-data; boundary=hxrt-go-boundary\")"),
				GoStmt.GoRaw("\t} else if self.postBytes != nil {"),
				GoStmt.GoRaw("\t\trawBody := make([]byte, len(self.postBytes.b))"),
				GoStmt.GoRaw("\t\tfor i := 0; i < len(self.postBytes.b); i++ {"),
				GoStmt.GoRaw("\t\t\trawBody[i] = byte(self.postBytes.b[i])"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t\tbodyReader = bytes.NewReader(rawBody)"),
				GoStmt.GoRaw("\t} else if self.postData != nil {"),
				GoStmt.GoRaw("\t\tbodyReader = strings.NewReader(*hxrt.StdString(self.postData))"),
				GoStmt.GoRaw("\t} else {"),
				GoStmt.GoRaw("\t\tencoded := query.Encode()"),
				GoStmt.GoRaw("\t\tbodyReader = strings.NewReader(encoded)"),
				GoStmt.GoRaw("\t\thasContentType := false"),
				GoStmt.GoRaw("\t\tfor _, header := range self.headers {"),
				GoStmt.GoRaw("\t\t\tif strings.EqualFold(*hxrt.StdString(header.name), \"Content-Type\") {"),
				GoStmt.GoRaw("\t\t\t\thasContentType = true"),
				GoStmt.GoRaw("\t\t\t\tbreak"),
				GoStmt.GoRaw("\t\t\t}"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t\tif !hasContentType {"),
				GoStmt.GoRaw("\t\t\tcontentTypeOverride = hxrt.StringFromLiteral(\"application/x-www-form-urlencoded\")"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("} else {"),
				GoStmt.GoRaw("\tparsedURL.RawQuery = query.Encode()"),
				GoStmt.GoRaw("}"),
				GoStmt.GoIf(GoExpr.GoRaw("parsedURL.Scheme == \"data\""), [
					GoStmt.GoVarDecl("payload", null, GoExpr.GoRaw("parsedURL.Opaque"), true),
					GoStmt.GoVarDecl("mediaType", null, GoExpr.GoStringLiteral("text/plain"), true),
					GoStmt.GoRaw("commaIndex := strings.Index(payload, \",\")"),
					GoStmt.GoRaw("if commaIndex >= 0 {"),
					GoStmt.GoRaw("\tif commaIndex > 0 {"),
					GoStmt.GoRaw("\t\tmediaType = payload[:commaIndex]"),
					GoStmt.GoRaw("\t}"),
					GoStmt.GoRaw("\tpayload = payload[commaIndex+1:]"),
					GoStmt.GoRaw("}"),
					GoStmt.GoRaw("if post {"),
					GoStmt.GoRaw("\tif self.fileUpload != nil {"),
					GoStmt.GoRaw("\t\tpayload = \"multipart file=\" + *hxrt.StdString(self.fileUpload.filename) + \";mime=\" + *hxrt.StdString(self.fileUpload.mimeType) + \";size=\" + *hxrt.StdString(self.fileUpload.size)"),
					GoStmt.GoRaw("\t} else if bodyReader != nil {"),
					GoStmt.GoRaw("\t\trawBody, readErr := io.ReadAll(bodyReader)"),
					GoStmt.GoRaw("\t\tif readErr == nil {"),
					GoStmt.GoRaw("\t\t\tpayload = string(rawBody)"),
					GoStmt.GoRaw("\t\t}"),
					GoStmt.GoRaw("\t}"),
					GoStmt.GoRaw("}"),
					GoStmt.GoRaw("decoded, decodeErr := url.QueryUnescape(payload)"),
					GoStmt.GoRaw("if decodeErr == nil {"),
					GoStmt.GoRaw("\tpayload = decoded"),
					GoStmt.GoRaw("}"),
					GoStmt.GoRaw("if methodOverride != nil {"),
					GoStmt.GoRaw("\tmethodToken := strings.ToUpper(*hxrt.StdString(methodOverride))"),
					GoStmt.GoRaw("\tif methodToken != \"\" && methodToken != \"NULL\" {"),
					GoStmt.GoRaw("\t\tpayload = methodToken + \" \" + payload"),
					GoStmt.GoRaw("\t}"),
					GoStmt.GoRaw("}"),
					GoStmt.GoRaw("rawPayload := []byte(payload)"),
					GoStmt.GoRaw("intPayload := make([]int, len(rawPayload))"),
					GoStmt.GoRaw("for i := 0; i < len(rawPayload); i++ {"),
					GoStmt.GoRaw("\tintPayload[i] = int(rawPayload[i])"),
					GoStmt.GoRaw("}"),
					GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("self"), "responseBytes"),
						GoExpr.GoRaw("&haxe__io__Bytes{b: intPayload, length: len(intPayload)}")),
					GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("self"), "responseAsString"),
						GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoIdent("payload")])),
					GoStmt.GoRaw("self.responseHeaders = New_haxe__ds__StringMap()"),
					GoStmt.GoRaw("self.responseHeaders.set(hxrt.StringFromLiteral(\"content-type\"), hxrt.StringFromLiteral(mediaType))"),
					GoStmt.GoRaw("self.responseHeaders.set(hxrt.StringFromLiteral(\"Content-Type\"), hxrt.StringFromLiteral(mediaType))"),
					GoStmt.GoRaw("self.responseHeadersSameKey = map[string][]*string{}"),
					GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("hxrt__http__captureApi"), [
						GoExpr.GoIdent("api"),
						GoExpr.GoSelector(GoExpr.GoIdent("self"), "responseBytes")
					])),
					GoStmt.GoIf(GoExpr.GoRaw("self.onStatus != nil"), [
						GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("self"), "onStatus"), [GoExpr.GoIntLiteral(200)]))
					], null),
					GoStmt.GoIf(GoExpr.GoRaw("self.onData != nil"), [
						GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("self"), "onData"),
							[GoExpr.GoSelector(GoExpr.GoIdent("self"), "responseAsString")]))
					], null),
					GoStmt.GoIf(GoExpr.GoRaw("self.onBytes != nil"), [
						GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("self"), "onBytes"),
							[GoExpr.GoSelector(GoExpr.GoIdent("self"), "responseBytes")]))
					], null),
					GoStmt.GoReturn(null)
				], null),
				GoStmt.GoIf(GoExpr.GoRaw("parsedURL.Scheme == \"\" || parsedURL.Host == \"\""), [
					GoStmt.GoIf(GoExpr.GoRaw("self.onError != nil"), [
						GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("self"), "onError"),
							[
								GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoStringLiteral("Invalid URL")])
							]))
					],
						null),
					GoStmt.GoReturn(null)
				],
					null),
				GoStmt.GoRaw("method := \"GET\""),
				GoStmt.GoRaw("if post {"),
				GoStmt.GoRaw("\tmethod = \"POST\""),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if methodOverride != nil {"),
				GoStmt.GoRaw("\tmethodToken := strings.ToUpper(*hxrt.StdString(methodOverride))"),
				GoStmt.GoRaw("\tif methodToken != \"\" && methodToken != \"NULL\" {"),
				GoStmt.GoRaw("\t\tmethod = methodToken"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("request, err := http.NewRequest(method, parsedURL.String(), bodyReader)"),
				GoStmt.GoIf(GoExpr.GoBinary("!=", GoExpr.GoIdent("err"), GoExpr.GoNil), [
					GoStmt.GoIf(GoExpr.GoRaw("self.onError != nil"), [
						GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("self"), "onError"),
							[
								GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoRaw("err.Error()")])
							]))
					],
						null),
					GoStmt.GoReturn(null)
				],
					null),
				GoStmt.GoRaw("for _, header := range self.headers {"),
				GoStmt.GoRaw("\trequest.Header.Set(*hxrt.StdString(header.name), *hxrt.StdString(header.value))"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if contentTypeOverride != nil && request.Header.Get(\"Content-Type\") == \"\" {"),
				GoStmt.GoRaw("\trequest.Header.Set(\"Content-Type\", *hxrt.StdString(contentTypeOverride))"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("transport := &http.Transport{}"),
				GoStmt.GoRaw("proxyURL := hxrt__http__proxyURL()"),
				GoStmt.GoRaw("if proxyURL != nil {"),
				GoStmt.GoRaw("\ttransport.Proxy = http.ProxyURL(proxyURL)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("var socketAdapter interface {"),
				GoStmt.GoRaw("\thxrt__socket_conn() net.Conn"),
				GoStmt.GoRaw("\thxrt__socket_setConn(net.Conn)"),
				GoStmt.GoRaw("\tclose()"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if candidate, ok := sock.(interface {"),
				GoStmt.GoRaw("\thxrt__socket_conn() net.Conn"),
				GoStmt.GoRaw("\thxrt__socket_setConn(net.Conn)"),
				GoStmt.GoRaw("\tclose()"),
				GoStmt.GoRaw("}); ok {"),
				GoStmt.GoRaw("\tsocketAdapter = candidate"),
				GoStmt.GoRaw("\ttransport.DisableKeepAlives = true"),
				GoStmt.GoRaw("\trequest.Close = true"),
				GoStmt.GoRaw("\tsocketConsumed := false"),
				GoStmt.GoRaw("\ttransport.Dial = func(network string, addr string) (net.Conn, error) {"),
				GoStmt.GoRaw("\t\tif socketConsumed {"),
				GoStmt.GoRaw("\t\t\treturn nil, io.EOF"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t\tsocketConsumed = true"),
				GoStmt.GoRaw("\t\tconn := socketAdapter.hxrt__socket_conn()"),
				GoStmt.GoRaw("\t\tif conn == nil {"),
				GoStmt.GoRaw("\t\t\tdialConn, dialErr := net.Dial(network, addr)"),
				GoStmt.GoRaw("\t\t\tif dialErr != nil {"),
				GoStmt.GoRaw("\t\t\t\treturn nil, dialErr"),
				GoStmt.GoRaw("\t\t\t}"),
				GoStmt.GoRaw("\t\t\tsocketAdapter.hxrt__socket_setConn(dialConn)"),
				GoStmt.GoRaw("\t\t\tconn = dialConn"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t\treturn conn, nil"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tdefer socketAdapter.close()"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("timeout := time.Duration(self.cnxTimeout * float64(time.Second))"),
				GoStmt.GoRaw("if timeout <= 0 {"),
				GoStmt.GoRaw("\ttimeout = 10 * time.Second"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("client := &http.Client{Transport: transport, Timeout: timeout}"),
				GoStmt.GoRaw("response, err := client.Do(request)"),
				GoStmt.GoIf(GoExpr.GoBinary("!=", GoExpr.GoIdent("err"), GoExpr.GoNil), [
					GoStmt.GoIf(GoExpr.GoRaw("self.onError != nil"), [
						GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("self"), "onError"),
							[
								GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoRaw("err.Error()")])
							]))
					],
						null),
					GoStmt.GoReturn(null)
				],
					null),
				GoStmt.GoRaw("defer response.Body.Close()"),
				GoStmt.GoRaw("self.responseHeaders = New_haxe__ds__StringMap()"),
				GoStmt.GoRaw("self.responseHeadersSameKey = map[string][]*string{}"),
				GoStmt.GoRaw("for name, values := range response.Header {"),
				GoStmt.GoRaw("\tif len(values) == 0 {"),
				GoStmt.GoRaw("\t\tcontinue"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tlowerKey := strings.ToLower(name)"),
				GoStmt.GoRaw("\tlastValue := hxrt.StringFromLiteral(values[len(values)-1])"),
				GoStmt.GoRaw("\tself.responseHeaders.set(hxrt.StringFromLiteral(name), lastValue)"),
				GoStmt.GoRaw("\tif lowerKey != name {"),
				GoStmt.GoRaw("\t\tself.responseHeaders.set(hxrt.StringFromLiteral(lowerKey), lastValue)"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tif len(values) > 1 {"),
				GoStmt.GoRaw("\t\tallValues := make([]*string, 0, len(values))"),
				GoStmt.GoRaw("\t\tfor _, rawValue := range values {"),
				GoStmt.GoRaw("\t\t\tallValues = append(allValues, hxrt.StringFromLiteral(rawValue))"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t\tself.responseHeadersSameKey[name] = allValues"),
				GoStmt.GoRaw("\t\tif lowerKey != name {"),
				GoStmt.GoRaw("\t\t\tself.responseHeadersSameKey[lowerKey] = allValues"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("}"),
				GoStmt.GoIf(GoExpr.GoRaw("self.onStatus != nil"), [
					GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("self"), "onStatus"), [GoExpr.GoRaw("response.StatusCode")]))
				], null),
				GoStmt.GoRaw("rawPayload, err := io.ReadAll(response.Body)"),
				GoStmt.GoIf(GoExpr.GoBinary("!=", GoExpr.GoIdent("err"), GoExpr.GoNil), [
					GoStmt.GoIf(GoExpr.GoRaw("self.onError != nil"), [
						GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("self"), "onError"),
							[
								GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoRaw("err.Error()")])
							]))
					],
						null),
					GoStmt.GoReturn(null)
				],
					null),
				GoStmt.GoRaw("intPayload := make([]int, len(rawPayload))"),
				GoStmt.GoRaw("for i := 0; i < len(rawPayload); i++ {"),
				GoStmt.GoRaw("\tintPayload[i] = int(rawPayload[i])"),
				GoStmt.GoRaw("}"),
				GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("self"), "responseBytes"),
					GoExpr.GoRaw("&haxe__io__Bytes{b: intPayload, length: len(intPayload)}")),
				GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("self"), "responseAsString"),
					GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoRaw("string(rawPayload)")])),
				GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("hxrt__http__captureApi"), [
					GoExpr.GoIdent("api"),
					GoExpr.GoSelector(GoExpr.GoIdent("self"), "responseBytes")
				])),
				GoStmt.GoIf(GoExpr.GoRaw("response.StatusCode >= 400"), [
					GoStmt.GoIf(GoExpr.GoRaw("self.onError != nil"), [
						GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("self"), "onError"), [
							GoExpr.GoRaw("hxrt.StringConcatAny(hxrt.StringFromLiteral(\"Http Error #\"), response.StatusCode)")
						]))
					], null),
					GoStmt.GoReturn(null)
				], null),
				GoStmt.GoIf(GoExpr.GoRaw("self.onData != nil"), [
					GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("self"), "onData"),
						[GoExpr.GoSelector(GoExpr.GoIdent("self"), "responseAsString")]))
				], null),
				GoStmt.GoIf(GoExpr.GoRaw("self.onBytes != nil"), [
					GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("self"), "onBytes"),
						[GoExpr.GoSelector(GoExpr.GoIdent("self"), "responseBytes")]))
				], null)
			]),
			GoDecl.GoFuncDecl("hxrt__http__captureApi", null, [
				{
					name: "api",
					typeName: "any"
				},
				{name: "payload", typeName: "*haxe__io__Bytes"}
			], [], [
				GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("sys__GoHttpHelpers_captureApi"), [GoExpr.GoIdent("api"), GoExpr.GoIdent("payload")]))
			]),
			GoDecl.GoFuncDecl("hxrt__http__proxyURL", null, [], ["*url.URL"], [
				GoStmt.GoIf(GoExpr.GoBinary("==", GoExpr.GoIdent("sys__Http_PROXY"), GoExpr.GoNil), [GoStmt.GoReturn(GoExpr.GoNil)], null),
				GoStmt.GoRaw("config, ok := sys__Http_PROXY.(map[string]any)"),
				GoStmt.GoIf(GoExpr.GoUnary("!", GoExpr.GoIdent("ok")), [GoStmt.GoReturn(GoExpr.GoNil)], null),
				GoStmt.GoRaw("host := *hxrt.StdString(config[\"host\"] )"),
				GoStmt.GoIf(GoExpr.GoBinary("==", GoExpr.GoIdent("host"), GoExpr.GoStringLiteral("")), [GoStmt.GoReturn(GoExpr.GoNil)], null),
				GoStmt.GoIf(GoExpr.GoBinary("==", GoExpr.GoIdent("host"), GoExpr.GoStringLiteral("null")), [GoStmt.GoReturn(GoExpr.GoNil)], null),
				GoStmt.GoRaw("port := *hxrt.StdString(config[\"port\"] )"),
				GoStmt.GoRaw("hostPort := host"),
				GoStmt.GoRaw("if port != \"\" && port != \"null\" && !strings.Contains(hostPort, \":\") {"),
				GoStmt.GoRaw("\thostPort = hostPort + \":\" + port"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("proxyURL, err := url.Parse(\"http://\" + hostPort)"),
				GoStmt.GoIf(GoExpr.GoBinary("!=", GoExpr.GoIdent("err"), GoExpr.GoNil), [GoStmt.GoReturn(GoExpr.GoNil)], null),
				GoStmt.GoRaw("if authValue, ok := config[\"auth\"]; ok {"),
				GoStmt.GoRaw("\tif authMap, ok := authValue.(map[string]any); ok {"),
				GoStmt.GoRaw("\t\tuser := *hxrt.StdString(authMap[\"user\"])"),
				GoStmt.GoRaw("\t\tpass := *hxrt.StdString(authMap[\"pass\"])"),
				GoStmt.GoRaw("\t\tif user != \"\" && user != \"null\" {"),
				GoStmt.GoRaw("\t\t\tif pass == \"null\" {"),
				GoStmt.GoRaw("\t\t\t\tpass = \"\""),
				GoStmt.GoRaw("\t\t\t}"),
				GoStmt.GoRaw("\t\t\tproxyURL.User = url.UserPassword(user, pass)"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("}"),
				GoStmt.GoReturn(GoExpr.GoIdent("proxyURL"))
			]),
			GoDecl.GoFuncDecl("sys__Http_hxrt_proxyDescriptor", null, [], ["*string"], [
				GoStmt.GoVarDecl("proxyURL", null, GoExpr.GoCall(GoExpr.GoIdent("hxrt__http__proxyURL"), []), true),
				GoStmt.GoIf(GoExpr.GoBinary("==", GoExpr.GoIdent("proxyURL"), GoExpr.GoNil), [
					GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoStringLiteral("null")]))
				],
					null),
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"),
					[GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("proxyURL"), "String"), [])]))
			]),
			GoDecl.GoFuncDecl("sys__Http_requestUrl", null, [
				{
					name: "url",
					typeName: "*string"
				}
			], ["*string"], [
				GoStmt.GoVarDecl("self", null, GoExpr.GoCall(GoExpr.GoIdent("New_sys__Http"), [GoExpr.GoIdent("url")]), true),
				GoStmt.GoVarDecl("result", null, GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoStringLiteral("")]), true),
				GoStmt.GoRaw("self.onData = func(data *string) { result = data }"),
				GoStmt.GoRaw("self.onError = func(msg *string) { result = msg }"),
				GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("self"), "request"), [])),
				GoStmt.GoReturn(GoExpr.GoIdent("result"))
			])
		];
	}

	function lowerSysStdlibShimDecls():Array<GoDecl> {
		var decls = [
			GoDecl.GoStructDecl("Sys", []),
			GoDecl.GoStructDecl("sys__io__File", []),
			GoDecl.GoInterfaceDecl("I_sys__io__FileInput", []),
			GoDecl.GoStructDecl("sys__io__FileInput", [
				{
					name: "__hx_this",
					typeName: "I_sys__io__FileInput"
				},
				{name: "__hx_io_bigEndian", typeName: "bool"}
			]),
			GoDecl.GoFuncDecl("New_sys__io__FileInput", null, [], ["*sys__io__FileInput"], [
				GoStmt.GoRaw("self := &sys__io__FileInput{}"),
				GoStmt.GoRaw("self.__hx_this = self"),
				GoStmt.GoReturn(GoExpr.GoIdent("self"))
			]),
			GoDecl.GoInterfaceDecl("I_sys__io__FileOutput", []),
			GoDecl.GoStructDecl("sys__io__FileOutput", [
				{
					name: "__hx_this",
					typeName: "I_sys__io__FileOutput"
				},
				{name: "__hx_io_bigEndian", typeName: "bool"}
			]),
			GoDecl.GoFuncDecl("New_sys__io__FileOutput", null, [], ["*sys__io__FileOutput"], [
				GoStmt.GoRaw("self := &sys__io__FileOutput{}"),
				GoStmt.GoRaw("self.__hx_this = self"),
				GoStmt.GoReturn(GoExpr.GoIdent("self"))
			]),
			GoDecl.GoStructDecl("sys__io__FileSeek",
				[
					{
						name: "tag",
						typeName: "int"
					},
					{name: "params", typeName: "[]any"}
				]),
			GoDecl.GoGlobalVarDecl("sys__io__FileSeek_SeekBegin", "*sys__io__FileSeek", GoExpr.GoRaw("&sys__io__FileSeek{tag: 0}")),
			GoDecl.GoGlobalVarDecl("sys__io__FileSeek_SeekCur", "*sys__io__FileSeek", GoExpr.GoRaw("&sys__io__FileSeek{tag: 1}")),
			GoDecl.GoGlobalVarDecl("sys__io__FileSeek_SeekEnd", "*sys__io__FileSeek", GoExpr.GoRaw("&sys__io__FileSeek{tag: 2}")),
			GoDecl.GoStructDecl("sys__io__ProcessOutput", [
				{
					name: "impl",
					typeName: "*hxrt.ProcessOutput"
				}
			]),
			GoDecl.GoStructDecl("sys__io__Process",
				[
					{name: "impl", typeName: "*hxrt.Process"},
					{name: "stdout", typeName: "*sys__io__ProcessOutput"}
				]),
			GoDecl.GoGlobalVarDecl("sys__io__fileInputHandles", "map[*sys__io__FileInput]*hxrt.FileInput",
				GoExpr.GoRaw("map[*sys__io__FileInput]*hxrt.FileInput{}")),
			GoDecl.GoGlobalVarDecl("sys__io__fileOutputHandles", "map[*sys__io__FileOutput]*hxrt.FileOutput",
				GoExpr.GoRaw("map[*sys__io__FileOutput]*hxrt.FileOutput{}")),
			GoDecl.GoFuncDecl("sys__io__fileSeekWhence", null, [
				{
					name: "pos",
					typeName: "*sys__io__FileSeek"
				}
			], ["int"], [
				GoStmt.GoRaw("if pos == nil {"),
				GoStmt.GoRaw("\treturn 0"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("switch pos.tag {"),
				GoStmt.GoRaw("case 1:"),
				GoStmt.GoRaw("\treturn 1"),
				GoStmt.GoRaw("case 2:"),
				GoStmt.GoRaw("\treturn 2"),
				GoStmt.GoRaw("default:"),
				GoStmt.GoRaw("\treturn 0"),
				GoStmt.GoRaw("}")
			]),
			GoDecl.GoFuncDecl("Sys_getCwd", null, [], ["*string"], [
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("hxrt"), "SysGetCwd"), []))
			]),
			GoDecl.GoFuncDecl("Sys_args", null, [], ["[]*string"], [
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("hxrt"), "SysArgs"), []))
			]),
			GoDecl.GoFuncDecl("Sys_getEnv", null, [
				{
					name: "key",
					typeName: "*string"
				}
			], ["*string"], [
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("hxrt"), "SysGetEnv"), [GoExpr.GoIdent("key")]))
			]),
			GoDecl.GoFuncDecl("Sys_putEnv", null, [
				{
					name: "key",
					typeName: "*string"
				},
				{name: "value", typeName: "*string"}
			], [], [
				GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("hxrt"), "SysPutEnv"), [GoExpr.GoIdent("key"), GoExpr.GoIdent("value")]))
			]),
			GoDecl.GoFuncDecl("Sys_systemName", null, [], ["*string"], [
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("hxrt"), "SysSystemName"), []))
			]),
			GoDecl.GoFuncDecl("sys__io__File_saveContent", null, [
				{
					name: "path",
					typeName: "*string"
				},
				{name: "content", typeName: "*string"}
			], [], [
				GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("hxrt"), "FileSaveContent"),
					[GoExpr.GoIdent("path"), GoExpr.GoIdent("content")]))
			]),
			GoDecl.GoFuncDecl("sys__io__File_getContent", null, [
				{
					name: "path",
					typeName: "*string"
				}
			], ["*string"], [
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("hxrt"), "FileGetContent"), [GoExpr.GoIdent("path")]))
			]),
			GoDecl.GoFuncDecl("sys__io__File_getBytes", null, [
				{
					name: "path",
					typeName: "*string"
				}
			], ["*haxe__io__Bytes"], [
				GoStmt.GoRaw("raw, err := hxrt.FileGetBytes(path)"),
				GoStmt.GoRaw("if err != nil {"),
				GoStmt.GoRaw("\thxrt.Throw(hxrt.StringFromLiteral(err.Error()))"),
				GoStmt.GoRaw("\treturn &haxe__io__Bytes{b: []int{}, length: 0}"),
				GoStmt.GoRaw("}"),
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt_rawToHaxeBytes"), [GoExpr.GoIdent("raw")]))
			]),
			GoDecl.GoFuncDecl("sys__io__File_saveBytes", null, [
				{
					name: "path",
					typeName: "*string"
				},
				{name: "bytes", typeName: "*haxe__io__Bytes"}
			], [], [
				GoStmt.GoRaw("if err := hxrt.FileSaveBytes(path, hxrt_haxeBytesToRaw(bytes)); err != nil {"),
				GoStmt.GoRaw("\thxrt.Throw(hxrt.StringFromLiteral(err.Error()))"),
				GoStmt.GoRaw("}")
			]),
			GoDecl.GoFuncDecl("sys__io__File_copy", null, [
				{
					name: "srcPath",
					typeName: "*string"
				},
				{name: "dstPath", typeName: "*string"}
			], [], [
				GoStmt.GoRaw("if err := hxrt.FileCopy(srcPath, dstPath); err != nil {"),
				GoStmt.GoRaw("\thxrt.Throw(hxrt.StringFromLiteral(err.Error()))"),
				GoStmt.GoRaw("}")
			]),
			GoDecl.GoFuncDecl("sys__io__File_read", null, [
				{
					name: "path",
					typeName: "*string"
				},
				{name: "binary", typeName: "bool"}
			], ["*sys__io__FileInput"], [
				GoStmt.GoAssign(GoExpr.GoIdent("_"), GoExpr.GoIdent("binary")),
				GoStmt.GoRaw("impl, err := hxrt.OpenFileInput(path)"),
				GoStmt.GoRaw("if err != nil {"),
				GoStmt.GoRaw("\thxrt.Throw(hxrt.StringFromLiteral(err.Error()))"),
				GoStmt.GoRaw("\treturn New_sys__io__FileInput()"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("self := New_sys__io__FileInput()"),
				GoStmt.GoRaw("sys__io__fileInputHandles[self] = impl"),
				GoStmt.GoReturn(GoExpr.GoIdent("self"))
			]),
			GoDecl.GoFuncDecl("sys__io__File_write", null, [
				{
					name: "path",
					typeName: "*string"
				},
				{name: "binary", typeName: "bool"}
			], ["*sys__io__FileOutput"], [
				GoStmt.GoAssign(GoExpr.GoIdent("_"), GoExpr.GoIdent("binary")),
				GoStmt.GoRaw("impl, err := hxrt.OpenFileWriteOutput(path)"),
				GoStmt.GoRaw("if err != nil {"),
				GoStmt.GoRaw("\thxrt.Throw(hxrt.StringFromLiteral(err.Error()))"),
				GoStmt.GoRaw("\treturn New_sys__io__FileOutput()"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("self := New_sys__io__FileOutput()"),
				GoStmt.GoRaw("sys__io__fileOutputHandles[self] = impl"),
				GoStmt.GoReturn(GoExpr.GoIdent("self"))
			]),
			GoDecl.GoFuncDecl("sys__io__File_append", null, [
				{
					name: "path",
					typeName: "*string"
				},
				{name: "binary", typeName: "bool"}
			], ["*sys__io__FileOutput"], [
				GoStmt.GoAssign(GoExpr.GoIdent("_"), GoExpr.GoIdent("binary")),
				GoStmt.GoRaw("impl, err := hxrt.OpenFileAppendOutput(path)"),
				GoStmt.GoRaw("if err != nil {"),
				GoStmt.GoRaw("\thxrt.Throw(hxrt.StringFromLiteral(err.Error()))"),
				GoStmt.GoRaw("\treturn New_sys__io__FileOutput()"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("self := New_sys__io__FileOutput()"),
				GoStmt.GoRaw("sys__io__fileOutputHandles[self] = impl"),
				GoStmt.GoReturn(GoExpr.GoIdent("self"))
			]),
			GoDecl.GoFuncDecl("sys__io__File_update", null, [
				{
					name: "path",
					typeName: "*string"
				},
				{name: "binary", typeName: "bool"}
			], ["*sys__io__FileOutput"], [
				GoStmt.GoAssign(GoExpr.GoIdent("_"), GoExpr.GoIdent("binary")),
				GoStmt.GoRaw("impl, err := hxrt.OpenFileUpdateOutput(path)"),
				GoStmt.GoRaw("if err != nil {"),
				GoStmt.GoRaw("\thxrt.Throw(hxrt.StringFromLiteral(err.Error()))"),
				GoStmt.GoRaw("\treturn New_sys__io__FileOutput()"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("self := New_sys__io__FileOutput()"),
				GoStmt.GoRaw("sys__io__fileOutputHandles[self] = impl"),
				GoStmt.GoReturn(GoExpr.GoIdent("self"))
			]),
			GoDecl.GoFuncDecl("tell", {
				name: "self",
				typeName: "*sys__io__FileInput"
			}, [], ["int"], [
				GoStmt.GoRaw("impl := sys__io__fileInputHandles[self]"),
				GoStmt.GoRaw("if impl == nil {"),
				GoStmt.GoRaw("\treturn 0"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("pos, err := impl.Tell()"),
				GoStmt.GoRaw("if err != nil {"),
				GoStmt.GoRaw("\thxrt.Throw(hxrt.StringFromLiteral(err.Error()))"),
				GoStmt.GoRaw("\treturn 0"),
				GoStmt.GoRaw("}"),
				GoStmt.GoReturn(GoExpr.GoIdent("pos"))
			]),
			GoDecl.GoFuncDecl("seek", {
				name: "self",
				typeName: "*sys__io__FileInput"
			},
				[{name: "p", typeName: "int"}, {name: "pos", typeName: "*sys__io__FileSeek"}], [], [
					GoStmt.GoRaw("impl := sys__io__fileInputHandles[self]"),
					GoStmt.GoRaw("if impl == nil {"),
					GoStmt.GoRaw("\treturn"),
					GoStmt.GoRaw("}"),
					GoStmt.GoRaw("if err := impl.Seek(p, sys__io__fileSeekWhence(pos)); err != nil {"),
					GoStmt.GoRaw("\thxrt.Throw(hxrt.StringFromLiteral(err.Error()))"),
					GoStmt.GoRaw("}")
				]),
			GoDecl.GoFuncDecl("eof", {
				name: "self",
				typeName: "*sys__io__FileInput"
			}, [], ["bool"], [
				GoStmt.GoRaw("impl := sys__io__fileInputHandles[self]"),
				GoStmt.GoRaw("if impl == nil {"),
				GoStmt.GoRaw("\treturn true"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("eof, err := impl.Eof()"),
				GoStmt.GoRaw("if err != nil {"),
				GoStmt.GoRaw("\thxrt.Throw(hxrt.StringFromLiteral(err.Error()))"),
				GoStmt.GoRaw("\treturn true"),
				GoStmt.GoRaw("}"),
				GoStmt.GoReturn(GoExpr.GoIdent("eof"))
			]),
			GoDecl.GoFuncDecl("tell", {
				name: "self",
				typeName: "*sys__io__FileOutput"
			}, [], ["int"], [
				GoStmt.GoRaw("impl := sys__io__fileOutputHandles[self]"),
				GoStmt.GoRaw("if impl == nil {"),
				GoStmt.GoRaw("\treturn 0"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("pos, err := impl.Tell()"),
				GoStmt.GoRaw("if err != nil {"),
				GoStmt.GoRaw("\thxrt.Throw(hxrt.StringFromLiteral(err.Error()))"),
				GoStmt.GoRaw("\treturn 0"),
				GoStmt.GoRaw("}"),
				GoStmt.GoReturn(GoExpr.GoIdent("pos"))
			]),
			GoDecl.GoFuncDecl("seek", {
				name: "self",
				typeName: "*sys__io__FileOutput"
			},
				[{name: "p", typeName: "int"}, {name: "pos", typeName: "*sys__io__FileSeek"}], [], [
					GoStmt.GoRaw("impl := sys__io__fileOutputHandles[self]"),
					GoStmt.GoRaw("if impl == nil {"),
					GoStmt.GoRaw("\treturn"),
					GoStmt.GoRaw("}"),
					GoStmt.GoRaw("if err := impl.Seek(p, sys__io__fileSeekWhence(pos)); err != nil {"),
					GoStmt.GoRaw("\thxrt.Throw(hxrt.StringFromLiteral(err.Error()))"),
					GoStmt.GoRaw("}")
				]),
			GoDecl.GoFuncDecl("New_sys__io__Process", null, [
				{
					name: "command",
					typeName: "*string"
				},
				{name: "args", typeName: "[]*string"}
			], ["*sys__io__Process"], [
				GoStmt.GoVarDecl("impl", null,
					GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("hxrt"), "NewProcess"), [GoExpr.GoIdent("command"), GoExpr.GoIdent("args")]), true),
				GoStmt.GoVarDecl("stdout", null, GoExpr.GoRaw("&sys__io__ProcessOutput{}"), true),
				GoStmt.GoRaw("if impl != nil {"),
				GoStmt.GoRaw("\tstdout.impl = impl.Stdout()"),
				GoStmt.GoRaw("}"),
				GoStmt.GoReturn(GoExpr.GoRaw("&sys__io__Process{impl: impl, stdout: stdout}"))
			]),
			GoDecl.GoFuncDecl("readLine", {
				name: "self",
				typeName: "*sys__io__ProcessOutput"
			}, [], ["*string"], [
				GoStmt.GoIf(GoExpr.GoRaw("self == nil || self.impl == nil"), [
					GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoStringLiteral("")]))
				],
					null),
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoSelector(GoExpr.GoIdent("self"), "impl"), "ReadLine"), []))
			]),
			GoDecl.GoFuncDecl("close", {
				name: "self",
				typeName: "*sys__io__Process"
			}, [], [], [
				GoStmt.GoIf(GoExpr.GoRaw("self == nil || self.impl == nil"), [GoStmt.GoReturn(null)], null),
				GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoSelector(GoExpr.GoIdent("self"), "impl"), "Close"), []))
			])
		];

		if (requiresSysCommandSurface) {
			decls = decls.concat([
				GoDecl.GoFuncDecl("Sys_command", null, [
					{
						name: "command",
						typeName: "*string"
					},
					{name: "args", typeName: "[]*string"}
				], ["int"], [
					GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("hxrt"), "SysCommand"), [GoExpr.GoIdent("command"), GoExpr.GoIdent("args")]))
				]),
				GoDecl.GoFuncDecl("Sys_exit", null, [
					{
						name: "code",
						typeName: "int"
					}
				], [], [
					GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("hxrt"), "SysExit"), [GoExpr.GoIdent("code")]))
				])
			]);
		}

		decls = decls.concat([
			GoDecl.GoFuncDecl("get_bigEndian", {
				name: "self",
				typeName: "*sys__io__FileInput"
			}, [], ["bool"], [
				GoStmt.GoIf(GoExpr.GoBinary("==", GoExpr.GoIdent("self"), GoExpr.GoNil), [GoStmt.GoReturn(GoExpr.GoBoolLiteral(false))], null),
				GoStmt.GoReturn(GoExpr.GoSelector(GoExpr.GoIdent("self"), "__hx_io_bigEndian"))
			]),
			GoDecl.GoFuncDecl("set_bigEndian", {
				name: "self",
				typeName: "*sys__io__FileInput"
			}, [{name: "e", typeName: "bool"}], ["bool"], [
				GoStmt.GoIf(GoExpr.GoBinary("!=", GoExpr.GoIdent("self"), GoExpr.GoNil), [
					GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("self"), "__hx_io_bigEndian"), GoExpr.GoIdent("e"))
				], null),
				GoStmt.GoReturn(GoExpr.GoIdent("e"))
			]),
			GoDecl.GoFuncDecl("close", {
				name: "self",
				typeName: "*sys__io__FileInput"
			}, [], [], [
				GoStmt.GoRaw("impl := sys__io__fileInputHandles[self]"),
				GoStmt.GoRaw("if impl == nil {"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if err := impl.Close(); err != nil {"),
				GoStmt.GoRaw("\thxrt.Throw(hxrt.StringFromLiteral(err.Error()))"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("delete(sys__io__fileInputHandles, self)")
			]),
			GoDecl.GoFuncDecl("readByte", {
				name: "self",
				typeName: "*sys__io__FileInput"
			}, [], ["int"], [
				GoStmt.GoRaw("impl := sys__io__fileInputHandles[self]"),
				GoStmt.GoRaw("if impl == nil {"),
				GoStmt.GoRaw("\thxrt.Throw(&haxe__io__Eof{})"),
				GoStmt.GoRaw("\treturn 0"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("value, eof, err := impl.ReadByte()"),
				GoStmt.GoRaw("if err != nil {"),
				GoStmt.GoRaw("\thxrt.Throw(hxrt.StringFromLiteral(err.Error()))"),
				GoStmt.GoRaw("\treturn 0"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if eof {"),
				GoStmt.GoRaw("\thxrt.Throw(&haxe__io__Eof{})"),
				GoStmt.GoRaw("\treturn 0"),
				GoStmt.GoRaw("}"),
				GoStmt.GoReturn(GoExpr.GoIdent("value"))
			]),
			GoDecl.GoFuncDecl("readBytes", {
				name: "self",
				typeName: "*sys__io__FileInput"
			}, [
				{name: "buf", typeName: "*haxe__io__Bytes"},
				{name: "pos", typeName: "int"},
				{name: "len", typeName: "int"}
			], ["int"], [
				GoStmt.GoRaw("if buf == nil || pos < 0 || len < 0 || pos+len > buf.length {"),
				GoStmt.GoRaw("\thxrt.Throw(haxe__io__Error_OutsideBounds)"),
				GoStmt.GoRaw("\treturn 0"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if self == nil {"),
				GoStmt.GoRaw("\thxrt.Throw(haxe__io__Error_Blocked)"),
				GoStmt.GoRaw("\treturn 0"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("k := 0"),
				GoStmt.GoRaw("for k < len {"),
				GoStmt.GoRaw("\tvalue := 0"),
				GoStmt.GoRaw("\tthrew := false"),
				GoStmt.GoRaw("\tvar thrown any"),
				GoStmt.GoRaw("\tfunc() {"),
				GoStmt.GoRaw("\t\tdefer func() {"),
				GoStmt.GoRaw("\t\t\tif recovered := recover(); recovered != nil {"),
				GoStmt.GoRaw("\t\t\t\tthrew = true"),
				GoStmt.GoRaw("\t\t\t\tthrown = hxrt.UnwrapException(recovered)"),
				GoStmt.GoRaw("\t\t\t}"),
				GoStmt.GoRaw("\t\t}()"),
				GoStmt.GoRaw("\t\tvalue = self.readByte()"),
				GoStmt.GoRaw("\t}()"),
				GoStmt.GoRaw("\tif threw {"),
				GoStmt.GoRaw("\t\tif haxe__io__input_isEof(thrown) {"),
				GoStmt.GoRaw("\t\t\tif k > 0 {"),
				GoStmt.GoRaw("\t\t\t\treturn k"),
				GoStmt.GoRaw("\t\t\t}"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t\thxrt.Throw(thrown)"),
				GoStmt.GoRaw("\t\treturn 0"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tbuf.b[pos+k] = value"),
				GoStmt.GoRaw("\tk++"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("return len")
			]),
			GoDecl.GoFuncDecl("get_bigEndian", {
				name: "self",
				typeName: "*sys__io__FileOutput"
			}, [], ["bool"], [
				GoStmt.GoIf(GoExpr.GoBinary("==", GoExpr.GoIdent("self"), GoExpr.GoNil), [GoStmt.GoReturn(GoExpr.GoBoolLiteral(false))], null),
				GoStmt.GoReturn(GoExpr.GoSelector(GoExpr.GoIdent("self"), "__hx_io_bigEndian"))
			]),
			GoDecl.GoFuncDecl("set_bigEndian", {
				name: "self",
				typeName: "*sys__io__FileOutput"
			}, [{name: "e", typeName: "bool"}], ["bool"], [
				GoStmt.GoIf(GoExpr.GoBinary("!=", GoExpr.GoIdent("self"), GoExpr.GoNil), [
					GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("self"), "__hx_io_bigEndian"), GoExpr.GoIdent("e"))
				], null),
				GoStmt.GoReturn(GoExpr.GoIdent("e"))
			]),
			GoDecl.GoFuncDecl("flush", {
				name: "self",
				typeName: "*sys__io__FileOutput"
			}, [], [], [
				GoStmt.GoRaw("impl := sys__io__fileOutputHandles[self]"),
				GoStmt.GoRaw("if impl == nil {"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if err := impl.Flush(); err != nil {"),
				GoStmt.GoRaw("\thxrt.Throw(hxrt.StringFromLiteral(err.Error()))"),
				GoStmt.GoRaw("}")
			]),
			GoDecl.GoFuncDecl("close", {
				name: "self",
				typeName: "*sys__io__FileOutput"
			}, [], [], [
				GoStmt.GoRaw("impl := sys__io__fileOutputHandles[self]"),
				GoStmt.GoRaw("if impl == nil {"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if err := impl.Close(); err != nil {"),
				GoStmt.GoRaw("\thxrt.Throw(hxrt.StringFromLiteral(err.Error()))"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("delete(sys__io__fileOutputHandles, self)")
			]),
			GoDecl.GoFuncDecl("prepare", {
				name: "self",
				typeName: "*sys__io__FileOutput"
			}, [{name: "nbytes", typeName: "int"}], [], [
				GoStmt.GoAssign(GoExpr.GoIdent("_"), GoExpr.GoIdent("self")),
				GoStmt.GoAssign(GoExpr.GoIdent("_"), GoExpr.GoIdent("nbytes"))
			]),
			GoDecl.GoFuncDecl("writeByte", {
				name: "self",
				typeName: "*sys__io__FileOutput"
			}, [{name: "c", typeName: "int"}], [], [
				GoStmt.GoRaw("impl := sys__io__fileOutputHandles[self]"),
				GoStmt.GoRaw("if impl == nil {"),
				GoStmt.GoRaw("\thxrt.Throw(hxrt.StringFromLiteral(\"FileOutput is closed\"))"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if err := impl.WriteByte(c); err != nil {"),
				GoStmt.GoRaw("\thxrt.Throw(hxrt.StringFromLiteral(err.Error()))"),
				GoStmt.GoRaw("}")
			]),
			GoDecl.GoFuncDecl("writeBytes", {
				name: "self",
				typeName: "*sys__io__FileOutput"
			}, [
				{name: "s", typeName: "*haxe__io__Bytes"},
				{name: "pos", typeName: "int"},
				{name: "len", typeName: "int"}
			], ["int"], [
				GoStmt.GoRaw("if s == nil || pos < 0 || len < 0 || pos+len > s.length {"),
				GoStmt.GoRaw("\thxrt.Throw(haxe__io__Error_OutsideBounds)"),
				GoStmt.GoRaw("\treturn 0"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("n := len"),
				GoStmt.GoRaw("for len > 0 {"),
				GoStmt.GoRaw("\tself.writeByte(s.b[pos])"),
				GoStmt.GoRaw("\tpos++"),
				GoStmt.GoRaw("\tlen--"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("return n")
			])
		]);

		for (methodName in [
			"readAll",
			"readFullBytes",
			"read",
			"readUntil",
			"readLine",
			"readFloat",
			"readDouble",
			"readInt8",
			"readInt16",
			"readUInt16",
			"readInt24",
			"readUInt24",
			"readInt32",
			"readString"
		]) {
			decls.push(lowerIoInputSyntheticHelper("*sys__io__FileInput", methodName));
		}

		for (methodName in [
			"write",
			"writeFullBytes",
			"writeFloat",
			"writeDouble",
			"writeInt8",
			"writeInt16",
			"writeUInt16",
			"writeInt24",
			"writeUInt24",
			"writeInt32",
			"writeInput",
			"writeString"
		]) {
			decls.push(lowerIoOutputSyntheticHelper("*sys__io__FileOutput", methodName));
		}

		if (requiredStdlibShimGroups.exists("ds")) {
			decls.push(GoDecl.GoFuncDecl("Sys_environment", null, [], ["*haxe__ds__StringMap"], [
				GoStmt.GoRaw("env := New_haxe__ds__StringMap()"),
				GoStmt.GoRaw("for key, value := range hxrt.SysEnvironment() {"),
				GoStmt.GoRaw("\tenv.h[key] = hxrt.StringFromLiteral(value)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoReturn(GoExpr.GoIdent("env"))
			]));
		}

		return decls;
	}

	function lowerFileSystemShimDecls():Array<GoDecl> {
		return [
			GoDecl.GoFuncDecl("sys__FileSystem_exists", null, [{name: "path", typeName: "*string"}], ["bool"], [
				GoStmt.GoRaw("_, err := os.Stat(*hxrt.StdString(path))"),
				GoStmt.GoReturn(GoExpr.GoRaw("err == nil"))
			]),
			GoDecl.GoFuncDecl("sys__FileSystem_rename", null, [
				{
					name: "path",
					typeName: "*string"
				},
				{name: "newPath", typeName: "*string"}
			], [], [
				GoStmt.GoRaw("if err := os.Rename(*hxrt.StdString(path), *hxrt.StdString(newPath)); err != nil {"),
				GoStmt.GoRaw("\thxrt.Throw(err)"),
				GoStmt.GoRaw("}")
			]),
			GoDecl.GoFuncDecl("sys__FileSystem_stat", null, [
				{
					name: "path",
					typeName: "*string"
				}
			], ["map[string]any"], [
				GoStmt.GoRaw("info, err := os.Stat(*hxrt.StdString(path))"),
				GoStmt.GoIf(GoExpr.GoBinary("!=", GoExpr.GoIdent("err"), GoExpr.GoNil), [
					GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("hxrt.Throw"), [GoExpr.GoIdent("err")])),
					GoStmt.GoReturn(GoExpr.GoRaw("map[string]any{}"))
				],
					null),
				GoStmt.GoVarDecl("modTime", null, GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("info"), "ModTime"), []), true),
				GoStmt.GoVarDecl("timeValue", null, GoExpr.GoRaw("&Date{value: modTime}"), true),
				GoStmt.GoReturn(GoExpr.GoRaw("map[string]any{\"gid\": 0, \"uid\": 0, \"atime\": timeValue, \"mtime\": timeValue, \"ctime\": timeValue, \"dev\": 0, \"ino\": 0, \"nlink\": 1, \"rdev\": 0, \"size\": int(info.Size()), \"mode\": int(info.Mode())}"))
			]),
			GoDecl.GoFuncDecl("sys__FileSystem_fullPath", null, [
				{
					name: "path",
					typeName: "*string"
				}
			], ["*string"], [
				GoStmt.GoRaw("resolved, err := filepath.Abs(*hxrt.StdString(path))"),
				GoStmt.GoIf(GoExpr.GoBinary("!=", GoExpr.GoIdent("err"), GoExpr.GoNil), [
					GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("hxrt.Throw"), [GoExpr.GoIdent("err")])),
					GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoStringLiteral("")]))
				], null),
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [
					GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("filepath"), "ToSlash"), [GoExpr.GoIdent("resolved")])
				]))
			]),
			GoDecl.GoFuncDecl("sys__FileSystem_isDirectory", null, [
				{
					name: "path",
					typeName: "*string"
				}
			], ["bool"], [
				GoStmt.GoRaw("info, err := os.Stat(*hxrt.StdString(path))"),
				GoStmt.GoIf(GoExpr.GoBinary("!=", GoExpr.GoIdent("err"), GoExpr.GoNil), [GoStmt.GoReturn(GoExpr.GoBoolLiteral(false))], null),
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("info"), "IsDir"), []))
			]),
			GoDecl.GoFuncDecl("sys__FileSystem_createDirectory", null, [
				{
					name: "path",
					typeName: "*string"
				}
			], [], [
				GoStmt.GoRaw("if err := os.MkdirAll(*hxrt.StdString(path), 0o755); err != nil {"),
				GoStmt.GoRaw("\thxrt.Throw(err)"),
				GoStmt.GoRaw("}")
			]),
			GoDecl.GoFuncDecl("sys__FileSystem_deleteFile", null, [
				{
					name: "path",
					typeName: "*string"
				}
			], [], [
				GoStmt.GoRaw("if err := os.Remove(*hxrt.StdString(path)); err != nil {"),
				GoStmt.GoRaw("\thxrt.Throw(err)"),
				GoStmt.GoRaw("}")
			]),
			GoDecl.GoFuncDecl("sys__FileSystem_deleteDirectory", null, [
				{
					name: "path",
					typeName: "*string"
				}
			], [], [
				GoStmt.GoRaw("if err := os.Remove(*hxrt.StdString(path)); err != nil {"),
				GoStmt.GoRaw("\thxrt.Throw(err)"),
				GoStmt.GoRaw("}")
			]),
			GoDecl.GoFuncDecl("sys__FileSystem_readDirectory", null, [
				{
					name: "path",
					typeName: "*string"
				}
			], ["[]*string"], [
				GoStmt.GoRaw("entries, err := os.ReadDir(*hxrt.StdString(path))"),
				GoStmt.GoIf(GoExpr.GoBinary("!=", GoExpr.GoIdent("err"), GoExpr.GoNil), [
					GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("hxrt.Throw"), [GoExpr.GoIdent("err")])),
					GoStmt.GoReturn(GoExpr.GoRaw("[]*string{}"))
				],
					null),
				GoStmt.GoVarDecl("out", null, GoExpr.GoRaw("make([]*string, 0, len(entries))"), true),
				GoStmt.GoRaw("for _, entry := range entries {"),
				GoStmt.GoRaw("\tout = append(out, hxrt.StringFromLiteral(entry.Name()))"),
				GoStmt.GoRaw("}"),
				GoStmt.GoReturn(GoExpr.GoIdent("out"))
			])
		];
	}

	function lowerStdlibSymbolShimDecls():Array<GoDecl> {
		var decls = [
			GoDecl.GoStructDecl("Std", []),
			GoDecl.GoFuncDecl("_UnicodeString__UnicodeString_Impl__get_length", null, [
				{
					name: "value",
					typeName: "any"
				}
			],
				["int"], [GoStmt.GoReturn(GoExpr.GoRaw("len([]rune(*hxrt.StdString(value)))"))]),
			GoDecl.GoFuncDecl("_UnicodeString__UnicodeString_Impl__charAt", null, [
				{
					name: "value",
					typeName: "any"
				},
				{name: "index", typeName: "int"}
			], ["*string"], [
				GoStmt.GoRaw("if index < 0 {"),
				GoStmt.GoRaw("\treturn hxrt.StringFromLiteral(\"\")"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("runes := []rune(*hxrt.StdString(value))"),
				GoStmt.GoRaw("if index >= len(runes) {"),
				GoStmt.GoRaw("\treturn hxrt.StringFromLiteral(\"\")"),
				GoStmt.GoRaw("}"),
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoRaw("string(runes[index])")]))
			]),
			GoDecl.GoFuncDecl("_UnicodeString__UnicodeString_Impl__charCodeAt", null, [
				{
					name: "value",
					typeName: "any"
				},
				{name: "index", typeName: "int"}
			], ["any"], [
				GoStmt.GoRaw("if index < 0 {"),
				GoStmt.GoRaw("\treturn nil"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("runes := []rune(*hxrt.StdString(value))"),
				GoStmt.GoRaw("if index >= len(runes) {"),
				GoStmt.GoRaw("\treturn nil"),
				GoStmt.GoRaw("}"),
				GoStmt.GoReturn(GoExpr.GoRaw("int(runes[index])"))
			]),
			GoDecl.GoFuncDecl("_UnicodeString__UnicodeString_Impl__substring", null, [
				{
					name: "value",
					typeName: "any"
				},
				{name: "startIndex", typeName: "int"},
				{name: "endIndex", typeName: "...int"}
			], ["*string"], [
				GoStmt.GoRaw("runes := []rune(*hxrt.StdString(value))"),
				GoStmt.GoRaw("end := len(runes)"),
				GoStmt.GoRaw("if len(endIndex) > 0 {"),
				GoStmt.GoRaw("\tend = endIndex[0]"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if startIndex < 0 {"),
				GoStmt.GoRaw("\tstartIndex = 0"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if end < 0 {"),
				GoStmt.GoRaw("\tend = 0"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if startIndex == end {"),
				GoStmt.GoRaw("\treturn hxrt.StringFromLiteral(\"\")"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if startIndex > end {"),
				GoStmt.GoRaw("\tstartIndex, end = end, startIndex"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if startIndex > len(runes) {"),
				GoStmt.GoRaw("\treturn hxrt.StringFromLiteral(\"\")"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if end > len(runes) {"),
				GoStmt.GoRaw("\tend = len(runes)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoRaw("string(runes[startIndex:end])")]))
			]),
			GoDecl.GoFuncDecl("_UnicodeString__UnicodeString_Impl__substr", null, [
				{
					name: "value",
					typeName: "any"
				},
				{name: "pos", typeName: "int"},
				{name: "lengthArgs", typeName: "...int"}
			], ["*string"], [
				GoStmt.GoRaw("runes := []rune(*hxrt.StdString(value))"),
				GoStmt.GoRaw("unicodeLength := len(runes)"),
				GoStmt.GoRaw("if pos < 0 {"),
				GoStmt.GoRaw("\tpos = unicodeLength + pos"),
				GoStmt.GoRaw("\tif pos < 0 {"),
				GoStmt.GoRaw("\t\tpos = 0"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if pos > unicodeLength {"),
				GoStmt.GoRaw("\treturn hxrt.StringFromLiteral(\"\")"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if len(lengthArgs) == 0 {"),
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoRaw("string(runes[pos:])")])),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("lengthValue := lengthArgs[0]"),
				GoStmt.GoRaw("end := unicodeLength"),
				GoStmt.GoRaw("if lengthValue < 0 {"),
				GoStmt.GoRaw("\tend = unicodeLength + lengthValue"),
				GoStmt.GoRaw("} else {"),
				GoStmt.GoRaw("\tend = pos + lengthValue"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if end < pos {"),
				GoStmt.GoRaw("\treturn hxrt.StringFromLiteral(\"\")"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if end > unicodeLength {"),
				GoStmt.GoRaw("\tend = unicodeLength"),
				GoStmt.GoRaw("}"),
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoRaw("string(runes[pos:end])")]))
			]),
			GoDecl.GoFuncDecl("_UnicodeString__UnicodeString_Impl__indexOf", null, [
				{
					name: "value",
					typeName: "any"
				},
				{name: "str", typeName: "*string"},
				{name: "startIndex", typeName: "...int"}
			], ["int"], [
				GoStmt.GoRaw("runes := []rune(*hxrt.StdString(value))"),
				GoStmt.GoRaw("needle := []rune(*hxrt.StdString(str))"),
				GoStmt.GoRaw("start := 0"),
				GoStmt.GoRaw("if len(startIndex) > 0 {"),
				GoStmt.GoRaw("\tstart = startIndex[0]"),
				GoStmt.GoRaw("\tif start < 0 {"),
				GoStmt.GoRaw("\t\tstart = len(runes) + start"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if start < 0 {"),
				GoStmt.GoRaw("\tstart = 0"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if len(needle) == 0 {"),
				GoStmt.GoRaw("\tif start > len(runes) {"),
				GoStmt.GoRaw("\t\treturn len(runes)"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn start"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if start > len(runes) || len(needle) > len(runes) {"),
				GoStmt.GoRaw("\treturn -1"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("for i := start; i+len(needle) <= len(runes); i++ {"),
				GoStmt.GoRaw("\tmatched := true"),
				GoStmt.GoRaw("\tfor j := 0; j < len(needle); j++ {"),
				GoStmt.GoRaw("\t\tif runes[i+j] != needle[j] {"),
				GoStmt.GoRaw("\t\t\tmatched = false"),
				GoStmt.GoRaw("\t\t\tbreak"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tif matched {"),
				GoStmt.GoRaw("\t\treturn i"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("}"),
				GoStmt.GoReturn(GoExpr.GoIntLiteral(-1))
			]),
			GoDecl.GoFuncDecl("_UnicodeString__UnicodeString_Impl__lastIndexOf", null, [
				{
					name: "value",
					typeName: "any"
				},
				{name: "str", typeName: "*string"},
				{name: "startIndex", typeName: "...int"}
			], ["int"], [
				GoStmt.GoRaw("runes := []rune(*hxrt.StdString(value))"),
				GoStmt.GoRaw("needle := []rune(*hxrt.StdString(str))"),
				GoStmt.GoRaw("if len(needle) == 0 {"),
				GoStmt.GoRaw("\tif len(startIndex) == 0 {"),
				GoStmt.GoRaw("\t\treturn len(runes)"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tstart := startIndex[0]"),
				GoStmt.GoRaw("\tif start < 0 {"),
				GoStmt.GoRaw("\t\tstart = 0"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tif start > len(runes) {"),
				GoStmt.GoRaw("\t\treturn len(runes)"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn start"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("start := len(runes)"),
				GoStmt.GoRaw("if len(startIndex) > 0 {"),
				GoStmt.GoRaw("\tstart = startIndex[0]"),
				GoStmt.GoRaw("\tif start < 0 {"),
				GoStmt.GoRaw("\t\tstart = 0"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("limit := start + len(needle)"),
				GoStmt.GoRaw("if limit > len(runes) {"),
				GoStmt.GoRaw("\tlimit = len(runes)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("for i := limit - len(needle); i >= 0; i-- {"),
				GoStmt.GoRaw("\tmatched := true"),
				GoStmt.GoRaw("\tfor j := 0; j < len(needle); j++ {"),
				GoStmt.GoRaw("\t\tif runes[i+j] != needle[j] {"),
				GoStmt.GoRaw("\t\t\tmatched = false"),
				GoStmt.GoRaw("\t\t\tbreak"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tif matched {"),
				GoStmt.GoRaw("\t\treturn i"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("}"),
				GoStmt.GoReturn(GoExpr.GoIntLiteral(-1))
			]),
			GoDecl.GoFuncDecl("_UnicodeString__UnicodeString_Impl__validate", null, [
				{
					name: "value",
					typeName: "*haxe__io__Bytes"
				},
				{name: "encoding", typeName: "*haxe__io__Encoding"}
			], ["bool"], [
				GoStmt.GoRaw("if haxe__io__resolveEncodingTag(encoding) == 1 {"),
				GoStmt.GoRaw("\thxrt.Throw(hxrt.StringFromLiteral(\"UnicodeString.validate: RawNative encoding is not supported\"))"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("raw := hxrt_haxeBytesToRaw(value)"),
				GoStmt.GoRaw("pos := 0"),
				GoStmt.GoRaw("max := len(raw)"),
				GoStmt.GoRaw("for pos < max {"),
				GoStmt.GoRaw("\tc := int(raw[pos])"),
				GoStmt.GoRaw("\tpos++"),
				GoStmt.GoRaw("\tif c < 0x80 {"),
				GoStmt.GoRaw("\t\tcontinue"),
				GoStmt.GoRaw("\t} else if c < 0xC2 {"),
				GoStmt.GoRaw("\t\treturn false"),
				GoStmt.GoRaw("\t} else if c < 0xE0 {"),
				GoStmt.GoRaw("\t\tif pos+1 > max {"),
				GoStmt.GoRaw("\t\t\treturn false"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t\tc2 := int(raw[pos])"),
				GoStmt.GoRaw("\t\tpos++"),
				GoStmt.GoRaw("\t\tif c2 < 0x80 || c2 > 0xBF {"),
				GoStmt.GoRaw("\t\t\treturn false"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t} else if c < 0xF0 {"),
				GoStmt.GoRaw("\t\tif pos+2 > max {"),
				GoStmt.GoRaw("\t\t\treturn false"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t\tc2 := int(raw[pos])"),
				GoStmt.GoRaw("\t\tpos++"),
				GoStmt.GoRaw("\t\tif c == 0xE0 {"),
				GoStmt.GoRaw("\t\t\tif c2 < 0xA0 || c2 > 0xBF {"),
				GoStmt.GoRaw("\t\t\t\treturn false"),
				GoStmt.GoRaw("\t\t\t}"),
				GoStmt.GoRaw("\t\t} else if c2 < 0x80 || c2 > 0xBF {"),
				GoStmt.GoRaw("\t\t\treturn false"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t\tc3 := int(raw[pos])"),
				GoStmt.GoRaw("\t\tpos++"),
				GoStmt.GoRaw("\t\tif c3 < 0x80 || c3 > 0xBF {"),
				GoStmt.GoRaw("\t\t\treturn false"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t\tc = (c << 16) | (c2 << 8) | c3"),
				GoStmt.GoRaw("\t\tif 0xEDA080 <= c && c <= 0xEDBFBF {"),
				GoStmt.GoRaw("\t\t\treturn false"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t} else if c > 0xF4 {"),
				GoStmt.GoRaw("\t\treturn false"),
				GoStmt.GoRaw("\t} else {"),
				GoStmt.GoRaw("\t\tif pos+3 > max {"),
				GoStmt.GoRaw("\t\t\treturn false"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t\tc2 := int(raw[pos])"),
				GoStmt.GoRaw("\t\tpos++"),
				GoStmt.GoRaw("\t\tif c == 0xF0 {"),
				GoStmt.GoRaw("\t\t\tif c2 < 0x90 || c2 > 0xBF {"),
				GoStmt.GoRaw("\t\t\t\treturn false"),
				GoStmt.GoRaw("\t\t\t}"),
				GoStmt.GoRaw("\t\t} else if c == 0xF4 {"),
				GoStmt.GoRaw("\t\t\tif c2 < 0x80 || c2 > 0x8F {"),
				GoStmt.GoRaw("\t\t\t\treturn false"),
				GoStmt.GoRaw("\t\t\t}"),
				GoStmt.GoRaw("\t\t} else if c2 < 0x80 || c2 > 0xBF {"),
				GoStmt.GoRaw("\t\t\treturn false"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t\tc3 := int(raw[pos])"),
				GoStmt.GoRaw("\t\tpos++"),
				GoStmt.GoRaw("\t\tif c3 < 0x80 || c3 > 0xBF {"),
				GoStmt.GoRaw("\t\t\treturn false"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t\tc4 := int(raw[pos])"),
				GoStmt.GoRaw("\t\tpos++"),
				GoStmt.GoRaw("\t\tif c4 < 0x80 || c4 > 0xBF {"),
				GoStmt.GoRaw("\t\t\treturn false"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("}"),
				GoStmt.GoReturn(GoExpr.GoBoolLiteral(true))
			]),
			GoDecl.GoStructDecl("Date", [
				{
					name: "value",
					typeName: "time.Time"
				}
			]),
			GoDecl.GoFuncDecl("Date_fromString", null, [{name: "source", typeName: "*string"}], ["*Date"], [
				GoStmt.GoVarDecl("raw", null, GoExpr.GoRaw("*hxrt.StdString(source)"), true),
				GoStmt.GoRaw("parsed, err := time.ParseInLocation(\"2006-01-02 15:04:05\", raw, time.Local)"),
				GoStmt.GoIf(GoExpr.GoBinary("!=", GoExpr.GoIdent("err"), GoExpr.GoNil), [
					GoStmt.GoRaw("parsedDateOnly, errDateOnly := time.ParseInLocation(\"2006-01-02\", raw, time.Local)"),
					GoStmt.GoIf(GoExpr.GoBinary("==", GoExpr.GoIdent("errDateOnly"), GoExpr.GoNil),
						[GoStmt.GoAssign(GoExpr.GoIdent("parsed"), GoExpr.GoIdent("parsedDateOnly"))],
						[GoStmt.GoAssign(GoExpr.GoIdent("parsed"), GoExpr.GoRaw("time.Unix(0, 0)"))])
				],
					null),
				GoStmt.GoReturn(GoExpr.GoRaw("&Date{value: parsed}"))
			]),
			GoDecl.GoFuncDecl("Date_now", null, [], ["*Date"], [GoStmt.GoReturn(GoExpr.GoRaw("&Date{value: time.Now()}"))]),
			GoDecl.GoFuncDecl("Date_fromTime", null, [
				{
					name: "ms",
					typeName: "float64"
				}
			], ["*Date"], [
				GoStmt.GoRaw("nanos := int64(ms * 1e6)"),
				GoStmt.GoReturn(GoExpr.GoRaw("&Date{value: time.Unix(0, nanos).In(time.Local)}"))
			]),
			GoDecl.GoFuncDecl("getFullYear", {
				name: "self",
				typeName: "*Date"
			}, [], ["int"], [
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoSelector(GoExpr.GoIdent("self"), "value"), "Year"), []))
			]),
			GoDecl.GoFuncDecl("getMonth", {
				name: "self",
				typeName: "*Date"
			},
				[], ["int"], [GoStmt.GoReturn(GoExpr.GoRaw("int(self.value.Month()) - 1"))]),
			GoDecl.GoFuncDecl("getDate", {name: "self", typeName: "*Date"}, [], ["int"], [
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoSelector(GoExpr.GoIdent("self"), "value"), "Day"), []))
			]),
			GoDecl.GoFuncDecl("getDay", {
				name: "self",
				typeName: "*Date"
			}, [], ["int"],
				[GoStmt.GoReturn(GoExpr.GoRaw("int(self.value.Weekday())"))]),
			GoDecl.GoFuncDecl("getHours", {
				name: "self",
				typeName: "*Date"
			}, [], ["int"], [
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoSelector(GoExpr.GoIdent("self"), "value"), "Hour"), []))
			]),
			GoDecl.GoFuncDecl("getMinutes", {
				name: "self",
				typeName: "*Date"
			}, [], ["int"], [
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoSelector(GoExpr.GoIdent("self"), "value"), "Minute"), []))
			]),
			GoDecl.GoFuncDecl("getSeconds", {
				name: "self",
				typeName: "*Date"
			}, [], ["int"], [
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoSelector(GoExpr.GoIdent("self"), "value"), "Second"), []))
			]),
			GoDecl.GoFuncDecl("getTime", {
				name: "self",
				typeName: "*Date"
			},
				[], ["float64"], [GoStmt.GoReturn(GoExpr.GoRaw("float64(self.value.UnixNano()) / 1e6"))]),
			GoDecl.GoStructDecl("Math", []),
			GoDecl.GoFuncDecl("Math_floor", null, [
				{
					name: "value",
					typeName: "float64"
				}
			],
				["int"], [GoStmt.GoReturn(GoExpr.GoRaw("int(math.Floor(value))"))]),
			GoDecl.GoFuncDecl("Math_ceil", null, [{name: "value", typeName: "float64"}], ["int"], [GoStmt.GoReturn(GoExpr.GoRaw("int(math.Ceil(value))"))]),
			GoDecl.GoFuncDecl("Math_round", null, [{name: "value", typeName: "float64"}], ["int"],
				[GoStmt.GoReturn(GoExpr.GoRaw("int(math.Floor(value + 0.5))"))]),
			GoDecl.GoFuncDecl("Math_abs", null, [{name: "value", typeName: "float64"}], ["float64"], [
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("math"), "Abs"), [GoExpr.GoIdent("value")]))
			]),
			GoDecl.GoFuncDecl("Math_isNaN", null, [
				{
					name: "value",
					typeName: "float64"
				}
			], ["bool"], [
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("math"), "IsNaN"), [GoExpr.GoIdent("value")]))
			]),
			GoDecl.GoFuncDecl("Math_isFinite", null, [
				{
					name: "value",
					typeName: "float64"
				}
			], ["bool"], [
				GoStmt.GoReturn(GoExpr.GoUnary("!",
					GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("math"), "IsInf"), [GoExpr.GoIdent("value"), GoExpr.GoIntLiteral(0)])))
			]),
			GoDecl.GoFuncDecl("Math_min", null, [
				{
					name: "a",
					typeName: "float64"
				},
				{name: "b", typeName: "float64"}
			], ["float64"], [
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("math"), "Min"), [GoExpr.GoIdent("a"), GoExpr.GoIdent("b")]))
			]),
			GoDecl.GoFuncDecl("Math_max", null, [
				{
					name: "a",
					typeName: "float64"
				},
				{name: "b", typeName: "float64"}
			], ["float64"], [
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("math"), "Max"), [GoExpr.GoIdent("a"), GoExpr.GoIdent("b")]))
			]),
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
			], [],
				[
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
			GoDecl.GoGlobalVarDecl("Xml_Element", "int", GoExpr.GoIntLiteral(0)),
			GoDecl.GoGlobalVarDecl("Xml_PCData", "int", GoExpr.GoIntLiteral(1)),
			GoDecl.GoGlobalVarDecl("Xml_CData", "int", GoExpr.GoIntLiteral(2)),
			GoDecl.GoGlobalVarDecl("Xml_Comment", "int", GoExpr.GoIntLiteral(3)),
			GoDecl.GoGlobalVarDecl("Xml_DocType", "int", GoExpr.GoIntLiteral(4)),
			GoDecl.GoGlobalVarDecl("Xml_ProcessingInstruction", "int", GoExpr.GoIntLiteral(5)),
			GoDecl.GoGlobalVarDecl("Xml_Document", "int", GoExpr.GoIntLiteral(6)),
			GoDecl.GoStructDecl("Xml", [
				{
					name: "nodeType",
					typeName: "int"
				},
				{name: "nodeName", typeName: "*string"},
				{name: "nodeValue", typeName: "*string"},
				{name: "parent", typeName: "*Xml"},
				{name: "children", typeName: "[]*Xml"},
				{name: "attributeMap", typeName: "map[string]string"},
				{name: "attributeOrder", typeName: "[]string"}
			]),
			GoDecl.GoFuncDecl("_Xml__XmlType_Impl__toString", null, [{name: "value", typeName: "int"}], ["*string"], [
				GoStmt.GoRaw("switch value {"),
				GoStmt.GoRaw("case Xml_Element:"),
				GoStmt.GoRaw("\treturn hxrt.StringFromLiteral(\"Element\")"),
				GoStmt.GoRaw("case Xml_PCData:"),
				GoStmt.GoRaw("\treturn hxrt.StringFromLiteral(\"PCData\")"),
				GoStmt.GoRaw("case Xml_CData:"),
				GoStmt.GoRaw("\treturn hxrt.StringFromLiteral(\"CData\")"),
				GoStmt.GoRaw("case Xml_Comment:"),
				GoStmt.GoRaw("\treturn hxrt.StringFromLiteral(\"Comment\")"),
				GoStmt.GoRaw("case Xml_DocType:"),
				GoStmt.GoRaw("\treturn hxrt.StringFromLiteral(\"DocType\")"),
				GoStmt.GoRaw("case Xml_ProcessingInstruction:"),
				GoStmt.GoRaw("\treturn hxrt.StringFromLiteral(\"ProcessingInstruction\")"),
				GoStmt.GoRaw("case Xml_Document:"),
				GoStmt.GoRaw("\treturn hxrt.StringFromLiteral(\"Document\")"),
				GoStmt.GoRaw("default:"),
				GoStmt.GoRaw("\treturn hxrt.StringFromLiteral(\"XmlType\")"),
				GoStmt.GoRaw("}")
			]),
			GoDecl.GoFuncDecl("New_Xml", null, [
				{
					name: "nodeType",
					typeName: "int"
				}
			], ["*Xml"], [
				GoStmt.GoReturn(GoExpr.GoRaw("&Xml{nodeType: nodeType, children: []*Xml{}, attributeMap: map[string]string{}, attributeOrder: []string{}}"))
			]),
			GoDecl.GoFuncDecl("Xml_createElement", null, [
				{
					name: "name",
					typeName: "*string"
				}
			], ["*Xml"], [
				GoStmt.GoRaw("xml := New_Xml(Xml_Element)"),
				GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("xml"), "nodeName"), GoExpr.GoIdent("name")),
				GoStmt.GoReturn(GoExpr.GoIdent("xml"))
			]),
			GoDecl.GoFuncDecl("Xml_createPCData", null, [
				{
					name: "data",
					typeName: "*string"
				}
			], ["*Xml"], [
				GoStmt.GoRaw("xml := New_Xml(Xml_PCData)"),
				GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("xml"), "nodeValue"), GoExpr.GoIdent("data")),
				GoStmt.GoReturn(GoExpr.GoIdent("xml"))
			]),
			GoDecl.GoFuncDecl("Xml_createCData", null, [
				{
					name: "data",
					typeName: "*string"
				}
			], ["*Xml"], [
				GoStmt.GoRaw("xml := New_Xml(Xml_CData)"),
				GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("xml"), "nodeValue"), GoExpr.GoIdent("data")),
				GoStmt.GoReturn(GoExpr.GoIdent("xml"))
			]),
			GoDecl.GoFuncDecl("Xml_createComment", null, [
				{
					name: "data",
					typeName: "*string"
				}
			], ["*Xml"], [
				GoStmt.GoRaw("xml := New_Xml(Xml_Comment)"),
				GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("xml"), "nodeValue"), GoExpr.GoIdent("data")),
				GoStmt.GoReturn(GoExpr.GoIdent("xml"))
			]),
			GoDecl.GoFuncDecl("Xml_createDocType", null, [
				{
					name: "data",
					typeName: "*string"
				}
			], ["*Xml"], [
				GoStmt.GoRaw("xml := New_Xml(Xml_DocType)"),
				GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("xml"), "nodeValue"), GoExpr.GoIdent("data")),
				GoStmt.GoReturn(GoExpr.GoIdent("xml"))
			]),
			GoDecl.GoFuncDecl("Xml_createProcessingInstruction", null, [
				{
					name: "data",
					typeName: "*string"
				}
			], ["*Xml"], [
				GoStmt.GoRaw("xml := New_Xml(Xml_ProcessingInstruction)"),
				GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("xml"), "nodeValue"), GoExpr.GoIdent("data")),
				GoStmt.GoReturn(GoExpr.GoIdent("xml"))
			]),
			GoDecl.GoFuncDecl("Xml_createDocument", null, [], ["*Xml"], [
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("New_Xml"), [GoExpr.GoIdent("Xml_Document")]))
			]),
			GoDecl.GoFuncDecl("Xml_ensureElementType", null, [
				{
					name: "self",
					typeName: "*Xml"
				}
			], [], [
				GoStmt.GoRaw("if self.nodeType != Xml_Document && self.nodeType != Xml_Element {"),
				GoStmt.GoRaw("\thxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(\"Bad node type, expected Element or Document but found \"), _Xml__XmlType_Impl__toString(self.nodeType)))"),
				GoStmt.GoRaw("}")
			]),
			GoDecl.GoFuncDecl("Xml_parse", null, [
				{
					name: "source",
					typeName: "*string"
				}
			], ["*Xml"], [
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("haxe__xml__Parser_parse"), [GoExpr.GoIdent("source")]))
			]),
			GoDecl.GoFuncDecl("get", {
				name: "self",
				typeName: "*Xml"
			}, [{name: "att", typeName: "*string"}], ["*string"], [
				GoStmt.GoRaw("if self.nodeType != Xml_Element {"),
				GoStmt.GoRaw("\thxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(\"Bad node type, expected Element but found \"), _Xml__XmlType_Impl__toString(self.nodeType)))"),
				GoStmt.GoRaw("\treturn nil"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("key := *hxrt.StdString(att)"),
				GoStmt.GoRaw("value, ok := self.attributeMap[key]"),
				GoStmt.GoRaw("if !ok {"),
				GoStmt.GoRaw("\treturn nil"),
				GoStmt.GoRaw("}"),
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoIdent("value")]))
			]),
			GoDecl.GoFuncDecl("set", {
				name: "self",
				typeName: "*Xml"
			}, [{name: "att", typeName: "*string"}, {name: "value", typeName: "*string"}],
				[], [
					GoStmt.GoRaw("if self.nodeType != Xml_Element {"),
					GoStmt.GoRaw("\thxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(\"Bad node type, expected Element but found \"), _Xml__XmlType_Impl__toString(self.nodeType)))"),
					GoStmt.GoRaw("\treturn"),
					GoStmt.GoRaw("}"),
					GoStmt.GoRaw("key := *hxrt.StdString(att)"),
					GoStmt.GoRaw("if _, ok := self.attributeMap[key]; !ok {"),
					GoStmt.GoRaw("\tself.attributeOrder = append(self.attributeOrder, key)"),
					GoStmt.GoRaw("}"),
					GoStmt.GoRaw("self.attributeMap[key] = *hxrt.StdString(value)")
				]),
			GoDecl.GoFuncDecl("remove", {
				name: "self",
				typeName: "*Xml"
			}, [{name: "att", typeName: "*string"}], [], [
				GoStmt.GoRaw("if self.nodeType != Xml_Element {"),
				GoStmt.GoRaw("\thxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(\"Bad node type, expected Element but found \"), _Xml__XmlType_Impl__toString(self.nodeType)))"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("key := *hxrt.StdString(att)"),
				GoStmt.GoRaw("delete(self.attributeMap, key)"),
				GoStmt.GoRaw("filtered := make([]string, 0, len(self.attributeOrder))"),
				GoStmt.GoRaw("for _, existing := range self.attributeOrder {"),
				GoStmt.GoRaw("\tif existing != key {"),
				GoStmt.GoRaw("\t\tfiltered = append(filtered, existing)"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("self.attributeOrder = filtered")
			]),
			GoDecl.GoFuncDecl("exists", {
				name: "self",
				typeName: "*Xml"
			}, [{name: "att", typeName: "*string"}], ["bool"], [
				GoStmt.GoRaw("if self.nodeType != Xml_Element {"),
				GoStmt.GoRaw("\thxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(\"Bad node type, expected Element but found \"), _Xml__XmlType_Impl__toString(self.nodeType)))"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("_, ok := self.attributeMap[*hxrt.StdString(att)]"),
				GoStmt.GoReturn(GoExpr.GoIdent("ok"))
			]),
			GoDecl.GoFuncDecl("attributes", {
				name: "self",
				typeName: "*Xml"
			}, [], ["map[string]any"], [
				GoStmt.GoRaw("if self.nodeType != Xml_Element {"),
				GoStmt.GoRaw("\thxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(\"Bad node type, expected Element but found \"), _Xml__XmlType_Impl__toString(self.nodeType)))"),
				GoStmt.GoRaw("\treturn map[string]any{}"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("index := 0"),
				GoStmt.GoRaw("iter := map[string]any{}"),
				GoStmt.GoRaw("iter[\"hasNext\"] = func() bool { return index < len(self.attributeOrder) }"),
				GoStmt.GoRaw("iter[\"next\"] = func() *string { key := self.attributeOrder[index]; index++; return hxrt.StringFromLiteral(key) }"),
				GoStmt.GoReturn(GoExpr.GoIdent("iter"))
			]),
			GoDecl.GoFuncDecl("iterator", {
				name: "self",
				typeName: "*Xml"
			}, [], ["map[string]any"], [
				GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("Xml_ensureElementType"), [GoExpr.GoIdent("self")])),
				GoStmt.GoRaw("index := 0"),
				GoStmt.GoRaw("iter := map[string]any{}"),
				GoStmt.GoRaw("iter[\"hasNext\"] = func() bool { return index < len(self.children) }"),
				GoStmt.GoRaw("iter[\"next\"] = func() *Xml { child := self.children[index]; index++; return child }"),
				GoStmt.GoReturn(GoExpr.GoIdent("iter"))
			]),
			GoDecl.GoFuncDecl("elements", {
				name: "self",
				typeName: "*Xml"
			}, [], ["map[string]any"], [
				GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("Xml_ensureElementType"), [GoExpr.GoIdent("self")])),
				GoStmt.GoRaw("matches := make([]*Xml, 0, len(self.children))"),
				GoStmt.GoRaw("for _, child := range self.children {"),
				GoStmt.GoRaw("\tif child.nodeType == Xml_Element {"),
				GoStmt.GoRaw("\t\tmatches = append(matches, child)"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("index := 0"),
				GoStmt.GoRaw("iter := map[string]any{}"),
				GoStmt.GoRaw("iter[\"hasNext\"] = func() bool { return index < len(matches) }"),
				GoStmt.GoRaw("iter[\"next\"] = func() *Xml { child := matches[index]; index++; return child }"),
				GoStmt.GoReturn(GoExpr.GoIdent("iter"))
			]),
			GoDecl.GoFuncDecl("elementsNamed", {
				name: "self",
				typeName: "*Xml"
			}, [{name: "name", typeName: "*string"}], ["map[string]any"], [
				GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("Xml_ensureElementType"), [GoExpr.GoIdent("self")])),
				GoStmt.GoRaw("wanted := *hxrt.StdString(name)"),
				GoStmt.GoRaw("matches := make([]*Xml, 0, len(self.children))"),
				GoStmt.GoRaw("for _, child := range self.children {"),
				GoStmt.GoRaw("\tif child.nodeType == Xml_Element && *hxrt.StdString(child.nodeName) == wanted {"),
				GoStmt.GoRaw("\t\tmatches = append(matches, child)"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("index := 0"),
				GoStmt.GoRaw("iter := map[string]any{}"),
				GoStmt.GoRaw("iter[\"hasNext\"] = func() bool { return index < len(matches) }"),
				GoStmt.GoRaw("iter[\"next\"] = func() *Xml { child := matches[index]; index++; return child }"),
				GoStmt.GoReturn(GoExpr.GoIdent("iter"))
			]),
			GoDecl.GoFuncDecl("firstChild", {
				name: "self",
				typeName: "*Xml"
			}, [], ["*Xml"], [
				GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("Xml_ensureElementType"), [GoExpr.GoIdent("self")])),
				GoStmt.GoRaw("if len(self.children) == 0 {"),
				GoStmt.GoRaw("\treturn nil"),
				GoStmt.GoRaw("}"),
				GoStmt.GoReturn(GoExpr.GoIndex(GoExpr.GoSelector(GoExpr.GoIdent("self"), "children"), GoExpr.GoIntLiteral(0)))
			]),
			GoDecl.GoFuncDecl("firstElement", {
				name: "self",
				typeName: "*Xml"
			}, [], ["*Xml"], [
				GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("Xml_ensureElementType"), [GoExpr.GoIdent("self")])),
				GoStmt.GoRaw("for _, child := range self.children {"),
				GoStmt.GoRaw("\tif child.nodeType == Xml_Element {"),
				GoStmt.GoRaw("\t\treturn child"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("}"),
				GoStmt.GoReturn(GoExpr.GoNil)
			]),
			GoDecl.GoFuncDecl("addChild", {
				name: "self",
				typeName: "*Xml"
			}, [{name: "x", typeName: "*Xml"}], [], [
				GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("Xml_ensureElementType"), [GoExpr.GoIdent("self")])),
				GoStmt.GoRaw("if x == nil {"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if x.parent != nil {"),
				GoStmt.GoRaw("\tx.parent.removeChild(x)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("self.children = append(self.children, x)"),
				GoStmt.GoRaw("x.parent = self")
			]),
			GoDecl.GoFuncDecl("removeChild", {
				name: "self",
				typeName: "*Xml"
			}, [{name: "x", typeName: "*Xml"}], ["bool"], [
				GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("Xml_ensureElementType"), [GoExpr.GoIdent("self")])),
				GoStmt.GoRaw("for i, child := range self.children {"),
				GoStmt.GoRaw("\tif child == x {"),
				GoStmt.GoRaw("\t\tself.children = append(self.children[:i], self.children[i+1:]...)"),
				GoStmt.GoRaw("\t\tx.parent = nil"),
				GoStmt.GoRaw("\t\treturn true"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("}"),
				GoStmt.GoReturn(GoExpr.GoBoolLiteral(false))
			]),
			GoDecl.GoFuncDecl("insertChild", {
				name: "self",
				typeName: "*Xml"
			}, [{name: "x", typeName: "*Xml"}, {name: "pos", typeName: "int"}], [], [
				GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("Xml_ensureElementType"), [GoExpr.GoIdent("self")])),
				GoStmt.GoRaw("if x == nil {"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if x.parent != nil {"),
				GoStmt.GoRaw("\tx.parent.removeChild(x)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if pos < 0 {"),
				GoStmt.GoRaw("\tpos = 0"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if pos > len(self.children) {"),
				GoStmt.GoRaw("\tpos = len(self.children)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("self.children = append(self.children, nil)"),
				GoStmt.GoRaw("copy(self.children[pos+1:], self.children[pos:])"),
				GoStmt.GoRaw("self.children[pos] = x"),
				GoStmt.GoRaw("x.parent = self")
			]),
			GoDecl.GoFuncDecl("toString", {
				name: "self",
				typeName: "*Xml"
			}, [], ["*string"],
				[
					GoStmt.GoIf(GoExpr.GoBinary("==", GoExpr.GoIdent("self"), GoExpr.GoNil), [
						GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoStringLiteral("")]))
					],
						null),
					GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("haxe__xml__Printer_print"), [GoExpr.GoIdent("self")]))
				]),
			GoDecl.GoStructDecl("haxe__crypto__Base64", []),
			GoDecl.GoStructDecl("haxe__crypto__Md5", []),
			GoDecl.GoStructDecl("haxe__crypto__Sha1", []),
			GoDecl.GoStructDecl("haxe__crypto__Sha224", []),
			GoDecl.GoStructDecl("haxe__crypto__Sha256", []),
			GoDecl.GoFuncDecl("hxrt_haxeBytesToRaw", null, [
				{
					name: "value",
					typeName: "*haxe__io__Bytes"
				}
			], ["[]byte"], [
				GoStmt.GoIf(GoExpr.GoBinary("==", GoExpr.GoIdent("value"), GoExpr.GoNil), [GoStmt.GoReturn(GoExpr.GoRaw("[]byte{}"))], null),
				GoStmt.GoRaw("if value.__hx_rawValid && len(value.__hx_raw) == len(value.b) {"),
				GoStmt.GoRaw("\treturn value.__hx_raw"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("raw := make([]byte, len(value.b))"),
				GoStmt.GoRaw("for i := 0; i < len(value.b); i++ {"),
				GoStmt.GoRaw("\traw[i] = byte(value.b[i])"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("value.__hx_raw = raw"),
				GoStmt.GoRaw("value.__hx_rawValid = true"),
				GoStmt.GoReturn(GoExpr.GoIdent("raw"))
			]),
			GoDecl.GoFuncDecl("hxrt_rawToHaxeBytes", null, [
				{
					name: "value",
					typeName: "[]byte"
				}
			], ["*haxe__io__Bytes"], [
				GoStmt.GoRaw("converted := make([]int, len(value))"),
				GoStmt.GoRaw("for i := 0; i < len(value); i++ {"),
				GoStmt.GoRaw("\tconverted[i] = int(value[i])"),
				GoStmt.GoRaw("}"),
				GoStmt.GoReturn(GoExpr.GoRaw("&haxe__io__Bytes{b: converted, length: len(converted), __hx_raw: value, __hx_rawValid: true}"))
			]),
			GoDecl.GoFuncDecl("haxe__crypto__Base64_encode", null, [
				{
					name: "bytes",
					typeName: "*haxe__io__Bytes"
				},
				{name: "complement", typeName: "...bool"}
			], ["*string"], [
				GoStmt.GoVarDecl("useComplement", null, GoExpr.GoBoolLiteral(true), true),
				GoStmt.GoIf(GoExpr.GoBinary(">", GoExpr.GoCall(GoExpr.GoIdent("len"), [GoExpr.GoIdent("complement")]), GoExpr.GoIntLiteral(0)), [
					GoStmt.GoAssign(GoExpr.GoIdent("useComplement"), GoExpr.GoIndex(GoExpr.GoIdent("complement"), GoExpr.GoIntLiteral(0)))
				],
					null),
				GoStmt.GoVarDecl("encoded", null, GoExpr.GoRaw("base64.StdEncoding.EncodeToString(hxrt_haxeBytesToRaw(bytes))"), true),
				GoStmt.GoIf(GoExpr.GoUnary("!", GoExpr.GoIdent("useComplement")), [
					GoStmt.GoAssign(GoExpr.GoIdent("encoded"),
						GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("strings"), "TrimRight"), [GoExpr.GoIdent("encoded"), GoExpr.GoStringLiteral("=")]))
				],
					null),
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoIdent("encoded")]))
			]),
			GoDecl.GoFuncDecl("haxe__crypto__Base64_decode", null, [
				{
					name: "value",
					typeName: "*string"
				},
				{name: "complement", typeName: "...bool"}
			], ["*haxe__io__Bytes"], [
				GoStmt.GoVarDecl("useComplement", null, GoExpr.GoBoolLiteral(true), true),
				GoStmt.GoIf(GoExpr.GoBinary(">", GoExpr.GoCall(GoExpr.GoIdent("len"), [GoExpr.GoIdent("complement")]), GoExpr.GoIntLiteral(0)), [
					GoStmt.GoAssign(GoExpr.GoIdent("useComplement"), GoExpr.GoIndex(GoExpr.GoIdent("complement"), GoExpr.GoIntLiteral(0)))
				], null),
				GoStmt.GoVarDecl("rawValue", null, GoExpr.GoRaw("*hxrt.StdString(value)"), true),
				GoStmt.GoIf(GoExpr.GoIdent("useComplement"), [
					GoStmt.GoAssign(GoExpr.GoIdent("rawValue"),
						GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("strings"), "TrimRight"), [GoExpr.GoIdent("rawValue"), GoExpr.GoStringLiteral("=")]))
				],
					null),
				GoStmt.GoRaw("decoded, err := base64.RawStdEncoding.DecodeString(rawValue)"),
				GoStmt.GoIf(GoExpr.GoBinary("!=", GoExpr.GoIdent("err"), GoExpr.GoNil), [
					GoStmt.GoRaw("decoded, err = base64.StdEncoding.DecodeString(*hxrt.StdString(value))"),
					GoStmt.GoIf(GoExpr.GoBinary("!=", GoExpr.GoIdent("err"), GoExpr.GoNil), [
						GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("hxrt.Throw"), [GoExpr.GoIdent("err")])),
						GoStmt.GoReturn(GoExpr.GoRaw("&haxe__io__Bytes{b: []int{}, length: 0}"))
					],
						null)
				],
					null),
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt_rawToHaxeBytes"), [GoExpr.GoIdent("decoded")]))
			]),
			GoDecl.GoFuncDecl("haxe__crypto__Base64_urlEncode", null, [
				{
					name: "bytes",
					typeName: "*haxe__io__Bytes"
				},
				{name: "complement", typeName: "...bool"}
			], ["*string"], [
				GoStmt.GoVarDecl("useComplement", null, GoExpr.GoBoolLiteral(false), true),
				GoStmt.GoIf(GoExpr.GoBinary(">", GoExpr.GoCall(GoExpr.GoIdent("len"), [GoExpr.GoIdent("complement")]), GoExpr.GoIntLiteral(0)), [
					GoStmt.GoAssign(GoExpr.GoIdent("useComplement"), GoExpr.GoIndex(GoExpr.GoIdent("complement"), GoExpr.GoIntLiteral(0)))
				],
					null),
				GoStmt.GoVarDecl("encoded", null, GoExpr.GoRaw("base64.RawURLEncoding.EncodeToString(hxrt_haxeBytesToRaw(bytes))"), true),
				GoStmt.GoIf(GoExpr.GoIdent("useComplement"), [
					GoStmt.GoVarDecl("missing", null, GoExpr.GoRaw("len(encoded) % 4"), true),
					GoStmt.GoIf(GoExpr.GoBinary("!=", GoExpr.GoIdent("missing"), GoExpr.GoIntLiteral(0)), [
						GoStmt.GoAssign(GoExpr.GoIdent("encoded"),
							GoExpr.GoBinary("+", GoExpr.GoIdent("encoded"),
								GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("strings"), "Repeat"),
									[
										GoExpr.GoStringLiteral("="),
										GoExpr.GoBinary("-", GoExpr.GoIntLiteral(4), GoExpr.GoIdent("missing"))
									])))
					],
						null)
				],
					null),
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoIdent("encoded")]))
			]),
			GoDecl.GoFuncDecl("haxe__crypto__Base64_urlDecode", null, [
				{
					name: "value",
					typeName: "*string"
				},
				{name: "complement", typeName: "...bool"}
			], ["*haxe__io__Bytes"], [
				GoStmt.GoVarDecl("rawValue", null, GoExpr.GoRaw("*hxrt.StdString(value)"), true),
				GoStmt.GoRaw("decoded, err := base64.RawURLEncoding.DecodeString(strings.TrimRight(rawValue, \"=\"))"),
				GoStmt.GoIf(GoExpr.GoBinary("!=", GoExpr.GoIdent("err"), GoExpr.GoNil), [
					GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("hxrt.Throw"), [GoExpr.GoIdent("err")])),
					GoStmt.GoReturn(GoExpr.GoRaw("&haxe__io__Bytes{b: []int{}, length: 0}"))
				],
					null),
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt_rawToHaxeBytes"), [GoExpr.GoIdent("decoded")]))
			]),
			GoDecl.GoFuncDecl("haxe__crypto__Md5_encode", null, [
				{
					name: "value",
					typeName: "*string"
				}
			], ["*string"], [
				GoStmt.GoRaw("sum := md5.Sum([]byte(*hxrt.StdString(value)))"),
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [
					GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("hex"), "EncodeToString"), [GoExpr.GoRaw("sum[:]")])
				]))
			]),
			GoDecl.GoFuncDecl("haxe__crypto__Md5_make", null, [
				{
					name: "value",
					typeName: "*haxe__io__Bytes"
				}
			], ["*haxe__io__Bytes"], [
				GoStmt.GoRaw("sum := md5.Sum(hxrt_haxeBytesToRaw(value))"),
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt_rawToHaxeBytes"), [GoExpr.GoRaw("sum[:]")]))
			]),
			GoDecl.GoFuncDecl("haxe__crypto__Sha1_encode", null, [
				{
					name: "value",
					typeName: "*string"
				}
			], ["*string"], [
				GoStmt.GoRaw("sum := sha1.Sum([]byte(*hxrt.StdString(value)))"),
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [
					GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("hex"), "EncodeToString"), [GoExpr.GoRaw("sum[:]")])
				]))
			]),
			GoDecl.GoFuncDecl("haxe__crypto__Sha1_make", null, [
				{
					name: "value",
					typeName: "*haxe__io__Bytes"
				}
			], ["*haxe__io__Bytes"], [
				GoStmt.GoRaw("sum := sha1.Sum(hxrt_haxeBytesToRaw(value))"),
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt_rawToHaxeBytes"), [GoExpr.GoRaw("sum[:]")]))
			]),
			GoDecl.GoFuncDecl("haxe__crypto__Sha224_encode", null, [
				{
					name: "value",
					typeName: "*string"
				}
			], ["*string"], [
				GoStmt.GoRaw("sum := sha256.Sum224([]byte(*hxrt.StdString(value)))"),
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [
					GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("hex"), "EncodeToString"), [GoExpr.GoRaw("sum[:]")])
				]))
			]),
			GoDecl.GoFuncDecl("haxe__crypto__Sha224_make", null, [
				{
					name: "value",
					typeName: "*haxe__io__Bytes"
				}
			], ["*haxe__io__Bytes"], [
				GoStmt.GoRaw("sum := sha256.Sum224(hxrt_haxeBytesToRaw(value))"),
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt_rawToHaxeBytes"), [GoExpr.GoRaw("sum[:]")]))
			]),
			GoDecl.GoFuncDecl("haxe__crypto__Sha256_encode", null, [
				{
					name: "value",
					typeName: "*string"
				}
			], ["*string"], [
				GoStmt.GoRaw("sum := sha256.Sum256([]byte(*hxrt.StdString(value)))"),
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [
					GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("hex"), "EncodeToString"), [GoExpr.GoRaw("sum[:]")])
				]))
			]),
			GoDecl.GoFuncDecl("haxe__crypto__Sha256_make", null, [
				{
					name: "value",
					typeName: "*haxe__io__Bytes"
				}
			], ["*haxe__io__Bytes"], [
				GoStmt.GoRaw("sum := sha256.Sum256(hxrt_haxeBytesToRaw(value))"),
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt_rawToHaxeBytes"), [GoExpr.GoRaw("sum[:]")]))
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
			],
				["*haxe__ds__Option"], [GoStmt.GoReturn(GoExpr.GoRaw("&haxe__ds__Option{tag: 0, params: []any{value}}"))]),
			GoDecl.GoStructDecl("haxe__xml__Parser", []),
			GoDecl.GoStructDecl("haxe__xml__Printer", []),
			GoDecl.GoFuncDecl("haxe__xml__Parser_parse", null, [
				{
					name: "source",
					typeName: "*string"
				},
				{name: "strict", typeName: "...bool"}
			], ["*Xml"], [
				GoStmt.GoRaw("raw := *hxrt.StdString(source)"),
				GoStmt.GoRaw("doc := Xml_createDocument()"),
				GoStmt.GoRaw("stack := []*Xml{doc}"),
				GoStmt.GoRaw("decoder := xml.NewDecoder(strings.NewReader(raw))"),
				GoStmt.GoRaw("for {"),
				GoStmt.GoRaw("\ttokenStart := decoder.InputOffset()"),
				GoStmt.GoRaw("\ttoken, err := decoder.Token()"),
				GoStmt.GoRaw("\ttokenEnd := decoder.InputOffset()"),
				GoStmt.GoRaw("\tif err == io.EOF {"),
				GoStmt.GoRaw("\t\tbreak"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tif err != nil {"),
				GoStmt.GoRaw("\t\thxrt.Throw(err)"),
				GoStmt.GoRaw("\t\treturn Xml_createDocument()"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tcurrent := stack[len(stack)-1]"),
				GoStmt.GoRaw("\tswitch value := token.(type) {"),
				GoStmt.GoRaw("\tcase xml.StartElement:"),
				GoStmt.GoRaw("\t\tnode := Xml_createElement(hxrt.StringFromLiteral(value.Name.Local))"),
				GoStmt.GoRaw("\t\tfor _, attr := range value.Attr {"),
				GoStmt.GoRaw("\t\t\tnode.set(hxrt.StringFromLiteral(attr.Name.Local), hxrt.StringFromLiteral(attr.Value))"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t\tcurrent.addChild(node)"),
				GoStmt.GoRaw("\t\tstack = append(stack, node)"),
				GoStmt.GoRaw("\tcase xml.EndElement:"),
				GoStmt.GoRaw("\t\tif len(stack) > 1 {"),
				GoStmt.GoRaw("\t\t\tstack = stack[:len(stack)-1]"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\tcase xml.CharData:"),
				GoStmt.GoRaw("\t\ttext := string([]byte(value))"),
				GoStmt.GoRaw("\t\tif len(text) != 0 {"),
				GoStmt.GoRaw("\t\t\ttokenSource := raw[tokenStart:tokenEnd]"),
				GoStmt.GoRaw("\t\t\tif strings.HasPrefix(tokenSource, \"<![CDATA[\") && strings.HasSuffix(tokenSource, \"]]>\") {"),
				GoStmt.GoRaw("\t\t\t\tcurrent.addChild(Xml_createCData(hxrt.StringFromLiteral(text)))"),
				GoStmt.GoRaw("\t\t\t} else {"),
				GoStmt.GoRaw("\t\t\t\tcurrent.addChild(Xml_createPCData(hxrt.StringFromLiteral(text)))"),
				GoStmt.GoRaw("\t\t\t}"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\tcase xml.Comment:"),
				GoStmt.GoRaw("\t\tcurrent.addChild(Xml_createComment(hxrt.StringFromLiteral(string([]byte(value)))))"),
				GoStmt.GoRaw("\tcase xml.Directive:"),
				GoStmt.GoRaw("\t\tdirective := strings.TrimSpace(string([]byte(value)))"),
				GoStmt.GoRaw("\t\tupper := strings.ToUpper(directive)"),
				GoStmt.GoRaw("\t\tif strings.HasPrefix(upper, \"DOCTYPE\") {"),
				GoStmt.GoRaw("\t\t\tdirective = strings.TrimSpace(directive[len(\"DOCTYPE\"):])"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t\tcurrent.addChild(Xml_createDocType(hxrt.StringFromLiteral(directive)))"),
				GoStmt.GoRaw("\tcase xml.ProcInst:"),
				GoStmt.GoRaw("\t\tpayload := value.Target"),
				GoStmt.GoRaw("\t\tif len(value.Inst) != 0 {"),
				GoStmt.GoRaw("\t\t\tpayload += \" \" + string(value.Inst)"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t\tcurrent.addChild(Xml_createProcessingInstruction(hxrt.StringFromLiteral(strings.TrimSpace(payload))))"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("}"),
				GoStmt.GoReturn(GoExpr.GoIdent("doc"))
			]),
			GoDecl.GoFuncDecl("haxe__xml__Printer_escapeText", null, [
				{
					name: "value",
					typeName: "string"
				}
			], ["string"], [
				GoStmt.GoReturn(GoExpr.GoRaw("strings.NewReplacer(\"&\", \"&amp;\", \"<\", \"&lt;\", \">\", \"&gt;\").Replace(value)"))
			]),
			GoDecl.GoFuncDecl("haxe__xml__Printer_escapeAttr", null, [
				{
					name: "value",
					typeName: "string"
				}
			], ["string"], [
				GoStmt.GoReturn(GoExpr.GoRaw("strings.NewReplacer(\"&\", \"&amp;\", \"<\", \"&lt;\", \">\", \"&gt;\", \"\\\"\", \"&quot;\").Replace(value)"))
			]),
			GoDecl.GoFuncDecl("haxe__xml__Printer_hasChildren", null, [
				{
					name: "value",
					typeName: "*Xml"
				}
			], ["bool"], [
				GoStmt.GoRaw("for _, child := range value.children {"),
				GoStmt.GoRaw("\tswitch child.nodeType {"),
				GoStmt.GoRaw("\tcase Xml_Element, Xml_PCData:"),
				GoStmt.GoRaw("\t\treturn true"),
				GoStmt.GoRaw("\tcase Xml_CData, Xml_Comment:"),
				GoStmt.GoRaw("\t\tif len(strings.TrimLeft(*hxrt.StdString(child.nodeValue), \" \\n\\r\\t\")) != 0 {"),
				GoStmt.GoRaw("\t\t\treturn true"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("}"),
				GoStmt.GoReturn(GoExpr.GoBoolLiteral(false))
			]),
			GoDecl.GoFuncDecl("haxe__xml__Printer_writeNode", null, [
				{
					name: "output",
					typeName: "*strings.Builder"
				},
				{name: "value", typeName: "*Xml"},
				{name: "tabs", typeName: "string"},
				{name: "pretty", typeName: "bool"}
			], [], [
				GoStmt.GoRaw("newline := func() {"),
				GoStmt.GoRaw("\tif pretty {"),
				GoStmt.GoRaw("\t\toutput.WriteString(\"\\n\")"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("switch value.nodeType {"),
				GoStmt.GoRaw("case Xml_CData:"),
				GoStmt.GoRaw("\toutput.WriteString(tabs + \"<![CDATA[\")"),
				GoStmt.GoRaw("\toutput.WriteString(*hxrt.StdString(value.nodeValue))"),
				GoStmt.GoRaw("\toutput.WriteString(\"]]>\")"),
				GoStmt.GoRaw("\tnewline()"),
				GoStmt.GoRaw("case Xml_Comment:"),
				GoStmt.GoRaw("\tcommentContent := strings.NewReplacer(\"\\n\", \"\", \"\\r\", \"\", \"\\t\", \"\").Replace(*hxrt.StdString(value.nodeValue))"),
				GoStmt.GoRaw("\toutput.WriteString(tabs)"),
				GoStmt.GoRaw("\toutput.WriteString(strings.TrimSpace(\"<!--\" + commentContent + \"-->\"))"),
				GoStmt.GoRaw("\tnewline()"),
				GoStmt.GoRaw("case Xml_Document:"),
				GoStmt.GoRaw("\tfor _, child := range value.children {"),
				GoStmt.GoRaw("\t\thaxe__xml__Printer_writeNode(output, child, tabs, pretty)"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("case Xml_Element:"),
				GoStmt.GoRaw("\toutput.WriteString(tabs + \"<\")"),
				GoStmt.GoRaw("\toutput.WriteString(*hxrt.StdString(value.nodeName))"),
				GoStmt.GoRaw("\tfor _, attribute := range value.attributeOrder {"),
				GoStmt.GoRaw("\t\toutput.WriteString(\" \" + attribute + \"=\\\"\")"),
				GoStmt.GoRaw("\t\toutput.WriteString(haxe__xml__Printer_escapeAttr(value.attributeMap[attribute]))"),
				GoStmt.GoRaw("\t\toutput.WriteString(\"\\\"\")"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tif haxe__xml__Printer_hasChildren(value) {"),
				GoStmt.GoRaw("\t\toutput.WriteString(\">\")"),
				GoStmt.GoRaw("\t\tnewline()"),
				GoStmt.GoRaw("\t\tchildTabs := tabs"),
				GoStmt.GoRaw("\t\tif pretty {"),
				GoStmt.GoRaw("\t\t\tchildTabs = tabs + \"\\t\""),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t\tfor _, child := range value.children {"),
				GoStmt.GoRaw("\t\t\thaxe__xml__Printer_writeNode(output, child, childTabs, pretty)"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t\toutput.WriteString(tabs + \"</\")"),
				GoStmt.GoRaw("\t\toutput.WriteString(*hxrt.StdString(value.nodeName))"),
				GoStmt.GoRaw("\t\toutput.WriteString(\">\")"),
				GoStmt.GoRaw("\t\tnewline()"),
				GoStmt.GoRaw("\t} else {"),
				GoStmt.GoRaw("\t\toutput.WriteString(\"/>\")"),
				GoStmt.GoRaw("\t\tnewline()"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("case Xml_PCData:"),
				GoStmt.GoRaw("\tnodeValue := *hxrt.StdString(value.nodeValue)"),
				GoStmt.GoRaw("\tif len(nodeValue) != 0 {"),
				GoStmt.GoRaw("\t\toutput.WriteString(tabs + haxe__xml__Printer_escapeText(nodeValue))"),
				GoStmt.GoRaw("\t\tnewline()"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("case Xml_ProcessingInstruction:"),
				GoStmt.GoRaw("\toutput.WriteString(\"<?\" + *hxrt.StdString(value.nodeValue) + \"?>\")"),
				GoStmt.GoRaw("\tnewline()"),
				GoStmt.GoRaw("case Xml_DocType:"),
				GoStmt.GoRaw("\toutput.WriteString(\"<!DOCTYPE \" + *hxrt.StdString(value.nodeValue) + \">\")"),
				GoStmt.GoRaw("\tnewline()"),
				GoStmt.GoRaw("}"),
			]),
			GoDecl.GoFuncDecl("haxe__xml__Printer_print", null, [
				{
					name: "value",
					typeName: "*Xml"
				},
				{name: "pretty", typeName: "...bool"}
			], ["*string"], [
				GoStmt.GoIf(GoExpr.GoBinary("==", GoExpr.GoIdent("value"), GoExpr.GoNil), [
					GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoStringLiteral("")]))
				],
					null),
				GoStmt.GoRaw("usePretty := false"),
				GoStmt.GoRaw("if len(pretty) > 0 {"),
				GoStmt.GoRaw("\tusePretty = pretty[0]"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("var output strings.Builder"),
				GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("haxe__xml__Printer_writeNode"),
					[
						GoExpr.GoRaw("&output"),
						GoExpr.GoIdent("value"),
						GoExpr.GoStringLiteral(""),
						GoExpr.GoIdent("usePretty")
					])),
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"),
					[GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("output"), "String"), [])]))
			]),
			GoDecl.GoStructDecl("haxe__zip__Compress", []),
			GoDecl.GoStructDecl("haxe__zip__Uncompress", []),
			GoDecl.GoFuncDecl("haxe__zip__Compress_run", null, [
				{
					name: "src",
					typeName: "*haxe__io__Bytes"
				},
				{name: "level", typeName: "int"}
			], ["*haxe__io__Bytes"], [
				GoStmt.GoVarDecl("raw", null, GoExpr.GoCall(GoExpr.GoIdent("hxrt_haxeBytesToRaw"), [GoExpr.GoIdent("src")]), true),
				GoStmt.GoRaw("var buffer bytes.Buffer"),
				GoStmt.GoRaw("writer, err := zlib.NewWriterLevel(&buffer, level)"),
				GoStmt.GoIf(GoExpr.GoBinary("!=", GoExpr.GoIdent("err"), GoExpr.GoNil), [
					GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("hxrt.Throw"), [GoExpr.GoIdent("err")])),
					GoStmt.GoReturn(GoExpr.GoNil)
				],
					null),
				GoStmt.GoRaw("if _, err := writer.Write(raw); err != nil {"),
				GoStmt.GoRaw("\t_ = writer.Close()"),
				GoStmt.GoRaw("\thxrt.Throw(err)"),
				GoStmt.GoRaw("\treturn nil"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if err := writer.Close(); err != nil {"),
				GoStmt.GoRaw("\thxrt.Throw(err)"),
				GoStmt.GoRaw("\treturn nil"),
				GoStmt.GoRaw("}"),
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt_rawToHaxeBytes"), [GoExpr.GoRaw("buffer.Bytes()")]))
			]),
			GoDecl.GoFuncDecl("haxe__zip__Uncompress_run", null, [
				{
					name: "src",
					typeName: "*haxe__io__Bytes"
				},
				{name: "bufsize", typeName: "...int"}
			], ["*haxe__io__Bytes"], [
				GoStmt.GoVarDecl("raw", null, GoExpr.GoCall(GoExpr.GoIdent("hxrt_haxeBytesToRaw"), [GoExpr.GoIdent("src")]), true),
				GoStmt.GoRaw("reader, err := zlib.NewReader(bytes.NewReader(raw))"),
				GoStmt.GoIf(GoExpr.GoBinary("!=", GoExpr.GoIdent("err"), GoExpr.GoNil), [
					GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("hxrt.Throw"), [GoExpr.GoIdent("err")])),
					GoStmt.GoReturn(GoExpr.GoNil)
				],
					null),
				GoStmt.GoRaw("defer reader.Close()"),
				GoStmt.GoRaw("decoded, err := io.ReadAll(reader)"),
				GoStmt.GoIf(GoExpr.GoBinary("!=", GoExpr.GoIdent("err"), GoExpr.GoNil), [
					GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("hxrt.Throw"), [GoExpr.GoIdent("err")])),
					GoStmt.GoReturn(GoExpr.GoNil)
				], null),
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt_rawToHaxeBytes"), [GoExpr.GoIdent("decoded")]))
			]),
			GoDecl.GoStructDecl("sys__FileSystem", [])
		];
		if (requiresReflectFieldsShim) {
			decls.push(reflectFieldsShimDecl());
		}
		decls = decls.concat(lowerTypeReflectionShimDecls());
		return decls;
	}

	function reflectFieldsShimDecl():GoDecl {
		return GoDecl.GoFuncDecl("Reflect_fields", null, [
			{
				name: "obj",
				typeName: "any"
			}
		], ["[]*string"], [
			GoStmt.GoRaw("if obj == nil {"),
			GoStmt.GoRaw("\treturn []*string{}"),
			GoStmt.GoRaw("}"),
			GoStmt.GoRaw("switch value := obj.(type) {"),
			GoStmt.GoRaw("case map[string]any:"),
			GoStmt.GoRaw("\tkeys := make([]*string, 0, len(value))"),
			GoStmt.GoRaw("\tfor key := range value {"),
			GoStmt.GoRaw("\t\tkeys = append(keys, hxrt.StringFromLiteral(key))"),
			GoStmt.GoRaw("\t}"),
			GoStmt.GoRaw("\treturn keys"),
			GoStmt.GoRaw("case map[any]any:"),
			GoStmt.GoRaw("\tkeys := make([]*string, 0, len(value))"),
			GoStmt.GoRaw("\tfor key := range value {"),
			GoStmt.GoRaw("\t\tkeys = append(keys, hxrt.StdString(key))"),
			GoStmt.GoRaw("\t}"),
			GoStmt.GoRaw("\treturn keys"),
			GoStmt.GoRaw("case *map[string]any:"),
			GoStmt.GoRaw("\tif value == nil {"),
			GoStmt.GoRaw("\t\treturn []*string{}"),
			GoStmt.GoRaw("\t}"),
			GoStmt.GoRaw("\tkeys := make([]*string, 0, len(*value))"),
			GoStmt.GoRaw("\tfor key := range *value {"),
			GoStmt.GoRaw("\t\tkeys = append(keys, hxrt.StringFromLiteral(key))"),
			GoStmt.GoRaw("\t}"),
			GoStmt.GoRaw("\treturn keys"),
			GoStmt.GoRaw("case *map[any]any:"),
			GoStmt.GoRaw("\tif value == nil {"),
			GoStmt.GoRaw("\t\treturn []*string{}"),
			GoStmt.GoRaw("\t}"),
			GoStmt.GoRaw("\tkeys := make([]*string, 0, len(*value))"),
			GoStmt.GoRaw("\tfor key := range *value {"),
			GoStmt.GoRaw("\t\tkeys = append(keys, hxrt.StdString(key))"),
			GoStmt.GoRaw("\t}"),
			GoStmt.GoRaw("\treturn keys"),
			GoStmt.GoRaw("}"),
			GoStmt.GoRaw("rv := reflect.ValueOf(obj)"),
			GoStmt.GoRaw("if !rv.IsValid() {"),
			GoStmt.GoRaw("\treturn []*string{}"),
			GoStmt.GoRaw("}"),
			GoStmt.GoRaw("if rv.Kind() == reflect.Pointer {"),
			GoStmt.GoRaw("\tif rv.IsNil() {"),
			GoStmt.GoRaw("\t\treturn []*string{}"),
			GoStmt.GoRaw("\t}"),
			GoStmt.GoRaw("\trv = rv.Elem()"),
			GoStmt.GoRaw("}"),
			GoStmt.GoRaw("if rv.Kind() != reflect.Struct {"),
			GoStmt.GoRaw("\treturn []*string{}"),
			GoStmt.GoRaw("}"),
			GoStmt.GoRaw("rt := rv.Type()"),
			GoStmt.GoRaw("keys := make([]*string, 0, rv.NumField())"),
			GoStmt.GoRaw("for i := 0; i < rv.NumField(); i++ {"),
			GoStmt.GoRaw("\tfield := rt.Field(i)"),
			GoStmt.GoRaw("\tif field.PkgPath != \"\" {"),
			GoStmt.GoRaw("\t\tcontinue"),
			GoStmt.GoRaw("\t}"),
			GoStmt.GoRaw("\tkeys = append(keys, hxrt.StringFromLiteral(field.Name))"),
			GoStmt.GoRaw("}"),
			GoStmt.GoRaw("return keys")
		]);
	}

	function lowerTypeReflectionShimDecls():Array<GoDecl> {
		return GoTypeReflectionEmitter.emit(typeReflectionClassMetadata(), typeReflectionEnumMetadata(), goRawQuotedString, goStringPointerArrayLiteral)
			.concat(GoRttiMetadataEmitter.emit(rttiClassMetadata(), goRawQuotedString));
	}

	function lowerTemplateSupportShimDecls():Array<GoDecl> {
		return [
			GoDecl.GoFuncDecl("haxe__Template_anyArrayToSlice_runtime", null, [{name: "value", typeName: "any"}], ["[]any"], [
				GoStmt.GoRaw("if value == nil {"),
				GoStmt.GoRaw("\treturn nil"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("rv := reflect.ValueOf(value)"),
				GoStmt.GoRaw("if !rv.IsValid() {"),
				GoStmt.GoRaw("\treturn nil"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if rv.Kind() == reflect.Pointer {"),
				GoStmt.GoRaw("\tif rv.IsNil() {"),
				GoStmt.GoRaw("\t\treturn nil"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\trv = rv.Elem()"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if rv.Kind() != reflect.Slice && rv.Kind() != reflect.Array {"),
				GoStmt.GoRaw("\treturn nil"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("out := make([]any, rv.Len())"),
				GoStmt.GoRaw("for i := 0; i < rv.Len(); i++ {"),
				GoStmt.GoRaw("\titem := rv.Index(i)"),
				GoStmt.GoRaw("\tif item.CanInterface() {"),
				GoStmt.GoRaw("\t\tout[i] = item.Interface()"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("return out")
			]),
			GoDecl.GoFuncDecl("Reflect_getProperty", null, [
				{
					name: "obj",
					typeName: "any"
				},
				{name: "field", typeName: "*string"}
			], ["any"], [
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("Reflect_field"), [GoExpr.GoIdent("obj"), GoExpr.GoIdent("field")]))
			]),
			GoDecl.GoFuncDecl("Reflect_isObject", null, [
				{
					name: "obj",
					typeName: "any"
				}
			], ["bool"], [
				GoStmt.GoRaw("if obj == nil {"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("rv := reflect.ValueOf(obj)"),
				GoStmt.GoRaw("if !rv.IsValid() {"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("switch rv.Kind() {"),
				GoStmt.GoRaw("case reflect.Pointer, reflect.Interface:"),
				GoStmt.GoRaw("\tif rv.IsNil() {"),
				GoStmt.GoRaw("\t\treturn false"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn Reflect_isObject(rv.Elem().Interface())"),
				GoStmt.GoRaw("case reflect.Struct, reflect.Map:"),
				GoStmt.GoRaw("\treturn true"),
				GoStmt.GoRaw("default:"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("}")
			]),
			GoDecl.GoFuncDecl("Reflect_callMethod", null, [
				{
					name: "obj",
					typeName: "any"
				},
				{name: "funcValue", typeName: "any"},
				{name: "args", typeName: "[]any"}
			], ["any"], [
				GoStmt.GoRaw("_ = obj"),
				GoStmt.GoRaw("if funcValue == nil {"),
				GoStmt.GoRaw("\treturn nil"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("fn := reflect.ValueOf(funcValue)"),
				GoStmt.GoRaw("if !fn.IsValid() || fn.Kind() != reflect.Func {"),
				GoStmt.GoRaw("\treturn nil"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("callArgs := make([]reflect.Value, 0, len(args))"),
				GoStmt.GoRaw("for _, arg := range args {"),
				GoStmt.GoRaw("\tcallArgs = append(callArgs, reflect.ValueOf(arg))"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("results := fn.Call(callArgs)"),
				GoStmt.GoRaw("if len(results) == 0 {"),
				GoStmt.GoRaw("\treturn nil"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("return results[0].Interface()")
			])
		];
	}

	function lowerRegexSerializerShimDecls():Array<GoDecl> {
		return GoRegexSerializerEmitter.emit(serializerClassMetadata(), serializerEnumMetadata(), goRawQuotedString);
	}

	function lowerNetSocketShimDecls():Array<GoDecl> {
		return GoNetSocketEmitter.emit(requiresUdpSocketSurface);
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
		var ioSubclassKind = ioStdlibSubclassKind(classType);

		var instanceDataFields = new Array<GoParam>();
		var instanceMethods = new Array<{name:String, func:TFunc, fieldType:Type}>();
		for (field in classType.fields.get()) {
			switch (field.kind) {
				case FVar(_, _):
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
		if (ioSubclassKind != null && !hasStructField(instanceDataFields, "__hx_io_bigEndian")) {
			instanceDataFields.push({
				name: "__hx_io_bigEndian",
				typeName: "bool"
			});
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

			var dispatchMethods = collectDispatchMethods(classType);
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
		}

		for (method in instanceMethods) {
			decls.push(lowerInstanceMethodDecl(classType, method.name, method.func, method.fieldType));
		}
		decls = decls.concat(lowerIoSubclassSyntheticDecls(classType, ioSubclassKind, instanceMethods));

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
						valueExpr == null ? null : lowerExpr(valueExpr).expr;
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

	function lowerConstructorDecl(classType:ClassType, ctorFunc:Null<TFunc>, ctorType:Null<Type>, superClass:Null<ClassType>):GoDecl {
		pushFunctionVarNameScope();
		var typeName = classTypeName(classType);
		var params = ctorFunc == null ? [] : lowerFunctionParams(ctorFunc, typedFunctionArgs(ctorType == null ? ctorFunc.t : ctorType));
		var body = new Array<GoStmt>();
		body.push(GoStmt.GoVarDecl("self", null, GoExpr.GoRaw("&" + typeName + "{}"), true));

		var loweredCtorBody:ConstructorBodyLowering = {
			superArgs: null,
			body: []
		};
		if (ctorFunc != null) {
			loweredCtorBody = lowerConstructorBody(ctorFunc.expr);
		}

		if (superClass != null) {
			var superTypeName = classTypeName(superClass);
			var superCtorArgs = loweredCtorBody.superArgs == null ? [] : [for (arg in loweredCtorBody.superArgs) lowerExpr(arg).expr];
			body.push(GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("self"), superTypeName),
				GoExpr.GoCall(GoExpr.GoIdent(constructorSymbol(superClass)), superCtorArgs)));
			if (!GoStdlibOwnership.isCompilerOwnedAuthority(fullClassName(superClass))) {
				body.push(GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoSelector(GoExpr.GoIdent("self"), superTypeName), "__hx_this"), GoExpr.GoIdent("self")));
			}
		} else if (directHaxeExceptionSuperClass(classType)) {
			var exceptionCtorArgs = loweredCtorBody.superArgs == null ? [] : [for (arg in loweredCtorBody.superArgs) lowerExpr(arg).expr];
			while (exceptionCtorArgs.length < 3) {
				exceptionCtorArgs.push(GoExpr.GoNil);
			}
			body.push(GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("self"), "__hx_exception"),
				GoExpr.GoCall(GoExpr.GoIdent("hxrt.BindException"), [GoExpr.GoIdent("self")].concat(exceptionCtorArgs))));
		}
		body.push(GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("self"), "__hx_this"), GoExpr.GoIdent("self")));
		if (ctorFunc != null) {
			prependLineDirective(loweredCtorBody.body, ctorFunc.expr.pos, classType.module);
		}
		body = body.concat(loweredCtorBody.body);
		body.push(GoStmt.GoReturn(GoExpr.GoIdent("self")));
		popFunctionVarNameScope();
		return GoDecl.GoFuncDecl(constructorSymbol(classType), null, params, ["*" + typeName], body);
	}

	function lowerInstanceMethodDecl(classType:ClassType, fieldName:String, func:TFunc, fieldType:Type):GoDecl {
		return lowerFunctionDecl(normalizeIdent(fieldName), func, {
			name: "self",
			typeName: "*" + classTypeName(classType)
		}, classType.module, fieldType);
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

	function ioStdlibSubclassKind(classType:ClassType):Null<String> {
		var cursor = classType.superClass;
		while (cursor != null) {
			var superType = cursor.t.get();
			var pack = superType.pack.join(".");
			if (pack == "haxe.io" && superType.name == "Input") {
				return "input";
			}
			if (pack == "haxe.io" && superType.name == "Output") {
				return "output";
			}
			cursor = superType.superClass;
		}
		return null;
	}

	function ioStdlibClassOrSubclassKind(classType:ClassType):Null<String> {
		var pack = classType.pack.join(".");
		if (pack == "haxe.io" && classType.name == "Input") {
			return "input";
		}
		if (pack == "haxe.io" && classType.name == "Output") {
			return "output";
		}
		return ioStdlibSubclassKind(classType);
	}

	function ioSyntheticMethod(receiverType:String, name:String, params:Array<GoParam>, results:Array<String>, body:Array<GoStmt>):GoDecl {
		return GoDecl.GoFuncDecl(name, {name: "self", typeName: receiverType}, params, results, body);
	}

	function lowerIoInputSyntheticHelper(receiverType:String, methodName:String):GoDecl {
		return switch (methodName) {
			case "readAll":
				ioSyntheticMethod(receiverType, "readAll", [{name: "bufsize", typeName: "...int"}], ["*haxe__io__Bytes"],
					[GoStmt.GoRaw("return haxe__io__input_readAll(self, bufsize...)")]);
			case "readFullBytes":
				ioSyntheticMethod(receiverType, "readFullBytes", [
					{name: "s", typeName: "*haxe__io__Bytes"},
					{name: "pos", typeName: "int"},
					{name: "len", typeName: "int"}
				], [], [GoStmt.GoRaw("haxe__io__input_readFullBytes(self, s, pos, len)")]);
			case "read":
				ioSyntheticMethod(receiverType, "read", [{name: "nbytes", typeName: "int"}], ["*haxe__io__Bytes"],
					[GoStmt.GoRaw("return haxe__io__input_read(self, nbytes)")]);
			case "readUntil":
				ioSyntheticMethod(receiverType, "readUntil", [{name: "end", typeName: "int"}], ["*string"],
					[GoStmt.GoRaw("return haxe__io__input_readUntil(self, end)")]);
			case "readLine":
				ioSyntheticMethod(receiverType, "readLine", [], ["*string"], [GoStmt.GoRaw("return haxe__io__input_readLine(self)")]);
			case "readFloat":
				ioSyntheticMethod(receiverType, "readFloat", [], ["float64"], [GoStmt.GoRaw("return haxe__io__input_readFloat(self)")]);
			case "readDouble":
				ioSyntheticMethod(receiverType, "readDouble", [], ["float64"], [GoStmt.GoRaw("return haxe__io__input_readDouble(self)")]);
			case "readInt8":
				ioSyntheticMethod(receiverType, "readInt8", [], ["int"], [GoStmt.GoRaw("return haxe__io__input_readInt8(self)")]);
			case "readInt16":
				ioSyntheticMethod(receiverType, "readInt16", [], ["int"], [GoStmt.GoRaw("return haxe__io__input_readInt16(self)")]);
			case "readUInt16":
				ioSyntheticMethod(receiverType, "readUInt16", [], ["int"], [GoStmt.GoRaw("return haxe__io__input_readUInt16(self)")]);
			case "readInt24":
				ioSyntheticMethod(receiverType, "readInt24", [], ["int"], [GoStmt.GoRaw("return haxe__io__input_readInt24(self)")]);
			case "readUInt24":
				ioSyntheticMethod(receiverType, "readUInt24", [], ["int"], [GoStmt.GoRaw("return haxe__io__input_readUInt24(self)")]);
			case "readInt32":
				ioSyntheticMethod(receiverType, "readInt32", [], ["int"], [GoStmt.GoRaw("return haxe__io__input_readInt32(self)")]);
			case "readString":
				ioSyntheticMethod(receiverType, "readString", [
					{name: "len", typeName: "int"},
					{name: "encoding", typeName: "...*haxe__io__Encoding"}
				], ["*string"],
					[GoStmt.GoRaw("return haxe__io__input_readString(self, len, encoding...)")]);
			case _:
				Context.fatalError("Unsupported io input helper synthetic method: " + methodName, Context.currentPos());
				ioSyntheticMethod(receiverType, methodName, [], [], []);
		};
	}

	function lowerIoOutputSyntheticHelper(receiverType:String, methodName:String):GoDecl {
		return switch (methodName) {
			case "write":
				ioSyntheticMethod(receiverType, "write", [{name: "s", typeName: "*haxe__io__Bytes"}], [], [GoStmt.GoRaw("haxe__io__output_write(self, s)")]);
			case "writeFullBytes":
				ioSyntheticMethod(receiverType, "writeFullBytes", [
					{name: "s", typeName: "*haxe__io__Bytes"},
					{name: "pos", typeName: "int"},
					{name: "len", typeName: "int"}
				], [], [GoStmt.GoRaw("haxe__io__output_writeFullBytes(self, s, pos, len)")]);
			case "writeFloat":
				ioSyntheticMethod(receiverType, "writeFloat", [{name: "x", typeName: "float64"}], [], [GoStmt.GoRaw("haxe__io__output_writeFloat(self, x)")]);
			case "writeDouble":
				ioSyntheticMethod(receiverType, "writeDouble", [{name: "x", typeName: "float64"}], [], [GoStmt.GoRaw("haxe__io__output_writeDouble(self, x)")]);
			case "writeInt8":
				ioSyntheticMethod(receiverType, "writeInt8", [{name: "x", typeName: "int"}], [], [GoStmt.GoRaw("haxe__io__output_writeInt8(self, x)")]);
			case "writeInt16":
				ioSyntheticMethod(receiverType, "writeInt16", [{name: "x", typeName: "int"}], [], [GoStmt.GoRaw("haxe__io__output_writeInt16(self, x)")]);
			case "writeUInt16":
				ioSyntheticMethod(receiverType, "writeUInt16", [{name: "x", typeName: "int"}], [], [GoStmt.GoRaw("haxe__io__output_writeUInt16(self, x)")]);
			case "writeInt24":
				ioSyntheticMethod(receiverType, "writeInt24", [{name: "x", typeName: "int"}], [], [GoStmt.GoRaw("haxe__io__output_writeInt24(self, x)")]);
			case "writeUInt24":
				ioSyntheticMethod(receiverType, "writeUInt24", [{name: "x", typeName: "int"}], [], [GoStmt.GoRaw("haxe__io__output_writeUInt24(self, x)")]);
			case "writeInt32":
				ioSyntheticMethod(receiverType, "writeInt32", [{name: "x", typeName: "int"}], [], [GoStmt.GoRaw("haxe__io__output_writeInt32(self, x)")]);
			case "writeInput":
				ioSyntheticMethod(receiverType, "writeInput", [{name: "i", typeName: "haxe__io__Input"}, {name: "bufsize", typeName: "...int"}], [],
					[GoStmt.GoRaw("haxe__io__output_writeInput(self, i, bufsize...)")]);
			case "writeString":
				ioSyntheticMethod(receiverType, "writeString", [
					{name: "s", typeName: "*string"},
					{name: "encoding", typeName: "...*haxe__io__Encoding"}
				], [], [GoStmt.GoRaw("haxe__io__output_writeString(self, s, encoding...)")]);
			case _:
				Context.fatalError("Unsupported io output helper synthetic method: " + methodName, Context.currentPos());
				ioSyntheticMethod(receiverType, methodName, [], [], []);
		};
	}

	function lowerIoSubclassSyntheticDecls(classType:ClassType, ioSubclassKind:Null<String>, instanceMethods:Array<{name:String, func:TFunc}>):Array<GoDecl> {
		if (ioSubclassKind == null) {
			return [];
		}

		var out = new Array<GoDecl>();
		var receiverType = "*" + classTypeName(classType);
		var fullName = fullClassName(classType);
		var isSysFileInput = fullName == "sys.io.FileInput";
		var isSysFileOutput = fullName == "sys.io.FileOutput";
		var methodNames = new Map<String, Bool>();
		for (method in instanceMethods) {
			methodNames.set(normalizeIdent(method.name), true);
		}

		inline function hasMethod(name:String):Bool {
			return methodNames.exists(name);
		}

		inline function addMethod(decl:GoDecl, name:String):Void {
			out.push(decl);
			methodNames.set(name, true);
		}

		if (ioSubclassKind == "input") {
			if (!hasMethod("get_bigEndian")) {
				addMethod(ioSyntheticMethod(receiverType, "get_bigEndian", [], ["bool"], [
					GoStmt.GoIf(GoExpr.GoBinary("==", GoExpr.GoIdent("self"), GoExpr.GoNil), [GoStmt.GoReturn(GoExpr.GoBoolLiteral(false))], null),
					GoStmt.GoReturn(GoExpr.GoSelector(GoExpr.GoIdent("self"), "__hx_io_bigEndian"))
				]), "get_bigEndian");
			}
			if (!hasMethod("set_bigEndian")) {
				addMethod(ioSyntheticMethod(receiverType, "set_bigEndian", [{name: "e", typeName: "bool"}], ["bool"], [
					GoStmt.GoIf(GoExpr.GoBinary("!=", GoExpr.GoIdent("self"), GoExpr.GoNil), [
						GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("self"), "__hx_io_bigEndian"), GoExpr.GoIdent("e"))
					], null),
					GoStmt.GoReturn(GoExpr.GoIdent("e"))
				]), "set_bigEndian");
			}
			if (!hasMethod("close")) {
				if (isSysFileInput) {
					addMethod(ioSyntheticMethod(receiverType, "close", [], [], [
						GoStmt.GoRaw("impl := sys__io__fileInputHandles[self]"),
						GoStmt.GoRaw("if impl == nil {"),
						GoStmt.GoRaw("\treturn"),
						GoStmt.GoRaw("}"),
						GoStmt.GoRaw("if err := impl.Close(); err != nil {"),
						GoStmt.GoRaw("\thxrt.Throw(hxrt.StringFromLiteral(err.Error()))"),
						GoStmt.GoRaw("\treturn"),
						GoStmt.GoRaw("}"),
						GoStmt.GoRaw("delete(sys__io__fileInputHandles, self)")
					]), "close");
				} else {
					addMethod(ioSyntheticMethod(receiverType, "close", [], [], [GoStmt.GoAssign(GoExpr.GoIdent("_"), GoExpr.GoIdent("self"))]), "close");
				}
			}
			if (!hasMethod("readByte")) {
				if (isSysFileInput) {
					addMethod(ioSyntheticMethod(receiverType, "readByte", [], ["int"], [
						GoStmt.GoRaw("impl := sys__io__fileInputHandles[self]"),
						GoStmt.GoRaw("if impl == nil {"),
						GoStmt.GoRaw("\thxrt.Throw(&haxe__io__Eof{})"),
						GoStmt.GoRaw("\treturn 0"),
						GoStmt.GoRaw("}"),
						GoStmt.GoRaw("value, eof, err := impl.ReadByte()"),
						GoStmt.GoRaw("if err != nil {"),
						GoStmt.GoRaw("\thxrt.Throw(hxrt.StringFromLiteral(err.Error()))"),
						GoStmt.GoRaw("\treturn 0"),
						GoStmt.GoRaw("}"),
						GoStmt.GoRaw("if eof {"),
						GoStmt.GoRaw("\thxrt.Throw(&haxe__io__Eof{})"),
						GoStmt.GoRaw("\treturn 0"),
						GoStmt.GoRaw("}"),
						GoStmt.GoReturn(GoExpr.GoIdent("value"))
					]), "readByte");
				} else {
					addMethod(ioSyntheticMethod(receiverType, "readByte", [], ["int"], [
						GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("hxrt.Throw"), [
							GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoStringLiteral("Not implemented")])
						])),
						GoStmt.GoReturn(GoExpr.GoIntLiteral(0))
					]), "readByte");
				}
			}
			if (!hasMethod("readBytes")) {
				addMethod(ioSyntheticMethod(receiverType, "readBytes", [
					{name: "buf", typeName: "*haxe__io__Bytes"},
					{name: "pos", typeName: "int"},
					{name: "len", typeName: "int"}
				], ["int"], [
					GoStmt.GoRaw("if buf == nil || pos < 0 || len < 0 || pos+len > buf.length {"),
					GoStmt.GoRaw("\thxrt.Throw(haxe__io__Error_OutsideBounds)"),
					GoStmt.GoRaw("\treturn 0"),
					GoStmt.GoRaw("}"),
					GoStmt.GoRaw("if self == nil {"),
					GoStmt.GoRaw("\thxrt.Throw(haxe__io__Error_Blocked)"),
					GoStmt.GoRaw("\treturn 0"),
					GoStmt.GoRaw("}"),
					GoStmt.GoRaw("k := 0"),
					GoStmt.GoRaw("for k < len {"),
					GoStmt.GoRaw("\tvalue := 0"),
					GoStmt.GoRaw("\tthrew := false"),
					GoStmt.GoRaw("\tvar thrown any"),
					GoStmt.GoRaw("\tfunc() {"),
					GoStmt.GoRaw("\t\tdefer func() {"),
					GoStmt.GoRaw("\t\t\tif recovered := recover(); recovered != nil {"),
					GoStmt.GoRaw("\t\t\t\tthrew = true"),
					GoStmt.GoRaw("\t\t\t\tthrown = hxrt.UnwrapException(recovered)"),
					GoStmt.GoRaw("\t\t\t}"),
					GoStmt.GoRaw("\t\t}()"),
					GoStmt.GoRaw("\t\tvalue = self.readByte()"),
					GoStmt.GoRaw("\t}()"),
					GoStmt.GoRaw("\tif threw {"),
					GoStmt.GoRaw("\t\tif haxe__io__input_isEof(thrown) {"),
					GoStmt.GoRaw("\t\t\tif k > 0 {"),
					GoStmt.GoRaw("\t\t\t\treturn k"),
					GoStmt.GoRaw("\t\t\t}"),
					GoStmt.GoRaw("\t\t}"),
					GoStmt.GoRaw("\t\thxrt.Throw(thrown)"),
					GoStmt.GoRaw("\t\treturn 0"),
					GoStmt.GoRaw("\t}"),
					GoStmt.GoRaw("\tbuf.b[pos+k] = value"),
					GoStmt.GoRaw("\tk++"),
					GoStmt.GoRaw("}"),
					GoStmt.GoRaw("return len")
				]), "readBytes");
			}
			for (methodName in [
				"readAll",
				"readFullBytes",
				"read",
				"readUntil",
				"readLine",
				"readFloat",
				"readDouble",
				"readInt8",
				"readInt16",
				"readUInt16",
				"readInt24",
				"readUInt24",
				"readInt32",
				"readString"
			]) {
				if (!hasMethod(methodName)) {
					addMethod(lowerIoInputSyntheticHelper(receiverType, methodName), methodName);
				}
			}
		}

		if (ioSubclassKind == "output") {
			if (!hasMethod("get_bigEndian")) {
				addMethod(ioSyntheticMethod(receiverType, "get_bigEndian", [], ["bool"], [
					GoStmt.GoIf(GoExpr.GoBinary("==", GoExpr.GoIdent("self"), GoExpr.GoNil), [GoStmt.GoReturn(GoExpr.GoBoolLiteral(false))], null),
					GoStmt.GoReturn(GoExpr.GoSelector(GoExpr.GoIdent("self"), "__hx_io_bigEndian"))
				]), "get_bigEndian");
			}
			if (!hasMethod("set_bigEndian")) {
				addMethod(ioSyntheticMethod(receiverType, "set_bigEndian", [{name: "e", typeName: "bool"}], ["bool"], [
					GoStmt.GoIf(GoExpr.GoBinary("!=", GoExpr.GoIdent("self"), GoExpr.GoNil), [
						GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("self"), "__hx_io_bigEndian"), GoExpr.GoIdent("e"))
					], null),
					GoStmt.GoReturn(GoExpr.GoIdent("e"))
				]), "set_bigEndian");
			}
			if (!hasMethod("flush")) {
				if (isSysFileOutput) {
					addMethod(ioSyntheticMethod(receiverType, "flush", [], [], [
						GoStmt.GoRaw("impl := sys__io__fileOutputHandles[self]"),
						GoStmt.GoRaw("if impl == nil {"),
						GoStmt.GoRaw("\treturn"),
						GoStmt.GoRaw("}"),
						GoStmt.GoRaw("if err := impl.Flush(); err != nil {"),
						GoStmt.GoRaw("\thxrt.Throw(hxrt.StringFromLiteral(err.Error()))"),
						GoStmt.GoRaw("}")
					]), "flush");
				} else {
					addMethod(ioSyntheticMethod(receiverType, "flush", [], [], [GoStmt.GoAssign(GoExpr.GoIdent("_"), GoExpr.GoIdent("self"))]), "flush");
				}
			}
			if (!hasMethod("close")) {
				if (isSysFileOutput) {
					addMethod(ioSyntheticMethod(receiverType, "close", [], [], [
						GoStmt.GoRaw("impl := sys__io__fileOutputHandles[self]"),
						GoStmt.GoRaw("if impl == nil {"),
						GoStmt.GoRaw("\treturn"),
						GoStmt.GoRaw("}"),
						GoStmt.GoRaw("if err := impl.Close(); err != nil {"),
						GoStmt.GoRaw("\thxrt.Throw(hxrt.StringFromLiteral(err.Error()))"),
						GoStmt.GoRaw("\treturn"),
						GoStmt.GoRaw("}"),
						GoStmt.GoRaw("delete(sys__io__fileOutputHandles, self)")
					]), "close");
				} else {
					addMethod(ioSyntheticMethod(receiverType, "close", [], [], [GoStmt.GoAssign(GoExpr.GoIdent("_"), GoExpr.GoIdent("self"))]), "close");
				}
			}
			if (!hasMethod("prepare")) {
				addMethod(ioSyntheticMethod(receiverType, "prepare", [{name: "nbytes", typeName: "int"}], [], [
					GoStmt.GoAssign(GoExpr.GoIdent("_"), GoExpr.GoIdent("self")),
					GoStmt.GoAssign(GoExpr.GoIdent("_"), GoExpr.GoIdent("nbytes"))
				]), "prepare");
			}
			if (!hasMethod("writeByte")) {
				if (isSysFileOutput) {
					addMethod(ioSyntheticMethod(receiverType, "writeByte", [{name: "c", typeName: "int"}], [], [
						GoStmt.GoRaw("impl := sys__io__fileOutputHandles[self]"),
						GoStmt.GoRaw("if impl == nil {"),
						GoStmt.GoRaw("\thxrt.Throw(hxrt.StringFromLiteral(\"FileOutput is closed\"))"),
						GoStmt.GoRaw("\treturn"),
						GoStmt.GoRaw("}"),
						GoStmt.GoRaw("if err := impl.WriteByte(c); err != nil {"),
						GoStmt.GoRaw("\thxrt.Throw(hxrt.StringFromLiteral(err.Error()))"),
						GoStmt.GoRaw("}")
					]), "writeByte");
				} else {
					addMethod(ioSyntheticMethod(receiverType, "writeByte", [{name: "c", typeName: "int"}], [], [
						GoStmt.GoAssign(GoExpr.GoIdent("_"), GoExpr.GoIdent("c")),
						GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("hxrt.Throw"), [
							GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoStringLiteral("Not implemented")])
						]))
					]), "writeByte");
				}
			}
			if (!hasMethod("writeBytes")) {
				addMethod(ioSyntheticMethod(receiverType, "writeBytes", [
					{name: "s", typeName: "*haxe__io__Bytes"},
					{name: "pos", typeName: "int"},
					{name: "len", typeName: "int"}
				], ["int"], [
					GoStmt.GoRaw("if s == nil || pos < 0 || len < 0 || pos+len > s.length {"),
					GoStmt.GoRaw("\thxrt.Throw(haxe__io__Error_OutsideBounds)"),
					GoStmt.GoRaw("\treturn 0"),
					GoStmt.GoRaw("}"),
					GoStmt.GoRaw("n := len"),
					GoStmt.GoRaw("for len > 0 {"),
					GoStmt.GoRaw("\tself.writeByte(s.b[pos])"),
					GoStmt.GoRaw("\tpos++"),
					GoStmt.GoRaw("\tlen--"),
					GoStmt.GoRaw("}"),
					GoStmt.GoRaw("return n")
				]), "writeBytes");
			}
			for (methodName in [
				"write",
				"writeFullBytes",
				"writeFloat",
				"writeDouble",
				"writeInt8",
				"writeInt16",
				"writeUInt16",
				"writeInt24",
				"writeUInt24",
				"writeInt32",
				"writeInput",
				"writeString"
			]) {
				if (!hasMethod(methodName)) {
					addMethod(lowerIoOutputSyntheticHelper(receiverType, methodName), methodName);
				}
			}
		}

		return out;
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

	function lowerFunctionResults(returnType:Type):Array<String> {
		if (isVoidType(returnType)) {
			return [];
		}
		return [scalarGoType(returnType)];
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

				var lowered = value == null ? null : lowerExprWithPrefix(value);
				var prefix = lowered == null ? [] : lowered.prefix;
				var loweredValue = lowered == null ? null : lowered.expr;
				if (value != null && loweredValue != null) {
					loweredValue = upcastIfNeeded(loweredValue, value.t, variable.t);
					var valueKnownNonNullPrimitive = nonNullPrimitiveExprGoType(value) != null;
					loweredValue = coerceAnyExprToType(loweredValue, value.t, variable.t,
						(exprBackedByAny(value) && !valueKnownNonNullPrimitive)
						|| shouldForceAnyCoerce(value.t, variable.t));
				}
				var goType = valueStorageGoType(variable.t);
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
						var loweredRight = lowerExprWithPrefix(right);
						var lengthAssignStmts = lowerArrayLengthAssign(left, loweredRight.expr);
						var assignStmts = if (lengthAssignStmts != null) {
							lengthAssignStmts;
						} else {
							var rightExpr = upcastIfNeeded(loweredRight.expr, right.t, left.t);
							[GoStmt.GoAssign(lowerLValue(left), rightExpr)];
						};
						if (loweredRight.prefix.length > 0) {
							loweredRight.prefix.concat(assignStmts);
						} else {
							assignStmts;
						}
					case OpAssignOp(assignOp):
						var loweredRight = lowerExprWithPrefix(right);
						var rightExpr = upcastIfNeeded(loweredRight.expr, right.t, left.t);
						var targetExpr = lowerLValue(left);
						var assignExpr = lowerAssignOpExpr(assignOp, targetExpr, rightExpr, left.t, right.t, expr.pos);
						var assignStmt = GoStmt.GoAssign(targetExpr, assignExpr);
						if (loweredRight.prefix.length > 0) {
							loweredRight.prefix.concat([assignStmt]);
						} else {
							[assignStmt];
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
					[GoStmt.GoWhile(lowerExpr(condition).expr, lowerToStatements(body))];
				} else {
					var firstPassVar = freshTempName("hx_do_first");
					var loweredCondition = lowerExpr(condition).expr;
					var loopCondition = GoExpr.GoBinary("||", GoExpr.GoIdent(firstPassVar), loweredCondition);
					var loopBody = [GoStmt.GoAssign(GoExpr.GoIdent(firstPassVar), GoExpr.GoBoolLiteral(false))].concat(lowerToStatements(body));
					[
						GoStmt.GoVarDecl(firstPassVar, null, GoExpr.GoBoolLiteral(true), true),
						GoStmt.GoWhile(loopCondition, loopBody)
					];
				}
			case TBreak:
				[GoStmt.GoBreak];
			case TContinue:
				[GoStmt.GoContinue];
			case TUnop(op, _, value):
				switch (op) {
					case OpIncrement:
						var target = lowerLValue(value);
						[GoStmt.GoAssign(target, unitStepExpr(target, "+", value.t, expr.pos))];
					case OpDecrement:
						var target = lowerLValue(value);
						[GoStmt.GoAssign(target, unitStepExpr(target, "-", value.t, expr.pos))];
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
						var loweredReturn = lowerExprWithPrefix(value);
						var returnExpr = loweredReturn.expr;
						redirected = redirected.concat(loweredReturn.prefix);
						if (redirect.valueName != null && redirect.valueType != null) {
							returnExpr = upcastIfNeeded(returnExpr, value.t, redirect.valueType);
							redirected.push(GoStmt.GoAssign(GoExpr.GoIdent(redirect.valueName), returnExpr));
						} else {
							redirected.push(GoStmt.GoExprStmt(returnExpr));
						}
					}
					redirected.push(GoStmt.GoAssign(GoExpr.GoIdent(redirect.flagName), GoExpr.GoBoolLiteral(true)));
					redirected.push(GoStmt.GoReturn(null));
					redirected;
				} else if (value == null) {
					[GoStmt.GoReturn(null)];
				} else {
					var loweredReturn = lowerExprWithPrefix(value);
					var returnExpr = loweredReturn.expr;
					var expectedReturnType = currentFunctionReturnType();
					if (expectedReturnType != null) {
						returnExpr = upcastIfNeeded(returnExpr, value.t, expectedReturnType);
					}
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
						var appendArgs = [site.tempExpr];
						for (arg in args) {
							var appendValue = lowerExpr(arg).expr;
							if (shouldMaskToByte) {
								appendValue = GoExpr.GoBinary("&", appendValue, GoExpr.GoIntLiteral(255));
							}
							appendArgs.push(appendValue);
						}
						site.prefix.concat([
							GoStmt.GoAssign(site.tempExpr, GoExpr.GoCall(GoExpr.GoIdent("append"), appendArgs))
						]).concat(site.writeBack(site.tempExpr));
					} else if (arrayCall != null && arrayCall.methodName == "pop") {
						var site = lowerArrayMutationSite(arrayCall.target);
						var lenExpr = GoExpr.GoCall(GoExpr.GoIdent("len"), [site.tempExpr]);
						site.prefix.concat([
							GoStmt.GoIf(GoExpr.GoBinary(">", lenExpr, GoExpr.GoIntLiteral(0)), [
								GoStmt.GoAssign(site.tempExpr, GoExpr.GoSlice(site.tempExpr, null, GoExpr.GoBinary("-", lenExpr, GoExpr.GoIntLiteral(1))))
							], null)
						]).concat(site.writeBack(site.tempExpr));
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

	function lowerArrayLengthAssign(left:TypedExpr, rightExpr:GoExpr):Null<Array<GoStmt>> {
		return switch (left.expr) {
			case TField(target, access):
				var fieldName = fieldAccessName(access);
				if (fieldName != "length" || !isArrayType(target.t)) {
					null;
				} else {
					var targetExpr = lowerExpr(target).expr;
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
		var loweredCases = new Array<GoSwitchCase>();
		for (caseEntry in cases) {
			loweredCases.push({
				values: [for (caseValue in caseEntry.values) lowerExpr(caseValue).expr],
				body: lowerToStatements(caseEntry.expr)
			});
		}

		return GoStmt.GoSwitch(lowerExpr(value).expr, loweredCases, defaultExpr == null ? null : lowerToStatements(defaultExpr));
	}

	function lowerSwitchExpr(value:TypedExpr, cases:Array<{values:Array<TypedExpr>, expr:TypedExpr}>, defaultExpr:Null<TypedExpr>,
			resultType:Type):LoweredExprWithPrefix {
		var temp = freshTempName("hx_switch");
		var loweredCases = new Array<GoSwitchCase>();

		for (caseEntry in cases) {
			var loweredCase = lowerExprWithPrefix(caseEntry.expr);
			var caseBody = loweredCase.prefix.concat([GoStmt.GoAssign(GoExpr.GoIdent(temp), loweredCase.expr)]);
			loweredCases.push({
				values: [for (caseValue in caseEntry.values) lowerExpr(caseValue).expr],
				body: caseBody
			});
		}

		var defaultBody:Null<Array<GoStmt>> = null;
		if (defaultExpr != null) {
			var loweredDefault = lowerExprWithPrefix(defaultExpr);
			defaultBody = loweredDefault.prefix.concat([GoStmt.GoAssign(GoExpr.GoIdent(temp), loweredDefault.expr)]);
		}

		return {
			prefix: [
				GoStmt.GoVarDecl(temp, valueStorageGoType(resultType), null, false),
				GoStmt.GoSwitch(lowerExpr(value).expr, loweredCases, defaultBody)
			],
			expr: GoExpr.GoIdent(temp),
			isStringLike: isStringType(resultType)
		};
	}

	function lowerIfExpr(condition:TypedExpr, thenBranch:TypedExpr, elseBranch:Null<TypedExpr>, resultType:Type):LoweredExprWithPrefix {
		var elseExpr = elseBranch;
		if (elseExpr == null) {
			Context.fatalError("If-expression requires an else branch", condition.pos);
		}

		var loweredCondition = lowerExprWithPrefix(condition);
		var facts = conditionNonNullFacts(condition);
		var loweredThen = lowerWithNonNullPrimitiveFacts(facts.thenFacts, function() return lowerExprWithPrefix(thenBranch));
		var loweredElse = lowerWithNonNullPrimitiveFacts(facts.elseFacts, function() return lowerExprWithPrefix(elseExpr));
		var temp = freshTempName("hx_if");
		var loweredThenValue = upcastIfNeeded(loweredThen.expr, thenBranch.t, resultType);
		var loweredElseValue = upcastIfNeeded(loweredElse.expr, elseExpr.t, resultType);
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
		var loweredTry = lowerExprWithPrefix(tryExpr);
		var loweredTryValue = upcastIfNeeded(loweredTry.expr, tryExpr.t, resultType);
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
			var loweredCatch = lowerExprWithPrefix(catchEntry.expr);
			var loweredCatchValue = upcastIfNeeded(loweredCatch.expr, catchEntry.expr.t, resultType);
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
		var prefix = [GoStmt.GoVarDecl(temp, "map[string]any", GoExpr.GoRaw("map[string]any{}"), true)];
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

	function lowerBlock(exprs:Array<TypedExpr>):Array<GoStmt> {
		pushLocalScope();
		var out = new Array<GoStmt>();
		var appliedNonNullFacts = new Map<Int, Null<String>>();
		for (index in 0...exprs.length) {
			var inner = exprs[index];
			registerBlockLocalReassignmentInfo(inner, exprs, index);
			out = out.concat(lowerToStatements(inner));
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

	function pushFunctionReturnType(returnType:Type):Void {
		functionReturnTypeScopes.push(returnType);
	}

	function popFunctionReturnType():Void {
		if (functionReturnTypeScopes.length > 0) {
			functionReturnTypeScopes.pop();
		}
	}

	function currentFunctionReturnType():Null<Type> {
		if (functionReturnTypeScopes.length == 0) {
			return null;
		}
		return functionReturnTypeScopes[functionReturnTypeScopes.length - 1];
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
		var returnType = currentFunctionReturnType();
		if (returnType == null || isVoidType(returnType)) {
			return [];
		}

		var zeroName = freshTempName("hx_throw_zero");
		var returnTypeName = typeToGoType(returnType);
		return [
			GoStmt.GoVarDecl(zeroName, returnTypeName, null, false),
			GoStmt.GoReturn(GoExpr.GoIdent(zeroName))
		];
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

	function calleeReceiverClass(target:TypedExpr):Null<ClassType> {
		return switch (Context.follow(target.t)) {
			case TInst(classRef, _):
				classRef.get();
			case _:
				null;
		};
	}

	// Some staged std/compiler-owned methods already lower optional arguments into
	// Go-native rest/vararg shapes. Source-level padding would duplicate those
	// omitted args and change call arity at the emitted Go boundary.
	function shouldSkipInstanceDefaultArgPadding(classType:ClassType, fieldName:String):Bool {
		if (classType.pack.length == 0 && classType.name == "EReg" && fieldName == "matchSub") {
			return true;
		}

		if (classType.pack.length == 1 && classType.pack[0] == "sys" && classType.name == "Http") {
			return switch (fieldName) {
				case "request", "customRequest":
					true;
				case _:
					false;
			};
		}

		return switch (ioStdlibClassOrSubclassKind(classType)) {
			case "input":
				switch (fieldName) {
					case "readAll", "readString":
						true;
					case _:
						false;
				}
			case "output":
				switch (fieldName) {
					case "writeInput", "writeString":
						true;
					case _:
						false;
				}
			case _:
				false;
		};
	}

	function shouldSkipStaticDefaultArgPadding(classType:ClassType, fieldName:String):Bool {
		var pack = classType.pack.join(".");
		return switch (pack) {
			case "_UnicodeString":
				classType.name == "UnicodeString_Impl_";
			case "haxe.crypto":
				classType.name == "Base64";
			case "haxe.xml": classType.name == "Parser" || classType.name == "Printer";
			case "haxe.zip": classType.name == "Compress" || classType.name == "Uncompress";
			case _:
				false;
		};
	}

	function shouldApplySourceDefaultArgPadding(callee:TypedExpr):Bool {
		return switch (callee.expr) {
			case TField(_, FStatic(classRef, field)):
				!shouldSkipStaticDefaultArgPadding(classRef.get(), field.get().name);
			case TField(target, FInstance(classRef, _, field)):
				!shouldSkipInstanceDefaultArgPadding(classRef.get(), field.get().name);
			case TField(target, FAnon(field)) | TField(target, FClosure(_, field)):
				var classType = calleeReceiverClass(target);
				classType == null ? true : !shouldSkipInstanceDefaultArgPadding(classType, field.get().name);
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

	/**
		What:
		Builds a safe array mutation plan for push/pop-style operations.

		Why:
		Anonymous-record field lvalues need a temporary slice plus explicit
		write-back to stay correct on Go. Plain local array lvalues can mutate the
		local directly, which keeps generated Go readable without weakening the
		anonymous-field safety fix.

		How:
		Returns prefix statements that capture the target once when needed, the slice
		expression to mutate, and a write-back closure that stores the final slice in
		the right place after mutation when direct assignment is not enough.
	**/
	function lowerArrayMutationSite(target:TypedExpr):ArrayMutationSite {
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
			case TField(parent, FAnon(field)) if (isAnonymousObjectType(parent.t)):
				var objectName = freshTempName("hx_obj");
				var objectExpr = GoExpr.GoIdent(objectName);
				var fieldName = field.get().name;
				{
					prefix: [
						GoStmt.GoVarDecl(objectName, typeToGoType(parent.t), lowerExpr(parent).expr, true),
						GoStmt.GoVarDecl(tempName, sliceType, lowerExpr(target).expr, true)
					],
					tempExpr: tempExpr,
					sliceType: sliceType,
					writeBack: function(value:GoExpr):Array<GoStmt> {
						return [
							GoStmt.GoAssign(GoExpr.GoIndex(objectExpr, GoExpr.GoStringLiteral(fieldName)), value)
						];
					}
				};
			case TField(parent, FDynamic(name)) if (isAnonymousObjectType(parent.t)):
				var objectName = freshTempName("hx_obj");
				var objectExpr = GoExpr.GoIdent(objectName);
				{
					prefix: [
						GoStmt.GoVarDecl(objectName, typeToGoType(parent.t), lowerExpr(parent).expr, true),
						GoStmt.GoVarDecl(tempName, sliceType, lowerExpr(target).expr, true)
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
				var targetExpr = lowerLValue(target);
				{
					prefix: [GoStmt.GoVarDecl(tempName, sliceType, lowerExpr(target).expr, true)],
					tempExpr: tempExpr,
					sliceType: sliceType,
					writeBack: function(value:GoExpr):Array<GoStmt> {
						return [GoStmt.GoAssign(targetExpr, value)];
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
					expr: GoExpr.GoArrayLiteral(arrayElementGoType(expr.t), [for (value in values) lowerExpr(value).expr]),
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
					expr: GoExpr.GoIndex(lowerExpr(target).expr, lowerExpr(index).expr),
					isStringLike: isStringType(expr.t)
				};
			case TEnumIndex(inner):
				{
					expr: GoExpr.GoSelector(lowerExpr(inner).expr, "tag"),
					isStringLike: false
				};
			case TEnumParameter(target, _, index):
				var payload = GoExpr.GoIndex(GoExpr.GoSelector(lowerExpr(target).expr, "params"), GoExpr.GoIntLiteral(index));
				var payloadType = scalarGoType(expr.t);
				{
					expr: payloadType == "any" ? payload : GoExpr.GoTypeAssert(payload, payloadType),
					isStringLike: isStringType(expr.t)
				};
			case TNew(classRef, _, args):
				var classType = classRef.get();
				noteSourceOwnedStdlibUsage(classType);
				var loweredArgs = [for (arg in args) lowerExpr(arg).expr];
				var ctorInfo = resolveConstructorInfo(classType);
				if (ctorInfo != null
					&& !GoStdlibOwnership.isCompilerOwnedAuthority(fullClassName(classType))
					&& loweredArgs.length < ctorInfo.defaults.length) {
					for (i in loweredArgs.length...ctorInfo.defaults.length) {
						var defaultValue = ctorInfo.defaults[i];
						if (defaultValue == null) {
							Context.fatalError("Missing required constructor argument at position " + i, expr.pos);
						}
						loweredArgs.push(lowerExpr(defaultValue).expr);
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
						registerMetalChanElementGoType(elementGoType);
						notePortableConcurrencyFastpathHit(expr.pos);
						noteLoweringSuccess("go.concurrency.typed", "go_chan_new", expr.pos,
							'Applied typed go.Chan constructor specialization (element type: ' + elementGoType + ").");
						{
							expr: GoExpr.GoCall(GoExpr.GoIdent(metalChanShimName("go__concurrency_newChan", elementGoType)), [GoExpr.GoIntLiteral(0)]),
							isStringLike: false
						};
					} else {
						notePortableConcurrencyFastpathFallback(expr.pos);
						noteLoweringFallback("go.concurrency.typed", "go_chan_new_unmorphable", expr.pos,
							withEligibilityReason("Could not monomorphize go.Chan element type for constructor specialization.", elementEligibility));
						{
							expr: GoExpr.GoCall(GoExpr.GoIdent(constructorSymbol(classType)), loweredArgs),
							isStringLike: false
						};
					}
				} else if (classType.pack.length == 0 && classType.name == "Array") {
					{
						expr: GoExpr.GoArrayLiteral(arrayElementGoType(expr.t), []),
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
				var variableGoType = narrowedStorageGoType == null ? valueStorageGoType(variable.t) : narrowedStorageGoType;
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
				if (innerGoType != castGoType) {
					if (castGoType != "any" && innerGoType == "any") {
						castExpr = lowerNullableAwareTypeAssertExpr(castExpr, expr.t);
					} else if (castGoType == "any" && innerGoType != "any") {
						castExpr = GoExpr.GoCall(GoExpr.GoIdent("any"), [castExpr]);
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
				var loweredValue = lowerExprWithPrefix(value);
				var throwBody = loweredValue.prefix.concat([
					GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("hxrt.Throw"), [loweredValue.expr]))
				]);
				if (isVoidType(expr.t)) {
					{
						expr: GoExpr.GoCall(GoExpr.GoFuncLiteral([], [], throwBody), []),
						isStringLike: false
					};
				} else {
					var resultTypeName = typeToGoType(expr.t);
					var zeroName = freshTempName("hx_throw_zero");
					{
						expr: GoExpr.GoCall(GoExpr.GoFuncLiteral([], [resultTypeName], throwBody.concat([
							GoStmt.GoVarDecl(zeroName, resultTypeName, null, false),
							GoStmt.GoReturn(GoExpr.GoIdent(zeroName))
						])), []),
						isStringLike: isStringType(expr.t)
					};
				}
			case TTypeExpr(moduleType):
				lowerTypeExpr(moduleType);
			case TBinop(op, left, right):
				switch (op) {
					case OpAssign:
						var targetExpr = lowerLValue(left);
						var loweredRight = lowerExprWithPrefix(right);
						var rightExpr = upcastIfNeeded(loweredRight.expr, right.t, left.t);
						{
							expr: GoExpr.GoCall(GoExpr.GoFuncLiteral([], [typeToGoType(left.t)],
								loweredRight.prefix.concat([GoStmt.GoAssign(targetExpr, rightExpr), GoStmt.GoReturn(targetExpr)])),
								[]),
							isStringLike: isStringType(left.t)
						};
					case OpAssignOp(assignOp):
						var targetExpr = lowerLValue(left);
						var loweredRight = lowerExprWithPrefix(right);
						var rightExpr = upcastIfNeeded(loweredRight.expr, right.t, left.t);
						var assignExpr = lowerAssignOpExpr(assignOp, targetExpr, rightExpr, left.t, right.t, expr.pos);
						{
							expr: GoExpr.GoCall(GoExpr.GoFuncLiteral([], [typeToGoType(left.t)],
								loweredRight.prefix.concat([GoStmt.GoAssign(targetExpr, assignExpr), GoStmt.GoReturn(targetExpr)])),
								[]),
							isStringLike: isStringType(left.t)
						};
					case _:
						lowerBinop(op, left, right, expr.t);
				}
			case TUnop(op, postFix, value):
				if (postFix) {
					return switch (op) {
						case OpIncrement, OpDecrement:
							var target = lowerLValue(value);
							var temp = freshTempName("hx_post");
							var opSymbol = op == OpIncrement ? "+" : "-";
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
						var opSymbol = op == OpIncrement ? "+" : "-";
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
				if (exprs.length == 0) {
					{prefix: [], expr: GoExpr.GoNil, isStringLike: false};
				} else {
					var prefix = new Array<GoStmt>();
					for (index in 0...exprs.length - 1) {
						prefix = prefix.concat(lowerToStatements(exprs[index]));
					}
					var tail = lowerExprWithPrefix(exprs[exprs.length - 1]);
					prefix = prefix.concat(tail.prefix);
					{
						prefix: prefix,
						expr: tail.expr,
						isStringLike: tail.isStringLike
					};
				}
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
					expr: GoExpr.GoIndex(loweredTarget.expr, loweredIndex.expr),
					isStringLike: isStringType(expr.t)
				};
			case TUnop(op, postFix, value):
				if (postFix) {
					switch (op) {
						case OpIncrement, OpDecrement:
							var target = lowerLValue(value);
							var temp = freshTempName("hx_post");
							var opSymbol = op == OpIncrement ? "+" : "-";
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
		var markerName = goRawQuotedString(moduleTypeDisplayName(moduleType));
		return {
			expr: GoExpr.GoRaw("&" + markerType + "{name: hxrt.StringFromLiteral(" + markerName + ")}"),
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
				noteStaticStdlibFieldUsage(classType, resolved.name);
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
				noteIoHelperFieldUsage(classType, resolved.name);
				var loweredTarget = lowerExpr(target).expr;

				if (isSuperTarget(target) && isMethodField(resolved)) {
					var baseSelector = GoExpr.GoSelector(GoExpr.GoIdent("self"), classTypeName(classType));
					return {
						expr: GoExpr.GoSelector(baseSelector, normalizeIdent(resolved.name)),
						isStringLike: isStringType(resolved.type)
					};
				}

				if (classType.isInterface) {
					return {
						expr: GoExpr.GoSelector(loweredTarget, normalizeIdent(resolved.name)),
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
						expr: GoExpr.GoCall(GoExpr.GoIdent("len"), [loweredTarget]),
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
				} else if (shouldUseVirtualDispatch(classType, resolved)) {
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
						expr: GoExpr.GoCall(GoExpr.GoIdent("len"), [loweredTarget]),
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
				if (isAnonymousObjectType(target.t)) {
					return {
						expr: GoExpr.GoIndex(loweredTarget, GoExpr.GoStringLiteral(name)),
						isStringLike: false
					};
				}
				var dynamicExpr = if (name == "length" && isArrayType(target.t)) {
					GoExpr.GoCall(GoExpr.GoIdent("len"), [loweredTarget]);
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

		var lambdaStaticCall = lowerLambdaStaticCall(callee, args, returnType);
		if (lambdaStaticCall != null) {
			return lambdaStaticCall;
		}

		var lambdaFunctionValueCall = lowerLambdaFunctionValueCall(callee, args, returnType);
		if (lambdaFunctionValueCall != null) {
			return lambdaFunctionValueCall;
		}

		var dsSortHelperCall = lowerDsSortHelperCall(callee, args, returnType);
		if (dsSortHelperCall != null) {
			return dsSortHelperCall;
		}

		var metalChanCall = lowerMetalGoChanCall(callee, args, returnType);
		if (metalChanCall != null) {
			return metalChanCall;
		}

		var metalSliceCall = lowerMetalGoSliceCall(callee, args, returnType);
		if (metalSliceCall != null) {
			return metalSliceCall;
		}

		var metalMapCall = lowerMetalGoMapCall(callee, args, returnType);
		if (metalMapCall != null) {
			return metalMapCall;
		}

		var metalResultCall = lowerMetalGoResultCall(callee, args, returnType);
		if (metalResultCall != null) {
			return metalResultCall;
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

		if (isStaticCall(callee, "Sys", [], "println")) {
			var arg = args.length > 0 ? lowerExpr(args[0]).expr : GoExpr.GoNil;
			return {
				expr: GoExpr.GoCall(GoExpr.GoIdent("hxrt.Println"), [arg]),
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

		if (isStaticCall(callee, "AtomicObject_Impl_", ["haxe", "atomic", "_AtomicObject"], "load")) {
			requireStdlibShimGroup("atomic");
			var atom = args.length > 0 ? lowerExpr(args[0]).expr : GoExpr.GoNil;
			var rawCall = GoExpr.GoCall(GoExpr.GoIdent("haxe__atomic___AtomicObject__AtomicObject_Impl__load"), [atom]);
			var expectedType = typeToGoType(returnType);
			return {
				expr: expectedType == "any" ? rawCall : GoExpr.GoTypeAssert(rawCall, expectedType),
				isStringLike: isStringType(returnType)
			};
		}

		if (isStaticCall(callee, "AtomicObject_Impl_", ["haxe", "atomic", "_AtomicObject"], "store")) {
			requireStdlibShimGroup("atomic");
			var atom = args.length > 0 ? lowerExpr(args[0]).expr : GoExpr.GoNil;
			var value = args.length > 1 ? lowerExpr(args[1]).expr : GoExpr.GoNil;
			var rawCall = GoExpr.GoCall(GoExpr.GoIdent("haxe__atomic___AtomicObject__AtomicObject_Impl__store"), [atom, value]);
			var expectedType = typeToGoType(returnType);
			return {
				expr: expectedType == "any" ? rawCall : GoExpr.GoTypeAssert(rawCall, expectedType),
				isStringLike: isStringType(returnType)
			};
		}

		if (isStaticCall(callee, "AtomicObject_Impl_", ["haxe", "atomic", "_AtomicObject"], "exchange")) {
			requireStdlibShimGroup("atomic");
			var atom = args.length > 0 ? lowerExpr(args[0]).expr : GoExpr.GoNil;
			var value = args.length > 1 ? lowerExpr(args[1]).expr : GoExpr.GoNil;
			var rawCall = GoExpr.GoCall(GoExpr.GoIdent("haxe__atomic___AtomicObject__AtomicObject_Impl__exchange"), [atom, value]);
			var expectedType = typeToGoType(returnType);
			return {
				expr: expectedType == "any" ? rawCall : GoExpr.GoTypeAssert(rawCall, expectedType),
				isStringLike: isStringType(returnType)
			};
		}

		if (isStaticCall(callee, "AtomicObject_Impl_", ["haxe", "atomic", "_AtomicObject"], "compareExchange")) {
			requireStdlibShimGroup("atomic");
			var atom = args.length > 0 ? lowerExpr(args[0]).expr : GoExpr.GoNil;
			var expected = args.length > 1 ? lowerExpr(args[1]).expr : GoExpr.GoNil;
			var replacement = args.length > 2 ? lowerExpr(args[2]).expr : GoExpr.GoNil;
			var rawCall = GoExpr.GoCall(GoExpr.GoIdent("haxe__atomic___AtomicObject__AtomicObject_Impl__compareExchange"), [atom, expected, replacement]);
			var expectedType = typeToGoType(returnType);
			return {
				expr: expectedType == "any" ? rawCall : GoExpr.GoTypeAssert(rawCall, expectedType),
				isStringLike: isStringType(returnType)
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
			var nullablePrimitiveArg = paramType != null
				&& isNullablePrimitiveType(paramType) ? lowerNullablePrimitiveCallArgExpr(arg) : null;
			var loweredArg = nullablePrimitiveArg == null ? lowerCallArgExpr(arg) : nullablePrimitiveArg;
			if (paramType != null) {
				loweredArg = upcastIfNeeded(loweredArg, arg.t, paramType);
				if (!isNullablePrimitiveType(paramType)) {
					var argKnownNonNullPrimitive = nonNullPrimitiveExprGoType(arg) != null;
					loweredArg = coerceAnyExprToType(loweredArg, arg.t, paramType, !argKnownNonNullPrimitive && (exprBackedByAny(arg)
						|| shouldForceAnyCoerce(arg.t, paramType)));
				}
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
		For `ArraySort.sort`, box the typed slice to `[]any`, invoke the existing
		helper, then copy sorted values back into the original slice. For `ListSort`,
		adapt the comparator to erased `any` parameters and type-assert the returned
		head back to the expected node type.
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
			var sliceType = typeToGoType(arrayType);
			var sliceExpr = lowerExpr(args[0]).expr;
			var comparatorExpr = lowerExpr(args[1]).expr;
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
		if (!isArrayType(sourceArrayType) || arrayElementGoType(sourceArrayType) == "any") {
			return typedSliceExpr;
		}
		var sourceName = freshTempName("hx_sort_src");
		var itemName = freshTempName("hx_sort_item");
		var outName = freshTempName("hx_sort_out");
		var sourceType = typeToGoType(sourceArrayType);
		return GoExpr.GoCall(GoExpr.GoFuncLiteral([{name: sourceName, typeName: sourceType}], ["[]any"], [
			GoStmt.GoVarDecl(outName, "[]any", GoExpr.GoRaw("make([]any, 0, len(" + sourceName + "))"), true),
			GoStmt.GoRaw("for _, " + itemName + " := range " + sourceName + " {"),
			GoStmt.GoAssign(GoExpr.GoIdent(outName), GoExpr.GoCall(GoExpr.GoIdent("append"), [GoExpr.GoIdent(outName), GoExpr.GoIdent(itemName)])),
			GoStmt.GoRaw("}"),
			GoStmt.GoReturn(GoExpr.GoIdent(outName))
		]), [typedSliceExpr]);
	}

	function lowerAnyArrayCopyBack(rawSliceExpr:GoExpr, targetSliceExpr:GoExpr, targetArrayType:Type):Array<GoStmt> {
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
				GoStmt.GoRaw("for " + indexName + ", " + itemName + " := range " + rawName + " {"),
				GoStmt.GoAssign(GoExpr.GoIndex(GoExpr.GoIdent(targetName), GoExpr.GoIdent(indexName)), convertedItemExpr),
				GoStmt.GoRaw("}")
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

	function lowerMetalGoChanCall(callee:TypedExpr, args:Array<TypedExpr>, returnType:Type):Null<LoweredExpr> {
		if (!useTypedGoConcurrencySpecialization()) {
			return null;
		}

		if (isStaticCall(callee, "Go", ["go"], "newChan")) {
			noteLoweringAttempt("go.concurrency.typed", "go_chan_new", callee.pos, "Attempt typed go.Go.newChan specialization.");
			var elementEligibility = goChanElementEligibility(returnType, "Could not resolve go.Go.newChan return type for metal specialization.");
			if (!elementEligibility.eligible) {
				notePortableConcurrencyFastpathFallback(callee.pos);
				noteLoweringFallback("go.concurrency.typed", "go_chan_new_unmorphable", callee.pos,
					withEligibilityReason("Could not monomorphize go.Go.newChan return type for metal specialization.", elementEligibility));
				return null;
			}
			var elementGoType = elementEligibility.goType;
			if (elementGoType == null) {
				notePortableConcurrencyFastpathFallback(callee.pos);
				noteLoweringFallback("go.concurrency.typed", "go_chan_new_unmorphable", callee.pos,
					"Could not monomorphize go.Go.newChan return type for metal specialization.");
				return null;
			}
			requireStdlibShimGroup("go_concurrency");
			registerMetalChanElementGoType(elementGoType);
			notePortableConcurrencyFastpathHit(callee.pos);
			noteLoweringSuccess("go.concurrency.typed", "go_chan_new", callee.pos,
				'Applied typed go.Go.newChan specialization (element type: ' + elementGoType + ").");
			var buffer = args.length > 0 ? lowerExpr(args[0]).expr : GoExpr.GoIntLiteral(0);
			return {
				expr: GoExpr.GoCall(GoExpr.GoIdent(metalChanShimName("go__concurrency_newChan", elementGoType)), [buffer]),
				isStringLike: false
			};
		}

		var methodCall = asGoChanMethodCall(callee);
		if (methodCall == null) {
			return null;
		}
		var methodKind = "go_chan_method_" + methodCall.methodName;
		noteLoweringAttempt("go.concurrency.typed", methodKind, callee.pos, "Attempt typed go.Chan method specialization.");

		var elementEligibility = metalTypeEligibility(methodCall.elementType, GoMetalEligibilityRole.ChanElement,
			"Could not resolve go.Chan method element type for metal specialization.");
		if (!elementEligibility.eligible) {
			notePortableConcurrencyFastpathFallback(callee.pos);
			noteLoweringFallback("go.concurrency.typed", "go_chan_method_unmorphable", callee.pos,
				withEligibilityReason("Could not monomorphize go.Chan method call for metal specialization.", elementEligibility));
			return null;
		}
		var elementGoType = elementEligibility.goType;
		if (elementGoType == null) {
			notePortableConcurrencyFastpathFallback(callee.pos);
			noteLoweringFallback("go.concurrency.typed", "go_chan_method_unmorphable", callee.pos,
				"Could not monomorphize go.Chan method call for metal specialization.");
			return null;
		}

		requireStdlibShimGroup("go_concurrency");
		registerMetalChanElementGoType(elementGoType);
		notePortableConcurrencyFastpathHit(callee.pos);
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
					expr: GoExpr.GoCall(GoExpr.GoIdent(metalChanShimName("go__concurrency_send", elementGoType)), [channelNative, value]),
					isStringLike: false
				};
			case "trySend":
				var value = args.length > 0 ? lowerExpr(args[0]).expr : GoExpr.GoNil;
				if (args.length > 0 && elementGoType != "any" && exprBackedByAny(args[0])) {
					value = GoExpr.GoTypeAssert(value, elementGoType);
				}
				return {
					expr: GoExpr.GoCall(GoExpr.GoIdent(metalChanShimName("go__concurrency_trySend", elementGoType)), [channelNative, value]),
					isStringLike: false
				};
			case "recv":
				return {
					expr: GoExpr.GoCall(GoExpr.GoIdent(metalChanShimName("go__concurrency_recv", elementGoType)), [channelNative]),
					isStringLike: isStringType(returnType)
				};
			case "recvOr":
				var defaultValue = args.length > 0 ? lowerExpr(args[0]).expr : GoExpr.GoNil;
				if (args.length > 0 && elementGoType != "any" && exprBackedByAny(args[0])) {
					defaultValue = GoExpr.GoTypeAssert(defaultValue, elementGoType);
				}
				return {
					expr: GoExpr.GoCall(GoExpr.GoIdent(metalChanShimName("go__concurrency_recvOr", elementGoType)), [channelNative, defaultValue]),
					isStringLike: isStringType(returnType)
				};
			case "tryRecv":
				return {
					expr: GoExpr.GoCall(GoExpr.GoIdent(metalChanShimName("go__concurrency_tryRecv", elementGoType)), [channelNative]),
					isStringLike: false
				};
			case "close":
				return {
					expr: GoExpr.GoCall(GoExpr.GoIdent(metalChanShimName("go__concurrency_close", elementGoType)), [channelNative]),
					isStringLike: false
				};
			case "__hx_setBuffer":
				var buffer = args.length > 0 ? lowerExpr(args[0]).expr : GoExpr.GoIntLiteral(0);
				return {
					expr: GoExpr.GoCall(GoExpr.GoIdent(metalChanShimName("go__concurrency_setBuffer", elementGoType)), [channel, buffer]),
					isStringLike: false
				};
			case _:
				return null;
		}
	}

	function lowerMetalGoSliceCall(callee:TypedExpr, args:Array<TypedExpr>, returnType:Type):Null<LoweredExpr> {
		if (!useTypedGoCollectionsSpecialization()) {
			return null;
		}

		var methodCall = asGoSliceMethodCall(callee);
		if (methodCall == null) {
			return null;
		}
		var methodKind = "go_slice_method_" + methodCall.methodName;
		noteLoweringAttempt("go.collections.typed", methodKind, callee.pos, "Attempt typed go.Slice method specialization.");

		var elementEligibility = goSliceElementEligibility(methodCall.target.t, "Could not resolve go.Slice element type for metal specialization.");
		if (!elementEligibility.eligible) {
			noteLoweringFallback("go.collections.typed", "go_slice_method_unmorphable", callee.pos,
				withEligibilityReason("Could not monomorphize go.Slice element type for metal specialization.", elementEligibility));
			return null;
		}
		var elementGoType = elementEligibility.goType;
		if (elementGoType == null) {
			noteLoweringFallback("go.collections.typed", "go_slice_method_unmorphable", callee.pos,
				"Could not monomorphize go.Slice element type for metal specialization.");
			return null;
		}

		requireStdlibShimGroup("go_collections");
		registerMetalSliceElementGoType(elementGoType);
		noteLoweringSuccess("go.collections.typed", methodKind, callee.pos,
			'Applied typed go.Slice method specialization (element type: ' + elementGoType + ").");
		var sliceExpr = lowerExpr(methodCall.target).expr;

		switch (methodCall.methodName) {
			case "push":
				var value = args.length > 0 ? lowerExpr(args[0]).expr : GoExpr.GoNil;
				return {
					expr: GoExpr.GoCall(GoExpr.GoIdent(metalSliceShimName("go__slice_push", elementGoType)), [sliceExpr, value]),
					isStringLike: false
				};
			case "set":
				var index = args.length > 0 ? lowerExpr(args[0]).expr : GoExpr.GoIntLiteral(0);
				var value = args.length > 1 ? lowerExpr(args[1]).expr : GoExpr.GoNil;
				return {
					expr: GoExpr.GoCall(GoExpr.GoIdent(metalSliceShimName("go__slice_set", elementGoType)), [sliceExpr, index, value]),
					isStringLike: false
				};
			case "get":
				var index = args.length > 0 ? lowerExpr(args[0]).expr : GoExpr.GoIntLiteral(0);
				return {
					expr: GoExpr.GoCall(GoExpr.GoIdent(metalSliceShimName("go__slice_get", elementGoType)), [sliceExpr, index]),
					isStringLike: isStringType(returnType)
				};
			case "get_length":
				return {
					expr: GoExpr.GoCall(GoExpr.GoIdent(metalSliceShimName("go__slice_length", elementGoType)), [sliceExpr]),
					isStringLike: false
				};
			case "toArray":
				return {
					expr: GoExpr.GoCall(GoExpr.GoIdent(metalSliceShimName("go__slice_toArray", elementGoType)), [sliceExpr]),
					isStringLike: false
				};
			case _:
				return null;
		}
	}

	function lowerMetalGoMapCall(callee:TypedExpr, args:Array<TypedExpr>, returnType:Type):Null<LoweredExpr> {
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
				"Could not monomorphize go.Map key/value types for metal specialization.");
			return null;
		}
		var keyEligibility = metalTypeEligibility(pair.keyType, GoMetalEligibilityRole.MapKey, "Could not resolve go.Map key type for metal specialization.");
		if (!keyEligibility.eligible) {
			noteLoweringFallback("go.collections.typed", "go_map_method_unmorphable", callee.pos,
				withEligibilityReason("Could not monomorphize go.Map key/value types for metal specialization.", keyEligibility));
			return null;
		}
		var valueEligibility = metalTypeEligibility(pair.valueType, GoMetalEligibilityRole.MapValue,
			"Could not resolve go.Map value type for metal specialization.");
		if (!valueEligibility.eligible) {
			noteLoweringFallback("go.collections.typed", "go_map_method_unmorphable", callee.pos,
				withEligibilityReason("Could not monomorphize go.Map key/value types for metal specialization.", valueEligibility));
			return null;
		}
		var keyGoType = keyEligibility.goType;
		var valueGoType = valueEligibility.goType;
		if (keyGoType == null || valueGoType == null) {
			noteLoweringFallback("go.collections.typed", "go_map_method_unmorphable", callee.pos,
				"Could not monomorphize go.Map key/value types for metal specialization.");
			return null;
		}

		requireStdlibShimGroup("go_collections");
		registerMetalMapTypePair(keyGoType, valueGoType);
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
					expr: GoExpr.GoCall(GoExpr.GoIdent(metalMapShimName("go__map_set", keyGoType, valueGoType)), [mapExpr, keyExpr, valueExpr]),
					isStringLike: false
				};
			case "get":
				var keyExpr = args.length > 0 ? lowerExpr(args[0]).expr : GoExpr.GoNil;
				var rawCall = GoExpr.GoCall(GoExpr.GoIdent(metalMapShimName("go__map_get", keyGoType, valueGoType)), [mapExpr, keyExpr]);
				return {
					expr: lowerNullableAwareTypeAssertExpr(rawCall, returnType),
					isStringLike: isStringType(returnType)
				};
			case "exists":
				var keyExpr = args.length > 0 ? lowerExpr(args[0]).expr : GoExpr.GoNil;
				return {
					expr: GoExpr.GoCall(GoExpr.GoIdent(metalMapShimName("go__map_exists", keyGoType, valueGoType)), [mapExpr, keyExpr]),
					isStringLike: false
				};
			case _:
				return null;
		}
	}

	function lowerMetalGoResultCall(callee:TypedExpr, args:Array<TypedExpr>, returnType:Type):Null<LoweredExpr> {
		if (!useTypedGoResultSpecialization()) {
			return null;
		}

		var returnElementEligibility = goResultElementEligibility(returnType, "Could not resolve go.Result<T> return type for metal specialization.");

		if (isStaticCall(callee, "Result", ["go"], "ok") || isStaticCall(callee, "Go", ["go"], "ok")) {
			noteLoweringAttempt("go.result.typed", "go_result_static_ok", callee.pos, "Attempt typed go.Result.ok specialization.");
			if (!returnElementEligibility.eligible || returnElementEligibility.goType == null) {
				noteLoweringFallback("go.result.typed", "go_result_static_ok_unmorphable", callee.pos,
					withEligibilityReason("Could not monomorphize go.Result<T>.ok return type for metal specialization.", returnElementEligibility));
				return null;
			}
			var returnElementGoType = returnElementEligibility.goType;
			requireStdlibShimGroup("go_result");
			registerMetalResultElementGoType(returnElementGoType);
			noteLoweringSuccess("go.result.typed", "go_result_static_ok", callee.pos,
				'Applied typed go.Result.ok specialization (element type: ' + returnElementGoType + ").");
			var value = args.length > 0 ? lowerExpr(args[0]).expr : GoExpr.GoNil;
			return {
				expr: GoExpr.GoCall(GoExpr.GoIdent(metalResultShimName("go__result_ok", returnElementGoType)), [value]),
				isStringLike: false
			};
		}

		if (isStaticCall(callee, "Result", ["go"], "failure") || isStaticCall(callee, "Go", ["go"], "fail")) {
			noteLoweringAttempt("go.result.typed", "go_result_static_failure", callee.pos, "Attempt typed go.Result.failure specialization.");
			if (!returnElementEligibility.eligible || returnElementEligibility.goType == null) {
				noteLoweringFallback("go.result.typed", "go_result_static_failure_unmorphable", callee.pos,
					withEligibilityReason("Could not monomorphize go.Result<T>.failure return type for metal specialization.", returnElementEligibility));
				return null;
			}
			var returnElementGoType = returnElementEligibility.goType;
			requireStdlibShimGroup("go_result");
			registerMetalResultElementGoType(returnElementGoType);
			noteLoweringSuccess("go.result.typed", "go_result_static_failure", callee.pos,
				'Applied typed go.Result.failure specialization (element type: ' + returnElementGoType + ").");
			var message = args.length > 0 ? lowerExpr(args[0]).expr : GoExpr.GoNil;
			return {
				expr: GoExpr.GoCall(GoExpr.GoIdent(metalResultShimName("go__result_failure", returnElementGoType)), [message]),
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
			"Could not resolve go.Result<T> method receiver type for metal specialization.");
		if (!receiverEligibility.eligible || receiverEligibility.goType == null) {
			noteLoweringFallback("go.result.typed", "go_result_method_unmorphable", callee.pos,
				withEligibilityReason("Could not monomorphize go.Result<T> method receiver for metal specialization.", receiverEligibility));
			return null;
		}
		var elementGoType = receiverEligibility.goType;

		requireStdlibShimGroup("go_result");
		registerMetalResultElementGoType(elementGoType);
		noteLoweringSuccess("go.result.typed", methodKind, callee.pos, 'Applied typed go.Result method specialization (element type: ' + elementGoType + ").");
		var resultExpr = lowerExpr(methodCall.target).expr;

		switch (methodCall.methodName) {
			case "isOk":
				return {
					expr: GoExpr.GoCall(GoExpr.GoIdent(metalResultShimName("go__result_isOk", elementGoType)), [resultExpr]),
					isStringLike: false
				};
			case "isErr":
				return {
					expr: GoExpr.GoCall(GoExpr.GoIdent(metalResultShimName("go__result_isErr", elementGoType)), [resultExpr]),
					isStringLike: false
				};
			case "unwrap":
				return {
					expr: GoExpr.GoCall(GoExpr.GoIdent(metalResultShimName("go__result_unwrap", elementGoType)), [resultExpr]),
					isStringLike: isStringType(returnType)
				};
			case "error":
				return {
					expr: GoExpr.GoCall(GoExpr.GoIdent(metalResultShimName("go__result_error", elementGoType)), [resultExpr]),
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
							loweredArg = upcastIfNeeded(loweredArg, arg.t, paramType);
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
				expr: lambdaIterableLowering.anyArrayCoerce(mappedAnyExpr, returnType),
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

	function lowerLambdaStaticCall(callee:TypedExpr, args:Array<TypedExpr>, returnType:Type):Null<LoweredExpr> {
		if (isStaticCall(callee, "Lambda", [], "count") || lambdaIterableLowering.isGeneratedCall(callee, "count")) {
			var supportsOptimizedCount = args.length == 1 || (args.length == 2 && isNullLiteralExpr(args[1]));
			if (!supportsOptimizedCount) {
				return null;
			}
			var sourcePlan = lambdaIterableLowering.trySourcePlan(args[0]);
			if (sourcePlan == null) {
				var dynamicSourceExpr = lambdaIterableLowering.dynamicIterableSource(args[0]);
				return {
					expr: GoExpr.GoCall(GoExpr.GoIdent("Lambda_count"), [dynamicSourceExpr, GoExpr.GoNil]),
					isStringLike: false
				};
			}
			var sourceExpr = sourcePlan.domain == "list" ? GoExpr.GoSelector(sourcePlan.sourceExpr, "items") : sourcePlan.sourceExpr;
			return {
				expr: GoExpr.GoCall(GoExpr.GoIdent("len"), [sourceExpr]),
				isStringLike: false
			};
		}

		if (isStaticCall(callee, "Lambda", [], "empty") || lambdaIterableLowering.isGeneratedCall(callee, "empty")) {
			if (args.length != 1) {
				Context.fatalError("Lambda.empty expects exactly 1 argument", callee.pos);
			}
			var sourcePlan = lambdaIterableLowering.trySourcePlan(args[0]);
			if (sourcePlan == null) {
				var dynamicSourceExpr = lambdaIterableLowering.dynamicIterableSource(args[0]);
				return {
					expr: GoExpr.GoCall(GoExpr.GoIdent("Lambda_empty"), [dynamicSourceExpr]),
					isStringLike: false
				};
			}
			var sourceExpr = sourcePlan.domain == "list" ? GoExpr.GoSelector(sourcePlan.sourceExpr, "items") : sourcePlan.sourceExpr;
			return {
				expr: GoExpr.GoBinary("==", GoExpr.GoCall(GoExpr.GoIdent("len"), [sourceExpr]), GoExpr.GoIntLiteral(0)),
				isStringLike: false
			};
		}

		if (isStaticCall(callee, "Lambda", [], "exists") || lambdaIterableLowering.isGeneratedCall(callee, "exists")) {
			if (args.length != 2) {
				Context.fatalError("Lambda.exists expects exactly 2 arguments", callee.pos);
			}
			var sourcePlan = lambdaIterableLowering.trySourcePlan(args[0]);
			if (sourcePlan == null) {
				var dynamicSourceExpr = lambdaIterableLowering.dynamicIterableSource(args[0]);
				var predicateExpr = lowerExpr(args[1]).expr;
				var adaptedPredicateExpr = lambdaIterableLowering.predicateAnyAdapter(predicateExpr, args[1].t);
				return {
					expr: GoExpr.GoCall(GoExpr.GoIdent("Lambda_exists"), [dynamicSourceExpr, adaptedPredicateExpr]),
					isStringLike: false
				};
			}
			var elementType = sourcePlan.elementType;
			var sourceExpr = sourcePlan.sourceExpr;
			var predicateExpr = lowerExpr(args[1]).expr;
			var sourceName = freshTempName("hx_lambda_items");
			var predicateName = freshTempName("hx_lambda_predicate");
			var itemName = freshTempName("hx_lambda_item");
			var sourceType = sourcePlan.sourceType;
			var predicateType = "func(" + elementType + ") bool";
			var loopBody = switch (sourcePlan.domain) {
				case "array":
					[
						GoStmt.GoRaw("for _, " + itemName + " := range " + sourceName + " {"),
						GoStmt.GoRaw("\tif " + predicateName + "(" + itemName + ") {"),
						GoStmt.GoReturn(GoExpr.GoBoolLiteral(true)),
						GoStmt.GoRaw("\t}"),
						GoStmt.GoRaw("}")
					];
				case "list":
					if (elementType == "any") {
						[
							GoStmt.GoRaw("for _, " + itemName + " := range " + sourceName + ".items {"),
							GoStmt.GoRaw("\tif " + predicateName + "(" + itemName + ") {"),
							GoStmt.GoReturn(GoExpr.GoBoolLiteral(true)),
							GoStmt.GoRaw("\t}"),
							GoStmt.GoRaw("}")
						];
					} else {
						var itemAnyName = freshTempName("hx_lambda_item_any");
						var itemZeroName = freshTempName("hx_lambda_item_zero");
						[
							GoStmt.GoRaw("for _, " + itemAnyName + " := range " + sourceName + ".items {"),
							GoStmt.GoRaw("\tvar " + itemName + " " + elementType),
							GoStmt.GoRaw("\tif " + itemAnyName + " == nil {"),
							GoStmt.GoRaw("\t\tvar " + itemZeroName + " " + elementType),
							GoStmt.GoRaw("\t\t" + itemName + " = " + itemZeroName),
							GoStmt.GoRaw("\t} else {"),
							GoStmt.GoRaw("\t\t" + itemName + " = " + itemAnyName + ".(" + elementType + ")"),
							GoStmt.GoRaw("\t}"),
							GoStmt.GoRaw("\tif " + predicateName + "(" + itemName + ") {"),
							GoStmt.GoReturn(GoExpr.GoBoolLiteral(true)),
							GoStmt.GoRaw("\t}"),
							GoStmt.GoRaw("}")
						];
					}
				case _:
					[];
			};
			return {
				expr: GoExpr.GoCall(GoExpr.GoFuncLiteral([
					{name: sourceName, typeName: sourceType},
					{name: predicateName, typeName: predicateType}
				], ["bool"],
					loopBody.concat([GoStmt.GoReturn(GoExpr.GoBoolLiteral(false))])), [sourceExpr, predicateExpr]),
				isStringLike: false
			};
		}

		if (isStaticCall(callee, "Lambda", [], "has") || lambdaIterableLowering.isGeneratedCall(callee, "has")) {
			if (args.length != 2) {
				Context.fatalError("Lambda.has expects exactly 2 arguments", callee.pos);
			}
			requireStdlibShimGroup("stdlib_symbols");
			var sourcePlan = lambdaIterableLowering.trySourcePlan(args[0]);
			if (sourcePlan == null) {
				var dynamicSourceExpr = lambdaIterableLowering.dynamicIterableSource(args[0]);
				var needleExpr = lowerExpr(args[1]).expr;
				return {
					expr: GoExpr.GoCall(GoExpr.GoIdent("Lambda_has"), [dynamicSourceExpr, needleExpr]),
					isStringLike: false
				};
			}
			var elementType = sourcePlan.elementType;
			var sourceExpr = sourcePlan.sourceExpr;
			var needleExpr = lowerExpr(args[1]).expr;
			var sourceName = freshTempName("hx_lambda_items");
			var needleName = freshTempName("hx_lambda_needle");
			var itemName = freshTempName("hx_lambda_item");
			var sourceType = sourcePlan.sourceType;
			var loopBody = switch (sourcePlan.domain) {
				case "array":
					[
						GoStmt.GoRaw("for _, " + itemName + " := range " + sourceName + " {"),
						GoStmt.GoRaw("\tif reflect.DeepEqual(" + itemName + ", " + needleName + ") {"),
						GoStmt.GoReturn(GoExpr.GoBoolLiteral(true)),
						GoStmt.GoRaw("\t}"),
						GoStmt.GoRaw("}")
					];
				case "list":
					if (elementType == "any") {
						[
							GoStmt.GoRaw("for _, " + itemName + " := range " + sourceName + ".items {"),
							GoStmt.GoRaw("\tif reflect.DeepEqual(" + itemName + ", " + needleName + ") {"),
							GoStmt.GoReturn(GoExpr.GoBoolLiteral(true)),
							GoStmt.GoRaw("\t}"),
							GoStmt.GoRaw("}")
						];
					} else {
						var itemAnyName = freshTempName("hx_lambda_item_any");
						var itemZeroName = freshTempName("hx_lambda_item_zero");
						[
							GoStmt.GoRaw("for _, " + itemAnyName + " := range " + sourceName + ".items {"),
							GoStmt.GoRaw("\tvar " + itemName + " " + elementType),
							GoStmt.GoRaw("\tif " + itemAnyName + " == nil {"),
							GoStmt.GoRaw("\t\tvar " + itemZeroName + " " + elementType),
							GoStmt.GoRaw("\t\t" + itemName + " = " + itemZeroName),
							GoStmt.GoRaw("\t} else {"),
							GoStmt.GoRaw("\t\t" + itemName + " = " + itemAnyName + ".(" + elementType + ")"),
							GoStmt.GoRaw("\t}"),
							GoStmt.GoRaw("\tif reflect.DeepEqual(" + itemName + ", " + needleName + ") {"),
							GoStmt.GoReturn(GoExpr.GoBoolLiteral(true)),
							GoStmt.GoRaw("\t}"),
							GoStmt.GoRaw("}")
						];
					}
				case _:
					[];
			};
			return {
				expr: GoExpr.GoCall(GoExpr.GoFuncLiteral([{name: sourceName, typeName: sourceType}, {name: needleName, typeName: "any"}], ["bool"],
					loopBody.concat([GoStmt.GoReturn(GoExpr.GoBoolLiteral(false))])),
					[sourceExpr, needleExpr]),
				isStringLike: false
			};
		}

		if (isStaticCall(callee, "Lambda", [], "iter") || lambdaIterableLowering.isGeneratedCall(callee, "iter")) {
			if (args.length != 2) {
				Context.fatalError("Lambda.iter expects exactly 2 arguments", callee.pos);
			}
			var sourcePlan = lambdaIterableLowering.trySourcePlan(args[0]);
			var iteratorSourceExpr = sourcePlan == null ? lambdaIterableLowering.dynamicIterableSource(args[0]) : lambdaIterableLowering.manualIteratorProtocolSource(sourcePlan.sourceExpr,
				sourcePlan);
			var consumerExpr = lowerExpr(args[1]).expr;
			var adaptedConsumerExpr = lambdaIterableLowering.consumerAnyAdapter(consumerExpr, args[1].t);
			return {
				expr: GoExpr.GoCall(GoExpr.GoIdent("Lambda_iter"), [iteratorSourceExpr, adaptedConsumerExpr]),
				isStringLike: false
			};
		}

		if (isStaticCall(callee, "Lambda", [], "filter") || lambdaIterableLowering.isGeneratedCall(callee, "filter")) {
			if (args.length != 2) {
				Context.fatalError("Lambda.filter expects exactly 2 arguments", callee.pos);
			}
			var sourcePlan = lambdaIterableLowering.trySourcePlan(args[0]);
			if (sourcePlan == null) {
				var dynamicSourceExpr = lambdaIterableLowering.dynamicIterableSource(args[0]);
				var predicateExpr = lowerExpr(args[1]).expr;
				var adaptedPredicateExpr = lambdaIterableLowering.predicateAnyAdapter(predicateExpr, args[1].t);
				var filteredAnyExpr = GoExpr.GoCall(GoExpr.GoIdent("Lambda_filter"), [dynamicSourceExpr, adaptedPredicateExpr]);
				return {
					expr: lambdaIterableLowering.anyArrayCoerce(filteredAnyExpr, returnType),
					isStringLike: false
				};
			}
			var elementType = sourcePlan.elementType;
			var sourceExpr = sourcePlan.sourceExpr;
			var predicateExpr = lowerExpr(args[1]).expr;
			var sourceName = freshTempName("hx_lambda_items");
			var predicateName = freshTempName("hx_lambda_predicate");
			var outName = freshTempName("hx_lambda_out");
			var itemName = freshTempName("hx_lambda_item");
			var sourceType = sourcePlan.sourceType;
			var outType = "[]" + elementType;
			var predicateType = "func(" + elementType + ") bool";
			var loopBody = switch (sourcePlan.domain) {
				case "array":
					[
						GoStmt.GoRaw("for _, " + itemName + " := range " + sourceName + " {"),
						GoStmt.GoRaw("\tif " + predicateName + "(" + itemName + ") {"),
						GoStmt.GoRaw("\t\t" + outName + " = append(" + outName + ", " + itemName + ")"),
						GoStmt.GoRaw("\t}"),
						GoStmt.GoRaw("}")
					];
				case "list":
					if (elementType == "any") {
						[
							GoStmt.GoRaw("for _, " + itemName + " := range " + sourceName + ".items {"),
							GoStmt.GoRaw("\tif " + predicateName + "(" + itemName + ") {"),
							GoStmt.GoRaw("\t\t" + outName + " = append(" + outName + ", " + itemName + ")"),
							GoStmt.GoRaw("\t}"),
							GoStmt.GoRaw("}")
						];
					} else {
						var itemAnyName = freshTempName("hx_lambda_item_any");
						var itemZeroName = freshTempName("hx_lambda_item_zero");
						[
							GoStmt.GoRaw("for _, " + itemAnyName + " := range " + sourceName + ".items {"),
							GoStmt.GoRaw("\tvar " + itemName + " " + elementType),
							GoStmt.GoRaw("\tif " + itemAnyName + " == nil {"),
							GoStmt.GoRaw("\t\tvar " + itemZeroName + " " + elementType),
							GoStmt.GoRaw("\t\t" + itemName + " = " + itemZeroName),
							GoStmt.GoRaw("\t} else {"),
							GoStmt.GoRaw("\t\t" + itemName + " = " + itemAnyName + ".(" + elementType + ")"),
							GoStmt.GoRaw("\t}"),
							GoStmt.GoRaw("\tif " + predicateName + "(" + itemName + ") {"),
							GoStmt.GoRaw("\t\t" + outName + " = append(" + outName + ", " + itemName + ")"),
							GoStmt.GoRaw("\t}"),
							GoStmt.GoRaw("}")
						];
					}
				case _:
					[];
			};
			return {
				expr: GoExpr.GoCall(GoExpr.GoFuncLiteral([
					{name: sourceName, typeName: sourceType},
					{name: predicateName, typeName: predicateType}
				], [outType], [
					GoStmt.GoVarDecl(outName, outType,
						GoExpr.GoRaw("make(" + outType + ", 0, len(" + (sourcePlan.domain == "list" ? sourceName + ".items" : sourceName) + "))"), true)
				].concat(loopBody).concat([GoStmt.GoReturn(GoExpr.GoIdent(outName))])), [sourceExpr, predicateExpr]),
				isStringLike: false
			};
		}

		if (isStaticCall(callee, "Lambda", [], "map") || lambdaIterableLowering.isGeneratedCall(callee, "map")) {
			if (args.length != 2) {
				Context.fatalError("Lambda.map expects exactly 2 arguments", callee.pos);
			}
			var sourcePlan = lambdaIterableLowering.trySourcePlan(args[0]);
			if (sourcePlan == null) {
				var dynamicSourceExpr = lambdaIterableLowering.dynamicIterableSource(args[0]);
				var mapperExpr = lowerExpr(args[1]).expr;
				var adaptedMapperExpr = lambdaIterableLowering.mapperAnyAdapter(mapperExpr, args[1].t);
				var foldValueName = freshTempName("hx_lambda_value");
				var foldAccName = freshTempName("hx_lambda_acc");
				var mappedAnyExpr = GoExpr.GoCall(GoExpr.GoIdent("Lambda_fold"), [
					dynamicSourceExpr,
					GoExpr.GoFuncLiteral([{name: foldValueName, typeName: "any"}, {name: foldAccName, typeName: "any"}], ["any"], [
						GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("append"), [
							GoExpr.GoTypeAssert(GoExpr.GoIdent(foldAccName), "[]any"),
							GoExpr.GoCall(adaptedMapperExpr, [GoExpr.GoIdent(foldValueName)])
						]))
					]),
					GoExpr.GoRaw("[]any{}")
				]);
				return {
					expr: lambdaIterableLowering.anyArrayCoerce(mappedAnyExpr, returnType),
					isStringLike: false
				};
			}
			var sourceElementType = sourcePlan.elementType;
			var mappedElementType = arrayElementGoType(returnType);
			var sourceExpr = sourcePlan.sourceExpr;
			var mapperExpr = lowerExpr(args[1]).expr;
			var sourceName = freshTempName("hx_lambda_items");
			var mapperName = freshTempName("hx_lambda_mapper");
			var outName = freshTempName("hx_lambda_out");
			var itemName = freshTempName("hx_lambda_item");
			var sourceType = sourcePlan.sourceType;
			var mappedType = "[]" + mappedElementType;
			var mapperType = "func(" + sourceElementType + ") " + mappedElementType;
			var loopBody = switch (sourcePlan.domain) {
				case "array":
					[
						GoStmt.GoRaw("for _, " + itemName + " := range " + sourceName + " {"),
						GoStmt.GoRaw("\t" + outName + " = append(" + outName + ", " + mapperName + "(" + itemName + "))"),
						GoStmt.GoRaw("}")
					];
				case "list":
					if (sourceElementType == "any") {
						[
							GoStmt.GoRaw("for _, " + itemName + " := range " + sourceName + ".items {"),
							GoStmt.GoRaw("\t" + outName + " = append(" + outName + ", " + mapperName + "(" + itemName + "))"),
							GoStmt.GoRaw("}")
						];
					} else {
						var itemAnyName = freshTempName("hx_lambda_item_any");
						var itemZeroName = freshTempName("hx_lambda_item_zero");
						[
							GoStmt.GoRaw("for _, " + itemAnyName + " := range " + sourceName + ".items {"),
							GoStmt.GoRaw("\tvar " + itemName + " " + sourceElementType),
							GoStmt.GoRaw("\tif " + itemAnyName + " == nil {"),
							GoStmt.GoRaw("\t\tvar " + itemZeroName + " " + sourceElementType),
							GoStmt.GoRaw("\t\t" + itemName + " = " + itemZeroName),
							GoStmt.GoRaw("\t} else {"),
							GoStmt.GoRaw("\t\t" + itemName + " = " + itemAnyName + ".(" + sourceElementType + ")"),
							GoStmt.GoRaw("\t}"),
							GoStmt.GoRaw("\t" + outName + " = append(" + outName + ", " + mapperName + "(" + itemName + "))"),
							GoStmt.GoRaw("}")
						];
					}
				case _:
					[];
			};
			return {
				expr: GoExpr.GoCall(GoExpr.GoFuncLiteral([
					{name: sourceName, typeName: sourceType},
					{name: mapperName, typeName: mapperType}
				], [mappedType], [
					GoStmt.GoVarDecl(outName, mappedType,
						GoExpr.GoRaw("make(" + mappedType + ", 0, len(" + (sourcePlan.domain == "list" ? sourceName + ".items" : sourceName) + "))"), true)
				].concat(loopBody).concat([GoStmt.GoReturn(GoExpr.GoIdent(outName))])), [sourceExpr, mapperExpr]),
				isStringLike: false
			};
		}

		if (isStaticCall(callee, "Lambda", [], "fold") || lambdaIterableLowering.isGeneratedCall(callee, "fold")) {
			if (args.length != 3) {
				Context.fatalError("Lambda.fold expects exactly 3 arguments", callee.pos);
			}
			var sourcePlan = lambdaIterableLowering.trySourcePlan(args[0]);
			if (sourcePlan == null) {
				var dynamicSourceExpr = lambdaIterableLowering.dynamicIterableSource(args[0]);
				var folderExpr = lowerExpr(args[1]).expr;
				var initExpr = lowerExpr(args[2]).expr;
				var adaptedFolderExpr = lambdaIterableLowering.folderAnyAdapter(folderExpr, args[1].t);
				var foldedAnyExpr = GoExpr.GoCall(GoExpr.GoIdent("Lambda_fold"), [dynamicSourceExpr, adaptedFolderExpr, initExpr]);
				return {
					expr: lowerNullableAwareTypeAssertExpr(foldedAnyExpr, returnType),
					isStringLike: false
				};
			}
			var elementType = sourcePlan.elementType;
			var accType = typeToGoType(returnType);
			var sourceExpr = sourcePlan.sourceExpr;
			var folderExpr = lowerExpr(args[1]).expr;
			var initExpr = lowerExpr(args[2]).expr;
			var sourceName = freshTempName("hx_lambda_items");
			var folderName = freshTempName("hx_lambda_folder");
			var accName = freshTempName("hx_lambda_acc");
			var itemName = freshTempName("hx_lambda_item");
			var sourceType = sourcePlan.sourceType;
			var folderType = "func(" + elementType + ", " + accType + ") " + accType;
			var loopBody = switch (sourcePlan.domain) {
				case "array":
					[
						GoStmt.GoRaw("for _, " + itemName + " := range " + sourceName + " {"),
						GoStmt.GoRaw("\t" + accName + " = " + folderName + "(" + itemName + ", " + accName + ")"),
						GoStmt.GoRaw("}")
					];
				case "list":
					if (elementType == "any") {
						[
							GoStmt.GoRaw("for _, " + itemName + " := range " + sourceName + ".items {"),
							GoStmt.GoRaw("\t" + accName + " = " + folderName + "(" + itemName + ", " + accName + ")"),
							GoStmt.GoRaw("}")
						];
					} else {
						var itemAnyName = freshTempName("hx_lambda_item_any");
						var itemZeroName = freshTempName("hx_lambda_item_zero");
						[
							GoStmt.GoRaw("for _, " + itemAnyName + " := range " + sourceName + ".items {"),
							GoStmt.GoRaw("\tvar " + itemName + " " + elementType),
							GoStmt.GoRaw("\tif " + itemAnyName + " == nil {"),
							GoStmt.GoRaw("\t\tvar " + itemZeroName + " " + elementType),
							GoStmt.GoRaw("\t\t" + itemName + " = " + itemZeroName),
							GoStmt.GoRaw("\t} else {"),
							GoStmt.GoRaw("\t\t" + itemName + " = " + itemAnyName + ".(" + elementType + ")"),
							GoStmt.GoRaw("\t}"),
							GoStmt.GoRaw("\t" + accName + " = " + folderName + "(" + itemName + ", " + accName + ")"),
							GoStmt.GoRaw("}")
						];
					}
				case _:
					[];
			};
			return {
				expr: GoExpr.GoCall(GoExpr.GoFuncLiteral([
					{name: sourceName, typeName: sourceType},
					{name: folderName, typeName: folderType},
					{name: accName, typeName: accType}
				], [accType],
					loopBody.concat([GoStmt.GoReturn(GoExpr.GoIdent(accName))])), [sourceExpr, folderExpr, initExpr]),
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
		var seen = new Map<String, Bool>();
		var out = new Array<String>();
		var seed = ["[]any", "[]int", "[]float64", "[]bool", "[]*string"];
		for (typeName in seed) {
			if (!seen.exists(typeName)) {
				seen.set(typeName, true);
				out.push(typeName);
			}
		}

		for (classType in projectClasses) {
			if (!hasInstanceLayout(classType)) {
				continue;
			}
			var classArray = "[]*" + classTypeName(classType);
			if (!seen.exists(classArray)) {
				seen.set(classArray, true);
				out.push(classArray);
			}
		}

		for (enumType in projectEnums) {
			var enumArray = "[]*" + enumTypeName(enumType);
			if (!seen.exists(enumArray)) {
				seen.set(enumArray, true);
				out.push(enumArray);
			}
		}

		out.sort(function(a, b) return Reflect.compare(a, b));
		return out;
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

	function isMetalProfile():Bool {
		return compilationContext.profile == GoProfile.Metal;
	}

	function isPortableProfile():Bool {
		return compilationContext.profile == GoProfile.Portable;
	}

	function useStringFastpath():Bool {
		return compilationContext.buildContext.portableStringFastpathEnabled;
	}

	function usePortableConcurrencyFastpath():Bool {
		return isPortableProfile() && compilationContext.buildContext.portableConcurrencyFastpathEnabled;
	}

	function useAutoLoweringPlannerMode():Bool {
		return isPortableProfile() && compilationContext.buildContext.autoLoweringMode != GoAutoLoweringMode.Off;
	}

	function useTypedGoConcurrencySpecialization():Bool {
		return isMetalProfile() || usePortableConcurrencyFastpath();
	}

	function useTypedGoCollectionsSpecialization():Bool {
		return isMetalProfile() || useAutoLoweringPlannerMode();
	}

	function useTypedGoResultSpecialization():Bool {
		return isMetalProfile() || useAutoLoweringPlannerMode();
	}

	function notePortableConcurrencyFastpathHit(pos:haxe.macro.Expr.Position):Void {
		if (!usePortableConcurrencyFastpath()) {
			return;
		}
		if (isFrameworkInternalPos(pos)) {
			return;
		}
		compilationContext.optimizerPortableConcurrencyTypedFastpathHits++;
	}

	function notePortableConcurrencyFastpathFallback(pos:haxe.macro.Expr.Position):Void {
		if (!usePortableConcurrencyFastpath()) {
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
		if (shouldSuppressPortableInternalFallbackReport(feature, pos)) {
			return;
		}
		noteLoweringDecision(feature, kind, "fallback", pos, detail);
		noteOptimizerTypedLoweringFallback(feature);
		noteMetalFallback(kind, pos, detail);
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

	function shouldSuppressPortableInternalFallbackReport(feature:String, pos:haxe.macro.Expr.Position):Bool {
		if (!isPortableProfile()) {
			return false;
		}
		if (!usePortableConcurrencyFastpath()) {
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
			inMetalLane: metadata.inMetalLane
		});
	}

	function noteMetalFallback(kind:String, pos:haxe.macro.Expr.Position, detail:String):Void {
		if (!isMetalProfile()) {
			return;
		}
		var metadata = loweringDecisionMetadata(pos);
		var violation = {
			kind: kind,
			detail: detail,
			location: metadata.location,
			module: metadata.moduleName,
			inMetalLane: metadata.inMetalLane
		};
		compilationContext.metalFallbackViolations.push(violation);
		var hardError = compilationContext.buildContext.metalContractHardError && !isFrameworkInternalPos(pos);
		if (hardError) {
			Context.error("Metal contract fallback is not allowed: "
				+ detail
				+ " Use `-D reflaxe_go_metal_allow_fallback` to permit fallback for this build.", pos);
		}
	}

	function loweringDecisionMetadata(pos:haxe.macro.Expr.Position):{moduleName:String, inMetalLane:Bool, location:String} {
		var moduleName = sourceModuleRegistry.sourceModuleForPos(pos);
		var inMetalLane = compilationContext.buildContext.metalLaneModules.indexOf(moduleName) != -1;
		var location = fallbackLocationLabel(pos, moduleName);
		return {
			moduleName: moduleName,
			inMetalLane: inMetalLane,
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

	function isMonomorphizableMetalElementType(elementGoType:String):Bool {
		return elementGoType != null && elementGoType != "" && elementGoType != "any";
	}

	function isMonomorphizableMetalChanElementType(elementGoType:String):Bool {
		return isMonomorphizableMetalElementType(elementGoType);
	}

	function goChanElementEligibility(type:Type, missingMessage:String):GoMetalTypeEligibilityResult {
		return metalTypeEligibility(goChanElementType(type), GoMetalEligibilityRole.ChanElement, missingMessage);
	}

	function goChanElementGoType(type:Type):Null<String> {
		var eligibility = goChanElementEligibility(type, "Could not resolve go.Chan element type for metal specialization.");
		return eligibility.eligible ? eligibility.goType : null;
	}

	function goSliceElementEligibility(type:Type, missingMessage:String):GoMetalTypeEligibilityResult {
		return metalTypeEligibility(goSliceElementType(type), GoMetalEligibilityRole.SliceElement, missingMessage);
	}

	function goSliceElementGoType(type:Type):Null<String> {
		var eligibility = goSliceElementEligibility(type, "Could not resolve go.Slice element type for metal specialization.");
		return eligibility.eligible ? eligibility.goType : null;
	}

	function goMapTypePairGoTypes(type:Type):Null<MetalMapTypePair> {
		var pair = goMapTypePair(type);
		if (pair == null) {
			return null;
		}
		var keyEligibility = metalTypeEligibility(pair.keyType, GoMetalEligibilityRole.MapKey, "Could not resolve go.Map key type for metal specialization.");
		if (!keyEligibility.eligible || keyEligibility.goType == null) {
			return null;
		}
		var valueEligibility = metalTypeEligibility(pair.valueType, GoMetalEligibilityRole.MapValue,
			"Could not resolve go.Map value type for metal specialization.");
		if (!valueEligibility.eligible || valueEligibility.goType == null) {
			return null;
		}
		return {
			keyGoType: keyEligibility.goType,
			valueGoType: valueEligibility.goType
		};
	}

	function goResultElementEligibility(type:Type, missingMessage:String):GoMetalTypeEligibilityResult {
		return metalTypeEligibility(goResultElementType(type), GoMetalEligibilityRole.ResultElement, missingMessage);
	}

	function goResultElementGoType(type:Type):Null<String> {
		var eligibility = goResultElementEligibility(type, "Could not resolve go.Result<T> element type for metal specialization.");
		return eligibility.eligible ? eligibility.goType : null;
	}

	function metalTypeEligibility(type:Null<Type>, role:GoMetalEligibilityRole, missingMessage:String):GoMetalTypeEligibilityResult {
		if (type == null) {
			return {
				eligible: false,
				goType: null,
				reasonCode: "missing_type",
				reason: missingMessage
			};
		}
		return GoMetalTypeEligibility.resolve(type, role, classTypeName, enumTypeName);
	}

	function withEligibilityReason(base:String, eligibility:GoMetalTypeEligibilityResult):String {
		var reason = eligibility.reason;
		if (reason == null || StringTools.trim(reason) == "") {
			return base;
		}
		var prefix = StringTools.endsWith(base, ".") ? base.substr(0, base.length - 1) : base;
		return prefix + ": " + reason;
	}

	function registerMetalChanElementGoType(elementGoType:String):Void {
		if (!useTypedGoConcurrencySpecialization()) {
			return;
		}
		if (!isMonomorphizableMetalChanElementType(elementGoType)) {
			return;
		}
		requiredMetalChanElementTypes.set(elementGoType, true);
	}

	function registerMetalSliceElementGoType(elementGoType:String):Void {
		if (!useTypedGoCollectionsSpecialization()) {
			return;
		}
		if (!isMonomorphizableMetalElementType(elementGoType)) {
			return;
		}
		requiredMetalSliceElementTypes.set(elementGoType, true);
	}

	function registerMetalMapTypePair(keyGoType:String, valueGoType:String):Void {
		if (!useTypedGoCollectionsSpecialization()) {
			return;
		}
		if (!isMonomorphizableMetalElementType(keyGoType) || !isMonomorphizableMetalElementType(valueGoType)) {
			return;
		}
		var signature = metalMapTypeSignature(keyGoType, valueGoType);
		requiredMetalMapTypePairs.set(signature, {
			keyGoType: keyGoType,
			valueGoType: valueGoType
		});
	}

	function registerMetalResultElementGoType(elementGoType:String):Void {
		if (!useTypedGoResultSpecialization()) {
			return;
		}
		if (!isMonomorphizableMetalElementType(elementGoType)) {
			return;
		}
		requiredMetalResultElementTypes.set(elementGoType, true);
	}

	function metalTypeHash(value:String):String {
		var hash = 0x811C9DC5;
		for (index in 0...value.length) {
			hash ^= value.charCodeAt(index);
			hash *= 0x01000193;
		}
		return StringTools.hex(hash, 8).toLowerCase();
	}

	function metalTypeSuffix(typeKey:String):String {
		var normalized = GoNaming.normalizeIdent(typeKey);
		if (normalized == "" || normalized == "hx_tmp") {
			normalized = "t";
		}
		return normalized + "_" + metalTypeHash(typeKey);
	}

	function metalMapTypeSignature(keyGoType:String, valueGoType:String):String {
		return keyGoType + "__" + valueGoType;
	}

	function metalChanShimName(base:String, elementGoType:String):String {
		return base + "__" + metalTypeSuffix(elementGoType);
	}

	function metalSliceShimName(base:String, elementGoType:String):String {
		return base + "__" + metalTypeSuffix(elementGoType);
	}

	function metalMapShimName(base:String, keyGoType:String, valueGoType:String):String {
		return base + "__" + metalTypeSuffix(metalMapTypeSignature(keyGoType, valueGoType));
	}

	function metalResultShimName(base:String, elementGoType:String):String {
		return base + "__" + metalTypeSuffix(elementGoType);
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

	function normalizeExternStringCallResult(callee:TypedExpr, returnType:Type, callExpr:GoExpr):GoExpr {
		if (!isStringType(returnType)) {
			return callExpr;
		}
		if (!isGoImportExternCall(callee)) {
			return callExpr;
		}
		return GoExpr.GoCall(GoExpr.GoIdent("hxrt.StdString"), [callExpr]);
	}

	function normalizeExternCallArg(callee:TypedExpr, argExpr:GoExpr, paramType:Null<Type>, returnType:Type):GoExpr {
		if (paramType == null || !isExternValueErrorCall(callee, returnType)) {
			return argExpr;
		}
		if (isStringType(paramType)) {
			return GoExpr.GoUnary("*", GoExpr.GoCall(GoExpr.GoIdent("hxrt.StdString"), [argExpr]));
		}
		return argExpr;
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
			case FMethod(_):
				true;
			case _:
				false;
		};
	}

	function shouldUseVirtualDispatch(classType:ClassType, field:ClassField):Bool {
		if (!isProjectClass(classType)) {
			return false;
		}
		var className = fullClassName(classType);
		if (GoStdlibOwnership.isCompilerOwnedAuthority(className)) {
			return false;
		}
		if (className == "haxe.io.Input" || className == "haxe.io.Output") {
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

	function lowerBinop(op:Binop, left:TypedExpr, right:TypedExpr, resultType:Type):LoweredExpr {
		var leftLowered = lowerExpr(left);
		var rightLowered = lowerExpr(right);
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
		var leftExprForOperator = nullComparison ? leftLowered.expr : coerceNullablePrimitiveOperandForUse(leftLowered.expr, left);
		var rightExprForOperator = nullComparison ? rightLowered.expr : coerceNullablePrimitiveOperandForUse(rightLowered.expr, right);
		var useStringEquality = stringMode && (!nullComparison || isStringType(left.t) || isStringType(right.t));
		var typedStringOps = isStringType(left.t) && isStringType(right.t);
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
				{
					expr: GoExpr.GoCall(GoExpr.GoIdent(typedStringOps ? "hxrt.StringConcatStringPtr" : "hxrt.StringConcatAny"),
						[leftLowered.expr, rightLowered.expr]),
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
			case OpAdd | OpSub | OpMult | OpDiv if (floatMode):
				{
					expr: GoExpr.GoBinary(binopSymbol(op), floatOperandExpr(leftLowered.expr, left.t, left),
						floatOperandExpr(rightLowered.expr, right.t, right)),
					isStringLike: false
				};
			case OpMod if (floatMode):
				{
					expr: GoExpr.GoCall(GoExpr.GoIdent("hxrt.FloatMod"), [
						floatOperandExpr(leftLowered.expr, left.t, left),
						floatOperandExpr(rightLowered.expr, right.t, right)
					]),
					isStringLike: false
				};
			case OpUShr if (int32Mode):
				var int32Left = coerceNullableIntOperandExpr(leftLowered.expr, left.t, left);
				var int32Right = coerceNullableIntOperandExpr(rightLowered.expr, right.t, right);
				{
					expr: lowerHaxeInt32BinopExpr(op, int32Left, int32Right),
					isStringLike: false
				};
			case OpAdd | OpSub | OpMult | OpMod | OpAnd | OpOr | OpXor | OpShl | OpShr if (int32Mode):
				var int32Left = coerceNullableIntOperandExpr(leftLowered.expr, left.t, left);
				var int32Right = coerceNullableIntOperandExpr(rightLowered.expr, right.t, right);
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

	function lowerAssignOpExpr(op:Binop, leftExpr:GoExpr, rightExpr:GoExpr, leftType:Type, rightType:Type, ?sourcePos:haxe.macro.Expr.Position):GoExpr {
		if (op == OpAssign) {
			return rightExpr;
		}
		if (op == OpAdd && (isStringType(leftType) || isStringType(rightType))) {
			var typedStringOps = isStringType(leftType) && isStringType(rightType);
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
		return coerceAnyExprToType(expr, operand.t, operand.t, exprBackedByAny(operand));
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

	function wrapInt32Expr(expr:GoExpr):GoExpr {
		return GoExprOperatorOps.wrapInt32Expr(expr);
	}

	function floatOperandExpr(expr:GoExpr, operandType:Type, ?operand:TypedExpr):GoExpr {
		if (isNullableFloatType(operandType)) {
			return coerceNullableFloatOperandExpr(expr, operandType, operand);
		}
		return GoExprOperatorOps.floatOperandExpr(expr, isFloatType(operandType));
	}

	function unitStepExpr(target:GoExpr, opSymbol:String, valueType:Type, ?sourcePos:haxe.macro.Expr.Position):GoExpr {
		return GoExprOperatorOps.unitStepExpr(target, opSymbol, isInt32SemanticType(valueType, sourcePos));
	}

	function binopSymbol(op:Binop):String {
		return GoExprOperatorOps.binopSymbol(op);
	}

	function unopSymbol(op:Unop):String {
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

	function upcastIfNeeded(expr:GoExpr, fromType:Type, toType:Type):GoExpr {
		var fromClass = classFromType(fromType);
		var toClass = classFromType(toType);
		if (fromClass == null || toClass == null) {
			return expr;
		}
		var toClassName = fullClassName(toClass);
		if (toClassName == "haxe.io.Input" || toClassName == "haxe.io.Output") {
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

	function typeToGoType(type:Type):String {
		return GoTypeMapper.typeToGoType(type, classTypeName, enumTypeName);
	}

	function valueStorageGoType(type:Type):String {
		// Nullable primitive expression temps must be able to hold Go nil; keep the
		// broader type mapper unchanged so signatures and eligibility stay stable.
		return isNullablePrimitiveType(type) ? "any" : typeToGoType(type);
	}

	function isStringType(type:Type):Bool {
		return GoTypeMapper.isStringType(type);
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

	function arrayElementGoType(type:Type):String {
		return GoTypeMapper.arrayElementGoType(type, classTypeName, enumTypeName);
	}

	function scalarGoType(type:Type):String {
		return GoTypeMapper.scalarGoType(type, classTypeName, enumTypeName);
	}

	function goFunctionType(args:Array<{name:String, opt:Bool, t:Type}>, returnType:Type):String {
		return GoTypeMapper.goFunctionType(args, returnType, classTypeName, enumTypeName);
	}

	function restElementType(type:Type):Null<Type> {
		return GoTypeMapper.restElementType(type);
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
							return GoExpr.GoArrayLiteral(arrayElementGoType(right.t), [for (value in values) lowerExpr(value).expr]);
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

	function isHaxeIoBaseClass(classType:ClassType):Bool {
		return GoTypeMapper.isHaxeIoBaseClass(classType);
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
		if (group == "http") {
			sourceOwnedStdlibPlanner.requireSourceOwnedStdlibClass("sys.GoHttpHelpers");
		} else if (group == "ds") {
			sourceOwnedStdlibPlanner.requireSourceOwnedStdlibModule("haxe.Constraints");
		}
	}

	function noteIoHelperFieldUsage(classType:ClassType, fieldName:String):Void {
		if (GoStdlibShimClassifier.needsIoHelperSurface(classType, fieldName, isIoInputHelperMethodName, isIoOutputHelperMethodName)) {
			requireIoSourceOwnedHelperSurface();
		}
	}

	function noteStaticStdlibFieldUsage(classType:ClassType, fieldName:String):Void {
		if (classType.pack.length == 0 && classType.name == "Sys" && (fieldName == "command" || fieldName == "exit")) {
			requiresSysCommandSurface = true;
		}
		if (classType.pack.length == 0 && classType.name == "Sys" && fieldName == "environment") {
			requireStdlibShimGroup("ds");
		}
		sourceOwnedStdlibPlanner.noteSourceOwnedStdlibUsage(classType);
	}

	function noteSourceOwnedStdlibUsage(classType:ClassType):Void {
		sourceOwnedStdlibPlanner.noteSourceOwnedStdlibUsage(classType);
	}

	function requireIoSourceOwnedHelperSurface():Void {
		requiresIoHelperSurface = true;
		sourceOwnedStdlibPlanner.requireIoSourceOwnedHelperClass();
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

	function isBytesOwnedBySourceStdDecl(decl:GoDecl):Bool {
		return switch (decl) {
			case GoDecl.GoStructDecl(name, _):
				name == "haxe__io__Bytes";
			case GoDecl.GoFuncDecl(name, receiver, _, _, _): (receiver != null && receiver.typeName == "*haxe__io__Bytes") || name == "New_haxe__io__Bytes" || StringTools.startsWith(name,
					"haxe__io__Bytes_");
			case _:
				false;
		};
	}

	function isEnumValueMapOwnedBySourceStdDecl(decl:GoDecl):Bool {
		return switch (decl) {
			case GoDecl.GoStructDecl(name, _):
				name == "haxe__ds__EnumValueMap";
			case GoDecl.GoFuncDecl(name, receiver, _, _, _): (receiver != null
					&& receiver.typeName == "*haxe__ds__EnumValueMap") || name == "New_haxe__ds__EnumValueMap";
			case _:
				false;
		};
	}

	function noteStdlibClass(classType:ClassType):Void {
		if (classType.pack.join(".") == "haxe.io") {
			switch (classType.name) {
				case "StringInput":
					requiresIoStringInputSurface = true;
					requireIoSourceOwnedHelperSurface();
				case "BufferInput":
					requiresIoBufferInputSurface = true;
					requireIoSourceOwnedHelperSurface();
				case "Eof":
					requiresIoEofStringSurface = true;
				case _:
			}
		}
		if (classType.pack.join(".") == "sys.net" && classType.name == "UdpSocket") {
			requiresUdpSocketSurface = true;
		}
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
		var classPaths = [for (classType in projectClasses) fullClassName(classType)];
		classPaths.sort(Reflect.compare);
		var enumPaths = [for (enumType in projectEnums) fullEnumName(enumType)];
		enumPaths.sort(Reflect.compare);
		return GoHxrtFeatureAnalyzer.inferWithReasons(classPaths, enumPaths, requiredShimGroups, requiresIoHelperSurface);
	}

	function resetExternImportPaths():Void {
		for (path in externImportPaths.keys()) {
			externImportPaths.remove(path);
		}
		for (path in externImportPackages.keys()) {
			externImportPackages.remove(path);
		}
	}

	function resetRequiredMetalChanElementTypes():Void {
		for (elementType in requiredMetalChanElementTypes.keys()) {
			requiredMetalChanElementTypes.remove(elementType);
		}
	}

	function resetRequiredMetalSliceElementTypes():Void {
		for (elementType in requiredMetalSliceElementTypes.keys()) {
			requiredMetalSliceElementTypes.remove(elementType);
		}
	}

	function resetRequiredMetalMapTypePairs():Void {
		for (signature in requiredMetalMapTypePairs.keys()) {
			requiredMetalMapTypePairs.remove(signature);
		}
	}

	function resetRequiredMetalResultElementTypes():Void {
		for (elementType in requiredMetalResultElementTypes.keys()) {
			requiredMetalResultElementTypes.remove(elementType);
		}
	}

	function noteExternImportPath(classType:ClassType, packageName:String):Void {
		var path = externClassImportPath(classType);
		if (path == null || path == "") {
			return;
		}
		externImportPaths.set(path, true);
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

	function lowerArrayInstanceCall(callee:TypedExpr, args:Array<TypedExpr>, returnType:Type):Null<LoweredExpr> {
		var methodCall = asArrayMethodCall(callee);
		if (methodCall == null || !isArrayType(methodCall.target.t)) {
			return null;
		}

		return switch (methodCall.methodName) {
			case "copy" if (args.length == 0):
				{
					expr: cloneArrayExpr(lowerExpr(methodCall.target).expr, methodCall.target.t),
					isStringLike: false
				};
			case "push":
				var site = lowerArrayMutationSite(methodCall.target);
				var appendArgs = [site.tempExpr];
				var shouldMaskToByte = isBytesBufferStorageArray(methodCall.target);
				for (arg in args) {
					var appendValue = lowerExpr(arg).expr;
					if (shouldMaskToByte) {
						appendValue = GoExpr.GoBinary("&", appendValue, GoExpr.GoIntLiteral(255));
					}
					appendArgs.push(appendValue);
				}
				var body = site.prefix.concat([
					GoStmt.GoAssign(site.tempExpr, GoExpr.GoCall(GoExpr.GoIdent("append"), appendArgs))
				]).concat(site.writeBack(site.tempExpr)).concat([GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("len"), [site.tempExpr]))]);
				{
					expr: GoExpr.GoCall(GoExpr.GoFuncLiteral([], ["int"], body), []),
					isStringLike: false
				};
			case "pop" if (args.length == 0):
				var site = lowerArrayMutationSite(methodCall.target);
				var lenName = freshTempName("hx_len");
				var valueName = freshTempName("hx_value");
				var zeroName = freshTempName("hx_zero");
				var resultType = typeToGoType(returnType);
				var body = site.prefix.concat([
					GoStmt.GoVarDecl(lenName, "int", GoExpr.GoCall(GoExpr.GoIdent("len"), [site.tempExpr]), true),
					GoStmt.GoIf(GoExpr.GoBinary("==", GoExpr.GoIdent(lenName), GoExpr.GoIntLiteral(0)), [
						GoStmt.GoVarDecl(zeroName, resultType, null, false),
						GoStmt.GoReturn(GoExpr.GoIdent(zeroName))
					],
						null),
					GoStmt.GoVarDecl(valueName, resultType,
						GoExpr.GoIndex(site.tempExpr, GoExpr.GoBinary("-", GoExpr.GoIdent(lenName), GoExpr.GoIntLiteral(1))), true),
					GoStmt.GoAssign(site.tempExpr, GoExpr.GoSlice(site.tempExpr, null, GoExpr.GoBinary("-", GoExpr.GoIdent(lenName), GoExpr.GoIntLiteral(1))))
				]).concat(site.writeBack(site.tempExpr)).concat([GoStmt.GoReturn(GoExpr.GoIdent(valueName))]);
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
		return GoNaming.staticSymbol(classType.pack, classType.name, fieldName, fullClassName(classType) == "Main");
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
