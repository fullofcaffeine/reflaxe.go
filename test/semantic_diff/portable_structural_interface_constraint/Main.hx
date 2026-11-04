interface HasNext {
	public function next():Int;
}

class Cursor implements HasNext {
	public function new() {}

	public function next():Int
		return 3;
}

class MutableCursor {
	public function new() {}

	public dynamic function next():Int
		return 11;
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
		var erased:Dynamic = source;
		var recovered:{function next():Int;} = erased;
		Sys.println(recovered.next());
		Sys.println((cast recovered : Dynamic) == erased);
		var missing:Dynamic = null;
		var nullView:{function next():Int;} = missing;
		Sys.println(nullView == null);
		var mutable = new MutableCursor();
		var mutableErased:Dynamic = mutable;
		var live:{function next():Int;} = mutableErased;
		var saved = live.next;
		mutable.next = () -> 19;
		Sys.println(saved());
		Sys.println(live.next());
		Sys.println((cast live : Dynamic) == mutableErased);
	}
}
