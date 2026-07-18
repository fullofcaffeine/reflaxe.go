package haxe.crypto;

import haxe.io.Bytes;
import hxrt.crypto.NativeCrypto;

/**
	What:
	- Implements the complete Haxe `haxe.crypto.Base64` API for the Go target.

	Why:
	- The mainstream Haxe stdlib implementation cannot be used unchanged on `haxe.go` because its
	  `BaseCode` dependency grows arrays through indexed assignment, while generated
	  Go slices require their length to be allocated before an indexed write.
	- Base64 defaults and padding are Haxe library behavior and must not live in a
	  compiler-generated Go helper.

	How:
	- Keep standard versus URL-safe alphabets and padding defaults here; delegate
	  only raw unpadded codec execution through the opaque cached byte view.
**/
class Base64 {
	public static var CHARS(default, null) = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/";
	public static var BYTES(default, null) = Bytes.ofString(CHARS);

	public static var URL_CHARS(default, null) = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_";
	public static var URL_BYTES(default, null) = Bytes.ofString(URL_CHARS);

	public static function encode(bytes:Bytes, complement:Bool = true):String {
		return addPadding(NativeCrypto.base64Encode(bytes.__hx_nativeView(), false), bytes.length, complement);
	}

	public static function decode(value:String, complement:Bool = true):Bytes {
		return Bytes.__hx_fromNativeView(NativeCrypto.base64Decode(removePadding(value, complement), false));
	}

	public static function urlEncode(bytes:Bytes, complement:Bool = false):String {
		return addPadding(NativeCrypto.base64Encode(bytes.__hx_nativeView(), true), bytes.length, complement);
	}

	public static function urlDecode(value:String, complement:Bool = false):Bytes {
		return Bytes.__hx_fromNativeView(NativeCrypto.base64Decode(removePadding(value, complement), true));
	}

	static function addPadding(value:String, byteLength:Int, complement:Bool):String {
		if (!complement)
			return value;
		switch (byteLength % 3) {
			case 1:
				return value + "==";
			case 2:
				return value + "=";
			default:
				return value;
		}
	}

	static function removePadding(value:String, complement:Bool):String {
		if (complement)
			while (value.length > 0 && value.charCodeAt(value.length - 1) == "=".code)
				value = value.substr(0, -1);
		return value;
	}
}
