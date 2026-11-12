/** Invalid bindings must fail before generating a native field access. */
@:goNative
class Main {
	static function main():Void {
		var dialer = new Dialer();
		Sys.println(dialer.timeout);
	}
}

@:go.import("time")
@:go.name("Duration")
@:go.valueType
extern class Duration {}

@:go.import("net")
@:go.name("Dialer")
@:go.struct
extern class Dialer {
	public function new();
	@:go.name("Timeout")
	@:go.valueField
	public var timeout:Int;
}
