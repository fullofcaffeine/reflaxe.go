class Main {
	static function probe(value:Dynamic):Void {
		try {
			throw value;
		} catch (caught:haxe.Exception) {
			Sys.println("msg=" + caught.message);
			Sys.println("isEx=" + Std.string(Std.isOfType(caught, haxe.Exception)));
		}
	}

	static function main() {
		probe("alpha");
		probe(42);
	}
}
