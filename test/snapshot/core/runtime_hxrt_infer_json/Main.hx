class Main {
	static function main() {
		var value:Dynamic = haxe.Json.parse('{"name":"reflaxe.go"}');
		Sys.println(Std.string(value != null));
	}
}
