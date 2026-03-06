package haxe.iterators;

class StringKeyValueIteratorUnicode {
	var byteOffset = 0;
	var charOffset = 0;
	var s:String;

	public inline function new(s:String) {
		this.s = s;
	}

	public inline function hasNext() {
		return byteOffset < s.length;
	}

	static inline function codeAt(value:String, index:Int):Int {
		return GoStringRuntime.charCodeAt(value, index);
	}

	public inline function next() {
		#if utf16
		var c = codeAt(s, byteOffset++);
		if (c >= 0xD800 && c <= 0xDBFF && byteOffset < s.length) {
			c = ((c - 0xD7C0) << 10) | (codeAt(s, byteOffset++) & 0x3FF);
		}
		return {key: charOffset++, value: c};
		#else
		return {key: charOffset++, value: codeAt(s, byteOffset++)};
		#end
	}

	static public inline function unicodeKeyValueIterator(s:String) {
		return new StringKeyValueIteratorUnicode(s);
	}
}
