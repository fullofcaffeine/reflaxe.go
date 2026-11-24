import go.NativeUintptr;

@:goNative class Main {
	static function optional(?value:NativeUintptr):Void {
		if (value == null)
			Sys.println("none");
	}

	static function main():Void {
		optional();
	}
}
