class Main {
	static function restDigest(...args:Int):String {
		var rest:haxe.Rest<Int> = args;
		var copied = rest.toArray();
		var appended = rest.append(9);
		var prepended = rest.prepend(-1);
		return copied.length + ":" + (copied.length > 0 ? copied[0] : -99) + ":" + (copied.length > 0 ? copied[copied.length - 1] : -99) + "|append="
			+ appended[appended.length - 1] + "|prepend=" + prepended[0];
	}

	static function main() {
		var map:haxe.Constraints.IMap<String, Int> = new haxe.ds.StringMap<Int>();
		map.set("alpha", 1);
		var copied = map.copy();
		copied.set("copied", 7);
		Sys.println("imap=" + map.exists("copied") + ":" + copied.exists("copied") + ":" + copied.get("copied"));
		Sys.println("rest=" + restDigest(3, 1, 4));
		Sys.println("rest.empty=" + restDigest());
	}
}
