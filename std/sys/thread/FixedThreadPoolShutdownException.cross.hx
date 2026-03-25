package sys.thread;

import haxe.Exception;

/**
	Internal shutdown sentinel for `FixedThreadPool`.
**/
class FixedThreadPoolShutdownException extends Exception {}
