import go.NativeUint32;

extern class NullableSource {
	public static function value():Null<NativeUint32>;
}

@:goNative class Main {
	static function main():Void {
		final value:Null<NativeUint32> = NullableSource.value();
		Sys.println(value == null);
	}
}
