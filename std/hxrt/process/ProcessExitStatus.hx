package hxrt.process;

/**
	What
	- Typed carrier for a polled or awaited child-process exit status.

	Why
	- The public API returns `Null<Int>`, but a native `Dynamic`/`Any` result would
	  conflate a running child with a completed child whose exit code is zero.

	How
	- Map the explicit `Available` and `Code` fields of `hxrt.ProcessExitStatus`;
	  staged Haxe performs the final `Null<Int>` conversion.
**/
@:go.import("hxrt")
@:go.package("hxrt")
@:go.name("ProcessExitStatus")
extern class ProcessExitStatus {
	@:go.name("Code")
	public var code:Int;

	@:go.name("Available")
	public var available:Bool;
}
