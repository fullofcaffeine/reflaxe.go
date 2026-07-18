package hxrt.reflect;

/**
	What:
	- Carries one dynamic native field or method lookup result across the typed
	  `hxrt` boundary.

	Why:
	- A present Haxe field may legitimately contain `null`, so a nullable value
	  alone cannot distinguish absence from presence.

	How:
	- Map the exact exported fields of `hxrt.ReflectFieldLookup`; staged `Reflect`
	  consumes the carrier immediately and keeps lookup precedence in Haxe source.
**/
@:go.import("hxrt")
@:go.package("hxrt")
@:go.name("ReflectFieldLookup")
extern class ReflectFieldLookup {
	@:go.name("Found")
	public var found:Bool;

	@:go.name("Value")
	public var value:Dynamic;
}
