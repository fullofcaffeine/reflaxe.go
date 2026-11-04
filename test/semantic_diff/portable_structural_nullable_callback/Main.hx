class Box<T> {
	public function new() {}

	public function read(value:T):Null<Int>
		return null;
}

typedef Reader = {function read(value:Int):Null<Int>;}

class Main {
	static function main() {
		var view:Reader = new Box<Int>();
		Sys.println(view.read(1) == null);
	}
}
