class Base {
	public var n:Int = 3;

	public function new(value:Dynamic) {
		Main.observed = ping();
	}

	public function ping():Int
		return n;
}

class Child extends Base {
	public override function ping():Int
		return n + 2;

	public function new() {
		super(Main.remember(this));
	}
}

class Main {
	static var keys = new haxe.ds.ObjectMap<Dynamic, Int>();
	static var saved:Dynamic;
	public static var observed:Int;

	public static function remember(value:Dynamic):Dynamic {
		saved = value;
		keys.set(value, 7);
		return value;
	}

	static function main() {
		var child = new Child();
		Sys.println(saved == child);
		Sys.println(keys.exists(child));
		Sys.println(keys.get(child));
		Sys.println(observed);
	}
}
