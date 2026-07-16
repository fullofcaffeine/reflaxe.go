import haxe.iterators.ArrayIterator;
import haxe.iterators.MapKeyValueIterator;

/**
	What: A user-defined generic iterator used by the structural assignment contract.
	Why: Built-in iterators alone would not prove that the adapter follows typed method
	metadata instead of recognizing a closed set of stdlib class names.
	How: Retain the ordinary Haxe `hasNext` / `next` shape over a shared array cursor.
**/
class GenericIterator<T> {
	final first:T;
	final second:T;
	var index:Int;

	public function new(first:T, second:T) {
		this.first = first;
		this.second = second;
		this.index = 0;
	}

	public function hasNext():Bool {
		return index < 2;
	}

	public function next():T {
		return index++ == 0 ? first : second;
	}
}

class BaseStringIterator {
	public final values:Array<String>;
	public var index:Int;

	public function new(values:Array<String>) {
		this.values = values;
		this.index = 0;
	}

	public function hasNext():Bool {
		return index < values.length;
	}

	public function next():String {
		return "base:" + values[index++];
	}
}

/**
	What: An overriding iterator observed through its generated base-class carrier.
	Why: A structural adapter that captures the base method directly would silently
	bypass Haxe virtual dispatch.
	How: Override `next` and assign the instance to `BaseStringIterator` before the
	structural `Iterator<String>` conversion.
**/
class SpecializedStringIterator extends BaseStringIterator {
	public function new(values:Array<String>) {
		super(values);
	}

	override public function next():String {
		return "special:" + values[index++];
	}
}

class Main {
	static function collectInts(iterator:Iterator<Int>):String {
		var values = [];
		while (iterator.hasNext()) {
			values.push(iterator.next());
		}
		return values.join(",");
	}

	static function collectStrings(iterator:Iterator<String>):String {
		var values = [];
		while (iterator.hasNext()) {
			values.push(iterator.next());
		}
		return values.join(",");
	}

	static function makeGenericIterator():Iterator<String> {
		return new GenericIterator("g1", "g2");
	}

	static function main() {
		var arrayValues = [1, 2, 3];
		var arrayIterator:Iterator<Int> = new ArrayIterator(arrayValues);
		arrayValues[0] = 9;
		Sys.println("array=" + collectInts(arrayIterator));
		Sys.println("arrayArgument=" + collectInts(new ArrayIterator([6, 7])));

		var genericIterator = makeGenericIterator();
		Sys.println("generic=" + collectStrings(genericIterator));
		Sys.println("genericArgument=" + collectStrings(new GenericIterator("c1", "c2")));

		var baseIterator:BaseStringIterator = new SpecializedStringIterator(["a", "b"]);
		var virtualIterator:Iterator<String> = baseIterator;
		Sys.println("virtual=" + collectStrings(virtualIterator));

		var map:Map<String, Int> = [];
		map.set("only", 7);
		var mapIterator:Iterator<{key:String, value:Int}> = new MapKeyValueIterator(map);
		var entry = mapIterator.next();
		Sys.println("map=" + entry.key + ":" + entry.value + ":" + mapIterator.hasNext());
	}
}
