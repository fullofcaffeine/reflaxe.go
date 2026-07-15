package hxrt.sys;

/**
	What
	- Typed carrier for one process environment key/value pair.

	Why
	- Go maps cannot stand in for the generated Haxe `Map<String, String>`
	  representation, and exposing the runtime result as `Dynamic` would erase the
	  source/runtime ownership boundary.

	How
	- Map the exported fields of `hxrt.SysEnvironmentEntry`; staged `Sys.hx`
	  constructs the public Haxe map from an array of these values.
**/
@:go.import("hxrt")
@:go.package("hxrt")
@:go.name("SysEnvironmentEntry")
extern class SysEnvironmentEntry {
	@:go.name("Key")
	public var key:String;

	@:go.name("Value")
	public var value:String;
}
