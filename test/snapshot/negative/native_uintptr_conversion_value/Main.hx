import go.NativeUintptr;

@:goNative class Main {
	static function main():Void {
		final convert:NativeUintptr->Int = NativeUintptr.toInt;
		Sys.println(convert(NativeUintptr.fromInt(7)));
	}
}
