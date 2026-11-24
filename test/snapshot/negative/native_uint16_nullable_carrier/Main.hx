import go.NativeUint16;

extern class NullableSource {
	public static function value():Null<NativeUint16>;
}

@:goNative class Main {
	static function main():Void {
		final value:Null<NativeUint16> = NullableSource.value();
		Sys.println(value == null);
	}
}
