class InlineGenericIterator<T> {
	final first:T;
	final second:T;
	var index = 0;

	public function new(first:T, second:T) {
		Main.note("generic:new");
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

class InlineBaseStringIterator {
	public final values:Array<String>;
	public var index = 0;

	public function new(values:Array<String>) {
		Main.note("virtual:new");
		this.values = values;
	}

	public function hasNext():Bool {
		return index < values.length;
	}

	public function next():String {
		return "base:" + values[index++];
	}
}

class InlineSpecializedStringIterator extends InlineBaseStringIterator {
	public function new(values:Array<String>) {
		super(values);
	}

	override public function next():String {
		return "special:" + values[index++];
	}
}

/**
	What: Exercises effectful inline blocks that end in user concrete iterators.
	Why: The inline block has an anonymous Iterator result type even though its tail
	still carries the concrete class authority needed by the Go structural adapter.
	How: Pin effect-before-construction order, generic result recovery, and inherited
	virtual dispatch in both declaration and ordinary-call argument contexts.
**/
class Main {
	static final events:Array<String> = [];

	public static function note(event:String) {
		events.push(event);
	}

	static inline function checkedGeneric<T>(first:T, second:T):Iterator<T> {
		note("generic:effect");
		return new InlineGenericIterator(first, second);
	}

	static inline function checkedVirtual(values:Array<String>):Iterator<String> {
		note("virtual:effect");
		var iterator:InlineBaseStringIterator = new InlineSpecializedStringIterator(values);
		return iterator;
	}

	static function collect(iterator:Iterator<String>):String {
		var values = [];
		while (iterator.hasNext()) {
			values.push(iterator.next());
		}
		return values.join(",");
	}

	static function main() {
		var generic:Iterator<String> = checkedGeneric("g1", "g2");
		Sys.println("events-generic=" + events.join(","));
		Sys.println("generic=" + collect(generic));
		Sys.println("virtual=" + collect(checkedVirtual(["v"])));
		Sys.println("events-final=" + events.join(","));
	}
}
