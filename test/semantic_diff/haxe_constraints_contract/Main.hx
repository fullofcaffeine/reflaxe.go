class Main {
	static function digestMap(map:haxe.Constraints.IMap<String, Int>):String {
		var copied:haxe.Constraints.IMap<String, Int> = map.copy();
		copied.set("copied", 99);
		return "alpha=" + map.get("alpha") + "|beta=" + map.get("beta") + "|origHasCopied=" + map.exists("copied") + "||copyHasCopied="
			+ copied.exists("copied") + "|copyValue=" + copied.get("copied");
	}

	static function main() {
		var direct:haxe.Constraints.IMap<String, Int> = new haxe.ds.StringMap<Int>();
		direct.set("beta", 2);
		direct.set("alpha", 1);
		Sys.println("stringMap=" + digestMap(direct));

		var ints:haxe.Constraints.IMap<Int, String> = new haxe.ds.IntMap<String>();
		ints.set(4, "four");
		ints.set(7, "seven");
		var copied:haxe.Constraints.IMap<Int, String> = ints.copy();
		copied.remove(4);
		Sys.println("intMap.orig=" + ints.exists(4) + ":" + ints.get(7));
		Sys.println("intMap.copy=" + copied.exists(4) + ":" + copied.get(7));
	}
}
