package app.runtime;

import app.core.PulseEvent;

class CoreRuntime implements PulseRuntime {
	public function new() {}

	public function profileId():String {
		return BuildConfig.PROFILE;
	}

	public function variantId():String {
		return BuildConfig.VARIANT;
	}

	public function capabilityId():String {
		return "core_loop";
	}

	public function processScore(events:Array<PulseEvent>):Int {
		var score = 0;
		for (event in events) {
			score += event.value;
			if (event.isAlert()) {
				score += 3;
			}
		}
		return score;
	}
}
