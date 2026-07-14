package haxe.iterators;

/**
	What:
	- Owns indexed code-unit iteration for the Go target string representation.

	Why:
	- The mainstream Haxe stdlib implementation cannot be used unchanged on `haxe.go`. because Haxe strings lower to pointer-backed Go runtime values and need the typed `GoStringRuntime` character bridge.

	How:
	- Advance one logical offset and return the upstream index/value record shape.
**/
class StringKeyValueIterator {
	var offset = 0;
	var s:String;

	public inline function new(s:String) {
		this.s = s;
	}

	public inline function hasNext() {
		return offset < s.length;
	}

	public inline function next() {
		var current = offset;
		offset++;
		var code:Int = GoStringRuntime.charCodeAt(s, current);
		return {key: current, value: code};
	}
}
