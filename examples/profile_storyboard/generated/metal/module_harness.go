package main

import "examples_profile_storyboard_metal/hxrt"

var Harness_STATE_DOING *string = hxrt.StringFromLiteral("doing")

var Harness_STATE_DONE *string = hxrt.StringFromLiteral("done")

var Harness_STATE_TODO *string = hxrt.StringFromLiteral("todo")

func Harness_assertContract(runtime profile__StoryboardRuntime) *string {
	cards := Harness_buildCards()
	summary := hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("cards="), cards.length), hxrt.StringFromLiteral(",points=")), Harness_totalPoints(cards)), hxrt.StringFromLiteral(",done_points=")), Harness_donePoints(cards)), hxrt.StringFromLiteral(",open_points=")), Harness_openPoints(cards)), hxrt.StringFromLiteral(",readiness=")), Harness_readinessPercent(Harness_donePoints(cards), Harness_totalPoints(cards)))
	if !hxrt.StringEqualStringPtr(summary, hxrt.StringFromLiteral("cards=5,points=21,done_points=8,open_points=13,readiness=38")) {
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("baseline drift: "), summary))
		var hx_throw_zero_1 *string
		return hx_throw_zero_1
	}
	extra := runtime.extraSignal(cards)
	if hxrt.StringEqualStringPtr(extra, nil) || hxrt.StringEqualStringPtr(extra, hxrt.StringFromLiteral("")) {
		hxrt.Throw(hxrt.StringFromLiteral("missing extra signal"))
		var hx_throw_zero_2 *string
		return hx_throw_zero_2
	}
	if runtime.velocityPerSprint() <= 0 {
		hxrt.Throw(hxrt.StringFromLiteral("invalid velocity"))
		var hx_throw_zero_3 *string
		return hx_throw_zero_3
	}
	return hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("OK "), runtime.profileId())
}

func Harness_buildCards() *haxe__ds__List {
	cards := New_haxe__ds__List()
	cards.add(Harness_card(1, hxrt.StringFromLiteral("Ship profile docs"), 3, Harness_makeTags(hxrt.StringFromLiteral("docs"), hxrt.StringFromLiteral("profiles")), hxrt.StringFromLiteral("done"), hxrt.StringFromLiteral("Alex")))
	cards.add(Harness_card(2, hxrt.StringFromLiteral("Backfill regression snapshots"), 5, Harness_makeTags(hxrt.StringFromLiteral("tests"), nil), hxrt.StringFromLiteral("done"), hxrt.StringFromLiteral("Mira")))
	cards.add(Harness_card(3, hxrt.StringFromLiteral("Wire release artifacts"), 5, Harness_makeTags(hxrt.StringFromLiteral("ci"), hxrt.StringFromLiteral("release")), hxrt.StringFromLiteral("doing"), hxrt.StringFromLiteral("Noah")))
	cards.add(Harness_card(4, hxrt.StringFromLiteral("CLI polish for dev:hx"), 3, Harness_makeTags(hxrt.StringFromLiteral("devex"), nil), hxrt.StringFromLiteral("todo"), hxrt.StringFromLiteral("Jules")))
	cards.add(Harness_card(5, hxrt.StringFromLiteral("Interactive tui_todo demo"), 5, Harness_makeTags(hxrt.StringFromLiteral("examples"), hxrt.StringFromLiteral("release")), hxrt.StringFromLiteral("doing"), hxrt.StringFromLiteral("Sam")))
	return cards
}

func Harness_card(id int, title *string, points int, tags *haxe__ds__List, state *string, owner *string) *domain__StoryCard {
	return New_domain__StoryCard(id, title, points, tags, state, owner)
}

func Harness_countByState(cards *haxe__ds__List, state *string) int {
	total := 0
	_ = total
	count := cards.length
	_ = count
	i := 0
	for i < count {
		value := func(hx_value_4 any) *domain__StoryCard {
			if hx_value_4 == nil {
				var hx_zero_5 *domain__StoryCard
				return hx_zero_5
			}
			return hx_value_4.(*domain__StoryCard)
		}(cards.pop())
		if value == nil {
			break
		}
		card := value
		if hxrt.StringEqualStringPtr(card.state, state) {
			total = int(int32((total + 1)))
		}
		cards.add(card)
		i = int(int32((i + 1)))
	}
	return total
}

