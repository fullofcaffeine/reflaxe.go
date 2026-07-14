package sys.thread;

/**
	What
	- Repo-authored worker support for the staged `sys.thread.ElasticThreadPool`.

	Why
	- The mainstream worker's independently synchronized `task`, `dead`, timeout,
	  and shutdown fields become data races on generated Go and permit queued work
	  to be stranded between queue draining and idle publication.

	How
	- Workers consume task tokens from one pool-owned counted lock, ask the pool to
	  resolve timeout boundaries, execute callbacks outside pool synchronization,
	  and report every completion or failure back to the pool state machine.
**/
@:allow(sys.thread.ElasticThreadPool)
class ElasticThreadPoolWorker {
	var task:Null<() -> Void> = null;
	var dead = true;

	final owner:ElasticThreadPool;
	final available:Lock;
	final timeout:Float;

	/**
		What: create a reusable worker object without starting a thread.
		Why: the pool must publish lifecycle counters before a worker can observe them.
		How: `ElasticThreadPool.startWorkerLocked` performs the later start transition.
	**/
	public function new(owner:ElasticThreadPool, available:Lock, timeout:Float) {
		this.owner = owner;
		this.available = available;
		this.timeout = timeout;
	}

	/**
		What: begin one worker generation.
		Why: timed-out worker objects are reused without allocating unbounded state.
		How: the pool calls this only while holding its lifecycle mutex.
	**/
	function start():Void {
		dead = false;
		task = null;
		Thread.create(loop);
	}

	/**
		What: consume tasks until timeout or a shutdown exit token.
		Why: callbacks must not run while holding the pool lifecycle mutex.
		How: each iteration performs a timed token wait, lets the pool atomically
		assign a queued task, then reports callback completion.
	**/
	function loop():Void {
		while (true) {
			var woke = available.wait(timeout);
			if (!owner.workerResolveWait(this, woke)) {
				return;
			}
			switch task {
				case null:
					continue;
				case fn:
					try {
						fn();
					} catch (err:Dynamic) {
						// Dynamic catch is intentional: Haxe tasks may throw any Haxe value. The
						// pool accounts for this task before preserving the original throw.
						owner.workerTaskFailed(this);
						throw err;
					}
					owner.workerTaskFinished(this);
			}
		}
	}
}
