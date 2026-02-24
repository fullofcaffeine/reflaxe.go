package app.core;

import app.runtime.PulseRuntime;

class PulsePipeline {
	final runtime:PulseRuntime;

	public function new(runtime:PulseRuntime) {
		this.runtime = runtime;
	}

	public function run(events:Array<PulseEvent>):PulseReport {
		var alertCount = 0;
		var totalValue = 0;
		for (event in events) {
			totalValue += event.value;
			if (event.isAlert()) {
				alertCount++;
			}
		}

		var runtimeScore = runtime.processScore(events);
		return new PulseReport(runtime.profileId(), runtime.variantId(), runtime.capabilityId(), events.length, totalValue, alertCount, runtimeScore);
	}
}
