package reflaxe.go.ast;

import reflaxe.go.ast.GoAST.GoStmt;

/**
	What: Go's "terminating statement" rule, applied to a lowered statement list.

	Why: Go rejects a function with results unless its body ends in a terminating
	statement ("missing return"). Go's rule is narrower than Haxe's: Haxe accepts
	a body that ends in an exhaustive enum `switch` whose every case returns, or in
	a `while (true)` loop that only leaves through `return`, but in Go a `switch`
	without `default` and a `for` loop with a condition are not terminating. The
	compiler must find these bodies to end them with an unreachable statement,
	and must leave every body that Go already accepts byte-for-byte unchanged.

	How: Follow the rule in the Go specification ("Terminating statements"):
	- `return` and a call to the built-in `panic` terminate.
	- An `if` terminates when it has an `else` and both branches terminate.
	- A `for` terminates when it has no condition, is not a range loop, and no
	  `break` refers to it.
	- A `switch` or type switch terminates when it has a `default`, every case
	  (with the default) ends in a terminating statement, and no `break` refers
	  to it. A `select` needs the same, without the `default` requirement.
	- A labeled statement terminates when the statement it labels terminates.
	The compiler-generated `while` form always prints a condition, so it never
	terminates. Raw Go text cannot be inspected, so a `GoRaw` statement counts as
	terminating: the answer is then "no change", which is what the compiler did
	before this check existed.
**/
class GoTermination {
	/** True when `stmts` ends in a statement that Go treats as terminating. */
	public static function endsInTerminatingStatement(stmts:Array<GoStmt>):Bool {
		return stmts.length > 0 && isTerminating(stmts[stmts.length - 1], null);
	}

	/** `label` is the label on `stmt` itself; a `break label` inside `stmt` leaves it. */
	static function isTerminating(stmt:GoStmt, label:Null<String>):Bool {
		return switch (stmt) {
			case GoReturn(_) | GoReturnValues(_):
				true;
			case GoRaw(_):
				// Unknown text: keep the output the compiler produced before this check.
				true;
			case GoExprStmt(GoCall(GoIdent("panic"), _)):
				true;
			case GoIf(_, thenBody, elseBody): elseBody != null && endsInTerminatingStatement(thenBody) && endsInTerminatingStatement(elseBody);
			case GoFor(_, null, _, body):
				!breaksOut(body, label, false);
			case GoSwitch(_, cases, defaultBody): defaultBody != null && casesTerminate([for (entry in cases) entry.body].concat([defaultBody]), label);
			case GoTypeSwitch(_, _, cases, defaultBody): defaultBody != null && casesTerminate([for (entry in cases) entry.body].concat([defaultBody]), label);
			case GoSelect(cases):
				casesTerminate([for (entry in cases) entry.body], label);
			case GoLabeled(name, inner):
				isTerminating(inner, name);
			case _:
				false;
		};
	}

	static function casesTerminate(bodies:Array<Array<GoStmt>>, label:Null<String>):Bool {
		for (body in bodies)
			if (!endsInTerminatingStatement(body) || breaksOut(body, label, false))
				return false;
		return true;
	}

	/**
		True when `stmts` contain a `break` that leaves the statement owning them:
		a `break` with that statement's `label`, or an unlabeled `break` that is not
		inside a nested `for`, `switch` or `select` (`nested`). Function literals are
		expressions and are not entered, because a `break` cannot cross them.
	**/
	static function breaksOut(stmts:Array<GoStmt>, label:Null<String>, nested:Bool):Bool {
		for (stmt in stmts) {
			var found = switch (stmt) {
				case GoBreak(target):
					target == null ? !nested : target == label;
				case GoIf(_, thenBody, elseBody): breaksOut(thenBody, label, nested) || (elseBody != null && breaksOut(elseBody, label, nested));
				case GoLabeled(_, inner):
					breaksOut([inner], label, nested);
				case GoWhile(_, body) | GoFor(_, _, _, body) | GoRangeStmt(_, _, _, _, body):
					breaksOut(body, label, true);
				case GoSwitch(_, cases, defaultBody): anyBreaksOut([for (entry in cases) entry.body],
						label) || (defaultBody != null && breaksOut(defaultBody, label, true));
				case GoTypeSwitch(_, _, cases, defaultBody): anyBreaksOut([for (entry in cases) entry.body],
						label) || (defaultBody != null && breaksOut(defaultBody, label, true));
				case GoSelect(cases):
					anyBreaksOut([for (entry in cases) entry.body], label);
				case _:
					false;
			};
			if (found)
				return true;
		}
		return false;
	}

	/** Bodies of a nested `switch` or `select`: only a labeled `break` can leave the outer statement. */
	static function anyBreaksOut(bodies:Array<Array<GoStmt>>, label:Null<String>):Bool {
		for (body in bodies)
			if (breaksOut(body, label, true))
				return true;
		return false;
	}
}
