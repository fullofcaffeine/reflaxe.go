class Main {
	static function failInt():Int {
		throw "boom-int";
	}

	static function failString():String {
		throw "boom-string";
	}

	static function main() {
		var a = try failInt() catch (_:Dynamic) 7;
		var b = try failString() catch (_:Dynamic) "ok";
		Sys.println(Std.string(a));
		Sys.println(b);
	}
}
