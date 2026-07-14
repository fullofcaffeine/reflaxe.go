package haxe.iterators;

/**
	What:
	- Owns generic key/value iteration over the Go-target `IMap` bridge.

	Why:
	- The mainstream Haxe stdlib implementation cannot be used unchanged on `haxe.go`. because compiler-emitted native map carriers expose the portable surface through `haxe.Constraints.IMap`.

	How:
	- Iterate keys through `IMap` and pair each key with the value returned by the same carrier.
**/
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
