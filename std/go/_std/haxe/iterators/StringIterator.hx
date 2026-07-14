package haxe.iterators;

import hxrt.string.GoStringRuntime;

/**
	What:
	- Owns code-unit iteration for the Go target string representation.

	Why:
	- The mainstream Haxe stdlib implementation cannot be used unchanged on `haxe.go`. because Haxe strings lower to pointer-backed Go runtime values and need the typed `GoStringRuntime` character bridge.

	How:
	- Track the logical offset and read each code unit through the shared runtime binding.
**/
class StringIterator {
	var offset = 0;
	var s:String;

	public inline function new(s:String) {
		this.s = s;
	}

	public inline function hasNext() {
		return offset < s.length;
	}

	public inline function next() {
		return GoStringRuntime.charCodeAt(s, offset++);
	}
}
