class ConstructorIntIterator {
	public function new() {}

	public function hasNext():Bool {
		return false;
	}

	public function next():Int {
		return 0;
	}
}

class GenericConsumer<T> {
	public function new(iterator:Iterator<T>) {}
}

class Main {
	static function main() {
		new GenericConsumer<String>(new ConstructorIntIterator());
	}
}
