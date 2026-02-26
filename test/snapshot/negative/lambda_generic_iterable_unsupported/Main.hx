import Lambda;

class CounterIterable {
	final data:Array<Int>;

	public function new(data:Array<Int>) {
		this.data = data;
	}

	public function iterator():Iterator<Int> {
		return data.iterator();
	}
}

class Main {
	static function main() {
		var source:Iterable<Int> = new CounterIterable([1, 2, 3]);
		Sys.println(Lambda.count(source));
	}
}
