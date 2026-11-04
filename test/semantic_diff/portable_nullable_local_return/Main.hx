class Main {
	static function integer(value:Null<Int>):Int {
		final result = value == null ? 9 : value;
		return result;
	}

	static function decimal(value:Null<Float>):Float {
		final result = value == null ? 1.5 : value;
		return result;
	}

	static function flag(value:Null<Bool>):Bool {
		final result = value == null ? true : value;
		return result;
	}

	static function redirected(value:Null<Int>):Int {
		try {
			final result = value == null ? 7 : value;
			return result;
		} catch (_:haxe.Exception) {
			return -1;
		}
	}

	static function nullable(value:Null<Int>):Null<Int> {
		final result = value;
		return result;
	}

	static function main() {
		Sys.println(integer(null));
		Sys.println(integer(0));
		Sys.println(integer(4));
		Sys.println(decimal(null));
		Sys.println(decimal(0.0));
		Sys.println(decimal(2.5));
		Sys.println(flag(null));
		Sys.println(flag(false));
		Sys.println(flag(true));
		Sys.println(redirected(null));
		Sys.println(redirected(0));
		Sys.println(nullable(null) == null);
		Sys.println(nullable(3));
	}
}
