package haxe;

import haxe.CallStack.StackItem;

/**
	Target-owned fallback for native stack bridging on `haxe.go`.

	Why
	- The mainstream stdlib version is an extern facade that assumes target runtime support.
	- `haxe.go` does not have native stack capture wired into `hxrt` yet.

	What
	- Provides a deterministic implementation of the `haxe.NativeStackTrace` API surface.

	How
	- `callStack()` and `exceptionStack()` return empty arrays.
	- `saveStack()` is a no-op.
	- `toHaxe()` accepts an existing `Array<StackItem>` and applies `skip`, otherwise returns `[]`.
**/
@:dox(hide)
@:noCompletion
class NativeStackTrace {
	static public function saveStack(_exception:Any):Void {}

	static public function callStack():Any {
		var stack:Array<StackItem> = [];
		return stack;
	}

	static public function exceptionStack():Any {
		var stack:Array<StackItem> = [];
		return stack;
	}

	static public function toHaxe(nativeStackTrace:Any, skip:Int = 0):Array<StackItem> {
		if (!Std.isOfType(nativeStackTrace, Array)) {
			return [];
		}
		var stack:Array<StackItem> = cast nativeStackTrace;
		if (skip <= 0) {
			return stack.copy();
		}
		var out = [];
		for (index in skip...stack.length) {
			out.push(stack[index]);
		}
		return out;
	}
}
