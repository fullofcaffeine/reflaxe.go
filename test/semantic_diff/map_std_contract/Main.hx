class Main {
	static function main():Void {
		var map = new Map<String, Int>();
		map.set("alpha", 2);
		map.set("beta", 5);

		var alpha = map.get("alpha");
		var missing = map.get("missing");

		Sys.println("map.alpha=" + Std.string(alpha));
		Sys.println("map.alpha.null=" + (alpha == null));
		Sys.println("map.missing=" + Std.string(missing));
		Sys.println("map.missing.null=" + (missing == null));
		Sys.println("map.exists.missing=" + map.exists("missing"));
	}
}
