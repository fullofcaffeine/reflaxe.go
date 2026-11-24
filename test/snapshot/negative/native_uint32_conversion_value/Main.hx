import go.NativeUint32;

@:goNative class Main {
	static function main():Void {
		final convert:NativeUint32->Int = NativeUint32.toInt;
		Sys.println(convert(NativeUint32.fromInt(7)));
	}
}
