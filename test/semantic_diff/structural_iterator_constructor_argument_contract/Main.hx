/** A generic user iterator proving constructor coercion is not stdlib-specific. **/
class GenericIterator<T> {
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

class BaseStringIterator {
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

class SpecializedStringIterator extends BaseStringIterator {
	public function new(values:Array<String>) {
		super(values);
	}

	override public function next():String {
		return "special:" + values[index++];
	}
}

/** Stores an integer iterator after three ordered constructor arguments. **/
class IntConsumer {
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

class GenericConsumer<T> {
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

/**
	What: Covers structural iterator coercion at constructor parameter boundaries.
	Why: Constructor arguments use a distinct lowering path from ordinary calls.
	How: Observe argument order, live array capture, generic result recovery, and
	virtual dispatch through values stored by constructors.
**/
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
		var arrayConsumer = new IntConsumer(mark("before"), checkedIterator(arrayValues), mark("after"));
		arrayValues[0] = 9;
		Sys.println("order=" + events.join(","));
		Sys.println("array=" + arrayConsumer.collect());

		var genericConsumer = new GenericConsumer<String>(new GenericIterator("g1", "g2"));
		Sys.println("generic=" + genericConsumer.collect());

		var baseIterator:BaseStringIterator = new SpecializedStringIterator(["v"]);
		var virtualConsumer = new GenericConsumer<String>(baseIterator);
		Sys.println("virtual=" + virtualConsumer.collect());
	}
}
