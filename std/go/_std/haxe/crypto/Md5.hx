package haxe.crypto;

import haxe.io.Bytes;
import hxrt.crypto.NativeCrypto;

/**
	What:
	- Implements `Md5.encode` and `Md5.make` as staged Haxe source.

	Why:
	- The mainstream Haxe stdlib implementation cannot be used unchanged on `haxe.go` while generated
	  Go lacks Haxe array-growth semantics for indexed writes; the public API still
	  does not justify compiler-owned declarations.

	How:
	- Reuse `Bytes`' opaque cached byte view and delegate only MD5 execution to
	  Go's standard runtime implementation.
**/
class Md5 {
	public static function encode(value:String):String {
		return NativeCrypto.md5String(value);
	}

	public static function make(value:Bytes):Bytes {
		return Bytes.__hx_fromNativeView(NativeCrypto.md5Values(value.__hx_nativeView()));
	}
}
