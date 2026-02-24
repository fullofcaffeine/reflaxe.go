class Node {
	public var id:Int;

	public function new(id:Int) {
		this.id = id;
	}
}

class Main {
	static function main() {
		var node:Node = null;
		var dyn:Dynamic = node;

		Sys.println(Std.string(node));
		Sys.println("" + node);
		Sys.println(dyn == null);
	}
}
