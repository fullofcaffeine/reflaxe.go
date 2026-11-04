class Box<T> {
	public function new() {}

	public function apply(f:T->T, value:T):T
		return f(value);
}

typedef Applied = {function apply(f:Int->Int, value:Int):Int;}

class Main {
	static function main() {
		var view:Applied = new Box<Int>();
		Sys.println(view.apply(x -> x + 1, 2));
	}
}
