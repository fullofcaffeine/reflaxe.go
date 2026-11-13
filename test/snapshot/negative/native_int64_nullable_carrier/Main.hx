import go.NativeInt64;

@:go.import("strconv")
extern class NativeStrconv {
	@:go.name("FormatInt")
	public static function formatInt(value:NativeInt64, base:Int):String;
}

extern class NativeOptionalInt64 {
	public static function value():Null<NativeInt64>;
}

@:goNative
class Main {
	static function main():Void {
		final absent:Null<NativeInt64> = NativeOptionalInt64.value();
		Sys.println(NativeStrconv.formatInt(absent, 10));
	}
}
