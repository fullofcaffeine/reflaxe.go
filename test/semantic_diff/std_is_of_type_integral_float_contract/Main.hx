/**
	`Std.isOfType(value, Int)` is true for a Float with an Int value (haxe_go-jxxd).

	`--interp`, neko, JS, Java, Python and PHP all answer `Std.isOfType(4.0,
	Int)` with true. On `--interp`, a Float counts as an Int when it is integral
	and within the 32-bit range: `4.0`, `-0.0`, `0.0`, `2147483647.0` and
	`-2147483648.0`, but not `2147483648.0`, NaN or an infinity. `Type.typeof`
	still reports `TFloat` for such a value, and a checked `cast(value, Int)`
	accepts it. haxe.go answered false for every Float, from a `Dynamic`,
	`Any` or `Null<Float>` value, from a value typed Float, and through a
	runtime `Int` token.
**/
class Main {
	static var calls = 0;

	/** Returns a Float and counts the calls, to show the test reads its operand once. */
	static function nextFloat():Float {
		calls++;
		return 3.0;
	}

	/**
		Names the ValueType of a value with a switch on Type.typeof. Printing a
		`TClass` value directly shows a pointer on haxe.go (haxe_go-snqo).
	**/
	static function kind(value:Dynamic):String {
		return switch (Type.typeof(value)) {
			case TInt: "TInt";
			case TFloat: "TFloat";
			case TNull: "TNull";
			case TBool: "TBool";
			case TClass(_): "TClass";
			case _: "other";
		};
	}

	static function describe(label:String, value:Dynamic):Void {
		Sys.println(label + ": int=" + Std.isOfType(value, Int) + " float=" + Std.isOfType(value, Float) + " typeof=" + kind(value) + " str="
			+ Std.string(value));
	}

	static function main() {
		var zero = 0.0;
		var negativeZero = -zero;
		var nan = Math.NaN;
		var labels = [
			"4.0",
			"-0.0",
			"0.0",
			"2147483647.0",
			"-2147483648.0",
			"2147483648.0",
			"-2147483649.0",
			"4294967296.0",
			"4294967297.0",
			"NaN",
			"+Inf",
			"-Inf",
			"1e20",
			"0.5",
			"-1.5",
			"2^53",
			"1/3*3",
			"int 4",
			"int -7",
			"null",
			"string 4",
			"true"
		];
		var values:Array<Dynamic> = [
			4.0,
			negativeZero,
			0.0,
			2147483647.0,
			-2147483648.0,
			2147483648.0,
			-2147483649.0,
			4294967296.0,
			4294967297.0,
			nan,
			Math.POSITIVE_INFINITY,
			Math.NEGATIVE_INFINITY,
			1e20,
			0.5,
			-1.5,
			9007199254740992.0,
			1.0 / 3.0 * 3.0,
			4,
			-7,
			null,
			"4",
			true
		];
		for (index in 0...values.length)
			describe(labels[index], values[index]);

		var typed:Float = 4.0;
		var fraction:Float = 4.5;
		Sys.println("Float 4.0: "
			+ Std.isOfType(typed, Int)
			+ ", Float 4.5: "
			+ Std.isOfType(fraction, Int)
			+ ", literal 10.0: "
			+ Std.isOfType(10.0, Int));
		var optional:Null<Float> = 5.0;
		var missing:Null<Float> = null;
		Sys.println("Null<Float> 5.0: " + Std.isOfType(optional, Int) + ", Null<Float> null: " + Std.isOfType(missing, Int));
		var any:Any = 6.0;
		Sys.println("Any 6.0: " + Std.isOfType(any, Int));
		Sys.println("call: " + Std.isOfType(nextFloat(), Int) + " calls=" + calls);

		var token:Dynamic = Int;
		Sys.println("token 7.0: "
			+ Std.isOfType(7.0, token)
			+ ", token 7.5: "
			+ Std.isOfType(7.5, token)
			+ ", token 7: "
			+ Std.isOfType(7, token));
		var floatToken:Dynamic = Float;
		Sys.println("Float token 7: " + Std.isOfType(7, floatToken) + ", Float token 7.5: " + Std.isOfType(7.5, floatToken));

		var whole:Dynamic = 9.0;
		var castWhole:Int = cast(whole, Int);
		Sys.println("cast(9.0, Int) = " + castWhole + " plus one = " + (castWhole + 1));
		try {
			var half:Dynamic = 9.5;
			var castHalf:Int = cast(half, Int);
			Sys.println("cast(9.5, Int) = " + castHalf);
		} catch (error:String) {
			Sys.println("cast(9.5, Int) throws " + error);
		}

		var parsed:Dynamic = haxe.Json.parse('[4.0,-0.0,0.0,2.5,4]');
		Sys.println("JSON: " + [
			for (index in 0...5)
				Std.isOfType(parsed[index], Int) + "/" + kind(parsed[index])
		].join(", "));
	}
}
