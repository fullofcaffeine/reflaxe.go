import haxe.iterators.StringIteratorUnicode;

class PortableSurfaceDigest {
	public static function compute(seed:Int):String {
		var list = new haxe.ds.List<Int>();
		list.add(seed);
		list.push(seed + 1);
		list.push(seed + 3);
		var popValue = list.pop();
		var listDigest = list.length + "|" + Std.string(list.first()) + "|" + Std.string(popValue);

		var map:Map<String, Int> = ["a" => seed, "bb" => seed + 2, "ccc" => seed + 4];
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

		var unicode = "A☺B";
		var unicodeIter:StringIteratorUnicode = new StringIteratorUnicode(unicode);
		var unicodeCount = 0;
		var unicodeSum = 0;
		while (unicodeIter.hasNext()) {
			unicodeCount++;
			unicodeSum += unicodeIter.next();
		}

		return listDigest + "|" + mapCount + "|" + mapKeyLen + "|" + mapValueSum + "|" + unicodeCount + "|" + unicodeSum;
	}
}
