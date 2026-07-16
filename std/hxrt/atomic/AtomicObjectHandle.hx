package hxrt.atomic;

/**
	What
	- Typed opaque binding for one native atomic object cell.

	Why
	- The mutex and object reference stored by Go cannot be represented as portable
	  Haxe data, and exposing the cell itself as `Dynamic` would allow callers to
	  bypass the atomic API.

	How
	- Map directly to `hxrt.AtomicObjectCell`; only `NativeAtomicObject` operations
	  consume the handle.
**/
@:go.import("hxrt")
@:go.package("hxrt")
@:go.name("AtomicObjectCell")
extern class AtomicObjectHandle {}
