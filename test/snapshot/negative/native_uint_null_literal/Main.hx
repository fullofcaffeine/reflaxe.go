import go.NativeUint;

@:goNative class Main {
	static function main():Void {
		final value:NativeUint = null;
		Sys.println(NativeUint.toInt(value));
	}
}
