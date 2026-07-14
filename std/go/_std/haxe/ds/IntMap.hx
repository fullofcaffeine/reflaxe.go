package haxe.ds;

/**
	What:
	- Declares the Go-target `IntMap` API and key/value iterator bridge.

	Why:
	- The mainstream Haxe stdlib implementation cannot be used unchanged on `haxe.go`. because the compiler-owned native Go map carrier needs a target-specific extern shape and typed `copyIMap` bridge.

	How:
	- Expose the upstream map contract while routing iteration and copies through the generated carrier methods.
**/
extern class IntMap<T> implements haxe.Constraints.IMap<Int, T> {
	function new():Void;
	function set(key:Int, value:T):Void;
	function get(key:Int):Null<T>;
	function exists(key:Int):Bool;
	function remove(key:Int):Bool;
	function keys():Iterator<Int>;
	function iterator():Iterator<T>;

	@:runtime inline function keyValueIterator():KeyValueIterator<Int, T> {
		return new haxe.iterators.MapKeyValueIterator(this);
	}

	@:native("copyIMap") private function copyIMap():haxe.Constraints.IMap<Int, T>;

	@:runtime inline function copy():IntMap<T> {
		return cast copyIMap();
	}

	function toString():String;
	function clear():Void;
}
