class ConcreteIterator {
	public final values:Array<String>;
	public var index:Int;

	public function new(values:Array<String>) {
		this.values = values;
		this.index = 0;
	}

	public function hasNext():Bool {
		return index < values.length;
	}

	public function next():String {
		return "base:" + values[index++];
	}
}

class SpecializedIterator extends ConcreteIterator {
	public function new(values:Array<String>) {
		super(values);
	}

	override public function next():String {
		return "special:" + values[index++];
	}
}

class ConcreteIterable {
	final values:Array<String>;

	public function new(values:Array<String>) {
		this.values = values;
	}

	public function iterator():SpecializedIterator {
		return new SpecializedIterator(values);
	}
}

interface NameContract {
	function macroName(resolve:String->Dynamic, ignored:Dynamic):String;
}

class MethodBase implements NameContract {
	var total:Int;

	public function new() {
		total = 0;
	}

	public function virtualName():String {
		return "base";
	}

	public function macroName(resolve:String->Dynamic, ignored:Dynamic):String {
		return virtualName();
	}

	public function describe(resolve:String->Dynamic, key:String):String {
		return Std.string(resolve(key)) + ":" + virtualName();
	}

	public function bump(resolve:String->Dynamic, key:String):Int {
		var amount:Int = resolve(key);
		total += amount;
		return total;
	}

	public function type(resolve:String->Dynamic, ignored:Dynamic):String {
		return "keyword";
	}

	private function secret(resolve:String->Dynamic, ignored:Dynamic):String {
		return "secret";
	}

	public function retainSecret():String {
		return secret(null, null);
	}
}

class MethodMiddle extends MethodBase {
	public function new() {
		super();
	}

	public function middleOnly(resolve:String->Dynamic, ignored:Dynamic):String {
		return "middle";
	}
}

class MethodLeaf extends MethodMiddle {
	public function new() {
		super();
	}

	override public function virtualName():String {
		return "leaf";
	}

	public function leafOnly(resolve:String->Dynamic, ignored:Dynamic):String {
		return "leaf-only";
	}
}

class Main {
	static function call0(obj:Dynamic, key:String):String {
		var template = new haxe.Template("$$invoke(dummy)");
		return template.execute({dummy: null}, {invoke: Reflect.field(obj, key)});
	}

	static function call1(obj:Dynamic, key:String, value:Dynamic):String {
		var template = new haxe.Template("$$invoke(value)");
		return template.execute({value: value}, {invoke: Reflect.field(obj, key)});
	}

	static function main() {
		var template = new haxe.Template("::foreach items::::__current__::;::end::");
		Sys.println(template.execute({items: new ConcreteIterable(["a", "b"])}));
		Sys.println(template.execute({items: new ConcreteIterator(["x", "y"])}));

		var leaf = new MethodLeaf();
		var middle:MethodMiddle = leaf;
		var base:MethodBase = leaf;
		var named:NameContract = leaf;

		Sys.println(Reflect.hasField(leaf, "leaf" + "Only"));
		Sys.println(Reflect.field(leaf, "leaf" + "Only") != null);
		Sys.println(call0(leaf, "leaf" + "Only"));
		Sys.println(Reflect.hasField(base, "leafOnly"));
		Sys.println(Reflect.field(base, "leafOnly") != null);
		Sys.println(call0(base, "leafOnly"));
		Sys.println(call0(middle, "middleOnly"));
		Sys.println(call0(leaf, "macroName")
			+ ":"
			+ call0(middle, "macroName")
			+ ":"
			+ call0(base, "macroName")
			+ ":"
			+ call0(named, "macroName"));
		Sys.println(call1(base, "describe", "bound"));
		Sys.println(call0(base, "type"));
		Sys.println(leaf.retainSecret());
		Sys.println(call0(base, "secret"));
		Sys.println(call1(base, "bump", 2));
		Sys.println(call1(base, "bump", 3));
		Sys.println(Reflect.hasField(base, "missing"));
		Sys.println(Std.string(Reflect.field(base, "missing")));

		var absent:MethodBase = null;
		Sys.println(Reflect.hasField(absent, "virtualName"));
		Sys.println(Std.string(Reflect.field(absent, "virtualName")));
	}
}
