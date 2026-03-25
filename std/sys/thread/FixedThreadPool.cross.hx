package sys.thread;

import haxe.Exception;

/**
	Thread pool with a constant amount of threads.

	Why
	This is ordinary upstream library behavior once `Thread` and `Deque` are real.

	What
	A fixed number of workers block on the shared queue until shutdown.

	How
	Shutdown is implemented by enqueueing a private task that throws a private
	exception caught inside the worker loop.
**/
@:coreApi
class FixedThreadPool implements IThreadPool {
	public var threadsCount(get, null):Int;
	public var isShutdown(get, never):Bool;

	var _isShutdown = false;
	final pool:Array<FixedThreadPoolWorker>;
	final queue = new Deque<() -> Void>();

	function get_threadsCount():Int {
		return pool.length;
	}

	function get_isShutdown():Bool {
		return _isShutdown;
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
		if (_isShutdown) {
			throw new ThreadPoolException("Task is rejected. Thread pool is shut down.");
		}
		if (task == null) {
			throw new ThreadPoolException("Task to run must not be null.");
		}
		queue.add(task);
	}

	public function shutdown():Void {
		if (_isShutdown) {
			return;
		}
		_isShutdown = true;
		for (_ in pool) {
			queue.add(shutdownTask);
		}
	}

	static function shutdownTask():Void {
		throw new FixedThreadPoolShutdownException("");
	}
}
