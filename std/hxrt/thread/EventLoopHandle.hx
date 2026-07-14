package hxrt.thread;

/**
	What
	- Opaque typed runtime handle for `sys.thread.EventLoop` state on `haxe.go`.

	Why
	- Queue, timer, and wait state is implemented by real Go runtime primitives in
	  `hxrt`. The carrier therefore belongs in runtime support rather than the
	  override-only stdlib tree.

	How
	- Metadata maps this extern to `hxrt.EventLoopHandle`; staged `sys.thread` code
	  passes it only through typed `NativeThread` operations.
**/
@:go.import("hxrt")
@:go.package("hxrt")
@:go.name("EventLoopHandle")
extern class EventLoopHandle {}
