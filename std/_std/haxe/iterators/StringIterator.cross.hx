package haxe.iterators;

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
