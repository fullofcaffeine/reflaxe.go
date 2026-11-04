typedef Ping = {function ping():Int;}

class Base {
	public function new() {
		var view:Ping = this;
		Main.saved = view;
		Main.keys.set(view, 7);
	}

	public function ping():Int
		return 1;
}

class Child extends Base {
	public function new() {
		super();
	}
}

class Main {
	public static var saved:Ping;
	public static var keys = new haxe.ds.ObjectMap<Dynamic, Int>();

	static function same(a:Dynamic, b:Dynamic):Bool
		return a == b;

	static function main() {
		var child = new Child();
		Sys.println(same(saved, child));
		Sys.println(keys.exists(saved));
		Sys.println(keys.exists(child));
	}
}
