import haxe.iterators.ArrayIterator;

class SnapshotGenericIterator<T> {
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

class SnapshotBaseIterator {
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

class SnapshotSpecializedIterator extends SnapshotBaseIterator {
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
		return new SnapshotGenericIterator("u", "v");
	}

	static function main() {
		var arrayValues = [4, 5];
		var arrayIterator:Iterator<Int> = new ArrayIterator(arrayValues);
		arrayValues[0] = 8;
		Sys.println("array=" + collectInts(arrayIterator));
		Sys.println("arrayArgument=" + collectInts(new ArrayIterator([6, 7])));

		var genericIterator = makeGenericIterator();
		Sys.println("generic=" + collectStrings(genericIterator));
		Sys.println("genericArgument=" + collectStrings(new SnapshotGenericIterator("c", "d")));

		var baseIterator:SnapshotBaseIterator = new SnapshotSpecializedIterator(["z"]);
		var virtualIterator:Iterator<String> = baseIterator;
		Sys.println("virtual=" + collectStrings(virtualIterator));
	}
}
