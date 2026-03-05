package haxe.iterators;

import haxe.ds.HashMap;

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
