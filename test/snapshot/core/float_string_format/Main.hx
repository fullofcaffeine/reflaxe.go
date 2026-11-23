enum Holder {
	Hold(value:Float);
}

class Sample {
	public var value:Float;

	public function new(value:Float) {
		this.value = value;
	}

	public function half():Float {
		return value / 2;
	}
}

/**
	Float to text (haxe_go-rlvo).

	haxe.go printed Floats with Go's `fmt.Sprint`, which switches to exponent
	form at 1e6 and prints the shortest digits: 2147483647.0 became
	`2.147483647e+09` and 0.1 + 0.2 became `0.30000000000000004`, where
	`--interp` prints `2147483647` and `0.300000000000000044`. NaN and the
	infinities printed `NaN`, `+Inf` and `-Inf` instead of `nan`, `infinity`
	and `neg_infinity`. haxe.Json used Go's JSON number format, and one NaN
	turned the whole document into `null`.

	Every path below must print what `--interp` prints (expected.stdout):
	Std.string, string concatenation and interpolation, Sys.println,
	StringBuf.add, Float values held in Dynamic, Arrays, enum arguments,
	anonymous objects and class fields, typed Floats widened from Int, trace's
	formatOutput and haxe.Json. The semantic-diff case
	float_string_format_contract also covers haxe.Serializer. Left out: a
	negative zero written as a literal (`-0.0`), which haxe.go still emits as
	a Go constant that equals +0.
**/
class Main {
	static function values():Array<Float> {
		var zero = 0.0;
		return [
			0.0,
			-zero,
			1.0,
			-1.0,
			0.5,
			1.5,
			100.0,
			3.14159,
			0.1,
			0.000123,
			0.0001234567,
			0.99609375,
			1e5,
			1e6,
			1e7,
			12345678.9,
			1e8,
			-1e8,
			123456789.0,
			999999999.0,
			2147483647.0,
			-2147483648.0,
			4294967296.0,
			1e9 + 0.5,
			123456789.123,
			100000000000.5,
			999999999999.0,
			1e12,
			-1e12,
			1e12 + 1,
			2e12,
			1234567890000.0,
			99999999999999.0,
			1e14 + 0.25,
			999999999999999.0,
			1e15,
			1e15 + 0.5,
			9007199254740992.0,
			1e16,
			1e17,
			123456789012345678.0,
			1234567890123456789.0,
			1e20,
			1e21,
			-1e21,
			1e22,
			1e23,
			1e100,
			1.5e300,
			1.7976931348623157e308,
			2.2250738585072014e-308,
			5e-324,
			0.1 + 0.2,
			1 / 3,
			-2 / 3,
			0.1 * 3,
			1e-4,
			1e-5,
			1e-6,
			1e-7,
			1.5e-7,
			2.5e-5,
			9e-6,
			1e-10,
			123e-20,
			1e-300,
			Math.NaN,
			Math.POSITIVE_INFINITY,
			Math.NEGATIVE_INFINITY
		];
	}

	static function main() {
		var floats = values();
		for (value in floats) {
			var dyn:Dynamic = value;
			var buffer = new StringBuf();
			buffer.add(value);
			Sys.println(Std.string(value) + " | " + ("" + value) + " | " + '$value' + " | " + buffer.toString() + " | " + Std.string(dyn) + " | "
				+ Std.string([value]) + " | " + Std.string(Hold(value)) + " | " + Std.string({
					value: value
				}) + " | " + haxe.Log.formatOutput(value, null));
		}

		Sys.println("--- Sys.println");
		for (value in floats) {
			Sys.println(value);
		}

		Sys.println("--- typed");
		var count:Int = 2147483647;
		var widened:Float = count;
		Sys.println(widened);
		Sys.println("widened=" + widened);
		var lowest:Float = -2147483648;
		Sys.println('lowest=$lowest');
		var sample = new Sample(3e8);
		Sys.println("field=" + sample.value + " half=" + sample.half());
		var maybe:Null<Float> = 1e8;
		Sys.println("nullable=" + maybe);
		var floatArray:Array<Float> = [1e8, 0.1 + 0.2, 1e-7];
		Sys.println("array=" + floatArray);
		var total = 0.0;
		for (step in 0...10) {
			total += 1e8;
		}
		Sys.println("sum=" + total);

		Sys.println("--- json");
		var zero = 0.0;
		var jsonValues:Array<Float> = [1e8, 2147483647.0, 0.1 + 0.2, 1e15, 1e21, 1e-7, -zero, 1.5];
		Sys.println(haxe.Json.stringify(jsonValues));
		Sys.println(haxe.Json.stringify({big: 1e8}));
		var nonFinite:Array<Float> = [1.5, Math.NaN, Math.POSITIVE_INFINITY];
		Sys.println(haxe.Json.stringify(nonFinite));
	}
}
