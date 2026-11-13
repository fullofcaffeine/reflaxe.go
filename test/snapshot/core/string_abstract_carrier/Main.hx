/** A string carrier must retain pointer storage without erasing its value. */
abstract Token(String) {
	private inline function new(raw:String) {
		this = raw;
	}

	public static function make(raw:String):Token {
		return new Token(raw);
	}

	public inline function text():String {
		return this;
	}
}

/** Nested and applied abstract carriers preserve the same string storage. */
abstract Tagged<T>(Token) {
	public inline function new(raw:String) {
		this = Token.make(raw);
	}

	public inline function text():String {
		return this.text();
	}
}

/** A non-string abstract remains outside this string-only specialization. */
abstract Count(Int) {
	public inline function new(raw:Int) {
		this = raw;
	}

	public inline function value():Int {
		return this;
	}
}

/** Applied type arguments must distinguish strings from unsupported scalar storage. */
abstract Box<T>(T) {
	public inline function new(value:T) {
		this = value;
	}

	public inline function value():T {
		return this;
	}
}

typedef TextAlias = String;

class Main {
	static function main() {
		final token:Token = Token.make("exact bytes");
		final nested:Tagged<Int> = new Tagged<Int>("nested");
		final absent:Null<Token> = null;
		final nullBacked:Token = Token.make(null);
		final count = new Count(7);
		final boxed:Box<String> = new Box<String>("applied");
		final repeated:Box<Box<String>> = new Box<Box<String>>(new Box<String>("repeated"));
		if (repeated.value().value() != "repeated")
			throw "repeated applied backing changed";
		final nullableBox:Box<Null<TextAlias>> = new Box<Null<TextAlias>>(null);
		if (nullableBox.value() != null)
			throw "nullable alias changed";
		final boxedInt:Box<Int> = new Box<Int>(9);
		if (boxed.value() != "applied" || boxedInt.value() != 9)
			throw "applied backing changed";
		if (token.text() != "exact bytes" || nested.text() != "nested")
			throw "string carrier changed";
		if (absent != null || nullBacked.text() != null)
			throw "nil string changed";
		if (count.value() != 7)
			throw "non-string fallback changed";
		Sys.println("string abstract carrier: PASS");
	}
}
