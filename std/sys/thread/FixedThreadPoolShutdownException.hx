package sys.thread;

import haxe.Exception;

/**
	What
	- Private shutdown sentinel shared by the staged fixed-pool implementation and
	  its worker.

	Why
	- The sentinel is repo-authored target support, not an upstream stdlib override.
	  It therefore remains ordinary source beside the worker instead of entering the
	  override-only `_std` package mapping.

	How
	- The pool enqueues this typed exception and the worker catches only this type,
	  separating intentional shutdown from task failures.
**/
class FixedThreadPoolShutdownException extends Exception {}
