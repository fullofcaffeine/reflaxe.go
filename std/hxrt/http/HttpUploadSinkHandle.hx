package hxrt.http;

/**
	What
	- Opaque typed destination for one caller-driven native HTTP upload.

	Why
	- Staged `sys.Http` must read its Haxe `Input` on the synchronous public
	  caller, while timeout, cancellation, and early responses must still unblock
	  a native body write.

	How
	- Map directly to `hxrt.HttpUploadSink`. Staged source writes bounded immutable
	  chunks, then finishes exact-size input or aborts with its source error.
**/
@:go.import("hxrt")
@:go.package("hxrt")
@:go.name("HttpUploadSink")
extern class HttpUploadSinkHandle {}
