package sys.net;

/**
	What
	Direct `sys.net.Address` support for `haxe.go`.

	Why
	The mainstream Haxe stdlib implementation cannot be used unchanged on `haxe.go`.
	- `sys.net.UdpSocket` and related socket APIs address peers via `{host, port}`.
	- The upstream module is part of the Haxe stdlib surface, but `haxe.go`
	  emits only repo-owned staged overrides for direct stdlib modules.

	How
	- Keep the upstream field layout and helper methods unchanged.
	- `getHost()` rebuilds a `Host` instance from the stored IPv4 integer without
	  requiring additional DNS resolution.
**/
class Address {
	public var host:Int;
	public var port:Int;

	static inline function octet(value:Int, shift:Int):String {
		return Std.string((value >>> shift) & 0xFF);
	}

	public function new() {
		host = 0;
		port = 0;
	}

	public function getHost():Host {
		var h = new Host(octet(host, 24) + "." + octet(host, 16) + "." + octet(host, 8) + "." + octet(host, 0));
		untyped h.ip = host;
		return h;
	}

	public function compare(a:Address):Int {
		var dh = a.host - host;
		if (dh != 0) {
			return dh;
		}
		var dp = a.port - port;
		if (dp != 0) {
			return dp;
		}
		return 0;
	}

	public function clone():Address {
		var c = new Address();
		c.host = host;
		c.port = port;
		return c;
	}
}
