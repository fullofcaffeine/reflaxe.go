/** A private enum that shares its name with one in `GoTypeShapeContractOtherFixture`. */
private enum Hidden {
	HiddenA;
}

/** A private abstract that shares its name with one in `GoTypeShapeContractOtherFixture`. */
private abstract Score(Int) from Int to Int {}

/** A private typedef that shares its name with one in `GoTypeShapeContractOtherFixture`. */
private typedef Pair = {left:Int, ?right:String};

/** Field and expression types that `GoTypeShapeContract` converts into shapes. */
class GoTypeShapeContractFixture<T> {
	public var hidden:Hidden;
	public var score:Score;
	public var pair:Pair;
	public var nested:Array<Array<Null<Int>>>;
	public var iterators:Iterator<Iterator<Int>>;
	public var callback:(Int, ?String) -> Array<T>;
	public var dynamicScores:Dynamic<Score>;
	public var plainDynamic:Dynamic;
	public var map:Map<String, Hidden>;

	public function new() {
		hidden = HiddenA;
		score = 1;
		pair = {left: 1};
	}

	public function wrap<S>(value:S):Array<S> {
		return [value];
	}

	/** Its local `pick<V>` prints as `pick.V`, like the one in the other fixture. */
	public static function local():Int {
		function pick<V>(value:V):V {
			return value;
		}
		var scored:Score = pick(2);
		return pick(1) + scored;
	}
}
