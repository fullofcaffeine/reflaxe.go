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
	public static function unixViaReceiver(self:GoTime):Int;
}

@:go.import("net/http")
extern class GoHttp {
	@:go.name("StatusText")
	public static function statusText(code:Int):String;
}

class Main {
	static function main() {
		var now = GoTime.now();
		var direct = now.unix();
		var viaReceiver = GoTime.unixViaReceiver(now);
		var statusOk = GoHttp.statusText(200) == "OK";

		if (direct == viaReceiver && direct > 0 && statusOk) {
			GoFmt.println(321);
		} else {
			Sys.println(-1);
		}
	}
}
