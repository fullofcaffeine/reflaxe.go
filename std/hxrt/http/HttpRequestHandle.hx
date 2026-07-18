package hxrt.http;

/**
	What
	- Opaque typed handle for one native HTTP request under construction.

	Why
	- Go URL parsing, transport configuration, and socket dialing need native
	  state, but exposing that state as an untyped value would erase the boundary
	  between staged `sys.Http` policy and `hxrt` transport.

	How
	- Map directly to `hxrt.HttpRequest`. Only `NativeHttp` builder capabilities
	  create or mutate the handle before one synchronous execution.
**/
@:go.import("hxrt")
@:go.package("hxrt")
@:go.name("HttpRequest")
extern class HttpRequestHandle {}
