package reflaxe.go.ast;

/**
	What: The closed set of Go assignment tokens used after an lvalue.
	Why: Compound assignments previously survived only as raw statement text, so
	transforms could not inspect either the target or the value.
	How: `GoStmt.GoAssign` and `GoSimpleStmt.GoSimpleAssign` store one left/right
	pair; an omitted operator remains ordinary `=` for source compatibility.
**/
enum abstract GoAssignmentOperator(String) {
	var Assign = "=";
	var AddAssign = "+=";
	var SubtractAssign = "-=";
	var MultiplyAssign = "*=";
	var DivideAssign = "/=";
	var RemainderAssign = "%=";
	var BitwiseAndAssign = "&=";
	var BitwiseOrAssign = "|=";
	var BitwiseXorAssign = "^=";
	var BitClearAssign = "&^=";
	var ShiftLeftAssign = "<<=";
	var ShiftRightAssign = ">>=";

	/** Return the target token; only the Go printer should normally need it. */
	public inline function token():String {
		return this;
	}
}