func Harness_donePoints(cards *haxe__ds__List) int {
	total := 0
	_ = total
	count := cards.length
	_ = count
	i := 0
	for i < count {
		value := func(hx_value_6 any) *domain__StoryCard {
			if hx_value_6 == nil {
				var hx_zero_7 *domain__StoryCard
				return hx_zero_7
			}
			return hx_value_6.(*domain__StoryCard)
		}(cards.pop())
		if value == nil {
			break
		}
		card := value
		if hxrt.StringEqualStringPtr(card.state, hxrt.StringFromLiteral("done")) {
			total = int(int32((hxrt.Int32Wrap(total) + hxrt.Int32Wrap(card.points))))
		}
		cards.add(card)
		i = int(int32((i + 1)))
	}
	return total
}

func Harness_formatCard(card *domain__StoryCard, runtime profile__StoryboardRuntime) *string {
	tags := New_haxe__ds__List()
	_ = tags
	tagCount := card.tags.length
	_ = tagCount
	j := 0
	for j < tagCount {
		tagValue := func(hx_value_8 any) *string {
			if hx_value_8 == nil {
				var hx_zero_9 *string
				return hx_zero_9
			}
			return hx_value_8.(*string)
		}(card.tags.pop())
		if hxrt.StringEqualStringPtr(tagValue, nil) {
			break
		}
		tag := tagValue
		tags.add(runtime.highlightTag(tag))
		card.tags.add(tag)
		j = int(int32((j + 1)))
	}
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("#"), card.id), hxrt.StringFromLiteral(" p")), card.points), hxrt.StringFromLiteral(" ")), runtime.decorateTitle(card.title)), hxrt.StringFromLiteral(" owner:")), card.owner), hxrt.StringFromLiteral(" tags:")), Harness_joinStringList(tags, hxrt.StringFromLiteral("|")))
}

func Harness_formatLane(cards *haxe__ds__List, state *string, title *string, runtime profile__StoryboardRuntime) *string {
	out := hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(title, hxrt.StringFromLiteral(" (")), Harness_countByState(cards, state)), hxrt.StringFromLiteral(")\n"))
	_ = out
	hasEntries := false
	_ = hasEntries
	cardCount := cards.length
	_ = cardCount
	i := 0
	for i < cardCount {
		cardValue := func(hx_value_10 any) *domain__StoryCard {
			if hx_value_10 == nil {
				var hx_zero_11 *domain__StoryCard
				return hx_zero_11
			}
			return hx_value_10.(*domain__StoryCard)
		}(cards.pop())
		if cardValue == nil {
			break
		}
		card := cardValue
		if hxrt.StringEqualStringPtr(card.state, state) {
			out = hxrt.StringConcatStringPtr(out, hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("  - "), Harness_formatCard(card, runtime)), hxrt.StringFromLiteral("\n")))
			hasEntries = true
		}
		cards.add(card)
		i = int(int32((i + 1)))
	}
	if !hasEntries {
		out = hxrt.StringConcatStringPtr(out, hxrt.StringFromLiteral("  - none\n"))
	}
	return out
}

func Harness_hasTag(card *domain__StoryCard, needle *string) bool {
	found := false
	_ = found
	count := card.tags.length
	_ = count
	i := 0
	for i < count {
		value := func(hx_value_12 any) *string {
			if hx_value_12 == nil {
				var hx_zero_13 *string
				return hx_zero_13
			}
			return hx_value_12.(*string)
		}(card.tags.pop())
		if hxrt.StringEqualStringPtr(value, nil) {
			break
		}
		tag := value
		if hxrt.StringEqualStringPtr(tag, needle) {
			found = true
		}
		card.tags.add(tag)
		i = int(int32((i + 1)))
	}
	return found
}

func Harness_intFloorDiv(numerator int, denominator int) int {
	if denominator <= 0 {
		return 0
	}
	quotient := 0
	_ = quotient
	remaining := numerator
	for remaining >= denominator {
		remaining = int(int32((hxrt.Int32Wrap(remaining) - hxrt.Int32Wrap(denominator))))
		quotient = int(int32((quotient + 1)))
	}
	return quotient
}

func Harness_joinStringList(values *haxe__ds__List, separator *string) *string {
	out := hxrt.StringFromLiteral("")
	_ = out
	first := true
	_ = first
	count := values.length
	_ = count
	i := 0
	for i < count {
		raw := func(hx_value_14 any) *string {
			if hx_value_14 == nil {
				var hx_zero_15 *string
				return hx_zero_15
			}
			return hx_value_14.(*string)
		}(values.pop())
		if hxrt.StringEqualStringPtr(raw, nil) {
			break
		}
		value := raw
		_ = value
		if !first {
			out = hxrt.StringConcatStringPtr(out, separator)
		}
		out = hxrt.StringConcatStringPtr(out, value)
		values.add(value)
		first = false
		i = int(int32((i + 1)))
	}
	return out
}

