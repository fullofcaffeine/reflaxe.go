class Main {
	static function main():Void {
		var map = new Map<String, Int>();
		map.set("alpha", 2);
		map.set("beta", 5);
		Sys.println("map.alpha=" + Std.string(map.get("alpha")));
		Sys.println("map.exists.beta0=" + Std.string(map.exists("beta")));
		Sys.println("map.remove.beta=" + Std.string(map.remove("beta")));
		Sys.println("map.exists.beta1=" + Std.string(map.exists("beta")));
	}
}
