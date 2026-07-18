package hxrt.ssl;

import go.NativeSlice;

/**
	What: Exposes native hashing, signing, and verification over integer byte slices.
	Why: Cryptographic implementations require Go libraries while public Bytes
	conversion and algorithm selection remain staged Haxe policy.
	How: Map directly to typed `runtime/hxrt/ssl.go` functions with no raw injection.
**/
@:go.import("hxrt")
@:go.package("hxrt")
extern class NativeDigest {
	@:go.name("SslDigestMakeValues")
	public static function make(data:NativeSlice<Int>, algorithm:String):NativeSlice<Int>;

	@:go.name("SslDigestSignValues")
	public static function sign(data:NativeSlice<Int>, key:KeyHandle, algorithm:String):NativeSlice<Int>;

	@:go.name("SslDigestVerifyValues")
	public static function verify(data:NativeSlice<Int>, signature:NativeSlice<Int>, key:KeyHandle, algorithm:String):Bool;
}
