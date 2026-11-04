class Main {
	static function invoke(values:Array<Int->Int>, index:Int):Int {
		return values[index](20);
	}

	static function main() {
		Sys.println(invoke([value -> value + 3], 0));
		// Dynamic is intentional: this tests recovery at an explicit typed boundary.
		final boxed:Dynamic = function(value:Int):Int return value * 2;
		Sys.println((cast boxed : Int->Int)(11));
	}
}
