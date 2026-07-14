package hxrt.string;

/**
	What
	- Typed bridge to the Go runtime character lookup used by staged string
	  iterators and UTF-8 compatibility code.

	Why
	- Haxe strings lower to pointer-backed `hxrt` values on this target, so code-unit
	  lookup needs the real Go runtime helper. This binding is runtime support, not
	  an upstream stdlib override, and therefore belongs under `std/hxrt`.

	How
	- Go import and package metadata map `charCodeAt` directly to
	  `hxrt.StringCharCodeAtStringPtr` while callers remain typed Haxe source.
**/
@:go.import("hxrt")
@:go.package("hxrt")
extern class GoStringRuntime {
	@:go.name("StringCharCodeAtStringPtr")
	public static function charCodeAt(value:String, index:Int):Int;
}
