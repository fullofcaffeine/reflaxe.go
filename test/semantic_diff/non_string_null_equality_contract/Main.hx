class Node {
	public var id:Int;

	public function new(id:Int) {
		this.id = id;
	}
}

class Holder {
	public var node:Node;

	public function new() {
		node = null;
	}
}

class Main {
	static function emit(label:String, value:Bool):Void {
		Sys.println(label + "=" + value);
	}

	static function main() {
		var missing:Node = null;
		emit("node.eqNull", missing == null);
		emit("node.neNull", missing != null);

		var present:Node = new Node(7);
		emit("present.eqNull", present == null);
		emit("present.neNull", present != null);

		var dynNil:Dynamic = null;
		emit("dyn.eqNull", dynNil == null);
		emit("dyn.neNull", dynNil != null);

		var holder = new Holder();
		emit("field.eqNull", holder.node == null);
		holder.node = new Node(9);
		emit("field.eqNull.afterSet", holder.node == null);
		Sys.println("field.id=" + (holder.node == null ? -1 : holder.node.id));

		var maybeString:String = null;
		emit("string.eqNull", maybeString == null);
		emit("string.neNull", maybeString != null);
	}
}
