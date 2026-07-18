package hxrt.net;

/**
	What
	- Typed opaque binding for one native TCP or UDP socket resource.

	Why
	- Go connections, listeners, buffered readers, deadlines, and socket options
	  cannot be represented as portable Haxe data. Using `Dynamic` would erase the
	  resource ownership boundary and make lifecycle mistakes easy to hide.

	How
	- Map directly to `hxrt.SocketHandle`; only `NativeSocket` capabilities consume
	  the value and staged `sys.net` classes retain it privately.
**/
@:go.import("hxrt")
@:go.package("hxrt")
@:go.name("SocketHandle")
extern class SocketHandle {}
