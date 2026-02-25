package haxe;

@:go.import("hxrt")
@:go.package("hxrt")
private extern class HxrtJson {
	@:go.name("JsonParse")
	static function parse(source:String):Dynamic;
	@:go.name("JsonStringify")
	static function stringify(value:Dynamic):String;
}

/**
	Go staged stdlib override.
	Owns the Json API surface in std/_std while delegating runtime behavior to hxrt.
**/
class Json {
	public static inline function parse(text:String):Dynamic {
		return HxrtJson.parse(text);
	}

	public static inline function stringify(value:Dynamic, ?replacer:(key:Dynamic, value:Dynamic) -> Dynamic, ?space:String):String {
		return HxrtJson.stringify(value);
	}
}
