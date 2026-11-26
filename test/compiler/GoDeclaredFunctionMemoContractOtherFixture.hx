/** Shares its name with the private class in `GoDeclaredFunctionMemoContractFixture`. */
@:build(GoDeclaredFunctionMemoContractMacro.addMade(11))
private class MemoHidden {
	public static function run(step:Int = 5):Int {
		return step + 5;
	}
}

/** The second module of `GoDeclaredFunctionMemoContract`. */
class GoDeclaredFunctionMemoContractOtherFixture {
	public function new() {}
}
