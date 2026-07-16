package;

import haxe.iterators.StringIterator;
import haxe.iterators.StringKeyValueIterator;

/**
	What:
	- Owns the Go-target implementation of the portable `StringTools` helper surface.

	Why:
	- The mainstream Haxe stdlib implementation cannot be used unchanged on `haxe.go`. because several mainstream helpers still fall through unsupported string-member lowering over the Go target string representation.

	How:
	- Keep macro-safe Haxe fallbacks and use typed Go `strings` bindings only for the target runtime subset.
**/
class StringTools {
	public static inline var MIN_SURROGATE_CODE_POINT = 65536;
	public static inline var MIN_HIGH_SURROGATE_CODE_POINT = 0xD800;
	public static inline var MAX_HIGH_SURROGATE_CODE_POINT = 0xDBFF;

	public static function urlEncode(s:String):String {
		var bytes = haxe.io.Bytes.ofString(s);
		var out = new StringBuf();
		var ascii = haxe.io.Bytes.alloc(1);
		for (index in 0...bytes.length) {
			var b = bytes.get(index);
			var isUnreserved = (b >= 0x41 && b <= 0x5A) || (b >= 0x61 && b <= 0x7A) || (b >= 0x30 && b <= 0x39) || b == 0x2D || b == 0x5F || b == 0x2E
				|| b == 0x7E;
			if (isUnreserved) {
				ascii.set(0, b);
				out.add(ascii.toString());
			} else if (b == 0x20) {
				out.add("%20");
			} else {
				out.add("%");
				out.add(hex(b, 2));
			}
		}
		return out.toString();
	}

	public static function urlDecode(s:String):String {
		var input = replace(s, "+", " ");
		var bytes:Array<Int> = [];
		var index = 0;
		while (index < input.length) {
			var c = input.substr(index, 1);
			if (c == "%" && index + 2 < input.length) {
				var hi = hexDigitValue(input.substr(index + 1, 1));
				var lo = hexDigitValue(input.substr(index + 2, 1));
				if (hi >= 0 && lo >= 0) {
					bytes.push((hi << 4) | lo);
					index += 3;
					continue;
				}
			}

			var chunk = haxe.io.Bytes.ofString(input.charAt(index));
			for (chunkIndex in 0...chunk.length) {
				bytes.push(chunk.get(chunkIndex));
			}
			index++;
		}

		var out = haxe.io.Bytes.alloc(bytes.length);
		for (byteIndex in 0...bytes.length) {
			out.set(byteIndex, bytes[byteIndex]);
		}
		return out.toString();
	}

	public static function contains(s:String, value:String):Bool {
		return containsImpl(s, value);
	}

	/**
		What: preserve the optional `quotes` behavior of `StringTools.htmlEscape`.

		Why: an omitted optional `Bool` is `null` in Haxe and therefore behaves as
		false, while treating the erased Go value as an unconditional `bool` panics.

		How: guard `null` before testing the Boolean value, so both `null` and
		`false` take the non-quote-escaping path used by the mainstream Haxe stdlib.
	**/
	public static function htmlEscape(s:String, ?quotes:Bool):String {
		s = replace(s, "&", "&amp;");
		s = replace(s, "<", "&lt;");
		s = replace(s, ">", "&gt;");
		if (quotes != null && quotes == true) {
			s = replace(s, "\"", "&quot;");
			s = replace(s, "'", "&#039;");
		}
		return s;
	}

	public static function htmlUnescape(s:String):String {
		s = replace(s, "&gt;", ">");
		s = replace(s, "&lt;", "<");
		s = replace(s, "&quot;", "\"");
		s = replace(s, "&#039;", "'");
		s = replace(s, "&amp;", "&");
		return s;
	}

	public static function startsWith(s:String, start:String):Bool {
		return startsWithImpl(s, start);
	}

	public static function endsWith(s:String, end:String):Bool {
		return endsWithImpl(s, end);
	}

	public static function isSpace(s:String, pos:Int):Bool {
		if (s.length == 0 || pos < 0 || pos >= s.length) {
			return false;
		}
		var c:Int = fastCodeAt(s, pos);
		return c != -1 && ((c > 8 && c < 14) || c == 32);
	}

