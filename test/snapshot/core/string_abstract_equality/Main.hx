abstract Token(String) {
	public inline function new(raw:String)
		this = raw;
}

class Box<T> {
	public final value:T;

	public function new(value:T) {
		this.value = value;
	}
}

class Main {
	static function same<T>(left:T, right:T):Bool {
		return left == right;
	}

	static function main() {
		final parts = ["ab", "c"];
		final a = new Token(parts.join(""));
		final b = new Token("a" + "bc".substr(0));
		final other = new Token("abd");
		Sys.println(a == b);
		Sys.println(a != b);
		Sys.println(a == other);
		final tokens = [a, other];
		Sys.println(tokens.indexOf(b));
		Sys.println(tokens.remove(b));
		Sys.println(tokens.length);
		Sys.println(same(a, b));
		Sys.println(new Box(a).value == b);
		final maybe:Null<Token> = b;
		Sys.println(a == maybe);
		final absent:Null<Token> = null;
		Sys.println(absent == null);
		Sys.println(a == absent);
		final record = {token: a};
		Sys.println(record.token == b);
		final read = () -> a;
		Sys.println(read() == b);
		final dynamicToken:Dynamic = a;
		Sys.println(dynamicToken == b);
	}
}
