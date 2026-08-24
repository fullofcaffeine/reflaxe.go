class Main {
	static function main():Void {
		final source = [1, 2];
		final tail = [3, 2];
		final combined = source.concat(tail);
		Sys.println(combined.join(","));
		Sys.println(combined.indexOf(2));
		Sys.println(combined.indexOf(2, 2));
		Sys.println(combined.indexOf(2, -1));
		Sys.println(combined.indexOf(2, -8));
		Sys.println(combined.indexOf(9));
		source[0] = 9;
		tail[0] = 8;
		Sys.println(combined.join(","));

		final rebuiltA = "alpha".substr(0, 1);
		final words = [rebuiltA, "b"];
		Sys.println(words.indexOf("a"));
	}
}
