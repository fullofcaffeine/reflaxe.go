import haxe.ds.List;
import utest.Assert;
import utest.Assertation;

/**
	What: Runs one active method from each selected official Haxe test family.
	Why: A target smoke must execute the official assertions on generated Go, while
	the full asynchronous utest runner would also claim unrelated timer, stack-trace,
	and report behavior that this deliberately narrow tracer does not select.
	How: Invoke the exact official methods through a tiny target adapter, use utest's
	own fail-fast assertion implementation, and publish active method/assertion counts.
**/
@:access(OfficialIssue6369Case)
class OfficialTargetSmokeMain {
	static inline final RECORD_PREFIX = "OFFICIAL_HAXE_SMOKE_RECORD\t";
	static inline final SUMMARY_PREFIX = "OFFICIAL_HAXE_SMOKE_SUMMARY\t";
	static inline final CONTROL_PREFIX = "OFFICIAL_HAXE_SMOKE_CONTROL\t";

	static function main():Void {
		#if official_haxe_smoke_runtime_failure
		OfficialSmokeObserver.line(CONTROL_PREFIX + "runtime");
		throw "official-haxe-smoke deliberate runtime failure";
		#end
		#if official_haxe_smoke_timeout_failure
		OfficialSmokeObserver.line(CONTROL_PREFIX + "timeout");
		// A real blocking call cannot be optimized away as an empty loop can.
		Sys.sleep(60.0);
		return;
		#end

		var completed = 0;
		runCase("unit.TestNumericSuffixes.testFloatSuffixes", function() new OfficialNumericSuffixCase().testFloatSuffixes());
		completed++;
		runCase("unit.spec.TestIntIterator.test", function() new OfficialIntIteratorCase().test());
		completed++;
		runCase("unit.issues.Issue6369.test", function() new OfficialIssue6369Case().test());
		completed++;

		#if official_haxe_smoke_assertion_failure
		OfficialSmokeObserver.line(CONTROL_PREFIX + "assertion");
		runCase("OfficialSmokeDeliberateFailure.testDeliberateFailure", function() {
			Assert.fail("official-haxe-smoke deliberate assertion failure");
		});
		#end

		OfficialSmokeObserver.line(SUMMARY_PREFIX + completed + "\t3\tpass");
	}

	static function runCase(id:String, action:Void->Void):Void {
		Assert.results = new List<Assertation>();
		action();
		var assertions = 0;
		for (_ in Assert.results) {
			assertions++;
		}
		if (assertions == 0) {
			throw "official Haxe target smoke method produced no active assertions: " + id;
		}
		OfficialSmokeObserver.line(RECORD_PREFIX + id + "\tpass\t" + assertions);
	}
}
