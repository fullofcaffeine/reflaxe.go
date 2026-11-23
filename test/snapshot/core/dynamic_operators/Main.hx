class Holder {
	public static var shared:Dynamic = 1;

	public var n:Dynamic;

	public function new(n:Dynamic) {
		this.n = n;
	}
}

/**
	Operators on `Dynamic` values (haxe_go-wqb4).

	Haxe decides at run time what an operator on a `Dynamic` value means.
	haxe.go applied Go operators to the Go `any` that holds the value, which
	does not build: `d + 1` and `d += 5` became `hxrt.Int32Wrap(d)`, `d - 1`
	became `float64(d)`, `d++` became `d + 1`, `!d` and `-d` used Go
	operators that `any` lacks, `d < 3` compared an `any`, and
	`while (fnReturningDynamic())` gave Go a non-boolean condition.

	Now each such operator calls a `runtime/hxrt/dynamic_ops.go` helper with
	the rules of `--interp`: Int arithmetic wraps to 32 bits, a Float value
	makes a Float result, a String value concatenates, `/` gives a Float,
	`Int % 0` is NaN, bitwise and shift operators need Int values, a
	condition is true only for `true`, `!` is true only for `false` and null,
	and any other operand raises `Invalid operation`. Operands run once, left
	to right, and a field's receiver runs once. Typed operands keep the Go
	operators (`plain` and `real` below). Left out: `&&`/`||` results with
	non-Bool operands, postfix `++` on a `Dynamic` local holding a String or
	an inexact Float (`--interp` computes the value as `(d = d + 1) - 1`
	there), and the run-time type of an Int value in an expression that Haxe
	types as Float.
**/
class Main {
	static var order = "";
	static var counter = 0;

	static function dyn(value:Dynamic):Dynamic {
		return value;
	}

	static function step(name:String, value:Dynamic):Dynamic {
		order += name;
		return value;
	}

	static function more():Dynamic {
		counter++;
		return counter < 3;
	}

	static function attempt(name:String, run:() -> Dynamic) {
		var text = try Std.string(run()) catch (e:String) "error " + e;
		Sys.println(name + " = " + text);
	}

	static function main() {
		arithmetic();
		typedOperands();
		bitwise();
		errors();
		increments();
		compound();
		unary();
		conditions();
		comparisons();
		evaluationOrder();
	}

	static function arithmetic() {
		var i:Dynamic = dyn(7);
		var f:Dynamic = dyn(2.5);
		var s:Dynamic = dyn("x");
		var max:Dynamic = dyn(2147483647);
		var min:Dynamic = dyn(-2147483648);
		Sys.println("add " + (i + 2) + " " + (i + 0.5) + " " + (f + 2) + " " + (i + i) + " " + (i + f) + " " + (s + i) + " " + (i + s));
		Sys.println("add types " + Type.typeof(i + 2) + " " + Type.typeof(i + 0.5) + " " + Type.typeof(i + i));
		Sys.println("sub " + (i - 2) + " " + (2 - i) + " " + (f - 0.5) + " " + (i - f));
		Sys.println("mul " + (i * 2) + " " + (f * 2) + " " + (i * f));
		Sys.println("div " + (i / 2) + " " + (f / 0.5) + " " + (i / i));
		Sys.println("mod " + (i % 4) + " " + (dyn(-7) % 2) + " " + (f % 1) + " " + (dyn(-7.5) % 2));
		Sys.println("zero " + Math.isNaN(dyn(0) / 0) + " " + (dyn(1) / 0 == Math.POSITIVE_INFINITY) + " " + Math.isNaN(dyn(5) % 0));
		Sys.println("overflow " + (max + 1) + " " + (dyn(65536) * dyn(65536)) + " " + (max * 2) + " " + (min - 1 == 2147483647) + " "
			+ (max + 0.5 == 2147483647.5));
		var none:Dynamic = dyn(null);
		Sys.println("null text " + (none + "t") + " " + ("t" + none) + " " + (none + s));
	}

