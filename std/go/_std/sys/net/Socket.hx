package sys.net;

import go.NativeSlice;
import haxe.io.Eof;
import haxe.io.Error;
import hxrt.net.NativeSocket;
import hxrt.net.SocketAddress;
import hxrt.net.SocketHandle;
import sys.net._SocketIO.SocketInput;
import sys.net._SocketIO.SocketOutput;

/**
	What
	- Implements the complete Haxe 4.3.7 `sys.net.Socket` API for the Go target.

	Why
	- The mainstream Haxe stdlib implementation cannot be used unchanged on `haxe.go`
	  because `sys.net.Socket` is extern. Socket objects, stream policy, result identity, and Haxe exceptions
	  are library semantics; only OS resources/deadlines/readiness require Go.

	How
	- Retain one opaque typed `SocketHandle`, expose ordinary source-owned Input and
	  Output wrappers, and delegate native capabilities through `NativeSocket`.
**/
class Socket {
	public var input(default, null):haxe.io.Input;
	public var output(default, null):haxe.io.Output;

	/**
		What: Stores the application payload returned with this exact Socket after `select`.
		Why: The upstream public API intentionally declares `custom:Dynamic`; narrowing it would break source compatibility.
		How: Keep the unavoidable untyped value localized to this public compatibility field; no native capability consumes it.
	**/
	public var custom:Dynamic;

	@:allow(sys.net.UdpSocket)
	@:allow(sys.ssl.Socket)
	private var handle:SocketHandle;

	public function new():Void {
		replaceHandle(NativeSocket.newTcp());
	}

	public function close():Void {
		NativeSocket.close(handle);
	}

	public function read():String {
		return input.readAll().toString();
	}

	public function write(content:String):Void {
		output.writeString(content);
	}

	public function connect(host:Host, port:Int):Void {
		if (host == null)
			throw "socket connect requires host";
		NativeSocket.connectTcp(handle, host.toString(), port);
	}

	public function listen(connections:Int):Void {
		NativeSocket.listen(handle, connections);
	}

	public function shutdown(read:Bool, write:Bool):Void {
		NativeSocket.shutdown(handle, read, write);
	}

	public function bind(host:Host, port:Int):Void {
		if (host == null)
			throw "socket bind requires host";
		NativeSocket.bindTcp(handle, host.toString(), port);
	}

	public function accept():Socket {
		var result = NativeSocket.accept(handle);
		if (result.status == NativeSocket.IO_BLOCKED)
			throw Error.Blocked;
		if (result.status == NativeSocket.IO_EOF || result.handle == null)
			throw new Eof();
		var accepted = new Socket();
		accepted.replaceHandle(result.handle);
		return accepted;
	}

	public function peer():{host:Host, port:Int} {
		return publicAddress(NativeSocket.peer(handle));
	}

	public function host():{host:Host, port:Int} {
		return publicAddress(NativeSocket.host(handle));
	}

	public function setTimeout(timeout:Float):Void {
		NativeSocket.setTimeout(handle, timeout);
	}

	public function waitForRead():Void {
		NativeSocket.waitForRead(handle);
	}

	public function setBlocking(blocking:Bool):Void {
		NativeSocket.setBlocking(handle, blocking);
	}

	public function setFastSend(fastSend:Bool):Void {
		NativeSocket.setFastSend(handle, fastSend);
	}

	public static function select(read:Array<Socket>, write:Array<Socket>, others:Array<Socket>,
			?timeout:Float):{read:Array<Socket>, write:Array<Socket>, others:Array<Socket>} {
		if (read == null)
			read = [];
		if (write == null)
			write = [];
		if (others == null)
			others = [];
		var readHandles = NativeSlice.fromArray([for (socket in read) socket.handle]);
		var writeHandles = NativeSlice.fromArray([for (socket in write) socket.handle]);
		var otherHandles = NativeSlice.fromArray([for (socket in others) socket.handle]);
		var result = NativeSocket.select(readHandles, writeHandles, otherHandles, timeout == null ? 0.0 : timeout, timeout != null);
		return {
			read: pick(read, result.read),
			write: pick(write, result.write),
			others: pick(others, result.others)
		};
	}

	/**
		What: Replaces the opaque resource and rebuilds both stream wrappers.
		Why: Accepted sockets and the UDP subclass must preserve one shared handle
		across the public Socket object, Input, and Output without leaking the empty
		handle allocated by the ordinary constructor.
		How: Close a distinct prior handle, install the typed replacement, then make
		fresh source-owned wrappers that retain that same replacement.
	**/
	@:allow(sys.net.UdpSocket)
	@:allow(sys.ssl.Socket)
	private function replaceHandle(next:SocketHandle):Void {
		if (handle != null && handle != next)
			NativeSocket.close(handle);
		handle = next;
		input = new SocketInput(next);
		output = new SocketOutput(next);
	}

	/**
		What: Converts a typed native address carrier into the public anonymous shape.
		Why: `hxrt` must not construct generated Haxe objects or know their layout.
		How: Build the Host in staged source and copy the native port.
	**/
	private static function publicAddress(address:SocketAddress):{host:Host, port:Int} {
		if (address == null)
			return {host: Host.fromIPv4(0), port: 0};
		return {host: Host.fromIPv4(address.host), port: address.port};
	}

	/**
		What: Maps native readiness indexes back to their exact public Socket objects.
		Why: Object identity and each Socket's `custom` payload are source semantics.
		How: Validate each native index and select from the original caller array.
	**/
	private static function pick(source:Array<Socket>, indexes:NativeSlice<Int>):Array<Socket> {
		var selected = new Array<Socket>();
		for (index in 0...indexes.length) {
			var sourceIndex = indexes[index];
			if (sourceIndex >= 0 && sourceIndex < source.length)
				selected.push(source[sourceIndex]);
		}
		return selected;
	}
}
