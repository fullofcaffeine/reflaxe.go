package sys.thread;

/**
	Internal worker used by `FixedThreadPool`.
**/
class FixedThreadPoolWorker {
	final queue:Deque<() -> Void>;

	public function new(queue:Deque<() -> Void>) {
		this.queue = queue;
		Thread.create(loop);
	}

	function loop():Void {
		try {
			while (true) {
				var task = queue.pop(true);
				if (task != null) {
					task();
				}
			}
		} catch (_:FixedThreadPoolShutdownException) {}
	}
}
