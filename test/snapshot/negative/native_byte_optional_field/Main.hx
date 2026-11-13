import go.NativeByte;

typedef OptionalByte = {?value:NativeByte};

@:goNative class Main {
	static function main():Void {
		final row:OptionalByte = {};
		Sys.println(row.value == null);
	}
}
