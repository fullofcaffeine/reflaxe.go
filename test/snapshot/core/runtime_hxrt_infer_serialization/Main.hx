class SerializationSnapshotBase {
	var baseValue:Int;

	public function new(baseValue:Int) {
		this.baseValue = baseValue;
	}

	public function readBase():Int {
		return baseValue;
	}

	public function label():String {
		return "base";
	}

	public function dispatch():String {
		return label();
	}
}

class SerializationSnapshotChild extends SerializationSnapshotBase {
	var childValue:String;

	public function new(baseValue:Int, childValue:String) {
		super(baseValue);
		this.childValue = childValue;
	}

	public function readChild():String {
		return childValue;
	}

	override public function label():String {
		return "child:" + childValue;
	}
}

class Main {
	static function main() {
		var encoded = haxe.Serializer.run(new SerializationSnapshotChild(1, "ok"));
		var decoded:SerializationSnapshotChild = cast haxe.Unserializer.run(encoded);
		Sys.println(decoded.readBase() == 1);
		Sys.println(decoded.readChild() == "ok");
		Sys.println(decoded.dispatch() == "child:ok");
	}
}
