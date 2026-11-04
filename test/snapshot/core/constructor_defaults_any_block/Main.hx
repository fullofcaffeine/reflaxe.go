class Base {
	public final state:String;

	public function new(label:String, ?number:Int, ?suffix:String) {
		state = label + ":" + Std.string(number) + ":" + Std.string(suffix);
	}
}

class Child extends Base {
	public function new() {
		super("x", 4);
	}
}

class Main {
	static function erasedValue(value:Child):Any
		return value;

	static function main():Void {
		var value = new Child();
		// Any is the explicit input boundary whose implicit conversion is tested.
		var erased:Any = value;
		Sys.println("state=" + value.state);
		Sys.println("identity=" + Std.string(erased == value));
		Sys.println("call-identity=" + Std.string(erasedValue(value) == value));
		var base:Base = value;
		Sys.println("base-identity=" + Std.string(base == value));
	}
}
