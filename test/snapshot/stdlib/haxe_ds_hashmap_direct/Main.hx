import haxe.ds.HashMap;

class Key {
	public final id:Int;

	public function new(id:Int) {
		this.id = id;
	}

	public function hashCode():Int {
		return id;
	}
}

class Main {
	static function main() {
		var map = new HashMap<Key, String>();
		var key3 = new Key(3);
		var key4 = new Key(4);
		map.set(key3, "three");
		map.set(key4, "four");
		Sys.println("three=" + map.get(key3));
		Sys.println("four=" + map.get(key4));
	}
}
