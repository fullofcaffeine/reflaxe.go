class Box {
	public var s:Null<String>;

	public function new(v:Null<String>) {
		s = v;
	}
}

class Main {
	static function show(v:Null<String>):String {
		return v == null ? "null" : v;
	}

	static function main() {
		var a = new Box(null);
		var b = new Box("ok");

		Sys.println(show(a.s));
		Sys.println(show(b.s));

		b.s = null;
		Sys.println(show(b.s));
	}
}
