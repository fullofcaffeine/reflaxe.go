package reflaxe.go.ast.transformers.passes;

import reflaxe.go.CompilationContext;
import reflaxe.go.GoProfile;
import reflaxe.go.ast.GoAST.GoDecl;
import reflaxe.go.ast.GoAST.GoExpr;
import reflaxe.go.ast.GoAST.GoFile;
import reflaxe.go.ast.GoAST.GoStmt;
import reflaxe.go.ast.GoAST.GoSwitchCase;
import reflaxe.go.ast.GoAST.GoTypeSwitchCase;
import reflaxe.go.ast.transformers.registry.RegistryCore.IGoASTPass;

class RewriteStringOpsPass implements IGoASTPass {
  public function new() {}

  public function getName():String {
    return "rewrite_string_ops";
  }

  public function getDependencies():Array<String> {
    return ["normalize_names"];
  }

  public function run(file:GoFile, context:CompilationContext):GoFile {
    if (context.profile == GoProfile.Portable) {
      return file;
    }

    return {
      packageName: file.packageName,
      imports: file.imports,
      decls: [for (decl in file.decls) rewriteDecl(decl)]
    };
  }

  function rewriteDecl(decl:GoDecl):GoDecl {
    return switch (decl) {
      case GoDecl.GoFuncDecl(name, receiver, params, results, body):
        GoDecl.GoFuncDecl(name, receiver, params, results, [for (stmt in body) rewriteStmt(stmt)]);
      case GoDecl.GoGlobalVarDecl(name, typeName, value):
        GoDecl.GoGlobalVarDecl(name, typeName, value == null ? null : rewriteExpr(value));
      case _:
        decl;
    };
  }

  function rewriteStmt(stmt:GoStmt):GoStmt {
    return switch (stmt) {
      case GoStmt.GoVarDecl(name, typeName, value, useShort):
        GoStmt.GoVarDecl(name, typeName, value == null ? null : rewriteExpr(value), useShort);
      case GoStmt.GoAssign(left, right):
        GoStmt.GoAssign(rewriteExpr(left), rewriteExpr(right));
      case GoStmt.GoExprStmt(expr):
        GoStmt.GoExprStmt(rewriteExpr(expr));
      case GoStmt.GoWhile(cond, body):
        GoStmt.GoWhile(rewriteExpr(cond), [for (inner in body) rewriteStmt(inner)]);
      case GoStmt.GoIf(cond, thenBody, elseBody):
        GoStmt.GoIf(
          rewriteExpr(cond),
          [for (inner in thenBody) rewriteStmt(inner)],
          elseBody == null ? null : [for (inner in elseBody) rewriteStmt(inner)]
        );
      case GoStmt.GoSwitch(value, cases, defaultBody):
        var rewrittenCases:Array<GoSwitchCase> = [];
        for (entry in cases) {
          rewrittenCases.push({
            values: [for (valueExpr in entry.values) rewriteExpr(valueExpr)],
            body: [for (bodyStmt in entry.body) rewriteStmt(bodyStmt)]
          });
        }
        GoStmt.GoSwitch(
          rewriteExpr(value),
          rewrittenCases,
          defaultBody == null ? null : [for (inner in defaultBody) rewriteStmt(inner)]
        );
      case GoStmt.GoTypeSwitch(value, bindingName, cases, defaultBody):
        var rewrittenCases:Array<GoTypeSwitchCase> = [];
        for (entry in cases) {
          rewrittenCases.push({
            typeName: entry.typeName,
            body: [for (bodyStmt in entry.body) rewriteStmt(bodyStmt)]
          });
        }
        GoStmt.GoTypeSwitch(
          rewriteExpr(value),
          bindingName,
          rewrittenCases,
          defaultBody == null ? null : [for (inner in defaultBody) rewriteStmt(inner)]
        );
      case GoStmt.GoReturn(expr):
        GoStmt.GoReturn(expr == null ? null : rewriteExpr(expr));
      case _:
        stmt;
    };
  }

