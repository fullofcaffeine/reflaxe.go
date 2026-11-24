package go;

/**
	What: One exact Go `uint`, the unsigned platform word, at a typed native
	boundary.

	Why: Haxe `Int` lowers to Go `int`, and Go has no implicit conversion between
	`int` and `uint`, so APIs such as `bits.Len` or `bits.OnesCount` cannot be
	declared with `Int`. `@:go.valueType` cannot name the builtin type either.

	How: The compiler lowers this opaque extern marker to Go `uint`, which has 32
	or 64 bits as the Go target does. It has no constructor, arithmetic,
	implicit numeric conversion or nullable representation. The two static
	conversions are the only bridge to Haxe `Int`; call them directly.
**/
extern class NativeUint {
	/**
		Converts with Go `int(value)`. A value above the Go `int` range wraps, as
		in Go, and the result is not reduced to 32 bits. Call directly; method
		values are not supported.
	**/
	public static function toInt(value:NativeUint):Int;

	/**
		Converts with Go `uint(value)`. A negative `Int` wraps, as in Go:
		`fromInt(-1)` has every bit set. Call directly; method values are not
		supported.
	**/
	public static function fromInt(value:Int):NativeUint;
}
