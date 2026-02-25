import go.ContextPkg;
import go.Fmt;
import go.Http;
import go.Time;

@:go.import("time")
@:go.name("Time")
extern class UserGoTime {
	@:go.name("Now")
	public static function now():UserGoTime;

	@:go.name("Unix")
	public function unix():Int;

	@:go.receiver
	@:go.name("Unix")
	public static function unixViaReceiver(value:UserGoTime):Int;
}

@:go.import("context")
@:go.name("Context")
extern interface UserGoContext {}

@:go.import("context")
extern class UserGoContextPkg {
	@:go.name("Background")
	public static function background():UserGoContext;
}

@:go.import("net/http")
extern class UserGoHttp {
	@:go.name("StatusText")
	public static function statusText(code:Int):String;
}

class Main {
	static function main() {
		var wrappedNow = Time.now();
		var wrappedUnixDirect = wrappedNow.unix();
		var wrappedUnixReceiver = Time.unixViaReceiver(wrappedNow);
		var wrappedCtx = ContextPkg.background();
		var wrappedStatusOk = Http.statusText(200) == "OK";
		var wrappedOk = wrappedUnixDirect == wrappedUnixReceiver && wrappedUnixDirect > 0 && wrappedCtx != null && wrappedStatusOk;

		var externNow = UserGoTime.now();
		var externUnixDirect = externNow.unix();
		var externUnixReceiver = UserGoTime.unixViaReceiver(externNow);
		var externCtx = UserGoContextPkg.background();
		var externStatusOk = UserGoHttp.statusText(200) == "OK";
		var externOk = externUnixDirect == externUnixReceiver && externUnixDirect > 0 && externCtx != null && externStatusOk;

		Fmt.println(wrappedOk && externOk ? 1 : 0);
	}
}
