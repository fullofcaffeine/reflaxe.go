import go.NativeUint32;

@:goNative class Main {
	static function main():Void {
		final value:NativeUint32 = null;
		Sys.println(NativeUint32.toInt(value));
	}
}
