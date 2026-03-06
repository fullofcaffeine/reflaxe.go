import haxe.iterators.ArrayIterator;
import haxe.iterators.DynamicAccessIterator;
import haxe.iterators.StringIteratorUnicode;

@:goMetal
class LanePortableSurfaceOps {
	public static function digest(seed:Int):String {
		var arrayValues = [seed, seed + 1, seed + 3];
		var arrayIter:ArrayIterator<Int> = new ArrayIterator(arrayValues);
		var arrayCount = 0;
		var arraySum = 0;
		while (arrayIter.hasNext()) {
			arrayCount++;
			arraySum += arrayIter.next();
		}

		var dynamicValues:haxe.DynamicAccess<Int> = {};
		dynamicValues.set("alpha", seed);
		dynamicValues.set("beta", seed + 2);
		var dynIter:DynamicAccessIterator<Int> = new DynamicAccessIterator(dynamicValues);
		var dynCount = 0;
		var dynSum = 0;
		while (dynIter.hasNext()) {
			dynCount++;
			dynSum += dynIter.next();
		}

		var list = new haxe.ds.List<Int>();
		list.add(seed + 4);
		list.push(seed + 5);
		var listDigest = list.length + "|" + Std.string(list.first()) + "|" + Std.string(list.last());

		var map:Map<String, Int> = ["x" => seed, "yy" => seed + 3, "zzz" => seed + 6];
		var mapCount = 0;
		var mapKeyLen = 0;
		var mapValueSum = 0;
		for (entry in map.keyValueIterator()) {
			var key:String = cast entry.key;
			var value:Int = cast entry.value;
			mapCount++;
			mapKeyLen += key.length;
			mapValueSum += value;
		}

		var unicode = "G☺H";
		var unicodeIter:StringIteratorUnicode = new StringIteratorUnicode(unicode);
		var unicodeCount = 0;
		var unicodeSum = 0;
		while (unicodeIter.hasNext()) {
			unicodeCount++;
			unicodeSum += unicodeIter.next();
		}

		return arrayCount + "|" + arraySum + "|" + dynCount + "|" + dynSum + "|" + listDigest + "|" + mapCount + "|" + mapKeyLen + "|" + mapValueSum + "|"
			+ unicodeCount + "|" + unicodeSum;
	}
}

class Main {
	static function main() {
		Sys.println(LanePortableSurfaceOps.digest(7));
	}
}
