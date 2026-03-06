import haxe.iterators.ArrayIterator;
import haxe.iterators.ArrayKeyValueIterator;
import haxe.iterators.DynamicAccessIterator;
import haxe.iterators.DynamicAccessKeyValueIterator;
import haxe.iterators.StringIterator;
import haxe.iterators.StringIteratorUnicode;
import haxe.iterators.StringKeyValueIterator;
import haxe.iterators.StringKeyValueIteratorUnicode;

class Main {
	static function summarizeArrayIterator(values:Array<Int>):String {
		var iterator:ArrayIterator<Int> = new ArrayIterator(values);
		var count = 0;
		var valueSum = 0;
		var weighted = 0;
		while (iterator.hasNext()) {
			var value = iterator.next();
			count++;
			valueSum += value;
			weighted += value * count;
		}
		return count + "|" + valueSum + "|" + weighted;
	}

	static function summarizeArrayKeyValueIterator(values:Array<Int>):String {
		var iterator:ArrayKeyValueIterator<Int> = new ArrayKeyValueIterator(values);
		var count = 0;
		var keySum = 0;
		var valueSum = 0;
		var keyValueMix = 0;
		while (iterator.hasNext()) {
			var entry = iterator.next();
			var key:Int = cast entry.key;
			var value:Int = cast entry.value;
			count++;
			keySum += key;
			valueSum += value;
			keyValueMix += key * value;
		}
		return count + "|" + keySum + "|" + valueSum + "|" + keyValueMix;
	}

	static function summarizeDynamicAccessIterator(values:haxe.DynamicAccess<Int>):String {
		var iterator:DynamicAccessIterator<Int> = new DynamicAccessIterator(values);
		var count = 0;
		var valueSum = 0;
		var squareSum = 0;
		while (iterator.hasNext()) {
			var value = iterator.next();
			count++;
			valueSum += value;
			squareSum += value * value;
		}
		return count + "|" + valueSum + "|" + squareSum;
	}

	static function summarizeDynamicAccessKeyValueIterator(values:haxe.DynamicAccess<Int>):String {
		var iterator:DynamicAccessKeyValueIterator<Int> = new DynamicAccessKeyValueIterator(values);
		var count = 0;
		var keyLengthSum = 0;
		var valueSum = 0;
		var keyHeadCodeSum = 0;
		while (iterator.hasNext()) {
			var entry = iterator.next();
			var value:Int = cast entry.value;
			count++;
			keyLengthSum += entry.key.length;
			valueSum += value;
			keyHeadCodeSum += entry.key.charCodeAt(0);
		}
		return count + "|" + keyLengthSum + "|" + valueSum + "|" + keyHeadCodeSum;
	}

	static function summarizeRestIterator(...args:Int):String {
		var rest:haxe.Rest<Int> = args;
		var iterator:haxe.iterators.RestIterator<Int> = rest.iterator();
		var count = 0;
		var valueSum = 0;
		var weighted = 0;
		while (iterator.hasNext()) {
			var value = iterator.next();
			count++;
			valueSum += value;
			weighted += value * count;
		}
		return count + "|" + valueSum + "|" + weighted;
	}

	static function summarizeRestKeyValueIterator(...args:Int):String {
		var rest:haxe.Rest<Int> = args;
		var iterator:haxe.iterators.RestKeyValueIterator<Int> = rest.keyValueIterator();
		var count = 0;
		var keySum = 0;
		var valueSum = 0;
		var keyValueMix = 0;
		while (iterator.hasNext()) {
			var entry = iterator.next();
			var key:Int = cast entry.key;
			var value:Int = cast entry.value;
			count++;
			keySum += key;
			valueSum += value;
			keyValueMix += key * value;
		}
		return count + "|" + keySum + "|" + valueSum + "|" + keyValueMix;
	}

	static function summarizeStringIterator(value:String):String {
		var iterator:StringIterator = new StringIterator(value);
		var count = 0;
		var valueSum = 0;
		while (iterator.hasNext()) {
			count++;
			valueSum += iterator.next();
		}
		return count + "|" + valueSum;
	}

	static function summarizeStringKeyValueIterator(value:String):String {
		var iterator:StringKeyValueIterator = new StringKeyValueIterator(value);
		var count = 0;
		var keySum = 0;
		var valueSum = 0;
		var keyValueMix = 0;
		while (iterator.hasNext()) {
			var entry = iterator.next();
			count++;
			keySum += entry.key;
			valueSum += entry.value;
			keyValueMix += entry.key * entry.value;
		}
		return count + "|" + keySum + "|" + valueSum + "|" + keyValueMix;
	}

	static function summarizeStringIteratorUnicode(value:String):String {
		var iterator:StringIteratorUnicode = new StringIteratorUnicode(value);
		var count = 0;
		var valueSum = 0;
		while (iterator.hasNext()) {
			count++;
			valueSum += iterator.next();
		}
		return count + "|" + valueSum;
	}

	static function summarizeStringKeyValueIteratorUnicode(value:String):String {
		var iterator:StringKeyValueIteratorUnicode = new StringKeyValueIteratorUnicode(value);
		var count = 0;
		var keySum = 0;
		var valueSum = 0;
		var keyValueMix = 0;
		while (iterator.hasNext()) {
			var entry = iterator.next();
			count++;
			keySum += entry.key;
			valueSum += entry.value;
			keyValueMix += entry.key * entry.value;
		}
		return count + "|" + keySum + "|" + valueSum + "|" + keyValueMix;
	}

	static function main() {
		var arrayValues = [4, 7, 9];
		Sys.println("array.iter=" + summarizeArrayIterator(arrayValues));
		Sys.println("array.kv=" + summarizeArrayKeyValueIterator(arrayValues));

		var dynamicValues:haxe.DynamicAccess<Int> = {};
		dynamicValues.set("alpha", 2);
		dynamicValues.set("beta", 5);
		dynamicValues.set("gamma", 11);
		Sys.println("dynamic.iter=" + summarizeDynamicAccessIterator(dynamicValues));
		Sys.println("dynamic.kv=" + summarizeDynamicAccessKeyValueIterator(dynamicValues));

		Sys.println("rest.iter=" + summarizeRestIterator(3, 1, 4, 1, 5));
		Sys.println("rest.kv=" + summarizeRestKeyValueIterator(3, 1, 4, 1, 5));

		var ascii = "haxe";
		Sys.println("string.iter=" + summarizeStringIterator(ascii));
		Sys.println("string.kv=" + summarizeStringKeyValueIterator(ascii));

		var unicode = "A☺B";
		Sys.println("string.unicode.iter=" + summarizeStringIteratorUnicode(unicode));
		Sys.println("string.unicode.kv=" + summarizeStringKeyValueIteratorUnicode(unicode));
	}
}
