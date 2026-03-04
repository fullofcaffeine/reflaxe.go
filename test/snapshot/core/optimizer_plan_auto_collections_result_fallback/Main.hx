class Main {
	static function main() {
		var slice:go.Slice<Null<Int>> = go.Go.newSlice();
		slice.push(cast null);

		var map:go.Map<Array<Int>, Int> = go.Go.newMap();
		var key = [1, 2];
		map.set(key, 7);
		var exists = map.exists(key);

		var result:go.Result<Null<Int>> = go.Go.ok((null : Null<Int>));
		var ok = result.isOk();

		Sys.println((exists ? 1 : 0) + (ok ? 1 : 0));
	}
}
