package hxrt.collections;

import go.NativeSlice;

/**
	What
	- Typed-handle bridge to the native identity-keyed storage used by staged
	  `ObjectMap`.

	Why
	- Go must derive and retain stable reference identities for object keys. Both
	  keys and values are generic Haxe values, so they are deliberately `Dynamic`
	  only at this narrow erased runtime boundary.

	How
	- Keep the carrier typed as `ObjectMapHandle`, expose storage primitives plus a
	  deterministic native-slice key snapshot, and let `haxe.ds.ObjectMap`
	  immediately restore each key/value to its declared generic type.
**/
@:go.import("hxrt")
@:go.package("hxrt")
extern class NativeObjectMap {
	@:go.name("ObjectMapNew")
	public static function create():ObjectMapHandle;

	@:go.name("ObjectMapSet")
	public static function set(handle:ObjectMapHandle, key:Dynamic, value:Dynamic):Void;

	@:go.name("ObjectMapGet")
	public static function get(handle:ObjectMapHandle, key:Dynamic):Dynamic;

	@:go.name("ObjectMapExists")
	public static function exists(handle:ObjectMapHandle, key:Dynamic):Bool;

	@:go.name("ObjectMapRemove")
	public static function remove(handle:ObjectMapHandle, key:Dynamic):Bool;

	@:go.name("ObjectMapKeys")
	public static function keys(handle:ObjectMapHandle):NativeSlice<Dynamic>;

	@:go.name("ObjectMapClear")
	public static function clear(handle:ObjectMapHandle):Void;
}
