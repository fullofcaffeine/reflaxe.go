package hxrt.process;

/**
	What
	- Typed opaque binding for one native child-process input pipe.

	Why
	- The live Go pipe and its close state are native resources that cannot be
	  represented safely by a Haxe structural value.

	How
	- Map directly to `hxrt.ProcessInput`; staged Process uses it only through
	  `NativeProcess` byte-transfer and lifecycle capabilities.
**/
@:go.import("hxrt")
@:go.package("hxrt")
@:go.name("ProcessInput")
extern class ProcessInputHandle {}
