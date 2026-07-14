package hxrt.thread;

/**
	What
	- Opaque typed runtime handle for `sys.thread.Semaphore` on `haxe.go`.

	Why
	- Token accounting and blocking require the Go runtime implementation in `hxrt`.
	  This carrier is runtime support rather than an upstream stdlib override and
	  therefore stays outside `_std`.

	How
	- Metadata maps the Haxe extern to `hxrt.SemaphoreHandle`; staged std accesses
	  it through typed `NativeThread` semaphore operations.
**/
@:go.import("hxrt")
@:go.package("hxrt")
@:go.name("SemaphoreHandle")
extern class SemaphoreHandle {}
