package sys.thread;

import hxrt.thread.NativeThread;
import hxrt.thread.ThreadLocalHandle;

/**
	Thread-local storage for `haxe.go`.

	Why
	The mainstream Haxe stdlib implementation cannot be used unchanged on `haxe.go`:
	Go exposes no supported goroutine-local storage key, while portable
	`sys.thread.Tls<T>` values must remain isolated and be released when a supported
	thread lifecycle ends.

	What
	A typed staged wrapper around a runtime-owned thread-local slot.

	How
	Each `Tls<T>` owns one opaque `ThreadLocalHandle`. Values live inside the
	current runtime `ThreadState`, so portable workers and compiler-owned detached
	`go.Go.spawn` goroutines release every stored value when their lifecycle ends.
	The unavoidable `Dynamic` representation is confined to the typed runtime
	bridge. Arbitrary foreign goroutines created outside those compiler-owned
	boundaries do not have an automatic detach transition.
**/
class Tls<T> {
	final __handle:ThreadLocalHandle;

	public var value(get, set):T;

	public function new():Void {
		__handle = NativeThread.localNew();
	}

	function get_value():T {
		return cast NativeThread.localGet(__handle);
	}

	function set_value(v:T):T {
		NativeThread.localSet(__handle, v);
		return v;
	}
}
