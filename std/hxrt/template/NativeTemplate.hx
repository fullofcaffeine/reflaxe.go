package hxrt.template;

/**
	What
	- Typed bridge to the three runtime representation operations needed by staged
	  `haxe.Template` on `haxe.go`.

	Why
	- The mainstream Haxe stdlib implementation cannot be used unchanged yet:
	  portable arrays become invariant Go slices, object classification depends on
	  the erased Go carrier, and a Template macro supplies its arguments at runtime.
	  None of those operations needs compiler metadata or justifies compiler-owned
	  Template functions.
	- `Dynamic` is intentional and localized here because Template accepts dynamic
	  contexts, iterators, and macro functions as part of its public Haxe contract.

	How
	- Map each operation one-for-one to `runtime/hxrt/template.go`. The staged
	  Template source keeps field lookup, stack fallback, iteration, macro argument
	  construction, error handling, and rendering policy.
**/
@:go.import("hxrt")
@:go.package("hxrt")
extern class NativeTemplate {
	@:go.name("TemplateArrayValues")
	public static function arrayValues(value:Dynamic):Null<Array<Dynamic>>;

	@:go.name("TemplateIsObject")
	public static function isObject(value:Dynamic):Bool;

	@:go.name("TemplateCall")
	public static function call(funcValue:Dynamic, args:Array<Dynamic>):Dynamic;
}
