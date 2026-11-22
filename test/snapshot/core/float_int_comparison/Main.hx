typedef Sample = {
	var whole:Int;
	var real:Float;
	var maybeWhole:Null<Int>;
	var maybeReal:Null<Float>;
}

class Meter {
	public static var limit:Int = 3;
	public static var scale:Float = 2.5;
	public static var spare:Null<Int> = null;

	public var count:Int;
	public var ratio:Float;
	public var maybeCount:Null<Int> = null;
	public var maybeRatio:Null<Float> = null;

	public function new(count:Int, ratio:Float) {
		this.count = count;
		this.ratio = ratio;
	}
}

/**
	`<`, `<=`, `>`, `>=`, `==` and `!=` between a Float and an Int
	(haxe_go-i91r).

	Haxe types such a comparison as Float. haxe.go compared the Go `float64`
	and `int` operands directly, which Go rejects, so a comparison of a Float
	local, field, call result or Array element with an Int one did not
	compile. A `Null<Float>` value (Go `any`) compared with `==` to an Int
	constant or a `Null<Int>` compared Go types as well as values, so
	`Null<Float>` 2.0 was not equal to `2`. The Int operand must now be
	converted to Float: `float64(...)` for a typed Int, a Float constant such
	as `2.0` next to `any`, and a nil-preserving read for a `Null<Int>`. An
	Int constant next to a typed `float64` stays as written. An ordered
	comparison with a null `Null<Int>` is false, as for a null `Null<Float>`.
	The `order` line checks that both operands still run once, in source
	order. `UInt` operands already compared through `UInt.toFloat` and keep
	their output.

	Each row prints one `t` (true) or `f` (false) per comparison.
**/
class Main {
	static inline var MAX_INT:Int = 2147483647;
	static inline var HALF:Float = 0.5;

	static var trail:Array<String> = [];

	static function flag(value:Bool):String {
		return value ? "t" : "f";
	}

	static function row(label:String, values:Array<Bool>):Void {
		Sys.println(label + "=" + values.map(flag).join(""));
	}

	// The twelve comparisons: Float on the left, then Int on the left.
	// (Comparisons inside Array literals are parenthesized because the
	// formatter reads `a < b, c > d` as type parameters.)
	static function all(f:Float, i:Int):String {
		return [
			(f < i), (f <= i), (f > i), (f >= i), (f == i), (f != i),
			(i < f), (i <= f), (i > f), (i >= f), (i == f), (i != f)
		].map(flag).join("");
	}

	static function whole(value:Int):Int {
		return value;
	}

	static function real(value:Float):Float {
		return value;
	}

	static function someWhole(value:Int):Null<Int> {
		return value;
	}

	static function noWhole():Null<Int> {
		return null;
	}

	static function someReal(value:Float):Null<Float> {
		return value;
	}

	static function noReal():Null<Float> {
		return null;
	}

	static function tickWhole(label:String, value:Int):Int {
		trail.push(label);
		return value;
	}

	static function tickReal(label:String, value:Float):Float {
		trail.push(label);
		return value;
	}

	static function tickMaybeWhole(label:String, value:Null<Int>):Null<Int> {
		trail.push(label);
		return value;
	}

