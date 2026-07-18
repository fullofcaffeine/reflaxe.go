package reflaxe.go.compiler;

#if macro
import reflaxe.go.ast.GoAST.GoDecl;
import reflaxe.go.ast.GoAST.GoExpr;
import reflaxe.go.ast.GoAST.GoSelectClause;
import reflaxe.go.ast.GoAST.GoStmt;
import reflaxe.go.ast.GoAssignmentOperator;
import reflaxe.go.ast.GoBinaryOperator;
import reflaxe.go.ast.GoBuiltinType;
import reflaxe.go.ast.GoChannelDirection;
import reflaxe.go.ast.GoCompositeElement;
import reflaxe.go.ast.GoImportPath;
import reflaxe.go.ast.GoIncDecOperator;
import reflaxe.go.ast.GoPackageName;
import reflaxe.go.ast.GoSimpleStmt;
import reflaxe.go.ast.GoType;
import reflaxe.go.ast.GoUnaryOperator;

class GoTestAstFixtureEmitter {
	/**
		What: Candidate package imports required only by synthetic typed-AST fixtures.
		Why: Import filtering must prove package use from structural `GoType` nodes;
		the fixture should exercise that production path instead of embedding raw Go.
		How: The compiler adds these paths to the normal candidate set, after which
		structural usage analysis keeps or removes them exactly like real extern imports.
	**/
	public static function imports(testCase:String):Array<String> {
		return switch (testCase) {
			case "typed_types": ["sync/atomic"];
			case _: [];
		};
	}

