class Main {
	static function main() {
		var encoded = haxe.Serializer.run({value: 1});
		Sys.println(encoded.length > 0);
	}
}
