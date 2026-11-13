package go;

/**
	What: One exact Go `byte` (`uint8`) at a typed native boundary.

	Why: Haxe `Int` lowers to Go `int`; treating it as a byte makes generated
	`[]byte` signatures become incompatible `[]int` signatures.

	How: The compiler lowers this opaque extern marker to Go `byte`. Obtain a
	value through a precise native extern or a native byte slice. It has no
	constructor, implicit numeric conversion, arithmetic, or nullable scalar
	representation; `go.NativeSlice<NativeByte>` is a nil-capable Go `[]byte`.
**/
extern class NativeByte {}
