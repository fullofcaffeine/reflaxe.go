package profile;

import domain.StoryCard;
import haxe.ds.List;

class GopherRuntime implements StoryboardRuntime {
	public function new() {}

	public function profileId():String {
		return "gopher";
	}

	public function decorateTitle(title:String):String {
		return "[go] " + title;
	}

	public function highlightTag(tag:String):String {
		return "go-" + tag;
	}

	public function extraSignal(cards:List<StoryCard>):String {
		var total = 0;
		var releaseTagged = 0;
		var count = cards.length;
		var i = 0;
		while (i < count) {
			var value = cards.pop();
			if (value == null) {
				break;
			}
			var card:StoryCard = cast value;
			total += card.points;

			var tagCount = card.tags.length;
			var j = 0;
			while (j < tagCount) {
				var tagValue = card.tags.pop();
				if (tagValue == null) {
					break;
				}
				var tag:String = cast tagValue;
				if (tag == "release") {
					releaseTagged++;
				}
				card.tags.add(tag);
				j++;
			}

			cards.add(card);
			i++;
		}
		return "interop_lane=typed,total_points=" + total + ",parallel_streams=3,release_cards=" + releaseTagged;
	}

	public function supportsVelocityHint():Bool {
		return true;
	}

	public function velocityPerSprint():Int {
		return 8;
	}

	public function riskThreshold():Int {
		return 5;
	}
}
