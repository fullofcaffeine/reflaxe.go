class Main {
	static function main() {
		var parsed = haxe.Json.parse("[1,true,\"x\"]");
		Sys.println(haxe.Json.stringify(parsed));
		var pretty = haxe.Json.parse("{\"items\":[1,2]}");
		Sys.println(haxe.Json.stringify(pretty, null, "  "));
		var object:Dynamic = haxe.Json.parse("{\"name\":\"alpha\"}");
		var name:String = cast Reflect.field(object, "name");
		Sys.println(name);
	}
}
