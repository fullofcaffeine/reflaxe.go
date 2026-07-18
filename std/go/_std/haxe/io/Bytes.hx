package haxe.io;

import hxrt.io.ByteView;
import hxrt.io.NativeBytes;

/**
	What: Implements Haxe byte storage, numeric access, slicing, and encoding policy.

	Why: The mainstream Haxe stdlib implementation cannot be used unchanged on `haxe.go`,
	where `BytesData` is an aliasing native `[]int`
	view. Keeping the whole API in `GoCompiler` nevertheless obscured portable
	validation and algorithms behind raw generated declarations.

	How: Keep public behavior in ordinary staged Haxe over `BytesData`. Use the
	typed `NativeBytes` boundary only for native allocation and string conversion,
	and cache an opaque `ByteView`; every mutating method invalidates that cache.
**/
class Bytes {
	public var length(default, null):Int;

	var b:BytesData;
	var __hx_raw:ByteView;
	var __hx_rawValid:Bool;
	var __hx_dataExposed:Bool;

	function new(length:Int, b:BytesData, ?raw:ByteView) {
		this.length = length;
		this.b = b;
		this.__hx_raw = raw;
		this.__hx_rawValid = raw != null;
		this.__hx_dataExposed = false;
	}

	public inline function get(pos:Int):Int {
		return b[pos];
	}

	public inline function set(pos:Int, value:Int):Void {
		b[pos] = value & 0xFF;
		__hx_rawValid = false;
	}

	public function blit(pos:Int, src:Bytes, srcpos:Int, len:Int):Void {
		if (pos < 0 || srcpos < 0 || len < 0 || pos + len > length || srcpos + len > src.length)
			throw Error.OutsideBounds;
		if (len == 0)
			return;
		NativeBytes.blitValues(b, pos, src.b, srcpos, len);
		__hx_rawValid = false;
	}

	public function fill(pos:Int, len:Int, value:Int):Void {
		if (pos < 0 || len < 0 || pos + len > length)
			throw Error.OutsideBounds;
		var masked = value & 0xFF;
		for (index in 0...len)
			b[pos + index] = masked;
		__hx_rawValid = false;
	}

	public function sub(pos:Int, len:Int):Bytes {
		if (pos < 0 || len < 0 || pos + len > length)
			throw Error.OutsideBounds;
		var out = alloc(len);
		for (index in 0...len)
			out.b[index] = b[pos + index];
		return out;
	}

	public function compare(other:Bytes):Int {
		var limit = length < other.length ? length : other.length;
		for (index in 0...limit) {
			if (b[index] < other.b[index])
				return -1;
			if (b[index] > other.b[index])
				return 1;
		}
		return length < other.length ? -1 : length > other.length ? 1 : 0;
	}

	public function getDouble(pos:Int):Float {
		return FPHelper.i64ToDouble(getInt32(pos), getInt32(pos + 4));
	}

	public function getFloat(pos:Int):Float {
		return FPHelper.i32ToFloat(getInt32(pos));
	}

	public function setDouble(pos:Int, value:Float):Void {
		var bits = FPHelper.doubleToI64(value);
		setInt32(pos, bits.low);
		setInt32(pos + 4, bits.high);
	}

	public function setFloat(pos:Int, value:Float):Void {
		setInt32(pos, FPHelper.floatToI32(value));
	}

	public inline function getUInt16(pos:Int):Int {
		return get(pos) | (get(pos + 1) << 8);
	}

	public inline function setUInt16(pos:Int, value:Int):Void {
		set(pos, value);
		set(pos + 1, value >> 8);
	}

	public inline function getInt32(pos:Int):Int {
		return get(pos) | (get(pos + 1) << 8) | (get(pos + 2) << 16) | (get(pos + 3) << 24);
	}

	public inline function getInt64(pos:Int):haxe.Int64 {
		return haxe.Int64.make(getInt32(pos + 4), getInt32(pos));
	}

	public inline function setInt32(pos:Int, value:Int):Void {
		set(pos, value);
		set(pos + 1, value >> 8);
		set(pos + 2, value >> 16);
		set(pos + 3, value >>> 24);
	}

	public inline function setInt64(pos:Int, value:haxe.Int64):Void {
		setInt32(pos, value.low);
		setInt32(pos + 4, value.high);
	}

	public function getString(pos:Int, len:Int, ?encoding:Encoding):String {
		if (pos < 0 || len < 0 || pos + len > length)
			throw Error.OutsideBounds;
		return NativeBytes.stringFromView(__hx_nativeView(), pos, len, encoding == RawNative && rawNativeUsesUtf16LE());
	}

	@:deprecated("readString is deprecated, use getString instead")
	@:noCompletion
	public inline function readString(pos:Int, len:Int):String {
		return getString(pos, len);
	}

	public function toString():String {
		return getString(0, length);
	}

	public function toHex():String {
		var out = new StringBuf();
		var digits = "0123456789abcdef";
		for (index in 0...length) {
			var value = get(index);
			out.addChar(digits.charCodeAt(value >> 4));
			out.addChar(digits.charCodeAt(value & 15));
		}
		return out.toString();
	}

	public function getData():BytesData {
		__hx_dataExposed = true;
		return b;
	}

	public static function alloc(length:Int):Bytes {
		return new Bytes(length, NativeBytes.allocValues(length));
	}

	@:pure
	public static function ofString(value:String, ?encoding:Encoding):Bytes {
		var view = NativeBytes.viewFromString(value, encoding == RawNative && rawNativeUsesUtf16LE());
		return __hx_fromNativeView(view);
	}

	public static function ofData(data:BytesData):Bytes {
		if (data == null)
			data = NativeBytes.allocValues(0);
		var bytes = new Bytes(data.length, data);
		bytes.__hx_dataExposed = true;
		return bytes;
	}

	public static function ofHex(value:String):Bytes {
		var textLength = value.length;
		if ((textLength & 1) != 0)
			throw "Not a hex string (odd number of digits)";
		var out = alloc(textLength >> 1);
		for (index in 0...out.length) {
			var high = StringTools.fastCodeAt(value, index * 2);
			var low = StringTools.fastCodeAt(value, index * 2 + 1);
			high = (high & 0xF) + (((high & 0x40) >> 6) * 9);
			low = (low & 0xF) + (((low & 0x40) >> 6) * 9);
			out.set(index, ((high << 4) | low) & 0xFF);
		}
		return out;
	}

	public inline static function fastGet(data:BytesData, pos:Int):Int {
		return data[pos];
	}

	@:allow(haxe.crypto.Base64)
	@:allow(haxe.crypto.Md5)
	@:allow(haxe.crypto.Sha1)
	@:allow(haxe.crypto.Sha224)
	@:allow(haxe.crypto.Sha256)
	function __hx_nativeView():ByteView {
		if (!__hx_rawValid || (__hx_dataExposed && !NativeBytes.viewMatchesValues(__hx_raw, b))) {
			__hx_raw = NativeBytes.viewFromValues(b);
			__hx_rawValid = true;
		}
		return __hx_raw;
	}

	@:allow(haxe.crypto.Base64)
	@:allow(haxe.crypto.Md5)
	@:allow(haxe.crypto.Sha1)
	@:allow(haxe.crypto.Sha224)
	@:allow(haxe.crypto.Sha256)
	static function __hx_fromNativeView(view:ByteView):Bytes {
		return new Bytes(NativeBytes.viewLength(view), NativeBytes.valuesFromView(view), view);
	}

	static inline function rawNativeUsesUtf16LE():Bool {
		#if (reflaxe_go_raw_native_mode == "utf16le")
		return true;
		#else
		return false;
		#end
	}
}
