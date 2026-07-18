package sys.ssl;

import go.NativeSlice;
import haxe.io.Bytes;
import hxrt.ssl.KeyHandle;
import hxrt.ssl.NativeKey;

/**
	What: Implements the Haxe 4.3.7 `sys.ssl.Key` constructors in staged source.
	Why: The mainstream Haxe stdlib implementation cannot be used unchanged on `haxe.go`
	because `sys.ssl.Key` is extern, and native PEM/DER key material must remain
	opaque without reducing the boundary to `Dynamic`.
	How: Store a typed `KeyHandle`, copy DER bytes to a native slice explicitly,
	and delegate only parsing to `NativeKey`.
**/
class Key {
	@:noCompletion
	@:dox(hide)
	public var handle(default, null):KeyHandle;

	private function new(handle:KeyHandle) {
		this.handle = handle;
	}

	public static function loadFile(file:String, ?isPublic:Bool, ?pass:String):Key {
		return new Key(NativeKey.loadFile(file, isPublic == true, pass));
	}

	public static function readPEM(data:String, isPublic:Bool, ?pass:String):Key {
		return new Key(NativeKey.readPem(data, isPublic, pass));
	}

	public static function readDER(data:Bytes, isPublic:Bool):Key {
		var values = new Array<Int>();
		for (index in 0...data.length)
			values.push(data.get(index));
		return new Key(NativeKey.readDer(NativeSlice.fromArray(values), isPublic));
	}
}
