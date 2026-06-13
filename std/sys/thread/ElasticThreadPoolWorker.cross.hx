package sys.thread;

/**
	Internal worker used by `ElasticThreadPool`.
**/
class ElasticThreadPoolWorker {
	public var task(default, null):Null<() -> Void>;
	public var dead(default, null) = false;

	final deathMutex = new Mutex();
	final waiter = new Lock();
	final queue:Deque<() -> Void>;
	final timeout:Float;
	var isShutdown = false;

	public function new(queue:Deque<() -> Void>, timeout:Float) {
		this.queue = queue;
		this.timeout = timeout;
		start();
	}

	public function wakeup(task:() -> Void):Void {
		deathMutex.acquire();
		if (dead) {
			start();
		}
		this.task = task;
		waiter.release();
		deathMutex.release();
	}

	public function shutdown():Void {
		isShutdown = true;
		waiter.release();
	}

	function start():Void {
		dead = false;
		Thread.create(loop);
	}

	function loop():Void {
		try {
			while (waiter.wait(timeout)) {
				switch task {
					case null:
						if (isShutdown) {
							break;
						}
					case fn:
						fn();
						while (true) {
							switch queue.pop(false) {
								case null:
									break;
								case queued:
									queued();
							}
						}
						task = null;
				}
			}
			deathMutex.acquire();
			if (task != null) {
				start();
			} else {
				dead = true;
			}
			deathMutex.release();
		} catch (err:Dynamic) {
			// Dynamic catch is intentional: worker tasks may throw any Haxe value, and
			// the pool must restart the worker before rethrowing that value unchanged.
			task = null;
			start();
			throw err;
		}
	}
}
