class Box {
	public var f:Float;
	public var b:Bool = false;
	public var s:String = "init";
	public var nf:Null<Float> = 0.5;

	public function new(f:Float) {
		this.f = f;
	}
}

/**
	`Dynamic` values stored in Float, Bool and String slots must behave as in
	`--interp` (haxe_go-bme1, haxe_go-h078).

	An Int in a Float slot reads as the same number (`f / 2` is 1.5 for 3) in
	a local, a `Dynamic` field read, an argument, a return value, a class or
	static field, an Array element, a Map value, a structure field, a ternary,
	a cast, an `Array<Dynamic>` element, `Reflect.field` and a `Dynamic` call
	result. A null value stays null in String, `Null<Float>` and `Null<Bool>`
	slots. A value that the slot cannot hold must not stop the program: a Haxe
	`try`/`catch` stays in control (`handled`). Left out: a null value in a
	Float or Bool slot, an Int compared as a String, and Int overflow in
	`f * 1000000000`, where haxe.go follows the static Haxe targets and
	`--interp` keeps the Int.
**/
class Main {
	static var staticFloat:Float = 0.5;
	static var staticString:String = "init";
	static var consumed = "";
	static var always = true;

	static function dyn(value:Dynamic):Dynamic {
		return value;
	}

	static function half(f:Float):String {
		return Std.string(f / 2);
	}

	static function isNull(s:String):Bool {
		return s == null;
	}

	static function yesNo(b:Bool):String {
		return b ? "yes" : "no";
	}

	static function floatOf(d:Dynamic):Float {
		return d;
	}

	static function floatOfCall(d:Dynamic):Float {
		return dyn(d);
	}

	static function boolOf(d:Dynamic):Bool {
		return d;
	}

	static function stringOf(d:Dynamic):String {
		return d;
	}

	static function stringOfCall(d:Dynamic):String {
		return dyn(d);
	}

	static function nullableFloatOf(d:Dynamic):Null<Float> {
		return d;
	}

	static function floatText(f:Float):String {
		return Std.string(f) + " " + Std.string(f / 2) + " " + Std.string(f + 0.25) + " " + Std.string(f == 3);
	}

	static function stringText(s:String):String {
		return Std.string(s == null) + " " + Std.string(s);
	}

	static function main() {
		floatSlots(3);
		floatSlots(2.5);
		boolSlots(true);
		stringSlots(null);
		stringSlots("x");
		otherValuesInStringSlots();
		nullableSlots();
		incompatibleValues();
	}

	static function floatSlots(d:Dynamic) {
		Sys.println("-- Float slots, " + Std.string(d));
		var local:Float = d;
		Sys.println("local " + floatText(local));
		var o:Dynamic = {n: d};
		var field:Float = o.n;
		Sys.println("dynamic field " + floatText(field));
		Sys.println("argument " + half(d));
		Sys.println("return " + floatText(floatOf(d)));
		Sys.println("return of call " + floatText(floatOfCall(d)));
		var box = new Box(d);
		Sys.println("constructor " + floatText(box.f));
		box.f = 0.5;
		box.f = d;
		Sys.println("field write " + floatText(box.f));
		staticFloat = d;
		Sys.println("static write " + floatText(staticFloat));
		var reassigned:Float = 0.5;
		reassigned = d;
		Sys.println("reassign " + floatText(reassigned));
		var fs:Array<Float> = [0.5];
		fs[0] = d;
		Sys.println("array write " + floatText(fs[0]));
		var literal:Array<Float> = [d, 0.5];
		Sys.println("array literal " + floatText(literal[0]));
		var m = new Map<String, Float>();
		m["k"] = d;
		Sys.println("map write " + floatText(m["k"]));
		var st:{f:Float} = {f: d};
		Sys.println("structure " + floatText(st.f));
		var picked:Float = always ? d : 0.5;
		Sys.println("ternary " + floatText(picked));
		var casted:Float = cast d;
		Sys.println("cast " + floatText(casted));
		var arr:Array<Dynamic> = [d];
		var element:Float = arr[0];
		Sys.println("Array<Dynamic> element " + floatText(element));
		var reflected:Float = Reflect.field({n: d}, "n");
		Sys.println("Reflect.field " + floatText(reflected));
		var called:Float = dyn(d);
		Sys.println("call " + floatText(called));
		var nullable:Null<Float> = d;
		var unwrapped:Float = nullable;
		Sys.println("Null<Float> then Float " + floatText(unwrapped));
		var lambda = (x:Float) -> x / 2;
		Sys.println("lambda argument " + Std.string(lambda(d)));
	}

