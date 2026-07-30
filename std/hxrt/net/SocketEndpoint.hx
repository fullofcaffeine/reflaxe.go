package hxrt.net;

/**
	What
	- Opaque typed carrier for a resolved network address and its logical hostname.

	Why
	- `sys.net.Host.toString()` returns the numeric IPv4 routing address, while
	  TLS verification and SNI must retain the original hostname identity.

	How
	- Staged socket code constructs the carrier through `NativeSocket.endpoint`;
	  only `hxrt` reads its native fields.
**/
@:go.import("hxrt")
@:go.package("hxrt")
@:go.name("SocketEndpoint")
extern class SocketEndpoint {}
