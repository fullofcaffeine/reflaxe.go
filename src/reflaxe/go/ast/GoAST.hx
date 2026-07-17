package reflaxe.go.ast;

/**
	What: One generated Go source file before target printing.
	Why: Package and import identity must remain validated and traversable through
	all AST passes, rather than becoming trusted target text at file assembly.
	How: Lowering supplies typed names/paths, transforms preserve them, and the
	printer owns package/import syntax.
**/
typedef GoFile = {
	final packageName:GoPackageName;
	final imports:Array<GoImportPath>;
	final decls:Array<GoDecl>;
}

/**
	What: A named structural Go type slot reused by parameters, receivers, and
	struct fields.
	Why: Those positions share printing/import behavior even though an empty name
	is meaningful only for embedded fields.
	How: The containing declaration supplies the context; `typeName` is always a
	validated `GoType`.
**/
typedef GoParam = {
	final name:String;
	final typeName:GoType;
}

/**
	What: One method signature carried by a declared Go interface.
	Why: Structural parameter/result types keep method-set imports visible.
	How: The interface printer renders the typed lists in declaration context.
**/
typedef GoInterfaceMethod = {
	final name:String;
	final params:Array<GoParam>;
	final results:Array<GoType>;
}

typedef GoSwitchCase = {
	final values:Array<GoExpr>;
	final body:Array<GoStmt>;
}

/**
	What: One typed case arm in a Go type switch.
	Why: Case types can reference imports and must not bypass type validation.
	How: Transforms preserve the `GoType`; the switch printer owns case syntax.
**/
typedef GoTypeSwitchCase = {
	final typeName:GoType;
	final body:Array<GoStmt>;
}

typedef GoSelectCase = {
	final clause:GoSelectClause;
	final body:Array<GoStmt>;
}

enum GoSelectClause {
	GoSelectSend(channel:GoExpr, value:GoExpr);
	GoSelectRecv(recv:GoExpr);
	GoSelectRecvAssign(target:GoExpr, recv:GoExpr, useShort:Bool);

	/**
		Why: non-blocking channel lowering must distinguish a temporarily empty
		channel from a selected receive on a drained closed channel.
		What: a select receive assignment with both value and comma-ok targets.
		How: the printer emits `case value, received := <-channel:` (or `=` when
		`useShort` is false), and every AST transform rewrites both targets.
	**/
	GoSelectRecvAssignOk(target:GoExpr, okTarget:GoExpr, recv:GoExpr, useShort:Bool);

	GoSelectDefault;
}

enum GoDecl {
	GoInterfaceDecl(name:String, methods:Array<GoInterfaceMethod>);
	GoStructDecl(name:String, fields:Array<GoParam>);
	GoGlobalVarDecl(name:String, typeName:GoType, value:Null<GoExpr>);
	GoFuncDecl(name:String, receiver:Null<GoParam>, params:Array<GoParam>, results:Array<GoType>, body:Array<GoStmt>);
}

enum GoStmt {
	GoVarDecl(name:String, typeName:Null<GoType>, value:Null<GoExpr>, useShort:Bool);

	/**
		What: Assign one multi-result Go expression to two or more named targets.
		Why: Typed runtime boundaries commonly return `(value, error)` or
		`(value, eof, error)`; representing that syntax as `GoRaw` hides expressions
		from import analysis and transform passes.
		How: `useShort` selects `:=` when every name is new in the current scope or
		`=` when every name already exists. The printer emits the comma-separated
		names before the typed value expression; mixed new/existing short declarations
		are intentionally outside this node's contract.
	**/
	GoMultiAssign(names:Array<String>, value:GoExpr, useShort:Bool);

	GoAssign(left:GoExpr, right:GoExpr);
	GoExprStmt(expr:GoExpr);
	GoGoStmt(call:GoExpr);
	GoDeferStmt(call:GoExpr);
	GoSendStmt(channel:GoExpr, value:GoExpr);
	GoRaw(code:String);
	GoWhile(cond:GoExpr, body:Array<GoStmt>);
	GoLabeled(label:String, stmt:GoStmt);
	GoRangeStmt(keyName:Null<String>, valueName:Null<String>, source:GoExpr, useShort:Bool, body:Array<GoStmt>);
	GoIf(cond:GoExpr, thenBody:Array<GoStmt>, elseBody:Null<Array<GoStmt>>);
	GoSwitch(value:GoExpr, cases:Array<GoSwitchCase>, defaultBody:Null<Array<GoStmt>>);
	GoTypeSwitch(value:GoExpr, bindingName:Null<String>, cases:Array<GoTypeSwitchCase>, defaultBody:Null<Array<GoStmt>>);
	GoSelect(cases:Array<GoSelectCase>);
	GoBreak(label:Null<String>);
	GoContinue;
	GoReturn(expr:Null<GoExpr>);
}

enum GoExpr {
	GoIdent(name:String);
	GoIntLiteral(value:Int);
	GoFloatLiteral(value:String);
	GoBoolLiteral(value:Bool);
	GoStringLiteral(value:String);
	GoNil;
	GoSelector(target:GoExpr, field:String);
	GoIndex(target:GoExpr, index:GoExpr);
	GoSlice(target:GoExpr, start:Null<GoExpr>, end:Null<GoExpr>);
	GoArrayLiteral(elementType:GoType, elements:Array<GoExpr>);

	/**
		What: A typed Go slice allocation with a required length and optional capacity.
		Why: Boxing native slices should preserve preallocation without hiding the
		element type or size expressions inside `GoRaw`.
		How: The printer emits `make([]T, length)` or `make([]T, length, capacity)`,
		while transforms traverse both size expressions and import analysis traverses
		the structural element type.
	**/
	GoMakeSlice(elementType:GoType, length:GoExpr, capacity:Null<GoExpr>);

	GoFuncLiteral(params:Array<GoParam>, results:Array<GoType>, body:Array<GoStmt>);
	GoRaw(code:String);
	GoTypeAssert(expr:GoExpr, typeName:GoType);
	GoRecvExpr(channel:GoExpr);
	GoUnary(op:GoUnaryOperator, expr:GoExpr);
	GoBinary(op:GoBinaryOperator, left:GoExpr, right:GoExpr);
	GoCall(callee:GoExpr, args:Array<GoExpr>);
}
