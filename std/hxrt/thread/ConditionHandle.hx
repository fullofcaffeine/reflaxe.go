package hxrt.thread;

/**
	What
	- Opaque typed runtime handle for `sys.thread.Condition` on `haxe.go`.

	Why
	- Condition state and blocking operations require real Go synchronization
	  primitives in `hxrt`. This carrier is a runtime binding rather than an upstream
	  stdlib override, so it lives under `std/hxrt`.

	How
	- Metadata maps the empty Haxe extern to the exported Go
	  `hxrt.ConditionHandle` type used by `NativeThread`.
**/
@:go.import("hxrt")
@:go.package("hxrt")
@:go.name("ConditionHandle")
extern class ConditionHandle {}
