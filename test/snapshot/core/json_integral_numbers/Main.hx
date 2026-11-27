typedef Reading = {count:Int, price:Float, ?discount:Float, samples:Array<Float>};

/**
	JSON numbers without a fraction are Haxe Ints (haxe_go-2bge).

	haxe.Json.parse returns an Int for `4`, `-0` and `1e3`, as Haxe's
	JsonParser does in `--interp`, and a Float for `4.0`, `2147483648` and
	`1e20` (expected.stdout is the `--interp` output). haxe.go returned a Go
	float64 for every number, so the typed `count` read below panicked.

	A Float structure field, `Null<Float>` field, Array element or sort
	argument can now hold that Int. The generated Go reads each of them with
	`hxrt.DynamicToFloatSlot` or `hxrt.DynamicToNullableFloat`, which convert
	an Int, where it asserted `.(float64)` before. The Int field keeps its
	plain `.(int)` assertion.
**/
class Main {
	static function main() {
		var reading:Reading = haxe.Json.parse('{"count":4,"price":10,"discount":2,"samples":[3,1,2.5]}');
		Sys.println("count+1=" + (reading.count + 1));
		Sys.println("price+0.5=" + (reading.price + 0.5));
		Sys.println("discount/4=" + (reading.discount / 4));
		var samples = reading.samples;
		Sys.println("samples[0]+0.5=" + (samples[0] + 0.5));
		samples.sort((left, right) -> left < right ? -1 : (left > right ? 1 : 0));
		Sys.println("sorted=" + samples);

		for (text in ["4", "-0", "1e3", "4.0", "2147483647", "2147483648", "1e20"]) {
			var value:Dynamic = haxe.Json.parse(text);
			Sys.println(text + " typeof=" + Type.typeof(value) + " str=" + Std.string(value) + " json=" + haxe.Json.stringify(value));
		}
	}
}
