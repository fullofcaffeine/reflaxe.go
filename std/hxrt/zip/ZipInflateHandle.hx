package hxrt.zip;

/**
	What
	- Typed opaque binding for one progressive native zlib or raw-DEFLATE
	  inflater.

	Why
	- `haxe.zip.Uncompress` needs persistent partial-input state, but a
	  `Dynamic` handle or generated `haxe.io.Bytes` reference would blur the
	  staged-source/runtime ownership boundary.

	How
	- Name the exported Go `hxrt.ZipInflateHandle` type and expose no native
	  representation details to Haxe.
**/
@:go.import("hxrt")
@:go.package("hxrt")
@:go.name("ZipInflateHandle")
extern class ZipInflateHandle {}
