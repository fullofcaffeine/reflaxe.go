package hxrt.atomic;

/**
	What
	- Typed opaque binding for one native atomic integer cell.

	Why
	- Go's `sync/atomic.Int64` storage cannot be represented as portable Haxe data,
	  and exposing the cell as `Dynamic` would erase the ownership boundary.

	How
	- Map directly to `hxrt.AtomicIntCell`; only `NativeAtomicInt` operations consume
	  the handle.
**/
@:go.import("hxrt")
@:go.package("hxrt")
@:go.name("AtomicIntCell")
extern class AtomicIntHandle {}
