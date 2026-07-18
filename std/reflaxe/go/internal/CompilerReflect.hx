package reflaxe.go.internal;

/**
	What:
	- Declares the narrow same-package metadata adapter consumed by staged
	  `Reflect`.

	Why:
	- The separate `hxrt` Go package cannot select generated lowercase Haxe fields
	  or methods, and it cannot know the compiler's closed class, enum, RTTI, and
	  metadata tables. Moving those facts into a runtime registry would duplicate
	  compiler authority and increase every program's footprint.
	- `Dynamic` is deliberate only at this public reflection boundary; the emitted
	  adapters recover exact generated receivers with typed switches.

	How:
	- Leave this internal extern unimported from Go so calls resolve to
	  compiler-emitted functions in the generated program package. Each function
	  belongs to the exact `reflect_metadata` capability, with generated member
	  adapters remaining separate from ordinary runtime representation inspection.
**/
extern class CompilerReflect {
	public static function typeField(object:Dynamic, field:String):Dynamic;
	public static function hasTypeField(object:Dynamic, field:String):Bool;
	public static function generatedField(object:Dynamic, field:String):Dynamic;
	public static function hasGeneratedField(object:Dynamic, field:String):Bool;
	public static function setGeneratedField(object:Dynamic, field:String, value:Dynamic):Bool;
	public static function generatedFields(object:Dynamic):Array<String>;
	public static function generatedMethod(object:Dynamic, field:String):Dynamic;
	public static function isEnumValue(value:Dynamic):Bool;
}
