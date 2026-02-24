import go.Go;
import go.Result;

class Main {
	static function showMaybe(value:Null<String>):String {
		return value == null ? "null" : value;
	}

	static function main() {
		var ok:Result<Int> = Result.ok(7);
		Sys.println("ok.isOk=" + ok.isOk());
		Sys.println("ok.isErr=" + ok.isErr());
		Sys.println("ok.unwrap=" + ok.unwrap());
		Sys.println("ok.error=" + showMaybe(ok.error()));

		var err:Result<Int> = Result.failure("boom");
		Sys.println("err.isOk=" + err.isOk());
		Sys.println("err.isErr=" + err.isErr());
		Sys.println("err.error=" + showMaybe(err.error()));
		try {
			Sys.println("err.unwrap=" + err.unwrap());
		} catch (e:Dynamic) {
			Sys.println("err.unwrap=throw:" + Std.string(e));
		}

		var viaGoOk:Result<String> = Go.ok("done");
		Sys.println("go.ok.unwrap=" + viaGoOk.unwrap());
		var viaGoErr:Result<String> = Go.fail("broken");
		Sys.println("go.fail.isErr=" + viaGoErr.isErr());
		Sys.println("go.fail.error=" + showMaybe(viaGoErr.error()));
	}
}
