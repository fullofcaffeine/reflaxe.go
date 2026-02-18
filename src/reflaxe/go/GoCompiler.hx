package reflaxe.go;

#if macro
import haxe.macro.Context;
import haxe.macro.Expr.Binop;
import haxe.macro.Expr.Unop;
import haxe.macro.Type;
import reflaxe.go.ast.GoAST.GoDecl;
import reflaxe.go.ast.GoAST.GoExpr;
import reflaxe.go.ast.GoAST.GoFile;
import reflaxe.go.ast.GoAST.GoStmt;
import reflaxe.go.ast.GoASTPrinter;
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

private typedef ArrayMethodCall = {
  final target:TypedExpr;
  final methodName:String;
}
#end

class GoCompiler {
  public function new() {}

  #if macro
  public function compileModule():Array<GoGeneratedFile> {
    var mainFile:GoFile = {
      packageName: "main",
      imports: ["snapshot/hxrt"],
      decls: [GoDecl.GoFuncDecl("main", [], [], lowerMainBody())]
    };

    return [{
      relativePath: "main.go",
      contents: GoASTPrinter.printFile(mainFile)
    }];
  }

  function lowerMainBody():Array<GoStmt> {
    var mainField = findMainField();
    var expr = mainField.expr();
    if (expr == null) {
      return [];
    }
    return switch (expr.expr) {
      case TFunction(func):
        lowerToStatements(func.expr);
      case TMeta(_, inner):
        lowerToStatements(inner);
      case _:
        lowerToStatements(expr);
    };
  }

  function findMainField():ClassField {
    var mainType = Context.getType("Main");
    return switch (Context.follow(mainType)) {
      case TInst(classRef, _):
        var classType = classRef.get();
        var field = findField(classType.statics.get(), "main");
        if (field == null) {
          Context.fatalError("Main.main was not found", Context.currentPos());
        }
        field;
      case _:
        Context.fatalError("Main must be a class type", Context.currentPos());
    };
  }

  function findField(fields:Array<ClassField>, name:String):Null<ClassField> {
    for (field in fields) {
      if (field.name == name) {
        return field;
      }
    }
    return null;
  }

  function lowerToStatements(expr:TypedExpr):Array<GoStmt> {
    return switch (expr.expr) {
      case TBlock(exprs):
        var out = new Array<GoStmt>();
        for (inner in exprs) {
          out = out.concat(lowerToStatements(inner));
        }
        out;
      case TMeta(_, inner):
        lowerToStatements(inner);
      case TParenthesis(inner):
        lowerToStatements(inner);
      case TCast(inner, _):
        lowerToStatements(inner);
      case TVar(variable, value):
        var loweredValue = value == null ? null : lowerExpr(value).expr;
        var goType = typeToGoType(variable.t);
        var useShort = loweredValue != null && !isNilExpr(loweredValue);
        [GoStmt.GoVarDecl(normalizeIdent(variable.name), goType, loweredValue, useShort)];
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
      case TArray(target, index):
        {
          expr: GoExpr.GoIndex(lowerExpr(target).expr, lowerExpr(index).expr),
          isStringLike: isStringType(expr.t)
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
      case _:
        Context.fatalError("Unsupported constant", Context.currentPos());
        {expr: GoExpr.GoNil, isStringLike: false};
    };
  }

  function lowerField(target:TypedExpr, access:FieldAccess):LoweredExpr {
    var loweredTarget = lowerExpr(target).expr;
    return switch (access) {
      case FStatic(_, field):
        var resolved = field.get();
        {
          expr: GoExpr.GoIdent(normalizeIdent(resolved.name)),
          isStringLike: isStringType(resolved.type)
        };
      case FInstance(_, _, field):
        var resolved = field.get();
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
    var followed = Context.follow(type);
    return switch (followed) {
      case TInst(classRef, params):
        var classType = classRef.get();
        if (classType.pack.length == 0 && classType.name == "String") {
          "*string";
        } else if (classType.pack.length == 0 && classType.name == "Array" && params.length == 1) {
          "[]" + scalarGoType(params[0]);
        } else {
          "any";
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
    var followed = Context.follow(type);
    return switch (followed) {
      case TInst(classRef, params):
        var classType = classRef.get();
        if (classType.pack.length == 0 && classType.name == "String") {
          "*string";
        } else if (classType.pack.length == 0 && classType.name == "Array" && params.length == 1) {
          "[]" + scalarGoType(params[0]);
        } else {
          "any";
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

  function isNilExpr(expr:GoExpr):Bool {
    return switch (expr) {
      case GoNil: true;
      case _: false;
    };
  }

  function normalizeIdent(name:String):String {
    var sanitized = new StringBuf();
    for (index in 0...name.length) {
      var ch = name.charCodeAt(index);
      var isLower = ch >= "a".code && ch <= "z".code;
      var isUpper = ch >= "A".code && ch <= "Z".code;
      var isDigit = ch >= "0".code && ch <= "9".code;
      if (isLower || isUpper || isDigit || ch == "_".code) {
        sanitized.addChar(ch);
      } else {
        sanitized.add("_");
      }
    }

    var normalized = sanitized.toString();
    if (normalized == "") {
      normalized = "hx_tmp";
    }
    var hasNonUnderscore = false;
    for (index in 0...normalized.length) {
      if (normalized.charCodeAt(index) != "_".code) {
        hasNonUnderscore = true;
        break;
      }
    }
    if (!hasNonUnderscore) {
      normalized = "hx_tmp";
    }

    var first = normalized.charCodeAt(0);
    var startsWithDigit = first >= "0".code && first <= "9".code;
    if (startsWithDigit) {
      normalized = "hx_" + normalized;
    }

    return switch (normalized) {
      case "func", "type", "var", "map", "range", "package", "return", "if", "else", "for", "go", "defer", "select", "chan", "switch", "fallthrough", "default", "case":
        normalized + "_";
      case _:
        normalized;
    };
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
