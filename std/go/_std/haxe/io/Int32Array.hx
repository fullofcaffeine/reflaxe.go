package haxe.io;

/**
	What
	A staged `haxe.io.Int32Array` override for `haxe.go`.

	Why
	The mainstream Haxe stdlib implementation cannot be used unchanged on `haxe.go`.
	- Keeps the public typed-array abstraction library-owned.
	- Uses explicit `Null<Int>` optional lengths so omitted-argument behavior stays
	  correct on `haxe.go`.

	How
	- Preserve the same `ArrayBufferView` carrier and little-endian Int32 access.
**/
typedef Int32ArrayData = ArrayBufferView.ArrayBufferViewData;

abstract Int32Array(Int32ArrayData) {
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

	@:arrayAccess public inline function get(index:Int):Int {
		return this.bytes.getInt32((index << 2) + this.byteOffset);
	}

	@:arrayAccess public inline function set(index:Int, value:Int):Int {
		if (index >= 0 && index < length) {
			this.bytes.setInt32((index << 2) + this.byteOffset, value);
			return value;
		}
		return 0;
	}

	public inline function sub(begin:Int, length:Null<Int> = null):Int32Array {
		var scaledLength:Null<Int> = length == null ? null : length << 2;
		return fromData(this.sub(begin << 2, scaledLength));
	}

	public inline function subarray(begin:Null<Int> = null, end:Null<Int> = null):Int32Array {
		var scaledBegin:Null<Int> = begin == null ? null : begin << 2;
		var scaledEnd:Null<Int> = end == null ? null : end << 2;
		return fromData(this.subarray(scaledBegin, scaledEnd));
	}

	public inline function getData():Int32ArrayData {
		return this;
	}

	public static function fromData(d:Int32ArrayData):Int32Array {
		return cast d;
	}

	public static function fromArray(a:Array<Int>, pos:Int = 0, length:Null<Int> = null):Int32Array {
		var resolvedLength:Int = length == null ? a.length - pos : length;
		if (pos < 0 || resolvedLength < 0 || pos + resolvedLength > a.length) {
			throw Error.OutsideBounds;
		}
		var out = new Int32Array(resolvedLength);
		for (idx in 0...resolvedLength) {
			out[idx] = a[idx + pos];
		}
		return out;
	}

	public static function fromBytes(bytes:haxe.io.Bytes, bytePos:Int = 0, length:Null<Int> = null):Int32Array {
		var resolvedLength:Null<Int> = length == null ? null : length << 2;
		return fromData(ArrayBufferView.fromBytes(bytes, bytePos, resolvedLength).getData());
	}
}
