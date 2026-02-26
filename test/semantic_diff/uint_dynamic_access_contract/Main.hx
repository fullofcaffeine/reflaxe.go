class Main {
	static function main():Void {
		var base:UInt = cast 3;
		var scaled:UInt = base * 7;
		var shifted:UInt = scaled >> 1;
		Sys.println("uint=" + Std.string(base) + "," + Std.string(scaled) + "," + Std.string(shifted));

		var values:haxe.DynamicAccess<Int> = {};
		values["left"] = base;
		values.set("right", shifted);

		Sys.println("dyn_exists=" + Std.string(values.exists("left")) + "," + Std.string(values.exists("missing")));
		Sys.println("dyn_values=" + Std.string(values["left"]) + "," + Std.string(values.get("right")));
	}
}
