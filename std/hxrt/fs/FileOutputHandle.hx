package hxrt.fs;

/**
	What
	- Typed opaque binding for the native writable file resource used by staged std.

	Why
	- Go file ownership, standard-stream lifetime, and flush policy are native state
	  that cannot be represented safely as a Haxe structural value.

	How
	- Map directly to `hxrt.FileOutput`; only `NativeFile` operations consume it.
**/
@:go.import("hxrt")
@:go.package("hxrt")
@:go.name("FileOutput")
extern class FileOutputHandle {}
