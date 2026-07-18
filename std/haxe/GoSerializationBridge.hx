package haxe;

/**
	What:
	- Declares the exact generated same-package hooks used by staged serialization.

	Why:
	- Public Haxe methods compile to package-private Go methods. The external hxrt
	  package cannot legally invoke those hooks, and modeling them as anonymous maps
	  would reject ordinary class instances.

	How:
	- Bind to the small compiler-generated interface-assertion bridge. These calls
	  carry only erased public hook receivers, typed Serializer/Unserializer values,
	  resolver names, and resolver results; all policy remains in staged source.
**/
extern class GoSerializationBridge {
	public static function hasSerializeHook(value:Dynamic):Bool;
	public static function callSerializeHook(value:Dynamic, serializer:Serializer):Bool;
	public static function callUnserializeHook(value:Dynamic, unserializer:Unserializer):Bool;
	public static function resolveClass(resolver:Dynamic, name:String):Class<Dynamic>;
	public static function resolveEnum(resolver:Dynamic, name:String):Enum<Dynamic>;
}
