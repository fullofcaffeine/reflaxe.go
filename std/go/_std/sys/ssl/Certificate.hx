package sys.ssl;

import go.NativeSlice;
import haxe.io.Bytes;
import hxrt.ssl.CertificateHandle;
import hxrt.ssl.NativeCertificate;

/**
	What
	- Implements the Haxe 4.3.7 `sys.ssl.Certificate` API as staged source.

	Why
	- The mainstream Haxe stdlib implementation cannot be used unchanged on `haxe.go`
	  because `sys.ssl.Certificate` is extern. Certificate parsing and trust stores are native, but public chain
	  traversal, Array conversion, and Date construction belong in Haxe source.

	How
	- Retain one opaque typed certificate handle and delegate native capabilities
	  through `NativeCertificate`; DER and alt-name slices are copied explicitly.
**/
class Certificate {
	@:noCompletion
	@:dox(hide)
	public var handle(default, null):CertificateHandle;

	@:allow(sys.ssl.Socket)
	private function new(handle:CertificateHandle) {
		this.handle = handle;
	}

	public static function loadFile(file:String):Certificate {
		return new Certificate(NativeCertificate.loadFile(file));
	}

	public static function loadPath(path:String):Certificate {
		return new Certificate(NativeCertificate.loadPath(path));
	}

	public static function fromString(value:String):Certificate {
		return new Certificate(NativeCertificate.fromString(value));
	}

	public static function loadDefaults():Certificate {
		return new Certificate(NativeCertificate.loadDefaults());
	}

	public var commonName(get, null):Null<String>;
	public var altNames(get, null):Array<String>;
	public var notBefore(get, null):Date;
	public var notAfter(get, null):Date;

	public function subject(field:String):Null<String> {
		return NativeCertificate.subject(handle, field);
	}

	public function issuer(field:String):Null<String> {
		return NativeCertificate.issuer(handle, field);
	}

	public function next():Null<Certificate> {
		var nextHandle = NativeCertificate.next(handle);
		return nextHandle == null ? null : new Certificate(nextHandle);
	}

	public function add(pem:String):Void {
		NativeCertificate.addPem(handle, pem);
	}

	public function addDER(der:Bytes):Void {
		var values = new Array<Int>();
		for (index in 0...der.length)
			values.push(der.get(index));
		NativeCertificate.addDer(handle, NativeSlice.fromArray(values));
	}

	private function get_commonName():Null<String> {
		return NativeCertificate.commonName(handle);
	}

	private function get_altNames():Array<String> {
		return NativeCertificate.altNames(handle).toArray();
	}

	private function get_notBefore():Date {
		return Date.fromTime(NativeCertificate.notBeforeMs(handle));
	}

	private function get_notAfter():Date {
		return Date.fromTime(NativeCertificate.notAfterMs(handle));
	}
}
