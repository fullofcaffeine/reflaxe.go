package sys.thread;

import hxrt.thread.NativeThread;

/**
	Logical thread API for `haxe.go`.

	Why
	The mainstream Haxe stdlib implementation cannot be used unchanged on `haxe.go`.
	The portable `sys.thread.Thread` surface requires message queues and optional
	event-loop ownership per thread.

	What
	`create` spawns a worker without an event loop, `createWithEventLoop` spawns a
	worker that drains its loop after `job`, and `sendMessage` / `readMessage`
	implement the upstream dynamic message queue contract. Portable workers are
	foreground threads: generated `main` waits for them and for nested workers.
	An uncaught Haxe throw ends only that worker and is reported as an uncaught
	exception; a foreign Go panic remains a fatal native panic.

	How
	Thread identity and queues live in `hxrt.thread.NativeThread`. Event-loop
	availability is checked explicitly so `Thread.events` still throws
	`NoEventLoopException` for threads that were not created with loop support.
	The compiler adds the foreground drain only when the thread runtime feature is
	selected. In that same feature-gated case, explicit `go.Go.spawn` callbacks get
	only a detached identity-cleanup scope: they remain non-joined and native panics
	remain fatal.
**/
class Thread {
	final __id:Int;

	function new(id:Int) {
		__id = id;
	}

	public var events(get, never):EventLoop;

	function get_events():EventLoop {
		if (!NativeThread.hasEventLoop(__id)) {
			throw new NoEventLoopException();
		}
		return EventLoop.__fromHandle(NativeThread.events(__id));
	}

	public function sendMessage(msg:Dynamic):Void {
		NativeThread.sendMessage(__id, msg);
	}

	public static function current():Thread {
		return new Thread(NativeThread.currentId());
	}

	public static function create(job:() -> Void):Thread {
		return new Thread(NativeThread.spawn(job));
	}

	public static function runWithEventLoop(job:() -> Void):Void {
		NativeThread.runWithEventLoop(job);
	}

	public static function createWithEventLoop(job:() -> Void):Thread {
		return new Thread(NativeThread.spawnWithEventLoop(job));
	}

	public static function readMessage(block:Bool):Dynamic {
		return NativeThread.readMessage(block);
	}

	private static function processEvents():Void {
		var current = Thread.current();
		if (NativeThread.hasEventLoop(current.__id)) {
			current.events.progress();
		}
	}
}
