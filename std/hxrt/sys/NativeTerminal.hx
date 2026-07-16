package hxrt.sys;

/**
	What
	- Typed bridge to the native character-at-a-time standard-input capability.

	Why
	- Interactive terminal mode is OS state that mainstream Haxe source cannot
	  control unchanged. EOF construction and echo policy still belong in staged
	  `Sys`, not in this binding or a compiler shim.

	How
	- Expose one integer-valued operation from the footprint-explicit `terminal`
	  hxrt slice. `-1` is the typed EOF sentinel; native failures cross the normal
	  Haxe exception boundary after terminal state has been restored.
**/
@:go.import("hxrt")
@:go.package("hxrt")
extern class NativeTerminal {
	@:go.name("SysReadCharValue")
	public static function readChar():Int;
}
