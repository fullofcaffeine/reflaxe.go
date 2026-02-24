import go.Go;
import go.Result;

class Main {
	static function okResult():Result<Int> {
		return Result.ok(7);
	}

	static function errResult():Result<Int> {
		return Result.failure("boom");
	}

	static function main() {
		var ok = okResult();
		Sys.println(ok.isOk());
		Sys.println(ok.unwrap());

		var err = errResult();
		Sys.println(err.isErr());
		Sys.println(err.error());

		try {
			err.unwrap();
			Sys.println("unexpected");
		} catch (e:Dynamic) {
			Sys.println("caught");
			Sys.println(Std.string(e));
		}

		var okViaGo:Result<String> = Go.ok("done");
		Sys.println(okViaGo.unwrap());

		var errViaGo:Result<String> = Go.fail("bad");
		Sys.println(errViaGo.error());
	}
}
