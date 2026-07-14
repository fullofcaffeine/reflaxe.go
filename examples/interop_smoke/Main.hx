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

@:go.import("strconv")
extern class UserGoStrconv {
	@:go.name("Atoi")
	@:go.valueError
	public static function atoi(value:String):go.Result<Int>;
}

@:goNative
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

		var valueErrorOk = UserGoStrconv.atoi("42");
		var valueErrorErr = UserGoStrconv.atoi("oops");
		var valueErrorPass = valueErrorOk.isOk() && !valueErrorOk.isErr() && valueErrorOk.unwrap() == 42 && !valueErrorErr.isOk() && valueErrorErr.isErr()
			&& valueErrorErr.error() != null;

		Fmt.println(wrappedOk && externOk && valueErrorPass ? 1 : 0);
	}
}
