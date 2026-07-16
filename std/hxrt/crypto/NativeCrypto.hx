package hxrt.crypto;

/**
	What:
	- Typed access to Go's standard Base64 and digest implementations for the staged
	  `haxe.crypto` overrides.

	Why:
	- The mainstream Haxe algorithms rely on arrays growing through indexed writes.
	  Generated Go uses fixed-length slices for indexed storage, so those algorithms
	  cannot be used unchanged until that general representation gap is closed.
	- Codec and digest execution is runtime work, but public padding defaults, Haxe
	  `Bytes` conversion, and API behavior still belong in staged Haxe source.

	How:
	- Pass only `Array<Int>` byte values, strings, and a URL-safe selector to
	  `runtime/hxrt/crypto.go`; no generated `haxe.io.Bytes` layout crosses packages.
**/
@:go.import("hxrt")
@:go.package("hxrt")
extern class NativeCrypto {
	@:go.name("CryptoBase64Encode")
	public static function base64Encode(values:Array<Int>, urlSafe:Bool):String;

	@:go.name("CryptoBase64Decode")
	public static function base64Decode(value:String, urlSafe:Bool):Array<Int>;

	@:go.name("CryptoMd5String")
	public static function md5String(value:String):String;

	@:go.name("CryptoMd5Values")
	public static function md5Values(values:Array<Int>):Array<Int>;

	@:go.name("CryptoSha1String")
	public static function sha1String(value:String):String;

	@:go.name("CryptoSha1Values")
	public static function sha1Values(values:Array<Int>):Array<Int>;

	@:go.name("CryptoSha224String")
	public static function sha224String(value:String):String;

	@:go.name("CryptoSha224Values")
	public static function sha224Values(values:Array<Int>):Array<Int>;

	@:go.name("CryptoSha256String")
	public static function sha256String(value:String):String;

	@:go.name("CryptoSha256Values")
	public static function sha256Values(values:Array<Int>):Array<Int>;
}
