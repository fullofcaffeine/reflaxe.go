package sys.thread;

/**
	What
	- Repo-authored worker support for the staged `sys.thread.ElasticThreadPool`
	  implementation.

	Why
	- The worker is a target support type rather than an upstream Haxe stdlib module,
	  so packaging it as an override-only `.cross.hx` file gave it the wrong
	  ownership. Its restart and timeout behavior still belongs in ordinary Haxe
	  beside the public pool override.

	How
	- A lock wakes one assigned task, the shared deque drains queued work, and the
	  worker either marks itself dead after the timeout or restarts before propagating
	  a task failure.
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