	static function main() {
		var minInt = -MAX_INT - 1;

		// Locals (function parameters).
		Sys.println("fraction=" + all(2.5, 2));
		Sys.println("equal=" + all(2.0, 2));
		Sys.println("below.zero=" + all(-0.5, 0));
		Sys.println("negative=" + all(-3.0, -3));
		Sys.println("negative.fraction=" + all(-2.5, -3));
		Sys.println("max=" + all(2147483647.0, MAX_INT));
		Sys.println("max.fraction=" + all(2147483646.5, MAX_INT));
		Sys.println("min=" + all(-2147483648.0, minInt));
		Sys.println("float32.gap=" + all(16777216.0, 16777217));
		Sys.println("nan=" + all(Math.NaN, 0));
		Sys.println("inf=" + all(Math.POSITIVE_INFINITY, MAX_INT));
		Sys.println("neg.inf=" + all(Math.NEGATIVE_INFINITY, minInt));

		// The Int operand keeps Int overflow before it is converted.
		var top = MAX_INT;
		var topReal:Float = MAX_INT;
		row("wrapped", [
			(topReal < top + 1),
			(top + 1 < topReal),
			(topReal == top + 1),
			(top * 2 == -2.0)
		]);

		// Constants on either side.
		var i = 2;
		var f = 2.5;
		row("constant", [
			(i < 2.5),
			(2.5 > i),
			(i <= 2.0),
			(i == 2.0),
			(2.0 == i),
			(i != 2.5),
			(f < 3),
			(3 > f),
			(f == 2),
			(2 != f),
			(f >= -1),
			(-1 <= f),
			(i < HALF),
			(HALF < i),
			(MAX_INT == 2147483647.0),
			(i < Math.POSITIVE_INFINITY),
			(Math.NaN != i),
			(Math.NEGATIVE_INFINITY < i)
		]);

		// Instance, static and structure fields.
		var meter = new Meter(2, 2.5);
		var sample:Sample = {
			whole: 3,
			real: 3.0,
			maybeWhole: null,
			maybeReal: null
		};
		row("field", [
			  (meter.count < meter.ratio),  (meter.ratio >= meter.count),  (meter.count == meter.ratio),
			  (Meter.limit > Meter.scale),  (Meter.scale == Meter.limit),  (Meter.scale <= Meter.limit),
			(sample.whole <= sample.real), (sample.real != sample.whole), (sample.whole == sample.real),
			  (meter.count < Meter.scale),  (Meter.limit == sample.real),  (sample.whole > meter.ratio)
		]);

		// Call results.
		row("call", [
			(whole(2) < real(2.5)),
			(real(2.0) == whole(2)),
			(whole(3) > 2.5),
			(2.5 < whole(3)),
			(real(-1.5) > whole(-2)),
			(whole(MAX_INT) != real(2147483647.0)),
			(real(f) >= whole(i))
		]);

		// Array<Int> and Array<Float> elements.
		var ints = [1, 2, 3];
		var reals = [1.5, 2.0, 2.5];
		row("array", [
			(ints[1] < reals[2]),
			(reals[1] == ints[1]),
			(ints[2] >= reals[2]),
			(reals[0] > ints[0]),
			(ints[0] != reals[0]),
			(ints[2] > f),
			(f < ints[2]),
			(reals[1] == i),
			(i < reals[2]),
			(ints[i] == 3.0),
			(reals[i - 1] <= ints[1])
		]);

		// Null<Int> and Null<Float> operands, null and non-null.
		var noInt = noWhole();
		var two = someWhole(2);
		var noFloat = noReal();
		var twoReal = someReal(2.0);
		var half = someReal(2.5);
		row("null.ordered", [
			(noInt < f),
			(f > noInt),
			(noInt <= 2.5),
			(2.5 >= noInt),
			(noFloat < i),
			(i > noFloat),
			(noFloat < noInt),
			(noInt > noFloat),
			(noFloat <= 2),
			!(noInt < f)
		]);
		row("value.ordered", [
			(two < f),
			(f > two),
			(two >= twoReal),
			(half > two),
			(two <= 2.5),
			(twoReal <= i),
			(i < half),
			(half >= 2),
			(two < half),
			(half < two)
		]);
		row("null.equal", [
			(noFloat == noInt),
			(noInt == noFloat),
			(noFloat != noInt),
			(noFloat == two),
			(two != noFloat),
			(noInt == f),
			(f != noInt),
			(noFloat == i),
			(i != noFloat),
			(noFloat == 2),
			(noInt == 2.0)
		]);
		row("value.equal", [
			(twoReal == two), (two == twoReal),  (half != two), (twoReal == i), (i == twoReal), (two == 2.0),
			    (2.0 == two),     (two != 2.5), (twoReal == 2), (2 == twoReal),    (half != 2),   (two == f)
		]);

		// Null<Int> and Null<Float> fields, structure fields, elements and calls.
		var maybeInts:Array<Null<Int>> = [null, 2];
		var maybeReals:Array<Null<Float>> = [null, 2.0];
		row("null.field", [
			(meter.maybeCount < meter.ratio),
			(meter.maybeRatio == meter.count),
			(meter.maybeRatio != meter.count),
			(Meter.spare <= Meter.scale),
			(Meter.spare == Meter.scale),
			(sample.maybeWhole < sample.real),
			(sample.maybeReal == sample.whole),
			(sample.maybeWhole == sample.maybeReal)
		]);
		meter.maybeCount = 2;
		meter.maybeRatio = 2.0;
		Meter.spare = 3;
		sample.maybeWhole = 3;
		sample.maybeReal = 2.5;
		row("value.field", [
			(meter.maybeCount < meter.ratio),
			(meter.maybeRatio == meter.count),
			(meter.maybeCount == meter.maybeRatio),
			(Meter.spare > Meter.scale),
			(Meter.spare == Meter.scale),
			(sample.maybeWhole > sample.real),
			(sample.maybeReal != sample.whole),
			(sample.maybeWhole >= sample.maybeReal)
		]);
		row("null.array", [
			(maybeInts[0] < f),
			(maybeInts[0] == f),
			(maybeInts[1] == 2.0),
			(maybeInts[1] < reals[2]),
			(maybeReals[0] > i),
			(maybeReals[1] == i),
			(maybeReals[1] >= ints[1]),
			(maybeInts[1] == maybeReals[1]),
			(maybeInts[0] == maybeReals[0])
		]);
		row("null.call", [
			(noWhole() < 1.5),
			(someWhole(2) == 2.0),
			(someReal(2.0) == 2),
			(noReal() >= 0),
			(someReal(2.5) > someWhole(2)),
			(noWhole() == noReal()),
			(someWhole(2) != someReal(2.0)),
			(noReal() == i)
		]);

		// Narrowed Null<Int> and Null<Float> locals are plain values.
		var maybeTwo = someWhole(2);
		var maybeHalf = someReal(2.5);
		row("narrowed", [
			(maybeTwo != null && maybeTwo < f),
			(maybeTwo != null && maybeTwo == 2.0),
			(maybeHalf != null && maybeHalf > i),
			(maybeHalf != null && maybeHalf != i),
			(maybeTwo != null && maybeHalf != null && maybeTwo < maybeHalf)
		]);

		// Conditions, ternaries, loops and guards.
		var taken = [];
		if (f > i) {
			taken.push("if");
		}
		if (i >= f) {
			taken.push("never");
		} else {
			taken.push("else");
		}
		taken.push(i < f ? "ternary" : "never");
		taken.push(noInt < f ? "never" : "null.ternary");
		taken.push(f < i && i < 3 ? "never" : "and");
		taken.push(f > i || noInt < f ? "or" : "never");
		var steps = 0;
		while (steps < f) {
			steps++;
		}
		var level = 10.0;
		var drops = 0;
		while (level > i) {
			level -= 1.5;
			drops++;
		}
		var rounds = 0;
		do {
			rounds++;
		} while (rounds <= HALF * 5);
		var counter:Null<Int> = 0;
		while (counter < 1.5) {
			counter = counter + 1;
		}
		var guard = switch (i) {
			case value if (value > f): "above";
			case value if (value == 2.0): "two";
			case _: "other";
		};
		var stored:Bool = f >= i;
		Sys.println("flow=" + taken.join(",") + ":" + steps + ":" + drops + ":" + rounds + ":" + counter + ":" + guard + ":" + flag(stored));

		// Both operands run once, in source order.
		var order = [
			(tickReal("a", 1.5) < tickWhole("b", 2)),
			(tickWhole("c", 2) >= tickReal("d", 2.0)),
			(tickWhole("e", 2) == tickReal("f", 2.0)),
			(tickReal("g", 1.0) != tickWhole("h", 1)),
			(tickMaybeWhole("i", null) < tickReal("j", 1.0)),
			(tickReal("k", 2.0) == tickMaybeWhole("l", 2))
		];
		Sys.println("order=" + order.map(flag).join("") + ":" + trail.join(","));

		// UInt and haxe.Int32 operands.
		var small:UInt = 3;
		var huge:UInt = 0xFFFFFFFF;
		var narrow:haxe.Int32 = 5;
		row("uint", [
			(small > f),
			(f < small),
			(small == 3.0),
			(huge > f),
			(huge > 4000000000.0),
			(huge == 4294967295.0)
		]);
		row("int32", [(narrow > f), (f < narrow), (narrow == 5.0), (narrow != f), (narrow <= 5.0)]);
	}
}
