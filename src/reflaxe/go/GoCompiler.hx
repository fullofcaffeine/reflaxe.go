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
import reflaxe.go.ast.GoAST.GoParam;
import reflaxe.go.ast.GoAST.GoStmt;
import reflaxe.go.ast.GoASTPrinter;
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
#end

class GoCompiler {
  #if macro
  final staticFunctionInfos:Map<String, FunctionInfo>;
  final localFunctionScopes:Array<Map<String, FunctionInfo>>;
  final localRestIteratorScopes:Array<Array<String>>;
  var tempVarCounter:Int;
  #end

  public function new() {
    #if macro
    staticFunctionInfos = new Map<String, FunctionInfo>();
    localFunctionScopes = [];
    localRestIteratorScopes = [];
    tempVarCounter = 0;
    #end
  }

  #if macro
  public function compileModule(types:Array<ModuleType>):Array<GoGeneratedFile> {
    var classes = collectProjectClasses(types);
    buildStaticFunctionInfoTable(classes);
    var mainFile:GoFile = {
      packageName: "main",
      imports: ["snapshot/hxrt"],
      decls: lowerClasses(classes)
    };

    return [{
      relativePath: "main.go",
      contents: GoASTPrinter.printFile(mainFile)
    }];
  }

  function collectProjectClasses(types:Array<ModuleType>):Array<ClassType> {
    var classes = new Array<ClassType>();
    for (moduleType in types) {
      switch (moduleType) {
        case TClassDecl(classRef):
          var classType = classRef.get();
          if (isProjectClass(classType)) {
            classes.push(classType);
          }
        case _:
      }
    }

    classes.sort(function(a, b) return Reflect.compare(fullClassName(a), fullClassName(b)));

    var hasMain = false;
    for (classType in classes) {
      if (fullClassName(classType) == "Main") {
        hasMain = true;
        break;
      }
    }
    if (!hasMain) {
      Context.fatalError("Main class was not found among project modules", Context.currentPos());
    }

    return classes;
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

  function fullClassName(classType:ClassType):String {
    return classType.pack.length == 0 ? classType.name : classType.pack.join(".") + "." + classType.name;
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

  function lowerClasses(classes:Array<ClassType>):Array<GoDecl> {
    var decls = new Array<GoDecl>();
    for (classType in classes) {
      decls = decls.concat(lowerClassDecls(classType));
    }
    return decls;
  }

  function lowerClassDecls(classType:ClassType):Array<GoDecl> {
    var decls = new Array<GoDecl>();
    var typeName = classTypeName(classType);

    var instanceFields = new Array<GoParam>();
    var instanceMethods = new Array<{name:String, func:TFunc}>();
    for (field in classType.fields.get()) {
      switch (field.kind) {
        case FVar(_, _):
          instanceFields.push({
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

    if (instanceFields.length > 0 || instanceMethods.length > 0 || ctorFunc != null) {
      decls.push(GoDecl.GoStructDecl(typeName, instanceFields));
      decls.push(lowerConstructorDecl(classType, ctorFunc));
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
    var params = lowerFunctionParams(func);
    var results = lowerFunctionResults(func.t);
    var body = lowerFunctionBody(func.expr);
    return GoDecl.GoFuncDecl(name, receiver, params, results, body);
  }

  function lowerConstructorDecl(classType:ClassType, ctorFunc:Null<TFunc>):GoDecl {
    var typeName = classTypeName(classType);
    var params = ctorFunc == null ? [] : lowerFunctionParams(ctorFunc);
    var body = new Array<GoStmt>();
    body.push(GoStmt.GoVarDecl("self", null, GoExpr.GoRaw("&" + typeName + "{}"), true));
    if (ctorFunc != null) {
      body = body.concat(lowerFunctionBody(ctorFunc.expr));
    }
    body.push(GoStmt.GoReturn(GoExpr.GoIdent("self")));
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
        name: normalizeIdent(arg.v.name),
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
        var variableName = normalizeIdent(variable.name);
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
        var goType = typeToGoType(variable.t);
        var useShort = loweredValue != null && !isNilExpr(loweredValue);
        var decl = GoStmt.GoVarDecl(variableName, goType, loweredValue, useShort);

        if (prefix.length > 0) {
          prefix.push(decl);
          prefix;
        } else {
          [decl];
        }
      case TBinop(op, left, right):
        switch (op) {
          case OpAssign:
            [GoStmt.GoAssign(lowerLValue(left), lowerExpr(right).expr)];
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
      case TWhile(condition, body, _):
        [GoStmt.GoWhile(lowerExpr(condition).expr, lowerToStatements(body))];
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
      case TReturn(value):
        [GoStmt.GoReturn(value == null ? null : lowerExpr(value).expr)];
      case TCall(callee, args):
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
        lookupLocalFunction(normalizeIdent(variable.name));
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
        var name = normalizeIdent(variable.name);
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
        GoExpr.GoIdent(normalizeIdent(variable.name));
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
      case TNew(classRef, _, args):
        {
          expr: GoExpr.GoCall(
            GoExpr.GoIdent(constructorSymbol(classRef.get())),
            [for (arg in args) lowerExpr(arg).expr]
          ),
          isStringLike: false
        };
      case TFunction(func):
        {
          expr: GoExpr.GoFuncLiteral(
            lowerFunctionParams(func),
            lowerFunctionResults(func.t),
            lowerFunctionBody(func.expr)
          ),
          isStringLike: false
        };
      case TLocal(variable):
        {
          expr: GoExpr.GoIdent(normalizeIdent(variable.name)),
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
        lowerBinop(op, left, right, expr.t);
      case TUnop(op, postFix, value):
        if (postFix) {
          unsupportedExpr(expr, "Postfix unary operations are not supported yet");
        }
        {
          expr: GoExpr.GoUnary(unopSymbol(op), lowerExpr(value).expr),
          isStringLike: isStringType(expr.t)
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
      case FInstance(_, _, field):
        var resolved = field.get();
        var restTargetName = resolveRestIteratorTargetName(target);
        var restFieldName = restIteratorFieldName(restTargetName, resolved.name);
        if (restFieldName != null) {
          return {
            expr: GoExpr.GoIdent(restFieldName),
            isStringLike: isStringType(resolved.type)
          };
        }
        var loweredTarget = lowerExpr(target).expr;
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
      case FAnon(field):
        var resolved = field.get();
        var restTargetName = resolveRestIteratorTargetName(target);
        var restFieldName = restIteratorFieldName(restTargetName, resolved.name);
        if (restFieldName != null) {
          return {
            expr: GoExpr.GoIdent(restFieldName),
            isStringLike: isStringType(resolved.type)
          };
        }
        var loweredTarget = lowerExpr(target).expr;
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
      case FEnum(_, field):
        {
          expr: GoExpr.GoIdent(normalizeIdent(field.name)),
          isStringLike: false
        };
    };
  }

  function lowerCall(callee:TypedExpr, args:Array<TypedExpr>, returnType:Type):LoweredExpr {
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

    var loweredArgs = [for (arg in args) lowerExpr(arg).expr];
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

    return {
      expr: GoExpr.GoCall(lowerExpr(callee).expr, loweredArgs),
      isStringLike: isStringType(returnType)
    };
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

  function lowerBinop(op:Binop, left:TypedExpr, right:TypedExpr, resultType:Type):LoweredExpr {
    var leftLowered = lowerExpr(left);
    var rightLowered = lowerExpr(right);
    var stringMode = leftLowered.isStringLike || rightLowered.isStringLike || isStringType(left.t) || isStringType(right.t);

    return switch (op) {
      case OpAdd if (stringMode):
        {
          expr: GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringConcatAny"), [leftLowered.expr, rightLowered.expr]),
          isStringLike: true
        };
      case OpEq if (stringMode):
        {
          expr: GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringEqualAny"), [leftLowered.expr, rightLowered.expr]),
          isStringLike: false
        };
      case OpNotEq if (stringMode):
        {
          expr: GoExpr.GoUnary("!", GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringEqualAny"), [leftLowered.expr, rightLowered.expr])),
          isStringLike: false
        };
      case _:
        {
          expr: GoExpr.GoBinary(binopSymbol(op), leftLowered.expr, rightLowered.expr),
          isStringLike: isStringType(resultType)
        };
    };
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

  function typeToGoType(type:Type):String {
    var restElement = restElementType(type);
    if (restElement != null) {
      return "[]" + scalarGoType(restElement);
    }

    var followed = Context.follow(type);
    return switch (followed) {
      case TInst(classRef, params):
        var classType = classRef.get();
        if (classType.pack.length == 0 && classType.name == "String") {
          "*string";
        } else if (classType.pack.length == 0 && classType.name == "Array" && params.length == 1) {
          "[]" + scalarGoType(params[0]);
        } else {
          "*" + classTypeName(classType);
        }
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
      case TInst(classRef, params):
        var classType = classRef.get();
        if (classType.pack.length == 0 && classType.name == "String") {
          "*string";
        } else if (classType.pack.length == 0 && classType.name == "Array" && params.length == 1) {
          "[]" + scalarGoType(params[0]);
        } else {
          "*" + classTypeName(classType);
        }
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

  function isNilExpr(expr:GoExpr):Bool {
    return switch (expr) {
      case GoNil: true;
      case _: false;
    };
  }

  function classTypeName(classType:ClassType):String {
    return GoNaming.typeSymbol(classType.pack, classType.name);
  }

  function constructorSymbol(classType:ClassType):String {
    return GoNaming.constructorSymbol(classType.pack, classType.name);
  }

  function staticSymbol(classType:ClassType, fieldName:String):String {
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

  function unsupportedExpr(expr:TypedExpr, message:String):LoweredExpr {
    Context.fatalError(message + " :: " + Std.string(expr.expr), expr.pos);
    return {expr: GoExpr.GoNil, isStringLike: false};
  }
  #end
}
