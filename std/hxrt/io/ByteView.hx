package hxrt.io;

/**
	What: An opaque handle to one immutable native Go byte view.

	Why: `haxe.io.BytesData` intentionally remains `[]int` so ordinary Haxe byte
	indexing keeps `Int` semantics, while Go codecs consume `[]byte`. Exposing the
	raw slice as an Array would erase that representation boundary.

	How: Runtime helpers create and consume this handle. Staged `haxe.io.Bytes`
	caches it and invalidates the cache whenever source-visible storage changes.
**/
@:go.import("hxrt")
@:go.package("hxrt")
@:go.name("ByteView")
extern class ByteView {}
