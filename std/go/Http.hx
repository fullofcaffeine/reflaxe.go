package go;

/**
	What
	- Typed facade for the Go `net/http` package helper used by interop coverage.

	Why
	- This models a real Go-native API rather than replacing an upstream Haxe stdlib
	  module, so it remains a normal public `go.*` module outside `_std`.

	How
	- Import and member metadata lower `statusText` directly to
	  `net/http.StatusText`.
**/
@:go.import("net/http")
extern class Http {
	@:go.name("StatusText")
	public static function statusText(code:Int):String;
}