  function rewriteExpr(expr:GoExpr):GoExpr {
    var rewritten = switch (expr) {
      case GoExpr.GoSelector(target, field):
        GoExpr.GoSelector(rewriteExpr(target), field);
      case GoExpr.GoIndex(target, index):
        GoExpr.GoIndex(rewriteExpr(target), rewriteExpr(index));
      case GoExpr.GoSlice(target, start, end):
        GoExpr.GoSlice(rewriteExpr(target), start == null ? null : rewriteExpr(start), end == null ? null : rewriteExpr(end));
      case GoExpr.GoArrayLiteral(elementType, elements):
        GoExpr.GoArrayLiteral(elementType, [for (element in elements) rewriteExpr(element)]);
      case GoExpr.GoFuncLiteral(params, results, body):
        GoExpr.GoFuncLiteral(params, results, [for (stmt in body) rewriteStmt(stmt)]);
      case GoExpr.GoTypeAssert(inner, typeName):
        GoExpr.GoTypeAssert(rewriteExpr(inner), typeName);
      case GoExpr.GoUnary(op, inner):
        GoExpr.GoUnary(op, rewriteExpr(inner));
      case GoExpr.GoBinary(op, left, right):
        GoExpr.GoBinary(op, rewriteExpr(left), rewriteExpr(right));
      case GoExpr.GoCall(callee, args):
        GoExpr.GoCall(rewriteExpr(callee), [for (arg in args) rewriteExpr(arg)]);
      case _:
        expr;
    };

    return foldExpr(rewritten);
  }

  function foldExpr(expr:GoExpr):GoExpr {
    return switch (expr) {
      case GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringConcatAny"), args):
        if (args.length == 2) {
          var left = staticStringValue(args[0]);
          var right = staticStringValue(args[1]);
          if (left != null && right != null) {
            GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoStringLiteral(left + right)]);
          } else {
            expr;
          }
        } else {
          expr;
        }
      case GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringConcatStringPtr"), args):
        if (args.length == 2) {
          var left = staticStringValue(args[0]);
          var right = staticStringValue(args[1]);
          if (left != null && right != null) {
            GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoStringLiteral(left + right)]);
          } else {
            expr;
          }
        } else {
          expr;
        }
      case GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringEqualAny"), args):
        if (args.length == 2) {
          var left = staticStringValue(args[0]);
          var right = staticStringValue(args[1]);
          if (left != null && right != null) {
            GoExpr.GoBoolLiteral(left == right);
          } else {
            expr;
          }
        } else {
          expr;
        }
      case GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringEqualStringPtr"), args):
        if (args.length == 2) {
          var left = staticStringValue(args[0]);
          var right = staticStringValue(args[1]);
          if (left != null && right != null) {
            GoExpr.GoBoolLiteral(left == right);
          } else {
            expr;
          }
        } else {
          expr;
        }
      case GoExpr.GoUnary("!", inner):
        switch (inner) {
          case GoExpr.GoBoolLiteral(value):
            GoExpr.GoBoolLiteral(!value);
          case GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringEqualAny"), args):
            if (args.length == 2) {
              var left = staticStringValue(args[0]);
              var right = staticStringValue(args[1]);
              if (left != null && right != null) {
                GoExpr.GoBoolLiteral(left != right);
              } else {
                expr;
              }
            } else {
              expr;
            }
          case GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringEqualStringPtr"), args):
            if (args.length == 2) {
              var left = staticStringValue(args[0]);
              var right = staticStringValue(args[1]);
              if (left != null && right != null) {
                GoExpr.GoBoolLiteral(left != right);
              } else {
                expr;
              }
            } else {
              expr;
            }
          case _:
            expr;
        }
      case _:
        expr;
    };
  }

  function staticStringValue(expr:GoExpr):Null<String> {
    return switch (expr) {
      case GoExpr.GoNil:
        "null";
      case GoExpr.GoStringLiteral(value):
        value;
      case GoExpr.GoBoolLiteral(value):
        value ? "true" : "false";
      case GoExpr.GoIntLiteral(value):
        Std.string(value);
      case GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), args):
        if (args.length == 1) {
          switch (args[0]) {
            case GoExpr.GoStringLiteral(value):
              value;
            case _:
              null;
          }
        } else {
          null;
        }
      case _:
        null;
    };
  }
}
