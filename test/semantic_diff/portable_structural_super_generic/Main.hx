class Base<T> {
	var it:Iterator<T>;

	public function new(it:Iterator<T>) {
		this.it = it;
	}

	public function read():T
		return it.next();
}

class Child extends Base<Int> {
	public function new(it:Iterator<Int>) {
		super(it);
	}
}

class Main {
	static function main() {
		var source = {hasNext: () -> true, next: () -> 3};
		var child = new Child(source);
		Sys.println(child.read());
		source.next = () -> 9;
		Sys.println(child.read());
	}
}
