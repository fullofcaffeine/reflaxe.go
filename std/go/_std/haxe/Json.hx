package haxe;

@:go.import("hxrt")
@:go.package("hxrt")
private extern class HxrtJson {
	@:go.name("JsonParse")
	static function parse(source:String):Dynamic;
	@:go.name("JsonStringify")
	static function stringify(value:Dynamic, space:Null<String>):String;
}

/**
	What:
	- Owns the `haxe.Json` API surface selected by the Go target.

	Why:
	- The mainstream Haxe stdlib implementation cannot be used unchanged on `haxe.go`. because dynamic JSON conversion is implemented by the typed `hxrt` bridge over Go JSON values.

	How:
	- Keep the Haxe API in staged std and delegate parse/stringify behavior through typed runtime externs.
**/
class Json {
	public static inline function parse(text:String):Dynamic {
		return HxrtJson.parse(text);
	}

	public static inline function stringify(value:Dynamic, ?replacer:(key:Dynamic, value:Dynamic) -> Dynamic, ?space:String):String {
		return HxrtJson.stringify(value, space);
	}
}
