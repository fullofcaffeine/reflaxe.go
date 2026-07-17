package main

import "examples_profile_storyboard_portable/hxrt"

var Harness_STATE_DOING *string = hxrt.StringFromLiteral("doing")

var Harness_STATE_DONE *string = hxrt.StringFromLiteral("done")

var Harness_STATE_TODO *string = hxrt.StringFromLiteral("todo")

func Harness_assertContract(runtime profile__StoryboardRuntime) *string {
	cards := Harness_buildCards()
	summary := hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("cards="), len(cards)), hxrt.StringFromLiteral(",points=")), Harness_totalPoints(cards)), hxrt.StringFromLiteral(",done_points=")), Harness_donePoints(cards)), hxrt.StringFromLiteral(",open_points=")), Harness_openPoints(cards)), hxrt.StringFromLiteral(",readiness=")), Harness_readinessPercent(Harness_donePoints(cards), Harness_totalPoints(cards)))
	if !hxrt.StringEqualStringPtr(summary, hxrt.StringFromLiteral("cards=5,points=21,done_points=8,open_points=13,readiness=38")) {
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("baseline drift: "), summary))
	}
	extra := runtime.extraSignal(Harness_buildSignalMetrics(cards))
	if hxrt.StringEqualStringPtr(extra, nil) || hxrt.StringEqualStringPtr(extra, hxrt.StringFromLiteral("")) {
		hxrt.Throw(hxrt.StringFromLiteral("missing extra signal"))
	}
	if runtime.velocityPerSprint() <= 0 {
		hxrt.Throw(hxrt.StringFromLiteral("invalid velocity"))
	}
	return hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("OK "), runtime.profileId())
}

func Harness_buildCards() []*domain__StoryCard {
	cards := []*domain__StoryCard{}
	cards = append(cards, Harness_card(1, hxrt.StringFromLiteral("Ship profile docs"), 3, Harness_makeTags(hxrt.StringFromLiteral("docs"), hxrt.StringFromLiteral("profiles")), hxrt.StringFromLiteral("done"), hxrt.StringFromLiteral("Alex")))
	cards = append(cards, Harness_card(2, hxrt.StringFromLiteral("Backfill regression snapshots"), 5, Harness_makeTags(hxrt.StringFromLiteral("tests"), nil), hxrt.StringFromLiteral("done"), hxrt.StringFromLiteral("Mira")))
	cards = append(cards, Harness_card(3, hxrt.StringFromLiteral("Wire release artifacts"), 5, Harness_makeTags(hxrt.StringFromLiteral("ci"), hxrt.StringFromLiteral("release")), hxrt.StringFromLiteral("doing"), hxrt.StringFromLiteral("Noah")))
	cards = append(cards, Harness_card(4, hxrt.StringFromLiteral("CLI polish for dev:hx"), 3, Harness_makeTags(hxrt.StringFromLiteral("devex"), nil), hxrt.StringFromLiteral("todo"), hxrt.StringFromLiteral("Jules")))
	cards = append(cards, Harness_card(5, hxrt.StringFromLiteral("Interactive tui_todo demo"), 5, Harness_makeTags(hxrt.StringFromLiteral("examples"), hxrt.StringFromLiteral("release")), hxrt.StringFromLiteral("doing"), hxrt.StringFromLiteral("Sam")))
	return cards
}

func Harness_buildSignalMetrics(cards []*domain__StoryCard) *profile__StorySignalMetrics {
	highValue := 0
	openHighValue := 0
	_g := 0
	for _g < len(cards) {
		card := cards[_g]
		_g = int(int32((_g + 1)))
		if card.points >= 5 {
			highValue = int(int32((highValue + 1)))
			if !hxrt.StringEqualStringPtr(card.state, hxrt.StringFromLiteral("done")) {
				openHighValue = int(int32((openHighValue + 1)))
			}
		}
	}
	return New_profile__StorySignalMetrics(len(cards), highValue, openHighValue)
}

