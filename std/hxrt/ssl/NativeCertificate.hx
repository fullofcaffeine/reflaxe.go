package hxrt.ssl;

import go.NativeSlice;

/**
	What
	Typed bridge for native certificate values that cannot be expressed by the
	portable `sys.ssl.Certificate` source alone.

	Why
	The mainstream Haxe stdlib API returns `Array<String>`, while Go's TLS runtime
	naturally returns a native string slice. Declaring that slice as a Haxe Array
	would hide incompatible identity and growth semantics.

	How
	Expose the Go result as `NativeSlice<String>`. The staged public API performs
	the explicit copy into a portable Array at the boundary.
**/
@:go.import("hxrt")
@:go.package("hxrt")
extern class NativeCertificate {
	@:go.name("SslCertLoadFile")
	public static function loadFile(file:String):CertificateHandle;

	@:go.name("SslCertLoadPath")
	public static function loadPath(path:String):CertificateHandle;

	@:go.name("SslCertFromString")
	public static function fromString(value:String):CertificateHandle;

	@:go.name("SslCertLoadDefaults")
	public static function loadDefaults():CertificateHandle;

	@:go.name("SslCertCommonName")
	public static function commonName(handle:CertificateHandle):String;

	@:go.name("SslCertAltNames")
	public static function altNames(handle:CertificateHandle):NativeSlice<String>;

	@:go.name("SslCertSubject")
	public static function subject(handle:CertificateHandle, field:String):String;

	@:go.name("SslCertIssuer")
	public static function issuer(handle:CertificateHandle, field:String):String;

	@:go.name("SslCertNotBeforeMs")
	public static function notBeforeMs(handle:CertificateHandle):Float;

	@:go.name("SslCertNotAfterMs")
	public static function notAfterMs(handle:CertificateHandle):Float;

	@:go.name("SslCertNext")
	public static function next(handle:CertificateHandle):CertificateHandle;

	@:go.name("SslCertAddPEM")
	public static function addPem(handle:CertificateHandle, value:String):Void;

	@:go.name("SslCertAddDERValues")
	public static function addDer(handle:CertificateHandle, values:NativeSlice<Int>):Void;
}
