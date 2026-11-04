class Holder<T> {
	public var value:Iterator<T>;

	public function new(value:Iterator<T>) {
		this.value = value;
	}

	public function read():Iterator<T>
		return value;

	public function erase():Dynamic
		return value;
}

class Main {
	static function main() {
		var source:Iterator<Int> = {hasNext: () -> true, next: () -> 3};
		var recovered:Iterator<Int> = new Holder<Int>(source).erase();
		Sys.println(recovered.next());
	}
}
