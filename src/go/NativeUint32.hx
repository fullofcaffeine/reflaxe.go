package go;

/**
	What: One exact Go `uint32` at a typed native boundary.

	Why: Haxe `Int` lowers to Go `int`, and Go has no implicit conversion between
	`int` and `uint32`. An `Int` declaration of a `uint32` parameter, such as the
	`perm` of `unix.Open` or the byte counts of `windows.LockFileEx`, builds only
	while every caller passes a Go constant; a variable is a Go build error.
	`@:go.valueType` cannot name the builtin type either. A named type such as
	`fs.FileMode` stays its own `@:go.valueType` extern.

	How: The compiler lowers this opaque extern marker to Go `uint32`. Obtain a
	value through a precise native extern or `fromInt`, store it in locals,
	parameters, fields and tuple carriers, and pass it back to Go unchanged. It
	has no constructor, arithmetic, implicit numeric conversion or nullable
	representation. The two static conversions are the only bridge to Haxe
	`Int`; call them directly.
**/
extern class NativeUint32 {
	/**
		Converts with Go `int(value)`. The result is not reduced to 32 bits, so a
		value above `0x7FFFFFFF` stays positive. Call directly; method values are
		not supported.
	**/
	public static function toInt(value:NativeUint32):Int;

	/**
		Converts with Go `uint32(value)`. A negative `Int` wraps, as in Go:
		`fromInt(-1)` is `0xFFFFFFFF`. Call directly; method values are not
		supported.
	**/
	public static function fromInt(value:Int):NativeUint32;
}
