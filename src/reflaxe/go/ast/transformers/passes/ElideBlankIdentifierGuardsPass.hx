package reflaxe.go.ast.transformers.passes;

import reflaxe.go.CompilationContext;
import reflaxe.go.ast.GoAST.GoDecl;
import reflaxe.go.ast.GoAST.GoExpr;
import reflaxe.go.ast.GoAST.GoFile;
import reflaxe.go.ast.GoAST.GoSelectCase;
import reflaxe.go.ast.GoAST.GoSelectClause;
import reflaxe.go.ast.GoAST.GoStmt;
import reflaxe.go.ast.GoAST.GoSwitchCase;
import reflaxe.go.ast.GoAST.GoTypeSwitchCase;
import reflaxe.go.ast.transformers.registry.RegistryCore.IGoASTPass;

class ElideBlankIdentifierGuardsPass implements IGoASTPass {
	var scopes:Array<Map<String, Int>>;
	var nextBindingId:Int;
	var readCounts:Map<Int, Int>;
	var unsafeBindings:Map<Int, Bool>;

	public function new() {}

	public function getName():String {
		return "elide_blank_identifier_guards";
	}

	public function getDependencies():Array<String> {
		return ["insert_runtime_prelude"];
	}

	public function run(file:GoFile, _context:CompilationContext):GoFile {
		resetCollectionState();
		collectFile(file);

		resetRewriteState();
		return rewriteFile(file);
	}

	function resetCollectionState():Void {
		scopes = [];
		nextBindingId = 0;
		readCounts = new Map<Int, Int>();
		unsafeBindings = new Map<Int, Bool>();
	}

	function resetRewriteState():Void {
		scopes = [];
		nextBindingId = 0;
	}

	function pushScope():Void {
		scopes.push(new Map<String, Int>());
	}

	function popScope():Void {
		if (scopes.length > 0) {
			scopes.pop();
		}
	}

	function declareBinding(name:String):Void {
		if (name == "_" || scopes.length == 0) {
			return;
		}
		var scope = scopes[scopes.length - 1];
		scope.set(name, nextBindingId++);
	}

	function resolveBinding(name:String):Null<Int> {
		if (name == "_") {
			return null;
		}
		var index = scopes.length - 1;
		while (index >= 0) {
			var scope = scopes[index];
			if (scope.exists(name)) {
				return scope.get(name);
			}
			index--;
		}
		return null;
	}

	function markBindingsUnsafeFromRaw(code:String):Void {
		for (scope in scopes) {
			for (name in scope.keys()) {
				if (!rawMentionsIdent(code, name)) {
					continue;
				}
				unsafeBindings.set(scope.get(name), true);
			}
		}
	}

	function rawMentionsIdent(code:String, ident:String):Bool {
		var nameLength = ident.length;
		if (nameLength == 0 || ident == "_") {
			return false;
		}
		var offset = 0;
		while (true) {
			var index = code.indexOf(ident, offset);
			if (index < 0) {
				return false;
			}
			var before = index == 0 ? null : code.charCodeAt(index - 1);
			var afterIndex = index + nameLength;
			var after = afterIndex >= code.length ? null : code.charCodeAt(afterIndex);
			if (!isIdentCode(before) && !isIdentCode(after)) {
				return true;
			}
			offset = index + nameLength;
		}
	}

	function isIdentCode(code:Null<Int>):Bool {
		if (code == null) {
			return false;
		}
		return (code >= "A".code && code <= "Z".code)
			|| (code >= "a".code && code <= "z".code)
			|| (code >= "0".code && code <= "9".code)
			|| code == "_".code;
	}

	function addRead(bindingId:Null<Int>):Void {
		if (bindingId == null) {
			return;
		}
		var current = readCounts.exists(bindingId) ? readCounts.get(bindingId) : 0;
		readCounts.set(bindingId, current + 1);
	}

	function collectFile(file:GoFile):Void {
		for (decl in file.decls) {
			collectDecl(decl);
		}
	}

	function collectDecl(decl:GoDecl):Void {
		switch (decl) {
			case GoDecl.GoFuncDecl(_, receiver, params, _, body):
				pushScope();
				if (receiver != null) {
					declareBinding(receiver.name);
				}
				for (param in params) {
					declareBinding(param.name);
				}
				collectStmtList(body);
				popScope();
			case GoDecl.GoGlobalVarDecl(_, _, value):
				if (value != null) {
					collectExprReads(value);
				}
			case _:
		}
	}