	public static function emit(testCase:String):Null<Array<GoDecl>> {
		return switch (testCase) {
			case "go_defer":
				[
					GoDecl.GoFuncDecl("hxrt__test_ast_go_defer_stmt_smoke", null, [], [], [
						GoStmt.GoVarDecl("fn", "func()", GoExpr.GoFuncLiteral([], [], []), true),
						GoStmt.GoDeferStmt(GoExpr.GoCall(GoExpr.GoIdent("fn"), [])),
						GoStmt.GoGoStmt(GoExpr.GoCall(GoExpr.GoIdent("fn"), []))
					])
				];
			case "send_recv":
				[
					GoDecl.GoFuncDecl("hxrt__test_ast_send_recv_stmt_smoke", null, [], [], [
						GoStmt.GoVarDecl("ch", null, GoExpr.GoCall(GoExpr.GoIdent("make"), [GoExpr.GoRaw("chan int"), GoExpr.GoIntLiteral(1)]), true),
						GoStmt.GoSendStmt(GoExpr.GoIdent("ch"), GoExpr.GoIntLiteral(7)),
						GoStmt.GoAssign(GoExpr.GoIdent("_"), GoExpr.GoRecvExpr(GoExpr.GoIdent("ch")))
					])
				];
			case "select":
				[
					GoDecl.GoFuncDecl("hxrt__test_ast_select_stmt_smoke", null, [], [], [
						GoStmt.GoVarDecl("in", null, GoExpr.GoCall(GoExpr.GoIdent("make"), [GoExpr.GoRaw("chan int"), GoExpr.GoIntLiteral(1)]), true),
						GoStmt.GoVarDecl("out", null, GoExpr.GoCall(GoExpr.GoIdent("make"), [GoExpr.GoRaw("chan int"), GoExpr.GoIntLiteral(1)]), true),
						GoStmt.GoSelect([
							{
								clause: GoSelectClause.GoSelectSend(GoExpr.GoIdent("out"), GoExpr.GoIntLiteral(1)),
								body: [GoStmt.GoAssign(GoExpr.GoIdent("_"), GoExpr.GoIntLiteral(11))]
							},
							{
								clause: GoSelectClause.GoSelectRecvAssign(GoExpr.GoIdent("value"), GoExpr.GoRecvExpr(GoExpr.GoIdent("in")), true),
								body: [GoStmt.GoAssign(GoExpr.GoIdent("_"), GoExpr.GoIdent("value"))]
							},
							{
								clause: GoSelectClause.GoSelectRecvAssignOk(GoExpr.GoIdent("value"), GoExpr.GoIdent("received"),
									GoExpr.GoRecvExpr(GoExpr.GoIdent("in")), true),
								body: [
									GoStmt.GoAssign(GoExpr.GoIdent("_"), GoExpr.GoIdent("value")),
									GoStmt.GoAssign(GoExpr.GoIdent("_"), GoExpr.GoIdent("received"))
								]
							},
							{
								clause: GoSelectClause.GoSelectRecv(GoExpr.GoRecvExpr(GoExpr.GoIdent("in"))),
								body: [GoStmt.GoAssign(GoExpr.GoIdent("_"), GoExpr.GoIntLiteral(22))]
							},
							{
								clause: GoSelectClause.GoSelectDefault,
								body: [GoStmt.GoAssign(GoExpr.GoIdent("_"), GoExpr.GoIntLiteral(0))]
							}
						])
					])
				];
			case "range":
				var intType = GoType.builtin(GoBuiltinType.Int);
				var mapType = GoType.map(GoType.builtin(GoBuiltinType.StringType), intType);
				[
					GoDecl.GoFuncDecl("hxrt__test_ast_range_stmt_smoke", null, [], [], [
						GoStmt.GoVarDecl("items", null, GoExpr.GoCompositeLiteral(mapType, [
							GoCompositeElement.GoCompositeKeyValue(GoExpr.GoStringLiteral("a"), GoExpr.GoIntLiteral(1)),
							GoCompositeElement.GoCompositeKeyValue(GoExpr.GoStringLiteral("b"), GoExpr.GoIntLiteral(2))
						]), true),
						GoStmt.GoRangeStmt("key", "value", GoExpr.GoIdent("items"), true,
							[
								GoStmt.GoAssign(GoExpr.GoIdent("_"), GoExpr.GoIdent("key")),
								GoStmt.GoAssign(GoExpr.GoIdent("_"), GoExpr.GoIdent("value"))
							]),
						GoStmt.GoVarDecl("seenKey", "string", null, false),
						GoStmt.GoRangeStmt("seenKey", null, GoExpr.GoIdent("items"), false, [GoStmt.GoAssign(GoExpr.GoIdent("_"), GoExpr.GoIdent("seenKey"))]),
						GoStmt.GoRangeStmt("index", null, GoExpr.GoCompositeLiteral(GoType.slice(intType), [
							GoCompositeElement.GoCompositeValue(GoExpr.GoIntLiteral(1)),
							GoCompositeElement.GoCompositeValue(GoExpr.GoIntLiteral(2)),
							GoCompositeElement.GoCompositeValue(GoExpr.GoIntLiteral(3))
						]), true, [GoStmt.GoAssign(GoExpr.GoIdent("_"), GoExpr.GoIdent("index"))])
					])
				];
			case "multi_assign":
				[
					GoDecl.GoFuncDecl("hxrt__test_ast_multi_assign_stmt_smoke", null, [], [], [
						GoStmt.GoVarDecl("items", "map[string]int", null, false),
						GoStmt.GoMultiAssign(["value", "found"], GoExpr.GoIndex(GoExpr.GoIdent("items"), GoExpr.GoStringLiteral("present")), true),
						GoStmt.GoAssign(GoExpr.GoIdent("_"), GoExpr.GoIdent("value")),
						GoStmt.GoAssign(GoExpr.GoIdent("_"), GoExpr.GoIdent("found")),
						GoStmt.GoVarDecl("missing", "int", null, false),
						GoStmt.GoMultiAssign(["missing", "found"], GoExpr.GoIndex(GoExpr.GoIdent("items"), GoExpr.GoStringLiteral("missing")), false),
						GoStmt.GoAssign(GoExpr.GoIdent("_"), GoExpr.GoIdent("missing")),
						GoStmt.GoAssign(GoExpr.GoIdent("_"), GoExpr.GoIdent("found"))
					])
				];
			case "typed_types":
				var intType = GoType.builtin(GoBuiltinType.Int);
				var stringType = GoType.builtin(GoBuiltinType.StringType);
				var localType = GoType.named("hxrt__test_ast_Local");
				var constraintType = GoType.interfaceType([
					GoType.interfaceMethod("Apply", [intType], [stringType, GoType.builtin(GoBuiltinType.Error)])
				]);
				[
					GoDecl.GoStructDecl("hxrt__test_ast_Local", []),
					GoDecl.GoStructDecl("hxrt__test_ast_type_shapes", [
						{name: "Builtin", typeName: GoType.builtin(GoBuiltinType.Bool)},
						{name: "Named", typeName: localType},
						{name: "Pointer", typeName: GoType.pointer(stringType)},
						{name: "Slice", typeName: GoType.slice(GoType.pointer(localType))},
						{name: "Array", typeName: GoType.array(3, GoType.builtin(GoBuiltinType.Byte))},
						{name: "Map", typeName: GoType.map(stringType, GoType.slice(intType))},
						{name: "Channel", typeName: GoType.channel(GoChannelDirection.Bidirectional, intType)},
						{name: "Receive", typeName: GoType.channel(GoChannelDirection.ReceiveOnly, stringType)},
						{name: "Send", typeName: GoType.channel(GoChannelDirection.SendOnly, GoType.builtin(GoBuiltinType.Bool))},
						{
							name: "Callback",
							typeName: GoType.functionType([intType, GoType.variadic(stringType)],
								[GoType.builtin(GoBuiltinType.Bool), GoType.builtin(GoBuiltinType.Error)])
						},
						{name: "Constraint", typeName: constraintType},
						{name: "Empty", typeName: GoType.emptyInterface()},
						{
							name: "Generic",
							typeName: GoType.generic(GoType.qualified(GoPackageName.named("atomic"), "Pointer"), [intType])
						}
					]),
					GoDecl.GoFuncDecl("hxrt__test_ast_typed_operator_smoke", null, [
						{name: "left", typeName: intType},
						{name: "right", typeName: intType},
						{name: "boxed", typeName: GoType.builtin(GoBuiltinType.AnyType)}
					], [GoType.builtin(GoBuiltinType.Bool)], [
						GoStmt.GoAssign(GoExpr.GoIdent("_"), GoExpr.GoUnary(GoUnaryOperator.Negate, GoExpr.GoIdent("left"))),
						GoStmt.GoAssign(GoExpr.GoIdent("_"), GoExpr.GoUnary(GoUnaryOperator.BitwiseNot, GoExpr.GoIdent("right"))),
						GoStmt.GoAssign(GoExpr.GoIdent("_"),
							GoExpr.GoCompositeLiteral(GoType.slice(intType),
								[
									GoCompositeElement.GoCompositeValue(GoExpr.GoIdent("left")),
									GoCompositeElement.GoCompositeValue(GoExpr.GoIdent("right"))
								])),
						GoStmt.GoAssign(GoExpr.GoIdent("_"), GoExpr.GoMakeSlice(intType, GoExpr.GoIntLiteral(0), GoExpr.GoIntLiteral(2))),
						GoStmt.GoAssign(GoExpr.GoIdent("_"), GoExpr.GoTypeAssert(GoExpr.GoIdent("boxed"), constraintType)),
						GoStmt.GoReturn(GoExpr.GoBinary(GoBinaryOperator.LogicalAnd,
							GoExpr.GoBinary(GoBinaryOperator.LessThan, GoExpr.GoIdent("left"), GoExpr.GoIdent("right")),
							GoExpr.GoBinary(GoBinaryOperator.NotEqual, GoExpr.GoIdent("left"), GoExpr.GoIdent("right"))))
					])
				];
			case "structured_composites":
				var intType = GoType.builtin(GoBuiltinType.Int);
				var stringType = GoType.builtin(GoBuiltinType.StringType);
				var pointType = GoType.named("hxrt__test_ast_Point");
				var sliceType = GoType.slice(intType);
				var arrayType = GoType.array(3, intType);
				var mapType = GoType.map(stringType, intType);
				[
					GoDecl.GoStructDecl("hxrt__test_ast_Point", [
						{name: "X", typeName: intType},
						{name: "Y", typeName: intType},
						{name: "Label", typeName: GoType.pointer(stringType)}
					]),
					GoDecl.GoFuncDecl("hxrt__test_ast_structured_composite_control_smoke", null, [], [], [
						GoStmt.GoVarDecl("point", null, GoExpr.GoUnary(GoUnaryOperator.AddressOf, GoExpr.GoCompositeLiteral(pointType, [
							GoCompositeElement.GoCompositeField("X", GoExpr.GoIntLiteral(1)),
							GoCompositeElement.GoCompositeField("Y", GoExpr.GoIntLiteral(2)),
							GoCompositeElement.GoCompositeField("Label",
								GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringConcatAny"), [GoExpr.GoStringLiteral("a"), GoExpr.GoStringLiteral("b")]))
						])), true),
						GoStmt.GoVarDecl("values", null, GoExpr.GoCompositeLiteral(sliceType, [
							GoCompositeElement.GoCompositeValue(GoExpr.GoIntLiteral(1)),
							GoCompositeElement.GoCompositeValue(GoExpr.GoIntLiteral(2)),
							GoCompositeElement.GoCompositeValue(GoExpr.GoIntLiteral(3))
						]), true),
						GoStmt.GoVarDecl("indexed", null, GoExpr.GoCompositeLiteral(arrayType, [
							GoCompositeElement.GoCompositeKeyValue(GoExpr.GoIntLiteral(0), GoExpr.GoIntLiteral(4)),
							GoCompositeElement.GoCompositeKeyValue(GoExpr.GoIntLiteral(2), GoExpr.GoIntLiteral(6))
						]), true),
						GoStmt.GoVarDecl("lookup", null, GoExpr.GoCompositeLiteral(mapType,
							[
								GoCompositeElement.GoCompositeKeyValue(GoExpr.GoStringLiteral("one"), GoExpr.GoIntLiteral(1)),
								GoCompositeElement.GoCompositeKeyValue(GoExpr.GoStringLiteral("two"), GoExpr.GoIntLiteral(2))
							]),
							true),
						GoStmt.GoVarDecl("total", null, GoExpr.GoIntLiteral(0), true),
						GoStmt.GoFor(GoSimpleStmt.GoSimpleDeclare("index", GoExpr.GoIntLiteral(0)),
							GoExpr.GoBinary(GoBinaryOperator.LessThan, GoExpr.GoIdent("index"),
								GoExpr.GoCall(GoExpr.GoIdent("len"), [GoExpr.GoIdent("values")])),
							GoSimpleStmt.GoSimpleIncDec(GoExpr.GoIdent("index"), GoIncDecOperator.Increment),
							[
								GoStmt.GoAssign(GoExpr.GoIdent("total"), GoExpr.GoIndex(GoExpr.GoIdent("values"), GoExpr.GoIdent("index")),
									GoAssignmentOperator.AddAssign)
							]),
						GoStmt.GoIncDec(GoExpr.GoIdent("total"), GoIncDecOperator.Increment),
						GoStmt.GoIncDec(GoExpr.GoIdent("total"), GoIncDecOperator.Decrement),
						GoStmt.GoAssign(GoExpr.GoIdent("_"), GoExpr.GoIdent("point")),
						GoStmt.GoAssign(GoExpr.GoIdent("_"), GoExpr.GoIdent("indexed")),
						GoStmt.GoAssign(GoExpr.GoIdent("_"), GoExpr.GoIdent("lookup")),
						GoStmt.GoAssign(GoExpr.GoIdent("_"), GoExpr.GoIdent("total"))
					])
				];
			case "invalid_composite_literal":
				[
					GoDecl.GoGlobalVarDecl("invalid", GoType.builtin(GoBuiltinType.Int), GoExpr.GoCompositeLiteral(GoType.builtin(GoBuiltinType.Int), []))
				];
			case "invalid_for_post":
				[
					GoDecl.GoFuncDecl("invalid", null, [], [], [
						GoStmt.GoFor(null, null, GoSimpleStmt.GoSimpleDeclare("value", GoExpr.GoIntLiteral(0)), [])
					])
				];
			case "invalid_type":
				GoType.parse("map[string");
				[];
			case "invalid_type_combination":
				GoType.map(GoType.slice(GoType.builtin(GoBuiltinType.Int)), GoType.builtin(GoBuiltinType.StringType));
				[];
			case "invalid_operator":
				GoBinaryOperator.parse("=>");
				[];
			case "invalid_import_path":
				GoImportPath.parse("bad import path");
				[];
			case "invalid_import_path_character":
				GoImportPath.parse("example.com/bad?path");
				[];
			case _:
				null;
		};
	}
}
#end
