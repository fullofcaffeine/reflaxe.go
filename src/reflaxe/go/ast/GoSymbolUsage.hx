package reflaxe.go.ast;

import reflaxe.go.ast.GoAST.GoDecl;
import reflaxe.go.ast.GoAST.GoExpr;
import reflaxe.go.ast.GoAST.GoParam;
import reflaxe.go.ast.GoAST.GoSelectClause;
import reflaxe.go.ast.GoAST.GoStmt;
import reflaxe.go.ast.GoCompositeElement;
import reflaxe.go.ast.GoSimpleStmt;

/**
	What: The package names and `package.Member` symbols that a group of Go
	declarations refers to.

	Why: Before a file is printed, the compiler drops each candidate import that
	the file does not use and turns on the runtime features whose `hxrt.*`
	symbols the file names. Asking the syntax tree about one name at a time
	walked every file once per candidate import and once per runtime symbol.

	How: `collect` walks the declarations once and records every name found:
	- A selector `x.f` whose target is the identifier `x` records `x` and `x.f`.
	- Identifier text and raw Go text (`GoRaw`, comments and strings included)
	  record each word that is followed by a dot, and each `word.word` pair.
	  A word is a maximal run of regular-expression word characters, and white
	  space may surround the dot.
	- A type records the package of each named type inside it, and the
	  `word.word` pairs of its printed form.
	String literals, the type of a short declaration (`x := value` prints no
	type), and declared names such as labels, parameters and fields are not read.
	`uses` then answers with one map lookup.

	For text, the result is the same as testing the text with `\bname\s*\.`
	for a package and `\bpkg\s*\.\s*Member\b` for a symbol, when the key is
	made of ASCII letters, digits and `_`: such a pattern matches exactly where
	the key is a whole word followed by a dot (and the member word). The scan
	uses the same regular-expression engine, so `\b`, `\w` and `\s` agree.
	`GoSymbolUsageContract` checks this against those patterns.
**/
class GoSymbolUsage {
	/**
		A word followed by optional white space and a dot: `fmt.` or `fmt .`.
		The `g` flag makes `map` visit every match. Scans use `map`, not a
		`matchSub` loop, because the Haxe 4.3 interpreter mixes byte and character
		offsets in `matchSub` once the text holds a non-ASCII character.
	**/
	static final packageReference = ~/\b(\w+)\s*\./g;

	/** A word, a dot and a second word; the lookahead lets `a.b.c` yield `a.b` and `b.c`. */
	static final memberReference = ~/\b(\w+)(?=\s*\.\s*(\w+)\b)/g;

	/** The accepted key shapes. ASCII word characters are word characters in every regular-expression mode. */
	static final keyShape = ~/^[A-Za-z0-9_]+(\.[A-Za-z0-9_]+)?$/;

	final names:Map<String, Bool>;

	function new() {
		names = new Map<String, Bool>();
	}

	/** Walk `decls` once and record every package name and qualified symbol they refer to. */
	public static function collect(decls:Array<GoDecl>):GoSymbolUsage {
		final usage = new GoSymbolUsage();
		for (decl in decls)
			usage.visitDecl(decl);
		return usage;
	}

	/**
		Whether the declarations refer to the package `key` (for example `fmt`) or
		to the qualified symbol `key` (for example `hxrt.ArraySort`).

		Throws for any other key shape, because the word scan answers only for
		keys made of ASCII letters, digits and `_`.
	**/
	public function uses(key:String):Bool {
		if (key == null || !keyShape.match(key))
			throw 'GoSymbolUsage key must be "name" or "name.Member": ' + key;
		return names.exists(key);
	}

	function visitDecl(decl:GoDecl):Void {
		switch (decl) {
			case GoInterfaceDecl(_, methods):
				for (method in methods) {
					visitParams(method.params);
					visitTypes(method.results);
				}
			case GoStructDecl(_, fields):
				visitParams(fields);
			case GoGlobalVarDecl(_, typeName, value):
				visitType(typeName);
				visitOptionalExpr(value);
			case GoFuncDecl(_, receiver, parameters, results, body):
				if (receiver != null)
					visitType(receiver.typeName);
				visitParams(parameters);
				visitTypes(results);
				visitStmts(body);
		}
	}

	function visitStmt(stmt:GoStmt):Void {
		switch (stmt) {
			case GoVarDecl(_, typeName, value, useShort):
				// A short declaration prints only `name := value`; its retained Haxe
				// type guides lowering but does not create a Go import reference.
				if (!useShort)
					visitType(typeName);
				visitOptionalExpr(value);
			case GoMultiAssign(_, value, _):
				visitExpr(value);
			case GoAssign(left, right, _):
				visitExpr(left);
				visitExpr(right);
			case GoIncDec(target, _):
				visitExpr(target);
			case GoExprStmt(value) | GoGoStmt(value) | GoDeferStmt(value):
				visitExpr(value);
			case GoSendStmt(channel, value):
				visitExpr(channel);
				visitExpr(value);
			case GoRaw(code):
				visitText(code);
			case GoWhile(cond, body):
				visitExpr(cond);
				visitStmts(body);
			case GoFor(initializer, condition, post, body):
				visitSimpleStmt(initializer);
				visitOptionalExpr(condition);
				visitSimpleStmt(post);
				visitStmts(body);
			case GoLabeled(_, child):
				visitStmt(child);
			case GoRangeStmt(_, _, source, _, body):
				visitExpr(source);
				visitStmts(body);
			case GoIf(cond, thenBody, elseBody):
				visitExpr(cond);
				visitStmts(thenBody);
				visitStmts(elseBody);
			case GoSwitch(value, cases, defaultBody):
				visitExpr(value);
				for (switchCase in cases) {
					visitExprs(switchCase.values);
					visitStmts(switchCase.body);
				}
				visitStmts(defaultBody);
			case GoTypeSwitch(value, _, cases, defaultBody):
				visitExpr(value);
				for (typeCase in cases) {
					visitType(typeCase.typeName);
					visitStmts(typeCase.body);
				}
				visitStmts(defaultBody);
			case GoSelect(cases):
				for (selectCase in cases) {
					visitSelectClause(selectCase.clause);
					visitStmts(selectCase.body);
				}
			case GoBreak(_), GoContinue:
			case GoReturn(value):
				visitOptionalExpr(value);
			case GoReturnValues(values):
				visitExprs(values);
		}
	}

