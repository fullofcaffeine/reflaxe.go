package sys.ssl;

import haxe.io.Bytes;

/**
	What
	Direct `sys.ssl.Digest` support for `haxe.go`.

	Why
	- The Haxe sys stdlib exposes digest and signature helpers under
	  `sys.ssl.Digest`.
	- Hashing/signing is target-runtime behavior, but the public Haxe API should
	  remain portable.

	How
	- Keep the public Haxe static API here in staged std.
	- Delegate hashing/signing/verifying to `hxrt` helpers and bridge `Bytes`
	  through the existing raw-byte conversion helpers.
**/
@:goAllowRaw
class Digest {
	public static function make(data:Bytes, alg:DigestAlgorithm):Bytes {
		var algName:String = cast alg;
		return untyped __go__("hxrt_rawToHaxeBytes(hxrt.SslDigestMake(hxrt_haxeBytesToRaw({0}), {1}))", data, algName);
	}

	public static function sign(data:Bytes, privKey:Key, alg:DigestAlgorithm):Bytes {
		var algName:String = cast alg;
		return untyped __go__("hxrt_rawToHaxeBytes(hxrt.SslDigestSign(hxrt_haxeBytesToRaw({0}), {1}.handle, {2}))", data, privKey, algName);
	}

	public static function verify(data:Bytes, signature:Bytes, pubKey:Key, alg:DigestAlgorithm):Bool {
		var algName:String = cast alg;
		return untyped __go__("hxrt.SslDigestVerify(hxrt_haxeBytesToRaw({0}), hxrt_haxeBytesToRaw({1}), {2}.handle, {3})", data, signature, pubKey, algName);
	}
}
