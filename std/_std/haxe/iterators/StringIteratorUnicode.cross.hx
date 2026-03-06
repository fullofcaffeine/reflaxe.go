package haxe.iterators;

class StringIteratorUnicode {
	var offset = 0;
	var s:String;

	public inline function new(s:String) {
		this.s = s;
	}

	public inline function hasNext() {
		return offset < s.length;
	}

	static inline function codeAt(value:String, index:Int):Int {
		return GoStringRuntime.charCodeAt(value, index);
	}

	public inline function next() {
		#if utf16
		var c = codeAt(s, offset++);
		if (c >= 0xD800 && c <= 0xDBFF && offset < s.length) {
			c = ((c - 0xD7C0) << 10) | (codeAt(s, offset++) & 0x3FF);
		}
		return c;
		#else
		return codeAt(s, offset++);
		#end
	}

	static public inline function unicodeIterator(s:String) {
		return new StringIteratorUnicode(s);
	}
}