func Harness_makeTags(a *string, b *string) *haxe__ds__List {
	tags := New_haxe__ds__List()
	tags.add(a)
	if !hxrt.StringEqualStringPtr(b, nil) {
		tags.add(b)
	}
	return tags
}

func Harness_openHighRisk(cards *haxe__ds__List, threshold int) int {
	total := 0
	_ = total
	count := cards.length
	_ = count
	i := 0
	for i < count {
		value := func(hx_value_16 any) *domain__StoryCard {
			if hx_value_16 == nil {
				var hx_zero_17 *domain__StoryCard
				return hx_zero_17
			}
			return hx_value_16.(*domain__StoryCard)
		}(cards.pop())
		if value == nil {
			break
		}
		card := value
		if !hxrt.StringEqualStringPtr(card.state, hxrt.StringFromLiteral("done")) && (card.points >= threshold) {
			total = int(int32((total + 1)))
		}
		cards.add(card)
		i = int(int32((i + 1)))
	}
	return total
}

func Harness_openOwnerFocus(cards *haxe__ds__List) *string {
	owners := New_haxe__ds__List()
	_ = owners
	cardCount := cards.length
	_ = cardCount
	i := 0
	for i < cardCount {
		cardValue := func(hx_value_18 any) *domain__StoryCard {
			if hx_value_18 == nil {
				var hx_zero_19 *domain__StoryCard
				return hx_zero_19
			}
			return hx_value_18.(*domain__StoryCard)
		}(cards.pop())
		if cardValue == nil {
			break
		}
		card := cardValue
		if !hxrt.StringEqualStringPtr(card.state, hxrt.StringFromLiteral("done")) {
			owners.add(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(card.owner, hxrt.StringFromLiteral("(p")), card.points), hxrt.StringFromLiteral(")")))
		}
		cards.add(card)
		i = int(int32((i + 1)))
	}
	if owners.length == 0 {
		return hxrt.StringFromLiteral("none")
	}
	return Harness_joinStringList(owners, hxrt.StringFromLiteral(", "))
}

