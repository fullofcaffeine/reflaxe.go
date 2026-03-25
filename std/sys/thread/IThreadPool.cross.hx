package sys.thread;

/**
	Shared thread-pool interface.

	Why
	This public abstraction belongs in staged std even while concrete Go pool
	implementations are still pending.

	What
	Exposes thread-count, shutdown state, task submission, and shutdown.

	How
	Concrete implementations can stay source-owned and build on top of the lower
	`sys.thread` primitives as they land.
**/
interface IThreadPool {
	var threadsCount(get, never):Int;
	var isShutdown(get, never):Bool;
	function run(task:() -> Void):Void;
	function shutdown():Void;
}
