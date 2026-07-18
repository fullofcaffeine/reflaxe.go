package hxrt.reflect;

import go.NativeSlice;

/**
	What:
	- Typed bridge to the erased Go representation operations required by staged
	  `Reflect`.

	Why:
	- The mainstream Haxe stdlib declaration cannot inspect Go maps, structs, or
	  function values by itself. Those host representation facts need a narrow
	  runtime owner, but lookup order and public reflection policy remain in Haxe.
	- `Dynamic` is unavoidable and localized here because the public `Reflect` API
	  deliberately erases both object and field value types.

	How:
	- Map each method one-for-one to `runtime/hxrt/reflect.go`. Field and method
	  probes return an explicit presence carrier; argument and field-name slices
	  use typed native slice boundaries.
**/
@:go.import("hxrt")
@:go.package("hxrt")
extern class NativeReflect {
	@:go.name("ReflectLookupField")
	public static function lookupField(object:Dynamic, field:String):ReflectFieldLookup;

	@:go.name("ReflectLookupMethod")
	public static function lookupMethod(object:Dynamic, field:String):ReflectFieldLookup;

	@:go.name("ReflectSetField")
	public static function setField(object:Dynamic, field:String, value:Dynamic):Bool;

	@:go.name("ReflectCallMethod")
	public static function callMethod(method:Dynamic, arguments:NativeSlice<Dynamic>):Dynamic;

	@:go.name("ReflectFields")
	public static function fields(object:Dynamic):NativeSlice<String>;

	@:go.name("ReflectIsFunction")
	public static function isFunction(value:Dynamic):Bool;

	@:go.name("ReflectCompare")
	public static function compare(left:Dynamic, right:Dynamic):Int;

	@:go.name("ReflectCompareMethods")
	public static function compareMethods(left:Dynamic, right:Dynamic):Bool;

	@:go.name("ReflectIsObject")
	public static function isObject(value:Dynamic):Bool;

	@:go.name("ReflectDeleteField")
	public static function deleteField(object:Dynamic, field:String):Bool;

	@:go.name("ReflectCopy")
	public static function copy(object:Dynamic):Dynamic;

	@:go.name("ReflectMakeVarArgs")
	public static function makeVarArgs(callback:Dynamic):Dynamic;
}
