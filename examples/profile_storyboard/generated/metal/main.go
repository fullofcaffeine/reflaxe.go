package main

import "examples_profile_storyboard_metal/hxrt"

func Harness_assertContract(runtime profile__StoryboardRuntime) *string {
	cards := Harness_buildCards()
	baseline := Harness_baselineSummary(cards)
	if !hxrt.StringEqualStringPtr(baseline, hxrt.StringFromLiteral("cards=3,points=13,open=3")) {
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("baseline drift: "), baseline))
	}
	extra := runtime.extraSignal(cards)
	if hxrt.StringEqualStringPtr(extra, nil) || hxrt.StringEqualStringPtr(extra, hxrt.StringFromLiteral("")) {
		hxrt.Throw(hxrt.StringFromLiteral("missing extra signal"))
	}
	return hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("OK "), runtime.profileId())
}

func Harness_baselineSummary(cards *haxe__ds__List) *string {
	totalPoints := 0
	_ = totalPoints
	open := 0
	_ = open
	count := cards.length
	_ = count
	i := 0
	for i < count {
		value := cards.pop().(*domain__StoryCard)
		if hxrt.StringEqualAny(value, nil) {
			break
		}
		card := value
		totalPoints = (totalPoints + card.points)
		open = (open + 1)
		cards.add(card)
		i = (i + 1)
	}
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("cards="), open), hxrt.StringFromLiteral(",points=")), totalPoints), hxrt.StringFromLiteral(",open=")), open)
}

func Harness_buildCards() *haxe__ds__List {
	cards := New_haxe__ds__List()
	cards.add(Harness_card(1, hxrt.StringFromLiteral("Ship profile docs"), 3, Harness_makeTags(hxrt.StringFromLiteral("docs"), hxrt.StringFromLiteral("profiles"))))
	cards.add(Harness_card(2, hxrt.StringFromLiteral("Backfill regression snapshots"), 5, Harness_makeTags(hxrt.StringFromLiteral("tests"), nil)))
	cards.add(Harness_card(3, hxrt.StringFromLiteral("Wire release artifacts"), 5, Harness_makeTags(hxrt.StringFromLiteral("ci"), hxrt.StringFromLiteral("release"))))
	return cards
}

func Harness_card(id int, title *string, points int, tags *haxe__ds__List) *domain__StoryCard {
	return New_domain__StoryCard(id, title, points, tags)
}

func Harness_formatCards(cards *haxe__ds__List, runtime profile__StoryboardRuntime) *string {
	out := hxrt.StringFromLiteral("")
	_ = out
	firstCard := true
	_ = firstCard
	cardCount := cards.length
	_ = cardCount
	i := 0
	for i < cardCount {
		cardValue := cards.pop().(*domain__StoryCard)
		if hxrt.StringEqualAny(cardValue, nil) {
			break
		}
		card := cardValue
		_ = card
		tags := New_haxe__ds__List()
		_ = tags
		tagCount := card.tags.length
		_ = tagCount
		j := 0
		for j < tagCount {
			tagValue := card.tags.pop().(*string)
			if hxrt.StringEqualStringPtr(tagValue, nil) {
				break
			}
			tag := tagValue
			tags.add(runtime.highlightTag(tag))
			card.tags.add(tag)
			j = (j + 1)
		}
		if !firstCard {
			out = hxrt.StringConcatStringPtr(out, hxrt.StringFromLiteral(";"))
		}
		out = hxrt.StringConcatStringPtr(out, hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("#"), card.id), hxrt.StringFromLiteral(":")), runtime.decorateTitle(card.title)), hxrt.StringFromLiteral(":p")), card.points), hxrt.StringFromLiteral(":")), Harness_joinStringList(tags, hxrt.StringFromLiteral("|"))))
		firstCard = false
		cards.add(card)
		i = (i + 1)
	}
	return out
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
		raw := values.pop().(*string)
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
		i = (i + 1)
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

func Harness_render(runtime profile__StoryboardRuntime) *string {
	cards := Harness_buildCards()
	_ = cards
	velocity := hxrt.StringFromLiteral("off")
	if runtime.supportsVelocityHint() {
		velocity = hxrt.StringFromLiteral("on")
	}
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("profile="), runtime.profileId()), hxrt.StringFromLiteral("\nbaseline=")), Harness_baselineSummary(cards)), hxrt.StringFromLiteral("\nview=")), Harness_formatCards(cards, runtime)), hxrt.StringFromLiteral("\nextra=")), runtime.extraSignal(cards)), hxrt.StringFromLiteral("\nvelocity_hint=")), velocity)
}

func main() {
	var runtime profile__StoryboardRuntime = profile__RuntimeFactory_create()
	hxrt.Println(Harness_render(runtime))
}

type I_domain__StoryCard interface {
}

type domain__StoryCard struct {
	__hx_this I_domain__StoryCard
	id        int
	title     *string
	points    int
	tags      *haxe__ds__List
}

func New_domain__StoryCard(id int, title *string, points int, tags *haxe__ds__List) *domain__StoryCard {
	self := &domain__StoryCard{}
	self.__hx_this = self
	self.id = id
	self.title = title
	self.points = points
	self.tags = tags
	return self
}

type I_profile__MetalRuntime interface {
	profileId() *string
	decorateTitle(title *string) *string
	highlightTag(tag *string) *string
	extraSignal(cards *haxe__ds__List) *string
	supportsVelocityHint() bool
}

