package hxrt.serialization;

/**
	What:
	- Exposes the bounded host floating-point parser needed by staged
	  Unserializer.

	Why:
	- Haxe's portable parser selects one numeric token, while the final decimal to
	  `float64` conversion is a Go standard-library representation operation.

	How:
	- Call exactly `strconv.ParseFloat`; generated field and method access now
	  reuses staged Reflect plus compiler-generated typed adapters.
**/
@:go.import("hxrt")
@:go.package("hxrt")
extern class NativeSerialization {
	@:go.name("SerializationParseFloat")
	public static function parseFloat(value:String):Float;
}
