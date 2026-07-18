package haxe.crypto;

import haxe.io.Bytes;
import hxrt.crypto.NativeCrypto;

/**
	What:
	- Implements `Sha1.encode` and `Sha1.make` as staged Haxe source.

	Why:
	- The mainstream Haxe stdlib implementation cannot be used unchanged on `haxe.go` while generated
	  Go lacks Haxe array-growth semantics for indexed writes; the public API still
	  does not justify compiler-owned declarations.

	How:
	- Reuse `Bytes`' opaque cached byte view and delegate only SHA-1 execution to
	  Go's standard runtime implementation.
**/
class Sha1 {
	public static function encode(value:String):String {
		return NativeCrypto.sha1String(value);
	}

	public static function make(value:Bytes):Bytes {
		return Bytes.__hx_fromNativeView(NativeCrypto.sha1Values(value.__hx_nativeView()));
	}
}
