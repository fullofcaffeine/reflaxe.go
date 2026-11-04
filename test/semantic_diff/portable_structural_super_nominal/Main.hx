class Cursor {
	public function new() {}

	public function hasNext():Bool
		return true;

	public function next():Int
		return 3;
}

class Base {
	var it:Iterator<Int>;

	public function new(it:Iterator<Int>) {
		this.it = it;
	}

	public function read():Int
		return it.next();
}

class Child extends Base {
	public function new() {
		super(new Cursor());
	}
}

class Main {
	static function main() {
		Sys.println(new Child().read());
	}
}