type profile__MetalRuntime struct {
	__hx_this I_profile__MetalRuntime
}

func New_profile__MetalRuntime() *profile__MetalRuntime {
	self := &profile__MetalRuntime{}
	self.__hx_this = self
	return self
}

func (self *profile__MetalRuntime) profileId() *string {
	return hxrt.StringFromLiteral("metal")
}

func (self *profile__MetalRuntime) decorateTitle(title *string) *string {
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("["), title), hxrt.StringFromLiteral("]"))
}

func (self *profile__MetalRuntime) highlightTag(tag *string) *string {
	return hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("metal-"), tag)
}

func (self *profile__MetalRuntime) extraSignal(cards *haxe__ds__List) *string {
	highValue := 0
	_ = highValue
	count := cards.length
	_ = count
	i := 0
	for i < count {
		value := cards.pop().(*domain__StoryCard)
		if hxrt.StringEqualAny(value, nil) {
			break
		}
		card := value
		if card.points >= 5 {
			highValue = (highValue + 1)
		}
		cards.add(card)
		i = (i + 1)
	}
	return hxrt.StringConcatAny(hxrt.StringFromLiteral("interop_lane=typed+strict,high_value="), highValue)
}

func (self *profile__MetalRuntime) supportsVelocityHint() bool {
	return true
}

func profile__RuntimeFactory_create() profile__StoryboardRuntime {
	return New_profile__MetalRuntime()
}

type profile__StoryboardRuntime interface {
	profileId() *string
	decorateTitle(title *string) *string
	highlightTag(tag *string) *string
	extraSignal(cards *haxe__ds__List) *string
	supportsVelocityHint() bool
}

type haxe__ds__IntMap struct {
	h map[int]any
}

type haxe__ds__StringMap struct {
	h map[string]any
}

type haxe__ds__ObjectMap struct {
	h map[any]any
}

type haxe__ds__EnumValueMap struct {
	h map[any]any
}

type haxe__ds__List struct {
	items  []any
	length int
}

func New_haxe__ds__IntMap() *haxe__ds__IntMap {
	return &haxe__ds__IntMap{h: map[int]any{}}
}

func (self *haxe__ds__IntMap) set(key int, value any) {
	self.h[key] = value
}

func (self *haxe__ds__IntMap) get(key int) any {
	value := self.h[key]
	return value
}

func (self *haxe__ds__IntMap) exists(key int) bool {
	_, ok := self.h[key]
	return ok
}

func (self *haxe__ds__IntMap) remove(key int) bool {
	_, ok := self.h[key]
	delete(self.h, key)
	return ok
}

func New_haxe__ds__StringMap() *haxe__ds__StringMap {
	return &haxe__ds__StringMap{h: map[string]any{}}
}

func (self *haxe__ds__StringMap) set(key *string, value any) {
	self.h[*hxrt.StdString(key)] = value
}

func (self *haxe__ds__StringMap) get(key *string) any {
	value := self.h[*hxrt.StdString(key)]
	return value
}

func (self *haxe__ds__StringMap) exists(key *string) bool {
	_, ok := self.h[*hxrt.StdString(key)]
	return ok
}

func (self *haxe__ds__StringMap) remove(key *string) bool {
	_, ok := self.h[*hxrt.StdString(key)]
	delete(self.h, *hxrt.StdString(key))
	return ok
}

func New_haxe__ds__ObjectMap() *haxe__ds__ObjectMap {
	return &haxe__ds__ObjectMap{h: map[any]any{}}
}

func (self *haxe__ds__ObjectMap) set(key any, value any) {
	self.h[key] = value
}

func (self *haxe__ds__ObjectMap) get(key any) any {
	return self.h[key]
}

func (self *haxe__ds__ObjectMap) exists(key any) bool {
	_, ok := self.h[key]
	return ok
}

func (self *haxe__ds__ObjectMap) remove(key any) bool {
	_, ok := self.h[key]
	delete(self.h, key)
	return ok
}

func New_haxe__ds__EnumValueMap() *haxe__ds__EnumValueMap {
	return &haxe__ds__EnumValueMap{h: map[any]any{}}
}

func (self *haxe__ds__EnumValueMap) set(key any, value any) {
	self.h[key] = value
}

func (self *haxe__ds__EnumValueMap) get(key any) any {
	return self.h[key]
}

func (self *haxe__ds__EnumValueMap) exists(key any) bool {
	_, ok := self.h[key]
	return ok
}

func (self *haxe__ds__EnumValueMap) remove(key any) bool {
	_, ok := self.h[key]
	delete(self.h, key)
	return ok
}

func New_haxe__ds__List() *haxe__ds__List {
	return &haxe__ds__List{items: []any{}, length: 0}
}

func (self *haxe__ds__List) add(item any) {
	self.items = append(self.items, item)
	self.length = len(self.items)
}

func (self *haxe__ds__List) push(item any) {
	self.add(item)
}

func (self *haxe__ds__List) pop() any {
	if len(self.items) == 0 {
		return nil
	}
	head := self.items[0]
	self.items = self.items[1:]
	self.length = len(self.items)
	return head
}

func (self *haxe__ds__List) first() any {
	if len(self.items) == 0 {
		return nil
	}
	return self.items[0]
}

func (self *haxe__ds__List) last() any {
	size := len(self.items)
	if size == 0 {
		return nil
	}
	return self.items[(size - 1)]
}
