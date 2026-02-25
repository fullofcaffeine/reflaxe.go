package go;

/**
	Typed wrapper for the Go `fmt` package.
	Keep this in framework-owned stdlib so app code does not need raw extern declarations
	for baseline interop.
**/
@:go.import("fmt")
extern class Fmt {
	@:go.name("Println")
	public static function println<T>(value:T):Void;
}
