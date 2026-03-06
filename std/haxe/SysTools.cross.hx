package haxe;

import haxe.ds.ReadOnlyArray;

/**
	What:
	- Go-target staged override for `haxe.SysTools`.
	- Preserves the direct command-line quoting helpers `quoteUnixArg` and
	  `quoteWinArg`, plus the public `winMetaCharacters` constant.

	Why:
	- The upstream Haxe stdlib implementation is the right semantic contract,
	  but `haxe.go` cannot reuse it unchanged yet.
	- The mainstream implementation depends on source-level operations this
	  backend still does not lower reliably in staged std code:
	  `String.indexOf`, `String.fromCharCode`, and `Array.indexOf`.
	- Keeping this behavior in staged std is still preferable to compiler-core
	  ownership because argument quoting is library semantics, not compiler
	  semantics.

	How:
	- Keep the public API aligned with upstream `haxe.SysTools`.
	- Re-express the same quoting rules using primitives this backend already
	  lowers well: `StringTools.replace`, `StringTools.fastCodeAt`, `charAt`,
	  simple loops, and string concatenation.
	- Delete this override once the missing string/array lowerings are solid
	  enough for the upstream stdlib file to compile unchanged on `haxe.go`.
**/
class SysTools {
	public static final winMetaCharacters:ReadOnlyArray<Int> = [
		" ".code,
		"(".code,
		")".code,
		"%".code,
		"!".code,
		"^".code,
		"\"".code,
		"<".code,
		">".code,
		"&".code,
		"|".code,
		"\n".code,
		"\r".code,
		",".code,
		";".code
	];

	public static function quoteUnixArg(argument:String):String {
		if (argument == "") {
			return "''";
		}
		if (!~/[^a-zA-Z0-9_@%+=:,.\/-]/.match(argument)) {
			return argument;
		}
		return "'" + StringTools.replace(argument, "'", "'\"'\"'") + "'";
	}

	public static function quoteWinArg(argument:String, escapeMetaCharacters:Bool):String {
		if (needsWinQuoting(argument)) {
			var result = "";
			var needQuote = hasSpaceOrTab(argument) || argument == "" || hasSlashAfterFirst(argument);
			if (needQuote) {
				result += "\"";
			}

			var backslashes = "";
			for (index in 0...argument.length) {
				var code = StringTools.fastCodeAt(argument, index);
				switch (code) {
					case "\\".code:
						backslashes += "\\";
					case "\"".code:
						result += backslashes;
						result += backslashes;
						backslashes = "";
						result += "\\\"";
					case _:
						if (backslashes != "") {
							result += backslashes;
							backslashes = "";
						}
						result += argument.charAt(index);
				}
			}

			result += backslashes;
			if (needQuote) {
				result += backslashes;
				result += "\"";
			}
			argument = result;
		}

		if (!escapeMetaCharacters) {
			return argument;
		}

		var escaped = "";
		for (index in 0...argument.length) {
			var code = StringTools.fastCodeAt(argument, index);
			if (isWinMetaCharacter(code)) {
				escaped += "^";
			}
			escaped += argument.charAt(index);
		}
		return escaped;
	}

	static function needsWinQuoting(argument:String):Bool {
		if (argument == "") {
			return true;
		}
		for (index in 0...argument.length) {
			var code = StringTools.fastCodeAt(argument, index);
			if (code == " ".code || code == "\t".code || code == "\\".code || code == "\"".code) {
				return true;
			}
			if (code == "/".code && index > 0) {
				return true;
			}
		}
		return false;
	}

	static function hasSpaceOrTab(argument:String):Bool {
		for (index in 0...argument.length) {
			var code = StringTools.fastCodeAt(argument, index);
			if (code == " ".code || code == "\t".code) {
				return true;
			}
		}
		return false;
	}

	static function hasSlashAfterFirst(argument:String):Bool {
		for (index in 1...argument.length) {
			if (StringTools.fastCodeAt(argument, index) == "/".code) {
				return true;
			}
		}
		return false;
	}

	static function isWinMetaCharacter(code:Int):Bool {
		return switch (code) {
			case 32, 40, 41, 37, 33, 94, 34, 60, 62, 38, 124, 10, 13, 44, 59:
				true;
			case _:
				false;
		};
	}
}
