package sys.thread;

import haxe.Exception;

/**
	Thrown when code requires an event loop but the current thread does not have
	one.

	Why
	The mainstream Haxe stdlib implementation cannot be used unchanged on `haxe.go`.
	This exception is part of the public `sys.thread` API even before the Go
	event-loop contract lands.

	What
	Straight staged exception subtype.

	How
	Keeps the upstream constructor/message shape unchanged.
**/
class NoEventLoopException extends Exception {
	public function new(msg:String = "Event loop is not available. Refer to sys.thread.Thread.runWithEventLoop.", ?previous:Exception) {
		super(msg, previous);
	}
}
