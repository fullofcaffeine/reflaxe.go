package reflaxe.go.ast;

/**
	What: The closed set of Go unary-expression operators emitted by this IR.
	Why: A typed operator lets transforms distinguish logical negation, address
	operations, dereference, and receive without inspecting printed text.
	How: Lowering selects a named value; legacy tokens cross the validating parser
	and the printer owns their spelling.
**/
enum abstract GoUnaryOperator(String) {
	var Positive = "+";
	var Negate = "-";
	var LogicalNot = "!";
	var BitwiseNot = "^";
	var Dereference = "*";
	var AddressOf = "&";
	var Receive = "<-";

	/** Convert a legacy operator token without admitting an unchecked fallback. */
	@:from
	public static function parse(token:String):GoUnaryOperator {
		return switch (token) {
			case "+": Positive;
			case "-": Negate;
			case "!": LogicalNot;
			case "^": BitwiseNot;
			case "*": Dereference;
			case "&": AddressOf;
			case "<-": Receive;
			case _: throw 'Invalid Go unary operator "' + token + '"';
		};
	}

	/** Return the target token; only the Go printer should normally need this. */
	public inline function token():String {
		return this;
	}
}
