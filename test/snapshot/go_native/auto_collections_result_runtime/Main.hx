import go.Result;

class Main {
	static function typedPath():String {
		var slice:go.Slice<Int> = go.Go.newSlice();
		slice.push(3);
		slice.push(4);
		slice.set(1, 9);

		var map:go.Map<String, Int> = go.Go.newMap();
		map.set("alpha", 7);
		map.set("beta", 5);

		var ok:Result<Int> = go.Go.ok(42);
		var err:Result<Int> = go.Go.fail("typed");
		var errText = err.error();

		return slice.length + "," + slice.get(1) + "," + (map.exists("alpha") ? 1 : 0) + "," + (map.exists("gamma") ? 1 : 0) + "," + map.get("beta") + ","
			+ (map.get("gamma") == null ? 1 : 0) + "," + (ok.isOk() ? 1 : 0) + "," + ok.unwrap() + "," + (err.isErr() ? 1 : 0) + ","
			+ (errText != null && errText == "typed" ? 1 : 0);
	}

	static function fallbackPath():String {
		var slice:go.Slice<Null<Int>> = go.Go.newSlice();
		slice.push(cast null);
		slice.push(5);

		var map:go.Map<Array<Int>, Int> = go.Go.newMap();
		var key = [1, 2];
		map.set(key, 7);

		var okResult:Result<Null<Int>> = go.Go.ok((null : Null<Int>));
		var errResult:Result<Null<Int>> = go.Go.fail("lane");
		var errText = errResult.error();

		return (slice.get(0) == null ? 1 : 0) + "," + slice.get(1) + "," + (map.exists(key) ? 1 : 0) + "," + (map.exists([1, 2]) ? 1 : 0) + ","
			+ (map.get([3, 4]) == null ? 1 : 0) + "," + (okResult.isOk() ? 1 : 0) + "," + (okResult.unwrap() == null ? 1 : 0) + ","
			+ (errResult.isErr() ? 1 : 0) + "," + (errText != null && errText == "lane" ? 1 : 0);
	}

	static function main() {
		Sys.println("typed=" + typedPath());
		Sys.println("fallback=" + fallbackPath());
	}
}
