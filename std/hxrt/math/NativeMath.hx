package hxrt.math;

/**
	What:
	- Maps staged Haxe `Math` operations directly to Go's typed `math` package.

	Why:
	- These are already representation-neutral Go functions, so adding compiler
	  shims or an `hxrt` Go wrapper would create ownership with no native gap to
	  bridge.

	How:
	- Use explicit import and member metadata for every native operation. Haxe-only
	  rounding, finiteness, min/max, and signed-zero policy stays in `Math.hx`.
**/
@:go.import("math")
extern class NativeMath {
	@:go.name("Abs")
	public static function abs(value:Float):Float;

	@:go.name("Sin")
	public static function sin(value:Float):Float;

	@:go.name("Cos")
	public static function cos(value:Float):Float;

	@:go.name("Tan")
	public static function tan(value:Float):Float;

	@:go.name("Asin")
	public static function asin(value:Float):Float;

	@:go.name("Acos")
	public static function acos(value:Float):Float;

	@:go.name("Atan")
	public static function atan(value:Float):Float;

	@:go.name("Atan2")
	public static function atan2(y:Float, x:Float):Float;

	@:go.name("Exp")
	public static function exp(value:Float):Float;

	@:go.name("Log")
	public static function log(value:Float):Float;

	@:go.name("Pow")
	public static function pow(value:Float, exponent:Float):Float;

	@:go.name("Sqrt")
	public static function sqrt(value:Float):Float;

	@:go.name("Floor")
	public static function floor(value:Float):Float;

	@:go.name("Ceil")
	public static function ceil(value:Float):Float;

	@:go.name("IsInf")
	public static function isInf(value:Float, sign:Int):Bool;

	@:go.name("IsNaN")
	public static function isNaN(value:Float):Bool;

	@:go.name("Inf")
	public static function inf(sign:Int):Float;

	@:go.name("NaN")
	public static function nan():Float;
}
