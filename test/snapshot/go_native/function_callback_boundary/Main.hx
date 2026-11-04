/** Exact native callbacks retain invocation and Haxe aliases at export. */
@:go.import("sort")
extern class NativeSearch {
	@:go.name("Search")
	public static function search(length:Int, predicate:Int->Bool):Int;
}

class Main {
	static function main() {
		var threshold = 3;
		final predicate = index -> index >= threshold;
		final alias = predicate;
		Sys.println(NativeSearch.search(10, predicate));
		threshold = 4;
		Sys.println(NativeSearch.search(10, predicate));
		Sys.println(Reflect.compareMethods(predicate, alias));
		final search = NativeSearch.search;
		Sys.println(search(10, predicate));
	}
}