	static function typedOperands() {
		var n = 3;
		var x = 0.5;
		var t = "t";
		var d:Dynamic = dyn(4);
		Sys.println("typed "
			+ (d + n)
			+ " "
			+ (n + d)
			+ " "
			+ (d - x)
			+ " "
			+ (x * d)
			+ " "
			+ (d + t)
			+ " "
			+ (t + d)
			+ " "
			+ (n - d)
			+ " "
			+ (n / d));
		var plain = n * 2 + 1;
		var real = x * 4 - 1;
		Sys.println("plain " + plain + " " + real);
	}

	static function bitwise() {
		var b:Dynamic = dyn(12);
		Sys.println("bits " + (b & 10) + " " + (b | 3) + " " + (b ^ 5) + " " + (b << 2) + " " + (b >> 1) + " " + (6 & b));
		Sys.println("shifts " + (dyn(-8) >> 1) + " " + (dyn(-8) >>> 1) + " " + (dyn(1) << 33) + " " + (dyn(2147483647) << 1));
	}

	static function errors() {
		attempt("null + 1", () -> dyn(null) + 1);
		attempt("true + 1", () -> dyn(true) + 1);
		attempt("x - 1", () -> dyn("x") - 1);
		attempt("null * 2", () -> dyn(null) * 2);
		attempt("2.5 & 1", () -> dyn(2.5) & 1);
		attempt("2.0 | 1", () -> dyn(2.0) | 1);
		attempt("-null", () -> -dyn(null));
		attempt("~x", () -> ~dyn("x"));
		attempt("null++", () -> {
			var d:Dynamic = dyn(null);
			d++;
			d;
		});
		attempt("ok", () -> dyn(1) + 1);
	}

	static function increments() {
		var d:Dynamic = dyn(7);
		d++;
		++d;
		d--;
		Sys.println("local " + d);
		var post = d++;
		var pre = ++d;
		var postDown = d--;
		var preDown = --d;
		Sys.println("values " + post + " " + pre + " " + postDown + " " + preDown + " " + d);
		var f:Dynamic = dyn(2.5);
		f++;
		var floatPost = f--;
		Sys.println("float " + f + " " + floatPost);
		var max:Dynamic = dyn(2147483647);
		max++;
		var text:Dynamic = dyn("s");
		++text;
		Sys.println("wrap " + max + " text " + text);
		var o:Dynamic = {n: 1};
		o.n++;
		++o.n;
		var fieldPost = o.n++;
		var fieldPre = --o.n;
		Sys.println("anon field " + o.n + " " + fieldPost + " " + fieldPre);
		var h = new Holder(5);
		h.n++;
		var held = h.n--;
		Sys.println("class field " + h.n + " " + held);
		var viaDynamic:Dynamic = new Holder(1.5);
		viaDynamic.n++;
		Sys.println("class via dynamic " + viaDynamic.n);
		Holder.shared++;
		var sharedPre = ++Holder.shared;
		Sys.println("static " + Holder.shared + " " + sharedPre);
		var values:Array<Dynamic> = [1, 2.5];
		values[0]++;
		var elementPost = values[1]++;
		Sys.println("array " + values + " " + elementPost);
	}

	static function compound() {
		var d:Dynamic = dyn(10);
		d += 5;
		d -= 3;
		d *= 2;
		d %= 7;
		d <<= 2;
		d |= 1;
		d &= 7;
		d ^= 3;
		d >>= 1;
		d >>>= 1;
		Sys.println("local " + d);
		d /= 4;
		Sys.println("divided " + d);
		var result = (d += 1);
		Sys.println("value " + result + " " + d);
		var text:Dynamic = dyn("a");
		text += 1;
		text += "b";
		Sys.println("text " + text);
		var wrap:Dynamic = dyn(2147483647);
		wrap += 1;
		Sys.println("wrap " + wrap);
		var o:Dynamic = {n: 1, s: "s"};
		o.n += 1;
		o.n *= 10;
		o.s += 2;
		var fieldResult = (o.n -= 5);
		Sys.println("field " + o.n + " " + o.s + " " + fieldResult);
		var h = new Holder(2);
		h.n += 0.5;
		Sys.println("class field " + h.n);
		var values:Array<Dynamic> = [1];
		values[0] += 2;
		values[0] *= 3;
		Sys.println("array " + values[0]);
		var typedInt = 1;
		typedInt += dyn(2);
		var typedFloat = 1.5;
		typedFloat += dyn(2);
		Sys.println("typed targets " + typedInt + " " + typedFloat);
	}

