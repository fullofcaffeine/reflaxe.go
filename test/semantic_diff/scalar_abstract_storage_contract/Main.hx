import haxe.ds.Vector;

/** The reported shape (haxe_go-ymvq): an abstract over Int with casts both ways. */
abstract Count(Int) from Int to Int {}

/** An abstract over Int with forwarded and implemented operators and methods. */
abstract Score(Int) from Int to Int {
	public inline function new(value:Int) {
		this = value;
	}

	@:op(A + B) static function add(a:Score, b:Score):Score;

	@:op(A * B) static function times(a:Score, b:Int):Score;

	@:op(A > B) static function above(a:Score, b:Score):Bool;

	@:op(-A) static function negate(a:Score):Score;

	@:op(A - B) function minus(b:Score):Score {
		return this - (b : Int);
	}

	@:op(A++) inline function increment():Score {
		return this
		++;
	}

	public function doubled():Score {
		return new Score(this * 2);
	}

	public inline function plus(n:Int):Score {
		return this + n;
	}

	public function toInt():Int {
		return this;
	}
}

/** An abstract over Float with forwarded and implemented operators. */
abstract Meters(Float) from Float to Float {
	@:op(A + B) static function add(a:Meters, b:Meters):Meters;

	@:op(A / B) static function ratio(a:Meters, b:Meters):Float;

	@:op(A < B) static function below(a:Meters, b:Meters):Bool;

	@:op(-A) static function negate(a:Meters):Meters;

	@:op(A * B) static function scale(a:Meters, b:Float):Meters {
		return (a : Float) * b;
	}

	public inline function half():Meters {
		return this / 2;
	}

	public function km():Float {
		return this / 1000;
	}
}

/** An Int enters an abstract over Float through `from Int`. */
abstract Real(Float) from Float from Int to Float {}

/** An abstract over Bool with forwarded operators. */
abstract Flag(Bool) from Bool to Bool {
	@:op(!A) static function not(a:Flag):Flag;

	@:op(A && B) static function and(a:Flag, b:Flag):Flag;

	@:op(A || B) static function or(a:Flag, b:Flag):Flag;

	public function text():String {
		return this ? "on" : "off";
	}
}

/** User-defined `@:from` and `@:to` conversions on an abstract over Int. */
abstract Percent(Int) {
	public inline function new(value:Int) {
		this = value;
	}

	@:from static function fromFloat(f:Float):Percent {
		return new Percent(Math.round(f * 100));
	}

	@:to function toFloat():Float {
		return this / 100;
	}

	@:to function toText():String {
		return this + "%";
	}
}

/** An abstract over another abstract over Int. */
abstract Total(Count) from Count to Count from Int to Int {}

/** An applied type parameter: `Box<Int>` stores an Int. */
abstract Box<T>(T) from T to T {}

/** Control: an abstract over `Null<Int>` can hold null and keeps nil-capable storage. */
abstract MaybeId(Null<Int>) from Null<Int> to Null<Int> {}

enum abstract Level(Int) to Int {
	var Low = 1;
	var Mid = 2;
	var High = 3;
}

enum abstract Code(Int) from Int to Int {
	var Ok = 200;
	var Missing = 404;
}

enum abstract Ratio(Float) to Float {
	var Half = 0.5;
	var Whole = 1.0;
}

enum abstract Toggle(Bool) to Bool {
	var On = true;
	var Off = false;
}

enum Slot {
	Filled(count:Count, level:Level);
	Empty;
}

typedef Entry = {
	count:Count,
	?meters:Meters,
	flag:Flag
}

interface Scorer {
	function score(c:Count):Score;
	function level():Level;
}

class Base {
	public function new() {}

	public function value():Count {
		return 1;
	}
}

class Derived extends Base {
	override public function value():Count {
		return 2;
	}
}

class Holder<T> {
	public var item:T;

	public function new(item:T) {
		this.item = item;
	}

	public function get():T {
		return item;
	}
}

class Player implements Scorer {
	public static var defaultCount:Count = 7;
	public static var top:Level = High;
	public static inline var LIMIT:Count = 9;
	public static final RATE:Meters = 2.5;
	public static var byCode:Map<Count, String> = new Map();

