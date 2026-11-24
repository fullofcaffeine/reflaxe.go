import go.NativeUint;

extern class NullableSource {
	public static function value():Null<NativeUint>;
}

@:goNative class Main {
	static function main():Void {
		final value:Null<NativeUint> = NullableSource.value();
		Sys.println(value == null);
	}
}
