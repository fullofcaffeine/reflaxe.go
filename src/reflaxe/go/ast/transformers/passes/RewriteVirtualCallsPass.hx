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
    var leafReturnCallTargets = detectLeafReturningFunctions(file.decls, leafReceivers);
    return {
      packageName: file.packageName,
      imports: file.imports,
      decls: [for (decl in file.decls) rewriteDecl(decl, leafReceivers, leafReturnCallTargets)]
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

  function detectLeafReturningFunctions(decls:Array<GoDecl>, leafReceivers:Map<String, Bool>):Map<String, Bool> {
    var out = new Map<String, Bool>();
    for (decl in decls) {
      switch (decl) {
        case GoDecl.GoFuncDecl(name, receiver, _, results, _):
          if (receiver == null && results.length == 1 && leafReceivers.exists(results[0])) {
            out.set(name, true);
          }
        case _:
      }
    }
    return out;
  }

  function rewriteDecl(
    decl:GoDecl,
    leafReceivers:Map<String, Bool>,
    leafReturnCallTargets:Map<String, Bool>
  ):GoDecl {
    return switch (decl) {
      case GoDecl.GoFuncDecl(name, receiver, params, results, body):
        var receiverName = receiver == null ? null : receiver.name;
        var canDevirtualizeSelf = receiver != null && leafReceivers.exists(receiver.typeName);
        var localLeafVars = new Map<String, Bool>();
        GoDecl.GoFuncDecl(
          name,
          receiver,
          params,
          results,
          rewriteStmtList(body, receiverName, canDevirtualizeSelf, localLeafVars, leafReceivers, leafReturnCallTargets)
        );
      case GoDecl.GoGlobalVarDecl(name, typeName, value):
        GoDecl.GoGlobalVarDecl(
          name,
          typeName,
          value == null ? null : rewriteExpr(value, null, false, new Map<String, Bool>(), leafReceivers, leafReturnCallTargets)
        );
      case _:
        decl;
    };
  }

  function rewriteStmtList(
    stmts:Array<GoStmt>,
    receiverName:Null<String>,
    canDevirtualizeSelf:Bool,
    localLeafVars:Map<String, Bool>,
    leafReceivers:Map<String, Bool>,
    leafReturnCallTargets:Map<String, Bool>
  ):Array<GoStmt> {
    var out = new Array<GoStmt>();
    for (stmt in stmts) {
      out.push(rewriteStmt(stmt, receiverName, canDevirtualizeSelf, localLeafVars, leafReceivers, leafReturnCallTargets));
    }
    return out;
  }

  function rewriteStmt(
    stmt:GoStmt,
    receiverName:Null<String>,
    canDevirtualizeSelf:Bool,
    localLeafVars:Map<String, Bool>,
    leafReceivers:Map<String, Bool>,
    leafReturnCallTargets:Map<String, Bool>
  ):GoStmt {
    return switch (stmt) {
      case GoStmt.GoVarDecl(name, typeName, value, useShort):
        var rewrittenValue = value == null ? null : rewriteExpr(
          value,
          receiverName,
          canDevirtualizeSelf,
          localLeafVars,
          leafReceivers,
          leafReturnCallTargets
        );
        if (rewrittenValue != null && isLeafCandidateValue(rewrittenValue, leafReceivers, localLeafVars, leafReturnCallTargets)) {
          localLeafVars.set(name, true);
        } else {
          localLeafVars.remove(name);
        }
        GoStmt.GoVarDecl(name, typeName, rewrittenValue, useShort);
      case GoStmt.GoAssign(left, right):
        var rewrittenLeft = rewriteExpr(left, receiverName, canDevirtualizeSelf, localLeafVars, leafReceivers, leafReturnCallTargets);
        var rewrittenRight = rewriteExpr(right, receiverName, canDevirtualizeSelf, localLeafVars, leafReceivers, leafReturnCallTargets);
        switch (rewrittenLeft) {
          case GoExpr.GoIdent(name):
            if (isLeafCandidateValue(rewrittenRight, leafReceivers, localLeafVars, leafReturnCallTargets)) {
              localLeafVars.set(name, true);
            } else {
              localLeafVars.remove(name);
            }
          case _:
        }
        GoStmt.GoAssign(
          rewrittenLeft,
          rewrittenRight
        );
      case GoStmt.GoExprStmt(expr):
        GoStmt.GoExprStmt(rewriteExpr(expr, receiverName, canDevirtualizeSelf, localLeafVars, leafReceivers, leafReturnCallTargets));
      case GoStmt.GoRaw(code):
        clearCandidates(localLeafVars);
        GoStmt.GoRaw(code);
      case GoStmt.GoWhile(cond, body):
        var loopCandidates = cloneCandidates(localLeafVars);
        var rewrittenBody = rewriteStmtList(body, receiverName, canDevirtualizeSelf, loopCandidates, leafReceivers, leafReturnCallTargets);
        clearCandidates(localLeafVars);
        GoStmt.GoWhile(
          rewriteExpr(cond, receiverName, canDevirtualizeSelf, localLeafVars, leafReceivers, leafReturnCallTargets),
          rewrittenBody
        );
      case GoStmt.GoIf(cond, thenBody, elseBody):
        var thenCandidates = cloneCandidates(localLeafVars);
        var elseCandidates = cloneCandidates(localLeafVars);
        var rewrittenThen = rewriteStmtList(thenBody, receiverName, canDevirtualizeSelf, thenCandidates, leafReceivers, leafReturnCallTargets);
        var rewrittenElse = elseBody == null
          ? null
          : rewriteStmtList(elseBody, receiverName, canDevirtualizeSelf, elseCandidates, leafReceivers, leafReturnCallTargets);
        clearCandidates(localLeafVars);
        GoStmt.GoIf(
          rewriteExpr(cond, receiverName, canDevirtualizeSelf, localLeafVars, leafReceivers, leafReturnCallTargets),
          rewrittenThen,
          rewrittenElse
        );
      case GoStmt.GoSwitch(value, cases, defaultBody):
        var switchCandidates = cloneCandidates(localLeafVars);
        var rewrittenCases:Array<GoSwitchCase> = [];
        for (entry in cases) {
          var caseCandidates = cloneCandidates(switchCandidates);
          rewrittenCases.push({
            values: [
              for (valueExpr in entry.values)
                rewriteExpr(valueExpr, receiverName, canDevirtualizeSelf, caseCandidates, leafReceivers, leafReturnCallTargets)
            ],
            body: rewriteStmtList(entry.body, receiverName, canDevirtualizeSelf, caseCandidates, leafReceivers, leafReturnCallTargets)
          });
        }
        var rewrittenDefault = if (defaultBody == null) {
          null;
        } else {
          var defaultCandidates = cloneCandidates(switchCandidates);
          rewriteStmtList(
            defaultBody,
            receiverName,
            canDevirtualizeSelf,
            defaultCandidates,
            leafReceivers,
            leafReturnCallTargets
          );
        }
        clearCandidates(localLeafVars);
        GoStmt.GoSwitch(
          rewriteExpr(value, receiverName, canDevirtualizeSelf, localLeafVars, leafReceivers, leafReturnCallTargets),
          rewrittenCases,
          rewrittenDefault
        );
      case GoStmt.GoTypeSwitch(value, bindingName, cases, defaultBody):
        var typeSwitchCandidates = cloneCandidates(localLeafVars);
        var rewrittenCases:Array<GoTypeSwitchCase> = [];
        for (entry in cases) {
          var caseCandidates = cloneCandidates(typeSwitchCandidates);
          rewrittenCases.push({
            typeName: entry.typeName,
            body: rewriteStmtList(entry.body, receiverName, canDevirtualizeSelf, caseCandidates, leafReceivers, leafReturnCallTargets)
          });
        }
        var rewrittenDefault = if (defaultBody == null) {
          null;
        } else {
          var defaultCandidates = cloneCandidates(typeSwitchCandidates);
          rewriteStmtList(
            defaultBody,
            receiverName,
            canDevirtualizeSelf,
            defaultCandidates,
            leafReceivers,
            leafReturnCallTargets
          );
        }
        clearCandidates(localLeafVars);
        GoStmt.GoTypeSwitch(
          rewriteExpr(value, receiverName, canDevirtualizeSelf, localLeafVars, leafReceivers, leafReturnCallTargets),
          bindingName,
          rewrittenCases,
          rewrittenDefault
        );
      case GoStmt.GoReturn(expr):
        GoStmt.GoReturn(
          expr == null ? null : rewriteExpr(expr, receiverName, canDevirtualizeSelf, localLeafVars, leafReceivers, leafReturnCallTargets)
        );
      case _:
        stmt;
    };
  }

  function rewriteExpr(
    expr:GoExpr,
    receiverName:Null<String>,
    canDevirtualizeSelf:Bool,
    localLeafVars:Map<String, Bool>,
    leafReceivers:Map<String, Bool>,
    leafReturnCallTargets:Map<String, Bool>
  ):GoExpr {
    var rewritten = switch (expr) {
      case GoExpr.GoSelector(target, field):
        GoExpr.GoSelector(rewriteExpr(target, receiverName, canDevirtualizeSelf, localLeafVars, leafReceivers, leafReturnCallTargets), field);
      case GoExpr.GoIndex(target, index):
        GoExpr.GoIndex(
          rewriteExpr(target, receiverName, canDevirtualizeSelf, localLeafVars, leafReceivers, leafReturnCallTargets),
          rewriteExpr(index, receiverName, canDevirtualizeSelf, localLeafVars, leafReceivers, leafReturnCallTargets)
        );
      case GoExpr.GoSlice(target, start, end):
        GoExpr.GoSlice(
          rewriteExpr(target, receiverName, canDevirtualizeSelf, localLeafVars, leafReceivers, leafReturnCallTargets),
          start == null ? null : rewriteExpr(start, receiverName, canDevirtualizeSelf, localLeafVars, leafReceivers, leafReturnCallTargets),
          end == null ? null : rewriteExpr(end, receiverName, canDevirtualizeSelf, localLeafVars, leafReceivers, leafReturnCallTargets)
        );
      case GoExpr.GoArrayLiteral(elementType, elements):
        GoExpr.GoArrayLiteral(
          elementType,
          [for (element in elements) rewriteExpr(element, receiverName, canDevirtualizeSelf, localLeafVars, leafReceivers, leafReturnCallTargets)]
        );
      case GoExpr.GoFuncLiteral(params, results, body):
        // Nested closures may shadow names, so do not apply receiver-specific rewrites inside.
        GoExpr.GoFuncLiteral(
          params,
          results,
          rewriteStmtList(body, null, false, new Map<String, Bool>(), leafReceivers, leafReturnCallTargets)
        );
      case GoExpr.GoTypeAssert(inner, typeName):
        GoExpr.GoTypeAssert(
          rewriteExpr(inner, receiverName, canDevirtualizeSelf, localLeafVars, leafReceivers, leafReturnCallTargets),
          typeName
        );
      case GoExpr.GoUnary(op, inner):
        GoExpr.GoUnary(op, rewriteExpr(inner, receiverName, canDevirtualizeSelf, localLeafVars, leafReceivers, leafReturnCallTargets));
      case GoExpr.GoBinary(op, left, right):
        GoExpr.GoBinary(
          op,
          rewriteExpr(left, receiverName, canDevirtualizeSelf, localLeafVars, leafReceivers, leafReturnCallTargets),
          rewriteExpr(right, receiverName, canDevirtualizeSelf, localLeafVars, leafReceivers, leafReturnCallTargets)
        );
      case GoExpr.GoCall(callee, args):
        GoExpr.GoCall(
          rewriteExpr(callee, receiverName, canDevirtualizeSelf, localLeafVars, leafReceivers, leafReturnCallTargets),
          [for (arg in args) rewriteExpr(arg, receiverName, canDevirtualizeSelf, localLeafVars, leafReceivers, leafReturnCallTargets)]
        );
      case _:
        expr;
    };

    return switch (rewritten) {
      case GoExpr.GoSelector(GoExpr.GoSelector(target, "__hx_this"), field):
        if (shouldDevirtualizeTarget(
          target,
          receiverName,
          canDevirtualizeSelf,
          localLeafVars,
          leafReceivers,
          leafReturnCallTargets
        )) {
          GoExpr.GoSelector(target, field);
        } else {
          rewritten;
        }
      case _:
        rewritten;
    };
  }

  function shouldDevirtualizeTarget(
    target:GoExpr,
    receiverName:Null<String>,
    canDevirtualizeSelf:Bool,
    localLeafVars:Map<String, Bool>,
    leafReceivers:Map<String, Bool>,
    leafReturnCallTargets:Map<String, Bool>
  ):Bool {
    return switch (target) {
      case GoExpr.GoIdent(name):
        (canDevirtualizeSelf && receiverName != null && name == receiverName) || localLeafVars.exists(name);
      case _:
        isLeafTargetExpr(target, leafReceivers, leafReturnCallTargets);
    };
  }

  function isLeafConstructorCall(expr:GoExpr, leafReceivers:Map<String, Bool>):Bool {
    return switch (expr) {
      case GoExpr.GoCall(GoExpr.GoIdent(callee), _) if (StringTools.startsWith(callee, "New_")):
        var typeName = callee.substr("New_".length);
        leafReceivers.exists("*" + typeName);
      case _:
        false;
    };
  }

  function isLeafCandidateValue(
    expr:GoExpr,
    leafReceivers:Map<String, Bool>,
    localLeafVars:Map<String, Bool>,
    leafReturnCallTargets:Map<String, Bool>
  ):Bool {
    if (isLeafTargetExpr(expr, leafReceivers, leafReturnCallTargets)) {
      return true;
    }
    return switch (expr) {
      case GoExpr.GoIdent(name):
        localLeafVars.exists(name);
      case _:
        false;
    };
  }

  function isLeafReturningCallExpr(expr:GoExpr, leafReturnCallTargets:Map<String, Bool>):Bool {
    return switch (expr) {
      case GoExpr.GoCall(GoExpr.GoIdent(callee), _):
        leafReturnCallTargets.exists(callee);
      case _:
        false;
    };
  }

  function isLeafTargetExpr(
    expr:GoExpr,
    leafReceivers:Map<String, Bool>,
    leafReturnCallTargets:Map<String, Bool>
  ):Bool {
    return isLeafConstructorCall(expr, leafReceivers) || isLeafReturningCallExpr(expr, leafReturnCallTargets);
  }

  function cloneCandidates(source:Map<String, Bool>):Map<String, Bool> {
    var out = new Map<String, Bool>();
    for (name in source.keys()) {
      out.set(name, true);
    }
    return out;
  }

  function clearCandidates(source:Map<String, Bool>):Void {
    var names = [for (name in source.keys()) name];
    for (name in names) {
      source.remove(name);
    }
  }
}
