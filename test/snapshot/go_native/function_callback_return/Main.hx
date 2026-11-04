/** Returned native callbacks become typed Haxe function values at import. */
@:go.import("sync")
extern class NativeOnce {
	@:go.name("OnceFunc")
	public static function once(callback:Void->Void):Void->Void;
	@:go.name("OnceValue")
	public static function onceValue<T>(callback:Void->T):Void->T;
}

class Main {
	static function main() {
		var calls = 0;
		final callback = NativeOnce.once(() -> calls++);
		final alias = callback;
		callback();
		alias();
		final factory = NativeOnce.once;
		final second = factory(() -> calls++);
		second();
		second();
		Sys.println(calls);
		Sys.println(Reflect.isFunction(callback));
		Sys.println(Reflect.compareMethods(callback, alias));
		var builds = 0;
		var total = 0;
		final nested = NativeOnce.onceValue(function() {
			builds++;
			return (delta:Int) -> {
				total += delta;
				return total;
			};
		});
		final first = nested();
		final second = nested();
		Sys.println(first(2));
		Sys.println(second(3));
		Sys.println(builds);
		final absent = NativeOnce.onceValue(() -> (null : Int->Int));
		Sys.println(absent() == null);
		final nestedFactory:(Void->(Int->Int))->(Void->(Int->Int)) = NativeOnce.onceValue;
		final fromFactory = nestedFactory(() -> (value:Int) -> value + 3);
		Sys.println(fromFactory()(7));
	}
}
