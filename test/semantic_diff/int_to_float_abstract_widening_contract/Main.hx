abstract Real(Float) from Float to Float {}

abstract Meters(Float) from Float from Int to Float {
	public inline function halfInline():Float
		return this / 2;

	public function halfCall():Float
		return this / 2;
}

abstract Distance(Real) from Real to Real {}
abstract Count(Int) from Int to Int {}

enum abstract Level(Int) to Int {
	var Low = 1;
	var High = 3;
}

enum abstract Code(Int) from Int to Int {
	var Ok = 200;
}

class Track {
	public var length:Real;
	public var meters:Meters;
	public var total:Float;
	public var maybe:Null<Real>;

	public function new() {}
}

/**
	An Int value reaches a Float through an abstract or a `Dynamic` value
	(haxe_go-gjqe).

	Haxe stores an abstract's value as its underlying type and adds no
	conversion when an Int enters a slot of an abstract over Float (`from Int`,
	or `from Float` that Int unifies with), when an abstract over Int (an enum
	abstract too) enters a Float slot, or when an unsafe `cast` or a `Dynamic`
	value puts an Int into an abstract over Float. `--interp` keeps every such
	value usable as a Float. Each line prints the value, half of it, two
	comparisons and its `Std.string` text, which must match `--interp`.

	UInt and `haxe.Int32` reach Float through their own `@:to` conversions;
	their lines guard that the abstract-following widening leaves them alone.
	`haxe.Int64` has no implicit conversion to Float, so it has no line here.
**/
class Main {
	static function show(label:String, value:Float):Void {
		Sys.println(label
			+ ": "
			+ value
			+ " half="
			+ (value / 2)
			+ " over="
			+ (value > 2.5)
			+ " seven="
			+ (value == 7)
			+ " text="
			+ Std.string(value));
	}

	static function halfReal(value:Real):Float
		return (value : Float) / 2;

	static function countAsFloat(count:Count):Float
		return count;

	static function intAsReal(n:Int):Real
		return n;

	static function levelAsFloat(level:Level):Float
		return level;

	static function unsafeFloat<T>(value:T):Float
		return cast value;

	static function main() {
		var seven:Dynamic = 7;
		var three = 3;

		// A Dynamic or Any value that holds an Int, read as a Float.
		var direct:Float = seven;
		show("dynamic local", direct);
		show("dynamic cast", cast seven);
		show("dynamic check", (seven : Float));
		var anyNine:Any = 9;
		var fromAny:Float = anyNine;
		show("any local", fromAny);

		// An abstract over Float that holds an Int from a Dynamic value.
		var unsafeReal:Real = cast seven;
		show("real from unsafe cast", unsafeReal);
		var dynamicReal:Real = seven;
		show("real from dynamic", dynamicReal);
		show("real from checked cast", cast(seven, Real));
		var distance:Distance = cast seven;
		show("abstract over real", (distance : Real));
		var castMeters:Meters = cast seven;
		show("inline method of cast meters", castMeters.halfInline());
		show("method of cast meters", castMeters.halfCall());
		show("type parameter unsafe cast", unsafeFloat(7));
		Sys.println("real text: " + Std.string(unsafeReal) + " " + unsafeReal);

		// Int values that enter abstracts over Float.
		var literalMeters:Meters = 7;
		show("meters from int literal", literalMeters);
		var localMeters:Meters = three;
		show("meters from int local", localMeters);
		show("meters inline method", localMeters.halfInline());
		show("meters method", localMeters.halfCall());
		var localReal:Real = three;
		show("real from int local", localReal);
		show("int argument for real", halfReal(three));
		show("int return as real", intAsReal(9));
		var track = new Track();
		track.length = three;
		track.meters = 4;
		show("real field", track.length);
		show("meters field", track.meters);
		var reals:Array<Real> = [three, 1.5];
		reals.push(5);
		show("real array literal", reals[0]);
		show("real array push", reals[2]);
		var realMap = new Map<String, Real>();
		realMap.set("three", three);
		show("real map value", realMap.get("three"));
		var maybeReal:Null<Real> = three;
		show("nullable real", maybeReal);
		track.maybe = 5;
		show("nullable real field", track.maybe);
		var noReal:Null<Real> = null;
		Sys.println("null real: " + (noReal == null));
		var castNullable:Null<Real> = cast seven;
		var asNullableFloat:Null<Float> = castNullable;
		Sys.println("nullable real to nullable float: " + (asNullableFloat + 0.5));

		// Abstracts over Int, enum abstracts included, that enter Float slots.
		var count:Count = 5;
		var countLocal:Float = count;
		show("count local", countLocal);
		show("count argument", count);
		show("count return", countAsFloat(count));
		var levelLocal:Float = High;
		show("level local", levelLocal);
		show("level argument", Low);
		show("level return", levelAsFloat(High));
		var code:Code = 404;
		var codeFloat:Float = code;
		show("code from int", codeFloat);
		show("code constant", Ok);
		var floats:Array<Float> = [count, High, 0.5];
		show("float array of count", floats[0]);
		show("float array of level", floats[1]);
		track.total = count;
		show("float field of count", track.total);
		var maybeCount:Null<Float> = count;
		show("nullable float of count", maybeCount);
		var maybeLevel:Null<Float> = High;
		show("nullable float of level", maybeLevel);
		var floatMap = new Map<String, Float>();
		floatMap.set("count", count);
		show("float map of count", floatMap.get("count"));
		Sys.println("sum of widened values: " + (countLocal + levelLocal + codeFloat + (unsafeReal : Float)));

		// UInt and Int32 keep their own conversions.
		var big:UInt = cast 0xFFFFFFFF;
		var bigFloat:Float = big;
		show("uint", bigFloat);
		var small:haxe.Int32 = 6;
		var smallFloat:Float = small;
		show("int32", smallFloat);
	}
}
