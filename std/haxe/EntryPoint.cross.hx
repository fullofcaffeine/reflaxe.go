package haxe;

import sys.thread.Mutex;
import sys.thread.Thread;

/**
	Entry point bridge for direct `haxe.EntryPoint` use on `haxe.go`.

	What
	Provides the public `haxe.EntryPoint` API used by `haxe.MainLoop` and
	`haxe.Timer`: scheduling work on the main thread, tracking worker lifetimes,
	and running the main event loop.

	Why
	The upstream implementation delegates threaded targets to the target runtime.
	On Go, that runtime is the `sys.thread.EventLoop` wrapper backed by
	`runtime/hxrt/thread.go`, so direct `haxe.EntryPoint` has to connect to it
	explicitly instead of relying on a compiler-inserted loop.

	How
	The main logical thread is captured once. `runInMainThread` queues callbacks on
	that thread's event loop, `addThread` promises a future main-loop wakeup while
	the worker is alive, and `run` drains the main event loop until the runtime has
	no blocking work left.
**/
class EntryPoint {
	static var mutex = new Mutex();
	static var mainThread:Thread = Thread.current();

	public static var threadCount(default, null):Int = 0;

	public static function wakeup():Void {
		mainThread.events.run(function() {});
	}

	public static function runInMainThread(f:Void->Void):Void {
		if (f != null) {
			mainThread.events.run(f);
		}
	}

	public static function addThread(f:Void->Void):Void {
		mutex.acquire();
		threadCount++;
		mutex.release();
		mainThread.events.promise();

		Thread.create(function() {
			if (f != null) {
				f();
			}
			mutex.acquire();
			threadCount--;
			mutex.release();
			mainThread.events.runPromised(function() {});
		});
	}

	@:keep public static function run():Void {
		mainThread.events.loop();
	}
}
