enum Color {
	Red;
	RGB(r:Int, g:Int, b:Int);
}

class Plain {
	public function new() {}
}

class Described {
	public function new() {}

	public function toString():String {
		return "described";
	}
}

/**
	Generated-shape contract for Std.string: every enum carrier gets
	`HxEnumValue()`, a class without `toString()` gets `HxClassName()`, and a
	class with `toString()` keeps its `String()` adapter.
**/
class Main {
	static function main() {
		var rgb:Dynamic = RGB(1, 2, 3);
		Sys.println(Std.string(Red));
		Sys.println(Std.string(rgb));
		Sys.println(Std.string(new Plain()));
		Sys.println(Std.string(new Described()));
		Sys.println(Std.string({color: Red, plain: new Plain()}));
		Sys.println(Std.string([Red, RGB(4, 5, 6)]));
	}
}