	function collectStmtList(stmts:Array<GoStmt>):Void {
		for (stmt in stmts) {
			collectStmt(stmt);
		}
	}

	function collectStmt(stmt:GoStmt):Void {
		switch (stmt) {
			case GoStmt.GoVarDecl(name, _, value, _):
				if (value != null) {
					collectExprReads(value);
				}
				declareBinding(name);
			case GoStmt.GoMultiAssign(names, value, useShort):
				collectExprReads(value);
				if (useShort) {
					for (name in names) {
						declareBinding(name);
					}
				}
			case GoStmt.GoAssign(GoExpr.GoIdent("_"), GoExpr.GoIdent(_)):
				// Ignore blank-identifier guards while counting real reads.
			case GoStmt.GoAssign(left, right):
				collectLValueReads(left);
				collectExprReads(right);
			case GoStmt.GoExprStmt(expr):
				collectExprReads(expr);
			case GoStmt.GoGoStmt(call):
				collectExprReads(call);
			case GoStmt.GoDeferStmt(call):
				collectExprReads(call);
			case GoStmt.GoSendStmt(channel, value):
				collectExprReads(channel);
				collectExprReads(value);
			case GoStmt.GoRaw(code):
				markBindingsUnsafeFromRaw(code);
			case GoStmt.GoWhile(cond, body):
				collectExprReads(cond);
				pushScope();
				collectStmtList(body);
				popScope();
			case GoStmt.GoLabeled(_, child):
				collectStmt(child);
			case GoStmt.GoRangeStmt(keyName, valueName, source, useShort, body):
				collectExprReads(source);
				pushScope();
				if (useShort) {
					if (keyName != null) {
						declareBinding(keyName);
					}
					if (valueName != null) {
						declareBinding(valueName);
					}
				}
				collectStmtList(body);
				popScope();
			case GoStmt.GoIf(cond, thenBody, elseBody):
				collectExprReads(cond);
				pushScope();
				collectStmtList(thenBody);
				popScope();
				if (elseBody != null) {
					pushScope();
					collectStmtList(elseBody);
					popScope();
				}
			case GoStmt.GoSwitch(value, cases, defaultBody):
				collectExprReads(value);
				for (entry in cases) {
					for (caseValue in entry.values) {
						collectExprReads(caseValue);
					}
					pushScope();
					collectStmtList(entry.body);
					popScope();
				}
				if (defaultBody != null) {
					pushScope();
					collectStmtList(defaultBody);
					popScope();
				}
			case GoStmt.GoTypeSwitch(value, bindingName, cases, defaultBody):
				collectExprReads(value);
				for (entry in cases) {
					pushScope();
					if (bindingName != null) {
						declareBinding(bindingName);
					}
					collectStmtList(entry.body);
					popScope();
				}
				if (defaultBody != null) {
					pushScope();
					if (bindingName != null) {
						declareBinding(bindingName);
					}
					collectStmtList(defaultBody);
					popScope();
				}
			case GoStmt.GoSelect(cases):
				for (entry in cases) {
					pushScope();
					collectSelectClause(entry.clause);
					collectStmtList(entry.body);
					popScope();
				}
			case GoStmt.GoBreak(_):
			case GoStmt.GoContinue:
			case GoStmt.GoReturn(expr):
				if (expr != null) {
					collectExprReads(expr);
				}
		}
	}

	function collectSelectClause(clause:GoSelectClause):Void {
		switch (clause) {
			case GoSelectClause.GoSelectSend(channel, value):
				collectExprReads(channel);
				collectExprReads(value);
			case GoSelectClause.GoSelectRecv(recv):
				collectExprReads(recv);
			case GoSelectClause.GoSelectRecvAssign(target, recv, useShort):
				collectExprReads(recv);
				if (useShort) {
					switch (target) {
						case GoExpr.GoIdent(name):
							declareBinding(name);
						case _:
							collectLValueReads(target);
					}
				} else {
					collectLValueReads(target);
				}
			case GoSelectClause.GoSelectRecvAssignOk(target, okTarget, recv, useShort):
				collectExprReads(recv);
				for (binding in [target, okTarget]) {
					if (useShort) {
						switch (binding) {
							case GoExpr.GoIdent(name):
								declareBinding(name);
							case _:
								collectLValueReads(binding);
						}
					} else {
						collectLValueReads(binding);
					}
				}
			case GoSelectClause.GoSelectDefault:
		}
	}

