package sys.thread;

import hxrt.thread.NativeThread;
import hxrt.thread.SemaphoreHandle;

/**
	Counting semaphore for `haxe.go`.

	Why
	The upstream threaded stdlib uses `Semaphore` directly and as a building block
	for higher-level synchronization.

	What
	Typed staged wrapper over the Go runtime semaphore handle.

	How
	The runtime owns the blocking counter semantics. The wrapper selects the timed
	or untimed fast path based on whether `timeout` was provided.
**/
@:coreApi
class Semaphore {
	final __h:SemaphoreHandle;

	public function new(value:Int):Void {
		__h = NativeThread.semaphoreNew(value);
	}

	public function acquire():Void {
		NativeThread.semaphoreAcquire(__h);
	}

	public function tryAcquire(timeout:Null<Float> = null):Bool {
		if (timeout == null) {
			return NativeThread.semaphoreTryAcquire(__h);
		}
		return NativeThread.semaphoreTryAcquireTimeoutDynamic(__h, timeout);
	}

	public function release():Void {
		NativeThread.semaphoreRelease(__h);
	}
}
