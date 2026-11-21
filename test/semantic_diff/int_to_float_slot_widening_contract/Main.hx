typedef Point = {x:Float, y:Float};

enum Reading {
	Value(amount:Float);
	Missing;
}

@:structInit
class Sample {
	public var weight:Float;
	public var label:String;
}

class Holder {
	public static var total:Float = 0;
	public static var base:Float = 1;

	public var value:Float;
	public var rate:Float = 2;

	public function new(value:Float) {
		this.value = value;
	}

	public function scale(factor:Float):Float {
		return value * factor * rate;
	}
}

class Box<T> {
	public var item:T;

	public function new(item:T) {
		this.item = item;
	}
}

class FloatBox extends Box<Float> {
	public function new(n:Int) {
		super(n);
	}
}

class Weighted extends Holder {
	public function new(n:Int) {
		super(n);
	}
}

interface Scaler {
	function scale(x:Float):Float;
}

class Doubler implements Scaler {
	public function new() {}

	public function scale(x:Float):Float {
		return x * 2;
	}
}

/**
	Int values that flow into Float slots must behave as Floats (haxe_go-ecxy).

	Every slot below has a Float (or `Null<Float>`) type and receives an Int:
	returns, assignments, call and constructor arguments, collection writes,
	structure fields, value branches and nullable storage. Each line prints
	divisions, fractions and `Std.string` results of those values, which must
	match `--interp` exactly. Null stays null where the slot is `Null<Float>`.
**/
class Main {
	static function plain(n:Int):Float {
		return n;
	}

	static function literal():Float {
		return 4;
	}

	static function viaIf(n:Int, flag:Bool):Float {
		return if (flag) n else 0.5;
	}

	static function viaSwitch(n:Int):Float {
		return switch n {
			case 0: 0.25;
			case _: n;
		};
	}

	static function maybe(n:Int, present:Bool):Null<Float> {
		if (!present)
			return null;
		return n;
	}

	static function maybeLiteral():Null<Float> {
		return 6;
	}

	static function sum(a:Float, ?b:Float = 2, ?c:Float):Float {
		return a + b + (c == null ? 0 : c);
	}

	static function total(...values:Float):Float {
		var out = 0.0;
		for (value in values)
			out += value / 2;
		return out;
	}

	static function describe(value:Null<Float>):String {
		return value == null ? "null" : Std.string(value / 4);
	}

	static function norm(p:Point):Float {
		return p.x * p.x + p.y / 2;
	}

	static function makePoint(n:Int):Point {
		return {x: n, y: 1};
	}

	static function amount(reading:Reading):Float {
		return switch reading {
			case Value(value): value / 2;
			case Missing: -1;
		};
	}

