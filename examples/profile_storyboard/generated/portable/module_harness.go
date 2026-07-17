package main

import "examples_profile_storyboard_portable/hxrt"

var Harness_STATE_DOING *string = hxrt.StringFromLiteral("doing")

var Harness_STATE_DONE *string = hxrt.StringFromLiteral("done")

var Harness_STATE_TODO *string = hxrt.StringFromLiteral("todo")

func Harness_assertContract(runtime profile__StoryboardRuntime) *string {
	cards := Harness_buildCards()
	summary := hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("cards="), cards.Len()), hxrt.StringFromLiteral(",points=")), Harness_totalPoints(cards)), hxrt.StringFromLiteral(",done_points=")), Harness_donePoints(cards)), hxrt.StringFromLiteral(",open_points=")), Harness_openPoints(cards)), hxrt.StringFromLiteral(",readiness=")), Harness_readinessPercent(Harness_donePoints(cards), Harness_totalPoints(cards)))
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

func Harness_buildCards() *hxrt.Array {
	cards := hxrt.NewArray()
	cards.Push(Harness_card(1, hxrt.StringFromLiteral("Ship profile docs"), 3, Harness_makeTags(hxrt.StringFromLiteral("docs"), hxrt.StringFromLiteral("profiles")), hxrt.StringFromLiteral("done"), hxrt.StringFromLiteral("Alex")))
	cards.Push(Harness_card(2, hxrt.StringFromLiteral("Backfill regression snapshots"), 5, Harness_makeTags(hxrt.StringFromLiteral("tests"), nil), hxrt.StringFromLiteral("done"), hxrt.StringFromLiteral("Mira")))
	cards.Push(Harness_card(3, hxrt.StringFromLiteral("Wire release artifacts"), 5, Harness_makeTags(hxrt.StringFromLiteral("ci"), hxrt.StringFromLiteral("release")), hxrt.StringFromLiteral("doing"), hxrt.StringFromLiteral("Noah")))
	cards.Push(Harness_card(4, hxrt.StringFromLiteral("CLI polish for dev:hx"), 3, Harness_makeTags(hxrt.StringFromLiteral("devex"), nil), hxrt.StringFromLiteral("todo"), hxrt.StringFromLiteral("Jules")))
	cards.Push(Harness_card(5, hxrt.StringFromLiteral("Interactive tui_todo demo"), 5, Harness_makeTags(hxrt.StringFromLiteral("examples"), hxrt.StringFromLiteral("release")), hxrt.StringFromLiteral("doing"), hxrt.StringFromLiteral("Sam")))
	return cards
}

func Harness_buildSignalMetrics(cards *hxrt.Array) *profile__StorySignalMetrics {
	highValue := 0
	openHighValue := 0
	_g := 0
	for _g < cards.Len() {
		card := func(hx_value_6 any) *domain__StoryCard {
			if hx_value_6 == nil {
				var hx_zero_7 *domain__StoryCard
				return hx_zero_7
			}
			return hx_value_6.(*domain__StoryCard)
		}(cards.Get(_g))
		_g = int(int32((_g + 1)))
		if card.points >= 5 {
			highValue = int(int32((highValue + 1)))
			if !hxrt.StringEqualStringPtr(card.state, hxrt.StringFromLiteral("done")) {
				openHighValue = int(int32((openHighValue + 1)))
			}
		}
	}
	return New_profile__StorySignalMetrics(cards.Len(), highValue, openHighValue)
}

func Harness_card(id int, title *string, points int, tags *hxrt.Array, state *string, owner *string) *domain__StoryCard {
	return New_domain__StoryCard(id, title, points, tags, state, owner)
}

func Harness_countByState(cards *hxrt.Array, state *string) int {
	total := 0
	_g := 0
	for _g < cards.Len() {
		card := func(hx_value_8 any) *domain__StoryCard {
			if hx_value_8 == nil {
				var hx_zero_9 *domain__StoryCard
				return hx_zero_9
			}
			return hx_value_8.(*domain__StoryCard)
		}(cards.Get(_g))
		_g = int(int32((_g + 1)))
		if hxrt.StringEqualStringPtr(card.state, state) {
			total = int(int32((total + 1)))
		}
	}
	return total
}