func Harness_openPoints(cards *haxe__ds__List) int {
	total := 0
	_ = total
	count := cards.length
	_ = count
	i := 0
	for i < count {
		value := func(hx_value_20 any) *domain__StoryCard {
			if hx_value_20 == nil {
				var hx_zero_21 *domain__StoryCard
				return hx_zero_21
			}
			return hx_value_20.(*domain__StoryCard)
		}(cards.pop())
		if value == nil {
			break
		}
		card := value
		if !hxrt.StringEqualStringPtr(card.state, hxrt.StringFromLiteral("done")) {
			total = int(int32((hxrt.Int32Wrap(total) + hxrt.Int32Wrap(card.points))))
		}
		cards.add(card)
		i = int(int32((i + 1)))
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

func Harness_releaseTaggedOpen(cards *haxe__ds__List) int {
	total := 0
	_ = total
	count := cards.length
	_ = count
	i := 0
	for i < count {
		value := func(hx_value_22 any) *domain__StoryCard {
			if hx_value_22 == nil {
				var hx_zero_23 *domain__StoryCard
				return hx_zero_23
			}
			return hx_value_22.(*domain__StoryCard)
		}(cards.pop())
		if value == nil {
			break
		}
		card := value
		if !hxrt.StringEqualStringPtr(card.state, hxrt.StringFromLiteral("done")) && Harness_hasTag(card, hxrt.StringFromLiteral("release")) {
			total = int(int32((total + 1)))
		}
		cards.add(card)
		i = int(int32((i + 1)))
	}
	return total
}

func Harness_render(runtime profile__StoryboardRuntime) *string {
	cards := Harness_buildCards()
	total := Harness_totalPoints(cards)
	_ = total
	done := Harness_donePoints(cards)
	_ = done
	open := Harness_openPoints(cards)
	_ = open
	readiness := Harness_readinessPercent(done, total)
	_ = readiness
	doneCards := Harness_countByState(cards, hxrt.StringFromLiteral("done"))
	_ = doneCards
	doingCards := Harness_countByState(cards, hxrt.StringFromLiteral("doing"))
	_ = doingCards
	todoCards := Harness_countByState(cards, hxrt.StringFromLiteral("todo"))
	_ = todoCards
	velocity := runtime.velocityPerSprint()
	forecast := Harness_sprintForecast(open, velocity)
	_ = forecast
	riskThreshold := runtime.riskThreshold()
	highRisk := Harness_openHighRisk(cards, riskThreshold)
	_ = highRisk
	releaseOpen := Harness_releaseTaggedOpen(cards)
	_ = releaseOpen
	var hx_if_24 *string
	if runtime.supportsVelocityHint() {
		hx_if_24 = hxrt.StringFromLiteral("adaptive")
	} else {
		hx_if_24 = hxrt.StringFromLiteral("baseline")
	}
	velocityHint := hx_if_24
	_ = velocityHint
	bar := Harness_progressBar(done, total, 24)
	_ = bar
	action := hxrt.StringFromLiteral("ready to cut release")
	if highRisk > 0 {
		action = hxrt.StringFromLiteral("ship high-risk open cards first")
	} else {
		if open > 0 {
			action = hxrt.StringFromLiteral("clear remaining open queue")
		}
	}
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("==============================================\nRelease Command Center :: "), runtime.profileId()), hxrt.StringFromLiteral("\nWindow: sprint-24")), hxrt.StringFromLiteral("\n==============================================")), hxrt.StringFromLiteral("\nHealth")), hxrt.StringFromLiteral("\n  Readiness  ")), bar), hxrt.StringFromLiteral(" ")), readiness), hxrt.StringFromLiteral("% (")), done), hxrt.StringFromLiteral("/")), total), hxrt.StringFromLiteral(" points)")), hxrt.StringFromLiteral("\n  Cards      total=")), cards.length), hxrt.StringFromLiteral(", todo=")), todoCards), hxrt.StringFromLiteral(", doing=")), doingCards), hxrt.StringFromLiteral(", done=")), doneCards), hxrt.StringFromLiteral("\n  Open Load  ")), open), hxrt.StringFromLiteral(" points | velocity=")), velocity), hxrt.StringFromLiteral(" points/sprint | eta=")), forecast), hxrt.StringFromLiteral(" sprint(s)")), hxrt.StringFromLiteral("\n  Team Focus ")), Harness_openOwnerFocus(cards)), hxrt.StringFromLiteral("\n  Velocity Hint: ")), velocityHint), hxrt.StringFromLiteral("\n\nBoard")), hxrt.StringFromLiteral("\n")), Harness_formatLane(cards, hxrt.StringFromLiteral("todo"), hxrt.StringFromLiteral("TODO"), runtime)), Harness_formatLane(cards, hxrt.StringFromLiteral("doing"), hxrt.StringFromLiteral("DOING"), runtime)), Harness_formatLane(cards, hxrt.StringFromLiteral("done"), hxrt.StringFromLiteral("DONE"), runtime)), hxrt.StringFromLiteral("\nRisk Radar")), hxrt.StringFromLiteral("\n  High-Risk Open (>= p")), riskThreshold), hxrt.StringFromLiteral("): ")), highRisk), hxrt.StringFromLiteral("\n  Release-Tagged Open: ")), releaseOpen), hxrt.StringFromLiteral("\n  Profile Signal: ")), runtime.extraSignal(cards)), hxrt.StringFromLiteral("\n\nDecision")), hxrt.StringFromLiteral("\n  ")), action)
}

func Harness_repeatChar(ch *string, count int) *string {
	if count <= 0 {
		return hxrt.StringFromLiteral("")
	}
	out := hxrt.StringFromLiteral("")
	_ = out
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

func Harness_totalPoints(cards *haxe__ds__List) int {
	totalPoints := 0
	_ = totalPoints
	count := cards.length
	_ = count
	i := 0
	for i < count {
		value := func(hx_value_25 any) *domain__StoryCard {
			if hx_value_25 == nil {
				var hx_zero_26 *domain__StoryCard
				return hx_zero_26
			}
			return hx_value_25.(*domain__StoryCard)
		}(cards.pop())
		if value == nil {
			break
		}
		card := value
		totalPoints = int(int32((hxrt.Int32Wrap(totalPoints) + hxrt.Int32Wrap(card.points))))
		cards.add(card)
		i = int(int32((i + 1)))
	}
	return totalPoints
}
