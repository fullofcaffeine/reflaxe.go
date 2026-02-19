package reflaxe.go;

#if macro
import haxe.macro.Context;
import haxe.macro.Expr.Binop;
import haxe.macro.Expr.Unop;
import haxe.macro.PositionTools;
import haxe.macro.Type;
import reflaxe.go.ast.GoAST.GoDecl;
import reflaxe.go.ast.GoAST.GoExpr;
import reflaxe.go.ast.GoAST.GoFile;
import reflaxe.go.ast.GoAST.GoInterfaceMethod;
import reflaxe.go.ast.GoAST.GoParam;
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

private typedef FunctionInfo = {
	final defaults:Array<Null<TypedExpr>>;
}

private typedef ConstructorBodyLowering = {
	final superArgs:Null<Array<TypedExpr>>;
	final body:Array<GoStmt>;
}
#end

class GoCompiler {
	#if macro
	final compilationContext:CompilationContext;
	final staticFunctionInfos:Map<String, FunctionInfo>;
	final localFunctionScopes:Array<Map<String, FunctionInfo>>;
	final localRestIteratorScopes:Array<Array<String>>;
	final requiredStdlibShimGroups:Map<String, Bool>;
	final functionVarNameScopes:Array<Map<Int, String>>;
	final functionVarNameCountScopes:Array<Map<String, Int>>;
	final functionReturnTypeScopes:Array<Type>;
	var projectClasses:Array<ClassType>;
	var projectEnums:Array<EnumType>;
	var tempVarCounter:Int;
	#end

