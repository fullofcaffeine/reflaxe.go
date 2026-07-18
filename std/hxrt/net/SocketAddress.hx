package hxrt.net;

/**
	What: Carries one native IPv4 address and port.
	Why: The runtime cannot construct source-owned `sys.net.Address` or `Host` objects.
	How: Typed hxrt calls return this value and staged source performs the public conversion.
**/
@:go.import("hxrt")
@:go.package("hxrt")
@:go.name("SocketAddress")
extern class SocketAddress {
	@:go.name("Host")
	public var host:Int;

	@:go.name("Port")
	public var port:Int;
}
