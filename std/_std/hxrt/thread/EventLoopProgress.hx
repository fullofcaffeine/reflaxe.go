package hxrt.thread;

/**
	Typed bridge result for `sys.thread.EventLoop.progress()`.
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
