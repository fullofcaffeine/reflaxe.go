package haxe;

/**
	What
	- Go-target override for `haxe.Constraints`.
	- Defines the portable constraint aliases used by the stdlib.

	Why
	The mainstream Haxe stdlib implementation cannot be used unchanged on `haxe.go`.
	- `haxe.go` cannot rely on upstream std source being emitted automatically for
	  these constraint aliases.
	- `IMap<K, V>` needs Go-specific erased bridges for key operations and `copy()`:
	  Haxe permits concrete maps to specialize key parameters and return their own
	  type, while Go interfaces require exact parameter and result signatures.

	How
	- Keep the compile-time constraint aliases local to this staged std module.
	- Bind the affected methods to `*IMap` selectors. Concrete staged maps expose
	  narrow, kept bridge methods while their public Haxe signatures remain precise.
**/
@:callable
abstract Function(Dynamic) {}

abstract FlatEnum(Dynamic) {}
abstract NotVoid(Dynamic) {}
abstract Constructible<T>(Dynamic) {}

interface IMap<K, V> {
	@:go.name("getIMap")
	function get(k:K):Null<V>;
	@:go.name("setIMap")
	function set(k:K, v:V):Void;
	@:go.name("existsIMap")
	function exists(k:K):Bool;
	@:go.name("removeIMap")
	function remove(k:K):Bool;
	function keys():Iterator<K>;
	function iterator():Iterator<V>;
	function keyValueIterator():KeyValueIterator<K, V>;
	@:go.name("copyIMap") function copy():IMap<K, V>;
	function toString():String;
	function clear():Void;
}
