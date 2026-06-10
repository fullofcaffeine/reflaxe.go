import haxe.CallStack;

class Main {
	static function nested():Void {
		var stack = CallStack.callStack();
		Sys.println("call.nonEmpty=" + (stack.length > 0));
		var rendered = CallStack.toString(stack);
		Sys.println("call.renderedNonEmpty=" + (rendered != ""));

		var native = haxe.NativeStackTrace.callStack();
		var nativeHaxe = haxe.NativeStackTrace.toHaxe(native, 0);
		var skipped = haxe.NativeStackTrace.toHaxe(native, 1);
		Sys.println("native.nonEmpty=" + (nativeHaxe.length > 0));
		Sys.println("native.skipOk=" + (skipped.length <= nativeHaxe.length));
		Sys.println("native.bogus.len=" + haxe.NativeStackTrace.toHaxe("not a stack", 0).length);
	}

	static function main() {
		nested();
	}
}
