package sys.thread;

/**
	Double-ended queue with optional blocking pop.

	Why
	The upstream stdlib uses `Deque<T>` as a portable synchronization helper and
	as the queue behind thread-pool implementations.

	What
	Pure Haxe implementation on top of `Mutex` and `Lock`.

	How
	Each enqueue releases the availability lock exactly once. `pop(true)` blocks on
	that lock, while `pop(false)` returns `null` when no item is available.
**/
@:coreApi
class Deque<T> {
	final __mutex = new Mutex();
	final __available = new Lock();
	final __items = new haxe.ds.List<T>();

	public function new():Void {}

	public function add(i:T):Void {
		__mutex.acquire();
		__items.add(i);
		__mutex.release();
		__available.release();
	}

	public function push(i:T):Void {
		__mutex.acquire();
		__items.push(i);
		__mutex.release();
		__available.release();
	}

	public function pop(block:Bool):Null<T> {
		if (block) {
			__available.wait();
		} else if (!__available.wait(0.0)) {
			return null;
		}

		__mutex.acquire();
		var value = __items.pop();
		__mutex.release();
		return value;
	}
}