	public var count:Count;
	public var meters:Meters;
	public var flag:Flag;
	public var lvl:Level;
	public var maybe:Null<Count>;
	public var maybeFlag:Null<Flag>;
	public var history:Array<Count>;

	public function new(count:Count, ?lvl:Level, flag:Flag = true) {
		this.count = count;
		this.meters = 1.5;
		this.flag = flag;
		this.lvl = lvl == null ? Low : lvl;
		this.maybe = null;
		this.maybeFlag = null;
		this.history = [count];
	}

	public function score(c:Count):Score {
		return (c : Int) + count;
	}

	public function level():Level {
		return lvl;
	}
}

/**
	Values of abstracts over Int, Float and Bool are stored as Go `int`,
	`float64` and `bool` wherever Haxe treats them as those types
	(haxe_go-5yxq). Before this change they were stored as Go `any`, and Int
	arithmetic or a forwarded `@:op` on them produced Go that did not build
	(haxe_go-ymvq): `count + 1`, `Low + 2`, `score + score`, `meters + meters`.

	Each line covers one shape: arithmetic and operators, locals, parameters
	(plain, optional, default, `Null<T>`), returns, instance and static
	fields, Arrays, Maps, `Vector`, closures, type parameters, enum
	arguments, structures, interfaces, `Dynamic` and `Any` boundaries,
	reflection, JSON, serialization and `Null<T>` wrapping. Every line must
	match `--interp`.
**/
class Main {
	static function countOf(c:Count):Int {
		return c;
	}

	static function twice(c:Count):Count {
		return c * 2;
	}

	static function optCount(?c:Count):String {
		return c == null ? "none" : Std.string(c + 1);
	}

	static function defCount(c:Count = 3):Count {
		return c;
	}

	static function nullCount(c:Null<Count>):String {
		return c == null ? "null" : Std.string(c);
	}

	static function optLevel(?l:Level):String {
		return l == null ? "none" : Std.string(l);
	}

	static function defLevel(l:Level = Mid):Int {
		return l + 0;
	}

	static function optFlag(?f:Flag):String {
		return f == null ? "none" : f.text();
	}

	static function defFlag(f:Flag = true):Flag {
		return !f;
	}

	static function defMeters(m:Meters = 1.5):Meters {
		return m + m;
	}

	static function realHalf(r:Real):Float {
		return (r : Float) / 2;
	}

	static function apply(f:Count->Count, c:Count):Count {
		return f(c);
	}

	static function makeAdder(n:Count):Count->Count {
		return c -> (c : Int) + n;
	}

	static function identity<T>(value:T):T {
		return value;
	}

	static function first<T>(values:Array<T>):T {
		return values[0];
	}

	static function describe(value:Dynamic):String {
		return Std.string(value);
	}

	static function levelName(l:Level):String {
		return switch l {
			case Low: "low";
			case Mid: "mid";
			case High: "high";
			case _: "other";
		}
	}

	static function codeName(c:Code):String {
		return switch c {
			case Ok: "ok";
			case Missing: "missing";
			case _: "code" + (c : Int);
		}
	}

	static function countName(c:Count):String {
		return switch c {
			case 0: "zero";
			case 1 | 2: "few";
			case _: "many";
		}
	}

	static function toggleName(t:Toggle):String {
		return switch t {
			case On: "ON";
			case Off: "OFF";
		}
	}

