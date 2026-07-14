package sys.thread;

/**
	Thread pool with a varying amount of threads.

	Why
	The mainstream Haxe stdlib implementation cannot be used unchanged on `haxe.go`.
	Its worker fields are ordinary shared values, which become Go data races when
	`run`, timeout, task completion, and `shutdown` overlap.

	What
	Accepted tasks enter one shared queue. The pool grows while no worker is idle,
	shrinks after `threadTimeout`, drains accepted work during shutdown, and rejects
	every task whose admission loses the serialized shutdown race.

	How
	One pool mutex owns shutdown state, admission, worker lifecycle counters, and
	task completion accounting. A shared counted `Lock` pairs one wake token with
	each queued task; pending work grows the pool when it exceeds live capacity,
	and shutdown adds one exit token per live worker. Callbacks run outside the pool
	mutex, while every lifecycle transition returns through it.
**/
@:allow(sys.thread.ElasticThreadPoolWorker)
@:coreApi
class ElasticThreadPool implements IThreadPool {
	public var threadsCount(get, null):Int;

	/**
		Maximum live workers. This stays a plain writable field to match Haxe's core
		API; callers must externally synchronize concurrent mutation with pool use.
	**/
	public var maxThreadsCount:Int;

	public var isShutdown(get, never):Bool;

	var _isShutdown = false;
	var liveWorkers = 0;
	var pendingTasks = 0;

	final pool:Array<ElasticThreadPoolWorker> = [];
	final queue = new Deque<() -> Void>();
	final available = new Lock();
	final mutex = new Mutex();
	final threadTimeout:Float;

	/**
		What: create an empty elastic pool with a bounded worker count.
		Why: at least one possible worker is required to uphold accepted-task delivery.
		How: workers start lazily when `run` first observes no idle capacity.
	**/
	public function new(maxThreadsCount:Int, threadTimeout:Float = 60):Void {
		if (maxThreadsCount < 1) {
			throw new ThreadPoolException("ElasticThreadPool needs maxThreadsCount to be at least 1.");
		}
		this.maxThreadsCount = maxThreadsCount;
		this.threadTimeout = threadTimeout;
	}

	function get_isShutdown():Bool {
		mutex.acquire();
		var result = _isShutdown;
		mutex.release();
		return result;
	}

	function get_threadsCount():Int {
		mutex.acquire();
		var result = liveWorkers;
		mutex.release();
		return result;
	}

	/**
		What: submit one task or reject it after shutdown.
		Why: returning normally must mean the task is durably owned exactly once.
		How: queue insertion and the shutdown check share the pool mutex; a wake token
		is published before another worker starts when pending work exceeds live capacity.
	**/
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

		pendingTasks++;
		queue.add(task);
		available.release();
		if (pendingTasks > liveWorkers && liveWorkers < maxThreadsCount) {
			startWorkerLocked();
		}
		mutex.release();
	}

	/**
		What: close admission and wake every live worker for draining or exit.
		Why: shutdown must be idempotent and must not strand previously accepted work.
		How: the same mutex that admits tasks publishes shutdown, then contributes one
		extra wake token per live worker after all accepted task tokens.
	**/
	public function shutdown():Void {
		mutex.acquire();
		if (_isShutdown) {
			mutex.release();
			return;
		}
		_isShutdown = true;
		for (_ in 0...liveWorkers) {
			available.release();
		}
		mutex.release();
	}

	/**
		What: start a new or previously retired worker.
		Why: retired worker objects can be reused without exposing mutable state to
		concurrent admission code.
		How: the caller holds `mutex`, updates the live count, and publishes the new
		thread before releasing pool state.
	**/
	function startWorkerLocked():Void {
		var selected:Null<ElasticThreadPoolWorker> = null;
		for (worker in pool) {
			if (worker.dead) {
				selected = worker;
				break;
			}
		}
		if (selected == null) {
			selected = new ElasticThreadPoolWorker(this, available, threadTimeout);
			pool.push(selected);
		}
		liveWorkers++;
		selected.start();
	}

	/**
		What: resolve one worker wakeup to a task, retry, or retirement.
		Why: a task can arrive at the timeout boundary after `wait` reports false.
		How: under `mutex`, a zero-time token check closes that boundary before the
		worker is marked dead; task and wake tokens are consumed as a pair.
	**/
	function workerResolveWait(worker:ElasticThreadPoolWorker, woke:Bool):Bool {
		mutex.acquire();
		var hasToken = woke;
		if (!hasToken) {
			hasToken = available.wait(0);
		}
		if (hasToken) {
			var nextTask = queue.pop(false);
			if (nextTask != null) {
				worker.task = nextTask;
				mutex.release();
				return true;
			}
			if (!_isShutdown) {
				mutex.release();
				return true;
			}
		}
		retireWorkerLocked(worker);
		mutex.release();
		return false;
	}

	/**
		What: retire a worker that timed out or consumed its shutdown exit token.
		Why: live-worker accounting drives future growth and shutdown wake counts.
		How: only pool-mutex code mutates `dead` and `liveWorkers`.
	**/
	function retireWorkerLocked(worker:ElasticThreadPoolWorker):Void {
		if (!worker.dead) {
			worker.dead = true;
			worker.task = null;
			liveWorkers--;
		}
	}

	/**
		What: account for one callback that returned normally.
		Why: pending work includes queued and executing callbacks until completion.
		How: clear the worker task and decrement the count under the pool mutex.
	**/
	function workerTaskFinished(worker:ElasticThreadPoolWorker):Void {
		mutex.acquire();
		worker.task = null;
		pendingTasks--;
		mutex.release();
	}

	/**
		What: account for a callback that threw and replace its worker when needed.
		Why: one failing task must not strand other tasks accepted before shutdown.
		How: retire under the pool mutex, then start replacement capacity when pending
		work exceeds the remaining live worker count.
	**/
	function workerTaskFailed(worker:ElasticThreadPoolWorker):Void {
		mutex.acquire();
		worker.task = null;
		pendingTasks--;
		retireWorkerLocked(worker);
		if (pendingTasks > liveWorkers && liveWorkers < maxThreadsCount) {
			startWorkerLocked();
		}
		mutex.release();
	}
}
