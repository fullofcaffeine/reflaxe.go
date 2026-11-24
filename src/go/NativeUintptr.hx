package go;

/**
	What: One exact Go `uintptr` at a typed native boundary.

	Why: Haxe `Int` lowers to Go `int`, and Go has no implicit conversion between
	`int` and `uintptr`. APIs such as `(*os.File).Fd`, `os.NewFile`,
	`runtime.Caller` and `runtime.FuncForPC` therefore cannot be declared with
	`Int`. `@:go.valueType` cannot name the builtin type either.

	How: The compiler lowers this opaque extern marker to Go `uintptr`. Obtain a
	value through a precise native extern, store it in locals, parameters, fields
	and tuple carriers, and pass it back to Go unchanged. It has no constructor,
	arithmetic, implicit numeric conversion or nullable representation. The two
	static conversions are the only bridge to Haxe `Int`; call them directly.
**/
extern class NativeUintptr {
	/**
		Converts with Go `int(value)`, as in `term.GetSize(int(os.Stdout.Fd()))`.

		A value above the Go `int` range wraps, as in Go. The result is not reduced
		to 32 bits. Call directly; method values are not supported.
	**/
	public static function toInt(value:NativeUintptr):Int;

	/**
		Converts with Go `uintptr(value)`. A negative `Int` wraps, as in Go.
		Call directly; method values are not supported.
	**/
	public static function fromInt(value:Int):NativeUintptr;
}
