package reflaxe.go.ast;

/**
	What: Go's two postfix increment/decrement statement forms.
	Why: `++` and `--` are statements in Go, not expressions, and raw spellings hide
	their target from lvalue and transform analysis.
	How: `GoIncDec` and typed `for` clauses retain the direction; the printer owns
	the final token.
**/
enum abstract GoIncDecOperator(String) {
	var Increment = "++";
	var Decrement = "--";

	/** Return the target token; only the Go printer should normally need it. */
	public inline function token():String {
		return this;
	}
}
