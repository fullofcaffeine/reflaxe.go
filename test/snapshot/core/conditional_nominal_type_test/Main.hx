import TaggedError;

class SameShape {
	public function new() {}

	public function code():Int {
		return 7;
	}
}

class Main {
	static function main():Void {
		var value:Any = new SameShape();
		if (Std.isOfType(value, Tag))
			throw "structural interface match violated nominal membership";
		Sys.println(Registry.count);
	}
}
