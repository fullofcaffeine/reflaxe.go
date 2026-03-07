package haxe.ds;

import haxe.iterators.HashMapKeyValueIterator;

/**
	What
	Target-owned `HashMap` compatibility override for `haxe.go`.

	Why
	The upstream stdlib implementation is an abstract over private backing state and
	assumes direct constrained `k.hashCode()` lowering. `haxe.go` currently erases
	that constraint too aggressively in this staged-stdlib path, so the target uses
	a compatibility bridge here instead of inheriting the upstream implementation
	unchanged.

	How
	Store keys and values in `IntMap`s keyed by the constrained `hashCode()` result,
	and expose the normal `Map`-compatible surface plus the dedicated key/value
	iterator helper.
**/
class HashMap<K:{function hashCode():Int;}, V> {
	final keysByHash:IntMap<K>;
	final valuesByHash:IntMap<V>;

	public inline function new() {
		keysByHash = new IntMap();
		valuesByHash = new IntMap();
	}

	static inline function hashOf(key:Dynamic):Int {
		var hashCode:Void->Int = cast Reflect.field(key, "hashCode");
		if (hashCode == null) {
			hashCode = cast Reflect.field(key, "HashCode");
		}
		return hashCode == null ? 0 : hashCode();
	}

	@:arrayAccess public inline function set(k:K, v:V):Void {
		var hash = hashOf(k);
		keysByHash.set(hash, k);
		valuesByHash.set(hash, v);
	}

	@:arrayAccess public inline function get(k:K):Null<V> {
		return valuesByHash.get(hashOf(k));
	}

	public inline function exists(k:K):Bool {
		return valuesByHash.exists(hashOf(k));
	}

	public inline function remove(k:K):Bool {
		var hash = hashOf(k);
		valuesByHash.remove(hash);
		return keysByHash.remove(hash);
	}

	public inline function keys():Iterator<K> {
		var hashes = keysByHash.keys();
		return {
			hasNext: function() return hashes.hasNext(),
			next: function() return cast keysByHash.get(hashes.next())
		};
	}

	public inline function iterator():Iterator<V> {
		var hashes = valuesByHash.keys();
		return {
			hasNext: function() return hashes.hasNext(),
			next: function() return cast valuesByHash.get(hashes.next())
		};
	}

	public inline function keyValueIterator():HashMapKeyValueIterator<K, V> {
		return new HashMapKeyValueIterator(this);
	}

	public function copy():HashMap<K, V> {
		var copied = new HashMap<K, V>();
		for (hash in keysByHash.keys()) {
			copied.keysByHash.set(hash, cast keysByHash.get(hash));
			copied.valuesByHash.set(hash, cast valuesByHash.get(hash));
		}
		return copied;
	}

	public inline function clear():Void {
		keysByHash.clear();
		valuesByHash.clear();
	}
}
