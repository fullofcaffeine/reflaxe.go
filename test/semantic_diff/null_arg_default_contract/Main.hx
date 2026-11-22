enum Shade {
	Light;
	Dark;
}

enum abstract Level(Int) to Int {
	var Low = 1;
	var High = 2;
}

interface Pricer {
	function price(rate:Float = 1.0, count:Int = 1, label:String = "iface"):String;
}

class BasePricer implements Pricer {
	public function new() {}

	public function price(rate:Float = 2.5, count:Int = 4, label:String = "base"):String {
		return "base(" + rate + "," + count + "," + label + ")";
	}
}

class ChildPricer extends BasePricer {
	public function new() {
		super();
	}

	override public function price(rate:Float = 3.5, count:Int = 6, label:String = "child"):String {
		var noRate:Null<Float> = null;
		var noCount:Null<Int> = null;
		return "child(" + rate + "," + count + "," + label + ")>" + super.price(noRate, noCount, null);
	}
}

class Counter {
	public final text:String;

	public function new(start:Float = 0.25, ?step:Int = 2, on:Bool = true, label:String = "ctr") {
		text = start + "," + step + "," + on + "," + label;
	}

	public function bump(by:Int = 10, scale:Float = 1.5, on:Bool = true, tag:String = "tag"):String {
		return by + "," + scale + "," + on + "," + tag;
	}
}

class SubCounter extends Counter {
	public function new(start:Null<Float>, on:Null<Bool>, label:String) {
		super(start, null, on, label);
	}
}

class Holder {
	public var f:Null<Float>;
	public var i:Null<Int>;
	public var b:Null<Bool>;

	public function new() {}
}

abstract Meters(Float) {
	public inline function new(value:Float) {
		this = value;
	}

	public function plus(by:Float = 0.5, times:Int = 2):Float {
		return this + by * times;
	}
}

/**
	A null argument for a parameter with a default value selects that default,
	as in `--interp` (haxe_go-jnvy). Haxe applies a default whenever the
	argument is null, not only when it is omitted.

	Each row passes the argument omitted, as a `null` literal (where a static
	target allows one), as a null `Null<T>` local, field, function result or
	Array element, and as a non-null value, including `0`, `false` and `""`,
	which must stay unchanged. The forms are `x:T = d`, `?x:T = d` and
	`x:Null<T> = d` parameters of type Int, Float, Bool, String, enum and enum
	abstract, in static functions, instance methods, constructors, superclass
	constructor calls, interface calls (the implementation's default wins),
	`super` calls, local functions, closures, abstract methods, function
	values and `Dynamic` calls. A `Null<Int>` argument for a Float parameter
	also selects the default when it is null. Haxe only accepts constant
	default values, so selecting one cannot run other code.
**/
class Main {
	static function staticFloat(x:Float = 1.5):Float {
		return x;
	}

	static function staticInt(x:Int = 3):Int {
		return x;
	}

	static function staticBool(x:Bool = true):Bool {
		return x;
	}

	static function staticString(x:String = "dflt"):String {
		return x;
	}

	static function staticShade(x:Shade = Dark):String {
		return Std.string(x);
	}

	static function staticLevel(x:Level = High):Int {
		return x;
	}

	static function optFloat(?x:Float = 4.5):Float {
		return x;
	}

	static function optInt(?x:Int = 6):Int {
		return x;
	}

	static function optBool(?x:Bool = true):Bool {
		return x;
	}

	static function optString(?x:String = "opt"):String {
		return x;
	}

	static function nullInt(x:Null<Int> = 9):Null<Int> {
		return x;
	}

	static function mixed(a:Int, b:Float = 0.5, c:Bool = true, d:String = "d"):String {
		return a + "/" + b + "/" + c + "/" + d;
	}

	static function forward(?f:Float, ?i:Int, ?b:Bool):String {
		return staticFloat(f) + "," + staticInt(i) + "," + staticBool(b);
	}

	static function nothingFloat():Null<Float> {
		return null;
	}

	static function nothingInt():Null<Int> {
		return null;
	}

	static function row(label:String, values:Array<Dynamic>):Void {
		Sys.println(label + "=" + values.join(","));
	}

