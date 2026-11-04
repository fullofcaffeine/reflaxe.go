interface Named {
	public function name():String;
}

class Base {
	public var n = 3;

	public function new() {}
}

class Child extends Base implements Named {
	public function new()
		super();

	public function name():String
		return "child";
}

class Grandchild extends Child {
	public function new()
		super();
}

class Main {
	static function call<T>(callback:T->Int, value:T):Int
		return callback(value);

	static function main() {
		var original:Base->Int = value -> value.n;
		Sys.println(call(original, new Child()));
		Sys.println(Reflect.callMethod(null, original, [new Child()]));
		var child = new Grandchild();
		var base:Base = child;
		var read:Base->Int = value -> value == null ? -1 : value.n;
		Sys.println(Reflect.callMethod(null, read, [child]));
		Sys.println(Reflect.callMethod(null, read, [base]));
		Sys.println(Reflect.callMethod(null, read, [null]));
		var name:Named->String = value -> value.name();
		Sys.println(Reflect.callMethod(null, name, [base]));
		var create:Void->Base = () -> child;
		Sys.println((Reflect.callMethod(null, create, []) : Base).n);
	}
}
