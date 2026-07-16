package haxe.iterators;

/**
	What
	- Owns key/value iteration over the Go-target erased `IMap` bridge.

	Why
	- The mainstream Haxe stdlib implementation cannot be used unchanged on
	  `haxe.go` yet because its generic `{key:K, value:V}` result can reach generated
	  concrete assignments before the erased interface value receives a Go type
	  assertion. Keeping the pair erased at this one boundary avoids invalid Go.

	How
	- Iterate through the ordinary `IMap` API, return the pair as its runtime
	  `Dynamic` fields, and let each typed consumer restore `K` and `V`. The map
	  implementations and their algorithms remain source-owned.
**/
class MapKeyValueIterator<K, V> {
	private var map:haxe.Constraints.IMap<K, V>;
	private var keys:Iterator<K>;

	public inline function new(map:haxe.Constraints.IMap<K, V>) {
		this.map = map;
		this.keys = map.keys();
	}

	public inline function hasNext():Bool {
		return keys.hasNext();
	}

	public inline function next():{key:Dynamic, value:Dynamic} {
		var key = keys.next();
		return {key: key, value: map.get(key)};
	}
}
