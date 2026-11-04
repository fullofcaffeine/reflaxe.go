/** Native callbacks must translate string representations in both directions. */
@:go.import("regexp")
@:go.name("Regexp")
extern class NativeRegexp {
	@:go.name("MustCompile") public static function compile(pattern:String):NativeRegexp;
	@:go.name("ReplaceAllStringFunc") public function replace(text:String, callback:String->String):String;
}

class Main {
	static function main() {
		var calls = 0;
		final callback = function(value:String):String {
			calls++;
			return "[" + value + "]";
		};
		final expression = NativeRegexp.compile("x");
		Sys.println(expression.replace("x x", callback));
		final replace = expression.replace;
		Sys.println(replace("x", callback));
		Sys.println(calls);
		final missing:String->String = null;
		Sys.println(expression.replace("z", missing));
		Sys.println(replace("z", missing));
	}
}
