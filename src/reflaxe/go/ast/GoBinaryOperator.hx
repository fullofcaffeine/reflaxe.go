package reflaxe.go.ast;

/**
	What: The closed set of Go binary-expression operators emitted by this IR.
	Why: Free-form operator text can create malformed syntax or bypass exhaustive
	operator transforms.
	How: Lowering selects a named value. The temporary string conversion validates
	legacy call sites, and the printer alone reads the target token.
**/
enum abstract GoBinaryOperator(String) {
	var Add = "+";
	var Subtract = "-";
	var Multiply = "*";
	var Divide = "/";
	var Remainder = "%";
	var Equal = "==";
	var NotEqual = "!=";
	var GreaterThan = ">";
	var GreaterThanOrEqual = ">=";
	var LessThan = "<";
	var LessThanOrEqual = "<=";
	var LogicalAnd = "&&";
	var LogicalOr = "||";
	var BitwiseAnd = "&";
	var BitwiseOr = "|";
	var BitwiseXor = "^";
	var BitClear = "&^";
	var ShiftLeft = "<<";
	var ShiftRight = ">>";

	/** Convert a legacy operator token without admitting an unchecked fallback. */
	@:from
	public static function parse(token:String):GoBinaryOperator {
		return switch (token) {
			case "+": Add;
			case "-": Subtract;
			case "*": Multiply;
			case "/": Divide;
			case "%": Remainder;
			case "==": Equal;
			case "!=": NotEqual;
			case ">": GreaterThan;
			case ">=": GreaterThanOrEqual;
			case "<": LessThan;
			case "<=": LessThanOrEqual;
			case "&&": LogicalAnd;
			case "||": LogicalOr;
			case "&": BitwiseAnd;
			case "|": BitwiseOr;
			case "^": BitwiseXor;
			case "&^": BitClear;
			case "<<": ShiftLeft;
			case ">>": ShiftRight;
			case _: throw 'Invalid Go binary operator "' + token + '"';
		};
	}

	/** Return the target token; only the Go printer should normally need this. */
	public inline function token():String {
		return this;
	}
}
