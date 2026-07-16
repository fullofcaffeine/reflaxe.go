package hxrt.collections;

/**
	What
	- Typed opaque binding for one native integer-keyed map.

	Why
	- Go's hash-map storage cannot be represented as ordinary portable Haxe data,
	  and exposing the storage as `Dynamic` would erase the collection boundary.

	How
	- Map directly to `hxrt.IntMapCell`; only `NativeIntMap` operations consume the
	  handle while staged `haxe.ds.IntMap` owns the public algorithms.
**/
@:go.import("hxrt")
@:go.package("hxrt")
@:go.name("IntMapCell")
extern class IntMapHandle {}
