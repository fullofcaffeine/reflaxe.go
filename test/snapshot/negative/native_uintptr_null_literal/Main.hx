import go.NativeUintptr;

@:goNative class Main {
	static function main():Void {
		final value:NativeUintptr = null;
		Sys.println(NativeUintptr.toInt(value));
	}
}
