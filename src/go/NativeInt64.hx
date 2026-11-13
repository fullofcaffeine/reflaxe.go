package go;

/**
	What: An exact signed Go `int64` at a typed native boundary.

	Why: Haxe `Int` and `haxe.Int64` have different source and target semantics;
	neither can safely stand in for an imported Go scalar field or parameter.

	How: The compiler lowers this opaque extern marker to a Go `int64`. Obtain a
	value through a precise extern such as `strconv.ParseInt` or a native field,
	and use `strconv.FormatInt` for exact decimal text. It has no constructor,
	implicit numeric conversion, or nullable representation.
**/
extern class NativeInt64 {}
