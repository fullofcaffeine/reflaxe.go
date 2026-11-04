class Box<T> {
	public function new() {}

	public function keep(f:T->T):T->T
		return f;
}

typedef Applied = {function keep(f:Int->Int):Int->Int;}

class Main {
	static function main() {
		var view:Applied = new Box<Int>();
		var f = view.keep(x -> x + 1);
		Sys.println(f(2));
	}
}
