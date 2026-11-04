/** Runtime classification and serialization must see the Haxe function, not its Go carrier. */
class Main {
	static function main() {
		final callback = () -> 1;
		Sys.println("type=" + Type.enumConstructor(Type.typeof(callback)));
		Sys.println("string=" + Std.string(callback));
		Sys.println("json=" + haxe.Json.stringify(callback));
		Sys.println("json-field=" + haxe.Json.stringify({call: callback}));
		Sys.println("json-array=" + haxe.Json.stringify([callback]));
		try {
			haxe.Serializer.run(callback);
		} catch (error:Dynamic) {
			Sys.println("serialize=" + Std.string(error));
		}
	}
}
