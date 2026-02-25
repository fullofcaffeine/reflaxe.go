package haxe.format;

/**
	Go staged stdlib override.
	Printer behavior delegates to haxe.Json.stringify.
**/
class JsonPrinter {
	public static inline function print(value:Dynamic, ?replacer:(key:Dynamic, value:Dynamic) -> Dynamic, ?space:String):String {
		return haxe.Json.stringify(value, replacer, space);
	}
}
