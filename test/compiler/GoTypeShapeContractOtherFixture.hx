/** Shares its name with the private enum in `GoTypeShapeContractFixture`. */
private enum Hidden {
	HiddenB(value:String);
}

/** Shares its name with the private abstract in `GoTypeShapeContractFixture`. */
private abstract Score(String) {
	public inline function new(value:String) {
		this = value;
	}
}

/** Shares its name with the private typedef in `GoTypeShapeContractFixture`. */
private typedef Pair = {other:Bool};

/** The second module of `GoTypeShapeContract`. */
class GoTypeShapeContractOtherFixture {
	public var hidden:Hidden;
	public var score:Score;
	public var pair:Pair;

	public function new() {
		hidden = HiddenB("b");
		score = new Score("b");
		pair = {other: true};
	}

	/** Its local `pick<V>` prints as `pick.V`, like the one in the first fixture. */
	public static function local():String {
		function pick<V>(value:V):V {
			return value;
		}
		return pick("x");
	}
}
