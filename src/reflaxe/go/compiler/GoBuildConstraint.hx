package reflaxe.go.compiler;

/**
	A validated Go file-selection expression.

	Why: Go build constraints have target syntax that Haxe cannot express with an
	ordinary declaration. Free-form header text would bypass the typed output
	boundary and could inject arbitrary source.
	What: This value accepts only identifiers, `!`, `&&`, `||`, and balanced
	parentheses in Go's build-constraint grammar.
	How: A small recursive-descent validator consumes a closed token stream, then
	renders those validated tokens with deterministic spacing.
**/
class GoBuildConstraint {
	final tokens:Array<GoBuildConstraintToken>;

	function new(tokens:Array<GoBuildConstraintToken>) {
		this.tokens = tokens;
	}

	public static function parse(source:String):GoBuildConstraint {
		final parser = new GoBuildConstraintParser(source);
		return new GoBuildConstraint(parser.parse());
	}

	public function render():String {
		final output = new StringBuf();
		var previous:Null<GoBuildConstraintToken> = null;
		for (token in tokens) {
			switch (token) {
				case And:
					output.add(" && ");
				case Or:
					output.add(" || ");
				case RightParen:
					output.add(")");
				case LeftParen:
					if (previous != null)
						switch (previous) {
							case Identifier(_) | RightParen: output.add(" ");
							case _:
						}
					output.add("(");
				case Not:
					output.add("!");
				case Identifier(value):
					if (previous != null)
						switch (previous) {
							case Identifier(_) | RightParen: output.add(" ");
							case _:
						}
					output.add(value);
			}
			previous = token;
		}
		return output.toString();
	}
}

private enum GoBuildConstraintToken {
	Identifier(value:String);
	Not;
	And;
	Or;
	LeftParen;
	RightParen;
}

private class GoBuildConstraintParser {
	final source:String;
	final tokens:Array<GoBuildConstraintToken> = [];
	var index = 0;

	public function new(source:String) {
		this.source = source == null ? "" : source;
	}

	public function parse():Array<GoBuildConstraintToken> {
		tokenize();
		if (tokens.length == 0)
			invalid();
		parseOr();
		if (index != tokens.length)
			invalid();
		return tokens.copy();
	}

	function tokenize():Void {
		var offset = 0;
		while (offset < source.length) {
			final code = source.charCodeAt(offset);
			if (code != null && (code == 32 || code == 9 || code == 10 || code == 13)) {
				offset++;
				continue;
			}
			final tail = source.substr(offset);
			if (StringTools.startsWith(tail, "&&")) {
				tokens.push(And);
				offset += 2;
			} else if (StringTools.startsWith(tail, "||")) {
				tokens.push(Or);
				offset += 2;
			} else
				switch (source.charAt(offset)) {
					case "!":
						tokens.push(Not);
						offset++;
					case "(":
						tokens.push(LeftParen);
						offset++;
					case ")":
						tokens.push(RightParen);
						offset++;
					case _:
						final identifier = ~/^[A-Za-z0-9_.]+/;
						if (!identifier.match(tail))
							invalid();
						final value = identifier.matched(0);
						tokens.push(Identifier(value));
						offset += value.length;
				}
		}
	}

	function parseOr():Void {
		parseAnd();
		while (accept(Or))
			parseAnd();
	}

	function parseAnd():Void {
		parseUnary();
		while (accept(And))
			parseUnary();
	}

	function parseUnary():Void {
		if (accept(Not)) {
			parseUnary();
			return;
		}
		if (accept(LeftParen)) {
			parseOr();
			if (!accept(RightParen))
				invalid();
			return;
		}
		if (index >= tokens.length)
			invalid();
		switch (tokens[index]) {
			case Identifier(_):
				index++;
			case _:
				invalid();
		}
	}

	function accept(expected:GoBuildConstraintToken):Bool {
		if (index >= tokens.length || Type.enumIndex(tokens[index]) != Type.enumIndex(expected))
			return false;
		index++;
		return true;
	}

	function invalid():Void {
		throw new haxe.Exception("the value is not a valid Go build constraint");
	}
}
