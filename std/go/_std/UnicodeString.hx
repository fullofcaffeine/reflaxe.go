/*
 * Copyright (C)2005-2019 Haxe Foundation
 *
 * Permission is hereby granted, free of charge, to any person obtaining a
 * copy of this software and associated documentation files (the "Software"),
 * to deal in the Software without restriction, including without limitation
 * the rights to use, copy, modify, merge, publish, distribute, sublicense,
 * and/or sell copies of the Software, and to permit persons to whom the
 * Software is furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
 * FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER
 * DEALINGS IN THE SOFTWARE.
 */

import haxe.io.Bytes;
import haxe.io.Encoding;
import haxe.iterators.StringIteratorUnicode;
import haxe.iterators.StringKeyValueIteratorUnicode;
import hxrt.string.GoStringRuntime;

/**
	What:
	- Provides the standard `UnicodeString` API with code-point indexing,
	  searching, slicing, iteration, comparison, concatenation, and validation.

	Why:
	- The mainstream Haxe stdlib implementation cannot be used unchanged on
	  `haxe.go`: its UTF-16 branch assumes native code-unit indexing, while Go
	  runtime strings are already rune-indexed. Its declaration-only abstract
	  operators also lose their string types during Go lowering.

	How:
	- Ordinary Haxe owns every library rule, including bounds, empty searches,
	  negative positions, and UTF-8 validation. A narrow typed `hxrt` binding is
	  used only to count, read, or slice the underlying Go rune representation.
**/
@:forward
@:access(StringTools)
abstract UnicodeString(String) from String to String {
	/** The number of Unicode code points in this string. **/
	public var length(get, never):Int;

	/** Creates a Unicode view of `value`. **/
	public inline function new(value:String):Void {
		this = value;
	}

	/** Returns an iterator over this string's Unicode code points. **/
	public inline function iterator():StringIteratorUnicode {
		return new StringIteratorUnicode(this);
	}

	/** Returns an iterator over logical indices and Unicode code points. **/
	public inline function keyValueIterator():StringKeyValueIteratorUnicode {
		return new StringKeyValueIteratorUnicode(this);
	}

	/**
		Returns the character at `index`, or the empty string when the index is
		outside this string.
	**/
	public function charAt(index:Int):String {
		if (index < 0 || index >= get_length()) {
			return "";
		}
		return GoStringRuntime.sliceCodePoints(this, index, index + 1);
	}

	/**
		Returns the Unicode code point at `index`, or `null` when the index is
		outside this string.
	**/
	public function charCodeAt(index:Int):Null<Int> {
		if (index < 0 || index >= get_length()) {
			return null;
		}
		return GoStringRuntime.charCodeAt(this, index);
	}

	/**
		Returns the leftmost occurrence of `str` at or after `startIndex`.
		Negative starts are relative to the end and then clamped to zero.
	**/
	public function indexOf(str:String, ?startIndex:Int):Int {
		var total = get_length();
		var start:Int = startIndex == null ? 0 : startIndex;
		if (start < 0) {
			start = total + start;
			if (start < 0) {
				start = 0;
			}
		}
		if (start > total) {
			start = total;
		}

		var needleLength = GoStringRuntime.length(str);
		if (needleLength == 0) {
			return start;
		}
		if (needleLength > total - start) {
			return -1;
		}

		var candidate = start;
		var lastCandidate = total - needleLength;
		while (candidate <= lastCandidate) {
			if (matchesAt(str, needleLength, candidate)) {
				return candidate;
			}
			candidate++;
		}
		return -1;
	}

	/**
		Returns the rightmost occurrence of `str` whose start is no greater than
		`startIndex`. A negative explicit start cannot match.
	**/
	public function lastIndexOf(str:String, ?startIndex:Int):Int {
		var total = get_length();
		var start:Int = startIndex == null ? total : startIndex;
		if (start < 0) {
			return -1;
		}
		if (start > total) {
			start = total;
		}

		var needleLength = GoStringRuntime.length(str);
		if (needleLength == 0) {
			return start;
		}
		if (needleLength > total) {
			return -1;
		}

		var candidate = start;
		var lastCandidate = total - needleLength;
		if (candidate > lastCandidate) {
			candidate = lastCandidate;
		}
		while (candidate >= 0) {
			if (matchesAt(str, needleLength, candidate)) {
				return candidate;
			}
			candidate--;
		}
		return -1;
	}

	/**
		Returns `len` code points starting at `pos`. Negative positions are
		relative to the end; a negative length excludes code points from the end.
	**/
	public function substr(pos:Int, ?len:Int):String {
		var total = get_length();
		var start = pos;
		if (start < 0) {
			start = total + start;
			if (start < 0) {
				start = 0;
			}
		}
		if (start > total) {
			return "";
		}

		var end = total;
		if (len != null) {
			end = len < 0 ? total + len : start + len;
			if (end > total) {
				end = total;
			}
			if (end <= start) {
				return "";
			}
		}
		return GoStringRuntime.sliceCodePoints(this, start, end);
	}

	/**
		Returns code points from `startIndex` up to, but not including,
		`endIndex`. Negative bounds become zero and explicit reversed bounds swap.
	**/
	public function substring(startIndex:Int, ?endIndex:Int):String {
		var total = get_length();
		var start = startIndex < 0 ? 0 : startIndex;
		var end = total;
		if (endIndex != null) {
			end = endIndex < 0 ? 0 : endIndex;
			if (start == end) {
				return "";
			}
			if (start > end) {
				var swap = start;
				start = end;
				end = swap;
			}
		}
		if (start > total) {
			return "";
		}
		if (end > total) {
			end = total;
		}
		return GoStringRuntime.sliceCodePoints(this, start, end);
	}

	inline function get_length():Int {
		return GoStringRuntime.length(this);
	}

	function matchesAt(str:String, needleLength:Int, candidate:Int):Bool {
		var offset = 0;
		while (offset < needleLength) {
			if (GoStringRuntime.charCodeAt(this, candidate + offset) != GoStringRuntime.charCodeAt(str, offset)) {
				return false;
			}
			offset++;
		}
		return true;
	}

	/** Tells whether `bytes` is a correctly encoded byte sequence. **/
	public static function validate(bytes:Bytes, encoding:Encoding):Bool {
		switch (encoding) {
			case RawNative:
				throw "UnicodeString.validate: RawNative encoding is not supported";
			case UTF8:
				var data = bytes.getData();
				var pos = 0;
				var max = bytes.length;
				while (pos < max) {
					var c:Int = Bytes.fastGet(data, pos++);
					if (c < 0x80) {} else if (c < 0xC2) {
						return false;
					} else if (c < 0xE0) {
						if (pos + 1 > max) {
							return false;
						}
						var c2:Int = Bytes.fastGet(data, pos++);
						if (c2 < 0x80 || c2 > 0xBF) {
							return false;
						}
					} else if (c < 0xF0) {
						if (pos + 2 > max) {
							return false;
						}
						var c2:Int = Bytes.fastGet(data, pos++);
						if (c == 0xE0) {
							if (c2 < 0xA0 || c2 > 0xBF) {
								return false;
							}
						} else if (c2 < 0x80 || c2 > 0xBF) {
							return false;
						}
						var c3:Int = Bytes.fastGet(data, pos++);
						if (c3 < 0x80 || c3 > 0xBF) {
							return false;
						}
						c = (c << 16) | (c2 << 8) | c3;
						if (0xEDA080 <= c && c <= 0xEDBFBF) {
							return false;
						}
					} else if (c > 0xF4) {
						return false;
					} else {
						if (pos + 3 > max) {
							return false;
						}
						var c2:Int = Bytes.fastGet(data, pos++);
						if (c == 0xF0) {
							if (c2 < 0x90 || c2 > 0xBF) {
								return false;
							}
						} else if (c == 0xF4) {
							if (c2 < 0x80 || c2 > 0x8F) {
								return false;
							}
						} else if (c2 < 0x80 || c2 > 0xBF) {
							return false;
						}
						var c3:Int = Bytes.fastGet(data, pos++);
						if (c3 < 0x80 || c3 > 0xBF) {
							return false;
						}
						var c4:Int = Bytes.fastGet(data, pos++);
						if (c4 < 0x80 || c4 > 0xBF) {
							return false;
						}
					}
				}
				return true;
		}
		return false;
	}

	static function compare(a:UnicodeString, b:UnicodeString):Int {
		var left:String = a;
		var right:String = b;
		var leftLength = GoStringRuntime.length(left);
		var rightLength = GoStringRuntime.length(right);
		var limit = leftLength < rightLength ? leftLength : rightLength;
		var index = 0;
		while (index < limit) {
			var leftCode = GoStringRuntime.charCodeAt(left, index);
			var rightCode = GoStringRuntime.charCodeAt(right, index);
			if (leftCode < rightCode) {
				return -1;
			}
			if (leftCode > rightCode) {
				return 1;
			}
			index++;
		}
		return leftLength < rightLength ? -1 : (leftLength > rightLength ? 1 : 0);
	}

	@:op(A < B)
	static function lessThan(a:UnicodeString, b:UnicodeString):Bool {
		return compare(a, b) < 0;
	}

	@:op(A <= B)
	static function lessThanOrEqual(a:UnicodeString, b:UnicodeString):Bool {
		return compare(a, b) <= 0;
	}

	@:op(A > B)
	static function greaterThan(a:UnicodeString, b:UnicodeString):Bool {
		return compare(a, b) > 0;
	}

	@:op(A >= B)
	static function greaterThanOrEqual(a:UnicodeString, b:UnicodeString):Bool {
		return compare(a, b) >= 0;
	}

	@:op(A == B)
	static inline function equal(a:UnicodeString, b:UnicodeString):Bool {
		return (a : String) == (b : String);
	}

	@:op(A != B)
	static inline function notEqual(a:UnicodeString, b:UnicodeString):Bool {
		return (a : String) != (b : String);
	}

	/**
		`+=` intentionally uses Haxe's normal compound-assignment rewrite through
		these typed `+` operators. A source-bodied `A += B` overload would lower as
		a discarded function call on Go instead of assigning the returned string.
	**/
	@:op(A + B)
	static inline function addUnicode(a:UnicodeString, b:UnicodeString):UnicodeString {
		return (a : String) + (b : String);
	}

	@:op(A + B)
	static inline function addString(a:UnicodeString, b:String):UnicodeString {
		return (a : String) + b;
	}
}
