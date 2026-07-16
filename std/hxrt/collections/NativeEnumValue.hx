package hxrt.collections;

/**
	What
	- Typed bridge for recognizing an erased Haxe enum value at runtime.

	Why
	- `EnumValueMap` recursively compares dynamic enum parameters, but Haxe cannot
	  use the `EnumValue` abstract itself as a runtime type value. Recognition must
	  inspect the generated Go enum carrier.

	How
	- Keep the erased input `Dynamic` only at this representation boundary and map
	  to `hxrt.IsEnumValue`; all comparison and tree behavior remains staged Haxe.
**/
@:go.import("hxrt")
@:go.package("hxrt")
extern class NativeEnumValue {
	@:go.name("IsEnumValue")
	public static function isEnumValue(value:Dynamic):Bool;
}
