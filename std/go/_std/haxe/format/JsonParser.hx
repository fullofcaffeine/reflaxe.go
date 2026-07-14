package haxe.format;

/**
	What:
	- Owns the Go-target `haxe.format.JsonParser` compatibility surface.

	Why:
	- The mainstream Haxe stdlib implementation cannot be used unchanged on `haxe.go`. because the target centralizes dynamic JSON semantics in the typed `haxe.Json` to `hxrt` bridge.

	How:
	- Retain the expected parser entrypoints and delegate them to `haxe.Json.parse`.
**/
class JsonParser {
	public var str:String;
	public var pos:Int;

	public function new(str:String) {
		this.str = str;
		this.pos = 0;
	}

	public function doParse(?fileName:String):Dynamic {
		return haxe.Json.parse(str);
	}

	public function parseRec():Dynamic {
		return doParse(null);
	}

	public function doParseRec():Dynamic {
		return doParse(null);
	}

	public static inline function parse(str:String):Dynamic {
		return haxe.Json.parse(str);
	}
}
