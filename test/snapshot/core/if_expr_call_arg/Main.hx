class Main {
	static function id(value:Int):Int {
		return value;
	}

	static function main() {
		var cond = true;
		Sys.println(Std.string(id(if (cond) 7 else 9)));

		cond = false;
		Sys.println(Std.string(id(if (cond) 7 else 9)));
	}
}
