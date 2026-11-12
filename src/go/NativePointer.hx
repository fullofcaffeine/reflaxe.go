package go;

/**
	What: Owns one addressable native Go scalar for an explicit pointer boundary.
	Why: Go APIs such as variadic scanners mutate `*int`, `*string`, or `*bool`
		destinations. Passing an ordinary Haxe local through `...any` boxes its
		value instead, so native writes cannot reach it.
	How: `of` copies the input into separate Go storage and `get` reads that
		storage. Passing the carrier to an imported Go extern retains its pointer
		identity. Only `Int`, `String`, and `Bool` are admitted; this does not alias
		or mutate the local used to create it. Go owns the pointer's escape lifetime.
**/
extern class NativePointer<T> {
	/** Copies a supported scalar into fresh addressable native storage. Call directly; method values are not supported. */
	public static function of<T>(value:T):NativePointer<T>;

	/** Reads the current native value; String results are detached snapshots. Call directly; method values are not supported. */
	public function get():T;
}
