import Lambda;

class CounterIterator {
	final data:Array<Int>;
	var index:Int;

	public function new(data:Array<Int>) {
		this.data = data;
		this.index = 0;
	}

	public function hasNext():Bool {
		return index < data.length;
	}

	public function next():Int {
		return data[index++];
	}
}

class CounterIterable {
	final data:Array<Int>;

	public function new(data:Array<Int>) {
		this.data = data;
	}

	public function iterator():CounterIterator {
		return new CounterIterator(data);
	}
}

class Main {
	static function main() {
		var values:Iterable<Int> = new CounterIterable([1, 2, 3]);
		var emptyValues:Iterable<Int> = new CounterIterable([]);

		Sys.println("count.values=" + Lambda.count(values));
		Sys.println("count.empty=" + Lambda.count(emptyValues));
		Sys.println("empty.values=" + Lambda.empty(values));
		Sys.println("empty.empty=" + Lambda.empty(emptyValues));
	}
}
