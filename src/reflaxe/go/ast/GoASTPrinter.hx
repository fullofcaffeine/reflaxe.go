package reflaxe.go.ast;

import reflaxe.go.ast.GoAST.GoDecl;
import reflaxe.go.ast.GoAST.GoExpr;
import reflaxe.go.ast.GoAST.GoFile;
import reflaxe.go.ast.GoAST.GoInterfaceMethod;
import reflaxe.go.ast.GoAST.GoParam;
import reflaxe.go.ast.GoAST.GoStmt;

class GoASTPrinter {
  public static function printFile(file:GoFile):String {
    var out = new StringBuf();
    out.add("package ");
    out.add(file.packageName);
    out.add("\n\n");

    if (file.imports.length == 1) {
      out.add('import "');
      out.add(file.imports[0]);
      out.add('"\n\n');
    } else if (file.imports.length > 1) {
      out.add("import (\n");
      for (path in file.imports) {
        out.add('\t"');
        out.add(path);
        out.add('"\n');
      }
      out.add(")\n\n");
    }

    var isFirst = true;
    for (decl in file.decls) {
      if (!isFirst) {
        out.add("\n");
      }
      isFirst = false;
      out.add(printDecl(decl));
    }

    return out.toString();
  }

  static function printDecl(decl:GoDecl):String {
    return switch (decl) {
      case GoInterfaceDecl(name, methods):
        var out = new StringBuf();
        out.add("type ");
        out.add(name);
        out.add(" interface {\n");
        for (method in methods) {
          out.add("\t");
          out.add(printInterfaceMethod(method));
          out.add("\n");
        }
        out.add("}\n");
        out.toString();
      case GoStructDecl(name, fields):
        var out = new StringBuf();
        out.add("type ");
        out.add(name);
        out.add(" struct {\n");
        for (field in fields) {
          out.add("\t");
          if (field.name != "") {
            out.add(field.name);
            out.add(" ");
          }
          out.add(field.typeName);
          out.add("\n");
        }
        out.add("}\n");
        out.toString();
      case GoGlobalVarDecl(name, typeName, value):
        if (value == null) {
          "var " + name + " " + typeName + "\n";
        } else {
          "var " + name + " " + typeName + " = " + printExpr(value) + "\n";
        }
      case GoFuncDecl(name, receiver, params, results, body):
        var out = new StringBuf();
        out.add("func ");
        if (receiver != null) {
          out.add("(");
          out.add(receiver.name);
          out.add(" ");
          out.add(receiver.typeName);
          out.add(") ");
        }
        out.add(name);
        out.add("(");
        out.add(printParams(params));
        out.add(")");

        if (results.length == 1) {
          out.add(" ");
          out.add(results[0]);
        } else if (results.length > 1) {
          out.add(" (");
          out.add(results.join(", "));
          out.add(")");
        }

        out.add(" {\n");
        for (stmt in body) {
          out.add("\t");
          out.add(printStmt(stmt));
          out.add("\n");
        }
        out.add("}\n");
        out.toString();
    }
  }

  static function printParams(params:Array<GoParam>):String {
    var rendered = new Array<String>();
    for (param in params) {
      rendered.push(param.name + " " + param.typeName);
    }
    return rendered.join(", ");
  }

  static function printInterfaceMethod(method:GoInterfaceMethod):String {
    var out = new StringBuf();
    out.add(method.name);
    out.add("(");
    out.add(printParams(method.params));
    out.add(")");
    if (method.results.length == 1) {
      out.add(" ");
      out.add(method.results[0]);
    } else if (method.results.length > 1) {
      out.add(" (");
      out.add(method.results.join(", "));
      out.add(")");
    }
    return out.toString();
  }

  static function printStmt(stmt:GoStmt):String {
    return switch (stmt) {
      case GoVarDecl(name, typeName, value, useShort):
        if (useShort && value != null) {
          name + " := " + printExpr(value);
        } else if (value == null) {
          typeName == null ? "var " + name : "var " + name + " " + typeName;
        } else if (typeName == null) {
          "var " + name + " = " + printExpr(value);
        } else {
          "var " + name + " " + typeName + " = " + printExpr(value);
        }
      case GoAssign(left, right):
        printExpr(left) + " = " + printExpr(right);
      case GoExprStmt(expr): printExpr(expr);
      case GoRaw(code): code;
      case GoWhile(cond, body):
        var out = new StringBuf();
        out.add("for ");
        out.add(printExpr(cond));
        out.add(" {\n");
        for (stmt in body) {
          out.add("\t");
          out.add(printStmt(stmt));
          out.add("\n");
        }
        out.add("}");
        out.toString();
      case GoIf(cond, thenBody, elseBody):
        var out = new StringBuf();
        out.add("if ");
        out.add(printExpr(cond));
        out.add(" {\n");
        for (stmt in thenBody) {
          out.add("\t");
          out.add(printStmt(stmt));
          out.add("\n");
        }
        out.add("}");
        if (elseBody != null) {
          out.add(" else {\n");
          for (stmt in elseBody) {
            out.add("\t");
            out.add(printStmt(stmt));
            out.add("\n");
          }
          out.add("}");
        }
        out.toString();
      case GoReturn(expr): expr == null ? "return" : "return " + printExpr(expr);
    }
  }

  static function printExpr(expr:GoExpr):String {
    return switch (expr) {
      case GoIdent(name): name;
      case GoIntLiteral(value): Std.string(value);
      case GoFloatLiteral(value): value;
      case GoBoolLiteral(value): value ? "true" : "false";
      case GoStringLiteral(value): '"' + escapeString(value) + '"';
      case GoNil: "nil";
      case GoSelector(target, field): printExpr(target) + "." + field;
      case GoIndex(target, index): printExpr(target) + "[" + printExpr(index) + "]";
      case GoSlice(target, start, end):
        var startCode = start == null ? "" : printExpr(start);
        var endCode = end == null ? "" : printExpr(end);
        printExpr(target) + "[" + startCode + ":" + endCode + "]";
      case GoArrayLiteral(elementType, elements):
        "[]" + elementType + "{" + [for (element in elements) printExpr(element)].join(", ") + "}";
      case GoFuncLiteral(params, results, body):
        var out = new StringBuf();
        out.add("func(");
        out.add(printParams(params));
        out.add(")");
        if (results.length == 1) {
          out.add(" ");
          out.add(results[0]);
        } else if (results.length > 1) {
          out.add(" (");
          out.add(results.join(", "));
          out.add(")");
        }
        out.add(" {\n");
        for (stmt in body) {
          out.add("\t");
          out.add(printStmt(stmt));
          out.add("\n");
        }
        out.add("}");
        out.toString();
      case GoRaw(code): code;
      case GoUnary(op, inner): op + printExpr(inner);
      case GoBinary(op, left, right): "(" + printExpr(left) + " " + op + " " + printExpr(right) + ")";
      case GoCall(callee, args):
        var renderedArgs = [for (arg in args) printExpr(arg)].join(", ");
        printExpr(callee) + "(" + renderedArgs + ")";
    }
  }

  static function escapeString(value:String):String {
    return value
      .split("\\").join("\\\\")
      .split("\"").join("\\\"")
      .split("\n").join("\\n")
      .split("\r").join("\\r")
      .split("\t").join("\\t");
  }
}
