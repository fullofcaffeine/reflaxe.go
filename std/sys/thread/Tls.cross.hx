package sys.thread;

import haxe.ds.IntMap;
import hxrt.thread.NativeThread;

/**
	Thread-local storage for `haxe.go`.

	Why
	The portable `sys.thread.Tls<T>` API needs per-thread values even before the
	full public `sys.thread.Thread` surface is promoted.

	What
	Pure Haxe storage keyed by the runtime thread id.

	How
	A staged `IntMap` stores values by `hxrt.thread.NativeThread.currentId()`, and
	a small `Mutex` keeps the map safe across concurrent access.
**/
class Tls<T> {
	final __mutex = new Mutex();
	final __values:IntMap<T> = new IntMap();

	public var value(get, set):T;

	public function new():Void {}

	inline function currentId():Int {
		return NativeThread.currentId();
	}

	function get_value():T {
		var id = currentId();
		__mutex.acquire();
		var value = __values.get(id);
		__mutex.release();
		return value;
	}

	function set_value(v:T):T {
		var id = currentId();
		__mutex.acquire();
		if (v == null) {
			__values.remove(id);
		} else {
			__values.set(id, v);
		}
		__mutex.release();
		return v;
	}
}
