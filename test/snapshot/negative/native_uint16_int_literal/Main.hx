import go.NativeUint16;

@:goNative class Main {
	static function main():Void {
		final value:NativeUint16 = 8;
		Sys.println(NativeUint16.toInt(value));
	}
}
