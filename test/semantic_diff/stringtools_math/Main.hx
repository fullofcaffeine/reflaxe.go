class Main {
	static function main() {
		var s = "  hi  ";
		var x = 3.8;

		Sys.println(StringTools.trim(s));
		Sys.println(StringTools.startsWith("hello", "he"));
		Sys.println(StringTools.replace("a-b-c", "-", ":"));

		Sys.println(Math.floor(x));
		Sys.println(Math.ceil(x));
		Sys.println(Math.round(x));
		Sys.println(Math.fround(x));
		Sys.println(Math.abs(x));
		Sys.println(Math.min(x, 2.1));
		Sys.println(Math.max(x, 2.1));
	}
}
