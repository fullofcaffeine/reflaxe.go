class Counter {
	public var value(get, set):Int;
	public var reads:Int = 0;
	public var writes:Int = 0;

	var stored:Int = 4;

	public function new() {}

	function get_value():Int {
		reads++;
		return stored;
	}

	function set_value(value:Int):Int {
		writes++;
		stored = value;
		return value;
	}
}

class Main {
	static function main():Void {
		var source = new Counter();
		var erased:Dynamic = source;
		var view:{var value(get, set):Int;} = erased;
		Sys.println("before=" + view.value);
		var assigned = view.value = 9;
		Sys.println("assigned=" + assigned);
		view.value += 3;
		Sys.println("after=" + view.value);
		Sys.println("reads=" + source.reads);
		Sys.println("writes=" + source.writes);
		Sys.println("raw-absent=" + Std.string(!Reflect.hasField(view, "value")));
		Sys.println("raw-null=" + Std.string(Reflect.field(view, "value") == null));
		var raw:{final value:Null<Int>;} = erased;
		Sys.println("plain-null=" + Std.string(raw.value == null));
	}
}
