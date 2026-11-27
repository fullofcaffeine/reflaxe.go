typedef Counter = {count:Int};
typedef Priced = {price:Float, ?discount:Float, ?limit:Int};
typedef Item = {id:Int, weight:Float};
typedef Nested = {outer:{inner:{count:Int, ratio:Float}}, items:Array<Item>};

/**
	JSON numbers without a fraction are Haxe Ints (haxe_go-2bge).

	Haxe's JsonParser, which `--interp`, neko, Java, C++ and HashLink use,
	returns an Int for a number that has no `.`, has an integral value and
	fits in 32 bits (`4`, `-0`, `1e3`). Every other number is a Float (`4.0`,
	`2147483648`, `1e20`). haxe.go returned a Go float64 for every number, so
	a typed `{count:Int}` read panicked, `Type.typeof` said `TFloat`, the
	Serializer wrote `d4` and `-0` printed as `-0`. A Float field or Array
	element that receives a JSON Int must still read as a Float.
**/
class Main {
	/** Prints how a parsed number answers the Int and Float type tests. */
	static function describe(text:String):Void {
		var value:Dynamic = haxe.Json.parse(text);
		Sys.println(text + " str=" + Std.string(value) + " int=" + Std.isOfType(value, Int) + " float=" + Std.isOfType(value, Float) + " typeof="
			+ Std.string(Type.typeof(value)) + " json=" + haxe.Json.stringify(value));
	}

	/**
		Prints a parsed number written with a fraction. `--interp` also answers
		`Std.isOfType(4.0, Int)` with true for any integral Float; that check is
		not about JSON, so it is left out here.
	**/
	static function describeFraction(text:String):Void {
		var value:Dynamic = haxe.Json.parse(text);
		Sys.println(text + " str=" + Std.string(value) + " float=" + Std.isOfType(value, Float) + " typeof=" + Std.string(Type.typeof(value)) + " json="
			+ haxe.Json.stringify(value));
	}

