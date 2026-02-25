package haxe.format;

/**
	Go staged stdlib override.
	Parser behavior delegates to haxe.Json.parse so runtime policy remains centralized.
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
