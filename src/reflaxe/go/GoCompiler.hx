package reflaxe.go;

#if macro
import haxe.macro.Context;
import haxe.macro.Expr;
import haxe.macro.Expr.Binop;
import haxe.macro.Expr.Unop;
import haxe.macro.PositionTools;
import haxe.macro.Type;
import reflaxe.go.compiler.GoAutoLoweringMode;
import reflaxe.go.compiler.GoExprOperatorOps;
import reflaxe.go.compiler.GoHxrtFeatureAnalyzer;
import reflaxe.go.compiler.GoMetalTypeEligibility;
import reflaxe.go.compiler.GoMetalTypeEligibility.GoMetalEligibilityRole;
import reflaxe.go.compiler.GoMetalTypeEligibility.GoMetalTypeEligibilityResult;
import reflaxe.go.compiler.GoStdlibShimClassifier;
import reflaxe.go.compiler.GoTestAstFixtureEmitter;
import reflaxe.go.compiler.GoTypeMapper;
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
	final superHaxeTypeName:Null<String>;
	final staticFieldNames:Array<String>;
	final instanceFieldNames:Array<String>;
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

private typedef LambdaSourcePlan = {
	final domain:String;
	final elementType:String;
	final sourceExpr:GoExpr;
	final sourceType:String;
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
	final sourceFileToModule:Map<String, String>;
	final sourceModuleBySuffix:Map<String, String>;
	final functionVarNameScopes:Array<Map<Int, String>>;
	final functionVarNameCountScopes:Array<Map<String, Int>>;
	final optionalPrimitiveParamScopes:Array<Map<Int, Bool>>;
	final functionReturnTypeScopes:Array<Type>;
	final returnRedirectScopes:Array<Null<ReturnRedirect>>;
	var sourceModuleSuffixes:Array<String>;
	var cachedVoidType:Null<Type>;
	var requiresIoHelperSurface:Bool;
	var projectClasses:Array<ClassType>;
	var projectEnums:Array<EnumType>;
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
		sourceFileToModule = new Map<String, String>();
		sourceModuleBySuffix = new Map<String, String>();
		functionVarNameScopes = [];
		functionVarNameCountScopes = [];
		optionalPrimitiveParamScopes = [];
		functionReturnTypeScopes = [];
		returnRedirectScopes = [];
		sourceModuleSuffixes = [];
		cachedVoidType = null;
		requiresIoHelperSurface = false;
		projectClasses = [];
		projectEnums = [];
		globalLeafReceiverTypes = new Map<String, Bool>();
		tempVarCounter = 0;
		requiresTypeValueSupport = false;
		#end
	}

	#if macro
	public function compileModule(types:Array<ModuleType>):Array<GoGeneratedFile> {
		return compileResolvedTypes(collectProjectClasses(types), collectProjectEnums(types));
	}

	public function compileSelectedTypes(classes:Array<ClassType>, enums:Array<EnumType>):Array<GoGeneratedFile> {
		return compileResolvedTypes(normalizeProjectClasses(classes), normalizeProjectEnums(enums));
	}

	function compileResolvedTypes(classes:Array<ClassType>, enums:Array<EnumType>):Array<GoGeneratedFile> {
		projectClasses = classes.copy();
		projectEnums = enums.copy();
		rebuildSourceModuleLookup(classes, enums);
		globalLeafReceiverTypes = buildGlobalLeafReceiverTypes(projectClasses);
		syncCompilationContextLeafReceivers();
		clearBoolMap(compilationContext.leafReturningFunctions);
		requiresIoHelperSurface = false;
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
		for (classType in classes) {
			appendModuleDecls(moduleDecls, classType.module, lowerClassDecls(classType));
		}
		applyStdlibShimGroupDependencies();

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
		compilationContext.inferredHxrtFeatures = inferRuntimeFeatures(requiredShimGroups);

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
			if (classType.isExtern || classType.isInterface || !classHasInstanceLayout(classType)) {
				continue;
			}
			var superName:Null<String> = null;
			if (classType.superClass != null) {
				superName = fullClassName(classType.superClass.t.get());
			}
			entries.push({
				goTypeName: classTypeName(classType),
				haxeTypeName: fullClassName(classType),
				constructorSymbol: constructorSymbol(classType),
				superHaxeTypeName: superName,
				staticFieldNames: collectClassStaticFieldNames(classType),
				instanceFieldNames: collectClassInstanceFieldNames(classType)
			});
		}
		entries.sort(function(a, b) return Reflect.compare(a.goTypeName, b.goTypeName));
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
		return isProjectClass(superType) ? superType : null;
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

	function rebuildSourceModuleLookup(classes:Array<ClassType>, enums:Array<EnumType>):Void {
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

	function sourceModuleForPos(pos:haxe.macro.Expr.Position):String {
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

	static function pathEndsWithSuffix(path:String, suffix:String):Bool {
		if (path == suffix) {
			return true;
		}
		return path != "" && suffix != "" && StringTools.endsWith(path, "/" + suffix);
	}

	static function compareSuffixBySpecificity(a:String, b:String):Int {
		if (a.length != b.length) {
			return b.length - a.length;
		}
		return Reflect.compare(a, b);
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
		if (requiredStdlibShimGroups.exists("http")) {
			// Http request shims expose and consume haxe.io.Bytes payloads.
			requireStdlibShimGroup("io");
			requireStdlibShimGroup("ds");
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
			GoDecl.GoStructDecl("haxe__io__BytesInput", [
				{name: "bigEndian", typeName: "bool"},
				{name: "b", typeName: "[]int"},
				{name: "pos", typeName: "int"},
				{name: "len", typeName: "int"},
				{name: "totlen", typeName: "int"}
			]),
			GoDecl.GoStructDecl("haxe__io__BytesOutput", [
				{name: "bigEndian", typeName: "bool"},
				{name: "b", typeName: "*haxe__io__BytesBuffer"}
			]),
			GoDecl.GoFuncDecl("New_haxe__io__Input", null, [], ["haxe__io__Input"],
				[
					GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("New_haxe__io__BytesInput"), [GoExpr.GoRaw("&haxe__io__Bytes{b: []int{}, length: 0}")]))
				]),
			GoDecl.GoFuncDecl("New_haxe__io__Output", null, [], ["haxe__io__Output"],
				[GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("New_haxe__io__BytesOutput"), []))]),
			GoDecl.GoFuncDecl("New_haxe__io__Eof", null, [], ["*haxe__io__Eof"], [GoStmt.GoReturn(GoExpr.GoRaw("&haxe__io__Eof{}"))]),
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
				GoStmt.GoRaw("raw := *hxrt.StdString(s)"),
				GoStmt.GoRaw("lenValue := len(raw)"),
				GoStmt.GoRaw("if (lenValue & 1) != 0 {"),
				GoStmt.GoRaw("\thxrt.Throw(hxrt.StringFromLiteral(\"Not a hex string (odd number of digits)\"))"),
				GoStmt.GoRaw("\treturn &haxe__io__Bytes{b: []int{}, length: 0}"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("ret := haxe__io__Bytes_alloc(lenValue >> 1)"),
				GoStmt.GoRaw("for i := 0; i < ret.length; i++ {"),
				GoStmt.GoRaw("\thigh := int(raw[i*2])"),
				GoStmt.GoRaw("\tlow := int(raw[i*2+1])"),
				GoStmt.GoRaw("\thigh = (high & 0xF) + ((high & 0x40) >> 6) * 9"),
				GoStmt.GoRaw("\tlow = (low & 0xF) + ((low & 0x40) >> 6) * 9"),
				GoStmt.GoRaw("\tret.set(i, ((high << 4) | low) & 0xFF)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoReturn(GoExpr.GoIdent("ret"))
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
				GoStmt.GoRaw("hexChars := \"0123456789abcdef\""),
				GoStmt.GoRaw("out := make([]byte, self.length*2)"),
				GoStmt.GoRaw("for i := 0; i < self.length; i++ {"),
				GoStmt.GoRaw("\tc := self.b[i] & 0xFF"),
				GoStmt.GoRaw("\tout[i*2] = hexChars[c>>4]"),
				GoStmt.GoRaw("\tout[i*2+1] = hexChars[c&15]"),
				GoStmt.GoRaw("}"),
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoRaw("string(out)")]))
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
				GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("self"), "b"), GoExpr.GoCall(GoExpr.GoIdent("append"), [
					GoExpr.GoSelector(GoExpr.GoIdent("self"), "b"),
					GoExpr.GoBinary("&", GoExpr.GoIdent("value"), GoExpr.GoIntLiteral(255))
				]))
			]),
			GoDecl.GoFuncDecl("add", {
				name: "self",
				typeName: "*haxe__io__BytesBuffer"
			}, [{name: "src", typeName: "*haxe__io__Bytes"}], [], [
				GoStmt.GoIf(GoExpr.GoBinary("==", GoExpr.GoIdent("src"), GoExpr.GoNil), [GoStmt.GoReturn(null)], null),
				GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("self"), "b"), GoExpr.GoRaw("append(self.b, src.b...)"))
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
				GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("self"), "b"), GoExpr.GoRaw("append(self.b, src.b[pos:pos+len]...)"))
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
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("len"), [GoExpr.GoSelector(GoExpr.GoIdent("self"), "b")]))
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
				GoStmt.GoRaw("if self == nil {"),
				GoStmt.GoRaw("\treturn &haxe__io__Bytes{b: []int{}, length: 0}"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("resolved := 1 << 14"),
				GoStmt.GoRaw("if len(bufsize) > 0 {"),
				GoStmt.GoRaw("\tresolved = bufsize[0]"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("buf := haxe__io__Bytes_alloc(resolved)"),
				GoStmt.GoRaw("total := New_haxe__io__BytesBuffer()"),
				GoStmt.GoRaw("for {"),
				GoStmt.GoRaw("\tchunk := 0"),
				GoStmt.GoRaw("\tthrew := false"),
				GoStmt.GoRaw("\tvar thrown any"),
				GoStmt.GoRaw("\tfunc() {"),
				GoStmt.GoRaw("\t\tdefer func() {"),
				GoStmt.GoRaw("\t\t\tif recovered := recover(); recovered != nil {"),
				GoStmt.GoRaw("\t\t\t\tthrew = true"),
				GoStmt.GoRaw("\t\t\t\tthrown = hxrt.UnwrapException(recovered)"),
				GoStmt.GoRaw("\t\t\t}"),
				GoStmt.GoRaw("\t\t}()"),
				GoStmt.GoRaw("\t\tchunk = self.readBytes(buf, 0, resolved)"),
				GoStmt.GoRaw("\t}()"),
				GoStmt.GoRaw("\tif threw {"),
				GoStmt.GoRaw("\t\tif haxe__io__input_isEof(thrown) {"),
				GoStmt.GoRaw("\t\t\tbreak"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t\thxrt.Throw(thrown)"),
				GoStmt.GoRaw("\t\treturn &haxe__io__Bytes{b: []int{}, length: 0}"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tif chunk == 0 {"),
				GoStmt.GoRaw("\t\thxrt.Throw(haxe__io__Error_Blocked)"),
				GoStmt.GoRaw("\t\treturn &haxe__io__Bytes{b: []int{}, length: 0}"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\ttotal.addBytes(buf, 0, chunk)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("total"), "getBytes"), []))
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
				GoStmt.GoRaw("if self == nil {"),
				GoStmt.GoRaw("\thxrt.Throw(haxe__io__Error_Blocked)"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("for len > 0 {"),
				GoStmt.GoRaw("\tk := self.readBytes(s, pos, len)"),
				GoStmt.GoRaw("\tif k == 0 {"),
				GoStmt.GoRaw("\t\thxrt.Throw(haxe__io__Error_Blocked)"),
				GoStmt.GoRaw("\t\treturn"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tpos += k"),
				GoStmt.GoRaw("\tlen -= k"),
				GoStmt.GoRaw("}")
			]),
			GoDecl.GoFuncDecl("haxe__io__input_read", null, [
				{
					name: "self",
					typeName: "haxe__io__Input"
				},
				{name: "nbytes", typeName: "int"}
			], ["*haxe__io__Bytes"], [
				GoStmt.GoRaw("s := haxe__io__Bytes_alloc(nbytes)"),
				GoStmt.GoRaw("p := 0"),
				GoStmt.GoRaw("for nbytes > 0 {"),
				GoStmt.GoRaw("\tk := self.readBytes(s, p, nbytes)"),
				GoStmt.GoRaw("\tif k == 0 {"),
				GoStmt.GoRaw("\t\thxrt.Throw(haxe__io__Error_Blocked)"),
				GoStmt.GoRaw("\t\treturn &haxe__io__Bytes{b: []int{}, length: 0}"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tp += k"),
				GoStmt.GoRaw("\tnbytes -= k"),
				GoStmt.GoRaw("}"),
				GoStmt.GoReturn(GoExpr.GoIdent("s"))
			]),
			GoDecl.GoFuncDecl("haxe__io__input_readUntil", null, [
				{
					name: "self",
					typeName: "haxe__io__Input"
				},
				{name: "end", typeName: "int"}
			], ["*string"], [
				GoStmt.GoRaw("buf := New_haxe__io__BytesBuffer()"),
				GoStmt.GoRaw("for {"),
				GoStmt.GoRaw("\tlast := self.readByte()"),
				GoStmt.GoRaw("\tif last == end {"),
				GoStmt.GoRaw("\t\tbreak"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tbuf.addByte(last)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("buf"), "getBytes"), []), "toString"), []))
			]),
			GoDecl.GoFuncDecl("haxe__io__input_readLine", null, [
				{
					name: "self",
					typeName: "haxe__io__Input"
				}
			], ["*string"], [
				GoStmt.GoRaw("buf := New_haxe__io__BytesBuffer()"),
				GoStmt.GoRaw("for {"),
				GoStmt.GoRaw("\tlast := 0"),
				GoStmt.GoRaw("\tthrew := false"),
				GoStmt.GoRaw("\tvar thrown any"),
				GoStmt.GoRaw("\tfunc() {"),
				GoStmt.GoRaw("\t\tdefer func() {"),
				GoStmt.GoRaw("\t\t\tif recovered := recover(); recovered != nil {"),
				GoStmt.GoRaw("\t\t\t\tthrew = true"),
				GoStmt.GoRaw("\t\t\t\tthrown = hxrt.UnwrapException(recovered)"),
				GoStmt.GoRaw("\t\t\t}"),
				GoStmt.GoRaw("\t\t}()"),
				GoStmt.GoRaw("\t\tlast = self.readByte()"),
				GoStmt.GoRaw("\t}()"),
				GoStmt.GoRaw("\tif threw {"),
				GoStmt.GoRaw("\t\tif haxe__io__input_isEof(thrown) {"),
				GoStmt.GoRaw("\t\t\ts := buf.getBytes().toString()"),
				GoStmt.GoRaw("\t\t\traw := *hxrt.StdString(s)"),
				GoStmt.GoRaw("\t\t\tif len(raw) == 0 {"),
				GoStmt.GoRaw("\t\t\t\thxrt.Throw(thrown)"),
				GoStmt.GoRaw("\t\t\t\treturn hxrt.StringFromLiteral(\"\")"),
				GoStmt.GoRaw("\t\t\t}"),
				GoStmt.GoRaw("\t\t\treturn s"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t\thxrt.Throw(thrown)"),
				GoStmt.GoRaw("\t\treturn hxrt.StringFromLiteral(\"\")"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tif last == 10 {"),
				GoStmt.GoRaw("\t\tbreak"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tbuf.addByte(last)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("s := buf.getBytes().toString()"),
				GoStmt.GoRaw("raw := *hxrt.StdString(s)"),
				GoStmt.GoRaw("if len(raw) > 0 && raw[len(raw)-1] == 13 {"),
				GoStmt.GoRaw("\traw = raw[:len(raw)-1]"),
				GoStmt.GoRaw("}"),
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoIdent("raw")]))
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
				GoStmt.GoRaw("if self == nil || s == nil {"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("remaining := s.length"),
				GoStmt.GoRaw("p := 0"),
				GoStmt.GoRaw("for remaining > 0 {"),
				GoStmt.GoRaw("\tk := self.writeBytes(s, p, remaining)"),
				GoStmt.GoRaw("\tif k == 0 {"),
				GoStmt.GoRaw("\t\thxrt.Throw(haxe__io__Error_Blocked)"),
				GoStmt.GoRaw("\t\treturn"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tp += k"),
				GoStmt.GoRaw("\tremaining -= k"),
				GoStmt.GoRaw("}")
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
				GoStmt.GoRaw("for len > 0 {"),
				GoStmt.GoRaw("\tk := self.writeBytes(s, pos, len)"),
				GoStmt.GoRaw("\tpos += k"),
				GoStmt.GoRaw("\tlen -= k"),
				GoStmt.GoRaw("}")
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
				GoStmt.GoRaw("if self == nil || i == nil {"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("resolved := 4096"),
				GoStmt.GoRaw("if len(bufsize) > 0 {"),
				GoStmt.GoRaw("\tresolved = bufsize[0]"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("buf := haxe__io__Bytes_alloc(resolved)"),
				GoStmt.GoRaw("for {"),
				GoStmt.GoRaw("\tlenRead := 0"),
				GoStmt.GoRaw("\tthrew := false"),
				GoStmt.GoRaw("\tvar thrown any"),
				GoStmt.GoRaw("\tfunc() {"),
				GoStmt.GoRaw("\t\tdefer func() {"),
				GoStmt.GoRaw("\t\t\tif recovered := recover(); recovered != nil {"),
				GoStmt.GoRaw("\t\t\t\tthrew = true"),
				GoStmt.GoRaw("\t\t\t\tthrown = hxrt.UnwrapException(recovered)"),
				GoStmt.GoRaw("\t\t\t}"),
				GoStmt.GoRaw("\t\t}()"),
				GoStmt.GoRaw("\t\tlenRead = i.readBytes(buf, 0, resolved)"),
				GoStmt.GoRaw("\t}()"),
				GoStmt.GoRaw("\tif threw {"),
				GoStmt.GoRaw("\t\tif haxe__io__input_isEof(thrown) {"),
				GoStmt.GoRaw("\t\t\tbreak"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t\thxrt.Throw(thrown)"),
				GoStmt.GoRaw("\t\treturn"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tif lenRead == 0 {"),
				GoStmt.GoRaw("\t\thxrt.Throw(haxe__io__Error_Blocked)"),
				GoStmt.GoRaw("\t\treturn"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tp := 0"),
				GoStmt.GoRaw("\tfor lenRead > 0 {"),
				GoStmt.GoRaw("\t\tk := self.writeBytes(buf, p, lenRead)"),
				GoStmt.GoRaw("\t\tif k == 0 {"),
				GoStmt.GoRaw("\t\t\thxrt.Throw(haxe__io__Error_Blocked)"),
				GoStmt.GoRaw("\t\t\treturn"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t\tp += k"),
				GoStmt.GoRaw("\t\tlenRead -= k"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("}")
			]),
			GoDecl.GoFuncDecl("haxe__io__output_writeString", null, [
				{
					name: "self",
					typeName: "haxe__io__Output"
				},
				{name: "s", typeName: "*string"},
				{name: "encoding", typeName: "...*haxe__io__Encoding"}
			], [], [
				GoStmt.GoRaw("if s == nil {"),
				GoStmt.GoRaw("\ts = hxrt.StringFromLiteral(\"\")"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("b := haxe__io__Bytes_ofString(s, encoding...)"),
				GoStmt.GoRaw("self.writeFullBytes(b, 0, b.length)")
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
		if (!requiresIoHelperSurface) {
			decls = trimIoShimToCoreSurface(decls);
		}
		return decls;
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
					if (receiver != null && receiver.typeName == "*haxe__io__BytesInput" && isIoInputHelperMethodName(name)) {
						continue;
					}
					if (receiver != null && receiver.typeName == "*haxe__io__BytesOutput" && isIoOutputHelperMethodName(name)) {
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

	function isIoOutputHelperMethodName(name:String):Bool {
		return switch (name) {
			case "write", "writeFullBytes", "writeFloat", "writeDouble", "writeInt8", "writeInt16", "writeUInt16", "writeInt24", "writeUInt24", "writeInt32",
				"prepare", "writeInput", "writeString":
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
		return [
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
			GoDecl.GoFuncDecl("copy", {
				name: "self",
				typeName: "*haxe__ds__IntMap"
			}, [], ["*haxe__ds__IntMap"], [
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
			GoDecl.GoFuncDecl("copy", {
				name: "self",
				typeName: "*haxe__ds__StringMap"
			}, [], ["*haxe__ds__StringMap"], [
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
			GoDecl.GoFuncDecl("copy", {
				name: "self",
				typeName: "*haxe__ds__ObjectMap"
			}, [], ["*haxe__ds__ObjectMap"], [
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
			GoDecl.GoFuncDecl("copy", {
				name: "self",
				typeName: "*haxe__ds__EnumValueMap"
			}, [], ["*haxe__ds__EnumValueMap"], [
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
				GoStmt.GoIf(GoExpr.GoBinary("==", GoExpr.GoIdent("self"), GoExpr.GoNil), [GoStmt.GoReturn(GoExpr.GoNil)], null),
				GoStmt.GoVarDecl("rawKey", null, GoExpr.GoRaw("*hxrt.StdString(key)"), true),
				GoStmt.GoVarDecl("normalized", null, GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("strings"), "ToLower"), [GoExpr.GoIdent("rawKey")]), true),
				GoStmt.GoRaw("if self.responseHeadersSameKey != nil {"),
				GoStmt.GoRaw("\tif values, ok := self.responseHeadersSameKey[rawKey]; ok {"),
				GoStmt.GoRaw("\t\treturn values"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tif values, ok := self.responseHeadersSameKey[normalized]; ok {"),
				GoStmt.GoRaw("\t\treturn values"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("}"),
				GoStmt.GoIf(GoExpr.GoBinary("==", GoExpr.GoSelector(GoExpr.GoIdent("self"), "responseHeaders"), GoExpr.GoNil),
					[GoStmt.GoReturn(GoExpr.GoNil)], null),
				GoStmt.GoVarDecl("single", null, GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoSelector(GoExpr.GoIdent("self"), "responseHeaders"), "get"),
					[
						GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoIdent("rawKey")])
					]),
					true),
				GoStmt.GoRaw("if single == nil && rawKey != normalized {"),
				GoStmt.GoRaw("\tsingle = self.responseHeaders.get(hxrt.StringFromLiteral(normalized))"),
				GoStmt.GoRaw("}"),
				GoStmt.GoIf(GoExpr.GoBinary("==", GoExpr.GoIdent("single"), GoExpr.GoNil), [GoStmt.GoReturn(GoExpr.GoNil)], null),
				GoStmt.GoReturn(GoExpr.GoRaw("[]*string{hxrt.StdString(single)}"))
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
				GoStmt.GoRaw("if api == nil || payload == nil {"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("switch out := api.(type) {"),
				GoStmt.GoRaw("case *haxe__io__BytesBuffer:"),
				GoStmt.GoRaw("\tout.add(payload)"),
				GoStmt.GoRaw("case interface{ add(*haxe__io__Bytes) }:"),
				GoStmt.GoRaw("\tout.add(payload)"),
				GoStmt.GoRaw("case interface{ writeBytes(*haxe__io__Bytes, int, int) int }:"),
				GoStmt.GoRaw("\tout.writeBytes(payload, 0, payload.length)"),
				GoStmt.GoRaw("case interface{ writeFullBytes(*haxe__io__Bytes, int, int) }:"),
				GoStmt.GoRaw("\tout.writeFullBytes(payload, 0, payload.length)"),
				GoStmt.GoRaw("case interface{ writeString(*string) }:"),
				GoStmt.GoRaw("\tout.writeString(payload.toString())"),
				GoStmt.GoRaw("}"),
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
		return [
			GoDecl.GoStructDecl("Sys", []),
			GoDecl.GoStructDecl("sys__io__File", []),
			GoDecl.GoStructDecl("sys__io__ProcessOutput", [{name: "impl", typeName: "*hxrt.ProcessOutput"}]),
			GoDecl.GoStructDecl("sys__io__Process", [
				{name: "impl", typeName: "*hxrt.Process"},
				{name: "stdout", typeName: "*sys__io__ProcessOutput"}
			]),
			GoDecl.GoFuncDecl("Sys_getCwd", null, [], ["*string"], [
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("hxrt"), "SysGetCwd"), []))
			]),
			GoDecl.GoFuncDecl("Sys_args", null, [], ["[]*string"], [
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("hxrt"), "SysArgs"), []))
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
			GoDecl.GoStructDecl("StringTools", []),
			GoDecl.GoFuncDecl("StringTools_trim", null, [{name: "value", typeName: "*string"}], ["*string"], [
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [
					GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("strings"), "TrimSpace"), [GoExpr.GoRaw("*hxrt.StdString(value)")])
				]))
			]),
			GoDecl.GoFuncDecl("StringTools_startsWith", null, [
				{
					name: "value",
					typeName: "*string"
				},
				{name: "prefix", typeName: "*string"}
			], ["bool"], [
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("strings"), "HasPrefix"),
					[GoExpr.GoRaw("*hxrt.StdString(value)"), GoExpr.GoRaw("*hxrt.StdString(prefix)")]))
			]),
			GoDecl.GoFuncDecl("StringTools_replace", null, [
				{
					name: "value",
					typeName: "*string"
				},
				{name: "sub", typeName: "*string"},
				{name: "by", typeName: "*string"}
			], ["*string"], [
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [
					GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("strings"), "ReplaceAll"), [
						GoExpr.GoRaw("*hxrt.StdString(value)"),
						GoExpr.GoRaw("*hxrt.StdString(sub)"),
						GoExpr.GoRaw("*hxrt.StdString(by)")
					])
				]))
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
			GoDecl.GoFuncDecl("getHours", {
				name: "self",
				typeName: "*Date"
			}, [], ["int"], [
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoSelector(GoExpr.GoIdent("self"), "value"), "Hour"), []))
			]),
			GoDecl.GoFuncDecl("getTime", {
				name: "self",
				typeName: "*Date"
			},
				[], ["float64"], [GoStmt.GoReturn(GoExpr.GoRaw("float64(self.value.UnixNano()) / 1e6"))]),
			GoDecl.GoStructDecl("DateTools", []),
			GoDecl.GoFuncDecl("DateTools_format", null, [
				{
					name: "date",
					typeName: "*Date"
				},
				{name: "format", typeName: "*string"}
			], ["*string"], [
				GoStmt.GoRaw("layout := *hxrt.StdString(format)"),
				GoStmt.GoRaw("layout = strings.ReplaceAll(layout, \"%%\", \"__HX_PERCENT__\")"),
				GoStmt.GoRaw("layout = strings.ReplaceAll(layout, \"%Y\", \"2006\")"),
				GoStmt.GoRaw("layout = strings.ReplaceAll(layout, \"%m\", \"01\")"),
				GoStmt.GoRaw("layout = strings.ReplaceAll(layout, \"%d\", \"02\")"),
				GoStmt.GoRaw("layout = strings.ReplaceAll(layout, \"%H\", \"15\")"),
				GoStmt.GoRaw("layout = strings.ReplaceAll(layout, \"%M\", \"04\")"),
				GoStmt.GoRaw("layout = strings.ReplaceAll(layout, \"%S\", \"05\")"),
				GoStmt.GoRaw("layout = strings.ReplaceAll(layout, \"__HX_PERCENT__\", \"%\")"),
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [
					GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoSelector(GoExpr.GoIdent("date"), "value"), "Format"), [GoExpr.GoIdent("layout")])
				]))
			]),
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
			GoDecl.GoStructDecl("Xml", [
				{
					name: "raw",
					typeName: "*string"
				}
			]),
			GoDecl.GoFuncDecl("Xml_parse", null, [{name: "source", typeName: "*string"}], ["*Xml"], [
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("haxe__xml__Parser_parse"), [GoExpr.GoIdent("source")]))
			]),
			GoDecl.GoFuncDecl("toString", {
				name: "self",
				typeName: "*Xml"
			}, [], ["*string"],
				[
					GoStmt.GoIf(GoExpr.GoRaw("self == nil || self.raw == nil"), [
						GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoStringLiteral("")]))
					],
						null),
					GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoRaw("*self.raw")]))
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
			GoDecl.GoStructDecl("haxe__ds__BalancedTree", []),
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
			GoDecl.GoStructDecl("haxe__io__Path", [
				{
					name: "dir",
					typeName: "*string"
				},
				{name: "file", typeName: "*string"},
				{name: "ext", typeName: "*string"},
				{name: "backslash", typeName: "bool"}
			]),
			GoDecl.GoFuncDecl("New_haxe__io__Path", null, [{name: "path", typeName: "*string"}], ["*haxe__io__Path"], [
				GoStmt.GoVarDecl("raw", null, GoExpr.GoRaw("*hxrt.StdString(path)"), true),
				GoStmt.GoVarDecl("dir", null, GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("filepath"), "Dir"), [GoExpr.GoIdent("raw")]), true),
				GoStmt.GoIf(GoExpr.GoBinary("==", GoExpr.GoIdent("dir"), GoExpr.GoStringLiteral(".")),
					[GoStmt.GoAssign(GoExpr.GoIdent("dir"), GoExpr.GoStringLiteral(""))], null),
				GoStmt.GoVarDecl("base", null, GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("filepath"), "Base"), [GoExpr.GoIdent("raw")]), true),
				GoStmt.GoVarDecl("dotExt", null, GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("filepath"), "Ext"), [GoExpr.GoIdent("base")]), true),
				GoStmt.GoVarDecl("file", null, GoExpr.GoIdent("base"), true),
				GoStmt.GoIf(GoExpr.GoBinary("!=", GoExpr.GoIdent("dotExt"), GoExpr.GoStringLiteral("")), [
					GoStmt.GoAssign(GoExpr.GoIdent("file"),
						GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("strings"), "TrimSuffix"), [GoExpr.GoIdent("base"), GoExpr.GoIdent("dotExt")]))
				],
					null),
				GoStmt.GoVarDecl("ext", null,
					GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("strings"), "TrimPrefix"), [GoExpr.GoIdent("dotExt"), GoExpr.GoStringLiteral(".")]), true),
				GoStmt.GoReturn(GoExpr.GoRaw("&haxe__io__Path{dir: hxrt.StringFromLiteral(dir), file: hxrt.StringFromLiteral(file), ext: hxrt.StringFromLiteral(ext), backslash: strings.Contains(raw, \"\\\\\")}"))
			]),
			GoDecl.GoFuncDecl("haxe__io__Path_join", null, [
				{
					name: "parts",
					typeName: "[]*string"
				}
			], ["*string"],
				[
					GoStmt.GoIf(GoExpr.GoBinary("==", GoExpr.GoCall(GoExpr.GoIdent("len"), [GoExpr.GoIdent("parts")]), GoExpr.GoIntLiteral(0)), [
						GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoStringLiteral("")]))
					],
						null),
					GoStmt.GoVarDecl("joined", null, GoExpr.GoRaw("filepath.ToSlash(filepath.Join(hxrt.StringSlice(parts)...))"), true),
					GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoIdent("joined")]))
				]),
			GoDecl.GoStructDecl("haxe__io__StringInput", []),
			GoDecl.GoStructDecl("haxe__xml__Parser", []),
			GoDecl.GoStructDecl("haxe__xml__Printer", []),
			GoDecl.GoFuncDecl("haxe__xml__Parser_parse", null, [
				{
					name: "source",
					typeName: "*string"
				},
				{name: "strict", typeName: "...bool"}
			], ["*Xml"], [
				GoStmt.GoVarDecl("raw", null, GoExpr.GoRaw("*hxrt.StdString(source)"), true),
				GoStmt.GoRaw("decoder := xml.NewDecoder(strings.NewReader(raw))"),
				GoStmt.GoRaw("for {"),
				GoStmt.GoRaw("\t_, err := decoder.Token()"),
				GoStmt.GoRaw("\tif err == io.EOF {"),
				GoStmt.GoRaw("\t\tbreak"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tif err != nil {"),
				GoStmt.GoRaw("\t\thxrt.Throw(err)"),
				GoStmt.GoRaw("\t\treturn &Xml{raw: hxrt.StringFromLiteral(\"\")}"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("}"),
				GoStmt.GoReturn(GoExpr.GoRaw("&Xml{raw: hxrt.StringFromLiteral(raw)}"))
			]),
			GoDecl.GoFuncDecl("haxe__xml__Printer_print", null, [
				{
					name: "value",
					typeName: "*Xml"
				},
				{name: "pretty", typeName: "...bool"}
			], ["*string"], [
				GoStmt.GoIf(GoExpr.GoRaw("value == nil || value.raw == nil"), [
					GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoStringLiteral("")]))
				],
					null),
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoRaw("*value.raw")]))
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
		decls = decls.concat(lowerTypeReflectionShimDecls());
		return decls;
	}

	function lowerTypeReflectionShimDecls():Array<GoDecl> {
		var classMetadata = typeReflectionClassMetadata();
		var enumMetadata = typeReflectionEnumMetadata();
		enumMetadata.push({
			goTypeName: "ValueType",
			haxeTypeName: "ValueType",
			constructors: [
				{
					name: "TNull",
					index: 0,
					symbol: "ValueType_TNull",
					arity: 0
				},
				{
					name: "TInt",
					index: 1,
					symbol: "ValueType_TInt",
					arity: 0
				},
				{
					name: "TFloat",
					index: 2,
					symbol: "ValueType_TFloat",
					arity: 0
				},
				{
					name: "TBool",
					index: 3,
					symbol: "ValueType_TBool",
					arity: 0
				},
				{
					name: "TObject",
					index: 4,
					symbol: "ValueType_TObject",
					arity: 0
				},
				{
					name: "TFunction",
					index: 5,
					symbol: "ValueType_TFunction",
					arity: 0
				},
				{
					name: "TClass",
					index: 6,
					symbol: "ValueType_TClass",
					arity: 1
				},
				{
					name: "TEnum",
					index: 7,
					symbol: "ValueType_TEnum",
					arity: 1
				},
				{
					name: "TUnknown",
					index: 8,
					symbol: "ValueType_TUnknown",
					arity: 0
				}
			]
		});
		enumMetadata.sort(function(a, b) return Reflect.compare(a.goTypeName, b.goTypeName));
		var classResolveBody = [
			GoStmt.GoRaw("if name == nil {"),
			GoStmt.GoRaw("\treturn nil"),
			GoStmt.GoRaw("}")
		];
		classResolveBody.push(GoStmt.GoRaw("rawName := *hxrt.StdString(name)"));
		classResolveBody.push(GoStmt.GoRaw("switch rawName {"));
		for (entry in classMetadata) {
			classResolveBody.push(GoStmt.GoRaw("case " + goRawQuotedString(entry.haxeTypeName) + ":"));
			classResolveBody.push(GoStmt.GoRaw("\treturn &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}"));
		}
		classResolveBody.push(GoStmt.GoRaw("default:"));
		classResolveBody.push(GoStmt.GoRaw("\treturn nil"));
		classResolveBody.push(GoStmt.GoRaw("}"));

		var enumResolveBody = [
			GoStmt.GoRaw("if name == nil {"),
			GoStmt.GoRaw("\treturn nil"),
			GoStmt.GoRaw("}")
		];
		enumResolveBody.push(GoStmt.GoRaw("rawName := *hxrt.StdString(name)"));
		enumResolveBody.push(GoStmt.GoRaw("switch rawName {"));
		for (entry in enumMetadata) {
			enumResolveBody.push(GoStmt.GoRaw("case " + goRawQuotedString(entry.haxeTypeName) + ":"));
			enumResolveBody.push(GoStmt.GoRaw("\treturn &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral(rawName)}"));
		}
		enumResolveBody.push(GoStmt.GoRaw("default:"));
		enumResolveBody.push(GoStmt.GoRaw("\treturn nil"));
		enumResolveBody.push(GoStmt.GoRaw("}"));

		var classCreateBody = [GoStmt.GoRaw("switch className {")];
		for (entry in classMetadata) {
			classCreateBody.push(GoStmt.GoRaw("case " + goRawQuotedString(entry.haxeTypeName) + ":"));
			classCreateBody.push(GoStmt.GoRaw("\treturn hxrt_typeCallAny(" + entry.constructorSymbol + ", args)"));
		}
		classCreateBody.push(GoStmt.GoRaw("default:"));
		classCreateBody.push(GoStmt.GoRaw("\treturn nil, false"));
		classCreateBody.push(GoStmt.GoRaw("}"));

		var enumCreateBody = [GoStmt.GoRaw("switch enumName {")];
		for (entry in enumMetadata) {
			enumCreateBody.push(GoStmt.GoRaw("case " + goRawQuotedString(entry.haxeTypeName) + ":"));
			enumCreateBody.push(GoStmt.GoRaw("\tif useIndex {"));
			enumCreateBody.push(GoStmt.GoRaw("\t\tswitch constructorIndex {"));
			for (constructor in entry.constructors) {
				enumCreateBody.push(GoStmt.GoRaw("\t\tcase " + constructor.index + ":"));
				enumCreateBody.push(GoStmt.GoRaw("\t\t\tif len(args) != " + constructor.arity + " {"));
				enumCreateBody.push(GoStmt.GoRaw("\t\t\t\treturn nil, false"));
				enumCreateBody.push(GoStmt.GoRaw("\t\t\t}"));
				if (constructor.arity == 0) {
					enumCreateBody.push(GoStmt.GoRaw("\t\t\treturn " + constructor.symbol + ", true"));
				} else {
					enumCreateBody.push(GoStmt.GoRaw("\t\t\treturn hxrt_typeCallAny(" + constructor.symbol + ", args)"));
				}
			}
			enumCreateBody.push(GoStmt.GoRaw("\t\tdefault:"));
			enumCreateBody.push(GoStmt.GoRaw("\t\t\treturn nil, false"));
			enumCreateBody.push(GoStmt.GoRaw("\t\t}"));
			enumCreateBody.push(GoStmt.GoRaw("\t}"));
			enumCreateBody.push(GoStmt.GoRaw("\tswitch constructorName {"));
			for (constructor in entry.constructors) {
				enumCreateBody.push(GoStmt.GoRaw("\tcase " + goRawQuotedString(constructor.name) + ":"));
				enumCreateBody.push(GoStmt.GoRaw("\t\tif len(args) != " + constructor.arity + " {"));
				enumCreateBody.push(GoStmt.GoRaw("\t\t\treturn nil, false"));
				enumCreateBody.push(GoStmt.GoRaw("\t\t}"));
				if (constructor.arity == 0) {
					enumCreateBody.push(GoStmt.GoRaw("\t\treturn " + constructor.symbol + ", true"));
				} else {
					enumCreateBody.push(GoStmt.GoRaw("\t\treturn hxrt_typeCallAny(" + constructor.symbol + ", args)"));
				}
			}
			enumCreateBody.push(GoStmt.GoRaw("\tdefault:"));
			enumCreateBody.push(GoStmt.GoRaw("\t\treturn nil, false"));
			enumCreateBody.push(GoStmt.GoRaw("\t}"));
		}
		enumCreateBody.push(GoStmt.GoRaw("default:"));
		enumCreateBody.push(GoStmt.GoRaw("\treturn nil, false"));
		enumCreateBody.push(GoStmt.GoRaw("}"));

		var enumConstructorBody = new Array<GoStmt>();
		if (enumMetadata.length == 0) {
			enumConstructorBody = [
				GoStmt.GoRaw("if hxrt.AnyEqualsNull(e) {"),
				GoStmt.GoRaw("\treturn nil"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("return nil")
			];
		} else {
			enumConstructorBody = [
				GoStmt.GoRaw("if hxrt.AnyEqualsNull(e) {"),
				GoStmt.GoRaw("\treturn nil"),
				GoStmt.GoRaw("}")
			];
			enumConstructorBody.push(GoStmt.GoRaw("switch value := e.(type) {"));
			for (entry in enumMetadata) {
				enumConstructorBody.push(GoStmt.GoRaw("case *" + entry.goTypeName + ":"));
				enumConstructorBody.push(GoStmt.GoRaw("\tif value == nil {"));
				enumConstructorBody.push(GoStmt.GoRaw("\t\treturn nil"));
				enumConstructorBody.push(GoStmt.GoRaw("\t}"));
				enumConstructorBody.push(GoStmt.GoRaw("\tswitch value.tag {"));
				for (constructor in entry.constructors) {
					enumConstructorBody.push(GoStmt.GoRaw("\tcase " + constructor.index + ":"));
					enumConstructorBody.push(GoStmt.GoRaw("\t\treturn hxrt.StringFromLiteral(" + goRawQuotedString(constructor.name) + ")"));
				}
				enumConstructorBody.push(GoStmt.GoRaw("\tdefault:"));
				enumConstructorBody.push(GoStmt.GoRaw("\t\treturn nil"));
				enumConstructorBody.push(GoStmt.GoRaw("\t}"));
			}
			enumConstructorBody.push(GoStmt.GoRaw("default:"));
			enumConstructorBody.push(GoStmt.GoRaw("\treturn nil"));
			enumConstructorBody.push(GoStmt.GoRaw("}"));
		}

		var enumIndexBody = new Array<GoStmt>();
		if (enumMetadata.length == 0) {
			enumIndexBody = [
				GoStmt.GoRaw("if hxrt.AnyEqualsNull(e) {"),
				GoStmt.GoRaw("\treturn -1"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("return -1")
			];
		} else {
			enumIndexBody = [
				GoStmt.GoRaw("if hxrt.AnyEqualsNull(e) {"),
				GoStmt.GoRaw("\treturn -1"),
				GoStmt.GoRaw("}")
			];
			enumIndexBody.push(GoStmt.GoRaw("switch value := e.(type) {"));
			for (entry in enumMetadata) {
				enumIndexBody.push(GoStmt.GoRaw("case *" + entry.goTypeName + ":"));
				enumIndexBody.push(GoStmt.GoRaw("\tif value == nil {"));
				enumIndexBody.push(GoStmt.GoRaw("\t\treturn -1"));
				enumIndexBody.push(GoStmt.GoRaw("\t}"));
				enumIndexBody.push(GoStmt.GoRaw("\treturn value.tag"));
			}
			enumIndexBody.push(GoStmt.GoRaw("default:"));
			enumIndexBody.push(GoStmt.GoRaw("\treturn -1"));
			enumIndexBody.push(GoStmt.GoRaw("}"));
		}

		var enumParametersBody = new Array<GoStmt>();
		if (enumMetadata.length == 0) {
			enumParametersBody = [
				GoStmt.GoRaw("if hxrt.AnyEqualsNull(e) {"),
				GoStmt.GoRaw("\treturn []any{}"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("return []any{}")
			];
		} else {
			enumParametersBody = [
				GoStmt.GoRaw("if hxrt.AnyEqualsNull(e) {"),
				GoStmt.GoRaw("\treturn []any{}"),
				GoStmt.GoRaw("}")
			];
			enumParametersBody.push(GoStmt.GoRaw("switch value := e.(type) {"));
			for (entry in enumMetadata) {
				enumParametersBody.push(GoStmt.GoRaw("case *" + entry.goTypeName + ":"));
				enumParametersBody.push(GoStmt.GoRaw("\tif value == nil || value.params == nil {"));
				enumParametersBody.push(GoStmt.GoRaw("\t\treturn []any{}"));
				enumParametersBody.push(GoStmt.GoRaw("\t}"));
				enumParametersBody.push(GoStmt.GoRaw("\tout := make([]any, len(value.params))"));
				enumParametersBody.push(GoStmt.GoRaw("\tcopy(out, value.params)"));
				enumParametersBody.push(GoStmt.GoRaw("\treturn out"));
			}
			enumParametersBody.push(GoStmt.GoRaw("default:"));
			enumParametersBody.push(GoStmt.GoRaw("\treturn []any{}"));
			enumParametersBody.push(GoStmt.GoRaw("}"));
		}

		var getClassBody = [
			GoStmt.GoRaw("if hxrt.AnyEqualsNull(o) {"),
			GoStmt.GoRaw("\treturn nil"),
			GoStmt.GoRaw("}")
		];
		getClassBody.push(GoStmt.GoRaw("switch value := o.(type) {"));
		getClassBody.push(GoStmt.GoRaw("case *hxrt__TypeClassValue:"));
		getClassBody.push(GoStmt.GoRaw("\tif value == nil {"));
		getClassBody.push(GoStmt.GoRaw("\t\treturn nil"));
		getClassBody.push(GoStmt.GoRaw("\t}"));
		getClassBody.push(GoStmt.GoRaw("\treturn value"));
		getClassBody.push(GoStmt.GoRaw("case hxrt__TypeClassValue:"));
		getClassBody.push(GoStmt.GoRaw("\tcopyValue := value"));
		getClassBody.push(GoStmt.GoRaw("\treturn &copyValue"));
		for (entry in classMetadata) {
			getClassBody.push(GoStmt.GoRaw("case *" + entry.goTypeName + ":"));
			getClassBody.push(GoStmt.GoRaw("\tif value == nil {"));
			getClassBody.push(GoStmt.GoRaw("\t\treturn nil"));
			getClassBody.push(GoStmt.GoRaw("\t}"));
			getClassBody.push(GoStmt.GoRaw("\treturn &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(" + goRawQuotedString(entry.haxeTypeName) + ")}"));
		}
		getClassBody.push(GoStmt.GoRaw("default:"));
		getClassBody.push(GoStmt.GoRaw("\treturn nil"));
		getClassBody.push(GoStmt.GoRaw("}"));

		var getEnumBody = [
			GoStmt.GoRaw("if hxrt.AnyEqualsNull(o) {"),
			GoStmt.GoRaw("\treturn nil"),
			GoStmt.GoRaw("}")
		];
		getEnumBody.push(GoStmt.GoRaw("switch value := o.(type) {"));
		getEnumBody.push(GoStmt.GoRaw("case *hxrt__TypeEnumValue:"));
		getEnumBody.push(GoStmt.GoRaw("\tif value == nil {"));
		getEnumBody.push(GoStmt.GoRaw("\t\treturn nil"));
		getEnumBody.push(GoStmt.GoRaw("\t}"));
		getEnumBody.push(GoStmt.GoRaw("\treturn value"));
		getEnumBody.push(GoStmt.GoRaw("case hxrt__TypeEnumValue:"));
		getEnumBody.push(GoStmt.GoRaw("\tcopyValue := value"));
		getEnumBody.push(GoStmt.GoRaw("\treturn &copyValue"));
		for (entry in enumMetadata) {
			getEnumBody.push(GoStmt.GoRaw("case *" + entry.goTypeName + ":"));
			getEnumBody.push(GoStmt.GoRaw("\tif value == nil {"));
			getEnumBody.push(GoStmt.GoRaw("\t\treturn nil"));
			getEnumBody.push(GoStmt.GoRaw("\t}"));
			getEnumBody.push(GoStmt.GoRaw("\treturn &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral(" + goRawQuotedString(entry.haxeTypeName) + ")}"));
		}
		getEnumBody.push(GoStmt.GoRaw("default:"));
		getEnumBody.push(GoStmt.GoRaw("\treturn nil"));
		getEnumBody.push(GoStmt.GoRaw("}"));

		var getSuperClassBody = [
			GoStmt.GoRaw("className, ok := hxrt_typeResolvedClassName(c)"),
			GoStmt.GoRaw("if !ok {"),
			GoStmt.GoRaw("\treturn nil"),
			GoStmt.GoRaw("}"),
			GoStmt.GoRaw("switch className {")
		];
		for (entry in classMetadata) {
			getSuperClassBody.push(GoStmt.GoRaw("case " + goRawQuotedString(entry.haxeTypeName) + ":"));
			if (entry.superHaxeTypeName == null) {
				getSuperClassBody.push(GoStmt.GoRaw("\treturn nil"));
			} else {
				getSuperClassBody.push(GoStmt.GoRaw("\treturn &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("
					+ goRawQuotedString(entry.superHaxeTypeName)
					+ ")}"));
			}
		}
		getSuperClassBody.push(GoStmt.GoRaw("default:"));
		getSuperClassBody.push(GoStmt.GoRaw("\treturn nil"));
		getSuperClassBody.push(GoStmt.GoRaw("}"));

		var getClassFieldsBody = [
			GoStmt.GoRaw("className, ok := hxrt_typeResolvedClassName(c)"),
			GoStmt.GoRaw("if !ok {"),
			GoStmt.GoRaw("\treturn []*string{}"),
			GoStmt.GoRaw("}"),
			GoStmt.GoRaw("switch className {")
		];
		for (entry in classMetadata) {
			getClassFieldsBody.push(GoStmt.GoRaw("case " + goRawQuotedString(entry.haxeTypeName) + ":"));
			getClassFieldsBody.push(GoStmt.GoRaw("\treturn " + goStringPointerArrayLiteral(entry.staticFieldNames)));
		}
		getClassFieldsBody.push(GoStmt.GoRaw("default:"));
		getClassFieldsBody.push(GoStmt.GoRaw("\treturn []*string{}"));
		getClassFieldsBody.push(GoStmt.GoRaw("}"));

		var getInstanceFieldsBody = [
			GoStmt.GoRaw("className, ok := hxrt_typeResolvedClassName(c)"),
			GoStmt.GoRaw("if !ok {"),
			GoStmt.GoRaw("\treturn []*string{}"),
			GoStmt.GoRaw("}"),
			GoStmt.GoRaw("switch className {")
		];
		for (entry in classMetadata) {
			getInstanceFieldsBody.push(GoStmt.GoRaw("case " + goRawQuotedString(entry.haxeTypeName) + ":"));
			getInstanceFieldsBody.push(GoStmt.GoRaw("\treturn " + goStringPointerArrayLiteral(entry.instanceFieldNames)));
		}
		getInstanceFieldsBody.push(GoStmt.GoRaw("default:"));
		getInstanceFieldsBody.push(GoStmt.GoRaw("\treturn []*string{}"));
		getInstanceFieldsBody.push(GoStmt.GoRaw("}"));

		var getEnumConstructsBody = [
			GoStmt.GoRaw("enumName, ok := hxrt_typeResolvedEnumName(e)"),
			GoStmt.GoRaw("if !ok {"),
			GoStmt.GoRaw("\treturn []*string{}"),
			GoStmt.GoRaw("}"),
			GoStmt.GoRaw("switch enumName {")
		];
		for (entry in enumMetadata) {
			getEnumConstructsBody.push(GoStmt.GoRaw("case " + goRawQuotedString(entry.haxeTypeName) + ":"));
			var constructorNames = [for (constructor in entry.constructors) constructor.name];
			getEnumConstructsBody.push(GoStmt.GoRaw("\treturn " + goStringPointerArrayLiteral(constructorNames)));
		}
		getEnumConstructsBody.push(GoStmt.GoRaw("default:"));
		getEnumConstructsBody.push(GoStmt.GoRaw("\treturn []*string{}"));
		getEnumConstructsBody.push(GoStmt.GoRaw("}"));

		var classCreateEmptyBody = [GoStmt.GoRaw("switch className {")];
		for (entry in classMetadata) {
			classCreateEmptyBody.push(GoStmt.GoRaw("case " + goRawQuotedString(entry.haxeTypeName) + ":"));
			classCreateEmptyBody.push(GoStmt.GoRaw("\treturn &" + entry.goTypeName + "{}, true"));
		}
		classCreateEmptyBody.push(GoStmt.GoRaw("default:"));
		classCreateEmptyBody.push(GoStmt.GoRaw("\treturn nil, false"));
		classCreateEmptyBody.push(GoStmt.GoRaw("}"));

		var allEnumsBody = [
			GoStmt.GoRaw("enumName, ok := hxrt_typeResolvedEnumName(e)"),
			GoStmt.GoRaw("if !ok {"),
			GoStmt.GoRaw("\treturn []any{}"),
			GoStmt.GoRaw("}"),
			GoStmt.GoRaw("switch enumName {")
		];
		for (entry in enumMetadata) {
			allEnumsBody.push(GoStmt.GoRaw("case " + goRawQuotedString(entry.haxeTypeName) + ":"));
			var zeroAritySymbols = [
				for (constructor in entry.constructors)
					if (constructor.arity == 0) constructor.symbol
			];
			if (zeroAritySymbols.length == 0) {
				allEnumsBody.push(GoStmt.GoRaw("\treturn []any{}"));
			} else {
				allEnumsBody.push(GoStmt.GoRaw("\treturn []any{" + zeroAritySymbols.join(", ") + "}"));
			}
		}
		allEnumsBody.push(GoStmt.GoRaw("default:"));
		allEnumsBody.push(GoStmt.GoRaw("\treturn []any{}"));
		allEnumsBody.push(GoStmt.GoRaw("}"));

		var typeOfBody = [
			GoStmt.GoRaw("if hxrt.AnyEqualsNull(v) {"),
			GoStmt.GoRaw("\treturn ValueType_TNull"),
			GoStmt.GoRaw("}"),
			GoStmt.GoRaw("if enumValue := Type_getEnum(v); enumValue != nil {"),
			GoStmt.GoRaw("\treturn ValueType_TEnum(enumValue)"),
			GoStmt.GoRaw("}"),
			GoStmt.GoRaw("if classValue := Type_getClass(v); classValue != nil {"),
			GoStmt.GoRaw("\treturn ValueType_TClass(classValue)"),
			GoStmt.GoRaw("}"),
			GoStmt.GoRaw("switch v.(type) {"),
			GoStmt.GoRaw("case bool:"),
			GoStmt.GoRaw("\treturn ValueType_TBool"),
			GoStmt.GoRaw("case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr:"),
			GoStmt.GoRaw("\treturn ValueType_TInt"),
			GoStmt.GoRaw("case float32, float64:"),
			GoStmt.GoRaw("\treturn ValueType_TFloat"),
			GoStmt.GoRaw("case string, *string:"),
			GoStmt.GoRaw("\treturn ValueType_TClass(&hxrt__TypeClassValue{name: hxrt.StringFromLiteral(\"String\")})"),
			GoStmt.GoRaw("}"),
			GoStmt.GoRaw("ref := reflect.ValueOf(v)"),
			GoStmt.GoRaw("if !ref.IsValid() {"),
			GoStmt.GoRaw("\treturn ValueType_TNull"),
			GoStmt.GoRaw("}"),
			GoStmt.GoRaw("switch ref.Kind() {"),
			GoStmt.GoRaw("case reflect.Func:"),
			GoStmt.GoRaw("\treturn ValueType_TFunction"),
			GoStmt.GoRaw("case reflect.Slice, reflect.Array:"),
			GoStmt.GoRaw("\treturn ValueType_TClass(&hxrt__TypeClassValue{name: hxrt.StringFromLiteral(\"Array\")})"),
			GoStmt.GoRaw("case reflect.Map, reflect.Struct, reflect.Interface, reflect.Pointer:"),
			GoStmt.GoRaw("\treturn ValueType_TObject"),
			GoStmt.GoRaw("default:"),
			GoStmt.GoRaw("\treturn ValueType_TUnknown"),
			GoStmt.GoRaw("}")
		];

		return [
			GoDecl.GoStructDecl("ValueType", [{name: "tag", typeName: "int"}, {name: "params", typeName: "[]any"}]),
			GoDecl.GoGlobalVarDecl("ValueType_TNull", "*ValueType", GoExpr.GoRaw("&ValueType{tag: 0, params: []any{}}")),
			GoDecl.GoGlobalVarDecl("ValueType_TInt", "*ValueType", GoExpr.GoRaw("&ValueType{tag: 1, params: []any{}}")),
			GoDecl.GoGlobalVarDecl("ValueType_TFloat", "*ValueType", GoExpr.GoRaw("&ValueType{tag: 2, params: []any{}}")),
			GoDecl.GoGlobalVarDecl("ValueType_TBool", "*ValueType", GoExpr.GoRaw("&ValueType{tag: 3, params: []any{}}")),
			GoDecl.GoGlobalVarDecl("ValueType_TObject", "*ValueType", GoExpr.GoRaw("&ValueType{tag: 4, params: []any{}}")),
			GoDecl.GoGlobalVarDecl("ValueType_TFunction", "*ValueType", GoExpr.GoRaw("&ValueType{tag: 5, params: []any{}}")),
			GoDecl.GoGlobalVarDecl("ValueType_TUnknown", "*ValueType", GoExpr.GoRaw("&ValueType{tag: 8, params: []any{}}")),
			GoDecl.GoFuncDecl("ValueType_TClass", null, [
				{
					name: "c",
					typeName: "any"
				}
			],
				["*ValueType"], [GoStmt.GoReturn(GoExpr.GoRaw("&ValueType{tag: 6, params: []any{c}}"))]),
			GoDecl.GoFuncDecl("ValueType_TEnum", null, [{name: "e", typeName: "any"}], ["*ValueType"],
				[GoStmt.GoReturn(GoExpr.GoRaw("&ValueType{tag: 7, params: []any{e}}"))]),
			GoDecl.GoFuncDecl("hxrt_typeCallAny", null, [{name: "callable", typeName: "any"}, {name: "args", typeName: "[]any"}], ["any", "bool"], [
				GoStmt.GoRaw("result := any(nil)"),
				GoStmt.GoRaw("ok := false"),
				GoStmt.GoRaw("defer func() {"),
				GoStmt.GoRaw("\tif recover() != nil {"),
				GoStmt.GoRaw("\t\tresult = nil"),
				GoStmt.GoRaw("\t\tok = false"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("}()"),
				GoStmt.GoRaw("if callable == nil {"),
				GoStmt.GoRaw("\treturn nil, false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("fn := reflect.ValueOf(callable)"),
				GoStmt.GoRaw("if !fn.IsValid() || fn.Kind() != reflect.Func {"),
				GoStmt.GoRaw("\treturn nil, false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("fnType := fn.Type()"),
				GoStmt.GoRaw("if fnType.NumIn() != len(args) {"),
				GoStmt.GoRaw("\treturn nil, false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("in := make([]reflect.Value, len(args))"),
				GoStmt.GoRaw("for i := 0; i < len(args); i++ {"),
				GoStmt.GoRaw("\tparamType := fnType.In(i)"),
				GoStmt.GoRaw("\targ := args[i]"),
				GoStmt.GoRaw("\tif arg == nil {"),
				GoStmt.GoRaw("\t\tin[i] = reflect.Zero(paramType)"),
				GoStmt.GoRaw("\t\tcontinue"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tv := reflect.ValueOf(arg)"),
				GoStmt.GoRaw("\tif v.IsValid() && v.Type().AssignableTo(paramType) {"),
				GoStmt.GoRaw("\t\tin[i] = v"),
				GoStmt.GoRaw("\t\tcontinue"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tif v.IsValid() && v.Type().ConvertibleTo(paramType) {"),
				GoStmt.GoRaw("\t\tin[i] = v.Convert(paramType)"),
				GoStmt.GoRaw("\t\tcontinue"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tif paramType.Kind() == reflect.Interface && v.IsValid() {"),
				GoStmt.GoRaw("\t\tin[i] = v"),
				GoStmt.GoRaw("\t\tcontinue"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn nil, false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("out := fn.Call(in)"),
				GoStmt.GoRaw("if len(out) == 0 {"),
				GoStmt.GoRaw("\treturn nil, true"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("first := out[0]"),
				GoStmt.GoRaw("if !first.IsValid() {"),
				GoStmt.GoRaw("\treturn nil, true"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("result = first.Interface()"),
				GoStmt.GoRaw("ok = true"),
				GoStmt.GoRaw("return result, ok")
			]),
			GoDecl.GoFuncDecl("hxrt_typeResolvedClassName", null, [
				{
					name: "value",
					typeName: "any"
				}
			], ["string", "bool"], [
				GoStmt.GoRaw("switch current := value.(type) {"),
				GoStmt.GoRaw("case *hxrt__TypeClassValue:"),
				GoStmt.GoRaw("\tif current == nil || current.name == nil {"),
				GoStmt.GoRaw("\t\treturn \"\", false"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn *current.name, true"),
				GoStmt.GoRaw("case hxrt__TypeClassValue:"),
				GoStmt.GoRaw("\tif current.name == nil {"),
				GoStmt.GoRaw("\t\treturn \"\", false"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn *current.name, true"),
				GoStmt.GoRaw("case string:"),
				GoStmt.GoRaw("\treturn current, true"),
				GoStmt.GoRaw("case *string:"),
				GoStmt.GoRaw("\tif current == nil {"),
				GoStmt.GoRaw("\t\treturn \"\", false"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn *current, true"),
				GoStmt.GoRaw("default:"),
				GoStmt.GoRaw("\treturn \"\", false"),
				GoStmt.GoRaw("}")
			]),
			GoDecl.GoFuncDecl("hxrt_typeResolvedEnumName", null, [
				{
					name: "value",
					typeName: "any"
				}
			], ["string", "bool"], [
				GoStmt.GoRaw("switch current := value.(type) {"),
				GoStmt.GoRaw("case *hxrt__TypeEnumValue:"),
				GoStmt.GoRaw("\tif current == nil || current.name == nil {"),
				GoStmt.GoRaw("\t\treturn \"\", false"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn *current.name, true"),
				GoStmt.GoRaw("case hxrt__TypeEnumValue:"),
				GoStmt.GoRaw("\tif current.name == nil {"),
				GoStmt.GoRaw("\t\treturn \"\", false"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn *current.name, true"),
				GoStmt.GoRaw("case string:"),
				GoStmt.GoRaw("\treturn current, true"),
				GoStmt.GoRaw("case *string:"),
				GoStmt.GoRaw("\tif current == nil {"),
				GoStmt.GoRaw("\t\treturn \"\", false"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn *current, true"),
				GoStmt.GoRaw("default:"),
				GoStmt.GoRaw("\treturn \"\", false"),
				GoStmt.GoRaw("}")
			]),
			GoDecl.GoFuncDecl("hxrt_typeCreateClassInstance", null, [
				{
					name: "className",
					typeName: "string"
				},
				{name: "args", typeName: "[]any"}
			],
				["any", "bool"], classCreateBody),
			GoDecl.GoFuncDecl("hxrt_typeCreateClassEmptyInstance", null, [{name: "className", typeName: "string"}], ["any", "bool"], classCreateEmptyBody),
			GoDecl.GoFuncDecl("hxrt_typeCreateEnumInstance", null, [
				{name: "enumName", typeName: "string"},
				{name: "constructorName", typeName: "string"},
				{name: "constructorIndex", typeName: "int"},
				{name: "useIndex", typeName: "bool"},
				{name: "args", typeName: "[]any"}
			],
				["any", "bool"], enumCreateBody),
			GoDecl.GoFuncDecl("Type_getClass", null, [{name: "o", typeName: "any"}], ["any"], getClassBody),
			GoDecl.GoFuncDecl("Type_getEnum", null, [{name: "o", typeName: "any"}], ["any"], getEnumBody),
			GoDecl.GoFuncDecl("Type_getSuperClass", null, [{name: "c", typeName: "any"}], ["any"], getSuperClassBody),
			GoDecl.GoFuncDecl("Type_getClassName", null, [{name: "c", typeName: "any"}], ["*string"], [
				GoStmt.GoRaw("className, ok := hxrt_typeResolvedClassName(c)"),
				GoStmt.GoRaw("if !ok {"),
				GoStmt.GoRaw("\treturn nil"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("return hxrt.StringFromLiteral(className)")
			]),
			GoDecl.GoFuncDecl("Type_getClassFields", null, [
				{
					name: "c",
					typeName: "any"
				}
			],
				["[]*string"], getClassFieldsBody),
			GoDecl.GoFuncDecl("Type_getInstanceFields", null, [{name: "c", typeName: "any"}], ["[]*string"], getInstanceFieldsBody),
			GoDecl.GoFuncDecl("Type_getEnumName", null, [
				{
					name: "e",
					typeName: "any"
				}
			], ["*string"], [
				GoStmt.GoRaw("enumName, ok := hxrt_typeResolvedEnumName(e)"),
				GoStmt.GoRaw("if !ok {"),
				GoStmt.GoRaw("\treturn nil"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("return hxrt.StringFromLiteral(enumName)")
			]),
			GoDecl.GoFuncDecl("Type_resolveClass", null, [
				{
					name: "name",
					typeName: "*string"
				}
			],
				["any"], classResolveBody),
			GoDecl.GoFuncDecl("Type_resolveEnum", null, [{name: "name", typeName: "*string"}], ["any"], enumResolveBody),
			GoDecl.GoFuncDecl("Type_createInstance", null, [{name: "cl", typeName: "any"}, {name: "args", typeName: "[]any"}], ["any"], [
				GoStmt.GoRaw("className, ok := hxrt_typeResolvedClassName(cl)"),
				GoStmt.GoRaw("if !ok {"),
				GoStmt.GoRaw("\treturn nil"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("instance, ok := hxrt_typeCreateClassInstance(className, args)"),
				GoStmt.GoRaw("if !ok {"),
				GoStmt.GoRaw("\treturn nil"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("return instance")
			]),
			GoDecl.GoFuncDecl("Type_createEmptyInstance", null, [
				{
					name: "cl",
					typeName: "any"
				}
			], ["any"], [
				GoStmt.GoRaw("className, ok := hxrt_typeResolvedClassName(cl)"),
				GoStmt.GoRaw("if !ok {"),
				GoStmt.GoRaw("\treturn nil"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("instance, ok := hxrt_typeCreateClassEmptyInstance(className)"),
				GoStmt.GoRaw("if !ok {"),
				GoStmt.GoRaw("\treturn nil"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("return instance")
			]),
			GoDecl.GoFuncDecl("Type_createEnum", null, [
				{
					name: "e",
					typeName: "any"
				},
				{name: "constr", typeName: "*string"},
				{name: "params", typeName: "[]any"}
			], ["any"], [
				GoStmt.GoRaw("enumName, ok := hxrt_typeResolvedEnumName(e)"),
				GoStmt.GoRaw("if !ok {"),
				GoStmt.GoRaw("\treturn nil"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("constructorName := \"\""),
				GoStmt.GoRaw("if constr != nil {"),
				GoStmt.GoRaw("\tconstructorName = *hxrt.StdString(constr)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("enumValue, ok := hxrt_typeCreateEnumInstance(enumName, constructorName, 0, false, params)"),
				GoStmt.GoRaw("if !ok {"),
				GoStmt.GoRaw("\treturn nil"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("return enumValue")
			]),
			GoDecl.GoFuncDecl("Type_createEnumIndex", null, [
				{
					name: "e",
					typeName: "any"
				},
				{name: "index", typeName: "int"},
				{name: "params", typeName: "[]any"}
			], ["any"], [
				GoStmt.GoRaw("enumName, ok := hxrt_typeResolvedEnumName(e)"),
				GoStmt.GoRaw("if !ok {"),
				GoStmt.GoRaw("\treturn nil"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("enumValue, ok := hxrt_typeCreateEnumInstance(enumName, \"\", index, true, params)"),
				GoStmt.GoRaw("if !ok {"),
				GoStmt.GoRaw("\treturn nil"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("return enumValue")
			]),
			GoDecl.GoFuncDecl("Type_enumConstructor", null, [
				{
					name: "e",
					typeName: "any"
				}
			],
				["*string"], enumConstructorBody),
			GoDecl.GoFuncDecl("Type_enumIndex", null, [{name: "e", typeName: "any"}], ["int"], enumIndexBody),
			GoDecl.GoFuncDecl("Type_getEnumConstructs", null, [{name: "e", typeName: "any"}], ["[]*string"], getEnumConstructsBody),
			GoDecl.GoFuncDecl("Type_enumParameters", null, [{name: "e", typeName: "any"}], ["[]any"], enumParametersBody),
			GoDecl.GoFuncDecl("Type_allEnums", null, [{name: "e", typeName: "any"}], ["[]any"], allEnumsBody),
			GoDecl.GoFuncDecl("Type_typeof", null, [{name: "v", typeName: "any"}], ["any"], typeOfBody),
			GoDecl.GoFuncDecl("Type_enumEq", null, [{name: "a", typeName: "any"}, {name: "b", typeName: "any"}], ["bool"],
				[GoStmt.GoReturn(GoExpr.GoRaw("reflect.DeepEqual(a, b)"))])
		];
	}

	function lowerRegexSerializerShimDecls():Array<GoDecl> {
		var classMetadata = serializerClassMetadata();
		var classLookupBody = [GoStmt.GoRaw("switch typeName {")];
		for (entry in classMetadata) {
			classLookupBody.push(GoStmt.GoRaw("case " + goRawQuotedString(entry.goTypeName) + ":"));
			classLookupBody.push(GoStmt.GoRaw("\treturn " + goRawQuotedString(entry.haxeTypeName) + ", true"));
		}
		classLookupBody.push(GoStmt.GoRaw("default:"));
		classLookupBody.push(GoStmt.GoRaw("\treturn \"\", false"));
		classLookupBody.push(GoStmt.GoRaw("}"));

		var enumMetadata = serializerEnumMetadata();
		var enumLookupBody = [GoStmt.GoRaw("switch typeName {")];
		for (entry in enumMetadata) {
			var constructorLiterals = [for (constructor in entry.constructors) goRawQuotedString(constructor)].join(", ");
			enumLookupBody.push(GoStmt.GoRaw("case " + goRawQuotedString(entry.goTypeName) + ":"));
			enumLookupBody.push(GoStmt.GoRaw("\tconstructors := []string{" + constructorLiterals + "}"));
			enumLookupBody.push(GoStmt.GoRaw("\tif tag < 0 || tag >= len(constructors) {"));
			enumLookupBody.push(GoStmt.GoRaw("\t\treturn \"\", \"\", false"));
			enumLookupBody.push(GoStmt.GoRaw("\t}"));
			enumLookupBody.push(GoStmt.GoRaw("\treturn " + goRawQuotedString(entry.haxeTypeName) + ", constructors[tag], true"));
		}
		enumLookupBody.push(GoStmt.GoRaw("default:"));
		enumLookupBody.push(GoStmt.GoRaw("\treturn \"\", \"\", false"));
		enumLookupBody.push(GoStmt.GoRaw("}"));

		var enumLookupByNameBody = [GoStmt.GoRaw("switch enumName {")];
		for (entry in enumMetadata) {
			var constructorLiterals = [for (constructor in entry.constructors) goRawQuotedString(constructor)].join(", ");
			enumLookupByNameBody.push(GoStmt.GoRaw("case " + goRawQuotedString(entry.haxeTypeName) + ":"));
			enumLookupByNameBody.push(GoStmt.GoRaw("\tconstructors := []string{" + constructorLiterals + "}"));
			enumLookupByNameBody.push(GoStmt.GoRaw("\tif index < 0 || index >= len(constructors) {"));
			enumLookupByNameBody.push(GoStmt.GoRaw("\t\treturn \"\", false"));
			enumLookupByNameBody.push(GoStmt.GoRaw("\t}"));
			enumLookupByNameBody.push(GoStmt.GoRaw("\treturn constructors[index], true"));
		}
		enumLookupByNameBody.push(GoStmt.GoRaw("default:"));
		enumLookupByNameBody.push(GoStmt.GoRaw("\treturn \"\", false"));
		enumLookupByNameBody.push(GoStmt.GoRaw("}"));

		var enumLookupIndexBody = [GoStmt.GoRaw("switch enumName {")];
		for (entry in enumMetadata) {
			enumLookupIndexBody.push(GoStmt.GoRaw("case " + goRawQuotedString(entry.haxeTypeName) + ":"));
			for (index in 0...entry.constructors.length) {
				enumLookupIndexBody.push(GoStmt.GoRaw("\tif constructorName == " + goRawQuotedString(entry.constructors[index]) + " {"));
				enumLookupIndexBody.push(GoStmt.GoRaw("\t\treturn " + index + ", true"));
				enumLookupIndexBody.push(GoStmt.GoRaw("\t}"));
			}
			enumLookupIndexBody.push(GoStmt.GoRaw("\treturn 0, false"));
		}
		enumLookupIndexBody.push(GoStmt.GoRaw("default:"));
		enumLookupIndexBody.push(GoStmt.GoRaw("\treturn 0, false"));
		enumLookupIndexBody.push(GoStmt.GoRaw("}"));

		var classExistsBody = [GoStmt.GoRaw("switch className {")];
		for (entry in classMetadata) {
			classExistsBody.push(GoStmt.GoRaw("case " + goRawQuotedString(entry.haxeTypeName) + ":"));
			classExistsBody.push(GoStmt.GoRaw("\treturn true"));
		}
		classExistsBody.push(GoStmt.GoRaw("default:"));
		classExistsBody.push(GoStmt.GoRaw("\treturn false"));
		classExistsBody.push(GoStmt.GoRaw("}"));

		var classCreateBody = [GoStmt.GoRaw("switch className {")];
		for (entry in classMetadata) {
			classCreateBody.push(GoStmt.GoRaw("case " + goRawQuotedString(entry.haxeTypeName) + ":"));
			classCreateBody.push(GoStmt.GoRaw("\tinstance := &" + entry.goTypeName + "{}"));
			classCreateBody.push(GoStmt.GoRaw("\thxrt_unserializerBindSelf(instance)"));
			classCreateBody.push(GoStmt.GoRaw("\treturn instance, true"));
		}
		classCreateBody.push(GoStmt.GoRaw("default:"));
		classCreateBody.push(GoStmt.GoRaw("\treturn nil, false"));
		classCreateBody.push(GoStmt.GoRaw("}"));

		var enumExistsBody = [GoStmt.GoRaw("switch enumName {")];
		for (entry in enumMetadata) {
			enumExistsBody.push(GoStmt.GoRaw("case " + goRawQuotedString(entry.haxeTypeName) + ":"));
			enumExistsBody.push(GoStmt.GoRaw("\treturn true"));
		}
		enumExistsBody.push(GoStmt.GoRaw("default:"));
		enumExistsBody.push(GoStmt.GoRaw("\treturn false"));
		enumExistsBody.push(GoStmt.GoRaw("}"));

		var enumCreateBody = [GoStmt.GoRaw("switch enumName {")];
		for (entry in enumMetadata) {
			enumCreateBody.push(GoStmt.GoRaw("case " + goRawQuotedString(entry.haxeTypeName) + ":"));
			enumCreateBody.push(GoStmt.GoRaw("\tenumValue := &" + entry.goTypeName + "{tag: constructorIndex}"));
			enumCreateBody.push(GoStmt.GoRaw("\tif len(args) > 0 {"));
			enumCreateBody.push(GoStmt.GoRaw("\t\tcopied := make([]any, len(args))"));
			enumCreateBody.push(GoStmt.GoRaw("\t\tcopy(copied, args)"));
			enumCreateBody.push(GoStmt.GoRaw("\t\tenumValue.params = copied"));
			enumCreateBody.push(GoStmt.GoRaw("\t}"));
			enumCreateBody.push(GoStmt.GoRaw("\treturn enumValue, true"));
		}
		enumCreateBody.push(GoStmt.GoRaw("default:"));
		enumCreateBody.push(GoStmt.GoRaw("\treturn nil, false"));
		enumCreateBody.push(GoStmt.GoRaw("}"));

		return [
			GoDecl.GoStructDecl("EReg", [
				{name: "regex", typeName: "*regexp.Regexp"},
				{name: "global", typeName: "bool"},
				{name: "lastSource", typeName: "*string"},
				{name: "lastIndices", typeName: "[]int"}
			]),
			GoDecl.GoFuncDecl("New_EReg", null, [
				{
					name: "pattern",
					typeName: "*string"
				},
				{name: "options", typeName: "*string"}
			], ["*EReg"], [
				GoStmt.GoRaw("rawPattern := *hxrt.StdString(pattern)"),
				GoStmt.GoRaw("rawOptions := *hxrt.StdString(options)"),
				GoStmt.GoRaw("global := false"),
				GoStmt.GoRaw("flagI := false"),
				GoStmt.GoRaw("flagM := false"),
				GoStmt.GoRaw("flagS := false"),
				GoStmt.GoRaw("for _, option := range rawOptions {"),
				GoStmt.GoRaw("\tswitch option {"),
				GoStmt.GoRaw("\tcase 'g':"),
				GoStmt.GoRaw("\t\tglobal = true"),
				GoStmt.GoRaw("\tcase 'i':"),
				GoStmt.GoRaw("\t\tflagI = true"),
				GoStmt.GoRaw("\tcase 'm':"),
				GoStmt.GoRaw("\t\tflagM = true"),
				GoStmt.GoRaw("\tcase 's':"),
				GoStmt.GoRaw("\t\tflagS = true"),
				GoStmt.GoRaw("\tcase 'u':"),
				GoStmt.GoRaw("\t\t// RE2 is UTF-8 aware by default; keep parity by accepting and ignoring this option."),
				GoStmt.GoRaw("\tdefault:"),
				GoStmt.GoRaw("\t\thxrt.Throw(hxrt.StringFromLiteral(\"Unsupported regexp option '\" + string(option) + \"'\"))"),
				GoStmt.GoRaw("\t\treturn &EReg{regex: regexp.MustCompile(\"a^\"), global: false, lastSource: nil, lastIndices: nil}"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("inlineFlags := \"\""),
				GoStmt.GoRaw("if flagI {"),
				GoStmt.GoRaw("\tinlineFlags += \"i\""),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if flagM {"),
				GoStmt.GoRaw("\tinlineFlags += \"m\""),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if flagS {"),
				GoStmt.GoRaw("\tinlineFlags += \"s\""),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if inlineFlags != \"\" {"),
				GoStmt.GoRaw("\trawPattern = \"(?\" + inlineFlags + \")\" + rawPattern"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("compiled, err := regexp.Compile(rawPattern)"),
				GoStmt.GoIf(GoExpr.GoBinary("!=", GoExpr.GoIdent("err"), GoExpr.GoNil), [
					GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("hxrt.Throw"), [GoExpr.GoIdent("err")])),
					GoStmt.GoAssign(GoExpr.GoIdent("compiled"),
						GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("regexp"), "MustCompile"), [GoExpr.GoStringLiteral("a^")]))
				],
					null),
				GoStmt.GoReturn(GoExpr.GoRaw("&EReg{regex: compiled, global: global, lastSource: nil, lastIndices: nil}"))
			]),
			GoDecl.GoFuncDecl("hxrt_eregHasMatch", null, [
				{
					name: "self",
					typeName: "*EReg"
				}
			], ["bool"], [
				GoStmt.GoRaw("return self != nil && self.lastSource != nil && len(self.lastIndices) >= 2 && self.lastIndices[0] >= 0 && self.lastIndices[1] >= self.lastIndices[0]")
			]),
			GoDecl.GoFuncDecl("hxrt_eregThrowNoMatch", null, [], [], [
				GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("hxrt.Throw"), [
					GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoStringLiteral("Invalid regex operation because no match was made")])
				]))
			]),
			GoDecl.GoFuncDecl("hxrt_eregThrowInvalidGroup", null, [], [], [
				GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("hxrt.Throw"), [
					GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoStringLiteral("Invalid group")])
				]))
			]),
			GoDecl.GoFuncDecl("match", {
				name: "self",
				typeName: "*EReg"
			}, [{name: "source", typeName: "*string"}], ["bool"], [
				GoStmt.GoRaw("if self == nil || self.regex == nil {"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("raw := *hxrt.StdString(source)"),
				GoStmt.GoRaw("found := self.regex.FindStringSubmatchIndex(raw)"),
				GoStmt.GoRaw("if found == nil {"),
				GoStmt.GoRaw("\tself.lastSource = nil"),
				GoStmt.GoRaw("\tself.lastIndices = nil"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("indices := make([]int, len(found))"),
				GoStmt.GoRaw("copy(indices, found)"),
				GoStmt.GoRaw("self.lastSource = hxrt.StringFromLiteral(raw)"),
				GoStmt.GoRaw("self.lastIndices = indices"),
				GoStmt.GoReturn(GoExpr.GoBoolLiteral(true))
			]),
			GoDecl.GoFuncDecl("matchSub", {
				name: "self",
				typeName: "*EReg"
			}, [
				{
					name: "source",
					typeName: "*string"
				},
				{name: "pos", typeName: "int"}
			], ["bool"], [
				GoStmt.GoRaw("if self == nil || self.regex == nil {"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("raw := *hxrt.StdString(source)"),
				GoStmt.GoRaw("if pos < 0 {"),
				GoStmt.GoRaw("\tpos = 0"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if pos > len(raw) {"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("found := self.regex.FindStringSubmatchIndex(raw[pos:])"),
				GoStmt.GoRaw("if found == nil {"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("shifted := make([]int, len(found))"),
				GoStmt.GoRaw("for i := 0; i < len(found); i++ {"),
				GoStmt.GoRaw("\tif found[i] >= 0 {"),
				GoStmt.GoRaw("\t\tshifted[i] = found[i] + pos"),
				GoStmt.GoRaw("\t} else {"),
				GoStmt.GoRaw("\t\tshifted[i] = -1"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("self.lastSource = hxrt.StringFromLiteral(raw)"),
				GoStmt.GoRaw("self.lastIndices = shifted"),
				GoStmt.GoReturn(GoExpr.GoBoolLiteral(true))
			]),
			GoDecl.GoFuncDecl("matched", {
				name: "self",
				typeName: "*EReg"
			}, [{name: "index", typeName: "int"}], ["*string"], [
				GoStmt.GoRaw("if !hxrt_eregHasMatch(self) {"),
				GoStmt.GoRaw("\thxrt_eregThrowNoMatch()"),
				GoStmt.GoRaw("\treturn nil"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if index < 0 {"),
				GoStmt.GoRaw("\thxrt_eregThrowInvalidGroup()"),
				GoStmt.GoRaw("\treturn nil"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("offset := index * 2"),
				GoStmt.GoRaw("if offset+1 >= len(self.lastIndices) {"),
				GoStmt.GoRaw("\thxrt_eregThrowInvalidGroup()"),
				GoStmt.GoRaw("\treturn nil"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("start := self.lastIndices[offset]"),
				GoStmt.GoRaw("end := self.lastIndices[offset+1]"),
				GoStmt.GoRaw("if start < 0 || end < start {"),
				GoStmt.GoRaw("\treturn nil"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("raw := *hxrt.StdString(self.lastSource)"),
				GoStmt.GoRaw("if end > len(raw) {"),
				GoStmt.GoRaw("\tend = len(raw)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoRaw("raw[start:end]")]))
			]),
			GoDecl.GoFuncDecl("matchedPos", {
				name: "self",
				typeName: "*EReg"
			}, [], ["map[string]any"], [
				GoStmt.GoRaw("if !hxrt_eregHasMatch(self) {"),
				GoStmt.GoRaw("\thxrt_eregThrowNoMatch()"),
				GoStmt.GoRaw("\treturn map[string]any{\"pos\": 0, \"len\": 0}"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("start := self.lastIndices[0]"),
				GoStmt.GoRaw("end := self.lastIndices[1]"),
				GoStmt.GoReturn(GoExpr.GoRaw("map[string]any{\"pos\": start, \"len\": end - start}"))
			]),
			GoDecl.GoFuncDecl("matchedLeft", {
				name: "self",
				typeName: "*EReg"
			}, [], ["*string"], [
				GoStmt.GoRaw("if !hxrt_eregHasMatch(self) {"),
				GoStmt.GoRaw("\thxrt_eregThrowNoMatch()"),
				GoStmt.GoRaw("\treturn nil"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("raw := *hxrt.StdString(self.lastSource)"),
				GoStmt.GoRaw("start := self.lastIndices[0]"),
				GoStmt.GoRaw("if start > len(raw) {"),
				GoStmt.GoRaw("\tstart = len(raw)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoRaw("raw[:start]")]))
			]),
			GoDecl.GoFuncDecl("matchedRight", {
				name: "self",
				typeName: "*EReg"
			}, [], ["*string"], [
				GoStmt.GoRaw("if !hxrt_eregHasMatch(self) {"),
				GoStmt.GoRaw("\thxrt_eregThrowNoMatch()"),
				GoStmt.GoRaw("\treturn nil"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("raw := *hxrt.StdString(self.lastSource)"),
				GoStmt.GoRaw("end := self.lastIndices[1]"),
				GoStmt.GoRaw("if end > len(raw) {"),
				GoStmt.GoRaw("\tend = len(raw)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoRaw("raw[end:]")]))
			]),
			GoDecl.GoFuncDecl("split", {
				name: "self",
				typeName: "*EReg"
			}, [{name: "source", typeName: "*string"}], ["[]*string"], [
				GoStmt.GoRaw("raw := *hxrt.StdString(source)"),
				GoStmt.GoRaw("if self == nil || self.regex == nil {"),
				GoStmt.GoRaw("\treturn []*string{hxrt.StringFromLiteral(raw)}"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("parts := self.regex.Split(raw, -1)"),
				GoStmt.GoRaw("out := make([]*string, 0, len(parts))"),
				GoStmt.GoRaw("for _, part := range parts {"),
				GoStmt.GoRaw("\tout = append(out, hxrt.StringFromLiteral(part))"),
				GoStmt.GoRaw("}"),
				GoStmt.GoReturn(GoExpr.GoIdent("out"))
			]),
			GoDecl.GoFuncDecl("replace", {
				name: "self",
				typeName: "*EReg"
			}, [
				{
					name: "source",
					typeName: "*string"
				},
				{name: "by", typeName: "*string"}
			], ["*string"], [
				GoStmt.GoRaw("if self == nil || self.regex == nil {"),
				GoStmt.GoRaw("\treturn source"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("rawSource := *hxrt.StdString(source)"),
				GoStmt.GoRaw("rawBy := *hxrt.StdString(by)"),
				GoStmt.GoRaw("if self.global {"),
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"),
					[
						GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoSelector(GoExpr.GoIdent("self"), "regex"), "ReplaceAllString"),
							[GoExpr.GoIdent("rawSource"), GoExpr.GoIdent("rawBy")])
					])),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("first := self.regex.FindStringSubmatchIndex(rawSource)"),
				GoStmt.GoRaw("if first == nil {"),
				GoStmt.GoRaw("\treturn hxrt.StringFromLiteral(rawSource)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("replacement := self.regex.ExpandString(nil, rawBy, rawSource, first)"),
				GoStmt.GoRaw("out := rawSource[:first[0]] + string(replacement) + rawSource[first[1]:]"),
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoIdent("out")]))
			]),
			GoDecl.GoFuncDecl("map_", {
				name: "self",
				typeName: "*EReg"
			}, [
				{
					name: "source",
					typeName: "*string"
				},
				{name: "callback", typeName: "func(*EReg) *string"}
			], ["*string"], [
				GoStmt.GoRaw("if self == nil || self.regex == nil {"),
				GoStmt.GoRaw("\treturn source"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("raw := *hxrt.StdString(source)"),
				GoStmt.GoRaw("if callback == nil {"),
				GoStmt.GoRaw("\treturn hxrt.StringFromLiteral(raw)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("var matches [][]int"),
				GoStmt.GoRaw("if self.global {"),
				GoStmt.GoRaw("\tmatches = self.regex.FindAllStringSubmatchIndex(raw, -1)"),
				GoStmt.GoRaw("} else {"),
				GoStmt.GoRaw("\tif first := self.regex.FindStringSubmatchIndex(raw); first != nil {"),
				GoStmt.GoRaw("\t\tmatches = [][]int{first}"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if len(matches) == 0 {"),
				GoStmt.GoRaw("\treturn hxrt.StringFromLiteral(raw)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("var builder strings.Builder"),
				GoStmt.GoRaw("cursor := 0"),
				GoStmt.GoRaw("for _, match := range matches {"),
				GoStmt.GoRaw("\tif len(match) < 2 {"),
				GoStmt.GoRaw("\t\tcontinue"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tstart := match[0]"),
				GoStmt.GoRaw("\tend := match[1]"),
				GoStmt.GoRaw("\tif start < cursor {"),
				GoStmt.GoRaw("\t\tstart = cursor"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tif end < start {"),
				GoStmt.GoRaw("\t\tend = start"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tif start > len(raw) {"),
				GoStmt.GoRaw("\t\tstart = len(raw)"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tif end > len(raw) {"),
				GoStmt.GoRaw("\t\tend = len(raw)"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tbuilder.WriteString(raw[cursor:start])"),
				GoStmt.GoRaw("\tindices := make([]int, len(match))"),
				GoStmt.GoRaw("\tcopy(indices, match)"),
				GoStmt.GoRaw("\tself.lastSource = hxrt.StringFromLiteral(raw)"),
				GoStmt.GoRaw("\tself.lastIndices = indices"),
				GoStmt.GoRaw("\treplacement := callback(self)"),
				GoStmt.GoRaw("\tbuilder.WriteString(*hxrt.StdString(replacement))"),
				GoStmt.GoRaw("\tcursor = end"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("builder.WriteString(raw[cursor:])"),
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"),
					[GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("builder"), "String"), [])]))
			]),
			GoDecl.GoStructDecl("haxe__SerializedDate",
				[
					{
						name: "ms",
						typeName: "float64"
					}
				]),
			GoDecl.GoStructDecl("haxe__SerializedBytes", [{name: "data", typeName: "[]byte"}]),
			GoDecl.GoStructDecl("haxe__SerializedClassRef", [{name: "name", typeName: "string"}]),
			GoDecl.GoStructDecl("haxe__SerializedEnumRef", [{name: "name", typeName: "string"}]),
			GoDecl.GoStructDecl("haxe__SerializedClass", [
				{
					name: "name",
					typeName: "string"
				},
				{name: "fieldNames", typeName: "[]string"},
				{name: "fieldValues", typeName: "[]any"}
			]),
			GoDecl.GoStructDecl("haxe__SerializedEnum",
				[
					{
						name: "name",
						typeName: "string"
					},
					{name: "constructor", typeName: "string"},
					{name: "constructorIndex", typeName: "int"},
					{name: "hasConstructorIndex", typeName: "bool"},
					{name: "args", typeName: "[]any"}
				]),
			GoDecl.GoStructDecl("haxe__Unserializer__DefaultResolver", []),
			GoDecl.GoStructDecl("haxe__Unserializer__NullResolver", []),
			GoDecl.GoGlobalVarDecl("haxe__Serializer_USE_CACHE", "bool", GoExpr.GoBoolLiteral(false)),
			GoDecl.GoGlobalVarDecl("haxe__Serializer_USE_ENUM_INDEX", "bool", GoExpr.GoBoolLiteral(false)),
			GoDecl.GoGlobalVarDecl("haxe__Unserializer_DEFAULT_RESOLVER", "any", GoExpr.GoRaw("&haxe__Unserializer__DefaultResolver{}")),
			GoDecl.GoGlobalVarDecl("haxe__Unserializer_NULL_RESOLVER", "any", GoExpr.GoRaw("&haxe__Unserializer__NullResolver{}")),
			GoDecl.GoFuncDecl("hxrt_serializerLookupClassName", null, [
				{
					name: "typeName",
					typeName: "string"
				}
			], ["string", "bool"], classLookupBody),
			GoDecl.GoFuncDecl("hxrt_serializerLookupEnumConstructor", null, [
				{
					name: "typeName",
					typeName: "string"
				},
				{name: "tag", typeName: "int"}
			], ["string", "string", "bool"],
				enumLookupBody),
			GoDecl.GoFuncDecl("hxrt_serializerLookupEnumConstructorByName", null, [
				{
					name: "enumName",
					typeName: "string"
				},
				{name: "index", typeName: "int"}
			], ["string", "bool"], enumLookupByNameBody),
			GoDecl.GoFuncDecl("hxrt_serializerLookupEnumIndexByName", null, [
				{
					name: "enumName",
					typeName: "string"
				},
				{name: "constructorName", typeName: "string"}
			], ["int", "bool"], enumLookupIndexBody),
			GoDecl.GoFuncDecl("resolveClass", {
				name: "self",
				typeName: "*haxe__Unserializer__DefaultResolver"
			}, [{name: "name", typeName: "*string"}], ["any"], [
				GoStmt.GoRaw("className := *hxrt.StdString(name)"),
				GoStmt.GoRaw("if !hxrt_unserializerHasClass(className) {"),
				GoStmt.GoRaw("\treturn nil"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("return &haxe__SerializedClassRef{name: className}")
			]),
			GoDecl.GoFuncDecl("resolveEnum", {
				name: "self",
				typeName: "*haxe__Unserializer__DefaultResolver"
			}, [{name: "name", typeName: "*string"}], ["any"], [
				GoStmt.GoRaw("enumName := *hxrt.StdString(name)"),
				GoStmt.GoRaw("if !hxrt_unserializerHasEnum(enumName) {"),
				GoStmt.GoRaw("\treturn nil"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("return &haxe__SerializedEnumRef{name: enumName}")
			]),
			GoDecl.GoFuncDecl("resolveClass", {
				name: "self",
				typeName: "*haxe__Unserializer__NullResolver"
			}, [{name: "name", typeName: "*string"}],
				["any"], [GoStmt.GoReturn(GoExpr.GoNil)]),
			GoDecl.GoFuncDecl("resolveEnum", {
				name: "self",
				typeName: "*haxe__Unserializer__NullResolver"
			},
				[{name: "name", typeName: "*string"}], ["any"], [GoStmt.GoReturn(GoExpr.GoNil)]),
			GoDecl.GoFuncDecl("hxrt_unserializerHasClass", null, [{name: "className", typeName: "string"}], ["bool"], classExistsBody),
			GoDecl.GoFuncDecl("hxrt_unserializerHasEnum", null, [{name: "enumName", typeName: "string"}], ["bool"], enumExistsBody),
			GoDecl.GoFuncDecl("hxrt_unserializerBindSelf", null, [{name: "instance", typeName: "any"}], [], [
				GoStmt.GoRaw("if instance == nil {"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("rv := reflect.ValueOf(instance)"),
				GoStmt.GoRaw("if !rv.IsValid() || rv.Kind() != reflect.Pointer || rv.IsNil() {"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("elem := rv.Elem()"),
				GoStmt.GoRaw("if !elem.IsValid() || elem.Kind() != reflect.Struct {"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("field := elem.FieldByName(\"__hx_this\")"),
				GoStmt.GoRaw("if !field.IsValid() {"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if !rv.Type().AssignableTo(field.Type()) {"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if field.CanSet() {"),
				GoStmt.GoRaw("\tfield.Set(rv)"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if !field.CanAddr() {"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("lifted := reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem()"),
				GoStmt.GoRaw("lifted.Set(rv)")
			]),
			GoDecl.GoFuncDecl("hxrt_unserializerInvokeResolver", null, [
				{
					name: "resolver",
					typeName: "any"
				},
				{name: "methodName", typeName: "string"},
				{name: "name", typeName: "string"}
			], ["any", "bool"], [
				GoStmt.GoRaw("if resolver == nil {"),
				GoStmt.GoRaw("\treturn nil, false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("result := any(nil)"),
				GoStmt.GoRaw("ok := false"),
				GoStmt.GoRaw("defer func() {"),
				GoStmt.GoRaw("\tif recover() != nil {"),
				GoStmt.GoRaw("\t\tresult = nil"),
				GoStmt.GoRaw("\t\tok = false"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("}()"),
				GoStmt.GoRaw("rv := reflect.ValueOf(resolver)"),
				GoStmt.GoRaw("if !rv.IsValid() {"),
				GoStmt.GoRaw("\treturn nil, false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("for rv.IsValid() && rv.Kind() == reflect.Interface {"),
				GoStmt.GoRaw("\tif rv.IsNil() {"),
				GoStmt.GoRaw("\t\treturn nil, false"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\trv = rv.Elem()"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("method := reflect.Value{}"),
				GoStmt.GoRaw("if rv.IsValid() {"),
				GoStmt.GoRaw("\tmethod = rv.MethodByName(methodName)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if !method.IsValid() && rv.IsValid() && rv.Kind() != reflect.Pointer && rv.CanAddr() {"),
				GoStmt.GoRaw("\tmethod = rv.Addr().MethodByName(methodName)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if !method.IsValid() && rv.IsValid() && rv.Kind() == reflect.Pointer && !rv.IsNil() {"),
				GoStmt.GoRaw("\tmethod = rv.Elem().MethodByName(methodName)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if !method.IsValid() && rv.IsValid() {"),
				GoStmt.GoRaw("\tswitch rv.Kind() {"),
				GoStmt.GoRaw("\tcase reflect.Struct:"),
				GoStmt.GoRaw("\t\tfield := rv.FieldByName(methodName)"),
				GoStmt.GoRaw("\t\tif field.IsValid() {"),
				GoStmt.GoRaw("\t\t\tfor field.IsValid() && field.Kind() == reflect.Interface {"),
				GoStmt.GoRaw("\t\t\t\tif field.IsNil() {"),
				GoStmt.GoRaw("\t\t\t\t\tbreak"),
				GoStmt.GoRaw("\t\t\t\t}"),
				GoStmt.GoRaw("\t\t\t\tfield = field.Elem()"),
				GoStmt.GoRaw("\t\t\t}"),
				GoStmt.GoRaw("\t\t\tif field.IsValid() && field.Kind() == reflect.Func {"),
				GoStmt.GoRaw("\t\t\t\tmethod = field"),
				GoStmt.GoRaw("\t\t\t}"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\tcase reflect.Pointer:"),
				GoStmt.GoRaw("\t\tif !rv.IsNil() && rv.Elem().Kind() == reflect.Struct {"),
				GoStmt.GoRaw("\t\t\tfield := rv.Elem().FieldByName(methodName)"),
				GoStmt.GoRaw("\t\t\tif field.IsValid() {"),
				GoStmt.GoRaw("\t\t\t\tfor field.IsValid() && field.Kind() == reflect.Interface {"),
				GoStmt.GoRaw("\t\t\t\t\tif field.IsNil() {"),
				GoStmt.GoRaw("\t\t\t\t\t\tbreak"),
				GoStmt.GoRaw("\t\t\t\t\t}"),
				GoStmt.GoRaw("\t\t\t\t\tfield = field.Elem()"),
				GoStmt.GoRaw("\t\t\t\t}"),
				GoStmt.GoRaw("\t\t\t\tif field.IsValid() && field.Kind() == reflect.Func {"),
				GoStmt.GoRaw("\t\t\t\t\tmethod = field"),
				GoStmt.GoRaw("\t\t\t\t}"),
				GoStmt.GoRaw("\t\t\t}"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\tcase reflect.Map:"),
				GoStmt.GoRaw("\t\tif rv.Type().Key().Kind() == reflect.String {"),
				GoStmt.GoRaw("\t\t\tfield := rv.MapIndex(reflect.ValueOf(methodName))"),
				GoStmt.GoRaw("\t\t\tif field.IsValid() {"),
				GoStmt.GoRaw("\t\t\t\tfor field.IsValid() && field.Kind() == reflect.Interface {"),
				GoStmt.GoRaw("\t\t\t\t\tif field.IsNil() {"),
				GoStmt.GoRaw("\t\t\t\t\t\tbreak"),
				GoStmt.GoRaw("\t\t\t\t\t}"),
				GoStmt.GoRaw("\t\t\t\t\tfield = field.Elem()"),
				GoStmt.GoRaw("\t\t\t\t}"),
				GoStmt.GoRaw("\t\t\t\tif field.IsValid() && field.Kind() == reflect.Func {"),
				GoStmt.GoRaw("\t\t\t\t\tmethod = field"),
				GoStmt.GoRaw("\t\t\t\t}"),
				GoStmt.GoRaw("\t\t\t}"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if !method.IsValid() || method.Kind() != reflect.Func {"),
				GoStmt.GoRaw("\treturn nil, false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("methodType := method.Type()"),
				GoStmt.GoRaw("if methodType.NumIn() != 1 || methodType.NumOut() < 1 {"),
				GoStmt.GoRaw("\treturn nil, false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("argType := methodType.In(0)"),
				GoStmt.GoRaw("nameValue := reflect.ValueOf(name)"),
				GoStmt.GoRaw("var arg reflect.Value"),
				GoStmt.GoRaw("if nameValue.Type().AssignableTo(argType) {"),
				GoStmt.GoRaw("\targ = nameValue"),
				GoStmt.GoRaw("} else if nameValue.Type().ConvertibleTo(argType) {"),
				GoStmt.GoRaw("\targ = nameValue.Convert(argType)"),
				GoStmt.GoRaw("} else if argType.Kind() == reflect.Pointer && argType.Elem().Kind() == reflect.String {"),
				GoStmt.GoRaw("\tnameCopy := name"),
				GoStmt.GoRaw("\targ = reflect.ValueOf(&nameCopy)"),
				GoStmt.GoRaw("\tif !arg.Type().AssignableTo(argType) {"),
				GoStmt.GoRaw("\t\tif arg.Type().ConvertibleTo(argType) {"),
				GoStmt.GoRaw("\t\t\targ = arg.Convert(argType)"),
				GoStmt.GoRaw("\t\t} else {"),
				GoStmt.GoRaw("\t\t\treturn nil, false"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("} else {"),
				GoStmt.GoRaw("\treturn nil, false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("out := method.Call([]reflect.Value{arg})"),
				GoStmt.GoRaw("if len(out) == 0 {"),
				GoStmt.GoRaw("\treturn nil, true"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("value := out[0]"),
				GoStmt.GoRaw("if !value.IsValid() {"),
				GoStmt.GoRaw("\treturn nil, true"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("for value.IsValid() && value.Kind() == reflect.Interface {"),
				GoStmt.GoRaw("\tif value.IsNil() {"),
				GoStmt.GoRaw("\t\treturn nil, true"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tvalue = value.Elem()"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if !value.IsValid() {"),
				GoStmt.GoRaw("\treturn nil, true"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("switch value.Kind() {"),
				GoStmt.GoRaw("case reflect.Pointer, reflect.Map, reflect.Slice, reflect.Func, reflect.Chan:"),
				GoStmt.GoRaw("\tif value.IsNil() {"),
				GoStmt.GoRaw("\t\treturn nil, true"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("result = value.Interface()"),
				GoStmt.GoRaw("ok = true"),
				GoStmt.GoRaw("return result, ok")
			]),
			GoDecl.GoFuncDecl("hxrt_unserializerResolveClass", null, [
				{
					name: "self",
					typeName: "*haxe__Unserializer"
				},
				{name: "name", typeName: "string"}
			], ["any"], [
				GoStmt.GoRaw("var resolver any"),
				GoStmt.GoRaw("if self != nil {"),
				GoStmt.GoRaw("\tresolver = self.resolver"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if resolver == nil {"),
				GoStmt.GoRaw("\tresolver = haxe__Unserializer_DEFAULT_RESOLVER"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if resolver == nil {"),
				GoStmt.GoRaw("\treturn nil"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("switch current := resolver.(type) {"),
				GoStmt.GoRaw("case interface{ resolveClass(*string) any }:"),
				GoStmt.GoRaw("\treturn current.resolveClass(hxrt.StringFromLiteral(name))"),
				GoStmt.GoRaw("case interface{ resolveClass(string) any }:"),
				GoStmt.GoRaw("\treturn current.resolveClass(name)"),
				GoStmt.GoRaw("case interface{ resolveClass(any) any }:"),
				GoStmt.GoRaw("\treturn current.resolveClass(name)"),
				GoStmt.GoRaw("case interface{ ResolveClass(*string) any }:"),
				GoStmt.GoRaw("\treturn current.ResolveClass(hxrt.StringFromLiteral(name))"),
				GoStmt.GoRaw("case interface{ ResolveClass(string) any }:"),
				GoStmt.GoRaw("\treturn current.ResolveClass(name)"),
				GoStmt.GoRaw("case interface{ ResolveClass(any) any }:"),
				GoStmt.GoRaw("\treturn current.ResolveClass(name)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("resolved, ok := hxrt_unserializerInvokeResolver(resolver, \"resolveClass\", name)"),
				GoStmt.GoRaw("if ok {"),
				GoStmt.GoRaw("\treturn resolved"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("resolved, ok = hxrt_unserializerInvokeResolver(resolver, \"ResolveClass\", name)"),
				GoStmt.GoRaw("if ok {"),
				GoStmt.GoRaw("\treturn resolved"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("return nil")
			]),
			GoDecl.GoFuncDecl("hxrt_unserializerResolveEnum", null, [
				{
					name: "self",
					typeName: "*haxe__Unserializer"
				},
				{name: "name", typeName: "string"}
			], ["any"], [
				GoStmt.GoRaw("var resolver any"),
				GoStmt.GoRaw("if self != nil {"),
				GoStmt.GoRaw("\tresolver = self.resolver"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if resolver == nil {"),
				GoStmt.GoRaw("\tresolver = haxe__Unserializer_DEFAULT_RESOLVER"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if resolver == nil {"),
				GoStmt.GoRaw("\treturn nil"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("switch current := resolver.(type) {"),
				GoStmt.GoRaw("case interface{ resolveEnum(*string) any }:"),
				GoStmt.GoRaw("\treturn current.resolveEnum(hxrt.StringFromLiteral(name))"),
				GoStmt.GoRaw("case interface{ resolveEnum(string) any }:"),
				GoStmt.GoRaw("\treturn current.resolveEnum(name)"),
				GoStmt.GoRaw("case interface{ resolveEnum(any) any }:"),
				GoStmt.GoRaw("\treturn current.resolveEnum(name)"),
				GoStmt.GoRaw("case interface{ ResolveEnum(*string) any }:"),
				GoStmt.GoRaw("\treturn current.ResolveEnum(hxrt.StringFromLiteral(name))"),
				GoStmt.GoRaw("case interface{ ResolveEnum(string) any }:"),
				GoStmt.GoRaw("\treturn current.ResolveEnum(name)"),
				GoStmt.GoRaw("case interface{ ResolveEnum(any) any }:"),
				GoStmt.GoRaw("\treturn current.ResolveEnum(name)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("resolved, ok := hxrt_unserializerInvokeResolver(resolver, \"resolveEnum\", name)"),
				GoStmt.GoRaw("if ok {"),
				GoStmt.GoRaw("\treturn resolved"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("resolved, ok = hxrt_unserializerInvokeResolver(resolver, \"ResolveEnum\", name)"),
				GoStmt.GoRaw("if ok {"),
				GoStmt.GoRaw("\treturn resolved"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("return nil")
			]),
			GoDecl.GoFuncDecl("hxrt_unserializerExtractNameField", null, [
				{
					name: "resolved",
					typeName: "any"
				}
			], ["string", "bool"], [
				GoStmt.GoRaw("rv := reflect.ValueOf(resolved)"),
				GoStmt.GoRaw("if !rv.IsValid() {"),
				GoStmt.GoRaw("\treturn \"\", false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("for rv.IsValid() && (rv.Kind() == reflect.Interface || rv.Kind() == reflect.Pointer) {"),
				GoStmt.GoRaw("\tif rv.IsNil() {"),
				GoStmt.GoRaw("\t\treturn \"\", false"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\trv = rv.Elem()"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if !rv.IsValid() || rv.Kind() != reflect.Struct {"),
				GoStmt.GoRaw("\treturn \"\", false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("field := rv.FieldByName(\"name\")"),
				GoStmt.GoRaw("if !field.IsValid() {"),
				GoStmt.GoRaw("\treturn \"\", false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("for field.IsValid() && field.Kind() == reflect.Interface {"),
				GoStmt.GoRaw("\tif field.IsNil() {"),
				GoStmt.GoRaw("\t\treturn \"\", false"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tfield = field.Elem()"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if !field.IsValid() {"),
				GoStmt.GoRaw("\treturn \"\", false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if field.Kind() == reflect.String {"),
				GoStmt.GoRaw("\treturn field.String(), true"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if field.Kind() == reflect.Pointer {"),
				GoStmt.GoRaw("\tif field.IsNil() {"),
				GoStmt.GoRaw("\t\treturn \"\", false"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tif field.Elem().Kind() == reflect.String {"),
				GoStmt.GoRaw("\t\treturn field.Elem().String(), true"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("return \"\", false")
			]),
			GoDecl.GoFuncDecl("hxrt_unserializerResolvedClassName", null, [
				{
					name: "resolved",
					typeName: "any"
				}
			], ["string", "bool"], [
				GoStmt.GoRaw("switch current := resolved.(type) {"),
				GoStmt.GoRaw("case haxe__SerializedClassRef:"),
				GoStmt.GoRaw("\treturn current.name, true"),
				GoStmt.GoRaw("case *haxe__SerializedClassRef:"),
				GoStmt.GoRaw("\tif current == nil {"),
				GoStmt.GoRaw("\t\treturn \"\", false"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn current.name, true"),
				GoStmt.GoRaw("case string:"),
				GoStmt.GoRaw("\treturn current, true"),
				GoStmt.GoRaw("case *string:"),
				GoStmt.GoRaw("\tif current == nil {"),
				GoStmt.GoRaw("\t\treturn \"\", false"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn *current, true"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("return hxrt_unserializerExtractNameField(resolved)")
			]),
			GoDecl.GoFuncDecl("hxrt_unserializerResolvedEnumName", null, [
				{
					name: "resolved",
					typeName: "any"
				}
			], ["string", "bool"], [
				GoStmt.GoRaw("switch current := resolved.(type) {"),
				GoStmt.GoRaw("case haxe__SerializedEnumRef:"),
				GoStmt.GoRaw("\treturn current.name, true"),
				GoStmt.GoRaw("case *haxe__SerializedEnumRef:"),
				GoStmt.GoRaw("\tif current == nil {"),
				GoStmt.GoRaw("\t\treturn \"\", false"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn current.name, true"),
				GoStmt.GoRaw("case string:"),
				GoStmt.GoRaw("\treturn current, true"),
				GoStmt.GoRaw("case *string:"),
				GoStmt.GoRaw("\tif current == nil {"),
				GoStmt.GoRaw("\t\treturn \"\", false"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn *current, true"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("return hxrt_unserializerExtractNameField(resolved)")
			]),
			GoDecl.GoFuncDecl("hxrt_unserializerCreateClassInstance", null, [
				{
					name: "className",
					typeName: "string"
				}
			], ["any", "bool"], classCreateBody),
			GoDecl.GoFuncDecl("hxrt_unserializerCreateEnumInstance", null, [
				{
					name: "enumName",
					typeName: "string"
				},
				{name: "constructorName", typeName: "string"},
				{name: "constructorIndex", typeName: "int"},
				{name: "hasConstructorIndex", typeName: "bool"},
				{name: "args", typeName: "[]any"}
			], ["any", "bool"], [
				GoStmt.GoRaw("if hasConstructorIndex {"),
				GoStmt.GoRaw("\tif _, ok := hxrt_serializerLookupEnumConstructorByName(enumName, constructorIndex); !ok {"),
				GoStmt.GoRaw("\t\treturn nil, false"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("} else {"),
				GoStmt.GoRaw("\tresolvedIndex, ok := hxrt_serializerLookupEnumIndexByName(enumName, constructorName)"),
				GoStmt.GoRaw("\tif !ok {"),
				GoStmt.GoRaw("\t\treturn nil, false"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tconstructorIndex = resolvedIndex"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("switch enumName {")
			].concat(enumCreateBody.slice(1))),
			GoDecl.GoStructDecl("haxe__Serializer", [
				{
					name: "buf",
					typeName: "*string"
				},
				{name: "useCache", typeName: "bool"},
				{name: "useEnumIndex", typeName: "bool"},
				{name: "stringCache", typeName: "map[string]int"},
				{name: "cacheRefs", typeName: "map[uintptr]int"}
			]),
			GoDecl.GoFuncDecl("New_haxe__Serializer", null, [], ["*haxe__Serializer"], [
				GoStmt.GoReturn(GoExpr.GoRaw("&haxe__Serializer{buf: hxrt.StringFromLiteral(\"\"), useCache: haxe__Serializer_USE_CACHE, useEnumIndex: haxe__Serializer_USE_ENUM_INDEX, stringCache: map[string]int{}, cacheRefs: map[uintptr]int{}}"))
			]),
			GoDecl.GoFuncDecl("serialize", {
				name: "self",
				typeName: "*haxe__Serializer"
			}, [{name: "value", typeName: "any"}], [], [
				GoStmt.GoIf(GoExpr.GoBinary("==", GoExpr.GoIdent("self"), GoExpr.GoNil), [GoStmt.GoReturn(null)], null),
				GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("hxrt_serializerWriteValue"), [GoExpr.GoIdent("self"), GoExpr.GoIdent("value")]))
			]),
			GoDecl.GoFuncDecl("serializeException", {
				name: "self",
				typeName: "*haxe__Serializer"
			}, [{name: "value", typeName: "any"}], [], [
				GoStmt.GoIf(GoExpr.GoBinary("==", GoExpr.GoIdent("self"), GoExpr.GoNil), [GoStmt.GoReturn(null)], null),
				GoStmt.GoRaw("hxrt_serializerAppend(self, \"x\")"),
				GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("hxrt_serializerWriteValue"), [GoExpr.GoIdent("self"), GoExpr.GoIdent("value")]))
			]),
			GoDecl.GoFuncDecl("toString", {
				name: "self",
				typeName: "*haxe__Serializer"
			}, [], ["*string"], [
				GoStmt.GoIf(GoExpr.GoRaw("self == nil || self.buf == nil"), [
					GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoStringLiteral("")]))
				], null),
				GoStmt.GoReturn(GoExpr.GoSelector(GoExpr.GoIdent("self"), "buf"))
			]),
			GoDecl.GoFuncDecl("hxrt_serializerAppend", null, [
				{
					name: "self",
					typeName: "*haxe__Serializer"
				},
				{name: "chunk", typeName: "string"}
			], [], [
				GoStmt.GoRaw("if self == nil {"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if self.buf == nil {"),
				GoStmt.GoRaw("\tself.buf = hxrt.StringFromLiteral(\"\")"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("self.buf = hxrt.StringFromLiteral(*hxrt.StdString(self.buf) + chunk)")
			]),
			GoDecl.GoFuncDecl("hxrt_serializerEscape", null, [
				{
					name: "value",
					typeName: "*string"
				}
			], ["string"], [
				GoStmt.GoRaw("raw := *hxrt.StdString(value)"),
				GoStmt.GoRaw("var builder strings.Builder"),
				GoStmt.GoRaw("hex := \"0123456789ABCDEF\""),
				GoStmt.GoRaw("for i := 0; i < len(raw); i++ {"),
				GoStmt.GoRaw("\tb := raw[i]"),
				GoStmt.GoRaw("\tif (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || (b >= '0' && b <= '9') || b == '_' || b == '.' || b == '-' {"),
				GoStmt.GoRaw("\t\tbuilder.WriteByte(b)"),
				GoStmt.GoRaw("\t\tcontinue"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tbuilder.WriteByte('%')"),
				GoStmt.GoRaw("\tbuilder.WriteByte(hex[b>>4])"),
				GoStmt.GoRaw("\tbuilder.WriteByte(hex[b&0x0F])"),
				GoStmt.GoRaw("}"),
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("builder"), "String"), []))
			]),
			GoDecl.GoFuncDecl("hxrt_serializerWriteStringToken", null, [
				{
					name: "self",
					typeName: "*haxe__Serializer"
				},
				{name: "value", typeName: "string"}
			], [], [
				GoStmt.GoRaw("if self == nil {"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if self.stringCache == nil {"),
				GoStmt.GoRaw("\tself.stringCache = map[string]int{}"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if index, ok := self.stringCache[value]; ok {"),
				GoStmt.GoRaw("\thxrt_serializerAppend(self, \"R\" + strconv.Itoa(index))"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("escaped := hxrt_serializerEscape(hxrt.StringFromLiteral(value))"),
				GoStmt.GoRaw("self.stringCache[value] = len(self.stringCache)"),
				GoStmt.GoRaw("hxrt_serializerAppend(self, \"y\" + strconv.Itoa(len(escaped)) + \":\" + escaped)")
			]),
			GoDecl.GoFuncDecl("hxrt_serializerWriteIntToken", null, [
				{
					name: "self",
					typeName: "*haxe__Serializer"
				},
				{name: "value", typeName: "int64"}
			], [], [
				GoStmt.GoRaw("if value == 0 {"),
				GoStmt.GoRaw("\thxrt_serializerAppend(self, \"z\")"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("hxrt_serializerAppend(self, \"i\" + strconv.FormatInt(value, 10))")
			]),
			GoDecl.GoFuncDecl("hxrt_serializerWriteBytesToken", null, [
				{
					name: "self",
					typeName: "*haxe__Serializer"
				},
				{name: "raw", typeName: "[]byte"}
			], [], [
				GoStmt.GoRaw("encoded := base64.StdEncoding.EncodeToString(raw)"),
				GoStmt.GoRaw("encoded = strings.TrimRight(encoded, \"=\")"),
				GoStmt.GoRaw("hxrt_serializerAppend(self, \"s\" + strconv.Itoa(len(encoded)) + \":\" + encoded)")
			]),
			GoDecl.GoFuncDecl("hxrt_serializerWriteEnumToken", null, [
				{
					name: "self",
					typeName: "*haxe__Serializer"
				},
				{name: "enumName", typeName: "string"},
				{name: "constructorName", typeName: "string"},
				{name: "constructorIndex", typeName: "int"},
				{name: "hasConstructorIndex", typeName: "bool"},
				{name: "args", typeName: "[]any"}
			], [], [
				GoStmt.GoRaw("if self != nil && self.useEnumIndex && hasConstructorIndex {"),
				GoStmt.GoRaw("\thxrt_serializerAppend(self, \"j\")"),
				GoStmt.GoRaw("\thxrt_serializerWriteStringToken(self, enumName)"),
				GoStmt.GoRaw("\thxrt_serializerAppend(self, \":\" + strconv.Itoa(constructorIndex))"),
				GoStmt.GoRaw("} else {"),
				GoStmt.GoRaw("\thxrt_serializerAppend(self, \"w\")"),
				GoStmt.GoRaw("\thxrt_serializerWriteStringToken(self, enumName)"),
				GoStmt.GoRaw("\thxrt_serializerWriteStringToken(self, constructorName)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("hxrt_serializerAppend(self, \":\" + strconv.Itoa(len(args)))"),
				GoStmt.GoRaw("for _, arg := range args {"),
				GoStmt.GoRaw("\thxrt_serializerWriteValue(self, arg)"),
				GoStmt.GoRaw("}")
			]),
			GoDecl.GoFuncDecl("hxrt_serializerWriteListToken", null, [
				{
					name: "self",
					typeName: "*haxe__Serializer"
				},
				{name: "items", typeName: "[]any"}
			], [], [
				GoStmt.GoRaw("hxrt_serializerAppend(self, \"l\")"),
				GoStmt.GoRaw("for _, item := range items {"),
				GoStmt.GoRaw("\thxrt_serializerWriteValue(self, item)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("hxrt_serializerAppend(self, \"h\")")
			]),
			GoDecl.GoFuncDecl("hxrt_serializerWriteStringMapToken", null, [
				{
					name: "self",
					typeName: "*haxe__Serializer"
				},
				{name: "entries", typeName: "map[string]any"}
			], [], [
				GoStmt.GoRaw("hxrt_serializerAppend(self, \"b\")"),
				GoStmt.GoRaw("keys := make([]string, 0, len(entries))"),
				GoStmt.GoRaw("for key := range entries {"),
				GoStmt.GoRaw("\tkeys = append(keys, key)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("sort.Strings(keys)"),
				GoStmt.GoRaw("for _, key := range keys {"),
				GoStmt.GoRaw("\thxrt_serializerWriteStringToken(self, key)"),
				GoStmt.GoRaw("\thxrt_serializerWriteValue(self, entries[key])"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("hxrt_serializerAppend(self, \"h\")")
			]),
			GoDecl.GoFuncDecl("hxrt_serializerWriteIntMapToken", null, [
				{
					name: "self",
					typeName: "*haxe__Serializer"
				},
				{name: "entries", typeName: "map[int]any"}
			], [], [
				GoStmt.GoRaw("hxrt_serializerAppend(self, \"q\")"),
				GoStmt.GoRaw("keys := make([]int, 0, len(entries))"),
				GoStmt.GoRaw("for key := range entries {"),
				GoStmt.GoRaw("\tkeys = append(keys, key)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("sort.Ints(keys)"),
				GoStmt.GoRaw("for _, key := range keys {"),
				GoStmt.GoRaw("\thxrt_serializerAppend(self, \":\" + strconv.Itoa(key))"),
				GoStmt.GoRaw("\thxrt_serializerWriteValue(self, entries[key])"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("hxrt_serializerAppend(self, \"h\")")
			]),
			GoDecl.GoFuncDecl("hxrt_serializerWriteObjectMapToken", null, [
				{
					name: "self",
					typeName: "*haxe__Serializer"
				},
				{name: "entries", typeName: "map[any]any"}
			], [], [
				GoStmt.GoRaw("hxrt_serializerAppend(self, \"M\")"),
				GoStmt.GoRaw("for key, value := range entries {"),
				GoStmt.GoRaw("\thxrt_serializerWriteValue(self, key)"),
				GoStmt.GoRaw("\thxrt_serializerWriteValue(self, value)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("hxrt_serializerAppend(self, \"h\")")
			]),
			GoDecl.GoFuncDecl("hxrt_serializerReflectAny", null, [
				{
					name: "value",
					typeName: "reflect.Value"
				}
			], ["any", "bool"], [
				GoStmt.GoRaw("defer func() {"),
				GoStmt.GoRaw("\tif recover() != nil {"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("}()"),
				GoStmt.GoRaw("if !value.IsValid() {"),
				GoStmt.GoRaw("\treturn nil, false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if value.CanInterface() {"),
				GoStmt.GoRaw("\treturn value.Interface(), true"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if !value.CanAddr() {"),
				GoStmt.GoRaw("\treturn nil, false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("lifted := reflect.NewAt(value.Type(), unsafe.Pointer(value.UnsafeAddr())).Elem()"),
				GoStmt.GoRaw("if !lifted.IsValid() || !lifted.CanInterface() {"),
				GoStmt.GoRaw("\treturn nil, false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("return lifted.Interface(), true")
			]),
			GoDecl.GoFuncDecl("hxrt_serializerWriteSerializedClass", null, [
				{
					name: "self",
					typeName: "*haxe__Serializer"
				},
				{name: "serialized", typeName: "*haxe__SerializedClass"}
			], [], [
				GoStmt.GoRaw("if serialized == nil {"),
				GoStmt.GoRaw("\thxrt_serializerAppend(self, \"n\")"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("hxrt_serializerAppend(self, \"c\")"),
				GoStmt.GoRaw("hxrt_serializerWriteStringToken(self, serialized.name)"),
				GoStmt.GoRaw("limit := len(serialized.fieldNames)"),
				GoStmt.GoRaw("if len(serialized.fieldValues) < limit {"),
				GoStmt.GoRaw("\tlimit = len(serialized.fieldValues)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("for i := 0; i < limit; i++ {"),
				GoStmt.GoRaw("\thxrt_serializerWriteStringToken(self, serialized.fieldNames[i])"),
				GoStmt.GoRaw("\thxrt_serializerWriteValue(self, serialized.fieldValues[i])"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("hxrt_serializerAppend(self, \"g\")")
			]),
			GoDecl.GoFuncDecl("hxrt_serializerWriteSerializedEnum", null, [
				{
					name: "self",
					typeName: "*haxe__Serializer"
				},
				{name: "serialized", typeName: "*haxe__SerializedEnum"}
			], [], [
				GoStmt.GoRaw("if serialized == nil {"),
				GoStmt.GoRaw("\thxrt_serializerAppend(self, \"n\")"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("constructorIndex := serialized.constructorIndex"),
				GoStmt.GoRaw("hasConstructorIndex := serialized.hasConstructorIndex"),
				GoStmt.GoRaw("if !hasConstructorIndex {"),
				GoStmt.GoRaw("\tif resolvedIndex, ok := hxrt_serializerLookupEnumIndexByName(serialized.name, serialized.constructor); ok {"),
				GoStmt.GoRaw("\t\tconstructorIndex = resolvedIndex"),
				GoStmt.GoRaw("\t\thasConstructorIndex = true"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("hxrt_serializerWriteEnumToken(self, serialized.name, serialized.constructor, constructorIndex, hasConstructorIndex, serialized.args)")
			]),
			GoDecl.GoFuncDecl("hxrt_serializerTryDsListStruct", null, [
				{
					name: "self",
					typeName: "*haxe__Serializer"
				},
				{name: "ref", typeName: "reflect.Value"}
			], ["bool"], [
				GoStmt.GoRaw("if !ref.IsValid() || ref.Kind() != reflect.Struct || ref.Type().Name() != \"haxe__ds__List\" {"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("itemsField := ref.FieldByName(\"items\")"),
				GoStmt.GoRaw("if !itemsField.IsValid() || itemsField.Kind() != reflect.Slice {"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("items := make([]any, 0, itemsField.Len())"),
				GoStmt.GoRaw("for i := 0; i < itemsField.Len(); i++ {"),
				GoStmt.GoRaw("\titem, ok := hxrt_serializerReflectAny(itemsField.Index(i))"),
				GoStmt.GoRaw("\tif !ok {"),
				GoStmt.GoRaw("\t\treturn false"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\titems = append(items, item)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("hxrt_serializerWriteListToken(self, items)"),
				GoStmt.GoRaw("return true")
			]),
			GoDecl.GoFuncDecl("hxrt_serializerTryDsStringMapStruct", null, [
				{
					name: "self",
					typeName: "*haxe__Serializer"
				},
				{name: "ref", typeName: "reflect.Value"}
			], ["bool"], [
				GoStmt.GoRaw("if !ref.IsValid() || ref.Kind() != reflect.Struct || ref.Type().Name() != \"haxe__ds__StringMap\" {"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("mapField := ref.FieldByName(\"h\")"),
				GoStmt.GoRaw("if !mapField.IsValid() || mapField.Kind() != reflect.Map {"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("entries := map[string]any{}"),
				GoStmt.GoRaw("for _, key := range mapField.MapKeys() {"),
				GoStmt.GoRaw("\tif key.Kind() != reflect.String {"),
				GoStmt.GoRaw("\t\treturn false"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tvalue, ok := hxrt_serializerReflectAny(mapField.MapIndex(key))"),
				GoStmt.GoRaw("\tif !ok {"),
				GoStmt.GoRaw("\t\treturn false"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tentries[key.String()] = value"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("hxrt_serializerWriteStringMapToken(self, entries)"),
				GoStmt.GoRaw("return true")
			]),
			GoDecl.GoFuncDecl("hxrt_serializerTryDsIntMapStruct", null, [
				{
					name: "self",
					typeName: "*haxe__Serializer"
				},
				{name: "ref", typeName: "reflect.Value"}
			], ["bool"], [
				GoStmt.GoRaw("if !ref.IsValid() || ref.Kind() != reflect.Struct || ref.Type().Name() != \"haxe__ds__IntMap\" {"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("mapField := ref.FieldByName(\"h\")"),
				GoStmt.GoRaw("if !mapField.IsValid() || mapField.Kind() != reflect.Map {"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("entries := map[int]any{}"),
				GoStmt.GoRaw("for _, key := range mapField.MapKeys() {"),
				GoStmt.GoRaw("\tvar intKey int"),
				GoStmt.GoRaw("\tswitch key.Kind() {"),
				GoStmt.GoRaw("\tcase reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:"),
				GoStmt.GoRaw("\t\tintKey = int(key.Int())"),
				GoStmt.GoRaw("\tdefault:"),
				GoStmt.GoRaw("\t\treturn false"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tvalue, ok := hxrt_serializerReflectAny(mapField.MapIndex(key))"),
				GoStmt.GoRaw("\tif !ok {"),
				GoStmt.GoRaw("\t\treturn false"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tentries[intKey] = value"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("hxrt_serializerWriteIntMapToken(self, entries)"),
				GoStmt.GoRaw("return true")
			]),
			GoDecl.GoFuncDecl("hxrt_serializerTryDsObjectMapStruct", null, [
				{
					name: "self",
					typeName: "*haxe__Serializer"
				},
				{name: "ref", typeName: "reflect.Value"}
			], ["bool"], [
				GoStmt.GoRaw("if !ref.IsValid() || ref.Kind() != reflect.Struct || ref.Type().Name() != \"haxe__ds__ObjectMap\" {"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("mapField := ref.FieldByName(\"h\")"),
				GoStmt.GoRaw("if !mapField.IsValid() || mapField.Kind() != reflect.Map {"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("entries := map[any]any{}"),
				GoStmt.GoRaw("for _, key := range mapField.MapKeys() {"),
				GoStmt.GoRaw("\tkeyAny, ok := hxrt_serializerReflectAny(key)"),
				GoStmt.GoRaw("\tif !ok {"),
				GoStmt.GoRaw("\t\treturn false"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tvalueAny, ok := hxrt_serializerReflectAny(mapField.MapIndex(key))"),
				GoStmt.GoRaw("\tif !ok {"),
				GoStmt.GoRaw("\t\treturn false"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tentries[keyAny] = valueAny"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("hxrt_serializerWriteObjectMapToken(self, entries)"),
				GoStmt.GoRaw("return true")
			]),
			GoDecl.GoFuncDecl("hxrt_serializerTryClassStruct", null, [
				{
					name: "self",
					typeName: "*haxe__Serializer"
				},
				{name: "value", typeName: "any"},
				{name: "ref", typeName: "reflect.Value"}
			], ["bool"], [
				GoStmt.GoRaw("defer func() {"),
				GoStmt.GoRaw("\tif recover() != nil {"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("}()"),
				GoStmt.GoRaw("if !ref.IsValid() || ref.Kind() != reflect.Struct {"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("className, ok := hxrt_serializerLookupClassName(ref.Type().Name())"),
				GoStmt.GoRaw("if !ok {"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if hxrt_serializerTrackRef(self, value) {"),
				GoStmt.GoRaw("\treturn true"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if custom, ok := value.(interface{ hxSerialize(*haxe__Serializer) }); ok {"),
				GoStmt.GoRaw("\thxrt_serializerAppend(self, \"C\")"),
				GoStmt.GoRaw("\thxrt_serializerWriteStringToken(self, className)"),
				GoStmt.GoRaw("\tcustom.hxSerialize(self)"),
				GoStmt.GoRaw("\thxrt_serializerAppend(self, \"g\")"),
				GoStmt.GoRaw("\treturn true"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if custom, ok := value.(interface{ HxSerialize(*haxe__Serializer) }); ok {"),
				GoStmt.GoRaw("\thxrt_serializerAppend(self, \"C\")"),
				GoStmt.GoRaw("\thxrt_serializerWriteStringToken(self, className)"),
				GoStmt.GoRaw("\tcustom.HxSerialize(self)"),
				GoStmt.GoRaw("\thxrt_serializerAppend(self, \"g\")"),
				GoStmt.GoRaw("\treturn true"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if ref.CanAddr() {"),
				GoStmt.GoRaw("\taddr := ref.Addr().Interface()"),
				GoStmt.GoRaw("\tif custom, ok := addr.(interface{ hxSerialize(*haxe__Serializer) }); ok {"),
				GoStmt.GoRaw("\t\thxrt_serializerAppend(self, \"C\")"),
				GoStmt.GoRaw("\t\thxrt_serializerWriteStringToken(self, className)"),
				GoStmt.GoRaw("\t\tcustom.hxSerialize(self)"),
				GoStmt.GoRaw("\t\thxrt_serializerAppend(self, \"g\")"),
				GoStmt.GoRaw("\t\treturn true"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tif custom, ok := addr.(interface{ HxSerialize(*haxe__Serializer) }); ok {"),
				GoStmt.GoRaw("\t\thxrt_serializerAppend(self, \"C\")"),
				GoStmt.GoRaw("\t\thxrt_serializerWriteStringToken(self, className)"),
				GoStmt.GoRaw("\t\tcustom.HxSerialize(self)"),
				GoStmt.GoRaw("\t\thxrt_serializerAppend(self, \"g\")"),
				GoStmt.GoRaw("\t\treturn true"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("hxrt_serializerAppend(self, \"c\")"),
				GoStmt.GoRaw("hxrt_serializerWriteStringToken(self, className)"),
				GoStmt.GoRaw("refType := ref.Type()"),
				GoStmt.GoRaw("for i := 0; i < ref.NumField(); i++ {"),
				GoStmt.GoRaw("\tfieldInfo := refType.Field(i)"),
				GoStmt.GoRaw("\tfieldName := fieldInfo.Name"),
				GoStmt.GoRaw("\tif strings.HasPrefix(fieldName, \"__hx_\") {"),
				GoStmt.GoRaw("\t\tcontinue"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tfieldValue, ok := hxrt_serializerReflectAny(ref.Field(i))"),
				GoStmt.GoRaw("\tif !ok {"),
				GoStmt.GoRaw("\t\treturn false"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\thxrt_serializerWriteStringToken(self, fieldName)"),
				GoStmt.GoRaw("\thxrt_serializerWriteValue(self, fieldValue)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("hxrt_serializerAppend(self, \"g\")"),
				GoStmt.GoRaw("return true")
			]),
			GoDecl.GoFuncDecl("hxrt_serializerTryEnumStruct", null, [
				{
					name: "self",
					typeName: "*haxe__Serializer"
				},
				{name: "value", typeName: "any"},
				{name: "ref", typeName: "reflect.Value"}
			], ["bool"], [
				GoStmt.GoRaw("defer func() {"),
				GoStmt.GoRaw("\tif recover() != nil {"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("}()"),
				GoStmt.GoRaw("if !ref.IsValid() || ref.Kind() != reflect.Struct {"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("tagField := ref.FieldByName(\"tag\")"),
				GoStmt.GoRaw("paramsField := ref.FieldByName(\"params\")"),
				GoStmt.GoRaw("if !tagField.IsValid() || !paramsField.IsValid() || paramsField.Kind() != reflect.Slice {"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("var tag int"),
				GoStmt.GoRaw("switch tagField.Kind() {"),
				GoStmt.GoRaw("case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:"),
				GoStmt.GoRaw("\ttag = int(tagField.Int())"),
				GoStmt.GoRaw("default:"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("enumName, constructorName, ok := hxrt_serializerLookupEnumConstructor(ref.Type().Name(), tag)"),
				GoStmt.GoRaw("if !ok {"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if hxrt_serializerTrackRef(self, value) {"),
				GoStmt.GoRaw("\treturn true"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("args := make([]any, 0, paramsField.Len())"),
				GoStmt.GoRaw("for i := 0; i < paramsField.Len(); i++ {"),
				GoStmt.GoRaw("\tvalue, ok := hxrt_serializerReflectAny(paramsField.Index(i))"),
				GoStmt.GoRaw("\tif !ok {"),
				GoStmt.GoRaw("\t\treturn false"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\targs = append(args, value)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("hxrt_serializerWriteEnumToken(self, enumName, constructorName, tag, true, args)"),
				GoStmt.GoRaw("return true")
			]),
			GoDecl.GoFuncDecl("hxrt_serializerTryDateStruct", null, [
				{
					name: "self",
					typeName: "*haxe__Serializer"
				},
				{name: "ref", typeName: "reflect.Value"}
			], ["bool"], [
				GoStmt.GoRaw("defer func() {"),
				GoStmt.GoRaw("\tif recover() != nil {"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("}()"),
				GoStmt.GoRaw("if !ref.IsValid() || ref.Kind() != reflect.Struct {"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("valueField := ref.FieldByName(\"value\")"),
				GoStmt.GoRaw("if !valueField.IsValid() || !valueField.CanAddr() {"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("fieldType := valueField.Type()"),
				GoStmt.GoRaw("if fieldType.PkgPath() != \"time\" || fieldType.Name() != \"Time\" {"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("timeAny := reflect.NewAt(fieldType, unsafe.Pointer(valueField.UnsafeAddr())).Elem().Interface()"),
				GoStmt.GoRaw("timeValue, ok := timeAny.(time.Time)"),
				GoStmt.GoRaw("if !ok {"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("ms := float64(timeValue.UnixNano()) / 1000000.0"),
				GoStmt.GoRaw("hxrt_serializerAppend(self, \"v\" + strconv.FormatFloat(ms, 'g', -1, 64))"),
				GoStmt.GoRaw("return true")
			]),
			GoDecl.GoFuncDecl("hxrt_serializerTrySpecialReflect", null, [
				{
					name: "self",
					typeName: "*haxe__Serializer"
				},
				{name: "value", typeName: "any"}
			], ["bool"], [
				GoStmt.GoRaw("ref := reflect.ValueOf(value)"),
				GoStmt.GoRaw("if !ref.IsValid() {"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("for ref.Kind() == reflect.Interface {"),
				GoStmt.GoRaw("\tif ref.IsNil() {"),
				GoStmt.GoRaw("\t\treturn false"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tref = ref.Elem()"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("for ref.Kind() == reflect.Pointer {"),
				GoStmt.GoRaw("\tif ref.IsNil() {"),
				GoStmt.GoRaw("\t\treturn false"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tref = ref.Elem()"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if !ref.IsValid() || ref.Kind() != reflect.Struct {"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("typeName := ref.Type().Name()"),
				GoStmt.GoRaw("if typeName == \"Date\" {"),
				GoStmt.GoRaw("\tif hxrt_serializerTryDateStruct(self, ref) {"),
				GoStmt.GoRaw("\t\treturn true"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if typeName == \"haxe__io__Bytes\" {"),
				GoStmt.GoRaw("\tbytesField := ref.FieldByName(\"b\")"),
				GoStmt.GoRaw("\tif !bytesField.IsValid() || bytesField.Kind() != reflect.Slice {"),
				GoStmt.GoRaw("\t\treturn false"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\traw := make([]byte, bytesField.Len())"),
				GoStmt.GoRaw("\tfor i := 0; i < bytesField.Len(); i++ {"),
				GoStmt.GoRaw("\t\tentry := bytesField.Index(i)"),
				GoStmt.GoRaw("\t\tif !entry.IsValid() {"),
				GoStmt.GoRaw("\t\t\treturn false"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t\tswitch entry.Kind() {"),
				GoStmt.GoRaw("\t\tcase reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:"),
				GoStmt.GoRaw("\t\t\traw[i] = byte(entry.Int())"),
				GoStmt.GoRaw("\t\tcase reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:"),
				GoStmt.GoRaw("\t\t\traw[i] = byte(entry.Uint())"),
				GoStmt.GoRaw("\t\tdefault:"),
				GoStmt.GoRaw("\t\t\treturn false"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\thxrt_serializerWriteBytesToken(self, raw)"),
				GoStmt.GoRaw("\treturn true"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if hxrt_serializerTryDsListStruct(self, ref) {"),
				GoStmt.GoRaw("\treturn true"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if hxrt_serializerTryDsStringMapStruct(self, ref) {"),
				GoStmt.GoRaw("\treturn true"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if hxrt_serializerTryDsIntMapStruct(self, ref) {"),
				GoStmt.GoRaw("\treturn true"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if hxrt_serializerTryDsObjectMapStruct(self, ref) {"),
				GoStmt.GoRaw("\treturn true"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if hxrt_serializerTryEnumStruct(self, value, ref) {"),
				GoStmt.GoRaw("\treturn true"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if hxrt_serializerTryClassStruct(self, value, ref) {"),
				GoStmt.GoRaw("\treturn true"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("return false")
			]),
			GoDecl.GoFuncDecl("hxrt_serializerTryTypeValueRef", null, [
				{
					name: "self",
					typeName: "*haxe__Serializer"
				},
				{name: "value", typeName: "any"}
			], ["bool"], [
				GoStmt.GoRaw("ref := reflect.ValueOf(value)"),
				GoStmt.GoRaw("if !ref.IsValid() {"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("for ref.Kind() == reflect.Interface || ref.Kind() == reflect.Pointer {"),
				GoStmt.GoRaw("\tif ref.IsNil() {"),
				GoStmt.GoRaw("\t\treturn false"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tref = ref.Elem()"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if !ref.IsValid() || ref.Kind() != reflect.Struct {"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("typeName := ref.Type().Name()"),
				GoStmt.GoRaw("if typeName != \"hxrt__TypeClassValue\" && typeName != \"hxrt__TypeEnumValue\" {"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("nameField := ref.FieldByName(\"name\")"),
				GoStmt.GoRaw("if !nameField.IsValid() {"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("for nameField.IsValid() && nameField.Kind() == reflect.Interface {"),
				GoStmt.GoRaw("\tif nameField.IsNil() {"),
				GoStmt.GoRaw("\t\treturn false"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tnameField = nameField.Elem()"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if !nameField.IsValid() {"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("resolvedName := \"\""),
				GoStmt.GoRaw("if nameField.Kind() == reflect.String {"),
				GoStmt.GoRaw("\tresolvedName = nameField.String()"),
				GoStmt.GoRaw("} else if nameField.Kind() == reflect.Pointer {"),
				GoStmt.GoRaw("\tif nameField.IsNil() || nameField.Elem().Kind() != reflect.String {"),
				GoStmt.GoRaw("\t\treturn false"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tresolvedName = nameField.Elem().String()"),
				GoStmt.GoRaw("} else {"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if resolvedName == \"\" {"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if typeName == \"hxrt__TypeClassValue\" {"),
				GoStmt.GoRaw("\thxrt_serializerAppend(self, \"A\")"),
				GoStmt.GoRaw("} else {"),
				GoStmt.GoRaw("\thxrt_serializerAppend(self, \"B\")"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("hxrt_serializerWriteStringToken(self, resolvedName)"),
				GoStmt.GoRaw("return true")
			]),
			GoDecl.GoFuncDecl("hxrt_serializerTrackRef", null, [
				{
					name: "self",
					typeName: "*haxe__Serializer"
				},
				{name: "value", typeName: "any"}
			], ["bool"], [
				GoStmt.GoRaw("if self == nil || !self.useCache {"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if self.cacheRefs == nil {"),
				GoStmt.GoRaw("\tself.cacheRefs = map[uintptr]int{}"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("ref := reflect.ValueOf(value)"),
				GoStmt.GoRaw("if !ref.IsValid() {"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("for ref.Kind() == reflect.Interface {"),
				GoStmt.GoRaw("\tif ref.IsNil() {"),
				GoStmt.GoRaw("\t\treturn false"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tref = ref.Elem()"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("switch ref.Kind() {"),
				GoStmt.GoRaw("case reflect.Map, reflect.Slice, reflect.Pointer:"),
				GoStmt.GoRaw("\tif ref.IsNil() {"),
				GoStmt.GoRaw("\t\treturn false"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tkey := ref.Pointer()"),
				GoStmt.GoRaw("\tif index, ok := self.cacheRefs[key]; ok {"),
				GoStmt.GoRaw("\t\thxrt_serializerAppend(self, \"r\" + strconv.Itoa(index))"),
				GoStmt.GoRaw("\t\treturn true"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tself.cacheRefs[key] = len(self.cacheRefs)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("return false")
			]),
			GoDecl.GoFuncDecl("haxe__Serializer_run", null, [
				{
					name: "value",
					typeName: "any"
				}
			], ["*string"], [
				GoStmt.GoRaw("serializer := New_haxe__Serializer()"),
				GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("serializer"), "serialize"), [GoExpr.GoIdent("value")])),
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("serializer"), "toString"), []))
			]),
			GoDecl.GoFuncDecl("hxrt_serializerWriteValue", null, [
				{
					name: "self",
					typeName: "*haxe__Serializer"
				},
				{name: "value", typeName: "any"}
			], [], [
				GoStmt.GoRaw("if self == nil {"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if value == nil {"),
				GoStmt.GoRaw("\thxrt_serializerAppend(self, \"n\")"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if hxrt_serializerTryTypeValueRef(self, value) {"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("switch current := value.(type) {"),
				GoStmt.GoRaw("case bool:"),
				GoStmt.GoRaw("\tif current {"),
				GoStmt.GoRaw("\t\thxrt_serializerAppend(self, \"t\")"),
				GoStmt.GoRaw("\t} else {"),
				GoStmt.GoRaw("\t\thxrt_serializerAppend(self, \"f\")"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("case string:"),
				GoStmt.GoRaw("\thxrt_serializerWriteStringToken(self, current)"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("case *string:"),
				GoStmt.GoRaw("\tif current == nil {"),
				GoStmt.GoRaw("\t\thxrt_serializerAppend(self, \"n\")"),
				GoStmt.GoRaw("\t} else {"),
				GoStmt.GoRaw("\t\thxrt_serializerWriteStringToken(self, *current)"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("case haxe__SerializedDate:"),
				GoStmt.GoRaw("\thxrt_serializerAppend(self, \"v\" + strconv.FormatFloat(current.ms, 'g', -1, 64))"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("case *haxe__SerializedDate:"),
				GoStmt.GoRaw("\tif current == nil {"),
				GoStmt.GoRaw("\t\thxrt_serializerAppend(self, \"n\")"),
				GoStmt.GoRaw("\t} else {"),
				GoStmt.GoRaw("\t\thxrt_serializerAppend(self, \"v\" + strconv.FormatFloat(current.ms, 'g', -1, 64))"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("case haxe__SerializedBytes:"),
				GoStmt.GoRaw("\thxrt_serializerWriteBytesToken(self, current.data)"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("case *haxe__SerializedBytes:"),
				GoStmt.GoRaw("\tif current == nil {"),
				GoStmt.GoRaw("\t\thxrt_serializerAppend(self, \"n\")"),
				GoStmt.GoRaw("\t} else {"),
				GoStmt.GoRaw("\t\thxrt_serializerWriteBytesToken(self, current.data)"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("case haxe__SerializedClass:"),
				GoStmt.GoRaw("\thxrt_serializerWriteSerializedClass(self, &current)"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("case *haxe__SerializedClass:"),
				GoStmt.GoRaw("\thxrt_serializerWriteSerializedClass(self, current)"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("case haxe__SerializedEnum:"),
				GoStmt.GoRaw("\thxrt_serializerWriteSerializedEnum(self, &current)"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("case *haxe__SerializedEnum:"),
				GoStmt.GoRaw("\thxrt_serializerWriteSerializedEnum(self, current)"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("case haxe__SerializedClassRef:"),
				GoStmt.GoRaw("\thxrt_serializerAppend(self, \"A\")"),
				GoStmt.GoRaw("\thxrt_serializerWriteStringToken(self, current.name)"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("case *haxe__SerializedClassRef:"),
				GoStmt.GoRaw("\tif current == nil {"),
				GoStmt.GoRaw("\t\thxrt_serializerAppend(self, \"n\")"),
				GoStmt.GoRaw("\t} else {"),
				GoStmt.GoRaw("\t\thxrt_serializerAppend(self, \"A\")"),
				GoStmt.GoRaw("\t\thxrt_serializerWriteStringToken(self, current.name)"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("case haxe__SerializedEnumRef:"),
				GoStmt.GoRaw("\thxrt_serializerAppend(self, \"B\")"),
				GoStmt.GoRaw("\thxrt_serializerWriteStringToken(self, current.name)"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("case *haxe__SerializedEnumRef:"),
				GoStmt.GoRaw("\tif current == nil {"),
				GoStmt.GoRaw("\t\thxrt_serializerAppend(self, \"n\")"),
				GoStmt.GoRaw("\t} else {"),
				GoStmt.GoRaw("\t\thxrt_serializerAppend(self, \"B\")"),
				GoStmt.GoRaw("\t\thxrt_serializerWriteStringToken(self, current.name)"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("case *haxe__ds__List:"),
				GoStmt.GoRaw("\tif current == nil {"),
				GoStmt.GoRaw("\t\thxrt_serializerAppend(self, \"n\")"),
				GoStmt.GoRaw("\t} else {"),
				GoStmt.GoRaw("\t\tif hxrt_serializerTrackRef(self, current) {"),
				GoStmt.GoRaw("\t\t\treturn"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t\thxrt_serializerWriteListToken(self, current.items)"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("case *haxe__ds__StringMap:"),
				GoStmt.GoRaw("\tif current == nil {"),
				GoStmt.GoRaw("\t\thxrt_serializerAppend(self, \"n\")"),
				GoStmt.GoRaw("\t} else {"),
				GoStmt.GoRaw("\t\tif hxrt_serializerTrackRef(self, current) {"),
				GoStmt.GoRaw("\t\t\treturn"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t\thxrt_serializerWriteStringMapToken(self, current.h)"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("case *haxe__ds__IntMap:"),
				GoStmt.GoRaw("\tif current == nil {"),
				GoStmt.GoRaw("\t\thxrt_serializerAppend(self, \"n\")"),
				GoStmt.GoRaw("\t} else {"),
				GoStmt.GoRaw("\t\tif hxrt_serializerTrackRef(self, current) {"),
				GoStmt.GoRaw("\t\t\treturn"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t\thxrt_serializerWriteIntMapToken(self, current.h)"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("case *haxe__ds__ObjectMap:"),
				GoStmt.GoRaw("\tif current == nil {"),
				GoStmt.GoRaw("\t\thxrt_serializerAppend(self, \"n\")"),
				GoStmt.GoRaw("\t} else {"),
				GoStmt.GoRaw("\t\tif hxrt_serializerTrackRef(self, current) {"),
				GoStmt.GoRaw("\t\t\treturn"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t\thxrt_serializerWriteObjectMapToken(self, current.h)"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("case int:"),
				GoStmt.GoRaw("\thxrt_serializerWriteIntToken(self, int64(current))"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("case int8:"),
				GoStmt.GoRaw("\thxrt_serializerWriteIntToken(self, int64(current))"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("case int16:"),
				GoStmt.GoRaw("\thxrt_serializerWriteIntToken(self, int64(current))"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("case int32:"),
				GoStmt.GoRaw("\thxrt_serializerWriteIntToken(self, int64(current))"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("case int64:"),
				GoStmt.GoRaw("\thxrt_serializerWriteIntToken(self, current)"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("case uint:"),
				GoStmt.GoRaw("\thxrt_serializerWriteIntToken(self, int64(current))"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("case uint8:"),
				GoStmt.GoRaw("\thxrt_serializerWriteIntToken(self, int64(current))"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("case uint16:"),
				GoStmt.GoRaw("\thxrt_serializerWriteIntToken(self, int64(current))"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("case uint32:"),
				GoStmt.GoRaw("\thxrt_serializerWriteIntToken(self, int64(current))"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("case uint64:"),
				GoStmt.GoRaw("\thxrt_serializerWriteIntToken(self, int64(current))"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("case float32:"),
				GoStmt.GoRaw("\tvalue64 := float64(current)"),
				GoStmt.GoRaw("\tif math.IsNaN(value64) {"),
				GoStmt.GoRaw("\t\thxrt_serializerAppend(self, \"k\")"),
				GoStmt.GoRaw("\t\treturn"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tif math.IsInf(value64, 1) {"),
				GoStmt.GoRaw("\t\thxrt_serializerAppend(self, \"p\")"),
				GoStmt.GoRaw("\t\treturn"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tif math.IsInf(value64, -1) {"),
				GoStmt.GoRaw("\t\thxrt_serializerAppend(self, \"m\")"),
				GoStmt.GoRaw("\t\treturn"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\thxrt_serializerAppend(self, \"d\" + strconv.FormatFloat(value64, 'g', -1, 64))"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("case float64:"),
				GoStmt.GoRaw("\tif math.IsNaN(current) {"),
				GoStmt.GoRaw("\t\thxrt_serializerAppend(self, \"k\")"),
				GoStmt.GoRaw("\t\treturn"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tif math.IsInf(current, 1) {"),
				GoStmt.GoRaw("\t\thxrt_serializerAppend(self, \"p\")"),
				GoStmt.GoRaw("\t\treturn"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tif math.IsInf(current, -1) {"),
				GoStmt.GoRaw("\t\thxrt_serializerAppend(self, \"m\")"),
				GoStmt.GoRaw("\t\treturn"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\thxrt_serializerAppend(self, \"d\" + strconv.FormatFloat(current, 'g', -1, 64))"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("ref := reflect.ValueOf(value)"),
				GoStmt.GoRaw("if !ref.IsValid() {"),
				GoStmt.GoRaw("\thxrt_serializerAppend(self, \"n\")"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if hxrt_serializerTrySpecialReflect(self, value) {"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("switch ref.Kind() {"),
				GoStmt.GoRaw("case reflect.Slice, reflect.Array:"),
				GoStmt.GoRaw("\tif ref.Kind() == reflect.Slice && ref.IsNil() {"),
				GoStmt.GoRaw("\t\thxrt_serializerAppend(self, \"n\")"),
				GoStmt.GoRaw("\t\treturn"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tif hxrt_serializerTrackRef(self, value) {"),
				GoStmt.GoRaw("\t\treturn"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\thxrt_serializerAppend(self, \"a\")"),
				GoStmt.GoRaw("\tnullRun := 0"),
				GoStmt.GoRaw("\tfor i := 0; i < ref.Len(); i++ {"),
				GoStmt.GoRaw("\t\titem := ref.Index(i).Interface()"),
				GoStmt.GoRaw("\t\tif item == nil {"),
				GoStmt.GoRaw("\t\t\tnullRun++"),
				GoStmt.GoRaw("\t\t\tcontinue"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t\tif nullRun > 1 {"),
				GoStmt.GoRaw("\t\t\thxrt_serializerAppend(self, \"u\" + strconv.Itoa(nullRun))"),
				GoStmt.GoRaw("\t\t} else if nullRun == 1 {"),
				GoStmt.GoRaw("\t\t\thxrt_serializerAppend(self, \"n\")"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t\tnullRun = 0"),
				GoStmt.GoRaw("\t\thxrt_serializerWriteValue(self, item)"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tif nullRun > 1 {"),
				GoStmt.GoRaw("\t\thxrt_serializerAppend(self, \"u\" + strconv.Itoa(nullRun))"),
				GoStmt.GoRaw("\t} else if nullRun == 1 {"),
				GoStmt.GoRaw("\t\thxrt_serializerAppend(self, \"n\")"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\thxrt_serializerAppend(self, \"h\")"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("case reflect.Map:"),
				GoStmt.GoRaw("\tif ref.IsNil() {"),
				GoStmt.GoRaw("\t\thxrt_serializerAppend(self, \"n\")"),
				GoStmt.GoRaw("\t\treturn"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tif ref.Type().Key().Kind() != reflect.String {"),
				GoStmt.GoRaw("\t\thxrt.Throw(hxrt.StringFromLiteral(\"Serializer map keys must be strings\"))"),
				GoStmt.GoRaw("\t\thxrt_serializerAppend(self, \"n\")"),
				GoStmt.GoRaw("\t\treturn"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tif hxrt_serializerTrackRef(self, value) {"),
				GoStmt.GoRaw("\t\treturn"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\thxrt_serializerAppend(self, \"o\")"),
				GoStmt.GoRaw("\tkeys := ref.MapKeys()"),
				GoStmt.GoRaw("\tsortedKeys := make([]string, 0, len(keys))"),
				GoStmt.GoRaw("\tfor _, key := range keys {"),
				GoStmt.GoRaw("\t\tsortedKeys = append(sortedKeys, key.String())"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tsort.Strings(sortedKeys)"),
				GoStmt.GoRaw("\tfor _, key := range sortedKeys {"),
				GoStmt.GoRaw("\t\thxrt_serializerWriteStringToken(self, key)"),
				GoStmt.GoRaw("\t\tvalueRef := ref.MapIndex(reflect.ValueOf(key))"),
				GoStmt.GoRaw("\t\tif valueRef.IsValid() {"),
				GoStmt.GoRaw("\t\t\thxrt_serializerWriteValue(self, valueRef.Interface())"),
				GoStmt.GoRaw("\t\t} else {"),
				GoStmt.GoRaw("\t\t\thxrt_serializerWriteValue(self, nil)"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\thxrt_serializerAppend(self, \"g\")"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("case reflect.Pointer:"),
				GoStmt.GoRaw("\tif ref.IsNil() {"),
				GoStmt.GoRaw("\t\thxrt_serializerAppend(self, \"n\")"),
				GoStmt.GoRaw("\t\treturn"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tif hxrt_serializerTrackRef(self, value) {"),
				GoStmt.GoRaw("\t\treturn"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\thxrt_serializerWriteValue(self, ref.Elem().Interface())"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("hxrt.Throw(hxrt.StringFromLiteral(\"Unsupported serializer value type\"))"),
				GoStmt.GoRaw("hxrt_serializerAppend(self, \"n\")")
			]),
			GoDecl.GoStructDecl("haxe__Unserializer", [
				{
					name: "buf",
					typeName: "*string"
				},
				{name: "pos", typeName: "int"},
				{name: "stringCache", typeName: "[]*string"},
				{name: "cache", typeName: "[]any"},
				{name: "resolver", typeName: "any"}
			]),
			GoDecl.GoFuncDecl("New_haxe__Unserializer", null, [{name: "buf", typeName: "*string"}], ["*haxe__Unserializer"], [
				GoStmt.GoRaw("resolver := haxe__Unserializer_DEFAULT_RESOLVER"),
				GoStmt.GoRaw("if resolver == nil {"),
				GoStmt.GoRaw("\tresolver = &haxe__Unserializer__DefaultResolver{}"),
				GoStmt.GoRaw("\thaxe__Unserializer_DEFAULT_RESOLVER = resolver"),
				GoStmt.GoRaw("}"),
				GoStmt.GoReturn(GoExpr.GoRaw("&haxe__Unserializer{buf: buf, pos: 0, stringCache: []*string{}, cache: []any{}, resolver: resolver}"))
			]),
			GoDecl.GoFuncDecl("setResolver", {
				name: "self",
				typeName: "*haxe__Unserializer"
			}, [{name: "resolver", typeName: "any"}], [], [
				GoStmt.GoRaw("if self == nil {"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if resolver == nil {"),
				GoStmt.GoRaw("\tself.resolver = haxe__Unserializer_NULL_RESOLVER"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("self.resolver = resolver")
			]),
			GoDecl.GoFuncDecl("getResolver", {
				name: "self",
				typeName: "*haxe__Unserializer"
			}, [], ["any"], [
				GoStmt.GoRaw("if self == nil {"),
				GoStmt.GoRaw("\treturn nil"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("return self.resolver")
			]),
			GoDecl.GoFuncDecl("unserialize", {
				name: "self",
				typeName: "*haxe__Unserializer"
			}, [], ["any"], [
				GoStmt.GoIf(GoExpr.GoRaw("self == nil || self.buf == nil"), [GoStmt.GoReturn(GoExpr.GoNil)], null),
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("haxe__Unserializer_readValue"), [GoExpr.GoIdent("self")]))
			]),
			GoDecl.GoFuncDecl("haxe__Unserializer_readUInt", null, [
				{
					name: "self",
					typeName: "*haxe__Unserializer"
				}
			], ["int"], [
				GoStmt.GoRaw("raw := *hxrt.StdString(self.buf)"),
				GoStmt.GoRaw("start := self.pos"),
				GoStmt.GoRaw("for self.pos < len(raw) && raw[self.pos] >= '0' && raw[self.pos] <= '9' {"),
				GoStmt.GoRaw("\tself.pos++"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if self.pos == start {"),
				GoStmt.GoRaw("\thxrt.Throw(hxrt.StringFromLiteral(\"Invalid serialized integer\"))"),
				GoStmt.GoRaw("\treturn 0"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("parsed, err := strconv.Atoi(raw[start:self.pos])"),
				GoStmt.GoRaw("if err != nil {"),
				GoStmt.GoRaw("\thxrt.Throw(err)"),
				GoStmt.GoRaw("\treturn 0"),
				GoStmt.GoRaw("}"),
				GoStmt.GoReturn(GoExpr.GoIdent("parsed"))
			]),
			GoDecl.GoFuncDecl("haxe__Unserializer_readDigits", null, [
				{
					name: "self",
					typeName: "*haxe__Unserializer"
				}
			], ["int"], [
				GoStmt.GoRaw("raw := *hxrt.StdString(self.buf)"),
				GoStmt.GoRaw("start := self.pos"),
				GoStmt.GoRaw("if self.pos < len(raw) && (raw[self.pos] == '-' || raw[self.pos] == '+') {"),
				GoStmt.GoRaw("\tself.pos++"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("digitStart := self.pos"),
				GoStmt.GoRaw("for self.pos < len(raw) && raw[self.pos] >= '0' && raw[self.pos] <= '9' {"),
				GoStmt.GoRaw("\tself.pos++"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if self.pos == digitStart {"),
				GoStmt.GoRaw("\thxrt.Throw(hxrt.StringFromLiteral(\"Invalid serialized integer\"))"),
				GoStmt.GoRaw("\treturn 0"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("parsed, err := strconv.Atoi(raw[start:self.pos])"),
				GoStmt.GoRaw("if err != nil {"),
				GoStmt.GoRaw("\thxrt.Throw(err)"),
				GoStmt.GoRaw("\treturn 0"),
				GoStmt.GoRaw("}"),
				GoStmt.GoReturn(GoExpr.GoIdent("parsed"))
			]),
			GoDecl.GoFuncDecl("hxrt_unserializerSetField", null, [
				{
					name: "target",
					typeName: "any"
				},
				{name: "fieldName", typeName: "string"},
				{name: "value", typeName: "any"}
			], [], [
				GoStmt.GoRaw("if target == nil {"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("switch obj := target.(type) {"),
				GoStmt.GoRaw("case map[string]any:"),
				GoStmt.GoRaw("\tobj[fieldName] = value"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("case map[any]any:"),
				GoStmt.GoRaw("\tobj[fieldName] = value"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("case *map[string]any:"),
				GoStmt.GoRaw("\tif obj != nil {"),
				GoStmt.GoRaw("\t\t(*obj)[fieldName] = value"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("case *map[any]any:"),
				GoStmt.GoRaw("\tif obj != nil {"),
				GoStmt.GoRaw("\t\t(*obj)[fieldName] = value"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("rv := reflect.ValueOf(target)"),
				GoStmt.GoRaw("if !rv.IsValid() || rv.Kind() != reflect.Pointer || rv.IsNil() {"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("elem := rv.Elem()"),
				GoStmt.GoRaw("if !elem.IsValid() || elem.Kind() != reflect.Struct {"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("field := elem.FieldByName(fieldName)"),
				GoStmt.GoRaw("if !field.IsValid() {"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("targetField := field"),
				GoStmt.GoRaw("if !targetField.CanSet() {"),
				GoStmt.GoRaw("\tif !targetField.CanAddr() {"),
				GoStmt.GoRaw("\t\treturn"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\ttargetField = reflect.NewAt(targetField.Type(), unsafe.Pointer(targetField.UnsafeAddr())).Elem()"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if value == nil {"),
				GoStmt.GoRaw("\ttargetField.Set(reflect.Zero(targetField.Type()))"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("incoming := reflect.ValueOf(value)"),
				GoStmt.GoRaw("if incoming.IsValid() && incoming.Type().AssignableTo(targetField.Type()) {"),
				GoStmt.GoRaw("\ttargetField.Set(incoming)"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if incoming.IsValid() && incoming.Type().ConvertibleTo(targetField.Type()) {"),
				GoStmt.GoRaw("\ttargetField.Set(incoming.Convert(targetField.Type()))"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if targetField.Kind() == reflect.Interface && incoming.IsValid() {"),
				GoStmt.GoRaw("\ttargetField.Set(incoming)"),
				GoStmt.GoRaw("}")
			]),
			GoDecl.GoFuncDecl("hxrt_unserializerReadObjectFields", null, [
				{
					name: "self",
					typeName: "*haxe__Unserializer"
				},
				{name: "target", typeName: "any"},
				{name: "invalidMessage", typeName: "string"}
			], [], [
				GoStmt.GoRaw("raw := *hxrt.StdString(self.buf)"),
				GoStmt.GoRaw("for {"),
				GoStmt.GoRaw("\tif self.pos >= len(raw) {"),
				GoStmt.GoRaw("\t\thxrt.Throw(hxrt.StringFromLiteral(invalidMessage))"),
				GoStmt.GoRaw("\t\treturn"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tif raw[self.pos] == 'g' {"),
				GoStmt.GoRaw("\t\tself.pos++"),
				GoStmt.GoRaw("\t\treturn"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tfieldNameAny := haxe__Unserializer_readValue(self)"),
				GoStmt.GoRaw("\tfieldName := *hxrt.StdString(fieldNameAny)"),
				GoStmt.GoRaw("\tfieldValue := haxe__Unserializer_readValue(self)"),
				GoStmt.GoRaw("\thxrt_unserializerSetField(target, fieldName, fieldValue)"),
				GoStmt.GoRaw("}")
			]),
			GoDecl.GoFuncDecl("haxe__Unserializer_readHexNibble", null, [
				{
					name: "ch",
					typeName: "byte"
				}
			], ["int"], [
				GoStmt.GoRaw("switch {"),
				GoStmt.GoRaw("case ch >= '0' && ch <= '9':"),
				GoStmt.GoRaw("\treturn int(ch - '0')"),
				GoStmt.GoRaw("case ch >= 'A' && ch <= 'F':"),
				GoStmt.GoRaw("\treturn int(ch-'A') + 10"),
				GoStmt.GoRaw("case ch >= 'a' && ch <= 'f':"),
				GoStmt.GoRaw("\treturn int(ch-'a') + 10"),
				GoStmt.GoRaw("default:"),
				GoStmt.GoRaw("\treturn -1"),
				GoStmt.GoRaw("}")
			]),
			GoDecl.GoFuncDecl("haxe__Unserializer_unescape", null, [
				{
					name: "value",
					typeName: "string"
				}
			], ["*string"], [
				GoStmt.GoRaw("out := make([]byte, 0, len(value))"),
				GoStmt.GoRaw("for i := 0; i < len(value); i++ {"),
				GoStmt.GoRaw("\tif value[i] != '%' {"),
				GoStmt.GoRaw("\t\tout = append(out, value[i])"),
				GoStmt.GoRaw("\t\tcontinue"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tif i+2 >= len(value) {"),
				GoStmt.GoRaw("\t\thxrt.Throw(hxrt.StringFromLiteral(\"Invalid serialized string escape\"))"),
				GoStmt.GoRaw("\t\treturn hxrt.StringFromLiteral(\"\")"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\thigh := haxe__Unserializer_readHexNibble(value[i+1])"),
				GoStmt.GoRaw("\tlow := haxe__Unserializer_readHexNibble(value[i+2])"),
				GoStmt.GoRaw("\tif high < 0 || low < 0 {"),
				GoStmt.GoRaw("\t\thxrt.Throw(hxrt.StringFromLiteral(\"Invalid serialized string escape\"))"),
				GoStmt.GoRaw("\t\treturn hxrt.StringFromLiteral(\"\")"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tout = append(out, byte((high<<4)|low))"),
				GoStmt.GoRaw("\ti += 2"),
				GoStmt.GoRaw("}"),
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoRaw("string(out)")]))
			]),
			GoDecl.GoFuncDecl("haxe__Unserializer_readValue", null, [
				{
					name: "self",
					typeName: "*haxe__Unserializer"
				}
			], ["any"], [
				GoStmt.GoRaw("if self == nil || self.buf == nil {"),
				GoStmt.GoRaw("\treturn nil"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("raw := *hxrt.StdString(self.buf)"),
				GoStmt.GoRaw("if self.pos >= len(raw) {"),
				GoStmt.GoRaw("\thxrt.Throw(hxrt.StringFromLiteral(\"Invalid serialized string\"))"),
				GoStmt.GoRaw("\treturn nil"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("token := raw[self.pos]"),
				GoStmt.GoRaw("self.pos++"),
				GoStmt.GoRaw("switch token {"),
				GoStmt.GoRaw("case 'n':"),
				GoStmt.GoRaw("\treturn nil"),
				GoStmt.GoRaw("case 't':"),
				GoStmt.GoRaw("\treturn true"),
				GoStmt.GoRaw("case 'f':"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("case 'z':"),
				GoStmt.GoRaw("\treturn 0"),
				GoStmt.GoRaw("case 'k':"),
				GoStmt.GoRaw("\treturn math.NaN()"),
				GoStmt.GoRaw("case 'p':"),
				GoStmt.GoRaw("\treturn math.Inf(1)"),
				GoStmt.GoRaw("case 'm':"),
				GoStmt.GoRaw("\treturn math.Inf(-1)"),
				GoStmt.GoRaw("case 'i':"),
				GoStmt.GoRaw("\tstart := self.pos"),
				GoStmt.GoRaw("\tif self.pos < len(raw) && (raw[self.pos] == '-' || raw[self.pos] == '+') {"),
				GoStmt.GoRaw("\t\tself.pos++"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tfor self.pos < len(raw) && raw[self.pos] >= '0' && raw[self.pos] <= '9' {"),
				GoStmt.GoRaw("\t\tself.pos++"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tif self.pos == start || (self.pos == start+1 && (raw[start] == '-' || raw[start] == '+')) {"),
				GoStmt.GoRaw("\t\thxrt.Throw(hxrt.StringFromLiteral(\"Invalid serialized integer\"))"),
				GoStmt.GoRaw("\t\treturn 0"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tparsed, err := strconv.Atoi(raw[start:self.pos])"),
				GoStmt.GoRaw("\tif err != nil {"),
				GoStmt.GoRaw("\t\thxrt.Throw(err)"),
				GoStmt.GoRaw("\t\treturn 0"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn parsed"),
				GoStmt.GoRaw("case 'd':"),
				GoStmt.GoRaw("\tstart := self.pos"),
				GoStmt.GoRaw("\thasDigit := false"),
				GoStmt.GoRaw("\tfor self.pos < len(raw) {"),
				GoStmt.GoRaw("\t\tch := raw[self.pos]"),
				GoStmt.GoRaw("\t\tif ch >= '0' && ch <= '9' {"),
				GoStmt.GoRaw("\t\t\thasDigit = true"),
				GoStmt.GoRaw("\t\t\tself.pos++"),
				GoStmt.GoRaw("\t\t\tcontinue"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t\tif ch == '+' || ch == '-' || ch == '.' || ch == 'e' || ch == 'E' {"),
				GoStmt.GoRaw("\t\t\tself.pos++"),
				GoStmt.GoRaw("\t\t\tcontinue"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t\tbreak"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tif !hasDigit {"),
				GoStmt.GoRaw("\t\thxrt.Throw(hxrt.StringFromLiteral(\"Invalid serialized float\"))"),
				GoStmt.GoRaw("\t\treturn 0.0"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tparsed, err := strconv.ParseFloat(raw[start:self.pos], 64)"),
				GoStmt.GoRaw("\tif err != nil {"),
				GoStmt.GoRaw("\t\thxrt.Throw(err)"),
				GoStmt.GoRaw("\t\treturn 0.0"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn parsed"),
				GoStmt.GoRaw("case 'v':"),
				GoStmt.GoRaw("\tstart := self.pos"),
				GoStmt.GoRaw("\thasDigit := false"),
				GoStmt.GoRaw("\tfor self.pos < len(raw) {"),
				GoStmt.GoRaw("\t\tch := raw[self.pos]"),
				GoStmt.GoRaw("\t\tif ch >= '0' && ch <= '9' {"),
				GoStmt.GoRaw("\t\t\thasDigit = true"),
				GoStmt.GoRaw("\t\t\tself.pos++"),
				GoStmt.GoRaw("\t\t\tcontinue"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t\tif ch == '+' || ch == '-' || ch == '.' || ch == 'e' || ch == 'E' {"),
				GoStmt.GoRaw("\t\t\tself.pos++"),
				GoStmt.GoRaw("\t\t\tcontinue"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t\tbreak"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tif !hasDigit {"),
				GoStmt.GoRaw("\t\thxrt.Throw(hxrt.StringFromLiteral(\"Invalid serialized date\"))"),
				GoStmt.GoRaw("\t\treturn &haxe__SerializedDate{ms: 0}"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tparsed, err := strconv.ParseFloat(raw[start:self.pos], 64)"),
				GoStmt.GoRaw("\tif err != nil {"),
				GoStmt.GoRaw("\t\thxrt.Throw(err)"),
				GoStmt.GoRaw("\t\treturn &haxe__SerializedDate{ms: 0}"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn &haxe__SerializedDate{ms: parsed}"),
				GoStmt.GoRaw("case 's':"),
				GoStmt.GoRaw("\tlength := haxe__Unserializer_readUInt(self)"),
				GoStmt.GoRaw("\tif self.pos >= len(raw) || raw[self.pos] != ':' {"),
				GoStmt.GoRaw("\t\thxrt.Throw(hxrt.StringFromLiteral(\"Invalid serialized bytes\"))"),
				GoStmt.GoRaw("\t\treturn &haxe__SerializedBytes{data: []byte{}}"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tself.pos++"),
				GoStmt.GoRaw("\tif length < 0 || self.pos+length > len(raw) {"),
				GoStmt.GoRaw("\t\thxrt.Throw(hxrt.StringFromLiteral(\"Invalid serialized bytes length\"))"),
				GoStmt.GoRaw("\t\treturn &haxe__SerializedBytes{data: []byte{}}"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tencoded := raw[self.pos : self.pos+length]"),
				GoStmt.GoRaw("\tself.pos += length"),
				GoStmt.GoRaw("\tdecoded, err := base64.RawStdEncoding.DecodeString(encoded)"),
				GoStmt.GoRaw("\tif err != nil {"),
				GoStmt.GoRaw("\t\tdecoded, err = base64.StdEncoding.DecodeString(encoded)"),
				GoStmt.GoRaw("\t\tif err != nil {"),
				GoStmt.GoRaw("\t\t\thxrt.Throw(err)"),
				GoStmt.GoRaw("\t\t\treturn &haxe__SerializedBytes{data: []byte{}}"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tout := make([]byte, len(decoded))"),
				GoStmt.GoRaw("\tcopy(out, decoded)"),
				GoStmt.GoRaw("\treturn &haxe__SerializedBytes{data: out}"),
				GoStmt.GoRaw("case 'y':"),
				GoStmt.GoRaw("\tlength := haxe__Unserializer_readUInt(self)"),
				GoStmt.GoRaw("\tif self.pos >= len(raw) || raw[self.pos] != ':' {"),
				GoStmt.GoRaw("\t\thxrt.Throw(hxrt.StringFromLiteral(\"Invalid serialized string\"))"),
				GoStmt.GoRaw("\t\treturn hxrt.StringFromLiteral(\"\")"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tself.pos++"),
				GoStmt.GoRaw("\tif length < 0 || self.pos+length > len(raw) {"),
				GoStmt.GoRaw("\t\thxrt.Throw(hxrt.StringFromLiteral(\"Invalid serialized string length\"))"),
				GoStmt.GoRaw("\t\treturn hxrt.StringFromLiteral(\"\")"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tdecoded := haxe__Unserializer_unescape(raw[self.pos : self.pos+length])"),
				GoStmt.GoRaw("\tself.pos += length"),
				GoStmt.GoRaw("\tself.stringCache = append(self.stringCache, decoded)"),
				GoStmt.GoRaw("\treturn decoded"),
				GoStmt.GoRaw("case 'R':"),
				GoStmt.GoRaw("\tindex := haxe__Unserializer_readUInt(self)"),
				GoStmt.GoRaw("\tif index < 0 || index >= len(self.stringCache) {"),
				GoStmt.GoRaw("\t\thxrt.Throw(hxrt.StringFromLiteral(\"Invalid string reference\"))"),
				GoStmt.GoRaw("\t\treturn hxrt.StringFromLiteral(\"\")"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn self.stringCache[index]"),
				GoStmt.GoRaw("case 'x':"),
				GoStmt.GoRaw("\thxrt.Throw(haxe__Unserializer_readValue(self))"),
				GoStmt.GoRaw("\treturn nil"),
				GoStmt.GoRaw("case 'l':"),
				GoStmt.GoRaw("\tlist := New_haxe__ds__List()"),
				GoStmt.GoRaw("\tcacheIndex := len(self.cache)"),
				GoStmt.GoRaw("\tself.cache = append(self.cache, list)"),
				GoStmt.GoRaw("\tfor {"),
				GoStmt.GoRaw("\t\tif self.pos >= len(raw) {"),
				GoStmt.GoRaw("\t\t\thxrt.Throw(hxrt.StringFromLiteral(\"Invalid serialized list\"))"),
				GoStmt.GoRaw("\t\t\treturn list"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t\tif raw[self.pos] == 'h' {"),
				GoStmt.GoRaw("\t\t\tself.pos++"),
				GoStmt.GoRaw("\t\t\tbreak"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t\tlist.items = append(list.items, haxe__Unserializer_readValue(self))"),
				GoStmt.GoRaw("\t\tlist.length = len(list.items)"),
				GoStmt.GoRaw("\t\tself.cache[cacheIndex] = list"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn list"),
				GoStmt.GoRaw("case 'b':"),
				GoStmt.GoRaw("\tstringMap := New_haxe__ds__StringMap()"),
				GoStmt.GoRaw("\tcacheIndex := len(self.cache)"),
				GoStmt.GoRaw("\tself.cache = append(self.cache, stringMap)"),
				GoStmt.GoRaw("\tfor {"),
				GoStmt.GoRaw("\t\tif self.pos >= len(raw) {"),
				GoStmt.GoRaw("\t\t\thxrt.Throw(hxrt.StringFromLiteral(\"Invalid serialized StringMap\"))"),
				GoStmt.GoRaw("\t\t\treturn stringMap"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t\tif raw[self.pos] == 'h' {"),
				GoStmt.GoRaw("\t\t\tself.pos++"),
				GoStmt.GoRaw("\t\t\tbreak"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t\tkeyAny := haxe__Unserializer_readValue(self)"),
				GoStmt.GoRaw("\t\tkey := *hxrt.StdString(keyAny)"),
				GoStmt.GoRaw("\t\tstringMap.h[key] = haxe__Unserializer_readValue(self)"),
				GoStmt.GoRaw("\t\tself.cache[cacheIndex] = stringMap"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn stringMap"),
				GoStmt.GoRaw("case 'q':"),
				GoStmt.GoRaw("\tintMap := New_haxe__ds__IntMap()"),
				GoStmt.GoRaw("\tcacheIndex := len(self.cache)"),
				GoStmt.GoRaw("\tself.cache = append(self.cache, intMap)"),
				GoStmt.GoRaw("\tfor {"),
				GoStmt.GoRaw("\t\tif self.pos >= len(raw) {"),
				GoStmt.GoRaw("\t\t\thxrt.Throw(hxrt.StringFromLiteral(\"Invalid serialized IntMap\"))"),
				GoStmt.GoRaw("\t\t\treturn intMap"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t\tif raw[self.pos] == 'h' {"),
				GoStmt.GoRaw("\t\t\tself.pos++"),
				GoStmt.GoRaw("\t\t\tbreak"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t\tif raw[self.pos] != ':' {"),
				GoStmt.GoRaw("\t\t\thxrt.Throw(hxrt.StringFromLiteral(\"Invalid serialized IntMap format\"))"),
				GoStmt.GoRaw("\t\t\treturn intMap"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t\tself.pos++"),
				GoStmt.GoRaw("\t\tkey := haxe__Unserializer_readDigits(self)"),
				GoStmt.GoRaw("\t\tintMap.h[key] = haxe__Unserializer_readValue(self)"),
				GoStmt.GoRaw("\t\tself.cache[cacheIndex] = intMap"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn intMap"),
				GoStmt.GoRaw("case 'M':"),
				GoStmt.GoRaw("\tobjectMap := New_haxe__ds__ObjectMap()"),
				GoStmt.GoRaw("\tcacheIndex := len(self.cache)"),
				GoStmt.GoRaw("\tself.cache = append(self.cache, objectMap)"),
				GoStmt.GoRaw("\tfor {"),
				GoStmt.GoRaw("\t\tif self.pos >= len(raw) {"),
				GoStmt.GoRaw("\t\t\thxrt.Throw(hxrt.StringFromLiteral(\"Invalid serialized ObjectMap\"))"),
				GoStmt.GoRaw("\t\t\treturn objectMap"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t\tif raw[self.pos] == 'h' {"),
				GoStmt.GoRaw("\t\t\tself.pos++"),
				GoStmt.GoRaw("\t\t\tbreak"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t\tkey := haxe__Unserializer_readValue(self)"),
				GoStmt.GoRaw("\t\tobjectMap.h[key] = haxe__Unserializer_readValue(self)"),
				GoStmt.GoRaw("\t\tself.cache[cacheIndex] = objectMap"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn objectMap"),
				GoStmt.GoRaw("case 'a':"),
				GoStmt.GoRaw("\tarr := make([]any, 0)"),
				GoStmt.GoRaw("\tcacheIndex := len(self.cache)"),
				GoStmt.GoRaw("\tself.cache = append(self.cache, arr)"),
				GoStmt.GoRaw("\tfor {"),
				GoStmt.GoRaw("\t\tif self.pos >= len(raw) {"),
				GoStmt.GoRaw("\t\t\thxrt.Throw(hxrt.StringFromLiteral(\"Invalid serialized array\"))"),
				GoStmt.GoRaw("\t\t\treturn arr"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t\tif raw[self.pos] == 'h' {"),
				GoStmt.GoRaw("\t\t\tself.pos++"),
				GoStmt.GoRaw("\t\t\tbreak"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t\tif raw[self.pos] == 'u' {"),
				GoStmt.GoRaw("\t\t\tself.pos++"),
				GoStmt.GoRaw("\t\t\tskip := haxe__Unserializer_readUInt(self)"),
				GoStmt.GoRaw("\t\t\tfor i := 0; i < skip; i++ {"),
				GoStmt.GoRaw("\t\t\t\tarr = append(arr, nil)"),
				GoStmt.GoRaw("\t\t\t}"),
				GoStmt.GoRaw("\t\t\tself.cache[cacheIndex] = arr"),
				GoStmt.GoRaw("\t\t\tcontinue"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t\tarr = append(arr, haxe__Unserializer_readValue(self))"),
				GoStmt.GoRaw("\t\tself.cache[cacheIndex] = arr"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn arr"),
				GoStmt.GoRaw("case 'o':"),
				GoStmt.GoRaw("\tobj := map[string]any{}"),
				GoStmt.GoRaw("\tcacheIndex := len(self.cache)"),
				GoStmt.GoRaw("\tself.cache = append(self.cache, obj)"),
				GoStmt.GoRaw("\thxrt_unserializerReadObjectFields(self, obj, \"Invalid serialized object\")"),
				GoStmt.GoRaw("\tself.cache[cacheIndex] = obj"),
				GoStmt.GoRaw("\treturn obj"),
				GoStmt.GoRaw("case 'C':"),
				GoStmt.GoRaw("\tclassNameAny := haxe__Unserializer_readValue(self)"),
				GoStmt.GoRaw("\trequestedName := *hxrt.StdString(classNameAny)"),
				GoStmt.GoRaw("\tresolvedClass := hxrt_unserializerResolveClass(self, requestedName)"),
				GoStmt.GoRaw("\tif resolvedClass == nil {"),
				GoStmt.GoRaw("\t\thxrt.Throw(hxrt.StringFromLiteral(\"Class not found \" + requestedName))"),
				GoStmt.GoRaw("\t\treturn nil"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tclassName, ok := hxrt_unserializerResolvedClassName(resolvedClass)"),
				GoStmt.GoRaw("\tif !ok || className == \"\" {"),
				GoStmt.GoRaw("\t\thxrt.Throw(hxrt.StringFromLiteral(\"Class not found \" + requestedName))"),
				GoStmt.GoRaw("\t\treturn nil"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tinstance, ok := hxrt_unserializerCreateClassInstance(className)"),
				GoStmt.GoRaw("\tif !ok {"),
				GoStmt.GoRaw("\t\thxrt.Throw(hxrt.StringFromLiteral(\"Class not found \" + className))"),
				GoStmt.GoRaw("\t\treturn nil"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tcacheIndex := len(self.cache)"),
				GoStmt.GoRaw("\tself.cache = append(self.cache, instance)"),
				GoStmt.GoRaw("\tif custom, ok := instance.(interface{ hxUnserialize(*haxe__Unserializer) }); ok {"),
				GoStmt.GoRaw("\t\tcustom.hxUnserialize(self)"),
				GoStmt.GoRaw("\t} else if custom, ok := instance.(interface{ HxUnserialize(*haxe__Unserializer) }); ok {"),
				GoStmt.GoRaw("\t\tcustom.HxUnserialize(self)"),
				GoStmt.GoRaw("\t} else {"),
				GoStmt.GoRaw("\t\thxrt.Throw(hxrt.StringFromLiteral(\"Invalid custom data\"))"),
				GoStmt.GoRaw("\t\treturn instance"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tif self.pos >= len(raw) || raw[self.pos] != 'g' {"),
				GoStmt.GoRaw("\t\thxrt.Throw(hxrt.StringFromLiteral(\"Invalid custom data\"))"),
				GoStmt.GoRaw("\t\treturn instance"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tself.pos++"),
				GoStmt.GoRaw("\tself.cache[cacheIndex] = instance"),
				GoStmt.GoRaw("\treturn instance"),
				GoStmt.GoRaw("case 'A':"),
				GoStmt.GoRaw("\tnameAny := haxe__Unserializer_readValue(self)"),
				GoStmt.GoRaw("\trequestedName := *hxrt.StdString(nameAny)"),
				GoStmt.GoRaw("\tresolvedClass := hxrt_unserializerResolveClass(self, requestedName)"),
				GoStmt.GoRaw("\tif resolvedClass == nil {"),
				GoStmt.GoRaw("\t\thxrt.Throw(hxrt.StringFromLiteral(\"Class not found \" + requestedName))"),
				GoStmt.GoRaw("\t\treturn nil"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn resolvedClass"),
				GoStmt.GoRaw("case 'B':"),
				GoStmt.GoRaw("\tnameAny := haxe__Unserializer_readValue(self)"),
				GoStmt.GoRaw("\trequestedName := *hxrt.StdString(nameAny)"),
				GoStmt.GoRaw("\tresolvedEnum := hxrt_unserializerResolveEnum(self, requestedName)"),
				GoStmt.GoRaw("\tif resolvedEnum == nil {"),
				GoStmt.GoRaw("\t\thxrt.Throw(hxrt.StringFromLiteral(\"Enum not found \" + requestedName))"),
				GoStmt.GoRaw("\t\treturn nil"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn resolvedEnum"),
				GoStmt.GoRaw("case 'c':"),
				GoStmt.GoRaw("\tclassNameAny := haxe__Unserializer_readValue(self)"),
				GoStmt.GoRaw("\trequestedName := *hxrt.StdString(classNameAny)"),
				GoStmt.GoRaw("\tresolvedClass := hxrt_unserializerResolveClass(self, requestedName)"),
				GoStmt.GoRaw("\tif resolvedClass == nil {"),
				GoStmt.GoRaw("\t\thxrt.Throw(hxrt.StringFromLiteral(\"Class not found \" + requestedName))"),
				GoStmt.GoRaw("\t\treturn nil"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tclassName, ok := hxrt_unserializerResolvedClassName(resolvedClass)"),
				GoStmt.GoRaw("\tif !ok || className == \"\" {"),
				GoStmt.GoRaw("\t\thxrt.Throw(hxrt.StringFromLiteral(\"Class not found \" + requestedName))"),
				GoStmt.GoRaw("\t\treturn nil"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tinstance, ok := hxrt_unserializerCreateClassInstance(className)"),
				GoStmt.GoRaw("\tif !ok {"),
				GoStmt.GoRaw("\t\thxrt.Throw(hxrt.StringFromLiteral(\"Class not found \" + className))"),
				GoStmt.GoRaw("\t\treturn nil"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tcacheIndex := len(self.cache)"),
				GoStmt.GoRaw("\tself.cache = append(self.cache, instance)"),
				GoStmt.GoRaw("\thxrt_unserializerReadObjectFields(self, instance, \"Invalid serialized class\")"),
				GoStmt.GoRaw("\tself.cache[cacheIndex] = instance"),
				GoStmt.GoRaw("\treturn instance"),
				GoStmt.GoRaw("case 'j':"),
				GoStmt.GoRaw("\tenumNameAny := haxe__Unserializer_readValue(self)"),
				GoStmt.GoRaw("\trequestedEnumName := *hxrt.StdString(enumNameAny)"),
				GoStmt.GoRaw("\tresolvedEnum := hxrt_unserializerResolveEnum(self, requestedEnumName)"),
				GoStmt.GoRaw("\tif resolvedEnum == nil {"),
				GoStmt.GoRaw("\t\thxrt.Throw(hxrt.StringFromLiteral(\"Enum not found \" + requestedEnumName))"),
				GoStmt.GoRaw("\t\treturn nil"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tenumName, ok := hxrt_unserializerResolvedEnumName(resolvedEnum)"),
				GoStmt.GoRaw("\tif !ok || enumName == \"\" {"),
				GoStmt.GoRaw("\t\thxrt.Throw(hxrt.StringFromLiteral(\"Enum not found \" + requestedEnumName))"),
				GoStmt.GoRaw("\t\treturn nil"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tif self.pos >= len(raw) || raw[self.pos] != ':' {"),
				GoStmt.GoRaw("\t\thxrt.Throw(hxrt.StringFromLiteral(\"Invalid serialized enum index\"))"),
				GoStmt.GoRaw("\t\treturn nil"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tself.pos++"),
				GoStmt.GoRaw("\tenumIndex := haxe__Unserializer_readDigits(self)"),
				GoStmt.GoRaw("\tif self.pos >= len(raw) || raw[self.pos] != ':' {"),
				GoStmt.GoRaw("\t\thxrt.Throw(hxrt.StringFromLiteral(\"Invalid serialized enum\"))"),
				GoStmt.GoRaw("\t\treturn nil"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tself.pos++"),
				GoStmt.GoRaw("\targCount := haxe__Unserializer_readUInt(self)"),
				GoStmt.GoRaw("\tif argCount < 0 {"),
				GoStmt.GoRaw("\t\thxrt.Throw(hxrt.StringFromLiteral(\"Invalid serialized enum arity\"))"),
				GoStmt.GoRaw("\t\treturn nil"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\targs := make([]any, 0, argCount)"),
				GoStmt.GoRaw("\tfor i := 0; i < argCount; i++ {"),
				GoStmt.GoRaw("\t\targs = append(args, haxe__Unserializer_readValue(self))"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tenumValue, ok := hxrt_unserializerCreateEnumInstance(enumName, \"\", enumIndex, true, args)"),
				GoStmt.GoRaw("\tif !ok {"),
				GoStmt.GoRaw("\t\thxrt.Throw(hxrt.StringFromLiteral(\"Unknown enum index \" + enumName + \"@\" + strconv.Itoa(enumIndex)))"),
				GoStmt.GoRaw("\t\treturn nil"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tself.cache = append(self.cache, enumValue)"),
				GoStmt.GoRaw("\treturn enumValue"),
				GoStmt.GoRaw("case 'w':"),
				GoStmt.GoRaw("\tenumNameAny := haxe__Unserializer_readValue(self)"),
				GoStmt.GoRaw("\tconstructorAny := haxe__Unserializer_readValue(self)"),
				GoStmt.GoRaw("\trequestedEnumName := *hxrt.StdString(enumNameAny)"),
				GoStmt.GoRaw("\tresolvedEnum := hxrt_unserializerResolveEnum(self, requestedEnumName)"),
				GoStmt.GoRaw("\tif resolvedEnum == nil {"),
				GoStmt.GoRaw("\t\thxrt.Throw(hxrt.StringFromLiteral(\"Enum not found \" + requestedEnumName))"),
				GoStmt.GoRaw("\t\treturn nil"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tenumName, ok := hxrt_unserializerResolvedEnumName(resolvedEnum)"),
				GoStmt.GoRaw("\tif !ok || enumName == \"\" {"),
				GoStmt.GoRaw("\t\thxrt.Throw(hxrt.StringFromLiteral(\"Enum not found \" + requestedEnumName))"),
				GoStmt.GoRaw("\t\treturn nil"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tconstructorName := *hxrt.StdString(constructorAny)"),
				GoStmt.GoRaw("\tif self.pos >= len(raw) || raw[self.pos] != ':' {"),
				GoStmt.GoRaw("\t\thxrt.Throw(hxrt.StringFromLiteral(\"Invalid serialized enum\"))"),
				GoStmt.GoRaw("\t\treturn nil"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tself.pos++"),
				GoStmt.GoRaw("\targCount := haxe__Unserializer_readUInt(self)"),
				GoStmt.GoRaw("\tif argCount < 0 {"),
				GoStmt.GoRaw("\t\thxrt.Throw(hxrt.StringFromLiteral(\"Invalid serialized enum arity\"))"),
				GoStmt.GoRaw("\t\treturn nil"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\targs := make([]any, 0, argCount)"),
				GoStmt.GoRaw("\tfor i := 0; i < argCount; i++ {"),
				GoStmt.GoRaw("\t\targs = append(args, haxe__Unserializer_readValue(self))"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tenumValue, ok := hxrt_unserializerCreateEnumInstance(enumName, constructorName, 0, false, args)"),
				GoStmt.GoRaw("\tif !ok {"),
				GoStmt.GoRaw("\t\thxrt.Throw(hxrt.StringFromLiteral(\"Invalid serialized enum\"))"),
				GoStmt.GoRaw("\t\treturn nil"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tself.cache = append(self.cache, enumValue)"),
				GoStmt.GoRaw("\treturn enumValue"),
				GoStmt.GoRaw("case 'r':"),
				GoStmt.GoRaw("\tindex := haxe__Unserializer_readUInt(self)"),
				GoStmt.GoRaw("\tif index < 0 || index >= len(self.cache) {"),
				GoStmt.GoRaw("\t\thxrt.Throw(hxrt.StringFromLiteral(\"Invalid object reference\"))"),
				GoStmt.GoRaw("\t\treturn nil"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn self.cache[index]"),
				GoStmt.GoRaw("default:"),
				GoStmt.GoRaw("\thxrt.Throw(hxrt.StringFromLiteral(\"Invalid serialized token\"))"),
				GoStmt.GoRaw("\treturn nil"),
				GoStmt.GoRaw("}")
			]),
			GoDecl.GoFuncDecl("haxe__Unserializer_run", null, [
				{
					name: "source",
					typeName: "*string"
				}
			], ["any"], [
				GoStmt.GoRaw("if source == nil {"),
				GoStmt.GoRaw("\treturn nil"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("decoder := New_haxe__Unserializer(source)"),
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("decoder"), "unserialize"), []))
			])
		];
	}

	function lowerNetSocketShimDecls():Array<GoDecl> {
		return [
			GoDecl.GoStructDecl("sys__net__Host", [
				{name: "host", typeName: "*string"},
				{name: "ip", typeName: "int"},
				{
					name: "resolved",
					typeName: "*string"
				}
			]),
			GoDecl.GoFuncDecl("hxrt__host_empty", null, [], ["*sys__net__Host"], [
				GoStmt.GoReturn(GoExpr.GoRaw("&sys__net__Host{host: hxrt.StringFromLiteral(\"\"), ip: 0, resolved: hxrt.StringFromLiteral(\"\")}"))
			]),
			GoDecl.GoFuncDecl("New_sys__net__Host", null, [
				{
					name: "name",
					typeName: "*string"
				}
			], ["*sys__net__Host"], [
				GoStmt.GoRaw("if name == nil {"),
				GoStmt.GoRaw("\thxrt.Throw(hxrt.StringFromLiteral(\"Could not resolve host\"))"),
				GoStmt.GoRaw("\treturn hxrt__host_empty()"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("rawName := *hxrt.StdString(name)"),
				GoStmt.GoRaw("ips, err := net.LookupIP(rawName)"),
				GoStmt.GoRaw("if err != nil || len(ips) == 0 {"),
				GoStmt.GoRaw("\thxrt.Throw(hxrt.StringFromLiteral(\"Could not resolve host\"))"),
				GoStmt.GoRaw("\treturn hxrt__host_empty()"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("selected := ips[0]"),
				GoStmt.GoRaw("for _, candidate := range ips {"),
				GoStmt.GoRaw("\tif v4 := candidate.To4(); v4 != nil {"),
				GoStmt.GoRaw("\t\tselected = v4"),
				GoStmt.GoRaw("\t\tbreak"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("resolved := hxrt.StringFromLiteral(selected.String())"),
				GoStmt.GoReturn(GoExpr.GoRaw("&sys__net__Host{host: name, ip: 0, resolved: resolved}"))
			]),
			GoDecl.GoFuncDecl("toString", {
				name: "self",
				typeName: "*sys__net__Host"
			}, [], ["*string"], [
				GoStmt.GoIf(GoExpr.GoRaw("self == nil || self.resolved == nil"), [
					GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoStringLiteral("")]))
				], null),
				GoStmt.GoReturn(GoExpr.GoSelector(GoExpr.GoIdent("self"), "resolved"))
			]),
			GoDecl.GoFuncDecl("reverse", {
				name: "self",
				typeName: "*sys__net__Host"
			}, [], ["*string"], [
				GoStmt.GoRaw("if self == nil || self.resolved == nil {"),
				GoStmt.GoRaw("\thxrt.Throw(hxrt.StringFromLiteral(\"Could not reverse host\"))"),
				GoStmt.GoRaw("\treturn hxrt.StringFromLiteral(\"\")"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("names, err := net.LookupAddr(*hxrt.StdString(self.resolved))"),
				GoStmt.GoRaw("if err != nil || len(names) == 0 {"),
				GoStmt.GoRaw("\thxrt.Throw(hxrt.StringFromLiteral(\"Could not reverse host\"))"),
				GoStmt.GoRaw("\treturn hxrt.StringFromLiteral(\"\")"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("resolved := strings.TrimSuffix(names[0], \".\")"),
				GoStmt.GoReturn(GoExpr.GoRaw("hxrt.StringFromLiteral(resolved)"))
			]),
			GoDecl.GoFuncDecl("sys__net__Host_localhost", null, [], ["*string"], [
				GoStmt.GoRaw("name, err := os.Hostname()"),
				GoStmt.GoRaw("if err != nil || name == \"\" {"),
				GoStmt.GoRaw("\treturn hxrt.StringFromLiteral(\"localhost\")"),
				GoStmt.GoRaw("}"),
				GoStmt.GoReturn(GoExpr.GoRaw("hxrt.StringFromLiteral(name)"))
			]),
			GoDecl.GoStructDecl("sys__net__SocketInput", [
				{
					name: "reader",
					typeName: "*bufio.Reader"
				},
				{name: "socket", typeName: "*sys__net__Socket"}
			]),
			GoDecl.GoStructDecl("sys__net__SocketOutput", [
				{name: "writer", typeName: "*bufio.Writer"},
				{name: "socket", typeName: "*sys__net__Socket"}
			]),
			GoDecl.GoStructDecl("sys__net__Socket", [
				{name: "input", typeName: "*sys__net__SocketInput"},
				{name: "output", typeName: "*sys__net__SocketOutput"},
				{name: "custom", typeName: "any"},
				{name: "conn", typeName: "net.Conn"},
				{name: "listener", typeName: "net.Listener"},
				{name: "timeout", typeName: "float64"},
				{name: "hasTimeout", typeName: "bool"},
				{name: "blocking", typeName: "bool"},
				{name: "fastSend", typeName: "bool"}
			]),
			GoDecl.GoFuncDecl("New_sys__net__Socket", null, [], ["*sys__net__Socket"], [
				GoStmt.GoReturn(GoExpr.GoRaw("&sys__net__Socket{input: &sys__net__SocketInput{}, output: &sys__net__SocketOutput{}, blocking: true}"))
			]),
			GoDecl.GoFuncDecl("hxrt__socket_deadline", null, [
				{
					name: "timeout",
					typeName: "float64"
				}
			], ["time.Time"], [
				GoStmt.GoRaw("duration := time.Duration(timeout * float64(time.Second))"),
				GoStmt.GoRaw("return time.Now().Add(duration)")
			]),
			GoDecl.GoFuncDecl("hxrt__socket_applyConnDeadline", {
				name: "self",
				typeName: "*sys__net__Socket"
			}, [], [], [
				GoStmt.GoRaw("if self == nil || self.conn == nil {"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if !self.blocking {"),
				GoStmt.GoRaw("\t_ = self.conn.SetDeadline(time.Now())"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if self.hasTimeout {"),
				GoStmt.GoRaw("\t_ = self.conn.SetDeadline(hxrt__socket_deadline(self.timeout))"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("_ = self.conn.SetDeadline(time.Time{})")
			]),
			GoDecl.GoFuncDecl("hxrt__socket_applyListenerDeadline", {
				name: "self",
				typeName: "*sys__net__Socket"
			}, [], [], [
				GoStmt.GoRaw("if self == nil || self.listener == nil {"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("tcpListener, ok := self.listener.(*net.TCPListener)"),
				GoStmt.GoRaw("if !ok {"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if !self.blocking {"),
				GoStmt.GoRaw("\t_ = tcpListener.SetDeadline(time.Now())"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if self.hasTimeout {"),
				GoStmt.GoRaw("\t_ = tcpListener.SetDeadline(hxrt__socket_deadline(self.timeout))"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("_ = tcpListener.SetDeadline(time.Time{})")
			]),
			GoDecl.GoFuncDecl("hxrt__socket_applyFastSend", {
				name: "self",
				typeName: "*sys__net__Socket"
			}, [], [], [
				GoStmt.GoRaw("if self == nil || self.conn == nil {"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("tcpConn, ok := self.conn.(*net.TCPConn)"),
				GoStmt.GoRaw("if !ok {"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if err := tcpConn.SetNoDelay(self.fastSend); err != nil {"),
				GoStmt.GoRaw("\thxrt.Throw(err)"),
				GoStmt.GoRaw("}")
			]),
			GoDecl.GoFuncDecl("hxrt__socket_setConn", {
				name: "self",
				typeName: "*sys__net__Socket"
			}, [{name: "conn", typeName: "net.Conn"}], [], [
				GoStmt.GoIf(GoExpr.GoRaw("self == nil || conn == nil"), [GoStmt.GoReturn(null)], null),
				GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("self"), "conn"), GoExpr.GoIdent("conn")),
				GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("self"), "input"),
					GoExpr.GoRaw("&sys__net__SocketInput{reader: bufio.NewReader(conn), socket: self}")),
				GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("self"), "output"),
					GoExpr.GoRaw("&sys__net__SocketOutput{writer: bufio.NewWriter(conn), socket: self}")),
				GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("self"), "hxrt__socket_applyFastSend"), [])),
				GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("self"), "hxrt__socket_applyConnDeadline"), []))
			]),
			GoDecl.GoFuncDecl("hxrt__socket_conn", {
				name: "self",
				typeName: "*sys__net__Socket"
			}, [], ["net.Conn"], [
				GoStmt.GoIf(GoExpr.GoRaw("self == nil"), [GoStmt.GoReturn(GoExpr.GoNil)], null),
				GoStmt.GoReturn(GoExpr.GoSelector(GoExpr.GoIdent("self"), "conn"))
			]),
			GoDecl.GoFuncDecl("close", {
				name: "self",
				typeName: "*sys__net__Socket"
			}, [], [], [
				GoStmt.GoIf(GoExpr.GoRaw("self == nil"), [GoStmt.GoReturn(null)], null),
				GoStmt.GoIf(GoExpr.GoRaw("self.conn != nil"), [
					GoStmt.GoRaw("_ = self.conn.Close()"),
					GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("self"), "conn"), GoExpr.GoNil)
				], null),
				GoStmt.GoIf(GoExpr.GoRaw("self.listener != nil"), [
					GoStmt.GoRaw("_ = self.listener.Close()"),
					GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("self"), "listener"), GoExpr.GoNil)
				], null)
			]),
			GoDecl.GoFuncDecl("connect", {
				name: "self",
				typeName: "*sys__net__Socket"
			}, [
				{
					name: "host",
					typeName: "*sys__net__Host"
				},
				{name: "port", typeName: "int"}
			], [], [
				GoStmt.GoIf(GoExpr.GoRaw("self == nil || host == nil"), [
					GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("hxrt.Throw"),
						[
							GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoStringLiteral("socket connect requires host")])
						])),
					GoStmt.GoReturn(null)
				],
					null),
				GoStmt.GoVarDecl("resolvedHost", "*string", GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("host"), "toString"), []), true),
				GoStmt.GoIf(GoExpr.GoRaw("resolvedHost == nil"), [
					GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("hxrt.Throw"), [
						GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoStringLiteral("socket connect requires host")])
					])),
					GoStmt.GoReturn(null)
				], null),
				GoStmt.GoVarDecl("address", null, GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("net"), "JoinHostPort"), [
					GoExpr.GoRaw("*hxrt.StdString(resolvedHost)"),
					GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("strconv"), "Itoa"), [GoExpr.GoIdent("port")])
				]), true),
				GoStmt.GoRaw("conn, err := net.Dial(\"tcp\", address)"),
				GoStmt.GoIf(GoExpr.GoBinary("!=", GoExpr.GoIdent("err"), GoExpr.GoNil), [
					GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("hxrt.Throw"), [GoExpr.GoIdent("err")])),
					GoStmt.GoReturn(null)
				],
					null),
				GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("self"), "hxrt__socket_setConn"), [GoExpr.GoIdent("conn")]))
			]),
			GoDecl.GoFuncDecl("bind", {
				name: "self",
				typeName: "*sys__net__Socket"
			}, [
				{
					name: "host",
					typeName: "*sys__net__Host"
				},
				{name: "port", typeName: "int"}
			], [], [
				GoStmt.GoIf(GoExpr.GoRaw("self == nil || host == nil"), [
					GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("hxrt.Throw"),
						[
							GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoStringLiteral("socket bind requires host")])
						])),
					GoStmt.GoReturn(null)
				],
					null),
				GoStmt.GoVarDecl("resolvedHost", "*string", GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("host"), "toString"), []), true),
				GoStmt.GoIf(GoExpr.GoRaw("resolvedHost == nil"), [
					GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("hxrt.Throw"), [
						GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoStringLiteral("socket bind requires host")])
					])),
					GoStmt.GoReturn(null)
				], null),
				GoStmt.GoVarDecl("address", null, GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("net"), "JoinHostPort"), [
					GoExpr.GoRaw("*hxrt.StdString(resolvedHost)"),
					GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("strconv"), "Itoa"), [GoExpr.GoIdent("port")])
				]), true),
				GoStmt.GoRaw("listener, err := net.Listen(\"tcp\", address)"),
				GoStmt.GoIf(GoExpr.GoBinary("!=", GoExpr.GoIdent("err"), GoExpr.GoNil), [
					GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("hxrt.Throw"), [GoExpr.GoIdent("err")])),
					GoStmt.GoReturn(null)
				],
					null),
				GoStmt.GoIf(GoExpr.GoRaw("self.listener != nil"), [GoStmt.GoRaw("_ = self.listener.Close()")], null),
				GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("self"), "listener"), GoExpr.GoIdent("listener")),
				GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("self"), "hxrt__socket_applyListenerDeadline"), []))
			]),
			GoDecl.GoFuncDecl("listen", {
				name: "self",
				typeName: "*sys__net__Socket"
			}, [{name: "connections", typeName: "int"}], [],
				[GoStmt.GoRaw("_ = connections")]),
			GoDecl.GoFuncDecl("accept", {
				name: "self",
				typeName: "*sys__net__Socket"
			}, [], ["*sys__net__Socket"], [
				GoStmt.GoIf(GoExpr.GoRaw("self == nil || self.listener == nil"), [
					GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("hxrt.Throw"),
						[
							GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoStringLiteral("socket accept requires listener")])
						])),
					GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("New_sys__net__Socket"), []))
				],
					null),
				GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("self"), "hxrt__socket_applyListenerDeadline"), [])),
				GoStmt.GoRaw("conn, err := self.listener.Accept()"),
				GoStmt.GoIf(GoExpr.GoBinary("!=", GoExpr.GoIdent("err"), GoExpr.GoNil), [
					GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("hxrt.Throw"), [GoExpr.GoIdent("err")])),
					GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("New_sys__net__Socket"), []))
				],
					null),
				GoStmt.GoVarDecl("accepted", null, GoExpr.GoCall(GoExpr.GoIdent("New_sys__net__Socket"), []), true),
				GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("accepted"), "timeout"), GoExpr.GoSelector(GoExpr.GoIdent("self"), "timeout")),
				GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("accepted"), "hasTimeout"), GoExpr.GoSelector(GoExpr.GoIdent("self"), "hasTimeout")),
				GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("accepted"), "blocking"), GoExpr.GoSelector(GoExpr.GoIdent("self"), "blocking")),
				GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("accepted"), "fastSend"), GoExpr.GoSelector(GoExpr.GoIdent("self"), "fastSend")),
				GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("accepted"), "hxrt__socket_setConn"), [GoExpr.GoIdent("conn")])),
				GoStmt.GoReturn(GoExpr.GoIdent("accepted"))
			]),
			GoDecl.GoFuncDecl("read", {
				name: "self",
				typeName: "*sys__net__Socket"
			}, [], ["*string"], [
				GoStmt.GoIf(GoExpr.GoRaw("self == nil || self.input == nil"), [
					GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoStringLiteral("")]))
				],
					null),
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoSelector(GoExpr.GoIdent("self"), "input"), "readLine"), []))
			]),
			GoDecl.GoFuncDecl("write", {
				name: "self",
				typeName: "*sys__net__Socket"
			}, [{name: "content", typeName: "*string"}], [], [
				GoStmt.GoIf(GoExpr.GoRaw("self == nil || self.output == nil"), [GoStmt.GoReturn(null)], null),
				GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoSelector(GoExpr.GoIdent("self"), "output"), "writeString"),
					[GoExpr.GoIdent("content")])),
				GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoSelector(GoExpr.GoIdent("self"), "output"), "flush"), []))
			]),
			GoDecl.GoFuncDecl("shutdown", {
				name: "self",
				typeName: "*sys__net__Socket"
			}, [
				{
					name: "read",
					typeName: "bool"
				},
				{name: "write", typeName: "bool"}
			], [], [
				GoStmt.GoRaw("if self == nil || self.conn == nil || (!read && !write) {"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if tcpConn, ok := self.conn.(*net.TCPConn); ok {"),
				GoStmt.GoRaw("\tif read {"),
				GoStmt.GoRaw("\t\tif err := tcpConn.CloseRead(); err != nil {"),
				GoStmt.GoRaw("\t\t\thxrt.Throw(err)"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tif write {"),
				GoStmt.GoRaw("\t\tif err := tcpConn.CloseWrite(); err != nil {"),
				GoStmt.GoRaw("\t\t\thxrt.Throw(err)"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if err := self.conn.Close(); err != nil {"),
				GoStmt.GoRaw("\thxrt.Throw(err)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("self.conn = nil")
			]),
			GoDecl.GoFuncDecl("hxrt__socket_addrInfo", null, [
				{
					name: "addr",
					typeName: "net.Addr"
				}
			], ["map[string]any"], [
				GoStmt.GoRaw("if addr == nil {"),
				GoStmt.GoRaw("\treturn map[string]any{\"host\": hxrt__host_empty(), \"port\": 0}"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("rawHost := \"\""),
				GoStmt.GoRaw("rawPort := \"0\""),
				GoStmt.GoRaw("hostPart, portPart, err := net.SplitHostPort(addr.String())"),
				GoStmt.GoRaw("if err == nil {"),
				GoStmt.GoRaw("\trawHost = hostPart"),
				GoStmt.GoRaw("\trawPort = portPart"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("port, _ := strconv.Atoi(rawPort)"),
				GoStmt.GoRaw("if rawHost == \"\" {"),
				GoStmt.GoRaw("\treturn map[string]any{\"host\": hxrt__host_empty(), \"port\": port}"),
				GoStmt.GoRaw("}"),
				GoStmt.GoReturn(GoExpr.GoRaw("map[string]any{\"host\": New_sys__net__Host(hxrt.StringFromLiteral(rawHost)), \"port\": port}"))
			]),
			GoDecl.GoFuncDecl("peer", {
				name: "self",
				typeName: "*sys__net__Socket"
			}, [], ["map[string]any"], [
				GoStmt.GoIf(GoExpr.GoRaw("self == nil || self.conn == nil"), [
					GoStmt.GoReturn(GoExpr.GoRaw("map[string]any{\"host\": hxrt__host_empty(), \"port\": 0}"))
				], null),
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt__socket_addrInfo"), [
					GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoSelector(GoExpr.GoIdent("self"), "conn"), "RemoteAddr"), [])
				]))
			]),
			GoDecl.GoFuncDecl("host", {
				name: "self",
				typeName: "*sys__net__Socket"
			}, [], ["map[string]any"], [
				GoStmt.GoIf(GoExpr.GoRaw("self == nil"), [
					GoStmt.GoReturn(GoExpr.GoRaw("map[string]any{\"host\": hxrt__host_empty(), \"port\": 0}"))
				], null),
				GoStmt.GoIf(GoExpr.GoRaw("self.conn != nil"), [
					GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt__socket_addrInfo"), [
						GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoSelector(GoExpr.GoIdent("self"), "conn"), "LocalAddr"), [])
					]))
				], null),
				GoStmt.GoIf(GoExpr.GoRaw("self.listener != nil"), [
					GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt__socket_addrInfo"), [
						GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoSelector(GoExpr.GoIdent("self"), "listener"), "Addr"), [])
					]))
				], null),
				GoStmt.GoReturn(GoExpr.GoRaw("map[string]any{\"host\": hxrt__host_empty(), \"port\": 0}"))
			]),
			GoDecl.GoFuncDecl("setTimeout", {
				name: "self",
				typeName: "*sys__net__Socket"
			}, [{name: "timeout", typeName: "float64"}], [], [
				GoStmt.GoRaw("if self == nil {"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if timeout < 0 {"),
				GoStmt.GoRaw("\tself.hasTimeout = false"),
				GoStmt.GoRaw("\tself.timeout = 0"),
				GoStmt.GoRaw("} else {"),
				GoStmt.GoRaw("\tself.hasTimeout = true"),
				GoStmt.GoRaw("\tself.timeout = timeout"),
				GoStmt.GoRaw("}"),
				GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("self"), "hxrt__socket_applyConnDeadline"), [])),
				GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("self"), "hxrt__socket_applyListenerDeadline"), []))
			]),
			GoDecl.GoFuncDecl("waitForRead", {
				name: "self",
				typeName: "*sys__net__Socket"
			}, [], [], [
				GoStmt.GoRaw("if self == nil {"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("_ = sys__net__Socket_select_([]*sys__net__Socket{self}, []*sys__net__Socket{}, []*sys__net__Socket{}, -1)")
			]),
			GoDecl.GoFuncDecl("setBlocking", {
				name: "self",
				typeName: "*sys__net__Socket"
			}, [{name: "b", typeName: "bool"}], [], [
				GoStmt.GoRaw("if self == nil {"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("self.blocking = b"),
				GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("self"), "hxrt__socket_applyConnDeadline"), [])),
				GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("self"), "hxrt__socket_applyListenerDeadline"), []))
			]),
			GoDecl.GoFuncDecl("setFastSend", {
				name: "self",
				typeName: "*sys__net__Socket"
			}, [{name: "b", typeName: "bool"}], [], [
				GoStmt.GoRaw("if self == nil {"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("self.fastSend = b"),
				GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("self"), "hxrt__socket_applyFastSend"), []))
			]),
			GoDecl.GoFuncDecl("sys__net__Socket_select_", null, [
				{
					name: "read",
					typeName: "[]*sys__net__Socket"
				},
				{name: "write", typeName: "[]*sys__net__Socket"},
				{name: "others", typeName: "[]*sys__net__Socket"},
				{name: "timeout", typeName: "...float64"}
			], ["map[string]any"], [
				GoStmt.GoIf(GoExpr.GoBinary("==", GoExpr.GoIdent("read"), GoExpr.GoNil),
					[GoStmt.GoAssign(GoExpr.GoIdent("read"), GoExpr.GoRaw("[]*sys__net__Socket{}"))], null),
				GoStmt.GoIf(GoExpr.GoBinary("==", GoExpr.GoIdent("write"), GoExpr.GoNil),
					[GoStmt.GoAssign(GoExpr.GoIdent("write"), GoExpr.GoRaw("[]*sys__net__Socket{}"))], null),
				GoStmt.GoIf(GoExpr.GoBinary("==", GoExpr.GoIdent("others"), GoExpr.GoNil),
					[GoStmt.GoAssign(GoExpr.GoIdent("others"), GoExpr.GoRaw("[]*sys__net__Socket{}"))], null),
				GoStmt.GoRaw("effectiveTimeout := -1.0"),
				GoStmt.GoRaw("if len(timeout) > 0 {"),
				GoStmt.GoRaw("\teffectiveTimeout = timeout[0]"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("readyRead := make([]*sys__net__Socket, 0, len(read))"),
				GoStmt.GoRaw("readyWrite := make([]*sys__net__Socket, 0, len(write))"),
				GoStmt.GoRaw("readyOther := make([]*sys__net__Socket, 0, len(others))"),
				GoStmt.GoRaw("for _, socket := range read {"),
				GoStmt.GoRaw("\tif socket == nil || socket.conn == nil || socket.input == nil || socket.input.reader == nil {"),
				GoStmt.GoRaw("\t\tcontinue"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treader := socket.input.reader"),
				GoStmt.GoRaw("\tif reader.Buffered() > 0 {"),
				GoStmt.GoRaw("\t\treadyRead = append(readyRead, socket)"),
				GoStmt.GoRaw("\t\tcontinue"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tif effectiveTimeout >= 0 {"),
				GoStmt.GoRaw("\t\tdeadline := time.Now()"),
				GoStmt.GoRaw("\t\tif effectiveTimeout > 0 {"),
				GoStmt.GoRaw("\t\t\tdeadline = time.Now().Add(time.Duration(effectiveTimeout * float64(time.Second)))"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t\t_ = socket.conn.SetReadDeadline(deadline)"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\t_, err := reader.Peek(1)"),
				GoStmt.GoRaw("\tsocket.hxrt__socket_applyConnDeadline()"),
				GoStmt.GoRaw("\tif err == nil {"),
				GoStmt.GoRaw("\t\treadyRead = append(readyRead, socket)"),
				GoStmt.GoRaw("\t\tcontinue"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tif netErr, ok := err.(net.Error); ok && netErr.Timeout() {"),
				GoStmt.GoRaw("\t\tcontinue"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treadyOther = append(readyOther, socket)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("for _, socket := range write {"),
				GoStmt.GoRaw("\tif socket == nil || socket.conn == nil {"),
				GoStmt.GoRaw("\t\tcontinue"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treadyWrite = append(readyWrite, socket)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("for _, socket := range others {"),
				GoStmt.GoRaw("\tif socket == nil {"),
				GoStmt.GoRaw("\t\tcontinue"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treadyOther = append(readyOther, socket)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoReturn(GoExpr.GoRaw("map[string]any{\"read\": readyRead, \"write\": readyWrite, \"others\": readyOther}"))
			]),
			GoDecl.GoFuncDecl("readLine", {
				name: "self",
				typeName: "*sys__net__SocketInput"
			}, [], ["*string"], [
				GoStmt.GoRaw("if self == nil || self.reader == nil {"),
				GoStmt.GoRaw("\treturn hxrt.StringFromLiteral(\"\")"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if self.socket != nil {"),
				GoStmt.GoRaw("\tself.socket.hxrt__socket_applyConnDeadline()"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("line, err := self.reader.ReadString('\\n')"),
				GoStmt.GoRaw("if err != nil && len(line) == 0 {"),
				GoStmt.GoRaw("\treturn hxrt.StringFromLiteral(\"\")"),
				GoStmt.GoRaw("}"),
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [
					GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("strings"), "TrimRight"), [GoExpr.GoIdent("line"), GoExpr.GoStringLiteral("\r\n")])
				]))
			]),
			GoDecl.GoFuncDecl("writeString", {
				name: "self",
				typeName: "*sys__net__SocketOutput"
			}, [{name: "value", typeName: "*string"}], [], [
				GoStmt.GoRaw("if self == nil || self.writer == nil || value == nil {"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if self.socket != nil {"),
				GoStmt.GoRaw("\tself.socket.hxrt__socket_applyConnDeadline()"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if _, err := self.writer.WriteString(*hxrt.StdString(value)); err != nil {"),
				GoStmt.GoRaw("\thxrt.Throw(err)"),
				GoStmt.GoRaw("}")
			]),
			GoDecl.GoFuncDecl("flush", {
				name: "self",
				typeName: "*sys__net__SocketOutput"
			}, [], [], [
				GoStmt.GoRaw("if self == nil || self.writer == nil {"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if self.socket != nil {"),
				GoStmt.GoRaw("\tself.socket.hxrt__socket_applyConnDeadline()"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if err := self.writer.Flush(); err != nil {"),
				GoStmt.GoRaw("\thxrt.Throw(err)"),
				GoStmt.GoRaw("}")
			])
		];
	}

	function lowerClassDecls(classType:ClassType):Array<GoDecl> {
		if (classType.isInterface) {
			return lowerInterfaceDecls(classType);
		}

		var decls = new Array<GoDecl>();
		var typeName = classTypeName(classType);
		var superClass = projectSuperClass(classType);
		var ioSubclassKind = ioStdlibSubclassKind(classType);
		if (ioSubclassKind != null) {
			requiresIoHelperSurface = true;
		}

		var instanceDataFields = new Array<GoParam>();
		var instanceMethods = new Array<{name:String, func:TFunc}>();
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
							instanceMethods.push({name: field.name, func: methodFunc});
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
					params: lowerFunctionParams(method.func),
					results: lowerFunctionResults(method.func.t)
				});
			}
			decls.push(GoDecl.GoInterfaceDecl(interfaceSymbol(classType), interfaceMethods));
			decls.push(GoDecl.GoStructDecl(typeName, instanceFields));
			decls.push(lowerConstructorDecl(classType, ctorFunc, superClass));
		}

		for (method in instanceMethods) {
			decls.push(lowerInstanceMethodDecl(classType, method.name, method.func));
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
					var valueExpr = field.expr();
					decls.push(GoDecl.GoGlobalVarDecl(symbol, scalarGoType(field.type), valueExpr == null ? null : lowerExpr(valueExpr).expr));
				case FMethod(_):
					var func = unwrapFunction(field.expr());
					if (func != null) {
						decls.push(lowerFunctionDecl(symbol, func, null, classType.module));
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
					var method = lowerInterfaceMethod(field);
					if (method != null && !seen.exists(method.name)) {
						seen.set(method.name, true);
						methods.push(method);
					}
				case _:
			}
		}
		return [GoDecl.GoInterfaceDecl(classTypeName(classType), methods)];
	}

	function lowerInterfaceMethod(field:ClassField):Null<GoInterfaceMethod> {
		var followed = Context.follow(field.type);
		return switch (followed) {
			case TFun(args, returnType):
				{
					name: normalizeIdent(field.name),
					params: lowerTypedFunArgs(args),
					results: lowerFunctionResults(returnType)
				};
			case _:
				var methodFunc = unwrapFunction(field.expr());
				if (methodFunc == null) {
					null;
				} else {
					{
						name: normalizeIdent(field.name),
						params: lowerFunctionParams(methodFunc),
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
				typeName: scalarGoType(arg.t)
			});
		}
		return out;
	}

	function lowerFunctionDecl(name:String, func:TFunc, receiver:Null<GoParam>, ?sourceModule:String):GoDecl {
		pushFunctionVarNameScope();
		var params = lowerFunctionParams(func);
		var results = lowerFunctionResults(func.t);
		pushFunctionReturnType(func.t);
		var body = lowerFunctionBody(func.expr);
		prependLineDirective(body, func.expr.pos, sourceModule);
		popFunctionReturnType();
		popFunctionVarNameScope();
		return GoDecl.GoFuncDecl(name, receiver, params, results, body);
	}

	function lowerConstructorDecl(classType:ClassType, ctorFunc:Null<TFunc>, superClass:Null<ClassType>):GoDecl {
		pushFunctionVarNameScope();
		var typeName = classTypeName(classType);
		var params = ctorFunc == null ? [] : lowerFunctionParams(ctorFunc);
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
			body.push(GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoSelector(GoExpr.GoIdent("self"), superTypeName), "__hx_this"), GoExpr.GoIdent("self")));
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

	function lowerInstanceMethodDecl(classType:ClassType, fieldName:String, func:TFunc):GoDecl {
		return lowerFunctionDecl(normalizeIdent(fieldName), func, {
			name: "self",
			typeName: "*" + classTypeName(classType)
		}, classType.module);
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
				addMethod(ioSyntheticMethod(receiverType, "close", [], [], [GoStmt.GoAssign(GoExpr.GoIdent("_"), GoExpr.GoIdent("self"))]), "close");
			}
			if (!hasMethod("readByte")) {
				addMethod(ioSyntheticMethod(receiverType, "readByte", [], ["int"], [
					GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("hxrt.Throw"), [
						GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoStringLiteral("Not implemented")])
					])),
					GoStmt.GoReturn(GoExpr.GoIntLiteral(0))
				]), "readByte");
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
				addMethod(ioSyntheticMethod(receiverType, "flush", [], [], [GoStmt.GoAssign(GoExpr.GoIdent("_"), GoExpr.GoIdent("self"))]), "flush");
			}
			if (!hasMethod("close")) {
				addMethod(ioSyntheticMethod(receiverType, "close", [], [], [GoStmt.GoAssign(GoExpr.GoIdent("_"), GoExpr.GoIdent("self"))]), "close");
			}
			if (!hasMethod("prepare")) {
				addMethod(ioSyntheticMethod(receiverType, "prepare", [{name: "nbytes", typeName: "int"}], [], [
					GoStmt.GoAssign(GoExpr.GoIdent("_"), GoExpr.GoIdent("self")),
					GoStmt.GoAssign(GoExpr.GoIdent("_"), GoExpr.GoIdent("nbytes"))
				]), "prepare");
			}
			if (!hasMethod("writeByte")) {
				addMethod(ioSyntheticMethod(receiverType, "writeByte", [{name: "c", typeName: "int"}], [], [
					GoStmt.GoAssign(GoExpr.GoIdent("_"), GoExpr.GoIdent("c")),
					GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("hxrt.Throw"), [
						GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoStringLiteral("Not implemented")])
					]))
				]), "writeByte");
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

	function collectDispatchMethods(classType:ClassType):Array<{name:String, func:TFunc}> {
		var orderedNames = new Array<String>();
		var methods = new Map<String, TFunc>();

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
						methods.set(methodName, methodFunc);
					case _:
				}
			}
		}

		collect(classType);

		var out = new Array<{name:String, func:TFunc}>();
		for (name in orderedNames) {
			out.push({
				name: name,
				func: methods.get(name)
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

	function lowerFunctionParams(func:TFunc):Array<GoParam> {
		var params = new Array<GoParam>();
		for (arg in func.args) {
			registerOptionalPrimitiveParam(arg.v, arg.value != null);
			params.push({
				name: localVarName(arg.v),
				typeName: scalarGoType(arg.v.t)
			});
		}
		return params;
	}

	function buildFunctionInfo(func:TFunc):FunctionInfo {
		return {
			defaults: [for (arg in func.args) arg.value]
		};
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
				var lambdaAlias = value == null ? null : lambdaFunctionAliasName(value);
				if (lambdaAlias != null) {
					registerLocalLambdaAlias(variableName, lambdaAlias);
				}

				var lowered = value == null ? null : lowerExprWithPrefix(value);
				var prefix = lowered == null ? [] : lowered.prefix;
				var loweredValue = lowered == null ? null : lowered.expr;
				if (value != null && loweredValue != null) {
					loweredValue = upcastIfNeeded(loweredValue, value.t, variable.t);
					loweredValue = coerceAnyExprToType(loweredValue, value.t, variable.t, exprBackedByAny(value)
						|| shouldForceAnyCoerce(value.t, variable.t));
				}
				var goType = typeToGoType(variable.t);
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
				[
					GoStmt.GoIf(lowerExpr(condition).expr, lowerToStatements(thenBranch), elseBranch == null ? null : lowerToStatements(elseBranch))
				];
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
						var targetExpr = lowerLValue(arrayCall.target);
						var shouldMaskToByte = isBytesBufferStorageArray(arrayCall.target);
						var appendArgs = [targetExpr];
						for (arg in args) {
							var appendValue = lowerExpr(arg).expr;
							if (shouldMaskToByte) {
								appendValue = GoExpr.GoBinary("&", appendValue, GoExpr.GoIntLiteral(255));
							}
							appendArgs.push(appendValue);
						}
						[GoStmt.GoAssign(targetExpr, GoExpr.GoCall(GoExpr.GoIdent("append"), appendArgs))];
					} else if (arrayCall != null && arrayCall.methodName == "pop") {
						var targetExpr = lowerLValue(arrayCall.target);
						var lenExpr = GoExpr.GoCall(GoExpr.GoIdent("len"), [targetExpr]);
						[
							GoStmt.GoIf(GoExpr.GoBinary(">", lenExpr, GoExpr.GoIntLiteral(0)), [
								GoStmt.GoAssign(targetExpr, GoExpr.GoSlice(targetExpr, null, GoExpr.GoBinary("-", lenExpr, GoExpr.GoIntLiteral(1))))
							], null)
						];
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
				GoStmt.GoVarDecl(temp, typeToGoType(resultType), null, false),
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
		var loweredThen = lowerExprWithPrefix(thenBranch);
		var loweredElse = lowerExprWithPrefix(elseExpr);
		var temp = freshTempName("hx_if");
		var loweredThenValue = upcastIfNeeded(loweredThen.expr, thenBranch.t, resultType);
		var loweredElseValue = upcastIfNeeded(loweredElse.expr, elseExpr.t, resultType);
		loweredThenValue = coerceAnyExprToType(loweredThenValue, thenBranch.t, resultType, exprBackedByAny(thenBranch) || shouldForceAnyCoerce(thenBranch.t,
			resultType));
		loweredElseValue = coerceAnyExprToType(loweredElseValue, elseExpr.t, resultType, exprBackedByAny(elseExpr) || shouldForceAnyCoerce(elseExpr.t,
			resultType));

		var prefix = [GoStmt.GoVarDecl(temp, typeToGoType(resultType), null, false)].concat(loweredCondition.prefix);

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
					GoStmt.GoVarDecl(temp, typeToGoType(resultType), null, false),
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
				GoStmt.GoVarDecl(temp, typeToGoType(resultType), null, false),
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
		for (inner in exprs) {
			out = out.concat(lowerToStatements(inner));
		}
		popLocalScope();
		return out;
	}

	function pushLocalScope():Void {
		localFunctionScopes.push(new Map<String, FunctionInfo>());
		localLambdaAliasScopes.push(new Map<String, String>());
		localRestIteratorScopes.push([]);
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
	}

	function pushFunctionVarNameScope():Void {
		functionVarNameScopes.push(new Map<Int, String>());
		functionVarNameCountScopes.push(new Map<String, Int>());
		optionalPrimitiveParamScopes.push(new Map<Int, Bool>());
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

	function registerOptionalPrimitiveParam(variable:TVar, hasDefaultExpr:Bool):Void {
		if (!hasDefaultExpr) {
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
		scope.set(variable.id, true);
	}

	function currentOptionalPrimitiveParamScope():Null<Map<Int, Bool>> {
		if (optionalPrimitiveParamScopes.length == 0) {
			return null;
		}
		return optionalPrimitiveParamScopes[optionalPrimitiveParamScopes.length - 1];
	}

	function isRegisteredOptionalPrimitiveParam(variable:TVar):Bool {
		var index = optionalPrimitiveParamScopes.length - 1;
		while (index >= 0) {
			var scope = optionalPrimitiveParamScopes[index];
			if (scope.exists(variable.id)) {
				return true;
			}
			index--;
		}
		return false;
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
				staticFunctionInfos.get(staticSymbol(classRef.get(), field.get().name));
			case TLocal(variable):
				lookupLocalFunction(localVarName(variable));
			case _:
				null;
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
				if (useTypedGoConcurrencySpecialization() && isGoChanClass(classType)) {
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
							expr: GoExpr.GoCall(GoExpr.GoIdent(constructorSymbol(classType)), [for (arg in args) lowerExpr(arg).expr]),
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
						expr: GoExpr.GoCall(GoExpr.GoIdent(constructorSymbol(classType)), [for (arg in args) lowerExpr(arg).expr]),
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
				var variableGoType = typeToGoType(variable.t);
				var exprGoType = typeToGoType(expr.t);
				if (variableGoType == "any" && exprGoType != "any") {
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
				lowerCall(callee, args, expr.t);
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
			expr: GoExpr.GoCall(GoExpr.GoFuncLiteral([], [typeToGoType(resultType)], lowered.prefix.concat([GoStmt.GoReturn(lowered.expr)])), []),
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

				if (isHaxeExceptionType(target.t) && resolved.name == "message") {
					return {
						expr: GoExpr.GoCall(GoExpr.GoIdent("hxrt.ExceptionMessage"), [loweredTarget]),
						isStringLike: true
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
				if (isHaxeExceptionType(target.t) && resolved.name == "message") {
					return {
						expr: GoExpr.GoCall(GoExpr.GoIdent("hxrt.ExceptionMessage"), [loweredTarget]),
						isStringLike: true
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

		var loweredArgs = new Array<GoExpr>();
		for (index in 0...args.length) {
			var arg = args[index];
			var loweredArg = lowerCallArgExpr(arg);
			var paramType = callParamType(callee.t, index);
			if (paramType != null) {
				loweredArg = upcastIfNeeded(loweredArg, arg.t, paramType);
			}
			loweredArg = normalizeExternCallArg(callee, loweredArg, paramType, returnType);
			loweredArgs.push(loweredArg);
		}
		var functionInfo = resolveFunctionInfo(callee);
		if (functionInfo != null && loweredArgs.length < functionInfo.defaults.length) {
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
				return {
					expr: GoExpr.GoCall(GoExpr.GoIdent(metalChanShimName("go__concurrency_send", elementGoType)), [channelNative, value]),
					isStringLike: false
				};
			case "trySend":
				var value = args.length > 0 ? lowerExpr(args[0]).expr : GoExpr.GoNil;
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

	function tryLambdaSourcePlan(sourceExpr:TypedExpr):Null<LambdaSourcePlan> {
		if (isArrayType(sourceExpr.t)) {
			var elementType = arrayElementGoType(sourceExpr.t);
			var loweredSourceExpr = lowerExpr(sourceExpr).expr;
			return {
				domain: "array",
				elementType: elementType,
				sourceExpr: loweredSourceExpr,
				sourceType: "[]" + elementType
			};
		}

		var listElement = haxeDsListElementType(sourceExpr.t);
		if (listElement != null) {
			requireStdlibShimGroup("ds");
			var loweredSourceExpr = lowerExpr(sourceExpr).expr;
			return {
				domain: "list",
				elementType: scalarGoType(listElement),
				sourceExpr: loweredSourceExpr,
				sourceType: "*haxe__ds__List"
			};
		}

		return null;
	}

	function lowerLambdaManualIteratorProtocolSource(sourceExpr:GoExpr, ?sourcePlan:Null<LambdaSourcePlan>):GoExpr {
		var sourceName = freshTempName("hx_lambda_source");
		var wrappedName = freshTempName("hx_lambda_wrapped");
		var iteratorFactoryBody:Array<GoStmt> = switch (sourcePlan == null ? "generic" : sourcePlan.domain) {
			case "array":
				var indexName = freshTempName("hx_lambda_index");
				var valueName = freshTempName("hx_lambda_value");
				var iteratorMapLiteral = "map[string]any{\"hasNext\": func() bool { return " + indexName + " < len(" + sourceName
					+ ") }, \"next\": func() any { " + valueName + " := " + sourceName + "[" + indexName + "]; " + indexName + "++; return " + valueName +
					" }}";
				[
					GoStmt.GoVarDecl(indexName, "int", GoExpr.GoIntLiteral(0), true),
					GoStmt.GoReturn(GoExpr.GoRaw(iteratorMapLiteral))
				];
			case "list":
				var indexName = freshTempName("hx_lambda_index");
				var valueName = freshTempName("hx_lambda_value");
				var iteratorMapLiteral = "map[string]any{\"hasNext\": func() bool { return " + indexName + " < len(" + sourceName
					+ ".items) }, \"next\": func() any { " + valueName + " := " + sourceName + ".items[" + indexName + "]; " + indexName + "++; return "
					+ valueName + " }}";
				[
					GoStmt.GoVarDecl(indexName, "int", GoExpr.GoIntLiteral(0), true),
					GoStmt.GoReturn(GoExpr.GoRaw(iteratorMapLiteral))
				];
			case _:
				var iteratorName = freshTempName("hx_lambda_iterator");
				var iteratorMapLiteral = "map[string]any{\"hasNext\": func() bool { return " + iteratorName + ".hasNext() }, \"next\": func() any { return "
					+ iteratorName + ".next() }}";
				[
					GoStmt.GoVarDecl(iteratorName, null, GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent(sourceName), "iterator"), []), true),
					GoStmt.GoReturn(GoExpr.GoRaw(iteratorMapLiteral))
				];
		};
		return GoExpr.GoCall(GoExpr.GoFuncLiteral([], ["map[string]any"], [
			GoStmt.GoVarDecl(sourceName, null, sourceExpr, true),
			GoStmt.GoVarDecl(wrappedName, "map[string]any", GoExpr.GoRaw("map[string]any{}"), true),
			GoStmt.GoAssign(GoExpr.GoIndex(GoExpr.GoIdent(wrappedName), GoExpr.GoStringLiteral("iterator")),
				GoExpr.GoFuncLiteral([], ["map[string]any"], iteratorFactoryBody)),
			GoStmt.GoReturn(GoExpr.GoIdent(wrappedName))
		]), []);
	}

	function lowerLambdaDynamicIterableSource(sourceExpr:TypedExpr):GoExpr {
		return lowerLambdaManualIteratorProtocolSource(lowerExpr(sourceExpr).expr);
	}

	function firstFunctionArgType(type:Type):Null<Type> {
		return switch (Context.follow(type)) {
			case TFun(args, _):
				args.length > 0 ? args[0].t : null;
			case _:
				null;
		};
	}

	function lambdaFunctionAliasName(expr:TypedExpr):Null<String> {
		return switch (expr.expr) {
			case TField(_, FStatic(classRef, fieldRef)):
				var classType = classRef.get();
				var field = fieldRef.get();
				if (classType.pack.length == 0 && classType.name == "Lambda" && (field.name == "map" || field.name == "fold")) {
					field.name;
				} else {
					null;
				}
			case TLocal(variable):
				lookupLocalLambdaAlias(localVarName(variable));
			case TMeta(_, inner):
				lambdaFunctionAliasName(inner);
			case TParenthesis(inner):
				lambdaFunctionAliasName(inner);
			case TCast(inner, _):
				lambdaFunctionAliasName(inner);
			case _:
				null;
		};
	}

	function lowerLambdaPredicateAnyAdapter(predicateExpr:GoExpr, predicateType:Type):GoExpr {
		var rawArgName = freshTempName("hx_lambda_arg");
		var adaptedArgExpr:GoExpr = GoExpr.GoIdent(rawArgName);
		var argType = firstFunctionArgType(predicateType);
		if (argType != null) {
			adaptedArgExpr = lowerNullableAwareTypeAssertExpr(adaptedArgExpr, argType);
		}
		return GoExpr.GoFuncLiteral([{name: rawArgName, typeName: "any"}], ["bool"], [GoStmt.GoReturn(GoExpr.GoCall(predicateExpr, [adaptedArgExpr]))]);
	}

	function lowerLambdaMapperAnyAdapter(mapperExpr:GoExpr, mapperType:Type):GoExpr {
		var rawArgName = freshTempName("hx_lambda_arg");
		var adaptedArgExpr:GoExpr = GoExpr.GoIdent(rawArgName);
		var argType = firstFunctionArgType(mapperType);
		if (argType != null) {
			adaptedArgExpr = lowerNullableAwareTypeAssertExpr(adaptedArgExpr, argType);
		}
		return GoExpr.GoFuncLiteral([{name: rawArgName, typeName: "any"}], ["any"], [GoStmt.GoReturn(GoExpr.GoCall(mapperExpr, [adaptedArgExpr]))]);
	}

	function lowerLambdaConsumerAnyAdapter(consumerExpr:GoExpr, consumerType:Type):GoExpr {
		var rawArgName = freshTempName("hx_lambda_arg");
		var adaptedArgExpr:GoExpr = GoExpr.GoIdent(rawArgName);
		var argType = firstFunctionArgType(consumerType);
		if (argType != null) {
			adaptedArgExpr = lowerNullableAwareTypeAssertExpr(adaptedArgExpr, argType);
		}
		return GoExpr.GoFuncLiteral([{name: rawArgName, typeName: "any"}], [], [GoStmt.GoExprStmt(GoExpr.GoCall(consumerExpr, [adaptedArgExpr]))]);
	}

	function lowerLambdaFolderAnyAdapter(folderExpr:GoExpr, folderType:Type):GoExpr {
		var rawValueName = freshTempName("hx_lambda_value");
		var rawAccName = freshTempName("hx_lambda_acc");
		var adaptedValueExpr:GoExpr = GoExpr.GoIdent(rawValueName);
		var adaptedAccExpr:GoExpr = GoExpr.GoIdent(rawAccName);
		switch (Context.follow(folderType)) {
			case TFun(args, _):
				if (args.length > 0) {
					adaptedValueExpr = lowerNullableAwareTypeAssertExpr(adaptedValueExpr, args[0].t);
				}
				if (args.length > 1) {
					adaptedAccExpr = lowerNullableAwareTypeAssertExpr(adaptedAccExpr, args[1].t);
				}
			case _:
		}
		return GoExpr.GoFuncLiteral([{name: rawValueName, typeName: "any"}, {name: rawAccName, typeName: "any"}], ["any"],
			[GoStmt.GoReturn(GoExpr.GoCall(folderExpr, [adaptedValueExpr, adaptedAccExpr]))]);
	}

	function lowerLambdaAnyArrayCoerce(anySliceExpr:GoExpr, targetArrayType:Type):GoExpr {
		if (!isArrayType(targetArrayType)) {
			return anySliceExpr;
		}
		var targetElementType = arrayElementType(targetArrayType);
		if (targetElementType == null) {
			return anySliceExpr;
		}
		var targetElementGoType = arrayElementGoType(targetArrayType);
		if (targetElementGoType == "any") {
			return anySliceExpr;
		}
		var rawName = freshTempName("hx_lambda_raw");
		var outName = freshTempName("hx_lambda_out");
		var itemName = freshTempName("hx_lambda_item");
		var convertedItemExpr = lowerNullableAwareTypeAssertExpr(GoExpr.GoIdent(itemName), targetElementType);
		var outType = "[]" + targetElementGoType;
		return GoExpr.GoCall(GoExpr.GoFuncLiteral([{name: rawName, typeName: "[]any"}], [outType], [
			GoStmt.GoVarDecl(outName, outType, GoExpr.GoRaw("make(" + outType + ", 0, len(" + rawName + "))"), true),
			GoStmt.GoRaw("for _, " + itemName + " := range " + rawName + " {"),
			GoStmt.GoAssign(GoExpr.GoIdent(outName), GoExpr.GoCall(GoExpr.GoIdent("append"), [GoExpr.GoIdent(outName), convertedItemExpr])),
			GoStmt.GoRaw("}"),
			GoStmt.GoReturn(GoExpr.GoIdent(outName))
		]), [anySliceExpr]);
	}

	function lowerLambdaFunctionValueCall(callee:TypedExpr, args:Array<TypedExpr>, returnType:Type):Null<LoweredExpr> {
		if (args.length < 2) {
			return null;
		}
		var alias = lambdaFunctionAliasName(callee);
		if (alias == null) {
			return null;
		}

		// Function-value Lambda.map alias call path.
		if (alias == "map") {
			if (args.length != 2) {
				Context.fatalError("Lambda.map expects exactly 2 arguments", callee.pos);
			}
			var calleeExpr = lowerExpr(callee).expr;
			var dynamicSourceExpr = lowerLambdaDynamicIterableSource(args[0]);
			var mapperExpr = lowerExpr(args[1]).expr;
			var adaptedMapperExpr = lowerLambdaMapperAnyAdapter(mapperExpr, args[1].t);
			var mappedAnyExpr = GoExpr.GoCall(calleeExpr, [dynamicSourceExpr, adaptedMapperExpr]);
			return {
				expr: lowerLambdaAnyArrayCoerce(mappedAnyExpr, returnType),
				isStringLike: false
			};
		}

		// Function-value Lambda.fold alias call path.
		if (alias == "fold") {
			if (args.length != 3) {
				Context.fatalError("Lambda.fold expects exactly 3 arguments", callee.pos);
			}
			var calleeExpr = lowerExpr(callee).expr;
			var dynamicSourceExpr = lowerLambdaDynamicIterableSource(args[0]);
			var folderExpr = lowerExpr(args[1]).expr;
			var initExpr = lowerExpr(args[2]).expr;
			var adaptedFolderExpr = lowerLambdaFolderAnyAdapter(folderExpr, args[1].t);
			var foldedAnyExpr = GoExpr.GoCall(calleeExpr, [dynamicSourceExpr, adaptedFolderExpr, initExpr]);
			return {
				expr: lowerNullableAwareTypeAssertExpr(foldedAnyExpr, returnType),
				isStringLike: false
			};
		}

		return null;
	}

	function lowerLambdaStaticCall(callee:TypedExpr, args:Array<TypedExpr>, returnType:Type):Null<LoweredExpr> {
		if (isStaticCall(callee, "Lambda", [], "count") || isGeneratedLambdaCall(callee, "count")) {
			var supportsOptimizedCount = args.length == 1 || (args.length == 2 && isNullLiteralExpr(args[1]));
			if (!supportsOptimizedCount) {
				return null;
			}
			var sourcePlan = tryLambdaSourcePlan(args[0]);
			if (sourcePlan == null) {
				var dynamicSourceExpr = lowerLambdaDynamicIterableSource(args[0]);
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

		if (isStaticCall(callee, "Lambda", [], "empty") || isGeneratedLambdaCall(callee, "empty")) {
			if (args.length != 1) {
				Context.fatalError("Lambda.empty expects exactly 1 argument", callee.pos);
			}
			var sourcePlan = tryLambdaSourcePlan(args[0]);
			if (sourcePlan == null) {
				var dynamicSourceExpr = lowerLambdaDynamicIterableSource(args[0]);
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

		if (isStaticCall(callee, "Lambda", [], "exists") || isGeneratedLambdaCall(callee, "exists")) {
			if (args.length != 2) {
				Context.fatalError("Lambda.exists expects exactly 2 arguments", callee.pos);
			}
			var sourcePlan = tryLambdaSourcePlan(args[0]);
			if (sourcePlan == null) {
				var dynamicSourceExpr = lowerLambdaDynamicIterableSource(args[0]);
				var predicateExpr = lowerExpr(args[1]).expr;
				var adaptedPredicateExpr = lowerLambdaPredicateAnyAdapter(predicateExpr, args[1].t);
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

		if (isStaticCall(callee, "Lambda", [], "has") || isGeneratedLambdaCall(callee, "has")) {
			if (args.length != 2) {
				Context.fatalError("Lambda.has expects exactly 2 arguments", callee.pos);
			}
			requireStdlibShimGroup("stdlib_symbols");
			var sourcePlan = tryLambdaSourcePlan(args[0]);
			if (sourcePlan == null) {
				var dynamicSourceExpr = lowerLambdaDynamicIterableSource(args[0]);
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

		if (isStaticCall(callee, "Lambda", [], "iter") || isGeneratedLambdaCall(callee, "iter")) {
			if (args.length != 2) {
				Context.fatalError("Lambda.iter expects exactly 2 arguments", callee.pos);
			}
			var sourcePlan = tryLambdaSourcePlan(args[0]);
			var iteratorSourceExpr = sourcePlan == null ? lowerLambdaDynamicIterableSource(args[0]) : lowerLambdaManualIteratorProtocolSource(sourcePlan.sourceExpr,
				sourcePlan);
			var consumerExpr = lowerExpr(args[1]).expr;
			var adaptedConsumerExpr = lowerLambdaConsumerAnyAdapter(consumerExpr, args[1].t);
			return {
				expr: GoExpr.GoCall(GoExpr.GoIdent("Lambda_iter"), [iteratorSourceExpr, adaptedConsumerExpr]),
				isStringLike: false
			};
		}

		if (isStaticCall(callee, "Lambda", [], "filter") || isGeneratedLambdaCall(callee, "filter")) {
			if (args.length != 2) {
				Context.fatalError("Lambda.filter expects exactly 2 arguments", callee.pos);
			}
			var sourcePlan = tryLambdaSourcePlan(args[0]);
			if (sourcePlan == null) {
				var dynamicSourceExpr = lowerLambdaDynamicIterableSource(args[0]);
				var predicateExpr = lowerExpr(args[1]).expr;
				var adaptedPredicateExpr = lowerLambdaPredicateAnyAdapter(predicateExpr, args[1].t);
				var filteredAnyExpr = GoExpr.GoCall(GoExpr.GoIdent("Lambda_filter"), [dynamicSourceExpr, adaptedPredicateExpr]);
				return {
					expr: lowerLambdaAnyArrayCoerce(filteredAnyExpr, returnType),
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

		if (isStaticCall(callee, "Lambda", [], "map") || isGeneratedLambdaCall(callee, "map")) {
			if (args.length != 2) {
				Context.fatalError("Lambda.map expects exactly 2 arguments", callee.pos);
			}
			var sourcePlan = tryLambdaSourcePlan(args[0]);
			if (sourcePlan == null) {
				var dynamicSourceExpr = lowerLambdaDynamicIterableSource(args[0]);
				var mapperExpr = lowerExpr(args[1]).expr;
				var adaptedMapperExpr = lowerLambdaMapperAnyAdapter(mapperExpr, args[1].t);
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
					expr: lowerLambdaAnyArrayCoerce(mappedAnyExpr, returnType),
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

		if (isStaticCall(callee, "Lambda", [], "fold") || isGeneratedLambdaCall(callee, "fold")) {
			if (args.length != 3) {
				Context.fatalError("Lambda.fold expects exactly 3 arguments", callee.pos);
			}
			var sourcePlan = tryLambdaSourcePlan(args[0]);
			if (sourcePlan == null) {
				var dynamicSourceExpr = lowerLambdaDynamicIterableSource(args[0]);
				var folderExpr = lowerExpr(args[1]).expr;
				var initExpr = lowerExpr(args[2]).expr;
				var adaptedFolderExpr = lowerLambdaFolderAnyAdapter(folderExpr, args[1].t);
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

	function isGeneratedLambdaCall(callee:TypedExpr, methodName:String):Bool {
		return switch (callee.expr) {
			case TIdent(name):
				name == ("Lambda_" + methodName);
			case TMeta(_, inner):
				isGeneratedLambdaCall(inner, methodName);
			case TParenthesis(inner):
				isGeneratedLambdaCall(inner, methodName);
			case TCast(inner, _):
				isGeneratedLambdaCall(inner, methodName);
			case _:
				false;
		};
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
		var moduleName = sourceModuleForPos(pos);
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
		var nullComparison = isNullLiteralExpr(left) || isNullLiteralExpr(right);
		var impossiblePrimitiveNullComparison = nullComparison
			&& ((isNullLiteralExpr(left) && isDefinitelyNonNullableType(right.t))
				|| (isNullLiteralExpr(right) && isDefinitelyNonNullableType(left.t)));
		var optionalPrimitiveLocalNullComparison = nullComparison
			&& ((isNullLiteralExpr(left) && isOptionalPrimitiveLocalExpr(right))
				|| (isNullLiteralExpr(right) && isOptionalPrimitiveLocalExpr(left)));
		impossiblePrimitiveNullComparison = impossiblePrimitiveNullComparison || optionalPrimitiveLocalNullComparison;
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
				if (isNullLiteralExpr(left) && isNullLiteralExpr(right)) {
					{
						expr: GoExpr.GoBoolLiteral(true),
						isStringLike: false
					};
				} else {
					var targetExpr = isNullLiteralExpr(left) ? rightLowered.expr : leftLowered.expr;
					{
						expr: GoExpr.GoCall(GoExpr.GoIdent("hxrt.AnyEqualsNull"), [targetExpr]),
						isStringLike: false
					};
				}
			case OpNotEq if (anyNullComparison):
				if (isNullLiteralExpr(left) && isNullLiteralExpr(right)) {
					{
						expr: GoExpr.GoBoolLiteral(false),
						isStringLike: false
					};
				} else {
					var targetExpr = isNullLiteralExpr(left) ? rightLowered.expr : leftLowered.expr;
					{
						expr: GoExpr.GoUnary("!", GoExpr.GoCall(GoExpr.GoIdent("hxrt.AnyEqualsNull"), [targetExpr])),
						isStringLike: false
					};
				}
			case OpAdd | OpSub | OpMult | OpDiv if (floatMode):
				{
					expr: GoExpr.GoBinary(binopSymbol(op), floatOperandExpr(leftLowered.expr, left.t), floatOperandExpr(rightLowered.expr, right.t)),
					isStringLike: false
				};
			case OpMod if (floatMode):
				{
					expr: GoExpr.GoCall(GoExpr.GoIdent("hxrt.FloatMod"), [
						floatOperandExpr(leftLowered.expr, left.t),
						floatOperandExpr(rightLowered.expr, right.t)
					]),
					isStringLike: false
				};
			case OpUShr if (int32Mode):
				var int32Left = coerceNullableIntOperandExpr(leftLowered.expr, left.t);
				var int32Right = coerceNullableIntOperandExpr(rightLowered.expr, right.t);
				{
					expr: lowerHaxeInt32BinopExpr(op, int32Left, int32Right),
					isStringLike: false
				};
			case OpAdd | OpSub | OpMult | OpMod | OpAnd | OpOr | OpXor | OpShl | OpShr if (int32Mode):
				var int32Left = coerceNullableIntOperandExpr(leftLowered.expr, left.t);
				var int32Right = coerceNullableIntOperandExpr(rightLowered.expr, right.t);
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
					expr: GoExpr.GoBinary(binopSymbol(op), leftLowered.expr, rightLowered.expr),
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

	function coerceNullableIntOperandExpr(expr:GoExpr, operandType:Type):GoExpr {
		if (!isNullableIntType(operandType)) {
			return expr;
		}
		return GoExpr.GoCall(GoExpr.GoIdent("hxrt.IntFromNullableAny"), [expr]);
	}

	function coerceNullableFloatOperandExpr(expr:GoExpr, operandType:Type):GoExpr {
		if (!isNullableFloatType(operandType)) {
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

	function floatOperandExpr(expr:GoExpr, operandType:Type):GoExpr {
		if (isNullableFloatType(operandType)) {
			return coerceNullableFloatOperandExpr(expr, operandType);
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
			case TLocal(variable):
				isRegisteredOptionalPrimitiveParam(variable);
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

	function exprBackedByAny(expr:TypedExpr):Bool {
		return switch (expr.expr) {
			case TLocal(variable):
				typeToGoType(variable.t) == "any";
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

	function isHaxeIoBaseClass(classType:ClassType):Bool {
		return GoTypeMapper.isHaxeIoBaseClass(classType);
	}

	function isHaxeExceptionType(type:Type):Bool {
		return GoTypeMapper.isHaxeExceptionType(type);
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

	function noteIoHelperFieldUsage(classType:ClassType, fieldName:String):Void {
		if (GoStdlibShimClassifier.needsIoHelperSurface(classType, fieldName, isIoInputHelperMethodName, isIoOutputHelperMethodName)) {
			requiresIoHelperSurface = true;
		}
	}

	function noteStdlibClass(classType:ClassType):Void {
		for (group in GoStdlibShimClassifier.requiredGroupsForClass(classType)) {
			requireStdlibShimGroup(group);
		}
	}

	function noteStdlibEnum(enumType:EnumType):Void {
		for (group in GoStdlibShimClassifier.requiredGroupsForEnum(enumType)) {
			requireStdlibShimGroup(group);
		}
	}

	function sortedRequiredStdlibShimGroups():Array<String> {
		var groups = [for (group in requiredStdlibShimGroups.keys()) group];
		groups.sort(Reflect.compare);
		return groups;
	}

	function inferRuntimeFeatures(requiredShimGroups:Array<String>):Array<String> {
		var classPaths = [for (classType in projectClasses) fullClassName(classType)];
		classPaths.sort(Reflect.compare);
		var enumPaths = [for (enumType in projectEnums) fullEnumName(enumType)];
		enumPaths.sort(Reflect.compare);
		return GoHxrtFeatureAnalyzer.inferFromUsage(classPaths, enumPaths, requiredShimGroups, requiresIoHelperSurface);
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
				if (isHaxeExceptionClass(classType) && field.get().name == "get_message") {
					target;
				} else {
					null;
				}
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
