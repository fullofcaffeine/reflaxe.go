package hxrt.net;

import go.NativeSlice;

/**
	What
	- Typed bridge to DNS, TCP, UDP, stream I/O, deadline, and readiness capabilities.

	Why
	- These operations require Go's networking runtime, but public socket objects,
	  Haxe exceptions, bounds checks, and object identity belong in staged std.

	How
	- Map one-for-one to `runtime/hxrt/socket.go`. Native slices make byte and handle
	  representation changes explicit; typed result carriers preserve EOF/blocked
	  states until staged Haxe translates them.
**/
@:go.import("hxrt")
@:go.package("hxrt")
extern class NativeSocket {
	@:go.name("SocketIOReady")
	public static var IO_READY(default, null):Int;

	@:go.name("SocketIOEOF")
	public static var IO_EOF(default, null):Int;

	@:go.name("SocketIOBlocked")
	public static var IO_BLOCKED(default, null):Int;

	@:go.name("SocketReadEOF")
	public static var READ_EOF(default, null):Int;

	@:go.name("SocketReadBlocked")
	public static var READ_BLOCKED(default, null):Int;

	@:go.name("HostResolve")
	public static function hostResolve(name:String):Int;

	@:go.name("HostToString")
	public static function hostToString(value:Int):String;

	@:go.name("HostReverse")
	public static function hostReverse(value:Int):String;

	@:go.name("HostLocal")
	public static function hostLocal():String;

	@:go.name("SocketNewTCP")
	public static function newTcp():SocketHandle;

	@:go.name("SocketNewUDP")
	public static function newUdp():SocketHandle;

	@:go.name("SocketConnectTCP")
	public static function connectTcp(handle:SocketHandle, host:String, port:Int):Void;

	@:go.name("SocketBindTCP")
	public static function bindTcp(handle:SocketHandle, host:String, port:Int):Void;

	@:go.name("SocketListen")
	public static function listen(handle:SocketHandle, connections:Int):Void;

	@:go.name("SocketAccept")
	public static function accept(handle:SocketHandle):SocketAcceptResult;

	@:go.name("SocketClose")
	public static function close(handle:SocketHandle):Void;

	@:go.name("SocketShutdown")
	public static function shutdown(handle:SocketHandle, read:Bool, write:Bool):Void;

	@:go.name("SocketPeer")
	public static function peer(handle:SocketHandle):SocketAddress;

	@:go.name("SocketHost")
	public static function host(handle:SocketHandle):SocketAddress;

	@:go.name("SocketSetTimeout")
	public static function setTimeout(handle:SocketHandle, timeout:Float):Void;

	@:go.name("SocketWaitForRead")
	public static function waitForRead(handle:SocketHandle):Void;

	@:go.name("SocketSetBlocking")
	public static function setBlocking(handle:SocketHandle, blocking:Bool):Void;

	@:go.name("SocketSetFastSend")
	public static function setFastSend(handle:SocketHandle, fastSend:Bool):Void;

	@:go.name("SocketReadByteValue")
	public static function readByte(handle:SocketHandle):Int;

	@:go.name("SocketReadValues")
	public static function readValues(handle:SocketHandle, length:Int):SocketIOResult;

	@:go.name("SocketWriteValues")
	public static function writeValues(handle:SocketHandle, values:NativeSlice<Int>):SocketIOResult;

	@:go.name("SocketFlush")
	public static function flush(handle:SocketHandle):Void;

	@:go.name("SocketSelect")
	public static function select(read:NativeSlice<SocketHandle>, write:NativeSlice<SocketHandle>, others:NativeSlice<SocketHandle>, timeout:Float,
		hasTimeout:Bool):SocketSelectResult;

	@:go.name("SocketUdpBind")
	public static function udpBind(handle:SocketHandle, host:String, port:Int):Void;

	@:go.name("SocketUdpSetBroadcast")
	public static function udpSetBroadcast(handle:SocketHandle, enabled:Bool):Void;

	@:go.name("SocketUdpSendTo")
	public static function udpSendTo(handle:SocketHandle, values:NativeSlice<Int>, host:Int, port:Int):SocketIOResult;

	@:go.name("SocketUdpReadFrom")
	public static function udpReadFrom(handle:SocketHandle, length:Int):SocketDatagramResult;
}
