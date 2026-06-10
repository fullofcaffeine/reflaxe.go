package haxe.io;

/**
	What
	A staged `haxe.io.Float32Array` override for `haxe.go`.

	Why
	- The public typed-array API belongs in staged std.
	- The upstream implementation depends on `Bytes.getFloat` / `setFloat`, which
	  `haxe.go` does not expose on the compiler-owned `Bytes` carrier.
	- Explicit `Null<Int>` optional lengths keep omitted-argument behavior correct.

	How
	- Reuse the same `ArrayBufferView` carrier.
	- Convert element bits through staged `FPHelper` plus existing Int32 byte
	  access instead of growing the compiler-owned `Bytes` API for this surface.
**/
typedef Float32ArrayData = ArrayBufferView.ArrayBufferViewData;

abstract Float32Array(Float32ArrayData) {
	public static inline var BYTES_PER_ELEMENT = 4;

	public var length(get, never):Int;
	public var view(get, never):ArrayBufferView;

	public inline function new(elements:Int) {
		this = new ArrayBufferView(elements * BYTES_PER_ELEMENT).getData();
	}

	inline function get_length():Int {
		return this.byteLength >> 2;
	}

	public inline function get_view():ArrayBufferView {
		return ArrayBufferView.fromData(this);
	}

	@:arrayAccess public inline function get(index:Int):Float {
		return FPHelper.i32ToFloat(this.bytes.getInt32((index << 2) + this.byteOffset));
	}

	@:arrayAccess public inline function set(index:Int, value:Float):Float {
		if (index >= 0 && index < length) {
			this.bytes.setInt32((index << 2) + this.byteOffset, FPHelper.floatToI32(value));
			return value;
		}
		return 0;
	}

	public inline function sub(begin:Int, length:Null<Int> = null):Float32Array {
		var scaledLength:Null<Int> = length == null ? null : length << 2;
		return fromData(this.sub(begin << 2, scaledLength));
	}

	public inline function subarray(begin:Null<Int> = null, end:Null<Int> = null):Float32Array {
		var scaledBegin:Null<Int> = begin == null ? null : begin << 2;
		var scaledEnd:Null<Int> = end == null ? null : end << 2;
		return fromData(this.subarray(scaledBegin, scaledEnd));
	}

	public inline function getData():Float32ArrayData {
		return this;
	}

	public static function fromData(d:Float32ArrayData):Float32Array {
		return cast d;
	}

	public static function fromArray(a:Array<Float>, pos:Int = 0, length:Null<Int> = null):Float32Array {
		var resolvedLength:Int = length == null ? a.length - pos : length;
		if (pos < 0 || resolvedLength < 0 || pos + resolvedLength > a.length) {
			throw Error.OutsideBounds;
		}
		var out = new Float32Array(resolvedLength);
		for (idx in 0...resolvedLength) {
			out[idx] = a[idx + pos];
		}
		return out;
	}

	public static function fromBytes(bytes:haxe.io.Bytes, bytePos:Int = 0, length:Null<Int> = null):Float32Array {
		var resolvedLength:Null<Int> = length == null ? null : length << 2;
		return fromData(ArrayBufferView.fromBytes(bytes, bytePos, resolvedLength).getData());
	}
}
