package haxe.io;

import haxe.Int64;

/**
	What
	A staged `haxe.io.Float64Array` override for `haxe.go`.

	Why
	- The public typed-array API should stay library-owned.
	- The upstream implementation depends on `Bytes.getDouble` / `setDouble`,
	  which `haxe.go` does not expose on the compiler-owned `Bytes` carrier.
	- Explicit `Null<Int>` optional lengths keep omitted-argument behavior correct.

	How
	- Reuse the same `ArrayBufferView` carrier.
	- Convert element bits through staged `FPHelper` and existing Int32 byte access.
**/
typedef Float64ArrayData = ArrayBufferView.ArrayBufferViewData;

abstract Float64Array(Float64ArrayData) {
	public static inline var BYTES_PER_ELEMENT = 8;

	public var length(get, never):Int;
	public var view(get, never):ArrayBufferView;

	public inline function new(elements:Int) {
		this = new ArrayBufferView(elements * BYTES_PER_ELEMENT).getData();
	}

	inline function get_length():Int {
		return this.byteLength >> 3;
	}

	public inline function get_view():ArrayBufferView {
		return ArrayBufferView.fromData(this);
	}

	@:arrayAccess public inline function get(index:Int):Float {
		var pos = (index << 3) + this.byteOffset;
		var low = this.bytes.getInt32(pos);
		var high = this.bytes.getInt32(pos + 4);
		return FPHelper.i64ToDouble(low, high);
	}

	@:arrayAccess public inline function set(index:Int, value:Float):Float {
		if (index >= 0 && index < length) {
			var pos = (index << 3) + this.byteOffset;
			var bits:Int64 = FPHelper.doubleToI64(value);
			this.bytes.setInt32(pos, bits.low);
			this.bytes.setInt32(pos + 4, bits.high);
			return value;
		}
		return 0;
	}

	public inline function sub(begin:Int, length:Null<Int> = null):Float64Array {
		var scaledLength:Null<Int> = length == null ? null : length << 3;
		return fromData(this.sub(begin << 3, scaledLength));
	}

	public inline function subarray(begin:Null<Int> = null, end:Null<Int> = null):Float64Array {
		var scaledBegin:Null<Int> = begin == null ? null : begin << 3;
		var scaledEnd:Null<Int> = end == null ? null : end << 3;
		return fromData(this.subarray(scaledBegin, scaledEnd));
	}

	public inline function getData():Float64ArrayData {
		return this;
	}

	public static function fromData(d:Float64ArrayData):Float64Array {
		return cast d;
	}

	public static function fromArray(a:Array<Float>, pos:Int = 0, length:Null<Int> = null):Float64Array {
		var resolvedLength:Int = length == null ? a.length - pos : length;
		if (pos < 0 || resolvedLength < 0 || pos + resolvedLength > a.length) {
			throw Error.OutsideBounds;
		}
		var out = new Float64Array(resolvedLength);
		for (idx in 0...resolvedLength) {
			out[idx] = a[idx + pos];
		}
		return out;
	}

	public static function fromBytes(bytes:haxe.io.Bytes, bytePos:Int = 0, length:Null<Int> = null):Float64Array {
		var resolvedLength:Null<Int> = length == null ? null : length << 3;
		return fromData(ArrayBufferView.fromBytes(bytes, bytePos, resolvedLength).getData());
	}
}
