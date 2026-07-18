class Base {
	public function new() {}

	public function who():String {
		return "base";
	}

	public function callWho():String {
		return who();
	}
}

class Middle extends Base {
	public function new() {
		super();
	}

	override public function who():String {
		return "middle";
	}
}

class Leaf extends Middle {
	public var constructedAs:String;

	public function new() {
		super();
		constructedAs = callWho();
	}

	override public function who():String {
		return "leaf";
	}
}

class Main {
	static function main() {
		var leaf = new Leaf();
		var middle:Middle = leaf;
		var base:Base = leaf;
		var boundBaseMethod = base.callWho;

		Sys.println(leaf.constructedAs);
		Sys.println(leaf.who());
		Sys.println(middle.who());
		Sys.println(base.who());
		Sys.println(base.callWho());
		Sys.println(boundBaseMethod());
	}
}
