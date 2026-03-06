class Main {
	static function main() {
		var error = new haxe.ValueException("boom");
		Sys.println(error.message);
	}
}
