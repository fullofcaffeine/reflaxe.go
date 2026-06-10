package sys.thread;

import hxrt.thread.LockHandle;
import hxrt.thread.NativeThread;

/**
	A lock that blocks until released.

	Why
	- `sys.thread.Lock` is part of the upstream threaded stdlib contract.
	- Go needs a real blocking primitive in the runtime to avoid busy loops.

	What
	- Thin staged wrapper around the Go runtime `hxrt.LockHandle`.

	How
	- `new()` creates the runtime handle.
	- `wait()` delegates to either the no-timeout or timed wait runtime path.
	- `release()` wakes exactly one waiter or increments the stored count.
**/
@:coreApi
class Lock {
	final __h:LockHandle;

	public function new():Void {
		__h = NativeThread.lockNew();
	}

	public function wait(timeout:Null<Float> = null):Bool {
		if (timeout == null) {
			return NativeThread.lockWait(__h);
		}
		return NativeThread.lockWaitTimeout(__h, timeout);
	}

	public function release():Void {
		NativeThread.lockRelease(__h);
	}
}
