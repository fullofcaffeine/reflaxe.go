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

class RewriteVirtualCallsPass implements IGoASTPass {
  public function new() {}

  public function getName():String {
    return "rewrite_virtual_calls";
  }

  public function getDependencies():Array<String> {
    return ["normalize_names"];
  }

  public function run(file:GoFile, context:CompilationContext):GoFile {
    if (context.profile == GoProfile.Portable) {
      return file;
    }

    var leafReceivers = detectLeafReceiverTypes(file.decls);
    return {
      packageName: file.packageName,
      imports: file.imports,
      decls: [for (decl in file.decls) rewriteDecl(decl, leafReceivers)]
    };
  }

  function detectLeafReceiverTypes(decls:Array<GoDecl>):Map<String, Bool> {
    var structHasVirtual = new Map<String, Bool>();
    var subclasses = new Map<String, Bool>();

    for (decl in decls) {
      switch (decl) {
        case GoDecl.GoStructDecl(name, fields):
          var hasVirtual = false;
          for (field in fields) {
            if (field.name == "__hx_this") {
              hasVirtual = true;
            } else if (field.name == "" && StringTools.startsWith(field.typeName, "*")) {
              subclasses.set(field.typeName.substr(1), true);
            }
          }
          structHasVirtual.set(name, hasVirtual);
        case _:
      }
    }

    var leafReceivers = new Map<String, Bool>();
    for (name in structHasVirtual.keys()) {
      if (structHasVirtual.get(name) && !subclasses.exists(name)) {
        leafReceivers.set("*" + name, true);
      }
    }
    return leafReceivers;
  }

  function rewriteDecl(decl:GoDecl, leafReceivers:Map<String, Bool>):GoDecl {
    return switch (decl) {
      case GoDecl.GoFuncDecl(name, receiver, params, results, body):
        var receiverName = receiver == null ? null : receiver.name;
        var canDevirtualizeSelf = receiver != null && leafReceivers.exists(receiver.typeName);
        GoDecl.GoFuncDecl(
          name,
          receiver,
          params,
          results,
          [for (stmt in body) rewriteStmt(stmt, receiverName, canDevirtualizeSelf)]
        );
      case GoDecl.GoGlobalVarDecl(name, typeName, value):
        GoDecl.GoGlobalVarDecl(name, typeName, value == null ? null : rewriteExpr(value, null, false));
      case _:
        decl;
    };
  }

  function rewriteStmt(stmt:GoStmt, receiverName:Null<String>, canDevirtualizeSelf:Bool):GoStmt {
    return switch (stmt) {
      case GoStmt.GoVarDecl(name, typeName, value, useShort):
        GoStmt.GoVarDecl(name, typeName, value == null ? null : rewriteExpr(value, receiverName, canDevirtualizeSelf), useShort);
      case GoStmt.GoAssign(left, right):
        GoStmt.GoAssign(
          rewriteExpr(left, receiverName, canDevirtualizeSelf),
          rewriteExpr(right, receiverName, canDevirtualizeSelf)
        );
      case GoStmt.GoExprStmt(expr):
        GoStmt.GoExprStmt(rewriteExpr(expr, receiverName, canDevirtualizeSelf));
      case GoStmt.GoWhile(cond, body):
        GoStmt.GoWhile(
          rewriteExpr(cond, receiverName, canDevirtualizeSelf),
          [for (inner in body) rewriteStmt(inner, receiverName, canDevirtualizeSelf)]
        );
      case GoStmt.GoIf(cond, thenBody, elseBody):
        GoStmt.GoIf(
          rewriteExpr(cond, receiverName, canDevirtualizeSelf),
          [for (inner in thenBody) rewriteStmt(inner, receiverName, canDevirtualizeSelf)],
          elseBody == null ? null : [for (inner in elseBody) rewriteStmt(inner, receiverName, canDevirtualizeSelf)]
        );
      case GoStmt.GoSwitch(value, cases, defaultBody):
        var rewrittenCases:Array<GoSwitchCase> = [];
        for (entry in cases) {
          rewrittenCases.push({
            values: [for (valueExpr in entry.values) rewriteExpr(valueExpr, receiverName, canDevirtualizeSelf)],
            body: [for (bodyStmt in entry.body) rewriteStmt(bodyStmt, receiverName, canDevirtualizeSelf)]
          });
        }
        GoStmt.GoSwitch(
          rewriteExpr(value, receiverName, canDevirtualizeSelf),
          rewrittenCases,
          defaultBody == null ? null : [for (inner in defaultBody) rewriteStmt(inner, receiverName, canDevirtualizeSelf)]
        );
      case GoStmt.GoTypeSwitch(value, bindingName, cases, defaultBody):
        var rewrittenCases:Array<GoTypeSwitchCase> = [];
        for (entry in cases) {
          rewrittenCases.push({
            typeName: entry.typeName,
            body: [for (bodyStmt in entry.body) rewriteStmt(bodyStmt, receiverName, canDevirtualizeSelf)]
          });
        }
        GoStmt.GoTypeSwitch(
          rewriteExpr(value, receiverName, canDevirtualizeSelf),
          bindingName,
          rewrittenCases,
          defaultBody == null ? null : [for (inner in defaultBody) rewriteStmt(inner, receiverName, canDevirtualizeSelf)]
        );
      case GoStmt.GoReturn(expr):
        GoStmt.GoReturn(expr == null ? null : rewriteExpr(expr, receiverName, canDevirtualizeSelf));
      case _:
        stmt;
    };
  }

