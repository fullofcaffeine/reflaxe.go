import utest.Assert;

/** Target adapter around the exact pinned UnitBuilder-expanded statement body. */
class OfficialIntIteratorCase {
	public function new() {}

	public function test():Void {
		OfficialUnitStdAdapter.body();
	}

	public function eq<T>(expected:T, actual:T, ?pos:haxe.PosInfos):Void {
		Assert.equals(expected, actual, pos);
	}

	public function t(value:Bool, ?pos:haxe.PosInfos):Void {
		Assert.isTrue(value, pos);
	}

	public function f(value:Bool, ?pos:haxe.PosInfos):Void {
		Assert.isFalse(value, pos);
	}
}
