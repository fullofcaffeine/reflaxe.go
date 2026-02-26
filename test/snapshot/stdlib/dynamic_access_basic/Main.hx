class Main {
	static function main():Void {
		var values:haxe.DynamicAccess<Int> = {};
		values["left"] = 2;
		values["right"] = 5;

		var sum = values["left"] + values["right"];

		Sys.println(Std.string(sum));
		Sys.println(Std.string(values.exists("missing")));
	}
}
