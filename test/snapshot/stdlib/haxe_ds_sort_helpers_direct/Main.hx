import haxe.ds.ArraySort;

class Main {
	static function cmp(a:Int, b:Int):Int {
		return a - b;
	}

	static function main() {
		var values = [5, 2, 4, 1, 3];
		ArraySort.sort(values, cmp);
		Sys.println(values[0] + "," + values[1] + "," + values[2] + "," + values[3] + "," + values[4]);
	}
}