	function collectExprReads(expr:GoExpr):Void {
		switch (expr) {
			case GoExpr.GoIdent(name):
				addRead(resolveBinding(name));
			case GoExpr.GoSelector(target, _):
				collectExprReads(target);
			case GoExpr.GoIndex(target, index):
				collectExprReads(target);
				collectExprReads(index);
			case GoExpr.GoSlice(target, start, end):
				collectExprReads(target);
				if (start != null) {
					collectExprReads(start);
				}
				if (end != null) {
					collectExprReads(end);
				}
			case GoExpr.GoArrayLiteral(_, elements):
				for (element in elements) {
					collectExprReads(element);
				}
			case GoExpr.GoFuncLiteral(params, _, body):
				pushScope();
				for (param in params) {
					declareBinding(param.name);
				}
				collectStmtList(body);
				popScope();
			case GoExpr.GoRaw(code):
				markBindingsUnsafeFromRaw(code);
			case GoExpr.GoTypeAssert(inner, _):
				collectExprReads(inner);
			case GoExpr.GoRecvExpr(channel):
				collectExprReads(channel);
			case GoExpr.GoUnary(_, inner):
				collectExprReads(inner);
			case GoExpr.GoBinary(_, left, right):
				collectExprReads(left);
				collectExprReads(right);
			case GoExpr.GoCall(callee, args):
				collectExprReads(callee);
				for (arg in args) {
					collectExprReads(arg);
				}
			case _:
		}
	}

	function collectLValueReads(expr:GoExpr):Void {
		switch (expr) {
			case GoExpr.GoIdent(_):
			case GoExpr.GoSelector(target, _):
				collectExprReads(target);
			case GoExpr.GoIndex(target, index):
				collectExprReads(target);
				collectExprReads(index);
			case GoExpr.GoSlice(target, start, end):
				collectExprReads(target);
				if (start != null) {
					collectExprReads(start);
				}
				if (end != null) {
					collectExprReads(end);
				}
			case GoExpr.GoRaw(code):
				markBindingsUnsafeFromRaw(code);
			case _:
				collectExprReads(expr);
		}
	}

	function rewriteFile(file:GoFile):GoFile {
		return {
			packageName: file.packageName,
			imports: file.imports,
			decls: [for (decl in file.decls) rewriteDecl(decl)]
		};
	}

	function rewriteDecl(decl:GoDecl):GoDecl {
		return switch (decl) {
			case GoDecl.GoFuncDecl(name, receiver, params, results, body):
				pushScope();
				if (receiver != null) {
					declareBinding(receiver.name);
				}
				for (param in params) {
					declareBinding(param.name);
				}
				var rewrittenBody = rewriteStmtList(body);
				popScope();
				GoDecl.GoFuncDecl(name, receiver, params, results, rewrittenBody);
			case GoDecl.GoGlobalVarDecl(name, typeName, value):
				GoDecl.GoGlobalVarDecl(name, typeName, value == null ? null : rewriteExpr(value));
			case _:
				decl;
		};
	}

	function rewriteStmtList(stmts:Array<GoStmt>):Array<GoStmt> {
		var out = new Array<GoStmt>();
		var keptNoReadGuards = new Map<Int, Bool>();
		for (stmt in stmts) {
			var rewritten = rewriteStmt(stmt, keptNoReadGuards);
			if (rewritten != null) {
				out.push(rewritten);
			}
		}
		return out;
	}

