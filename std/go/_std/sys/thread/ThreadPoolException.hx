package sys.thread;

import haxe.Exception;

/**
	Exception thrown by `sys.thread` pool implementations.

	Why
	The mainstream Haxe stdlib implementation cannot be used unchanged on `haxe.go`.
	The interface and direct user code need a concrete exception type even before
	the Go thread-pool implementations are promoted.

	What
	Plain staged exception subtype.

	How
	Mirrors the upstream type so staged and user implementations can throw it
	without backend-owned shims.
**/
class ThreadPoolException extends Exception {}
