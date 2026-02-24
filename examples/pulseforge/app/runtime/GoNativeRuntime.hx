package app.runtime;

import app.core.PulseEvent;
import go.Chan;
import go.Go;
import go.Select;

class GoNativeRuntime implements PulseRuntime {
	public function new() {}

	public function profileId():String {
		return BuildConfig.PROFILE;
	}

	public function variantId():String {
		return BuildConfig.VARIANT;
	}

	public function capabilityId():String {
		return "chan_select";
	}

	public function processScore(events:Array<PulseEvent>):Int {
		var inbox:Chan<Int> = Go.newChan(events.length);
		var score = 0;

		for (event in events) {
			score += switch (Select.send(inbox, event.value)) {
				case Sent: 2;
				case Defaulted: -20;
			};
		}

		var remaining = events.length;
		while (remaining > 0) {
			score += switch (Select.recv(inbox)) {
				case Received(value): value;
				case Defaulted: -10;
			};
			remaining--;
		}

		inbox.close();
		return score;
	}
}
