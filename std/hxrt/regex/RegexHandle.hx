package hxrt.regex;

/**
	What:
	- Represents one compiled native regular expression without exposing Go's
	  `regexp.Regexp` layout.

	Why:
	- Compiled RE2 state is a real target-runtime resource. Modeling it as
	  `Dynamic` would erase the boundary retained privately by staged `EReg`.

	How:
	- Map directly to `hxrt.RegexHandle`; only `NativeRegex` accepts this opaque
	  value.
**/
@:go.import("hxrt")
@:go.package("hxrt")
@:go.name("RegexHandle")
extern class RegexHandle {}
