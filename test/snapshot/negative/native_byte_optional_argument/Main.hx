import go.NativeByte;

@:goNative class Main {
	static function optional(?value:NativeByte):Void {
		if (value == null)
			Sys.println("none");
	}

	static function main():Void {
		optional();
	}
}
