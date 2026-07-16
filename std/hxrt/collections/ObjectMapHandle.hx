package hxrt.collections;

/**
	What
	- Typed opaque binding for one native identity-keyed object map.

	Why
	- Object identity and Go hash storage are runtime representation facts. Exposing
	  the carrier as `Dynamic` would allow callers to bypass `ObjectMap` semantics.

	How
	- Map directly to `hxrt.ObjectMapCell`; only `NativeObjectMap` operations consume
	  the handle while staged `haxe.ds.ObjectMap` owns the public algorithms.
**/
@:go.import("hxrt")
@:go.package("hxrt")
@:go.name("ObjectMapCell")
extern class ObjectMapHandle {}
