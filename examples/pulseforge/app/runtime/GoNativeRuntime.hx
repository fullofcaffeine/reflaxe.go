package app.runtime;

import app.core.PulseAlert;
import app.core.PulseCodec;
import app.core.PulseEnrichedEvent;
import app.core.PulseEvent;
import app.core.PulseIngressFrame;
import go.Chan;
import go.Go;
import go.Select;
import haxe.ds.IntMap;

class GoNativeRuntime implements PulseRuntime {
	public function profileId():String {
		return BuildConfig.PROFILE;
	}

	public function variantId():String {
		return BuildConfig.VARIANT;
	}

	public function capabilityId():String {
		return "chan_fanout_select";
	}

	public function parse(frames:Array<PulseIngressFrame>, workerCount:Int):Array<PulseEvent> {
		if (frames.length == 0) {
			return [];
		}

		var workers = normalizedWorkers(workerCount);
		var inbox:Chan<PulseIngressFrame> = Go.newChan(frames.length);
		var out:Chan<PulseEvent> = Go.newChan(frames.length);
		var done:Chan<Int> = Go.newChan(workers);

		for (frame in frames) {
			inbox.send(frame);
		}
		inbox.close();

		var index = 0;
		while (index < workers) {
			Go.spawn(function() {
				var processed = 0;
				while (true) {
					var frame = inbox.recv();
					if (frame == null) {
						break;
					}
					out.send(PulseCodec.parse(cast frame));
					processed++;
				}
				done.send(processed);
			});
			index++;
		}

		waitForWorkers(done, workers);

		return orderParsed(drainEvents(out, frames.length), frames.length);
	}

	public function enrich(events:Array<PulseEvent>, workerCount:Int):Array<PulseEnrichedEvent> {
		if (events.length == 0) {
			return [];
		}

		var workers = normalizedWorkers(workerCount);
		var inbox:Chan<PulseEvent> = Go.newChan(events.length);
		var out:Chan<PulseEnrichedEvent> = Go.newChan(events.length);
		var done:Chan<Int> = Go.newChan(workers);

		for (event in events) {
			inbox.send(event);
		}
		inbox.close();

		var index = 0;
		while (index < workers) {
			Go.spawn(function() {
				var processed = 0;
				while (true) {
					var event = inbox.recv();
					if (event == null) {
						break;
					}
					out.send(PulseCodec.enrich(cast event));
					processed++;
				}
				done.send(processed);
			});
			index++;
		}

		waitForWorkers(done, workers);

		return orderEnriched(drainEnriched(out, events.length), events.length);
	}

	public function stageScore(parsed:Array<PulseEvent>, enriched:Array<PulseEnrichedEvent>, alerts:Array<PulseAlert>, backpressureEvents:Int):Int {
		var inbox:Chan<Int> = Go.newChan(enriched.length);
		var score = 0;

		for (entry in enriched) {
			score += switch (Select.send(inbox, entry.weightedValue)) {
				case Sent: 2;
				case Defaulted: -25;
			};
		}

		var remaining = enriched.length;
		while (remaining > 0) {
			score += switch (Select.recv(inbox)) {
				case Received(value): value;
				case Defaulted:
					inbox.recv();
			};
			remaining--;
		}

		inbox.close();
		score += alerts.length * 7;
		score -= backpressureEvents * 3;
		score += parsed.length;
		return score;
	}

	function normalizedWorkers(workerCount:Int):Int {
		if (workerCount <= 0) {
			return 1;
		}
		return workerCount;
	}

	function waitForWorkers(done:Chan<Int>, workers:Int):Void {
		var completed = 0;
		while (completed < workers) {
			done.recv();
			completed++;
		}
		done.close();
	}

	function drainEvents(out:Chan<PulseEvent>, expected:Int):Array<PulseEvent> {
		var parsed = new Array<PulseEvent>();
		var remaining = expected;
		while (remaining > 0) {
			switch (Select.recv(out)) {
				case Received(event):
					parsed.push(event);
					remaining--;
				case Defaulted:
					var fallback = out.recv();
					if (fallback != null) {
						parsed.push(cast fallback);
						remaining--;
					}
			}
		}
		out.close();
		return parsed;
	}

	function drainEnriched(out:Chan<PulseEnrichedEvent>, expected:Int):Array<PulseEnrichedEvent> {
		var enriched = new Array<PulseEnrichedEvent>();
		var remaining = expected;
		while (remaining > 0) {
			switch (Select.recv(out)) {
				case Received(event):
					enriched.push(event);
					remaining--;
				case Defaulted:
					var fallback = out.recv();
					if (fallback != null) {
						enriched.push(cast fallback);
						remaining--;
					}
			}
		}
		out.close();
		return enriched;
	}

	function orderParsed(items:Array<PulseEvent>, expected:Int):Array<PulseEvent> {
		var byId:IntMap<PulseEvent> = new IntMap<PulseEvent>();
		for (item in items) {
			byId.set(item.id, item);
		}

		var ordered = new Array<PulseEvent>();
		var id = 1;
		while (id <= expected) {
			var event = byId.get(id);
			if (event != null) {
				ordered.push(event);
			}
			id++;
		}
		return ordered;
	}

	function orderEnriched(items:Array<PulseEnrichedEvent>, expected:Int):Array<PulseEnrichedEvent> {
		var byId:IntMap<PulseEnrichedEvent> = new IntMap<PulseEnrichedEvent>();
		for (item in items) {
			byId.set(item.event.id, item);
		}

		var ordered = new Array<PulseEnrichedEvent>();
		var id = 1;
		while (id <= expected) {
			var event = byId.get(id);
			if (event != null) {
				ordered.push(event);
			}
			id++;
		}
		return ordered;
	}
}
