class Main {
	static function main() {
		var slice:go.Slice<Int> = go.Go.newSlice();
		slice.push(3);
		slice.push(4);

		var map:go.Map<String, Int> = go.Go.newMap();
		map.set("k", slice.get(0));
		var fromMap = map.exists("k") ? 1 : 0;

		var result:go.Result<Int> = go.Go.ok(slice.get(1));
		var value = result.isOk() ? result.unwrap() : 0;

		Sys.println(slice.length + fromMap + value);
	}
}
