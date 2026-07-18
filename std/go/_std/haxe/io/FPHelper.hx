package haxe.io;

import haxe.Int64;
import hxrt.io.NativeFloatBits;

/**
	What
	A staged `haxe.io.FPHelper` override for `haxe.go`.

	Why
	The mainstream Haxe stdlib implementation cannot be used unchanged on `haxe.go`.
	- The public `FPHelper` API is portable stdlib behavior and does not need to
	  live in `GoCompiler`.
	- Implementing the conversion through `BytesInput` / `BytesOutput` would
	  recurse because those stream methods call `FPHelper` themselves.
	- Reinterpreting Go floating-point storage is a target capability that staged
	  Haxe cannot express with ordinary arithmetic.

	How
	- Delegate only raw IEEE-754 bit reinterpretation to the typed
	  `hxrt.io.NativeFloatBits` boundary.
	- Keep public word ordering and `Int64` construction in staged Haxe.
**/
class FPHelper {
	public static function i32ToFloat(i:Int):Float {
		return NativeFloatBits.float32FromBits(i);
	}

	public static function floatToI32(f:Float):Int {
		return NativeFloatBits.float32Bits(f);
	}

	public static function i64ToDouble(low:Int, high:Int):Float {
		return NativeFloatBits.float64FromWords(low, high);
	}

	public static function doubleToI64(v:Float):Int64 {
		return Int64.make(NativeFloatBits.float64HighWord(v), NativeFloatBits.float64LowWord(v));
	}
}
