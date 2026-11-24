import go.NativeUintptr;

@:goNative class Main {
	static function main():Void {
		final value:NativeUintptr = 1;
		Sys.println(NativeUintptr.toInt(value));
	}
}
