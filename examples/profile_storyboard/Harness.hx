import domain.StoryCard;
import haxe.ds.List;
import profile.StoryboardRuntime;

class Harness {
	static inline final STATE_TODO = "todo";
	static inline final STATE_DOING = "doing";
	static inline final STATE_DONE = "done";

	static function makeTags(a:String, ?b:String):List<String> {
		var tags = new List<String>();
		tags.add(a);
		if (b != null) {
			tags.add(b);
		}
		return tags;
	}

	static function joinStringList(values:List<String>, separator:String):String {
		var out = "";
		var first = true;
		var count = values.length;
		var i = 0;
		while (i < count) {
			var raw = values.pop();
			if (raw == null) {
				break;
			}
			var value:String = cast raw;
			if (!first) {
				out += separator;
			}
			out += value;
			values.add(value);
			first = false;
			i++;
		}
		return out;
	}

	static function card(id:Int, title:String, points:Int, tags:List<String>, state:String, owner:String):StoryCard {
		return new StoryCard(id, title, points, tags, state, owner);
	}

	public static function buildCards():List<StoryCard> {
		var cards = new List<StoryCard>();
		cards.add(card(1, "Ship profile docs", 3, makeTags("docs", "profiles"), STATE_DONE, "Alex"));
		cards.add(card(2, "Backfill regression snapshots", 5, makeTags("tests"), STATE_DONE, "Mira"));
		cards.add(card(3, "Wire release artifacts", 5, makeTags("ci", "release"), STATE_DOING, "Noah"));
		cards.add(card(4, "CLI polish for dev:hx", 3, makeTags("devex"), STATE_TODO, "Jules"));
		cards.add(card(5, "Interactive tui_todo demo", 5, makeTags("examples", "release"), STATE_DOING, "Sam"));
		return cards;
	}

	static function totalPoints(cards:List<StoryCard>):Int {
		var totalPoints = 0;
		var count = cards.length;
		var i = 0;
		while (i < count) {
			var value = cards.pop();
			if (value == null) {
				break;
			}
			var card:StoryCard = cast value;
			totalPoints += card.points;
			cards.add(card);
			i++;
		}
		return totalPoints;
	}

	static function countByState(cards:List<StoryCard>, state:String):Int {
		var total = 0;
		var count = cards.length;
		var i = 0;
		while (i < count) {
			var value = cards.pop();
			if (value == null) {
				break;
			}
			var card:StoryCard = cast value;
			if (card.state == state) {
				total++;
			}
			cards.add(card);
			i++;
		}
		return total;
	}

	static function donePoints(cards:List<StoryCard>):Int {
		var total = 0;
		var count = cards.length;
		var i = 0;
		while (i < count) {
			var value = cards.pop();
			if (value == null) {
				break;
			}
			var card:StoryCard = cast value;
			if (card.state == STATE_DONE) {
				total += card.points;
			}
			cards.add(card);
			i++;
		}
		return total;
	}

	static function openPoints(cards:List<StoryCard>):Int {
		var total = 0;
		var count = cards.length;
		var i = 0;
		while (i < count) {
			var value = cards.pop();
			if (value == null) {
				break;
			}
			var card:StoryCard = cast value;
			if (card.state != STATE_DONE) {
				total += card.points;
			}
			cards.add(card);
			i++;
		}
		return total;
	}

	static function hasTag(card:StoryCard, needle:String):Bool {
		var found = false;
		var count = card.tags.length;
		var i = 0;
		while (i < count) {
			var value = card.tags.pop();
			if (value == null) {
				break;
			}
			var tag:String = cast value;
			if (tag == needle) {
				found = true;
			}
			card.tags.add(tag);
			i++;
		}
		return found;
	}

	static function openHighRisk(cards:List<StoryCard>, threshold:Int):Int {
		var total = 0;
		var count = cards.length;
		var i = 0;
		while (i < count) {
			var value = cards.pop();
			if (value == null) {
				break;
			}
			var card:StoryCard = cast value;
			if (card.state != STATE_DONE && card.points >= threshold) {
				total++;
			}
			cards.add(card);
			i++;
		}
		return total;
	}

	static function releaseTaggedOpen(cards:List<StoryCard>):Int {
		var total = 0;
		var count = cards.length;
		var i = 0;
		while (i < count) {
			var value = cards.pop();
			if (value == null) {
				break;
			}
			var card:StoryCard = cast value;
			if (card.state != STATE_DONE && hasTag(card, "release")) {
				total++;
			}
			cards.add(card);
			i++;
		}
		return total;
	}

	static function readinessPercent(donePoints:Int, totalPoints:Int):Int {
		if (totalPoints <= 0) {
			return 0;
		}
		return intFloorDiv(donePoints * 100, totalPoints);
	}

	static function sprintForecast(openPoints:Int, velocityPerSprint:Int):Int {
		if (openPoints <= 0) {
			return 0;
		}
		return intFloorDiv(openPoints + velocityPerSprint - 1, velocityPerSprint);
	}

	static function intFloorDiv(numerator:Int, denominator:Int):Int {
		if (denominator <= 0) {
			return 0;
		}
		var quotient = 0;
		var remaining = numerator;
		while (remaining >= denominator) {
			remaining -= denominator;
			quotient++;
		}
		return quotient;
	}

