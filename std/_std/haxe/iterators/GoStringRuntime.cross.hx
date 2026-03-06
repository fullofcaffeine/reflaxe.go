package haxe.iterators;

@:go.import("hxrt")
@:go.package("hxrt")
extern class GoStringRuntime {
	@:go.name("StringCharCodeAtStringPtr")
	static function charCodeAt(value:String, index:Int):Int;
}
