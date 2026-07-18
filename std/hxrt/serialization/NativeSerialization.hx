package hxrt.serialization;

import go.NativeSlice;

/**
	What:
	- Exposes the three representation-sensitive object capabilities needed by
	  staged Serializer and Unserializer.

	Why:
	- Generated Haxe fields and the hidden virtual-dispatch self pointer are
	  package-private Go fields. Ordinary staged source cannot read or initialize
	  them, while token policy and recursive traversal still belong in Haxe.
	- `Dynamic` is unavoidable only at this narrow boundary because the public wire
	  format deliberately erases arbitrary class field types.

	How:
	- Snapshot deterministic fields into typed `SerializationField` carriers, set
	  one decoded field with reflected destination coercion, and repair `__hx_this`
	  after constructor-free allocation.
**/
@:go.import("hxrt")
@:go.package("hxrt")
extern class NativeSerialization {
	@:go.name("SerializationFields")
	public static function fields(value:Dynamic):NativeSlice<SerializationField>;

	@:go.name("SerializationSetField")
	public static function setField(target:Dynamic, name:String, value:Dynamic):Void;

	@:go.name("SerializationBindSelf")
	public static function bindSelf(instance:Dynamic):Void;

	@:go.name("SerializationParseFloat")
	public static function parseFloat(value:String):Float;
}
