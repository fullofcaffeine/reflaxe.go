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
		Sys.println("int.present=" + intBranch(4));
		Sys.println("int.missing=" + intBranch());
		Sys.println("float.present=" + Std.string(floatBranch(5.0)));
		Sys.println("float.missing=" + Std.string(floatBranch()));
		Sys.println("bool.present=" + boolBranch(true));
		Sys.println("bool.missing=" + boolBranch());
	}
}
