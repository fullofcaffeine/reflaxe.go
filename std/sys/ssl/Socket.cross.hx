package sys.ssl;

import haxe.exceptions.NotImplementedException;
import sys.net.Host;

/**
	What
	Direct `sys.ssl.Socket` support for `haxe.go`.

	Why
	- `sys.ssl.Socket` is the public TLS socket surface used by direct sys code
	  and by higher-level wrappers such as `sys.Http`.
	- The existing compiler-owned `sys.net.Socket` carrier already owns socket
	  lifecycle, deadlines, buffering, and select behavior on Go. Reusing that
	  carrier avoids growing a second socket stack just to add TLS.

	How
	- Keep the Haxe-facing TLS API here as a staged class that extends the
	  existing `sys.net.Socket` carrier.
	- Store only TLS-specific configuration on the staged class.
	- Delegate TLS dial/listen/handshake/peer-certificate behavior to `hxrt`
	  helpers while reusing the inherited socket fields and methods for the rest
	  of the contract.
**/
@:goAllowRaw
class Socket extends sys.net.Socket {
	public static var DEFAULT_VERIFY_CERT:Null<Bool> = true;
	public static var DEFAULT_CA:Null<Certificate>;

	public var verifyCert:Null<Bool>;

	var caCert:Null<Certificate>;
	var hostname:Null<String>;
	var ownCert:Null<Certificate>;
	var ownKey:Null<Key>;

	public function new() {
		super();
		if (DEFAULT_VERIFY_CERT == true && DEFAULT_CA == null) {
			try {
				DEFAULT_CA = Certificate.loadDefaults();
			} catch (_:Dynamic) {}
		}
		verifyCert = DEFAULT_VERIFY_CERT;
		caCert = DEFAULT_CA;
	}

	public function handshake():Void {
		untyped __go__("func() int { hxrt.SslSocketHandshake({0}.hxrt__socket_conn()); return 0 }()", this);
	}

	public function setCA(cert:Certificate):Void {
		caCert = cert;
	}

	public function setHostname(name:String):Void {
		hostname = name;
	}

	public function setCertificate(cert:Certificate, key:Key):Void {
		ownCert = cert;
		ownKey = key;
	}

	/**
		What
		Register additional SNI certificate selection logic.

		Why
		- The baseline `haxe.go` TLS socket support currently covers one active
		  certificate/key pair for direct portable usage.
		- Multi-certificate SNI callback dispatch needs a second slice so the
		  callback bridge stays typed and testable instead of landing as an ad hoc
		  raw closure.

		How
		- Fail fast with an explicit `NotImplementedException` until that follow-up
		  lands.
	**/
	public function addSNICertificate(cbServernameMatch:String->Bool, cert:Certificate, key:Key):Void {
		throw new NotImplementedException("sys.ssl.Socket.addSNICertificate is not implemented on haxe.go yet");
	}

	override public function connect(host:Host, port:Int):Void {
		if (host == null) {
			throw "socket connect requires host";
		}
		var resolvedHost = host.toString();
		if (resolvedHost == null) {
			throw "socket connect requires host";
		}
		untyped __go__("func() int { conn := hxrt.SslSocketConnect({1}, {2}, {3}, {4}, {5}, {6}, {7}); {0}.hxrt__socket_setConn(conn); return 0 }()", this,
			resolvedHost, port, verifyCert != false, caCert == null ? null : caCert.handle, hostname, ownCert == null ? null : ownCert.handle,
			ownKey == null ? null : ownKey.handle);
	}

	override public function bind(host:Host, port:Int):Void {
		if (host == null) {
			throw "socket bind requires host";
		}
		var resolvedHost = host.toString();
		if (resolvedHost == null) {
			throw "socket bind requires host";
		}
		untyped __go__("func() int { listener := hxrt.SslSocketListen({1}, {2}, {3}, {4}); if {0}.listener != nil { _ = {0}.listener.Close() }; {0}.listener = listener; {0}.hxrt__socket_applyListenerDeadline(); return 0 }()",
			this, resolvedHost, port, ownCert == null ? null : ownCert.handle, ownKey == null ? null : ownKey.handle);
	}

	override public function accept():sys.net.Socket {
		var accepted = super.accept();
		untyped __go__("func() int { hxrt.SslSocketHandshake({0}.hxrt__socket_conn()); return 0 }()", accepted);
		return accepted;
	}

	public function peerCertificate():Certificate {
		var handle = untyped __go__("hxrt.SslSocketPeerCertificate({0}.hxrt__socket_conn())", this);
		return handle == null ? null : new Certificate(handle);
	}
}
