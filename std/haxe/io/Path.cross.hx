package haxe.io;

/**
	What:
	- Go-target staged override for `haxe.io.Path`.
	- Preserves the mainstream stdlib helper surface for path parsing, joining,
	  normalization, trailing-slash helpers, and absolute-path checks.

	Why:
	- The upstream Haxe stdlib implementation is the right semantic model, but
	  this backend cannot reuse it unchanged yet.
	- The mainstream implementation depends on target-lowering that `haxe.go`
	  does not currently cover reliably in staged std code:
	  `String.lastIndexOf`, `String.split`, `Array.join`,
	  `String.fromCharCode`, and the constructor early-return shape.
	- Keeping this as a staged std override is still preferable to compiler-core
	  ownership because the behavior is library semantics, not compiler semantics.

	How:
	- Re-express the same helper behavior using string primitives this backend
	  already lowers cleanly: `substr`, `charAt`, `charCodeAt`, `StringBuf`,
	  simple loops, and explicit token accumulation.
	- Keep the API surface aligned with upstream `haxe.io.Path` so this override
	  can be deleted once the missing lowerings are implemented and upstream std
	  can be inherited directly.
**/
class Path {
	public var dir:Null<String>;
	public var file:String;
	public var ext:Null<String>;
	public var backslash:Bool;

	public function new(path:String) {
		dir = null;
		file = "";
		ext = null;
		backslash = false;

		if (path == "." || path == "..") {
			dir = path;
		} else {
			var slashIndex = lastIndexOfCode(path, "/".code);
			var backslashIndex = lastIndexOfCode(path, "\\".code);
			if (slashIndex < backslashIndex) {
				dir = path.substr(0, backslashIndex);
				path = path.substr(backslashIndex + 1);
				backslash = true;
			} else if (backslashIndex < slashIndex) {
				dir = path.substr(0, slashIndex);
				path = path.substr(slashIndex + 1);
			}

			var dotIndex = lastIndexOfCode(path, ".".code);
			if (dotIndex != -1) {
				ext = path.substr(dotIndex + 1);
				file = path.substr(0, dotIndex);
			} else {
				file = path;
			}
		}
	}

	public function toString():String {
		return (dir == null ? "" : dir + (backslash ? "\\" : "/")) + file + (ext == null ? "" : "." + ext);
	}

	public static function withoutExtension(path:String):String {
		var resolved = new Path(path);
		resolved.ext = null;
		return resolved.toString();
	}

	public static function withoutDirectory(path:String):String {
		var resolved = new Path(path);
		resolved.dir = null;
		return resolved.toString();
	}

	public static function directory(path:String):String {
		var resolved = new Path(path);
		return resolved.dir == null ? "" : resolved.dir;
	}

	public static function extension(path:String):String {
		var resolved = new Path(path);
		return resolved.ext == null ? "" : resolved.ext;
	}

	public static function withExtension(path:String, ext:String):String {
		var resolved = new Path(path);
		resolved.ext = ext;
		return resolved.toString();
	}

	public static function join(paths:Array<String>):String {
		var filtered = new Array<String>();
		for (segment in paths) {
			if (segment != null && segment != "") {
				filtered.push(segment);
			}
		}
		if (filtered.length == 0) {
			return "";
		}

		var path = filtered[0];
		for (index in 1...filtered.length) {
			path = addTrailingSlash(path);
			path += filtered[index];
		}
		return normalize(path);
	}

	public static function normalize(path:String):String {
		path = StringTools.replace(path, "\\", "/");
		if (path == "/") {
			return "/";
		}

		var target = new Array<String>();
		var absolute = path.length > 0 && StringTools.fastCodeAt(path, 0) == "/".code;
		for (token in splitOnSlash(path)) {
			if (token == ".." && target.length > 0 && target[target.length - 1] != "..") {
				target.pop();
			} else if (token == "") {
				if (target.length > 0 || absolute) {
					target.push(token);
				}
			} else if (token != ".") {
				target.push(token);
			}
		}

		var compact = joinWithSlash(target);
		var output = new StringBuf();
		var sawColon = false;
		var sawSlashes = false;
		for (index in 0...compact.length) {
			var code = StringTools.fastCodeAt(compact, index);
			if (code == ":".code) {
				output.add(compact.charAt(index));
				sawColon = true;
			} else if (code == "/".code && !sawColon) {
				sawSlashes = true;
			} else {
				sawColon = false;
				if (sawSlashes) {
					output.add("/");
					sawSlashes = false;
				}
				output.add(compact.charAt(index));
			}
		}

		return output.toString();
	}

	public static function addTrailingSlash(path:String):String {
		if (path.length == 0) {
			return "/";
		}
		var slashIndex = lastIndexOfCode(path, "/".code);
		var backslashIndex = lastIndexOfCode(path, "\\".code);
		return if (slashIndex < backslashIndex) {
			if (backslashIndex != path.length - 1)
				path + "\\"
			else
				path;
		} else {
			if (slashIndex != path.length - 1)
				path + "/"
			else
				path;
		}
	}

	public static function removeTrailingSlashes(path:String):String {
		while (path.length > 0) {
			var code = StringTools.fastCodeAt(path, path.length - 1);
			if (code == "/".code || code == "\\".code) {
				path = path.substr(0, path.length - 1);
				continue;
			}
			return path;
		}
		return path;
	}

	public static function isAbsolute(path:String):Bool {
		if (StringTools.startsWith(path, "/")) {
			return true;
		}
		if (path.length > 1 && StringTools.fastCodeAt(path, 1) == ":".code) {
			return true;
		}
		if (StringTools.startsWith(path, "\\\\")) {
			return true;
		}
		return false;
	}

	static function lastIndexOfCode(path:String, code:Int):Int {
		var found = -1;
		for (index in 0...path.length) {
			if (StringTools.fastCodeAt(path, index) == code) {
				found = index;
			}
		}
		return found;
	}

	static function splitOnSlash(path:String):Array<String> {
		var tokens = new Array<String>();
		var start = 0;
		for (index in 0...path.length) {
			if (StringTools.fastCodeAt(path, index) == "/".code) {
				tokens.push(path.substr(start, index - start));
				start = index + 1;
			}
		}
		tokens.push(path.substr(start, path.length - start));
		return tokens;
	}

	static function joinWithSlash(tokens:Array<String>):String {
		var output = new StringBuf();
		for (index in 0...tokens.length) {
			if (index > 0) {
				output.add("/");
			}
			output.add(tokens[index]);
		}
		return output.toString();
	}
}