func Harness_card(id int, title *string, points int, tags []*string, state *string, owner *string) *domain__StoryCard {
	return New_domain__StoryCard(id, title, points, tags, state, owner)
}

func Harness_countByState(cards []*domain__StoryCard, state *string) int {
	total := 0
	_g := 0
	for _g < len(cards) {
		card := cards[_g]
		_g = int(int32((_g + 1)))
		if hxrt.StringEqualStringPtr(card.state, state) {
			total = int(int32((total + 1)))
		}
	}
	return total
}

func Harness_donePoints(cards []*domain__StoryCard) int {
	total := 0
	_g := 0
	for _g < len(cards) {
		card := cards[_g]
		_g = int(int32((_g + 1)))
		if hxrt.StringEqualStringPtr(card.state, hxrt.StringFromLiteral("done")) {
			total = int(int32((hxrt.Int32Wrap(total) + hxrt.Int32Wrap(card.points))))
		}
	}
	return total
}

func Harness_formatCard(card *domain__StoryCard, runtime profile__StoryboardRuntime) *string {
	tags := []*string{}
	_g := 0
	_g1 := card.tags
	for _g < len(_g1) {
		tag := _g1[_g]
		_g = int(int32((_g + 1)))
		tags = append(tags, runtime.highlightTag(tag))
	}
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("#"), card.id), hxrt.StringFromLiteral(" p")), card.points), hxrt.StringFromLiteral(" ")), runtime.decorateTitle(card.title)), hxrt.StringFromLiteral(" owner:")), card.owner), hxrt.StringFromLiteral(" tags:")), Harness_joinStringList(tags, hxrt.StringFromLiteral("|")))
}

func Harness_formatLane(cards []*domain__StoryCard, state *string, title *string, runtime profile__StoryboardRuntime) *string {
	out := hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(title, hxrt.StringFromLiteral(" (")), Harness_countByState(cards, state)), hxrt.StringFromLiteral(")\n"))
	hasEntries := false
	_g := 0
	for _g < len(cards) {
		card := cards[_g]
		_g = int(int32((_g + 1)))
		if hxrt.StringEqualStringPtr(card.state, state) {
			out = hxrt.StringConcatStringPtr(out, hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("  - "), Harness_formatCard(card, runtime)), hxrt.StringFromLiteral("\n")))
			hasEntries = true
		}
	}
	if !hasEntries {
		out = hxrt.StringConcatStringPtr(out, hxrt.StringFromLiteral("  - none\n"))
	}
	return out
}

func Harness_hasTag(card *domain__StoryCard, needle *string) bool {
	_g := 0
	_g1 := card.tags
	for _g < len(_g1) {
		tag := _g1[_g]
		_g = int(int32((_g + 1)))
		if hxrt.StringEqualStringPtr(tag, needle) {
			return true
		}
	}
	return false
}

func Harness_intFloorDiv(numerator int, denominator int) int {
	if denominator <= 0 {
		return 0
	}
	quotient := 0
	remaining := numerator
	for remaining >= denominator {
		remaining = int(int32((hxrt.Int32Wrap(remaining) - hxrt.Int32Wrap(denominator))))
		quotient = int(int32((quotient + 1)))
	}
	return quotient
}

func Harness_joinStringList(values []*string, separator *string) *string {
	out := hxrt.StringFromLiteral("")
	first := true
	_g := 0
	for _g < len(values) {
		value := values[_g]
		_g = int(int32((_g + 1)))
		if !first {
			out = hxrt.StringConcatStringPtr(out, separator)
		}
		out = hxrt.StringConcatStringPtr(out, value)
		first = false
	}
	return out
}

func Harness_makeTags(a *string, b *string) []*string {
	tags := []*string{}
	tags = append(tags, a)
	if !hxrt.StringEqualStringPtr(b, nil) {
		tags = append(tags, b)
	}
	return tags
}

