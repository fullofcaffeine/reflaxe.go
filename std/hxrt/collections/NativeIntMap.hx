package hxrt.collections;

import go.NativeSlice;

/**
	What
	- Typed bridge to the native integer-keyed storage used by staged `IntMap`.

	Why
	- The public map API is ordinary Haxe library behavior, but Haxe source cannot
	  directly declare or operate on a Go `map[int]any` value.

	How
	- Expose only storage primitives and a deterministic native-slice key snapshot.
	  Generic map values use `Dynamic` solely at this erased runtime boundary and
	  are cast back to `T` immediately by `haxe.ds.IntMap`.
**/
@:go.import("hxrt")
@:go.package("hxrt")
extern class NativeIntMap {
	@:go.name("IntMapNew")
	public static function create():IntMapHandle;

	@:go.name("IntMapSet")
	public static function set(handle:IntMapHandle, key:Int, value:Dynamic):Void;

	@:go.name("IntMapGet")
	public static function get(handle:IntMapHandle, key:Int):Dynamic;

	@:go.name("IntMapExists")
	public static function exists(handle:IntMapHandle, key:Int):Bool;

	@:go.name("IntMapRemove")
	public static function remove(handle:IntMapHandle, key:Int):Bool;

	@:go.name("IntMapKeys")
	public static function keys(handle:IntMapHandle):NativeSlice<Int>;

	@:go.name("IntMapClear")
	public static function clear(handle:IntMapHandle):Void;
}
