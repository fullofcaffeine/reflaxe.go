package sys.thread;

import haxe.Exception;

/**
	Thread pool with a constant amount of threads.

	Why
	The mainstream Haxe stdlib implementation cannot be used unchanged on `haxe.go`.
	This is ordinary upstream library behavior once `Thread` and `Deque` are real.

	What
	A fixed number of workers block on the shared queue until shutdown. A task
	is accepted only when its queue insertion completes before shutdown's single
	serialized state transition.

	How
	A pool mutex linearizes `run`, `isShutdown`, and `shutdown`. Shutdown holds the
	mutex while enqueueing the private worker sentinels, so every accepted task is
	ahead of every sentinel and therefore executes exactly once before workers exit.
**/
@:coreApi
class FixedThreadPool implements IThreadPool {
	public var threadsCount(get, null):Int;
	public var isShutdown(get, never):Bool;

	var _isShutdown = false;
	final pool:Array<FixedThreadPoolWorker>;
	final queue = new Deque<() -> Void>();
	final mutex = new Mutex();

	function get_threadsCount():Int {
		return pool.length;
	}

	function get_isShutdown():Bool {
		mutex.acquire();
		var result = _isShutdown;
		mutex.release();
		return result;
	}

	public function new(threadsCount:Int):Void {
		if (threadsCount < 1) {
			throw new ThreadPoolException("FixedThreadPool needs threadsCount to be at least 1.");
		}
		var workers = new Array<FixedThreadPoolWorker>();
		for (_ in 0...threadsCount) {
			workers.push(new FixedThreadPoolWorker(queue));
		}
		pool = workers;
	}

	public function run(task:() -> Void):Void {
		mutex.acquire();
		if (_isShutdown) {
			mutex.release();
			throw new ThreadPoolException("Task is rejected. Thread pool is shut down.");
		}
		if (task == null) {
			mutex.release();
			throw new ThreadPoolException("Task to run must not be null.");
		}
		queue.add(task);
		mutex.release();
	}

	public function shutdown():Void {
		mutex.acquire();
		if (_isShutdown) {
			mutex.release();
			return;
		}
		_isShutdown = true;
		for (_ in pool) {
			queue.add(shutdownTask);
		}
		mutex.release();
	}

	static function shutdownTask():Void {
		throw new FixedThreadPoolShutdownException("");
	}
}
