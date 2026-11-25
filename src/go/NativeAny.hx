package go;

/**
	What: One native Go value held in a Go `any`, for data that a Go API reads
	with reflection, such as the data argument of `text/template` or
	`encoding/json`.

	Why: Haxe values keep their Haxe carriers when they become `Dynamic`: a
	String is a `*string`, an anonymous object is a runtime map, and an Array is
	a shared `hxrt.Array`. Go code that inspects an `any` sees those carriers,
	so `%T` prints `*string` and `eq .Type "blocks"` fails in a template. A port
	of Go code that builds `map[string]interface{}` data needs the values that
	Go code would store instead.

	How: The compiler stores the abstract as a Go `any`. Each constructor stores
	the Go value that Go code would use: `ofString` a Go `string`, `ofInt` an
	`int`, `ofFloat` a `float64` (also for an integral literal such as `2`),
	`ofBool` a `bool`, `ofMap` a `map[string]any`, `ofSlice` a `[]any`,
	`ofStrings` a `[]string`, and `nil` a nil interface. The constructors are
	implicit casts, so `map.set("Title", title)` converts `title`. Call them
	directly; method values are not supported. A null String becomes "null", as
	it does at every Go `string` parameter; use `nil()` for Go `nil`.

	A `NativeAny` holds Go data only. Pass it to Go; do not read it back as a
	Haxe value. A `Dynamic` value unifies with this type without conversion and
	keeps its Haxe carrier, so convert typed values instead of passing Dynamic.
	`Any` is the honest storage for a Go `any`, and it stays inside this
	explicit go.* boundary.
**/
abstract NativeAny(Any) {
	/** Stores the Go `string` value of `value`. A null String becomes "null". */
	@:from public static extern function ofString(value:String):NativeAny;

	/** Stores a Go `int`. */
	@:from public static extern function ofInt(value:Int):NativeAny;

	/** Stores a Go `float64`, also for an integral value such as `2`. */
	@:from public static extern function ofFloat(value:Float):NativeAny;

	/** Stores a Go `bool`. */
	@:from public static extern function ofBool(value:Bool):NativeAny;

	/** Stores the Go `map[string]any` itself, not a copy. */
	@:from public static extern function ofMap(value:NativeAnyMap):NativeAny;

	/** Stores the Go `[]any` itself, not a copy. */
	@:from public static extern function ofSlice(value:NativeSlice<NativeAny>):NativeAny;

	/** Stores the Go `[]string` itself, not a copy. */
	@:from public static extern function ofStrings(value:NativeStringSlice):NativeAny;

	/** The Go nil interface, as Go code stores `nil` in a map. */
	public static extern function nil():NativeAny;
}
