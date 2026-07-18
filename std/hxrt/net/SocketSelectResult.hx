package hxrt.net;

import go.NativeSlice;

/**
	What: Carries typed readiness indexes returned by native polling.
	Why: Indexes keep public Socket identity and `custom` payload mapping in staged
	Haxe instead of passing generated objects into the runtime package.
	How: Staged `Socket.select` maps each index back to its original Socket instance.
**/
@:go.import("hxrt")
@:go.package("hxrt")
@:go.name("SocketSelectResult")
extern class SocketSelectResult {
	@:go.name("Read")
	public var read:NativeSlice<Int>;

	@:go.name("Write")
	public var write:NativeSlice<Int>;

	@:go.name("Others")
	public var others:NativeSlice<Int>;
}
