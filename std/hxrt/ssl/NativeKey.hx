package hxrt.ssl;

import go.NativeSlice;

/**
	What: Exposes typed native PEM, DER, and file key parsing capabilities.
	Why: Key parsing requires Go crypto libraries, but public constructors and byte
	conversion belong to staged `sys.ssl.Key`.
	How: Map directly to `runtime/hxrt/ssl.go` and return opaque `KeyHandle` values.
**/
@:go.import("hxrt")
@:go.package("hxrt")
extern class NativeKey {
	@:go.name("SslKeyLoadFile")
	public static function loadFile(file:String, isPublic:Bool, pass:String):KeyHandle;

	@:go.name("SslKeyReadPEM")
	public static function readPem(data:String, isPublic:Bool, pass:String):KeyHandle;

	@:go.name("SslKeyReadDERValues")
	public static function readDer(values:NativeSlice<Int>, isPublic:Bool):KeyHandle;
}
