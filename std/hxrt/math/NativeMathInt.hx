package hxrt.math;

/**
	What:
	- Provides the three Go `int` results required by Haxe `Math.floor`, `ceil`,
	  and `round`.

	Why:
	- Go's `math` package returns `float64` for rounding functions, while the Haxe
	  4.3.7 public signatures return `Int`. The staged compiler path cannot use the
	  still-migrating `Std.int` shim as a library conversion.

	How:
	- Bind to three narrow `hxrt` functions that perform the native Go conversion.
	  All Float-returning operations remain direct typed Go `math` extern calls.
**/
@:go.import("hxrt")
@:go.package("hxrt")
extern class NativeMathInt {
	@:go.name("MathFloorInt")
	public static function floor(value:Float):Int;

	@:go.name("MathCeilInt")
	public static function ceil(value:Float):Int;

	@:go.name("MathRoundInt")
	public static function round(value:Float):Int;
}
