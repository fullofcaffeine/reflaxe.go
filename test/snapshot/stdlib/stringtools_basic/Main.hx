class Main {
	static function main() {
		var s = "  hi  ";
		Sys.println(StringTools.trim(s));
		Sys.println(StringTools.startsWith("hello", "he"));
		Sys.println(StringTools.replace("a-b-c", "-", ":"));
	}
}
