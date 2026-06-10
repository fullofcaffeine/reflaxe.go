package haxe;

import hxrt.thread.NativeThread;
import sys.thread.EventLoop.EventHandler;
import sys.thread.Thread;

/**
	Timer implementation for direct `haxe.Timer` use on `haxe.go`.

	What
	Implements repeating timers, one-shot `delay`, `measure`, and `stamp` for the
	Go target.

	Why
	The upstream threaded implementation expects the target to provide a real event
	loop for the current thread. Go now has that through `sys.thread.EventLoop`, so
	this override keeps timer behavior in staged std code instead of compiler raw
	shims.

	How
	Each timer registers a repeating event on `Thread.current().events`. `delay`
	creates one of those timers and stops it from inside the first callback.
**/
class Timer {
	var thread:Thread;
	var eventHandler:EventHandler;

	public var run:Void->Void;

	public function new(time_ms:Int) {
		run = function() {};
		thread = Thread.current();
		eventHandler = thread.events.repeat(function() this.run(), time_ms);
	}

	public function stop():Void {
		if (eventHandler != cast 0) {
			thread.events.cancel(eventHandler);
			eventHandler = cast 0;
		}
	}

	public static function delay(f:Void->Void, time_ms:Int):Timer {
		var timer = new Timer(time_ms);
		timer.run = function() {
			timer.stop();
			f();
		};
		return timer;
	}

	public static function measure<T>(f:Void->T, ?pos:PosInfos):T {
		var t0 = stamp();
		var result = f();
		Log.trace((stamp() - t0) + "s", pos);
		return result;
	}

	public static inline function stamp():Float {
		return NativeThread.nowSeconds();
	}
}
