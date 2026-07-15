package reflaxe.go.ast;

/**
	What: A validated Go package-clause name or imported package qualifier.
	Why: Package identifiers participate in both printing and structural import use,
	so punctuation or keywords must fail before source emission.
	How: Construction validates the backend's emitted ASCII identifier subset and
	Go keyword set once. Unicode source names are normalized before this boundary.
**/
abstract GoPackageName(String) {
	private inline function new(value:String) {
		this = value;
	}

	/** Validate one package-clause name or import qualifier. */
	@:from
	public static function named(value:String):GoPackageName {
		if (!isIdentifier(value)) {
			throw 'Invalid Go package name "' + value + '"';
		}
		return new GoPackageName(value);
	}

	/** Return the validated target identifier for printing or comparison. */
	public inline function value():String {
		return this;
	}

	public inline function toString():String {
		return this;
	}

	/** Test the normalized identifier subset shared by package and type names. */
	public static function isIdentifier(value:String):Bool {
		if (value == null || value == "" || isKeyword(value)) {
			return false;
		}
		for (index in 0...value.length) {
			var code = value.charCodeAt(index);
			var valid = code == 95 || (code >= 65 && code <= 90) || (code >= 97 && code <= 122) || (index > 0 && code >= 48 && code <= 57);
			if (!valid) {
				return false;
			}
		}
		return true;
	}

	static function isKeyword(value:String):Bool {
		return switch (value) {
			case "break", "default", "func", "interface", "select", "case", "defer", "go", "map", "struct", "chan", "else", "goto", "package", "switch",
				"const", "fallthrough", "if", "range", "type", "continue", "for", "import", "return", "var": true;
			case _: false;
		};
	}
}
