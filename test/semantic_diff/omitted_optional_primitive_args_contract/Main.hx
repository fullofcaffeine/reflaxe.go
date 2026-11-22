interface Scaler {
	function scale(?factor:Float, ?steps:Int, ?exact:Bool):String;
}

class BaseScaler implements Scaler {
	public function new() {}

	public function scale(?factor:Float, ?steps:Int, ?exact:Bool):String {
		return "base(" + Main.describeFloat(factor) + "," + Main.describeInt(steps) + "," + Main.describeBool(exact) + ")";
	}
}

class ChildScaler extends BaseScaler {
	public function new() {
		super();
	}

	override public function scale(?factor:Float, ?steps:Int, ?exact:Bool):String {
		return "child(" + factor + "," + steps + "," + exact + ")>" + super.scale(factor, steps, exact);
	}
}

class Box {
	public final label:String;

	public function new(?width:Float, ?count:Int, ?open:Bool) {
		label = Main.describeFloat(width) + "|" + Main.describeInt(count) + "|" + Main.describeBool(open);
	}
}

class Meter {
	final offset:Float;

	public function new(offset:Float) {
		this.offset = offset;
	}

	public function read(?extra:Float, ?times:Int):String {
		var shown = "" + extra + "/" + times;
		var total = Main.attempt(() -> "" + (offset + extra));
		var scaled = Main.attempt(() -> "" + (offset * times));
		return shown + ":" + total + ":" + scaled;
	}
}

class Pair {
	public final count:Int;
	public final ratio:Float;

	public function new(count:Int, ratio:Float) {
		this.count = count;
		this.ratio = ratio;
	}
}

class Cursor {
	public var pos:Int;
	public var scale:Float;
	public var live:Bool;
	public var pair:Pair;

	public function new(?pos:Int, ?scale:Float, ?live:Bool) {
		if (pos == null)
			pos = 1;
		if (scale == null)
			scale = 0.5;
		if (live == null)
			live = true;
		this.pos = pos;
		this.scale = scale;
		this.live = live;
		pair = new Pair(pos, scale);
	}

	public function describe():String {
		return pos + "," + scale + "," + live + "," + pair.count + "," + pair.ratio;
	}
}

/**
	Omitted optional `Int`, `Float` and `Bool` arguments must read as null,
	as in `--interp` (haxe_go-zv09).

	The argument is omitted, passed as `null` or passed a value, for static
	functions, instance methods, constructors, local functions, closures,
	interface calls, overrides and `Null<T>` parameters. A null argument
	prints `null` (concatenation, `Std.string`, interpolation), equals
	`null`, raises an error in arithmetic and compound assignment, and is
	false in an ordered comparison with a Float. It keeps null when it is
	forwarded to another optional parameter or captured by a closure, and
	a `null` check narrows it. After the function gives it a value, it can
	be stored in typed locals, fields, statics, structure fields and Array
	elements and passed to constructors and String methods. Default values
	still apply when an argument is omitted.
**/
class Main {
	public static function attempt(action:Void->String):String {
		try {
			return action();
		} catch (_:Dynamic) {
			return "threw";
		}
	}

	static function row(label:String, values:Array<String>):Void {
		Sys.println(label + "=" + values.join(":"));
	}

	public static function describeFloat(?x:Float):String {
		return "" + x + "," + (x == null) + "," + (x != null) + "," + Std.string(x) + "," + attempt(() -> "" + (x + 1)) + "," + attempt(() -> "" + (2 * x))
			+ "," + (x < 1) + "," + (x >= 0) + "," + (1 > x);
	}

	public static function describeInt(?x:Int):String {
		return "" + x + "," + (x == null) + "," + (x != null) + "," + Std.string(x) + "," + attempt(() -> "" + (x + 1)) + "," + attempt(() -> "" + (x - 2))
			+ "," + (x < 1.5) + "," + (x >= 0.5) + "," + (1.5 > x);
	}

	public static function describeBool(?x:Bool):String {
		return ""
			+ x
			+ ","
			+ (x == null)
			+ ","
			+ (x != null)
			+ ","
			+ Std.string(x)
			+ ","
			+ (x ? "yes" : "no")
			+ ","
			+ (x == true)
			+ ","
			+ (x == false);
	}

	static function nullableFloat(x:Null<Float>):String {
		return "" + x + "," + (x == null) + "," + attempt(() -> "" + (x - 1)) + "," + (x > 0);
	}

	static function nullableInt(x:Null<Int>):String {
		return "" + x + "," + (x == null) + "," + attempt(() -> "" + (x * 3)) + "," + (x <= 0.5);
	}

	static function forwardFloat(?x:Float):String {
		return describeFloat(x) + "|" + nullableFloat(x);
	}

	static function forwardInt(?x:Int):String {
		return describeInt(x) + "|" + nullableInt(x);
	}

	static function forwardBool(?x:Bool):String {
		return describeBool(x);
	}

	static function forwardToDefaults(?ratio:Float, ?count:Int, ?flag:Bool):String {
		return defaults(ratio, count, flag);
	}

	static function defaults(ratio:Float = 1.5, count:Int = 4, flag:Bool = true):String {
		return ratio
			+ ","
			+ count
			+ ","
			+ flag
			+ ","
			+ (ratio + 1)
			+ ","
			+ (count * 2)
			+ ","
			+ (flag ? "on" : "off")
			+ ","
			+ (ratio < count);
	}

	static function mixed(label:String, ?x:Float, ?y:Int):String {
		return label + ":" + x + ":" + y;
	}

