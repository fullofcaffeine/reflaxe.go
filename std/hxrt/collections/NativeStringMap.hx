package hxrt.collections;

/**
	What
	- Typed bridge to the native string-keyed storage used by staged `StringMap`.

	Why
	- The public map API is ordinary Haxe library behavior, but Haxe source cannot
	  directly declare or operate on a Go `map[string]any` value.

	How
	- Expose only storage primitives and a deterministic key snapshot. Generic map
	  values use `Dynamic` solely at this erased runtime boundary and are cast back
	  to `T` immediately by `haxe.ds.StringMap`.
**/
@:go.import("hxrt")
@:go.package("hxrt")
extern class NativeStringMap {
	@:go.name("StringMapNew")
	public static function create():StringMapHandle;

	@:go.name("StringMapSet")
	public static function set(handle:StringMapHandle, key:String, value:Dynamic):Void;

	@:go.name("StringMapGet")
	public static function get(handle:StringMapHandle, key:String):Dynamic;

	@:go.name("StringMapExists")
	public static function exists(handle:StringMapHandle, key:String):Bool;

	@:go.name("StringMapRemove")
	public static function remove(handle:StringMapHandle, key:String):Bool;

	@:go.name("StringMapKeys")
	public static function keys(handle:StringMapHandle):Array<String>;

	@:go.name("StringMapClear")
	public static function clear(handle:StringMapHandle):Void;
}
