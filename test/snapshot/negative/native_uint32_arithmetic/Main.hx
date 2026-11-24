import go.NativeUint32;

extern class PermissionSource {
	public static function value():NativeUint32;
}

@:goNative class Main {
	static function main():Void {
		final sum = PermissionSource.value() + PermissionSource.value();
		Sys.println(sum);
	}
}
