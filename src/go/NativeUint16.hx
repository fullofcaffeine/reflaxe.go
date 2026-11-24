package go;

/**
	What: One exact Go `uint16` at a typed native boundary.

	Why: Haxe `Int` lowers to Go `int`, and Go has no implicit conversion between
	`int` and `uint16`, so APIs such as `bits.ReverseBytes16`,
	`netip.AddrPortFrom` or the `uint16` fields of `zip.FileHeader` cannot be
	declared with `Int`. `@:go.valueType` cannot name the builtin type either.

	How: The compiler lowers this opaque extern marker to Go `uint16`. It has no
	constructor, arithmetic, implicit numeric conversion or nullable
	representation. The two static conversions are the only bridge to Haxe
	`Int`; call them directly.
**/
extern class NativeUint16 {
	/** Converts with Go `int(value)`. Call directly; method values are not supported. */
	public static function toInt(value:NativeUint16):Int;

	/**
		Converts with Go `uint16(value)`. Go keeps the low 16 bits, so a negative
		or larger `Int` wraps: `fromInt(-1)` is `0xFFFF`. Call directly; method
		values are not supported.
	**/
	public static function fromInt(value:Int):NativeUint16;
}
