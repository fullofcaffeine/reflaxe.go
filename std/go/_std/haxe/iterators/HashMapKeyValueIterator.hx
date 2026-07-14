package haxe.iterators;

import haxe.ds.HashMap;

/**
	What:
	- Owns key/value iteration for the target-owned `haxe.ds.HashMap` implementation.

	Why:
	- The mainstream Haxe stdlib implementation cannot be used unchanged on `haxe.go`. because the Go override stores entries through the constrained `hashCode()` map representation rather than the mainstream private backing state.

	How:
	- Iterate the staged map keys and resolve each value through its public target-owned map API.
**/
class HashMapKeyValueIterator<K:{function hashCode():Int;}, V> {
	final map:HashMap<K, V>;
	final keys:Iterator<K>;

	public inline function new(map:haxe.ds.HashMap<K, V>) {
		this.map = map;
		this.keys = map.keys();
	}

	public inline function hasNext():Bool {
		return keys.hasNext();
	}

	public inline function next():{key:K, value:Dynamic} {
		var key:K = cast keys.next();
		var value = map.get(key);
		return {key: key, value: value};
	}
}
