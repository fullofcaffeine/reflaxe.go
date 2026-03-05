package haxe.iterators;

class MapKeyValueIterator<K, V> {
	var map:haxe.Constraints.IMap<K, V>;
	var keys:Iterator<K>;

	public inline function new(map:haxe.Constraints.IMap<K, V>) {
		this.map = map;
		this.keys = map.keys();
	}

	public inline function hasNext():Bool {
		return keys.hasNext();
	}

	public inline function next():{key:Dynamic, value:Dynamic} {
		var key = keys.next();
		var value = map.get(key);
		return {key: key, value: value};
	}
}
