import hxrt.regex.NativeRegex;
import hxrt.regex.RegexHandle;
import hxrt.regex.RegexMatch;

/**
	What:
	- Implements the complete Haxe 4.3.7 `EReg` API as staged source for the Go
	  target.

	Why:
	- The mainstream Haxe stdlib implementation cannot be used unchanged on
	  `haxe.go` because `EReg` is a target-provided contract, but match state, group errors, global policy,
	  splitting, and callback mapping are public Haxe semantics rather than compiler
	  intrinsics.

	How:
	- Retain an opaque typed RE2 handle and delegate only compilation, match-index
	  discovery, and quoting to `NativeRegex`.
	- Store the last match in source and implement all public result/policy methods
	  over code-point offsets converted by the runtime capability.
**/
@:coreApi
class EReg {
	private var handle:RegexHandle;
	private var global:Bool;
	private var lastSource:Null<String>;
	private var lastMatch:Null<RegexMatch>;

	public function new(r:String, opt:String):Void {
		handle = NativeRegex.compile(r, opt);
		global = StringTools.contains(opt, "g");
		lastSource = null;
		lastMatch = null;
	}

	public function match(s:String):Bool {
		var found = NativeRegex.find(handle, s, 0);
		if (found == null) {
			lastSource = null;
			lastMatch = null;
			return false;
		}
		remember(s, found);
		return true;
	}

	public function matched(n:Int):String {
		var current = requireMatch();
		if (n < 0)
			throw "Invalid group";
		var offset = n * 2;
		if (offset + 1 >= current.indices.length)
			throw "Invalid group";
		var start = current.indices[offset];
		var end = current.indices[offset + 1];
		if (start < 0 || end < start)
			return null;
		return lastSource.substr(start, end - start);
	}

	public function matchedLeft():String {
		var current = requireMatch();
		return lastSource.substr(0, current.indices[0]);
	}

	public function matchedRight():String {
		var current = requireMatch();
		return lastSource.substr(current.indices[1]);
	}

	public function matchedPos():{pos:Int, len:Int} {
		var current = requireMatch();
		var start = current.indices[0];
		return {pos: start, len: current.indices[1] - start};
	}

	public function matchSub(s:String, pos:Int, len:Int = -1):Bool {
		var start = pos < 0 ? 0 : pos;
		var limit = len < 0 ? s.length : start + len;
		if (limit > s.length)
			limit = s.length;
		if (start > limit)
			return false;
		var searched = limit == s.length ? s : s.substr(0, limit);
		var found = NativeRegex.find(handle, searched, start);
		if (found == null)
			return false;
		remember(s, found);
		return true;
	}

	public function split(s:String):Array<String> {
		if (s.length == 0)
			return [s];
		if (!global) {
			var first = NativeRegex.find(handle, s, 0);
			if (first == null)
				return [s];
			return [s.substr(0, first.indices[0]), s.substr(first.indices[1])];
		}
		var parts = new Array<String>();
		var copyOffset = 0;
		var searchStart = 0;
		while (true) {
			var current = NativeRegex.find(handle, s, searchStart);
			if (current == null) {
				parts.push(s.substr(copyOffset));
				break;
			}
			var start = current.indices[0];
			var end = current.indices[1];
			parts.push(s.substr(copyOffset, start - copyOffset));
			copyOffset = end;
			var nextStart = end;
			if (start == end && nextStart == searchStart)
				nextStart++;
			if (nextStart >= s.length) {
				parts.push("");
				break;
			}
			searchStart = nextStart;
		}
		return parts;
	}

	public function replace(s:String, by:String):String {
		var out = new StringBuf();
		var copyOffset = 0;
		var searchStart = 0;
		while (true) {
			var current = NativeRegex.find(handle, s, searchStart);
			if (current == null)
				break;
			var start = current.indices[0];
			var end = current.indices[1];
			out.add(s.substr(copyOffset, start - copyOffset));
			out.add(expandReplacement(by, s, current));
			copyOffset = end;
			if (!global)
				break;
			if (start == end) {
				if (end >= s.length)
					break;
				searchStart = end + 1;
			} else {
				searchStart = end;
			}
		}
		out.add(s.substr(copyOffset));
		return out.toString();
	}

	public function map(s:String, f:EReg->String):String {
		var out = new StringBuf();
		var offset = 0;
		while (offset < s.length) {
			var current = NativeRegex.find(handle, s, offset);
			if (current == null) {
				out.add(s.substr(offset));
				break;
			}
			var start = current.indices[0];
			var end = current.indices[1];
			out.add(s.substr(offset, start - offset));
			remember(s, current);
			out.add(f(this));
			if (start == end) {
				out.add(s.substr(start, 1));
				offset = start + 1;
			} else {
				offset = end;
			}
			if (!global) {
				if (offset < s.length)
					out.add(s.substr(offset));
				break;
			}
		}
		return out.toString();
	}

	public static function escape(s:String):String {
		return NativeRegex.escape(s);
	}

	private inline function remember(source:String, match:RegexMatch):Void {
		lastSource = source;
		lastMatch = match;
	}

	private function requireMatch():RegexMatch {
		if (lastSource == null || lastMatch == null || lastMatch.indices.length < 2)
			throw "Invalid regex operation because no match was made";
		return lastMatch;
	}

	/**
		What:
		- Expands the Haxe replacement-template forms `$1` through `$9` and `$$`
		  against one explicitly supplied match snapshot.

		Why:
		- Go treats an identifier directly after `$1` as part of a named capture,
		  while Haxe requires `$1x` to mean capture one followed by literal `x`.
		  Delegating this policy to `regexp.ExpandString` therefore changes valid
		  portable programs.

		How:
		- Scan one code point at a time, consume only the documented two-character
		  forms, and omit an unmatched optional capture instead of appending `null`.
	**/
	private function expandReplacement(by:String, source:String, currentMatch:RegexMatch):String {
		var out = new StringBuf();
		var index = 0;
		while (index < by.length) {
			var current = by.charAt(index);
			if (current == "$" && index + 1 < by.length) {
				var next = by.charAt(index + 1);
				if (next == "$") {
					out.add("$");
					index += 2;
					continue;
				}
				var group = switch next {
					case "1": 1;
					case "2": 2;
					case "3": 3;
					case "4": 4;
					case "5": 5;
					case "6": 6;
					case "7": 7;
					case "8": 8;
					case "9": 9;
					default: 0;
				};
				if (group != 0) {
					var offset = group * 2;
					if (offset + 1 >= currentMatch.indices.length)
						throw "Invalid group";
					var start = currentMatch.indices[offset];
					var end = currentMatch.indices[offset + 1];
					var value = start < 0 || end < start ? null : source.substr(start, end - start);
					if (value != null)
						out.add(value);
					index += 2;
					continue;
				}
			}
			out.add(current);
			index++;
		}
		return out.toString();
	}
}