	function rewriteStmt(stmt:GoStmt, keptNoReadGuards:Map<Int, Bool>):Null<GoStmt> {
		return switch (stmt) {
			case GoStmt.GoVarDecl(name, typeName, value, useShort):
				var rewrittenValue = value == null ? null : rewriteExpr(value);
				declareBinding(name);
				GoStmt.GoVarDecl(name, typeName, rewrittenValue, useShort);
			case GoStmt.GoMultiAssign(names, value, useShort):
				var rewrittenValue = rewriteExpr(value);
				if (useShort) {
					for (name in names) {
						declareBinding(name);
					}
				}
				GoStmt.GoMultiAssign(names, rewrittenValue, useShort);
			case GoStmt.GoAssign(GoExpr.GoIdent("_"), GoExpr.GoIdent(name)):
				var bindingId = resolveBinding(name);
				if (bindingId == null || unsafeBindings.exists(bindingId)) {
					stmt;
				} else {
					var reads = readCounts.exists(bindingId) ? readCounts.get(bindingId) : 0;
					if (reads > 0) {
						null;
					} else if (keptNoReadGuards.exists(bindingId)) {
						null;
					} else {
						keptNoReadGuards.set(bindingId, true);
						stmt;
					}
				}
			case GoStmt.GoAssign(left, right):
				GoStmt.GoAssign(rewriteExpr(left), rewriteExpr(right));
			case GoStmt.GoExprStmt(expr):
				GoStmt.GoExprStmt(rewriteExpr(expr));
			case GoStmt.GoGoStmt(call):
				GoStmt.GoGoStmt(rewriteExpr(call));
			case GoStmt.GoDeferStmt(call):
				GoStmt.GoDeferStmt(rewriteExpr(call));
			case GoStmt.GoSendStmt(channel, value):
				GoStmt.GoSendStmt(rewriteExpr(channel), rewriteExpr(value));
			case GoStmt.GoRaw(_):
				stmt;
			case GoStmt.GoWhile(cond, body):
				var rewrittenCond = rewriteExpr(cond);
				pushScope();
				var rewrittenBody = rewriteStmtList(body);
				popScope();
				GoStmt.GoWhile(rewrittenCond, rewrittenBody);
			case GoStmt.GoLabeled(label, child):
				var rewrittenChild = rewriteStmt(child, new Map<Int, Bool>());
				rewrittenChild == null ? null : GoStmt.GoLabeled(label, rewrittenChild);
			case GoStmt.GoRangeStmt(keyName, valueName, source, useShort, body):
				var rewrittenSource = rewriteExpr(source);
				pushScope();
				if (useShort) {
					if (keyName != null) {
						declareBinding(keyName);
					}
					if (valueName != null) {
						declareBinding(valueName);
					}
				}
				var rewrittenBody = rewriteStmtList(body);
				popScope();
				GoStmt.GoRangeStmt(keyName, valueName, rewrittenSource, useShort, rewrittenBody);
			case GoStmt.GoIf(cond, thenBody, elseBody):
				var rewrittenCond = rewriteExpr(cond);
				pushScope();
				var rewrittenThen = rewriteStmtList(thenBody);
				popScope();
				var rewrittenElse = if (elseBody == null) {
					null;
				} else {
					pushScope();
					var body = rewriteStmtList(elseBody);
					popScope();
					body;
				}
				GoStmt.GoIf(rewrittenCond, rewrittenThen, rewrittenElse);
			case GoStmt.GoSwitch(value, cases, defaultBody):
				var rewrittenValue = rewriteExpr(value);
				var rewrittenCases = new Array<GoSwitchCase>();
				for (entry in cases) {
					var rewrittenValues = [for (caseValue in entry.values) rewriteExpr(caseValue)];
					pushScope();
					var rewrittenBody = rewriteStmtList(entry.body);
					popScope();
					rewrittenCases.push({
						values: rewrittenValues,
						body: rewrittenBody
					});
				}
				var rewrittenDefault = if (defaultBody == null) {
					null;
				} else {
					pushScope();
					var body = rewriteStmtList(defaultBody);
					popScope();
					body;
				}
				GoStmt.GoSwitch(rewrittenValue, rewrittenCases, rewrittenDefault);
			case GoStmt.GoTypeSwitch(value, bindingName, cases, defaultBody):
				var rewrittenValue = rewriteExpr(value);
				var rewrittenCases = new Array<GoTypeSwitchCase>();
				for (entry in cases) {
					pushScope();
					if (bindingName != null) {
						declareBinding(bindingName);
					}
					var rewrittenBody = rewriteStmtList(entry.body);
					popScope();
					rewrittenCases.push({
						typeName: entry.typeName,
						body: rewrittenBody
					});
				}
				var rewrittenDefault = if (defaultBody == null) {
					null;
				} else {
					pushScope();
					if (bindingName != null) {
						declareBinding(bindingName);
					}
					var body = rewriteStmtList(defaultBody);
					popScope();
					body;
				}
				GoStmt.GoTypeSwitch(rewrittenValue, bindingName, rewrittenCases, rewrittenDefault);
			case GoStmt.GoSelect(cases):
				var rewrittenCases = new Array<GoSelectCase>();
				for (entry in cases) {
					pushScope();
					var rewrittenClause = rewriteSelectClause(entry.clause);
					var rewrittenBody = rewriteStmtList(entry.body);
					popScope();
					rewrittenCases.push({
						clause: rewrittenClause,
						body: rewrittenBody
					});
				}
				GoStmt.GoSelect(rewrittenCases);
			case GoStmt.GoBreak(_):
				stmt;
			case GoStmt.GoContinue:
				stmt;
			case GoStmt.GoReturn(expr):
				GoStmt.GoReturn(expr == null ? null : rewriteExpr(expr));
		};
	}

