class Main {
	static function describe(value:Int):String {
		Sys.println("value:" + switch value {
			case 0: return "zero";
			case 1: "one";
			case _: "many";
		});
		return "done";
	}

	static function main():Void {
		Sys.println(describe(0));
	}
}
