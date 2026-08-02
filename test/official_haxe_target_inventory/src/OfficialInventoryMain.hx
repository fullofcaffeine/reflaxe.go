import haxe.ds.List;
import utest.Assert;
import utest.Assertation;
import utest.Async;

/**
	What: Runs one generated shard of the pinned official Haxe test inventory.
	Why: The complete upstream inventory is too large for one useful diagnostic
	compile, while per-file compiles would repeat package and compiler overhead.
	How: Python generates only the case constructors and unitstd symlink tree for
	the selected shard. Each upstream case's compile-time utest initializer owns
	method discovery and ordering; this adapter only executes its synchronous plan.
**/
class OfficialInventoryMain {
	static inline final RECORD_PREFIX = "OFFICIAL_HAXE_INVENTORY_RECORD\t";
	static inline final SUMMARY_PREFIX = "OFFICIAL_HAXE_INVENTORY_SUMMARY\t";

	static var completed = 0;
	static var expected = 0;

	static function main():Void {
		var cases = OfficialInventoryCases.build();
		for (testCase in cases) {
			runCase(testCase);
		}
		OfficialSmokeObserver.line(SUMMARY_PREFIX + completed + "\t" + expected + "\t0");
	}

	static function runCase(testCase:OfficialInventoryCase):Void {
		var plan = testCase.__initializeUtest__();
		expected += plan.tests.length;
		var owner = Type.getClassName(Type.getClass(testCase));
		if (owner == null)
			throw "official inventory case has no runtime class";

		runAccessory(plan.accessories.setupClass);
		for (test in plan.tests) {
			Assert.results = new List<Assertation>();
			runAccessory(plan.accessories.setup);
			requireResolved(test.execute(), owner + "." + test.name);
			runAccessory(plan.accessories.teardown);

			var assertions = 0;
			for (_ in Assert.results)
				assertions++;
			if (assertions == 0) {
				throw "official inventory method produced no active assertions: " + owner + "." + test.name;
			}
			completed++;
			OfficialSmokeObserver.line(RECORD_PREFIX + owner + "." + test.name + "\tpass\t" + assertions);
		}
		runAccessory(plan.accessories.teardownClass);
	}

	static function runAccessory(action:Null<Void->Async>):Void {
		if (action != null)
			requireResolved(action(), "utest accessory");
	}

	static function requireResolved(async:Async, owner:String):Void {
		if (!async.resolved) {
			throw "official inventory unexpectedly reached an asynchronous operation: " + owner;
		}
	}
}
