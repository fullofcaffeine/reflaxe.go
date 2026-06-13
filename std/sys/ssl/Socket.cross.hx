package sys.ssl;

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

	/**
		What
		Opaque handle for Go's TLS SNI certificate table.

		Why
		The actual certificate selector must be installed on Go's `tls.Config`,
		so the state lives in `hxrt` rather than in ordinary Haxe collections.

		How
		`addSNICertificate` creates or extends this handle through a typed hxrt
		helper, and `bind` passes it to the TLS listener helper.
	**/
	var sniConfig:Dynamic;

	public function new() {
		super();
		if (DEFAULT_VERIFY_CERT == true && DEFAULT_CA == null) {
			try {
				DEFAULT_CA = Certificate.loadDefaults();
			} catch (_:haxe.Exception) {}
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
		Register a certificate/key pair for Server Name Indication (SNI).

		Why
		SNI is the hostname a TLS client sends during handshake. Servers use it
		to choose the right certificate when one listener serves multiple names.

		How
		The Haxe callback remains the public API. `haxe.go` stores the callback
		and parsed certificate in an `hxrt` table, then Go's TLS listener asks that
		table for a matching certificate during handshake.
	**/
	public function addSNICertificate(cbServernameMatch:String->Bool, cert:Certificate, key:Key):Void {
		if (cbServernameMatch == null || cert == null || key == null) {
			throw "sys.ssl.Socket.addSNICertificate requires callback, certificate, and key";
		}
		sniConfig = untyped __go__("hxrt.SslSocketAddSNICertificate({0}, {1}, {2}, {3})", sniConfig, cbServernameMatch, cert.handle, key.handle);
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
		untyped __go__("func() int { listener := hxrt.SslSocketListen({1}, {2}, {3}, {4}, {5}); if {0}.listener != nil { _ = {0}.listener.Close() }; {0}.listener = listener; {0}.hxrt__socket_applyListenerDeadline(); return 0 }()",
			this, resolvedHost, port, ownCert == null ? null : ownCert.handle, ownKey == null ? null : ownKey.handle, sniConfig);
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