func Harness_donePoints(cards *hxrt.Array) int {
	total := 0
	_g := 0
	for _g < cards.Len() {
		card := func(hx_value_10 any) *domain__StoryCard {
			if hx_value_10 == nil {
				var hx_zero_11 *domain__StoryCard
				return hx_zero_11
			}
			return hx_value_10.(*domain__StoryCard)
		}(cards.Get(_g))
		_g = int(int32((_g + 1)))
		if hxrt.StringEqualStringPtr(card.state, hxrt.StringFromLiteral("done")) {
			total = int(int32((hxrt.Int32Wrap(total) + hxrt.Int32Wrap(card.points))))
		}
	}
	return total
}

func Harness_formatCard(card *domain__StoryCard, runtime profile__StoryboardRuntime) *string {
	tags := hxrt.NewArray()
	_g := 0
	_g1 := card.tags
	for _g < _g1.Len() {
		tag := func(hx_value_12 any) *string {
			if hx_value_12 == nil {
				var hx_zero_13 *string
				return hx_zero_13
			}
			return hx_value_12.(*string)
		}(_g1.Get(_g))
		_g = int(int32((_g + 1)))
		tags.Push(runtime.highlightTag(tag))
	}
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("#"), card.id), hxrt.StringFromLiteral(" p")), card.points), hxrt.StringFromLiteral(" ")), runtime.decorateTitle(card.title)), hxrt.StringFromLiteral(" owner:")), card.owner), hxrt.StringFromLiteral(" tags:")), Harness_joinStringList(tags, hxrt.StringFromLiteral("|")))
}

func Harness_formatLane(cards *hxrt.Array, state *string, title *string, runtime profile__StoryboardRuntime) *string {
	out := hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(title, hxrt.StringFromLiteral(" (")), Harness_countByState(cards, state)), hxrt.StringFromLiteral(")\n"))
	hasEntries := false
	_g := 0
	for _g < cards.Len() {
		card := func(hx_value_15 any) *domain__StoryCard {
			if hx_value_15 == nil {
				var hx_zero_16 *domain__StoryCard
				return hx_zero_16
			}
			return hx_value_15.(*domain__StoryCard)
		}(cards.Get(_g))
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
	for _g < _g1.Len() {
		tag := func(hx_value_17 any) *string {
			if hx_value_17 == nil {
				var hx_zero_18 *string
				return hx_zero_18
			}
			return hx_value_17.(*string)
		}(_g1.Get(_g))
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

func Harness_joinStringList(values *hxrt.Array, separator *string) *string {
	out := hxrt.StringFromLiteral("")
	first := true
	_g := 0
	for _g < values.Len() {
		value := func(hx_value_19 any) *string {
			if hx_value_19 == nil {
				var hx_zero_20 *string
				return hx_zero_20
			}
			return hx_value_19.(*string)
		}(values.Get(_g))
		_g = int(int32((_g + 1)))
		if !first {
			out = hxrt.StringConcatStringPtr(out, separator)
		}
		out = hxrt.StringConcatStringPtr(out, value)
		first = false
	}
	return out
}

func Harness_makeTags(a *string, b *string) *hxrt.Array {
	tags := hxrt.NewArray()
	tags.Push(a)
	if !hxrt.StringEqualStringPtr(b, nil) {
		tags.Push(b)
	}
	return tags
}

func Harness_openHighRisk(cards *hxrt.Array, threshold int) int {
	total := 0
	_g := 0
	for _g < cards.Len() {
		card := func(hx_value_23 any) *domain__StoryCard {
			if hx_value_23 == nil {
				var hx_zero_24 *domain__StoryCard
				return hx_zero_24
			}
			return hx_value_23.(*domain__StoryCard)
		}(cards.Get(_g))
		_g = int(int32((_g + 1)))
		if !hxrt.StringEqualStringPtr(card.state, hxrt.StringFromLiteral("done")) && (card.points >= threshold) {
			total = int(int32((total + 1)))
		}
	}
	return total
}

func Harness_openOwnerFocus(cards *hxrt.Array) *string {
	owners := hxrt.NewArray()
	_g := 0
	for _g < cards.Len() {
		card := func(hx_value_25 any) *domain__StoryCard {
			if hx_value_25 == nil {
				var hx_zero_26 *domain__StoryCard
				return hx_zero_26
			}
			return hx_value_25.(*domain__StoryCard)
		}(cards.Get(_g))
		_g = int(int32((_g + 1)))
		if !hxrt.StringEqualStringPtr(card.state, hxrt.StringFromLiteral("done")) {
			owners.Push(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(card.owner, hxrt.StringFromLiteral("(p")), card.points), hxrt.StringFromLiteral(")")))
		}
	}
	if owners.Len() == 0 {
		return hxrt.StringFromLiteral("none")
	}
	return Harness_joinStringList(owners, hxrt.StringFromLiteral(", "))
}