	static function main() {
		for (text in [
			"4",
			"0",
			"-5",
			"1e3",
			"1E2",
			"1e+2",
			"10e-1",
			"25e-1",
			"1.5",
			"-0",
			"2147483647",
			"2147483648",
			"-2147483648",
			"-2147483649",
			"4294967296",
			"9007199254740993",
			"12345678901234567890",
			"1e20",
			"1e-400",
			"1e400",
			"-1e400"
		])
			describe(text);
		for (text in ["4.0", "-0.0", "0.0", "2147483647.0"])
			describeFraction(text);

		var counter:Counter = haxe.Json.parse('{"count":4}');
		Sys.println("counter.count+1=" + (counter.count + 1));
		Sys.println("counter.count*3=" + (counter.count * 3));
		Sys.println("counter.count/8=" + (counter.count / 8));
		Sys.println("counter.json=" + haxe.Json.stringify(counter));
		var back:Counter = haxe.Json.parse(haxe.Json.stringify({count: 41}));
		Sys.println("back.count+1=" + (back.count + 1));

		var priced:Priced = haxe.Json.parse('{"price":10,"discount":2,"limit":3}');
		Sys.println("priced.price+0.5=" + (priced.price + 0.5));
		Sys.println("priced.price=" + priced.price);
		Sys.println("priced.discount=" + priced.discount);
		var discount:Float = priced.discount;
		Sys.println("priced.discount/4=" + (discount / 4));
		Sys.println("priced.limit+1=" + (priced.limit + 1));
		var sparse:Priced = haxe.Json.parse('{"price":1.25}');
		Sys.println("sparse.price*2=" + (sparse.price * 2));
		Sys.println("sparse.discount=" + sparse.discount + " sparse.limit=" + sparse.limit);

		var dynamicValue:Dynamic = haxe.Json.parse('{"count":4,"ratio":2.5,"whole":3.0,"zero":-0}');
		Sys.println("dynamic.count+1=" + (dynamicValue.count + 1));
		Sys.println("dynamic.count*2=" + (dynamicValue.count * 2));
		Sys.println("dynamic.count/8=" + (dynamicValue.count / 8));
		Sys.println("dynamic.ratio*2=" + (dynamicValue.ratio * 2));
		Sys.println("dynamic.count==4=" + (dynamicValue.count == 4));
		Sys.println("dynamic.count typeof=" + Type.typeof(dynamicValue.count) + " whole typeof=" + Type.typeof(dynamicValue.whole) + " zero typeof="
			+ Type.typeof(dynamicValue.zero));
		Sys.println("dynamic.whole=" + dynamicValue.whole + " dynamic.zero=" + dynamicValue.zero);
		var count:Int = dynamicValue.count;
		Sys.println("count<<1=" + (count << 1));
		var widened:Float = dynamicValue.count;
		Sys.println("widened/8=" + (widened / 8));
		Sys.println("Std.int(ratio)=" + Std.int(dynamicValue.ratio));
		Sys.println("reflect count isInt=" + Std.isOfType(Reflect.field(dynamicValue, "count"), Int));

		var ints:Array<Int> = haxe.Json.parse('[1,2,3]');
		var intTotal = 0;
		for (value in ints)
			intTotal += value;
		Sys.println("ints.total=" + intTotal + " ints[0]+ints[2]=" + (ints[0] + ints[2]) + " ints=" + ints);
		var floats:Array<Float> = haxe.Json.parse('[1,2.5,3]');
		var floatTotal = 0.0;
		for (value in floats)
			floatTotal += value;
		Sys.println("floats.total=" + floatTotal + " floats[0]+0.5=" + (floats[0] + 0.5) + " floats=" + floats);
		Sys.println("floats.map=" + floats.map(value -> value * 2));
		var unsorted:Array<Float> = haxe.Json.parse('[3,1,2.5]');
		var sorted = unsorted.copy();
		sorted.sort((left, right) -> left < right ? -1 : (left > right ? 1 : 0));
		Sys.println("sorted="
			+ sorted
			+ " filter="
			+ unsorted.filter(value -> value > 1)
			+ " fold="
			+ Lambda.fold(unsorted, (value, total) -> value + total, 0.5)
			+ " indexOf(1)="
			+ unsorted.indexOf(1));
		unsorted[0] += 0.5;
		var popped = unsorted.copy().pop();
		var shifted = unsorted.copy().shift();
		Sys.println("compound=" + unsorted[0] + " pop+0.25=" + (popped + 0.25) + " shift+0.25=" + (shifted + 0.25));
		var optional:Array<Null<Float>> = haxe.Json.parse('[1,null]');
		var first:Null<Float> = optional[0];
		Sys.println("optional[0]+0.5=" + (optional[0] + 0.5) + " first>0.5=" + (first > 0.5) + " optional[1]==null=" + (optional[1] == null));
		var mixed:Array<Dynamic> = haxe.Json.parse('[1,1.5,-0,2147483648]');
		Sys.println("mixed typeof=" + [for (value in mixed) Std.string(Type.typeof(value))].join(","));

		var nested:Nested = haxe.Json.parse('{"outer":{"inner":{"count":7,"ratio":1}},"items":[{"id":1,"weight":2},{"id":2,"weight":0.5}]}');
		Sys.println("nested.count+1=" + (nested.outer.inner.count + 1));
		Sys.println("nested.ratio/4=" + (nested.outer.inner.ratio / 4));
		var ids = 0;
		var weight = 0.0;
		for (item in nested.items) {
			ids += item.id;
			weight += item.weight;
		}
		Sys.println("nested.ids=" + ids + " nested.weight=" + weight);

		var access:haxe.DynamicAccess<Int> = haxe.Json.parse('{"x":1,"y":2}');
		Sys.println("access.sum=" + (access.get("x") + access.get("y")));

		// A `Dynamic` field write and Reflect.setField store an Int as it is, too.
		var written:Dynamic = {};
		written.price = 4;
		Reflect.setField(written, "rate", 2);
		var rated:{price:Float, rate:Float} = written;
		Sys.println("written.price+0.5=" + (rated.price + 0.5) + " written.rate/4=" + (rated.rate / 4));
		var optionalPriced:Priced = haxe.Json.parse('{"price":4,"discount":4}');
		Sys.println("optionalPriced.discount>3=" + (optionalPriced.discount > 3) + " discount+1=" + (optionalPriced.discount + 1));

		Sys.println("serialized=" + haxe.Serializer.run(haxe.Json.parse('[4,1.5,-0,2147483648,1e3,-2147483648]')));
		var unserialized:Dynamic = haxe.Unserializer.run(haxe.Serializer.run(haxe.Json.parse('{"count":4}')));
		Sys.println("unserialized.count typeof=" + Type.typeof(unserialized.count));

		// Objects with one field each: the field order of a larger parsed object differs between targets.
		var document = '[4,4.5,[1,2.0,-0,1e3,2147483648],{"e":-0.0},{"f":-7},{"g":[3,0.5]}]';
		var once = haxe.Json.stringify(haxe.Json.parse(document));
		Sys.println("roundtrip=" + once);
		// `-0.0` prints as `-0`, which parses back as the Int 0.
		Sys.println("roundtrip.twice=" + haxe.Json.stringify(haxe.Json.parse(once)));
	}
}
