package haxe.io;

import haxe.Int64;

/**
	What
	A staged `haxe.io.FPHelper` override for `haxe.go`.

	Why
	The mainstream Haxe stdlib implementation cannot be used unchanged on `haxe.go`.
	- The public `FPHelper` API is portable stdlib behavior and does not need to
	  live in `GoCompiler`.
	- The upstream fallback implementation assumes math helper surfaces that
	  `haxe.go` does not expose directly as public `Math_*` shims.
	- This backend already owns a little-endian `BytesInput` / `BytesOutput`
	  contract, so `FPHelper` can be expressed on top of that existing portable IO
	  surface instead of adding more raw Go.

	How
	- Encode values through `BytesOutput` with `bigEndian = false`.
	- Decode them again through `BytesInput`, preserving the same low-endian
	  contract the upstream helper documents.
**/
class FPHelper {
	static function littleEndianOutput():BytesOutput {
		var out = new BytesOutput();
		out.bigEndian = false;
		return out;
	}

	static function littleEndianInput(bytes:Bytes):BytesInput {
		var input = new BytesInput(bytes);
		input.bigEndian = false;
		return input;
	}

	public static function i32ToFloat(i:Int):Float {
		var out = littleEndianOutput();
		out.writeInt32(i);
		return littleEndianInput(out.getBytes()).readFloat();
	}

	public static function floatToI32(f:Float):Int {
		var out = littleEndianOutput();
		out.writeFloat(f);
		return littleEndianInput(out.getBytes()).readInt32();
	}

	public static function i64ToDouble(low:Int, high:Int):Float {
		var out = littleEndianOutput();
		out.writeInt32(low);
		out.writeInt32(high);
		return littleEndianInput(out.getBytes()).readDouble();
	}

	public static function doubleToI64(v:Float):Int64 {
		var out = littleEndianOutput();
		out.writeDouble(v);
		var input = littleEndianInput(out.getBytes());
		var low = input.readInt32();
		var high = input.readInt32();
		return Int64.make(high, low);
	}
}
