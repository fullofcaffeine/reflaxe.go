interface HasNext {
	public function next():Int;
}

class Cursor implements HasNext {
	public function new() {}

	public function next():Int
		return 3;
}

class Main {
	static function read<T:{function next():Int;}>(value:T):Int
		return value.next();

	static function main() {
		var source = new Cursor();
		var nominal:HasNext = source;
		var narrow:{function next():Int;} = nominal;
		Sys.println(narrow.next());
		Sys.println(read({next: () -> 8}));
		Sys.println(read(source));
	}
}
