package haxe.crypto;

import haxe.io.Bytes;
import hxrt.crypto.NativeCrypto;

/**
	What:
	- Implements `Sha256.encode` and `Sha256.make` as staged Haxe source.

	Why:
	- The mainstream Haxe stdlib implementation cannot be used unchanged on `haxe.go` while generated
	  Go lacks Haxe array-growth semantics for indexed writes; the public API still
	  does not justify compiler-owned declarations.

	How:
	- Reuse `Bytes`' opaque cached byte view and delegate only SHA-256 execution to
	  Go's standard runtime implementation.
**/
class Sha256 {
	public static function encode(value:String):String {
		return NativeCrypto.sha256String(value);
	}

	public static function make(value:Bytes):Bytes {
		return Bytes.__hx_fromNativeView(NativeCrypto.sha256Values(value.__hx_nativeView()));
	}
}
