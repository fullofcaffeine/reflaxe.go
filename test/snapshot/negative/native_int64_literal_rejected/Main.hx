import go.NativeInt64;

@:goNative
class Main {
	static function main():Void {
		final nativeValue:NativeInt64 = 9007199254740993;
		trace(nativeValue);
	}
}
