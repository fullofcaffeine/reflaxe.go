package hxrt.http;

/**
	What
	- Opaque typed handle for one completed native HTTP response.

	Why
	- Go response bodies and header maps are transport results, while callback
	  ordering, public maps, status policy, and Haxe bytes belong in staged
	  `sys.Http`.

	How
	- Map directly to `hxrt.HttpResponse`. Staged source reads scalar fields,
	  opaque byte views, and indexed header values through `NativeHttp`.
**/
@:go.import("hxrt")
@:go.package("hxrt")
@:go.name("HttpResponse")
extern class HttpResponseHandle {}
