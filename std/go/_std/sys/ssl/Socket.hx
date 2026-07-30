package sys.ssl;

import haxe.io.Eof;
import haxe.io.Error;
import hxrt.net.NativeSocket as NativeNetSocket;
import hxrt.ssl.NativeSocket as NativeSslSocket;
import hxrt.ssl.SNIConfigHandle;
import sys.net.Host;

/**
	What
	- Implements the Haxe 4.3.7 `sys.ssl.Socket` API as staged source.

	Why
	- The mainstream Haxe stdlib implementation cannot be used unchanged on `haxe.go`
	  because `sys.ssl.Socket` is extern. TLS policy belongs beside the public
	  Haxe API, while connection/listener installation, SNI selection, handshakes,
	  and peer certificates require native resources.

	How
	- Extend the source-owned `sys.net.Socket`, retain typed certificate/key/SNI
	  configuration, and compose TLS through the shared opaque socket handle.
**/
class Socket extends sys.net.Socket {
	public static var DEFAULT_VERIFY_CERT:Null<Bool> = true;
	public static var DEFAULT_CA:Null<Certificate>;

	public var verifyCert:Null<Bool>;

	private var caCert:Null<Certificate>;
	private var hostname:Null<String>;
	private var ownCert:Null<Certificate>;
	private var ownKey:Null<Key>;
	private var sniConfig:SNIConfigHandle;

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
		NativeSslSocket.handshake(handle);
	}

	/**
		What: Applies the inherited socket shutdown contract to a TLS connection.
		Why: TLS write shutdown must send a protocol `close_notify`, while a read-only
		TLS shutdown has no safe equivalent and must not silently look successful.
		How: Delegate to the shared typed handle: write-only preserves peer reads,
		read-only raises a deterministic unsupported error, and both sides close.
	**/
	override public function shutdown(read:Bool, write:Bool):Void {
		super.shutdown(read, write);
	}

	/**
		What: Applies the inherited fast-send preference to TLS's TCP transport.
		Why: TLS wraps the TCP connection, so a direct TCP-only native check would
		accept this call while leaving `TCP_NODELAY` unchanged.
		How: The typed runtime follows the TLS connection's `NetConn` boundary and
		updates the underlying TCP socket; a non-TCP transport reports unsupported.
	**/
	override public function setFastSend(fastSend:Bool):Void {
		super.setFastSend(fastSend);
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

	/** Register one typed Haxe SNI matcher and native certificate/key pair. **/
	public function addSNICertificate(matcher:String->Bool, cert:Certificate, key:Key):Void {
		if (matcher == null || cert == null || key == null)
			throw "sys.ssl.Socket.addSNICertificate requires callback, certificate, and key";
		sniConfig = NativeSslSocket.addSniCertificate(sniConfig, matcher, cert.handle, key.handle);
	}

	/**
		What: Connects TLS to the resolved IPv4 address while retaining the source hostname.
		Why: `Host.toString()` is the routing address, not necessarily the certificate
		identity or SNI name.
		How: Pass both values through a typed endpoint; an explicit `setHostname` value
		still overrides the retained source hostname.
	**/
	override public function connect(host:Host, port:Int):Void {
		if (host == null)
			throw "socket connect requires host";
		var endpoint = NativeNetSocket.endpoint(host.toString(), host.host);
		NativeSslSocket.connect(handle, endpoint, port, verifyCert != false, caCert == null ? null : caCert.handle, hostname,
			ownCert == null ? null : ownCert.handle, ownKey == null ? null : ownKey.handle);
	}

	override public function bind(host:Host, port:Int):Void {
		if (host == null)
			throw "socket bind requires host";
		NativeSslSocket.bind(handle, host.toString(), port, ownCert == null ? null : ownCert.handle, ownKey == null ? null : ownKey.handle, sniConfig);
	}

	/**
		What: Accepts and handshakes one TLS connection while preserving SSL object identity.
		Why: The inherited signature returns `sys.net.Socket`, but callers must still receive
		a dynamically typed `sys.ssl.Socket`, matching established Haxe target behavior.
		How: Accept the shared typed handle directly, install it on a new SSL instance,
		then perform the server handshake before returning its embedded base view.
	**/
	override public function accept():sys.net.Socket {
		var result = NativeNetSocket.accept(handle);
		if (result.status == NativeNetSocket.IO_BLOCKED)
			throw Error.Blocked;
		if (result.status == NativeNetSocket.IO_EOF || result.handle == null)
			throw new Eof();
		var accepted = new Socket();
		accepted.replaceHandle(result.handle);
		NativeSslSocket.handshake(accepted.handle);
		return accepted;
	}

	public function peerCertificate():Certificate {
		var certificateHandle = NativeSslSocket.peerCertificate(handle);
		return certificateHandle == null ? null : new Certificate(certificateHandle);
	}
}
