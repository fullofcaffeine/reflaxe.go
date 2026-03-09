class Main {
	static function main() {
		var pos:haxe.PosInfos = {
			fileName: "Main.hx",
			lineNumber: 7,
			className: "Main",
			methodName: "main"
		};

		var posError = new haxe.exceptions.PosException("boom", null, pos);
		var argError = new haxe.exceptions.ArgumentException("count", null, null, pos);
		var notImpl = new haxe.exceptions.NotImplementedException(null, null, pos);

		Sys.println(posError.message);
		Sys.println(argError.argument);
		Sys.println(notImpl.message);
	}
}
