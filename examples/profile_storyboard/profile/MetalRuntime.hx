package profile;

class MetalRuntime implements StoryboardRuntime {
	public function new() {}

	public function profileId():String {
		return "metal";
	}

	public function decorateTitle(title:String):String {
		return "[strict] " + title;
	}

	public function highlightTag(tag:String):String {
		return "metal-" + tag;
	}

	public function extraSignal(metrics:StorySignalMetrics):String {
		return "interop_lane=typed+strict,high_value="
			+ metrics.highValue
			+ ",open_high_value="
			+ metrics.openHighValue
			+ ",policy_gate=on";
	}

	public function supportsVelocityHint():Bool {
		return true;
	}

	public function velocityPerSprint():Int {
		return 9;
	}

	public function riskThreshold():Int {
		return 4;
	}
}
