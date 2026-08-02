class Main {
	static function main() {
		var parsed = haxe.Json.parse("[1,true,\"x\"]");
		Sys.println(haxe.Json.stringify(parsed));
		var object:Dynamic = haxe.Json.parse("{\"name\":\"alpha\"}");
		var name:String = cast Reflect.field(object, "name");
		Sys.println(name);
	}
}