	static function boolSlots(d:Dynamic) {
		Sys.println("-- Bool slots, " + Std.string(d));
		var local:Bool = d;
		Sys.println("local " + yesNo(local));
		var o:Dynamic = {n: d};
		var field:Bool = o.n;
		Sys.println("dynamic field " + yesNo(field));
		Sys.println("argument " + yesNo(d));
		Sys.println("return " + yesNo(boolOf(d)));
		var box = new Box(0.5);
		box.b = d;
		Sys.println("field write " + yesNo(box.b));
		var bs:Array<Bool> = [false];
		bs[0] = d;
		Sys.println("array write " + yesNo(bs[0]));
		var arr:Array<Dynamic> = [d];
		var element:Bool = arr[0];
		Sys.println("Array<Dynamic> element " + yesNo(element));
	}

	static function stringSlots(d:Dynamic) {
		Sys.println("-- String slots, " + Std.string(d));
		var local:String = d;
		Sys.println("local " + stringText(local));
		var o:Dynamic = {n: d};
		var field:String = o.n;
		Sys.println("dynamic field " + stringText(field));
		Sys.println("argument " + Std.string(isNull(d)));
		Sys.println("return " + stringText(stringOf(d)));
		Sys.println("return of call " + stringText(stringOfCall(d)));
		var box = new Box(0.5);
		box.s = d;
		Sys.println("field write " + stringText(box.s));
		staticString = d;
		Sys.println("static write " + stringText(staticString));
		var reassigned:String = "init";
		reassigned = d;
		Sys.println("reassign " + stringText(reassigned));
		var ss:Array<String> = ["init"];
		ss[0] = d;
		Sys.println("array write " + stringText(ss[0]));
		var m = new Map<String, String>();
		m["k"] = d;
		Sys.println("map write " + stringText(m["k"]));
		var st:{s:String} = {s: d};
		Sys.println("structure " + stringText(st.s));
		var picked:String = always ? d : "other";
		Sys.println("ternary " + stringText(picked));
		var casted:String = cast d;
		Sys.println("cast " + stringText(casted));
		var arr:Array<Dynamic> = [d];
		var element:String = arr[0];
		Sys.println("Array<Dynamic> element " + stringText(element));
		var reflected:String = Reflect.field({n: d}, "n");
		Sys.println("Reflect.field " + stringText(reflected));
		var called:String = dyn(d);
		Sys.println("call " + stringText(called));
	}

	static function otherValuesInStringSlots() {
		Sys.println("-- Int, Float and Bool values in String slots");
		var values:Array<Dynamic> = [3, 2.5, true];
		for (d in values) {
			var s:String = d;
			Sys.println(Std.string(s) + " " + (s + "!") + " " + Std.string(s == null));
		}
	}

	static function nullableSlots() {
		Sys.println("-- Null<Float> and Null<Bool> slots");
		var values:Array<Dynamic> = [3, 2.5, null];
		for (d in values) {
			var v:Null<Float> = d;
			var returned = nullableFloatOf(d);
			var text = v == null ? "null" : Std.string(v + 0.5);
			Sys.println("Null<Float> " + Std.string(v) + " " + text + " " + Std.string(returned == null));
			var box = new Box(0.5);
			box.nf = d;
			Sys.println("Null<Float> field " + Std.string(box.nf) + " " + Std.string(box.nf == null));
		}
		var flags:Array<Dynamic> = [true, false, null];
		for (d in flags) {
			var v:Null<Bool> = d;
			Sys.println("Null<Bool> " + Std.string(v) + " " + Std.string(v == null) + " " + (v == true ? "yes" : "no"));
		}
	}

	static function handled(label:String, run:() -> String) {
		try {
			consumed = run();
		} catch (e:Dynamic) {}
		Sys.println(label + " handled");
	}

	static function incompatibleValues() {
		Sys.println("-- values that the slot cannot hold");
		var values:Array<Dynamic> = ["x", true, {a: 1}];
		for (d in values) {
			handled("Float local", () -> {
				var f:Float = d;
				return Std.string(f);
			});
			handled("Float argument", () -> half(d));
			handled("Float return", () -> Std.string(floatOf(d)));
			handled("Null<Float> local", () -> {
				var v:Null<Float> = d;
				return Std.string(v);
			});
		}
		var others:Array<Dynamic> = [1, 2.5, "x", {a: 1}];
		for (d in others) {
			handled("Bool local", () -> {
				var b:Bool = d;
				return Std.string(b);
			});
			handled("Bool argument", () -> yesNo(d));
			handled("Bool return", () -> Std.string(boolOf(d)));
			handled("Bool field write", () -> {
				var box = new Box(0.5);
				box.b = d;
				return Std.string(box.b);
			});
			handled("Null<Bool> local", () -> {
				var v:Null<Bool> = d;
				return Std.string(v);
			});
		}
		Sys.println("done");
	}
}
