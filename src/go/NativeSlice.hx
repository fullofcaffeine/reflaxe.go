package go;

/**
	What: A typed view of one native Go slice at an explicit Go interop boundary.

	Why: Portable Haxe `Array<T>` has shared object identity and sparse-growth
	semantics that a Go slice does not provide. Runtime externs must therefore not
	declare native `[]T` values as `Array<T>` merely because both support indexing.

	How: The compiler maps this class directly to `[]T`. Indexing and `length` use
	ordinary Go slice operations. `fromArray` and `toArray` make the representation
	change explicit and return shallow copies, preserving element identity without
	sharing the incompatible collection headers.
**/
extern class NativeSlice<T> implements ArrayAccess<T> {
	/** The current native slice length. */
	public var length(default, null):Int;

	/** Copies a portable Haxe Array into native Go slice storage. */
	public static function fromArray<T>(value:Array<T>):NativeSlice<T>;

	/**
		Returns the native slice produced by Go's `append` built-in.

		The returned value must be retained because Go may replace the slice header
		when its backing storage grows.
	**/
	public static function append<T>(target:NativeSlice<T>, value:T):NativeSlice<T>;

	/** Copies this native Go slice into a portable shared-identity Haxe Array. */
	public function toArray():Array<T>;
}