	static function main() {
		// The reported shapes (haxe_go-ymvq).
		var count:Count = 5;
		var plusOne:Int = count + 1;
		var low = Low + 2;
		var score:Score = 4;
		var scoreSum = score + score;
		var meters:Meters = 1.5;
		var meterSum = meters + meters;
		var big:Count = 2147483647;
		var wrapped:Int = big + 1;
		var bigScore:Score = 2147483647;
		var wrappedScore = bigScore + bigScore;
		Sys.println("repro=" + plusOne + "," + low + "," + scoreSum + "," + meterSum + "," + wrapped + "," + wrappedScore);

		// Int arithmetic, bit operations and compound assignments.
		var other:Count = 3;
		var arith = [
			count * count,
			count - 2,
			count % 3,
			count & 1,
			count | 8,
			count ^ 3,
			count << 2,
			count >> 1,
			-count,
			~count,
			count * other,
			count >>> 1
		];
		Sys.println("arith=" + arith.join("|") + "," + (count / 2) + "," + (count + 0.5));
		var counter:Count = 1;
		counter += 2;
		counter -= 1;
		counter *= 3;
		counter++;
		++counter;
		counter--;
		var post = counter++;
		Sys.println("compound=" + counter + "," + post);
		Sys.println("compare=" + ((count : Int) > 3) + "," + (count == 5) + "," + (count != other) + "," + ((other : Int) < (count : Int)) + ","
			+ ((Low : Int) < (High : Int)) + "," + (Player.top == High) + "," + ((count : Int) >= 5) + "," + (Low == Low));

		// Operators of Score, forwarded and implemented, and methods using `this`.
		var s3 = scoreSum - 1;
		s3++;
		var scaled = score * 3;
		var negative = -score;
		Sys.println("score=" + s3 + "," + scaled + "," + negative + "," + (scoreSum > score) + "," + s3.doubled() + "," + s3.plus(2) + "," + s3.toInt()
			+ "," + new Score(6).doubled());

		// Float abstracts.
		var other2:Meters = 0.25;
		Sys.println("meters=" + (meters * 2) + "," + (meters / other2) + "," + (other2 < meters) + "," + (-meters) + "," + meters.half() + "," + meters.km()
			+ "," + Math.floor(meters) + "," + Std.int(meters) + "," + ((meters : Float) + 1) + "," + (meters * meters));
		var real:Real = 3;
		var realFloat:Real = 2.5;
		Sys.println("real="
			+ realHalf(real)
			+ ","
			+ realHalf(7)
			+ ","
			+ (real : Float)
			+ ","
			+ realFloat
			+ ","
			+ ((real : Float) + (realFloat : Float)));

		// Bool abstracts.
		var flag:Flag = true;
		var off:Flag = false;
		var loops = 0;
		var running:Flag = true;
		while (running) {
			loops++;
			running = loops < 3;
		}
		Sys.println("flag=" + (!flag) + "," + (flag && off) + "," + (flag || off) + "," + (flag ? 1 : 2) + "," + flag.text() + "," + off.text() + ","
			+ loops + "," + (if (off) "yes" else "no") + "," + (flag == off));

		// Enum abstracts.
		var code:Code = 404;
		var toggle:Toggle = On;
		var ratio:Ratio = Half;
		var levels:Array<Level> = [Low, High, Mid];
		Sys.println("enumAbstract=" + levelName(Mid) + "," + levelName(levels[1]) + "," + codeName(code) + "," + codeName(Ok) + "," + codeName(201) + ","
			+ toggleName(toggle) + "," + toggleName(Off) + "," + (ratio : Float) * 4 + "," + (Whole : Float) + "," + (High - Low) + ","
			+ (if (toggle) "on" else "off") + "," + levels.join("|") + "," + (code == Missing) + "," + countName(0) + "," + countName(2) + ","
			+ countName(count));

		// Parameters (plain, optional, default, Null<T>) and returns.
		var noCount:Null<Count> = null;
		Sys.println("params=" + countOf(9) + "," + twice(count) + "," + optCount() + "," + optCount(null) + "," + optCount(4) + "," + defCount() + ","
			+ defCount(8) + "," + nullCount(noCount) + "," + nullCount(6) + "," + optLevel() + "," + optLevel(High) + "," + defLevel() + "," + defLevel(Low)
			+ "," + optFlag() + "," + optFlag(false) + "," + defFlag() + "," + defFlag(false) + "," + defMeters() + "," + defMeters(0.25));

		// Instance and static fields.
		var player = new Player(4);
		var hero = new Player(6, High, false);
		player.count += 2;
		player.count++;
		player.meters = player.meters + 1.25;
		player.flag = !player.flag;
		Player.defaultCount = Player.defaultCount + 1;
		Player.top = Mid;
		player.maybe = 3;
		hero.maybeFlag = true;
		player.history.push(player.count);
		Player.byCode.set(count, "five");
		Sys.println("fields=" + player.count + "," + player.meters + "," + player.flag + "," + player.lvl + "," + hero.lvl + "," + hero.flag + ","
			+ Player.defaultCount + "," + Player.top + "," + Player.LIMIT + "," + Player.RATE + "," + player.maybe + "," + hero.maybe + ","
			+ player.maybeFlag + "," + hero.maybeFlag + "," + player.history.join("|") + "," + Player.byCode.get(5) + ","
			+ ((player.count : Int) + Player.LIMIT));

		// Arrays and Vector.
		var counts:Array<Count> = [3, 1, 2];
		counts.push(count);
		counts[1] = counts[1] + 10;
		var total = 0;
		for (c in counts) {
			total += c;
		}
		var sorted = counts.copy();
		sorted.sort((a, b) -> a - b);
		var mapped = counts.map(c -> c + 1);
		var kept = counts.filter(c -> (c : Int) > 2);
		var meterList:Array<Meters> = [1.5, 0.25];
		meterList.push(meters + meters);
		var flags:Array<Flag> = [true, false];
		var secondFlag:Flag = flags[1];
		var maybeCounts:Array<Null<Count>> = [1, null, 3];
		var grid:Array<Array<Count>> = [[1, 2], [3, 4]];
		var vector = new Vector<Count>(3);
		vector[0] = 7;
		vector[1] = vector[0] + 1;
		vector[2] = count;
		Sys.println("array="
			+ counts.join("|")
			+ ","
			+ total
			+ ","
			+ sorted.join("|")
			+ ","
			+ mapped.join("|")
			+ ","
			+ kept.join("|")
			+ ","
			+ counts.indexOf(5)
			+ ","
			+ counts.contains(11)
			+ ","
			+ (counts[0] * counts[2])
			+ ","
			+ meterList.join("|")
			+ ","
			+ (meterList[0] + meterList[1])
			+ ","
			+ flags.join("|")
			+ ","
			+ (!secondFlag)
			+ ","
			+ maybeCounts.join("|")
			+ ","
			+ (maybeCounts[1] == null)
			+ ","
			+ ((grid[1][0]:Int) + grid[0][1])
			+ ","
			+ [vector[0], vector[1], vector[2]].join("|") + "," + ((vector[1] : Int) + vector[2]));

		// Maps.
		var byName = new Map<String, Count>();
		byName.set("a", 1);
		byName["b"] = 2;
		var a:Int = byName["a"];
		var b:Count = byName["b"];
		byName["c"] = a + b;
		var mapTotal = 0;
		for (v in byName) {
			mapTotal += v;
		}
		var names = [for (k in byName.keys()) k];
		names.sort(Reflect.compare);
		var missing:Null<Count> = byName.get("zz");
		var meterMap:Map<String, Meters> = ["m" => 2.5];
		var flagMap:Map<String, Flag> = ["f" => true];
		var levelMap = new Map<Int, Level>();
		levelMap.set(1, High);
		var byLevel = new Map<Level, String>();
		byLevel.set(High, "h");
		byLevel.set(Low, "l");
		var byCount = new Map<Count, Count>();
		byCount.set(count, count + 1);
		var nextCount:Int = byCount[count];
		Sys.println("map=" + byName.get("c") + "," + mapTotal + "," + names.join("|") + "," + missing + "," + (meterMap["m"] * 2) + "," + (!flagMap["f"])
			+ "," + (levelMap.get(1) + 1) + "," + byLevel.get(High) + "," + byLevel.get(Low) + "," + byCount.get(5) + "," + byName.exists("a") + ","
			+ (nextCount + nextCount));

		// Closures and function values.
		var inc:Count->Count = c -> c + 1;
		var captured:Count = 10;
		var bump = function() {
			captured += 5;
			return captured;
		};
		bump();
		var adder = makeAdder(100);
		function local(c:Count, ?extra:Count):Count {
			return extra == null ? c : (c : Int) + (extra : Int);
		}
		var levelFn:Void->Level = () -> High;
		var flagFn:Flag->Flag = f -> !f;
		var meterFn:Meters->Meters = m -> m + m;
		var scoreFn = player.score;
		Sys.println("closure=" + inc(4) + "," + apply(inc, 7) + "," + captured + "," + bump() + "," + adder(5) + "," + local(1) + "," + local(1, 2) + ","
			+ levelFn() + "," + flagFn(true) + "," + meterFn(0.75) + "," + scoreFn(1) + "," + apply(twice, 21));

		// Type parameters instantiated with abstracts.
		var holder = new Holder<Count>(4);
		holder.item = holder.item + 1;
		var meterHolder = new Holder<Meters>(1.5);
		var intBox:Box<Int> = 9;
		var countBox:Box<Count> = (11 : Count);
		var floatBox:Box<Float> = 0.5;
		var boolBox:Box<Bool> = true;
		var folded = Lambda.fold(counts, (c, acc) -> acc + (c : Int), 0);
		var found:Null<Count> = Lambda.find(counts, c -> (c : Int) > 10);
		Sys.println("generic=" + identity(count) + "," + (identity(count) + 1) + "," + identity(meters) + "," + identity(flag) + "," + identity(High) + ","
			+ first(counts) + "," + (first(counts) * 2) + "," + holder.get() + "," + (holder.get() + 1) + "," + (meterHolder.get() + meterHolder.get()) + ","
			+ ((intBox : Int) + 1) + "," + ((countBox : Count) + 1) + "," + ((floatBox : Float) + 1) + "," + (!(boolBox : Bool)) + "," + folded + "," + found
			+ "," + Lambda.count(counts, c -> (c : Int) < 4));

		// Enum arguments and structures.
		var slot = Filled(count, High);
		var slotText = switch slot {
			case Filled(c, l): Std.string((c : Int) + l);
			case Empty: "empty";
		}
		var entry:Entry = {count: 2, meters: 0.5, flag: true};
		var sparse:Entry = {count: 3, flag: false};
		entry.count = entry.count + 4;
		var inferred = {c: count, m: meters};
		var entries:Array<Entry> = [entry, sparse];
		Sys.println("struct=" + slotText + "," + entry.count + "," + entry.meters + "," + entry.flag + "," + sparse.meters + "," + (sparse.meters == null)
			+ "," + (inferred.c + 1) + "," + (inferred.m + inferred.m) + "," + ((entries[1].count : Int) + entries[0].count) + "," + (!entries[1].flag));

		// Interfaces and overrides.
		var scorer:Scorer = hero;
		var base:Base = new Derived();
		Sys.println("interface=" + scorer.score(3) + "," + scorer.level() + "," + (scorer.score(1) + 1) + "," + base.value() + ","
			+ ((base.value() : Int) + new Base().value()));

		// Dynamic and Any boundaries.
		var dyn:Dynamic = count;
		var back:Count = dyn;
		var dynSum = dyn + 1;
		var checked = cast(dyn, Count);
		var unsafe:Count = cast dyn;
		var dynMeters:Dynamic = meters;
		var backMeters:Meters = dynMeters;
		var dynFlag:Dynamic = flag;
		var backFlag:Flag = dynFlag;
		var anyLevel:Any = High;
		var backLevel:Level = anyLevel;
		var mixed:Array<Dynamic> = [count, meters, flag, High];
		var playerDyn:Dynamic = player;
		playerDyn.count = 11;
		Sys.println("dynamic=" + (back + 1) + "," + dynSum + "," + (checked + 1) + "," + (unsafe * 2) + "," + (backMeters + backMeters) + "," + (!backFlag)
			+ "," + (backLevel + 1) + "," + mixed.join("|") + "," + describe(count) + "," + describe(meters) + "," + describe(flag) + "," + player.count
			+ "," + playerDyn.count + "," + (playerDyn.meters * 2));

		// Reflection.
		Reflect.setField(player, "count", 12);
		var fieldCount:Count = Reflect.field(player, "count");
		Sys.println("reflect=" + Reflect.field(player, "count") + "," + (fieldCount + 1) + "," + Reflect.field(player, "meters") + ","
			+ Reflect.field(player, "flag") + "," + Reflect.field(player, "lvl") + "," + Reflect.field(entry, "count") + "," + Std.string(count) + ","
			+ Std.string(meters) + "," + Std.string(flag) + "," + Std.string(High) + "," + Type.typeof(count) + "," + Type.typeof(meters) + ","
			+ Type.typeof(flag) + "," + Type.typeof(High) + "," + Std.isOfType(count, Int) + "," + Std.isOfType(meters, Float) + ","
			+ Std.isOfType(flag, Bool) + "," + Std.isOfType(count, Float) + "," + Reflect.callMethod(null, countOf, [8]));

		// JSON and serialization.
		var parsedJson:Dynamic = haxe.Json.parse('{"count":4,"meters":2.5,"flag":true}');
		var parsedCount:Count = parsedJson.count;
		var parsedMeters:Meters = parsedJson.meters;
		var parsedFlag:Flag = parsedJson.flag;
		var parsedCounts:Array<Count> = haxe.Json.parse("[1,2,3]");
		var serialized = haxe.Serializer.run([count, count]) + "|" + haxe.Serializer.run(meters) + "|" + haxe.Serializer.run(flag) + "|"
			+ haxe.Serializer.run(High);
		var restored:Array<Count> = haxe.Unserializer.run(haxe.Serializer.run(counts));
		var restoredMeters:Meters = haxe.Unserializer.run(haxe.Serializer.run(meters));
		Sys.println("json=" + haxe.Json.stringify(count) + "," + haxe.Json.stringify(meters) + "," + haxe.Json.stringify(flag) + ","
			+ haxe.Json.stringify(counts) + "," + haxe.Json.stringify({
				count: count
			}) + "," + (parsedCount + 1) + "," + (parsedMeters * 2) + "," + (!parsedFlag) + "," + ((parsedCounts[0] : Int) + parsedCounts[2]) + ","
			+ serialized + "," + (restored[3] + 1) + "," + (restoredMeters * 2));

		// Implicit casts: `from`/`to` on the header, `@:from`/`@:to` functions,
		// abstracts over abstracts, and text.
		var percent:Percent = 0.25;
		var percentFloat:Float = percent;
		var percentText:String = percent;
		var totalCount:Total = 6;
		var totalAsCount:Count = totalCount;
		var totalInt:Int = totalCount;
		var asFloat:Float = count;
		Sys.println("casts=" + percentFloat + "," + percentText + "," + (totalAsCount + 1) + "," + (totalInt + 2) + "," + asFloat
			+ ',c=$count m=$meters f=$flag l=$low' + "," + "x" + count + meters + flag);

		// Casts of nullable and Dynamic values to scalars and to their abstracts.
		var quarter:Null<Float> = 2.25;
		var maybeOn:Null<Bool> = true;
		var maybeFour:Null<Int> = 4;
		var sevenDyn:Dynamic = 7;
		var castFloat:Null<Float> = cast sevenDyn;
		var castMeters:Null<Meters> = cast sevenDyn;
		var widenedMeters:Null<Float> = castMeters;
		var nothingDyn:Dynamic = null;
		var nullCasts = try {
			cast(nothingDyn, Int);
			cast(nothingDyn, Count);
			cast(nothingDyn, Float);
			"passed";
		} catch (e:Dynamic) {
			"threw " + e;
		}
		Sys.println("nullableCasts=" + (cast(quarter, Float) * 2) + "," + ((cast(quarter, Meters) : Float) * 2) + "," + !cast(maybeOn, Bool) + ","
			+ !(cast(maybeOn, Flag) : Bool) + "," + (cast(maybeFour, Int) + 1) + "," + ((cast(maybeFour, Count) : Int) + 1) + "," + (castFloat + 0.5) + ","
			+ (widenedMeters + 0.5) + "," + nullCasts);

		// Iteration ranges and Null<T> wrapping.
		var rangeOut = new Array<Int>();
		for (i in 0...count) {
			rangeOut.push(i);
		}
		for (i in Low...High) {
			rangeOut.push(i * 10);
		}
		var maybe:Null<Count> = null;
		var wasNull = maybe == null;
		maybe = 4;
		var unwrapped:Count = maybe;
		var maybeMeters:Null<Meters> = null;
		var maybeFlag:Null<Flag> = null;
		var maybeLevel:Null<Level> = null;
		maybeLevel = Mid;
		var parsedInt:Count = Std.parseInt("12");
		var maybeId:MaybeId = null;
		var someId:MaybeId = 8;
		Sys.println("null=" + rangeOut.join("|") + "," + wasNull + "," + (unwrapped + 1) + "," + (maybe != null ? maybe + 1 : 0) + "," + maybeMeters + ","
			+ (maybeMeters == null) + "," + maybeFlag + "," + maybeLevel + "," + (maybeLevel == Mid) + "," + (parsedInt + 1) + "," + maybeId + "," + someId
			+ "," + (maybeId == null));
	}
}
