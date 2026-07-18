package hxrt.serialization;

/**
	What:
	- Carries one generated Haxe instance field and its erased serialization value.

	Why:
	- Haxe serialization is inherently dynamic, but returning an anonymous map would
	  additionally erase the native boundary and lose deterministic field order.

	How:
	- Map the exact `hxrt.SerializationField` carrier; staged Serializer immediately
	  consumes `name` and recursively serializes `value`.
**/
@:go.import("hxrt")
@:go.package("hxrt")
@:go.name("SerializationField")
extern class SerializationField {
	@:go.name("Name")
	public var name:String;

	@:go.name("Value")
	public var value:Dynamic;
}
