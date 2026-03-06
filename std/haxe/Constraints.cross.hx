package haxe;

/**
	What
	- Go-target override for `haxe.Constraints`.
	- Defines the portable constraint aliases used by the stdlib.

	Why
	- `haxe.go` cannot rely on upstream std source being emitted automatically for
	  these constraint aliases.
	- `IMap<K, V>` needs one Go-specific bridge on `copy()`: Haxe allows concrete
	  map implementations to return their own map type, but Go interfaces require
	  an exact method signature match.

	How
	- Keep the compile-time constraint aliases local to this staged std module.
	- Bind `IMap.copy()` to the native selector `copyIMap` so concrete maps can
	  expose an interface-typed bridge for Go while staged std wrappers keep the
	  user-facing `copy()` return type precise.
**/
@:callable
abstract Function(Dynamic) {}

abstract FlatEnum(Dynamic) {}
abstract NotVoid(Dynamic) {}
abstract Constructible<T>(Dynamic) {}

interface IMap<K, V> {
	function get(k:K):Null<V>;
	function set(k:K, v:V):Void;
	function exists(k:K):Bool;
	function remove(k:K):Bool;
	function keys():Iterator<K>;
	function iterator():Iterator<V>;
	function keyValueIterator():KeyValueIterator<K, V>;
	@:native("copyIMap") function copy():IMap<K, V>;
	function toString():String;
	function clear():Void;
}
