/** Returned native callbacks become typed Haxe function values at import. */
@:go.import("sync")
extern class NativeOnce {
	@:go.name("OnceFunc")
	public static function once(callback:Void->Void):Void->Void;
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
	}
}
