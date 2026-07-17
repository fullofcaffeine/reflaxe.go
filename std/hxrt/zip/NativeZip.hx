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
	- Cross only explicit `NativeSlice<Int>`, integers, and a raw-DEFLATE selector into
	  `runtime/hxrt/zip.go`; runtime errors return through the ordinary Haxe
	  exception carrier.
**/
@:go.import("hxrt")
@:go.package("hxrt")
extern class NativeZip {
	@:go.name("ZipCompress")
	public static function compress(values:NativeSlice<Int>, level:Int):NativeSlice<Int>;

	@:go.name("ZipUncompress")
	public static function uncompress(values:NativeSlice<Int>, raw:Bool, bufferSize:Int):NativeSlice<Int>;
}
