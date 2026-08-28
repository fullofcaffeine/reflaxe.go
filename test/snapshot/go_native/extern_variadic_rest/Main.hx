@:go.import("fmt")
extern class FmtPkg {
	@:go.name("Sprint")
	static function sprint(values:haxe.Rest<Dynamic>):String;
}

class Main {
	static function main() {
		final rendered = FmtPkg.sprint("left", 7, "right");
		if (rendered != "left7right")
			throw "variadic extern arguments were not expanded as native values";
	}
}
