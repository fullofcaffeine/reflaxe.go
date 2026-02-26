import Lambda;
import haxe.ds.List;

class Main {
	static function main() {
		var values = new List<Int>();
		values.add(1);
		values.add(2);
		values.add(3);
		values.add(4);
		values.add(5);

		var even = Lambda.filter(values, function(v:Int) return v % 2 == 0);
		var doubled = Lambda.map(even, function(v:Int) return v * 2);
		var total = Lambda.fold(doubled, function(v:Int, acc:Int) return acc + v, 0);

		Sys.println(Lambda.has(values, 3));
		Sys.println(Lambda.exists(values, function(v:Int) return v > 4));
		Sys.println(Lambda.count(values));
		Sys.println(Lambda.empty(new List<Int>()));
		Sys.println(total);
		Sys.println(doubled.length);
	}
}
