class Main {
	static function main() {
		var error = new haxe.ValueException("boom");
		Sys.println("msg=" + error.message);
		Sys.println("value=" + Std.string(error.value));
		Sys.println("str=" + error.toString());
	}
}