	public static function ltrim(s:String):String {
		var r = 0;
		while (r < s.length && isSpace(s, r)) {
			r++;
		}
		return r > 0 ? s.substr(r, s.length - r) : s;
	}

	public static function rtrim(s:String):String {
		var r = 0;
		while (r < s.length && isSpace(s, s.length - r - 1)) {
			r++;
		}
		return r > 0 ? s.substr(0, s.length - r) : s;
	}

	public static function trim(s:String):String {
		return ltrim(rtrim(s));
	}

	public static function lpad(s:String, c:String, l:Int):String {
		if (c.length <= 0) {
			return s;
		}
		var buf = "";
		while (buf.length + s.length < l) {
			buf += c;
		}
		return buf + s;
	}

	public static function rpad(s:String, c:String, l:Int):String {
		if (c.length <= 0) {
			return s;
		}
		var buf = s;
		while (buf.length < l) {
			buf += c;
		}
		return buf;
	}

	public static function replace(s:String, sub:String, by:String):String {
		if (sub.length == 0) {
			var every = new StringBuf();
			every.add(by);
			for (index in 0...s.length) {
				every.add(s.substr(index, 1));
				every.add(by);
			}
			return every.toString();
		}

		var out = new StringBuf();
		var index = 0;
		while (index < s.length) {
			if (index + sub.length <= s.length && s.substr(index, sub.length) == sub) {
				out.add(by);
				index += sub.length;
			} else {
				out.add(s.substr(index, 1));
				index++;
			}
		}
		return out.toString();
	}

	public static function hex(n:Int, ?digits:Int):String {
		var hexChars = "0123456789ABCDEF";
		var value = n;
		var out = "";
		do {
			out = hexChars.charAt(value & 15) + out;
			value = value >>> 4;
		} while (value > 0);

		var resolvedDigits:Int = digits == null ? 0 : digits;
		while (resolvedDigits != 0 && out.length < resolvedDigits) {
			out = "0" + out;
		}
		return out;
	}

	public static inline function fastCodeAt(s:String, index:Int):Int {
		var c:Null<Int> = s.charCodeAt(index);
		return c == null ? -1 : c;
	}

	public static inline function unsafeCodeAt(s:String, index:Int):Int {
		var c:Null<Int> = s.charCodeAt(index);
		return c == null ? -1 : c;
	}

	public static inline function iterator(s:String):StringIterator {
		return new StringIterator(s);
	}

	public static inline function keyValueIterator(s:String):StringKeyValueIterator {
		return new StringKeyValueIterator(s);
	}

	@:noUsing
	public static inline function isEof(c:Int):Bool {
		return c == -1;
	}

	public static inline function utf16CodePointAt(s:String, index:Int):Int {
		var c = fastCodeAt(s, index);
		if (c >= MIN_HIGH_SURROGATE_CODE_POINT && c <= MAX_HIGH_SURROGATE_CODE_POINT) {
			c = ((c - 0xD7C0) << 10) | (fastCodeAt(s, index + 1) & 0x3FF);
		}
		return c;
	}

	static function containsImpl(s:String, value:String):Bool {
		if (value.length == 0) {
			return true;
		}
		var limit = s.length - value.length;
		var index = 0;
		while (index <= limit) {
			if (s.substr(index, value.length) == value) {
				return true;
			}
			index++;
		}
		return false;
	}

	static function startsWithImpl(s:String, start:String):Bool {
		return s.length >= start.length && s.substr(0, start.length) == start;
	}

	static function endsWithImpl(s:String, end:String):Bool {
		var elen = end.length;
		var slen = s.length;
		return slen >= elen && s.substr(slen - elen, elen) == end;
	}

	static function hexDigitValue(value:String):Int {
		if (value == null || value.length == 0) {
			return -1;
		}
		var code:Int = fastCodeAt(value, 0);
		if (code == -1) {
			return -1;
		}
		if (code >= "0".code && code <= "9".code) {
			return code - "0".code;
		}
		if (code >= "A".code && code <= "F".code) {
			return code - "A".code + 10;
		}
		if (code >= "a".code && code <= "f".code) {
			return code - "a".code + 10;
		}
		return -1;
	}
}
