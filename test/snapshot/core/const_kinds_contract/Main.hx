class Base {
	public var id:Int;

	public function new(id:Int) {
		this.id = id;
	}

	public function asBase():Base {
		return this;
	}

	public function baseId():Int {
		return this.id;
	}
}

class Child extends Base {
	public function new() {
		super(7);
	}

	public function superId():Int {
		return super.baseId();
	}

	public function superAsBase():Base {
		return super.asBase();
	}
}

class Main {
	static function main() {
		var child = new Child();
		var asBase = child.superAsBase();
		Sys.println(Std.string(child.superId()));
		Sys.println(Std.string(asBase.id));
		Sys.println(Std.string(null));
		Sys.println(Std.string(3));
		Sys.println(Std.string(1.5));
		Sys.println(Std.string(true));
		Sys.println("ok");
	}
}
