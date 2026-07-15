package hxrt.process;

/**
	What
	- Typed opaque binding for one native child process.

	Why
	- Go's `exec.Cmd`, wait state, and pipes cannot be represented as portable Haxe
	  data, and exposing them as `Dynamic` would erase the ownership boundary.

	How
	- Map directly to `hxrt.Process`; only `NativeProcess` capabilities consume it.
**/
@:go.import("hxrt")
@:go.package("hxrt")
@:go.name("Process")
extern class ProcessHandle {}
