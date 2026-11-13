import go.NativeByte;

extern class ByteSource {
	public static function value():NativeByte;
}

@:goNative class Main {
	static function main():Void {
		final value:Int = ByteSource.value();
		Sys.println(value);
	}
}
