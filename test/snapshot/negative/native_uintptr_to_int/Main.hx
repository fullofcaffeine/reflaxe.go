import go.NativeUintptr;

extern class DescriptorSource {
	public static function value():NativeUintptr;
}

@:goNative class Main {
	static function main():Void {
		final value:Int = DescriptorSource.value();
		Sys.println(value);
	}
}
