/**
	First fixture module for `GoTypeUsageTrackerContract`.

	Why: The contract compares the compiler's type-usage tracker with Reflaxe's
	on module types that reach every tracker rule.

	What: A generic class with a generic superclass and interface, statics, a
	constructor, instance fields, and method bodies that construct, declare,
	access statics, build functions and anonymous structures, and follow
	abstracts over abstracts. An instance method has a local generic function
	`pick<V>`; the static `local()` inlines the other module's `pick<V>`, which
	prints alike, and its private declarations. Reflaxe keeps the first of two
	entries with one id and visits instance fields before statics, so the
	order of these two methods is part of the contract. Also an enum, a typedef,
	abstracts, and private declarations whose names the other fixture module
	reuses.

	How: The contract loads this module and `GoTypeUsageTrackerContractOtherFixture`
	with `Context.getModule(...)`.
**/
interface GoTypeUsageTrackerContractShape<T> {
	function shapeOf(value:T):Null<T>;
}

class GoTypeUsageTrackerContractBase<T> {
	public var stored:Array<T>;

	public function new(stored:Array<T>) {
		this.stored = stored;
	}
}

enum GoTypeUsageTrackerContractEvent<T> {
	Started;
	Moved(distance:GoTypeUsageTrackerContractKilometers, label:String);
	Carried(value:T, next:Null<GoTypeUsageTrackerContractEvent<T>>);
}

abstract GoTypeUsageTrackerContractMeters(Float) from Float to Float {}
abstract GoTypeUsageTrackerContractKilometers(GoTypeUsageTrackerContractMeters) from GoTypeUsageTrackerContractMeters {}

enum abstract GoTypeUsageTrackerContractMode(Int) to Int {
	var Fast = 1;
	var Slow = 2;
}

typedef GoTypeUsageTrackerContractPair = {
	final left:Int;
	final ?right:String;
	final callback:(Int, String) -> Bool;
}

private typedef Hidden = {
	final value:Int;
}

private abstract Score(Int) from Int {}

private enum Choice {
	Left(value:Int);
	Right(text:String);
}

class GoTypeUsageTrackerContractFixture<T> extends GoTypeUsageTrackerContractBase<Array<T>> implements GoTypeUsageTrackerContractShape<T> {
	public static var counter:Int = 0;
	public static var registry:Map<String, Array<Int>> = new Map();

	public var hidden:Hidden;
	public var score:Score;
	public var choice:Choice;
	public var mode:GoTypeUsageTrackerContractMode;
	public var listener:Null<(GoTypeUsageTrackerContractPair) -> Void>;

	public function new() {
		super([]);
		hidden = {value: 1};
		score = 2;
		choice = Left(1);
		mode = Fast;
		listener = null;
	}

	public function shapeOf(value:T):Null<T> {
		return value;
	}

	public function exercise(values:Array<T>, distance:GoTypeUsageTrackerContractKilometers):GoTypeUsageTrackerContractEvent<T> {
		counter++;
		var total:GoTypeUsageTrackerContractMeters = 0.0;
		final names = new List<String>();
		final structure:GoTypeUsageTrackerContractPair = {left: 1, callback: (number, text) -> number > text.length};
		for (value in values)
			names.add(Std.string(value));
		final box = new GoTypeUsageTrackerContractBase<Map<Int, String>>([]);
		final anything:Dynamic = box;
		final iterator:Iterator<T> = values.iterator();
		if (registry.exists("a") && iterator.hasNext() && anything != null)
			total = total + 1.0;
		return values.length > 0 ? Carried(values[0], Moved(distance, structure.right == null ? "" : structure.right)) : Started;
	}

	public function pickLocal():Int {
		function pick<V>(value:V):V
			return value;
		return pick(counter);
	}

	public static function local():Int {
		return GoTypeUsageTrackerContractOtherFixture.picked() + GoTypeUsageTrackerContractOtherFixture.hiddenLength();
	}
}
