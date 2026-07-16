package haxe.ds;

import hxrt.collections.NativeObjectMap;
import hxrt.collections.ObjectMapHandle;

/**
	What
	- Implements the Go-target `ObjectMap` API as ordinary staged Haxe source over
	  a typed opaque identity-map handle.

	Why
	- The mainstream Haxe stdlib implementation cannot be used unchanged on
	  `haxe.go` because it is an extern declaration with no operation bodies. Go
	  also needs explicit reference identity rather than structural key comparison.

	How
	- Delegate only native identity/hash-storage facts to `NativeObjectMap`; keep
	  iteration, copying, rendering, and the public `IMap` contract in Haxe source.
	  Key snapshots preserve insertion order and retain strong key references.
**/
class ObjectMap<K:{}, V> implements haxe.Constraints.IMap<K, V> {
	private var h:ObjectMapHandle;

	public function new():Void {
		h = NativeObjectMap.create();
	}

	public function set(key:K, value:V):Void {
		NativeObjectMap.set(h, key, value);
	}

	public function get(key:K):Null<V> {
		return cast NativeObjectMap.get(h, key);
	}

	public function exists(key:K):Bool {
		return NativeObjectMap.exists(h, key);
	}

	public function remove(key:K):Bool {
		return NativeObjectMap.remove(h, key);
	}

	public function keys():Iterator<K> {
		var keys = NativeObjectMap.keys(h);
		var index = 0;
		return {
			hasNext: function() {
				return index < keys.length;
			},
			next: function() {
				return cast keys[index++];
			}
		};
	}

	public function iterator():Iterator<V> {
		var keys = NativeObjectMap.keys(h);
		var index = 0;
		return {
			hasNext: function() {
				return index < keys.length;
			},
			next: function() {
				return cast NativeObjectMap.get(h, keys[index++]);
			}
		};
	}

	@:runtime public function keyValueIterator():KeyValueIterator<K, V> {
		var keys = keys();
		return {
			hasNext: function() {
				return keys.hasNext();
			},
			next: function() {
				var key = keys.next();
				return {key: key, value: cast get(key)};
			}
		};
	}

	@:keep private function getIMap(key:Dynamic):Dynamic {
		return get(cast key);
	}

	@:keep private function setIMap(key:Dynamic, value:Dynamic):Void {
		set(cast key, cast value);
	}

	@:keep private function existsIMap(key:Dynamic):Bool {
		return exists(cast key);
	}

	@:keep private function removeIMap(key:Dynamic):Bool {
		return remove(cast key);
	}

	@:keep private function copyIMap():haxe.Constraints.IMap<K, V> {
		return copy();
	}

	public function copy():ObjectMap<K, V> {
		var copied = new ObjectMap<K, V>();
		for (key in keys()) {
			copied.set(key, cast get(key));
		}
		return copied;
	}

	public function toString():String {
		var out = new StringBuf();
		out.add("[");
		var iterator = keys();
		while (iterator.hasNext()) {
			var key = iterator.next();
			out.add(Std.string(key));
			out.add(" => ");
			out.add(Std.string(get(key)));
			if (iterator.hasNext()) {
				out.add(", ");
			}
		}
		out.add("]");
		return out.toString();
	}

	public inline function clear():Void {
		NativeObjectMap.clear(h);
	}
}
