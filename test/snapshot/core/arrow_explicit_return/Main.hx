class Main {
	static function main():Void {
		final rows = [{index: 0, key: 2}, {index: 1, key: 1}, {index: 2, key: 1}];
		rows.sort((left, right) -> {
			final compared = left.key - right.key;
			return compared != 0 ? compared : (left.index < right.index ? -1 : (left.index == right.index ? 0 : 1));
		});
		Sys.println([for (row in rows) row.index].join(","));
	}
}