	function visitSimpleStmt(stmt:Null<GoSimpleStmt>):Void {
		if (stmt == null)
			return;
		switch (stmt) {
			case GoSimpleDeclare(_, value) | GoSimpleExpr(value):
				visitExpr(value);
			case GoSimpleAssign(left, right, _):
				visitExpr(left);
				visitExpr(right);
			case GoSimpleIncDec(target, _):
				visitExpr(target);
			case GoSimpleSend(channel, value):
				visitExpr(channel);
				visitExpr(value);
		}
	}

	function visitSelectClause(clause:GoSelectClause):Void {
		switch (clause) {
			case GoSelectSend(channel, value):
				visitExpr(channel);
				visitExpr(value);
			case GoSelectRecv(recv):
				visitExpr(recv);
			case GoSelectRecvAssign(target, recv, _):
				visitExpr(target);
				visitExpr(recv);
			case GoSelectRecvAssignOk(target, okTarget, recv, _):
				visitExpr(target);
				visitExpr(okTarget);
				visitExpr(recv);
			case GoSelectDefault:
		}
	}

	function visitExpr(expr:GoExpr):Void {
		switch (expr) {
			// A standalone identifier can be a local variable with the same name as
			// a package, so only qualified text inside it counts.
			case GoIdent(name):
				visitText(name);
			case GoIntLiteral(_), GoFloatLiteral(_), GoBoolLiteral(_), GoStringLiteral(_), GoNil:
			case GoSelector(target, field):
				switch (target) {
					case GoIdent(name):
						names.set(name, true);
						names.set(name + "." + field, true);
					case _:
				}
				visitExpr(target);
			case GoInstantiate(target, arguments):
				visitExpr(target);
				visitTypes(arguments);
			case GoIndex(target, index):
				visitExpr(target);
				visitExpr(index);
			case GoSlice(target, start, end):
				visitExpr(target);
				visitOptionalExpr(start);
				visitOptionalExpr(end);
			case GoCompositeLiteral(typeName, elements):
				visitType(typeName);
				for (element in elements)
					switch (element) {
						case GoCompositeValue(value), GoCompositeField(_, value):
							visitExpr(value);
						case GoCompositeKeyValue(key, value):
							visitExpr(key);
							visitExpr(value);
					}
			case GoMakeSlice(elementType, length, capacity):
				visitType(elementType);
				visitExpr(length);
				visitOptionalExpr(capacity);
			case GoFuncLiteral(parameters, results, body):
				visitParams(parameters);
				visitTypes(results);
				visitStmts(body);
			case GoRaw(code):
				visitText(code);
			case GoTypeAssert(inner, typeName):
				visitExpr(inner);
				visitType(typeName);
			case GoRecvExpr(channel):
				visitExpr(channel);
			case GoUnary(_, inner):
				visitExpr(inner);
			case GoBinary(_, left, right):
				visitExpr(left);
				visitExpr(right);
			case GoCall(callee, args) | GoVariadicCall(callee, args):
				visitExpr(callee);
				visitExprs(args);
		}
	}

	function visitStmts(values:Null<Array<GoStmt>>):Void {
		if (values != null)
			for (value in values)
				visitStmt(value);
	}

	function visitExprs(values:Array<GoExpr>):Void {
		for (value in values)
			visitExpr(value);
	}

	function visitOptionalExpr(value:Null<GoExpr>):Void {
		if (value != null)
			visitExpr(value);
	}

	function visitParams(values:Array<GoParam>):Void {
		for (value in values)
			visitType(value.typeName);
	}

	function visitTypes(values:Array<GoType>):Void {
		for (value in values)
			visitType(value);
	}

	function visitType(typeName:Null<GoType>):Void {
		if (typeName == null)
			return;
		typeName.addPackageNames(names);
		final printed = typeName.render();
		if (printed.indexOf(".") >= 0)
			addMemberReferences(printed);
	}

	/** Raw Go or identifier text; neither pattern can match without a dot. */
	function visitText(code:Null<String>):Void {
		if (code == null || code.indexOf(".") < 0)
			return;
		packageReference.map(code, notePackageReference);
		addMemberReferences(code);
	}

	function addMemberReferences(code:String):Void {
		memberReference.map(code, noteMemberReference);
	}

	function notePackageReference(match:EReg):String {
		names.set(match.matched(1), true);
		return "";
	}

	function noteMemberReference(match:EReg):String {
		names.set(match.matched(1) + "." + match.matched(2), true);
		return "";
	}
}
