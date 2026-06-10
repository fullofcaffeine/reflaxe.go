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
		Sys.println("int.present=" + intBranch(4));
		Sys.println("int.missing=" + intBranch(null));
		Sys.println("float.present=" + Std.string(floatBranch(5.0)));
		Sys.println("float.missing=" + Std.string(floatBranch(null)));
		Sys.println("bool.present=" + boolBranch(true));
		Sys.println("bool.missing=" + boolBranch(null));
	}
}
