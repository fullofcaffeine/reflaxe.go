class Main {
	static function describe(value:Int):String {
		final selected = switch value {
			case 0:
				return "zero";
			case 1:
				"one";
			case _:
				if (value < 0)
					return "negative";
				"many";
		};
		return "selected:" + selected;
	}

	static function main():Void {
		Sys.println(describe(0));
		Sys.println(describe(1));
		Sys.println(describe(-1));
		Sys.println(describe(2));
	}
}
