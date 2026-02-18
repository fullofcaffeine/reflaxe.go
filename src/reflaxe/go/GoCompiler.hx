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
    if (requiredStdlibShimGroups.exists("sys")) {
      imports.push("bufio");
      imports.push("os");
      imports.push("os/exec");
    }
    var mainFile:GoFile = {
      packageName: "main",
      imports: imports,
      decls: decls
    };
    var transformedFile = GoASTTransformer.transform(mainFile, compilationContext);

    return [{
      relativePath: "main.go",
      contents: GoASTPrinter.printFile(transformedFile)
    }];
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
    if (classType.isExtern || classType.isInterface) {
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
      decls.push(GoDecl.GoStructDecl(enumName, [
        {name: "tag", typeName: "int"},
        {name: "params", typeName: "[]any"}
      ]));

      var constructors = [for (field in enumType.constructs) field];
      constructors.sort(function(a, b) return a.index - b.index);

      for (constructor in constructors) {
        var symbol = enumConstructorSymbol(enumType, constructor.name);
        var ctorArgs = enumConstructorArgs(constructor.type);
        if (ctorArgs.length == 0) {
          decls.push(GoDecl.GoGlobalVarDecl(
            symbol,
            "*" + enumName,
            GoExpr.GoRaw("&" + enumName + "{tag: " + constructor.index + "}")
          ));
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

          decls.push(GoDecl.GoFuncDecl(
            symbol,
            null,
            params,
            ["*" + enumName],
            [
              GoStmt.GoVarDecl("enumValue", null, GoExpr.GoRaw("&" + enumName + "{tag: " + constructor.index + "}"), true),
              GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("enumValue"), "params"), GoExpr.GoArrayLiteral("any", payloadExprs)),
              GoStmt.GoReturn(GoExpr.GoIdent("enumValue"))
            ]
          ));
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
    if (requiredStdlibShimGroups.exists("json")) {
      decls = decls.concat(lowerJsonStdlibShimDecls());
    }
    if (requiredStdlibShimGroups.exists("io")) {
      decls = decls.concat(lowerIoStdlibShimDecls());
    }
    if (requiredStdlibShimGroups.exists("ds")) {
      decls = decls.concat(lowerDsStdlibShimDecls());
    }
    if (requiredStdlibShimGroups.exists("sys")) {
      decls = decls.concat(lowerSysStdlibShimDecls());
    }
    return decls;
  }

  function lowerJsonStdlibShimDecls():Array<GoDecl> {
    return [
      GoDecl.GoStructDecl("haxe__Json", []),
      GoDecl.GoStructDecl("haxe__format__JsonParser", [
        {name: "source", typeName: "*string"}
      ]),
      GoDecl.GoFuncDecl(
        "New_haxe__format__JsonParser",
        null,
        [{name: "source", typeName: "*string"}],
        ["*haxe__format__JsonParser"],
        [GoStmt.GoReturn(GoExpr.GoRaw("&haxe__format__JsonParser{source: source}"))]
      ),
      GoDecl.GoFuncDecl(
        "doParse",
        {name: "self", typeName: "*haxe__format__JsonParser"},
        [],
        ["any"],
        [GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.JsonParse"), [GoExpr.GoSelector(GoExpr.GoIdent("self"), "source")]))]
      ),
      GoDecl.GoFuncDecl(
        "haxe__format__JsonPrinter_print",
        null,
        [
          {name: "value", typeName: "any"},
          {name: "rest", typeName: "...any"}
        ],
        ["*string"],
        [GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.JsonStringify"), [GoExpr.GoIdent("value")]))]
      ),
      GoDecl.GoFuncDecl(
        "haxe__Json_parse",
        null,
        [{name: "source", typeName: "*string"}],
        ["any"],
        [GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.JsonParse"), [GoExpr.GoIdent("source")]))]
      ),
      GoDecl.GoFuncDecl(
        "haxe__Json_stringify",
        null,
        [
          {name: "value", typeName: "any"},
          {name: "rest", typeName: "...any"}
        ],
        ["*string"],
        [GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.JsonStringify"), [GoExpr.GoIdent("value")]))]
      )
    ];
  }

  function lowerIoStdlibShimDecls():Array<GoDecl> {
    return [
      GoDecl.GoStructDecl("haxe__io__Encoding", []),
      GoDecl.GoStructDecl("haxe__io__Input", []),
      GoDecl.GoStructDecl("haxe__io__Output", []),
      GoDecl.GoStructDecl("haxe__io__Bytes", [
        {name: "b", typeName: "[]int"},
        {name: "length", typeName: "int"}
      ]),
      GoDecl.GoStructDecl("haxe__io__BytesBuffer", [
        {name: "b", typeName: "[]int"}
      ]),
      GoDecl.GoFuncDecl(
        "New_haxe__io__Input",
        null,
        [],
        ["*haxe__io__Input"],
        [GoStmt.GoReturn(GoExpr.GoRaw("&haxe__io__Input{}"))]
      ),
      GoDecl.GoFuncDecl(
        "New_haxe__io__Output",
        null,
        [],
        ["*haxe__io__Output"],
        [GoStmt.GoReturn(GoExpr.GoRaw("&haxe__io__Output{}"))]
      ),
      GoDecl.GoFuncDecl(
        "New_haxe__io__Bytes",
        null,
        [
          {name: "length", typeName: "int"},
          {name: "b", typeName: "[]int"}
        ],
        ["*haxe__io__Bytes"],
        [
          GoStmt.GoIf(
            GoExpr.GoBinary("==", GoExpr.GoIdent("b"), GoExpr.GoNil),
            [GoStmt.GoAssign(GoExpr.GoIdent("b"), GoExpr.GoRaw("make([]int, length)"))],
            null
          ),
          GoStmt.GoReturn(GoExpr.GoRaw("&haxe__io__Bytes{b: b, length: len(b)}"))
        ]
      ),
      GoDecl.GoFuncDecl(
        "haxe__io__Bytes_alloc",
        null,
        [{name: "length", typeName: "int"}],
        ["*haxe__io__Bytes"],
        [GoStmt.GoReturn(GoExpr.GoRaw("&haxe__io__Bytes{b: make([]int, length), length: length}"))]
      ),
      GoDecl.GoFuncDecl(
        "haxe__io__Bytes_ofString",
        null,
        [
          {name: "value", typeName: "*string"},
          {name: "encoding", typeName: "...*haxe__io__Encoding"}
        ],
        ["*haxe__io__Bytes"],
        [
          GoStmt.GoVarDecl("raw", null, GoExpr.GoCall(GoExpr.GoIdent("hxrt.BytesFromString"), [GoExpr.GoIdent("value")]), true),
          GoStmt.GoReturn(GoExpr.GoRaw("&haxe__io__Bytes{b: raw, length: len(raw)}"))
        ]
      ),
      GoDecl.GoFuncDecl(
        "toString",
        {name: "self", typeName: "*haxe__io__Bytes"},
        [],
        ["*string"],
        [
          GoStmt.GoIf(
            GoExpr.GoBinary("==", GoExpr.GoIdent("self"), GoExpr.GoNil),
            [GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoStringLiteral("")]))],
            null
          ),
          GoStmt.GoReturn(GoExpr.GoCall(
            GoExpr.GoIdent("hxrt.BytesToString"),
            [GoExpr.GoSelector(GoExpr.GoIdent("self"), "b")]
          ))
        ]
      ),
      GoDecl.GoFuncDecl(
        "get",
        {name: "self", typeName: "*haxe__io__Bytes"},
        [{name: "pos", typeName: "int"}],
        ["int"],
        [GoStmt.GoReturn(GoExpr.GoIndex(GoExpr.GoSelector(GoExpr.GoIdent("self"), "b"), GoExpr.GoIdent("pos")))]
      ),
      GoDecl.GoFuncDecl(
        "set",
        {name: "self", typeName: "*haxe__io__Bytes"},
        [
          {name: "pos", typeName: "int"},
          {name: "value", typeName: "int"}
        ],
        [],
        [GoStmt.GoAssign(GoExpr.GoIndex(GoExpr.GoSelector(GoExpr.GoIdent("self"), "b"), GoExpr.GoIdent("pos")), GoExpr.GoIdent("value"))]
      ),
      GoDecl.GoFuncDecl(
        "New_haxe__io__BytesBuffer",
        null,
        [],
        ["*haxe__io__BytesBuffer"],
        [GoStmt.GoReturn(GoExpr.GoRaw("&haxe__io__BytesBuffer{b: []int{}}"))]
      ),
      GoDecl.GoFuncDecl(
        "addByte",
        {name: "self", typeName: "*haxe__io__BytesBuffer"},
        [{name: "value", typeName: "int"}],
        [],
        [GoStmt.GoAssign(
          GoExpr.GoSelector(GoExpr.GoIdent("self"), "b"),
          GoExpr.GoCall(GoExpr.GoIdent("append"), [
            GoExpr.GoSelector(GoExpr.GoIdent("self"), "b"),
            GoExpr.GoIdent("value")
          ])
        )]
      ),
      GoDecl.GoFuncDecl(
        "add",
        {name: "self", typeName: "*haxe__io__BytesBuffer"},
        [{name: "src", typeName: "*haxe__io__Bytes"}],
        [],
        [
          GoStmt.GoIf(
            GoExpr.GoBinary("==", GoExpr.GoIdent("src"), GoExpr.GoNil),
            [GoStmt.GoReturn(null)],
            null
          ),
          GoStmt.GoAssign(
            GoExpr.GoSelector(GoExpr.GoIdent("self"), "b"),
            GoExpr.GoRaw("append(self.b, src.b...)")
          )
        ]
      ),
      GoDecl.GoFuncDecl(
        "addString",
        {name: "self", typeName: "*haxe__io__BytesBuffer"},
        [
          {name: "value", typeName: "*string"},
          {name: "encoding", typeName: "...*haxe__io__Encoding"}
        ],
        [],
        [GoStmt.GoExprStmt(GoExpr.GoCall(
          GoExpr.GoSelector(GoExpr.GoIdent("self"), "add"),
          [GoExpr.GoCall(GoExpr.GoIdent("haxe__io__Bytes_ofString"), [GoExpr.GoIdent("value")])]
        ))]
      ),
      GoDecl.GoFuncDecl(
        "getBytes",
        {name: "self", typeName: "*haxe__io__BytesBuffer"},
        [],
        ["*haxe__io__Bytes"],
        [
          GoStmt.GoVarDecl("copied", null, GoExpr.GoCall(
            GoExpr.GoIdent("hxrt.BytesClone"),
            [GoExpr.GoSelector(GoExpr.GoIdent("self"), "b")]
          ), true),
          GoStmt.GoReturn(GoExpr.GoRaw("&haxe__io__Bytes{b: copied, length: len(copied)}"))
        ]
      ),
      GoDecl.GoFuncDecl(
        "get_length",
        {name: "self", typeName: "*haxe__io__BytesBuffer"},
        [],
        ["int"],
        [GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("len"), [GoExpr.GoSelector(GoExpr.GoIdent("self"), "b")]))]
      )
    ];
  }

  function lowerDsStdlibShimDecls():Array<GoDecl> {
    return [
      GoDecl.GoStructDecl("haxe__ds__IntMap", [
        {name: "h", typeName: "map[int]any"}
      ]),
      GoDecl.GoStructDecl("haxe__ds__StringMap", [
        {name: "h", typeName: "map[string]any"}
      ]),
      GoDecl.GoStructDecl("haxe__ds__ObjectMap", [
        {name: "h", typeName: "map[any]any"}
      ]),
      GoDecl.GoStructDecl("haxe__ds__EnumValueMap", [
        {name: "h", typeName: "map[any]any"}
      ]),
      GoDecl.GoStructDecl("haxe__ds__List", [
        {name: "items", typeName: "[]any"},
        {name: "length", typeName: "int"}
      ]),
      GoDecl.GoFuncDecl(
        "New_haxe__ds__IntMap",
        null,
        [],
        ["*haxe__ds__IntMap"],
        [GoStmt.GoReturn(GoExpr.GoRaw("&haxe__ds__IntMap{h: map[int]any{}}"))]
      ),
      GoDecl.GoFuncDecl(
        "set",
        {name: "self", typeName: "*haxe__ds__IntMap"},
        [
          {name: "key", typeName: "int"},
          {name: "value", typeName: "any"}
        ],
        [],
        [GoStmt.GoAssign(GoExpr.GoIndex(GoExpr.GoSelector(GoExpr.GoIdent("self"), "h"), GoExpr.GoIdent("key")), GoExpr.GoIdent("value"))]
      ),
      GoDecl.GoFuncDecl(
        "get",
        {name: "self", typeName: "*haxe__ds__IntMap"},
        [{name: "key", typeName: "int"}],
        ["any"],
        [
          GoStmt.GoVarDecl("value", null, GoExpr.GoIndex(GoExpr.GoSelector(GoExpr.GoIdent("self"), "h"), GoExpr.GoIdent("key")), true),
          GoStmt.GoReturn(GoExpr.GoIdent("value"))
        ]
      ),
      GoDecl.GoFuncDecl(
        "exists",
        {name: "self", typeName: "*haxe__ds__IntMap"},
        [{name: "key", typeName: "int"}],
        ["bool"],
        [
          GoStmt.GoRaw("_, ok := self.h[key]"),
          GoStmt.GoReturn(GoExpr.GoIdent("ok"))
        ]
      ),
      GoDecl.GoFuncDecl(
        "remove",
        {name: "self", typeName: "*haxe__ds__IntMap"},
        [{name: "key", typeName: "int"}],
        ["bool"],
        [
          GoStmt.GoRaw("_, ok := self.h[key]"),
          GoStmt.GoRaw("delete(self.h, key)"),
          GoStmt.GoReturn(GoExpr.GoIdent("ok"))
        ]
      ),
      GoDecl.GoFuncDecl(
        "New_haxe__ds__StringMap",
        null,
        [],
        ["*haxe__ds__StringMap"],
        [GoStmt.GoReturn(GoExpr.GoRaw("&haxe__ds__StringMap{h: map[string]any{}}"))]
      ),
      GoDecl.GoFuncDecl(
        "set",
        {name: "self", typeName: "*haxe__ds__StringMap"},
        [
          {name: "key", typeName: "*string"},
          {name: "value", typeName: "any"}
        ],
        [],
        [GoStmt.GoAssign(GoExpr.GoRaw("self.h[*hxrt.StdString(key)]"), GoExpr.GoIdent("value"))]
      ),
      GoDecl.GoFuncDecl(
        "get",
        {name: "self", typeName: "*haxe__ds__StringMap"},
        [{name: "key", typeName: "*string"}],
        ["any"],
        [
          GoStmt.GoRaw("value := self.h[*hxrt.StdString(key)]"),
          GoStmt.GoReturn(GoExpr.GoIdent("value"))
        ]
      ),
      GoDecl.GoFuncDecl(
        "exists",
        {name: "self", typeName: "*haxe__ds__StringMap"},
        [{name: "key", typeName: "*string"}],
        ["bool"],
        [
          GoStmt.GoRaw("_, ok := self.h[*hxrt.StdString(key)]"),
          GoStmt.GoReturn(GoExpr.GoIdent("ok"))
        ]
      ),
      GoDecl.GoFuncDecl(
        "remove",
        {name: "self", typeName: "*haxe__ds__StringMap"},
        [{name: "key", typeName: "*string"}],
        ["bool"],
        [
          GoStmt.GoRaw("_, ok := self.h[*hxrt.StdString(key)]"),
          GoStmt.GoRaw("delete(self.h, *hxrt.StdString(key))"),
          GoStmt.GoReturn(GoExpr.GoIdent("ok"))
        ]
      ),
      GoDecl.GoFuncDecl(
        "New_haxe__ds__ObjectMap",
        null,
        [],
        ["*haxe__ds__ObjectMap"],
        [GoStmt.GoReturn(GoExpr.GoRaw("&haxe__ds__ObjectMap{h: map[any]any{}}"))]
      ),
      GoDecl.GoFuncDecl(
        "set",
        {name: "self", typeName: "*haxe__ds__ObjectMap"},
        [
          {name: "key", typeName: "any"},
          {name: "value", typeName: "any"}
        ],
        [],
        [GoStmt.GoAssign(GoExpr.GoIndex(GoExpr.GoSelector(GoExpr.GoIdent("self"), "h"), GoExpr.GoIdent("key")), GoExpr.GoIdent("value"))]
      ),
      GoDecl.GoFuncDecl(
        "get",
        {name: "self", typeName: "*haxe__ds__ObjectMap"},
        [{name: "key", typeName: "any"}],
        ["any"],
        [GoStmt.GoReturn(GoExpr.GoIndex(GoExpr.GoSelector(GoExpr.GoIdent("self"), "h"), GoExpr.GoIdent("key")))]
      ),
      GoDecl.GoFuncDecl(
        "exists",
        {name: "self", typeName: "*haxe__ds__ObjectMap"},
        [{name: "key", typeName: "any"}],
        ["bool"],
        [
          GoStmt.GoRaw("_, ok := self.h[key]"),
          GoStmt.GoReturn(GoExpr.GoIdent("ok"))
        ]
      ),
      GoDecl.GoFuncDecl(
        "remove",
        {name: "self", typeName: "*haxe__ds__ObjectMap"},
        [{name: "key", typeName: "any"}],
        ["bool"],
        [
          GoStmt.GoRaw("_, ok := self.h[key]"),
          GoStmt.GoRaw("delete(self.h, key)"),
          GoStmt.GoReturn(GoExpr.GoIdent("ok"))
        ]
      ),
      GoDecl.GoFuncDecl(
        "New_haxe__ds__EnumValueMap",
        null,
        [],
        ["*haxe__ds__EnumValueMap"],
        [GoStmt.GoReturn(GoExpr.GoRaw("&haxe__ds__EnumValueMap{h: map[any]any{}}"))]
      ),
      GoDecl.GoFuncDecl(
        "set",
        {name: "self", typeName: "*haxe__ds__EnumValueMap"},
        [
          {name: "key", typeName: "any"},
          {name: "value", typeName: "any"}
        ],
        [],
        [GoStmt.GoAssign(GoExpr.GoIndex(GoExpr.GoSelector(GoExpr.GoIdent("self"), "h"), GoExpr.GoIdent("key")), GoExpr.GoIdent("value"))]
      ),
      GoDecl.GoFuncDecl(
        "get",
        {name: "self", typeName: "*haxe__ds__EnumValueMap"},
        [{name: "key", typeName: "any"}],
        ["any"],
        [GoStmt.GoReturn(GoExpr.GoIndex(GoExpr.GoSelector(GoExpr.GoIdent("self"), "h"), GoExpr.GoIdent("key")))]
      ),
      GoDecl.GoFuncDecl(
        "exists",
        {name: "self", typeName: "*haxe__ds__EnumValueMap"},
        [{name: "key", typeName: "any"}],
        ["bool"],
        [
          GoStmt.GoRaw("_, ok := self.h[key]"),
          GoStmt.GoReturn(GoExpr.GoIdent("ok"))
        ]
      ),
      GoDecl.GoFuncDecl(
        "remove",
        {name: "self", typeName: "*haxe__ds__EnumValueMap"},
        [{name: "key", typeName: "any"}],
        ["bool"],
        [
          GoStmt.GoRaw("_, ok := self.h[key]"),
          GoStmt.GoRaw("delete(self.h, key)"),
          GoStmt.GoReturn(GoExpr.GoIdent("ok"))
        ]
      ),
      GoDecl.GoFuncDecl(
        "New_haxe__ds__List",
        null,
        [],
        ["*haxe__ds__List"],
        [GoStmt.GoReturn(GoExpr.GoRaw("&haxe__ds__List{items: []any{}, length: 0}"))]
      ),
      GoDecl.GoFuncDecl(
        "add",
        {name: "self", typeName: "*haxe__ds__List"},
        [{name: "item", typeName: "any"}],
        [],
        [
          GoStmt.GoAssign(
            GoExpr.GoSelector(GoExpr.GoIdent("self"), "items"),
            GoExpr.GoCall(GoExpr.GoIdent("append"), [
              GoExpr.GoSelector(GoExpr.GoIdent("self"), "items"),
              GoExpr.GoIdent("item")
            ])
          ),
          GoStmt.GoAssign(
            GoExpr.GoSelector(GoExpr.GoIdent("self"), "length"),
            GoExpr.GoCall(GoExpr.GoIdent("len"), [GoExpr.GoSelector(GoExpr.GoIdent("self"), "items")])
          )
        ]
      ),
      GoDecl.GoFuncDecl(
        "push",
        {name: "self", typeName: "*haxe__ds__List"},
        [{name: "item", typeName: "any"}],
        [],
        [GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("self"), "add"), [GoExpr.GoIdent("item")]))]
      ),
      GoDecl.GoFuncDecl(
        "pop",
        {name: "self", typeName: "*haxe__ds__List"},
        [],
        ["any"],
        [
          GoStmt.GoIf(
            GoExpr.GoBinary("==", GoExpr.GoCall(GoExpr.GoIdent("len"), [GoExpr.GoSelector(GoExpr.GoIdent("self"), "items")]), GoExpr.GoIntLiteral(0)),
            [GoStmt.GoReturn(GoExpr.GoNil)],
            null
          ),
          GoStmt.GoVarDecl("head", null, GoExpr.GoIndex(GoExpr.GoSelector(GoExpr.GoIdent("self"), "items"), GoExpr.GoIntLiteral(0)), true),
          GoStmt.GoAssign(
            GoExpr.GoSelector(GoExpr.GoIdent("self"), "items"),
            GoExpr.GoSlice(GoExpr.GoSelector(GoExpr.GoIdent("self"), "items"), GoExpr.GoIntLiteral(1), null)
          ),
          GoStmt.GoAssign(
            GoExpr.GoSelector(GoExpr.GoIdent("self"), "length"),
            GoExpr.GoCall(GoExpr.GoIdent("len"), [GoExpr.GoSelector(GoExpr.GoIdent("self"), "items")])
          ),
          GoStmt.GoReturn(GoExpr.GoIdent("head"))
        ]
      ),
      GoDecl.GoFuncDecl(
        "first",
        {name: "self", typeName: "*haxe__ds__List"},
        [],
        ["any"],
        [
          GoStmt.GoIf(
            GoExpr.GoBinary("==", GoExpr.GoCall(GoExpr.GoIdent("len"), [GoExpr.GoSelector(GoExpr.GoIdent("self"), "items")]), GoExpr.GoIntLiteral(0)),
            [GoStmt.GoReturn(GoExpr.GoNil)],
            null
          ),
          GoStmt.GoReturn(GoExpr.GoIndex(GoExpr.GoSelector(GoExpr.GoIdent("self"), "items"), GoExpr.GoIntLiteral(0)))
        ]
      ),
      GoDecl.GoFuncDecl(
        "last",
        {name: "self", typeName: "*haxe__ds__List"},
        [],
        ["any"],
        [
          GoStmt.GoVarDecl("size", null, GoExpr.GoCall(GoExpr.GoIdent("len"), [GoExpr.GoSelector(GoExpr.GoIdent("self"), "items")]), true),
          GoStmt.GoIf(
            GoExpr.GoBinary("==", GoExpr.GoIdent("size"), GoExpr.GoIntLiteral(0)),
            [GoStmt.GoReturn(GoExpr.GoNil)],
            null
          ),
          GoStmt.GoReturn(GoExpr.GoIndex(
            GoExpr.GoSelector(GoExpr.GoIdent("self"), "items"),
            GoExpr.GoBinary("-", GoExpr.GoIdent("size"), GoExpr.GoIntLiteral(1))
          ))
        ]
      )
    ];
  }

  function lowerSysStdlibShimDecls():Array<GoDecl> {
    return [
      GoDecl.GoStructDecl("Sys", []),
      GoDecl.GoStructDecl("sys__io__File", []),
      GoDecl.GoStructDecl("sys__io__ProcessOutput", [
        {name: "scanner", typeName: "*bufio.Scanner"}
      ]),
      GoDecl.GoStructDecl("sys__io__Process", [
        {name: "cmd", typeName: "*exec.Cmd"},
        {name: "stdout", typeName: "*sys__io__ProcessOutput"}
      ]),
      GoDecl.GoFuncDecl(
        "Sys_getCwd",
        null,
        [],
        ["*string"],
        [
          GoStmt.GoRaw("cwd, err := os.Getwd()"),
          GoStmt.GoIf(
            GoExpr.GoBinary("!=", GoExpr.GoIdent("err"), GoExpr.GoNil),
            [GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoStringLiteral("")]))],
            null
          ),
          GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoIdent("cwd")]))
        ]
      ),
      GoDecl.GoFuncDecl(
        "Sys_args",
        null,
        [],
        ["[]*string"],
        [
          GoStmt.GoRaw("args := os.Args"),
          GoStmt.GoIf(
            GoExpr.GoBinary("<=", GoExpr.GoCall(GoExpr.GoIdent("len"), [GoExpr.GoIdent("args")]), GoExpr.GoIntLiteral(1)),
            [GoStmt.GoReturn(GoExpr.GoRaw("[]*string{}"))],
            null
          ),
          GoStmt.GoRaw("out := make([]*string, 0, len(args)-1)"),
          GoStmt.GoRaw("for _, arg := range args[1:] {"),
          GoStmt.GoRaw("\tout = append(out, hxrt.StringFromLiteral(arg))"),
          GoStmt.GoRaw("}"),
          GoStmt.GoReturn(GoExpr.GoIdent("out"))
        ]
      ),
      GoDecl.GoFuncDecl(
        "sys__io__File_saveContent",
        null,
        [
          {name: "path", typeName: "*string"},
          {name: "content", typeName: "*string"}
        ],
        [],
        [GoStmt.GoAssign(
          GoExpr.GoIdent("_"),
          GoExpr.GoRaw("os.WriteFile(*hxrt.StdString(path), []byte(*hxrt.StdString(content)), 0o644)")
        )]
      ),
      GoDecl.GoFuncDecl(
        "sys__io__File_getContent",
        null,
        [{name: "path", typeName: "*string"}],
        ["*string"],
        [
          GoStmt.GoRaw("raw, err := os.ReadFile(*hxrt.StdString(path))"),
          GoStmt.GoIf(
            GoExpr.GoBinary("!=", GoExpr.GoIdent("err"), GoExpr.GoNil),
            [GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoStringLiteral("")]))],
            null
          ),
          GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoRaw("string(raw)")]))
        ]
      ),
      GoDecl.GoFuncDecl(
        "New_sys__io__Process",
        null,
        [
          {name: "command", typeName: "*string"},
          {name: "args", typeName: "[]*string"}
        ],
        ["*sys__io__Process"],
        [
          GoStmt.GoRaw("cmd := exec.Command(*hxrt.StdString(command), hxrt.StringSlice(args)...)"),
          GoStmt.GoRaw("stdoutPipe, _ := cmd.StdoutPipe()"),
          GoStmt.GoRaw("_ = cmd.Start()"),
          GoStmt.GoRaw("scanner := bufio.NewScanner(stdoutPipe)"),
          GoStmt.GoReturn(GoExpr.GoRaw("&sys__io__Process{cmd: cmd, stdout: &sys__io__ProcessOutput{scanner: scanner}}"))
        ]
      ),
      GoDecl.GoFuncDecl(
        "readLine",
        {name: "self", typeName: "*sys__io__ProcessOutput"},
        [],
        ["*string"],
        [
          GoStmt.GoIf(
            GoExpr.GoRaw("self == nil || self.scanner == nil"),
            [GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoStringLiteral("")]))],
            null
          ),
          GoStmt.GoIf(
            GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoSelector(GoExpr.GoIdent("self"), "scanner"), "Scan"), []),
            [GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoSelector(GoExpr.GoIdent("self"), "scanner"), "Text"), [])]))],
            null
          ),
          GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoStringLiteral("")]))
        ]
      ),
      GoDecl.GoFuncDecl(
        "close",
        {name: "self", typeName: "*sys__io__Process"},
        [],
        [],
        [
          GoStmt.GoIf(
            GoExpr.GoRaw("self == nil || self.cmd == nil"),
            [GoStmt.GoReturn(null)],
            null
          ),
          GoStmt.GoRaw("if self.cmd.Process != nil {"),
          GoStmt.GoRaw("\t_ = self.cmd.Process.Kill()"),
          GoStmt.GoRaw("}"),
          GoStmt.GoAssign(GoExpr.GoIdent("_"), GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoSelector(GoExpr.GoIdent("self"), "cmd"), "Wait"), []))
        ]
      )
    ];
  }

  function lowerClassDecls(classType:ClassType):Array<GoDecl> {
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
          decls.push(GoDecl.GoGlobalVarDecl(
            symbol,
            scalarGoType(field.type),
            valueExpr == null ? null : lowerExpr(valueExpr).expr
          ));
        case FMethod(_):
          var func = unwrapFunction(field.expr());
          if (func != null) {
            decls.push(lowerFunctionDecl(symbol, func, null));
          }
      }
    }

    return decls;
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
      body.push(GoStmt.GoAssign(
        GoExpr.GoSelector(GoExpr.GoIdent("self"), superTypeName),
        GoExpr.GoCall(GoExpr.GoIdent(constructorSymbol(superClass)), superCtorArgs)
      ));
      body.push(GoStmt.GoAssign(
        GoExpr.GoSelector(GoExpr.GoSelector(GoExpr.GoIdent("self"), superTypeName), "__hx_this"),
        GoExpr.GoIdent("self")
      ));
    }
    body.push(GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("self"), "__hx_this"), GoExpr.GoIdent("self")));
    body = body.concat(loweredCtorBody.body);
    body.push(GoStmt.GoReturn(GoExpr.GoIdent("self")));
    popFunctionVarNameScope();
    return GoDecl.GoFuncDecl(constructorSymbol(classType), null, params, ["*" + typeName], body);
  }

  function lowerInstanceMethodDecl(classType:ClassType, fieldName:String, func:TFunc):GoDecl {
    return lowerFunctionDecl(
      normalizeIdent(fieldName),
      func,
      {
        name: "self",
        typeName: "*" + classTypeName(classType)
      }
    );
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
        var useShort = loweredValue != null && !isNilExpr(loweredValue) && goType != "any";
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
          GoStmt.GoIf(
            lowerExpr(condition).expr,
            lowerToStatements(thenBranch),
            elseBranch == null ? null : lowerToStatements(elseBranch)
          )
        ];
      case TWhile(condition, body, normalWhile):
        if (normalWhile) {
          [GoStmt.GoWhile(lowerExpr(condition).expr, lowerToStatements(body))];
        } else {
          var firstPassVar = freshTempName("hx_do_first");
          var loweredCondition = lowerExpr(condition).expr;
          var loopCondition = GoExpr.GoBinary("||", GoExpr.GoIdent(firstPassVar), loweredCondition);
          var loopBody = [GoStmt.GoAssign(GoExpr.GoIdent(firstPassVar), GoExpr.GoBoolLiteral(false))]
            .concat(lowerToStatements(body));
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
        [GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("hxrt.Throw"), [lowerExpr(value).expr]))];
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
            GoStmt.GoIf(
              GoExpr.GoBinary(">", lenExpr, GoExpr.GoIntLiteral(0)),
              [GoStmt.GoAssign(targetExpr, GoExpr.GoSlice(targetExpr, null, GoExpr.GoBinary("-", lenExpr, GoExpr.GoIntLiteral(1))))],
              null
            )
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

    return GoStmt.GoSwitch(
      lowerExpr(value).expr,
      loweredCases,
      defaultExpr == null ? null : lowerToStatements(defaultExpr)
    );
  }

  function lowerSwitchExpr(value:TypedExpr, cases:Array<{values:Array<TypedExpr>, expr:TypedExpr}>, defaultExpr:Null<TypedExpr>, resultType:Type):LoweredExprWithPrefix {
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
        var dynamicValueExpr = haxeExceptionCatch
          ? GoExpr.GoCall(GoExpr.GoIdent("hxrt.ExceptionCaught"), [GoExpr.GoIdent(caughtName)])
          : GoExpr.GoIdent(caughtName);
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
      dynamicBody = [GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("hxrt.Throw"), [GoExpr.GoIdent(caughtName)]))];
    }

    var catchBody:Array<GoStmt> = if (typedCases.length == 0) {
      dynamicBody;
    } else {
      [GoStmt.GoTypeSwitch(GoExpr.GoIdent(caughtName), typeBindingName, typedCases, dynamicBody)];
    };

    return GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("hxrt.TryCatch"), [
      GoExpr.GoFuncLiteral([], [], lowerToStatements(tryExpr)),
      GoExpr.GoFuncLiteral([{name: caughtName, typeName: "any"}], [], catchBody)
    ]));
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
        lowerField(target, access).expr;
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
        var payload = GoExpr.GoIndex(
          GoExpr.GoSelector(lowerExpr(target).expr, "params"),
          GoExpr.GoIntLiteral(index)
        );
        var payloadType = scalarGoType(expr.t);
        {
          expr: payloadType == "any" ? payload : GoExpr.GoTypeAssert(payload, payloadType),
          isStringLike: isStringType(expr.t)
        };
      case TNew(classRef, _, args):
        {
          expr: GoExpr.GoCall(
            GoExpr.GoIdent(constructorSymbol(classRef.get())),
            [for (arg in args) lowerExpr(arg).expr]
          ),
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
          expr: GoExpr.GoFuncLiteral(
            loweredParams,
            loweredResults,
            loweredBody
          ),
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
              expr: GoExpr.GoCall(
                GoExpr.GoFuncLiteral(
                  [],
                  [typeToGoType(left.t)],
                  loweredRight.prefix.concat([
                    GoStmt.GoAssign(targetExpr, assignExpr),
                    GoStmt.GoReturn(targetExpr)
                  ])
                ),
                []
              ),
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
                expr: GoExpr.GoCall(
                  GoExpr.GoFuncLiteral(
                    [],
                    [typeToGoType(value.t)],
                    [
                      GoStmt.GoVarDecl(temp, null, target, true),
                      GoStmt.GoAssign(target, GoExpr.GoBinary(opSymbol, target, GoExpr.GoIntLiteral(1))),
                      GoStmt.GoReturn(GoExpr.GoIdent(temp))
                    ]
                  ),
                  []
                ),
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
              expr: GoExpr.GoCall(
                GoExpr.GoFuncLiteral(
                  [],
                  [typeToGoType(value.t)],
                  [
                    GoStmt.GoAssign(target, GoExpr.GoBinary(opSymbol, target, GoExpr.GoIntLiteral(1))),
                    GoStmt.GoReturn(target)
                  ]
                ),
                []
              ),
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
      expr: GoExpr.GoCall(
        GoExpr.GoFuncLiteral(
          [],
          ["bool"],
          loweredValue.prefix.concat([GoStmt.GoReturn(loweredCheck)])
        ),
        []
      ),
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
      return stdIsOfTypeTypeSwitch(valueExpr, ["int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64", "uintptr"]);
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
        return GoExpr.GoCall(
          GoExpr.GoFuncLiteral(
            [{name: "hx_value", typeName: valueTypeName}],
            ["bool"],
            [
              GoStmt.GoIf(
                GoExpr.GoBinary("==", GoExpr.GoIdent("hx_value"), GoExpr.GoNil),
                [GoStmt.GoReturn(GoExpr.GoBoolLiteral(false))],
                null
              ),
              GoStmt.GoRaw("_, ok := hx_value.__hx_this.(" + targetPointerType + ")"),
              GoStmt.GoReturn(GoExpr.GoIdent("ok"))
            ]
          ),
          [valueExpr]
        );
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
      return fullEnumName(valueEnum) == fullEnumName(targetEnum)
        ? GoExpr.GoBinary("!=", valueExpr, GoExpr.GoNil)
        : GoExpr.GoBoolLiteral(false);
    }

    if (!isAnyLikeType(valueType)) {
      return GoExpr.GoBoolLiteral(false);
    }

    return stdIsOfTypeTypeSwitch(valueExpr, ["*" + enumTypeName(targetEnum)]);
  }

  function stdIsOfTypeTypeSwitch(valueExpr:GoExpr, typeNames:Array<String>):GoExpr {
    var bindingName = "hx_type";
    return GoExpr.GoCall(
      GoExpr.GoFuncLiteral(
        [{name: "hx_value", typeName: "any"}],
        ["bool"],
        [
          GoStmt.GoTypeSwitch(
            GoExpr.GoIdent("hx_value"),
            bindingName,
            [for (typeName in typeNames) {
              typeName: typeName,
              body: [
                GoStmt.GoAssign(GoExpr.GoIdent("_"), GoExpr.GoIdent(bindingName)),
                GoStmt.GoReturn(GoExpr.GoBoolLiteral(true))
              ]
            }],
            [
              GoStmt.GoAssign(GoExpr.GoIdent("_"), GoExpr.GoIdent(bindingName)),
              GoStmt.GoReturn(GoExpr.GoBoolLiteral(false))
            ]
          )
        ]
      ),
      [GoExpr.GoCall(GoExpr.GoIdent("any"), [valueExpr])]
    );
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
      case TField(_, FStatic(classRef, field)):
        var classType = classRef.get();
        classType.name == className && classType.pack.join(".") == classPack.join(".") && field.get().name == fieldName;
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
      case TFun(args, _):
        index >= 0 && index < args.length ? args[index].t : null;
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
          expr: GoExpr.GoCall(
            GoExpr.GoIdent(typedStringOps ? "hxrt.StringConcatStringPtr" : "hxrt.StringConcatAny"),
            [leftLowered.expr, rightLowered.expr]
          ),
          isStringLike: true
        };
      case OpEq if (stringMode):
        {
          expr: GoExpr.GoCall(
            GoExpr.GoIdent(typedStringOps ? "hxrt.StringEqualStringPtr" : "hxrt.StringEqualAny"),
            [leftLowered.expr, rightLowered.expr]
          ),
          isStringLike: false
        };
      case OpNotEq if (stringMode):
        {
          expr: GoExpr.GoUnary(
            "!",
            GoExpr.GoCall(
              GoExpr.GoIdent(typedStringOps ? "hxrt.StringEqualStringPtr" : "hxrt.StringEqualAny"),
              [leftLowered.expr, rightLowered.expr]
            )
          ),
          isStringLike: false
        };
      case OpUShr:
        {
          expr: GoExpr.GoCall(
            GoExpr.GoIdent("int"),
            [
              GoExpr.GoBinary(
                ">>",
                GoExpr.GoCall(GoExpr.GoIdent("uint32"), [leftLowered.expr]),
                GoExpr.GoCall(GoExpr.GoIdent("uint"), [rightLowered.expr])
              )
            ]
          ),
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
      return GoExpr.GoCall(
        GoExpr.GoIdent(typedStringOps ? "hxrt.StringConcatStringPtr" : "hxrt.StringConcatAny"),
        [leftExpr, rightExpr]
      );
    }
    if (op == OpUShr) {
      return GoExpr.GoCall(
        GoExpr.GoIdent("int"),
        [
          GoExpr.GoBinary(
            ">>",
            GoExpr.GoCall(GoExpr.GoIdent("uint32"), [leftExpr]),
            GoExpr.GoCall(GoExpr.GoIdent("uint"), [rightExpr])
          )
        ]
      );
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
        } else if (classType.pack.length == 0 && classType.name == "String") {
          "*string";
        } else if (classType.pack.length == 0 && classType.name == "Array" && params.length == 1) {
          "[]" + scalarGoType(params[0]);
        } else {
          "*" + classTypeName(classType);
        }
      case TEnum(enumRef, _):
        "*" + enumTypeName(enumRef.get());
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
      case TInst(classRef, _):
        var classType = classRef.get();
        classType.pack.length == 0 && classType.name == "String";
      case TAbstract(abstractRef, _):
        var abstractType = abstractRef.get();
        abstractType.pack.length == 0 && abstractType.name == "String";
      case _:
        false;
    };
  }

  function isBoolType(type:Type):Bool {
    var followed = Context.follow(type);
    return switch (followed) {
      case TAbstract(abstractRef, _):
        var abstractType = abstractRef.get();
        abstractType.pack.length == 0 && abstractType.name == "Bool";
      case _:
        false;
    };
  }

  function isIntType(type:Type):Bool {
    var followed = Context.follow(type);
    return switch (followed) {
      case TAbstract(abstractRef, _):
        var abstractType = abstractRef.get();
        abstractType.pack.length == 0 && abstractType.name == "Int";
      case _:
        false;
    };
  }

  function isFloatType(type:Type):Bool {
    var followed = Context.follow(type);
    return switch (followed) {
      case TAbstract(abstractRef, _):
        var abstractType = abstractRef.get();
        abstractType.pack.length == 0 && abstractType.name == "Float";
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

  function isArrayType(type:Type):Bool {
    if (restElementType(type) != null) {
      return true;
    }

    var followed = Context.follow(type);
    return switch (followed) {
      case TInst(classRef, _):
        var classType = classRef.get();
        classType.pack.length == 0 && classType.name == "Array";
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
        } else if (classType.pack.length == 0 && classType.name == "String") {
          "*string";
        } else if (classType.pack.length == 0 && classType.name == "Array" && params.length == 1) {
          "[]" + scalarGoType(params[0]);
        } else {
          "*" + classTypeName(classType);
        }
      case TEnum(enumRef, _):
        "*" + enumTypeName(enumRef.get());
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
      case TAbstract(abstractRef, _):
        var abstractType = abstractRef.get();
        abstractType.pack.length == 0 && abstractType.name == "Void";
      case _:
        false;
    };
  }

  function isDynamicCatchType(type:Type):Bool {
    var followed = Context.follow(type);
    return switch (followed) {
      case TDynamic(_):
        true;
      case TAbstract(abstractRef, _):
        var abstractType = abstractRef.get();
        abstractType.pack.length == 0 && abstractType.name == "Dynamic";
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
        case _:
      }
      return;
    }

    if (pack == "haxe.ds") {
      switch (classType.name) {
        case "IntMap", "StringMap", "ObjectMap", "EnumValueMap", "List":
          requireStdlibShimGroup("ds");
        case _:
      }
      return;
    }

    if ((pack == "" && classType.name == "Sys")
      || (pack == "sys.io" && (classType.name == "File" || classType.name == "Process"))) {
      requireStdlibShimGroup("sys");
      return;
    }

    if ((pack == "haxe" && classType.name == "Json")
      || (pack == "haxe.format" && (classType.name == "JsonParser" || classType.name == "JsonPrinter"))) {
      requireStdlibShimGroup("json");
    }
  }

  function classTypeName(classType:ClassType):String {
    noteStdlibClass(classType);
    return GoNaming.typeSymbol(classType.pack, classType.name);
  }

  function enumTypeName(enumType:EnumType):String {
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
