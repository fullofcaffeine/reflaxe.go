typedef Gauge = {
	var level:Null<Float>;
}

class Meter {
	public static var shared:Null<Float> = null;

	public var reading:Null<Float> = null;

	public function new() {}
}

/**
	Binary `+`, `-`, `*`, `/` and `%`, compound assignment, and `<`, `<=`,
	`>` and `>=` on a null `Null<Float>` (haxe_go-tmrg).

	A `Null<Float>` local, field, structure field or `Array<Null<Float>>`
	element is stored as Go `any`. The operators read null from that storage
	as `0`, so `null + 1` gave `1` and `null < 1` gave `true`, and a
	comparison of an `Array<Null<Float>>` element did not compile. As in
	`--interp`, an arithmetic operator on null must now raise an error that
	leaves the storage null, and an ordered comparison with null must be
	false. The `order` line checks that a comparison still evaluates both
	operands. `==`, `!=`, string concatenation (which prints `null`),
	narrowed values and non-null values keep their output.
**/
class Main {
	static var trail:Array<String> = [];

	static function attempt(action:Void->String):String {
		try {
			return action();
		} catch (_:Dynamic) {
			return "threw";
		}
	}

	static function row(label:String, values:Array<String>):Void {
		Sys.println(label + "=" + values.join(":"));
	}

	static function none():Null<Float> {
		return null;
	}

	static function some(value:Float):Null<Float> {
		return value;
	}

	static function cells():Array<Null<Float>> {
		return [null, 1.5];
	}

	static function tick(label:String, value:Float):Float {
		trail.push(label);
		return value;
	}

