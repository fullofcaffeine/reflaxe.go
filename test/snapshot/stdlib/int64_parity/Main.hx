import haxe.Int64;

class Main {
	static inline function emit(label:String, value:Int64):Void {
		Sys.println(label + "=" + Int64.toStr(value));
	}

	static function main() {
		var max = Int64.parseString("9223372036854775807");
		var min = Int64.parseString("-9223372036854775808");
		emit("wrap_add", max + Int64.ofInt(1));
		emit("wrap_sub", min - Int64.ofInt(1));

		var a = Int64.parseString("1234567890123");
		var b = Int64.parseString("-987654321");
		emit("sum", a + b);
		emit("diff", a - b);
		emit("mul", Int64.ofInt(30000) * Int64.ofInt(30000));

		var positive = Int64.divMod(Int64.parseString("123456789"), Int64.ofInt(97));
		emit("div_q", positive.quotient);
		emit("div_r", positive.modulus);
		var negative = Int64.divMod(Int64.parseString("-123456789"), Int64.ofInt(97));
		emit("div_neg_q", negative.quotient);
		emit("div_neg_r", negative.modulus);

		Sys.println("cmp=" + Int64.compare(Int64.ofInt(-1), Int64.ofInt(1)));
		Sys.println("ucmp=" + Int64.ucompare(Int64.ofInt(-1), Int64.ofInt(1)));

		emit("shl", Int64.shl(Int64.ofInt(1), 40));
		emit("shr", Int64.parseString("-8") >> 1);
		emit("ushr", Int64.parseString("-1") >>> 1);

		emit("from_float", Int64.fromFloat(9007199254740991.0));
		Sys.println("to_int_ok=" + Int64.toInt(Int64.ofInt(2147483647)));
		try {
			Int64.toInt(Int64.parseString("2147483648"));
			Sys.println("to_int_overflow=missing");
		} catch (e:Dynamic) {
			Sys.println("to_int_overflow=" + Std.string(e));
		}

		var round = Int64.make(0x7fffffff, -12345);
		Sys.println("round_high=" + round.high);
		Sys.println("round_low=" + round.low);
	}
}
