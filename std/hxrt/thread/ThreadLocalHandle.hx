package hxrt.thread;

/**
	What
	Opaque runtime handle for one staged `sys.thread.Tls<T>` instance.

	Why
	A raw integer slot would let ordinary Haxe code forge or reuse thread-local
	identity. The mainstream Haxe stdlib cannot supply a Go runtime handle whose
	values are reclaimed with the owning logical thread.

	How
	`NativeThread.localNew()` creates the matching `hxrt.ThreadLocalHandle`; only
	the typed bridge can read or write it.
**/
@:go.import("hxrt")
@:go.package("hxrt")
@:go.name("ThreadLocalHandle")
extern class ThreadLocalHandle {}