func Harness_openHighRisk(cards []*domain__StoryCard, threshold int) int {
	total := 0
	_g := 0
	for _g < len(cards) {
		card := cards[_g]
		_g = int(int32((_g + 1)))
		if !hxrt.StringEqualStringPtr(card.state, hxrt.StringFromLiteral("done")) && (card.points >= threshold) {
			total = int(int32((total + 1)))
		}
	}
	return total
}

func Harness_openOwnerFocus(cards []*domain__StoryCard) *string {
	owners := []*string{}
	_g := 0
	for _g < len(cards) {
		card := cards[_g]
		_g = int(int32((_g + 1)))
		if !hxrt.StringEqualStringPtr(card.state, hxrt.StringFromLiteral("done")) {
			owners = append(owners, hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(card.owner, hxrt.StringFromLiteral("(p")), card.points), hxrt.StringFromLiteral(")")))
		}
	}
	if len(owners) == 0 {
		return hxrt.StringFromLiteral("none")
	}
	return Harness_joinStringList(owners, hxrt.StringFromLiteral(", "))
}

func Harness_openPoints(cards []*domain__StoryCard) int {
	total := 0
	_g := 0
	for _g < len(cards) {
		card := cards[_g]
		_g = int(int32((_g + 1)))
		if !hxrt.StringEqualStringPtr(card.state, hxrt.StringFromLiteral("done")) {
			total = int(int32((hxrt.Int32Wrap(total) + hxrt.Int32Wrap(card.points))))
		}
	}
	return total
}

func Harness_progressBar(donePoints int, totalPoints int, width int) *string {
	if width <= 0 {
		return hxrt.StringFromLiteral("[]")
	}
	filled := 0
	if totalPoints > 0 {
		filled = Harness_intFloorDiv(int(int32((hxrt.Int32Wrap(donePoints) * hxrt.Int32Wrap(width)))), totalPoints)
	}
	if filled < 0 {
		filled = 0
	}
	if filled > width {
		filled = width
	}
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("["), Harness_repeatChar(hxrt.StringFromLiteral("#"), filled)), Harness_repeatChar(hxrt.StringFromLiteral("-"), int(int32((hxrt.Int32Wrap(width)-hxrt.Int32Wrap(filled)))))), hxrt.StringFromLiteral("]"))
}

func Harness_readinessPercent(donePoints int, totalPoints int) int {
	if totalPoints <= 0 {
		return 0
	}
	return Harness_intFloorDiv(int(int32((hxrt.Int32Wrap(donePoints) * hxrt.Int32Wrap(100)))), totalPoints)
}

func Harness_releaseTaggedOpen(cards []*domain__StoryCard) int {
	total := 0
	_g := 0
	for _g < len(cards) {
		card := cards[_g]
		_g = int(int32((_g + 1)))
		if !hxrt.StringEqualStringPtr(card.state, hxrt.StringFromLiteral("done")) && Harness_hasTag(card, hxrt.StringFromLiteral("release")) {
			total = int(int32((total + 1)))
		}
	}
	return total
}

