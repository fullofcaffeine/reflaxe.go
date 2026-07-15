package reflaxe.go.ast;

/**
	What: A validated, unquoted Go import path.
	Why: Imports previously accepted whitespace, quoting, and traversal-like segments
	that could only fail after target printing.
	How: Admit Go tooling's portable ASCII path characters and segment rules,
	store the canonical unquoted path, and let the file printer own quoting.
**/
abstract GoImportPath(String) {
	private inline function new(value:String) {
		this = value;
	}

	/** Validate and retain one canonical, unquoted import path. */
	@:from
	public static function parse(value:String):GoImportPath {
		if (!isValid(value)) {
			throw 'Invalid Go import path "' + value + '"';
		}
		return new GoImportPath(value);
	}

	/** Return the canonical unquoted path for printing or comparison. */
	public inline function value():String {
		return this;
	}

	public inline function toString():String {
		return this;
	}

	static function isValid(value:String):Bool {
		if (value == null || value == "" || StringTools.startsWith(value, "/") || StringTools.endsWith(value, "/") || StringTools.startsWith(value, "-")) {
			return false;
		}
		var segments = value.split("/");
		for (segment in segments) {
			if (!isValidSegment(segment)) {
				return false;
			}
		}
		return true;
	}

	static function isValidSegment(segment:String):Bool {
		if (segment == "" || segment == "." || segment == ".." || StringTools.endsWith(segment, ".")) {
			return false;
		}
		for (index in 0...segment.length) {
			var code = segment.charCodeAt(index);
			var valid = code == 43 || code == 45 || code == 46 || code == 95 || code == 126 || (code >= 48 && code <= 57) || (code >= 65 && code <= 90)
				|| (code >= 97 && code <= 122);
			if (!valid) {
				return false;
			}
		}

		var dot = segment.indexOf(".");
		var portableStem = (dot < 0 ? segment : segment.substr(0, dot)).toUpperCase();
		if (isReservedWindowsStem(portableStem) || hasWindowsShortNameSuffix(portableStem)) {
			return false;
		}
		return true;
	}

	static function isReservedWindowsStem(stem:String):Bool {
		return switch (stem) {
			case "CON", "PRN", "AUX", "NUL", "COM1", "COM2", "COM3", "COM4", "COM5", "COM6", "COM7", "COM8", "COM9", "LPT1", "LPT2", "LPT3", "LPT4", "LPT5",
				"LPT6", "LPT7", "LPT8", "LPT9": true;
			case _: false;
		};
	}

	static function hasWindowsShortNameSuffix(stem:String):Bool {
		var tilde = stem.lastIndexOf("~");
		if (tilde < 0 || tilde == stem.length - 1) {
			return false;
		}
		for (index in tilde + 1...stem.length) {
			var code = stem.charCodeAt(index);
			if (code < 48 || code > 57) {
				return false;
			}
		}
		return true;
	}
}
