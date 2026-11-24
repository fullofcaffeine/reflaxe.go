import go.NativeUint32;

typedef OptionalPermission = {?perm:NativeUint32};

@:goNative class Main {
	static function main():Void {
		final row:OptionalPermission = {};
		Sys.println(row.perm == null);
	}
}
