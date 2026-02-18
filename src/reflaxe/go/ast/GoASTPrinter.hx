package reflaxe.go.ast;

import reflaxe.go.ast.GoAST.GoDecl;
import reflaxe.go.ast.GoAST.GoExpr;
import reflaxe.go.ast.GoAST.GoFile;
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
      case GoFuncDecl(name, params, results, body):
        var out = new StringBuf();
        out.add("func ");
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

  static function printStmt(stmt:GoStmt):String {
    return switch (stmt) {
      case GoExprStmt(expr): printExpr(expr);
      case GoReturn(expr): expr == null ? "return" : "return " + printExpr(expr);
    }
  }

  static function printExpr(expr:GoExpr):String {
    return switch (expr) {
      case GoIdent(name): name;
      case GoStringLiteral(value): '"' + escapeString(value) + '"';
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
