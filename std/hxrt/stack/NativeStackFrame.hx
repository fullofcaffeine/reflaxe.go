package hxrt.stack;

/**
	What
	A typed extern for one Go runtime stack frame captured by `hxrt`.

	Why
	Native stack frames are useful diagnostics, but they are target-sensitive and
	must stay behind the explicit `reflaxe_go_native_stack_trace` opt-in.

	How
	The fields map directly to `hxrt.StackFrame` fields. Staged `haxe.CallStack`
	code converts these target-owned frames into portable `StackItem` values only
	for display/debugging.
**/
@:go.import("hxrt")
@:go.package("hxrt")
@:go.name("StackFrame")
extern class NativeStackFrame {
	@:go.name("Function")
	public var functionName:String;

	@:go.name("File")
	public var file:String;

	@:go.name("Line")
	public var line:Int;
}
