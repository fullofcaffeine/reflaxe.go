typedef Point = {
	var x:Int;
	var label:String;
}

class Base {
	public var inherited:String = "base";

	public function new() {}

	public function describe():String {
		return "describe " + inherited;
	}
}

class Box extends Base {
	public var value:Int;
	public var computed(get, never):Int;

	var hidden:String = "hidden";

	public function new(value:Int) {
		super();
		this.value = value;
	}

	function get_computed():Int {
		return value * 2;
	}

	public function add(n:Int):Int {
		return value + n;
	}

	public dynamic function scale(n:Int):Int {
		return n * value;
	}
}

/**
	Field reads and writes on a `Dynamic` value (haxe_go-lkgv, haxe_go-l0gk).

	Haxe reads `o.f` on a `Dynamic` receiver as `Reflect.field(o, "f")` and
	writes `o.f = v` as `Reflect.setField(o, "f", v)`. haxe.go emitted a Go
	selector on an `any` value, which does not build (`o.to undefined (type
	any has no field or method to)`), for anonymous objects, class instances,
	missing fields, nested reads, writes and calls in every context.

	As in `--interp`: a missing field reads as null, also on a null receiver;
	no property accessor runs, so a property without storage reads as null;
	a method read as a value stays bound to its object and sees later
	writes; a write to a missing field of an anonymous object adds the
	field. The receiver runs once, before the right side. Numeric `+=` and
	`++` on a `Dynamic` field are covered by `core/dynamic_operators`. Left
	out: `@:isVar` properties with accessors (`Reflect.field` does not see
	their storage on haxe.go yet).
**/
class Main {
	static var order = "";

	static function echo(value:Dynamic):Dynamic {
		return value;
	}

	static function step(name:String, value:Dynamic):Dynamic {
		order += name;
		return value;
	}

	static function stepText(name:String, value:String):String {
		order += name;
		return value;
	}

	static function main() {
		anonymousReads();
		builtAndTypedObjects();
		nestedReads();
		classInstanceReads();
		writes();
		calls();
		evaluationOrder();
	}

	static function anonymousReads() {
		var o:Dynamic = {to: 2, name: "anon", flag: true};
		Sys.println("anon " + o.to + " " + o.name);
		Sys.println("arg " + echo(o.name));
		if (o.flag)
			Sys.println("cond yes");
		Sys.println("ternary " + (o.flag ? "on" : "off"));
		Sys.println('interp to=${o.to} name=${o.name}');
		Sys.println("eq " + (o.to == 2) + " " + (o.name == "anon") + " " + (o.to != 3));
		var typed:Int = o.to;
		Sys.println("typed " + (typed + 1));
		var text:String = o.name;
		Sys.println("text " + text.length);
		Sys.println("missing " + o.missing + " " + (o.missing == null));
		var nothing:Dynamic = null;
		Sys.println("null receiver " + nothing.field);
	}

	static function builtAndTypedObjects() {
		var built:Dynamic = {};
		Reflect.setField(built, "a", 1);
		Sys.println("reflect built " + built.a);
		var point:Point = {x: 1, label: "p"};
		var view:Dynamic = point;
		Sys.println("typedef " + view.x + " " + view.label);
		view.x = 5;
		Sys.println("typedef shared " + point.x);
	}

	static function nestedReads() {
		var tree:Dynamic = {a: {b: 3}, list: [1, 2]};
		Sys.println("nested " + tree.a.b + " " + tree.a.missing + " " + tree.list);
		tree.a.b = 4;
		Sys.println("nested write " + tree.a.b);
	}

	static function classInstanceReads() {
		var box = new Box(4);
		var dynamicBox:Dynamic = box;
		Sys.println("field " + dynamicBox.value + " inherited " + dynamicBox.inherited + " private " + dynamicBox.hidden);
		Sys.println("property without storage " + dynamicBox.computed + " typed " + box.computed);
		Sys.println("class missing " + dynamicBox.missing);
		var bound = dynamicBox.add;
		var first:Int = bound(1);
		var dynamicBound:Dynamic = dynamicBox.add;
		box.value = 10;
		var second:Int = bound(1);
		Sys.println("bound " + first + " live " + second + " dynamic " + dynamicBound(2));
	}

	static function writes() {
		var o:Dynamic = {to: 2, name: "anon"};
		o.to = 9;
		o.extra = "added";
		Sys.println("anon write " + o.to + " " + o.extra + " " + Reflect.hasField(o, "extra"));
		o.name += "!";
		Sys.println("concat " + o.name);
		var assigned = (o.to = 11);
		Sys.println("assigned " + assigned + " " + o.to);
		var box = new Box(1);
		var dynamicBox:Dynamic = box;
		dynamicBox.value = 8;
		dynamicBox.inherited = "changed";
		Sys.println("class write " + box.value + " " + box.inherited);
	}

	static function calls() {
		var functions:Dynamic = {f: (x:Int) -> x + 100};
		Sys.println("anon call " + functions.f(1));
		var box:Dynamic = new Box(3);
		Sys.println("method call " + box.add(2) + " " + box.describe());
		Sys.println("dynamic method " + box.scale(5));
	}

	static function evaluationOrder() {
		var o:Dynamic = {name: "n", count: 0};
		step("t", o).name = step("v", "w");
		Sys.println("write order " + order + " " + o.name);
		order = "";
		step("t", o).name += stepText("v", "+");
		Sys.println("concat order " + order + " " + o.name);
		order = "";
		Sys.println("read order " + step("t", o).name + step("v", "") + " " + order);
	}
}
