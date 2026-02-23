package reflaxe.go.compiler;

#if macro
import reflaxe.go.ast.GoAST.GoDecl;
import reflaxe.go.ast.GoAST.GoExpr;
import reflaxe.go.ast.GoAST.GoSelectClause;
import reflaxe.go.ast.GoAST.GoStmt;

class GoTestAstFixtureEmitter {
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
				[
					GoDecl.GoFuncDecl("hxrt__test_ast_range_stmt_smoke", null, [], [], [
						GoStmt.GoVarDecl("items", null, GoExpr.GoRaw("map[string]int{\"a\": 1, \"b\": 2}"), true),
						GoStmt.GoRangeStmt("key", "value", GoExpr.GoIdent("items"), true,
							[
								GoStmt.GoAssign(GoExpr.GoIdent("_"), GoExpr.GoIdent("key")),
								GoStmt.GoAssign(GoExpr.GoIdent("_"), GoExpr.GoIdent("value"))
							]),
						GoStmt.GoVarDecl("seenKey", "string", null, false),
						GoStmt.GoRangeStmt("seenKey", null, GoExpr.GoIdent("items"), false, [GoStmt.GoAssign(GoExpr.GoIdent("_"), GoExpr.GoIdent("seenKey"))]),
						GoStmt.GoRangeStmt("index", null, GoExpr.GoRaw("[]int{1, 2, 3}"), true, [GoStmt.GoAssign(GoExpr.GoIdent("_"), GoExpr.GoIdent("index"))])
					])
				];
			case _:
				null;
		};
	}
}
#end