	static function main() {
		var nf:Null<Float> = null;
		var ni:Null<Int> = null;
		var nb:Null<Bool> = null;
		var ns:String = null;
		var nshade:Shade = null;
		var nlevel:Null<Level> = null;
		var zeroF:Null<Float> = 0.0;
		var zeroI:Null<Int> = 0;
		var falseB:Null<Bool> = false;
		var empty:String = "";

		row("static.float", [staticFloat(), staticFloat(nf), staticFloat(zeroF), staticFloat(2.25)]);
		row("static.int", [staticInt(), staticInt(ni), staticInt(zeroI), staticInt(5)]);
		row("static.bool", [staticBool(), staticBool(nb), staticBool(falseB), staticBool(false)]);
		row("static.string", [
			staticString(),
			staticString(null),
			staticString(ns),
			staticString(empty),
			staticString("x")
		]);
		row("static.enum", [staticShade(), staticShade(null), staticShade(nshade), staticShade(Light)]);
		row("static.level", [staticLevel(), staticLevel(nlevel), staticLevel(Low)]);
		row("opt.float", [optFloat(), optFloat(null), optFloat(nf), optFloat(zeroF), optFloat(0.75)]);
		row("opt.int", [optInt(), optInt(null), optInt(ni), optInt(zeroI), optInt(8)]);
		row("opt.bool", [optBool(), optBool(null), optBool(nb), optBool(falseB), optBool(false)]);
		row("opt.string", [optString(), optString(null), optString(ns), optString(empty)]);
		row("nullable.int", [nullInt(), nullInt(null), nullInt(ni), nullInt(zeroI), nullInt(12)]);
		row("mixed", [
			mixed(1),
			mixed(1, nf),
			mixed(1, nf, nb),
			mixed(1, nf, nb, ns),
			mixed(1, zeroF, falseB, empty)
		]);
		row("forward", [
			forward(),
			forward(null, null, null),
			forward(nf, ni, nb),
			forward(0.0, 0, false)
		]);

		var holder = new Holder();
		var nulls:Array<Null<Int>> = [null, 0];
		row("sources", [
			staticFloat(holder.f),
			staticInt(holder.i),
			staticBool(holder.b),
			staticFloat(nothingFloat()),
			staticInt(nothingInt()),
			staticInt(nulls[0]),
			staticInt(nulls[1])
		]);
		holder.f = 0.0;
		holder.i = 0;
		holder.b = false;
		row("sources.set", [staticFloat(holder.f), staticInt(holder.i), staticBool(holder.b)]);

		if (nf == null)
			nf = 6.5;
		if (ni != null)
			row("narrowed", [staticInt(ni)]);
		row("assigned", [staticFloat(nf)]);
		nf = null;

		var counter = new Counter();
		row("method", [
			counter.bump(),
			counter.bump(ni),
			counter.bump(ni, nf),
			counter.bump(ni, nf, nb),
			counter.bump(ni, nf, nb, ns),
			counter.bump(zeroI, zeroF, falseB, empty)
		]);
		row("ctor", [
			new Counter().text,
			new Counter(nf).text,
			new Counter(nf, null).text,
			new Counter(nf, ni, nb, ns).text,
			new Counter(zeroF, zeroI, falseB, empty).text
		]);
		row("super.ctor", [
			new SubCounter(null, null, null).text,
			new SubCounter(nf, nb, ns).text,
			new SubCounter(0.0, false, "").text
		]);

		var child:Pricer = new ChildPricer();
		var base:BasePricer = new ChildPricer();
		var plain = new BasePricer();
		row("interface", [child.price(), child.price(nf, ni, ns), child.price(zeroF, zeroI, empty)]);
		row("override", [base.price(nf, ni, ns), plain.price(), plain.price(nf, ni, ns)]);

		function localScale(x:Float = 8.5, n:Int = 2, s:String = "loc"):String {
			return x + "*" + n + "*" + s;
		}
		row("local", [
			localScale(),
			localScale(nf),
			localScale(nf, ni),
			localScale(nf, ni, ns),
			localScale(zeroF, zeroI, empty)
		]);

		var closure = function(s:String = "cl", b:Bool = false, k:Int = 11):String {
			return s + ":" + b + ":" + k;
		};
		row("closure", [
			closure(),
			closure(null),
			closure(ns, nb),
			closure(ns, nb, ni),
			closure(empty, falseB, zeroI)
		]);

		var intValue = staticInt;
		var stringValue = staticString;
		var bound = counter.bump;
		row("values", [
			intValue(),
			intValue(ni),
			intValue(zeroI),
			stringValue(ns),
			bound(),
			bound(ni, nf, nb, ns)
		]);

		var dynamicInt:Dynamic = staticInt;
		var dynamicString:Dynamic = staticString;
		row("dynamic", [dynamicInt(null), dynamicInt(0), dynamicString(null)]);

		var three:Null<Int> = 3;
		row("widen", [
			staticFloat(ni),
			staticFloat(three),
			optFloat(ni),
			optFloat(three),
			localScale(ni),
			child.price(ni, three)
		]);

		var meters = new Meters(10);
		row("abstract", [meters.plus(), meters.plus(nf), meters.plus(nf, ni), meters.plus(zeroF, zeroI)]);
	}
}
