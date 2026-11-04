import haxe.ds.ArraySort;

class Main {
	static var factories = 0;

	static function comparator():Int->Int->Int {
		factories++;
		return (left, right) -> left - right;
	}

	static function main() {
		final ordinary = [3, 1, 2];
		ordinary.sort(comparator());
		Sys.println(ordinary.join(","));
		Sys.println(factories);
		final staged = [6, 4, 5];
		ArraySort.sort(staged, comparator());
		Sys.println(staged.join(","));
		Sys.println(factories);
		final empty:Array<Int> = [];
		empty.sort(null);
		Sys.println("empty");
		try {
			[2, 1].sort(null);
		} catch (error:haxe.Exception) {
			Sys.println(error.message);
		}
	}
}
