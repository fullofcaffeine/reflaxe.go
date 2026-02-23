class Main {
	static function withTry(flag:Bool):String {
		var state = "start";
		try {
			state = "try";
			if (flag) {
				return "try:" + state;
			}
			throw "boom";
		} catch (e:Dynamic) {
			state = "catch";
			return "catch:" + state + ":" + Std.string(e);
		}
		state = "tail";
		return "tail:" + state;
	}

	static function main():Void {
		Sys.println(withTry(true));
		Sys.println(withTry(false));
	}
}
