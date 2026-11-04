class Leaf {
	public function new() {}
}

class Box<T> {
	public final value:T;

	public function new(value:T) {
		this.value = value;
	}
}

class Main {
	static function main() {
		final absent:Leaf = null;
		final empty = new Box<Leaf>(absent);
		final present = new Box<Leaf>(new Leaf());
		Sys.println(empty.value == null);
		Sys.println(null == empty.value);
		Sys.println(empty.value != null);
		Sys.println(null != empty.value);
		Sys.println(present.value == null);
		Sys.println(present.value != null);
	}
}
