package sys.thread;

/**
	What
	- Repo-authored worker support for the staged `sys.thread.FixedThreadPool`
	  implementation.

	Why
	- This helper is not an upstream Haxe stdlib module and must not package as an
	  override-only `.cross.hx` file. Keeping it as ordinary source also makes its
	  dependency on the private shutdown sentinel explicit.

	How
	- The worker blocks on the shared deque, executes each task, and exits only when
	  `FixedThreadPoolShutdownException` is raised by the pool shutdown path.
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
