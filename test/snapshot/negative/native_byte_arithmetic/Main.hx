import go.NativeByte;

extern class ByteSource {
	public static function value():NativeByte;
}

@:goNative class Main {
	static function main():Void {
		final sum = ByteSource.value() + ByteSource.value();
		Sys.println(sum);
	}
}
