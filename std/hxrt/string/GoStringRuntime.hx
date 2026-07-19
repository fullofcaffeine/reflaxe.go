package hxrt.string;

/**
	What
	- Typed access to the representation-level string operations and exact native
	  float conversion that staged Haxe source cannot express directly.

	Why
	- Haxe strings lower to pointer-backed `hxrt` values on this target. Counting,
	  reading, and slicing Go runes therefore need the real runtime representation,
	  but Haxe library rules such as bounds normalization must stay in staged std.

	How
	- Go metadata maps each typed method to one narrow `hxrt` helper. Callers decide
	  valid indices and ranges before crossing this boundary; `sliceCodePoints`
	  only converts an already-normalized code-point range into a Go string slice.
	- `parseFloatExact` receives a prefix already validated by staged `Std`; it does
	  not own whitespace, prefix scanning, or malformed-exponent policy.
**/
@:go.import("hxrt")
@:go.package("hxrt")
extern class GoStringRuntime {
	@:go.name("StringLengthStringPtr")
	public static function length(value:String):Int;

	@:go.name("StringCharCodeAtStringPtr")
	public static function charCodeAt(value:String, index:Int):Int;

	@:go.name("StringSliceCodePointsStringPtr")
	public static function sliceCodePoints(value:String, start:Int, end:Int):String;

	@:go.name("StringParseFloatExact")
	public static function parseFloatExact(value:String):Float;
}
