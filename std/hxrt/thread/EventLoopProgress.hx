package hxrt.thread;

/**
	What
	- Typed runtime result for `sys.thread.EventLoop.progress()`.

	Why
	- Progress state originates in the Go-backed event loop and must cross the
	  runtime boundary without an anonymous or dynamic carrier. It is runtime
	  support rather than an upstream stdlib override.

	How
	- The extern maps directly to `hxrt.EventLoopProgress` and exposes its progress
	  kind and next scheduled time as typed fields.
**/
@:go.import("hxrt")
@:go.package("hxrt")
@:go.name("EventLoopProgress")
extern class EventLoopProgress {
	@:go.name("Kind")
	public var kind:Int;

	@:go.name("Time")
	public var time:Float;
}
