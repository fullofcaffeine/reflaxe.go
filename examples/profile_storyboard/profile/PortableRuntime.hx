package profile;

import domain.StoryCard;
import haxe.ds.List;

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

	public function extraSignal(cards:List<StoryCard>):String {
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
