import sys.thread.Condition;
import sys.thread.Deque;
import sys.thread.IThreadPool;
import sys.thread.Lock;
import sys.thread.Mutex;
import sys.thread.NoEventLoopException;
import sys.thread.Semaphore;
import sys.thread.ThreadPoolException;
import sys.thread.Tls;
import sys.thread.Thread;

private class DummyPool implements IThreadPool {
	public var threadsCount(get, never):Int;
	public var isShutdown(get, never):Bool;

	var _isShutdown = false;
	var runs = 0;

	public function new() {}

	function get_threadsCount():Int {
		return 0;
	}

	function get_isShutdown():Bool {
		return _isShutdown;
	}

	public function run(task:() -> Void):Void {
		if (_isShutdown) {
			throw new ThreadPoolException("shutdown");
		}
		runs++;
		task();
	}

	public function shutdown():Void {
		_isShutdown = true;
	}

	public function runCount():Int {
		return runs;
	}
}

class Main {
	static function main() {
		var lock = new Lock();
		lock.release();
		Sys.println("lock.release_before_wait=" + lock.wait());
		Sys.println("lock.timeout_empty=" + lock.wait(0.0));

		var mutex = new Mutex();
		mutex.acquire();
		mutex.acquire();
		Sys.println("mutex.try_reentrant=" + mutex.tryAcquire());
		mutex.release();
		mutex.release();
		mutex.release();
		Sys.println("mutex.try_after_release=" + mutex.tryAcquire());
		mutex.release();

		var condition = new Condition();
		condition.acquire();
		Sys.println("condition.try_reentrant=" + condition.tryAcquire());
		condition.release();
		condition.release();

		var sem = new Semaphore(1);
		Sys.println("sem.try_first=" + sem.tryAcquire());
		Sys.println("sem.try_empty=" + sem.tryAcquire(0.0));
		sem.release();
		sem.acquire();
		Sys.println("sem.try_after_acquire=" + sem.tryAcquire(0.0));

		var deque = new Deque<String>();
		deque.add("tail");
		deque.push("head");
		Sys.println("deque.pop1=" + deque.pop(false));
		Sys.println("deque.pop2=" + deque.pop(false));
		Sys.println("deque.pop3=" + Std.string(deque.pop(false)));

		var tls = new Tls<String>();
		Sys.println("tls.initial=" + Std.string(tls.value));
		tls.value = "worker";
		Sys.println("tls.after_set=" + Std.string(tls.value));
		tls.value = null;
		Sys.println("tls.after_clear=" + Std.string(tls.value));
		tls.value = "main";
		var tlsDone = new Lock();
		Thread.create(function() {
			Sys.println("tls.worker_initial=" + Std.string(tls.value));
			tls.value = "child";
			Sys.println("tls.worker_set=" + Std.string(tls.value));
			tlsDone.release();
		});
		tlsDone.wait();
		Sys.println("tls.main_after_worker=" + Std.string(tls.value));
		tls.value = null;

		var noLoop = new NoEventLoopException();
		Sys.println("noLoop.msg=" + noLoop.message);

		var pool = new DummyPool();
		pool.run(function() Sys.println("pool.task=ran"));
		Sys.println("pool.runs=" + pool.runCount());
		pool.shutdown();
		Sys.println("pool.shutdown=" + pool.isShutdown);
		try {
			pool.run(function() {});
		} catch (err:ThreadPoolException) {
			Sys.println("pool.error=" + err.message);
		}
	}
}
