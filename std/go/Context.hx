package go;

/**
	What
	- Typed facade for Go's native `context.Context` interface.

	Why
	- `context.Context` is a real Go API, not an upstream Haxe stdlib module, so it
	  belongs in the public `go.*` facade root rather than the override-only `_std`
	  tree.

	How
	- Import and name metadata map this interface directly to `context.Context`
	  while keeping application code typed and free of raw injection.
**/
@:go.import("context")
@:go.name("Context")
extern interface Context {}
