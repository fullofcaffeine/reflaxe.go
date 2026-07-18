package reflaxe.go.ast;

import reflaxe.go.ast.GoAST.GoExpr;

/**
	What: The closed simple-statement subset admitted by a classic Go `for` clause.
	Why: Reusing arbitrary `GoStmt` values would allow blocks, returns, and other
	invalid syntax in initializer/post positions.
	How: Initializers may use every variant; the printer rejects a short declaration
	in the post position because the Go grammar forbids it there.
**/
enum GoSimpleStmt {
	GoSimpleDeclare(name:String, value:GoExpr);
	GoSimpleAssign(left:GoExpr, right:GoExpr, ?op:GoAssignmentOperator);
	GoSimpleIncDec(target:GoExpr, op:GoIncDecOperator);
	GoSimpleExpr(expr:GoExpr);
	GoSimpleSend(channel:GoExpr, value:GoExpr);
}
