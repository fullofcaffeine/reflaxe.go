package hxrt.thread;

/**
	What
	- Opaque typed runtime handle for `sys.thread.Mutex` on `haxe.go`.

	Why
	- The public mutex API stays in staged std, while the actual Go synchronization
	  primitive belongs in `hxrt`. This binding is target support, not an upstream
	  stdlib override.

	How
	- Metadata maps the extern to `hxrt.MutexHandle`, which is created and operated
	  on only through typed `NativeThread` functions.
**/
@:go.import("hxrt")
@:go.package("hxrt")
@:go.name("MutexHandle")
extern class MutexHandle {}
