package sys.thread;

import hxrt.thread.EventLoopHandle;
import hxrt.thread.EventLoopProgress;
import hxrt.thread.NativeThread;

/**
	When an event loop has an available event to execute.
**/
enum NextEventTime {
	Now;
	Never;
	AnyTime(time:Null<Float>);
	At(time:Float);
}

/**
	Event loop implementation used by `sys.thread.Thread` and direct portable code.

	Why
	The upstream threaded stdlib exposes an event-loop API in ordinary Haxe code.
	Go needs runtime-backed scheduling so timers and queued callbacks can block and
	reawaken without busy loops.

	What
	This wrapper supports one-shot events, promised events, repeating events, and
	blocking loop progression.

	How
	A typed `hxrt.thread.EventLoopHandle` owns the mutable runtime state. The Haxe
	API stays source-owned and maps the runtime progress markers back to the
	upstream `NextEventTime` enum.
**/
@:coreApi
@:allow(sys.thread.Thread)
class EventLoop {
	var __h:EventLoopHandle;

	public function new():Void {
		__h = NativeThread.eventLoopNew();
	}

	static function __fromHandle(handle:EventLoopHandle):EventLoop {
		var loop = new EventLoop();
		untyped loop.__h = handle;
		return loop;
	}

	public function repeat(event:() -> Void, intervalMs:Int):EventHandler {
		return cast NativeThread.eventLoopRepeat(__h, event, intervalMs);
	}

	public function cancel(eventHandler:EventHandler):Void {
		NativeThread.eventLoopCancel(__h, cast eventHandler);
	}

	public function promise():Void {
		NativeThread.eventLoopPromise(__h);
	}

	public function run(event:() -> Void):Void {
		NativeThread.eventLoopRun(__h, event);
	}

	public function runPromised(event:() -> Void):Void {
		NativeThread.eventLoopRunPromised(__h, event);
	}

	public function progress():NextEventTime {
		var result:EventLoopProgress = NativeThread.eventLoopProgress(__h);
		return switch (result.kind) {
			case 0: Now;
			case 1: Never;
			case 2:
				if (result.time < 0) AnyTime(-1.0) else AnyTime(result.time);
			case 3: At(result.time);
			case _: Never;
		};
	}

	public function wait(?timeout:Float):Bool {
		if (timeout == null) {
			return NativeThread.eventLoopWait(__h);
		}
		return NativeThread.eventLoopWaitTimeoutDynamic(__h, timeout);
	}

	public function loop():Void {
		NativeThread.eventLoopLoop(__h);
	}
}

abstract EventHandler(Int) from Int to Int {}
