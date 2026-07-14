package haxe.io;

/**
	What
	A staged `haxe.io.UInt8Array` override for `haxe.go`.

	Why
	The mainstream Haxe stdlib implementation cannot be used unchanged on `haxe.go`.
	- The public typed-array API is portable library behavior.
	- Explicit `Null<Int>` optional parameters keep omitted-argument semantics
	  intact on `haxe.go` without adding compiler-only special cases.

	How
	- Reuse `ArrayBufferView` as the carrier, exactly like upstream.
	- Keep byte access and aliasing semantics unchanged.
**/
typedef UInt8ArrayData = ArrayBufferView.ArrayBufferViewData;

abstract UInt8Array(UInt8ArrayData) {
	public static inline var BYTES_PER_ELEMENT = 1;

	public var length(get, never):Int;
	public var view(get, never):ArrayBufferView;

	public inline function new(elements:Int) {
		this = new ArrayBufferView(elements * BYTES_PER_ELEMENT).getData();
	}

	inline function get_length():Int {
		return this.byteLength;
	}

	public inline function get_view():ArrayBufferView {
		return ArrayBufferView.fromData(this);
	}

	@:arrayAccess public inline function get(index:Int):Int {
		return this.bytes.get(index + this.byteOffset);
	}

	@:arrayAccess public inline function set(index:Int, value:Int):Int {
		if (index >= 0 && index < length) {
			this.bytes.set(index + this.byteOffset, value);
			return value;
		}
		return 0;
	}

	public inline function sub(begin:Int, length:Null<Int> = null):UInt8Array {
		return fromData(this.sub(begin, length));
	}

	public inline function subarray(begin:Null<Int> = null, end:Null<Int> = null):UInt8Array {
		return fromData(this.subarray(begin, end));
	}

	public inline function getData():UInt8ArrayData {
		return this;
	}

	public static function fromData(d:UInt8ArrayData):UInt8Array {
		return cast d;
	}

	public static function fromArray(a:Array<Int>, pos:Int = 0, length:Null<Int> = null):UInt8Array {
		var resolvedLength:Int = length == null ? a.length - pos : length;
		if (pos < 0 || resolvedLength < 0 || pos + resolvedLength > a.length) {
			throw Error.OutsideBounds;
		}
		var out = new UInt8Array(resolvedLength);
		for (idx in 0...resolvedLength) {
			out[idx] = a[idx + pos];
		}
		return out;
	}

	public static function fromBytes(bytes:haxe.io.Bytes, bytePos:Int = 0, length:Null<Int> = null):UInt8Array {
		return fromData(ArrayBufferView.fromBytes(bytes, bytePos, length).getData());
	}
}
