import haxe.CallStack.StackItem;

typedef Sample = {
	var maybeWhole:Null<Int>;
	var maybeReal:Null<Float>;
}

class Counter {
	public static var spare:Null<Int> = null;

	public var maybeCount:Null<Int> = null;
	public var maybeRatio:Null<Float> = null;

	public function new() {}
}

/**
	`<`, `<=`, `>` and `>=` with a null `Null<Int>` or `Null<Float>` operand
	must behave as in `--interp` (haxe_go-wr4m).

	When exactly one operand is null, every ordered comparison is false and
	raises no error. When both operands are null, `<=` and `>=` are true and
	`<` and `>` are false. The operands are `Null<Int>` and `Null<Float>`
	parameters, omitted optional arguments, locals, instance and static
	fields, structure fields, `Array<Null<Int>>` and `Array<Null<Float>>`
	elements, call results and narrowed locals, next to Int, Float,
	`haxe.Int32` and constant operands, on either side. Non-null values, NaN,
	zero and negative numbers keep their ordinary results. The comparisons
	also run inside `if`, ternaries, `while` and `||`, and their operands run
	once, in source order.

	Each row prints one `t` (true) or `f` (false) per comparison.
**/
class Main {
	static var trail:Array<String> = [];

	static function flag(value:Bool):String {
		return value ? "t" : "f";
	}

	static function row(label:String, values:Array<Bool>):Void {
		Sys.println(label + "=" + values.map(flag).join(""));
	}

	// The eight comparisons: `a` on the left, then `b` on the left.
	// (Comparisons inside Array literals are parenthesized because the
	// formatter reads `a < b, c > d` as type parameters.)
	static function intInt(a:Null<Int>, b:Int):String {
		return [(a < b), (a <= b), (a > b), (a >= b), (b < a), (b <= a), (b > a), (b >= a)].map(flag).join("");
	}

	static function nullInts(a:Null<Int>, b:Null<Int>):String {
		return [(a < b), (a <= b), (a > b), (a >= b), (b < a), (b <= a), (b > a), (b >= a)].map(flag).join("");
	}

	static function nullReals(a:Null<Float>, b:Null<Float>):String {
		return [(a < b), (a <= b), (a > b), (a >= b), (b < a), (b <= a), (b > a), (b >= a)].map(flag).join("");
	}

	static function nullMixed(a:Null<Int>, b:Null<Float>):String {
		return [(a < b), (a <= b), (a > b), (a >= b), (b < a), (b <= a), (b > a), (b >= a)].map(flag).join("");
	}

	static function intReal(a:Null<Int>, b:Float):String {
		return [(a < b), (a <= b), (a > b), (a >= b), (b < a), (b <= a), (b > a), (b >= a)].map(flag).join("");
	}

	static function realInt(a:Null<Float>, b:Int):String {
		return [(a < b), (a <= b), (a > b), (a >= b), (b < a), (b <= a), (b > a), (b >= a)].map(flag).join("");
	}

	static function optional(?limit:Int):String {
		return [(limit < 1), (limit <= 0), (limit > -1), (limit >= 0), (1 > limit)].map(flag).join("");
	}

	// Inside the `if`, `value` is known to be non-null and `plain` is stored
	// as a plain Int; `other` can still be null.
	static function narrowed(?value:Int, ?other:Int):String {
		if (value != null) {
			var plain = value;
			return [
				(plain > 1),
				(plain <= other),
				(other >= plain),
				(value < 3),
				(value >= other),
				(other > value)
			].map(flag).join("");
		}
		return "missing";
	}

	static function someWhole(value:Int):Null<Int> {
		return value;
	}

	static function noWhole():Null<Int> {
		return null;
	}

	static function noReal():Null<Float> {
		return null;
	}

	static function tick(label:String, value:Null<Int>):Null<Int> {
		trail.push(label);
		return value;
	}

	static function tickWhole(label:String, value:Int):Int {
		trail.push(label);
		return value;
	}

	static function tickReal(label:String, value:Null<Float>):Null<Float> {
		trail.push(label);
		return value;
	}

