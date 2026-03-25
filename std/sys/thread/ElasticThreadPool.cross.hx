package sys.thread;

/**
	Thread pool with a varying amount of threads.

	Why
	The upstream threaded stdlib keeps elastic pool behavior in library code. The
	Go target should do the same once `Thread` and the primitive queue/lock
	surfaces are real.

	What
	Workers wake on demand, execute queued tasks, and time out when idle.

	How
	This is a staged std implementation on top of `Thread`, `Lock`, `Mutex`, and
	`Deque`, so the pool policy stays readable Haxe code instead of backend-owned
	Go emitters.
**/
@:coreApi
class ElasticThreadPool implements IThreadPool {
	public var threadsCount(get, null):Int;
	public var maxThreadsCount:Int;
	public var isShutdown(get, never):Bool;

	var _isShutdown = false;

	function get_isShutdown():Bool {
		return _isShutdown;
	}

	final pool:Array<ElasticThreadPoolWorker> = [];
	final queue = new Deque<() -> Void>();
	final mutex = new Mutex();
	final threadTimeout:Float;

	public function new(maxThreadsCount:Int, threadTimeout:Float = 60):Void {
		if (maxThreadsCount < 1) {
			throw new ThreadPoolException("ElasticThreadPool needs maxThreadsCount to be at least 1.");
		}
		this.maxThreadsCount = maxThreadsCount;
		this.threadTimeout = threadTimeout;
	}

	public function run(task:() -> Void):Void {
		if (_isShutdown) {
			throw new ThreadPoolException("Task is rejected. Thread pool is shut down.");
		}
		if (task == null) {
			throw new ThreadPoolException("Task to run must not be null.");
		}

		mutex.acquire();
		var submitted = false;
		var deadWorker:Null<ElasticThreadPoolWorker> = null;
		for (worker in pool) {
			if (deadWorker == null && worker.dead) {
				deadWorker = worker;
			}
			if (worker.task == null) {
				submitted = true;
				worker.wakeup(task);
				break;
			}
		}
		if (!submitted) {
			if (deadWorker != null) {
				deadWorker.wakeup(task);
			} else if (pool.length < maxThreadsCount) {
				var worker = new ElasticThreadPoolWorker(queue, threadTimeout);
				pool.push(worker);
				worker.wakeup(task);
			} else {
				queue.add(task);
			}
		}
		mutex.release();
	}

	public function shutdown():Void {
		if (_isShutdown) {
			return;
		}
		mutex.acquire();
		_isShutdown = true;
		for (worker in pool) {
			worker.shutdown();
		}
		mutex.release();
	}

	function get_threadsCount():Int {
		var result = 0;
		for (worker in pool) {
			if (!worker.dead) {
				result++;
			}
		}
		return result;
	}
}
