/** Owner of the methods that `GoDeclaredFunctionMemoContract` asks about. */
class GoDeclaredFunctionMemoContractFixture {
	public var count:Int;

	public function new(start:Int = 1) {
		count = start;
	}

	public static function run(step:Int = 1):Int {
		return step + 1;
	}

	public function apply(label:String = "base"):String {
		return label;
	}

	/** A dynamic method has no stable declared function. */
	public dynamic function hook(value:Int = 3):Int {
		return value;
	}
}

/** Declares a static `run` and an instance `apply` of its own, with other defaults. */
class MemoOther {
	public function new() {}

	public static function run(step:Int = 2):Int {
		return step * 2;
	}

	public function apply(label:String = "other"):String {
		return label + "!";
	}
}

/** Overrides `apply` with another default. */
class MemoChild extends GoDeclaredFunctionMemoContractFixture {
	override public function apply(label:String = "child"):String {
		return label + "?";
	}
}

/** Two overloads share their class and name. */
class MemoOverloads {
	public overload extern inline static function pick(value:Int, step:Int = 6):Int {
		return value + step;
	}

	public overload extern inline static function pick(value:String, suffix:String = "seven"):String {
		return value + suffix;
	}
}

/** Gets a generated `made` at the same position as `MemoMadeB`. */
@:build(GoDeclaredFunctionMemoContractMacro.addMade(8))
class MemoMadeA {}

/** Gets a generated `made` at the same position as `MemoMadeA`. */
@:build(GoDeclaredFunctionMemoContractMacro.addMade(9))
class MemoMadeB {}

/** Shares its name with the private class in `GoDeclaredFunctionMemoContractOtherFixture`. */
@:build(GoDeclaredFunctionMemoContractMacro.addMade(10))
private class MemoHidden {
	public static function run(step:Int = 4):Int {
		return step - 1;
	}
}
