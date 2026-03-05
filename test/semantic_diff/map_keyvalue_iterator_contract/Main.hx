import haxe.ds.HashMap;
import haxe.iterators.HashMapKeyValueIterator;
import haxe.iterators.MapKeyValueIterator;

class Key {
	public final id:Int;

	public function new(id:Int) {
		this.id = id;
	}

	public function hashCode():Int {
		return id;
	}

	public function HashCode():Int {
		return id;
	}
}

class Main {
	static function summarizeStringMapViaKeyValueIterator(map:Map<String, Int>):String {
		var count = 0;
		var keyLenSum = 0;
		var valueSum = 0;
		for (entry in map.keyValueIterator()) {
			var key:String = cast entry.key;
			var value:Int = cast entry.value;
			count++;
			keyLenSum += key.length;
			valueSum += value;
		}
		return count + "|" + keyLenSum + "|" + valueSum;
	}

	static function summarizeStringMapViaWrapper(map:Map<String, Int>):String {
		var iterator = new MapKeyValueIterator(map);
		var count = 0;
		var keyLenSum = 0;
		var valueSum = 0;
		while (iterator.hasNext()) {
			var entry = iterator.next();
			var key:String = cast entry.key;
			var value:Int = cast entry.value;
			count++;
			keyLenSum += key.length;
			valueSum += value;
		}
		return count + "|" + keyLenSum + "|" + valueSum;
	}

	static function summarizeHashMapViaKeyValueIterator(map:HashMap<Key, Int>):String {
		var count = 0;
		var keyIdSum = 0;
		var valueSum = 0;
		for (entry in map.keyValueIterator()) {
			var key:Key = cast entry.key;
			var value:Int = cast entry.value;
			count++;
			keyIdSum += key.id;
			valueSum += value;
		}
		return count + "|" + keyIdSum + "|" + valueSum;
	}

	static function summarizeHashMapViaWrapper(map:HashMap<Key, Int>):String {
		var iterator = new HashMapKeyValueIterator(map);
		var count = 0;
		var keyIdSum = 0;
		var valueSum = 0;
		while (iterator.hasNext()) {
			var entry = iterator.next();
			var key:Key = cast entry.key;
			var value:Int = cast entry.value;
			count++;
			keyIdSum += key.id;
			valueSum += value;
		}
		return count + "|" + keyIdSum + "|" + valueSum;
	}

	static function main() {
		var map:Map<String, Int> = ["alpha" => 1, "beta" => 2, "gamma" => 3];
		Sys.println("map.kv=" + summarizeStringMapViaKeyValueIterator(map));
		Sys.println("map.wrapper=" + summarizeStringMapViaWrapper(map));

		var hashMap = new HashMap<Key, Int>();
		hashMap.set(new Key(10), 3);
		hashMap.set(new Key(20), 4);
		hashMap.set(new Key(30), 5);
		Sys.println("hash.kv=" + summarizeHashMapViaKeyValueIterator(hashMap));
		Sys.println("hash.wrapper=" + summarizeHashMapViaWrapper(hashMap));
	}
}