  function rewriteExpr(expr:GoExpr, receiverName:Null<String>, canDevirtualizeSelf:Bool):GoExpr {
    var rewritten = switch (expr) {
      case GoExpr.GoSelector(target, field):
        GoExpr.GoSelector(rewriteExpr(target, receiverName, canDevirtualizeSelf), field);
      case GoExpr.GoIndex(target, index):
        GoExpr.GoIndex(
          rewriteExpr(target, receiverName, canDevirtualizeSelf),
          rewriteExpr(index, receiverName, canDevirtualizeSelf)
        );
      case GoExpr.GoSlice(target, start, end):
        GoExpr.GoSlice(
          rewriteExpr(target, receiverName, canDevirtualizeSelf),
          start == null ? null : rewriteExpr(start, receiverName, canDevirtualizeSelf),
          end == null ? null : rewriteExpr(end, receiverName, canDevirtualizeSelf)
        );
      case GoExpr.GoArrayLiteral(elementType, elements):
        GoExpr.GoArrayLiteral(elementType, [for (element in elements) rewriteExpr(element, receiverName, canDevirtualizeSelf)]);
      case GoExpr.GoFuncLiteral(params, results, body):
        // Nested closures may shadow names, so do not apply receiver-specific rewrites inside.
        GoExpr.GoFuncLiteral(params, results, [for (stmt in body) rewriteStmt(stmt, null, false)]);
      case GoExpr.GoTypeAssert(inner, typeName):
        GoExpr.GoTypeAssert(rewriteExpr(inner, receiverName, canDevirtualizeSelf), typeName);
      case GoExpr.GoUnary(op, inner):
        GoExpr.GoUnary(op, rewriteExpr(inner, receiverName, canDevirtualizeSelf));
      case GoExpr.GoBinary(op, left, right):
        GoExpr.GoBinary(
          op,
          rewriteExpr(left, receiverName, canDevirtualizeSelf),
          rewriteExpr(right, receiverName, canDevirtualizeSelf)
        );
      case GoExpr.GoCall(callee, args):
        GoExpr.GoCall(
          rewriteExpr(callee, receiverName, canDevirtualizeSelf),
          [for (arg in args) rewriteExpr(arg, receiverName, canDevirtualizeSelf)]
        );
      case _:
        expr;
    };

    if (!canDevirtualizeSelf || receiverName == null) {
      return rewritten;
    }

    return switch (rewritten) {
      case GoExpr.GoSelector(GoExpr.GoSelector(GoExpr.GoIdent(name), "__hx_this"), field) if (name == receiverName):
        GoExpr.GoSelector(GoExpr.GoIdent(name), field);
      case _:
        rewritten;
    };
  }
}
