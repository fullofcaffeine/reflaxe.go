import go.NativeUint;

@:goNative class Main {
	static function main():Void {
		final convert:NativeUint->Int = NativeUint.toInt;
		Sys.println(convert(NativeUint.fromInt(7)));
	}
}
