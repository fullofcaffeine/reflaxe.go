package haxe.io;

import hxrt.io.NativeBytes;

/**
	What: Accumulates bytes and numeric values into a growable buffer.

	Why: The mainstream Haxe stdlib implementation cannot be used unchanged on `haxe.go`
	because it expects a target-provided growable byte carrier. A native Go slice
	is the carrier here, but buffer policy and bounds remain library behavior.

	How: Keep a `BytesData` slice and call narrow typed append primitives only when
	Go must replace the slice header after growth.
**/
class BytesBuffer {
	var b:BytesData;

	public var length(get, never):Int;

	public function new() {
		b = NativeBytes.allocValues(0);
	}

	inline function get_length():Int {
		return b.length;
	}

	public inline function addByte(value:Int):Void {
		b = NativeBytes.appendByte(b, value);
	}

	public inline function add(source:Bytes):Void {
		b = NativeBytes.appendValues(b, source.getData());
	}

	public inline function addString(value:String, ?encoding:Encoding):Void {
		add(Bytes.ofString(value, encoding));
	}

	public function addInt32(value:Int):Void {
		addByte(value & 0xFF);
		addByte((value >> 8) & 0xFF);
		addByte((value >> 16) & 0xFF);
		addByte(value >>> 24);
	}

	public function addInt64(value:haxe.Int64):Void {
		addInt32(value.low);
		addInt32(value.high);
	}

	public inline function addFloat(value:Float):Void {
		addInt32(FPHelper.floatToI32(value));
	}

	public inline function addDouble(value:Float):Void {
		addInt64(FPHelper.doubleToI64(value));
	}

	public inline function addBytes(source:Bytes, pos:Int, len:Int):Void {
		if (pos < 0 || len < 0 || pos + len > source.length)
			throw Error.OutsideBounds;
		b = NativeBytes.appendSlice(b, source.getData(), pos, len);
	}

	public function getBytes():Bytes {
		return Bytes.ofData(NativeBytes.cloneValues(b));
	}
}
