package app.runtime;

import app.core.PulseEvent;

interface PulseRuntime {
	public function profileId():String;
	public function variantId():String;
	public function capabilityId():String;
	public function processScore(events:Array<PulseEvent>):Int;
}
