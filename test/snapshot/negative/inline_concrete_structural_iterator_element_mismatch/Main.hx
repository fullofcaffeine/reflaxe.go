class InlineMismatchIntIterator {
	public function new() {}

	public function hasNext():Bool {
		return false;
	}

	public function next():Int {
		return 0;
	}
}

class Main {
	static inline function intIterator():Iterator<Int> {
		return new InlineMismatchIntIterator();
	}

	static function collect(iterator:Iterator<String>) {}

	static function main() {
		collect(intIterator());
	}
}
