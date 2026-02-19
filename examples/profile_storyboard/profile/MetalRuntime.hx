package profile;

import domain.StoryCard;
import haxe.ds.List;

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

	public function extraSignal(cards:List<StoryCard>):String {
		var highValue = 0;
		var openHighValue = 0;
		var count = cards.length;
		var i = 0;
		while (i < count) {
			var value = cards.pop();
			if (value == null) {
				break;
			}
			var card:StoryCard = cast value;
			if (card.points >= 5) {
				highValue++;
				if (card.state != "done") {
					openHighValue++;
				}
			}
			cards.add(card);
			i++;
		}
		return "interop_lane=typed+strict,high_value=" + highValue + ",open_high_value=" + openHighValue + ",policy_gate=on";
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
