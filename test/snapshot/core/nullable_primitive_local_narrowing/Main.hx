class Main {
	static function intBranch(?value:Int):Int {
		if (value != null) {
			var narrowed = value;
			return narrowed + 1;
		}
		return -1;
	}

	static function floatBranch(?value:Float):Float {
		if (value != null) {
			var narrowed = value;
			return narrowed / 2.0;
		}
		return -1.0;
	}

	static function boolBranch(?value:Bool):String {
		if (value != null) {
			var narrowed = value;
			return narrowed ? "true" : "false";
		}
		return "missing";
	}

	static function main() {
		Sys.println(intBranch(4));
		Sys.println(intBranch());
		Sys.println(floatBranch(5.0));
		Sys.println(floatBranch());
		Sys.println(boolBranch(true));
		Sys.println(boolBranch());
	}
}
