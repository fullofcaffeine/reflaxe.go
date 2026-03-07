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
			var exceptionStack = CallStack.exceptionStack();
			Sys.println("exc.len=" + exceptionStack.length);
			Sys.println("exc.str=" + CallStack.toString(exceptionStack));
		}

		var nativeCall = haxe.NativeStackTrace.callStack();
		var nativeHaxe = haxe.NativeStackTrace.toHaxe(nativeCall);
		Sys.println("native.len=" + nativeHaxe.length);

		var entryRan = false;
		haxe.EntryPoint.runInMainThread(function() entryRan = true);
		haxe.EntryPoint.run();
		Sys.println("entry.ran=" + entryRan);

		var mainLoopRan = false;
		var event:haxe.MainLoop.MainEvent = null;
		event = haxe.MainLoop.add(function() {
			mainLoopRan = true;
			event.stop();
		});
		event.delay(null);
		haxe.EntryPoint.run();
		Sys.println("mainloop.ran=" + mainLoopRan);

		var timerRan = false;
		haxe.Timer.delay(function() timerRan = true, 10);
		haxe.EntryPoint.run();
		Sys.println("timer.delay=" + timerRan);

		var first = haxe.Timer.stamp();
		var second = haxe.Timer.stamp();
		Sys.println("timer.stamp.monotonic=" + (second >= first));

		var measured = haxe.Timer.measure(function() return 7);
		Sys.println("timer.measure=" + measured);
	}
}
