package hxrt.regex;

/**
	What:
	- Exposes the small set of native RE2 capabilities required by staged `EReg`.

	Why:
	- The mainstream Haxe stdlib declares `EReg` as a target-provided API, and Go's
	  compiled regex engine cannot be implemented in portable Haxe source.
	- Match state, group validation, splitting, mapping, and global-option policy
	  remain ordinary Haxe behavior and therefore do not belong in this binding.

	How:
	- Cross only strings, scalars, an opaque `RegexHandle`, and typed `RegexMatch`
	  snapshots into `runtime/hxrt/regex.go`.
**/
@:go.import("hxrt")
@:go.package("hxrt")
extern class NativeRegex {
	@:go.name("RegexCompile")
	public static function compile(pattern:String, options:String):RegexHandle;

	@:go.name("RegexFind")
	public static function find(handle:RegexHandle, source:String, position:Int):Null<RegexMatch>;

	@:go.name("RegexEscape")
	public static function escape(value:String):String;
}
