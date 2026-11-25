class Base {
	public function new() {}

	public function name():String
		return "base";
}

class Mid extends Base {
	override public function name():String
		return "mid";
}

/**
	An unchecked cast of an `Any` value to a class, stored in a typed local.
	The cast itself recovers the class value, so the local must not recover
	it a second time: the output shows one `hx_recover_class_Mid` call. The
	staged `haxe.Exception.catchValue` has the same shape for
	`ValueException`, and the `String` catch below makes the program use it.
**/
class Main {
	static function main():Void {
		var value:Any = new Mid();
		var mid:Mid = cast value;
		Sys.println(mid.name());
		try {
			throw "boom";
		} catch (message:String) {
			Sys.println(message);
		}
	}
}
