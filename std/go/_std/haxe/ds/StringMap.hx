package haxe.ds;

import hxrt.collections.NativeStringMap;
import hxrt.collections.StringMapHandle;

/**
	What
	- Implements the Go-target `StringMap` API as ordinary staged Haxe source over
	  a typed opaque native-storage handle.

	Why
	- The mainstream Haxe stdlib implementation cannot be used unchanged on
	  `haxe.go` because it is an extern declaration with no operation bodies. Go
	  needs a target implementation, but those bodies do not require compiler context.

	How
	- Delegate only native hash-storage facts to `NativeStringMap`; keep iteration,
	  copying, rendering, and the public `IMap` contract in Haxe source. Key
	  snapshots preserve insertion order for deterministic target behavior.
**/
class StringMap<T> implements haxe.Constraints.IMap<String, T> {
	private var h:StringMapHandle;

	public function new():Void {
		h = NativeStringMap.create();
	}

	public function set(key:String, value:T):Void {
		NativeStringMap.set(h, key, value);
	}

	public function get(key:String):Null<T> {
		return cast NativeStringMap.get(h, key);
	}

	public function exists(key:String):Bool {
		return NativeStringMap.exists(h, key);
	}

	public function remove(key:String):Bool {
		return NativeStringMap.remove(h, key);
	}

	public function keys():Iterator<String> {
		var keys = NativeStringMap.keys(h);
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
		var keys = NativeStringMap.keys(h);
		var index = 0;
		return {
			hasNext: function() {
				return index < keys.length;
			},
			next: function() {
				return cast NativeStringMap.get(h, keys[index++]);
			}
		};
	}

	@:runtime public function keyValueIterator():KeyValueIterator<String, T> {
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

	@:keep private function copyIMap():haxe.Constraints.IMap<String, T> {
		return copy();
	}

	public function copy():StringMap<T> {
		var copied = new StringMap<T>();
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
		NativeStringMap.clear(h);
	}
}
