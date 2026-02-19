class Base {
	public function new() {}

	public function who():Int {
		return 1;
	}

	public function callWho():Int {
		return who();
	}
}

class Child extends Base {
	override public function who():Int {
		return 2;
	}

	public function callSuperWho():Int {
		return super.who();
	}
}

class Main {
	static function main() {
		var child = new Child();
		Sys.println(child.callWho());
		Sys.println(child.callSuperWho());
	}
}
