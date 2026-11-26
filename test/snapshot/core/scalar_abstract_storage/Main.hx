abstract Count(Int) from Int to Int {}

abstract Score(Int) from Int to Int {
	@:op(A + B) static function add(a:Score, b:Score):Score;

	public function doubled():Score {
		return this * 2;
	}
}

abstract Meters(Float) from Float to Float {
	@:op(A + B) static function add(a:Meters, b:Meters):Meters;
}

abstract Flag(Bool) from Bool to Bool {
	@:op(!A) static function not(a:Flag):Flag;
}

abstract Box<T>(T) from T to T {}

enum abstract Level(Int) to Int {
	var Low = 1;
	var High = 3;
}

class Counter {
	public static var start:Count = 7;

	public var count:Count;
	public var meters:Meters;
	public var flag:Flag;
	public var maybe:Null<Count>;
	public var byKey:Map<Count, String>;

	public function new(count:Count) {
		this.count = count;
		this.meters = 1.5;
		this.flag = true;
		this.maybe = null;
		this.byKey = new Map();
	}
}

/**
	Values of abstracts over Int, Float and Bool use the Go storage of the
	underlying type (`int`, `float64`, `bool`) in locals, parameters, returns,
	fields, statics, Array and Map elements, function values, applied type
	parameters and enum abstract constants (haxe_go-5yxq). Int arithmetic and
	forwarded `@:op` operators on them build and wrap like Int arithmetic
	(haxe_go-ymvq). `Null<T>` of such an abstract keeps nil-capable `any`, a
	checked cast of an erased `Null<Float>` reads its `any` storage, a `cast`
	of a `Dynamic` Int into `Null<T>` of an abstract over Float stores a
	float64, and a null passes a checked cast to such an abstract.
**/
class Main {
	static function next(c:Count):Count {
		return c + 1;
	}

	static function optional(?c:Count):String {
		return c == null ? "none" : Std.string(c);
	}

	static function main() {
		var count:Count = 5;
		var plusOne:Int = count + 1;
		var low = Low + 2;
		var score:Score = 4;
		var meters:Meters = 1.5;
		var flag:Flag = true;
		var big:Count = 2147483647;
		var wrapped:Int = big + 1;
		Sys.println("ops="
			+ plusOne
			+ ","
			+ low
			+ ","
			+ (score + score)
			+ ","
			+ score.doubled()
			+ ","
			+ (meters + meters)
			+ ","
			+ (!flag)
			+ ","
			+ wrapped);

		var counter = new Counter(next(count));
		counter.count += 2;
		Counter.start = Counter.start + 1;
		counter.byKey.set(count, "five");
		Sys.println("fields=" + counter.count + "," + counter.meters + "," + counter.flag + "," + counter.maybe + "," + Counter.start + ","
			+ counter.byKey.get(5));

		var counts:Array<Count> = [3, 1, 2];
		counts.push(count);
		var levels:Map<String, Level> = ["top" => High];
		var inc:Count->Count = c -> c + 1;
		var box:Box<Int> = 9;
		var maybe:Null<Count> = null;
		var dyn:Dynamic = count;
		var back:Count = dyn;
		Sys.println("values=" + counts.join("|") + "," + (counts[3] + 1) + "," + levels.get("top") + "," + inc(count) + "," + ((box : Int) + 1) + ","
			+ optional() + "," + optional(4) + "," + maybe + "," + (back + 1));

		// Casts of nullable and Dynamic values keep the scalar storage.
		var quarter:Null<Float> = 2.25;
		var seven:Dynamic = 7;
		var castMeters:Null<Meters> = cast seven;
		var widened:Null<Float> = castMeters;
		var nothing:Dynamic = null;
		var nullCast = try {
			cast(nothing, Count);
			"passed";
		} catch (e:Dynamic) {
			"threw";
		}
		Sys.println("casts=" + ((cast(quarter, Meters) : Float) * 2) + "," + (widened + 0.5) + "," + nullCast);
	}
}