	static function main() {
		var n = 3;

		var half = function(k:Int):Float return k;
		var arrow:Int->Float = k -> k;
		Sys.println("return=" + plain(n) / 2 + " literal=" + literal() / 8 + " if=" + viaIf(n, true) / 2 + ":" + viaIf(n, false) + " switch="
			+ viaSwitch(n) / 2 + ":" + viaSwitch(0) + " closure=" + half(n) / 2 + " arrow=" + arrow(5) / 2 + " string=" + Std.string(plain(n) / 2));

		var x:Float = 1.5;
		x = n;
		var fromLiteral:Float = 1;
		fromLiteral += 0.5;
		var fromInt:Float = n;
		Holder.total = n;
		Holder.total += n;
		Holder.base = 5;
		var holder = new Holder(n);
		holder.value = n;
		holder.value += n;
		Sys.println("assign=" + x / 2 + " literal=" + fromLiteral + " local=" + fromInt / 2 + " static=" + Holder.total / 4 + ":" + Holder.base / 2
			+ " field=" + holder.value / 4 + " method=" + holder.scale(n) / 4);

		var scaleBy = function(f:Float):Float return f * 1.5;
		var scaler:Scaler = new Doubler();
		Sys.println("args=" + sum(n) / 2 + ":" + sum(n, n) / 2 + ":" + sum(n, n, n) / 2 + ":" + sum(1, 2, 3) / 4 + " rest=" + total(n, 1, n) + ":"
			+ total(...[n, 2]) + " value=" + scaleBy(n) + " interface=" + scaler.scale(n) / 4 + " math=" + Math.max(n, 0.5) / 2 + " nullable=" + describe(n)
			+ ":" + describe(null));

		var sample:Sample = {weight: n, label: "w"};
		var box = new Box<Float>(n);
		var boxed = box.item / 2;
		box.item = 5;
		var reading = Value(n);
		Sys.println("ctor=" + new Holder(n).value / 2 + " super=" + new Weighted(n).value / 2 + ":" + new FloatBox(n).item / 2 + " struct="
			+ sample.weight / 2 + " box=" + boxed + ":" + box.item / 2 + " enum=" + amount(reading) + ":" + amount(Value(7)) + ":" + amount(Missing));

		var values:Array<Float> = [1, n, 2.5];
		values[0] = n;
		values.push(n);
		values.push(4);
		values.unshift(n);
		values.insert(1, 7);
		values[1] += n;
		var halves = 0.0;
		for (value in values)
			halves += value / 2;
		var weights = new Map<String, Float>();
		weights.set("a", n);
		weights["b"] = n;
		weights.set("c", 2);
		var literalMap:Map<String, Float> = ["d" => n, "e" => 1];
		var optional:Array<Null<Float>> = [n, null];
		optional.push(2);
		var vector = new haxe.ds.Vector<Float>(2);
		vector[0] = n;
		vector[1] = 1;
		Sys.println("array="
			+ values.join(",")
			+ " halves="
			+ halves
			+ " first="
			+ values[0] / 2
			+ " find="
			+ values.indexOf(n)
			+ ":"
			+ values.contains(7)
			+ " optional="
			+ optional[0] / 2
			+ ":"
			+ optional[1]
			+ ":"
			+ optional[2] / 4 + " vector=" + vector[0] / 2 + ":" + vector[1] / 4
			+ " map="
			+ weights.get("a") / 2
			+ ":"
			+ weights["b"] * 1.5 + ":" + weights["c"] / 4 + ":" + literalMap["d"] / 2 + ":" + literalMap["e"] / 4);

		var p:Point = {x: n, y: 1};
		var made = makePoint(n);
		var p2:Point = {x: 1, y: 2};
		p2.x = n;
		p2.y += n;
		Sys.println("anon=" + norm(p) + ":" + norm({x: 1, y: n}) + ":" + norm(made) + ":" + made.x / 2 + ":" + p2.x / 2 + ":" + p2.y / 2);

		var ternary:Float = n > 1 ? n : 0.5;
		var switched:Float = switch n {
			case 3: n;
			case _: 0.75;
		};
		var mixed = n > 1 ? n : 0.5;
		Sys.println("arms=" + ternary / 2 + ":" + switched / 2 + ":" + mixed / 2);

		var nf:Null<Float> = n;
		var nfLiteral:Null<Float> = 4;
		var ni:Null<Int> = n;
		var nfFromNi:Null<Float> = ni;
		var reassigned:Null<Int> = null;
		if (n > 0)
			reassigned = n + 1;
		var nfFromReassigned:Null<Float> = reassigned;
		var noInt:Null<Int> = null;
		var nfNull:Null<Float> = noInt;
		var fromNi:Float = ni;
		var fromReassigned:Float = reassigned;
		Sys.println("nullable=" + (nf + 0.5) + ":" + nfLiteral / 8 + ":" + nfFromNi / 2 + ":" + nfFromReassigned / 2 + ":" + (nfNull == null) + ":"
			+ fromNi / 2 + ":" + fromReassigned / 2 + ":" + maybe(n, true) / 2 + ":" + maybe(n, false) + ":" + maybeLiteral() / 4);
	}
}
