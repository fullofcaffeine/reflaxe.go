import go.NativeUint;

@:goNative class Main {
	static function main():Void {
		final value:NativeUint = 9;
		Sys.println(NativeUint.toInt(value));
	}
}
