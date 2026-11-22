typedef Slot = {
	var amount:Null<Float>;
	var count:Null<Int>;
}

@:structInit
class Options {
	public var weight:Null<Float>;
	public var label:String;
}

class Gauge {
	public static var shared:Null<Float>;
	public static var limit:Null<Float> = 4;
	public static var hits:Null<Int> = null;
	public static var enabled:Null<Bool> = false;

	public var reading:Null<Float>;
	public var amount:Null<Float>;
	public var count:Null<Int> = 3;
	public var rate:Null<Float> = 2;
	public var flag:Null<Bool> = true;
	public var none:Null<Float> = null;

	public function new(start:Float, whole:Int) {
		amount = start;
		reading = whole;
	}

	public function current():Null<Float> {
		return amount;
	}

	public function plainAmount():Float {
		return amount;
	}
}

class SubGauge extends Gauge {
	public function new() {
		super(0.5, 7);
		count = null;
	}

	public function doubled():Float {
		return amount * 2;
	}
}

/**
	Nullable primitive class fields must behave as in `--interp` (haxe_go-0a8l).

	Every field below is `Null<Int>`, `Null<Float>` or `Null<Bool>`, as an
	instance or a static field. Each line prints the fields after writes of
	null, of values and of Ints, through initializers, constructors, compound
	assignment, parameters, `??`, a structure view, reflection, a subclass and
	a `@:structInit` class. The output must match `--interp` exactly: null
	stays null, and Int values written to a `Null<Float>` field read back as
	Floats.
**/
class Main {
	static function half(value:Float):Float {
		return value / 2;
	}

	static function twice(value:Int):Int {
		return value * 2;
	}

	static function describe(value:Null<Float>):String {
		return value == null ? "none" : "some(" + value / 4 + ")";
	}

	static function main() {
		var g = new Gauge(1.5, 5);
		Sys.println("init="
			+ g.amount
			+ ":"
			+ g.reading / 2
			+ ":"
			+ g.count
			+ ":"
			+ g.rate / 4
			+ ":"
			+ g.flag
			+ ":"
			+ g.none
			+ ":"
			+ (g.none == null));

		g.amount = null;
		Sys.println("null=" + g.amount + ":" + (g.amount == null) + ":" + (g.amount != null) + ":" + (null == g.amount) + ":" + describe(g.amount) + ":"
			+ (g.amount ?? 0.25));
		g.amount = 2.5;
		Sys.println("value=" + g.amount + ":" + (g.amount == null) + ":" + (g.amount == 2.5) + ":" + (g.amount > 1) + ":" + describe(g.amount) + ":"
			+ (g.amount ?? 0.25));
		var n = 3;
		g.amount = n;
		Sys.println("int=" + g.amount / 2 + ":" + half(g.amount));
		g.amount = 7;
		g.amount += 1;
		g.amount *= 0.5;
		g.amount -= n;
		g.count += 2;
		g.count += 1;
		Sys.println("compound=" + g.amount / 2 + ":" + g.count + ":" + twice(g.count));

		var local:Null<Float> = g.amount;
		var plain:Float = g.amount;
		g.rate = g.count;
		g.none = g.amount;
		Sys.println("copy=" + local / 2 + ":" + plain / 4 + ":" + g.rate / 4 + ":" + g.none / 2 + ":" + g.current() / 4);
		g.count = null;
		g.rate = g.count;
		Sys.println("nullcopy=" + g.count + ":" + g.rate + ":" + (g.rate == null));

		g.flag = null;
		Sys.println("bool=" + g.flag + ":" + (g.flag == null));
		g.flag = false;
		Sys.println("bool=" + g.flag + ":" + (g.flag == false) + ":" + (g.flag ?? true));
		g.flag = true;
		if (g.flag)
			Sys.println("bool=" + g.flag + ":" + g.plainAmount() / 4 + ":" + (g.amount == g.none));

		Sys.println("static=" + Gauge.shared + ":" + Gauge.limit / 8 + ":" + Gauge.hits + ":" + Gauge.enabled);
		Gauge.shared = 1;
		Gauge.shared += 0.5;
		Gauge.hits = 4;
		Gauge.hits += 1;
		Gauge.limit = null;
		Gauge.enabled = null;
		Sys.println("static=" + Gauge.shared / 2 + ":" + Gauge.limit + ":" + Gauge.hits + ":" + Gauge.enabled + ":" + (Gauge.limit ?? 0.75));

		var slot:Slot = {amount: null, count: 1};
		Sys.println("anon=" + slot.amount + ":" + slot.count);
		slot.amount = n;
		slot.count = null;
		Sys.println("anon=" + slot.amount / 2 + ":" + slot.count);
		var view:Slot = g;
		view.amount = 4.5;
		view.count = 9;
		Sys.println("view=" + g.amount + ":" + g.count + ":" + view.amount / 2);
		g.amount = null;
		Sys.println("view=" + g.amount + ":" + (view.amount == null));

		Sys.println("reflect=" + Reflect.field(g, "amount") + ":" + Reflect.field(g, "count"));
		Reflect.setField(g, "amount", 6.5);
		Reflect.setField(g, "count", null);
		Sys.println("reflect=" + g.amount / 2 + ":" + g.count + ":" + Reflect.field(g, "count"));
		Reflect.setField(g, "amount", 3);
		Sys.println("reflect=" + g.amount / 2 + ":" + Reflect.field(g, "amount"));
		Reflect.setField(g, "flag", false);
		Reflect.setField(g, "rate", null);
		Sys.println("reflect=" + g.flag + ":" + g.rate);
		Reflect.setField(g, "amount", null);
		Sys.println("reflect=" + g.amount + ":" + (g.amount == null));

		var sub = new SubGauge();
		Sys.println("sub=" + sub.amount + ":" + sub.reading / 2 + ":" + sub.count + ":" + sub.doubled());

		var options:Options = {weight: 3, label: "w"};
		Sys.println("struct=" + options.weight / 2 + ":" + options.label);
		options.weight = null;
		Sys.println("struct=" + options.weight);
	}
}
