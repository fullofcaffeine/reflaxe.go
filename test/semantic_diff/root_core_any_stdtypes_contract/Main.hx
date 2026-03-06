import haxe.iterators.ArrayIterator;
import haxe.iterators.ArrayKeyValueIterator;

class Main {
	static function sumAny(values:Array<Any>):Int {
		var total = 0;
		for (value in values)
			total += (cast value : Int);
		return total;
	}

	static function main() {
		var ints:Array<Any> = [1, 2, 3, 4];
		Sys.println("any.sum=" + sumAny(ints));

		var boxed:Any = {name: "portable", count: 4};
		Sys.println("any.obj=" + Std.string(Reflect.field(cast boxed, "name")) + ":" + Std.string(Reflect.field(cast boxed, "count")));

		var missing:Any = null;
		Sys.println("any.null=" + (missing == null));

		var intMap = new Map<String, Int>();
		intMap.set("nine", 9);
		var nullMissing:Null<Int> = intMap.get("missing");
		var nullPresent:Null<Int> = intMap.get("nine");
		Sys.println("null.missing=" + (nullMissing == null));
		Sys.println("null.present=" + (nullPresent == null ? -1 : nullPresent + 1));

		var iter:Iterator<Int> = cast new ArrayIterator([5, 6, 7]);
		var first = iter.next();
		var second = iter.next();
		Sys.println("iterator.first2=" + first + ":" + second + ":" + iter.hasNext());

		var kv:KeyValueIterator<Int, Int> = cast new ArrayKeyValueIterator([4, 8]);
		var firstEntry = kv.next();
		var secondEntry = kv.next();
		Sys.println("kv=" + firstEntry.key + ":" + firstEntry.value + "," + secondEntry.key + ":" + secondEntry.value);
	}
}
