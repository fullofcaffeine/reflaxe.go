class Main {
	static function main() {
		var value:Dynamic = haxe.Json.parse('{"name":"reflaxe.go"}');
		if (value == null) {
			throw "unexpected";
		}
	}
}
