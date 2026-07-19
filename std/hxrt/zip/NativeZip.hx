package hxrt.zip;

import go.NativeSlice;

/**
	What:
	- Provides typed access to Go zlib and raw-DEFLATE execution for the staged
	  `haxe.zip` overrides.

	Why:
	- Compression is a native runtime capability, but public levels, optional
	  buffer defaults, `Bytes` conversion, and instance API behavior belong in
	  Haxe source.
	- Passing generated `haxe.io.Bytes` objects into `hxrt` would couple the
	  runtime package to application-owned layouts and break isolated Go package
	  boundaries.

	How:
	- Cross only explicit `NativeSlice<Int>`, integers, opaque typed handles, and
	  a raw-DEFLATE selector into `runtime/hxrt/zip.go`; runtime errors return
	  through the ordinary Haxe exception carrier.
	- Return `ZipCodecStep` carriers so staged source retains destination writes,
	  offsets, public result construction, and lifecycle policy.
**/
@:go.import("hxrt")
@:go.package("hxrt")
extern class NativeZip {
	@:go.name("ZipFlushNo")
	public static var FLUSH_NO(default, null):Int;

	@:go.name("ZipFlushSync")
	public static var FLUSH_SYNC(default, null):Int;

	@:go.name("ZipFlushFinish")
	public static var FLUSH_FINISH(default, null):Int;

	@:go.name("ZipDeflateCreate")
	public static function createDeflate(level:Int):ZipDeflateHandle;

	@:go.name("ZipDeflateExecute")
	public static function executeDeflate(handle:ZipDeflateHandle, values:NativeSlice<Int>, outputLimit:Int, flushMode:Int):ZipCodecStep;

	@:go.name("ZipDeflateClose")
	public static function closeDeflate(handle:ZipDeflateHandle):Void;

	@:go.name("ZipInflateCreate")
	public static function createInflate(raw:Bool):ZipInflateHandle;

	@:go.name("ZipInflateExecute")
	public static function executeInflate(handle:ZipInflateHandle, values:NativeSlice<Int>, outputLimit:Int, flushMode:Int):ZipCodecStep;

	@:go.name("ZipInflateClose")
	public static function closeInflate(handle:ZipInflateHandle):Void;

	@:go.name("ZipCompress")
	public static function compress(values:NativeSlice<Int>, level:Int):NativeSlice<Int>;

	@:go.name("ZipUncompress")
	public static function uncompress(values:NativeSlice<Int>, raw:Bool, bufferSize:Int):NativeSlice<Int>;
}
