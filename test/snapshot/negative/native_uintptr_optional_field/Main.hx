import go.NativeUintptr;

typedef OptionalDescriptor = {?fd:NativeUintptr};

@:goNative class Main {
	static function main():Void {
		final row:OptionalDescriptor = {};
		Sys.println(row.fd == null);
	}
}
