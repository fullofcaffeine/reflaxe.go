import go.NativeUintptr;

extern class DescriptorSource {
	public static function value():NativeUintptr;
}

@:goNative class Main {
	static function main():Void {
		final sum = DescriptorSource.value() + DescriptorSource.value();
		Sys.println(sum);
	}
}
