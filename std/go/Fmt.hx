package go;

/**
	What
	- Typed facade for the Go `fmt` package operations used by baseline interop.

	Why
	- `fmt` is a real Go-native API, so its public wrapper belongs under ordinary
	  `std/go` rather than the upstream-override `_std` tree. Framework ownership
	  also keeps application code from needing raw extern declarations.

	How
	- Import and member metadata lower `println` directly to `fmt.Println`.
**/
@:go.import("fmt")
extern class Fmt {
	@:go.name("Println")
	public static function println<T>(value:T):Void;
}
