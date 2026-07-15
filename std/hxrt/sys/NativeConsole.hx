package hxrt.sys;

/**
	What
	- Typed bridge to the display capabilities used by staged `Sys.print` and
	  `Sys.println`.

	Why
	- The upstream API intentionally accepts `Dynamic`, but that untyped display
	  boundary must not spread into `NativeSys` or make selective builds acquire
	  unrelated OS-process capabilities.

	How
	- Forward each value once to the baseline `runtime/hxrt/print.go` slice.
	  `Dynamic` is confined to these two signatures because it is the public Haxe
	  contract being represented.
**/
@:go.import("hxrt")
@:go.package("hxrt")
extern class NativeConsole {
	@:go.name("Print")
	public static function print(value:Dynamic):Void;

	@:go.name("Println")
	public static function println(value:Dynamic):Void;
}
