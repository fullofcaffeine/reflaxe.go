package haxe.iterators;

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
