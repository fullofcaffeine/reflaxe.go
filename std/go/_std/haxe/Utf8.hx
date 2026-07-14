package haxe;

import haxe.io.Bytes;
import haxe.iterators.GoStringRuntime;

/**
	What
	- Go-target staged override for `haxe.Utf8`.
	- Preserves the deprecated legacy UTF-8 helper surface, including the
	  optional constructor size hint, without pushing more text-library behavior
	  into `GoCompiler`.

	Why
	The mainstream Haxe stdlib implementation cannot be used unchanged on `haxe.go`.
	- The upstream Haxe stdlib implementation is the right public API, but
	  `haxe.go` cannot reuse it unchanged yet.
	- The mainstream source path currently falls onto backend gaps that surfaced
	  in direct red probes: wrong optional-constructor lowering, missing
	  `String.fromCharCode` helper emission, bad string-compare lowering on Go
	  string pointers, and `iter` callback typing drift.
	- This is library semantics, not compiler semantics, so the fix belongs in a
	  staged std override instead of new compiler-owned decl blobs.

	How
	- Keep the public deprecated API aligned with upstream `haxe.Utf8`.
	- Use `StringBuf` for buffer writes so `addChar` stays target-owned std code.
	- Use `haxe.io.Bytes` for byte-oriented `encode`/`decode` and byte-length
	  behavior.
	- Use `UnicodeString` for codepoint iteration, indexed char access, and
	  UTF-8 character-position slicing.
**/
@:deprecated('haxe.Utf8 is deprecated. Use UnicodeString instead.')
class Utf8 {
	var __b:String;

	/**
		What
		Constructs an empty legacy UTF-8 buffer. The `size` argument is a
		deprecated capacity hint with a default value and does not affect the
		visible buffer contents.

		Why
		Upstream Haxe accepts `new haxe.Utf8(size)` and `new haxe.Utf8()` for old
		code even though the class itself is deprecated. A default `Int` keeps the
		generated Go constructor typed instead of widening the ignored hint to
		`any`.

		How
		The Go target stores the buffer as a `String`; capacity preallocation would
		not change public behavior, so the hint is intentionally ignored.
	**/
	public function new(size:Int = 0):Void {
		__b = "";
	}

	public inline function addChar(c:Int):Void {
		__b += codePointToString(c);
	}

	public inline function toString():String {
		return __b;
	}

	public static function iter(s:String, chars:Int->Void):Void {
		var unicode:UnicodeString = s;
		for (index in 0...unicode.length) {
			chars(GoStringRuntime.charCodeAt(s, index));
		}
	}

	public static function encode(s:String):String {
		var bytes = Bytes.ofString(s);
		var out = "";
		for (index in 0...bytes.length) {
			out += codePointToString(bytes.get(index));
		}
		return out;
	}

	public static function decode(s:String):String {
		var bytes = Bytes.alloc(s.length);
		for (index in 0...s.length) {
			var code = StringTools.fastCodeAt(s, index);
			bytes.set(index, code < 0 ? 0 : code & 0xFF);
		}
		return bytes.toString();
	}

	public static function charCodeAt(s:String, index:Int):Int {
		return GoStringRuntime.charCodeAt(s, index);
	}

	public static inline function validate(s:String):Bool {
		return UnicodeString.validate(Bytes.ofString(s), haxe.io.Encoding.UTF8);
	}

	public static inline function length(s:String):Int {
		return Bytes.ofString(s).length;
	}

	public static function compare(a:String, b:String):Int {
		var left = Bytes.ofString(a);
		var right = Bytes.ofString(b);
		var limit = left.length < right.length ? left.length : right.length;
		for (index in 0...limit) {
			var l = left.get(index);
			var r = right.get(index);
			if (l > r) {
				return 1;
			}
			if (l < r) {
				return -1;
			}
		}
		if (left.length > right.length) {
			return 1;
		}
		if (left.length < right.length) {
			return -1;
		}
		return 0;
	}

	public static inline function sub(s:String, pos:Int, len:Int):String {
		var unicode:UnicodeString = s;
		return unicode.substr(pos, len);
	}

	static function codePointToString(code:Int):String {
		var raw = if (code < 0x80) {
			[code];
		} else if (code < 0x800) {
			[0xC0 | (code >> 6), 0x80 | (code & 0x3F)];
		} else if (code < 0x10000) {
			[0xE0 | (code >> 12), 0x80 | ((code >> 6) & 0x3F), 0x80 | (code & 0x3F)];
		} else {
			[
				0xF0 | (code >> 18),
				0x80 | ((code >> 12) & 0x3F),
				0x80 | ((code >> 6) & 0x3F),
				0x80 | (code & 0x3F)
			];
		}
		var bytes = Bytes.alloc(raw.length);
		for (index in 0...raw.length) {
			bytes.set(index, raw[index]);
		}
		return bytes.toString();
	}
}
