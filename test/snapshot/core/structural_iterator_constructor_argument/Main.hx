class SnapshotGenericIterator<T> {
	final first:T;
	final second:T;
	var index = 0;

	public function new(first:T, second:T) {
		this.first = first;
		this.second = second;
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
	public var index = 0;

	public function new(values:Array<String>) {
		this.values = values;
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

class SnapshotIntConsumer {
	final iterator:Iterator<Int>;

	public function new(before:String, iterator:Iterator<Int>, after:String) {
		this.iterator = iterator;
	}

	public function collect():String {
		var values = [];
		while (iterator.hasNext()) {
			values.push(iterator.next());
		}
		return values.join(",");
	}
}

class SnapshotGenericConsumer<T> {
	final iterator:Iterator<T>;

	public function new(iterator:Iterator<T>) {
		this.iterator = iterator;
	}

	public function collect():String {
		var values = [];
		while (iterator.hasNext()) {
			values.push(iterator.next());
		}
		return values.join(",");
	}
}

/** Pins constructor argument order and structural iterator adapter output. **/
class Main {
	static final events:Array<String> = [];

	static function mark(label:String):String {
		events.push(label);
		return label;
	}

	static inline function checkedIterator(values:Array<Int>):Iterator<Int> {
		events.push("iterator");
		return values.iterator();
	}

	static function main() {
		var arrayValues = [1, 2];
		var arrayConsumer = new SnapshotIntConsumer(mark("before"), checkedIterator(arrayValues), mark("after"));
		arrayValues[0] = 9;
		Sys.println("order=" + events.join(","));
		Sys.println("array=" + arrayConsumer.collect());

		var genericConsumer = new SnapshotGenericConsumer<String>(new SnapshotGenericIterator("g1", "g2"));
		Sys.println("generic=" + genericConsumer.collect());

		var baseIterator:SnapshotBaseIterator = new SnapshotSpecializedIterator(["v"]);
		var virtualConsumer = new SnapshotGenericConsumer<String>(baseIterator);
		Sys.println("virtual=" + virtualConsumer.collect());
	}
}
