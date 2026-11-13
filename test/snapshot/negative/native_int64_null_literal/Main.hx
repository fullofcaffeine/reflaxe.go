import go.NativeInt64;

@:goNative
class Main {
	static function main():Void {
		final absent:NativeInt64 = null;
		trace(absent);
	}
}
