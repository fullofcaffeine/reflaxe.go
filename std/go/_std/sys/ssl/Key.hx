package sys.ssl;

import haxe.io.Bytes;

private typedef KeyHandle = Dynamic;

/**
	What
	Direct `sys.ssl.Key` support for `haxe.go`.

	Why
	The mainstream Haxe stdlib implementation cannot be used unchanged on `haxe.go`.
	- `sys.ssl.Digest` and `sys.ssl.Socket` need a concrete key representation on
	  sys targets.
	- Parsing PEM/DER is target-runtime work, not compiler lowering work.

	How
	- Keep the Haxe-facing constructors in staged std.
	- Store the native key handle behind one hidden localized `Dynamic` field and
	  delegate parsing to `hxrt`.
**/
@:goAllowRaw
class Key {
	@:noCompletion
	@:dox(hide)
	public var handle(default, null):KeyHandle;

	private function new(handle:KeyHandle) {
		this.handle = handle;
	}

	public static function loadFile(file:String, ?isPublic:Bool, ?pass:String):Key {
		return new Key(untyped __go__("hxrt.SslKeyLoadFile({0}, {1}, {2})", file, isPublic == true, pass));
	}

	public static function readPEM(data:String, isPublic:Bool, ?pass:String):Key {
		return new Key(untyped __go__("hxrt.SslKeyReadPEM({0}, {1}, {2})", data, isPublic, pass));
	}

	public static function readDER(data:Bytes, isPublic:Bool):Key {
		return new Key(untyped __go__("hxrt.SslKeyReadDER(hxrt_haxeBytesToRaw({0}), {1})", data, isPublic));
	}
}
