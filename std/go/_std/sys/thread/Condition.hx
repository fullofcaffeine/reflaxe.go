package sys.thread;

import hxrt.thread.ConditionHandle;
import hxrt.thread.NativeThread;

/**
	Condition variable with an internal mutex.

	Why
	The mainstream Haxe stdlib implementation cannot be used unchanged on `haxe.go`.
	The upstream threaded stdlib expects `Condition` to exist as a first-class
	synchronization primitive.

	What
	Typed staged wrapper over the Go runtime condition handle.

	How
	The runtime owns the internal mutex and signal state. The Haxe-facing API
	stays identical to upstream.
**/
@:coreApi
class Condition {
	final __h:ConditionHandle;

	public function new():Void {
		__h = NativeThread.conditionNew();
	}

	public function acquire():Void {
		NativeThread.conditionAcquire(__h);
	}

	public function tryAcquire():Bool {
		return NativeThread.conditionTryAcquire(__h);
	}

	public function release():Void {
		NativeThread.conditionRelease(__h);
	}

	public function wait():Void {
		NativeThread.conditionWait(__h);
	}

	public function signal():Void {
		NativeThread.conditionSignal(__h);
	}

	public function broadcast():Void {
		NativeThread.conditionBroadcast(__h);
	}
}
