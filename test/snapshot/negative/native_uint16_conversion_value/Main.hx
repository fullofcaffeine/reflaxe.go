import go.NativeUint16;

@:goNative class Main {
	static function main():Void {
		final convert:NativeUint16->Int = NativeUint16.toInt;
		Sys.println(convert(NativeUint16.fromInt(7)));
	}
}
