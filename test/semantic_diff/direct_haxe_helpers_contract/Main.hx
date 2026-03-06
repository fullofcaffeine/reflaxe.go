class Main {
	static function main() {
		Sys.println(haxe.Log.formatOutput("x", {
			fileName: "f",
			lineNumber: 1,
			className: null,
			methodName: null,
			customParams: ["y"]
		}));

		Sys.println(haxe.SysTools.quoteUnixArg("a b"));
		Sys.println(haxe.SysTools.quoteWinArg("a b", true));
	}
}
