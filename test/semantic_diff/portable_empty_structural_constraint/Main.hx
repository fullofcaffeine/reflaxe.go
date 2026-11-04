class Item {
	public var n:Int = 4;

	public function new() {}
}

class Main {
	static function identity<T:{}>(value:T):T
		return value;

	static function same(a:Dynamic, b:Dynamic):Bool
		return a == b;

	static function main() {
		var item = new Item();
		Sys.println(same(identity(item), item));
		Sys.println(identity(item).n);
		var object = {n: 7};
		Sys.println(same(identity(object), object));
		Sys.println(identity(object).n);
	}
}
