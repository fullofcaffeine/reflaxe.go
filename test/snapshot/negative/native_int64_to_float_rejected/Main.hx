import go.NativeInt64;

extern class NativeValueSource {
	public static function value():NativeInt64;
}

@:goNative
class Main {
	static function main():Void {
		final ordinaryFloat:Float = NativeValueSource.value();
		trace(ordinaryFloat);
	}
}
