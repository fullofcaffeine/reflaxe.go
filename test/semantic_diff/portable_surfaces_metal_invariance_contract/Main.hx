import haxe.iterators.ArrayIterator;
import haxe.iterators.ArrayKeyValueIterator;
import haxe.iterators.DynamicAccessIterator;
import haxe.iterators.DynamicAccessKeyValueIterator;
import haxe.iterators.StringIterator;
import haxe.iterators.StringIteratorUnicode;
import haxe.iterators.StringKeyValueIterator;
import haxe.iterators.StringKeyValueIteratorUnicode;

class Main {
	static function digestIterators(seed:Int):String {
		var arrayValues = [seed, seed + 3, seed + 6];
		var arrayIter:ArrayIterator<Int> = new ArrayIterator(arrayValues);
		var arrayCount = 0;
		var arraySum = 0;
		while (arrayIter.hasNext()) {
			arrayCount++;
			arraySum += arrayIter.next();
		}

		var arrayKv:ArrayKeyValueIterator<Int> = new ArrayKeyValueIterator(arrayValues);
		var arrayKeySum = 0;
		var arrayValueSum = 0;
		while (arrayKv.hasNext()) {
			var entry = arrayKv.next();
			arrayKeySum += entry.key;
			arrayValueSum += entry.value;
		}

		var dynamicValues:haxe.DynamicAccess<Int> = {};
		dynamicValues.set("alpha", seed + 1);
		dynamicValues.set("beta", seed + 2);
		dynamicValues.set("gamma", seed + 4);

		var dynIter:DynamicAccessIterator<Int> = new DynamicAccessIterator(dynamicValues);
		var dynCount = 0;
		var dynSum = 0;
		while (dynIter.hasNext()) {
			dynCount++;
			dynSum += dynIter.next();
		}

		var dynKv:DynamicAccessKeyValueIterator<Int> = new DynamicAccessKeyValueIterator(dynamicValues);
		var dynKeyLen = 0;
		var dynValueSum = 0;
		while (dynKv.hasNext()) {
			var entry = dynKv.next();
			dynKeyLen += entry.key.length;
			dynValueSum += cast entry.value;
		}

		return arrayCount
			+ "|"
			+ arraySum
			+ "|"
			+ arrayKeySum
			+ "|"
			+ arrayValueSum
			+ "|"
			+ dynCount
			+ "|"
			+ dynSum
			+ "|"
			+ dynKeyLen
			+ "|"
			+ dynValueSum;
	}

	static function digestListAndMap(seed:Int):String {
		var list = new haxe.ds.List<Int>();
		list.add(seed);
		list.push(seed + 2);
		list.push(seed + 5);
		var popValue = list.pop();
		var listDigest = list.length + "|" + Std.string(list.first()) + "|" + Std.string(list.last()) + "|" + Std.string(popValue);

		var map:Map<String, Int> = ["aa" => seed, "bbbb" => seed + 3, "ccc" => seed + 7];
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

		return listDigest + "|" + mapCount + "|" + mapKeyLen + "|" + mapValueSum;
	}

	static function digestStringIterators():String {
		var ascii = "haxe";
		var asciiIter:StringIterator = new StringIterator(ascii);
		var asciiSum = 0;
		while (asciiIter.hasNext()) {
			asciiSum += asciiIter.next();
		}

		var asciiKv:StringKeyValueIterator = new StringKeyValueIterator(ascii);
		var asciiKeySum = 0;
		var asciiValueSum = 0;
		while (asciiKv.hasNext()) {
			var entry = asciiKv.next();
			asciiKeySum += entry.key;
			asciiValueSum += entry.value;
		}

		var unicode = "A☺B";
		var unicodeIter:StringIteratorUnicode = new StringIteratorUnicode(unicode);
		var unicodeCount = 0;
		var unicodeSum = 0;
		while (unicodeIter.hasNext()) {
			unicodeCount++;
			unicodeSum += unicodeIter.next();
		}

		var unicodeKv:StringKeyValueIteratorUnicode = new StringKeyValueIteratorUnicode(unicode);
		var unicodeKeySum = 0;
		var unicodeValueSum = 0;
		while (unicodeKv.hasNext()) {
			var entry = unicodeKv.next();
			unicodeKeySum += entry.key;
			unicodeValueSum += entry.value;
		}

		return asciiSum
			+ "|"
			+ asciiKeySum
			+ "|"
			+ asciiValueSum
			+ "|"
			+ unicodeCount
			+ "|"
			+ unicodeSum
			+ "|"
			+ unicodeKeySum
			+ "|"
			+ unicodeValueSum;
	}

	static function digestRestIterators(...args:Int):String {
		var rest:haxe.Rest<Int> = args;
		var iter:haxe.iterators.RestIterator<Int> = rest.iterator();
		var count = 0;
		var sum = 0;
		while (iter.hasNext()) {
			count++;
			sum += iter.next();
		}

		var kv:haxe.iterators.RestKeyValueIterator<Int> = rest.keyValueIterator();
		var keySum = 0;
		var valueSum = 0;
		while (kv.hasNext()) {
			var entry = kv.next();
			keySum += entry.key;
			valueSum += entry.value;
		}

		return count + "|" + sum + "|" + keySum + "|" + valueSum;
	}

	static function main() {
		Sys.println("iter=" + digestIterators(10));
		Sys.println("listmap=" + digestListAndMap(5));
		Sys.println("string=" + digestStringIterators());
		Sys.println("rest=" + digestRestIterators(2, 4, 6, 8));
	}
}
