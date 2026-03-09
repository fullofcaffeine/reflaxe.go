class Main {
	static function main() {
		var pos:haxe.PosInfos = {
			fileName: "Main.hx",
			lineNumber: 7,
			className: "Main",
			methodName: "main"
		};

		var posError = new haxe.exceptions.PosException("boom", null, pos);
		Sys.println("pos.message=" + posError.message);
		Sys.println("pos.string=" + posError.toString());
		Sys.println("pos.isException=" + Std.string(Std.isOfType(posError, haxe.Exception)));

		var argError = new haxe.exceptions.ArgumentException("count", null, null, pos);
		Sys.println("arg.argument=" + argError.argument);
		Sys.println("arg.message=" + argError.message);
		Sys.println("arg.string=" + argError.toString());

		var notImpl = new haxe.exceptions.NotImplementedException(null, null, pos);
		Sys.println("notImpl.message=" + notImpl.message);
		Sys.println("notImpl.string=" + notImpl.toString());
	}
}
