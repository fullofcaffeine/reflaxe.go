package app.runtime;

import app.core.PulseAlert;
import app.core.PulseCodec;
import app.core.PulseEnrichedEvent;
import app.core.PulseEvent;
import app.core.PulseIngressFrame;

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
