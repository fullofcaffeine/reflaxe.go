package sys.net;

import go.NativeSlice;
import haxe.io.Bytes;
import haxe.io.Eof;
import haxe.io.Error;
import hxrt.net.NativeSocket;

/**
	What
	- Implements Haxe 4.3.7 `sys.net.UdpSocket` over the shared typed socket handle.

	Why
	- The mainstream Haxe stdlib implementation cannot be used unchanged on `haxe.go`
	  because its cross-target fallback throws “not available,” while UDP addressing,
	  broadcast options, and datagram I/O require target support. Bounds and public
	  Address mutation remain Haxe semantics rather than compiler/runtime policy.

	How
	- Replace the inherited empty TCP handle with a UDP handle, validate/copy Bytes
	  in source, and delegate only bind/socket-option/datagram operations to hxrt.
**/
class UdpSocket extends Socket {
	public function new() {
		super();
		replaceHandle(NativeSocket.newUdp());
	}

	override public function bind(host:Host, port:Int):Void {
		if (host == null)
			throw "udp bind requires host";
		NativeSocket.udpBind(handle, host.toString(), port);
	}

	public function setBroadcast(enabled:Bool):Void {
		NativeSocket.udpSetBroadcast(handle, enabled);
	}

	public function sendTo(bytes:Bytes, pos:Int, length:Int, address:Address):Int {
		if (bytes == null || address == null || pos < 0 || length < 0 || pos + length > bytes.length)
			throw Error.OutsideBounds;
		var values = new Array<Int>();
		for (index in 0...length)
			values.push(bytes.get(pos + index));
		var result = NativeSocket.udpSendTo(handle, NativeSlice.fromArray(values), address.host, address.port);
		if (result.status == NativeSocket.IO_BLOCKED)
			throw Error.Blocked;
		if (result.status == NativeSocket.IO_EOF)
			throw new Eof();
		return result.count;
	}

	public function readFrom(bytes:Bytes, pos:Int, length:Int, address:Address):Int {
		if (bytes == null || address == null || pos < 0 || length < 0 || pos + length > bytes.length)
			throw Error.OutsideBounds;
		if (length == 0)
			return 0;
		var result = NativeSocket.udpReadFrom(handle, length);
		if (result.status == NativeSocket.IO_BLOCKED)
			throw Error.Blocked;
		if (result.status == NativeSocket.IO_EOF)
			throw new Eof();
		for (index in 0...result.count)
			bytes.set(pos + index, result.values[index]);
		address.host = result.host;
		address.port = result.port;
		return result.count;
	}
}
