package go;

/**
	What
	- Package-level facade for Go `context` entry points.

	Why
	- These functions expose a real Go-native package surface. Keeping the wrapper
	  in ordinary `std/go` makes it publicly importable without treating it as an
	  upstream Haxe stdlib override.

	How
	- Package and member metadata lower `background()` to `context.Background`
	  and return the typed `go.Context` facade.
**/
@:go.import("context")
extern class ContextPkg {
	@:go.name("Background")
	public static function background():Context;
}
