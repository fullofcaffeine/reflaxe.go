import haxe.CallStack;

class Main {
	static function main() {
		var stack = CallStack.callStack();
		Sys.println("call.len=" + stack.length);
		Sys.println("call.str=" + CallStack.toString(stack));
		Sys.println("call.copy.len=" + stack.copy().length);

		try {
			throw "boom";
		} catch (error:Dynamic) {
			var exceptionStack = CallStack.exceptionStack(false);
			Sys.println("exc.len=" + exceptionStack.length);
			Sys.println("exc.str=" + CallStack.toString(exceptionStack));
		}

		var nativeCall = haxe.NativeStackTrace.callStack();
		var nativeHaxe = haxe.NativeStackTrace.toHaxe(nativeCall, 0);
		Sys.println("native.len=" + nativeHaxe.length);
	}
}
