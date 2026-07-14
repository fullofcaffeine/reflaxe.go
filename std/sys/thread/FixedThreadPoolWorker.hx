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
	  `FixedThreadPoolShutdownException` is raised by the pool shutdown path. Any
	  other Haxe throw starts a replacement worker before the original value is
	  rethrown for normal portable-thread reporting, so queued accepted work drains.
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
		} catch (_:FixedThreadPoolShutdownException) {
			return;
		} catch (err:Dynamic) {
			// Dynamic catch is intentional: Haxe tasks may throw any Haxe value. Replace
			// this worker before preserving that value for the thread-level reporter.
			Thread.create(loop);
			throw err;
		}
	}
}
