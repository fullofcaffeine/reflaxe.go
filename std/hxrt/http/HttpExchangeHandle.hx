package hxrt.http;

/**
	What
	- Opaque typed handle for one live native HTTP response exchange.

	Why
	- Go owns the response body, cancellation context, transport, and optional
	  socket until staged `sys.Http` finishes its Haxe-visible callback lifecycle.
	  A completed byte carrier would delay status and lose partial progress.

	How
	- Map directly to `hxrt.HttpExchange`. Staged source reads bounded immutable
	  chunks and must call `closeExchange` or `cancelExchange` exactly once.
**/
@:go.import("hxrt")
@:go.package("hxrt")
@:go.name("HttpExchange")
extern class HttpExchangeHandle {}
