package hxrt.io;

/**
	What: Typed access to the target's IEEE-754 bit reinterpretation primitives.

	Why: The mainstream `haxe.io.FPHelper` fallback cannot be used unchanged on
	`haxe.go`: implementing bit conversion through `BytesInput` and `BytesOutput`
	would recurse because those streams themselves call `FPHelper`. Reinterpreting
	Go floating-point storage is a genuine runtime capability, not portable stream
	policy and not a reason for a compiler-owned stdlib shim.

	How: Exchange only numeric scalar values with `runtime/hxrt/bytes.go`.
	`FPHelper` continues to own the public word ordering and `Int64` construction.
**/
@:go.import("hxrt")
@:go.package("hxrt")
extern class NativeFloatBits {
	@:go.name("Float32FromBits")
	public static function float32FromBits(value:Int):Float;

	@:go.name("Float32Bits")
	public static function float32Bits(value:Float):Int;

	@:go.name("Float64FromWords")
	public static function float64FromWords(low:Int, high:Int):Float;

	@:go.name("Float64LowWord")
	public static function float64LowWord(value:Float):Int;

	@:go.name("Float64HighWord")
	public static function float64HighWord(value:Float):Int;
}
