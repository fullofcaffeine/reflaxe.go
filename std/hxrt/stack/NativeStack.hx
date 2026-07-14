package hxrt.stack;

/**
	What
	Typed bridge to optional Go-native stack capture in `hxrt`.

	Why
	Stack capture needs Go runtime APIs, so the implementation belongs in `hxrt`,
	not in app code, the override-only stdlib tree, or compiler-owned raw emitters.
	`isFrameSlice` accepts `Any` only at this runtime boundary because it must
	classify an unchecked caught Haxe value before the public API can type it.

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
