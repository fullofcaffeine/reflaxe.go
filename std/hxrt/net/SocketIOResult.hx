package hxrt.net;

import go.NativeSlice;

/**
	What: Typed byte-progress result for native socket reads and writes.
	Why: EOF and blocked are public Haxe values, not Go runtime types, so a native
	`Dynamic` result or runtime-created generated enum would violate ownership.
	How: Carry values/count/status explicitly and translate status in `_SocketIO.hx`.
**/
@:go.import("hxrt")
@:go.package("hxrt")
@:go.name("SocketIOResult")
extern class SocketIOResult {
	@:go.name("Values")
	public var values:NativeSlice<Int>;

	@:go.name("Count")
	public var count:Int;

	@:go.name("Status")
	public var status:Int;
}
