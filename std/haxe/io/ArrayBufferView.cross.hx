package haxe.io;

/**
	What
	A staged `haxe.io.ArrayBufferView` override for `haxe.go`.

	Why
	- The public typed-array surface is library-level portable behavior and should
	  stay in staged std.
		- The upstream implementation relies on optional primitive arguments; `haxe.go`
		  now preserves those as nullable values at source-owned boundaries.

	How
	- Keep the same public API and carrier shape.
	- Spell nullable optional lengths/endpoints explicitly as `Null<Int>` so the
	  generated Go can preserve omitted-argument semantics without backend-only
	  special cases.
**/
typedef ArrayBufferViewData = ArrayBufferViewImpl;

class ArrayBufferViewImpl {
	public var bytes:haxe.io.Bytes;
	public var byteOffset:Int;
	public var byteLength:Int;

	public function new(bytes:Bytes, pos:Int, length:Int) {
		this.bytes = bytes;
		this.byteOffset = pos;
		this.byteLength = length;
	}

	public function sub(begin:Int, length:Null<Int> = null) {
		var resolvedLength:Int = length == null ? byteLength - begin : length;
		if (begin < 0 || resolvedLength < 0 || begin + resolvedLength > byteLength) {
			throw Error.OutsideBounds;
		}
		return new ArrayBufferViewImpl(bytes, byteOffset + begin, resolvedLength);
	}

	public function subarray(begin:Null<Int> = null, end:Null<Int> = null) {
		var resolvedBegin:Int = begin == null ? 0 : begin;
		var resolvedEnd:Int = end == null ? byteLength - resolvedBegin : end;
		return sub(resolvedBegin, resolvedEnd - resolvedBegin);
	}
}

abstract ArrayBufferView(ArrayBufferViewData) {
	public var buffer(get, never):haxe.io.Bytes;
	public var byteOffset(get, never):Int;
	public var byteLength(get, never):Int;

	public inline function new(size:Int) {
		this = new ArrayBufferViewData(haxe.io.Bytes.alloc(size), 0, size);
	}

	inline function get_byteOffset():Int
		return this.byteOffset;

	inline function get_byteLength():Int
		return this.byteLength;

	inline function get_buffer():haxe.io.Bytes
		return this.bytes;

	public inline function sub(begin:Int, length:Null<Int> = null):ArrayBufferView {
		return fromData(this.sub(begin, length));
	}

	public inline function subarray(begin:Null<Int> = null, end:Null<Int> = null):ArrayBufferView {
		return fromData(this.subarray(begin, end));
	}

	public inline function getData():ArrayBufferViewData {
		return this;
	}

	public static inline function fromData(a:ArrayBufferViewData):ArrayBufferView {
		return cast a;
	}

	public static function fromBytes(bytes:haxe.io.Bytes, pos:Int = 0, length:Null<Int> = null):ArrayBufferView {
		var resolvedLength:Int = length == null ? bytes.length - pos : length;
		if (pos < 0 || resolvedLength < 0 || pos + resolvedLength > bytes.length) {
			throw Error.OutsideBounds;
		}
		return fromData(new ArrayBufferViewData(bytes, pos, resolvedLength));
	}
}
