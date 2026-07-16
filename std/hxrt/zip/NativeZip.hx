package hxrt.zip;

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
	- Cross only `Array<Int>`, integers, and a raw-DEFLATE selector into
	  `runtime/hxrt/zip.go`; runtime errors return through the ordinary Haxe
	  exception carrier.
**/
@:go.import("hxrt")
@:go.package("hxrt")
extern class NativeZip {
	@:go.name("ZipCompress")
	public static function compress(values:Array<Int>, level:Int):Array<Int>;

	@:go.name("ZipUncompress")
	public static function uncompress(values:Array<Int>, raw:Bool, bufferSize:Int):Array<Int>;
}