	static function main() {
		var empty = none();
		var half = some(0.5);
		var two:Float = 2;
		var three:Int = 3;

		// Each operator with null on the left, on the right and on both sides,
		// then on non-null values.
		row("add", [
			attempt(() -> "" + (empty + 1)),
			attempt(() -> "" + (1 + empty)),
			attempt(() -> "" + (empty + empty)),
			attempt(() -> "" + (empty + two)),
			attempt(() -> "" + (three + empty)),
			"" + (half + 1),
			"" + (1 + half),
			"" + (half + half),
			"" + (three + half)
		]);
		row("sub", [
			attempt(() -> "" + (empty - 1)),
			attempt(() -> "" + (two - empty)),
			"" + (half - 1),
			"" + (two - half)
		]);
		row("mul", [
			attempt(() -> "" + (empty * 2)),
			attempt(() -> "" + (two * empty)),
			"" + (half * 4),
			"" + (two * half)
		]);
		row("div", [
			attempt(() -> "" + (empty / 2)),
			attempt(() -> "" + (two / empty)),
			"" + (half / 2),
			"" + (two / half)
		]);
		row("mod", [
			attempt(() -> "" + (empty % 2)),
			attempt(() -> "" + (two % empty)),
			"" + (half % 2),
			"" + (two % half)
		]);

		// Ordered comparisons with null are false; equality still sees null.
		row("compare.null", [
			"" + (empty < 1),
			"" + (1 < empty),
			"" + (empty <= 1),
			"" + (1 <= empty),
			"" + (empty > 1),
			"" + (1 > empty),
			"" + (empty >= 1),
			"" + (1 >= empty),
			"" + (empty < empty),
			"" + (empty >= two),
			"" + (two <= empty),
			"" + !(empty < 1)
		]);
		row("compare.value", [
			"" + (half < 1),
			"" + (1 < half),
			"" + (half <= 0.5),
			"" + (half >= two),
			"" + (two > half),
			"" + (half > half)
		]);
		row("condition", [
			empty < 1 ? "taken" : "skipped",
			empty >= 0 ? "taken" : "skipped",
			!(empty > 0) ? "negated" : "kept",
			half < 1 ? "taken" : "skipped"
		]);
		row("equal", [
			"" + (empty == null),
			"" + (empty != null),
			"" + (empty == 1),
			"" + (empty != 1),
			"" + (half == null),
			"" + (half == 0.5),
			"" + (half != 0.5)
		]);

		// Compound assignment raises before it stores anything.
		var sum = none();
		var diff = none();
		var product = none();
		var quotient = none();
		var rest = none();
		row("compound.null", [
			attempt(() -> {
				sum += 1;
				"stepped";
			}),
			attempt(() -> {
				diff -= 1;
				"stepped";
			}),
			attempt(() -> {
				product *= 2;
				"stepped";
			}),
			attempt(() -> {
				quotient /= 2;
				"stepped";
			}),
			attempt(() -> {
				rest %= 2;
				"stepped";
			}),
			"" + sum,
			"" + diff,
			"" + product,
			"" + quotient,
			"" + rest
		]);
		var plain:Float = 1;
		row("compound.rhs", [
			attempt(() -> {
				plain += empty;
				"" + plain;
			}),
			attempt(() -> "" + (plain *= empty)),
			"" + plain
		]);
		var level = some(1.5);
		level += 1;
		level -= 0.5;
		level *= 3;
		level /= 4;
		level %= 1;
		row("compound.value", ["" + level]);

		// Instance and static fields.
		var meter = new Meter();
		row("field", [
			attempt(() -> "" + (meter.reading + 1)),
			attempt(() -> "" + (2 * meter.reading)),
			attempt(() -> {
				meter.reading += 1;
				"stepped";
			}),
			attempt(() -> {
				meter.reading -= 1;
				"stepped";
			}),
			"" + (meter.reading < 1),
			"" + (meter.reading >= 0),
			"" + meter.reading
		]);
		row("static", [
			attempt(() -> "" + (Meter.shared - 1)),
			attempt(() -> {
				Meter.shared *= 2;
				"stepped";
			}),
			"" + (Meter.shared > 0),
			"" + Meter.shared
		]);
		meter.reading = 2.5;
		Meter.shared = 4;
		meter.reading += Meter.shared;
		row("field.value", [
			"" + meter.reading,
			"" + (Meter.shared * meter.reading),
			"" + (meter.reading > Meter.shared)
		]);

		// Structure fields.
		var gauge:Gauge = {level: null};
		row("struct", [
			attempt(() -> "" + (gauge.level + 1)),
			attempt(() -> {
				gauge.level /= 2;
				"stepped";
			}),
			"" + (gauge.level <= 0),
			"" + gauge.level
		]);
		gauge.level = 3;
		gauge.level -= 0.5;
		row("struct.value", ["" + (gauge.level * 2), "" + (gauge.level > 2)]);

		// Elements of an Array<Null<Float>>.
		var slots = cells();
		row("array", [
			attempt(() -> "" + (slots[0] + 1)),
			attempt(() -> "" + (3 - slots[0])),
			attempt(() -> {
				slots[0] += 1;
				"stepped";
			}),
			"" + (slots[0] < 2),
			"" + (2 > slots[0]),
			"" + slots[0],
			"" + (slots[1] + 1),
			"" + (slots[1] < 2),
			"" + (slots[1] * slots[1])
		]);

		// Call results, and Null<Int> operands mixed with Null<Float>.
		row("call", [
			attempt(() -> "" + (none() + 1)),
			attempt(() -> "" + (1 - none())),
			"" + (none() < 1),
			"" + (some(2) * 2)
		]);
		var count:Null<Int> = null;
		count = 2;
		row("mixed", [
			attempt(() -> "" + (count + empty)),
			attempt(() -> "" + (empty * count)),
			"" + (count + half),
			"" + (half * count)
		]);

		// Concatenation with a string prints null; it is not arithmetic.
		var text = "s";
		text += empty;
		row("text", ["x" + empty, empty + "x", text, "a" + empty + 1, half + "!", "v" + meter.reading]);

		// Narrowed values are plain Floats.
		var maybe = none();
		var known = some(1.25);
		row("narrowed", [
			maybe != null ? "" + (maybe + 1) : "none",
			known != null ? "" + (known + 1) : "none",
			known != null && known < 2 ? "small" : "large"
		]);

		// A comparison evaluates both operands; arithmetic raises after the
		// left operand runs.
		var order = [
			"" + (empty < tick("compare", 1)),
			attempt(() -> "" + (tick("left", 1) + empty)),
			attempt(() -> "" + (tick("scaled", 2) * empty))
		];
		row("order", order.concat([trail.join(",")]));
	}
}
