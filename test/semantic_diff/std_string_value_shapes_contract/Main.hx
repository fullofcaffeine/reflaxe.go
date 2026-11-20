import shapes.Marker;

enum Opt<T> {
	None;
	Some(v:T);
}

enum Shape {
	Circle(r:Float);
	Rect(w:Int, h:Int);
	Label(text:String, inner:Opt<Int>);
	Flag(b:Bool);
	Holder(o:Dynamic);
	Many(xs:Array<Int>);
}

enum Chain {
	End;
	Link(next:Chain);
}

class Point {
	public var x:Int;
	public var y:Int;

	public function new(x:Int, y:Int) {
		this.x = x;
		this.y = y;
	}
}

class SubPoint extends Point {
	public function new() {
		super(0, 0);
	}
}

class LabeledPoint extends Point {
	public function new() {
		super(5, 6);
	}

	public function toString():String {
		return "LabeledPoint(" + x + "," + y + ")";
	}
}

class Named {
	var name:String;

	public function new(name:String) {
		this.name = name;
	}

	public function toString():String {
		return "Named(" + name + ")";
	}
}

class SubNamed extends Named {
	public function new(name:String) {
		super(name);
	}
}

private class Hidden {
	public function new() {}
}

/**
	Std.string must print enums, anonymous objects, and class instances the way
	Haxe does. The expected text comes from `haxe --interp`:
	- enum values print as `Name` or `Name(arg,arg)`;
	- anonymous objects print as `{name: value, ...}` in the interpreter's field order;
	- class instances call `toString()` when they have one, else print the class path;
	- anything nested more than five levels deep prints as `<...>`.
**/
class Main {
	static function show(label:String, value:Dynamic):Void {
		Sys.println(label + "=" + Std.string(value));
	}

	static function main() {
		var none:Opt<Int> = None;
		var some:Opt<Int> = Some(1);
		Sys.println("enum.none=" + Std.string(none));
		Sys.println("enum.some=" + Std.string(some));
		show("enum.nested", Some(Some(1)));
		show("enum.string", Some("hello"));
		show("enum.float", Circle(1.5));
		show("enum.two", Rect(2, 3));
		show("enum.mixed", Label("x", Some(4)));
		show("enum.bool", Flag(true));
		show("enum.null", Holder(null));
		show("enum.object", Holder({a: 1}));
		show("enum.array", Many([1, 2]));
		show("enum.option", haxe.ds.Option.Some(3));
		show("enum.option.none", haxe.ds.Option.None);
		var missing:Null<Opt<Int>> = null;
		show("enum.typed.null", missing);

		show("anon.two", {a: 1, b: "hello"});
		show("anon.order", {b: 1, a: 2});
		show("anon.order3", {zeta: 1, alpha: 2, mid: 3});
		show("anon.empty", {});
		show("anon.nested", {
			inner: {x: 1, y: [1, 2]},
			e: Some(1),
			s: "q",
			f: 1.5,
			nul: null
		});
		show("anon.quoted", {"quoted key": 1, normal: 2});
		var grown:Dynamic = {x: 1};
		Reflect.setField(grown, "added", 2);
		show("anon.added", grown);
		Reflect.deleteField(grown, "x");
		show("anon.deleted", grown);
		show("anon.tostring", {toString: function() return "custom"});
		show("anon.objects", {
			e: None,
			arr: [None, Some(2)],
			obj: new Marker(),
			named: new Named("in")
		});

		show("class.plain", new Point(1, 2));
		show("class.sub", new SubPoint());
		show("class.package", new Marker());
		show("class.private", new Hidden());
		show("class.tostring", new Named("n"));
		show("class.inherited.tostring", new SubNamed("m"));
		show("class.sub.tostring", new LabeledPoint());
		var asBase:Point = new SubPoint();
		show("class.base.view", asBase);

		show("array.enums", [Some(1), None]);
		show("array.objects", [{a: 1}]);
		show("array.classes", [new Point(1, 2)]);
		show("array.tostring", [new Named("z")]);
		var list:Array<Opt<Int>> = [Some(1), None];
		Sys.println("array.toString=" + list.toString());
		Sys.println("array.concat=" + list);

		var e:Opt<Int> = Some(5);
		Sys.println('' + e);
		Sys.println('${e}');
		Sys.println("concat.enum=" + e);
		var o = {k: "v"};
		Sys.println("concat.anon=" + o);
		Sys.println('interp.anon=${o}');
		var p = new Point(3, 4);
		Sys.println("concat.class=" + p);
		Sys.println('interp.class=${p}');
		var d:Dynamic = Some(7);
		Sys.println(d);
		var byKey = new Map<String, Opt<Int>>();
		byKey.set("a", Some(1));
		Sys.println("map=" + byKey.toString());

		show("depth.array5", [[[[[1]]]]]);
		show("depth.array6", [[[[[[1]]]]]]);
		show("depth.enum5", Link(Link(Link(Link(Link(End))))));
		show("depth.enum6", Link(Link(Link(Link(Link(Link(End)))))));
		show("depth.anon5", {a: {b: {c: {d: {e: 1}}}}});
		show("depth.anon6", {a: {b: {c: {d: {e: {f: 1}}}}}});
		show("depth.tostring4", [[[[new Named("deep")]]]]);
		show("depth.tostring5", [[[[[new Named("deep")]]]]]);
		show("depth.string5", [[[[["s"]]]]]);
		var cyclic:Array<Dynamic> = [];
		cyclic.push(cyclic);
		show("depth.cyclic.array", cyclic);
		var loop:Dynamic = {};
		Reflect.setField(loop, "self", loop);
		show("depth.cyclic.anon", loop);
	}
}
