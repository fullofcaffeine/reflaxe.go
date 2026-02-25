package go;

/**
	Typed wrapper for the Go `net/http` package used by interop smoke coverage.
**/
@:go.import("net/http")
extern class Http {
	@:go.name("StatusText")
	public static function statusText(code:Int):String;
}
