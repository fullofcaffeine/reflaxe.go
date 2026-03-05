import Lambda;
import haxe.ds.List;

class Main {
	static function sumArrayWithIter(values:Array<Int>):Int {
		var sum = 0;
		Lambda.iter(values, function(value:Int):Void {
			sum += value;
		});
		return sum;
	}

	static function sumListWithIter(values:List<Int>):Int {
		var sum = 0;
		Lambda.iter(values, function(value:Int):Void {
			sum += value;
		});
		return sum;
	}

	static function main() {
		var arr = [1, 2, 3];
		var list = new List<Int>();
		list.add(4);
		list.add(5);
		list.add(6);

		Sys.println("arr.sum=" + sumArrayWithIter(arr));
		Sys.println("list.sum=" + sumListWithIter(list));
	}
}