func Harness_openPoints(cards *hxrt.Array) int {
	total := 0
	_g := 0
	for _g < cards.Len() {
		card := func(hx_value_28 any) *domain__StoryCard {
			if hx_value_28 == nil {
				var hx_zero_29 *domain__StoryCard
				return hx_zero_29
			}
			return hx_value_28.(*domain__StoryCard)
		}(cards.Get(_g))
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

func Harness_releaseTaggedOpen(cards *hxrt.Array) int {
	total := 0
	_g := 0
	for _g < cards.Len() {
		card := func(hx_value_30 any) *domain__StoryCard {
			if hx_value_30 == nil {
				var hx_zero_31 *domain__StoryCard
				return hx_zero_31
			}
			return hx_value_30.(*domain__StoryCard)
		}(cards.Get(_g))
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
	var hx_if_32 *string
	if runtime.supportsVelocityHint() {
		hx_if_32 = hxrt.StringFromLiteral("adaptive")
	} else {
		hx_if_32 = hxrt.StringFromLiteral("baseline")
	}
	velocityHint := hx_if_32
	bar := Harness_progressBar(done, total, 24)
	action := hxrt.StringFromLiteral("ready to cut release")
	if highRisk > 0 {
		action = hxrt.StringFromLiteral("ship high-risk open cards first")
	} else {
		if open > 0 {
			action = hxrt.StringFromLiteral("clear remaining open queue")
		}
	}
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("==============================================\nRelease Command Center :: "), runtime.profileId()), hxrt.StringFromLiteral("\nWindow: sprint-24")), hxrt.StringFromLiteral("\n==============================================")), hxrt.StringFromLiteral("\nHealth")), hxrt.StringFromLiteral("\n  Readiness  ")), bar), hxrt.StringFromLiteral(" ")), readiness), hxrt.StringFromLiteral("% (")), done), hxrt.StringFromLiteral("/")), total), hxrt.StringFromLiteral(" points)")), hxrt.StringFromLiteral("\n  Cards      total=")), cards.Len()), hxrt.StringFromLiteral(", todo=")), todoCards), hxrt.StringFromLiteral(", doing=")), doingCards), hxrt.StringFromLiteral(", done=")), doneCards), hxrt.StringFromLiteral("\n  Open Load  ")), open), hxrt.StringFromLiteral(" points | velocity=")), velocity), hxrt.StringFromLiteral(" points/sprint | eta=")), forecast), hxrt.StringFromLiteral(" sprint(s)")), hxrt.StringFromLiteral("\n  Team Focus ")), Harness_openOwnerFocus(cards)), hxrt.StringFromLiteral("\n  Velocity Hint: ")), velocityHint), hxrt.StringFromLiteral("\n\nBoard")), hxrt.StringFromLiteral("\n")), Harness_formatLane(cards, hxrt.StringFromLiteral("todo"), hxrt.StringFromLiteral("TODO"), runtime)), Harness_formatLane(cards, hxrt.StringFromLiteral("doing"), hxrt.StringFromLiteral("DOING"), runtime)), Harness_formatLane(cards, hxrt.StringFromLiteral("done"), hxrt.StringFromLiteral("DONE"), runtime)), hxrt.StringFromLiteral("\nRisk Radar")), hxrt.StringFromLiteral("\n  High-Risk Open (>= p")), riskThreshold), hxrt.StringFromLiteral("): ")), highRisk), hxrt.StringFromLiteral("\n  Release-Tagged Open: ")), releaseOpen), hxrt.StringFromLiteral("\n  Profile Signal: ")), runtime.extraSignal(signalMetrics)), hxrt.StringFromLiteral("\n\nDecision")), hxrt.StringFromLiteral("\n  ")), action)
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

func Harness_totalPoints(cards *hxrt.Array) int {
	totalPoints := 0
	_g := 0
	for _g < cards.Len() {
		card := func(hx_value_33 any) *domain__StoryCard {
			if hx_value_33 == nil {
				var hx_zero_34 *domain__StoryCard
				return hx_zero_34
			}
			return hx_value_33.(*domain__StoryCard)
		}(cards.Get(_g))
		_g = int(int32((_g + 1)))
		totalPoints = int(int32((hxrt.Int32Wrap(totalPoints) + hxrt.Int32Wrap(card.points))))
	}
	return totalPoints
}
