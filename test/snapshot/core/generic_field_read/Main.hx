class Box<T> {
	public var value:T;

	public function new(value:T) {
		this.value = value;
	}
}

class Leaf {
	public var id:Int;

	public function new(id:Int) {
		this.id = id;
	}
}

class Main {
	static function main():Void {
		final reference = new Box(new Leaf(7));
		Sys.println(reference.value.id);
		final number = new Box(4);
		Sys.println(number.value + 1);
		final text = new Box("word");
		Sys.println(text.value.toUpperCase());
		final missing:Box<Leaf> = new Box(null);
		final loaded:Leaf = missing.value;
		Sys.println(loaded == null);
		final callable:Box<Int->Int> = new Box(value -> value + 1);
		Sys.println(callable.value(8));
	}
}
