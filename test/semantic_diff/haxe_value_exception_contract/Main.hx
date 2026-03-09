class Main {
	static function main() {
		var error = new haxe.ValueException("boom");
		Sys.println("msg=" + error.message);
		Sys.println("value=" + Std.string(error.value));
		Sys.println("str=" + error.toString());
		Sys.println("isEx=" + Std.string(Std.isOfType(error, haxe.Exception)));
		Sys.println("isValueEx=" + Std.string(Std.isOfType(error, haxe.ValueException)));

		try {
			throw error;
		} catch (caught:haxe.Exception) {
			Sys.println("caught=" + caught.message);
		}
	}
}
