package sys.ssl;

import haxe.io.Bytes;

private typedef CertificateHandle = Dynamic;

/**
	What
	Direct `sys.ssl.Certificate` support for `haxe.go`.

	Why
	- TLS and certificate inspection are part of the Haxe sys stdlib surface.
	- The public API is ordinary Haxe code, but the actual certificate parsing and
	  platform trust-root handling are Go runtime concerns.

	How
	- Keep the Haxe-facing API in staged std.
	- Store the native handle behind a hidden localized `Dynamic` field and route
	  all heavy lifting through `hxrt` helpers.
**/
@:goAllowRaw
class Certificate {
	@:noCompletion
	@:dox(hide)
	public var handle(default, null):CertificateHandle;

	private function new(handle:CertificateHandle) {
		this.handle = handle;
	}

	public static function loadFile(file:String):Certificate {
		return new Certificate(untyped __go__("hxrt.SslCertLoadFile({0})", file));
	}

	public static function loadPath(path:String):Certificate {
		return new Certificate(untyped __go__("hxrt.SslCertLoadPath({0})", path));
	}

	public static function fromString(str:String):Certificate {
		return new Certificate(untyped __go__("hxrt.SslCertFromString({0})", str));
	}

	public static function loadDefaults():Certificate {
		return new Certificate(untyped __go__("hxrt.SslCertLoadDefaults()"));
	}

	public var commonName(get, null):Null<String>;
	public var altNames(get, null):Array<String>;
	public var notBefore(get, null):Date;
	public var notAfter(get, null):Date;

	public function subject(field:String):Null<String> {
		return untyped __go__("hxrt.SslCertSubject({0}, {1})", handle, field);
	}

	public function issuer(field:String):Null<String> {
		return untyped __go__("hxrt.SslCertIssuer({0}, {1})", handle, field);
	}

	public function next():Null<Certificate> {
		var nextHandle:CertificateHandle = untyped __go__("hxrt.SslCertNext({0})", handle);
		return nextHandle == null ? null : new Certificate(nextHandle);
	}

	public function add(pem:String):Void {
		untyped __go__("func() int { hxrt.SslCertAddPEM({0}, {1}); return 0 }()", handle, pem);
	}

	public function addDER(der:Bytes):Void {
		untyped __go__("func() int { hxrt.SslCertAddDER({0}, hxrt_haxeBytesToRaw({1})); return 0 }()", handle, der);
	}

	function get_commonName():Null<String> {
		return untyped __go__("hxrt.SslCertCommonName({0})", handle);
	}

	function get_altNames():Array<String> {
		return untyped __go__("hxrt.SslCertAltNames({0})", handle);
	}

	function get_notBefore():Date {
		return Date.fromTime(untyped __go__("hxrt.SslCertNotBeforeMs({0})", handle));
	}

	function get_notAfter():Date {
		return Date.fromTime(untyped __go__("hxrt.SslCertNotAfterMs({0})", handle));
	}
}
