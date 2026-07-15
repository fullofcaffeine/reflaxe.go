package hxrt.fs;

/**
	What
	- Typed opaque binding for the native readable file resource used by staged std.

	Why
	- A live Go `os.File` cannot be represented as portable Haxe data, and exposing it
	  as `Dynamic` would erase the ownership boundary.

	How
	- Map directly to `hxrt.FileInput`; only `NativeFile` operations consume it.
**/
@:go.import("hxrt")
@:go.package("hxrt")
@:go.name("FileInput")
extern class FileInputHandle {}
