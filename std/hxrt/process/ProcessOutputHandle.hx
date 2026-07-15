package hxrt.process;

/**
	What
	- Typed opaque binding for one native child-process output pipe.

	Why
	- Buffered Go pipe state cannot cross the target boundary as portable Haxe data,
	  while using `Dynamic` would hide which operations are valid.

	How
	- Map directly to `hxrt.ProcessOutput`; staged Process consumes it only through
	  the typed read and close capabilities in `NativeProcess`.
**/
@:go.import("hxrt")
@:go.package("hxrt")
@:go.name("ProcessOutput")
extern class ProcessOutputHandle {}
