class IntIterator {
	public function new() {}

	public function hasNext():Bool {
		return false;
	}

	public function next():Int {
		return 0;
	}
}

class Main {
	static function main() {
		var iterator:Iterator<String> = new IntIterator();
		Sys.println(iterator.hasNext());
	}
}
