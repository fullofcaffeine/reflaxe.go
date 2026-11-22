typedef Tally = {
	var count:Null<Int>;
	var weight:Null<Float>;
	var seen:Null<Bool>;
}

class Meter {
	public static var total:Null<Int> = 10;
	public static var scale:Null<Float> = 1.5;
	public static var ready:Null<Bool> = false;
	public static var missing:Null<Int>;

	public var hits:Null<Int> = 3;
	public var level:Null<Float> = 0.25;
	public var open:Null<Bool> = true;
	public var none:Null<Float> = null;

	public function new() {}
}

/**
	`++`, `--`, unary `-`, `~` and `!` on nullable primitive storage
	(haxe_go-zuq3).

	`Null<Int>`, `Null<Float>` and `Null<Bool>` locals, instance and static
	fields, `Array<Null<Int>>` elements and structure fields are stored as Go
	`any`. The operators used that storage directly (`x + 1`, `-x`, `!x`), so
	Go rejected the program. Each line uses the operators in statement
	position, in value position (a variable, a call argument, an Array index)
	and in conditions. The `order` lines check that a postfix operator yields
	the old value and that the operand is read and written once, in source
	order. The `null` lines apply each operator to a null operand: `!null` is
	true and the numeric operators raise an error, as in `--interp`.
**/
class Main {
	static var trail:Array<String> = [];
	static var sinkInt:Int = 0;
	static var sinkFloat:Float = 0;

	static function note(label:String, value:Int):Int {
		trail.push(label + "=" + value);
		return value;
	}

	static function pair(first:Int, second:Int):String {
		return first + "," + second;
	}

	static function pairFloat(first:Float, second:Float):String {
		return first + "," + second;
	}

	static function keep(value:Null<Int>):Null<Int> {
		return value;
	}

	static function flip(value:Bool):String {
		return value ? "yes" : "no";
	}

	static function spend(tally:Tally):String {
		tally.count++;
		var post = tally.count--;
		var pre = ++tally.weight;
		--tally.weight;
		return tally.count + ":" + post + ":" + pre + ":" + tally.weight + ":" + !tally.seen + ":" + (-tally.count);
	}

	static function bump(?start:Int):Int {
		start++;
		return start;
	}

	static function attempt(action:Void->Void):String {
		try {
			action();
			return "ok";
		} catch (_:Dynamic) {
			return "threw";
		}
	}

	static function main() {
		// Reassigned locals keep nil-capable storage.
		var count:Null<Int> = null;
		count = 5;
		count++;
		++count;
		count--;
		--count;
		var post = count++;
		var pre = ++count;
		var arg = note("arg", count--);
		var kept = keep(--count);
		Sys.println("local.int=" + count + ":" + post + ":" + pre + ":" + arg + ":" + kept);

		var amount:Null<Float> = null;
		amount = 0.5;
		amount++;
		--amount;
		var floatPost = amount++;
		var floatPre = --amount;
		Sys.println("local.float=" + amount + ":" + floatPost + ":" + floatPre + ":" + (-amount) + ":" + pairFloat(amount--, amount));

		// Instance and static fields.
		var meter = new Meter();
		meter.hits++;
		--meter.hits;
		var hitsPost = meter.hits--;
		var hitsPre = ++meter.hits;
		meter.level++;
		var levelPost = meter.level--;
		Sys.println("field="
			+ meter.hits
			+ ":"
			+ hitsPost
			+ ":"
			+ hitsPre
			+ ":"
			+ meter.level
			+ ":"
			+ levelPost
			+ ":"
			+ (-meter.hits)
			+ ":"
			+ (~meter.hits));

		Meter.total++;
		--Meter.total;
		var totalPost = Meter.total++;
		var totalPre = note("static", --Meter.total);
		Meter.scale--;
		var scalePre = ++Meter.scale;
		Sys.println("static=" + Meter.total + ":" + totalPost + ":" + totalPre + ":" + Meter.scale + ":" + scalePre + ":" + (-Meter.scale));

		// Elements of an Array<Null<Int>> and Null<Int> Array positions.
		var slots:Array<Null<Int>> = [1, null, 7];
		slots[0]++;
		--slots[2];
		var slotPost = slots[0]--;
		var slotPre = ++slots[2];
		var position:Null<Int> = null;
		position = 0;
		var values = [10, 20, 30];
		var picked = values[position++];
		values[position++] = 99;
		slots[position--]++;
		Sys.println("array=" + slots + ":" + slotPost + ":" + slotPre + ":" + picked + ":" + values + ":" + position + ":" + (-slots[0]));

		// Structure fields and optional arguments.
		var tally:Tally = {count: 1, weight: 2.5, seen: false};
		tally.count++;
		var countPost = tally.count--;
		var weightPre = ++tally.weight;
		Sys.println("struct=" + tally.count + ":" + countPost + ":" + weightPre + ":" + tally.weight + ":" + !tally.seen);
		Sys.println("struct.arg="
			+ spend({count: 4, weight: 0.5, seen: null})
			+ ":"
			+ attempt(() -> spend({count: null, weight: 1.5, seen: true})));
		Sys.println("optional=" + bump(4) + ":" + attempt(() -> bump()));

		// `!` on Null<Bool> storage, in conditions and as values.
		var flag:Null<Bool> = null;
		flag = false;
		if (!flag) {
			Sys.println("not.local");
		}
		var rounds = 0;
		var done:Null<Bool> = null;
		done = false;
		while (!done) {
			rounds++;
			done = rounds == 3;
		}
		var negated = !flag;
		Sys.println("bool=" + negated + ":" + flip(!meter.open) + ":" + (!Meter.ready ? "ready" : "waiting") + ":" + rounds + ":" + !done);

		// Postfix yields the old value; each operand is read and written once.
		var order:Null<Int> = null;
		order = 1;
		var sum = order++ + order;
		var calls = pair(order++, order) + "|" + pair(++order, order) + "|" + pair(order, order--) + "|" + (order + order++);
		var cursor:Null<Int> = null;
		cursor = 0;
		var counts:Array<Null<Int>> = [0, 0, 0];
		counts[cursor++]++;
		counts[cursor++] += 5;
		var nested = counts[cursor--]++;
		Sys.println("order=" + sum + ":" + calls + ":" + order + ":" + counts + ":" + cursor + ":" + nested);
		Sys.println("trail=" + trail.join(","));

		// A null operand: `!` reads null as false; the numeric operators throw
		// and leave the storage unchanged.
		var nothing:Null<Int> = null;
		nothing = null;
		var nothingFloat:Null<Float> = null;
		nothingFloat = null;
		var unknown:Null<Bool> = null;
		unknown = null;
		Sys.println("null.int=" + attempt(() -> nothing++) + ":" + attempt(() -> --nothing) + ":" + attempt(() -> note("null", nothing--)) + ":"
			+ attempt(() -> sinkInt = -nothing) + ":" + attempt(() -> sinkInt = ~nothing) + ":" + nothing);
		Sys.println("null.float=" + attempt(() -> nothingFloat++) + ":" + attempt(() -> ++nothingFloat) + ":" + attempt(() -> sinkFloat = -nothingFloat)
			+ ":" + nothingFloat);
		Sys.println("null.field=" + attempt(() -> Meter.missing++) + ":" + attempt(() -> ++meter.none) + ":" + attempt(() -> slots[1]--) + ":"
			+ Meter.missing + ":" + meter.none + ":" + slots[1]);
		Sys.println("null.bool=" + !unknown + ":" + (!unknown ? "taken" : "skipped") + ":" + unknown);
	}
}
