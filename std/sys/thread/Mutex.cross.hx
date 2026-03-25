package sys.thread;

import hxrt.thread.MutexHandle;
import hxrt.thread.NativeThread;

/**
	Re-entrant mutex for `haxe.go`.

	Why
	The upstream `sys.thread.Mutex` contract allows the owning thread to acquire
	the same mutex multiple times.

	What
	Typed staged wrapper over the Go runtime mutex handle.

	How
	The runtime tracks owner identity and recursion depth; this wrapper just keeps
	the public API source-owned.
**/
@:coreApi
class Mutex {
	final __h:MutexHandle;

	public function new():Void {
		__h = NativeThread.mutexNew();
	}

	public function acquire():Void {
		NativeThread.mutexAcquire(__h);
	}

	public function tryAcquire():Bool {
		return NativeThread.mutexTryAcquire(__h);
	}

	public function release():Void {
		NativeThread.mutexRelease(__h);
	}
}
