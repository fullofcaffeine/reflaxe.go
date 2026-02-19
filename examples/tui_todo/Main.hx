import profile.RuntimeFactory;

class Main {
	static function hasArg(flag:String):Bool {
		for (arg in Sys.args()) {
			if (arg == flag) {
				return true;
			}
		}
		return false;
	}

	static function main() {
		var runtime = RuntimeFactory.create();
		#if example_ci
		Sys.println(Harness.assertContract(runtime));
		#else
		if (hasArg("--scripted")) {
			Sys.println(Harness.run(runtime));
		} else {
			InteractiveCli.run(runtime);
		}
		#end
	}
}
