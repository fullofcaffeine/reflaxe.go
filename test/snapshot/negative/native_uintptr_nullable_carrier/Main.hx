import go.NativeUintptr;

extern class NullableSource {
	public static function value():Null<NativeUintptr>;
}

@:goNative class Main {
	static function main():Void {
		final value:Null<NativeUintptr> = NullableSource.value();
		Sys.println(value == null);
	}
}