	static function formatCard(card:StoryCard, runtime:StoryboardRuntime):String {
		var tags = new List<String>();
		var tagCount = card.tags.length;
		var j = 0;
		while (j < tagCount) {
			var tagValue = card.tags.pop();
			if (tagValue == null) {
				break;
			}
			var tag:String = cast tagValue;
			tags.add(runtime.highlightTag(tag));
			card.tags.add(tag);
			j++;
		}

		return "#"
			+ card.id
			+ " p"
			+ card.points
			+ " "
			+ runtime.decorateTitle(card.title)
			+ " owner:"
			+ card.owner
			+ " tags:"
			+ joinStringList(tags, "|");
	}

	static function repeatChar(ch:String, count:Int):String {
		if (count <= 0) {
			return "";
		}
		var out = "";
		var i = 0;
		while (i < count) {
			out += ch;
			i++;
		}
		return out;
	}

	static function progressBar(donePoints:Int, totalPoints:Int, width:Int):String {
		if (width <= 0) {
			return "[]";
		}
		var filled = 0;
		if (totalPoints > 0) {
			filled = intFloorDiv(donePoints * width, totalPoints);
		}
		if (filled < 0) {
			filled = 0;
		}
		if (filled > width) {
			filled = width;
		}
		return "[" + repeatChar("#", filled) + repeatChar("-", width - filled) + "]";
	}

	static function formatLane(cards:List<StoryCard>, state:String, title:String, runtime:StoryboardRuntime):String {
		var out = title + " (" + countByState(cards, state) + ")\n";
		var hasEntries = false;
		var cardCount = cards.length;
		var i = 0;
		while (i < cardCount) {
			var cardValue = cards.pop();
			if (cardValue == null) {
				break;
			}
			var card:StoryCard = cast cardValue;
			if (card.state == state) {
				out += "  - " + formatCard(card, runtime) + "\n";
				hasEntries = true;
			}
			cards.add(card);
			i++;
		}
		if (!hasEntries) {
			out += "  - none\n";
		}
		return out;
	}

	static function openOwnerFocus(cards:List<StoryCard>):String {
		var owners = new List<String>();
		var cardCount = cards.length;
		var i = 0;
		while (i < cardCount) {
			var cardValue = cards.pop();
			if (cardValue == null) {
				break;
			}
			var card:StoryCard = cast cardValue;
			if (card.state != STATE_DONE) {
				owners.add(card.owner + "(p" + card.points + ")");
			}
			cards.add(card);
			i++;
		}
		if (owners.length == 0) {
			return "none";
		}
		return joinStringList(owners, ", ");
	}

	public static function render(runtime:StoryboardRuntime):String {
		var cards = buildCards();
		var total = totalPoints(cards);
		var done = donePoints(cards);
		var open = openPoints(cards);
		var readiness = readinessPercent(done, total);
		var doneCards = countByState(cards, STATE_DONE);
		var doingCards = countByState(cards, STATE_DOING);
		var todoCards = countByState(cards, STATE_TODO);
		var velocity = runtime.velocityPerSprint();
		var forecast = sprintForecast(open, velocity);
		var riskThreshold = runtime.riskThreshold();
		var highRisk = openHighRisk(cards, riskThreshold);
		var releaseOpen = releaseTaggedOpen(cards);
		var velocityHint = runtime.supportsVelocityHint() ? "adaptive" : "baseline";
		var bar = progressBar(done, total, 24);

		var action = "ready to cut release";
		if (highRisk > 0) {
			action = "ship high-risk open cards first";
		} else if (open > 0) {
			action = "clear remaining open queue";
		}

		return "==============================================" + "\nRelease Command Center :: " + runtime.profileId() + "\nWindow: sprint-24"
			+ "\n==============================================" + "\nHealth" + "\n  Readiness  " + bar + " " + readiness + "% (" + done + "/" + total
			+ " points)" + "\n  Cards      total=" + cards.length + ", todo=" + todoCards + ", doing=" + doingCards + ", done=" + doneCards
			+ "\n  Open Load  " + open + " points | velocity=" + velocity + " points/sprint | eta=" + forecast + " sprint(s)" + "\n  Team Focus "
			+ openOwnerFocus(cards) + "\n  Velocity Hint: " + velocityHint + "\n\nBoard" + "\n" + formatLane(cards, STATE_TODO, "TODO", runtime)
			+ formatLane(cards, STATE_DOING, "DOING", runtime) + formatLane(cards, STATE_DONE, "DONE", runtime) + "\nRisk Radar" + "\n  High-Risk Open (>= p"
			+ riskThreshold + "): " + highRisk + "\n  Release-Tagged Open: " + releaseOpen + "\n  Profile Signal: " + runtime.extraSignal(cards)
			+ "\n\nDecision" + "\n  " + action;
	}

	public static function assertContract(runtime:StoryboardRuntime):String {
		var cards = buildCards();
		var summary = "cards=" + cards.length + ",points=" + totalPoints(cards) + ",done_points=" + donePoints(cards) + ",open_points=" + openPoints(cards)
			+ ",readiness=" + readinessPercent(donePoints(cards), totalPoints(cards));
		if (summary != "cards=5,points=21,done_points=8,open_points=13,readiness=38") {
			throw "baseline drift: " + summary;
		}
		var extra = runtime.extraSignal(cards);
		if (extra == null || extra == "") {
			throw "missing extra signal";
		}
		if (runtime.velocityPerSprint() <= 0) {
			throw "invalid velocity";
		}
		return "OK " + runtime.profileId();
	}
}
