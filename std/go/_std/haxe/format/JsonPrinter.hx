package haxe.format;

/**
	What:
	- Owns the Go-target `haxe.format.JsonPrinter` compatibility surface.

	Why:
	- The mainstream Haxe stdlib implementation cannot be used unchanged on `haxe.go`. because the target centralizes dynamic JSON semantics in the typed `haxe.Json` to `hxrt` bridge.

	How:
	- Retain the expected printer entrypoint and delegate it to `haxe.Json.stringify`.
**/
class JsonPrinter {
	public static inline function print(value:Dynamic, ?replacer:(key:Dynamic, value:Dynamic) -> Dynamic, ?space:String):String {
		return haxe.Json.stringify(value, replacer, space);
	}
}
