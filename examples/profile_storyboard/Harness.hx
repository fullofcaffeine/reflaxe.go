import domain.StoryCard;
import profile.StorySignalMetrics;
import profile.StoryboardRuntime;

class Harness {
	static inline final STATE_TODO = "todo";
	static inline final STATE_DOING = "doing";
	static inline final STATE_DONE = "done";

	static function makeTags(a:String, ?b:String):Array<String> {
		var tags = new Array<String>();
		tags.push(a);
		if (b != null) {
			tags.push(b);
		}
		return tags;
	}

	static function joinStringList(values:Array<String>, separator:String):String {
		var out = "";
		var first = true;
		for (value in values) {
			if (!first) {
				out += separator;
			}
			out += value;
			first = false;
		}
		return out;
	}

	static function card(id:Int, title:String, points:Int, tags:Array<String>, state:String, owner:String):StoryCard {
		return new StoryCard(id, title, points, tags, state, owner);
	}

	public static function buildCards():Array<StoryCard> {
		var cards = new Array<StoryCard>();
		cards.push(card(1, "Ship profile docs", 3, makeTags("docs", "profiles"), STATE_DONE, "Alex"));
		cards.push(card(2, "Backfill regression snapshots", 5, makeTags("tests"), STATE_DONE, "Mira"));
		cards.push(card(3, "Wire release artifacts", 5, makeTags("ci", "release"), STATE_DOING, "Noah"));
		cards.push(card(4, "CLI polish for dev:hx", 3, makeTags("devex"), STATE_TODO, "Jules"));
		cards.push(card(5, "Interactive tui_todo demo", 5, makeTags("examples", "release"), STATE_DOING, "Sam"));
		return cards;
	}

	static function totalPoints(cards:Array<StoryCard>):Int {
		var totalPoints = 0;
		for (card in cards) {
			totalPoints += card.points;
		}
		return totalPoints;
	}

	static function countByState(cards:Array<StoryCard>, state:String):Int {
		var total = 0;
		for (card in cards) {
			if (card.state == state) {
				total++;
			}
		}
		return total;
	}

	static function donePoints(cards:Array<StoryCard>):Int {
		var total = 0;
		for (card in cards) {
			if (card.state == STATE_DONE) {
				total += card.points;
			}
		}
		return total;
	}

	static function openPoints(cards:Array<StoryCard>):Int {
		var total = 0;
		for (card in cards) {
			if (card.state != STATE_DONE) {
				total += card.points;
			}
		}
		return total;
	}

	static function hasTag(card:StoryCard, needle:String):Bool {
		for (tag in card.tags) {
			if (tag == needle) {
				return true;
			}
		}
		return false;
	}

	static function openHighRisk(cards:Array<StoryCard>, threshold:Int):Int {
		var total = 0;
		for (card in cards) {
			if (card.state != STATE_DONE && card.points >= threshold) {
				total++;
			}
		}
		return total;
	}

	static function releaseTaggedOpen(cards:Array<StoryCard>):Int {
		var total = 0;
		for (card in cards) {
			if (card.state != STATE_DONE && hasTag(card, "release")) {
				total++;
			}
		}
		return total;
	}

	static function buildSignalMetrics(cards:Array<StoryCard>):StorySignalMetrics {
		var highValue = 0;
		var openHighValue = 0;
		for (card in cards) {
			if (card.points >= 5) {
				highValue++;
				if (card.state != STATE_DONE) {
					openHighValue++;
				}
			}
		}
		return new StorySignalMetrics(cards.length, highValue, openHighValue);
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
		var tags = new Array<String>();
		for (tag in card.tags) {
			tags.push(runtime.highlightTag(tag));
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

	static function formatLane(cards:Array<StoryCard>, state:String, title:String, runtime:StoryboardRuntime):String {
		var out = title + " (" + countByState(cards, state) + ")\n";
		var hasEntries = false;
		for (card in cards) {
			if (card.state == state) {
				out += "  - " + formatCard(card, runtime) + "\n";
				hasEntries = true;
			}
		}
		if (!hasEntries) {
			out += "  - none\n";
		}
		return out;
	}

	static function openOwnerFocus(cards:Array<StoryCard>):String {
		var owners = new Array<String>();
		for (card in cards) {
			if (card.state != STATE_DONE) {
				owners.push(card.owner + "(p" + card.points + ")");
			}
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
		var signalMetrics = buildSignalMetrics(cards);
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
			+ riskThreshold + "): " + highRisk + "\n  Release-Tagged Open: " + releaseOpen + "\n  Profile Signal: " + runtime.extraSignal(signalMetrics)
			+ "\n\nDecision" + "\n  " + action;
	}

	public static function assertContract(runtime:StoryboardRuntime):String {
		var cards = buildCards();
		var summary = "cards=" + cards.length + ",points=" + totalPoints(cards) + ",done_points=" + donePoints(cards) + ",open_points=" + openPoints(cards)
			+ ",readiness=" + readinessPercent(donePoints(cards), totalPoints(cards));
		if (summary != "cards=5,points=21,done_points=8,open_points=13,readiness=38") {
			throw "baseline drift: " + summary;
		}
		var extra = runtime.extraSignal(buildSignalMetrics(cards));
		if (extra == null || extra == "") {
			throw "missing extra signal";
		}
		if (runtime.velocityPerSprint() <= 0) {
			throw "invalid velocity";
		}
		return "OK " + runtime.profileId();
	}
}
