package app.runtime;

import app.core.PulseAlert;
import app.core.PulseCodec;
import app.core.PulseEnrichedEvent;
import app.core.PulseEvent;
import app.core.PulseIngressFrame;

/**
	Portable-baseline runtime adapter for PulseForge.

	This lane uses straightforward sequential loops. It is the reference behavior
	for portability and deterministic contract output.

	See `GoNativeRuntime` for the Go-first lane that changes execution strategy
	(using channels/select) while preserving the same domain contract.
**/
class CoreRuntime implements PulseRuntime {
	public function profileId():String {
		return BuildConfig.PROFILE;
	}

	public function variantId():String {
		return BuildConfig.VARIANT;
	}

	public function capabilityId():String {
		return "core_loop";
	}

	public function parse(frames:Array<PulseIngressFrame>, workerCount:Int):Array<PulseEvent> {
		var parsed = new Array<PulseEvent>();
		for (frame in frames) {
			parsed.push(PulseCodec.parse(frame));
		}
		return parsed;
	}

	public function enrich(events:Array<PulseEvent>, workerCount:Int):Array<PulseEnrichedEvent> {
		var enriched = new Array<PulseEnrichedEvent>();
		for (event in events) {
			enriched.push(PulseCodec.enrich(event));
		}
		return enriched;
	}

	public function stageScore(parsed:Array<PulseEvent>, enriched:Array<PulseEnrichedEvent>, alerts:Array<PulseAlert>, backpressureEvents:Int):Int {
		var score = 0;
		for (entry in enriched) {
			score += entry.weightedValue;
		}
		score += alerts.length * 5;
		score -= backpressureEvents * 2;
		score += parsed.length;
		return score;
	}
}
