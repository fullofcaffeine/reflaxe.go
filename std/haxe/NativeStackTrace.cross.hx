package haxe;

import haxe.CallStack.StackItem;
#if reflaxe_go_native_stack_trace
import hxrt.stack.NativeStack;
import hxrt.stack.NativeStackFrame;
#end

/**
	Target-owned fallback for native stack bridging on `haxe.go`.

	Why
	- The mainstream stdlib version is an extern facade that assumes target runtime support.
	- `haxe.go` keeps native stack capture explicit because stack frames are
	  target-sensitive diagnostics, not portable semantic-diff behavior.

	What
	- Provides a deterministic implementation of the `haxe.NativeStackTrace` API surface.

	How
	- By default, `callStack()` and `exceptionStack()` return empty arrays.
	- With `reflaxe_go_native_stack_trace`, they return an `hxrt` Go-frame carrier.
	- `saveStack()` is a no-op.
	- `toHaxe()` converts the active carrier to `Array<StackItem>` and applies `skip`.
**/
@:dox(hide)
@:noCompletion
class NativeStackTrace {
	static public function saveStack(_exception:Any):Void {}

	static public function callStack():Any {
		#if reflaxe_go_native_stack_trace
		return NativeStack.capture(1);
		#else
		var stack:Array<StackItem> = [];
		return stack;
		#end
	}

	static public function exceptionStack():Any {
		#if reflaxe_go_native_stack_trace
		return NativeStack.capture(1);
		#else
		var stack:Array<StackItem> = [];
		return stack;
		#end
	}

	static public function toHaxe(nativeStackTrace:Any, skip:Int = 0):Array<StackItem> {
		#if reflaxe_go_native_stack_trace
		if (nativeStackTrace == null || !NativeStack.isFrameSlice(nativeStackTrace)) {
			return [];
		}
		var frames:Array<NativeStackFrame> = cast nativeStackTrace;
		var out:Array<StackItem> = [];
		var start = skip < 0 ? 0 : skip;
		for (index in start...frames.length) {
			var frame = frames[index];
			out.push(FilePos(Method(null, frame.functionName), frame.file, frame.line, 0));
		}
		return out;
		#else
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
		#end
	}
}
