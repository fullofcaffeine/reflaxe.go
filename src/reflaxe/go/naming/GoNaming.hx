package reflaxe.go.naming;

class GoNaming {
	static final goKeywords = [
		"func",
		"type",
		"var",
		"map",
		"range",
		"package",
		"return",
		"if",
		"else",
		"for",
		"go",
		"defer",
		"select",
		"chan",
		"switch",
		"fallthrough",
		"default",
		"case"
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

		for (keyword in goKeywords) {
			if (normalized == keyword) {
				return normalized + "_";
			}
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
		package block with every import of every file in that Go package, so a
		static named like the entry symbol or an imported package (`strings`,
		`hxrt`) would not compile.

		What: returns the entry symbol for the main class's `main` field, the
		bare normalized name for its other statics, and the qualified
		`Type_field` form for every other class's statics.

		How: a main-class static whose bare name equals the entry symbol or any
		name in `reservedPackageNames` falls back to the qualified form. That
		form starts with the Haxe type name, which is uppercase, while Go
		package names are lowercase by convention. Callers pass the package
		names reserved before lowering, so every reference to one static gets
		the same symbol.
	**/
	public static function staticSymbol(pack:Array<String>, typeName:String, fieldName:String, selectedMainEntrySymbol:Null<String>,
			reservedPackageNames:Map<String, Bool>):String {
		if (selectedMainEntrySymbol != null) {
			if (fieldName == "main") {
				return selectedMainEntrySymbol;
			}
			final normalized = normalizeIdent(fieldName);
			if (normalized == selectedMainEntrySymbol || reservedPackageNames.exists(normalized)) {
				return typeSymbol(pack, typeName) + "_" + normalized;
			}
			return normalized;
		}
		return typeSymbol(pack, typeName) + "_" + normalizeIdent(fieldName);
	}
}
