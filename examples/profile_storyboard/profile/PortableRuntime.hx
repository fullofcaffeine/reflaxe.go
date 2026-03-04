package profile;

/**
	Portable runtime implementation for `profile_storyboard`.

	Why this file is simple:
	- the example is a portable reference app,
	- no separate metal runtime is shipped because prior metal-only differences
	  were synthetic and not tied to real performance/interop value.
**/
class PortableRuntime implements StoryboardRuntime {
	public function new() {}

	public function profileId():String {
		return "portable";
	}

	public function decorateTitle(title:String):String {
		return title;
	}

	public function highlightTag(tag:String):String {
		return tag;
	}

	public function extraSignal(metrics:StorySignalMetrics):String {
		return "interop_lane=off,optimizer=stable,policy_gate=off";
	}

	public function supportsVelocityHint():Bool {
		return false;
	}

	public function velocityPerSprint():Int {
		return 5;
	}

	public function riskThreshold():Int {
		return 5;
	}
}
