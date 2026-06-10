package hxrt.stack;

/**
	What
	Typed bridge to optional Go-native stack capture in `hxrt`.

	Why
	Stack capture needs Go runtime APIs, so the implementation belongs in `hxrt`,
	not in app code or compiler-owned raw emitters.

	How
	`capture(skip)` returns Go runtime frames after skipping this helper and the
	requested number of caller frames. Public Haxe APIs convert the result to
	`haxe.CallStack.StackItem` values.
**/
@:go.import("hxrt")
@:go.package("hxrt")
extern class NativeStack {
	@:go.name("NativeStackCapture")
	public static function capture(skip:Int):Array<NativeStackFrame>;

	@:go.name("NativeStackIsFrameSlice")
	public static function isFrameSlice(value:Any):Bool;
}
