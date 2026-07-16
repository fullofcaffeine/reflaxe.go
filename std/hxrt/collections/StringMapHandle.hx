package hxrt.collections;

/**
	What
	- Typed opaque binding for one native string-keyed map.

	Why
	- Go's hash-map storage cannot be represented as ordinary portable Haxe data,
	  and exposing the storage as `Dynamic` would erase the collection boundary.

	How
	- Map directly to `hxrt.StringMapCell`; only `NativeStringMap` operations
	  consume the handle while staged `haxe.ds.StringMap` owns the public algorithms.
**/
@:go.import("hxrt")
@:go.package("hxrt")
@:go.name("StringMapCell")
extern class StringMapHandle {}
