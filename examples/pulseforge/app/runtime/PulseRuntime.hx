package app.runtime;

import app.core.PulseAlert;
import app.core.PulseEnrichedEvent;
import app.core.PulseEvent;
import app.core.PulseIngressFrame;

/**
	Execution-strategy contract for PulseForge runtime lanes.

	`CoreRuntime` is the portable baseline implementation (simple loop stages).
	`GoNativeRuntime` is the Go-first implementation (channel/select worker fanout).

	Both implementations must keep the same domain-level output contract; only
	the execution strategy differs.
**/
interface PulseRuntime {
	public function profileId():String;
	public function variantId():String;
	public function capabilityId():String;
	public function parse(frames:Array<PulseIngressFrame>, workerCount:Int):Array<PulseEvent>;
	public function enrich(events:Array<PulseEvent>, workerCount:Int):Array<PulseEnrichedEvent>;
	public function stageScore(parsed:Array<PulseEvent>, enriched:Array<PulseEnrichedEvent>, alerts:Array<PulseAlert>, backpressureEvents:Int):Int;
}
