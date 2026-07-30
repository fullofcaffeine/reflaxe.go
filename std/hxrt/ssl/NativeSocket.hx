package hxrt.ssl;

import hxrt.net.SocketHandle;
import hxrt.net.SocketEndpoint;

/**
	What: Typed TLS composition capabilities over the shared native socket handle.
	Why: TLS connection/listener installation and handshake require Go resources,
	while public configuration and accepted Socket identity belong in staged Haxe.
	How: Map to the footprint-explicit `socket_ssl.go` bridge and pass only typed
	certificate, key, SNI-table, and socket handles.
**/
@:go.import("hxrt")
@:go.package("hxrt")
extern class NativeSocket {
	@:go.name("SslSocketAddSNICertificate")
	public static function addSniCertificate(config:SNIConfigHandle, matcher:String->Bool, cert:CertificateHandle, key:KeyHandle):SNIConfigHandle;

	@:go.name("SslSocketConnect")
	public static function connect(handle:SocketHandle, endpoint:SocketEndpoint, port:Int, verifyCert:Bool, ca:CertificateHandle, serverName:String,
		cert:CertificateHandle, key:KeyHandle):Void;

	/**
		What: Reserves a server endpoint and retains its typed TLS wrapping policy.
		Why: `sys.ssl.Socket.bind` must not start listening or discard the later
		`listen(connections)` backlog.
		How: Bind through `SslSocketBind`; the inherited typed network capability
		performs the actual listen transition.
	**/
	@:go.name("SslSocketBind")
	public static function bind(handle:SocketHandle, host:String, port:Int, cert:CertificateHandle, key:KeyHandle, sni:SNIConfigHandle):Void;

	@:go.name("SslSocketHandshake")
	public static function handshake(handle:SocketHandle):Void;

	@:go.name("SslSocketPeerCertificate")
	public static function peerCertificate(handle:SocketHandle):CertificateHandle;
}
