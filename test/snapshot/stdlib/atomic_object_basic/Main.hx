import haxe.atomic.AtomicObject;

class Node {
	public var id:String;

	public function new(id:String) {
		this.id = id;
	}
}

class Main {
	static function out(label:String, value:Dynamic):Void {
		Sys.println(label + "=" + Std.string(value));
	}

	static function nodeId(value:Node):String {
		return value == null ? "null" : value.id;
	}

	static function main() {
		var a = new Node("a");
		var b = new Node("a");
		var c = new Node("c");
		var d = new Node("d");

		var atom = new AtomicObject<Node>(a);
		out("load.0", nodeId(atom.load()));

		var oldMiss = atom.compareExchange(b, c);
		out("cmp.miss.old", nodeId(oldMiss));
		out("cmp.miss.now", nodeId(atom.load()));

		var oldHit = atom.compareExchange(a, c);
		out("cmp.hit.old", nodeId(oldHit));
		out("cmp.hit.now", nodeId(atom.load()));

		var oldExchange = atom.exchange(d);
		out("xchg.old", nodeId(oldExchange));
		out("xchg.now", nodeId(atom.load()));

		var stored = atom.store(a);
		out("store.ret", nodeId(stored));
		out("store.now", nodeId(atom.load()));

		var alias = atom.load();
		alias.id = "a_mut";
		out("alias.now", nodeId(atom.load()));

		var oldAlias = atom.compareExchange(alias, c);
		out("cmp.alias.old", nodeId(oldAlias));
		out("cmp.alias.now", nodeId(atom.load()));
	}
}
