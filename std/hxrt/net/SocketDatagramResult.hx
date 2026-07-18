package hxrt.net;

import go.NativeSlice;

/**
	What: Carries UDP byte progress, status, and the native peer address.
	Why: Datagram reads need both payload and address without passing generated Haxe
	objects or an untyped result through the runtime boundary.
	How: `NativeSocket.udpReadFrom` returns this value for staged UdpSocket to copy.
**/
@:go.import("hxrt")
@:go.package("hxrt")
@:go.name("SocketDatagramResult")
extern class SocketDatagramResult {
	@:go.name("Values")
	public var values:NativeSlice<Int>;

	@:go.name("Count")
	public var count:Int;

	@:go.name("Status")
	public var status:Int;

	@:go.name("Host")
	public var host:Int;

	@:go.name("Port")
	public var port:Int;
}
