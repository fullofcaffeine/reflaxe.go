@:go.import("fmt")
extern class GoFmt {
	@:go.name("Println")
	public static function println(value:Int):Void;
}

@:go.import("time")
@:go.name("Time")
extern class GoTime {
	@:go.name("Now")
	public static function now():GoTime;

	@:go.name("Unix")
	public function unix():Int;

	@:go.receiver
	@:go.name("Unix")
	public static function unixViaReceiver(value:GoTime):Int;
}

@:go.import("context")
@:go.name("Context")
extern interface GoContext {}

@:go.import("context")
extern class GoContextPkg {
	@:go.name("Background")
	public static function background():GoContext;
}

@:go.import("net/http")
extern class GoHttp {
	@:go.name("StatusText")
	public static function statusText(code:Int):Dynamic;
}

class Main {
	static function main() {
		var now = GoTime.now();
		var unixDirect = now.unix();
		var unixReceiver = GoTime.unixViaReceiver(now);
		var ctx = GoContextPkg.background();
		var statusOk = Std.string(GoHttp.statusText(200)) == "OK";

		var ok = unixDirect == unixReceiver && unixDirect > 0 && ctx != null && statusOk;
		GoFmt.println(ok ? 1 : 0);
	}
}
