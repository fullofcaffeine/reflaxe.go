package reflaxe.go.compiler;

/**
	Validated, factored selection conditions for generated Go files.
	Authored tags cross the parser once; generated helpers compose the same typed
	conditions without enumerating build configurations. Always has no header and
	Never is omitted by the file owner; neither is rendered as a Go tag literal.
**/
class GoBuildConstraint {
	final selection:GoBuildSelection;

	function new(selection:GoBuildSelection) {
		this.selection = selection;
	}

	public static function parse(source:String):GoBuildConstraint {
		return new GoBuildConstraintParser(source).parse();
	}

	public static function always():GoBuildConstraint
		return new GoBuildConstraint(Always);

	public static function never():GoBuildConstraint
		return new GoBuildConstraint(Never);

	/** A generated tag is subject to the same closed grammar as authored input. */
	public static function tag(name:String):GoBuildConstraint {
		if (name == null || !~/^[A-Za-z0-9_.]+$/.match(name))
			throw new haxe.Exception("the value is not a valid Go build constraint");
		return new GoBuildConstraint(Tag(name));
	}

	public function isAlways():Bool
		return switch (selection) {
			case Always: true;
			case _: false;
		};

	public function isNever():Bool
		return switch (selection) {
			case Never: true;
			case _: false;
		};

	public function negate():GoBuildConstraint {
		return switch (selection) {
			case Always: never();
			case Never: always();
			case Negation(value): value;
			case _: new GoBuildConstraint(Negation(this));
		};
	}

	public function and(other:GoBuildConstraint):GoBuildConstraint
		return combine(other, true);

	public function or(other:GoBuildConstraint):GoBuildConstraint
		return combine(other, false);

	/** Canonical fact keys include constants, which cannot be Go header expressions. */
	public function key():String {
		return switch (selection) {
			case Always: "@always";
			case Never: "@never";
			case _: render();
		};
	}

	function combine(other:GoBuildConstraint, conjunction:Bool):GoBuildConstraint {
		var terms = new Array<GoBuildConstraint>();
		for (value in [this, other])
			switch (value.selection) {
				case Conjunction(children) if (conjunction):
					terms = terms.concat(children);
				case Disjunction(children) if (!conjunction):
					terms = terms.concat(children);
				case _:
					terms.push(value);
			}
		var unique = new Map<String, GoBuildConstraint>();
		for (term in terms) {
			if ((conjunction && term.isNever()) || (!conjunction && term.isAlways()))
				return conjunction ? never() : always();
			if ((conjunction && term.isAlways()) || (!conjunction && term.isNever()))
				continue;
			if (unique.exists(term.negate().key()))
				return conjunction ? never() : always();
			unique.set(term.key(), term);
		}
		var keys = [for (key in unique.keys()) key];
		keys.sort((left, right) -> left < right ? -1 : left > right ? 1 : 0);
		var values = [for (key in keys) unique.get(key)];
		return switch (values.length) {
			case 0: conjunction ? always() : never();
			case 1: values[0];
			case _: new GoBuildConstraint(conjunction ? Conjunction(values) : Disjunction(values));
		};
	}

	/** Render a nonconstant condition with precedence-preserving parentheses. */
	public function render():String
		return renderAt(0);

	function renderAt(parentPrecedence:Int):String {
		var precedence = 4;
		var text = switch (selection) {
			case Always | Never:
				throw new haxe.Exception("constant build selection must be handled by the file owner");
			case Tag(name): name;
			case Negation(value):
				precedence = 3;
				"!" + value.renderAt(precedence);
			case Conjunction(values):
				precedence = 2;
				[for (value in values) value.renderAt(precedence)].join(" && ");
			case Disjunction(values):
				precedence = 1;
				[for (value in values) value.renderAt(precedence)].join(" || ");
		};
		return precedence < parentPrecedence ? "(" + text + ")" : text;
	}
}

private enum GoBuildSelection {
	Always;
	Never;
	Tag(name:String);
	Negation(value:GoBuildConstraint);
	Conjunction(values:Array<GoBuildConstraint>);
	Disjunction(values:Array<GoBuildConstraint>);
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

	public function parse():GoBuildConstraint {
		tokenize();
		if (tokens.length == 0)
			invalid();
		var value = parseOr();
		if (index != tokens.length)
			invalid();
		return value;
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

	function parseOr():GoBuildConstraint {
		var value = parseAnd();
		while (accept(Or))
			value = value.or(parseAnd());
		return value;
	}

	function parseAnd():GoBuildConstraint {
		var value = parseUnary();
		while (accept(And))
			value = value.and(parseUnary());
		return value;
	}

	function parseUnary():GoBuildConstraint {
		if (accept(Not))
			return parseUnary().negate();
		if (accept(LeftParen)) {
			var value = parseOr();
			if (!accept(RightParen))
				invalid();
			return value;
		}
		if (index >= tokens.length)
			invalid();
		return switch (tokens[index]) {
			case Identifier(name):
				index++;
				GoBuildConstraint.tag(name);
			case _:
				invalid();
				GoBuildConstraint.never();
		};
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
