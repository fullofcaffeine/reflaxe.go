import go.NativeAny;

/** A `@:goNative` class whose go-free static `plain` has a sibling with a `go.*` argument type. */
@:goNative
class Tools {
	public static function plain():String {
		return describe(NativeAny.ofString("class"));
	}

	static function describe(value:NativeAny):String {
		return value == null ? "nil" : "class value";
	}
}