	static function captured(?x:Float, ?y:Int, ?z:Bool):String {
		var read = function():String {
			return x + "/" + y + "/" + z;
		};
		return read();
	}

	static function updated(?x:Float, ?y:Int):String {
		var before = x + "/" + y;
		if (x == null) {
			x = 2.5;
		}
		if (y == null) {
			y = 3;
		}
		return before + ">" + x + "/" + y + ">" + (x + y);
	}

	static function compound(?x:Float, ?y:Int):String {
		var first = attempt(() -> {
			x += 1;
			"stepped";
		});
		var second = attempt(() -> {
			y *= 2;
			"stepped";
		});
		return first + "," + second + "," + x + "," + y;
	}

	static function interpolated(?x:Float, ?y:Int, ?z:Bool):String {
		return 'x=$x y=$y z=$z';
	}

	static function narrowed(?x:Float, ?y:Int):String {
		var left = x != null ? "" + (x + 1) : "none";
		var right = y == null ? "none" : "" + (y + 1);
		return left + "," + right;
	}

	static var lastPos:Int = 0;
	static var lastScale:Float = 0;

	static function asInt(value:Int):Int {
		return value;
	}

	static function slots(?p:Int, ?r:Float, ?f:Bool):String {
		if (p == null)
			p = 2;
		if (r == null)
			r = 1.5;
		if (f == null)
			f = false;
		var n:Int = 0;
		n = p;
		var ratio:Float = 0;
		ratio = r;
		var flag = true;
		flag = f;
		lastPos = p;
		lastScale = r;
		var record:{v:Int, w:Float} = {v: 0, w: 0};
		record.v = p;
		record.w = r;
		var cells:Array<Int> = [0, 0];
		cells[0] = p;
		cells[1] = asInt(p);
		var picked:Int = f ? 0 : p;
		return [
			"" + n,
			"" + ratio,
			"" + flag,
			"" + lastPos,
			"" + lastScale,
			"" + record.v,
			"" + record.w,
			cells.join("-"),
			"" + picked,
			"abcdef".charAt(p),
			"abcdef".substr(p, p),
			"" + "abcdef".charCodeAt(p),
			String.fromCharCode(64 + p)
		].join(",");
	}

	static function main() {
		row("static.float", [describeFloat(), describeFloat(null), describeFloat(0.5)]);
		row("static.int", [describeInt(), describeInt(null), describeInt(3)]);
		row("static.bool", [describeBool(), describeBool(null), describeBool(true), describeBool(false)]);
		row("nullable.param", [nullableFloat(null), nullableFloat(2.5), nullableInt(null), nullableInt(-2)]);
		row("forward", [
			forwardFloat(),
			forwardFloat(null),
			forwardFloat(1.25),
			forwardInt(),
			forwardInt(7),
			forwardBool(),
			forwardBool(false)
		]);
		row("defaults", [defaults(), defaults(0.25), defaults(0.25, 1, false)]);
		row("defaults.forward", [forwardToDefaults(3.5, 5, false), forwardToDefaults(0.5, 1, true)]);
		row("mixed", [mixed("a"), mixed("b", 1.5), mixed("c", null, 2), mixed("d", 0.5, 6)]);

		var meter = new Meter(10);
		row("method", [meter.read(), meter.read(null, null), meter.read(2.5), meter.read(1.5, 3)]);
		row("constructor", [new Box().label, new Box(null, null, null).label, new Box(4.5, 2, true).label]);

		function localFloat(?x:Float, ?y:Int, ?z:Bool):String {
			return x + "," + y + "," + z + "," + (x == null) + "," + attempt(() -> "" + (x + y)) + "," + (x < y);
		}
		row("local", [localFloat(), localFloat(null, null, null), localFloat(1.5, 2, false)]);
		function localDefaults(ratio:Float = 1.5, count:Int = 4, flag:Bool = true):String {
			return ratio + "," + count + "," + flag + "," + (ratio + count);
		}
		var relay = function(?ratio:Float, ?count:Int, ?flag:Bool):String {
			return localDefaults(ratio, count, flag);
		};
		row("local.defaults", [localDefaults(), localDefaults(0.5), relay(2.5, 1, false)]);

		var closure = function(?x:Float, ?y:Int, ?z:Bool):String {
			return x + "," + y + "," + z + "," + (y == null) + "," + attempt(() -> "" + (y - 1)) + "," + (y > 0.5);
		};
		row("closure", [closure(), closure(null, null, null), closure(0.5, 9, true)]);

		var scaler:Scaler = new BaseScaler();
		var child:Scaler = new ChildScaler();
		var direct = new ChildScaler();
		row("interface", [scaler.scale(), scaler.scale(null, null, null), scaler.scale(2.5, 3, true)]);
		row("override", [child.scale(), child.scale(1.5), direct.scale(), direct.scale(0.5, 1, false)]);

		row("captured", [captured(), captured(null, null, null), captured(1.5, 2, true)]);
		row("updated", [updated(), updated(null, 4), updated(0.5, 1)]);
		row("compound", [compound(), compound(null, null), compound(1.5, 2)]);
		row("interpolated", [interpolated(), interpolated(1.5, 2, true)]);
		row("narrowed", [narrowed(), narrowed(1.5, 2)]);
		row("typed.slots", [slots(), slots(3, 0.25, true)]);
		row("constructor.fields", [new Cursor().describe(), new Cursor(4, 2.5, false).describe()]);
	}
}
