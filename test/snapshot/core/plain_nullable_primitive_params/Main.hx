class Main {
	static function intBranch(value:Null<Int>):Int {
		if (value != null) {
			return value + 1;
		}
		return -1;
	}

	static function floatBranch(value:Null<Float>):Float {
		if (value != null) {
			return value / 2.0;
		}
		return -1.0;
	}

	static function boolBranch(value:Null<Bool>):String {
		if (value != null) {
			return value ? "true" : "false";
		}
		return "missing";
	}

	static function main() {
		Sys.println(intBranch(4));
		Sys.println(intBranch(null));
		Sys.println(floatBranch(5.0));
		Sys.println(floatBranch(null));
		Sys.println(boolBranch(true));
		Sys.println(boolBranch(null));
	}
}