func Harness_render(runtime profile__StoryboardRuntime) *string {
	cards := Harness_buildCards()
	total := Harness_totalPoints(cards)
	done := Harness_donePoints(cards)
	open := Harness_openPoints(cards)
	readiness := Harness_readinessPercent(done, total)
	doneCards := Harness_countByState(cards, hxrt.StringFromLiteral("done"))
	doingCards := Harness_countByState(cards, hxrt.StringFromLiteral("doing"))
	todoCards := Harness_countByState(cards, hxrt.StringFromLiteral("todo"))
	velocity := runtime.velocityPerSprint()
	forecast := Harness_sprintForecast(open, velocity)
	riskThreshold := runtime.riskThreshold()
	highRisk := Harness_openHighRisk(cards, riskThreshold)
	releaseOpen := Harness_releaseTaggedOpen(cards)
	signalMetrics := Harness_buildSignalMetrics(cards)
	var hx_if_10 *string
	if runtime.supportsVelocityHint() {
		hx_if_10 = hxrt.StringFromLiteral("adaptive")
	} else {
		hx_if_10 = hxrt.StringFromLiteral("baseline")
	}
	velocityHint := hx_if_10
	bar := Harness_progressBar(done, total, 24)
	action := hxrt.StringFromLiteral("ready to cut release")
	if highRisk > 0 {
		action = hxrt.StringFromLiteral("ship high-risk open cards first")
	} else {
		if open > 0 {
			action = hxrt.StringFromLiteral("clear remaining open queue")
		}
	}
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("==============================================\nRelease Command Center :: "), runtime.profileId()), hxrt.StringFromLiteral("\nWindow: sprint-24")), hxrt.StringFromLiteral("\n==============================================")), hxrt.StringFromLiteral("\nHealth")), hxrt.StringFromLiteral("\n  Readiness  ")), bar), hxrt.StringFromLiteral(" ")), readiness), hxrt.StringFromLiteral("% (")), done), hxrt.StringFromLiteral("/")), total), hxrt.StringFromLiteral(" points)")), hxrt.StringFromLiteral("\n  Cards      total=")), len(cards)), hxrt.StringFromLiteral(", todo=")), todoCards), hxrt.StringFromLiteral(", doing=")), doingCards), hxrt.StringFromLiteral(", done=")), doneCards), hxrt.StringFromLiteral("\n  Open Load  ")), open), hxrt.StringFromLiteral(" points | velocity=")), velocity), hxrt.StringFromLiteral(" points/sprint | eta=")), forecast), hxrt.StringFromLiteral(" sprint(s)")), hxrt.StringFromLiteral("\n  Team Focus ")), Harness_openOwnerFocus(cards)), hxrt.StringFromLiteral("\n  Velocity Hint: ")), velocityHint), hxrt.StringFromLiteral("\n\nBoard")), hxrt.StringFromLiteral("\n")), Harness_formatLane(cards, hxrt.StringFromLiteral("todo"), hxrt.StringFromLiteral("TODO"), runtime)), Harness_formatLane(cards, hxrt.StringFromLiteral("doing"), hxrt.StringFromLiteral("DOING"), runtime)), Harness_formatLane(cards, hxrt.StringFromLiteral("done"), hxrt.StringFromLiteral("DONE"), runtime)), hxrt.StringFromLiteral("\nRisk Radar")), hxrt.StringFromLiteral("\n  High-Risk Open (>= p")), riskThreshold), hxrt.StringFromLiteral("): ")), highRisk), hxrt.StringFromLiteral("\n  Release-Tagged Open: ")), releaseOpen), hxrt.StringFromLiteral("\n  Profile Signal: ")), runtime.extraSignal(signalMetrics)), hxrt.StringFromLiteral("\n\nDecision")), hxrt.StringFromLiteral("\n  ")), action)
}

func Harness_repeatChar(ch *string, count int) *string {
	if count <= 0 {
		return hxrt.StringFromLiteral("")
	}
	out := hxrt.StringFromLiteral("")
	i := 0
	for i < count {
		out = hxrt.StringConcatStringPtr(out, ch)
		i = int(int32((i + 1)))
	}
	return out
}

func Harness_sprintForecast(openPoints int, velocityPerSprint int) int {
	if openPoints <= 0 {
		return 0
	}
	return Harness_intFloorDiv(int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(openPoints) + hxrt.Int32Wrap(velocityPerSprint))))) - hxrt.Int32Wrap(1)))), velocityPerSprint)
}

func Harness_totalPoints(cards []*domain__StoryCard) int {
	totalPoints := 0
	_g := 0
	for _g < len(cards) {
		card := cards[_g]
		_g = int(int32((_g + 1)))
		totalPoints = int(int32((hxrt.Int32Wrap(totalPoints) + hxrt.Int32Wrap(card.points))))
	}
	return totalPoints
}
