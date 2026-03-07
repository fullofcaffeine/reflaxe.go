import haxe.ds.HashMap;
import haxe.iterators.HashMapKeyValueIterator;

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
	static function summarize(map:HashMap<Key, String>):String {
		var iterator = new HashMapKeyValueIterator(map);
		var count = 0;
		var keySum = 0;
		var valueLen = 0;
		while (iterator.hasNext()) {
			var entry = iterator.next();
			var key:Key = cast entry.key;
			var value:String = cast entry.value;
			count++;
			keySum += key.id;
			valueLen += value.length;
		}
		return count + "|" + keySum + "|" + valueLen;
	}

	static function main() {
		var map = new HashMap<Key, String>();
		var key10 = new Key(10);
		var key20 = new Key(20);
		map.set(key10, "ten");
		map.set(key20, "twenty");
		Sys.println("get10=" + map.get(key10));
		Sys.println("get20=" + map.get(key20));
		Sys.println("exists20=" + map.exists(key20));
		Sys.println("copy=" + summarize(map.copy()));
		Sys.println("iter=" + summarize(map));
		Sys.println("remove10=" + map.remove(key10));
		Sys.println("exists10=" + map.exists(key10));
	}
}
