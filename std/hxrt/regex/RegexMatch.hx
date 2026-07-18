package hxrt.regex;

import go.NativeSlice;

/**
	What:
	- Carries one native regex match as pairs of Haxe string offsets.

	Why:
	- Go reports UTF-8 byte indexes while haxe.go String APIs use code-point
	  indexes. The native boundary must make that representation conversion exact.

	How:
	- `runtime/hxrt/regex.go` converts every match and capture pair before staged
	  `EReg` consumes this typed native slice.
**/
@:go.import("hxrt")
@:go.package("hxrt")
@:go.name("RegexMatch")
extern class RegexMatch {
	@:go.name("Indices")
	public var indices:NativeSlice<Int>;
}
