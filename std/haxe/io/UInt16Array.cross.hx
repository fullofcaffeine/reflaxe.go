package haxe.io;

/**
	What
	A staged `haxe.io.UInt16Array` override for `haxe.go`.

	Why
	- This keeps the typed-array surface in staged std.
		- Explicit nullable optional lengths keep the source API honest while the
		  compiler preserves omitted primitive arguments at Go boundaries.

	How
	- Preserve the upstream `ArrayBufferView` carrier and little-endian indexing.
	- Express omitted lengths/endpoints with `Null<Int>` while keeping the same
	  public behavior.
**/
typedef UInt16ArrayData = ArrayBufferView.ArrayBufferViewData;

abstract UInt16Array(UInt16ArrayData) {
	public static inline var BYTES_PER_ELEMENT = 2;

	public var length(get, never):Int;
	public var view(get, never):ArrayBufferView;

	public inline function new(elements:Int) {
		this = new ArrayBufferView(elements * BYTES_PER_ELEMENT).getData();
	}

	inline function get_length():Int {
		return this.byteLength >> 1;
	}

	public inline function get_view():ArrayBufferView {
		return ArrayBufferView.fromData(this);
	}

	@:arrayAccess public inline function get(index:Int):Int {
		return this.bytes.getUInt16((index << 1) + this.byteOffset);
	}

	@:arrayAccess public inline function set(index:Int, value:Int):Int {
		if (index >= 0 && index < length) {
			this.bytes.setUInt16((index << 1) + this.byteOffset, value);
			return value;
		}
		return 0;
	}

	public inline function sub(begin:Int, length:Null<Int> = null):UInt16Array {
		var scaledLength:Null<Int> = length == null ? null : length << 1;
		return fromData(this.sub(begin << 1, scaledLength));
	}

	public inline function subarray(begin:Null<Int> = null, end:Null<Int> = null):UInt16Array {
		var scaledBegin:Null<Int> = begin == null ? null : begin << 1;
		var scaledEnd:Null<Int> = end == null ? null : end << 1;
		return fromData(this.subarray(scaledBegin, scaledEnd));
	}

	public inline function getData():UInt16ArrayData {
		return this;
	}

	public static function fromData(d:UInt16ArrayData):UInt16Array {
		return cast d;
	}

	public static function fromArray(a:Array<Int>, pos:Int = 0, length:Null<Int> = null):UInt16Array {
		var resolvedLength:Int = length == null ? a.length - pos : length;
		if (pos < 0 || resolvedLength < 0 || pos + resolvedLength > a.length) {
			throw Error.OutsideBounds;
		}
		var out = new UInt16Array(resolvedLength);
		for (idx in 0...resolvedLength) {
			out[idx] = a[idx + pos];
		}
		return out;
	}

	public static function fromBytes(bytes:haxe.io.Bytes, bytePos:Int = 0, length:Null<Int> = null):UInt16Array {
		var resolvedLength:Null<Int> = length == null ? null : length << 1;
		return fromData(ArrayBufferView.fromBytes(bytes, bytePos, resolvedLength).getData());
	}
}