	static function unary() {
		Sys.println("neg " + (-dyn(7)) + " " + (-dyn(2.5)) + " " + (-dyn(-2147483648) == -2147483648));
		Sys.println("not " + !dyn(true) + " " + !dyn(false) + " " + !dyn(null) + " " + !dyn(0) + " " + !dyn("") + " " + !dyn(1));
		Sys.println("complement " + (~dyn(7)) + " " + (~dyn(0)) + " " + (~dyn(-2147483648)));
		var flag:Dynamic = dyn(false);
		if (!flag)
			Sys.println("not local");
	}

	static function conditions() {
		if (dyn(true))
			Sys.println("if call true");
		if (dyn(1))
			Sys.println("if one true");
		else
			Sys.println("if one false");
		var none:Dynamic = dyn(null);
		if (none)
			Sys.println("if null true");
		else
			Sys.println("if null false");
		Sys.println("ternary " + (dyn(true) ? "yes" : "no") + " " + (dyn("x") ? "yes" : "no") + " " + (none ? "yes" : "no"));
		counter = 0;
		var loops = 0;
		while (more())
			loops++;
		Sys.println("while " + loops);
		counter = 0;
		var passes = 0;
		do {
			passes++;
		} while (more());
		Sys.println("do-while " + passes);
		var running:Dynamic = dyn(true);
		var steps = 0;
		while (running) {
			steps++;
			if (steps == 3)
				running = false;
		}
		Sys.println("while local " + steps);
		var high:Dynamic = dyn(4);
		var sum = 0;
		for (i in 0...high)
			sum += i;
		Sys.println("for range " + sum);
		var limit:Dynamic = dyn(3);
		var k = 0;
		while (k < limit)
			k++;
		Sys.println("compare loop " + k);
		Sys.println("logic "
			+ (dyn(true) && dyn(true))
			+ " "
			+ (dyn(false) || dyn(true))
			+ " "
			+ (dyn(true) && false)
			+ " "
			+ (true && dyn(false)));
		if (dyn(1) && true)
			Sys.println("and one true");
		else
			Sys.println("and one false");
		if (dyn(false) || dyn(true))
			Sys.println("or true");
	}

	static function comparisons() {
		var i:Dynamic = dyn(7);
		var none:Dynamic = dyn(null);
		var o:Dynamic = {a: 1};
		Sys.println("numbers " + (i < 8) + " " + (i > 8) + " " + (i <= 7) + " " + (i >= 7.5) + " " + (dyn(2.5) < i) + " " + (3 < i));
		Sys.println("strings " + (dyn("a") < dyn("b")) + " " + (dyn("10") < "9") + " " + (dyn("b") >= "a"));
		Sys.println("bools " + (dyn(false) < dyn(true)) + " " + (dyn(true) <= dyn(true)));
		Sys.println("mixed " + (dyn("x") < 1) + " " + (dyn("x") > 1) + " " + (i < dyn("x")));
		Sys.println("null " + (none < 1) + " " + (none >= 0) + " " + (none <= none) + " " + (none < none) + " " + (i >= none));
		Sys.println("objects " + (o <= o) + " " + (o < o) + " " + (o >= dyn({a: 1})));
		Sys.println("equality " + (i == 7) + " " + (i == 7.0) + " " + (i != dyn(8)) + " " + (none == null) + " " + (dyn("x") == "x"));
	}

	static function evaluationOrder() {
		order = "";
		var sum = step("a", 1) + step("b", 2);
		Sys.println("operands " + order + " " + sum);
		order = "";
		var less = step("l", 1) < step("r", 2);
		Sys.println("compare " + order + " " + less);
		var target:Dynamic = {n: 1};
		order = "";
		step("t", target).n += step("v", 2);
		Sys.println("field compound " + order + " " + target.n);
		order = "";
		step("t", target).n++;
		Sys.println("field increment " + order + " " + target.n);
		var y:Dynamic = dyn(10);
		var z = y - (y = 3);
		Sys.println("left first " + z + " " + y);
		var w:Dynamic = dyn(5);
		var bump = () -> {
			w = 100;
			return 1;
		};
		var captured = w + bump();
		Sys.println("captured left first " + captured + " " + w);
	}
}
