import go.Result;

class Main {
	static function main() {
		var slice:go.Slice<Int> = go.Go.newSlice();
		slice.push(3);
		slice.push(4);
		slice.set(1, 9);
		var sliceLen = slice.length;
		var sliceSecond = slice.get(1);

		var map:go.Map<String, Int> = go.Go.newMap();
		map.set("alpha", 7);
		map.set("beta", 5);
		var hasAlpha = map.exists("alpha");
		var hasGamma = map.exists("gamma");
		var beta = map.get("beta");
		var gammaIsNull = map.get("gamma") == null;

		var ok:Result<Int> = go.Go.ok(42);
		var okIsOk = ok.isOk();
		var okValue = ok.unwrap();

		var err:Result<Int> = go.Go.fail("typed");
		var errIsErr = err.isErr();
		var errText = err.error();
		var errMatches = errText != null && errText == "typed";

		Sys.println(sliceLen + "," + sliceSecond + "," + (hasAlpha ? 1 : 0) + "," + (hasGamma ? 1 : 0) + "," + beta + "," + (gammaIsNull ? 1 : 0) + ","
			+ (okIsOk ? 1 : 0) + "," + okValue + "," + (errIsErr ? 1 : 0) + "," + (errMatches ? 1 : 0));
	}
}
