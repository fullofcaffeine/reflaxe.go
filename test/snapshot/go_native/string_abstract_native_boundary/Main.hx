/** A native named string is not the portable Haxe string pointer carrier. */
@:go.import("encoding/json")
@:go.name("Number")
@:go.valueType
extern class NativeNumber {}

abstract NumberToken(NativeNumber) {}
abstract NestedNumber(NumberToken) {}

/** Extern abstracts must retain their existing representation boundary. */
extern abstract ForeignText(String) {}

abstract WrappedForeign(ForeignText) {}

class Main {
	static function keepNumber(value:NestedNumber):NestedNumber {
		return value;
	}

	static function keepForeign(value:WrappedForeign):WrappedForeign {
		return value;
	}

	static function main() {
		if (keepNumber(null) != null)
			throw "native carrier null changed";
		if (keepForeign(null) != null)
			throw "extern abstract null changed";
		Sys.println("string abstract native boundary: PASS");
	}
}
