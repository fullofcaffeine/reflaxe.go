import go.NativeByte;

extern class NullableSource {
	public static function value():Null<NativeByte>;
}

@:goNative class Main {
	static function main():Void {
		final value:Null<NativeByte> = NullableSource.value();
		Sys.println(value);
	}
}
