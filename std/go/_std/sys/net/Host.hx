package sys.net;

import hxrt.net.NativeSocket;

/**
	What
	- Implements the Haxe 4.3.7 `sys.net.Host` API for the Go target.

	Why
	- The mainstream Haxe stdlib implementation cannot be used unchanged on `haxe.go`
	  because `sys.net.Host` is extern. DNS and OS hostname lookup are native capabilities, but the public
	  host/name/IP model belongs in ordinary staged source rather than the compiler.

	How
	- Store the upstream `host` and network-order IPv4 `ip` fields, delegating only
	  resolution, rendering, reverse lookup, and local hostname access to typed hxrt.
**/
class Host {
	public var host(default, null):String;
	public var ip(default, null):Int;

	public function new(name:String):Void {
		host = name;
		ip = NativeSocket.hostResolve(name);
	}

	public function toString():String {
		return NativeSocket.hostToString(ip);
	}

	public function reverse():String {
		return NativeSocket.hostReverse(ip);
	}

	public static function localhost():String {
		return NativeSocket.hostLocal();
	}

	/**
		What: Builds a public Host from a native network-order IPv4 value.
		Why: Peer and local address calls already resolved the address in Go, so a
		second public DNS policy decision would be misleading.
		How: Render the integer through the typed runtime helper, construct the
		canonical public object, and retain the exact native integer.
	**/
	@:allow(sys.net.Address)
	@:allow(sys.net.Socket)
	private static function fromIPv4(value:Int):Host {
		var result = new Host(NativeSocket.hostToString(value));
		result.ip = value;
		return result;
	}
}