	static function main() {
		// Function parameters: Null<Int> with Int.
		Sys.println("int.null=" + intInt(null, 1));
		Sys.println("int.null.zero=" + intInt(null, 0));
		Sys.println("int.null.negative=" + intInt(null, -1));
		Sys.println("int.value=" + intInt(2, 1));
		Sys.println("int.equal=" + intInt(1, 1));

		// Two Null<Int>, two Null<Float>, and one of each.
		Sys.println("ints.both.null=" + nullInts(null, null));
		Sys.println("ints.left.null=" + nullInts(null, 0));
		Sys.println("ints.right.null=" + nullInts(0, null));
		Sys.println("ints.value=" + nullInts(1, 2));
		Sys.println("ints.equal=" + nullInts(2, 2));
		Sys.println("reals.both.null=" + nullReals(null, null));
		Sys.println("reals.left.null=" + nullReals(null, 1.5));
		Sys.println("reals.value=" + nullReals(1.5, 2.5));
		Sys.println("reals.nan=" + nullReals(Math.NaN, Math.NaN));
		Sys.println("reals.nan.null=" + nullReals(Math.NaN, null));
		Sys.println("mixed.both.null=" + nullMixed(null, null));
		Sys.println("mixed.left.null=" + nullMixed(null, 2.5));
		Sys.println("mixed.right.null=" + nullMixed(2, null));
		Sys.println("mixed.value=" + nullMixed(2, 2.5));
		Sys.println("mixed.equal=" + nullMixed(2, 2.0));

		// Null<Int> with Float and Null<Float> with Int.
		Sys.println("int.real.null=" + intReal(null, 1.5));
		Sys.println("real.int.null=" + realInt(null, 1));
		Sys.println("real.int.value=" + realInt(1.5, 1));

		// An omitted optional Int argument is null.
		Sys.println("optional.null=" + optional());
		Sys.println("optional.value=" + optional(0));

		// Locals and constants.
		var none = noWhole();
		var other = noWhole();
		var one = someWhole(1);
		var x = 1;
		row("local", [
			(none < x),
			(x > none),
			(none <= 1),
			(1 >= none),
			(none >= -1),
			(0 < none),
			(none < other),
			(none <= other),
			(none > other),
			(none >= other),
			(one < none),
			(none <= one),
			(one <= x),
			(one > 0)
		]);

		// Instance, static and structure fields.
		var counter = new Counter();
		var sample:Sample = {maybeWhole: null, maybeReal: null};
		row("null.field", [
			(counter.maybeCount < 1),
			(Counter.spare >= x),
			(sample.maybeWhole <= 0),
			(counter.maybeCount <= Counter.spare),
			(Counter.spare >= sample.maybeWhole),
			(counter.maybeCount < Counter.spare),
			(counter.maybeRatio <= sample.maybeReal),
			(sample.maybeReal >= counter.maybeCount),
			(sample.maybeReal > counter.maybeRatio)
		]);
		counter.maybeCount = 2;
		Counter.spare = 3;
		sample.maybeWhole = 2;
		sample.maybeReal = 2.5;
		row("value.field", [
			(counter.maybeCount < 3),
			(Counter.spare >= x),
			(sample.maybeWhole <= 2),
			(counter.maybeCount <= Counter.spare),
			(Counter.spare >= sample.maybeWhole),
			(counter.maybeCount < sample.maybeWhole),
			(counter.maybeRatio <= sample.maybeReal),
			(sample.maybeReal >= counter.maybeCount),
			(sample.maybeReal > counter.maybeRatio)
		]);

		// Array<Null<Int>> and Array<Null<Float>> elements.
		var maybeInts:Array<Null<Int>> = [null, 2, null];
		var maybeReals:Array<Null<Float>> = [null, 2.5];
		var i = 1;
		row("array", [
			(maybeInts[0] < 1),
			(1 > maybeInts[0]),
			(maybeInts[1] >= 2),
			(maybeInts[0] <= maybeInts[2]),
			(maybeInts[0] < maybeInts[2]),
			(maybeInts[1] > maybeInts[0]),
			(maybeInts[i] >= maybeInts[i + 1]),
			(maybeInts[0] >= maybeReals[0]),
			(maybeReals[0] <= maybeInts[2]),
			(maybeInts[1] < maybeReals[1]),
			(maybeReals[0] >= maybeReals[0])
		]);

		// Call results.
		row("call", [
			(noWhole() < 1),
			(0 >= noWhole()),
			(noWhole() <= noWhole()),
			(noWhole() > noWhole()),
			(someWhole(2) >= noWhole()),
			(noWhole() >= noReal()),
			(noReal() <= noReal()),
			(someWhole(1) < someWhole(2)),
			(2 >= someWhole(2))
		]);

		// A Null<Int> proven non-null is a plain Int.
		Sys.println("narrowed=" + narrowed(2, null) + "," + narrowed(2, 3) + "," + narrowed(null, 1));

		// haxe.Int32 with Null<Int>.
		var narrow:haxe.Int32 = 5;
		row("int32", [(narrow < none), (none >= narrow), (narrow > one), (one <= narrow)]);

		// `haxe.CallStack` prints a frame's optional column only when it is
		// above 0; a frame without a column must render instead of raising.
		var frames = [FilePos(null, "Main.hx", 3), FilePos(Method("Main", "run"), "Main.hx", 7, 2)];
		Sys.println("stack=" + StringTools.replace(haxe.CallStack.toString(frames), "\n", "|"));

		// Conditions, ternaries and loops.
		var taken = [];
		if (none < 3) {
			taken.push("never");
		} else {
			taken.push("else");
		}
		taken.push(none <= other ? "both" : "never");
		taken.push(none < other ? "never" : "strict");
		taken.push(one >= none || none >= other ? "or" : "never");
		var steps = 0;
		var limit:Null<Int> = null;
		while (steps < limit) {
			steps++;
		}
		var counted:Null<Int> = 0;
		while (counted < 3) {
			counted = counted + 1;
		}
		var stored:Bool = none >= other;
		Sys.println("flow=" + taken.join(",") + ":" + steps + ":" + counted + ":" + flag(stored));

		// Both operands run once, in source order.
		var order = [
			(tick("a", null) < tickWhole("b", 1)),
			(tickWhole("c", 1) >= tick("d", null)),
			(tick("e", null) <= tick("f", null)),
			(tick("g", null) > tick("h", null)),
			(tickReal("i", null) >= tickReal("j", null)),
			(tick("k", null) <= tickReal("l", null)),
			(tick("m", 1) < tick("n", 2)),
			(tickWhole("o", 2) <= tick("p", 2))
		];
		Sys.println("order=" + order.map(flag).join("") + ":" + trail.join(","));
	}
}
