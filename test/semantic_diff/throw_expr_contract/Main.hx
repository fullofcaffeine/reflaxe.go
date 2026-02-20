class Main {
	static function choose(flag:Bool):Int {
		return flag ? 1 : throw "boom";
	}

	static function main() {
		try {
			Sys.println("a=" + choose(true));
		} catch (e:Dynamic) {
			Sys.println("a_err=" + Std.string(e));
		}
		try {
			Sys.println("b=" + choose(false));
		} catch (e:Dynamic) {
			Sys.println("b_err=" + Std.string(e));
		}
	}
}
