import go.NativeUint32;

@:goNative class Main {
	static function optional(?value:NativeUint32):Void {
		if (value == null)
			Sys.println("none");
	}

	static function main():Void {
		optional();
	}
}
