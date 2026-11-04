interface Parent<T> {
	public function f(x:T):T;
}

interface Middle<T> extends Parent<T> {}
interface Child extends Middle<Int> {}

class Concrete implements Child {
	public function new() {}

	public function f(x:Int):Int
		return x + 1;
}

class Main {
	static function main() {
		var nominal:Child = null;
		var view:{function f(x:Int):Int;} = nominal;
		Sys.println(view == null);
		var instance:Child = new Concrete();
		var live:{function f(x:Int):Int;} = instance;
		Sys.println(live.f(4));
		var saved = live.f;
		Sys.println(saved(5));
	}
}
