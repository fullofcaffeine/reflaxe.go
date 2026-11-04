/**
	What: Pins method identity across source views and construction.
	Why: Receiver identity alone cannot explain overrides, saved constructor
	bindings, mutable function fields, or reflection. The interpreter supplies
	the expected behavior before the Go representation is changed.
	How: Compare bindings and invoke them after construction and replacement;
	ordinary direct calls remain controls for the generated callable path.
**/
interface MethodView {
	public function add(value:Int):Int;
}

class BaseCounter implements MethodView {
	public final saved:Int->Int;
	public var total:Int = 1;

	public function new() {
		saved = add;
	}

	public function add(value:Int):Int {
		return total + value;
	}

	public dynamic function callback(value:Int):Int {
		return total + value;
	}
}

class DerivedCounter extends BaseCounter {
	public function new() {
		super();
		total = 10;
	}

	override public function add(value:Int):Int {
		return total + value + 100;
	}
}

class Main {
	static function identity(left:Int->Int, right:Int->Int):Bool {
		return Reflect.compareMethods(left, right);
	}

	static function named(value:Int):Int {
		return value + 1;
	}

	static function main() {
		final child = new DerivedCounter();
		final base:BaseCounter = child;
		final view:MethodView = child;
		final direct = child.add;
		final reflected = Reflect.field(child, "add");
		Sys.println("method.views="
			+ identity(direct, base.add)
			+ ":"
			+ identity(direct, view.add)
			+ ":"
			+ Reflect.compareMethods(direct, reflected));
		Sys.println("method.constructor=" + identity(child.saved, direct) + ":" + child.saved(2) + ":" + direct(2));
		final other = new DerivedCounter();
		Sys.println("method.receiver=" + identity(direct, other.add));
		final before = child.callback;
		child.callback = named;
		Sys.println("method.dynamic="
			+ identity(before, child.callback)
			+ ":"
			+ identity(child.callback, named)
			+ ":"
			+ before(2)
			+ ":"
			+ child.callback(2));
		Sys.println("method.static=" + identity(named, named));
		Sys.println("method.equals=" + (named == named) + ":" + (direct == child.add) + ":" + (direct == other.add));
	}
}
