package reflaxe.go.ast;

/**
	What: The closed set of predeclared Go types admitted by the typed target AST.

	Why: Treating builtins as arbitrary strings allowed misspellings to survive until
	`go test`, after compiler transforms and import analysis had already run.

	How: AST builders use these values directly; the legacy type parser maps exact
	Go spellings here and rejects unknown builtin-like tokens.
**/
enum abstract GoBuiltinType(String) {
	var AnyType = "any";
	var Bool = "bool";
	var Byte = "byte";
	var Complex64 = "complex64";
	var Complex128 = "complex128";
	var Error = "error";
	var Float32 = "float32";
	var Float64 = "float64";
	var Int = "int";
	var Int8 = "int8";
	var Int16 = "int16";
	var Int32 = "int32";
	var Int64 = "int64";
	var Rune = "rune";
	var StringType = "string";
	var Uint = "uint";
	var Uint8 = "uint8";
	var Uint16 = "uint16";
	var Uint32 = "uint32";
	var Uint64 = "uint64";
	var Uintptr = "uintptr";

	public inline function token():String {
		return this;
	}

	/** Resolve an exact predeclared spelling, or return null for a named type. */
	public static function fromToken(token:String):Null<GoBuiltinType> {
		return switch (token) {
			case "any": AnyType;
			case "bool": Bool;
			case "byte": Byte;
			case "complex64": Complex64;
			case "complex128": Complex128;
			case "error": Error;
			case "float32": Float32;
			case "float64": Float64;
			case "int": Int;
			case "int8": Int8;
			case "int16": Int16;
			case "int32": Int32;
			case "int64": Int64;
			case "rune": Rune;
			case "string": StringType;
			case "uint": Uint;
			case "uint8": Uint8;
			case "uint16": Uint16;
			case "uint32": Uint32;
			case "uint64": Uint64;
			case "uintptr": Uintptr;
			case _: null;
		};
	}
}
