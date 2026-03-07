package haxe.ds;

import haxe.iterators.HashMapKeyValueIterator;

/**
	What
	Target-owned `HashMap` compatibility override for `haxe.go`.

	Why
	The upstream stdlib implementation is an abstract over private backing state.
	`haxe.go` keeps this target-owned staged override so the entrypoint and owned
	iterator surfaces stay explicit in repo-controlled compatibility code, while
	still honoring the mainstream lowercase `hashCode()` contract without requiring
	target-specific `HashCode` aliases.

	How
	Store keys and values in `IntMap`s keyed by the constrained `hashCode()` result,
	and expose the normal `Map`-compatible surface plus the dedicated key/value
	iterator helper. The compiler preserves the method-only constraint as a local Go
	interface, so `key.hashCode()` lowers directly instead of falling back to
	reflection.
**/
class HashMap<K:{function hashCode():Int;}, V> {
	final keysByHash:IntMap<K>;
	final valuesByHash:IntMap<V>;

	public inline function new() {
		keysByHash = new IntMap();
		valuesByHash = new IntMap();
	}

	static inline function hashOf(key:{function hashCode():Int;}):Int {
		return key.hashCode();
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
