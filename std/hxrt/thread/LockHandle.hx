package hxrt.thread;

/**
	What
	Opaque runtime handle for `sys.thread.Lock` on `haxe.go`.

	Why
	The public `Lock` API belongs in staged std, but the actual blocking primitive
	must live in the Go runtime so it can block without busy loops.

	How
	This extern names the exported Go `hxrt.LockHandle` type so the staged wrapper
	can stay typed and avoid `Dynamic` at the Haxe boundary.
**/
@:go.import("hxrt")
@:go.package("hxrt")
@:go.name("LockHandle")
extern class LockHandle {}
