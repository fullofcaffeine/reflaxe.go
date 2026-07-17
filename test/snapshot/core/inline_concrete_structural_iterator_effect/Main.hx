class SnapshotInlineGenericIterator<T> {
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

class SnapshotInlineBaseStringIterator {
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

class SnapshotInlineSpecializedStringIterator extends SnapshotInlineBaseStringIterator {
	public function new(values:Array<String>) {
		super(values);
	}

	override public function next():String {
		return "special:" + values[index++];
	}
}

/**
	What: Pins generated Go for effectful inline concrete iterator adaptation.
	Why: Prefix effects must stay ordered while the concrete tail keeps enough typed
	authority for the shared structural Iterator carrier.
	How: Exercise a generic constructor and a base-typed subclass tail through both
	a declaration and an ordinary call argument.
**/
class Main {
	static final events:Array<String> = [];

	public static function note(event:String) {
		events.push(event);
	}

	static inline function checkedGeneric<T>(first:T, second:T):Iterator<T> {
		note("generic:effect");
		return new SnapshotInlineGenericIterator(first, second);
	}

	static inline function checkedVirtual(values:Array<String>):Iterator<String> {
		note("virtual:effect");
		var iterator:SnapshotInlineBaseStringIterator = new SnapshotInlineSpecializedStringIterator(values);
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
