package hxrt.net;

/**
	What: Carries an accepted socket handle and its readiness status.
	Why: Timeout and nonblocking outcomes must remain explicit without returning
	`Dynamic` or constructing public Haxe Socket objects inside the runtime.
	How: `NativeSocket.accept` returns this native carrier for staged source to translate.
**/
@:go.import("hxrt")
@:go.package("hxrt")
@:go.name("SocketAcceptResult")
extern class SocketAcceptResult {
	@:go.name("Handle")
	public var handle:SocketHandle;

	@:go.name("Status")
	public var status:Int;
}