	function rewriteSelectClause(clause:GoSelectClause):GoSelectClause {
		return switch (clause) {
			case GoSelectClause.GoSelectSend(channel, value):
				GoSelectClause.GoSelectSend(rewriteExpr(channel), rewriteExpr(value));
			case GoSelectClause.GoSelectRecv(recv):
				GoSelectClause.GoSelectRecv(rewriteExpr(recv));
			case GoSelectClause.GoSelectRecvAssign(target, recv, useShort):
				var rewrittenTarget = rewriteExpr(target);
				var rewrittenRecv = rewriteExpr(recv);
				if (useShort) {
					switch (rewrittenTarget) {
						case GoExpr.GoIdent(name):
							declareBinding(name);
						case _:
					}
				}
				GoSelectClause.GoSelectRecvAssign(rewrittenTarget, rewrittenRecv, useShort);
			case GoSelectClause.GoSelectRecvAssignOk(target, okTarget, recv, useShort):
				var rewrittenTarget = rewriteExpr(target);
				var rewrittenOkTarget = rewriteExpr(okTarget);
				var rewrittenRecv = rewriteExpr(recv);
				if (useShort) {
					for (binding in [rewrittenTarget, rewrittenOkTarget]) {
						switch (binding) {
							case GoExpr.GoIdent(name):
								declareBinding(name);
							case _:
						}
					}
				}
				GoSelectClause.GoSelectRecvAssignOk(rewrittenTarget, rewrittenOkTarget, rewrittenRecv, useShort);
			case GoSelectClause.GoSelectDefault:
				GoSelectClause.GoSelectDefault;
		};
	}

	function rewriteExpr(expr:GoExpr):GoExpr {
		return switch (expr) {
			case GoExpr.GoSelector(target, field):
				GoExpr.GoSelector(rewriteExpr(target), field);
			case GoExpr.GoIndex(target, index):
				GoExpr.GoIndex(rewriteExpr(target), rewriteExpr(index));
			case GoExpr.GoSlice(target, start, end):
				GoExpr.GoSlice(rewriteExpr(target), start == null ? null : rewriteExpr(start), end == null ? null : rewriteExpr(end));
			case GoExpr.GoArrayLiteral(elementType, elements):
				GoExpr.GoArrayLiteral(elementType, [for (element in elements) rewriteExpr(element)]);
			case GoExpr.GoFuncLiteral(params, results, body):
				pushScope();
				for (param in params) {
					declareBinding(param.name);
				}
				var rewrittenBody = rewriteStmtList(body);
				popScope();
				GoExpr.GoFuncLiteral(params, results, rewrittenBody);
			case GoExpr.GoTypeAssert(inner, typeName):
				GoExpr.GoTypeAssert(rewriteExpr(inner), typeName);
			case GoExpr.GoRecvExpr(channel):
				GoExpr.GoRecvExpr(rewriteExpr(channel));
			case GoExpr.GoUnary(op, inner):
				GoExpr.GoUnary(op, rewriteExpr(inner));
			case GoExpr.GoBinary(op, left, right):
				GoExpr.GoBinary(op, rewriteExpr(left), rewriteExpr(right));
			case GoExpr.GoCall(callee, args):
				GoExpr.GoCall(rewriteExpr(callee), [for (arg in args) rewriteExpr(arg)]);
			case _:
				expr;
		};
	}
}
