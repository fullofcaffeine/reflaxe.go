package reflaxe.go.naming;

import reflaxe.go.ast.GoPackageName;

class GoNaming {
	/**
		Go's predeclared identifiers: the universe-block types, constants, zero
		value and builtin functions listed in the Go specification.

		Why: these are ordinary identifiers, not keywords, so Go lets a
		declaration shadow them. Generated code relies on them everywhere
		(`any`, `string`, `len`, `append`, `nil`), so a Haxe local or
		selected-main static with one of these names breaks code the compiler
		emits in that scope.

		How: `GoCompiler.reserveGoNames` reserves every entry; locals then take
		a `_local` suffix and selected-main statics the qualified `Type_field`
		form. Keep this list in step with the Go specification's "Predeclared
		identifiers" section when a Go release adds one.
	**/
	public static final goPredeclaredIdentifiers:Array<String> = [
		// Types.
		"any",
		"bool",
		"byte",
		"comparable",
		"complex64",
		"complex128",
		"error",
		"float32",
		"float64",
		"int",
		"int8",
		"int16",
		"int32",
		"int64",
		"rune",
		"string",
		"uint",
		"uint8",
		"uint16",
		"uint32",
		"uint64",
		"uintptr",
		// Constants and the zero value.
		"true",
		"false",
		"iota",
		"nil",
		// Functions.
		"append",
		"cap",
		"clear",
		"close",
		"complex",
		"copy",
		"delete",
		"imag",
		"len",
		"make",
		"max",
		"min",
		"new",
		"panic",
		"print",
		"println",
		"real",
		"recover"
	];

	public static function normalizeIdent(name:String):String {
		if (name == "this") {
			return "self";
		}

		var sanitized = new StringBuf();
		for (index in 0...name.length) {
			var ch = name.charCodeAt(index);
			var isLower = ch >= "a".code && ch <= "z".code;
			var isUpper = ch >= "A".code && ch <= "Z".code;
			var isDigit = ch >= "0".code && ch <= "9".code;
			if (isLower || isUpper || isDigit || ch == "_".code) {
				sanitized.addChar(ch);
			} else {
				sanitized.add("_");
			}
		}

		var normalized = sanitized.toString();
		if (normalized == "") {
			normalized = "hx_tmp";
		}

		var hasNonUnderscore = false;
		for (index in 0...normalized.length) {
			if (normalized.charCodeAt(index) != "_".code) {
				hasNonUnderscore = true;
				break;
			}
		}
		if (!hasNonUnderscore) {
			normalized = "hx_tmp";
		}

		var first = normalized.charCodeAt(0);
		var startsWithDigit = first >= "0".code && first <= "9".code;
		if (startsWithDigit) {
			normalized = "hx_" + normalized;
		}

		// A Go keyword cannot be an identifier. Haxe reserves most of them too,
		// but `struct`, `const` and `goto` are ordinary Haxe names, so check the
		// whole Go keyword set rather than a partial copy of it.
		if (GoPackageName.isKeyword(normalized)) {
			return normalized + "_";
		}

		return normalized;
	}

	public static function typeSymbol(pack:Array<String>, typeName:String):String {
		var parts = [for (part in pack) normalizeIdent(part)];
		parts.push(normalizeIdent(typeName));
		return parts.join("__");
	}

	public static function constructorSymbol(pack:Array<String>, typeName:String):String {
		return "New_" + typeSymbol(pack, typeName);
	}

	/**
		Names one static field as a Go package-level symbol.

		Why: statics of the selected main class print as bare names (`helper`)
		so the entry package reads like hand-written Go. A bare name shares the
		package block with every import of every file in that Go package and
		shadows Go's predeclared identifiers for the whole package, so a static
		named like the entry symbol, an imported package (`strings`, `hxrt`) or
		a predeclared identifier (`len`, `string`, `any`) would not compile.

		What: returns the entry symbol for the main class's `main` field, the
		bare normalized name for its other statics, and the qualified
		`Type_field` form for every other class's statics.

		How: a main-class static whose bare name equals the entry symbol or any
		name in `reservedGoNames` falls back to the qualified form. That form
		starts with the Haxe type name, which is uppercase, while Go package
		names are lowercase by convention and every predeclared identifier is
		lowercase. Callers pass the names reserved before lowering (see
		`GoCompiler.reserveGoNames`), so every reference to one static gets the
		same symbol.
	**/
	public static function staticSymbol(pack:Array<String>, typeName:String, fieldName:String, selectedMainEntrySymbol:Null<String>,
			reservedGoNames:Map<String, Bool>):String {
		if (selectedMainEntrySymbol != null) {
			if (fieldName == "main") {
				return selectedMainEntrySymbol;
			}
			final normalized = normalizeIdent(fieldName);
			if (normalized == selectedMainEntrySymbol || reservedGoNames.exists(normalized)) {
				return typeSymbol(pack, typeName) + "_" + normalized;
			}
			return normalized;
		}
		return typeSymbol(pack, typeName) + "_" + normalizeIdent(fieldName);
	}
}