	public function new(?compilationContext:CompilationContext) {
		#if macro
		this.compilationContext = compilationContext == null ? new CompilationContext(GoProfile.Portable, "snapshot") : compilationContext;
		staticFunctionInfos = new Map<String, FunctionInfo>();
		localFunctionScopes = [];
		localRestIteratorScopes = [];
		requiredStdlibShimGroups = new Map<String, Bool>();
		functionVarNameScopes = [];
		functionVarNameCountScopes = [];
		functionReturnTypeScopes = [];
		projectClasses = [];
		projectEnums = [];
		tempVarCounter = 0;
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
		buildStaticFunctionInfoTable(classes);
		var decls = lowerEnums(enums).concat(lowerClasses(classes)).concat(lowerStdlibShimDecls());
		var imports = [compilationContext.runtimeImportPath];
		if (requiredStdlibShimGroups.exists("http")) {
			imports.push("bytes");
			imports.push("io");
			imports.push("net/http");
			imports.push("net/url");
			imports.push("strings");
			imports.push("time");
		}
		if (requiredStdlibShimGroups.exists("sys")) {
			imports.push("bufio");
			imports.push("os");
			imports.push("os/exec");
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
		if (requiredStdlibShimGroups.exists("regex_serializer")) {
			imports.push("encoding/json");
			imports.push("regexp");
			imports.push("strings");
		}
		if (requiredStdlibShimGroups.exists("net_socket")) {
			imports.push("bufio");
			imports.push("net");
			imports.push("strconv");
			imports.push("strings");
		}
		var mainFile:GoFile = {
			packageName: "main",
			imports: imports,
			decls: decls
		};
		var transformedFile = GoASTTransformer.transform(mainFile, compilationContext);

		return [
			{
				relativePath: "main.go",
				contents: GoASTPrinter.printFile(transformedFile)
			}
		];
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

	function fullClassName(classType:ClassType):String {
		return classType.pack.length == 0 ? classType.name : classType.pack.join(".") + "." + classType.name;
	}

	function fullEnumName(enumType:EnumType):String {
		return enumType.pack.length == 0 ? enumType.name : enumType.pack.join(".") + "." + enumType.name;
	}

	function projectSuperClass(classType:ClassType):Null<ClassType> {
		if (classType.superClass == null) {
			return null;
		}
		var superType = classType.superClass.t.get();
		return isProjectClass(superType) ? superType : null;
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

	function lowerStdlibShimDecls():Array<GoDecl> {
		var decls = new Array<GoDecl>();
		if (requiredStdlibShimGroups.exists("stdlib_symbols")) {
			// Symbol shims now include crypto/xml/zip behavior that depends on haxe.io.Bytes.
			requireStdlibShimGroup("io");
		}
		if (requiredStdlibShimGroups.exists("http")) {
			// Http request shims expose and consume haxe.io.Bytes payloads.
			requireStdlibShimGroup("io");
			requireStdlibShimGroup("ds");
		}
		if (requiredStdlibShimGroups.exists("json")) {
			decls = decls.concat(lowerJsonStdlibShimDecls());
		}
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
		if (requiredStdlibShimGroups.exists("stdlib_symbols")) {
			decls = decls.concat(lowerStdlibSymbolShimDecls());
		}
		if (requiredStdlibShimGroups.exists("regex_serializer")) {
			decls = decls.concat(lowerRegexSerializerShimDecls());
		}
		if (requiredStdlibShimGroups.exists("net_socket")) {
			decls = decls.concat(lowerNetSocketShimDecls());
		}
		return decls;
	}

	function lowerJsonStdlibShimDecls():Array<GoDecl> {
		return [
			GoDecl.GoStructDecl("haxe__Json", []),
			GoDecl.GoStructDecl("haxe__format__JsonParser", [{name: "source", typeName: "*string"}]),
			GoDecl.GoFuncDecl("New_haxe__format__JsonParser", null, [{name: "source", typeName: "*string"}], ["*haxe__format__JsonParser"],
				[GoStmt.GoReturn(GoExpr.GoRaw("&haxe__format__JsonParser{source: source}"))]),
			GoDecl.GoFuncDecl("doParse", {name: "self", typeName: "*haxe__format__JsonParser"}, [], ["any"], [
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.JsonParse"), [GoExpr.GoSelector(GoExpr.GoIdent("self"), "source")]))
			]),
			GoDecl.GoFuncDecl("haxe__format__JsonPrinter_print", null, [
				{
					name: "value",
					typeName: "any"
				},
				{name: "rest", typeName: "...any"}
			], ["*string"], [
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.JsonStringify"), [GoExpr.GoIdent("value")]))
			]),
			GoDecl.GoFuncDecl("haxe__Json_parse", null, [
				{
					name: "source",
					typeName: "*string"
				}
			], ["any"], [
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.JsonParse"), [GoExpr.GoIdent("source")]))
			]),
			GoDecl.GoFuncDecl("haxe__Json_stringify", null, [
				{
					name: "value",
					typeName: "any"
				},
				{name: "rest", typeName: "...any"}
			], ["*string"], [
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.JsonStringify"), [GoExpr.GoIdent("value")]))
			])
		];
	}

	function lowerIoStdlibShimDecls():Array<GoDecl> {
		return [
			GoDecl.GoStructDecl("haxe__io__Encoding", []),
			GoDecl.GoStructDecl("haxe__io__Input", []),
			GoDecl.GoStructDecl("haxe__io__Output", []),
			GoDecl.GoStructDecl("haxe__io__Bytes",
				[
					{
						name: "b",
						typeName: "[]int"
					},
					{name: "length", typeName: "int"}
				]),
			GoDecl.GoStructDecl("haxe__io__BytesBuffer", [{name: "b", typeName: "[]int"}]),
			GoDecl.GoFuncDecl("New_haxe__io__Input", null, [], ["*haxe__io__Input"], [GoStmt.GoReturn(GoExpr.GoRaw("&haxe__io__Input{}"))]),
			GoDecl.GoFuncDecl("New_haxe__io__Output", null, [], ["*haxe__io__Output"], [GoStmt.GoReturn(GoExpr.GoRaw("&haxe__io__Output{}"))]),
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
				GoStmt.GoVarDecl("raw", null, GoExpr.GoCall(GoExpr.GoIdent("hxrt.BytesFromString"), [GoExpr.GoIdent("value")]), true),
				GoStmt.GoReturn(GoExpr.GoRaw("&haxe__io__Bytes{b: raw, length: len(raw)}"))
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
			GoDecl.GoFuncDecl("get", {
				name: "self",
				typeName: "*haxe__io__Bytes"
			}, [{name: "pos", typeName: "int"}], ["int"], [
				GoStmt.GoReturn(GoExpr.GoIndex(GoExpr.GoSelector(GoExpr.GoIdent("self"), "b"), GoExpr.GoIdent("pos")))
			]),
			GoDecl.GoFuncDecl("set", {name: "self", typeName: "*haxe__io__Bytes"}, [{name: "pos", typeName: "int"}, {name: "value", typeName: "int"}], [],
				[
					GoStmt.GoAssign(GoExpr.GoIndex(GoExpr.GoSelector(GoExpr.GoIdent("self"), "b"), GoExpr.GoIdent("pos")), GoExpr.GoIdent("value"))
				]),
			GoDecl.GoFuncDecl("New_haxe__io__BytesBuffer", null, [], ["*haxe__io__BytesBuffer"],
				[GoStmt.GoReturn(GoExpr.GoRaw("&haxe__io__BytesBuffer{b: []int{}}"))]),
			GoDecl.GoFuncDecl("addByte", {
				name: "self",
				typeName: "*haxe__io__BytesBuffer"
			}, [{name: "value", typeName: "int"}], [], [
				GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("self"), "b"),
					GoExpr.GoCall(GoExpr.GoIdent("append"), [GoExpr.GoSelector(GoExpr.GoIdent("self"), "b"), GoExpr.GoIdent("value")]))
			]),
			GoDecl.GoFuncDecl("add", {
				name: "self",
				typeName: "*haxe__io__BytesBuffer"
			}, [{name: "src", typeName: "*haxe__io__Bytes"}], [], [
				GoStmt.GoIf(GoExpr.GoBinary("==", GoExpr.GoIdent("src"), GoExpr.GoNil), [GoStmt.GoReturn(null)], null),
				GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("self"), "b"), GoExpr.GoRaw("append(self.b, src.b...)"))
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
			])
		];
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
			}, [{name: "key", typeName: "int"}, {name: "value", typeName: "any"}], [], [
				GoStmt.GoAssign(GoExpr.GoIndex(GoExpr.GoSelector(GoExpr.GoIdent("self"), "h"), GoExpr.GoIdent("key")), GoExpr.GoIdent("value"))
			]),
			GoDecl.GoFuncDecl("get", {
				name: "self",
				typeName: "*haxe__ds__IntMap"
			}, [{name: "key", typeName: "int"}], ["any"], [
				GoStmt.GoVarDecl("value", null, GoExpr.GoIndex(GoExpr.GoSelector(GoExpr.GoIdent("self"), "h"), GoExpr.GoIdent("key")), true),
				GoStmt.GoReturn(GoExpr.GoIdent("value"))
			]),
			GoDecl.GoFuncDecl("exists", {
				name: "self",
				typeName: "*haxe__ds__IntMap"
			},
				[{name: "key", typeName: "int"}], ["bool"], [GoStmt.GoRaw("_, ok := self.h[key]"), GoStmt.GoReturn(GoExpr.GoIdent("ok"))]),
			GoDecl.GoFuncDecl("remove", {name: "self", typeName: "*haxe__ds__IntMap"}, [{name: "key", typeName: "int"}], ["bool"],
				[
					GoStmt.GoRaw("_, ok := self.h[key]"),
					GoStmt.GoRaw("delete(self.h, key)"),
					GoStmt.GoReturn(GoExpr.GoIdent("ok"))
				]),
			GoDecl.GoFuncDecl("New_haxe__ds__StringMap", null, [], ["*haxe__ds__StringMap"],
				[GoStmt.GoReturn(GoExpr.GoRaw("&haxe__ds__StringMap{h: map[string]any{}}"))]),
			GoDecl.GoFuncDecl("set", {
				name: "self",
				typeName: "*haxe__ds__StringMap"
			}, [{name: "key", typeName: "*string"}, {name: "value", typeName: "any"}], [], [
				GoStmt.GoAssign(GoExpr.GoRaw("self.h[*hxrt.StdString(key)]"), GoExpr.GoIdent("value"))
			]),
			GoDecl.GoFuncDecl("get", {name: "self", typeName: "*haxe__ds__StringMap"}, [{name: "key", typeName: "*string"}], ["any"], [
				GoStmt.GoRaw("value := self.h[*hxrt.StdString(key)]"),
				GoStmt.GoReturn(GoExpr.GoIdent("value"))
			]),
			GoDecl.GoFuncDecl("exists", {name: "self", typeName: "*haxe__ds__StringMap"}, [{name: "key", typeName: "*string"}], ["bool"], [
				GoStmt.GoRaw("_, ok := self.h[*hxrt.StdString(key)]"),
				GoStmt.GoReturn(GoExpr.GoIdent("ok"))
			]),
			GoDecl.GoFuncDecl("remove", {name: "self", typeName: "*haxe__ds__StringMap"}, [{name: "key", typeName: "*string"}], ["bool"],
				[
					GoStmt.GoRaw("_, ok := self.h[*hxrt.StdString(key)]"),
					GoStmt.GoRaw("delete(self.h, *hxrt.StdString(key))"),
					GoStmt.GoReturn(GoExpr.GoIdent("ok"))
				]),
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
			GoDecl.GoFuncDecl("remove", {name: "self", typeName: "*haxe__ds__ObjectMap"}, [{name: "key", typeName: "any"}], ["bool"],
				[
					GoStmt.GoRaw("_, ok := self.h[key]"),
					GoStmt.GoRaw("delete(self.h, key)"),
					GoStmt.GoReturn(GoExpr.GoIdent("ok"))
				]),
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
			GoDecl.GoFuncDecl("remove", {name: "self", typeName: "*haxe__ds__EnumValueMap"}, [{name: "key", typeName: "any"}], ["bool"],
				[
					GoStmt.GoRaw("_, ok := self.h[key]"),
					GoStmt.GoRaw("delete(self.h, key)"),
					GoStmt.GoReturn(GoExpr.GoIdent("ok"))
				]),
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
				GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("self"), "add"), [GoExpr.GoIdent("item")]))
			]),
			GoDecl.GoFuncDecl("pop", {name: "self", typeName: "*haxe__ds__List"}, [], ["any"], [
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
				GoStmt.GoVarDecl("methodOverride", "*string", GoExpr.GoNil, false),
				GoStmt.GoRaw("if len(rest) >= 2 {"),
				GoStmt.GoRaw("\tswitch candidate := rest[1].(type) {"),
				GoStmt.GoRaw("\tcase *string:"),
				GoStmt.GoRaw("\t\tmethodOverride = candidate"),
				GoStmt.GoRaw("\tcase string:"),
				GoStmt.GoRaw("\t\tmethodOverride = hxrt.StringFromLiteral(candidate)"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("}"),
				GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("self"), "hxrt__http__requestWith"),
					[GoExpr.GoIdent("post"), GoExpr.GoIdent("methodOverride"), GoExpr.GoIdent("api")]))
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
					[GoExpr.GoIdent("isPost"), GoExpr.GoNil, GoExpr.GoNil]))
			]),
			GoDecl.GoFuncDecl("hxrt__http__requestWith", {
				name: "self",
				typeName: "*sys__Http"
			}, [
				{name: "post", typeName: "bool"},
				{name: "methodOverride", typeName: "*string"},
				{name: "api", typeName: "any"}
			], [], [
				GoStmt.GoAssign(GoExpr.GoIdent("_"), GoExpr.GoIdent("api")),
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
				GoStmt.GoRaw("if proxyURL := hxrt__http__proxyURL(); proxyURL != nil {"),
				GoStmt.GoRaw("\ttransport.Proxy = http.ProxyURL(proxyURL)"),
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
							GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoRaw("response.Status")])
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
			GoDecl.GoStructDecl("sys__io__ProcessOutput", [{name: "scanner", typeName: "*bufio.Scanner"}]),
			GoDecl.GoStructDecl("sys__io__Process", [
				{name: "cmd", typeName: "*exec.Cmd"},
				{name: "stdout", typeName: "*sys__io__ProcessOutput"}
			]),
			GoDecl.GoFuncDecl("Sys_getCwd", null, [], ["*string"], [
				GoStmt.GoRaw("cwd, err := os.Getwd()"),
				GoStmt.GoIf(GoExpr.GoBinary("!=", GoExpr.GoIdent("err"), GoExpr.GoNil), [
					GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoStringLiteral("")]))
				],
					null),
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoIdent("cwd")]))
			]),
			GoDecl.GoFuncDecl("Sys_args", null, [], ["[]*string"], [
				GoStmt.GoRaw("args := os.Args"),
				GoStmt.GoIf(GoExpr.GoBinary("<=", GoExpr.GoCall(GoExpr.GoIdent("len"), [GoExpr.GoIdent("args")]), GoExpr.GoIntLiteral(1)),
					[GoStmt.GoReturn(GoExpr.GoRaw("[]*string{}"))], null),
				GoStmt.GoRaw("out := make([]*string, 0, len(args)-1)"),
				GoStmt.GoRaw("for _, arg := range args[1:] {"),
				GoStmt.GoRaw("\tout = append(out, hxrt.StringFromLiteral(arg))"),
				GoStmt.GoRaw("}"),
				GoStmt.GoReturn(GoExpr.GoIdent("out"))
			]),
			GoDecl.GoFuncDecl("sys__io__File_saveContent", null, [
				{
					name: "path",
					typeName: "*string"
				},
				{name: "content", typeName: "*string"}
			], [], [
				GoStmt.GoAssign(GoExpr.GoIdent("_"), GoExpr.GoRaw("os.WriteFile(*hxrt.StdString(path), []byte(*hxrt.StdString(content)), 0o644)"))
			]),
			GoDecl.GoFuncDecl("sys__io__File_getContent", null, [
				{
					name: "path",
					typeName: "*string"
				}
			], ["*string"], [
				GoStmt.GoRaw("raw, err := os.ReadFile(*hxrt.StdString(path))"),
				GoStmt.GoIf(GoExpr.GoBinary("!=", GoExpr.GoIdent("err"), GoExpr.GoNil), [
					GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoStringLiteral("")]))
				],
					null),
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoRaw("string(raw)")]))
			]),
			GoDecl.GoFuncDecl("New_sys__io__Process", null, [
				{
					name: "command",
					typeName: "*string"
				},
				{name: "args", typeName: "[]*string"}
			], ["*sys__io__Process"], [
				GoStmt.GoRaw("cmd := exec.Command(*hxrt.StdString(command), hxrt.StringSlice(args)...)"),
				GoStmt.GoRaw("stdoutPipe, _ := cmd.StdoutPipe()"),
				GoStmt.GoRaw("_ = cmd.Start()"),
				GoStmt.GoRaw("scanner := bufio.NewScanner(stdoutPipe)"),
				GoStmt.GoReturn(GoExpr.GoRaw("&sys__io__Process{cmd: cmd, stdout: &sys__io__ProcessOutput{scanner: scanner}}"))
			]),
			GoDecl.GoFuncDecl("readLine", {
				name: "self",
				typeName: "*sys__io__ProcessOutput"
			}, [], ["*string"], [
				GoStmt.GoIf(GoExpr.GoRaw("self == nil || self.scanner == nil"), [
					GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoStringLiteral("")]))
				], null),
				GoStmt.GoIf(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoSelector(GoExpr.GoIdent("self"), "scanner"), "Scan"), []), [
					GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [
						GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoSelector(GoExpr.GoIdent("self"), "scanner"), "Text"), [])
					]))
				], null),
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoStringLiteral("")]))
			]),
			GoDecl.GoFuncDecl("close", {
				name: "self",
				typeName: "*sys__io__Process"
			}, [], [], [
				GoStmt.GoIf(GoExpr.GoRaw("self == nil || self.cmd == nil"), [GoStmt.GoReturn(null)], null),
				GoStmt.GoRaw("if self.cmd.Process != nil {"),
				GoStmt.GoRaw("\t_ = self.cmd.Process.Kill()"),
				GoStmt.GoRaw("}"),
				GoStmt.GoAssign(GoExpr.GoIdent("_"), GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoSelector(GoExpr.GoIdent("self"), "cmd"), "Wait"), []))
			])
		];
	}

	function lowerStdlibSymbolShimDecls():Array<GoDecl> {
		return [
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
				GoStmt.GoRaw("raw := make([]byte, len(value.b))"),
				GoStmt.GoRaw("for i := 0; i < len(value.b); i++ {"),
				GoStmt.GoRaw("\traw[i] = byte(value.b[i])"),
				GoStmt.GoRaw("}"),
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
				GoStmt.GoReturn(GoExpr.GoRaw("&haxe__io__Bytes{b: converted, length: len(converted)}"))
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
			GoDecl.GoStructDecl("haxe__io__BytesInput", []),
			GoDecl.GoStructDecl("haxe__io__BytesOutput", []),
			GoDecl.GoStructDecl("haxe__io__Eof", []),
			GoDecl.GoStructDecl("haxe__io__Error", []),
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
	}

	function lowerRegexSerializerShimDecls():Array<GoDecl> {
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
				GoStmt.GoRaw("if strings.Contains(rawOptions, \"i\") {"),
				GoStmt.GoRaw("\trawPattern = \"(?i)\" + rawPattern"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("compiled, err := regexp.Compile(rawPattern)"),
				GoStmt.GoIf(GoExpr.GoBinary("!=", GoExpr.GoIdent("err"), GoExpr.GoNil), [
					GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("hxrt.Throw"), [GoExpr.GoIdent("err")])),
					GoStmt.GoAssign(GoExpr.GoIdent("compiled"),
						GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("regexp"), "MustCompile"), [GoExpr.GoStringLiteral("a^")]))
				],
					null),
				GoStmt.GoReturn(GoExpr.GoRaw("&EReg{regex: compiled, global: strings.Contains(rawOptions, \"g\"), lastSource: hxrt.StringFromLiteral(\"\"), lastIndices: nil}"))
			]),
			GoDecl.GoFuncDecl("match", {
				name: "self",
				typeName: "*EReg"
			}, [{name: "source", typeName: "*string"}], ["bool"], [
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("self"), "matchSub"), [GoExpr.GoIdent("source"), GoExpr.GoIntLiteral(0)]))
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
				GoStmt.GoRaw("\tself.lastSource = hxrt.StringFromLiteral(raw)"),
				GoStmt.GoRaw("\tself.lastIndices = nil"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("found := self.regex.FindStringSubmatchIndex(raw[pos:])"),
				GoStmt.GoRaw("if found == nil {"),
				GoStmt.GoRaw("\tself.lastSource = hxrt.StringFromLiteral(raw)"),
				GoStmt.GoRaw("\tself.lastIndices = nil"),
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
				GoStmt.GoRaw("if self == nil || self.lastSource == nil || index < 0 {"),
				GoStmt.GoRaw("\treturn hxrt.StringFromLiteral(\"\")"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("offset := index * 2"),
				GoStmt.GoRaw("if offset+1 >= len(self.lastIndices) {"),
				GoStmt.GoRaw("\treturn hxrt.StringFromLiteral(\"\")"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("start := self.lastIndices[offset]"),
				GoStmt.GoRaw("end := self.lastIndices[offset+1]"),
				GoStmt.GoRaw("if start < 0 || end < start {"),
				GoStmt.GoRaw("\treturn hxrt.StringFromLiteral(\"\")"),
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
				GoStmt.GoRaw("if self == nil || len(self.lastIndices) < 2 || self.lastIndices[0] < 0 || self.lastIndices[1] < self.lastIndices[0] {"),
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
				GoStmt.GoRaw("if self == nil || self.lastSource == nil || len(self.lastIndices) < 2 || self.lastIndices[0] < 0 {"),
				GoStmt.GoRaw("\treturn hxrt.StringFromLiteral(\"\")"),
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
				GoStmt.GoRaw("if self == nil || self.lastSource == nil || len(self.lastIndices) < 2 || self.lastIndices[1] < 0 {"),
				GoStmt.GoRaw("\treturn hxrt.StringFromLiteral(\"\")"),
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
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [
					GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoSelector(GoExpr.GoIdent("self"), "regex"), "ReplaceAllString"),
						[GoExpr.GoRaw("*hxrt.StdString(source)"), GoExpr.GoRaw("*hxrt.StdString(by)")])
				]))
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
				GoStmt.GoRaw("matches := self.regex.FindAllStringSubmatchIndex(raw, -1)"),
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
			GoDecl.GoStructDecl("haxe__Serializer", [
				{
					name: "buf",
					typeName: "*string"
				},
				{name: "useEnumIndex", typeName: "bool"}
			]),
			GoDecl.GoFuncDecl("New_haxe__Serializer", null, [], ["*haxe__Serializer"], [
				GoStmt.GoReturn(GoExpr.GoRaw("&haxe__Serializer{buf: hxrt.StringFromLiteral(\"\"), useEnumIndex: false}"))
			]),
			GoDecl.GoFuncDecl("serialize", {
				name: "self",
				typeName: "*haxe__Serializer"
			}, [{name: "value", typeName: "any"}], [], [
				GoStmt.GoIf(GoExpr.GoBinary("==", GoExpr.GoIdent("self"), GoExpr.GoNil), [GoStmt.GoReturn(null)], null),
				GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("self"), "buf"),
					GoExpr.GoCall(GoExpr.GoIdent("haxe__Serializer_run"), [GoExpr.GoIdent("value")]))
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
			GoDecl.GoFuncDecl("haxe__Serializer_run", null, [
				{
					name: "value",
					typeName: "any"
				}
			], ["*string"], [
				GoStmt.GoRaw("encoded, err := json.Marshal(value)"),
				GoStmt.GoIf(GoExpr.GoBinary("!=", GoExpr.GoIdent("err"), GoExpr.GoNil), [
					GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("hxrt.Throw"), [GoExpr.GoIdent("err")])),
					GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoStringLiteral("")]))
				],
					null),
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoRaw("string(encoded)")]))
			]),
			GoDecl.GoStructDecl("haxe__Unserializer",
				[
					{
						name: "buf",
						typeName: "*string"
					}
				]),
			GoDecl.GoFuncDecl("New_haxe__Unserializer", null, [{name: "buf", typeName: "*string"}], ["*haxe__Unserializer"],
				[GoStmt.GoReturn(GoExpr.GoRaw("&haxe__Unserializer{buf: buf}"))]),
			GoDecl.GoFuncDecl("unserialize", {
				name: "self",
				typeName: "*haxe__Unserializer"
			}, [], ["any"], [
				GoStmt.GoIf(GoExpr.GoRaw("self == nil || self.buf == nil"), [GoStmt.GoReturn(GoExpr.GoNil)], null),
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("haxe__Unserializer_run"), [GoExpr.GoSelector(GoExpr.GoIdent("self"), "buf")]))
			]),
			GoDecl.GoFuncDecl("hxrt_jsonNormalize", null, [
				{
					name: "value",
					typeName: "any"
				}
			], ["any"], [
				GoStmt.GoRaw("switch current := value.(type) {"),
				GoStmt.GoRaw("case map[string]any:"),
				GoStmt.GoRaw("\tout := make(map[string]any, len(current))"),
				GoStmt.GoRaw("\tfor key, item := range current {"),
				GoStmt.GoRaw("\t\tout[key] = hxrt_jsonNormalize(item)"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn out"),
				GoStmt.GoRaw("case []any:"),
				GoStmt.GoRaw("\tout := make([]any, len(current))"),
				GoStmt.GoRaw("\tfor i, item := range current {"),
				GoStmt.GoRaw("\t\tout[i] = hxrt_jsonNormalize(item)"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn out"),
				GoStmt.GoRaw("case float64:"),
				GoStmt.GoRaw("\tasInt := int(current)"),
				GoStmt.GoRaw("\tif float64(asInt) == current {"),
				GoStmt.GoRaw("\t\treturn asInt"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn current"),
				GoStmt.GoRaw("default:"),
				GoStmt.GoRaw("\treturn current"),
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
				GoStmt.GoRaw("var decoded any"),
				GoStmt.GoRaw("if err := json.Unmarshal([]byte(*hxrt.StdString(source)), &decoded); err != nil {"),
				GoStmt.GoRaw("\thxrt.Throw(err)"),
				GoStmt.GoRaw("\treturn nil"),
				GoStmt.GoRaw("}"),
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt_jsonNormalize"), [GoExpr.GoIdent("decoded")]))
			])
		];
	}

	function lowerNetSocketShimDecls():Array<GoDecl> {
		return [
			GoDecl.GoStructDecl("sys__net__Host", [{name: "host", typeName: "*string"}, {name: "ip", typeName: "int"}]),
			GoDecl.GoFuncDecl("New_sys__net__Host", null, [{name: "name", typeName: "*string"}], ["*sys__net__Host"],
				[GoStmt.GoReturn(GoExpr.GoRaw("&sys__net__Host{host: name, ip: 0}"))]),
			GoDecl.GoFuncDecl("toString", {
				name: "self",
				typeName: "*sys__net__Host"
			}, [], ["*string"], [
				GoStmt.GoIf(GoExpr.GoRaw("self == nil || self.host == nil"), [
					GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoStringLiteral("")]))
				], null),
				GoStmt.GoReturn(GoExpr.GoSelector(GoExpr.GoIdent("self"), "host"))
			]),
			GoDecl.GoFuncDecl("reverse", {
				name: "self",
				typeName: "*sys__net__Host"
			}, [], ["*string"], [
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("self"), "toString"), []))
			]),
			GoDecl.GoFuncDecl("sys__net__Host_localhost", null, [], ["*string"], [
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoStringLiteral("localhost")]))
			]),
			GoDecl.GoStructDecl("sys__net__SocketInput", [
				{
					name: "reader",
					typeName: "*bufio.Reader"
				}
			]),
			GoDecl.GoStructDecl("sys__net__SocketOutput", [{name: "writer", typeName: "*bufio.Writer"}]),
			GoDecl.GoStructDecl("sys__net__Socket", [
				{name: "input", typeName: "*sys__net__SocketInput"},
				{name: "output", typeName: "*sys__net__SocketOutput"},
				{name: "custom", typeName: "any"},
				{name: "conn", typeName: "net.Conn"},
				{name: "listener", typeName: "net.Listener"}
			]),
			GoDecl.GoFuncDecl("New_sys__net__Socket", null, [], ["*sys__net__Socket"], [
				GoStmt.GoReturn(GoExpr.GoRaw("&sys__net__Socket{input: &sys__net__SocketInput{}, output: &sys__net__SocketOutput{}}"))
			]),
			GoDecl.GoFuncDecl("hxrt__socket_setConn", {
				name: "self",
				typeName: "*sys__net__Socket"
			}, [{name: "conn", typeName: "net.Conn"}], [], [
				GoStmt.GoIf(GoExpr.GoRaw("self == nil || conn == nil"), [GoStmt.GoReturn(null)], null),
				GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("self"), "conn"), GoExpr.GoIdent("conn")),
				GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("self"), "input"), GoExpr.GoRaw("&sys__net__SocketInput{reader: bufio.NewReader(conn)}")),
				GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("self"), "output"), GoExpr.GoRaw("&sys__net__SocketOutput{writer: bufio.NewWriter(conn)}"))
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
				GoStmt.GoIf(GoExpr.GoRaw("self == nil || host == nil || host.host == nil"), [
					GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("hxrt.Throw"), [
						GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoStringLiteral("socket connect requires host")])
					])),
					GoStmt.GoReturn(null)
				], null),
				GoStmt.GoVarDecl("address", null, GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("net"), "JoinHostPort"), [
					GoExpr.GoRaw("*hxrt.StdString(host.host)"),
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
				GoStmt.GoIf(GoExpr.GoRaw("self == nil || host == nil || host.host == nil"), [
					GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("hxrt.Throw"), [
						GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoStringLiteral("socket bind requires host")])
					])),
					GoStmt.GoReturn(null)
				], null),
				GoStmt.GoVarDecl("address", null, GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("net"), "JoinHostPort"), [
					GoExpr.GoRaw("*hxrt.StdString(host.host)"),
					GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("strconv"), "Itoa"), [GoExpr.GoIdent("port")])
				]), true),
				GoStmt.GoRaw("listener, err := net.Listen(\"tcp\", address)"),
				GoStmt.GoIf(GoExpr.GoBinary("!=", GoExpr.GoIdent("err"), GoExpr.GoNil), [
					GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("hxrt.Throw"), [GoExpr.GoIdent("err")])),
					GoStmt.GoReturn(null)
				],
					null),
				GoStmt.GoIf(GoExpr.GoRaw("self.listener != nil"), [GoStmt.GoRaw("_ = self.listener.Close()")], null),
				GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("self"), "listener"), GoExpr.GoIdent("listener"))
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
				], null),
				GoStmt.GoRaw("conn, err := self.listener.Accept()"),
				GoStmt.GoIf(GoExpr.GoBinary("!=", GoExpr.GoIdent("err"), GoExpr.GoNil), [
					GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("hxrt.Throw"), [GoExpr.GoIdent("err")])),
					GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("New_sys__net__Socket"), []))
				],
					null),
				GoStmt.GoVarDecl("accepted", null, GoExpr.GoCall(GoExpr.GoIdent("New_sys__net__Socket"), []), true),
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
			],
				[], [GoStmt.GoRaw("_ = read"), GoStmt.GoRaw("_ = write")]),
			GoDecl.GoFuncDecl("hxrt__socket_addrInfo", null, [{name: "addr", typeName: "net.Addr"}], ["map[string]any"], [
				GoStmt.GoRaw("if addr == nil {"),
				GoStmt.GoRaw("\treturn map[string]any{\"host\": New_sys__net__Host(hxrt.StringFromLiteral(\"\")), \"port\": 0}"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("rawHost := \"\""),
				GoStmt.GoRaw("rawPort := \"0\""),
				GoStmt.GoRaw("hostPart, portPart, err := net.SplitHostPort(addr.String())"),
				GoStmt.GoRaw("if err == nil {"),
				GoStmt.GoRaw("\trawHost = hostPart"),
				GoStmt.GoRaw("\trawPort = portPart"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("port, _ := strconv.Atoi(rawPort)"),
				GoStmt.GoReturn(GoExpr.GoRaw("map[string]any{\"host\": New_sys__net__Host(hxrt.StringFromLiteral(rawHost)), \"port\": port}"))
			]),
			GoDecl.GoFuncDecl("peer", {
				name: "self",
				typeName: "*sys__net__Socket"
			}, [], ["map[string]any"], [
				GoStmt.GoIf(GoExpr.GoRaw("self == nil || self.conn == nil"), [
					GoStmt.GoReturn(GoExpr.GoRaw("map[string]any{\"host\": New_sys__net__Host(hxrt.StringFromLiteral(\"\")), \"port\": 0}"))
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
					GoStmt.GoReturn(GoExpr.GoRaw("map[string]any{\"host\": New_sys__net__Host(hxrt.StringFromLiteral(\"\")), \"port\": 0}"))
				], null),
				GoStmt.GoIf(GoExpr.GoRaw("self.conn != nil"), [
					GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt__socket_addrInfo"), [
						GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoSelector(GoExpr.GoIdent("self"), "conn"), "LocalAddr"), [])
					]))
				], null),
				GoStmt.GoIf(GoExpr.GoRaw("self.listener != nil"), [
					GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt__socket_addrInfo"),
						[
							GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoSelector(GoExpr.GoIdent("self"), "listener"), "Addr"), [])
						]))
				],
					null),
				GoStmt.GoReturn(GoExpr.GoRaw("map[string]any{\"host\": New_sys__net__Host(hxrt.StringFromLiteral(\"\")), \"port\": 0}"))
			]),
			GoDecl.GoFuncDecl("setTimeout", {
				name: "self",
				typeName: "*sys__net__Socket"
			}, [{name: "timeout", typeName: "float64"}], [],
				[GoStmt.GoRaw("_ = timeout")]),
			GoDecl.GoFuncDecl("waitForRead", {
				name: "self",
				typeName: "*sys__net__Socket"
			}, [], [], []),
			GoDecl.GoFuncDecl("setBlocking", {
				name: "self",
				typeName: "*sys__net__Socket"
			}, [{name: "b", typeName: "bool"}], [], [GoStmt.GoRaw("_ = b")]),
			GoDecl.GoFuncDecl("setFastSend", {
				name: "self",
				typeName: "*sys__net__Socket"
			}, [{name: "b", typeName: "bool"}], [],
				[GoStmt.GoRaw("_ = b")]),
			GoDecl.GoFuncDecl("sys__net__Socket_select", null, [
				{
					name: "read",
					typeName: "[]*sys__net__Socket"
				},
				{name: "write", typeName: "[]*sys__net__Socket"},
				{name: "others", typeName: "[]*sys__net__Socket"},
				{name: "timeout", typeName: "...float64"}
			], ["map[string]any"], [
				GoStmt.GoRaw("_ = timeout"),
				GoStmt.GoIf(GoExpr.GoBinary("==", GoExpr.GoIdent("read"), GoExpr.GoNil),
					[GoStmt.GoAssign(GoExpr.GoIdent("read"), GoExpr.GoRaw("[]*sys__net__Socket{}"))], null),
				GoStmt.GoIf(GoExpr.GoBinary("==", GoExpr.GoIdent("write"), GoExpr.GoNil),
					[GoStmt.GoAssign(GoExpr.GoIdent("write"), GoExpr.GoRaw("[]*sys__net__Socket{}"))], null),
				GoStmt.GoIf(GoExpr.GoBinary("==", GoExpr.GoIdent("others"), GoExpr.GoNil),
					[GoStmt.GoAssign(GoExpr.GoIdent("others"), GoExpr.GoRaw("[]*sys__net__Socket{}"))], null),
				GoStmt.GoReturn(GoExpr.GoRaw("map[string]any{\"read\": read, \"write\": write, \"others\": others}"))
			]),
			GoDecl.GoFuncDecl("readLine", {
				name: "self",
				typeName: "*sys__net__SocketInput"
			}, [], ["*string"], [
				GoStmt.GoRaw("if self == nil || self.reader == nil {"),
				GoStmt.GoRaw("\treturn hxrt.StringFromLiteral(\"\")"),
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
				GoStmt.GoRaw("if self == nil || self.writer == nil {"),
				GoStmt.GoRaw("\treturn"),
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
						decls.push(lowerFunctionDecl(symbol, func, null));
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

	function lowerFunctionDecl(name:String, func:TFunc, receiver:Null<GoParam>):GoDecl {
		pushFunctionVarNameScope();
		var params = lowerFunctionParams(func);
		var results = lowerFunctionResults(func.t);
		pushFunctionReturnType(func.t);
		var body = lowerFunctionBody(func.expr);
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
		body = body.concat(loweredCtorBody.body);
		body.push(GoStmt.GoReturn(GoExpr.GoIdent("self")));
		popFunctionVarNameScope();
		return GoDecl.GoFuncDecl(constructorSymbol(classType), null, params, ["*" + typeName], body);
	}

	function lowerInstanceMethodDecl(classType:ClassType, fieldName:String, func:TFunc):GoDecl {
		return lowerFunctionDecl(normalizeIdent(fieldName), func, {
			name: "self",
			typeName: "*" + classTypeName(classType)
		});
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

				var lowered = value == null ? null : lowerExprWithPrefix(value);
				var prefix = lowered == null ? [] : lowered.prefix;
				var loweredValue = lowered == null ? null : lowered.expr;
				if (value != null && loweredValue != null) {
					loweredValue = upcastIfNeeded(loweredValue, value.t, variable.t);
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
						var rightExpr = upcastIfNeeded(loweredRight.expr, right.t, left.t);
						var assignStmt = GoStmt.GoAssign(lowerLValue(left), rightExpr);
						if (loweredRight.prefix.length > 0) {
							loweredRight.prefix.concat([assignStmt]);
						} else {
							[assignStmt];
						}
					case OpAssignOp(assignOp):
						var loweredRight = lowerExprWithPrefix(right);
						var rightExpr = upcastIfNeeded(loweredRight.expr, right.t, left.t);
						var targetExpr = lowerLValue(left);
						var assignExpr = lowerAssignOpExpr(assignOp, targetExpr, rightExpr, left.t, right.t);
						var assignStmt = GoStmt.GoAssign(targetExpr, assignExpr);
						if (loweredRight.prefix.length > 0) {
							loweredRight.prefix.concat([assignStmt]);
						} else {
							[assignStmt];
						}
					case _:
						[GoStmt.GoExprStmt(lowerExpr(expr).expr)];
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
						[GoStmt.GoAssign(target, GoExpr.GoBinary("+", target, GoExpr.GoIntLiteral(1)))];
					case OpDecrement:
						var target = lowerLValue(value);
						[GoStmt.GoAssign(target, GoExpr.GoBinary("-", target, GoExpr.GoIntLiteral(1)))];
					case _:
						[GoStmt.GoExprStmt(lowerExpr(expr).expr)];
				}
			case TSwitch(value, cases, defaultExpr):
				[lowerSwitchStmt(value, cases, defaultExpr)];
			case TThrow(value):
				[
					GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("hxrt.Throw"), [lowerExpr(value).expr]))
				];
			case TTry(tryExpr, catches):
				[lowerTryCatchStmt(tryExpr, catches)];
			case TReturn(value):
				if (value == null) {
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
						var appendArgs = [targetExpr].concat([for (arg in args) lowerExpr(arg).expr]);
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
						[GoStmt.GoExprStmt(lowerCall(callee, args, expr.t).expr)];
					}
				}
			case _:
				[GoStmt.GoExprStmt(lowerExpr(expr).expr)];
		};
	}

	function lowerFunctionBody(expr:TypedExpr):Array<GoStmt> {
		pushLocalScope();
		var out = lowerToStatements(expr);
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

		var prefix = [GoStmt.GoVarDecl(temp, typeToGoType(resultType), null, false)].concat(loweredCondition.prefix);

		prefix.push(GoStmt.GoIf(loweredCondition.expr, loweredThen.prefix.concat([GoStmt.GoAssign(GoExpr.GoIdent(temp), loweredThenValue)]),
			loweredElse.prefix.concat([GoStmt.GoAssign(GoExpr.GoIdent(temp), loweredElseValue)])));

		return {
			prefix: prefix,
			expr: GoExpr.GoIdent(temp),
			isStringLike: isStringType(resultType)
		};
	}

	function lowerTryCatchStmt(tryExpr:TypedExpr, catches:Array<{v:TVar, expr:TypedExpr}>):GoStmt {
		if (catches.length == 0) {
			var tryBody = lowerToStatements(tryExpr);
			return GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoFuncLiteral([], [], tryBody), []));
		}

		var caughtName = freshTempName("hx_caught");
		var typeBindingName = freshTempName("hx_typed");
		var typedCases = new Array<GoTypeSwitchCase>();
		var dynamicBody:Null<Array<GoStmt>> = null;

		for (index in 0...catches.length) {
			var catchEntry = catches[index];
			var catchVarName = localVarName(catchEntry.v);
			var catchType = typeToGoType(catchEntry.v.t);
			var catchExprBody = lowerToStatements(catchEntry.expr);
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

		return GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("hxrt.TryCatch"), [
			GoExpr.GoFuncLiteral([], [], lowerToStatements(tryExpr)),
			GoExpr.GoFuncLiteral([
				{
					name: caughtName,
					typeName: "any"
				}
			], [], catchBody)
		]));
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
		localRestIteratorScopes.push([]);
	}

	function popLocalScope():Void {
		if (localFunctionScopes.length > 0) {
			localFunctionScopes.pop();
		}
		if (localRestIteratorScopes.length > 0) {
			localRestIteratorScopes.pop();
		}
	}

	function pushFunctionVarNameScope():Void {
		functionVarNameScopes.push(new Map<Int, String>());
		functionVarNameCountScopes.push(new Map<String, Int>());
	}

	function popFunctionVarNameScope():Void {
		if (functionVarNameScopes.length > 0) {
			functionVarNameScopes.pop();
		}
		if (functionVarNameCountScopes.length > 0) {
			functionVarNameCountScopes.pop();
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
				} else if (exprs.length == 0) {
					{expr: GoExpr.GoNil, isStringLike: false};
				} else {
					lowerExpr(exprs[exprs.length - 1]);
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
				{
					expr: GoExpr.GoCall(GoExpr.GoIdent(constructorSymbol(classRef.get())), [for (arg in args) lowerExpr(arg).expr]),
					isStringLike: false
				};
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
				{
					expr: GoExpr.GoIdent(localVarName(variable)),
					isStringLike: isStringType(variable.t)
				};
			case TParenthesis(inner):
				lowerExpr(inner);
			case TMeta(_, inner):
				lowerExpr(inner);
			case TCast(inner, _):
				lowerExpr(inner);
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
			case TBinop(op, left, right):
				switch (op) {
					case OpAssignOp(assignOp):
						var targetExpr = lowerLValue(left);
						var loweredRight = lowerExprWithPrefix(right);
						var rightExpr = upcastIfNeeded(loweredRight.expr, right.t, left.t);
						var assignExpr = lowerAssignOpExpr(assignOp, targetExpr, rightExpr, left.t, right.t);
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
									GoStmt.GoAssign(target, GoExpr.GoBinary(opSymbol, target, GoExpr.GoIntLiteral(1))),
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
								GoStmt.GoAssign(target, GoExpr.GoBinary(opSymbol, target, GoExpr.GoIntLiteral(1))),
								GoStmt.GoReturn(target)
							]), []),
							isStringLike: isStringType(expr.t)
						};
					case _:
						{
							expr: GoExpr.GoUnary(unopSymbol(op), lowerExpr(value).expr),
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
									GoStmt.GoAssign(target, GoExpr.GoBinary(opSymbol, target, GoExpr.GoIntLiteral(1)))
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
				lowerExprWithPrefix(inner);
			case _:
				var lowered = lowerExpr(expr);
				{
					prefix: [],
					expr: lowered.expr,
					isStringLike: lowered.isStringLike
				};
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
			case _:
				Context.fatalError("Unsupported constant: " + Std.string(constant), Context.currentPos());
				{expr: GoExpr.GoNil, isStringLike: false};
		};
	}

	function lowerField(target:TypedExpr, access:FieldAccess):LoweredExpr {
		return switch (access) {
			case FStatic(classRef, field):
				var resolved = field.get();
				{
					expr: GoExpr.GoIdent(staticSymbol(classRef.get(), resolved.name)),
					isStringLike: isStringType(resolved.type)
				};
			case FInstance(classRef, _, field):
				var resolved = field.get();
				var classType = classRef.get();
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
					var anonField = GoExpr.GoIndex(loweredTarget, GoExpr.GoStringLiteral(resolved.name));
					var anonFieldType = scalarGoType(resolved.type);
					return {
						expr: anonFieldType == "any" ? anonField : GoExpr.GoTypeAssert(anonField, anonFieldType),
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

	function lowerCall(callee:TypedExpr, args:Array<TypedExpr>, returnType:Type):LoweredExpr {
		if (isStaticCall(callee, "Json", ["haxe"], "parse")) {
			var arg = args.length > 0 ? lowerExpr(args[0]).expr : GoExpr.GoNil;
			return {
				expr: GoExpr.GoCall(GoExpr.GoIdent("hxrt.JsonParse"), [arg]),
				isStringLike: false
			};
		}

		if (isStaticCall(callee, "Json", ["haxe"], "stringify")) {
			var arg = args.length > 0 ? lowerExpr(args[0]).expr : GoExpr.GoNil;
			return {
				expr: GoExpr.GoCall(GoExpr.GoIdent("hxrt.JsonStringify"), [arg]),
				isStringLike: true
			};
		}

		if (isStaticCall(callee, "JsonPrinter", ["haxe", "format"], "print")) {
			var arg = args.length > 0 ? lowerExpr(args[0]).expr : GoExpr.GoNil;
			return {
				expr: GoExpr.GoCall(GoExpr.GoIdent("hxrt.JsonStringify"), [arg]),
				isStringLike: true
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
			var loweredArg = lowerExpr(arg).expr;
			var paramType = callParamType(callee.t, index);
			if (paramType != null) {
				loweredArg = upcastIfNeeded(loweredArg, arg.t, paramType);
			}
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
		if (shouldAssertGenericCallResult(callee, returnType)) {
			var expectedType = typeToGoType(returnType);
			if (expectedType != "any") {
				callExpr = GoExpr.GoTypeAssert(callExpr, expectedType);
			}
		}

		return {
			expr: callExpr,
			isStringLike: isStringType(returnType)
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

		Context.fatalError("Unsupported Std.isOfType target type: " + typeToGoType(targetType), valueTypedExpr.pos);
		return GoExpr.GoBoolLiteral(false);
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
		var bindingName = "hx_type";
		return GoExpr.GoCall(GoExpr.GoFuncLiteral([{name: "hx_value", typeName: "any"}], ["bool"], [
			GoStmt.GoTypeSwitch(GoExpr.GoIdent("hx_value"), bindingName, [
				for (typeName in typeNames)
					{
						typeName: typeName,
						body: [
							GoStmt.GoAssign(GoExpr.GoIdent("_"), GoExpr.GoIdent(bindingName)),
							GoStmt.GoReturn(GoExpr.GoBoolLiteral(true))
						]
					}
			], [
				GoStmt.GoAssign(GoExpr.GoIdent("_"), GoExpr.GoIdent(bindingName)),
				GoStmt.GoReturn(GoExpr.GoBoolLiteral(false))
			])
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

	function shouldAssertGenericCallResult(callee:TypedExpr, returnType:Type):Bool {
		if (typeToGoType(returnType) == "any") {
			return false;
		}

		return switch (callee.expr) {
			case TField(_, FInstance(classRef, _, field)):
				var classType = classRef.get();
				var pack = classType.pack.join(".");
				var fieldName = field.get().name;
				if (pack == "haxe.ds") {
					if ((classType.name == "IntMap" || classType.name == "StringMap" || classType.name == "ObjectMap" || classType.name == "EnumValueMap")
						&& fieldName == "get") {
						true;
					} else if (classType.name == "List" && (fieldName == "pop" || fieldName == "first" || fieldName == "last")) {
						true;
					} else {
						false;
					}
				} else {
					false;
				}
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
		var typedStringOps = compilationContext.profile != GoProfile.Portable && isStringType(left.t) && isStringType(right.t);

		return switch (op) {
			case OpAdd if (stringMode):
				{
					expr: GoExpr.GoCall(GoExpr.GoIdent(typedStringOps ? "hxrt.StringConcatStringPtr" : "hxrt.StringConcatAny"),
						[leftLowered.expr, rightLowered.expr]),
					isStringLike: true
				};
			case OpEq if (stringMode):
				{
					expr: GoExpr.GoCall(GoExpr.GoIdent(typedStringOps ? "hxrt.StringEqualStringPtr" : "hxrt.StringEqualAny"),
						[leftLowered.expr, rightLowered.expr]),
					isStringLike: false
				};
			case OpNotEq if (stringMode):
				{
					expr: GoExpr.GoUnary("!",
						GoExpr.GoCall(GoExpr.GoIdent(typedStringOps ? "hxrt.StringEqualStringPtr" : "hxrt.StringEqualAny"),
							[leftLowered.expr, rightLowered.expr])),
					isStringLike: false
				};
			case OpUShr:
				{
					expr: GoExpr.GoCall(GoExpr.GoIdent("int"), [
						GoExpr.GoBinary(">>", GoExpr.GoCall(GoExpr.GoIdent("uint32"), [leftLowered.expr]),
							GoExpr.GoCall(GoExpr.GoIdent("uint"), [rightLowered.expr]))
					]),
					isStringLike: false
				};
			case _:
				{
					expr: GoExpr.GoBinary(binopSymbol(op), leftLowered.expr, rightLowered.expr),
					isStringLike: isStringType(resultType)
				};
		};
	}

	function lowerAssignOpExpr(op:Binop, leftExpr:GoExpr, rightExpr:GoExpr, leftType:Type, rightType:Type):GoExpr {
		if (op == OpAdd && (isStringType(leftType) || isStringType(rightType))) {
			var typedStringOps = compilationContext.profile != GoProfile.Portable && isStringType(leftType) && isStringType(rightType);
			return GoExpr.GoCall(GoExpr.GoIdent(typedStringOps ? "hxrt.StringConcatStringPtr" : "hxrt.StringConcatAny"), [leftExpr, rightExpr]);
		}
		if (op == OpUShr) {
			return GoExpr.GoCall(GoExpr.GoIdent("int"), [
				GoExpr.GoBinary(">>", GoExpr.GoCall(GoExpr.GoIdent("uint32"), [leftExpr]), GoExpr.GoCall(GoExpr.GoIdent("uint"), [rightExpr]))
			]);
		}
		return GoExpr.GoBinary(binopSymbol(op), leftExpr, rightExpr);
	}

	function binopSymbol(op:Binop):String {
		return switch (op) {
			case OpAdd: "+";
			case OpMult: "*";
			case OpDiv: "/";
			case OpSub: "-";
			case OpMod: "%";
			case OpEq: "==";
			case OpNotEq: "!=";
			case OpGt: ">";
			case OpGte: ">=";
			case OpLt: "<";
			case OpLte: "<=";
			case OpBoolAnd: "&&";
			case OpBoolOr: "||";
			case OpAnd: "&";
			case OpOr: "|";
			case OpXor: "^";
			case OpShl: "<<";
			case OpShr: ">>";
			case OpUShr: ">>";
			case OpAssign:
				Context.fatalError("Assignment is handled at statement level", Context.currentPos());
			case _:
				Context.fatalError("Unsupported binary operator", Context.currentPos());
		};
	}

	function unopSymbol(op:Unop):String {
		return switch (op) {
			case OpNot: "!";
			case OpNeg: "-";
			case OpNegBits: "^";
			case OpIncrement:
				Context.fatalError("Increment operator is not supported yet", Context.currentPos());
			case OpDecrement:
				Context.fatalError("Decrement operator is not supported yet", Context.currentPos());
			case _:
				Context.fatalError("Unsupported unary operator", Context.currentPos());
		};
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
		var restElement = restElementType(type);
		if (restElement != null) {
			return "[]" + scalarGoType(restElement);
		}

		var followed = Context.follow(type);
		return switch (followed) {
			case TFun(args, returnType):
				goFunctionType(args, returnType);
			case TInst(classRef, params):
				var classType = classRef.get();
				if (isTypeParameterClass(classType)) {
					"any";
				} else if (isHaxeExceptionClass(classType)) {
					"*hxrt.ExceptionValue";
				} else if (classType.isInterface) {
					classTypeName(classType);
				} else if (classType.pack.length == 0 && classType.name == "String") {
					"*string";
				} else if (classType.pack.length == 0 && classType.name == "Array" && params.length == 1) {
					"[]" + scalarGoType(params[0]);
				} else {
					"*" + classTypeName(classType);
				}
			case TEnum(enumRef, _):
				"*" + enumTypeName(enumRef.get());
			case TAnonymous(_):
				"map[string]any";
			case TAbstract(abstractRef, _):
				var abstractType = abstractRef.get();
				if (abstractType.pack.length == 0 && abstractType.name == "Int") {
					"int";
				} else if (abstractType.pack.length == 0 && abstractType.name == "Float") {
					"float64";
				} else if (abstractType.pack.length == 0 && abstractType.name == "Bool") {
					"bool";
				} else if (abstractType.pack.length == 0 && abstractType.name == "String") {
					"*string";
				} else {
					"any";
				}
			case _:
				"any";
		};
	}

	function isStringType(type:Type):Bool {
		var followed = Context.follow(type);
		return switch (followed) {
			case TInst(classRef, _): var classType = classRef.get(); classType.pack.length == 0 && classType.name == "String";
			case TAbstract(abstractRef, _): var abstractType = abstractRef.get(); abstractType.pack.length == 0 && abstractType.name == "String";
			case _:
				false;
		};
	}

	function isInterfaceType(type:Type):Bool {
		var followed = Context.follow(type);
		return switch (followed) {
			case TInst(classRef, _):
				classRef.get().isInterface;
			case _:
				false;
		};
	}

	function isBoolType(type:Type):Bool {
		var followed = Context.follow(type);
		return switch (followed) {
			case TAbstract(abstractRef, _): var abstractType = abstractRef.get(); abstractType.pack.length == 0 && abstractType.name == "Bool";
			case _:
				false;
		};
	}

	function isIntType(type:Type):Bool {
		var followed = Context.follow(type);
		return switch (followed) {
			case TAbstract(abstractRef, _): var abstractType = abstractRef.get(); abstractType.pack.length == 0 && abstractType.name == "Int";
			case _:
				false;
		};
	}

	function isFloatType(type:Type):Bool {
		var followed = Context.follow(type);
		return switch (followed) {
			case TAbstract(abstractRef, _): var abstractType = abstractRef.get(); abstractType.pack.length == 0 && abstractType.name == "Float";
			case _:
				false;
		};
	}

	function isAnyLikeType(type:Type):Bool {
		return isDynamicCatchType(type) || typeToGoType(type) == "any";
	}

	function isDefinitelyNonNullableType(type:Type):Bool {
		return isBoolType(type) || isIntType(type) || isFloatType(type);
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

	function isAnonymousObjectType(type:Type):Bool {
		var followed = Context.follow(type);
		return switch (followed) {
			case TAnonymous(_):
				true;
			case _:
				false;
		};
	}

	function isArrayType(type:Type):Bool {
		if (restElementType(type) != null) {
			return true;
		}

		var followed = Context.follow(type);
		return switch (followed) {
			case TInst(classRef, _): var classType = classRef.get(); classType.pack.length == 0 && classType.name == "Array";
			case _:
				false;
		};
	}

	function arrayElementGoType(type:Type):String {
		var restElement = restElementType(type);
		if (restElement != null) {
			return scalarGoType(restElement);
		}

		var followed = Context.follow(type);
		return switch (followed) {
			case TInst(classRef, params):
				var classType = classRef.get();
				if (classType.pack.length == 0 && classType.name == "Array" && params.length == 1) {
					scalarGoType(params[0]);
				} else {
					"any";
				}
			case _:
				"any";
		};
	}

	function scalarGoType(type:Type):String {
		var restElement = restElementType(type);
		if (restElement != null) {
			return "[]" + scalarGoType(restElement);
		}

		var followed = Context.follow(type);
		return switch (followed) {
			case TFun(args, returnType):
				goFunctionType(args, returnType);
			case TInst(classRef, params):
				var classType = classRef.get();
				if (isTypeParameterClass(classType)) {
					"any";
				} else if (isHaxeExceptionClass(classType)) {
					"*hxrt.ExceptionValue";
				} else if (classType.isInterface) {
					classTypeName(classType);
				} else if (classType.pack.length == 0 && classType.name == "String") {
					"*string";
				} else if (classType.pack.length == 0 && classType.name == "Array" && params.length == 1) {
					"[]" + scalarGoType(params[0]);
				} else {
					"*" + classTypeName(classType);
				}
			case TEnum(enumRef, _):
				"*" + enumTypeName(enumRef.get());
			case TAnonymous(_):
				"map[string]any";
			case TAbstract(abstractRef, _):
				var abstractType = abstractRef.get();
				if (abstractType.pack.length == 0 && abstractType.name == "Int") {
					"int";
				} else if (abstractType.pack.length == 0 && abstractType.name == "Float") {
					"float64";
				} else if (abstractType.pack.length == 0 && abstractType.name == "Bool") {
					"bool";
				} else if (abstractType.pack.length == 0 && abstractType.name == "String") {
					"*string";
				} else {
					"any";
				}
			case _:
				"any";
		};
	}

	function goFunctionType(args:Array<{name:String, opt:Bool, t:Type}>, returnType:Type):String {
		var params = [for (arg in args) scalarGoType(arg.t)].join(", ");
		if (isVoidType(returnType)) {
			return "func(" + params + ")";
		}
		return "func(" + params + ") " + scalarGoType(returnType);
	}

	function restElementType(type:Type):Null<Type> {
		var followed = Context.follow(type);
		return switch (followed) {
			case TAbstract(abstractRef, params):
				var abstractType = abstractRef.get();
				if (abstractType.pack.join(".") == "haxe" && abstractType.name == "Rest" && params.length == 1) {
					params[0];
				} else {
					null;
				}
			case TType(typeRef, params):
				var typeDef = typeRef.get();
				if (typeDef.pack.join(".") == "haxe._Rest" && typeDef.name == "NativeRest" && params.length == 1) {
					params[0];
				} else {
					null;
				}
			case _:
				null;
		};
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
		var followed = Context.follow(type);
		return switch (followed) {
			case TAbstract(abstractRef, _): var abstractType = abstractRef.get(); abstractType.pack.length == 0 && abstractType.name == "Void";
			case _:
				false;
		};
	}

	function isDynamicCatchType(type:Type):Bool {
		var followed = Context.follow(type);
		return switch (followed) {
			case TDynamic(_):
				true;
			case TAbstract(abstractRef, _): var abstractType = abstractRef.get(); abstractType.pack.length == 0 && abstractType.name == "Dynamic";
			case _:
				false;
		};
	}

	function isTypeParameterClass(classType:ClassType):Bool {
		return switch (classType.kind) {
			case KTypeParameter(_):
				true;
			case _:
				false;
		};
	}

	function isHaxeExceptionClass(classType:ClassType):Bool {
		return classType.pack.join(".") == "haxe" && classType.name == "Exception";
	}

	function isHaxeExceptionType(type:Type):Bool {
		var followed = Context.follow(type);
		return switch (followed) {
			case TInst(classRef, _):
				isHaxeExceptionClass(classRef.get());
			case _:
				false;
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

	function noteStdlibClass(classType:ClassType):Void {
		var pack = classType.pack.join(".");
		if (pack == "haxe.io") {
			switch (classType.name) {
				case "Bytes", "BytesBuffer", "Input", "Output", "Encoding":
					requireStdlibShimGroup("io");
					return;
				case "BytesInput", "BytesOutput", "Eof", "Error", "Path", "StringInput":
					requireStdlibShimGroup("stdlib_symbols");
					return;
				case _:
			}
		}

		if (pack == "haxe.ds") {
			switch (classType.name) {
				case "IntMap", "StringMap", "ObjectMap", "EnumValueMap", "List":
					requireStdlibShimGroup("ds");
					return;
				case "BalancedTree":
					requireStdlibShimGroup("stdlib_symbols");
					return;
				case _:
			}
		}

		if (pack == "sys" && classType.name == "Http") {
			requireStdlibShimGroup("http");
			return;
		}

		if ((pack == "" && classType.name == "Sys") || (pack == "sys.io" && (classType.name == "File" || classType.name == "Process"))) {
			requireStdlibShimGroup("sys");
			return;
		}

		if ((pack == "haxe" && classType.name == "Json")
			|| (pack == "haxe.format" && (classType.name == "JsonParser" || classType.name == "JsonPrinter"))) {
			requireStdlibShimGroup("json");
			return;
		}

		if (pack == "sys.net" && (classType.name == "Host" || classType.name == "Socket")) {
			requireStdlibShimGroup("net_socket");
			return;
		}

		if ((pack == "" && classType.name == "EReg")
			|| (pack == "haxe" && (classType.name == "Serializer" || classType.name == "Unserializer"))) {
			requireStdlibShimGroup("regex_serializer");
			return;
		}

		if ((pack == ""
			&& (classType.name == "Std" || classType.name == "StringTools" || classType.name == "Date" || classType.name == "Math"
				|| classType.name == "Type" || classType.name == "Reflect" || classType.name == "Xml"))
			|| (pack == "haxe.crypto"
				&& (classType.name == "Base64" || classType.name == "Md5" || classType.name == "Sha1" || classType.name == "Sha224"
					|| classType.name == "Sha256"))
			|| (pack == "haxe.xml" && (classType.name == "Parser" || classType.name == "Printer"))
			|| (pack == "haxe.zip" && (classType.name == "Compress" || classType.name == "Uncompress"))
			|| (pack == "sys" && classType.name == "FileSystem")) {
			requireStdlibShimGroup("stdlib_symbols");
		}
	}

	function noteStdlibEnum(enumType:EnumType):Void {
		var pack = enumType.pack.join(".");
		if ((pack == "haxe.io" && enumType.name == "Error") || (pack == "haxe.ds" && enumType.name == "Option")) {
			requireStdlibShimGroup("stdlib_symbols");
		}
	}

	function classTypeName(classType:ClassType):String {
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
