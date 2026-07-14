package haxe.ds;

/**
	What:
	- Declares the Go-target `ObjectMap` API and key/value iterator bridge.

	Why:
	- The mainstream Haxe stdlib implementation cannot be used unchanged on `haxe.go`. because the compiler-owned native Go map carrier needs a target-specific extern shape and typed `copyIMap` bridge.

	How:
	- Expose the upstream map contract while routing iteration and copies through the generated carrier methods.
**/
extern class ObjectMap<K:{}, V> implements haxe.Constraints.IMap<K, V> {
	function new():Void;
	function set(key:K, value:V):Void;
	function get(key:K):Null<V>;
	function exists(key:K):Bool;
	function remove(key:K):Bool;
	function keys():Iterator<K>;
	function iterator():Iterator<V>;

	@:runtime inline function keyValueIterator():KeyValueIterator<K, V> {
		return new haxe.iterators.MapKeyValueIterator(this);
	}

	@:native("copyIMap") private function copyIMap():haxe.Constraints.IMap<K, V>;

	@:runtime inline function copy():ObjectMap<K, V> {
		return cast copyIMap();
	}

	function toString():String;
	function clear():Void;
}
