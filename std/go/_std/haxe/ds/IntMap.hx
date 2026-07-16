package haxe.ds;

import hxrt.collections.IntMapHandle;
import hxrt.collections.NativeIntMap;

/**
	What
	- Implements the Go-target `IntMap` API as ordinary staged Haxe source over a
	  typed opaque native-storage handle.

	Why
	- The mainstream Haxe stdlib implementation cannot be used unchanged on
	  `haxe.go` because it is an extern declaration with no operation bodies. Go
	  needs a target implementation, but those bodies do not require compiler context.

	How
	- Delegate only native hash-storage facts to `NativeIntMap`; keep iteration,
	  copying, rendering, and the public `IMap` contract in Haxe source. Key
	  snapshots preserve insertion order for deterministic target behavior.
**/
class IntMap<T> implements haxe.Constraints.IMap<Int, T> {
	private var h:IntMapHandle;

	public function new():Void {
		h = NativeIntMap.create();
	}

	public function set(key:Int, value:T):Void {
		NativeIntMap.set(h, key, value);
	}

	public function get(key:Int):Null<T> {
		return cast NativeIntMap.get(h, key);
	}

	public function exists(key:Int):Bool {
		return NativeIntMap.exists(h, key);
	}

	public function remove(key:Int):Bool {
		return NativeIntMap.remove(h, key);
	}

	public function keys():Iterator<Int> {
		var keys = NativeIntMap.keys(h);
		var index = 0;
		return {
			hasNext: function() {
				return index < keys.length;
			},
			next: function() {
				return keys[index++];
			}
		};
	}

	public function iterator():Iterator<T> {
		var keys = NativeIntMap.keys(h);
		var index = 0;
		return {
			hasNext: function() {
				return index < keys.length;
			},
			next: function() {
				return cast NativeIntMap.get(h, keys[index++]);
			}
		};
	}

	@:runtime public function keyValueIterator():KeyValueIterator<Int, T> {
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

	@:keep private function copyIMap():haxe.Constraints.IMap<Int, T> {
		return copy();
	}

	public function copy():IntMap<T> {
		var copied = new IntMap<T>();
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
			out.add(key);
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
		NativeIntMap.clear(h);
	}
}
