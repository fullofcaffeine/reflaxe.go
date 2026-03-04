import go.Result;

class Main {
	static function main() {
		var slice:go.Slice<Null<Int>> = go.Go.newSlice();
		slice.push(cast null);
		slice.push(5);
		var firstIsNull = slice.get(0) == null;
		var second = slice.get(1);

		var map:go.Map<Array<Int>, Int> = go.Go.newMap();
		var key = [1, 2];
		map.set(key, 7);
		var existsSameRef = map.exists(key);
		var existsEqualValue = map.exists([1, 2]);
		var missingIsNull = map.get([3, 4]) == null;

		var okResult:Result<Null<Int>> = go.Go.ok((null : Null<Int>));
		var okIsOk = okResult.isOk();
		var okUnwrapNull = okResult.unwrap() == null;

		var errResult:Result<Null<Int>> = go.Go.fail("lane");
		var errIsErr = errResult.isErr();
		var errText = errResult.error();
		var errHasMsg = errText != null && errText == "lane";

		Sys.println((firstIsNull ? 1 : 0) + "," + second + "," + (existsSameRef ? 1 : 0) + "," + (existsEqualValue ? 1 : 0) + "," + (missingIsNull ? 1 : 0)
			+ "," + (okIsOk ? 1 : 0) + "," + (okUnwrapNull ? 1 : 0) + "," + (errIsErr ? 1 : 0) + "," + (errHasMsg ? 1 : 0));
	}
}
