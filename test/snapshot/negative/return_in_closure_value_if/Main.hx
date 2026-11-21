class Main {
	static function describe(value:Int):String {
		Sys.println("value:" + (if (value == 0) return "zero" else "other"));
		return "done";
	}

	static function main():Void {
		Sys.println(describe(0));
	}
}
