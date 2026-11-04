class Parent extends haxe.Exception {
	public function new() {
		super("parent");
	}
}

class Child extends Parent {
	public function new() {
		super();
	}
}

class Main {
	static var calls:Int = 0;

	static function absent():Child {
		calls++;
		return null;
	}

	static function main():Void {
		var child:Child = absent();
		var parent:Parent = child;
		var base:haxe.Exception = absent();
		Sys.println("parent-null=" + Std.string(parent == null));
		Sys.println("base-null=" + Std.string(base == null));
		Sys.println("calls=" + calls);
	}
}
