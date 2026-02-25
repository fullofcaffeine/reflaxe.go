@:goMetal
class LaneResultOps {
	public static function eval():String {
		var ok:go.Result<String> = go.Go.ok("done");
		var err:go.Result<String> = go.Go.fail("broken");
		var errValue = err.error();
		var errLabel = errValue == null ? "none" : errValue;
		return ok.unwrap() + "|" + (ok.isOk() ? "1" : "0") + "|" + (err.isErr() ? "1" : "0") + "|" + errLabel;
	}
}

class Main {
	static function main() {
		Sys.println(LaneResultOps.eval());
	}
}
