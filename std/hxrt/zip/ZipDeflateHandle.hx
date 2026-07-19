package hxrt.zip;

/**
	What
	- Typed opaque binding for one progressive native zlib compressor.

	Why
	- `haxe.zip.Compress` must retain codec state across partial source and
	  destination buffers without exposing that state as `Dynamic` or coupling
	  `hxrt` to generated Haxe object layouts.

	How
	- Name the exported Go `hxrt.ZipDeflateHandle` type while leaving all of its
	  fields inaccessible to staged Haxe source.
**/
@:go.import("hxrt")
@:go.package("hxrt")
@:go.name("ZipDeflateHandle")
extern class ZipDeflateHandle {}
